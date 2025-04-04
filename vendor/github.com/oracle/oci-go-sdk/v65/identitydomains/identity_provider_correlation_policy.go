// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity Domains API
//
// Use the Identity Domains API to manage resources within an identity domain, for example, users, dynamic resource groups, groups, and identity providers. For information about managing resources within identity domains, see Identity and Access Management (with identity domains) (https://docs.oracle.com/iaas/Content/Identity/home.htm).
// Use this pattern to construct endpoints for identity domains: `https://<domainURL>/admin/v1/`. See Finding an Identity Domain URL (https://docs.oracle.com/en-us/iaas/Content/Identity/api-getstarted/locate-identity-domain-url.htm) to locate the domain URL you need.
// Use the table of contents and search tool to explore the Identity Domains API.
//

package identitydomains

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// IdentityProviderCorrelationPolicy Correlation policy
// **Added In:** 20.1.3
// **SCIM++ Properties:**
//   - caseExact: true
//   - idcsSearchable: false
//   - multiValued: false
//   - mutability: immutable
//   - required: false
//   - returned: default
//   - type: complex
//   - uniqueness: none
type IdentityProviderCorrelationPolicy struct {

	// A label that indicates the type that this references.
	// **Added In:** 20.1.3
	// **SCIM++ Properties:**
	//  - idcsDefaultValue: Policy
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: immutable
	//  - required: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Type IdentityProviderCorrelationPolicyTypeEnum `mandatory:"true" json:"type"`

	// Policy identifier
	// **Added In:** 20.1.3
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Value *string `mandatory:"true" json:"value"`

	// Policy URI
	// **Added In:** 20.1.3
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: reference
	//  - uniqueness: none
	Ref *string `mandatory:"false" json:"$ref"`

	// Policy display name
	// **Added In:** 20.1.3
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Display *string `mandatory:"false" json:"display"`
}

func (m IdentityProviderCorrelationPolicy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m IdentityProviderCorrelationPolicy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingIdentityProviderCorrelationPolicyTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetIdentityProviderCorrelationPolicyTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// IdentityProviderCorrelationPolicyTypeEnum Enum with underlying type: string
type IdentityProviderCorrelationPolicyTypeEnum string

// Set of constants representing the allowable values for IdentityProviderCorrelationPolicyTypeEnum
const (
	IdentityProviderCorrelationPolicyTypePolicy IdentityProviderCorrelationPolicyTypeEnum = "Policy"
)

var mappingIdentityProviderCorrelationPolicyTypeEnum = map[string]IdentityProviderCorrelationPolicyTypeEnum{
	"Policy": IdentityProviderCorrelationPolicyTypePolicy,
}

var mappingIdentityProviderCorrelationPolicyTypeEnumLowerCase = map[string]IdentityProviderCorrelationPolicyTypeEnum{
	"policy": IdentityProviderCorrelationPolicyTypePolicy,
}

// GetIdentityProviderCorrelationPolicyTypeEnumValues Enumerates the set of values for IdentityProviderCorrelationPolicyTypeEnum
func GetIdentityProviderCorrelationPolicyTypeEnumValues() []IdentityProviderCorrelationPolicyTypeEnum {
	values := make([]IdentityProviderCorrelationPolicyTypeEnum, 0)
	for _, v := range mappingIdentityProviderCorrelationPolicyTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetIdentityProviderCorrelationPolicyTypeEnumStringValues Enumerates the set of values in String for IdentityProviderCorrelationPolicyTypeEnum
func GetIdentityProviderCorrelationPolicyTypeEnumStringValues() []string {
	return []string{
		"Policy",
	}
}

// GetMappingIdentityProviderCorrelationPolicyTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingIdentityProviderCorrelationPolicyTypeEnum(val string) (IdentityProviderCorrelationPolicyTypeEnum, bool) {
	enum, ok := mappingIdentityProviderCorrelationPolicyTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
