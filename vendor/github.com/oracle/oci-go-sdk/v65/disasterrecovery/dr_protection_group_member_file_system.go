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
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DrProtectionGroupMemberFileSystem The properties for a file system member of a DR protection group.
type DrProtectionGroupMemberFileSystem struct {

	// The OCID of the member.
	// Example: `ocid1.instance.oc1..uniqueID`
	MemberId *string `mandatory:"true" json:"memberId"`

	// The availability domain of the destination mount target.
	//
	// Example: `BBTh:region-AD`
	DestinationAvailabilityDomain *string `mandatory:"false" json:"destinationAvailabilityDomain"`

	// A list of mappings between the primary region file system export and destination region mount target.
	ExportMappings []FileSystemExportMapping `mandatory:"false" json:"exportMappings"`

	DestinationEncryptionKey *VaultAndEncryptionKey `mandatory:"false" json:"destinationEncryptionKey"`

	// The OCID of the snapshot policy to use in the destination region. This policy will be attached to the file system after it moves to the destination region.
	// Example: `ocid1.filesystemsnapshotpolicy.oc1..uniqueID`
	DestinationSnapshotPolicyId *string `mandatory:"false" json:"destinationSnapshotPolicyId"`
}

// GetMemberId returns MemberId
func (m DrProtectionGroupMemberFileSystem) GetMemberId() *string {
	return m.MemberId
}

func (m DrProtectionGroupMemberFileSystem) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DrProtectionGroupMemberFileSystem) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DrProtectionGroupMemberFileSystem) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDrProtectionGroupMemberFileSystem DrProtectionGroupMemberFileSystem
	s := struct {
		DiscriminatorParam string `json:"memberType"`
		MarshalTypeDrProtectionGroupMemberFileSystem
	}{
		"FILE_SYSTEM",
		(MarshalTypeDrProtectionGroupMemberFileSystem)(m),
	}

	return json.Marshal(&s)
}
