// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"

	// "strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_dbmulticloud "github.com/oracle/oci-go-sdk/v65/dbmulticloud"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"

	// "github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DbmulticloudOracleDbAzureVaultRequiredOnlyResource = DbmulticloudOracleDbAzureVaultResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_dbmulticloud_oracle_db_azure_vault", "test_oracle_db_azure_vault", acctest.Required, acctest.Create, DbmulticloudOracleDbAzureVaultRepresentation)

	DbmulticloudOracleDbAzureVaultResourceConfig = DbmulticloudOracleDbAzureVaultResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_dbmulticloud_oracle_db_azure_vault", "test_oracle_db_azure_vault", acctest.Optional, acctest.Update, DbmulticloudOracleDbAzureVaultRepresentation)

	DbmulticloudOracleDbAzureVaultSingularDataSourceRepresentation = map[string]interface{}{
		"oracle_db_azure_vault_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_dbmulticloud_oracle_db_azure_vault.test_oracle_db_azure_vault.id}`},
	}

	DbmulticloudOracleDbAzureVaultDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: `Discover_Tersi_Test`},
		// "oracle_db_azure_connector_id":   acctest.Representation{RepType: acctest.Required, Create: `${oci_dbmulticloud_oracle_db_azure_connector.test_oracle_db_azure_connector.id}`},
		"oracle_db_azure_resource_group": acctest.Representation{RepType: acctest.Optional, Create: `Prasanna.RG`},
		"state":                          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":                         acctest.RepresentationGroup{RepType: acctest.Required, Group: DbmulticloudOracleDbAzureVaultDataSourceFilterRepresentation}}
	DbmulticloudOracleDbAzureVaultDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_dbmulticloud_oracle_db_azure_vault.test_oracle_db_azure_vault.id}`}},
	}

	DbmulticloudOracleDbAzureVaultRepresentation = map[string]interface{}{
		"compartment_id":         acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":           acctest.Representation{RepType: acctest.Required, Create: `Discover_Tersi_Test`},
		"oracle_db_connector_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_dbmulticloud_oracle_db_azure_connector.test_oracle_db_azure_connector.id}`},
		"azure_vault_id":         acctest.Representation{RepType: acctest.Required, Create: `PrasannaHSM2`},
		// "defined_tags":                   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		// "freeform_tags":                  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"location":                       acctest.Representation{RepType: acctest.Required, Create: `eastus2`},
		"oracle_db_azure_resource_group": acctest.Representation{RepType: acctest.Required, Create: `Prasanna.RG`},
		// "properties":                     acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"properties": "properties"}, Update: map[string]string{"properties2": "properties2"}},
		"type": acctest.Representation{RepType: acctest.Required, Create: `managedHSMs`},
	}

	DbmulticloudOracleDbAzureVaultResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_dbmulticloud_oracle_db_azure_connector", "test_oracle_db_azure_connector", acctest.Required, acctest.Create, DbmulticloudOracleDbAzureConnectorRepresentation)
)

