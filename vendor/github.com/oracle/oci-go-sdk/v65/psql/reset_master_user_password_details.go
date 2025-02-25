// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// PGSQL Control Plane API
//
// Use the OCI Database with PostgreSQL API to manage resources such as database systems, database nodes, backups, and configurations.
// For information, see the user guide documentation for the service (https://docs.oracle.com/iaas/Content/postgresql/home.htm).
//

package psql

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ResetMasterUserPasswordDetails Password detail that will be used to reset the database system's master user.
// These details are not visible on any subsequent operation, such as GET /dbSystems/{dbSystemId}.
type ResetMasterUserPasswordDetails struct {
	PasswordDetails PasswordDetails `mandatory:"true" json:"passwordDetails"`
}

func (m ResetMasterUserPasswordDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ResetMasterUserPasswordDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *ResetMasterUserPasswordDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		PasswordDetails passworddetails `json:"passwordDetails"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.PasswordDetails.UnmarshalPolymorphicJSON(model.PasswordDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.PasswordDetails = nn.(PasswordDetails)
	} else {
		m.PasswordDetails = nil
	}

	return
}
