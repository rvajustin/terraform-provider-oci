// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// MySqlApplyErrorWorker If the replica is multi-threaded, error from worker threads. Otherwise, error from the applier thread.
type MySqlApplyErrorWorker struct {

	// The error number of the most recent error that caused the worker thread to stop.
	LastErrorNumber *int `mandatory:"true" json:"lastErrorNumber"`

	// The error message of the most recent error that caused the worker thread to stop.
	LastErrorMessage *string `mandatory:"true" json:"lastErrorMessage"`

	// The timestamp when the most recent worker error occurred.
	TimeLastError *common.SDKTime `mandatory:"true" json:"timeLastError"`
}

func (m MySqlApplyErrorWorker) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MySqlApplyErrorWorker) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
