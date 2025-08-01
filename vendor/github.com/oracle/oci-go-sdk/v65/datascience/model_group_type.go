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

// ModelGroupTypeEnum Enum with underlying type: string
type ModelGroupTypeEnum string

// Set of constants representing the allowable values for ModelGroupTypeEnum
const (
	ModelGroupTypeHomogeneous   ModelGroupTypeEnum = "HOMOGENEOUS"
	ModelGroupTypeHeterogeneous ModelGroupTypeEnum = "HETEROGENEOUS"
	ModelGroupTypeStacked       ModelGroupTypeEnum = "STACKED"
)

var mappingModelGroupTypeEnum = map[string]ModelGroupTypeEnum{
	"HOMOGENEOUS":   ModelGroupTypeHomogeneous,
	"HETEROGENEOUS": ModelGroupTypeHeterogeneous,
	"STACKED":       ModelGroupTypeStacked,
}

var mappingModelGroupTypeEnumLowerCase = map[string]ModelGroupTypeEnum{
	"homogeneous":   ModelGroupTypeHomogeneous,
	"heterogeneous": ModelGroupTypeHeterogeneous,
	"stacked":       ModelGroupTypeStacked,
}

// GetModelGroupTypeEnumValues Enumerates the set of values for ModelGroupTypeEnum
func GetModelGroupTypeEnumValues() []ModelGroupTypeEnum {
	values := make([]ModelGroupTypeEnum, 0)
	for _, v := range mappingModelGroupTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetModelGroupTypeEnumStringValues Enumerates the set of values in String for ModelGroupTypeEnum
func GetModelGroupTypeEnumStringValues() []string {
	return []string{
		"HOMOGENEOUS",
		"HETEROGENEOUS",
		"STACKED",
	}
}

// GetMappingModelGroupTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingModelGroupTypeEnum(val string) (ModelGroupTypeEnum, bool) {
	enum, ok := mappingModelGroupTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
