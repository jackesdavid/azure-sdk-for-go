//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatafactory

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ManagedVirtualNetworksClient contains the methods for the ManagedVirtualNetworks group.
// Don't use this type directly, use NewManagedVirtualNetworksClient() instead.
type ManagedVirtualNetworksClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewManagedVirtualNetworksClient creates a new instance of ManagedVirtualNetworksClient with the specified values.
// subscriptionID - The subscription identifier.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewManagedVirtualNetworksClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ManagedVirtualNetworksClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &ManagedVirtualNetworksClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// CreateOrUpdate - Creates or updates a managed Virtual Network.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The resource group name.
// factoryName - The factory name.
// managedVirtualNetworkName - Managed virtual network name
// managedVirtualNetwork - Managed Virtual Network resource definition.
// options - ManagedVirtualNetworksClientCreateOrUpdateOptions contains the optional parameters for the ManagedVirtualNetworksClient.CreateOrUpdate
// method.
func (client *ManagedVirtualNetworksClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, factoryName string, managedVirtualNetworkName string, managedVirtualNetwork ManagedVirtualNetworkResource, options *ManagedVirtualNetworksClientCreateOrUpdateOptions) (ManagedVirtualNetworksClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, factoryName, managedVirtualNetworkName, managedVirtualNetwork, options)
	if err != nil {
		return ManagedVirtualNetworksClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ManagedVirtualNetworksClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ManagedVirtualNetworksClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ManagedVirtualNetworksClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, factoryName string, managedVirtualNetworkName string, managedVirtualNetwork ManagedVirtualNetworkResource, options *ManagedVirtualNetworksClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/managedVirtualNetworks/{managedVirtualNetworkName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if factoryName == "" {
		return nil, errors.New("parameter factoryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{factoryName}", url.PathEscape(factoryName))
	if managedVirtualNetworkName == "" {
		return nil, errors.New("parameter managedVirtualNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedVirtualNetworkName}", url.PathEscape(managedVirtualNetworkName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header.Set("If-Match", *options.IfMatch)
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, managedVirtualNetwork)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ManagedVirtualNetworksClient) createOrUpdateHandleResponse(resp *http.Response) (ManagedVirtualNetworksClientCreateOrUpdateResponse, error) {
	result := ManagedVirtualNetworksClientCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedVirtualNetworkResource); err != nil {
		return ManagedVirtualNetworksClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Get - Gets a managed Virtual Network.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The resource group name.
// factoryName - The factory name.
// managedVirtualNetworkName - Managed virtual network name
// options - ManagedVirtualNetworksClientGetOptions contains the optional parameters for the ManagedVirtualNetworksClient.Get
// method.
func (client *ManagedVirtualNetworksClient) Get(ctx context.Context, resourceGroupName string, factoryName string, managedVirtualNetworkName string, options *ManagedVirtualNetworksClientGetOptions) (ManagedVirtualNetworksClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, factoryName, managedVirtualNetworkName, options)
	if err != nil {
		return ManagedVirtualNetworksClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ManagedVirtualNetworksClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ManagedVirtualNetworksClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ManagedVirtualNetworksClient) getCreateRequest(ctx context.Context, resourceGroupName string, factoryName string, managedVirtualNetworkName string, options *ManagedVirtualNetworksClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/managedVirtualNetworks/{managedVirtualNetworkName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if factoryName == "" {
		return nil, errors.New("parameter factoryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{factoryName}", url.PathEscape(factoryName))
	if managedVirtualNetworkName == "" {
		return nil, errors.New("parameter managedVirtualNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedVirtualNetworkName}", url.PathEscape(managedVirtualNetworkName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfNoneMatch != nil {
		req.Raw().Header.Set("If-None-Match", *options.IfNoneMatch)
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ManagedVirtualNetworksClient) getHandleResponse(resp *http.Response) (ManagedVirtualNetworksClientGetResponse, error) {
	result := ManagedVirtualNetworksClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedVirtualNetworkResource); err != nil {
		return ManagedVirtualNetworksClientGetResponse{}, err
	}
	return result, nil
}

// ListByFactory - Lists managed Virtual Networks.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The resource group name.
// factoryName - The factory name.
// options - ManagedVirtualNetworksClientListByFactoryOptions contains the optional parameters for the ManagedVirtualNetworksClient.ListByFactory
// method.
func (client *ManagedVirtualNetworksClient) ListByFactory(resourceGroupName string, factoryName string, options *ManagedVirtualNetworksClientListByFactoryOptions) *ManagedVirtualNetworksClientListByFactoryPager {
	return &ManagedVirtualNetworksClientListByFactoryPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByFactoryCreateRequest(ctx, resourceGroupName, factoryName, options)
		},
		advancer: func(ctx context.Context, resp ManagedVirtualNetworksClientListByFactoryResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ManagedVirtualNetworkListResponse.NextLink)
		},
	}
}

// listByFactoryCreateRequest creates the ListByFactory request.
func (client *ManagedVirtualNetworksClient) listByFactoryCreateRequest(ctx context.Context, resourceGroupName string, factoryName string, options *ManagedVirtualNetworksClientListByFactoryOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/managedVirtualNetworks"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if factoryName == "" {
		return nil, errors.New("parameter factoryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{factoryName}", url.PathEscape(factoryName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByFactoryHandleResponse handles the ListByFactory response.
func (client *ManagedVirtualNetworksClient) listByFactoryHandleResponse(resp *http.Response) (ManagedVirtualNetworksClientListByFactoryResponse, error) {
	result := ManagedVirtualNetworksClientListByFactoryResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedVirtualNetworkListResponse); err != nil {
		return ManagedVirtualNetworksClientListByFactoryResponse{}, err
	}
	return result, nil
}
