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

// Config A configuration item, which has a number of mutually exclusive properties that can be used to set specific
// portions of the configuration.
type Config interface {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the configuration item. An OCID is generated
	// when the item is created.
	GetId() *string

	// The time the resource was created, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339)
	// timestamp format.
	// Example: `2020-02-12T22:47:12.613Z`
	GetTimeCreated() *common.SDKTime

	// The time the resource was updated, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339)
	// timestamp format.
	// Example: `2020-02-13T22:47:12.613Z`
	GetTimeUpdated() *common.SDKTime

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a user.
	GetCreatedBy() *string

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a user.
	GetUpdatedBy() *string

	// For optimistic concurrency control. See `if-match`.
	GetEtag() *string

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}
}

type config struct {
	JsonData     []byte
	Id           *string                           `mandatory:"false" json:"id"`
	TimeCreated  *common.SDKTime                   `mandatory:"false" json:"timeCreated"`
	TimeUpdated  *common.SDKTime                   `mandatory:"false" json:"timeUpdated"`
	CreatedBy    *string                           `mandatory:"false" json:"createdBy"`
	UpdatedBy    *string                           `mandatory:"false" json:"updatedBy"`
	Etag         *string                           `mandatory:"false" json:"etag"`
	FreeformTags map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags  map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	ConfigType   string                            `json:"configType"`
}

// UnmarshalJSON unmarshals json
func (m *config) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerconfig config
	s := struct {
		Model Unmarshalerconfig
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Id = s.Model.Id
	m.TimeCreated = s.Model.TimeCreated
	m.TimeUpdated = s.Model.TimeUpdated
	m.CreatedBy = s.Model.CreatedBy
	m.UpdatedBy = s.Model.UpdatedBy
	m.Etag = s.Model.Etag
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.ConfigType = s.Model.ConfigType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *config) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ConfigType {
	case "AGENT":
		mm := AgentConfig{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OPTIONS":
		mm := Options{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "MACS_APM_EXTENSION":
		mm := MacsApmExtension{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "METRIC_GROUP":
		mm := MetricGroup{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "APDEX":
		mm := ApdexRules{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SPAN_FILTER":
		mm := SpanFilter{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for Config: %s.", m.ConfigType)
		return *m, nil
	}
}

// GetId returns Id
func (m config) GetId() *string {
	return m.Id
}

// GetTimeCreated returns TimeCreated
func (m config) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m config) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetCreatedBy returns CreatedBy
func (m config) GetCreatedBy() *string {
	return m.CreatedBy
}

// GetUpdatedBy returns UpdatedBy
func (m config) GetUpdatedBy() *string {
	return m.UpdatedBy
}

// GetEtag returns Etag
func (m config) GetEtag() *string {
	return m.Etag
}

// GetFreeformTags returns FreeformTags
func (m config) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m config) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m config) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m config) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
