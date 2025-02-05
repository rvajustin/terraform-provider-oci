// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseScheduledActionsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseScheduledActions,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"scheduling_plan_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"service_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"scheduled_action_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(DatabaseScheduledActionResource()),
						},
					},
				},
			},
		},
	}
}

func readDatabaseScheduledActions(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseScheduledActionsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseScheduledActionsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListScheduledActionsResponse
}

func (s *DatabaseScheduledActionsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseScheduledActionsDataSourceCrud) Get() error {
	request := oci_database.ListScheduledActionsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	if schedulingPlanId, ok := s.D.GetOkExists("scheduling_plan_id"); ok {
		tmp := schedulingPlanId.(string)
		request.SchedulingPlanId = &tmp
	}

	if serviceType, ok := s.D.GetOkExists("service_type"); ok {
		tmp := serviceType.(string)
		request.ServiceType = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_database.ScheduledActionSummaryLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.ListScheduledActions(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListScheduledActions(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseScheduledActionsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseScheduledActionsDataSource-", DatabaseScheduledActionsDataSource(), s.D))
	resources := []map[string]interface{}{}
	scheduledAction := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ScheduledActionSummaryToMap(item))
	}
	scheduledAction["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DatabaseScheduledActionsDataSource().Schema["scheduled_action_collection"].Elem.(*schema.Resource).Schema)
		scheduledAction["items"] = items
	}

	resources = append(resources, scheduledAction)
	if err := s.D.Set("scheduled_action_collection", resources); err != nil {
		return err
	}

	return nil
}
