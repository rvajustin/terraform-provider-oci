// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package resourcediscovery

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/terraform-providers/terraform-provider-oci/internal/service/bds"

	oci_core "github.com/oracle/oci-go-sdk/v54/core"

	tf_client "github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/globalvar"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-exec/tfexec"

	"github.com/hashicorp/hcl2/hclwrite"
)

var isInitDone bool
var initLock sync.Mutex

type TerraformResourceHints struct {
	// Information about this resource
	resourceClass        string // The name of the resource class (e.g. oci_core_vcn)
	resourceAbbreviation string // An abbreviated version of the resource class used for generating shorter resource names (e.g. vcn)

	// Hints to help with discovering this resource using data sources
	datasourceClass              string                  // The name of the data source class to use for discovering resources (e.g. oci_core_vcns)
	datasourceItemsAttr          string                  // The attribute with the data source that contains the discovered resources returned by the data source (e.g. virtual_networks)
	isDatasourceCollection       bool                    // True if list datasource is modeled as a collection with `items` field under datasourceItemsAttr
	requireResourceRefresh       bool                    // Whether to use the resource to fill in missing information from datasource (e.g. when datasources only return summary information)
	discoverableLifecycleStates  []string                // List of lifecycle states that should be discovered. If empty, then all lifecycle states are discoverable.
	processDiscoveredResourcesFn ProcessOCIResourcesFunc // Custom function for processing resources discovered by the data source
	alwaysExportable             bool                    // Some resources always need to be exportable, regardless of whether they are being targeted for export
	isDataSource                 bool
	getIdFn                      func(*OCIResource) (string, error) // If the resource has no OCID generated by services, then implement this to generate one from the OCIResource. Typically used for composite IDs.

	// Override function for discovering resources. To be used when there is no datasource implementation to help with discovery.
	findResourcesOverrideFn func(*resourceDiscoveryContext, *TerraformResourceAssociation, *OCIResource, *TerraformResourceGraph) ([]*OCIResource, error)

	// Hints to help with generating HCL representation from this resource
	getHCLStringOverrideFn func(*strings.Builder, *OCIResource, map[string]string) error // Custom function for generating HCL syntax for the resource

	// Hints for adding default value to HCL representation for attributes not found in resource discovery
	defaultValuesForMissingAttributes map[string]interface{}

	// Hints for adding resource attributes to `ignore_changes` in HCL representation
	// This is added to avoid plan failure/diff for attributes that service does not return in read response
	// The attributes references are interpolated in case of nested attributes
	ignorableRequiredMissingAttributes map[string]bool
}

func (h *TerraformResourceHints) DiscoversWithSingularDatasource() bool {
	return h.datasourceItemsAttr == ""
}

type TerraformResourceAssociation struct {
	*TerraformResourceHints
	datasourceQueryParams map[string]string // Mapping of data source inputs and the source attribute from a parent resource
}

// Wrapper around string value to differentiate strings from interpolations
// Differentiation needed to write oci_resource.resource_name vs "oci_resource.resource_name" for v0.12
type InterpolationString struct {
	resourceReference string
	interpolation     string
	value             string
}

type ResourceDiscoveryError struct {
	resourceType   string
	parentResource string
	error          error
	resourceGraph  *TerraformResourceGraph
}

type ErrorList struct {
	errors []*ResourceDiscoveryError
}

/*  ctxLock is the common lock for the whole struct
WARN: Make sure NOT to pass resourceDiscoveryContext as value,
as that would copy the struct and locks should not be copied
*/
type resourceDiscoveryContext struct {
	ctxLock                     sync.Mutex // common lock for the whole context, make sure to acquire the lock before modifying any field in the resourceDiscoveryContext
	terraformProviderBinaryPath string
	terraformCLIPath            string
	terraform                   *tfexec.Terraform
	clients                     *tf_client.OracleClients
	expectedResourceIds         map[string]bool
	tenancyOcid                 string
	discoveredResources         []*OCIResource
	summaryStatements           []string
	targetSpecificResources     bool
	resourceHintsLookup         map[string]*TerraformResourceHints
	*ExportCommandArgs
	errorList                    ErrorList
	missingAttributesPerResource map[string][]string
	isImportError                bool // flag indicates if there was an import failure and if reference map needs to be updated
	state                        interface{}
	timeTakenToDiscover          time.Duration
	timeTakenToGenerateState     time.Duration
	timeTakenForEntireExport     time.Duration
}

// Resource discovery Exit status
type Status int

const (
	// Exit statuses
	StatusSuccess Status = iota
	StatusFail
	StatusPartialSuccess Status = 64

	// Tags to filter resources
	OracleTagsCreatedBy           = "Oracle-Tags.CreatedBy"
	OkeTagValue                   = "oke"
	ResourceCreatedByInstancePool = "oci:compute:instancepool"

	// parallelism configs
	ChunkSize         int = 10
	MaxParallelChunks int = 4
)

var isAllDataSourceLock sync.Mutex

