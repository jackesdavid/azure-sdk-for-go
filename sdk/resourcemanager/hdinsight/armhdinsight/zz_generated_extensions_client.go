//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhdinsight

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

// ExtensionsClient contains the methods for the Extensions group.
// Don't use this type directly, use NewExtensionsClient() instead.
type ExtensionsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewExtensionsClient creates a new instance of ExtensionsClient with the specified values.
// subscriptionID - The subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID
// forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewExtensionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ExtensionsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &ExtensionsClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// BeginCreate - Creates an HDInsight cluster extension.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// clusterName - The name of the cluster.
// extensionName - The name of the cluster extension.
// parameters - The cluster extensions create request.
// options - ExtensionsClientBeginCreateOptions contains the optional parameters for the ExtensionsClient.BeginCreate method.
func (client *ExtensionsClient) BeginCreate(ctx context.Context, resourceGroupName string, clusterName string, extensionName string, parameters Extension, options *ExtensionsClientBeginCreateOptions) (ExtensionsClientCreatePollerResponse, error) {
	resp, err := client.create(ctx, resourceGroupName, clusterName, extensionName, parameters, options)
	if err != nil {
		return ExtensionsClientCreatePollerResponse{}, err
	}
	result := ExtensionsClientCreatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ExtensionsClient.Create", "location", resp, client.pl)
	if err != nil {
		return ExtensionsClientCreatePollerResponse{}, err
	}
	result.Poller = &ExtensionsClientCreatePoller{
		pt: pt,
	}
	return result, nil
}

