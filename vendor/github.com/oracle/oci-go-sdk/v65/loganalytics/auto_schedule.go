// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AutoSchedule Auto Fixed frequency schedule for a scheduled task.
type AutoSchedule struct {

	// The date and time the scheduled task should execute first time after create or update;
	// thereafter the task will execute as specified in the schedule.
	TimeOfFirstExecution *common.SDKTime `mandatory:"false" json:"timeOfFirstExecution"`

	// Schedule misfire retry policy.
	MisfirePolicy ScheduleMisfirePolicyEnum `mandatory:"false" json:"misfirePolicy,omitempty"`
}

// GetMisfirePolicy returns MisfirePolicy
func (m AutoSchedule) GetMisfirePolicy() ScheduleMisfirePolicyEnum {
	return m.MisfirePolicy
}

// GetTimeOfFirstExecution returns TimeOfFirstExecution
func (m AutoSchedule) GetTimeOfFirstExecution() *common.SDKTime {
	return m.TimeOfFirstExecution
}

func (m AutoSchedule) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AutoSchedule) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingScheduleMisfirePolicyEnum(string(m.MisfirePolicy)); !ok && m.MisfirePolicy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MisfirePolicy: %s. Supported values are: %s.", m.MisfirePolicy, strings.Join(GetScheduleMisfirePolicyEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m AutoSchedule) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeAutoSchedule AutoSchedule
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeAutoSchedule
	}{
		"AUTO",
		(MarshalTypeAutoSchedule)(m),
	}

	return json.Marshal(&s)
}
