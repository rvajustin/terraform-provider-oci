// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Autoscaling API
//
// Use the Autoscaling API to dynamically scale compute resources to meet application requirements. For more information about
// autoscaling, see Autoscaling (https://docs.oracle.com/iaas/Content/Compute/Tasks/autoscalinginstancepools.htm). For information about the
// Compute service, see Compute (https://docs.oracle.com/iaas/Content/Compute/home.htm).
//

package autoscaling

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Condition A rule that defines a specific autoscaling action to take (scale in or scale out) and the metric that triggers that action.
type Condition struct {
	Action *Action `mandatory:"true" json:"action"`

	Metric MetricBase `mandatory:"true" json:"metric"`

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// ID of the condition that is assigned after creation.
	Id *string `mandatory:"false" json:"id"`
}

func (m Condition) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Condition) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *Condition) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName *string    `json:"displayName"`
		Id          *string    `json:"id"`
		Action      *Action    `json:"action"`
		Metric      metricbase `json:"metric"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.Id = model.Id

	m.Action = model.Action

	nn, e = model.Metric.UnmarshalPolymorphicJSON(model.Metric.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Metric = nn.(MetricBase)
	} else {
		m.Metric = nil
	}

	return
}
