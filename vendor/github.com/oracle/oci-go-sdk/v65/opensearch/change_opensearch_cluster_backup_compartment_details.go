// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OpenSearch Service API
//
// The OpenSearch service API provides access to OCI Search Service with OpenSearch.
//

package opensearch

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ChangeOpensearchClusterBackupCompartmentDetails Details about the compartment that the cluster backup should move to.
type ChangeOpensearchClusterBackupCompartmentDetails struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment
	// into which the cluster backup should be moved.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`
}

func (m ChangeOpensearchClusterBackupCompartmentDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ChangeOpensearchClusterBackupCompartmentDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
