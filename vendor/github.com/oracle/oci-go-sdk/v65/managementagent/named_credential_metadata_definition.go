// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Management Agent API
//
// Use the Management Agent API to manage your infrastructure's management agents, including their plugins and install keys.
// For more information, see Management Agent (https://docs.oracle.com/iaas/management-agents/index.html).
//

package managementagent

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// NamedCredentialMetadataDefinition A named credential metadata definition
type NamedCredentialMetadataDefinition struct {

	// The type of the Named Credential.
	Type *string `mandatory:"true" json:"type"`

	// Display name for this type of Named Credential
	DisplayName *string `mandatory:"true" json:"displayName"`

	// This Named Credential type is supported on management agents at this version or above.
	MinimumAgentVersion *string `mandatory:"true" json:"minimumAgentVersion"`

	// The property definitions for this named credential metadata
	Properties []NamedCredentialFieldDefinition `mandatory:"true" json:"properties"`
}

func (m NamedCredentialMetadataDefinition) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m NamedCredentialMetadataDefinition) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
