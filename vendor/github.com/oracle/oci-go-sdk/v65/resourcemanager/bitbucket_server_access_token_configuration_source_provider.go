// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Resource Manager API
//
// Use the Resource Manager API to automate deployment and operations for all Oracle Cloud Infrastructure resources.
// Using the infrastructure-as-code (IaC) model, the service is based on Terraform, an open source industry standard that lets DevOps engineers develop and deploy their infrastructure anywhere.
// For more information, see
// the Resource Manager documentation (https://docs.oracle.com/iaas/Content/ResourceManager/home.htm).
//

package resourcemanager

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// BitbucketServerAccessTokenConfigurationSourceProvider The properties that define a configuration source provider of the type `BITBUCKET_SERVER_ACCESS_TOKEN`.
// This type corresponds to a configuration source provider in Bitbucket server that is authenticated with a personal access token.
type BitbucketServerAccessTokenConfigurationSourceProvider struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the configuration source provider.
	Id *string `mandatory:"false" json:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment where the configuration source provider is located.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// Human-readable display name for the configuration source provider.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Description of the configuration source provider.
	Description *string `mandatory:"false" json:"description"`

	// The date and time when the configuration source provider was created.
	// Format is defined by RFC3339.
	// Example: `2020-01-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	PrivateServerConfigDetails *PrivateServerConfigDetails `mandatory:"false" json:"privateServerConfigDetails"`

	// Username which is used to authorize the user.
	Username *string `mandatory:"false" json:"username"`

	// Secret ocid which is used to authorize the user.
	SecretId *string `mandatory:"false" json:"secretId"`

	// Free-form tags associated with this resource. Each tag is a key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The Bitbucket server service endpoint.
	// Example: `https://bitbucket.org/`
	ApiEndpoint *string `mandatory:"false" json:"apiEndpoint"`

	// The current lifecycle state of the configuration source provider.
	// For more information about configuration source provider lifecycle states in Resource Manager, see
	// Key Concepts (https://docs.oracle.com/iaas/Content/ResourceManager/Concepts/resourcemanager.htm#concepts__CSPStates).
	LifecycleState ConfigurationSourceProviderLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

// GetId returns Id
func (m BitbucketServerAccessTokenConfigurationSourceProvider) GetId() *string {
	return m.Id
}

// GetCompartmentId returns CompartmentId
func (m BitbucketServerAccessTokenConfigurationSourceProvider) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetDisplayName returns DisplayName
func (m BitbucketServerAccessTokenConfigurationSourceProvider) GetDisplayName() *string {
	return m.DisplayName
}

// GetDescription returns Description
func (m BitbucketServerAccessTokenConfigurationSourceProvider) GetDescription() *string {
	return m.Description
}

// GetTimeCreated returns TimeCreated
func (m BitbucketServerAccessTokenConfigurationSourceProvider) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetLifecycleState returns LifecycleState
func (m BitbucketServerAccessTokenConfigurationSourceProvider) GetLifecycleState() ConfigurationSourceProviderLifecycleStateEnum {
	return m.LifecycleState
}

// GetPrivateServerConfigDetails returns PrivateServerConfigDetails
func (m BitbucketServerAccessTokenConfigurationSourceProvider) GetPrivateServerConfigDetails() *PrivateServerConfigDetails {
	return m.PrivateServerConfigDetails
}

// GetUsername returns Username
func (m BitbucketServerAccessTokenConfigurationSourceProvider) GetUsername() *string {
	return m.Username
}

// GetSecretId returns SecretId
func (m BitbucketServerAccessTokenConfigurationSourceProvider) GetSecretId() *string {
	return m.SecretId
}

// GetFreeformTags returns FreeformTags
func (m BitbucketServerAccessTokenConfigurationSourceProvider) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m BitbucketServerAccessTokenConfigurationSourceProvider) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m BitbucketServerAccessTokenConfigurationSourceProvider) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BitbucketServerAccessTokenConfigurationSourceProvider) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingConfigurationSourceProviderLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetConfigurationSourceProviderLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m BitbucketServerAccessTokenConfigurationSourceProvider) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeBitbucketServerAccessTokenConfigurationSourceProvider BitbucketServerAccessTokenConfigurationSourceProvider
	s := struct {
		DiscriminatorParam string `json:"configSourceProviderType"`
		MarshalTypeBitbucketServerAccessTokenConfigurationSourceProvider
	}{
		"BITBUCKET_SERVER_ACCESS_TOKEN",
		(MarshalTypeBitbucketServerAccessTokenConfigurationSourceProvider)(m),
	}

	return json.Marshal(&s)
}
