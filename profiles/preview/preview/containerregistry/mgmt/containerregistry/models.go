//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/eng/tools/profileBuilder

package containerregistry

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/preview/containerregistry/mgmt/2021-09-01/containerregistry"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type Action = original.Action

const (
	ActionAllow Action = original.ActionAllow
)

type ActionsRequired = original.ActionsRequired

const (
	ActionsRequiredNone     ActionsRequired = original.ActionsRequiredNone
	ActionsRequiredRecreate ActionsRequired = original.ActionsRequiredRecreate
)

type Architecture = original.Architecture

const (
	ArchitectureAmd64         Architecture = original.ArchitectureAmd64
	ArchitectureArm           Architecture = original.ArchitectureArm
	ArchitectureArm64         Architecture = original.ArchitectureArm64
	ArchitectureThreeEightSix Architecture = original.ArchitectureThreeEightSix
	ArchitectureX86           Architecture = original.ArchitectureX86
)

type BaseImageDependencyType = original.BaseImageDependencyType

const (
	BaseImageDependencyTypeBuildTime BaseImageDependencyType = original.BaseImageDependencyTypeBuildTime
	BaseImageDependencyTypeRunTime   BaseImageDependencyType = original.BaseImageDependencyTypeRunTime
)

type BaseImageTriggerType = original.BaseImageTriggerType

const (
	BaseImageTriggerTypeAll     BaseImageTriggerType = original.BaseImageTriggerTypeAll
	BaseImageTriggerTypeRuntime BaseImageTriggerType = original.BaseImageTriggerTypeRuntime
)

type ConnectionStatus = original.ConnectionStatus

const (
	ConnectionStatusApproved     ConnectionStatus = original.ConnectionStatusApproved
	ConnectionStatusDisconnected ConnectionStatus = original.ConnectionStatusDisconnected
	ConnectionStatusPending      ConnectionStatus = original.ConnectionStatusPending
	ConnectionStatusRejected     ConnectionStatus = original.ConnectionStatusRejected
)

type CreatedByType = original.CreatedByType

const (
	CreatedByTypeApplication     CreatedByType = original.CreatedByTypeApplication
	CreatedByTypeKey             CreatedByType = original.CreatedByTypeKey
	CreatedByTypeManagedIdentity CreatedByType = original.CreatedByTypeManagedIdentity
	CreatedByTypeUser            CreatedByType = original.CreatedByTypeUser
)

type DefaultAction = original.DefaultAction

const (
	DefaultActionAllow DefaultAction = original.DefaultActionAllow
	DefaultActionDeny  DefaultAction = original.DefaultActionDeny
)

type EncryptionStatus = original.EncryptionStatus

const (
	EncryptionStatusDisabled EncryptionStatus = original.EncryptionStatusDisabled
	EncryptionStatusEnabled  EncryptionStatus = original.EncryptionStatusEnabled
)

type ExportPolicyStatus = original.ExportPolicyStatus

const (
	ExportPolicyStatusDisabled ExportPolicyStatus = original.ExportPolicyStatusDisabled
	ExportPolicyStatusEnabled  ExportPolicyStatus = original.ExportPolicyStatusEnabled
)

type ImportMode = original.ImportMode

const (
	ImportModeForce   ImportMode = original.ImportModeForce
	ImportModeNoForce ImportMode = original.ImportModeNoForce
)

type LastModifiedByType = original.LastModifiedByType

const (
	LastModifiedByTypeApplication     LastModifiedByType = original.LastModifiedByTypeApplication
	LastModifiedByTypeKey             LastModifiedByType = original.LastModifiedByTypeKey
	LastModifiedByTypeManagedIdentity LastModifiedByType = original.LastModifiedByTypeManagedIdentity
	LastModifiedByTypeUser            LastModifiedByType = original.LastModifiedByTypeUser
)

type NetworkRuleBypassOptions = original.NetworkRuleBypassOptions

const (
	NetworkRuleBypassOptionsAzureServices NetworkRuleBypassOptions = original.NetworkRuleBypassOptionsAzureServices
	NetworkRuleBypassOptionsNone          NetworkRuleBypassOptions = original.NetworkRuleBypassOptionsNone
)

type OS = original.OS

const (
	OSLinux   OS = original.OSLinux
	OSWindows OS = original.OSWindows
)

type PasswordName = original.PasswordName

const (
	PasswordNamePassword  PasswordName = original.PasswordNamePassword
	PasswordNamePassword2 PasswordName = original.PasswordNamePassword2
)

type PolicyStatus = original.PolicyStatus