func (ctx *resourceDiscoveryContext) addErrorToList(error *ResourceDiscoveryError) {
	ctx.ctxLock.Lock()
	defer ctx.ctxLock.Unlock()
	ctx.errorList.errors = append(ctx.errorList.errors, error)

}

func (ctx *resourceDiscoveryContext) postValidate() {
	// Check that all expected resource IDs were found, if any were given
	var missingResourceIds []string
	for resourceId, found := range ctx.expectedResourceIds {
		if !found {
			missingResourceIds = append(missingResourceIds, resourceId)
		}
	}

	if len(missingResourceIds) > 0 {
		ctx.summaryStatements = append(ctx.summaryStatements, "")
		ctx.summaryStatements = append(ctx.summaryStatements, "Warning: The following resource IDs were not found.")
		for _, resourceId := range missingResourceIds {
			ctx.summaryStatements = append(ctx.summaryStatements, fmt.Sprintf("- %s", resourceId))
		}
		rdError := &ResourceDiscoveryError{
			"",
			"",
			fmt.Errorf("[ERROR] one or more expected resource ids were not found"),
			nil}

		ctx.addErrorToList(rdError)
	}
}

func (ctx *resourceDiscoveryContext) printSummary() {

	ctx.summaryStatements = append(ctx.summaryStatements, "=== COMPLETED ===")

	for _, statement := range ctx.summaryStatements {
		utils.Logln(utils.Green(statement))
	}
	utils.Logln(utils.Green("========= PERFORMANCE SUMMARY =========="))
	utils.Logln(utils.Green(fmt.Sprintf("Total resources: %v", len(ctx.discoveredResources))))
	utils.Logln(utils.Green(fmt.Sprintf("Total time taken for discovering all services: %v", ctx.timeTakenToDiscover)))
	utils.Logln(utils.Green(fmt.Sprintf("Total time taken for generating state of all services: %v", ctx.timeTakenToGenerateState)))
	utils.Logln(utils.Green(fmt.Sprintf("Total time taken by entire export: %v", ctx.timeTakenForEntireExport)))
}

func (ctx *resourceDiscoveryContext) printErrors() {
	utils.Logln(utils.Yellow("\n\n[WARN] Resource discovery finished with errors listed below:\n"))
	for _, resourceDiscoveryError := range ctx.errorList.errors {
		if resourceDiscoveryError.resourceType == "" || ctx.targetSpecificResources {
			utils.Logln(utils.Yellow(resourceDiscoveryError.error.Error()))

		} else if resourceDiscoveryError.parentResource == "export" {
			utils.Logln(utils.Yellow(fmt.Sprintf("Error discovering `%s` resources: %s", resourceDiscoveryError.resourceType, resourceDiscoveryError.error.Error())))

		} else {
			utils.Logln(utils.Yellow(fmt.Sprintf("Error discovering `%s` resources for %s: %s", resourceDiscoveryError.resourceType, resourceDiscoveryError.parentResource, resourceDiscoveryError.error.Error())))
		}
		/* log child resources if exist and were not discovered because of error in parent resource discovery*/
		if resourceDiscoveryError.resourceGraph != nil && !ctx.targetSpecificResources {
			var notFoundChildren []string
			getNotFoundChildren(resourceDiscoveryError.resourceType, resourceDiscoveryError.resourceGraph, &notFoundChildren)
			if len(notFoundChildren) > 0 {
				utils.Logln(utils.Yellow(fmt.Sprintf("\tFollowing child resources were also not discovered due to parent error: %v", strings.Join(notFoundChildren, ", "))))
			}
		}
	}
}

func getNotFoundChildren(parent string, resourceGraph *TerraformResourceGraph, children *[]string) {
	childResources, exists := (*resourceGraph)[parent]
	if exists {
		for _, child := range childResources {
			*children = append(*children, child.resourceClass)
			// Avoid recursion if a resource can be nested within itself e.g. compartments
			if child.resourceClass != parent {
				getNotFoundChildren(child.resourceClass, resourceGraph, children)
			}
		}
	}
}

func createResourceDiscoveryContext(clients *tf_client.OracleClients, args *ExportCommandArgs, tenancyOcid string) (*resourceDiscoveryContext, error) {

	result := &resourceDiscoveryContext{
		clients:             clients,
		ExportCommandArgs:   args,
		tenancyOcid:         tenancyOcid,
		discoveredResources: []*OCIResource{},
		summaryStatements:   []string{},
		errorList: ErrorList{
			errors: []*ResourceDiscoveryError{},
		},
		targetSpecificResources: false,
		resourceHintsLookup:     createResourceHintsLookupMap(),
	}
	// Use user provided terraform-provider-oci executable
	if pluginDir := utils.GetEnvSettingWithBlankDefault("provider_bin_path"); pluginDir != "" {
		result.terraformProviderBinaryPath = pluginDir
		utils.Logf("[INFO] terraform provider binary path (pluginDir) set using `provider_bin_path`: '%s'", result.terraformProviderBinaryPath)
	}

	if *result.CompartmentId == "" {
		*result.CompartmentId = tenancyOcid
		vars["tenancy_ocid"] = fmt.Sprintf("\"%s\"", tenancyOcid)
		referenceMap[tenancyOcid] = tfHclVersion.getVarHclString("tenancy_ocid")
	} else {
		vars["compartment_ocid"] = fmt.Sprintf("\"%s\"", *result.CompartmentId)
		referenceMap[*result.CompartmentId] = tfHclVersion.getVarHclString("compartment_ocid")
	}

	result.expectedResourceIds = convertStringSliceToSet(args.IDs, true)

	re := regexp.MustCompile(`oci_([^:]+):(.+$)`)

	for id := range result.expectedResourceIds {
		subMatchAll := re.FindStringSubmatch(id)
		if subMatchAll != nil && len(subMatchAll) == 3 {
			result.targetSpecificResources = true
			break
		}
	}
	// validate terraform version and initialize terraform for import - only required if generating state file
	if args.GenerateState {
		if tf, terraformCLIPath, err := createTerraformStruct(args); err != nil {
			return result, err
		} else {
			result.terraform = tf
			result.terraformCLIPath = terraformCLIPath
		}
	}
	return result, nil
}

