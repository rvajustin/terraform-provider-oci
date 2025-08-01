// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_management

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database_management "github.com/oracle/oci-go-sdk/v65/databasemanagement"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseManagementCloudAsmConfigurationDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDatabaseManagementCloudAsmConfiguration,
		Schema: map[string]*schema.Schema{
			"cloud_asm_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"opc_named_credential_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			// Computed
			"init_parameters": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"asm_instance_display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"asm_instance_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"auto_mount_disk_groups": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"disk_discovery_path": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"preferred_read_failure_groups": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"rebalance_power": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readSingularDatabaseManagementCloudAsmConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementCloudAsmConfigurationDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.ReadResource(sync)
}

type DatabaseManagementCloudAsmConfigurationDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_management.DbManagementClient
	Res    *oci_database_management.GetCloudAsmConfigurationResponse
}

func (s *DatabaseManagementCloudAsmConfigurationDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseManagementCloudAsmConfigurationDataSourceCrud) Get() error {
	request := oci_database_management.GetCloudAsmConfigurationRequest{}

	if cloudAsmId, ok := s.D.GetOkExists("cloud_asm_id"); ok {
		tmp := cloudAsmId.(string)
		request.CloudAsmId = &tmp
	}

	if opcNamedCredentialId, ok := s.D.GetOkExists("opc_named_credential_id"); ok {
		tmp := opcNamedCredentialId.(string)
		request.OpcNamedCredentialId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_management")

	response, err := s.Client.GetCloudAsmConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseManagementCloudAsmConfigurationDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseManagementCloudAsmConfigurationDataSource-", DatabaseManagementCloudAsmConfigurationDataSource(), s.D))

	initParameters := []interface{}{}
	for _, item := range s.Res.InitParameters {
		initParameters = append(initParameters, CloudAsmInstanceParametersToMap(item))
	}
	s.D.Set("init_parameters", initParameters)

	return nil
}

func CloudAsmInstanceParametersToMap(obj oci_database_management.CloudAsmInstanceParameters) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AsmInstanceDisplayName != nil {
		result["asm_instance_display_name"] = string(*obj.AsmInstanceDisplayName)
	}

	if obj.AsmInstanceId != nil {
		result["asm_instance_id"] = string(*obj.AsmInstanceId)
	}

	result["auto_mount_disk_groups"] = obj.AutoMountDiskGroups

	if obj.DiskDiscoveryPath != nil {
		result["disk_discovery_path"] = string(*obj.DiskDiscoveryPath)
	}

	result["preferred_read_failure_groups"] = obj.PreferredReadFailureGroups

	if obj.RebalancePower != nil {
		result["rebalance_power"] = int(*obj.RebalancePower)
	}

	return result
}