const (
	PolicyStatusDisabled PolicyStatus = original.PolicyStatusDisabled
	PolicyStatusEnabled  PolicyStatus = original.PolicyStatusEnabled
)

type ProvisioningState = original.ProvisioningState

const (
	ProvisioningStateCanceled  ProvisioningState = original.ProvisioningStateCanceled
	ProvisioningStateCreating  ProvisioningState = original.ProvisioningStateCreating
	ProvisioningStateDeleting  ProvisioningState = original.ProvisioningStateDeleting
	ProvisioningStateFailed    ProvisioningState = original.ProvisioningStateFailed
	ProvisioningStateSucceeded ProvisioningState = original.ProvisioningStateSucceeded
	ProvisioningStateUpdating  ProvisioningState = original.ProvisioningStateUpdating
)

type PublicNetworkAccess = original.PublicNetworkAccess

const (
	PublicNetworkAccessDisabled PublicNetworkAccess = original.PublicNetworkAccessDisabled
	PublicNetworkAccessEnabled  PublicNetworkAccess = original.PublicNetworkAccessEnabled
)

type RegistryUsageUnit = original.RegistryUsageUnit

const (
	RegistryUsageUnitBytes RegistryUsageUnit = original.RegistryUsageUnitBytes
	RegistryUsageUnitCount RegistryUsageUnit = original.RegistryUsageUnitCount
)

type ResourceIdentityType = original.ResourceIdentityType

const (
	ResourceIdentityTypeNone                       ResourceIdentityType = original.ResourceIdentityTypeNone
	ResourceIdentityTypeSystemAssigned             ResourceIdentityType = original.ResourceIdentityTypeSystemAssigned
	ResourceIdentityTypeSystemAssignedUserAssigned ResourceIdentityType = original.ResourceIdentityTypeSystemAssignedUserAssigned
	ResourceIdentityTypeUserAssigned               ResourceIdentityType = original.ResourceIdentityTypeUserAssigned
)

type RunStatus = original.RunStatus

const (
	RunStatusCanceled  RunStatus = original.RunStatusCanceled
	RunStatusError     RunStatus = original.RunStatusError
	RunStatusFailed    RunStatus = original.RunStatusFailed
	RunStatusQueued    RunStatus = original.RunStatusQueued
	RunStatusRunning   RunStatus = original.RunStatusRunning
	RunStatusStarted   RunStatus = original.RunStatusStarted
	RunStatusSucceeded RunStatus = original.RunStatusSucceeded
	RunStatusTimeout   RunStatus = original.RunStatusTimeout
)

type RunType = original.RunType

const (
	RunTypeAutoBuild  RunType = original.RunTypeAutoBuild
	RunTypeAutoRun    RunType = original.RunTypeAutoRun
	RunTypeQuickBuild RunType = original.RunTypeQuickBuild
	RunTypeQuickRun   RunType = original.RunTypeQuickRun
)

type SecretObjectType = original.SecretObjectType

const (
	SecretObjectTypeOpaque      SecretObjectType = original.SecretObjectTypeOpaque
	SecretObjectTypeVaultsecret SecretObjectType = original.SecretObjectTypeVaultsecret
)

type SkuName = original.SkuName

const (
	SkuNameBasic    SkuName = original.SkuNameBasic
	SkuNameClassic  SkuName = original.SkuNameClassic
	SkuNamePremium  SkuName = original.SkuNamePremium
	SkuNameStandard SkuName = original.SkuNameStandard
)

type SkuTier = original.SkuTier

const (
	SkuTierBasic    SkuTier = original.SkuTierBasic
	SkuTierClassic  SkuTier = original.SkuTierClassic
	SkuTierPremium  SkuTier = original.SkuTierPremium
	SkuTierStandard SkuTier = original.SkuTierStandard
)

type SourceControlType = original.SourceControlType

const (
	SourceControlTypeGithub                  SourceControlType = original.SourceControlTypeGithub
	SourceControlTypeVisualStudioTeamService SourceControlType = original.SourceControlTypeVisualStudioTeamService
)

type SourceRegistryLoginMode = original.SourceRegistryLoginMode

const (
	SourceRegistryLoginModeDefault SourceRegistryLoginMode = original.SourceRegistryLoginModeDefault
	SourceRegistryLoginModeNone    SourceRegistryLoginMode = original.SourceRegistryLoginModeNone
)

type SourceTriggerEvent = original.SourceTriggerEvent

const (
	SourceTriggerEventCommit      SourceTriggerEvent = original.SourceTriggerEventCommit
	SourceTriggerEventPullrequest SourceTriggerEvent = original.SourceTriggerEventPullrequest
)

type TaskStatus = original.TaskStatus

