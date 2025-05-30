// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// Use the Core Services API to manage resources such as virtual cloud networks (VCNs),
// compute instances, and block storage volumes. For more information, see the console
// documentation for the Networking (https://docs.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.oracle.com/iaas/Content/Block/Concepts/overview.htm) services.
// The required permissions are documented in the
// Details for the Core Services (https://docs.oracle.com/iaas/Content/Identity/Reference/corepolicyreference.htm) article.
//

package core

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateDedicatedVmHostDetails The details for creating a new dedicated virtual machine host.
type CreateDedicatedVmHostDetails struct {

	// The availability domain of the dedicated virtual machine host.
	// Example: `Uocm:PHX-AD-1`
	AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain"`

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The dedicated virtual machine host shape. The shape determines the number of CPUs and
	// other resources available for VM instances launched on the dedicated virtual machine host.
	DedicatedVmHostShape *string `mandatory:"true" json:"dedicatedVmHostShape"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The fault domain for the dedicated virtual machine host's assigned instances.
	// For more information, see Fault Domains (https://docs.oracle.com/iaas/Content/General/Concepts/regions.htm#fault).
	// If you do not specify the fault domain, the system selects one for you. To change the fault domain for a dedicated virtual machine host,
	// delete it and create a new dedicated virtual machine host in the preferred fault domain.
	// To get a list of fault domains, use the `ListFaultDomains` operation in
	// the Identity and Access Management Service API (https://docs.oracle.com/iaas/api/#/en/identity/20160918/).
	// Example: `FAULT-DOMAIN-1`
	FaultDomain *string `mandatory:"false" json:"faultDomain"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	PlacementConstraintDetails PlacementConstraintDetails `mandatory:"false" json:"placementConstraintDetails"`
}

func (m CreateDedicatedVmHostDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateDedicatedVmHostDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CreateDedicatedVmHostDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DefinedTags                map[string]map[string]interface{} `json:"definedTags"`
		DisplayName                *string                           `json:"displayName"`
		FaultDomain                *string                           `json:"faultDomain"`
		FreeformTags               map[string]string                 `json:"freeformTags"`
		PlacementConstraintDetails placementconstraintdetails        `json:"placementConstraintDetails"`
		AvailabilityDomain         *string                           `json:"availabilityDomain"`
		CompartmentId              *string                           `json:"compartmentId"`
		DedicatedVmHostShape       *string                           `json:"dedicatedVmHostShape"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DefinedTags = model.DefinedTags

	m.DisplayName = model.DisplayName

	m.FaultDomain = model.FaultDomain

	m.FreeformTags = model.FreeformTags

	nn, e = model.PlacementConstraintDetails.UnmarshalPolymorphicJSON(model.PlacementConstraintDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.PlacementConstraintDetails = nn.(PlacementConstraintDetails)
	} else {
		m.PlacementConstraintDetails = nil
	}

	m.AvailabilityDomain = model.AvailabilityDomain

	m.CompartmentId = model.CompartmentId

	m.DedicatedVmHostShape = model.DedicatedVmHostShape

	return
}