// issue-routing-tag: dbmulticloud/default
func TestDbmulticloudOracleDbAzureVaultResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDbmulticloudOracleDbAzureVaultResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_dbmulticloud_oracle_db_azure_vault.test_oracle_db_azure_vault"
	datasourceName := "data.oci_dbmulticloud_oracle_db_azure_vaults.test_oracle_db_azure_vaults"
	singularDatasourceName := "data.oci_dbmulticloud_oracle_db_azure_vault.test_oracle_db_azure_vault"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DbmulticloudOracleDbAzureVaultResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_dbmulticloud_oracle_db_azure_vault", "test_oracle_db_azure_vault", acctest.Optional, acctest.Create, DbmulticloudOracleDbAzureVaultRepresentation), "dbmulticloud", "oracleDbAzureVault", t)

	acctest.ResourceTest(t, testAccCheckDbmulticloudOracleDbAzureVaultDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DbmulticloudOracleDbAzureVaultResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dbmulticloud_oracle_db_azure_vault", "test_oracle_db_azure_vault", acctest.Required, acctest.Create, DbmulticloudOracleDbAzureVaultRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "Discover_Tersi_Test"),
				resource.TestCheckResourceAttrSet(resourceName, "oracle_db_connector_id"),
				resource.TestCheckResourceAttr(resourceName, "oracle_db_azure_resource_group", "Prasanna.RG"),
				resource.TestCheckResourceAttr(resourceName, "location", "eastus2"),
				resource.TestCheckResourceAttr(resourceName, "type", "managedHSMs"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DbmulticloudOracleDbAzureVaultResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DbmulticloudOracleDbAzureVaultResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dbmulticloud_oracle_db_azure_vault", "test_oracle_db_azure_vault", acctest.Optional, acctest.Create, DbmulticloudOracleDbAzureVaultRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "azure_vault_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "Discover_Tersi_Test"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "oracle_db_azure_resource_group", "Prasanna.RG"),
				resource.TestCheckResourceAttrSet(resourceName, "oracle_db_connector_id"),
				resource.TestCheckResourceAttr(resourceName, "location", "eastus2"),
				resource.TestCheckResourceAttr(resourceName, "type", "managedHSMs"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					// if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
					// 	if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
					// 		return errExport
					// 	}
					// }
					return err
				},
			),
		},

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DbmulticloudOracleDbAzureVaultResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dbmulticloud_oracle_db_azure_vault", "test_oracle_db_azure_vault", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DbmulticloudOracleDbAzureVaultRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "azure_vault_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "Discover_Tersi_Test"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "oracle_db_azure_resource_group", "Prasanna.RG"),
				resource.TestCheckResourceAttrSet(resourceName, "oracle_db_connector_id"),
				resource.TestCheckResourceAttr(resourceName, "location", "eastus2"),
				resource.TestCheckResourceAttr(resourceName, "type", "managedHSMs"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + DbmulticloudOracleDbAzureVaultResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dbmulticloud_oracle_db_azure_vault", "test_oracle_db_azure_vault", acctest.Optional, acctest.Update, DbmulticloudOracleDbAzureVaultRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "azure_vault_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "Discover_Tersi_Test"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "oracle_db_azure_resource_group", "Prasanna.RG"),
				resource.TestCheckResourceAttrSet(resourceName, "oracle_db_connector_id"),
				resource.TestCheckResourceAttr(resourceName, "location", "eastus2"),
				resource.TestCheckResourceAttr(resourceName, "type", "managedHSMs"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_dbmulticloud_oracle_db_azure_vaults", "test_oracle_db_azure_vaults", acctest.Optional, acctest.Update, DbmulticloudOracleDbAzureVaultDataSourceRepresentation) +
				compartmentIdVariableStr + DbmulticloudOracleDbAzureVaultResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dbmulticloud_oracle_db_azure_vault", "test_oracle_db_azure_vault", acctest.Optional, acctest.Update, DbmulticloudOracleDbAzureVaultRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "Discover_Tersi_Test"),
				// resource.TestCheckResourceAttrSet(datasourceName, "oracle_db_azure_connector_id"),
				resource.TestCheckResourceAttr(datasourceName, "oracle_db_azure_resource_group", "Prasanna.RG"),
				// resource.TestCheckResourceAttrSet(datasourceName, "oracle_db_azure_vault_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttr(resourceName, "location", "eastus2"),
				resource.TestCheckResourceAttr(resourceName, "type", "managedHSMs"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_dbmulticloud_oracle_db_azure_vault", "test_oracle_db_azure_vault", acctest.Required, acctest.Create, DbmulticloudOracleDbAzureVaultSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DbmulticloudOracleDbAzureVaultResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "oracle_db_azure_vault_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "Discover_Tersi_Test"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				// resource.TestCheckResourceAttrSet(singularDatasourceName, "last_modification"),
				resource.TestCheckResourceAttr(singularDatasourceName, "oracle_db_azure_resource_group", "Prasanna.RG"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "location", "eastus2"),
				resource.TestCheckResourceAttr(resourceName, "type", "managedHSMs"),
			),
		},
		// verify resource import
		{
			Config:                  config + DbmulticloudOracleDbAzureVaultRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDbmulticloudOracleDbAzureVaultDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).OracleDbAzureVaultClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_dbmulticloud_oracle_db_azure_vault" {
			noResourceFound = false
			request := oci_dbmulticloud.GetOracleDbAzureVaultRequest{}

			tmp := rs.Primary.ID
			request.OracleDbAzureVaultId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "dbmulticloud")

			response, err := client.GetOracleDbAzureVault(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_dbmulticloud.OracleDbAzureVaultLifecycleStateDeleted): true,
				}
				if _, ok := deletedLifecycleStates[string(response.LifecycleState)]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.LifecycleState)
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
	if !acctest.InSweeperExcludeList("DbmulticloudOracleDbAzureVault") {
		resource.AddTestSweepers("DbmulticloudOracleDbAzureVault", &resource.Sweeper{
			Name:         "DbmulticloudOracleDbAzureVault",
			Dependencies: acctest.DependencyGraph["oracleDbAzureVault"],
			F:            sweepDbmulticloudOracleDbAzureVaultResource,
		})
	}
}

