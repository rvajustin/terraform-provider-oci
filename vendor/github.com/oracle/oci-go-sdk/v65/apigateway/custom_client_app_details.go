// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// API Gateway API
//
// API for the API Gateway service. Use this API to manage gateways, deployments, and related items.
// For more information, see
// Overview of API Gateway (https://docs.oracle.com/iaas/Content/APIGateway/Concepts/apigatewayoverview.htm).
//

package apigateway

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CustomClientAppDetails Client App Credentials to be provided again.
type CustomClientAppDetails struct {

	// Client ID for the OAuth2/OIDC app.
	ClientId *string `mandatory:"true" json:"clientId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Vault Service secret resource.
	ClientSecretId *string `mandatory:"true" json:"clientSecretId"`

	// The version number of the client secret to use.
	ClientSecretVersionNumber *int64 `mandatory:"true" json:"clientSecretVersionNumber"`
}

func (m CustomClientAppDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CustomClientAppDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CustomClientAppDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCustomClientAppDetails CustomClientAppDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeCustomClientAppDetails
	}{
		"CUSTOM",
		(MarshalTypeCustomClientAppDetails)(m),
	}

	return json.Marshal(&s)
}