type resourceDiscoveryStep interface {
	discover() error
	getOmittedResources() []*OCIResource
	writeTmpConfigurationForImport() error
	writeConfiguration() error
	writeTmpState() error
	getBaseStep() *resourceDiscoveryBaseStep
	mergeTempStateFiles(tmpStateOutputDir string) error
	mergeGeneratedStateFile() error
	getDiscoveredResources() []*OCIResource
	updateTimeTakenForDiscovery(timeTaken time.Duration)
	updateTimeTakenForGeneratingState(timeTaken time.Duration)
}

type resourceDiscoveryBaseStep struct {
	ctx                         *resourceDiscoveryContext
	name                        string
	discoveredResources         []*OCIResource
	omittedResources            []*OCIResource
	tempState                   interface{}
	timeTakenForDiscovery       time.Duration
	timeTakenForGeneratingState time.Duration
}

func (r *resourceDiscoveryBaseStep) mergeTempStateFiles(tmpStateOutputDir string) error {
	defer elapsed(fmt.Sprintf("merging temp state files for %v", r.name), nil, 0)()
	files, err := ioutil.ReadDir(tmpStateOutputDir)
	if err != nil {
		return err
	}
	// loop over tmp state files for each chunk and merge all to form temp State for the service
	for _, file := range files {
		var tempState interface{}
		tmpFilePath := filepath.Join(tmpStateOutputDir, file.Name())
		if !strings.HasSuffix(file.Name(), ".backup") { // ignore the backup file created by terraform
			if jsonState, err := ioutil.ReadFile(tmpFilePath); err != nil {
				return err
			} else {
				if err := json.Unmarshal(jsonState, &tempState); err != nil {
					return err
				}
			}
			if r.tempState == nil {
				r.tempState = tempState
			} else {
				r.tempState, _ = mergeState(r.tempState, tempState)
			}
		}
	}
	return nil
}
func (r *resourceDiscoveryBaseStep) writeTmpState() error {
	defer elapsed(fmt.Sprintf("writing temp state for %d '%s' resources", len(r.getDiscoveredResources()), r.name), nil, 0)()
	// Run terraform init if not already done
	if !isInitDone {
		utils.Debugf("[DEBUG] acquiring lock to run terraform init")
		initLock.Lock()
		// Check for existence of .terraform folder to make sure init is not run already by another thread
		if _, err := os.Stat(fmt.Sprintf("%s%s.terraform", *r.ctx.OutputDir, string(os.PathSeparator))); os.IsNotExist(err) {
			// Run init command if not already run
			utils.Debugf("[DEBUG] writeTmpState: running init")
			backgroundCtx := context.Background()

			var initArgs []tfexec.InitOption

			if r.ctx.terraformProviderBinaryPath != "" {
				utils.Logf("[INFO] plugin dir set to: '%s'", r.ctx.terraformProviderBinaryPath)
				initArgs = append(initArgs, tfexec.PluginDir(r.ctx.terraformProviderBinaryPath))
			}

			if err := r.ctx.terraform.Init(backgroundCtx, initArgs...); err != nil {
				return err
			}
			isInitDone = true
		}
		initLock.Unlock()
		utils.Debugf("[DEBUG] releasing lock")
	}
	tmpStateOutputDir := filepath.Join(*r.ctx.OutputDir, "tmp", r.name)
	tmpStateOutputFilePrefix := filepath.Join(tmpStateOutputDir, globalvar.DefaultTmpStateFile)

	if err := os.RemoveAll(tmpStateOutputDir); err != nil {
		utils.Logf("[WARN] unable to delete existing tmp state directory %s", tmpStateOutputDir)
		return err
	}

	isAllDataSources := true
	totalResources := len(r.discoveredResources)
	// divide list of discovered resources which is a slice into chunks
	// process each chunk in parallel
	chunkSize := ChunkSize // chunk size defines number of resources in each chunk.
	// if there are additional chunks required for left over resources.
	// For example, if chunk size is 5 and we have 8 resources then 8/5 gives int output as 1.
	// So chunk 1 will occupy 5 resources and 8 % 5 = 3. For remaining 3 resources we need additional chunk.
	additionalChunks := 1 // no. of additional chunks required to process (totalResources % chunkSize) resources.
	if totalResources%chunkSize == 0 {
		additionalChunks = 0
	}
	totalChunks := totalResources/chunkSize + additionalChunks
	var importWg sync.WaitGroup
	importWg.Add(totalChunks) // we need to wait for all chunks to finish importing resources
	// we Create buffered channel to control max parallel chunks that can be executed in parallel
	semImport := make(chan struct{}, MaxParallelChunks)
	// loop over chunks
	for chunkIdx := 0; chunkIdx < totalResources; chunkIdx += chunkSize {
		// position of last element of chunk
		lastPos := chunkIdx + chunkSize
		semImport <- struct{}{}
		// in case last chunk isn't full set the lastPos accordingly
		if lastPos > totalResources {
			lastPos = totalResources
		}
		go func(resources []*OCIResource, chunkIndex int) {
			for _, res := range resources {
				fileName := tmpStateOutputFilePrefix + fmt.Sprint(chunkIndex)
				importResource(r.ctx, res, fileName)
				if res.terraformTypeInfo != nil && !res.terraformTypeInfo.isDataSource {
					isAllDataSourceLock.Lock()
					isAllDataSources = false
					isAllDataSourceLock.Unlock()
				}
			}
			<-semImport
			importWg.Done()
			// take resources beginning at chunkIdx upto and excluding lastPos
		}(r.discoveredResources[chunkIdx:lastPos], chunkIdx)
	}
	// wait for all chunks to finish importing resources
	importWg.Wait()
	// The found resource only include the data sources (ADs and namespaces) that resource discovery adds
	if isAllDataSources {
		return nil
	}
	err := r.mergeTempStateFiles(tmpStateOutputDir)
	if err != nil {
		return err
	}
	return nil
}

