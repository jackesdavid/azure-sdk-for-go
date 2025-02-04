//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armtestbase

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

// AvailableOSClient contains the methods for the AvailableOS group.
// Don't use this type directly, use NewAvailableOSClient() instead.
type AvailableOSClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewAvailableOSClient creates a new instance of AvailableOSClient with the specified values.
// subscriptionID - The Azure subscription ID. This is a GUID-formatted string.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewAvailableOSClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *AvailableOSClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &AvailableOSClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// Get - Gets an available OS to run a package under a Test Base Account.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource.
// testBaseAccountName - The resource name of the Test Base Account.
// availableOSResourceName - The resource name of an Available OS.
// options - AvailableOSClientGetOptions contains the optional parameters for the AvailableOSClient.Get method.
func (client *AvailableOSClient) Get(ctx context.Context, resourceGroupName string, testBaseAccountName string, availableOSResourceName string, options *AvailableOSClientGetOptions) (AvailableOSClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, testBaseAccountName, availableOSResourceName, options)
	if err != nil {
		return AvailableOSClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AvailableOSClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AvailableOSClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *AvailableOSClient) getCreateRequest(ctx context.Context, resourceGroupName string, testBaseAccountName string, availableOSResourceName string, options *AvailableOSClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TestBase/testBaseAccounts/{testBaseAccountName}/availableOSs/{availableOSResourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if testBaseAccountName == "" {
		return nil, errors.New("parameter testBaseAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{testBaseAccountName}", url.PathEscape(testBaseAccountName))
	if availableOSResourceName == "" {
		return nil, errors.New("parameter availableOSResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{availableOSResourceName}", url.PathEscape(availableOSResourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-12-16-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AvailableOSClient) getHandleResponse(resp *http.Response) (AvailableOSClientGetResponse, error) {
	result := AvailableOSClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AvailableOSResource); err != nil {
		return AvailableOSClientGetResponse{}, err
	}
	return result, nil
}

// List - Lists all the available OSs to run a package under a Test Base Account.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource.
// testBaseAccountName - The resource name of the Test Base Account.
// osUpdateType - The type of the OS Update.
// options - AvailableOSClientListOptions contains the optional parameters for the AvailableOSClient.List method.
func (client *AvailableOSClient) List(resourceGroupName string, testBaseAccountName string, osUpdateType OsUpdateType, options *AvailableOSClientListOptions) *AvailableOSClientListPager {
	return &AvailableOSClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, testBaseAccountName, osUpdateType, options)
		},
		advancer: func(ctx context.Context, resp AvailableOSClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.AvailableOSListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *AvailableOSClient) listCreateRequest(ctx context.Context, resourceGroupName string, testBaseAccountName string, osUpdateType OsUpdateType, options *AvailableOSClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TestBase/testBaseAccounts/{testBaseAccountName}/availableOSs"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if testBaseAccountName == "" {
		return nil, errors.New("parameter testBaseAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{testBaseAccountName}", url.PathEscape(testBaseAccountName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("osUpdateType", string(osUpdateType))
	reqQP.Set("api-version", "2020-12-16-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *AvailableOSClient) listHandleResponse(resp *http.Response) (AvailableOSClientListResponse, error) {
	result := AvailableOSClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AvailableOSListResult); err != nil {
		return AvailableOSClientListResponse{}, err
	}
	return result, nil
}
