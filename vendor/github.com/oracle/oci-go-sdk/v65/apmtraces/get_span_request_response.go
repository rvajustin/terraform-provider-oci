// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package apmtraces

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// GetSpanRequest wrapper for the GetSpan operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/apmtraces/GetSpan.go.html to see an example of how to use GetSpanRequest.
type GetSpanRequest struct {

	// The APM Domain ID for the intended request.
	ApmDomainId *string `mandatory:"true" contributesTo:"query" name:"apmDomainId"`

	// Unique Application Performance Monitoring span identifier (spanId).
	SpanKey *string `mandatory:"true" contributesTo:"path" name:"spanKey"`

	// Unique Application Performance Monitoring trace identifier (traceId).
	TraceKey *string `mandatory:"true" contributesTo:"path" name:"traceKey"`

	// Unique Oracle-assigned identifier for the request.  If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Include spans that have a `spanStartTime` equal to or greater than this value.
	TimeSpanStartedGreaterThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeSpanStartedGreaterThanOrEqualTo"`

	// Include spans that have a `spanStartTime`less than this value.
	TimeSpanStartedLessThan *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeSpanStartedLessThan"`

	// Name space from which the span details need to be retrieved.
	SpanNamespace GetSpanSpanNamespaceEnum `mandatory:"false" contributesTo:"query" name:"spanNamespace" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetSpanRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetSpanRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetSpanRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetSpanRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GetSpanRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingGetSpanSpanNamespaceEnum(string(request.SpanNamespace)); !ok && request.SpanNamespace != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SpanNamespace: %s. Supported values are: %s.", request.SpanNamespace, strings.Join(GetGetSpanSpanNamespaceEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GetSpanResponse wrapper for the GetSpan operation
type GetSpanResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The Span instance
	Span `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetSpanResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetSpanResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetSpanSpanNamespaceEnum Enum with underlying type: string
type GetSpanSpanNamespaceEnum string

// Set of constants representing the allowable values for GetSpanSpanNamespaceEnum
const (
	GetSpanSpanNamespaceTraces    GetSpanSpanNamespaceEnum = "TRACES"
	GetSpanSpanNamespaceSynthetic GetSpanSpanNamespaceEnum = "SYNTHETIC"
)

var mappingGetSpanSpanNamespaceEnum = map[string]GetSpanSpanNamespaceEnum{
	"TRACES":    GetSpanSpanNamespaceTraces,
	"SYNTHETIC": GetSpanSpanNamespaceSynthetic,
}

var mappingGetSpanSpanNamespaceEnumLowerCase = map[string]GetSpanSpanNamespaceEnum{
	"traces":    GetSpanSpanNamespaceTraces,
	"synthetic": GetSpanSpanNamespaceSynthetic,
}

// GetGetSpanSpanNamespaceEnumValues Enumerates the set of values for GetSpanSpanNamespaceEnum
func GetGetSpanSpanNamespaceEnumValues() []GetSpanSpanNamespaceEnum {
	values := make([]GetSpanSpanNamespaceEnum, 0)
	for _, v := range mappingGetSpanSpanNamespaceEnum {
		values = append(values, v)
	}
	return values
}

// GetGetSpanSpanNamespaceEnumStringValues Enumerates the set of values in String for GetSpanSpanNamespaceEnum
func GetGetSpanSpanNamespaceEnumStringValues() []string {
	return []string{
		"TRACES",
		"SYNTHETIC",
	}
}

// GetMappingGetSpanSpanNamespaceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetSpanSpanNamespaceEnum(val string) (GetSpanSpanNamespaceEnum, bool) {
	enum, ok := mappingGetSpanSpanNamespaceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
