// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_identity_data_plane "github.com/oracle/oci-go-sdk/v54/identitydataplane"
)

func init() {
	RegisterResource("oci_identity_data_plane_generate_scoped_access_token", IdentityDataPlaneGenerateScopedAccessTokenResource())
}

func IdentityDataPlaneGenerateScopedAccessTokenResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: DefaultTimeout,
		Create:   createIdentityDataPlaneGenerateScopedAccessToken,
		Read:     readIdentityDataPlaneGenerateScopedAccessToken,
		Delete:   deleteIdentityDataPlaneGenerateScopedAccessToken,
		Schema: map[string]*schema.Schema{
			// Required
			"public_key": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"scope": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional

			// Computed
			"token": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createIdentityDataPlaneGenerateScopedAccessToken(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDataPlaneGenerateScopedAccessTokenResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).dataplaneClient()

	return CreateResource(d, sync)
}

func readIdentityDataPlaneGenerateScopedAccessToken(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteIdentityDataPlaneGenerateScopedAccessToken(d *schema.ResourceData, m interface{}) error {
	return nil
}

type IdentityDataPlaneGenerateScopedAccessTokenResourceCrud struct {
	BaseCrud
	Client                 *oci_identity_data_plane.DataplaneClient
	Res                    *oci_identity_data_plane.SecurityToken
	DisableNotFoundRetries bool
}

func (s *IdentityDataPlaneGenerateScopedAccessTokenResourceCrud) ID() string {
	return *s.Res.Token
}

func (s *IdentityDataPlaneGenerateScopedAccessTokenResourceCrud) Create() error {
	request := oci_identity_data_plane.GenerateScopedAccessTokenRequest{}

	if publicKey, ok := s.D.GetOkExists("public_key"); ok {
		tmp := publicKey.(string)
		request.PublicKey = &tmp
	}

	if scope, ok := s.D.GetOkExists("scope"); ok {
		tmp := scope.(string)
		request.Scope = &tmp
	}

	request.RequestMetadata.RetryPolicy = GetRetryPolicy(s.DisableNotFoundRetries, "identity_data_plane")

	response, err := s.Client.GenerateScopedAccessToken(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.SecurityToken
	return nil
}

func (s *IdentityDataPlaneGenerateScopedAccessTokenResourceCrud) SetData() error {
	if s.Res.Token != nil {
		s.D.Set("token", *s.Res.Token)
	}

	return nil
}
