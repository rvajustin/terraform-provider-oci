// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package managementdashboard

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListOobManagementSavedSearchesRequest wrapper for the ListOobManagementSavedSearches operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/managementdashboard/ListOobManagementSavedSearches.go.html to see an example of how to use ListOobManagementSavedSearchesRequest.
type ListOobManagementSavedSearchesRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page on which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListOobManagementSavedSearchesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is the default.
	SortBy ListOobManagementSavedSearchesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListOobManagementSavedSearchesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListOobManagementSavedSearchesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListOobManagementSavedSearchesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListOobManagementSavedSearchesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListOobManagementSavedSearchesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListOobManagementSavedSearchesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListOobManagementSavedSearchesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOobManagementSavedSearchesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListOobManagementSavedSearchesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListOobManagementSavedSearchesResponse wrapper for the ListOobManagementSavedSearches operation
type ListOobManagementSavedSearchesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ManagementSavedSearchCollection instances
	ManagementSavedSearchCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListOobManagementSavedSearchesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListOobManagementSavedSearchesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListOobManagementSavedSearchesSortOrderEnum Enum with underlying type: string
type ListOobManagementSavedSearchesSortOrderEnum string

// Set of constants representing the allowable values for ListOobManagementSavedSearchesSortOrderEnum
const (
	ListOobManagementSavedSearchesSortOrderAsc  ListOobManagementSavedSearchesSortOrderEnum = "ASC"
	ListOobManagementSavedSearchesSortOrderDesc ListOobManagementSavedSearchesSortOrderEnum = "DESC"
)

var mappingListOobManagementSavedSearchesSortOrderEnum = map[string]ListOobManagementSavedSearchesSortOrderEnum{
	"ASC":  ListOobManagementSavedSearchesSortOrderAsc,
	"DESC": ListOobManagementSavedSearchesSortOrderDesc,
}

var mappingListOobManagementSavedSearchesSortOrderEnumLowerCase = map[string]ListOobManagementSavedSearchesSortOrderEnum{
	"asc":  ListOobManagementSavedSearchesSortOrderAsc,
	"desc": ListOobManagementSavedSearchesSortOrderDesc,
}

// GetListOobManagementSavedSearchesSortOrderEnumValues Enumerates the set of values for ListOobManagementSavedSearchesSortOrderEnum
func GetListOobManagementSavedSearchesSortOrderEnumValues() []ListOobManagementSavedSearchesSortOrderEnum {
	values := make([]ListOobManagementSavedSearchesSortOrderEnum, 0)
	for _, v := range mappingListOobManagementSavedSearchesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListOobManagementSavedSearchesSortOrderEnumStringValues Enumerates the set of values in String for ListOobManagementSavedSearchesSortOrderEnum
func GetListOobManagementSavedSearchesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListOobManagementSavedSearchesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOobManagementSavedSearchesSortOrderEnum(val string) (ListOobManagementSavedSearchesSortOrderEnum, bool) {
	enum, ok := mappingListOobManagementSavedSearchesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListOobManagementSavedSearchesSortByEnum Enum with underlying type: string
type ListOobManagementSavedSearchesSortByEnum string

// Set of constants representing the allowable values for ListOobManagementSavedSearchesSortByEnum
const (
	ListOobManagementSavedSearchesSortByTimecreated ListOobManagementSavedSearchesSortByEnum = "timeCreated"
	ListOobManagementSavedSearchesSortByDisplayname ListOobManagementSavedSearchesSortByEnum = "displayName"
)

var mappingListOobManagementSavedSearchesSortByEnum = map[string]ListOobManagementSavedSearchesSortByEnum{
	"timeCreated": ListOobManagementSavedSearchesSortByTimecreated,
	"displayName": ListOobManagementSavedSearchesSortByDisplayname,
}

var mappingListOobManagementSavedSearchesSortByEnumLowerCase = map[string]ListOobManagementSavedSearchesSortByEnum{
	"timecreated": ListOobManagementSavedSearchesSortByTimecreated,
	"displayname": ListOobManagementSavedSearchesSortByDisplayname,
}

// GetListOobManagementSavedSearchesSortByEnumValues Enumerates the set of values for ListOobManagementSavedSearchesSortByEnum
func GetListOobManagementSavedSearchesSortByEnumValues() []ListOobManagementSavedSearchesSortByEnum {
	values := make([]ListOobManagementSavedSearchesSortByEnum, 0)
	for _, v := range mappingListOobManagementSavedSearchesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListOobManagementSavedSearchesSortByEnumStringValues Enumerates the set of values in String for ListOobManagementSavedSearchesSortByEnum
func GetListOobManagementSavedSearchesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListOobManagementSavedSearchesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOobManagementSavedSearchesSortByEnum(val string) (ListOobManagementSavedSearchesSortByEnum, bool) {
	enum, ok := mappingListOobManagementSavedSearchesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
