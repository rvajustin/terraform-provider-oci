// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps API to create DevOps projects, configure code repositories,  add artifacts to deploy, build and test software applications, configure  target deployment environments, and deploy software applications.  For more information, see DevOps (https://docs.oracle.com/iaas/Content/devops/using/home.htm).
//

package devops

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateComputeInstanceGroupDeployEnvironmentDetails Specifies the Compute instance group environment.
type CreateComputeInstanceGroupDeployEnvironmentDetails struct {

	// The OCID of a project.
	ProjectId *string `mandatory:"true" json:"projectId"`

	ComputeInstanceGroupSelectors *ComputeInstanceGroupSelectorCollection `mandatory:"true" json:"computeInstanceGroupSelectors"`

	// Optional description about the deployment environment.
	Description *string `mandatory:"false" json:"description"`

	// Deployment environment display name. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

// GetDescription returns Description
func (m CreateComputeInstanceGroupDeployEnvironmentDetails) GetDescription() *string {
	return m.Description
}

// GetDisplayName returns DisplayName
func (m CreateComputeInstanceGroupDeployEnvironmentDetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetProjectId returns ProjectId
func (m CreateComputeInstanceGroupDeployEnvironmentDetails) GetProjectId() *string {
	return m.ProjectId
}

// GetFreeformTags returns FreeformTags
func (m CreateComputeInstanceGroupDeployEnvironmentDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m CreateComputeInstanceGroupDeployEnvironmentDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m CreateComputeInstanceGroupDeployEnvironmentDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateComputeInstanceGroupDeployEnvironmentDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateComputeInstanceGroupDeployEnvironmentDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateComputeInstanceGroupDeployEnvironmentDetails CreateComputeInstanceGroupDeployEnvironmentDetails
	s := struct {
		DiscriminatorParam string `json:"deployEnvironmentType"`
		MarshalTypeCreateComputeInstanceGroupDeployEnvironmentDetails
	}{
		"COMPUTE_INSTANCE_GROUP",
		(MarshalTypeCreateComputeInstanceGroupDeployEnvironmentDetails)(m),
	}

	return json.Marshal(&s)
}
