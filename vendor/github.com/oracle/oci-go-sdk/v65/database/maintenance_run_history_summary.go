// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// MaintenanceRunHistorySummary Details of a maintenance run history.
type MaintenanceRunHistorySummary struct {

	// The OCID of the maintenance run history.
	Id *string `mandatory:"true" json:"id"`

	MaintenanceRunDetails *MaintenanceRunSummary `mandatory:"false" json:"maintenanceRunDetails"`

	// List of database server history details.
	DbServersHistoryDetails []DbServerHistorySummary `mandatory:"false" json:"dbServersHistoryDetails"`

	// The OCID of the current execution window.
	CurrentExecutionWindow *string `mandatory:"false" json:"currentExecutionWindow"`

	// The list of granular maintenance history details.
	GranularMaintenanceHistory []GranularMaintenanceHistoryDetails `mandatory:"false" json:"granularMaintenanceHistory"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m MaintenanceRunHistorySummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MaintenanceRunHistorySummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
