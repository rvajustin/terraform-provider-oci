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

func DatabaseManagementCloudClusterDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["cloud_cluster_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DatabaseManagementCloudClusterResource(), fieldMap, readSingularDatabaseManagementCloudCluster)
}

func readSingularDatabaseManagementCloudCluster(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementCloudClusterDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.ReadResource(sync)
}

type DatabaseManagementCloudClusterDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_management.DbManagementClient
	Res    *oci_database_management.GetCloudClusterResponse
}

func (s *DatabaseManagementCloudClusterDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseManagementCloudClusterDataSourceCrud) Get() error {
	request := oci_database_management.GetCloudClusterRequest{}

	if cloudClusterId, ok := s.D.GetOkExists("cloud_cluster_id"); ok {
		tmp := cloudClusterId.(string)
		request.CloudClusterId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_management")

	response, err := s.Client.GetCloudCluster(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseManagementCloudClusterDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	s.D.Set("additional_details", s.Res.AdditionalDetails)

	if s.Res.CloudConnectorId != nil {
		s.D.Set("cloud_connector_id", *s.Res.CloudConnectorId)
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

	if s.Res.GridHome != nil {
		s.D.Set("grid_home", *s.Res.GridHome)
	}

	if s.Res.IsFlexCluster != nil {
		s.D.Set("is_flex_cluster", *s.Res.IsFlexCluster)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	networkConfigurations := []interface{}{}
	for _, item := range s.Res.NetworkConfigurations {
		networkConfigurations = append(networkConfigurations, CloudClusterNetworkConfigurationToMap(item))
	}
	s.D.Set("network_configurations", networkConfigurations)

	if s.Res.OcrFileLocation != nil {
		s.D.Set("ocr_file_location", *s.Res.OcrFileLocation)
	}

	scanConfigurations := []interface{}{}
	for _, item := range s.Res.ScanConfigurations {
		scanConfigurations = append(scanConfigurations, CloudClusterScanListenerConfigurationToMap(item))
	}
	s.D.Set("scan_configurations", scanConfigurations)

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

	if s.Res.Version != nil {
		s.D.Set("version", *s.Res.Version)
	}

	vipConfigurations := []interface{}{}
	for _, item := range s.Res.VipConfigurations {
		vipConfigurations = append(vipConfigurations, CloudClusterVipConfigurationToMap(item))
	}
	s.D.Set("vip_configurations", vipConfigurations)

	return nil
}