// writeTmpConfigurationForImport writes temporary configuration to run terraform import on the discovered resources
// It only writes the resource block and skips the resource fields
// The configuration will be discarded and written again after import is completed for all resources
func (r *resourceDiscoveryBaseStep) writeTmpConfigurationForImport() error {
	defer elapsed(fmt.Sprintf("writing temp configuration for %d %s resources", len(r.getDiscoveredResources()), r.name), nil, 0)()
	configOutputFile := fmt.Sprintf("%s%s%s.tf", *r.ctx.OutputDir, string(os.PathSeparator), r.name)
	tmpConfigOutputFile := fmt.Sprintf("%s%s%s.tf.tmp", *r.ctx.OutputDir, string(os.PathSeparator), r.name)

	file, err := os.OpenFile(tmpConfigOutputFile, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return err
	}

	// Build the HCL config
	builder := &strings.Builder{}
	builder.WriteString("## This is tmp config to run import for resources\n\n")
	for _, resource := range r.discoveredResources {
		if resource.terraformTypeInfo != nil && resource.terraformTypeInfo.isDataSource {
			builder.WriteString(fmt.Sprintf("data %s %s {}\n\n", resource.terraformClass, resource.terraformName))
		} else {
			builder.WriteString(fmt.Sprintf("resource %s %s {}\n\n", resource.terraformClass, resource.terraformName))
		}

		r.ctx.ctxLock.Lock()
		r.ctx.discoveredResources = append(r.ctx.discoveredResources, resource)
		r.ctx.ctxLock.Unlock()
	}

	_, err = file.WriteString(string(builder.String()))
	if err != nil {
		_ = file.Close()
		return err
	}

	if fErr := file.Close(); fErr != nil {
		return fErr
	}

	if err := os.Rename(tmpConfigOutputFile, configOutputFile); err != nil {
		return err
	}
	return nil
}

