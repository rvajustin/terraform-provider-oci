// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Resource Manager API
//
// Use the Resource Manager API to automate deployment and operations for all Oracle Cloud Infrastructure resources.
// Using the infrastructure-as-code (IaC) model, the service is based on Terraform, an open source industry standard that lets DevOps engineers develop and deploy their infrastructure anywhere.
// For more information, see
// the Resource Manager documentation (https://docs.oracle.com/iaas/Content/ResourceManager/home.htm).
//

package resourcemanager

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreatePlanRollbackJobOperationDetails Job details that are specific to a plan rollback job. For more information about plan rollback jobs,
// see Creating a Plan Rollback Job (https://docs.oracle.com/iaas/Content/ResourceManager/Tasks/create-job-plan-rollback.htm).
type CreatePlanRollbackJobOperationDetails struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a successful apply job to use for the plan rollback job.
	TargetRollbackJobId *string `mandatory:"true" json:"targetRollbackJobId"`

	// Specifies whether or not to upgrade provider versions.
	// Within the version constraints of your Terraform configuration, use the latest versions available from the source of Terraform providers.
	// For more information about this option, see Dependency Lock File (terraform.io) (https://www.terraform.io/language/files/dependency-lock).
	IsProviderUpgradeRequired *bool `mandatory:"false" json:"isProviderUpgradeRequired"`

	TerraformAdvancedOptions *TerraformAdvancedOptions `mandatory:"false" json:"terraformAdvancedOptions"`
}

// GetIsProviderUpgradeRequired returns IsProviderUpgradeRequired
func (m CreatePlanRollbackJobOperationDetails) GetIsProviderUpgradeRequired() *bool {
	return m.IsProviderUpgradeRequired
}

func (m CreatePlanRollbackJobOperationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreatePlanRollbackJobOperationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreatePlanRollbackJobOperationDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreatePlanRollbackJobOperationDetails CreatePlanRollbackJobOperationDetails
	s := struct {
		DiscriminatorParam string `json:"operation"`
		MarshalTypeCreatePlanRollbackJobOperationDetails
	}{
		"PLAN_ROLLBACK",
		(MarshalTypeCreatePlanRollbackJobOperationDetails)(m),
	}

	return json.Marshal(&s)
}
