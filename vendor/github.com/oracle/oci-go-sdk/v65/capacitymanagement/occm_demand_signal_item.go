// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OCI Control Center Capacity Management API
//
// OCI Control Center (OCC) Capacity Management enables you to manage capacity requests in realms where OCI Control Center Capacity Management is available. For more information, see OCI Control Center (https://docs.oracle.com/iaas/Content/control-center/home.htm).
//

package capacitymanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OccmDemandSignalItem An occm demand signal item is a resource that is used to communicate the forecasting need for a particular resource with OCI. It's a sub-resource and need to be grouped inside a demand signal.
type OccmDemandSignalItem struct {

	// The OCID of the demand signal item.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the tenancy from which the demand signal item was created.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the demand signal under which this item will be grouped.
	DemandSignalId *string `mandatory:"true" json:"demandSignalId"`

	// The name of the OCI service in consideration for demand signal submission. For example: COMPUTE, NETWORK, GPU etc.
	DemandSignalNamespace DemandSignalNamespaceEnum `mandatory:"true" json:"demandSignalNamespace"`

	// The OCID of the corresponding demand signal catalog resource.
	DemandSignalCatalogResourceId *string `mandatory:"true" json:"demandSignalCatalogResourceId"`

	// The type of request (DEMAND or RETURN) made against a particular demand signal item.
	RequestType OccmDemandSignalItemRequestTypeEnum `mandatory:"true" json:"requestType"`

	// The name of the OCI resource that you want to request.
	ResourceName *string `mandatory:"true" json:"resourceName"`

	// The name of region for which you want to request the OCI resource.
	Region *string `mandatory:"true" json:"region"`

	// The quantity of the resource that you want to demand from OCI.
	DemandQuantity *int64 `mandatory:"true" json:"demandQuantity"`

	// the date before which you would ideally like the OCI resource to be delivered to you.
	TimeNeededBefore *common.SDKTime `mandatory:"true" json:"timeNeededBefore"`

	// A map of various properties associated with the OCI resource.
	ResourceProperties map[string]string `mandatory:"true" json:"resourceProperties"`

	// The current lifecycle state of the resource.
	LifecycleState OccmDemandSignalItemLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The name of the availability domain for which you want to request the OCI resource.
	AvailabilityDomain *string `mandatory:"false" json:"availabilityDomain"`

	// The OCID of the tenancy for which you want to request the OCI resource for. This is an optional parameter.
	TargetCompartmentId *string `mandatory:"false" json:"targetCompartmentId"`

	// This field will serve as notes section for you. You can use this section to convey a message to OCI regarding your resource request.
	// NOTE: The previous value gets overwritten with the new one for this once updated.
	Notes *string `mandatory:"false" json:"notes"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m OccmDemandSignalItem) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OccmDemandSignalItem) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDemandSignalNamespaceEnum(string(m.DemandSignalNamespace)); !ok && m.DemandSignalNamespace != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DemandSignalNamespace: %s. Supported values are: %s.", m.DemandSignalNamespace, strings.Join(GetDemandSignalNamespaceEnumStringValues(), ",")))
	}
	if _, ok := GetMappingOccmDemandSignalItemRequestTypeEnum(string(m.RequestType)); !ok && m.RequestType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RequestType: %s. Supported values are: %s.", m.RequestType, strings.Join(GetOccmDemandSignalItemRequestTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingOccmDemandSignalItemLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetOccmDemandSignalItemLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OccmDemandSignalItemLifecycleStateEnum Enum with underlying type: string
type OccmDemandSignalItemLifecycleStateEnum string

// Set of constants representing the allowable values for OccmDemandSignalItemLifecycleStateEnum
const (
	OccmDemandSignalItemLifecycleStateCreating OccmDemandSignalItemLifecycleStateEnum = "CREATING"
	OccmDemandSignalItemLifecycleStateUpdating OccmDemandSignalItemLifecycleStateEnum = "UPDATING"
	OccmDemandSignalItemLifecycleStateActive   OccmDemandSignalItemLifecycleStateEnum = "ACTIVE"
	OccmDemandSignalItemLifecycleStateDeleting OccmDemandSignalItemLifecycleStateEnum = "DELETING"
	OccmDemandSignalItemLifecycleStateDeleted  OccmDemandSignalItemLifecycleStateEnum = "DELETED"
	OccmDemandSignalItemLifecycleStateFailed   OccmDemandSignalItemLifecycleStateEnum = "FAILED"
)

var mappingOccmDemandSignalItemLifecycleStateEnum = map[string]OccmDemandSignalItemLifecycleStateEnum{
	"CREATING": OccmDemandSignalItemLifecycleStateCreating,
	"UPDATING": OccmDemandSignalItemLifecycleStateUpdating,
	"ACTIVE":   OccmDemandSignalItemLifecycleStateActive,
	"DELETING": OccmDemandSignalItemLifecycleStateDeleting,
	"DELETED":  OccmDemandSignalItemLifecycleStateDeleted,
	"FAILED":   OccmDemandSignalItemLifecycleStateFailed,
}

var mappingOccmDemandSignalItemLifecycleStateEnumLowerCase = map[string]OccmDemandSignalItemLifecycleStateEnum{
	"creating": OccmDemandSignalItemLifecycleStateCreating,
	"updating": OccmDemandSignalItemLifecycleStateUpdating,
	"active":   OccmDemandSignalItemLifecycleStateActive,
	"deleting": OccmDemandSignalItemLifecycleStateDeleting,
	"deleted":  OccmDemandSignalItemLifecycleStateDeleted,
	"failed":   OccmDemandSignalItemLifecycleStateFailed,
}

// GetOccmDemandSignalItemLifecycleStateEnumValues Enumerates the set of values for OccmDemandSignalItemLifecycleStateEnum
func GetOccmDemandSignalItemLifecycleStateEnumValues() []OccmDemandSignalItemLifecycleStateEnum {
	values := make([]OccmDemandSignalItemLifecycleStateEnum, 0)
	for _, v := range mappingOccmDemandSignalItemLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetOccmDemandSignalItemLifecycleStateEnumStringValues Enumerates the set of values in String for OccmDemandSignalItemLifecycleStateEnum
func GetOccmDemandSignalItemLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingOccmDemandSignalItemLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOccmDemandSignalItemLifecycleStateEnum(val string) (OccmDemandSignalItemLifecycleStateEnum, bool) {
	enum, ok := mappingOccmDemandSignalItemLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
