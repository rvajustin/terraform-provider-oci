// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Delegate Access Control API
//
// Oracle Delegate Access Control allows ExaCC and ExaCS customers to delegate management of their Exadata resources operators outside their tenancies.
// With Delegate Access Control, Support Providers can deliver managed services using comprehensive and robust tooling built on the OCI platform.
// Customers maintain control over who has access to the delegated resources in their tenancy and what actions can be taken.
// Enterprises managing resources across multiple tenants can use Delegate Access Control to streamline management tasks.
// Using logging service, customers can view a near real-time audit report of all actions performed by a Service Provider operator.
//

package delegateaccesscontrol

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateDelegationControlDetails Information about the new Delegation Control.
type UpdateDelegationControlDetails struct {

	// Name of the Delegation Control. The name does not need to be unique.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Description of the Delegation Control.
	Description *string `mandatory:"false" json:"description"`

	// number of approvals required.
	NumApprovalsRequired *int `mandatory:"false" json:"numApprovalsRequired"`

	// List of Delegation Subscription OCID that are allowed for this Delegation Control. The allowed subscriptions will determine the available Service Provider Actions. Only support operators for the allowed subscriptions are allowed to create Delegated Resource Access Request.
	DelegationSubscriptionIds []string `mandatory:"false" json:"delegationSubscriptionIds"`

	// Set to true to allow all Delegated Resource Access Request to be approved automatically during maintenance.
	IsAutoApproveDuringMaintenance *bool `mandatory:"false" json:"isAutoApproveDuringMaintenance"`

	// The OCID of the selected resources that this Delegation Control is applicable to.
	ResourceIds []string `mandatory:"false" json:"resourceIds"`

	// List of pre-approved Service Provider Action names. The list of pre-defined Service Provider Actions can be obtained from the ListServiceProviderActions API. Delegated Resource Access Requests associated with a resource governed by this Delegation Control will be
	// automatically approved if the Delegated Resource Access Request only contain Service Provider Actions in the pre-approved list.
	PreApprovedServiceProviderActionNames []string `mandatory:"false" json:"preApprovedServiceProviderActionNames"`

	// The OCID of the OCI Notification topic to publish messages related to this Delegation Control.
	NotificationTopicId *string `mandatory:"false" json:"notificationTopicId"`

	// The format of the OCI Notification messages for this Delegation Control.
	NotificationMessageFormat DelegationControlNotificationMessageFormatEnum `mandatory:"false" json:"notificationMessageFormat,omitempty"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateDelegationControlDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateDelegationControlDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDelegationControlNotificationMessageFormatEnum(string(m.NotificationMessageFormat)); !ok && m.NotificationMessageFormat != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for NotificationMessageFormat: %s. Supported values are: %s.", m.NotificationMessageFormat, strings.Join(GetDelegationControlNotificationMessageFormatEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
