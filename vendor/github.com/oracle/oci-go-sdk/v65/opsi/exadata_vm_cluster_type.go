// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Ops Insights API
//
// Use the Ops Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Ops Insights (https://docs.oracle.com/iaas/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"strings"
)

// ExadataVmClusterTypeEnum Enum with underlying type: string
type ExadataVmClusterTypeEnum string

// Set of constants representing the allowable values for ExadataVmClusterTypeEnum
const (
	ExadataVmClusterTypeVmCluster           ExadataVmClusterTypeEnum = "vmCluster"
	ExadataVmClusterTypeAutonomousVmCluster ExadataVmClusterTypeEnum = "autonomousVmCluster"
)

var mappingExadataVmClusterTypeEnum = map[string]ExadataVmClusterTypeEnum{
	"vmCluster":           ExadataVmClusterTypeVmCluster,
	"autonomousVmCluster": ExadataVmClusterTypeAutonomousVmCluster,
}

var mappingExadataVmClusterTypeEnumLowerCase = map[string]ExadataVmClusterTypeEnum{
	"vmcluster":           ExadataVmClusterTypeVmCluster,
	"autonomousvmcluster": ExadataVmClusterTypeAutonomousVmCluster,
}

// GetExadataVmClusterTypeEnumValues Enumerates the set of values for ExadataVmClusterTypeEnum
func GetExadataVmClusterTypeEnumValues() []ExadataVmClusterTypeEnum {
	values := make([]ExadataVmClusterTypeEnum, 0)
	for _, v := range mappingExadataVmClusterTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetExadataVmClusterTypeEnumStringValues Enumerates the set of values in String for ExadataVmClusterTypeEnum
func GetExadataVmClusterTypeEnumStringValues() []string {
	return []string{
		"vmCluster",
		"autonomousVmCluster",
	}
}

// GetMappingExadataVmClusterTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExadataVmClusterTypeEnum(val string) (ExadataVmClusterTypeEnum, bool) {
	enum, ok := mappingExadataVmClusterTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
