// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Usage API
//
// Use the Usage API to view your Oracle Cloud usage and costs. The API allows you to request data that meets the specified filter criteria, and to group that data by the chosen dimension. The Usage API is used by Cost Analysis (https://docs.oracle.com/iaas/Content/Billing/Concepts/costanalysisoverview.htm), Scheduled Reports (https://docs.oracle.com/iaas/Content/Billing/Concepts/scheduledreportoverview.htm), and Carbon Emissions Analysis (https://docs.oracle.com/iaas/Content/General/Concepts/emissions-management.htm) in the Console. Also see Using the Usage API (https://docs.oracle.com/iaas/Content/Billing/Concepts/costanalysisoverview.htm#cost_analysis_using_the_api) for more information.
//

package usageapi

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CustomTableSummary Custom table in the list request.
type CustomTableSummary struct {

	// The custom table OCID.
	Id *string `mandatory:"true" json:"id"`

	SavedCustomTable *SavedCustomTable `mandatory:"true" json:"savedCustomTable"`
}

func (m CustomTableSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CustomTableSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
