// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// API Gateway API
//
// API for the API Gateway service. Use this API to manage gateways, deployments, and related items.
// For more information, see
// Overview of API Gateway (https://docs.oracle.com/iaas/Content/APIGateway/Concepts/apigatewayoverview.htm).
//

package apigateway

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Subscriber A subscriber, which encapsulates a number of clients and usage plans that they are subscribed to.
type Subscriber struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the resource.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which the
	// resource is created.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The clients belonging to this subscriber.
	Clients []Client `mandatory:"true" json:"clients"`

	// An array of OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm)s of usage
	// plan resources.
	UsagePlans []string `mandatory:"true" json:"usagePlans"`

	// The time this resource was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time this resource was last updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The current state of the subscriber.
	LifecycleState SubscriberLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	// Example: `My new resource`
	DisplayName *string `mandatory:"false" json:"displayName"`

	// A message describing the current state in more detail.
	// For example, can be used to provide actionable information for a
	// resource in a Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Locks associated with this resource.
	Locks []ResourceLock `mandatory:"false" json:"locks"`

	// Free-form tags for this resource. Each tag is a simple key-value pair
	// with no predefined name, type, or namespace. For more information, see
	// Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see
	// Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m Subscriber) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Subscriber) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSubscriberLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetSubscriberLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SubscriberLifecycleStateEnum Enum with underlying type: string
type SubscriberLifecycleStateEnum string

// Set of constants representing the allowable values for SubscriberLifecycleStateEnum
const (
	SubscriberLifecycleStateCreating SubscriberLifecycleStateEnum = "CREATING"
	SubscriberLifecycleStateActive   SubscriberLifecycleStateEnum = "ACTIVE"
	SubscriberLifecycleStateUpdating SubscriberLifecycleStateEnum = "UPDATING"
	SubscriberLifecycleStateDeleting SubscriberLifecycleStateEnum = "DELETING"
	SubscriberLifecycleStateDeleted  SubscriberLifecycleStateEnum = "DELETED"
	SubscriberLifecycleStateFailed   SubscriberLifecycleStateEnum = "FAILED"
)

var mappingSubscriberLifecycleStateEnum = map[string]SubscriberLifecycleStateEnum{
	"CREATING": SubscriberLifecycleStateCreating,
	"ACTIVE":   SubscriberLifecycleStateActive,
	"UPDATING": SubscriberLifecycleStateUpdating,
	"DELETING": SubscriberLifecycleStateDeleting,
	"DELETED":  SubscriberLifecycleStateDeleted,
	"FAILED":   SubscriberLifecycleStateFailed,
}

var mappingSubscriberLifecycleStateEnumLowerCase = map[string]SubscriberLifecycleStateEnum{
	"creating": SubscriberLifecycleStateCreating,
	"active":   SubscriberLifecycleStateActive,
	"updating": SubscriberLifecycleStateUpdating,
	"deleting": SubscriberLifecycleStateDeleting,
	"deleted":  SubscriberLifecycleStateDeleted,
	"failed":   SubscriberLifecycleStateFailed,
}

// GetSubscriberLifecycleStateEnumValues Enumerates the set of values for SubscriberLifecycleStateEnum
func GetSubscriberLifecycleStateEnumValues() []SubscriberLifecycleStateEnum {
	values := make([]SubscriberLifecycleStateEnum, 0)
	for _, v := range mappingSubscriberLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetSubscriberLifecycleStateEnumStringValues Enumerates the set of values in String for SubscriberLifecycleStateEnum
func GetSubscriberLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingSubscriberLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSubscriberLifecycleStateEnum(val string) (SubscriberLifecycleStateEnum, bool) {
	enum, ok := mappingSubscriberLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
