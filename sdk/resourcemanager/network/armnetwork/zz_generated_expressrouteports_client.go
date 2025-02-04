//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

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

// ExpressRoutePortsClient contains the methods for the ExpressRoutePorts group.
// Don't use this type directly, use NewExpressRoutePortsClient() instead.
type ExpressRoutePortsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewExpressRoutePortsClient creates a new instance of ExpressRoutePortsClient with the specified values.
// subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
// ID forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewExpressRoutePortsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ExpressRoutePortsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &ExpressRoutePortsClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// BeginCreateOrUpdate - Creates or updates the specified ExpressRoutePort resource.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// expressRoutePortName - The name of the ExpressRoutePort resource.
// parameters - Parameters supplied to the create ExpressRoutePort operation.
// options - ExpressRoutePortsClientBeginCreateOrUpdateOptions contains the optional parameters for the ExpressRoutePortsClient.BeginCreateOrUpdate
// method.
func (client *ExpressRoutePortsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, expressRoutePortName string, parameters ExpressRoutePort, options *ExpressRoutePortsClientBeginCreateOrUpdateOptions) (ExpressRoutePortsClientCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, expressRoutePortName, parameters, options)
	if err != nil {
		return ExpressRoutePortsClientCreateOrUpdatePollerResponse{}, err
	}
	result := ExpressRoutePortsClientCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ExpressRoutePortsClient.CreateOrUpdate", "azure-async-operation", resp, client.pl)
	if err != nil {
		return ExpressRoutePortsClientCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &ExpressRoutePortsClientCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates the specified ExpressRoutePort resource.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ExpressRoutePortsClient) createOrUpdate(ctx context.Context, resourceGroupName string, expressRoutePortName string, parameters ExpressRoutePort, options *ExpressRoutePortsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, expressRoutePortName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ExpressRoutePortsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, expressRoutePortName string, parameters ExpressRoutePort, options *ExpressRoutePortsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ExpressRoutePorts/{expressRoutePortName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if expressRoutePortName == "" {
		return nil, errors.New("parameter expressRoutePortName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{expressRoutePortName}", url.PathEscape(expressRoutePortName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes the specified ExpressRoutePort resource.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// expressRoutePortName - The name of the ExpressRoutePort resource.
// options - ExpressRoutePortsClientBeginDeleteOptions contains the optional parameters for the ExpressRoutePortsClient.BeginDelete
// method.
func (client *ExpressRoutePortsClient) BeginDelete(ctx context.Context, resourceGroupName string, expressRoutePortName string, options *ExpressRoutePortsClientBeginDeleteOptions) (ExpressRoutePortsClientDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, expressRoutePortName, options)
	if err != nil {
		return ExpressRoutePortsClientDeletePollerResponse{}, err
	}
	result := ExpressRoutePortsClientDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ExpressRoutePortsClient.Delete", "location", resp, client.pl)
	if err != nil {
		return ExpressRoutePortsClientDeletePollerResponse{}, err
	}
	result.Poller = &ExpressRoutePortsClientDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes the specified ExpressRoutePort resource.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ExpressRoutePortsClient) deleteOperation(ctx context.Context, resourceGroupName string, expressRoutePortName string, options *ExpressRoutePortsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, expressRoutePortName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ExpressRoutePortsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, expressRoutePortName string, options *ExpressRoutePortsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ExpressRoutePorts/{expressRoutePortName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if expressRoutePortName == "" {
		return nil, errors.New("parameter expressRoutePortName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{expressRoutePortName}", url.PathEscape(expressRoutePortName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// GenerateLOA - Generate a letter of authorization for the requested ExpressRoutePort resource.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// expressRoutePortName - The name of ExpressRoutePort.
// request - Request parameters supplied to generate a letter of authorization.
// options - ExpressRoutePortsClientGenerateLOAOptions contains the optional parameters for the ExpressRoutePortsClient.GenerateLOA
// method.
func (client *ExpressRoutePortsClient) GenerateLOA(ctx context.Context, resourceGroupName string, expressRoutePortName string, request GenerateExpressRoutePortsLOARequest, options *ExpressRoutePortsClientGenerateLOAOptions) (ExpressRoutePortsClientGenerateLOAResponse, error) {
	req, err := client.generateLOACreateRequest(ctx, resourceGroupName, expressRoutePortName, request, options)
	if err != nil {
		return ExpressRoutePortsClientGenerateLOAResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ExpressRoutePortsClientGenerateLOAResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ExpressRoutePortsClientGenerateLOAResponse{}, runtime.NewResponseError(resp)
	}
	return client.generateLOAHandleResponse(resp)
}

// generateLOACreateRequest creates the GenerateLOA request.
func (client *ExpressRoutePortsClient) generateLOACreateRequest(ctx context.Context, resourceGroupName string, expressRoutePortName string, request GenerateExpressRoutePortsLOARequest, options *ExpressRoutePortsClientGenerateLOAOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRoutePorts/{expressRoutePortName}/generateLoa"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if expressRoutePortName == "" {
		return nil, errors.New("parameter expressRoutePortName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{expressRoutePortName}", url.PathEscape(expressRoutePortName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, request)
}

// generateLOAHandleResponse handles the GenerateLOA response.
func (client *ExpressRoutePortsClient) generateLOAHandleResponse(resp *http.Response) (ExpressRoutePortsClientGenerateLOAResponse, error) {
	result := ExpressRoutePortsClientGenerateLOAResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.GenerateExpressRoutePortsLOAResult); err != nil {
		return ExpressRoutePortsClientGenerateLOAResponse{}, err
	}
	return result, nil
}

// Get - Retrieves the requested ExpressRoutePort resource.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// expressRoutePortName - The name of ExpressRoutePort.
// options - ExpressRoutePortsClientGetOptions contains the optional parameters for the ExpressRoutePortsClient.Get method.
func (client *ExpressRoutePortsClient) Get(ctx context.Context, resourceGroupName string, expressRoutePortName string, options *ExpressRoutePortsClientGetOptions) (ExpressRoutePortsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, expressRoutePortName, options)
	if err != nil {
		return ExpressRoutePortsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ExpressRoutePortsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ExpressRoutePortsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ExpressRoutePortsClient) getCreateRequest(ctx context.Context, resourceGroupName string, expressRoutePortName string, options *ExpressRoutePortsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ExpressRoutePorts/{expressRoutePortName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if expressRoutePortName == "" {
		return nil, errors.New("parameter expressRoutePortName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{expressRoutePortName}", url.PathEscape(expressRoutePortName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ExpressRoutePortsClient) getHandleResponse(resp *http.Response) (ExpressRoutePortsClientGetResponse, error) {
	result := ExpressRoutePortsClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExpressRoutePort); err != nil {
		return ExpressRoutePortsClientGetResponse{}, err
	}
	return result, nil
}

// List - List all the ExpressRoutePort resources in the specified subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// options - ExpressRoutePortsClientListOptions contains the optional parameters for the ExpressRoutePortsClient.List method.
func (client *ExpressRoutePortsClient) List(options *ExpressRoutePortsClientListOptions) *ExpressRoutePortsClientListPager {
	return &ExpressRoutePortsClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp ExpressRoutePortsClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ExpressRoutePortListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *ExpressRoutePortsClient) listCreateRequest(ctx context.Context, options *ExpressRoutePortsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/ExpressRoutePorts"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ExpressRoutePortsClient) listHandleResponse(resp *http.Response) (ExpressRoutePortsClientListResponse, error) {
	result := ExpressRoutePortsClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExpressRoutePortListResult); err != nil {
		return ExpressRoutePortsClientListResponse{}, err
	}
	return result, nil
}

// ListByResourceGroup - List all the ExpressRoutePort resources in the specified resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// options - ExpressRoutePortsClientListByResourceGroupOptions contains the optional parameters for the ExpressRoutePortsClient.ListByResourceGroup
// method.
func (client *ExpressRoutePortsClient) ListByResourceGroup(resourceGroupName string, options *ExpressRoutePortsClientListByResourceGroupOptions) *ExpressRoutePortsClientListByResourceGroupPager {
	return &ExpressRoutePortsClientListByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp ExpressRoutePortsClientListByResourceGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ExpressRoutePortListResult.NextLink)
		},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *ExpressRoutePortsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *ExpressRoutePortsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ExpressRoutePorts"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *ExpressRoutePortsClient) listByResourceGroupHandleResponse(resp *http.Response) (ExpressRoutePortsClientListByResourceGroupResponse, error) {
	result := ExpressRoutePortsClientListByResourceGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExpressRoutePortListResult); err != nil {
		return ExpressRoutePortsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// UpdateTags - Update ExpressRoutePort tags.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// expressRoutePortName - The name of the ExpressRoutePort resource.
// parameters - Parameters supplied to update ExpressRoutePort resource tags.
// options - ExpressRoutePortsClientUpdateTagsOptions contains the optional parameters for the ExpressRoutePortsClient.UpdateTags
// method.
func (client *ExpressRoutePortsClient) UpdateTags(ctx context.Context, resourceGroupName string, expressRoutePortName string, parameters TagsObject, options *ExpressRoutePortsClientUpdateTagsOptions) (ExpressRoutePortsClientUpdateTagsResponse, error) {
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, expressRoutePortName, parameters, options)
	if err != nil {
		return ExpressRoutePortsClientUpdateTagsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ExpressRoutePortsClientUpdateTagsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ExpressRoutePortsClientUpdateTagsResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateTagsHandleResponse(resp)
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *ExpressRoutePortsClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, expressRoutePortName string, parameters TagsObject, options *ExpressRoutePortsClientUpdateTagsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ExpressRoutePorts/{expressRoutePortName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if expressRoutePortName == "" {
		return nil, errors.New("parameter expressRoutePortName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{expressRoutePortName}", url.PathEscape(expressRoutePortName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateTagsHandleResponse handles the UpdateTags response.
func (client *ExpressRoutePortsClient) updateTagsHandleResponse(resp *http.Response) (ExpressRoutePortsClientUpdateTagsResponse, error) {
	result := ExpressRoutePortsClientUpdateTagsResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExpressRoutePort); err != nil {
		return ExpressRoutePortsClientUpdateTagsResponse{}, err
	}
	return result, nil
}