func (r *resourceDiscoveryBaseStep) writeConfiguration() error {
	defer elapsed(fmt.Sprintf("writing actual configuration for %d %s resources", len(r.getDiscoveredResources()), r.name), nil, 0)()
	configOutputFile := fmt.Sprintf("%s%s%s.tf", *r.ctx.OutputDir, string(os.PathSeparator), r.name)
	tmpConfigOutputFile := fmt.Sprintf("%s%s%s.tf.tmp", *r.ctx.OutputDir, string(os.PathSeparator), r.name)

	file, err := os.OpenFile(tmpConfigOutputFile, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return err
	}

	// Build the HCL config
	// Note that we still build a TF file even if no resources were discovered for this TF file.
	// A user may run this command multiple times and may see stale resources if we don't overwrite the file with
	// an empty one.
	builder := &strings.Builder{}
	builder.WriteString("## This configuration was generated by terraform-provider-oci\n\n")

	exportedResourceCount := 0
	for _, resource := range r.discoveredResources {

		// Skip writing the config for resources for which import command failed
		if !resource.isErrorResource {
			utils.Logf("[INFO] ===> Generating resource '%s'", resource.getTerraformReference())
			if err := resource.getHCLString(builder, referenceMap); err != nil {
				_ = file.Close()
				return err
			}

			if resource.terraformTypeInfo != nil && len(resource.terraformTypeInfo.ignorableRequiredMissingAttributes) > 0 {
				attributes := make([]string, 0, len(resource.terraformTypeInfo.ignorableRequiredMissingAttributes))
				for attribute := range resource.terraformTypeInfo.ignorableRequiredMissingAttributes {
					attributes = append(attributes, attribute)
				}
				missingAttributesPerResourceLock.Lock()
				if r.ctx.missingAttributesPerResource == nil {
					r.ctx.missingAttributesPerResource = make(map[string][]string)
				}
				r.ctx.missingAttributesPerResource[resource.getTerraformReference()] = attributes
				missingAttributesPerResourceLock.Unlock()
			}

			r.ctx.discoveredResources = append(r.ctx.discoveredResources, resource)
			exportedResourceCount++
		} else {
			// remove missing attributes info if present for a failed resource
			missingAttributesPerResourceLock.Lock()
			if _, ok := r.ctx.missingAttributesPerResource[resource.getTerraformReference()]; ok {
				delete(r.ctx.missingAttributesPerResource, resource.getTerraformReference())
			}
			missingAttributesPerResourceLock.Unlock()
		}
	}

	// Format the HCL config
	formattedString := hclwrite.Format([]byte(builder.String()))

	_, err = file.WriteString(string(formattedString))
	if err != nil {
		_ = file.Close()
		return err
	}

	if fErr := file.Close(); fErr != nil {
		return fErr
	}

	if err := os.Rename(tmpConfigOutputFile, configOutputFile); err != nil {
		return err
	}

	if r.ctx.targetSpecificResources {
		r.ctx.summaryStatements = append(r.ctx.summaryStatements, fmt.Sprintf("Found %d resources. Generated under '%s'", exportedResourceCount, configOutputFile))
	} else {
		r.ctx.summaryStatements = append(r.ctx.summaryStatements, fmt.Sprintf("Found %d '%s' resources. Generated under '%s'", exportedResourceCount, r.name, configOutputFile))
	}
	r.ctx.summaryStatements = append(r.ctx.summaryStatements, fmt.Sprintf("Time taken for discovery: %v, generating state: %v", r.timeTakenForDiscovery, r.timeTakenForGeneratingState))
	return nil
}

func (r *resourceDiscoveryBaseStep) getOmittedResources() []*OCIResource {
	return r.omittedResources
}

func (r *resourceDiscoveryBaseStep) getDiscoveredResources() []*OCIResource {
	return r.discoveredResources
}

func (r *resourceDiscoveryBaseStep) updateTimeTakenForDiscovery(timeTaken time.Duration) {
	r.timeTakenForDiscovery = timeTaken
}
func (r *resourceDiscoveryBaseStep) updateTimeTakenForGeneratingState(timeTaken time.Duration) {
	r.timeTakenForGeneratingState = timeTaken
}

func (r *resourceDiscoveryBaseStep) getBaseStep() *resourceDiscoveryBaseStep {
	return r
}

type resourceDiscoveryWithGraph struct {
	resourceDiscoveryBaseStep
	root          *OCIResource
	resourceGraph TerraformResourceGraph
}

func (r *resourceDiscoveryWithGraph) discover() error {
	var err error
	var ociResources []*OCIResource

	ociResources, err = findResources(r.ctx, r.root, r.resourceGraph)
	if err != nil {
		return err
	}

	// Filter out omitted resources from export
	r.discoveredResources = []*OCIResource{}
	r.omittedResources = []*OCIResource{}
	for _, resource := range ociResources {
		if !resource.omitFromExport {

			refMapLock.Lock()
			referenceMap[resource.id] = resource.getHclReferenceIdString()
			refMapLock.Unlock()

			r.discoveredResources = append(r.discoveredResources, resource)
		} else {
			r.omittedResources = append(r.omittedResources, resource)
		}
	}
	utils.Logf("[INFO] Discovery complete for step root %s", r.name)
	return nil
}

type resourceDiscoveryWithTargetIds struct {
	resourceDiscoveryBaseStep
	exportIds map[string]string // map of IDs and their respective resource types
}

func createResourceHintsLookupMap() map[string]*TerraformResourceHints {
	result := map[string]*TerraformResourceHints{}

	for _, graphCollection := range []map[string]TerraformResourceGraph{compartmentResourceGraphs, tenancyResourceGraphs} {
		for _, graph := range graphCollection {
			for _, associations := range graph {
				for _, assoc := range associations {
					result[assoc.resourceClass] = assoc.TerraformResourceHints
				}
			}
		}
	}
	return result
}

func (ctx *resourceDiscoveryContext) getResourceHint(resourceClass string) (*TerraformResourceHints, error) {
	if hints, exists := ctx.resourceHintsLookup[resourceClass]; exists {
		return hints, nil
	}

	// If no resource hint could be found, just return a simple hint for now to unblock
	return nil, fmt.Errorf("[ERROR] resource type '%s' is not supported by resource discovery", resourceClass)
}

