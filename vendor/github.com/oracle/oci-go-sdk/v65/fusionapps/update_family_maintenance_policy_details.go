// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fusion Applications Environment Management API
//
// Use the Fusion Applications Environment Management API to manage the environments where your Fusion Applications run. For more information, see the Fusion Applications Environment Management documentation (https://docs.oracle.com/iaas/Content/fusion-applications/home.htm).
//

package fusionapps

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateFamilyMaintenancePolicyDetails The editable settings of the policy that specifies the maintenance and upgrade preferences for an environment.
type UpdateFamilyMaintenancePolicyDetails struct {

	// Whether the Fusion environment receives monthly patching.
	IsMonthlyPatchingEnabled *bool `mandatory:"false" json:"isMonthlyPatchingEnabled"`

	// Whether production and non-production environments are upgraded concurrently.
	ConcurrentMaintenance FamilyMaintenancePolicyConcurrentMaintenanceEnum `mandatory:"false" json:"concurrentMaintenance,omitempty"`
}

func (m UpdateFamilyMaintenancePolicyDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateFamilyMaintenancePolicyDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingFamilyMaintenancePolicyConcurrentMaintenanceEnum(string(m.ConcurrentMaintenance)); !ok && m.ConcurrentMaintenance != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ConcurrentMaintenance: %s. Supported values are: %s.", m.ConcurrentMaintenance, strings.Join(GetFamilyMaintenancePolicyConcurrentMaintenanceEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
