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

// PatchRemoveMemberModels This will specify the list of source model group models that are not to be associated with the cloned model group.
type PatchRemoveMemberModels struct {

	// Array of inference key and model OCID.
	Values []MemberModelDetails `mandatory:"true" json:"values"`
}

func (m PatchRemoveMemberModels) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PatchRemoveMemberModels) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m PatchRemoveMemberModels) MarshalJSON() (buff []byte, e error) {
	type MarshalTypePatchRemoveMemberModels PatchRemoveMemberModels
	s := struct {
		DiscriminatorParam string `json:"operation"`
		MarshalTypePatchRemoveMemberModels
	}{
		"REMOVE",
		(MarshalTypePatchRemoveMemberModels)(m),
	}

	return json.Marshal(&s)
}
