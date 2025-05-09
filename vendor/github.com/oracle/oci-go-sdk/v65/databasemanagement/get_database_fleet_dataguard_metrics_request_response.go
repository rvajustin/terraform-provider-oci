// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// GetDatabaseFleetDataguardMetricsRequest wrapper for the GetDatabaseFleetDataguardMetrics operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/GetDatabaseFleetDataguardMetrics.go.html to see an example of how to use GetDatabaseFleetDataguardMetricsRequest.
type GetDatabaseFleetDataguardMetricsRequest struct {

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Managed Database Group.
	ManagedDatabaseGroupId *string `mandatory:"false" contributesTo:"query" name:"managedDatabaseGroupId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The filter used to retrieve a specific set of metrics by passing the desired metric names with a comma separator. Note that, by default, the service returns all supported metrics.
	FilterByMetricNames *string `mandatory:"false" contributesTo:"query" name:"filterByMetricNames"`

	// The page token representing the page from where the next set of paginated results
	// are retrieved. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of records returned in the paginated response.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The field to sort information by. Only one sortOrder can be used. The
	// default sort order for `DATABASENAME` is ascending and it is case-sensitive.
	SortBy GetDatabaseFleetDataguardMetricsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The option to sort information in ascending (‘ASC’) or descending (‘DESC’) order. Ascending order is the default order.
	SortOrder GetDatabaseFleetDataguardMetricsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// A list of tag filters to apply.  Only resources with a defined tag matching the value will be returned.
	// Each item in the list has the format "{namespace}.{tagName}.{value}".  All inputs are case-insensitive.
	// Multiple values for the same key (i.e. same namespace and tag name) are interpreted as "OR".
	// Values for different keys (i.e. different namespaces, different tag names, or both) are interpreted as "AND".
	DefinedTagEquals []string `contributesTo:"query" name:"definedTagEquals" collectionFormat:"multi"`

	// A list of tag filters to apply.  Only resources with a freeform tag matching the value will be returned.
	// The key for each tag is "{tagName}.{value}".  All inputs are case-insensitive.
	// Multiple values for the same tag name are interpreted as "OR".  Values for different tag names are interpreted as "AND".
	FreeformTagEquals []string `contributesTo:"query" name:"freeformTagEquals" collectionFormat:"multi"`

	// A list of tag existence filters to apply.  Only resources for which the specified defined tags exist will be returned.
	// Each item in the list has the format "{namespace}.{tagName}.true" (for checking existence of a defined tag)
	// or "{namespace}.true".  All inputs are case-insensitive.
	// Currently, only existence ("true" at the end) is supported. Absence ("false" at the end) is not supported.
	// Multiple values for the same key (i.e. same namespace and tag name) are interpreted as "OR".
	// Values for different keys (i.e. different namespaces, different tag names, or both) are interpreted as "AND".
	DefinedTagExists []string `contributesTo:"query" name:"definedTagExists" collectionFormat:"multi"`

	// A list of tag existence filters to apply.  Only resources for which the specified freeform tags exist the value will be returned.
	// The key for each tag is "{tagName}.true".  All inputs are case-insensitive.
	// Currently, only existence ("true" at the end) is supported. Absence ("false" at the end) is not supported.
	// Multiple values for different tag names are interpreted as "AND".
	FreeformTagExists []string `contributesTo:"query" name:"freeformTagExists" collectionFormat:"multi"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetDatabaseFleetDataguardMetricsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetDatabaseFleetDataguardMetricsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetDatabaseFleetDataguardMetricsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetDatabaseFleetDataguardMetricsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GetDatabaseFleetDataguardMetricsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingGetDatabaseFleetDataguardMetricsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetGetDatabaseFleetDataguardMetricsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingGetDatabaseFleetDataguardMetricsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetGetDatabaseFleetDataguardMetricsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GetDatabaseFleetDataguardMetricsResponse wrapper for the GetDatabaseFleetDataguardMetrics operation
type GetDatabaseFleetDataguardMetricsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DatabaseFleetDataguardMetrics instances
	DatabaseFleetDataguardMetrics `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response GetDatabaseFleetDataguardMetricsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetDatabaseFleetDataguardMetricsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetDatabaseFleetDataguardMetricsSortByEnum Enum with underlying type: string
type GetDatabaseFleetDataguardMetricsSortByEnum string

// Set of constants representing the allowable values for GetDatabaseFleetDataguardMetricsSortByEnum
const (
	GetDatabaseFleetDataguardMetricsSortByDatabasename GetDatabaseFleetDataguardMetricsSortByEnum = "DATABASENAME"
)

var mappingGetDatabaseFleetDataguardMetricsSortByEnum = map[string]GetDatabaseFleetDataguardMetricsSortByEnum{
	"DATABASENAME": GetDatabaseFleetDataguardMetricsSortByDatabasename,
}

var mappingGetDatabaseFleetDataguardMetricsSortByEnumLowerCase = map[string]GetDatabaseFleetDataguardMetricsSortByEnum{
	"databasename": GetDatabaseFleetDataguardMetricsSortByDatabasename,
}

// GetGetDatabaseFleetDataguardMetricsSortByEnumValues Enumerates the set of values for GetDatabaseFleetDataguardMetricsSortByEnum
func GetGetDatabaseFleetDataguardMetricsSortByEnumValues() []GetDatabaseFleetDataguardMetricsSortByEnum {
	values := make([]GetDatabaseFleetDataguardMetricsSortByEnum, 0)
	for _, v := range mappingGetDatabaseFleetDataguardMetricsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetGetDatabaseFleetDataguardMetricsSortByEnumStringValues Enumerates the set of values in String for GetDatabaseFleetDataguardMetricsSortByEnum
func GetGetDatabaseFleetDataguardMetricsSortByEnumStringValues() []string {
	return []string{
		"DATABASENAME",
	}
}

// GetMappingGetDatabaseFleetDataguardMetricsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetDatabaseFleetDataguardMetricsSortByEnum(val string) (GetDatabaseFleetDataguardMetricsSortByEnum, bool) {
	enum, ok := mappingGetDatabaseFleetDataguardMetricsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// GetDatabaseFleetDataguardMetricsSortOrderEnum Enum with underlying type: string
type GetDatabaseFleetDataguardMetricsSortOrderEnum string

// Set of constants representing the allowable values for GetDatabaseFleetDataguardMetricsSortOrderEnum
const (
	GetDatabaseFleetDataguardMetricsSortOrderAsc  GetDatabaseFleetDataguardMetricsSortOrderEnum = "ASC"
	GetDatabaseFleetDataguardMetricsSortOrderDesc GetDatabaseFleetDataguardMetricsSortOrderEnum = "DESC"
)

var mappingGetDatabaseFleetDataguardMetricsSortOrderEnum = map[string]GetDatabaseFleetDataguardMetricsSortOrderEnum{
	"ASC":  GetDatabaseFleetDataguardMetricsSortOrderAsc,
	"DESC": GetDatabaseFleetDataguardMetricsSortOrderDesc,
}

var mappingGetDatabaseFleetDataguardMetricsSortOrderEnumLowerCase = map[string]GetDatabaseFleetDataguardMetricsSortOrderEnum{
	"asc":  GetDatabaseFleetDataguardMetricsSortOrderAsc,
	"desc": GetDatabaseFleetDataguardMetricsSortOrderDesc,
}

// GetGetDatabaseFleetDataguardMetricsSortOrderEnumValues Enumerates the set of values for GetDatabaseFleetDataguardMetricsSortOrderEnum
func GetGetDatabaseFleetDataguardMetricsSortOrderEnumValues() []GetDatabaseFleetDataguardMetricsSortOrderEnum {
	values := make([]GetDatabaseFleetDataguardMetricsSortOrderEnum, 0)
	for _, v := range mappingGetDatabaseFleetDataguardMetricsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetGetDatabaseFleetDataguardMetricsSortOrderEnumStringValues Enumerates the set of values in String for GetDatabaseFleetDataguardMetricsSortOrderEnum
func GetGetDatabaseFleetDataguardMetricsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingGetDatabaseFleetDataguardMetricsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetDatabaseFleetDataguardMetricsSortOrderEnum(val string) (GetDatabaseFleetDataguardMetricsSortOrderEnum, bool) {
	enum, ok := mappingGetDatabaseFleetDataguardMetricsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