func sweepDbmulticloudOracleDbAzureVaultResource(compartment string) error {
	oracleDbAzureVaultClient := acctest.GetTestClients(&schema.ResourceData{}).OracleDbAzureVaultClient()
	oracleDbAzureVaultIds, err := getDbmulticloudOracleDbAzureVaultIds(compartment)
	if err != nil {
		return err
	}
	for _, oracleDbAzureVaultId := range oracleDbAzureVaultIds {
		if ok := acctest.SweeperDefaultResourceId[oracleDbAzureVaultId]; !ok {
			deleteOracleDbAzureVaultRequest := oci_dbmulticloud.DeleteOracleDbAzureVaultRequest{}

			deleteOracleDbAzureVaultRequest.OracleDbAzureVaultId = &oracleDbAzureVaultId

			deleteOracleDbAzureVaultRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "dbmulticloud")
			_, error := oracleDbAzureVaultClient.DeleteOracleDbAzureVault(context.Background(), deleteOracleDbAzureVaultRequest)
			if error != nil {
				fmt.Printf("Error deleting OracleDbAzureVault %s %s, It is possible that the resource is already deleted. Please verify manually \n", oracleDbAzureVaultId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &oracleDbAzureVaultId, DbmulticloudOracleDbAzureVaultSweepWaitCondition, time.Duration(3*time.Minute),
				DbmulticloudOracleDbAzureVaultSweepResponseFetchOperation, "dbmulticloud", true)
		}
	}
	return nil
}

func getDbmulticloudOracleDbAzureVaultIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "OracleDbAzureVaultId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	oracleDbAzureVaultClient := acctest.GetTestClients(&schema.ResourceData{}).OracleDbAzureVaultClient()

	listOracleDbAzureVaultsRequest := oci_dbmulticloud.ListOracleDbAzureVaultsRequest{}
	listOracleDbAzureVaultsRequest.CompartmentId = &compartmentId
	listOracleDbAzureVaultsRequest.LifecycleState = oci_dbmulticloud.OracleDbAzureVaultLifecycleStateActive
	listOracleDbAzureVaultsResponse, err := oracleDbAzureVaultClient.ListOracleDbAzureVaults(context.Background(), listOracleDbAzureVaultsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting OracleDbAzureVault list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, oracleDbAzureVault := range listOracleDbAzureVaultsResponse.Items {
		id := *oracleDbAzureVault.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "OracleDbAzureVaultId", id)
	}
	return resourceIds, nil
}

func DbmulticloudOracleDbAzureVaultSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is ACTIVE beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if oracleDbAzureVaultResponse, ok := response.Response.(oci_dbmulticloud.GetOracleDbAzureVaultResponse); ok {
		return oracleDbAzureVaultResponse.LifecycleState != oci_dbmulticloud.OracleDbAzureVaultLifecycleStateDeleted
	}
	return false
}

func DbmulticloudOracleDbAzureVaultSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.OracleDbAzureVaultClient().GetOracleDbAzureVault(context.Background(), oci_dbmulticloud.GetOracleDbAzureVaultRequest{
		OracleDbAzureVaultId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
