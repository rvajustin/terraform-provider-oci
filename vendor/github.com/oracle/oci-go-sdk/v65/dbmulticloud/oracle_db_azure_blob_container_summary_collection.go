// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Database MultiCloud Data plane Integration
//
// 1. Oracle Azure Connector Resource: This is for installing Azure Arc Server in ExaCS VM Cluster.
//   There are two way to install Azure Arc Server (Azure Identity) in ExaCS VMCluster.
//     a. Using Bearer Access Token or
//     b. By providing Authentication token
// 2. Oracle Azure Blob Container Resource: This is for to capture Azure Container details
//    and same will be used in multiple ExaCS VMCluster to mount the Azure Container.
// 3. Oracle Azure Blob Mount Resource: This is for to mount Azure Container in ExaCS VMCluster
//    using Oracle Azure Connector and Oracle Azure Blob Container Resource.
//

package dbmulticloud

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OracleDbAzureBlobContainerSummaryCollection List of OracleDbAzureBlobContainerSummary resources.
type OracleDbAzureBlobContainerSummaryCollection struct {

	// This list contains the summary of Oracle DB Azure Blob Container resources.
	Items []OracleDbAzureBlobContainerSummary `mandatory:"true" json:"items"`
}

func (m OracleDbAzureBlobContainerSummaryCollection) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OracleDbAzureBlobContainerSummaryCollection) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
