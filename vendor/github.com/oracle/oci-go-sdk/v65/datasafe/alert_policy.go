// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AlertPolicy An Alert Policy is a set of alerting rules evaluated against a target.
// The alert policy is said to be satisfied when all rules in the policy evaulate to true.
// If there are three rules: rule1,rule2 and rule3, the policy is satisfied if rule1 AND rule2 AND rule3 is True.
type AlertPolicy struct {

	// The OCID of the alert policy.
	Id *string `mandatory:"true" json:"id"`

	// The display name of the alert policy.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID of the compartment that contains the alert policy.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Creation date and time of the alert policy, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Last date and time the alert policy was updated, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The current state of the alert.
	LifecycleState AlertPolicyLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The description of the alert policy.
	Description *string `mandatory:"false" json:"description"`

	// Indicates the Data Safe feature to which the alert policy belongs.
	AlertPolicyType AlertPolicyTypeEnum `mandatory:"false" json:"alertPolicyType,omitempty"`

	// Indicates if the alert policy is user-defined (true) or pre-defined (false).
	IsUserDefined *bool `mandatory:"false" json:"isUserDefined"`

	// Severity level of the alert raised by this policy.
	Severity AlertSeverityEnum `mandatory:"false" json:"severity,omitempty"`

	// Details about the current state of the alert policy.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m AlertPolicy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AlertPolicy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAlertPolicyLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetAlertPolicyLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingAlertPolicyTypeEnum(string(m.AlertPolicyType)); !ok && m.AlertPolicyType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AlertPolicyType: %s. Supported values are: %s.", m.AlertPolicyType, strings.Join(GetAlertPolicyTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAlertSeverityEnum(string(m.Severity)); !ok && m.Severity != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Severity: %s. Supported values are: %s.", m.Severity, strings.Join(GetAlertSeverityEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
