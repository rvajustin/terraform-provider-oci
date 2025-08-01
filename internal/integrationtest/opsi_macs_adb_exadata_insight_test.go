// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_opsi "github.com/oracle/oci-go-sdk/v65/opsi"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	OpsiMacsAdbExadataInsight = OpsiMacsAdbExadataInsightResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_opsi_exadata_insight", "test_exadata_insight", acctest.Required, acctest.Create, OpsiMacsAdbExadataInsightRepresentation)

	OpsiMacsAdbExadataInsightResourceConfig = OpsiMacsAdbExadataInsightResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_opsi_exadata_insight", "test_exadata_insight", acctest.Optional, acctest.Update, OpsiMacsAdbExadataInsightRepresentation)

	OpsiMacsAdbExadataInsightSingularDataSourceRepresentation = map[string]interface{}{
		"exadata_insight_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_opsi_exadata_insight.test_exadata_insight.id}`},
	}

	OpsiMacsAdbExadataInsightDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"compartment_id_in_subtree": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"exadata_type":              acctest.Representation{RepType: acctest.Optional, Create: []string{`EXACC`}},
		"id":                        acctest.Representation{RepType: acctest.Optional, Create: `${oci_opsi_exadata_insight.test_exadata_insight.id}`},
		"state":                     acctest.Representation{RepType: acctest.Optional, Create: []string{`ACTIVE`}},
		"status":                    acctest.Representation{RepType: acctest.Optional, Create: []string{`ENABLED`}, Update: []string{`DISABLED`}},
		"filter":                    acctest.RepresentationGroup{RepType: acctest.Required, Group: OpsiMacsAdbExadataInsightDataSourceFilterRepresentation}}

	OpsiMacsAdbExadataInsightDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_opsi_exadata_insight.test_exadata_insight.id}`}},
	}

	OpsiMacsAdbExadataInsightRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"exadata_infra_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.exadata_infra_id}`},
		"entity_source":             acctest.Representation{RepType: acctest.Required, Create: `MACS_MANAGED_CLOUD_EXADATA`},
		"status":                    acctest.Representation{RepType: acctest.Optional, Create: `ENABLED`, Update: `DISABLED`},
		"defined_tags":              acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":             acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":                 acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreChangesOpsiMacsAdbExadataInsightRepresentation},
		"member_vm_cluster_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: OpsiMacsAdbExadataInsightMemberVmClusterDetailsRepresentation},
	}

	OpsiMacsAdbExadataInsightMemberVmClusterDetailsRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"vmcluster_id":              acctest.Representation{RepType: acctest.Optional, Create: `${var.vmcluster_id}`},
		"member_autonomous_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: OpsiMacsAdbExadataInsightMemberVmClusterDetailsMemberAdbDetailsRepresentation},
	}

	OpsiMacsAdbExadataInsightMemberVmClusterDetailsMemberAdbDetailsRepresentation = map[string]interface{}{
		"compartment_id":                acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"connection_credential_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: OpsiMacsAdbExadataInsightNamedCredentialDetailsRepresentation},
		"connection_details":            acctest.RepresentationGroup{RepType: acctest.Required, Group: OpsiMacsAdbExadataInsightConnectionDetailsRepresentation},
		"database_id":                   acctest.Representation{RepType: acctest.Optional, Create: `${var.macs_adb_id}`},
		"database_resource_type":        acctest.Representation{RepType: acctest.Optional, Create: `autonomousdatabase`},
		"deployment_type":               acctest.Representation{RepType: acctest.Optional, Create: `EXACC`},
		"entity_source":                 acctest.Representation{RepType: acctest.Optional, Create: `MACS_MANAGED_AUTONOMOUS_DATABASE`},
		"freeform_tags":                 acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"freeformTags": "freeformTags"}},
		"management_agent_id":           acctest.Representation{RepType: acctest.Optional, Create: `${var.management_agent_id}`},
	}

	OpsiMacsAdbExadataInsightNamedCredentialDetailsRepresentation = map[string]interface{}{
		"credential_type":     acctest.Representation{RepType: acctest.Required, Create: `CREDENTIALS_BY_NAMED_CREDS`},
		"named_credential_id": acctest.Representation{RepType: acctest.Required, Create: `${var.named_credential_id}`},
	}

	OpsiMacsAdbExadataInsightConnectionDetailsRepresentation = map[string]interface{}{
		"host_name":    acctest.Representation{RepType: acctest.Required, Create: `${var.adb_host}`},
		"port":         acctest.Representation{RepType: acctest.Required, Create: `${var.adb_port}`},
		"protocol":     acctest.Representation{RepType: acctest.Required, Create: `TCP`},
		"service_name": acctest.Representation{RepType: acctest.Required, Create: `${var.service_name}`},
	}

	ignoreChangesOpsiMacsAdbExadataInsightRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`, `freeform_tags`}},
	}

	OpsiMacsAdbExadataInsightResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: opsi/controlPlane
func TestOpsiMacsAdbExadataInsightResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOpsiMacsAdbExadataInsightResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	exacsAvailable := utils.GetEnvSettingWithBlankDefault("opsi_exacs_available")
	if exacsAvailable == "" {
		// Suggested by Rakshith Siddanahalli for our case of required but scarce ExaCS systems
		t.Skip("Skipping tests which require ExaCS")
	}

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	exadataInfraId := utils.GetEnvSettingWithBlankDefault("exadata_infra_id")
	exadataInfraIdVariableStr := fmt.Sprintf("variable \"exadata_infra_id\" { default = \"%s\" }\n", exadataInfraId)

	vmClusterId := utils.GetEnvSettingWithBlankDefault("vmcluster_id")
	vmClusterIdVariableStr := fmt.Sprintf("variable \"vmcluster_id\" { default = \"%s\" }\n", vmClusterId)

	managementAgentId := utils.GetEnvSettingWithBlankDefault("management_agent_id")
	managementAgentIdVariableStr := fmt.Sprintf("variable \"management_agent_id\" { default = \"%s\" }\n", managementAgentId)

	namedCredentialId := utils.GetEnvSettingWithBlankDefault("named_credential_id")
	namedCredentialIdVariableStr := fmt.Sprintf("variable \"named_credential_id\" { default = \"%s\" }\n", namedCredentialId)

	macsDatabaseId := utils.GetEnvSettingWithBlankDefault("macs_adb_id")
	macsDatabaseIdVariableStr := fmt.Sprintf("variable \"macs_adb_id\" { default = \"%s\" }\n", macsDatabaseId)

	serviceName := utils.GetEnvSettingWithBlankDefault("service_name")
	serviceNameVariableStr := fmt.Sprintf("variable \"service_name\" { default = \"%s\" }\n", serviceName)

	adbHostName := utils.GetEnvSettingWithBlankDefault("adb_host")
	adbHostNameVariableStr := fmt.Sprintf("variable \"adb_host\" { default = \"%s\" }\n", adbHostName)

	adbPort := utils.GetEnvSettingWithBlankDefault("adb_port")
	adbPortVariableStr := fmt.Sprintf("variable \"adb_port\" { default = \"%s\" }\n", adbPort)

	envVarsString := compartmentIdVariableStr + compartmentIdUVariableStr + managementAgentIdVariableStr + exadataInfraIdVariableStr + vmClusterIdVariableStr + namedCredentialIdVariableStr + macsDatabaseIdVariableStr + serviceNameVariableStr +
		adbPortVariableStr + adbHostNameVariableStr

	resourceName := "oci_opsi_exadata_insight.test_exadata_insight"
	datasourceName := "data.oci_opsi_exadata_insights.test_exadata_insights"
	singularDatasourceName := "data.oci_opsi_exadata_insight.test_exadata_insight"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+envVarsString+OpsiMacsAdbExadataInsightResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_opsi_exadata_insight", "test_exadata_insight", acctest.Optional, acctest.Create, OpsiMacsAdbExadataInsightRepresentation), "opsi", "exadataInsight", t)

	acctest.ResourceTest(t, testAccCheckOpsiMacsAdbExadataInsightDestroy, []resource.TestStep{
		// verify Create with optionals - Step 0
		{
			Config: config + envVarsString + OpsiMacsAdbExadataInsightResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opsi_exadata_insight", "test_exadata_insight", acctest.Optional, acctest.Create, OpsiMacsAdbExadataInsightRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "exadata_infra_id", exadataInfraId),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "entity_source", "MACS_MANAGED_CLOUD_EXADATA"),
				resource.TestCheckResourceAttrSet(resourceName, "exadata_name"),
				resource.TestCheckResourceAttrSet(resourceName, "exadata_type"),
				resource.TestCheckResourceAttrSet(resourceName, "exadata_shape"),
				resource.TestCheckResourceAttrSet(resourceName, "exadata_infra_resource_type"),
				resource.TestCheckResourceAttrSet(resourceName, "exadata_rack_type"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "status"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},
		// verify update to the compartment (the compartment will be switched back in the next step) - Step 1
		{
			Config: config + envVarsString + OpsiMacsAdbExadataInsightResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opsi_exadata_insight", "test_exadata_insight", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(OpsiMacsAdbExadataInsightRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "exadata_infra_id", exadataInfraId),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "entity_source", "MACS_MANAGED_CLOUD_EXADATA"),
				resource.TestCheckResourceAttrSet(resourceName, "exadata_name"),
				resource.TestCheckResourceAttrSet(resourceName, "exadata_type"),
				resource.TestCheckResourceAttrSet(resourceName, "exadata_shape"),
				resource.TestCheckResourceAttrSet(resourceName, "exadata_infra_resource_type"),
				resource.TestCheckResourceAttrSet(resourceName, "exadata_rack_type"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "status"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Step 1: resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},
		// verify updates to updatable parameters - Step 2 (Update causes status to go to disabled)
		{
			Config: config + envVarsString + OpsiMacsAdbExadataInsightResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opsi_exadata_insight", "test_exadata_insight", acctest.Optional, acctest.Update, OpsiMacsAdbExadataInsightRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "exadata_infra_id", exadataInfraId),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "entity_source", "MACS_MANAGED_CLOUD_EXADATA"),
				resource.TestCheckResourceAttr(resourceName, "status", "DISABLED"),
				resource.TestCheckResourceAttrSet(resourceName, "exadata_name"),
				resource.TestCheckResourceAttrSet(resourceName, "exadata_type"),
				resource.TestCheckResourceAttrSet(resourceName, "exadata_shape"),
				resource.TestCheckResourceAttrSet(resourceName, "exadata_infra_resource_type"),
				resource.TestCheckResourceAttrSet(resourceName, "exadata_rack_type"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Step 2: resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},
		//verify enable - Step 3
		{
			Config: config + envVarsString + OpsiMacsAdbExadataInsightResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opsi_exadata_insight", "test_exadata_insight", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(OpsiMacsAdbExadataInsightRepresentation, map[string]interface{}{
						"status": acctest.Representation{RepType: acctest.Required, Update: `ENABLED`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "status", "ENABLED"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Step 3: resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},
		// verify datasource - Step 4
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_opsi_exadata_insights", "test_exadata_insights", acctest.Optional, acctest.Update, OpsiMacsAdbExadataInsightDataSourceRepresentation) +
				envVarsString + OpsiMacsAdbExadataInsightResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opsi_exadata_insight", "test_exadata_insight", acctest.Optional, acctest.Update, OpsiMacsAdbExadataInsightRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id_in_subtree", "false"),
				resource.TestCheckResourceAttr(datasourceName, "exadata_type.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "state.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "status.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "filter.#", "1"),
			),
		},
		//verify singular datasource - Step 5
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_opsi_exadata_insight", "test_exadata_insight", acctest.Required, acctest.Create, OpsiMacsAdbExadataInsightSingularDataSourceRepresentation) +
				envVarsString + OpsiMacsAdbExadataInsightResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "exadata_insight_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "exadata_infra_id", exadataInfraId),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "entity_source", "MACS_MANAGED_CLOUD_EXADATA"),
				resource.TestCheckResourceAttr(singularDatasourceName, "status", "DISABLED"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "exadata_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "exadata_display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "exadata_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "exadata_shape"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "exadata_infra_resource_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "exadata_rack_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import - Step 6
		{
			Config:            config + OpsiMacsAdbExadataInsight,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"member_vm_cluster_details",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckOpsiMacsAdbExadataInsightDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).OperationsInsightsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_opsi_exadata_insight" {
			noResourceFound = false
			request := oci_opsi.GetExadataInsightRequest{}

			tmp := rs.Primary.ID
			request.ExadataInsightId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "opsi")

			response, err := client.GetExadataInsight(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_opsi.ExadataInsightLifecycleStateDeleted): true,
				}
				if _, ok := deletedLifecycleStates[string(response.GetLifecycleState())]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.GetLifecycleState())
				}
				//resource lifecycle state is in expected deleted lifecycle states. continue with next one.
				continue
			}

			//Verify that exception is for '404 not found'.
			if failure, isServiceError := common.IsServiceError(err); !isServiceError || failure.GetHTTPStatusCode() != 404 {
				return err
			}
		}
	}
	if noResourceFound {
		return fmt.Errorf("at least one resource was expected from the state file, but could not be found")
	}

	return nil
}

func init() {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("OpsiMacsAdbExadataInsight") {
		resource.AddTestSweepers("OpsiMacsAdbExadataInsight", &resource.Sweeper{
			Name:         "OpsiMacsAdbExadataInsight",
			Dependencies: acctest.DependencyGraph["exadataInsight"],
			F:            sweepOpsiMacsAdbExadataInsightResource,
		})
	}
}

func sweepOpsiMacsAdbExadataInsightResource(compartment string) error {
	operationsInsightsClient := acctest.GetTestClients(&schema.ResourceData{}).OperationsInsightsClient()
	exadataInsightIds, err := getOpsiMacsAdbExadataInsightIds(compartment)
	if err != nil {
		return err
	}
	for _, exadataInsightId := range exadataInsightIds {
		if ok := acctest.SweeperDefaultResourceId[exadataInsightId]; !ok {
			deleteExadataInsightRequest := oci_opsi.DeleteExadataInsightRequest{}

			deleteExadataInsightRequest.ExadataInsightId = &exadataInsightId

			deleteExadataInsightRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "opsi")
			_, error := operationsInsightsClient.DeleteExadataInsight(context.Background(), deleteExadataInsightRequest)
			if error != nil {
				fmt.Printf("Error deleting ExadataInsight %s %s, It is possible that the resource is already deleted. Please verify manually \n", exadataInsightId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &exadataInsightId, OpsiMacsAdbExadataInsightSweepWaitCondition, time.Duration(3*time.Minute),
				OpsiMacsAdbExadataInsightSweepResponseFetchOperation, "opsi", true)
		}
	}
	return nil
}

func getOpsiMacsAdbExadataInsightIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ExadataInsightId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	operationsInsightsClient := acctest.GetTestClients(&schema.ResourceData{}).OperationsInsightsClient()

	listExadataInsightsRequest := oci_opsi.ListExadataInsightsRequest{}
	listExadataInsightsRequest.CompartmentId = &compartmentId
	listExadataInsightsRequest.LifecycleState = []oci_opsi.LifecycleStateEnum{oci_opsi.LifecycleStateActive}
	listExadataInsightsResponse, err := operationsInsightsClient.ListExadataInsights(context.Background(), listExadataInsightsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting ExadataInsight list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, exadataInsight := range listExadataInsightsResponse.Items {
		id := *exadataInsight.GetId()
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ExadataInsightId", id)
	}
	return resourceIds, nil
}

func OpsiMacsAdbExadataInsightSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if exadataInsightResponse, ok := response.Response.(oci_opsi.GetExadataInsightResponse); ok {
		return exadataInsightResponse.GetLifecycleState() != oci_opsi.ExadataInsightLifecycleStateDeleted
	}
	return false
}

func OpsiMacsAdbExadataInsightSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.OperationsInsightsClient().GetExadataInsight(context.Background(), oci_opsi.GetExadataInsightRequest{
		ExadataInsightId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