const (
	TaskStatusDisabled TaskStatus = original.TaskStatusDisabled
	TaskStatusEnabled  TaskStatus = original.TaskStatusEnabled
)

type TokenType = original.TokenType

const (
	TokenTypeOAuth TokenType = original.TokenTypeOAuth
	TokenTypePAT   TokenType = original.TokenTypePAT
)

type TriggerStatus = original.TriggerStatus

const (
	TriggerStatusDisabled TriggerStatus = original.TriggerStatusDisabled
	TriggerStatusEnabled  TriggerStatus = original.TriggerStatusEnabled
)

type TrustPolicyType = original.TrustPolicyType

const (
	TrustPolicyTypeNotary TrustPolicyType = original.TrustPolicyTypeNotary
)

type Type = original.Type

const (
	TypeDockerBuildRequest    Type = original.TypeDockerBuildRequest
	TypeEncodedTaskRunRequest Type = original.TypeEncodedTaskRunRequest
	TypeFileTaskRunRequest    Type = original.TypeFileTaskRunRequest
	TypeRunRequest            Type = original.TypeRunRequest
	TypeTaskRunRequest        Type = original.TypeTaskRunRequest
)

type TypeBasicTaskStepProperties = original.TypeBasicTaskStepProperties

const (
	TypeBasicTaskStepPropertiesTypeDocker             TypeBasicTaskStepProperties = original.TypeBasicTaskStepPropertiesTypeDocker
	TypeBasicTaskStepPropertiesTypeEncodedTask        TypeBasicTaskStepProperties = original.TypeBasicTaskStepPropertiesTypeEncodedTask
	TypeBasicTaskStepPropertiesTypeFileTask           TypeBasicTaskStepProperties = original.TypeBasicTaskStepPropertiesTypeFileTask
	TypeBasicTaskStepPropertiesTypeTaskStepProperties TypeBasicTaskStepProperties = original.TypeBasicTaskStepPropertiesTypeTaskStepProperties
)

type TypeBasicTaskStepUpdateParameters = original.TypeBasicTaskStepUpdateParameters

const (
	TypeBasicTaskStepUpdateParametersTypeDocker                   TypeBasicTaskStepUpdateParameters = original.TypeBasicTaskStepUpdateParametersTypeDocker
	TypeBasicTaskStepUpdateParametersTypeEncodedTask              TypeBasicTaskStepUpdateParameters = original.TypeBasicTaskStepUpdateParametersTypeEncodedTask
	TypeBasicTaskStepUpdateParametersTypeFileTask                 TypeBasicTaskStepUpdateParameters = original.TypeBasicTaskStepUpdateParametersTypeFileTask
	TypeBasicTaskStepUpdateParametersTypeTaskStepUpdateParameters TypeBasicTaskStepUpdateParameters = original.TypeBasicTaskStepUpdateParametersTypeTaskStepUpdateParameters
)

type UpdateTriggerPayloadType = original.UpdateTriggerPayloadType

const (
	UpdateTriggerPayloadTypeDefault UpdateTriggerPayloadType = original.UpdateTriggerPayloadTypeDefault
	UpdateTriggerPayloadTypeToken   UpdateTriggerPayloadType = original.UpdateTriggerPayloadTypeToken
)

type Variant = original.Variant

const (
	VariantV6 Variant = original.VariantV6
	VariantV7 Variant = original.VariantV7
	VariantV8 Variant = original.VariantV8
)

type WebhookAction = original.WebhookAction

const (
	WebhookActionChartDelete WebhookAction = original.WebhookActionChartDelete
	WebhookActionChartPush   WebhookAction = original.WebhookActionChartPush
	WebhookActionDelete      WebhookAction = original.WebhookActionDelete
	WebhookActionPush        WebhookAction = original.WebhookActionPush
	WebhookActionQuarantine  WebhookAction = original.WebhookActionQuarantine
)

type WebhookStatus = original.WebhookStatus

const (
	WebhookStatusDisabled WebhookStatus = original.WebhookStatusDisabled
	WebhookStatusEnabled  WebhookStatus = original.WebhookStatusEnabled
)

type ZoneRedundancy = original.ZoneRedundancy

const (
	ZoneRedundancyDisabled ZoneRedundancy = original.ZoneRedundancyDisabled
	ZoneRedundancyEnabled  ZoneRedundancy = original.ZoneRedundancyEnabled
)

