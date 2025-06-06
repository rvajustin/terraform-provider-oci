// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DNS API
//
// API for the DNS service. Use this API to manage DNS zones, records, and other DNS resources.
// For more information, see Overview of the DNS Service (https://docs.oracle.com/iaas/Content/DNS/Concepts/dnszonemanagement.htm).
//

package dns

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SteeringPolicyPriorityRule The representation of SteeringPolicyPriorityRule
type SteeringPolicyPriorityRule struct {

	// A user-defined description of the rule's purpose or behavior.
	Description *string `mandatory:"false" json:"description"`

	// An array of `caseConditions`. A rule may optionally include a sequence of cases defining alternate
	// configurations for how it should behave during processing for any given DNS query. When a rule has
	// no sequence of `cases`, it is always evaluated with the same configuration during processing. When
	// a rule has an empty sequence of `cases`, it is always ignored during processing. When a rule has a
	// non-empty sequence of `cases`, its behavior during processing is configured by the first matching
	// `case` in the sequence. When a rule has no matching cases the rule is ignored. A rule case with no
	// `caseCondition` always matches. A rule case with a `caseCondition` matches only when that expression
	// evaluates to true for the given query.
	Cases []SteeringPolicyPriorityRuleCase `mandatory:"false" json:"cases"`

	// Defines a default set of answer conditions and values that are applied to an answer when
	// `cases` is not defined for the rule or a matching case does not have any matching
	// `answerCondition`s in its `answerData`. `defaultAnswerData` is not applied if `cases` is
	// defined and there are no matching cases. In this scenario, the next rule will be processed.
	DefaultAnswerData []SteeringPolicyPriorityAnswerData `mandatory:"false" json:"defaultAnswerData"`
}

// GetDescription returns Description
func (m SteeringPolicyPriorityRule) GetDescription() *string {
	return m.Description
}

func (m SteeringPolicyPriorityRule) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SteeringPolicyPriorityRule) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m SteeringPolicyPriorityRule) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeSteeringPolicyPriorityRule SteeringPolicyPriorityRule
	s := struct {
		DiscriminatorParam string `json:"ruleType"`
		MarshalTypeSteeringPolicyPriorityRule
	}{
		"PRIORITY",
		(MarshalTypeSteeringPolicyPriorityRule)(m),
	}

	return json.Marshal(&s)
}
