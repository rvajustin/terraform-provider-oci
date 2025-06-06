// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// InputParameter Input Parameters for the Task
type InputParameter struct {

	// stepName for which the input parameters are provided
	StepName *string `mandatory:"true" json:"stepName"`

	// Arguments for the Task
	Arguments []TaskArgument `mandatory:"false" json:"arguments"`
}

func (m InputParameter) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InputParameter) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *InputParameter) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Arguments []taskargument `json:"arguments"`
		StepName  *string        `json:"stepName"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Arguments = make([]TaskArgument, len(model.Arguments))
	for i, n := range model.Arguments {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Arguments[i] = nn.(TaskArgument)
		} else {
			m.Arguments[i] = nil
		}
	}
	m.StepName = model.StepName

	return
}
