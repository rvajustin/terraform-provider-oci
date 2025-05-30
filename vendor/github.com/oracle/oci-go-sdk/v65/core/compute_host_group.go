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
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ComputeHostGroup Detail information for a compute host group.
type ComputeHostGroup struct {

	// The availability domain of a host group.
	// Example: `Uocm:PHX-AD-1`
	AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID of the compartment that contains host group.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The date and time the host group was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The lifecycle state of the host group
	LifecycleState ComputeHostGroupLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the Customer-unique host group
	Id *string `mandatory:"false" json:"id"`

	// A list of HostGroupConfiguration objects
	Configurations []HostGroupConfiguration `mandatory:"false" json:"configurations"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// The date and time the host group was updated, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A flag that allows customers to restrict placement for hosts attached to the group. If true, the only way to place on hosts is to target the specific host group.
	IsTargetedPlacementRequired *bool `mandatory:"false" json:"isTargetedPlacementRequired"`
}

func (m ComputeHostGroup) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ComputeHostGroup) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingComputeHostGroupLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetComputeHostGroupLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ComputeHostGroupLifecycleStateEnum Enum with underlying type: string
type ComputeHostGroupLifecycleStateEnum string

// Set of constants representing the allowable values for ComputeHostGroupLifecycleStateEnum
const (
	ComputeHostGroupLifecycleStateActive  ComputeHostGroupLifecycleStateEnum = "ACTIVE"
	ComputeHostGroupLifecycleStateDeleted ComputeHostGroupLifecycleStateEnum = "DELETED"
)

var mappingComputeHostGroupLifecycleStateEnum = map[string]ComputeHostGroupLifecycleStateEnum{
	"ACTIVE":  ComputeHostGroupLifecycleStateActive,
	"DELETED": ComputeHostGroupLifecycleStateDeleted,
}

var mappingComputeHostGroupLifecycleStateEnumLowerCase = map[string]ComputeHostGroupLifecycleStateEnum{
	"active":  ComputeHostGroupLifecycleStateActive,
	"deleted": ComputeHostGroupLifecycleStateDeleted,
}

// GetComputeHostGroupLifecycleStateEnumValues Enumerates the set of values for ComputeHostGroupLifecycleStateEnum
func GetComputeHostGroupLifecycleStateEnumValues() []ComputeHostGroupLifecycleStateEnum {
	values := make([]ComputeHostGroupLifecycleStateEnum, 0)
	for _, v := range mappingComputeHostGroupLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetComputeHostGroupLifecycleStateEnumStringValues Enumerates the set of values in String for ComputeHostGroupLifecycleStateEnum
func GetComputeHostGroupLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"DELETED",
	}
}

// GetMappingComputeHostGroupLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingComputeHostGroupLifecycleStateEnum(val string) (ComputeHostGroupLifecycleStateEnum, bool) {
	enum, ok := mappingComputeHostGroupLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
