// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Tools
//
// Use the Database Tools API to manage connections, private endpoints, and work requests in the Database Tools service.
//

package databasetools

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DatabaseToolsKeyStoreGenericJdbcSummary The summary of the key store.
type DatabaseToolsKeyStoreGenericJdbcSummary struct {

	// The key store type.
	KeyStoreType KeyStoreTypeGenericJdbcEnum `mandatory:"false" json:"keyStoreType,omitempty"`

	KeyStoreContent DatabaseToolsKeyStoreContentGenericJdbcSummary `mandatory:"false" json:"keyStoreContent"`

	KeyStorePassword DatabaseToolsKeyStorePasswordGenericJdbcSummary `mandatory:"false" json:"keyStorePassword"`
}

func (m DatabaseToolsKeyStoreGenericJdbcSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DatabaseToolsKeyStoreGenericJdbcSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingKeyStoreTypeGenericJdbcEnum(string(m.KeyStoreType)); !ok && m.KeyStoreType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for KeyStoreType: %s. Supported values are: %s.", m.KeyStoreType, strings.Join(GetKeyStoreTypeGenericJdbcEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *DatabaseToolsKeyStoreGenericJdbcSummary) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		KeyStoreType     KeyStoreTypeGenericJdbcEnum                     `json:"keyStoreType"`
		KeyStoreContent  databasetoolskeystorecontentgenericjdbcsummary  `json:"keyStoreContent"`
		KeyStorePassword databasetoolskeystorepasswordgenericjdbcsummary `json:"keyStorePassword"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.KeyStoreType = model.KeyStoreType

	nn, e = model.KeyStoreContent.UnmarshalPolymorphicJSON(model.KeyStoreContent.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.KeyStoreContent = nn.(DatabaseToolsKeyStoreContentGenericJdbcSummary)
	} else {
		m.KeyStoreContent = nil
	}

	nn, e = model.KeyStorePassword.UnmarshalPolymorphicJSON(model.KeyStorePassword.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.KeyStorePassword = nn.(DatabaseToolsKeyStorePasswordGenericJdbcSummary)
	} else {
		m.KeyStorePassword = nil
	}

	return
}
