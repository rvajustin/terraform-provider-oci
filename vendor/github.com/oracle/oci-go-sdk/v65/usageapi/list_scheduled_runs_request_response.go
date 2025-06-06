// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package usageapi

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListScheduledRunsRequest wrapper for the ListScheduledRuns operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/usageapi/ListScheduledRuns.go.html to see an example of how to use ListScheduledRunsRequest.
type ListScheduledRunsRequest struct {

	// The schedule unique ID.
	ScheduleId *string `mandatory:"true" contributesTo:"query" name:"scheduleId"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The page token representing the page at which to start retrieving results.
	// This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The field to sort by. If not specified, the default is timeCreated.
	SortBy ListScheduledRunsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, whether 'asc' or 'desc'.
	SortOrder ListScheduledRunsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListScheduledRunsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListScheduledRunsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListScheduledRunsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListScheduledRunsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListScheduledRunsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListScheduledRunsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListScheduledRunsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListScheduledRunsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListScheduledRunsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListScheduledRunsResponse wrapper for the ListScheduledRuns operation
type ListScheduledRunsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ScheduledRunCollection instances
	ScheduledRunCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListScheduledRunsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListScheduledRunsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListScheduledRunsSortByEnum Enum with underlying type: string
type ListScheduledRunsSortByEnum string

// Set of constants representing the allowable values for ListScheduledRunsSortByEnum
const (
	ListScheduledRunsSortByTimecreated ListScheduledRunsSortByEnum = "timeCreated"
)

var mappingListScheduledRunsSortByEnum = map[string]ListScheduledRunsSortByEnum{
	"timeCreated": ListScheduledRunsSortByTimecreated,
}

var mappingListScheduledRunsSortByEnumLowerCase = map[string]ListScheduledRunsSortByEnum{
	"timecreated": ListScheduledRunsSortByTimecreated,
}

// GetListScheduledRunsSortByEnumValues Enumerates the set of values for ListScheduledRunsSortByEnum
func GetListScheduledRunsSortByEnumValues() []ListScheduledRunsSortByEnum {
	values := make([]ListScheduledRunsSortByEnum, 0)
	for _, v := range mappingListScheduledRunsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListScheduledRunsSortByEnumStringValues Enumerates the set of values in String for ListScheduledRunsSortByEnum
func GetListScheduledRunsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
	}
}

// GetMappingListScheduledRunsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListScheduledRunsSortByEnum(val string) (ListScheduledRunsSortByEnum, bool) {
	enum, ok := mappingListScheduledRunsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListScheduledRunsSortOrderEnum Enum with underlying type: string
type ListScheduledRunsSortOrderEnum string

// Set of constants representing the allowable values for ListScheduledRunsSortOrderEnum
const (
	ListScheduledRunsSortOrderAsc  ListScheduledRunsSortOrderEnum = "ASC"
	ListScheduledRunsSortOrderDesc ListScheduledRunsSortOrderEnum = "DESC"
)

var mappingListScheduledRunsSortOrderEnum = map[string]ListScheduledRunsSortOrderEnum{
	"ASC":  ListScheduledRunsSortOrderAsc,
	"DESC": ListScheduledRunsSortOrderDesc,
}

var mappingListScheduledRunsSortOrderEnumLowerCase = map[string]ListScheduledRunsSortOrderEnum{
	"asc":  ListScheduledRunsSortOrderAsc,
	"desc": ListScheduledRunsSortOrderDesc,
}

// GetListScheduledRunsSortOrderEnumValues Enumerates the set of values for ListScheduledRunsSortOrderEnum
func GetListScheduledRunsSortOrderEnumValues() []ListScheduledRunsSortOrderEnum {
	values := make([]ListScheduledRunsSortOrderEnum, 0)
	for _, v := range mappingListScheduledRunsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListScheduledRunsSortOrderEnumStringValues Enumerates the set of values in String for ListScheduledRunsSortOrderEnum
func GetListScheduledRunsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListScheduledRunsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListScheduledRunsSortOrderEnum(val string) (ListScheduledRunsSortOrderEnum, bool) {
	enum, ok := mappingListScheduledRunsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
