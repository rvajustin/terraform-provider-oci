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

// ListCloudDbSystemConnectorsRequest wrapper for the ListCloudDbSystemConnectors operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListCloudDbSystemConnectors.go.html to see an example of how to use ListCloudDbSystemConnectorsRequest.
type ListCloudDbSystemConnectorsRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud DB system.
	CloudDbSystemId *string `mandatory:"false" contributesTo:"query" name:"cloudDbSystemId"`

	// A filter to only return the resources that match the entire display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The page token representing the page from where the next set of paginated results
	// are retrieved. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of records returned in the paginated response.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The field to sort information by. Only one sortOrder can be used. The default sort order
	// for `TIMECREATED` is descending and the default sort order for `DISPLAYNAME` is ascending.
	// The `DISPLAYNAME` sort order is case-sensitive.
	SortBy ListCloudDbSystemConnectorsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The option to sort information in ascending (‘ASC’) or descending (‘DESC’) order. Ascending order is the default order.
	SortOrder ListCloudDbSystemConnectorsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListCloudDbSystemConnectorsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListCloudDbSystemConnectorsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListCloudDbSystemConnectorsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListCloudDbSystemConnectorsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListCloudDbSystemConnectorsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListCloudDbSystemConnectorsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListCloudDbSystemConnectorsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCloudDbSystemConnectorsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListCloudDbSystemConnectorsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListCloudDbSystemConnectorsResponse wrapper for the ListCloudDbSystemConnectors operation
type ListCloudDbSystemConnectorsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of CloudDbSystemConnectorCollection instances
	CloudDbSystemConnectorCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListCloudDbSystemConnectorsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListCloudDbSystemConnectorsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListCloudDbSystemConnectorsSortByEnum Enum with underlying type: string
type ListCloudDbSystemConnectorsSortByEnum string

// Set of constants representing the allowable values for ListCloudDbSystemConnectorsSortByEnum
const (
	ListCloudDbSystemConnectorsSortByTimecreated ListCloudDbSystemConnectorsSortByEnum = "TIMECREATED"
	ListCloudDbSystemConnectorsSortByDisplayname ListCloudDbSystemConnectorsSortByEnum = "DISPLAYNAME"
)

var mappingListCloudDbSystemConnectorsSortByEnum = map[string]ListCloudDbSystemConnectorsSortByEnum{
	"TIMECREATED": ListCloudDbSystemConnectorsSortByTimecreated,
	"DISPLAYNAME": ListCloudDbSystemConnectorsSortByDisplayname,
}

var mappingListCloudDbSystemConnectorsSortByEnumLowerCase = map[string]ListCloudDbSystemConnectorsSortByEnum{
	"timecreated": ListCloudDbSystemConnectorsSortByTimecreated,
	"displayname": ListCloudDbSystemConnectorsSortByDisplayname,
}

// GetListCloudDbSystemConnectorsSortByEnumValues Enumerates the set of values for ListCloudDbSystemConnectorsSortByEnum
func GetListCloudDbSystemConnectorsSortByEnumValues() []ListCloudDbSystemConnectorsSortByEnum {
	values := make([]ListCloudDbSystemConnectorsSortByEnum, 0)
	for _, v := range mappingListCloudDbSystemConnectorsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListCloudDbSystemConnectorsSortByEnumStringValues Enumerates the set of values in String for ListCloudDbSystemConnectorsSortByEnum
func GetListCloudDbSystemConnectorsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListCloudDbSystemConnectorsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCloudDbSystemConnectorsSortByEnum(val string) (ListCloudDbSystemConnectorsSortByEnum, bool) {
	enum, ok := mappingListCloudDbSystemConnectorsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListCloudDbSystemConnectorsSortOrderEnum Enum with underlying type: string
type ListCloudDbSystemConnectorsSortOrderEnum string

// Set of constants representing the allowable values for ListCloudDbSystemConnectorsSortOrderEnum
const (
	ListCloudDbSystemConnectorsSortOrderAsc  ListCloudDbSystemConnectorsSortOrderEnum = "ASC"
	ListCloudDbSystemConnectorsSortOrderDesc ListCloudDbSystemConnectorsSortOrderEnum = "DESC"
)

var mappingListCloudDbSystemConnectorsSortOrderEnum = map[string]ListCloudDbSystemConnectorsSortOrderEnum{
	"ASC":  ListCloudDbSystemConnectorsSortOrderAsc,
	"DESC": ListCloudDbSystemConnectorsSortOrderDesc,
}

var mappingListCloudDbSystemConnectorsSortOrderEnumLowerCase = map[string]ListCloudDbSystemConnectorsSortOrderEnum{
	"asc":  ListCloudDbSystemConnectorsSortOrderAsc,
	"desc": ListCloudDbSystemConnectorsSortOrderDesc,
}

// GetListCloudDbSystemConnectorsSortOrderEnumValues Enumerates the set of values for ListCloudDbSystemConnectorsSortOrderEnum
func GetListCloudDbSystemConnectorsSortOrderEnumValues() []ListCloudDbSystemConnectorsSortOrderEnum {
	values := make([]ListCloudDbSystemConnectorsSortOrderEnum, 0)
	for _, v := range mappingListCloudDbSystemConnectorsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListCloudDbSystemConnectorsSortOrderEnumStringValues Enumerates the set of values in String for ListCloudDbSystemConnectorsSortOrderEnum
func GetListCloudDbSystemConnectorsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListCloudDbSystemConnectorsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCloudDbSystemConnectorsSortOrderEnum(val string) (ListCloudDbSystemConnectorsSortOrderEnum, bool) {
	enum, ok := mappingListCloudDbSystemConnectorsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
