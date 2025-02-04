//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armservicebus

import (
	"context"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"net/http"
	"time"
)

// DisasterRecoveryConfigsClientBreakPairingResponse contains the response from method DisasterRecoveryConfigsClient.BreakPairing.
type DisasterRecoveryConfigsClientBreakPairingResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DisasterRecoveryConfigsClientCheckNameAvailabilityResponse contains the response from method DisasterRecoveryConfigsClient.CheckNameAvailability.
type DisasterRecoveryConfigsClientCheckNameAvailabilityResponse struct {
	DisasterRecoveryConfigsClientCheckNameAvailabilityResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DisasterRecoveryConfigsClientCheckNameAvailabilityResult contains the result from method DisasterRecoveryConfigsClient.CheckNameAvailability.
type DisasterRecoveryConfigsClientCheckNameAvailabilityResult struct {
	CheckNameAvailabilityResult
}

// DisasterRecoveryConfigsClientCreateOrUpdateResponse contains the response from method DisasterRecoveryConfigsClient.CreateOrUpdate.
type DisasterRecoveryConfigsClientCreateOrUpdateResponse struct {
	DisasterRecoveryConfigsClientCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DisasterRecoveryConfigsClientCreateOrUpdateResult contains the result from method DisasterRecoveryConfigsClient.CreateOrUpdate.
type DisasterRecoveryConfigsClientCreateOrUpdateResult struct {
	ArmDisasterRecovery
}

// DisasterRecoveryConfigsClientDeleteResponse contains the response from method DisasterRecoveryConfigsClient.Delete.
type DisasterRecoveryConfigsClientDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DisasterRecoveryConfigsClientFailOverResponse contains the response from method DisasterRecoveryConfigsClient.FailOver.
type DisasterRecoveryConfigsClientFailOverResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DisasterRecoveryConfigsClientGetAuthorizationRuleResponse contains the response from method DisasterRecoveryConfigsClient.GetAuthorizationRule.
type DisasterRecoveryConfigsClientGetAuthorizationRuleResponse struct {
	DisasterRecoveryConfigsClientGetAuthorizationRuleResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DisasterRecoveryConfigsClientGetAuthorizationRuleResult contains the result from method DisasterRecoveryConfigsClient.GetAuthorizationRule.
type DisasterRecoveryConfigsClientGetAuthorizationRuleResult struct {
	SBAuthorizationRule
}

// DisasterRecoveryConfigsClientGetResponse contains the response from method DisasterRecoveryConfigsClient.Get.
type DisasterRecoveryConfigsClientGetResponse struct {
	DisasterRecoveryConfigsClientGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DisasterRecoveryConfigsClientGetResult contains the result from method DisasterRecoveryConfigsClient.Get.
type DisasterRecoveryConfigsClientGetResult struct {
	ArmDisasterRecovery
}

// DisasterRecoveryConfigsClientListAuthorizationRulesResponse contains the response from method DisasterRecoveryConfigsClient.ListAuthorizationRules.
type DisasterRecoveryConfigsClientListAuthorizationRulesResponse struct {
	DisasterRecoveryConfigsClientListAuthorizationRulesResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DisasterRecoveryConfigsClientListAuthorizationRulesResult contains the result from method DisasterRecoveryConfigsClient.ListAuthorizationRules.
type DisasterRecoveryConfigsClientListAuthorizationRulesResult struct {
	SBAuthorizationRuleListResult
}

// DisasterRecoveryConfigsClientListKeysResponse contains the response from method DisasterRecoveryConfigsClient.ListKeys.
type DisasterRecoveryConfigsClientListKeysResponse struct {
	DisasterRecoveryConfigsClientListKeysResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DisasterRecoveryConfigsClientListKeysResult contains the result from method DisasterRecoveryConfigsClient.ListKeys.
type DisasterRecoveryConfigsClientListKeysResult struct {
	AccessKeys
}

// DisasterRecoveryConfigsClientListResponse contains the response from method DisasterRecoveryConfigsClient.List.
type DisasterRecoveryConfigsClientListResponse struct {
	DisasterRecoveryConfigsClientListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DisasterRecoveryConfigsClientListResult contains the result from method DisasterRecoveryConfigsClient.List.
type DisasterRecoveryConfigsClientListResult struct {
	ArmDisasterRecoveryListResult
}

// MigrationConfigsClientCompleteMigrationResponse contains the response from method MigrationConfigsClient.CompleteMigration.
type MigrationConfigsClientCompleteMigrationResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// MigrationConfigsClientCreateAndStartMigrationPollerResponse contains the response from method MigrationConfigsClient.CreateAndStartMigration.
type MigrationConfigsClientCreateAndStartMigrationPollerResponse struct {
	// Poller contains an initialized poller.
	Poller *MigrationConfigsClientCreateAndStartMigrationPoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l MigrationConfigsClientCreateAndStartMigrationPollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (MigrationConfigsClientCreateAndStartMigrationResponse, error) {
	respType := MigrationConfigsClientCreateAndStartMigrationResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.MigrationConfigProperties)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a MigrationConfigsClientCreateAndStartMigrationPollerResponse from the provided client and resume token.
func (l *MigrationConfigsClientCreateAndStartMigrationPollerResponse) Resume(ctx context.Context, client *MigrationConfigsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("MigrationConfigsClient.CreateAndStartMigration", token, client.pl)
	if err != nil {
		return err
	}
	poller := &MigrationConfigsClientCreateAndStartMigrationPoller{
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

// MigrationConfigsClientCreateAndStartMigrationResponse contains the response from method MigrationConfigsClient.CreateAndStartMigration.
type MigrationConfigsClientCreateAndStartMigrationResponse struct {
	MigrationConfigsClientCreateAndStartMigrationResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// MigrationConfigsClientCreateAndStartMigrationResult contains the result from method MigrationConfigsClient.CreateAndStartMigration.
type MigrationConfigsClientCreateAndStartMigrationResult struct {
	MigrationConfigProperties
}

// MigrationConfigsClientDeleteResponse contains the response from method MigrationConfigsClient.Delete.
type MigrationConfigsClientDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// MigrationConfigsClientGetResponse contains the response from method MigrationConfigsClient.Get.
type MigrationConfigsClientGetResponse struct {
	MigrationConfigsClientGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// MigrationConfigsClientGetResult contains the result from method MigrationConfigsClient.Get.
type MigrationConfigsClientGetResult struct {
	MigrationConfigProperties
}

// MigrationConfigsClientListResponse contains the response from method MigrationConfigsClient.List.
type MigrationConfigsClientListResponse struct {
	MigrationConfigsClientListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// MigrationConfigsClientListResult contains the result from method MigrationConfigsClient.List.
type MigrationConfigsClientListResult struct {
	MigrationConfigListResult
}

// MigrationConfigsClientRevertResponse contains the response from method MigrationConfigsClient.Revert.
type MigrationConfigsClientRevertResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// NamespacesClientCheckNameAvailabilityResponse contains the response from method NamespacesClient.CheckNameAvailability.
type NamespacesClientCheckNameAvailabilityResponse struct {
	NamespacesClientCheckNameAvailabilityResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// NamespacesClientCheckNameAvailabilityResult contains the result from method NamespacesClient.CheckNameAvailability.
type NamespacesClientCheckNameAvailabilityResult struct {
	CheckNameAvailabilityResult
}

// NamespacesClientCreateOrUpdateAuthorizationRuleResponse contains the response from method NamespacesClient.CreateOrUpdateAuthorizationRule.
type NamespacesClientCreateOrUpdateAuthorizationRuleResponse struct {
	NamespacesClientCreateOrUpdateAuthorizationRuleResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// NamespacesClientCreateOrUpdateAuthorizationRuleResult contains the result from method NamespacesClient.CreateOrUpdateAuthorizationRule.
type NamespacesClientCreateOrUpdateAuthorizationRuleResult struct {
	SBAuthorizationRule
}

// NamespacesClientCreateOrUpdateNetworkRuleSetResponse contains the response from method NamespacesClient.CreateOrUpdateNetworkRuleSet.
type NamespacesClientCreateOrUpdateNetworkRuleSetResponse struct {
	NamespacesClientCreateOrUpdateNetworkRuleSetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// NamespacesClientCreateOrUpdateNetworkRuleSetResult contains the result from method NamespacesClient.CreateOrUpdateNetworkRuleSet.
type NamespacesClientCreateOrUpdateNetworkRuleSetResult struct {
	NetworkRuleSet
}

// NamespacesClientCreateOrUpdatePollerResponse contains the response from method NamespacesClient.CreateOrUpdate.
type NamespacesClientCreateOrUpdatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *NamespacesClientCreateOrUpdatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l NamespacesClientCreateOrUpdatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (NamespacesClientCreateOrUpdateResponse, error) {
	respType := NamespacesClientCreateOrUpdateResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.SBNamespace)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a NamespacesClientCreateOrUpdatePollerResponse from the provided client and resume token.
func (l *NamespacesClientCreateOrUpdatePollerResponse) Resume(ctx context.Context, client *NamespacesClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("NamespacesClient.CreateOrUpdate", token, client.pl)
	if err != nil {
		return err
	}
	poller := &NamespacesClientCreateOrUpdatePoller{
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

// NamespacesClientCreateOrUpdateResponse contains the response from method NamespacesClient.CreateOrUpdate.
type NamespacesClientCreateOrUpdateResponse struct {
	NamespacesClientCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// NamespacesClientCreateOrUpdateResult contains the result from method NamespacesClient.CreateOrUpdate.
type NamespacesClientCreateOrUpdateResult struct {
	SBNamespace
}

// NamespacesClientDeleteAuthorizationRuleResponse contains the response from method NamespacesClient.DeleteAuthorizationRule.
type NamespacesClientDeleteAuthorizationRuleResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// NamespacesClientDeletePollerResponse contains the response from method NamespacesClient.Delete.
type NamespacesClientDeletePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *NamespacesClientDeletePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l NamespacesClientDeletePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (NamespacesClientDeleteResponse, error) {
	respType := NamespacesClientDeleteResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a NamespacesClientDeletePollerResponse from the provided client and resume token.
func (l *NamespacesClientDeletePollerResponse) Resume(ctx context.Context, client *NamespacesClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("NamespacesClient.Delete", token, client.pl)
	if err != nil {
		return err
	}
	poller := &NamespacesClientDeletePoller{
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

// NamespacesClientDeleteResponse contains the response from method NamespacesClient.Delete.
type NamespacesClientDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// NamespacesClientGetAuthorizationRuleResponse contains the response from method NamespacesClient.GetAuthorizationRule.
type NamespacesClientGetAuthorizationRuleResponse struct {
	NamespacesClientGetAuthorizationRuleResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// NamespacesClientGetAuthorizationRuleResult contains the result from method NamespacesClient.GetAuthorizationRule.
type NamespacesClientGetAuthorizationRuleResult struct {
	SBAuthorizationRule
}

// NamespacesClientGetNetworkRuleSetResponse contains the response from method NamespacesClient.GetNetworkRuleSet.
type NamespacesClientGetNetworkRuleSetResponse struct {
	NamespacesClientGetNetworkRuleSetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// NamespacesClientGetNetworkRuleSetResult contains the result from method NamespacesClient.GetNetworkRuleSet.
type NamespacesClientGetNetworkRuleSetResult struct {
	NetworkRuleSet
}

// NamespacesClientGetResponse contains the response from method NamespacesClient.Get.
type NamespacesClientGetResponse struct {
	NamespacesClientGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// NamespacesClientGetResult contains the result from method NamespacesClient.Get.
type NamespacesClientGetResult struct {
	SBNamespace
}

// NamespacesClientListAuthorizationRulesResponse contains the response from method NamespacesClient.ListAuthorizationRules.
type NamespacesClientListAuthorizationRulesResponse struct {
	NamespacesClientListAuthorizationRulesResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// NamespacesClientListAuthorizationRulesResult contains the result from method NamespacesClient.ListAuthorizationRules.
type NamespacesClientListAuthorizationRulesResult struct {
	SBAuthorizationRuleListResult
}

// NamespacesClientListByResourceGroupResponse contains the response from method NamespacesClient.ListByResourceGroup.
type NamespacesClientListByResourceGroupResponse struct {
	NamespacesClientListByResourceGroupResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// NamespacesClientListByResourceGroupResult contains the result from method NamespacesClient.ListByResourceGroup.
type NamespacesClientListByResourceGroupResult struct {
	SBNamespaceListResult
}

// NamespacesClientListKeysResponse contains the response from method NamespacesClient.ListKeys.
type NamespacesClientListKeysResponse struct {
	NamespacesClientListKeysResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// NamespacesClientListKeysResult contains the result from method NamespacesClient.ListKeys.
type NamespacesClientListKeysResult struct {
	AccessKeys
}

// NamespacesClientListNetworkRuleSetsResponse contains the response from method NamespacesClient.ListNetworkRuleSets.
type NamespacesClientListNetworkRuleSetsResponse struct {
	NamespacesClientListNetworkRuleSetsResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// NamespacesClientListNetworkRuleSetsResult contains the result from method NamespacesClient.ListNetworkRuleSets.
type NamespacesClientListNetworkRuleSetsResult struct {
	NetworkRuleSetListResult
}

// NamespacesClientListResponse contains the response from method NamespacesClient.List.
type NamespacesClientListResponse struct {
	NamespacesClientListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// NamespacesClientListResult contains the result from method NamespacesClient.List.
type NamespacesClientListResult struct {
	SBNamespaceListResult
}

// NamespacesClientRegenerateKeysResponse contains the response from method NamespacesClient.RegenerateKeys.
type NamespacesClientRegenerateKeysResponse struct {
	NamespacesClientRegenerateKeysResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// NamespacesClientRegenerateKeysResult contains the result from method NamespacesClient.RegenerateKeys.
type NamespacesClientRegenerateKeysResult struct {
	AccessKeys
}

// NamespacesClientUpdateResponse contains the response from method NamespacesClient.Update.
type NamespacesClientUpdateResponse struct {
	NamespacesClientUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// NamespacesClientUpdateResult contains the result from method NamespacesClient.Update.
type NamespacesClientUpdateResult struct {
	SBNamespace
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
	PrivateEndpointConnectionListResult
}

// PrivateLinkResourcesClientGetResponse contains the response from method PrivateLinkResourcesClient.Get.
type PrivateLinkResourcesClientGetResponse struct {
	PrivateLinkResourcesClientGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateLinkResourcesClientGetResult contains the result from method PrivateLinkResourcesClient.Get.
type PrivateLinkResourcesClientGetResult struct {
	PrivateLinkResourcesListResult
}

// QueuesClientCreateOrUpdateAuthorizationRuleResponse contains the response from method QueuesClient.CreateOrUpdateAuthorizationRule.
type QueuesClientCreateOrUpdateAuthorizationRuleResponse struct {
	QueuesClientCreateOrUpdateAuthorizationRuleResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// QueuesClientCreateOrUpdateAuthorizationRuleResult contains the result from method QueuesClient.CreateOrUpdateAuthorizationRule.
type QueuesClientCreateOrUpdateAuthorizationRuleResult struct {
	SBAuthorizationRule
}

// QueuesClientCreateOrUpdateResponse contains the response from method QueuesClient.CreateOrUpdate.
type QueuesClientCreateOrUpdateResponse struct {
	QueuesClientCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// QueuesClientCreateOrUpdateResult contains the result from method QueuesClient.CreateOrUpdate.
type QueuesClientCreateOrUpdateResult struct {
	SBQueue
}

// QueuesClientDeleteAuthorizationRuleResponse contains the response from method QueuesClient.DeleteAuthorizationRule.
type QueuesClientDeleteAuthorizationRuleResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// QueuesClientDeleteResponse contains the response from method QueuesClient.Delete.
type QueuesClientDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// QueuesClientGetAuthorizationRuleResponse contains the response from method QueuesClient.GetAuthorizationRule.
type QueuesClientGetAuthorizationRuleResponse struct {
	QueuesClientGetAuthorizationRuleResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// QueuesClientGetAuthorizationRuleResult contains the result from method QueuesClient.GetAuthorizationRule.
type QueuesClientGetAuthorizationRuleResult struct {
	SBAuthorizationRule
}

// QueuesClientGetResponse contains the response from method QueuesClient.Get.
type QueuesClientGetResponse struct {
	QueuesClientGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// QueuesClientGetResult contains the result from method QueuesClient.Get.
type QueuesClientGetResult struct {
	SBQueue
}

// QueuesClientListAuthorizationRulesResponse contains the response from method QueuesClient.ListAuthorizationRules.
type QueuesClientListAuthorizationRulesResponse struct {
	QueuesClientListAuthorizationRulesResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// QueuesClientListAuthorizationRulesResult contains the result from method QueuesClient.ListAuthorizationRules.
type QueuesClientListAuthorizationRulesResult struct {
	SBAuthorizationRuleListResult
}

// QueuesClientListByNamespaceResponse contains the response from method QueuesClient.ListByNamespace.
type QueuesClientListByNamespaceResponse struct {
	QueuesClientListByNamespaceResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// QueuesClientListByNamespaceResult contains the result from method QueuesClient.ListByNamespace.
type QueuesClientListByNamespaceResult struct {
	SBQueueListResult
}

// QueuesClientListKeysResponse contains the response from method QueuesClient.ListKeys.
type QueuesClientListKeysResponse struct {
	QueuesClientListKeysResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// QueuesClientListKeysResult contains the result from method QueuesClient.ListKeys.
type QueuesClientListKeysResult struct {
	AccessKeys
}

// QueuesClientRegenerateKeysResponse contains the response from method QueuesClient.RegenerateKeys.
type QueuesClientRegenerateKeysResponse struct {
	QueuesClientRegenerateKeysResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// QueuesClientRegenerateKeysResult contains the result from method QueuesClient.RegenerateKeys.
type QueuesClientRegenerateKeysResult struct {
	AccessKeys
}

// RulesClientCreateOrUpdateResponse contains the response from method RulesClient.CreateOrUpdate.
type RulesClientCreateOrUpdateResponse struct {
	RulesClientCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// RulesClientCreateOrUpdateResult contains the result from method RulesClient.CreateOrUpdate.
type RulesClientCreateOrUpdateResult struct {
	Rule
}

// RulesClientDeleteResponse contains the response from method RulesClient.Delete.
type RulesClientDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// RulesClientGetResponse contains the response from method RulesClient.Get.
type RulesClientGetResponse struct {
	RulesClientGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// RulesClientGetResult contains the result from method RulesClient.Get.
type RulesClientGetResult struct {
	Rule
}

// RulesClientListBySubscriptionsResponse contains the response from method RulesClient.ListBySubscriptions.
type RulesClientListBySubscriptionsResponse struct {
	RulesClientListBySubscriptionsResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// RulesClientListBySubscriptionsResult contains the result from method RulesClient.ListBySubscriptions.
type RulesClientListBySubscriptionsResult struct {
	RuleListResult
}

// SubscriptionsClientCreateOrUpdateResponse contains the response from method SubscriptionsClient.CreateOrUpdate.
type SubscriptionsClientCreateOrUpdateResponse struct {
	SubscriptionsClientCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SubscriptionsClientCreateOrUpdateResult contains the result from method SubscriptionsClient.CreateOrUpdate.
type SubscriptionsClientCreateOrUpdateResult struct {
	SBSubscription
}

// SubscriptionsClientDeleteResponse contains the response from method SubscriptionsClient.Delete.
type SubscriptionsClientDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SubscriptionsClientGetResponse contains the response from method SubscriptionsClient.Get.
type SubscriptionsClientGetResponse struct {
	SubscriptionsClientGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SubscriptionsClientGetResult contains the result from method SubscriptionsClient.Get.
type SubscriptionsClientGetResult struct {
	SBSubscription
}

// SubscriptionsClientListByTopicResponse contains the response from method SubscriptionsClient.ListByTopic.
type SubscriptionsClientListByTopicResponse struct {
	SubscriptionsClientListByTopicResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SubscriptionsClientListByTopicResult contains the result from method SubscriptionsClient.ListByTopic.
type SubscriptionsClientListByTopicResult struct {
	SBSubscriptionListResult
}

// TopicsClientCreateOrUpdateAuthorizationRuleResponse contains the response from method TopicsClient.CreateOrUpdateAuthorizationRule.
type TopicsClientCreateOrUpdateAuthorizationRuleResponse struct {
	TopicsClientCreateOrUpdateAuthorizationRuleResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// TopicsClientCreateOrUpdateAuthorizationRuleResult contains the result from method TopicsClient.CreateOrUpdateAuthorizationRule.
type TopicsClientCreateOrUpdateAuthorizationRuleResult struct {
	SBAuthorizationRule
}

// TopicsClientCreateOrUpdateResponse contains the response from method TopicsClient.CreateOrUpdate.
type TopicsClientCreateOrUpdateResponse struct {
	TopicsClientCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// TopicsClientCreateOrUpdateResult contains the result from method TopicsClient.CreateOrUpdate.
type TopicsClientCreateOrUpdateResult struct {
	SBTopic
}

// TopicsClientDeleteAuthorizationRuleResponse contains the response from method TopicsClient.DeleteAuthorizationRule.
type TopicsClientDeleteAuthorizationRuleResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// TopicsClientDeleteResponse contains the response from method TopicsClient.Delete.
type TopicsClientDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// TopicsClientGetAuthorizationRuleResponse contains the response from method TopicsClient.GetAuthorizationRule.
type TopicsClientGetAuthorizationRuleResponse struct {
	TopicsClientGetAuthorizationRuleResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// TopicsClientGetAuthorizationRuleResult contains the result from method TopicsClient.GetAuthorizationRule.
type TopicsClientGetAuthorizationRuleResult struct {
	SBAuthorizationRule
}

// TopicsClientGetResponse contains the response from method TopicsClient.Get.
type TopicsClientGetResponse struct {
	TopicsClientGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// TopicsClientGetResult contains the result from method TopicsClient.Get.
type TopicsClientGetResult struct {
	SBTopic
}

// TopicsClientListAuthorizationRulesResponse contains the response from method TopicsClient.ListAuthorizationRules.
type TopicsClientListAuthorizationRulesResponse struct {
	TopicsClientListAuthorizationRulesResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// TopicsClientListAuthorizationRulesResult contains the result from method TopicsClient.ListAuthorizationRules.
type TopicsClientListAuthorizationRulesResult struct {
	SBAuthorizationRuleListResult
}

// TopicsClientListByNamespaceResponse contains the response from method TopicsClient.ListByNamespace.
type TopicsClientListByNamespaceResponse struct {
	TopicsClientListByNamespaceResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// TopicsClientListByNamespaceResult contains the result from method TopicsClient.ListByNamespace.
type TopicsClientListByNamespaceResult struct {
	SBTopicListResult
}

// TopicsClientListKeysResponse contains the response from method TopicsClient.ListKeys.
type TopicsClientListKeysResponse struct {
	TopicsClientListKeysResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// TopicsClientListKeysResult contains the result from method TopicsClient.ListKeys.
type TopicsClientListKeysResult struct {
	AccessKeys
}

// TopicsClientRegenerateKeysResponse contains the response from method TopicsClient.RegenerateKeys.
type TopicsClientRegenerateKeysResponse struct {
	TopicsClientRegenerateKeysResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// TopicsClientRegenerateKeysResult contains the result from method TopicsClient.RegenerateKeys.
type TopicsClientRegenerateKeysResult struct {
	AccessKeys
}
