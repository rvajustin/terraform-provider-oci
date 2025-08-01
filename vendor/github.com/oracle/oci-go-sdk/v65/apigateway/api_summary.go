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

// ApiSummary A summary of the API.
type ApiSummary struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the resource.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which the
	// resource is created.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	// Example: `My new resource`
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The time this resource was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time this resource was last updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Locks associated with this resource.
	Locks []ResourceLock `mandatory:"false" json:"locks"`

	// The current state of the API.
	LifecycleState ApiSummaryLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// A message describing the current lifecycleState in more detail. For ACTIVE
	// state it describes if the document has been validated and the possible values are:
	// - 'New' for just updated API Specifications
	// - 'Validating' for a document which is being validated.
	// - 'Valid' the document has been validated without any errors or warnings
	// - 'Warning' the document has been validated and contains warnings
	// - 'Error' the document has been validated and contains errors
	// - 'Failed' the document validation failed
	// - 'Canceled' the document validation was canceled
	// For other states it may provide more details like actionable information.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Type of API Specification file.
	SpecificationType *string `mandatory:"false" json:"specificationType"`

	// Status of each feature available from the API.
	ValidationResults []ApiValidationResult `mandatory:"false" json:"validationResults"`

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

func (m ApiSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ApiSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingApiSummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetApiSummaryLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ApiSummaryLifecycleStateEnum Enum with underlying type: string
type ApiSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for ApiSummaryLifecycleStateEnum
const (
	ApiSummaryLifecycleStateCreating ApiSummaryLifecycleStateEnum = "CREATING"
	ApiSummaryLifecycleStateActive   ApiSummaryLifecycleStateEnum = "ACTIVE"
	ApiSummaryLifecycleStateUpdating ApiSummaryLifecycleStateEnum = "UPDATING"
	ApiSummaryLifecycleStateDeleting ApiSummaryLifecycleStateEnum = "DELETING"
	ApiSummaryLifecycleStateDeleted  ApiSummaryLifecycleStateEnum = "DELETED"
	ApiSummaryLifecycleStateFailed   ApiSummaryLifecycleStateEnum = "FAILED"
)

var mappingApiSummaryLifecycleStateEnum = map[string]ApiSummaryLifecycleStateEnum{
	"CREATING": ApiSummaryLifecycleStateCreating,
	"ACTIVE":   ApiSummaryLifecycleStateActive,
	"UPDATING": ApiSummaryLifecycleStateUpdating,
	"DELETING": ApiSummaryLifecycleStateDeleting,
	"DELETED":  ApiSummaryLifecycleStateDeleted,
	"FAILED":   ApiSummaryLifecycleStateFailed,
}

var mappingApiSummaryLifecycleStateEnumLowerCase = map[string]ApiSummaryLifecycleStateEnum{
	"creating": ApiSummaryLifecycleStateCreating,
	"active":   ApiSummaryLifecycleStateActive,
	"updating": ApiSummaryLifecycleStateUpdating,
	"deleting": ApiSummaryLifecycleStateDeleting,
	"deleted":  ApiSummaryLifecycleStateDeleted,
	"failed":   ApiSummaryLifecycleStateFailed,
}

// GetApiSummaryLifecycleStateEnumValues Enumerates the set of values for ApiSummaryLifecycleStateEnum
func GetApiSummaryLifecycleStateEnumValues() []ApiSummaryLifecycleStateEnum {
	values := make([]ApiSummaryLifecycleStateEnum, 0)
	for _, v := range mappingApiSummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetApiSummaryLifecycleStateEnumStringValues Enumerates the set of values in String for ApiSummaryLifecycleStateEnum
func GetApiSummaryLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingApiSummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingApiSummaryLifecycleStateEnum(val string) (ApiSummaryLifecycleStateEnum, bool) {
	enum, ok := mappingApiSummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