// Create - Creates an HDInsight cluster extension.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ExtensionsClient) create(ctx context.Context, resourceGroupName string, clusterName string, extensionName string, parameters Extension, options *ExtensionsClientBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, clusterName, extensionName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createCreateRequest creates the Create request.
func (client *ExtensionsClient) createCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, extensionName string, parameters Extension, options *ExtensionsClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}/extensions/{extensionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if extensionName == "" {
		return nil, errors.New("parameter extensionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{extensionName}", url.PathEscape(extensionName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes the specified extension for HDInsight cluster.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// clusterName - The name of the cluster.
// extensionName - The name of the cluster extension.
// options - ExtensionsClientBeginDeleteOptions contains the optional parameters for the ExtensionsClient.BeginDelete method.
func (client *ExtensionsClient) BeginDelete(ctx context.Context, resourceGroupName string, clusterName string, extensionName string, options *ExtensionsClientBeginDeleteOptions) (ExtensionsClientDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, clusterName, extensionName, options)
	if err != nil {
		return ExtensionsClientDeletePollerResponse{}, err
	}
	result := ExtensionsClientDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ExtensionsClient.Delete", "location", resp, client.pl)
	if err != nil {
		return ExtensionsClientDeletePollerResponse{}, err
	}
	result.Poller = &ExtensionsClientDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes the specified extension for HDInsight cluster.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ExtensionsClient) deleteOperation(ctx context.Context, resourceGroupName string, clusterName string, extensionName string, options *ExtensionsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, clusterName, extensionName, options)
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
func (client *ExtensionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, extensionName string, options *ExtensionsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}/extensions/{extensionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if extensionName == "" {
		return nil, errors.New("parameter extensionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{extensionName}", url.PathEscape(extensionName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// BeginDisableAzureMonitor - Disables the Azure Monitor on the HDInsight cluster.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// clusterName - The name of the cluster.
// options - ExtensionsClientBeginDisableAzureMonitorOptions contains the optional parameters for the ExtensionsClient.BeginDisableAzureMonitor
// method.
func (client *ExtensionsClient) BeginDisableAzureMonitor(ctx context.Context, resourceGroupName string, clusterName string, options *ExtensionsClientBeginDisableAzureMonitorOptions) (ExtensionsClientDisableAzureMonitorPollerResponse, error) {
	resp, err := client.disableAzureMonitor(ctx, resourceGroupName, clusterName, options)
	if err != nil {
		return ExtensionsClientDisableAzureMonitorPollerResponse{}, err
	}
	result := ExtensionsClientDisableAzureMonitorPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ExtensionsClient.DisableAzureMonitor", "location", resp, client.pl)
	if err != nil {
		return ExtensionsClientDisableAzureMonitorPollerResponse{}, err
	}
	result.Poller = &ExtensionsClientDisableAzureMonitorPoller{
		pt: pt,
	}
	return result, nil
}

// DisableAzureMonitor - Disables the Azure Monitor on the HDInsight cluster.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ExtensionsClient) disableAzureMonitor(ctx context.Context, resourceGroupName string, clusterName string, options *ExtensionsClientBeginDisableAzureMonitorOptions) (*http.Response, error) {
	req, err := client.disableAzureMonitorCreateRequest(ctx, resourceGroupName, clusterName, options)
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

// disableAzureMonitorCreateRequest creates the DisableAzureMonitor request.
func (client *ExtensionsClient) disableAzureMonitorCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, options *ExtensionsClientBeginDisableAzureMonitorOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}/extensions/azureMonitor"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// BeginDisableMonitoring - Disables the Operations Management Suite (OMS) on the HDInsight cluster.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// clusterName - The name of the cluster.
// options - ExtensionsClientBeginDisableMonitoringOptions contains the optional parameters for the ExtensionsClient.BeginDisableMonitoring
// method.
func (client *ExtensionsClient) BeginDisableMonitoring(ctx context.Context, resourceGroupName string, clusterName string, options *ExtensionsClientBeginDisableMonitoringOptions) (ExtensionsClientDisableMonitoringPollerResponse, error) {
	resp, err := client.disableMonitoring(ctx, resourceGroupName, clusterName, options)
	if err != nil {
		return ExtensionsClientDisableMonitoringPollerResponse{}, err
	}
	result := ExtensionsClientDisableMonitoringPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ExtensionsClient.DisableMonitoring", "location", resp, client.pl)
	if err != nil {
		return ExtensionsClientDisableMonitoringPollerResponse{}, err
	}
	result.Poller = &ExtensionsClientDisableMonitoringPoller{
		pt: pt,
	}
	return result, nil
}

// DisableMonitoring - Disables the Operations Management Suite (OMS) on the HDInsight cluster.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ExtensionsClient) disableMonitoring(ctx context.Context, resourceGroupName string, clusterName string, options *ExtensionsClientBeginDisableMonitoringOptions) (*http.Response, error) {
	req, err := client.disableMonitoringCreateRequest(ctx, resourceGroupName, clusterName, options)
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

// disableMonitoringCreateRequest creates the DisableMonitoring request.
func (client *ExtensionsClient) disableMonitoringCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, options *ExtensionsClientBeginDisableMonitoringOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}/extensions/clustermonitoring"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// BeginEnableAzureMonitor - Enables the Azure Monitor on the HDInsight cluster.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// clusterName - The name of the cluster.
// parameters - The Log Analytics workspace parameters.
// options - ExtensionsClientBeginEnableAzureMonitorOptions contains the optional parameters for the ExtensionsClient.BeginEnableAzureMonitor
// method.
func (client *ExtensionsClient) BeginEnableAzureMonitor(ctx context.Context, resourceGroupName string, clusterName string, parameters AzureMonitorRequest, options *ExtensionsClientBeginEnableAzureMonitorOptions) (ExtensionsClientEnableAzureMonitorPollerResponse, error) {
	resp, err := client.enableAzureMonitor(ctx, resourceGroupName, clusterName, parameters, options)
	if err != nil {
		return ExtensionsClientEnableAzureMonitorPollerResponse{}, err
	}
	result := ExtensionsClientEnableAzureMonitorPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ExtensionsClient.EnableAzureMonitor", "location", resp, client.pl)
	if err != nil {
		return ExtensionsClientEnableAzureMonitorPollerResponse{}, err
	}
	result.Poller = &ExtensionsClientEnableAzureMonitorPoller{
		pt: pt,
	}
	return result, nil
}

// EnableAzureMonitor - Enables the Azure Monitor on the HDInsight cluster.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ExtensionsClient) enableAzureMonitor(ctx context.Context, resourceGroupName string, clusterName string, parameters AzureMonitorRequest, options *ExtensionsClientBeginEnableAzureMonitorOptions) (*http.Response, error) {
	req, err := client.enableAzureMonitorCreateRequest(ctx, resourceGroupName, clusterName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// enableAzureMonitorCreateRequest creates the EnableAzureMonitor request.
func (client *ExtensionsClient) enableAzureMonitorCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, parameters AzureMonitorRequest, options *ExtensionsClientBeginEnableAzureMonitorOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}/extensions/azureMonitor"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginEnableMonitoring - Enables the Operations Management Suite (OMS) on the HDInsight cluster.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// clusterName - The name of the cluster.
// parameters - The Operations Management Suite (OMS) workspace parameters.
// options - ExtensionsClientBeginEnableMonitoringOptions contains the optional parameters for the ExtensionsClient.BeginEnableMonitoring
// method.
func (client *ExtensionsClient) BeginEnableMonitoring(ctx context.Context, resourceGroupName string, clusterName string, parameters ClusterMonitoringRequest, options *ExtensionsClientBeginEnableMonitoringOptions) (ExtensionsClientEnableMonitoringPollerResponse, error) {
	resp, err := client.enableMonitoring(ctx, resourceGroupName, clusterName, parameters, options)
	if err != nil {
		return ExtensionsClientEnableMonitoringPollerResponse{}, err
	}
	result := ExtensionsClientEnableMonitoringPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ExtensionsClient.EnableMonitoring", "location", resp, client.pl)
	if err != nil {
		return ExtensionsClientEnableMonitoringPollerResponse{}, err
	}
	result.Poller = &ExtensionsClientEnableMonitoringPoller{
		pt: pt,
	}
	return result, nil
}

// EnableMonitoring - Enables the Operations Management Suite (OMS) on the HDInsight cluster.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ExtensionsClient) enableMonitoring(ctx context.Context, resourceGroupName string, clusterName string, parameters ClusterMonitoringRequest, options *ExtensionsClientBeginEnableMonitoringOptions) (*http.Response, error) {
	req, err := client.enableMonitoringCreateRequest(ctx, resourceGroupName, clusterName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// enableMonitoringCreateRequest creates the EnableMonitoring request.
func (client *ExtensionsClient) enableMonitoringCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, parameters ClusterMonitoringRequest, options *ExtensionsClientBeginEnableMonitoringOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}/extensions/clustermonitoring"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// Get - Gets the extension properties for the specified HDInsight cluster extension.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// clusterName - The name of the cluster.
// extensionName - The name of the cluster extension.
// options - ExtensionsClientGetOptions contains the optional parameters for the ExtensionsClient.Get method.
func (client *ExtensionsClient) Get(ctx context.Context, resourceGroupName string, clusterName string, extensionName string, options *ExtensionsClientGetOptions) (ExtensionsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, clusterName, extensionName, options)
	if err != nil {
		return ExtensionsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ExtensionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ExtensionsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ExtensionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, extensionName string, options *ExtensionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}/extensions/{extensionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if extensionName == "" {
		return nil, errors.New("parameter extensionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{extensionName}", url.PathEscape(extensionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ExtensionsClient) getHandleResponse(resp *http.Response) (ExtensionsClientGetResponse, error) {
	result := ExtensionsClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ClusterMonitoringResponse); err != nil {
		return ExtensionsClientGetResponse{}, err
	}
	return result, nil
}

// GetAzureAsyncOperationStatus - Gets the async operation status.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// clusterName - The name of the cluster.
// extensionName - The name of the cluster extension.
// operationID - The long running operation id.
// options - ExtensionsClientGetAzureAsyncOperationStatusOptions contains the optional parameters for the ExtensionsClient.GetAzureAsyncOperationStatus
// method.
func (client *ExtensionsClient) GetAzureAsyncOperationStatus(ctx context.Context, resourceGroupName string, clusterName string, extensionName string, operationID string, options *ExtensionsClientGetAzureAsyncOperationStatusOptions) (ExtensionsClientGetAzureAsyncOperationStatusResponse, error) {
	req, err := client.getAzureAsyncOperationStatusCreateRequest(ctx, resourceGroupName, clusterName, extensionName, operationID, options)
	if err != nil {
		return ExtensionsClientGetAzureAsyncOperationStatusResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ExtensionsClientGetAzureAsyncOperationStatusResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ExtensionsClientGetAzureAsyncOperationStatusResponse{}, runtime.NewResponseError(resp)
	}
	return client.getAzureAsyncOperationStatusHandleResponse(resp)
}

// getAzureAsyncOperationStatusCreateRequest creates the GetAzureAsyncOperationStatus request.
func (client *ExtensionsClient) getAzureAsyncOperationStatusCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, extensionName string, operationID string, options *ExtensionsClientGetAzureAsyncOperationStatusOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}/extensions/{extensionName}/azureAsyncOperations/{operationId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if extensionName == "" {
		return nil, errors.New("parameter extensionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{extensionName}", url.PathEscape(extensionName))
	if operationID == "" {
		return nil, errors.New("parameter operationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{operationId}", url.PathEscape(operationID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getAzureAsyncOperationStatusHandleResponse handles the GetAzureAsyncOperationStatus response.
func (client *ExtensionsClient) getAzureAsyncOperationStatusHandleResponse(resp *http.Response) (ExtensionsClientGetAzureAsyncOperationStatusResponse, error) {
	result := ExtensionsClientGetAzureAsyncOperationStatusResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AsyncOperationResult); err != nil {
		return ExtensionsClientGetAzureAsyncOperationStatusResponse{}, err
	}
	return result, nil
}

// GetAzureMonitorStatus - Gets the status of Azure Monitor on the HDInsight cluster.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// clusterName - The name of the cluster.
// options - ExtensionsClientGetAzureMonitorStatusOptions contains the optional parameters for the ExtensionsClient.GetAzureMonitorStatus
// method.
func (client *ExtensionsClient) GetAzureMonitorStatus(ctx context.Context, resourceGroupName string, clusterName string, options *ExtensionsClientGetAzureMonitorStatusOptions) (ExtensionsClientGetAzureMonitorStatusResponse, error) {
	req, err := client.getAzureMonitorStatusCreateRequest(ctx, resourceGroupName, clusterName, options)
	if err != nil {
		return ExtensionsClientGetAzureMonitorStatusResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ExtensionsClientGetAzureMonitorStatusResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ExtensionsClientGetAzureMonitorStatusResponse{}, runtime.NewResponseError(resp)
	}
	return client.getAzureMonitorStatusHandleResponse(resp)
}

// getAzureMonitorStatusCreateRequest creates the GetAzureMonitorStatus request.
func (client *ExtensionsClient) getAzureMonitorStatusCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, options *ExtensionsClientGetAzureMonitorStatusOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}/extensions/azureMonitor"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getAzureMonitorStatusHandleResponse handles the GetAzureMonitorStatus response.
func (client *ExtensionsClient) getAzureMonitorStatusHandleResponse(resp *http.Response) (ExtensionsClientGetAzureMonitorStatusResponse, error) {
	result := ExtensionsClientGetAzureMonitorStatusResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AzureMonitorResponse); err != nil {
		return ExtensionsClientGetAzureMonitorStatusResponse{}, err
	}
	return result, nil
}

// GetMonitoringStatus - Gets the status of Operations Management Suite (OMS) on the HDInsight cluster.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// clusterName - The name of the cluster.
// options - ExtensionsClientGetMonitoringStatusOptions contains the optional parameters for the ExtensionsClient.GetMonitoringStatus
// method.
func (client *ExtensionsClient) GetMonitoringStatus(ctx context.Context, resourceGroupName string, clusterName string, options *ExtensionsClientGetMonitoringStatusOptions) (ExtensionsClientGetMonitoringStatusResponse, error) {
	req, err := client.getMonitoringStatusCreateRequest(ctx, resourceGroupName, clusterName, options)
	if err != nil {
		return ExtensionsClientGetMonitoringStatusResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ExtensionsClientGetMonitoringStatusResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ExtensionsClientGetMonitoringStatusResponse{}, runtime.NewResponseError(resp)
	}
	return client.getMonitoringStatusHandleResponse(resp)
}

// getMonitoringStatusCreateRequest creates the GetMonitoringStatus request.
func (client *ExtensionsClient) getMonitoringStatusCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, options *ExtensionsClientGetMonitoringStatusOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}/extensions/clustermonitoring"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getMonitoringStatusHandleResponse handles the GetMonitoringStatus response.
func (client *ExtensionsClient) getMonitoringStatusHandleResponse(resp *http.Response) (ExtensionsClientGetMonitoringStatusResponse, error) {
	result := ExtensionsClientGetMonitoringStatusResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ClusterMonitoringResponse); err != nil {
		return ExtensionsClientGetMonitoringStatusResponse{}, err
	}
	return result, nil
}