func (r *resourceDiscoveryWithTargetIds) discover() error {
	sortedIds := make([]string, len(r.ctx.expectedResourceIds))
	idx := 0
	for id, _ := range r.ctx.expectedResourceIds {
		sortedIds[idx] = id
		idx++
	}
	sort.Strings(sortedIds)

	re := regexp.MustCompile(`(oci_[^:]+):(.+$)`)

	for _, id := range sortedIds {
		subMatchAll := re.FindStringSubmatch(id)
		if len(subMatchAll) != 3 {
			utils.Logf("[WARN] Encountered invalid ID tuple '%s'", id)
			continue
		}

		resourceClass := subMatchAll[1]
		resourceId, _ := url.PathUnescape(subMatchAll[2])

		utils.Logf("===> Finding resource with ID '%s' and type '%s'", resourceId, resourceClass)
		resourceSchema, exists := resourcesMap[resourceClass]
		if !exists || resourceSchema.Read == nil {
			utils.Logf("[WARN] No valid resource schema could be found. Skipping.")
			continue
		}

		d := resourceSchema.Data(nil)
		d.SetId(resourceId)
		if err := resourceSchema.Read(d, r.ctx.clients); err != nil {
			utils.Logf("[WARN] Unable to read resource due to error: %v", err)
			continue
		}

		if d.Id() == "" {
			utils.Logf("[WARN] Resource ID was voided because resource could not be found. Skipping.")
			continue
		}

		resourceHint, err := r.ctx.getResourceHint(resourceClass)
		if err != nil {
			continue
		}
		ociResource, err := getOciResource(d, resourceSchema.Schema, *r.ctx.CompartmentId, resourceHint, resourceId)
		if err != nil {
			return err
		}

		if resourceHint.processDiscoveredResourcesFn != nil {
			processResults, err := resourceHint.processDiscoveredResourcesFn(r.ctx, []*OCIResource{ociResource})
			if err != nil {
				return err
			}

			if len(processResults) != 1 {
				utils.Logf("[WARN] processing of single resource resulted in %v resources being returned", len(processResults))
				continue
			}
			ociResource = processResults[0]
		}

		if ociResource.terraformName, err = generateTerraformNameFromResource(ociResource.sourceAttributes, resourceSchema.Schema); err != nil {
			terraformName := fmt.Sprintf("export_%s", resourceHint.resourceAbbreviation)
			if count, resourceNameExists := resourceNameCount[terraformName]; resourceNameExists {
				resourceNameCount[terraformName] = count + 1
				terraformName = fmt.Sprintf("%s_%d", terraformName, count)
			} else {
				resourceNameCount[terraformName] = 1
			}
			ociResource.terraformName = terraformName
		}

		r.discoveredResources = append(r.discoveredResources, ociResource)

		r.ctx.expectedResourceIds[id] = true
		// expectedResourceIds contains tuples in case of export using ids and for related resources the ids will not be a tuple
		//delete(r.ctx.expectedResourceIds, id)
		//r.ctx.expectedResourceIds[ociResource.id] = true

		if _, hasRelatedResources := exportRelatedResourcesGraph[resourceHint.resourceClass]; hasRelatedResources && r.ctx.IsExportWithRelatedResources {
			utils.Logf("[INFO] resource discovery: finding related resources for %s\n", resourceHint.resourceClass)
			ociResources, err := findResources(r.ctx, ociResource, exportRelatedResourcesGraph)
			if err != nil {
				return err
			}
			/*
				 1. Current closure graph generates only related resources but we may need to filter resources in future as the graph grows
					Because hints use datasources and if data source does not take parent param then it may generate unrelated resources
				 2. With current implementation, resource.omitFromExport will be true for child resources but we do not filter resources. If we add filtering to handle #1,
				 	then logic to set resource.omitFromExport will also need Update to handle related resources
			*/
			r.discoveredResources = append(r.discoveredResources, ociResources...)
		}
		// Add resource reference to referenceMap for discovered resources
		// If there are more than 1 resources found, this will help generate the possible references if the resources are linked
		for _, resource := range r.discoveredResources {
			referenceMap[resource.id] = resource.getHclReferenceIdString()
		}
	}
	return nil
}

type TerraformResourceGraph map[string][]TerraformResourceAssociation

type ProcessOCIResourcesFunc func(*resourceDiscoveryContext, []*OCIResource) ([]*OCIResource, error)

