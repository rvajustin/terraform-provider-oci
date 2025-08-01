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

func DatabaseManagementCloudClusterInstanceDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["cloud_cluster_instance_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DatabaseManagementCloudClusterInstanceResource(), fieldMap, readSingularDatabaseManagementCloudClusterInstance)
}

func readSingularDatabaseManagementCloudClusterInstance(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementCloudClusterInstanceDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.ReadResource(sync)
}

type DatabaseManagementCloudClusterInstanceDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_management.DbManagementClient
	Res    *oci_database_management.GetCloudClusterInstanceResponse
}

func (s *DatabaseManagementCloudClusterInstanceDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseManagementCloudClusterInstanceDataSourceCrud) Get() error {
	request := oci_database_management.GetCloudClusterInstanceRequest{}

	if cloudClusterInstanceId, ok := s.D.GetOkExists("cloud_cluster_instance_id"); ok {
		tmp := cloudClusterInstanceId.(string)
		request.CloudClusterInstanceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_management")

	response, err := s.Client.GetCloudClusterInstance(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseManagementCloudClusterInstanceDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.AdrHomeDirectory != nil {
		s.D.Set("adr_home_directory", *s.Res.AdrHomeDirectory)
	}

	if s.Res.CloudClusterId != nil {
		s.D.Set("cloud_cluster_id", *s.Res.CloudClusterId)
	}

	if s.Res.CloudConnectorId != nil {
		s.D.Set("cloud_connector_id", *s.Res.CloudConnectorId)
	}

	if s.Res.CloudDbNodeId != nil {
		s.D.Set("cloud_db_node_id", *s.Res.CloudDbNodeId)
	}

	if s.Res.CloudDbSystemId != nil {
		s.D.Set("cloud_db_system_id", *s.Res.CloudDbSystemId)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ComponentName != nil {
		s.D.Set("component_name", *s.Res.ComponentName)
	}

	if s.Res.CrsBaseDirectory != nil {
		s.D.Set("crs_base_directory", *s.Res.CrsBaseDirectory)
	}

	if s.Res.DbaasId != nil {
		s.D.Set("dbaas_id", *s.Res.DbaasId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.HostName != nil {
		s.D.Set("host_name", *s.Res.HostName)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("node_role", s.Res.NodeRole)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
