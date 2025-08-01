// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Science API
//
// Use the Data Science API to organize your data science work, access data and computing resources, and build, train, deploy and manage models and model deployments. For more information, see Data Science (https://docs.oracle.com/iaas/data-science/using/data-science.htm).
//

package datascience

import (
	"strings"
)

// PatchInstructionTypeEnum Enum with underlying type: string
type PatchInstructionTypeEnum string

// Set of constants representing the allowable values for PatchInstructionTypeEnum
const (
	PatchInstructionTypeInsert PatchInstructionTypeEnum = "INSERT"
	PatchInstructionTypeRemove PatchInstructionTypeEnum = "REMOVE"
)

var mappingPatchInstructionTypeEnum = map[string]PatchInstructionTypeEnum{
	"INSERT": PatchInstructionTypeInsert,
	"REMOVE": PatchInstructionTypeRemove,
}

var mappingPatchInstructionTypeEnumLowerCase = map[string]PatchInstructionTypeEnum{
	"insert": PatchInstructionTypeInsert,
	"remove": PatchInstructionTypeRemove,
}

// GetPatchInstructionTypeEnumValues Enumerates the set of values for PatchInstructionTypeEnum
func GetPatchInstructionTypeEnumValues() []PatchInstructionTypeEnum {
	values := make([]PatchInstructionTypeEnum, 0)
	for _, v := range mappingPatchInstructionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetPatchInstructionTypeEnumStringValues Enumerates the set of values in String for PatchInstructionTypeEnum
func GetPatchInstructionTypeEnumStringValues() []string {
	return []string{
		"INSERT",
		"REMOVE",
	}
}

// GetMappingPatchInstructionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPatchInstructionTypeEnum(val string) (PatchInstructionTypeEnum, bool) {
	enum, ok := mappingPatchInstructionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
