// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Full Stack Disaster Recovery API
//
// Use the Full Stack Disaster Recovery (DR) API to manage disaster recovery for business applications.
// Full Stack DR is an OCI disaster recovery orchestration and management service that provides comprehensive disaster
// recovery capabilities for all layers of an application stack, including infrastructure, middleware, database,
// and application.
//

package disasterrecovery

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OkeBackupLocation The details for object storage backup location of an OKE Cluster
type OkeBackupLocation struct {

	// The namespace in object storage backup location(Note - this is usually the tenancy name).
	// Example: `myocitenancy`
	Namespace *string `mandatory:"true" json:"namespace"`

	// The bucket name inside the object storage namespace.
	// Example: `operation_logs`
	Bucket *string `mandatory:"true" json:"bucket"`

	// The object name inside the object storage bucket.
	// Example: `switchover_plan_executions`
	Object *string `mandatory:"false" json:"object"`
}

func (m OkeBackupLocation) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OkeBackupLocation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
