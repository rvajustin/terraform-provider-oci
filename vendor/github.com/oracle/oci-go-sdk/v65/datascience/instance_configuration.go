// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Science API
//
// Use the Data Science API to organize your data science work, access data and computing resources, and build, train, deploy and manage models and model deployments. For more information, see Data Science (https://docs.oracle.com/iaas/data-science/using/data-science.htm).
//

package datascience

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// InstanceConfiguration The model deployment instance configuration.
type InstanceConfiguration struct {

	// The shape used to launch the model deployment instances.
	InstanceShapeName *string `mandatory:"true" json:"instanceShapeName"`

	ModelDeploymentInstanceShapeConfigDetails *ModelDeploymentInstanceShapeConfigDetails `mandatory:"false" json:"modelDeploymentInstanceShapeConfigDetails"`

	// A model deployment instance is provided with a VNIC for network access.  This specifies the OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet to create a VNIC in.  The subnet should be in a VCN with a NAT/SGW gateway for egress.
	SubnetId *string `mandatory:"false" json:"subnetId"`

	// The OCID of a Data Science private endpoint.
	PrivateEndpointId *string `mandatory:"false" json:"privateEndpointId"`
}

func (m InstanceConfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InstanceConfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
