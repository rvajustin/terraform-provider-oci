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

// ListCloudListenerServicesRequest wrapper for the ListCloudListenerServices operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListCloudListenerServices.go.html to see an example of how to use ListCloudListenerServicesRequest.
type ListCloudListenerServicesRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud listener.
	CloudListenerId *string `mandatory:"true" contributesTo:"path" name:"cloudListenerId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Managed Database.
	ManagedDatabaseId *string `mandatory:"true" contributesTo:"query" name:"managedDatabaseId"`

	// The page token representing the page from where the next set of paginated results
	// are retrieved. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of records returned in the paginated response.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The field to sort information by. Only one sortOrder can be used. The
	// default sort order for `NAME` is ascending and it is case-sensitive.
	SortBy ListCloudListenerServicesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The option to sort information in ascending (‘ASC’) or descending (‘DESC’) order. Ascending order is the default order.
	SortOrder ListCloudListenerServicesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The OCID of the Named Credential.
	OpcNamedCredentialId *string `mandatory:"false" contributesTo:"header" name:"opc-named-credential-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListCloudListenerServicesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListCloudListenerServicesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListCloudListenerServicesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListCloudListenerServicesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListCloudListenerServicesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListCloudListenerServicesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListCloudListenerServicesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCloudListenerServicesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListCloudListenerServicesSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListCloudListenerServicesResponse wrapper for the ListCloudListenerServices operation
type ListCloudListenerServicesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of CloudListenerServiceCollection instances
	CloudListenerServiceCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListCloudListenerServicesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListCloudListenerServicesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListCloudListenerServicesSortByEnum Enum with underlying type: string
type ListCloudListenerServicesSortByEnum string

// Set of constants representing the allowable values for ListCloudListenerServicesSortByEnum
const (
	ListCloudListenerServicesSortByName ListCloudListenerServicesSortByEnum = "NAME"
)

var mappingListCloudListenerServicesSortByEnum = map[string]ListCloudListenerServicesSortByEnum{
	"NAME": ListCloudListenerServicesSortByName,
}

var mappingListCloudListenerServicesSortByEnumLowerCase = map[string]ListCloudListenerServicesSortByEnum{
	"name": ListCloudListenerServicesSortByName,
}

// GetListCloudListenerServicesSortByEnumValues Enumerates the set of values for ListCloudListenerServicesSortByEnum
func GetListCloudListenerServicesSortByEnumValues() []ListCloudListenerServicesSortByEnum {
	values := make([]ListCloudListenerServicesSortByEnum, 0)
	for _, v := range mappingListCloudListenerServicesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListCloudListenerServicesSortByEnumStringValues Enumerates the set of values in String for ListCloudListenerServicesSortByEnum
func GetListCloudListenerServicesSortByEnumStringValues() []string {
	return []string{
		"NAME",
	}
}

// GetMappingListCloudListenerServicesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCloudListenerServicesSortByEnum(val string) (ListCloudListenerServicesSortByEnum, bool) {
	enum, ok := mappingListCloudListenerServicesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListCloudListenerServicesSortOrderEnum Enum with underlying type: string
type ListCloudListenerServicesSortOrderEnum string

// Set of constants representing the allowable values for ListCloudListenerServicesSortOrderEnum
const (
	ListCloudListenerServicesSortOrderAsc  ListCloudListenerServicesSortOrderEnum = "ASC"
	ListCloudListenerServicesSortOrderDesc ListCloudListenerServicesSortOrderEnum = "DESC"
)

var mappingListCloudListenerServicesSortOrderEnum = map[string]ListCloudListenerServicesSortOrderEnum{
	"ASC":  ListCloudListenerServicesSortOrderAsc,
	"DESC": ListCloudListenerServicesSortOrderDesc,
}

var mappingListCloudListenerServicesSortOrderEnumLowerCase = map[string]ListCloudListenerServicesSortOrderEnum{
	"asc":  ListCloudListenerServicesSortOrderAsc,
	"desc": ListCloudListenerServicesSortOrderDesc,
}

// GetListCloudListenerServicesSortOrderEnumValues Enumerates the set of values for ListCloudListenerServicesSortOrderEnum
func GetListCloudListenerServicesSortOrderEnumValues() []ListCloudListenerServicesSortOrderEnum {
	values := make([]ListCloudListenerServicesSortOrderEnum, 0)
	for _, v := range mappingListCloudListenerServicesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListCloudListenerServicesSortOrderEnumStringValues Enumerates the set of values in String for ListCloudListenerServicesSortOrderEnum
func GetListCloudListenerServicesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListCloudListenerServicesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCloudListenerServicesSortOrderEnum(val string) (ListCloudListenerServicesSortOrderEnum, bool) {
	enum, ok := mappingListCloudListenerServicesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
