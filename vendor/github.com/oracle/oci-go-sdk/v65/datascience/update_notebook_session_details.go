// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Science API
//
// Use the Data Science API to organize your data science work, access data and computing resources, and build, train, deploy and manage models and model deployments. For more information, see Data Science (https://docs.oracle.com/iaas/data-science/using/data-science.htm).
//

package datascience

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateNotebookSessionDetails Details for updating a notebook session. `notebookSessionConfigurationDetails` can only be updated while the notebook session is in the `INACTIVE` state.
// Changes to the `notebookSessionConfigurationDetails` take effect the next time the `ActivateNotebookSession` action is invoked on the notebook session resource.
type UpdateNotebookSessionDetails struct {

	// A user-friendly display name for the resource. It does not have to be unique and can be modified. Avoid entering confidential information.
	// Example: `My NotebookSession`
	DisplayName *string `mandatory:"false" json:"displayName"`

	NotebookSessionConfigurationDetails *NotebookSessionConfigurationDetails `mandatory:"false" json:"notebookSessionConfigurationDetails"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. See Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	NotebookSessionRuntimeConfigDetails *NotebookSessionRuntimeConfigDetails `mandatory:"false" json:"notebookSessionRuntimeConfigDetails"`

	// Collection of NotebookSessionStorageMountConfigurationDetails.
	NotebookSessionStorageMountConfigurationDetailsList []StorageMountConfigurationDetails `mandatory:"false" json:"notebookSessionStorageMountConfigurationDetailsList"`
}

func (m UpdateNotebookSessionDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateNotebookSessionDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *UpdateNotebookSessionDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName                                         *string                              `json:"displayName"`
		NotebookSessionConfigurationDetails                 *NotebookSessionConfigurationDetails `json:"notebookSessionConfigurationDetails"`
		FreeformTags                                        map[string]string                    `json:"freeformTags"`
		DefinedTags                                         map[string]map[string]interface{}    `json:"definedTags"`
		NotebookSessionRuntimeConfigDetails                 *NotebookSessionRuntimeConfigDetails `json:"notebookSessionRuntimeConfigDetails"`
		NotebookSessionStorageMountConfigurationDetailsList []storagemountconfigurationdetails   `json:"notebookSessionStorageMountConfigurationDetailsList"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.NotebookSessionConfigurationDetails = model.NotebookSessionConfigurationDetails

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.NotebookSessionRuntimeConfigDetails = model.NotebookSessionRuntimeConfigDetails

	m.NotebookSessionStorageMountConfigurationDetailsList = make([]StorageMountConfigurationDetails, len(model.NotebookSessionStorageMountConfigurationDetailsList))
	for i, n := range model.NotebookSessionStorageMountConfigurationDetailsList {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.NotebookSessionStorageMountConfigurationDetailsList[i] = nn.(StorageMountConfigurationDetails)
		} else {
			m.NotebookSessionStorageMountConfigurationDetailsList[i] = nil
		}
	}
	return
}