func init() {
	// TODO: The following changes to resource hints are deviations from what can currently be handled by the core resource discovery/generation logic
	// We should strive to eliminate these deviations by either improving the core logic or code generator

	//exportObjectStorageNamespaceHints.isDataSource = true
	exportIdentityAvailabilityDomainHints.isDataSource = true

	// Custom overrides for generating composite Load Balancer IDs within the resource discovery framework
	exportLoadBalancerBackendHints.processDiscoveredResourcesFn = processLoadBalancerBackends
	exportLoadBalancerBackendSetHints.processDiscoveredResourcesFn = processLoadBalancerBackendSets
	exportLoadBalancerCertificateHints.processDiscoveredResourcesFn = processLoadBalancerCertificates
	exportLoadBalancerHostnameHints.processDiscoveredResourcesFn = processLoadBalancerHostnames
	exportLoadBalancerListenerHints.findResourcesOverrideFn = findLoadBalancerListeners
	exportLoadBalancerListenerHints.processDiscoveredResourcesFn = processLoadBalancerListeners
	exportLoadBalancerPathRouteSetHints.processDiscoveredResourcesFn = processLoadBalancerPathRouteSets
	exportLoadBalancerRuleSetHints.processDiscoveredResourcesFn = processLoadBalancerRuleSets
	exportLoadBalancerLoadBalancerRoutingPolicyHints.processDiscoveredResourcesFn = processLoadBalancerRoutingPolicies

	exportCoreBootVolumeHints.processDiscoveredResourcesFn = filterSourcedBootVolumes
	exportCoreCrossConnectGroupHints.discoverableLifecycleStates = append(exportCoreCrossConnectGroupHints.discoverableLifecycleStates, string(oci_core.CrossConnectGroupLifecycleStateInactive))
	exportCoreDhcpOptionsHints.processDiscoveredResourcesFn = processDefaultDhcpOptions
	exportCoreImageHints.processDiscoveredResourcesFn = filterCustomImages

	exportCoreInstanceHints.discoverableLifecycleStates = append(exportCoreInstanceHints.discoverableLifecycleStates, string(oci_core.InstanceLifecycleStateStopped))
	exportCoreInstanceHints.processDiscoveredResourcesFn = processInstances
	exportCorePublicIpHints.processDiscoveredResourcesFn = processCorePublicIp
	exportCorePrivateIpHints.processDiscoveredResourcesFn = processPrivateIps
	exportCoreInstanceHints.requireResourceRefresh = true
	exportCoreNetworkSecurityGroupSecurityRuleHints.datasourceClass = "oci_core_network_security_group_security_rules"
	exportCoreNetworkSecurityGroupSecurityRuleHints.datasourceItemsAttr = "security_rules"
	exportCoreNetworkSecurityGroupSecurityRuleHints.processDiscoveredResourcesFn = processNetworkSecurityGroupRules
	exportCoreRouteTableHints.processDiscoveredResourcesFn = processDefaultRouteTables
	exportCoreSecurityListHints.processDiscoveredResourcesFn = processDefaultSecurityLists
	exportCoreVcnHints.processDiscoveredResourcesFn = processCoreVcns
	exportCoreVnicAttachmentHints.requireResourceRefresh = true
	exportCoreVnicAttachmentHints.processDiscoveredResourcesFn = filterSecondaryVnicAttachments
	exportCoreVolumeGroupHints.processDiscoveredResourcesFn = processVolumeGroups

	exportDatabaseAutonomousContainerDatabaseHints.requireResourceRefresh = true
	exportDatabaseAutonomousDatabaseHints.requireResourceRefresh = true
	exportDatabaseAutonomousDatabaseHints.processDiscoveredResourcesFn = processAutonomousDatabaseSource
	exportDatabaseAutonomousExadataInfrastructureHints.requireResourceRefresh = true
	exportDatabaseDbSystemHints.requireResourceRefresh = true
	exportDatabaseDbSystemHints.processDiscoveredResourcesFn = processDbSystems
	exportDatabaseDbHomeHints.processDiscoveredResourcesFn = filterPrimaryDbHomes
	exportDatabaseDbHomeHints.requireResourceRefresh = true
	exportDatabaseDatabaseHints.requireResourceRefresh = true
	exportDatabaseDatabaseHints.processDiscoveredResourcesFn = filterPrimaryDatabases
	exportDatabaseDatabaseHints.defaultValuesForMissingAttributes = map[string]interface{}{
		"source": "NONE",
	}
	exportDatabaseDatabaseHints.processDiscoveredResourcesFn = processDatabases
	exportDatabaseExadataInfrastructureHints.processDiscoveredResourcesFn = processDatabaseExadataInfrastructures
	//exportDatascienceModelHints.defaultValuesForMissingAttributes = map[string]interface{}{
	//	"artifact_content_length": "0",
	//}
	exportIdentityAvailabilityDomainHints.resourceAbbreviation = "ad"
	exportIdentityAvailabilityDomainHints.alwaysExportable = true
	exportIdentityAvailabilityDomainHints.processDiscoveredResourcesFn = processAvailabilityDomains
	exportIdentityAvailabilityDomainHints.getHCLStringOverrideFn = getAvailabilityDomainHCLDatasource
	exportIdentityAuthenticationPolicyHints.processDiscoveredResourcesFn = processIdentityAuthenticationPolicies
	exportIdentityTagHints.findResourcesOverrideFn = findIdentityTags
	exportIdentityTagHints.processDiscoveredResourcesFn = processTagDefinitions

	exportLoggingLogHints.getIdFn = getLogId

	exportObjectStorageNamespaceHints.processDiscoveredResourcesFn = processObjectStorageNamespace
	exportObjectStorageNamespaceHints.getHCLStringOverrideFn = getObjectStorageNamespaceHCLDatasource
	exportObjectStorageNamespaceHints.alwaysExportable = true
	exportObjectStorageObjectHints.requireResourceRefresh = true
	exportObjectStoragePreauthenticatedRequestHints.processDiscoveredResourcesFn = processObjectStoragePreauthenticatedRequest
	exportObjectStorageReplicationPolicyHints.processDiscoveredResourcesFn = processObjectStorageReplicationPolicy

	exportStreamingStreamHints.processDiscoveredResourcesFn = processStreamingStream

	exportContainerengineNodePoolHints.processDiscoveredResourcesFn = processContainerengineNodePool

	exportNosqlIndexHints.processDiscoveredResourcesFn = processNosqlIndex
	//
	//exportFileStorageMountTargetHints.requireResourceRefresh = true
	//
	exportKmsKeyHints.processDiscoveredResourcesFn = processKmsKey
	exportKmsKeyVersionHints.processDiscoveredResourcesFn = processKmsKeyVersion

	exportDnsRrsetHints.findResourcesOverrideFn = findDnsRrset
	exportDnsRrsetHints.processDiscoveredResourcesFn = processDnsRrset

	exportMysqlMysqlBackupHints.requireResourceRefresh = true
	exportMysqlMysqlBackupHints.processDiscoveredResourcesFn = filterMysqlBackups
	exportMysqlMysqlDbSystemHints.processDiscoveredResourcesFn = processMysqlDbSystem
	//
	//// Custom overrides for generating composite Network Load Balancer IDs within the resource discovery framework
	//exportNetworkLoadBalancerBackendHints.processDiscoveredResourcesFn = processNetworkLoadBalancerBackends
	//exportNetworkLoadBalancerBackendSetHints.processDiscoveredResourcesFn = processNetworkLoadBalancerBackendSets
	//exportNetworkLoadBalancerListenerHints.findResourcesOverrideFn = findNetworkLoadBalancerListeners
	//exportNetworkLoadBalancerListenerHints.processDiscoveredResourcesFn = processNetworkLoadBalancerListeners

	exportCoreDrgRouteTableRouteRuleHints.datasourceClass = "oci_core_drg_route_table_route_rules"
	exportCoreDrgRouteTableRouteRuleHints.datasourceItemsAttr = "drg_route_rules"
	exportCoreDrgRouteTableRouteRuleHints.processDiscoveredResourcesFn = processDrgRouteTableRouteRules

	//exportLogAnalyticsLogAnalyticsObjectCollectionRuleHints.findResourcesOverrideFn = findLogAnalyticsObjectCollectionRules
	//exportLogAnalyticsLogAnalyticsObjectCollectionRuleHints.processDiscoveredResourcesFn = processLogAnalyticsObjectCollectionRules

	exportCertificatesManagementCertificateAuthorityHints.processDiscoveredResourcesFn = processCertificateAuthorities
	exportCertificatesManagementCertificateHints.processDiscoveredResourcesFn = processCertificates
	exportBdsBdsInstanceApiKeyHints.processDiscoveredResourcesFn = processBdsInstanceApiKeys
}
func processBdsInstanceApiKeys(ctx *resourceDiscoveryContext, resources []*OCIResource) ([]*OCIResource, error) {
	for _, resource := range resources {
		apiKeyId := resource.id
		bdsInstanceId := resource.sourceAttributes["bds_instance_id"].(string)
		resource.importId = bds.GetBdsInstanceApiKeyCompositeId(apiKeyId, bdsInstanceId)
	}
	return resources, nil
}

