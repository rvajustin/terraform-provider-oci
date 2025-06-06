// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GoldenGate API
//
// Use the Oracle Cloud Infrastructure GoldenGate APIs to perform data replication operations.
//

package goldengate

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateRestIcebergCatalogDetails The information to update a Rest catalog used in the Iceberg connection.
type UpdateRestIcebergCatalogDetails struct {

	// The base URL for the REST Catalog API.
	// e.g.: 'https://my-rest-catalog.example.com/api/v1'
	Uri *string `mandatory:"false" json:"uri"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret that stores the content
	// of the configuration file containing additional properties for the REST catalog.
	// See documentation: https://docs.oracle.com/en-us/iaas/Content/Identity/Tasks/managingcredentials.htm
	PropertiesSecretId *string `mandatory:"false" json:"propertiesSecretId"`
}

func (m UpdateRestIcebergCatalogDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateRestIcebergCatalogDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UpdateRestIcebergCatalogDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateRestIcebergCatalogDetails UpdateRestIcebergCatalogDetails
	s := struct {
		DiscriminatorParam string `json:"catalogType"`
		MarshalTypeUpdateRestIcebergCatalogDetails
	}{
		"REST",
		(MarshalTypeUpdateRestIcebergCatalogDetails)(m),
	}

	return json.Marshal(&s)
}
