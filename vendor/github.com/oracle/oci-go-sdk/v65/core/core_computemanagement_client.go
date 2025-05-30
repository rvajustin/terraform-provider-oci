// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// Use the Core Services API to manage resources such as virtual cloud networks (VCNs),
// compute instances, and block storage volumes. For more information, see the console
// documentation for the Networking (https://docs.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.oracle.com/iaas/Content/Block/Concepts/overview.htm) services.
// The required permissions are documented in the
// Details for the Core Services (https://docs.oracle.com/iaas/Content/Identity/Reference/corepolicyreference.htm) article.
//

package core

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// ComputeManagementClient a client for ComputeManagement
type ComputeManagementClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewComputeManagementClientWithConfigurationProvider Creates a new default ComputeManagement client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewComputeManagementClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client ComputeManagementClient, err error) {
	if enabled := common.CheckForEnabledServices("core"); !enabled {
		return client, fmt.Errorf("the Developer Tool configuration disabled this service, this behavior is controlled by OciSdkEnabledServicesMap variables. Please check if your local developer-tool-configuration.json file configured the service you're targeting or contact the cloud provider on the availability of this service")
	}
	provider, err := auth.GetGenericConfigurationProvider(configProvider)
	if err != nil {
		return client, err
	}
	baseClient, e := common.NewClientWithConfig(provider)
	if e != nil {
		return client, e
	}
	return newComputeManagementClientFromBaseClient(baseClient, provider)
}

// NewComputeManagementClientWithOboToken Creates a new default ComputeManagement client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewComputeManagementClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client ComputeManagementClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newComputeManagementClientFromBaseClient(baseClient, configProvider)
}

func newComputeManagementClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client ComputeManagementClient, err error) {
	// ComputeManagement service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("ComputeManagement"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = ComputeManagementClient{BaseClient: baseClient}
	client.BasePath = "20160918"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *ComputeManagementClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("iaas", "https://iaas.{region}.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *ComputeManagementClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
	if ok, err := common.IsConfigurationProviderValid(configProvider); !ok {
		return err
	}

	// Error has been checked already
	region, _ := configProvider.Region()
	client.SetRegion(region)
	if client.Host == "" {
		return fmt.Errorf("invalid region or Host. Endpoint cannot be constructed without endpointServiceName or serviceEndpointTemplate for a dotted region")
	}
	client.config = &configProvider
	return nil
}

// ConfigurationProvider the ConfigurationProvider used in this client, or null if none set
func (client *ComputeManagementClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// AttachInstancePoolInstance Attaches an instance to an instance pool. For information about the prerequisites
// that an instance must meet before you can attach it to a pool, see
// Attaching an Instance to an Instance Pool (https://docs.oracle.com/iaas/Content/Compute/Tasks/updatinginstancepool.htm#attach-instance).
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/core/AttachInstancePoolInstance.go.html to see an example of how to use AttachInstancePoolInstance API.
func (client ComputeManagementClient) AttachInstancePoolInstance(ctx context.Context, request AttachInstancePoolInstanceRequest) (response AttachInstancePoolInstanceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.attachInstancePoolInstance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = AttachInstancePoolInstanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = AttachInstancePoolInstanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(AttachInstancePoolInstanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into AttachInstancePoolInstanceResponse")
	}
	return
}

// attachInstancePoolInstance implements the OCIOperation interface (enables retrying operations)
func (client ComputeManagementClient) attachInstancePoolInstance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/instancePools/{instancePoolId}/instances", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response AttachInstancePoolInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "ComputeManagement", "AttachInstancePoolInstance", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// AttachLoadBalancer Attach a load balancer to the instance pool.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/core/AttachLoadBalancer.go.html to see an example of how to use AttachLoadBalancer API.
func (client ComputeManagementClient) AttachLoadBalancer(ctx context.Context, request AttachLoadBalancerRequest) (response AttachLoadBalancerResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.attachLoadBalancer, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = AttachLoadBalancerResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = AttachLoadBalancerResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(AttachLoadBalancerResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into AttachLoadBalancerResponse")
	}
	return
}

// attachLoadBalancer implements the OCIOperation interface (enables retrying operations)
func (client ComputeManagementClient) attachLoadBalancer(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/instancePools/{instancePoolId}/actions/attachLoadBalancer", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response AttachLoadBalancerResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/iaas/20160918/InstancePool/AttachLoadBalancer"
		err = common.PostProcessServiceError(err, "ComputeManagement", "AttachLoadBalancer", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeClusterNetworkCompartment Moves a cluster network with instance pools (https://docs.oracle.com/iaas/Content/Compute/Tasks/managingclusternetworks.htm)
// into a different compartment within the same tenancy. For
// information about moving resources between compartments, see
// Moving Resources to a Different Compartment (https://docs.oracle.com/iaas/Content/Identity/Tasks/managingcompartments.htm#moveRes).
// When you move a cluster network to a different compartment, associated resources such as the instances
// in the cluster network, boot volumes, and VNICs are not moved.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/core/ChangeClusterNetworkCompartment.go.html to see an example of how to use ChangeClusterNetworkCompartment API.
func (client ComputeManagementClient) ChangeClusterNetworkCompartment(ctx context.Context, request ChangeClusterNetworkCompartmentRequest) (response ChangeClusterNetworkCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.changeClusterNetworkCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeClusterNetworkCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeClusterNetworkCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeClusterNetworkCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeClusterNetworkCompartmentResponse")
	}
	return
}

// changeClusterNetworkCompartment implements the OCIOperation interface (enables retrying operations)
func (client ComputeManagementClient) changeClusterNetworkCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/clusterNetworks/{clusterNetworkId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeClusterNetworkCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/iaas/20160918/ClusterNetwork/ChangeClusterNetworkCompartment"
		err = common.PostProcessServiceError(err, "ComputeManagement", "ChangeClusterNetworkCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeInstanceConfigurationCompartment Moves an instance configuration into a different compartment within the same tenancy.
// For information about moving resources between compartments, see
// Moving Resources to a Different Compartment (https://docs.oracle.com/iaas/Content/Identity/Tasks/managingcompartments.htm#moveRes).
// When you move an instance configuration to a different compartment, associated resources such as
// instance pools are not moved.
// **Important:** Most of the properties for an existing instance configuration, including the compartment,
// cannot be modified after you create the instance configuration. Although you can move an instance configuration
// to a different compartment, you will not be able to use the instance configuration to manage instance pools
// in the new compartment. If you want to update an instance configuration to point to a different compartment,
// you should instead create a new instance configuration in the target compartment using
// CreateInstanceConfiguration (https://docs.oracle.com/iaas/api/#/en/iaas/20160918/InstanceConfiguration/CreateInstanceConfiguration).
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/core/ChangeInstanceConfigurationCompartment.go.html to see an example of how to use ChangeInstanceConfigurationCompartment API.
func (client ComputeManagementClient) ChangeInstanceConfigurationCompartment(ctx context.Context, request ChangeInstanceConfigurationCompartmentRequest) (response ChangeInstanceConfigurationCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.changeInstanceConfigurationCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeInstanceConfigurationCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeInstanceConfigurationCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeInstanceConfigurationCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeInstanceConfigurationCompartmentResponse")
	}
	return
}

// changeInstanceConfigurationCompartment implements the OCIOperation interface (enables retrying operations)
func (client ComputeManagementClient) changeInstanceConfigurationCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/instanceConfigurations/{instanceConfigurationId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeInstanceConfigurationCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/iaas/20160918/InstanceConfiguration/ChangeInstanceConfigurationCompartment"
		err = common.PostProcessServiceError(err, "ComputeManagement", "ChangeInstanceConfigurationCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeInstancePoolCompartment Moves an instance pool into a different compartment within the same tenancy. For
// information about moving resources between compartments, see
// Moving Resources to a Different Compartment (https://docs.oracle.com/iaas/Content/Identity/Tasks/managingcompartments.htm#moveRes).
// When you move an instance pool to a different compartment, associated resources such as the instances in
// the pool, boot volumes, VNICs, and autoscaling configurations are not moved.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/core/ChangeInstancePoolCompartment.go.html to see an example of how to use ChangeInstancePoolCompartment API.
func (client ComputeManagementClient) ChangeInstancePoolCompartment(ctx context.Context, request ChangeInstancePoolCompartmentRequest) (response ChangeInstancePoolCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.changeInstancePoolCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeInstancePoolCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeInstancePoolCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeInstancePoolCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeInstancePoolCompartmentResponse")
	}
	return
}

// changeInstancePoolCompartment implements the OCIOperation interface (enables retrying operations)
func (client ComputeManagementClient) changeInstancePoolCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/instancePools/{instancePoolId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeInstancePoolCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/iaas/20160918/InstancePool/ChangeInstancePoolCompartment"
		err = common.PostProcessServiceError(err, "ComputeManagement", "ChangeInstancePoolCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateClusterNetwork Creates a cluster network with instance pools (https://docs.oracle.com/iaas/Content/Compute/Tasks/managingclusternetworks.htm).
// A cluster network is a group of high performance computing (HPC), GPU, or optimized bare metal
// instances that are connected with an ultra low-latency remote direct memory access (RDMA) network.
// Cluster networks with instance pools use instance pools to manage groups of identical instances.
// Use cluster networks with instance pools when you want predictable capacity for a specific number of identical
// instances that are managed as a group.
// If you want to manage instances in the RDMA network independently of each other or use different types of instances
// in the network group, create a compute cluster by using the CreateComputeCluster
// operation.
// To determine whether capacity is available for a specific shape before you create a cluster network,
// use the CreateComputeCapacityReport
// operation.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/core/CreateClusterNetwork.go.html to see an example of how to use CreateClusterNetwork API.
func (client ComputeManagementClient) CreateClusterNetwork(ctx context.Context, request CreateClusterNetworkRequest) (response CreateClusterNetworkResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createClusterNetwork, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateClusterNetworkResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateClusterNetworkResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateClusterNetworkResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateClusterNetworkResponse")
	}
	return
}

// createClusterNetwork implements the OCIOperation interface (enables retrying operations)
func (client ComputeManagementClient) createClusterNetwork(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/clusterNetworks", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateClusterNetworkResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/iaas/20160918/ClusterNetwork/CreateClusterNetwork"
		err = common.PostProcessServiceError(err, "ComputeManagement", "CreateClusterNetwork", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateInstanceConfiguration Creates an instance configuration. An instance configuration is a template that defines the
// settings to use when creating Compute instances.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/core/CreateInstanceConfiguration.go.html to see an example of how to use CreateInstanceConfiguration API.
func (client ComputeManagementClient) CreateInstanceConfiguration(ctx context.Context, request CreateInstanceConfigurationRequest) (response CreateInstanceConfigurationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createInstanceConfiguration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateInstanceConfigurationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateInstanceConfigurationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateInstanceConfigurationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateInstanceConfigurationResponse")
	}
	return
}

// createInstanceConfiguration implements the OCIOperation interface (enables retrying operations)
func (client ComputeManagementClient) createInstanceConfiguration(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/instanceConfigurations", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateInstanceConfigurationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/iaas/20160918/InstanceConfiguration/CreateInstanceConfiguration"
		err = common.PostProcessServiceError(err, "ComputeManagement", "CreateInstanceConfiguration", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateInstancePool Creates an instance pool.
// To determine whether capacity is available for a specific shape before you create an instance pool,
// use the CreateComputeCapacityReport
// operation.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/core/CreateInstancePool.go.html to see an example of how to use CreateInstancePool API.
func (client ComputeManagementClient) CreateInstancePool(ctx context.Context, request CreateInstancePoolRequest) (response CreateInstancePoolResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createInstancePool, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateInstancePoolResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateInstancePoolResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateInstancePoolResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateInstancePoolResponse")
	}
	return
}

// createInstancePool implements the OCIOperation interface (enables retrying operations)
func (client ComputeManagementClient) createInstancePool(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/instancePools", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateInstancePoolResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/iaas/20160918/InstancePool/CreateInstancePool"
		err = common.PostProcessServiceError(err, "ComputeManagement", "CreateInstancePool", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteInstanceConfiguration Deletes an instance configuration.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/core/DeleteInstanceConfiguration.go.html to see an example of how to use DeleteInstanceConfiguration API.
func (client ComputeManagementClient) DeleteInstanceConfiguration(ctx context.Context, request DeleteInstanceConfigurationRequest) (response DeleteInstanceConfigurationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteInstanceConfiguration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteInstanceConfigurationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteInstanceConfigurationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteInstanceConfigurationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteInstanceConfigurationResponse")
	}
	return
}

// deleteInstanceConfiguration implements the OCIOperation interface (enables retrying operations)
func (client ComputeManagementClient) deleteInstanceConfiguration(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/instanceConfigurations/{instanceConfigurationId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteInstanceConfigurationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "ComputeManagement", "DeleteInstanceConfiguration", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DetachInstancePoolInstance Detaches an instance from an instance pool.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/core/DetachInstancePoolInstance.go.html to see an example of how to use DetachInstancePoolInstance API.
func (client ComputeManagementClient) DetachInstancePoolInstance(ctx context.Context, request DetachInstancePoolInstanceRequest) (response DetachInstancePoolInstanceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.detachInstancePoolInstance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DetachInstancePoolInstanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DetachInstancePoolInstanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DetachInstancePoolInstanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DetachInstancePoolInstanceResponse")
	}
	return
}

// detachInstancePoolInstance implements the OCIOperation interface (enables retrying operations)
func (client ComputeManagementClient) detachInstancePoolInstance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/instancePools/{instancePoolId}/actions/detachInstance", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DetachInstancePoolInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/iaas/20160918/InstancePoolInstance/DetachInstancePoolInstance"
		err = common.PostProcessServiceError(err, "ComputeManagement", "DetachInstancePoolInstance", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DetachLoadBalancer Detach a load balancer from the instance pool.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/core/DetachLoadBalancer.go.html to see an example of how to use DetachLoadBalancer API.
func (client ComputeManagementClient) DetachLoadBalancer(ctx context.Context, request DetachLoadBalancerRequest) (response DetachLoadBalancerResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.detachLoadBalancer, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DetachLoadBalancerResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DetachLoadBalancerResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DetachLoadBalancerResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DetachLoadBalancerResponse")
	}
	return
}

// detachLoadBalancer implements the OCIOperation interface (enables retrying operations)
func (client ComputeManagementClient) detachLoadBalancer(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/instancePools/{instancePoolId}/actions/detachLoadBalancer", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DetachLoadBalancerResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/iaas/20160918/InstancePool/DetachLoadBalancer"
		err = common.PostProcessServiceError(err, "ComputeManagement", "DetachLoadBalancer", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetClusterNetwork Gets information about a cluster network with instance pools (https://docs.oracle.com/iaas/Content/Compute/Tasks/managingclusternetworks.htm).
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/core/GetClusterNetwork.go.html to see an example of how to use GetClusterNetwork API.
func (client ComputeManagementClient) GetClusterNetwork(ctx context.Context, request GetClusterNetworkRequest) (response GetClusterNetworkResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getClusterNetwork, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetClusterNetworkResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetClusterNetworkResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetClusterNetworkResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetClusterNetworkResponse")
	}
	return
}

// getClusterNetwork implements the OCIOperation interface (enables retrying operations)
func (client ComputeManagementClient) getClusterNetwork(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/clusterNetworks/{clusterNetworkId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetClusterNetworkResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/iaas/20160918/ClusterNetwork/GetClusterNetwork"
		err = common.PostProcessServiceError(err, "ComputeManagement", "GetClusterNetwork", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetInstanceConfiguration Gets the specified instance configuration
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/core/GetInstanceConfiguration.go.html to see an example of how to use GetInstanceConfiguration API.
func (client ComputeManagementClient) GetInstanceConfiguration(ctx context.Context, request GetInstanceConfigurationRequest) (response GetInstanceConfigurationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getInstanceConfiguration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetInstanceConfigurationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetInstanceConfigurationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetInstanceConfigurationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetInstanceConfigurationResponse")
	}
	return
}

// getInstanceConfiguration implements the OCIOperation interface (enables retrying operations)
func (client ComputeManagementClient) getInstanceConfiguration(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/instanceConfigurations/{instanceConfigurationId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetInstanceConfigurationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/iaas/20160918/InstanceConfiguration/GetInstanceConfiguration"
		err = common.PostProcessServiceError(err, "ComputeManagement", "GetInstanceConfiguration", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetInstancePool Gets the specified instance pool
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/core/GetInstancePool.go.html to see an example of how to use GetInstancePool API.
func (client ComputeManagementClient) GetInstancePool(ctx context.Context, request GetInstancePoolRequest) (response GetInstancePoolResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getInstancePool, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetInstancePoolResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetInstancePoolResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetInstancePoolResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetInstancePoolResponse")
	}
	return
}

// getInstancePool implements the OCIOperation interface (enables retrying operations)
func (client ComputeManagementClient) getInstancePool(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/instancePools/{instancePoolId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetInstancePoolResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/iaas/20160918/InstancePool/GetInstancePool"
		err = common.PostProcessServiceError(err, "ComputeManagement", "GetInstancePool", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetInstancePoolInstance Gets information about an instance that belongs to an instance pool.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/core/GetInstancePoolInstance.go.html to see an example of how to use GetInstancePoolInstance API.
func (client ComputeManagementClient) GetInstancePoolInstance(ctx context.Context, request GetInstancePoolInstanceRequest) (response GetInstancePoolInstanceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getInstancePoolInstance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetInstancePoolInstanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetInstancePoolInstanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetInstancePoolInstanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetInstancePoolInstanceResponse")
	}
	return
}

// getInstancePoolInstance implements the OCIOperation interface (enables retrying operations)
func (client ComputeManagementClient) getInstancePoolInstance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/instancePools/{instancePoolId}/instances/{instanceId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetInstancePoolInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/iaas/20160918/InstancePoolInstance/GetInstancePoolInstance"
		err = common.PostProcessServiceError(err, "ComputeManagement", "GetInstancePoolInstance", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetInstancePoolLoadBalancerAttachment Gets information about a load balancer that is attached to the specified instance pool.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/core/GetInstancePoolLoadBalancerAttachment.go.html to see an example of how to use GetInstancePoolLoadBalancerAttachment API.
func (client ComputeManagementClient) GetInstancePoolLoadBalancerAttachment(ctx context.Context, request GetInstancePoolLoadBalancerAttachmentRequest) (response GetInstancePoolLoadBalancerAttachmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getInstancePoolLoadBalancerAttachment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetInstancePoolLoadBalancerAttachmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetInstancePoolLoadBalancerAttachmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetInstancePoolLoadBalancerAttachmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetInstancePoolLoadBalancerAttachmentResponse")
	}
	return
}

// getInstancePoolLoadBalancerAttachment implements the OCIOperation interface (enables retrying operations)
func (client ComputeManagementClient) getInstancePoolLoadBalancerAttachment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/instancePools/{instancePoolId}/loadBalancerAttachments/{instancePoolLoadBalancerAttachmentId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetInstancePoolLoadBalancerAttachmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/iaas/20160918/InstancePoolLoadBalancerAttachment/GetInstancePoolLoadBalancerAttachment"
		err = common.PostProcessServiceError(err, "ComputeManagement", "GetInstancePoolLoadBalancerAttachment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// LaunchInstanceConfiguration Creates an instance from an instance configuration.
// If the instance configuration does not include all of the parameters that are
// required to create an instance, such as the availability domain and subnet ID, you must
// provide these parameters when you create an instance from the instance configuration.
// For more information, see the InstanceConfiguration
// resource.
// To determine whether capacity is available for a specific shape before you create an instance,
// use the CreateComputeCapacityReport
// operation.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/core/LaunchInstanceConfiguration.go.html to see an example of how to use LaunchInstanceConfiguration API.
func (client ComputeManagementClient) LaunchInstanceConfiguration(ctx context.Context, request LaunchInstanceConfigurationRequest) (response LaunchInstanceConfigurationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.launchInstanceConfiguration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = LaunchInstanceConfigurationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = LaunchInstanceConfigurationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(LaunchInstanceConfigurationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into LaunchInstanceConfigurationResponse")
	}
	return
}

// launchInstanceConfiguration implements the OCIOperation interface (enables retrying operations)
func (client ComputeManagementClient) launchInstanceConfiguration(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/instanceConfigurations/{instanceConfigurationId}/actions/launch", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response LaunchInstanceConfigurationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/iaas/20160918/Instance/LaunchInstanceConfiguration"
		err = common.PostProcessServiceError(err, "ComputeManagement", "LaunchInstanceConfiguration", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListClusterNetworkInstances Lists the instances in a cluster network with instance pools (https://docs.oracle.com/iaas/Content/Compute/Tasks/managingclusternetworks.htm).
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/core/ListClusterNetworkInstances.go.html to see an example of how to use ListClusterNetworkInstances API.
func (client ComputeManagementClient) ListClusterNetworkInstances(ctx context.Context, request ListClusterNetworkInstancesRequest) (response ListClusterNetworkInstancesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listClusterNetworkInstances, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListClusterNetworkInstancesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListClusterNetworkInstancesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListClusterNetworkInstancesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListClusterNetworkInstancesResponse")
	}
	return
}

// listClusterNetworkInstances implements the OCIOperation interface (enables retrying operations)
func (client ComputeManagementClient) listClusterNetworkInstances(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/clusterNetworks/{clusterNetworkId}/instances", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListClusterNetworkInstancesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/iaas/20160918/ClusterNetwork/ListClusterNetworkInstances"
		err = common.PostProcessServiceError(err, "ComputeManagement", "ListClusterNetworkInstances", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListClusterNetworks Lists the cluster networks with instance pools (https://docs.oracle.com/iaas/Content/Compute/Tasks/managingclusternetworks.htm)
// in the specified compartment.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/core/ListClusterNetworks.go.html to see an example of how to use ListClusterNetworks API.
func (client ComputeManagementClient) ListClusterNetworks(ctx context.Context, request ListClusterNetworksRequest) (response ListClusterNetworksResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listClusterNetworks, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListClusterNetworksResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListClusterNetworksResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListClusterNetworksResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListClusterNetworksResponse")
	}
	return
}

// listClusterNetworks implements the OCIOperation interface (enables retrying operations)
func (client ComputeManagementClient) listClusterNetworks(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/clusterNetworks", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListClusterNetworksResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/iaas/20160918/ClusterNetwork/ListClusterNetworks"
		err = common.PostProcessServiceError(err, "ComputeManagement", "ListClusterNetworks", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListInstanceConfigurations Lists the instance configurations in the specified compartment.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/core/ListInstanceConfigurations.go.html to see an example of how to use ListInstanceConfigurations API.
func (client ComputeManagementClient) ListInstanceConfigurations(ctx context.Context, request ListInstanceConfigurationsRequest) (response ListInstanceConfigurationsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listInstanceConfigurations, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListInstanceConfigurationsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListInstanceConfigurationsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListInstanceConfigurationsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListInstanceConfigurationsResponse")
	}
	return
}

// listInstanceConfigurations implements the OCIOperation interface (enables retrying operations)
func (client ComputeManagementClient) listInstanceConfigurations(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/instanceConfigurations", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListInstanceConfigurationsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/iaas/20160918/InstanceConfigurationSummary/ListInstanceConfigurations"
		err = common.PostProcessServiceError(err, "ComputeManagement", "ListInstanceConfigurations", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListInstancePoolInstances List the instances in the specified instance pool.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/core/ListInstancePoolInstances.go.html to see an example of how to use ListInstancePoolInstances API.
func (client ComputeManagementClient) ListInstancePoolInstances(ctx context.Context, request ListInstancePoolInstancesRequest) (response ListInstancePoolInstancesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listInstancePoolInstances, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListInstancePoolInstancesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListInstancePoolInstancesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListInstancePoolInstancesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListInstancePoolInstancesResponse")
	}
	return
}

// listInstancePoolInstances implements the OCIOperation interface (enables retrying operations)
func (client ComputeManagementClient) listInstancePoolInstances(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/instancePools/{instancePoolId}/instances", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListInstancePoolInstancesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/iaas/20160918/InstanceSummary/ListInstancePoolInstances"
		err = common.PostProcessServiceError(err, "ComputeManagement", "ListInstancePoolInstances", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListInstancePools Lists the instance pools in the specified compartment.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/core/ListInstancePools.go.html to see an example of how to use ListInstancePools API.
func (client ComputeManagementClient) ListInstancePools(ctx context.Context, request ListInstancePoolsRequest) (response ListInstancePoolsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listInstancePools, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListInstancePoolsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListInstancePoolsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListInstancePoolsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListInstancePoolsResponse")
	}
	return
}

// listInstancePools implements the OCIOperation interface (enables retrying operations)
func (client ComputeManagementClient) listInstancePools(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/instancePools", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListInstancePoolsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/iaas/20160918/InstancePoolSummary/ListInstancePools"
		err = common.PostProcessServiceError(err, "ComputeManagement", "ListInstancePools", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ResetInstancePool Performs the reset (immediate power off and power on) action on the specified instance pool,
// which performs the action on all the instances in the pool.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/core/ResetInstancePool.go.html to see an example of how to use ResetInstancePool API.
func (client ComputeManagementClient) ResetInstancePool(ctx context.Context, request ResetInstancePoolRequest) (response ResetInstancePoolResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.resetInstancePool, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ResetInstancePoolResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ResetInstancePoolResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ResetInstancePoolResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ResetInstancePoolResponse")
	}
	return
}

// resetInstancePool implements the OCIOperation interface (enables retrying operations)
func (client ComputeManagementClient) resetInstancePool(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/instancePools/{instancePoolId}/actions/reset", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ResetInstancePoolResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/iaas/20160918/InstancePool/ResetInstancePool"
		err = common.PostProcessServiceError(err, "ComputeManagement", "ResetInstancePool", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SoftresetInstancePool Performs the softreset (ACPI shutdown and power on) action on the specified instance pool,
// which performs the action on all the instances in the pool.
// Softreset gracefully reboots the instances by sending a shutdown command to the operating systems.
// After waiting 15 minutes for the OS to shut down, the instances are powered off and then powered back on.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/core/SoftresetInstancePool.go.html to see an example of how to use SoftresetInstancePool API.
func (client ComputeManagementClient) SoftresetInstancePool(ctx context.Context, request SoftresetInstancePoolRequest) (response SoftresetInstancePoolResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.softresetInstancePool, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SoftresetInstancePoolResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SoftresetInstancePoolResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SoftresetInstancePoolResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SoftresetInstancePoolResponse")
	}
	return
}

// softresetInstancePool implements the OCIOperation interface (enables retrying operations)
func (client ComputeManagementClient) softresetInstancePool(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/instancePools/{instancePoolId}/actions/softreset", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SoftresetInstancePoolResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/iaas/20160918/InstancePool/SoftresetInstancePool"
		err = common.PostProcessServiceError(err, "ComputeManagement", "SoftresetInstancePool", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SoftstopInstancePool Performs the softstop (ACPI shutdown and power on) action on the specified instance pool,
// which performs the action on all the instances in the pool.
// Softstop gracefully reboots the instances by sending a shutdown command to the operating systems.
// After waiting 15 minutes for the OS to shutdown, the instances are powered off and then powered back on.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/core/SoftstopInstancePool.go.html to see an example of how to use SoftstopInstancePool API.
func (client ComputeManagementClient) SoftstopInstancePool(ctx context.Context, request SoftstopInstancePoolRequest) (response SoftstopInstancePoolResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.softstopInstancePool, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SoftstopInstancePoolResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SoftstopInstancePoolResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SoftstopInstancePoolResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SoftstopInstancePoolResponse")
	}
	return
}

// softstopInstancePool implements the OCIOperation interface (enables retrying operations)
func (client ComputeManagementClient) softstopInstancePool(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/instancePools/{instancePoolId}/actions/softstop", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SoftstopInstancePoolResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/iaas/20160918/InstancePool/SoftstopInstancePool"
		err = common.PostProcessServiceError(err, "ComputeManagement", "SoftstopInstancePool", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// StartInstancePool Performs the start (power on) action on the specified instance pool,
// which performs the action on all the instances in the pool.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/core/StartInstancePool.go.html to see an example of how to use StartInstancePool API.
func (client ComputeManagementClient) StartInstancePool(ctx context.Context, request StartInstancePoolRequest) (response StartInstancePoolResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.startInstancePool, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = StartInstancePoolResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = StartInstancePoolResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(StartInstancePoolResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into StartInstancePoolResponse")
	}
	return
}

// startInstancePool implements the OCIOperation interface (enables retrying operations)
func (client ComputeManagementClient) startInstancePool(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/instancePools/{instancePoolId}/actions/start", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response StartInstancePoolResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/iaas/20160918/InstancePool/StartInstancePool"
		err = common.PostProcessServiceError(err, "ComputeManagement", "StartInstancePool", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// StopInstancePool Performs the stop (immediate power off) action on the specified instance pool,
// which performs the action on all the instances in the pool.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/core/StopInstancePool.go.html to see an example of how to use StopInstancePool API.
func (client ComputeManagementClient) StopInstancePool(ctx context.Context, request StopInstancePoolRequest) (response StopInstancePoolResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.stopInstancePool, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = StopInstancePoolResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = StopInstancePoolResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(StopInstancePoolResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into StopInstancePoolResponse")
	}
	return
}

// stopInstancePool implements the OCIOperation interface (enables retrying operations)
func (client ComputeManagementClient) stopInstancePool(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/instancePools/{instancePoolId}/actions/stop", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response StopInstancePoolResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/iaas/20160918/InstancePool/StopInstancePool"
		err = common.PostProcessServiceError(err, "ComputeManagement", "StopInstancePool", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// TerminateClusterNetwork Deletes (terminates) a cluster network with instance pools (https://docs.oracle.com/iaas/Content/Compute/Tasks/managingclusternetworks.htm).
// When you delete a cluster network, all of its resources are permanently deleted,
// including associated instances and instance pools.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/core/TerminateClusterNetwork.go.html to see an example of how to use TerminateClusterNetwork API.
func (client ComputeManagementClient) TerminateClusterNetwork(ctx context.Context, request TerminateClusterNetworkRequest) (response TerminateClusterNetworkResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.terminateClusterNetwork, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = TerminateClusterNetworkResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = TerminateClusterNetworkResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(TerminateClusterNetworkResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into TerminateClusterNetworkResponse")
	}
	return
}

// terminateClusterNetwork implements the OCIOperation interface (enables retrying operations)
func (client ComputeManagementClient) terminateClusterNetwork(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/clusterNetworks/{clusterNetworkId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response TerminateClusterNetworkResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/iaas/20160918/ClusterNetwork/TerminateClusterNetwork"
		err = common.PostProcessServiceError(err, "ComputeManagement", "TerminateClusterNetwork", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// TerminateInstancePool Terminate the specified instance pool.
// **Warning:** When you delete an instance pool, the resources that were created by the pool are permanently
// deleted, including associated instances, attached boot volumes, and block volumes.
// If an autoscaling configuration applies to the instance pool, the autoscaling configuration will be deleted
// asynchronously after the pool is deleted. You can also manually delete the autoscaling configuration using
// the `DeleteAutoScalingConfiguration` operation in the Autoscaling API.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/core/TerminateInstancePool.go.html to see an example of how to use TerminateInstancePool API.
func (client ComputeManagementClient) TerminateInstancePool(ctx context.Context, request TerminateInstancePoolRequest) (response TerminateInstancePoolResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.terminateInstancePool, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = TerminateInstancePoolResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = TerminateInstancePoolResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(TerminateInstancePoolResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into TerminateInstancePoolResponse")
	}
	return
}

// terminateInstancePool implements the OCIOperation interface (enables retrying operations)
func (client ComputeManagementClient) terminateInstancePool(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/instancePools/{instancePoolId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response TerminateInstancePoolResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "ComputeManagement", "TerminateInstancePool", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateClusterNetwork Updates a cluster network with instance pools (https://docs.oracle.com/iaas/Content/Compute/Tasks/managingclusternetworks.htm).
// The OCID of the cluster network remains the same.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/core/UpdateClusterNetwork.go.html to see an example of how to use UpdateClusterNetwork API.
func (client ComputeManagementClient) UpdateClusterNetwork(ctx context.Context, request UpdateClusterNetworkRequest) (response UpdateClusterNetworkResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.updateClusterNetwork, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateClusterNetworkResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateClusterNetworkResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateClusterNetworkResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateClusterNetworkResponse")
	}
	return
}

// updateClusterNetwork implements the OCIOperation interface (enables retrying operations)
func (client ComputeManagementClient) updateClusterNetwork(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/clusterNetworks/{clusterNetworkId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateClusterNetworkResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/iaas/20160918/ClusterNetwork/UpdateClusterNetwork"
		err = common.PostProcessServiceError(err, "ComputeManagement", "UpdateClusterNetwork", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateInstanceConfiguration Updates the free-form tags, defined tags, and display name of an instance configuration.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/core/UpdateInstanceConfiguration.go.html to see an example of how to use UpdateInstanceConfiguration API.
func (client ComputeManagementClient) UpdateInstanceConfiguration(ctx context.Context, request UpdateInstanceConfigurationRequest) (response UpdateInstanceConfigurationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.updateInstanceConfiguration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateInstanceConfigurationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateInstanceConfigurationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateInstanceConfigurationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateInstanceConfigurationResponse")
	}
	return
}

// updateInstanceConfiguration implements the OCIOperation interface (enables retrying operations)
func (client ComputeManagementClient) updateInstanceConfiguration(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/instanceConfigurations/{instanceConfigurationId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateInstanceConfigurationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/iaas/20160918/InstanceConfiguration/UpdateInstanceConfiguration"
		err = common.PostProcessServiceError(err, "ComputeManagement", "UpdateInstanceConfiguration", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateInstancePool Update the specified instance pool.
// The OCID of the instance pool remains the same.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/core/UpdateInstancePool.go.html to see an example of how to use UpdateInstancePool API.
func (client ComputeManagementClient) UpdateInstancePool(ctx context.Context, request UpdateInstancePoolRequest) (response UpdateInstancePoolResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.updateInstancePool, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateInstancePoolResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateInstancePoolResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateInstancePoolResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateInstancePoolResponse")
	}
	return
}

// updateInstancePool implements the OCIOperation interface (enables retrying operations)
func (client ComputeManagementClient) updateInstancePool(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/instancePools/{instancePoolId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateInstancePoolResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/iaas/20160918/InstancePool/UpdateInstancePool"
		err = common.PostProcessServiceError(err, "ComputeManagement", "UpdateInstancePool", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
