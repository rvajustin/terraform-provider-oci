// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Performance Monitoring Configuration API
//
// Use the Application Performance Monitoring Configuration API to query and set Application Performance Monitoring
// configuration. For more information, see Application Performance Monitoring (https://docs.oracle.com/iaas/application-performance-monitoring/index.html).
//

package apmconfig

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateMacsApmExtensionDetails An object that represents APM Agent provisioning via a Management Agent.
type CreateMacsApmExtensionDetails struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Management Agent that will provision the APM Agent.
	ManagementAgentId *string `mandatory:"true" json:"managementAgentId"`

	// Filter patterns used to discover active Java processes for provisioning the APM Agent.
	ProcessFilter []string `mandatory:"true" json:"processFilter"`

	// The OS user that should be used to discover Java processes.
	RunAsUser *string `mandatory:"true" json:"runAsUser"`

	// The name of the service being monitored. This argument enables you to filter by
	// service and view traces and other signals in the APM Explorer user interface.
	ServiceName *string `mandatory:"true" json:"serviceName"`

	// The version of the referenced agent bundle.
	AgentVersion *string `mandatory:"true" json:"agentVersion"`

	// The directory owned by runAsUser.
	AttachInstallDir *string `mandatory:"true" json:"attachInstallDir"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The name by which a configuration entity is displayed to the end user.
	DisplayName *string `mandatory:"false" json:"displayName"`
}

// GetFreeformTags returns FreeformTags
func (m CreateMacsApmExtensionDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m CreateMacsApmExtensionDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m CreateMacsApmExtensionDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateMacsApmExtensionDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateMacsApmExtensionDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateMacsApmExtensionDetails CreateMacsApmExtensionDetails
	s := struct {
		DiscriminatorParam string `json:"configType"`
		MarshalTypeCreateMacsApmExtensionDetails
	}{
		"MACS_APM_EXTENSION",
		(MarshalTypeCreateMacsApmExtensionDetails)(m),
	}

	return json.Marshal(&s)
}
