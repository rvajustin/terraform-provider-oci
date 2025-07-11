// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package wlms

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListWlsDomainServerInstalledPatchesRequest wrapper for the ListWlsDomainServerInstalledPatches operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/wlms/ListWlsDomainServerInstalledPatches.go.html to see an example of how to use ListWlsDomainServerInstalledPatchesRequest.
type ListWlsDomainServerInstalledPatchesRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the WebLogic domain.
	WlsDomainId *string `mandatory:"true" contributesTo:"path" name:"wlsDomainId"`

	// The unique identifier of a server.
	// **Note:** Not an OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	ServerId *string `mandatory:"true" contributesTo:"path" name:"serverId"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token that represents the page at which to start retrieving results. The token is usually retrieved from a previous List call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order is either 'ASC' or 'DESC'.
	SortOrder ListWlsDomainServerInstalledPatchesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field by which to sort the resource. Only one sort order may be provided.
	// Default order for _displayName_ is **ascending**.
	// If no value is specified, _displayName_ is default.
	SortBy ListWlsDomainServerInstalledPatchesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListWlsDomainServerInstalledPatchesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListWlsDomainServerInstalledPatchesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListWlsDomainServerInstalledPatchesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListWlsDomainServerInstalledPatchesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListWlsDomainServerInstalledPatchesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListWlsDomainServerInstalledPatchesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListWlsDomainServerInstalledPatchesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListWlsDomainServerInstalledPatchesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListWlsDomainServerInstalledPatchesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListWlsDomainServerInstalledPatchesResponse wrapper for the ListWlsDomainServerInstalledPatches operation
type ListWlsDomainServerInstalledPatchesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of InstalledPatchCollection instances
	InstalledPatchCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListWlsDomainServerInstalledPatchesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListWlsDomainServerInstalledPatchesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListWlsDomainServerInstalledPatchesSortOrderEnum Enum with underlying type: string
type ListWlsDomainServerInstalledPatchesSortOrderEnum string

// Set of constants representing the allowable values for ListWlsDomainServerInstalledPatchesSortOrderEnum
const (
	ListWlsDomainServerInstalledPatchesSortOrderAsc  ListWlsDomainServerInstalledPatchesSortOrderEnum = "ASC"
	ListWlsDomainServerInstalledPatchesSortOrderDesc ListWlsDomainServerInstalledPatchesSortOrderEnum = "DESC"
)

var mappingListWlsDomainServerInstalledPatchesSortOrderEnum = map[string]ListWlsDomainServerInstalledPatchesSortOrderEnum{
	"ASC":  ListWlsDomainServerInstalledPatchesSortOrderAsc,
	"DESC": ListWlsDomainServerInstalledPatchesSortOrderDesc,
}

var mappingListWlsDomainServerInstalledPatchesSortOrderEnumLowerCase = map[string]ListWlsDomainServerInstalledPatchesSortOrderEnum{
	"asc":  ListWlsDomainServerInstalledPatchesSortOrderAsc,
	"desc": ListWlsDomainServerInstalledPatchesSortOrderDesc,
}

// GetListWlsDomainServerInstalledPatchesSortOrderEnumValues Enumerates the set of values for ListWlsDomainServerInstalledPatchesSortOrderEnum
func GetListWlsDomainServerInstalledPatchesSortOrderEnumValues() []ListWlsDomainServerInstalledPatchesSortOrderEnum {
	values := make([]ListWlsDomainServerInstalledPatchesSortOrderEnum, 0)
	for _, v := range mappingListWlsDomainServerInstalledPatchesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListWlsDomainServerInstalledPatchesSortOrderEnumStringValues Enumerates the set of values in String for ListWlsDomainServerInstalledPatchesSortOrderEnum
func GetListWlsDomainServerInstalledPatchesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListWlsDomainServerInstalledPatchesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListWlsDomainServerInstalledPatchesSortOrderEnum(val string) (ListWlsDomainServerInstalledPatchesSortOrderEnum, bool) {
	enum, ok := mappingListWlsDomainServerInstalledPatchesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListWlsDomainServerInstalledPatchesSortByEnum Enum with underlying type: string
type ListWlsDomainServerInstalledPatchesSortByEnum string

// Set of constants representing the allowable values for ListWlsDomainServerInstalledPatchesSortByEnum
const (
	ListWlsDomainServerInstalledPatchesSortByDisplayname ListWlsDomainServerInstalledPatchesSortByEnum = "displayName"
)

var mappingListWlsDomainServerInstalledPatchesSortByEnum = map[string]ListWlsDomainServerInstalledPatchesSortByEnum{
	"displayName": ListWlsDomainServerInstalledPatchesSortByDisplayname,
}

var mappingListWlsDomainServerInstalledPatchesSortByEnumLowerCase = map[string]ListWlsDomainServerInstalledPatchesSortByEnum{
	"displayname": ListWlsDomainServerInstalledPatchesSortByDisplayname,
}

// GetListWlsDomainServerInstalledPatchesSortByEnumValues Enumerates the set of values for ListWlsDomainServerInstalledPatchesSortByEnum
func GetListWlsDomainServerInstalledPatchesSortByEnumValues() []ListWlsDomainServerInstalledPatchesSortByEnum {
	values := make([]ListWlsDomainServerInstalledPatchesSortByEnum, 0)
	for _, v := range mappingListWlsDomainServerInstalledPatchesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListWlsDomainServerInstalledPatchesSortByEnumStringValues Enumerates the set of values in String for ListWlsDomainServerInstalledPatchesSortByEnum
func GetListWlsDomainServerInstalledPatchesSortByEnumStringValues() []string {
	return []string{
		"displayName",
	}
}

// GetMappingListWlsDomainServerInstalledPatchesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListWlsDomainServerInstalledPatchesSortByEnum(val string) (ListWlsDomainServerInstalledPatchesSortByEnum, bool) {
	enum, ok := mappingListWlsDomainServerInstalledPatchesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
