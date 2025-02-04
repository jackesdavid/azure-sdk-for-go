//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsynapse

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

// WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClient contains the methods for the WorkspaceManagedSQLServerExtendedBlobAuditingPolicies group.
// Don't use this type directly, use NewWorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClient() instead.
type WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewWorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClient creates a new instance of WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewWorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// BeginCreateOrUpdate - Create or Update a workspace managed sql server's extended blob auditing policy.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// blobAuditingPolicyName - The name of the blob auditing policy.
// parameters - Properties of extended blob auditing policy.
// options - WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClientBeginCreateOrUpdateOptions contains the optional parameters
// for the WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClient.BeginCreateOrUpdate method.
func (client *WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, blobAuditingPolicyName BlobAuditingPolicyName, parameters ExtendedServerBlobAuditingPolicy, options *WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClientBeginCreateOrUpdateOptions) (WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClientCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, workspaceName, blobAuditingPolicyName, parameters, options)
	if err != nil {
		return WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClientCreateOrUpdatePollerResponse{}, err
	}
	result := WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClientCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClient.CreateOrUpdate", "", resp, client.pl)
	if err != nil {
		return WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClientCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClientCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Create or Update a workspace managed sql server's extended blob auditing policy.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClient) createOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, blobAuditingPolicyName BlobAuditingPolicyName, parameters ExtendedServerBlobAuditingPolicy, options *WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, workspaceName, blobAuditingPolicyName, parameters, options)
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

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, blobAuditingPolicyName BlobAuditingPolicyName, parameters ExtendedServerBlobAuditingPolicy, options *WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/extendedAuditingSettings/{blobAuditingPolicyName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if blobAuditingPolicyName == "" {
		return nil, errors.New("parameter blobAuditingPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{blobAuditingPolicyName}", url.PathEscape(string(blobAuditingPolicyName)))
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

// Get - Get a workspace SQL server's extended blob auditing policy.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// blobAuditingPolicyName - The name of the blob auditing policy.
// options - WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClientGetOptions contains the optional parameters for the
// WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClient.Get method.
func (client *WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, blobAuditingPolicyName BlobAuditingPolicyName, options *WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClientGetOptions) (WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, blobAuditingPolicyName, options)
	if err != nil {
		return WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, blobAuditingPolicyName BlobAuditingPolicyName, options *WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/extendedAuditingSettings/{blobAuditingPolicyName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if blobAuditingPolicyName == "" {
		return nil, errors.New("parameter blobAuditingPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{blobAuditingPolicyName}", url.PathEscape(string(blobAuditingPolicyName)))
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
func (client *WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClient) getHandleResponse(resp *http.Response) (WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClientGetResponse, error) {
	result := WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExtendedServerBlobAuditingPolicy); err != nil {
		return WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClientGetResponse{}, err
	}
	return result, nil
}

// ListByWorkspace - List workspace managed sql server's extended blob auditing policies.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// options - WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClientListByWorkspaceOptions contains the optional parameters
// for the WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClient.ListByWorkspace method.
func (client *WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClient) ListByWorkspace(resourceGroupName string, workspaceName string, options *WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClientListByWorkspaceOptions) *WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClientListByWorkspacePager {
	return &WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClientListByWorkspacePager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByWorkspaceCreateRequest(ctx, resourceGroupName, workspaceName, options)
		},
		advancer: func(ctx context.Context, resp WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClientListByWorkspaceResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ExtendedServerBlobAuditingPolicyListResult.NextLink)
		},
	}
}

// listByWorkspaceCreateRequest creates the ListByWorkspace request.
func (client *WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClient) listByWorkspaceCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, options *WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClientListByWorkspaceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/extendedAuditingSettings"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
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

// listByWorkspaceHandleResponse handles the ListByWorkspace response.
func (client *WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClient) listByWorkspaceHandleResponse(resp *http.Response) (WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClientListByWorkspaceResponse, error) {
	result := WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClientListByWorkspaceResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExtendedServerBlobAuditingPolicyListResult); err != nil {
		return WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClientListByWorkspaceResponse{}, err
	}
	return result, nil
}
