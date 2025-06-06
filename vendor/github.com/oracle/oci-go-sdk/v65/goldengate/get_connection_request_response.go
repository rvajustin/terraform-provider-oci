// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package goldengate

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// GetConnectionRequest wrapper for the GetConnection operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/GetConnection.go.html to see an example of how to use GetConnectionRequest.
type GetConnectionRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a Connection.
	ConnectionId *string `mandatory:"true" contributesTo:"path" name:"connectionId"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Selects the connection fields returned in connection details.
	View GetConnectionViewEnum `mandatory:"false" contributesTo:"query" name:"view" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetConnectionRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetConnectionRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetConnectionRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetConnectionRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GetConnectionRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingGetConnectionViewEnum(string(request.View)); !ok && request.View != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for View: %s. Supported values are: %s.", request.View, strings.Join(GetGetConnectionViewEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GetConnectionResponse wrapper for the GetConnection operation
type GetConnectionResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The Connection instance
	Connection `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// A unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please include the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetConnectionResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetConnectionResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetConnectionViewEnum Enum with underlying type: string
type GetConnectionViewEnum string

// Set of constants representing the allowable values for GetConnectionViewEnum
const (
	GetConnectionViewFull    GetConnectionViewEnum = "FULL"
	GetConnectionViewCompact GetConnectionViewEnum = "COMPACT"
)

var mappingGetConnectionViewEnum = map[string]GetConnectionViewEnum{
	"FULL":    GetConnectionViewFull,
	"COMPACT": GetConnectionViewCompact,
}

var mappingGetConnectionViewEnumLowerCase = map[string]GetConnectionViewEnum{
	"full":    GetConnectionViewFull,
	"compact": GetConnectionViewCompact,
}

// GetGetConnectionViewEnumValues Enumerates the set of values for GetConnectionViewEnum
func GetGetConnectionViewEnumValues() []GetConnectionViewEnum {
	values := make([]GetConnectionViewEnum, 0)
	for _, v := range mappingGetConnectionViewEnum {
		values = append(values, v)
	}
	return values
}

// GetGetConnectionViewEnumStringValues Enumerates the set of values in String for GetConnectionViewEnum
func GetGetConnectionViewEnumStringValues() []string {
	return []string{
		"FULL",
		"COMPACT",
	}
}

// GetMappingGetConnectionViewEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetConnectionViewEnum(val string) (GetConnectionViewEnum, bool) {
	enum, ok := mappingGetConnectionViewEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
