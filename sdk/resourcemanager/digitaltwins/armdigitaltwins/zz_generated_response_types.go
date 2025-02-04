//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdigitaltwins

import (
	"context"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"net/http"
	"time"
)

// ClientCheckNameAvailabilityResponse contains the response from method Client.CheckNameAvailability.
type ClientCheckNameAvailabilityResponse struct {
	ClientCheckNameAvailabilityResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ClientCheckNameAvailabilityResult contains the result from method Client.CheckNameAvailability.
type ClientCheckNameAvailabilityResult struct {
	CheckNameResult
}

// ClientCreateOrUpdatePollerResponse contains the response from method Client.CreateOrUpdate.
type ClientCreateOrUpdatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *ClientCreateOrUpdatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l ClientCreateOrUpdatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (ClientCreateOrUpdateResponse, error) {
	respType := ClientCreateOrUpdateResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.Description)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a ClientCreateOrUpdatePollerResponse from the provided client and resume token.
func (l *ClientCreateOrUpdatePollerResponse) Resume(ctx context.Context, client *Client, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("Client.CreateOrUpdate", token, client.pl)
	if err != nil {
		return err
	}
	poller := &ClientCreateOrUpdatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// ClientCreateOrUpdateResponse contains the response from method Client.CreateOrUpdate.
type ClientCreateOrUpdateResponse struct {
	ClientCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ClientCreateOrUpdateResult contains the result from method Client.CreateOrUpdate.
type ClientCreateOrUpdateResult struct {
	Description
}

// ClientDeletePollerResponse contains the response from method Client.Delete.
type ClientDeletePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *ClientDeletePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l ClientDeletePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (ClientDeleteResponse, error) {
	respType := ClientDeleteResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.Description)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a ClientDeletePollerResponse from the provided client and resume token.
func (l *ClientDeletePollerResponse) Resume(ctx context.Context, client *Client, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("Client.Delete", token, client.pl)
	if err != nil {
		return err
	}
	poller := &ClientDeletePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// ClientDeleteResponse contains the response from method Client.Delete.
type ClientDeleteResponse struct {
	ClientDeleteResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ClientDeleteResult contains the result from method Client.Delete.
type ClientDeleteResult struct {
	Description
}

// ClientGetResponse contains the response from method Client.Get.
type ClientGetResponse struct {
	ClientGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ClientGetResult contains the result from method Client.Get.
type ClientGetResult struct {
	Description
}

// ClientListByResourceGroupResponse contains the response from method Client.ListByResourceGroup.
type ClientListByResourceGroupResponse struct {
	ClientListByResourceGroupResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ClientListByResourceGroupResult contains the result from method Client.ListByResourceGroup.
type ClientListByResourceGroupResult struct {
	DescriptionListResult
}

// ClientListResponse contains the response from method Client.List.
type ClientListResponse struct {
	ClientListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ClientListResult contains the result from method Client.List.
type ClientListResult struct {
	DescriptionListResult
}

// ClientUpdatePollerResponse contains the response from method Client.Update.
type ClientUpdatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *ClientUpdatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l ClientUpdatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (ClientUpdateResponse, error) {
	respType := ClientUpdateResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.Description)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a ClientUpdatePollerResponse from the provided client and resume token.
func (l *ClientUpdatePollerResponse) Resume(ctx context.Context, client *Client, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("Client.Update", token, client.pl)
	if err != nil {
		return err
	}
	poller := &ClientUpdatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// ClientUpdateResponse contains the response from method Client.Update.
type ClientUpdateResponse struct {
	ClientUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ClientUpdateResult contains the result from method Client.Update.
type ClientUpdateResult struct {
	Description
}

// EndpointClientCreateOrUpdatePollerResponse contains the response from method EndpointClient.CreateOrUpdate.
type EndpointClientCreateOrUpdatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *EndpointClientCreateOrUpdatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l EndpointClientCreateOrUpdatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (EndpointClientCreateOrUpdateResponse, error) {
	respType := EndpointClientCreateOrUpdateResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.EndpointResource)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a EndpointClientCreateOrUpdatePollerResponse from the provided client and resume token.
func (l *EndpointClientCreateOrUpdatePollerResponse) Resume(ctx context.Context, client *EndpointClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("EndpointClient.CreateOrUpdate", token, client.pl)
	if err != nil {
		return err
	}
	poller := &EndpointClientCreateOrUpdatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// EndpointClientCreateOrUpdateResponse contains the response from method EndpointClient.CreateOrUpdate.
type EndpointClientCreateOrUpdateResponse struct {
	EndpointClientCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// EndpointClientCreateOrUpdateResult contains the result from method EndpointClient.CreateOrUpdate.
type EndpointClientCreateOrUpdateResult struct {
	EndpointResource
}

// EndpointClientDeletePollerResponse contains the response from method EndpointClient.Delete.
type EndpointClientDeletePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *EndpointClientDeletePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l EndpointClientDeletePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (EndpointClientDeleteResponse, error) {
	respType := EndpointClientDeleteResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.EndpointResource)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a EndpointClientDeletePollerResponse from the provided client and resume token.
func (l *EndpointClientDeletePollerResponse) Resume(ctx context.Context, client *EndpointClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("EndpointClient.Delete", token, client.pl)
	if err != nil {
		return err
	}
	poller := &EndpointClientDeletePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// EndpointClientDeleteResponse contains the response from method EndpointClient.Delete.
type EndpointClientDeleteResponse struct {
	EndpointClientDeleteResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// EndpointClientDeleteResult contains the result from method EndpointClient.Delete.
type EndpointClientDeleteResult struct {
	EndpointResource
}

// EndpointClientGetResponse contains the response from method EndpointClient.Get.
type EndpointClientGetResponse struct {
	EndpointClientGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// EndpointClientGetResult contains the result from method EndpointClient.Get.
type EndpointClientGetResult struct {
	EndpointResource
}

// EndpointClientListResponse contains the response from method EndpointClient.List.
type EndpointClientListResponse struct {
	EndpointClientListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// EndpointClientListResult contains the result from method EndpointClient.List.
type EndpointClientListResult struct {
	EndpointResourceListResult
}

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	OperationsClientListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// OperationsClientListResult contains the result from method OperationsClient.List.
type OperationsClientListResult struct {
	OperationListResult
}

// PrivateEndpointConnectionsClientCreateOrUpdatePollerResponse contains the response from method PrivateEndpointConnectionsClient.CreateOrUpdate.
type PrivateEndpointConnectionsClientCreateOrUpdatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *PrivateEndpointConnectionsClientCreateOrUpdatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l PrivateEndpointConnectionsClientCreateOrUpdatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (PrivateEndpointConnectionsClientCreateOrUpdateResponse, error) {
	respType := PrivateEndpointConnectionsClientCreateOrUpdateResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.PrivateEndpointConnection)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a PrivateEndpointConnectionsClientCreateOrUpdatePollerResponse from the provided client and resume token.
func (l *PrivateEndpointConnectionsClientCreateOrUpdatePollerResponse) Resume(ctx context.Context, client *PrivateEndpointConnectionsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("PrivateEndpointConnectionsClient.CreateOrUpdate", token, client.pl)
	if err != nil {
		return err
	}
	poller := &PrivateEndpointConnectionsClientCreateOrUpdatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// PrivateEndpointConnectionsClientCreateOrUpdateResponse contains the response from method PrivateEndpointConnectionsClient.CreateOrUpdate.
type PrivateEndpointConnectionsClientCreateOrUpdateResponse struct {
	PrivateEndpointConnectionsClientCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateEndpointConnectionsClientCreateOrUpdateResult contains the result from method PrivateEndpointConnectionsClient.CreateOrUpdate.
type PrivateEndpointConnectionsClientCreateOrUpdateResult struct {
	PrivateEndpointConnection
}

// PrivateEndpointConnectionsClientDeletePollerResponse contains the response from method PrivateEndpointConnectionsClient.Delete.
type PrivateEndpointConnectionsClientDeletePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *PrivateEndpointConnectionsClientDeletePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l PrivateEndpointConnectionsClientDeletePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (PrivateEndpointConnectionsClientDeleteResponse, error) {
	respType := PrivateEndpointConnectionsClientDeleteResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a PrivateEndpointConnectionsClientDeletePollerResponse from the provided client and resume token.
func (l *PrivateEndpointConnectionsClientDeletePollerResponse) Resume(ctx context.Context, client *PrivateEndpointConnectionsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("PrivateEndpointConnectionsClient.Delete", token, client.pl)
	if err != nil {
		return err
	}
	poller := &PrivateEndpointConnectionsClientDeletePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// PrivateEndpointConnectionsClientDeleteResponse contains the response from method PrivateEndpointConnectionsClient.Delete.
type PrivateEndpointConnectionsClientDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateEndpointConnectionsClientGetResponse contains the response from method PrivateEndpointConnectionsClient.Get.
type PrivateEndpointConnectionsClientGetResponse struct {
	PrivateEndpointConnectionsClientGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateEndpointConnectionsClientGetResult contains the result from method PrivateEndpointConnectionsClient.Get.
type PrivateEndpointConnectionsClientGetResult struct {
	PrivateEndpointConnection
}

// PrivateEndpointConnectionsClientListResponse contains the response from method PrivateEndpointConnectionsClient.List.
type PrivateEndpointConnectionsClientListResponse struct {
	PrivateEndpointConnectionsClientListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateEndpointConnectionsClientListResult contains the result from method PrivateEndpointConnectionsClient.List.
type PrivateEndpointConnectionsClientListResult struct {
	PrivateEndpointConnectionsResponse
}

// PrivateLinkResourcesClientGetResponse contains the response from method PrivateLinkResourcesClient.Get.
type PrivateLinkResourcesClientGetResponse struct {
	PrivateLinkResourcesClientGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateLinkResourcesClientGetResult contains the result from method PrivateLinkResourcesClient.Get.
type PrivateLinkResourcesClientGetResult struct {
	GroupIDInformation
}

// PrivateLinkResourcesClientListResponse contains the response from method PrivateLinkResourcesClient.List.
type PrivateLinkResourcesClientListResponse struct {
	PrivateLinkResourcesClientListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateLinkResourcesClientListResult contains the result from method PrivateLinkResourcesClient.List.
type PrivateLinkResourcesClientListResult struct {
	GroupIDInformationResponse
}
