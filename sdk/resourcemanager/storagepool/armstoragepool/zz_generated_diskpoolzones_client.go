//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstoragepool

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

// DiskPoolZonesClient contains the methods for the DiskPoolZones group.
// Don't use this type directly, use NewDiskPoolZonesClient() instead.
type DiskPoolZonesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewDiskPoolZonesClient creates a new instance of DiskPoolZonesClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewDiskPoolZonesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *DiskPoolZonesClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &DiskPoolZonesClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// List - Lists available Disk Pool Skus in an Azure location.
// If the operation fails it returns an *azcore.ResponseError type.
// location - The location of the resource.
// options - DiskPoolZonesClientListOptions contains the optional parameters for the DiskPoolZonesClient.List method.
func (client *DiskPoolZonesClient) List(location string, options *DiskPoolZonesClientListOptions) *DiskPoolZonesClientListPager {
	return &DiskPoolZonesClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, location, options)
		},
		advancer: func(ctx context.Context, resp DiskPoolZonesClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.DiskPoolZoneListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *DiskPoolZonesClient) listCreateRequest(ctx context.Context, location string, options *DiskPoolZonesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.StoragePool/locations/{location}/diskPoolZones"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *DiskPoolZonesClient) listHandleResponse(resp *http.Response) (DiskPoolZonesClientListResponse, error) {
	result := DiskPoolZonesClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DiskPoolZoneListResult); err != nil {
		return DiskPoolZonesClientListResponse{}, err
	}
	return result, nil
}