type Actor = original.Actor
type AgentPool = original.AgentPool
type AgentPoolListResult = original.AgentPoolListResult
type AgentPoolListResultIterator = original.AgentPoolListResultIterator
type AgentPoolListResultPage = original.AgentPoolListResultPage
type AgentPoolProperties = original.AgentPoolProperties
type AgentPoolPropertiesUpdateParameters = original.AgentPoolPropertiesUpdateParameters
type AgentPoolQueueStatus = original.AgentPoolQueueStatus
type AgentPoolUpdateParameters = original.AgentPoolUpdateParameters
type AgentPoolsClient = original.AgentPoolsClient
type AgentPoolsCreateFuture = original.AgentPoolsCreateFuture
type AgentPoolsDeleteFuture = original.AgentPoolsDeleteFuture
type AgentPoolsUpdateFuture = original.AgentPoolsUpdateFuture
type AgentProperties = original.AgentProperties
type Argument = original.Argument
type AuthInfo = original.AuthInfo
type AuthInfoUpdateParameters = original.AuthInfoUpdateParameters
type BaseClient = original.BaseClient
type BaseImageDependency = original.BaseImageDependency
type BaseImageTrigger = original.BaseImageTrigger
type BaseImageTriggerUpdateParameters = original.BaseImageTriggerUpdateParameters
type BasicRunRequest = original.BasicRunRequest
type BasicTaskStepProperties = original.BasicTaskStepProperties
type BasicTaskStepUpdateParameters = original.BasicTaskStepUpdateParameters
type CallbackConfig = original.CallbackConfig
type Credentials = original.Credentials
type CustomRegistryCredentials = original.CustomRegistryCredentials
type DockerBuildRequest = original.DockerBuildRequest
type DockerBuildStep = original.DockerBuildStep
type DockerBuildStepUpdateParameters = original.DockerBuildStepUpdateParameters
type EncodedTaskRunRequest = original.EncodedTaskRunRequest
type EncodedTaskStep = original.EncodedTaskStep
type EncodedTaskStepUpdateParameters = original.EncodedTaskStepUpdateParameters
type EncryptionProperty = original.EncryptionProperty
type ErrorResponse = original.ErrorResponse
type ErrorResponseBody = original.ErrorResponseBody
type Event = original.Event
type EventContent = original.EventContent
type EventInfo = original.EventInfo
type EventListResult = original.EventListResult
type EventListResultIterator = original.EventListResultIterator
type EventListResultPage = original.EventListResultPage
type EventRequestMessage = original.EventRequestMessage
type EventResponseMessage = original.EventResponseMessage
type ExportPolicy = original.ExportPolicy
type FileTaskRunRequest = original.FileTaskRunRequest
type FileTaskStep = original.FileTaskStep
type FileTaskStepUpdateParameters = original.FileTaskStepUpdateParameters
type IPRule = original.IPRule
type IdentityProperties = original.IdentityProperties
type ImageDescriptor = original.ImageDescriptor
type ImageUpdateTrigger = original.ImageUpdateTrigger
type ImportImageParameters = original.ImportImageParameters
type ImportSource = original.ImportSource
type ImportSourceCredentials = original.ImportSourceCredentials
type InnerErrorDescription = original.InnerErrorDescription
type KeyVaultProperties = original.KeyVaultProperties
type NetworkRuleSet = original.NetworkRuleSet
type OperationDefinition = original.OperationDefinition
type OperationDisplayDefinition = original.OperationDisplayDefinition
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage
type OperationLogSpecificationDefinition = original.OperationLogSpecificationDefinition
type OperationMetricSpecificationDefinition = original.OperationMetricSpecificationDefinition
type OperationPropertiesDefinition = original.OperationPropertiesDefinition
type OperationServiceSpecificationDefinition = original.OperationServiceSpecificationDefinition
type OperationsClient = original.OperationsClient
type OverrideTaskStepProperties = original.OverrideTaskStepProperties
type PackageType = original.PackageType
type PlatformProperties = original.PlatformProperties
type PlatformUpdateParameters = original.PlatformUpdateParameters
type Policies = original.Policies
type PrivateEndpoint = original.PrivateEndpoint
type PrivateEndpointConnection = original.PrivateEndpointConnection
type PrivateEndpointConnectionListResult = original.PrivateEndpointConnectionListResult
type PrivateEndpointConnectionListResultIterator = original.PrivateEndpointConnectionListResultIterator
type PrivateEndpointConnectionListResultPage = original.PrivateEndpointConnectionListResultPage
type PrivateEndpointConnectionProperties = original.PrivateEndpointConnectionProperties
type PrivateEndpointConnectionsClient = original.PrivateEndpointConnectionsClient
type PrivateEndpointConnectionsCreateOrUpdateFuture = original.PrivateEndpointConnectionsCreateOrUpdateFuture
type PrivateEndpointConnectionsDeleteFuture = original.PrivateEndpointConnectionsDeleteFuture
type PrivateLinkResource = original.PrivateLinkResource
type PrivateLinkResourceListResult = original.PrivateLinkResourceListResult
type PrivateLinkResourceListResultIterator = original.PrivateLinkResourceListResultIterator
type PrivateLinkResourceListResultPage = original.PrivateLinkResourceListResultPage
type PrivateLinkResourceProperties = original.PrivateLinkResourceProperties
type PrivateLinkServiceConnectionState = original.PrivateLinkServiceConnectionState
type ProxyResource = original.ProxyResource
type QuarantinePolicy = original.QuarantinePolicy
type RegenerateCredentialParameters = original.RegenerateCredentialParameters
type RegistriesClient = original.RegistriesClient
type RegistriesCreateFuture = original.RegistriesCreateFuture
type RegistriesDeleteFuture = original.RegistriesDeleteFuture
type RegistriesImportImageFuture = original.RegistriesImportImageFuture
type RegistriesScheduleRunFuture = original.RegistriesScheduleRunFuture
type RegistriesUpdateFuture = original.RegistriesUpdateFuture
type Registry = original.Registry
type RegistryListCredentialsResult = original.RegistryListCredentialsResult
type RegistryListResult = original.RegistryListResult
type RegistryListResultIterator = original.RegistryListResultIterator
type RegistryListResultPage = original.RegistryListResultPage
type RegistryNameCheckRequest = original.RegistryNameCheckRequest
type RegistryNameStatus = original.RegistryNameStatus
type RegistryPassword = original.RegistryPassword
type RegistryProperties = original.RegistryProperties
type RegistryPropertiesUpdateParameters = original.RegistryPropertiesUpdateParameters
type RegistryUpdateParameters = original.RegistryUpdateParameters
type RegistryUsage = original.RegistryUsage
type RegistryUsageListResult = original.RegistryUsageListResult
type Replication = original.Replication
type ReplicationListResult = original.ReplicationListResult
type ReplicationListResultIterator = original.ReplicationListResultIterator
type ReplicationListResultPage = original.ReplicationListResultPage
type ReplicationProperties = original.ReplicationProperties
type ReplicationUpdateParameters = original.ReplicationUpdateParameters
type ReplicationUpdateParametersProperties = original.ReplicationUpdateParametersProperties
type ReplicationsClient = original.ReplicationsClient
type ReplicationsCreateFuture = original.ReplicationsCreateFuture
type ReplicationsDeleteFuture = original.ReplicationsDeleteFuture
type ReplicationsUpdateFuture = original.ReplicationsUpdateFuture
type Request = original.Request
type Resource = original.Resource
type RetentionPolicy = original.RetentionPolicy
type Run = original.Run
type RunFilter = original.RunFilter
type RunGetLogResult = original.RunGetLogResult
type RunListResult = original.RunListResult
type RunListResultIterator = original.RunListResultIterator
type RunListResultPage = original.RunListResultPage
type RunProperties = original.RunProperties
type RunRequest = original.RunRequest
type RunUpdateParameters = original.RunUpdateParameters
type RunsCancelFuture = original.RunsCancelFuture
type RunsClient = original.RunsClient
type RunsUpdateFuture = original.RunsUpdateFuture
type SecretObject = original.SecretObject
type SetValue = original.SetValue
type Sku = original.Sku
type Source = original.Source
type SourceProperties = original.SourceProperties
type SourceRegistryCredentials = original.SourceRegistryCredentials
type SourceTrigger = original.SourceTrigger
type SourceTriggerDescriptor = original.SourceTriggerDescriptor
type SourceTriggerUpdateParameters = original.SourceTriggerUpdateParameters
type SourceUpdateParameters = original.SourceUpdateParameters
type SourceUploadDefinition = original.SourceUploadDefinition
type Status = original.Status
type StorageAccountProperties = original.StorageAccountProperties
type SystemData = original.SystemData
type Target = original.Target
type Task = original.Task
type TaskListResult = original.TaskListResult
type TaskListResultIterator = original.TaskListResultIterator
type TaskListResultPage = original.TaskListResultPage
type TaskProperties = original.TaskProperties
type TaskPropertiesUpdateParameters = original.TaskPropertiesUpdateParameters
type TaskRun = original.TaskRun
type TaskRunListResult = original.TaskRunListResult
type TaskRunListResultIterator = original.TaskRunListResultIterator
type TaskRunListResultPage = original.TaskRunListResultPage
type TaskRunProperties = original.TaskRunProperties
type TaskRunPropertiesUpdateParameters = original.TaskRunPropertiesUpdateParameters
type TaskRunRequest = original.TaskRunRequest
type TaskRunUpdateParameters = original.TaskRunUpdateParameters
type TaskRunsClient = original.TaskRunsClient
type TaskRunsCreateFuture = original.TaskRunsCreateFuture
type TaskRunsDeleteFuture = original.TaskRunsDeleteFuture
type TaskRunsUpdateFuture = original.TaskRunsUpdateFuture
type TaskStepProperties = original.TaskStepProperties
type TaskStepUpdateParameters = original.TaskStepUpdateParameters
type TaskUpdateParameters = original.TaskUpdateParameters
type TasksClient = original.TasksClient
type TasksCreateFuture = original.TasksCreateFuture
type TasksDeleteFuture = original.TasksDeleteFuture
type TasksUpdateFuture = original.TasksUpdateFuture
type TimerTrigger = original.TimerTrigger
type TimerTriggerDescriptor = original.TimerTriggerDescriptor
type TimerTriggerUpdateParameters = original.TimerTriggerUpdateParameters
type TriggerProperties = original.TriggerProperties
type TriggerUpdateParameters = original.TriggerUpdateParameters
type TrustPolicy = original.TrustPolicy
type UserIdentityProperties = original.UserIdentityProperties
type Webhook = original.Webhook
type WebhookCreateParameters = original.WebhookCreateParameters
type WebhookListResult = original.WebhookListResult
type WebhookListResultIterator = original.WebhookListResultIterator
type WebhookListResultPage = original.WebhookListResultPage
type WebhookProperties = original.WebhookProperties
type WebhookPropertiesCreateParameters = original.WebhookPropertiesCreateParameters
type WebhookPropertiesUpdateParameters = original.WebhookPropertiesUpdateParameters
type WebhookUpdateParameters = original.WebhookUpdateParameters
type WebhooksClient = original.WebhooksClient
type WebhooksCreateFuture = original.WebhooksCreateFuture
type WebhooksDeleteFuture = original.WebhooksDeleteFuture
type WebhooksUpdateFuture = original.WebhooksUpdateFuture

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewAgentPoolListResultIterator(page AgentPoolListResultPage) AgentPoolListResultIterator {
	return original.NewAgentPoolListResultIterator(page)
}
func NewAgentPoolListResultPage(cur AgentPoolListResult, getNextPage func(context.Context, AgentPoolListResult) (AgentPoolListResult, error)) AgentPoolListResultPage {
	return original.NewAgentPoolListResultPage(cur, getNextPage)
}
func NewAgentPoolsClient(subscriptionID string) AgentPoolsClient {
	return original.NewAgentPoolsClient(subscriptionID)
}
func NewAgentPoolsClientWithBaseURI(baseURI string, subscriptionID string) AgentPoolsClient {
	return original.NewAgentPoolsClientWithBaseURI(baseURI, subscriptionID)
}
func NewEventListResultIterator(page EventListResultPage) EventListResultIterator {
	return original.NewEventListResultIterator(page)
}
func NewEventListResultPage(cur EventListResult, getNextPage func(context.Context, EventListResult) (EventListResult, error)) EventListResultPage {
	return original.NewEventListResultPage(cur, getNextPage)
}
func NewOperationListResultIterator(page OperationListResultPage) OperationListResultIterator {
	return original.NewOperationListResultIterator(page)
}
func NewOperationListResultPage(cur OperationListResult, getNextPage func(context.Context, OperationListResult) (OperationListResult, error)) OperationListResultPage {
	return original.NewOperationListResultPage(cur, getNextPage)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewPrivateEndpointConnectionListResultIterator(page PrivateEndpointConnectionListResultPage) PrivateEndpointConnectionListResultIterator {
	return original.NewPrivateEndpointConnectionListResultIterator(page)
}
func NewPrivateEndpointConnectionListResultPage(cur PrivateEndpointConnectionListResult, getNextPage func(context.Context, PrivateEndpointConnectionListResult) (PrivateEndpointConnectionListResult, error)) PrivateEndpointConnectionListResultPage {
	return original.NewPrivateEndpointConnectionListResultPage(cur, getNextPage)
}
func NewPrivateEndpointConnectionsClient(subscriptionID string) PrivateEndpointConnectionsClient {
	return original.NewPrivateEndpointConnectionsClient(subscriptionID)
}
func NewPrivateEndpointConnectionsClientWithBaseURI(baseURI string, subscriptionID string) PrivateEndpointConnectionsClient {
	return original.NewPrivateEndpointConnectionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewPrivateLinkResourceListResultIterator(page PrivateLinkResourceListResultPage) PrivateLinkResourceListResultIterator {
	return original.NewPrivateLinkResourceListResultIterator(page)
}
func NewPrivateLinkResourceListResultPage(cur PrivateLinkResourceListResult, getNextPage func(context.Context, PrivateLinkResourceListResult) (PrivateLinkResourceListResult, error)) PrivateLinkResourceListResultPage {
	return original.NewPrivateLinkResourceListResultPage(cur, getNextPage)
}
func NewRegistriesClient(subscriptionID string) RegistriesClient {
	return original.NewRegistriesClient(subscriptionID)
}
func NewRegistriesClientWithBaseURI(baseURI string, subscriptionID string) RegistriesClient {
	return original.NewRegistriesClientWithBaseURI(baseURI, subscriptionID)
}
func NewRegistryListResultIterator(page RegistryListResultPage) RegistryListResultIterator {
	return original.NewRegistryListResultIterator(page)
}
func NewRegistryListResultPage(cur RegistryListResult, getNextPage func(context.Context, RegistryListResult) (RegistryListResult, error)) RegistryListResultPage {
	return original.NewRegistryListResultPage(cur, getNextPage)
}
func NewReplicationListResultIterator(page ReplicationListResultPage) ReplicationListResultIterator {
	return original.NewReplicationListResultIterator(page)
}
func NewReplicationListResultPage(cur ReplicationListResult, getNextPage func(context.Context, ReplicationListResult) (ReplicationListResult, error)) ReplicationListResultPage {
	return original.NewReplicationListResultPage(cur, getNextPage)
}
func NewReplicationsClient(subscriptionID string) ReplicationsClient {
	return original.NewReplicationsClient(subscriptionID)
}
func NewReplicationsClientWithBaseURI(baseURI string, subscriptionID string) ReplicationsClient {
	return original.NewReplicationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewRunListResultIterator(page RunListResultPage) RunListResultIterator {
	return original.NewRunListResultIterator(page)
}
func NewRunListResultPage(cur RunListResult, getNextPage func(context.Context, RunListResult) (RunListResult, error)) RunListResultPage {
	return original.NewRunListResultPage(cur, getNextPage)
}
func NewRunsClient(subscriptionID string) RunsClient {
	return original.NewRunsClient(subscriptionID)
}
func NewRunsClientWithBaseURI(baseURI string, subscriptionID string) RunsClient {
	return original.NewRunsClientWithBaseURI(baseURI, subscriptionID)
}
func NewTaskListResultIterator(page TaskListResultPage) TaskListResultIterator {
	return original.NewTaskListResultIterator(page)
}
func NewTaskListResultPage(cur TaskListResult, getNextPage func(context.Context, TaskListResult) (TaskListResult, error)) TaskListResultPage {
	return original.NewTaskListResultPage(cur, getNextPage)
}
func NewTaskRunListResultIterator(page TaskRunListResultPage) TaskRunListResultIterator {
	return original.NewTaskRunListResultIterator(page)
}
func NewTaskRunListResultPage(cur TaskRunListResult, getNextPage func(context.Context, TaskRunListResult) (TaskRunListResult, error)) TaskRunListResultPage {
	return original.NewTaskRunListResultPage(cur, getNextPage)
}
func NewTaskRunsClient(subscriptionID string) TaskRunsClient {
	return original.NewTaskRunsClient(subscriptionID)
}
func NewTaskRunsClientWithBaseURI(baseURI string, subscriptionID string) TaskRunsClient {
	return original.NewTaskRunsClientWithBaseURI(baseURI, subscriptionID)
}
func NewTasksClient(subscriptionID string) TasksClient {
	return original.NewTasksClient(subscriptionID)
}
func NewTasksClientWithBaseURI(baseURI string, subscriptionID string) TasksClient {
	return original.NewTasksClientWithBaseURI(baseURI, subscriptionID)
}
func NewWebhookListResultIterator(page WebhookListResultPage) WebhookListResultIterator {
	return original.NewWebhookListResultIterator(page)
}
func NewWebhookListResultPage(cur WebhookListResult, getNextPage func(context.Context, WebhookListResult) (WebhookListResult, error)) WebhookListResultPage {
	return original.NewWebhookListResultPage(cur, getNextPage)
}
func NewWebhooksClient(subscriptionID string) WebhooksClient {
	return original.NewWebhooksClient(subscriptionID)
}
func NewWebhooksClientWithBaseURI(baseURI string, subscriptionID string) WebhooksClient {
	return original.NewWebhooksClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleActionValues() []Action {
	return original.PossibleActionValues()
}
func PossibleActionsRequiredValues() []ActionsRequired {
	return original.PossibleActionsRequiredValues()
}
func PossibleArchitectureValues() []Architecture {
	return original.PossibleArchitectureValues()
}
func PossibleBaseImageDependencyTypeValues() []BaseImageDependencyType {
	return original.PossibleBaseImageDependencyTypeValues()
}
func PossibleBaseImageTriggerTypeValues() []BaseImageTriggerType {
	return original.PossibleBaseImageTriggerTypeValues()
}
func PossibleConnectionStatusValues() []ConnectionStatus {
	return original.PossibleConnectionStatusValues()
}
func PossibleCreatedByTypeValues() []CreatedByType {
	return original.PossibleCreatedByTypeValues()
}
func PossibleDefaultActionValues() []DefaultAction {
	return original.PossibleDefaultActionValues()
}
func PossibleEncryptionStatusValues() []EncryptionStatus {
	return original.PossibleEncryptionStatusValues()
}
func PossibleExportPolicyStatusValues() []ExportPolicyStatus {
	return original.PossibleExportPolicyStatusValues()
}
func PossibleImportModeValues() []ImportMode {
	return original.PossibleImportModeValues()
}
func PossibleLastModifiedByTypeValues() []LastModifiedByType {
	return original.PossibleLastModifiedByTypeValues()
}
func PossibleNetworkRuleBypassOptionsValues() []NetworkRuleBypassOptions {
	return original.PossibleNetworkRuleBypassOptionsValues()
}
func PossibleOSValues() []OS {
	return original.PossibleOSValues()
}
func PossiblePasswordNameValues() []PasswordName {
	return original.PossiblePasswordNameValues()
}
func PossiblePolicyStatusValues() []PolicyStatus {
	return original.PossiblePolicyStatusValues()
}
func PossibleProvisioningStateValues() []ProvisioningState {
	return original.PossibleProvisioningStateValues()
}
func PossiblePublicNetworkAccessValues() []PublicNetworkAccess {
	return original.PossiblePublicNetworkAccessValues()
}
func PossibleRegistryUsageUnitValues() []RegistryUsageUnit {
	return original.PossibleRegistryUsageUnitValues()
}
func PossibleResourceIdentityTypeValues() []ResourceIdentityType {
	return original.PossibleResourceIdentityTypeValues()
}
func PossibleRunStatusValues() []RunStatus {
	return original.PossibleRunStatusValues()
}
func PossibleRunTypeValues() []RunType {
	return original.PossibleRunTypeValues()
}
func PossibleSecretObjectTypeValues() []SecretObjectType {
	return original.PossibleSecretObjectTypeValues()
}
func PossibleSkuNameValues() []SkuName {
	return original.PossibleSkuNameValues()
}
func PossibleSkuTierValues() []SkuTier {
	return original.PossibleSkuTierValues()
}
func PossibleSourceControlTypeValues() []SourceControlType {
	return original.PossibleSourceControlTypeValues()
}
func PossibleSourceRegistryLoginModeValues() []SourceRegistryLoginMode {
	return original.PossibleSourceRegistryLoginModeValues()
}
func PossibleSourceTriggerEventValues() []SourceTriggerEvent {
	return original.PossibleSourceTriggerEventValues()
}
func PossibleTaskStatusValues() []TaskStatus {
	return original.PossibleTaskStatusValues()
}
func PossibleTokenTypeValues() []TokenType {
	return original.PossibleTokenTypeValues()
}
func PossibleTriggerStatusValues() []TriggerStatus {
	return original.PossibleTriggerStatusValues()
}
func PossibleTrustPolicyTypeValues() []TrustPolicyType {
	return original.PossibleTrustPolicyTypeValues()
}
func PossibleTypeBasicTaskStepPropertiesValues() []TypeBasicTaskStepProperties {
	return original.PossibleTypeBasicTaskStepPropertiesValues()
}
func PossibleTypeBasicTaskStepUpdateParametersValues() []TypeBasicTaskStepUpdateParameters {
	return original.PossibleTypeBasicTaskStepUpdateParametersValues()
}
func PossibleTypeValues() []Type {
	return original.PossibleTypeValues()
}
func PossibleUpdateTriggerPayloadTypeValues() []UpdateTriggerPayloadType {
	return original.PossibleUpdateTriggerPayloadTypeValues()
}
func PossibleVariantValues() []Variant {
	return original.PossibleVariantValues()
}
func PossibleWebhookActionValues() []WebhookAction {
	return original.PossibleWebhookActionValues()
}
func PossibleWebhookStatusValues() []WebhookStatus {
	return original.PossibleWebhookStatusValues()
}
func PossibleZoneRedundancyValues() []ZoneRedundancy {
	return original.PossibleZoneRedundancyValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