/*
mergeState merges 2 json state files
*/
func mergeState(state1 interface{}, state2 interface{}) (interface{}, error) {

	state1Bytes, _ := json.MarshalIndent(state1, "", "\t")
	state2Bytes, _ := json.MarshalIndent(state2, "", "\t")

	out1 := map[string]interface{}{}
	if err := json.Unmarshal(state1Bytes, &out1); err != nil {
		return out1, fmt.Errorf("[ERROR] error occurred while generating state file for resource discovery: %s", err.Error())
	}

	out2 := map[string]interface{}{}
	if err := json.Unmarshal(state2Bytes, &out2); err != nil {
		return out1, fmt.Errorf("[ERROR] error occurred while generating state file for resource discovery: %s", err.Error())
	}

	state1resources, _ := out1["resources"].([]interface{})
	state2resources, _ := out2["resources"].([]interface{})

	out1["resources"] = append(state1resources, state2resources...)

	return out1, nil

}

func (r *resourceDiscoveryBaseStep) mergeGeneratedStateFile() error {
	if r.tempState == nil {
		return nil
	}
	utils.Debugf("[DEBUG] merging state file for %s", r.name)
	defer elapsed(fmt.Sprintf("[DEBUG] merging state file for %s", r.name), nil, 0)()
	if r.ctx.state == nil {
		// if state exists for the step, initialize the final state
		r.ctx.state = r.tempState
	} else {
		// merge the state for step to final state
		r.ctx.state, _ = mergeState(r.ctx.state, r.tempState)
	}

	return nil

}
