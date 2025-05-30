// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AutonomousDatabaseSqlWatchFeatureDetails The details required to enable the SQL Watch feature.
type AutonomousDatabaseSqlWatchFeatureDetails struct {
	DatabaseConnectionDetails *DatabaseConnectionDetails `mandatory:"true" json:"databaseConnectionDetails"`

	ConnectorDetails ConnectorDetails `mandatory:"false" json:"connectorDetails"`
}

// GetDatabaseConnectionDetails returns DatabaseConnectionDetails
func (m AutonomousDatabaseSqlWatchFeatureDetails) GetDatabaseConnectionDetails() *DatabaseConnectionDetails {
	return m.DatabaseConnectionDetails
}

// GetConnectorDetails returns ConnectorDetails
func (m AutonomousDatabaseSqlWatchFeatureDetails) GetConnectorDetails() ConnectorDetails {
	return m.ConnectorDetails
}

func (m AutonomousDatabaseSqlWatchFeatureDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AutonomousDatabaseSqlWatchFeatureDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m AutonomousDatabaseSqlWatchFeatureDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeAutonomousDatabaseSqlWatchFeatureDetails AutonomousDatabaseSqlWatchFeatureDetails
	s := struct {
		DiscriminatorParam string `json:"feature"`
		MarshalTypeAutonomousDatabaseSqlWatchFeatureDetails
	}{
		"SQLWATCH",
		(MarshalTypeAutonomousDatabaseSqlWatchFeatureDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *AutonomousDatabaseSqlWatchFeatureDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		ConnectorDetails          connectordetails           `json:"connectorDetails"`
		DatabaseConnectionDetails *DatabaseConnectionDetails `json:"databaseConnectionDetails"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.ConnectorDetails.UnmarshalPolymorphicJSON(model.ConnectorDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ConnectorDetails = nn.(ConnectorDetails)
	} else {
		m.ConnectorDetails = nil
	}

	m.DatabaseConnectionDetails = model.DatabaseConnectionDetails

	return
}
