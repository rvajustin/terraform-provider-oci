// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Network Firewall API
//
// Use the Network Firewall API to create network firewalls and configure policies that regulates network traffic in and across VCNs.
//

package networkfirewall

import (
	"strings"
)

// NatTypeEnum Enum with underlying type: string
type NatTypeEnum string

// Set of constants representing the allowable values for NatTypeEnum
const (
	NatTypeNatv4 NatTypeEnum = "NATV4"
)

var mappingNatTypeEnum = map[string]NatTypeEnum{
	"NATV4": NatTypeNatv4,
}

var mappingNatTypeEnumLowerCase = map[string]NatTypeEnum{
	"natv4": NatTypeNatv4,
}

// GetNatTypeEnumValues Enumerates the set of values for NatTypeEnum
func GetNatTypeEnumValues() []NatTypeEnum {
	values := make([]NatTypeEnum, 0)
	for _, v := range mappingNatTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetNatTypeEnumStringValues Enumerates the set of values in String for NatTypeEnum
func GetNatTypeEnumStringValues() []string {
	return []string{
		"NATV4",
	}
}

// GetMappingNatTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingNatTypeEnum(val string) (NatTypeEnum, bool) {
	enum, ok := mappingNatTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
