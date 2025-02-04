# Release History

## 0.2.0 (2022-01-14)
### Breaking Changes

- Function `*TagRulesClient.Get` parameter(s) have been changed from `(context.Context, string, string, string, *TagRulesGetOptions)` to `(context.Context, string, string, string, *TagRulesClientGetOptions)`
- Function `*TagRulesClient.Get` return value(s) have been changed from `(TagRulesGetResponse, error)` to `(TagRulesClientGetResponse, error)`
- Function `*TagRulesClient.BeginDelete` parameter(s) have been changed from `(context.Context, string, string, string, *TagRulesBeginDeleteOptions)` to `(context.Context, string, string, string, *TagRulesClientBeginDeleteOptions)`
- Function `*TagRulesClient.BeginDelete` return value(s) have been changed from `(TagRulesDeletePollerResponse, error)` to `(TagRulesClientDeletePollerResponse, error)`
- Function `*OperationsClient.List` parameter(s) have been changed from `(*OperationsListOptions)` to `(*OperationsClientListOptions)`
- Function `*OperationsClient.List` return value(s) have been changed from `(*OperationsListPager)` to `(*OperationsClientListPager)`
- Function `*MonitorsClient.List` parameter(s) have been changed from `(*MonitorsListOptions)` to `(*MonitorsClientListOptions)`
- Function `*MonitorsClient.List` return value(s) have been changed from `(*MonitorsListPager)` to `(*MonitorsClientListPager)`
- Function `*VMHostClient.List` parameter(s) have been changed from `(string, string, *VMHostListOptions)` to `(string, string, *VMHostClientListOptions)`
- Function `*VMHostClient.List` return value(s) have been changed from `(*VMHostListPager)` to `(*VMHostClientListPager)`
- Function `*MonitoredResourcesClient.List` parameter(s) have been changed from `(string, string, *MonitoredResourcesListOptions)` to `(string, string, *MonitoredResourcesClientListOptions)`
- Function `*MonitoredResourcesClient.List` return value(s) have been changed from `(*MonitoredResourcesListPager)` to `(*MonitoredResourcesClientListPager)`
- Function `*MonitorsClient.ListByResourceGroup` parameter(s) have been changed from `(string, *MonitorsListByResourceGroupOptions)` to `(string, *MonitorsClientListByResourceGroupOptions)`
- Function `*MonitorsClient.ListByResourceGroup` return value(s) have been changed from `(*MonitorsListByResourceGroupPager)` to `(*MonitorsClientListByResourceGroupPager)`
- Function `*DeploymentInfoClient.List` parameter(s) have been changed from `(context.Context, string, string, *DeploymentInfoListOptions)` to `(context.Context, string, string, *DeploymentInfoClientListOptions)`
- Function `*DeploymentInfoClient.List` return value(s) have been changed from `(DeploymentInfoListResponse, error)` to `(DeploymentInfoClientListResponse, error)`
- Function `*TagRulesClient.List` parameter(s) have been changed from `(string, string, *TagRulesListOptions)` to `(string, string, *TagRulesClientListOptions)`
- Function `*TagRulesClient.List` return value(s) have been changed from `(*TagRulesListPager)` to `(*TagRulesClientListPager)`
- Function `*MonitorsClient.BeginCreate` parameter(s) have been changed from `(context.Context, string, string, *MonitorsBeginCreateOptions)` to `(context.Context, string, string, *MonitorsClientBeginCreateOptions)`
- Function `*MonitorsClient.BeginCreate` return value(s) have been changed from `(MonitorsCreatePollerResponse, error)` to `(MonitorsClientCreatePollerResponse, error)`
- Function `*MonitorsClient.BeginDelete` parameter(s) have been changed from `(context.Context, string, string, *MonitorsBeginDeleteOptions)` to `(context.Context, string, string, *MonitorsClientBeginDeleteOptions)`
- Function `*MonitorsClient.BeginDelete` return value(s) have been changed from `(MonitorsDeletePollerResponse, error)` to `(MonitorsClientDeletePollerResponse, error)`
- Function `*MonitorsClient.Get` parameter(s) have been changed from `(context.Context, string, string, *MonitorsGetOptions)` to `(context.Context, string, string, *MonitorsClientGetOptions)`
- Function `*MonitorsClient.Get` return value(s) have been changed from `(MonitorsGetResponse, error)` to `(MonitorsClientGetResponse, error)`
- Function `*MonitorsClient.Update` parameter(s) have been changed from `(context.Context, string, string, *MonitorsUpdateOptions)` to `(context.Context, string, string, *MonitorsClientUpdateOptions)`
- Function `*MonitorsClient.Update` return value(s) have been changed from `(MonitorsUpdateResponse, error)` to `(MonitorsClientUpdateResponse, error)`
- Function `*VMCollectionClient.Update` parameter(s) have been changed from `(context.Context, string, string, *VMCollectionUpdateOptions)` to `(context.Context, string, string, *VMCollectionClientUpdateOptions)`
- Function `*VMCollectionClient.Update` return value(s) have been changed from `(VMCollectionUpdateResponse, error)` to `(VMCollectionClientUpdateResponse, error)`
- Function `*VMIngestionClient.Details` parameter(s) have been changed from `(context.Context, string, string, *VMIngestionDetailsOptions)` to `(context.Context, string, string, *VMIngestionClientDetailsOptions)`
- Function `*VMIngestionClient.Details` return value(s) have been changed from `(VMIngestionDetailsResponseEnvelope, error)` to `(VMIngestionClientDetailsResponse, error)`
- Function `*TagRulesClient.CreateOrUpdate` parameter(s) have been changed from `(context.Context, string, string, string, *TagRulesCreateOrUpdateOptions)` to `(context.Context, string, string, string, *TagRulesClientCreateOrUpdateOptions)`
- Function `*TagRulesClient.CreateOrUpdate` return value(s) have been changed from `(TagRulesCreateOrUpdateResponse, error)` to `(TagRulesClientCreateOrUpdateResponse, error)`
- Type of `MonitorProperties.ElasticProperties` has been changed from `*ElasticProperties` to `*Properties`
- Function `*MonitoredResourcesListPager.PageResponse` has been removed
- Function `*MonitorsCreatePoller.Poll` has been removed
- Function `*TagRulesDeletePoller.Done` has been removed
- Function `*OperationsListPager.Err` has been removed
- Function `MonitorsCreatePollerResponse.PollUntilDone` has been removed
- Function `*TagRulesListPager.NextPage` has been removed
- Function `MonitorsDeletePollerResponse.PollUntilDone` has been removed
- Function `*OperationsListPager.PageResponse` has been removed
- Function `*MonitorsCreatePoller.ResumeToken` has been removed
- Function `*MonitorsDeletePoller.Done` has been removed
- Function `*MonitorsDeletePollerResponse.Resume` has been removed
- Function `*VMHostListPager.NextPage` has been removed
- Function `*MonitorsDeletePoller.Poll` has been removed
- Function `*MonitorsDeletePoller.ResumeToken` has been removed
- Function `ElasticMonitorResource.MarshalJSON` has been removed
- Function `*OperationsListPager.NextPage` has been removed
- Function `*MonitorsCreatePollerResponse.Resume` has been removed
- Function `*MonitorsCreatePoller.FinalResponse` has been removed
- Function `*MonitoredResourcesListPager.NextPage` has been removed
- Function `*MonitorsListByResourceGroupPager.NextPage` has been removed
- Function `ElasticMonitorResourceUpdateParameters.MarshalJSON` has been removed
- Function `*MonitorsListPager.Err` has been removed
- Function `*MonitorsListByResourceGroupPager.PageResponse` has been removed
- Function `*TagRulesDeletePoller.FinalResponse` has been removed
- Function `*TagRulesListPager.PageResponse` has been removed
- Function `*TagRulesDeletePollerResponse.Resume` has been removed
- Function `*MonitoredResourcesListPager.Err` has been removed
- Function `*TagRulesDeletePoller.ResumeToken` has been removed
- Function `*MonitorsDeletePoller.FinalResponse` has been removed
- Function `*MonitorsListPager.PageResponse` has been removed
- Function `TagRulesDeletePollerResponse.PollUntilDone` has been removed
- Function `*MonitorsListPager.NextPage` has been removed
- Function `*VMHostListPager.Err` has been removed
- Function `*VMHostListPager.PageResponse` has been removed
- Function `*MonitorsCreatePoller.Done` has been removed
- Function `ElasticMonitorResourceListResponse.MarshalJSON` has been removed
- Function `*TagRulesDeletePoller.Poll` has been removed
- Function `*TagRulesListPager.Err` has been removed
- Function `*MonitorsListByResourceGroupPager.Err` has been removed
- Function `ResourceProviderDefaultErrorResponse.Error` has been removed
- Struct `DeploymentInfoListOptions` has been removed
- Struct `DeploymentInfoListResponse` has been removed
- Struct `DeploymentInfoListResult` has been removed
- Struct `ElasticCloudDeployment` has been removed
- Struct `ElasticCloudUser` has been removed
- Struct `ElasticMonitorResource` has been removed
- Struct `ElasticMonitorResourceListResponse` has been removed
- Struct `ElasticMonitorResourceUpdateParameters` has been removed
- Struct `ElasticProperties` has been removed
- Struct `MonitoredResourcesListOptions` has been removed
- Struct `MonitoredResourcesListPager` has been removed
- Struct `MonitoredResourcesListResponse` has been removed
- Struct `MonitoredResourcesListResult` has been removed
- Struct `MonitorsBeginCreateOptions` has been removed
- Struct `MonitorsBeginDeleteOptions` has been removed
- Struct `MonitorsCreatePoller` has been removed
- Struct `MonitorsCreatePollerResponse` has been removed
- Struct `MonitorsCreateResponse` has been removed
- Struct `MonitorsCreateResult` has been removed
- Struct `MonitorsDeletePoller` has been removed
- Struct `MonitorsDeletePollerResponse` has been removed
- Struct `MonitorsDeleteResponse` has been removed
- Struct `MonitorsGetOptions` has been removed
- Struct `MonitorsGetResponse` has been removed
- Struct `MonitorsGetResult` has been removed
- Struct `MonitorsListByResourceGroupOptions` has been removed
- Struct `MonitorsListByResourceGroupPager` has been removed
- Struct `MonitorsListByResourceGroupResponse` has been removed
- Struct `MonitorsListByResourceGroupResult` has been removed
- Struct `MonitorsListOptions` has been removed
- Struct `MonitorsListPager` has been removed
- Struct `MonitorsListResponse` has been removed
- Struct `MonitorsListResult` has been removed
- Struct `MonitorsUpdateOptions` has been removed
- Struct `MonitorsUpdateResponse` has been removed
- Struct `MonitorsUpdateResult` has been removed
- Struct `OperationsListOptions` has been removed
- Struct `OperationsListPager` has been removed
- Struct `OperationsListResponse` has been removed
- Struct `OperationsListResult` has been removed
- Struct `TagRulesBeginDeleteOptions` has been removed
- Struct `TagRulesCreateOrUpdateOptions` has been removed
- Struct `TagRulesCreateOrUpdateResponse` has been removed
- Struct `TagRulesCreateOrUpdateResult` has been removed
- Struct `TagRulesDeletePoller` has been removed
- Struct `TagRulesDeletePollerResponse` has been removed
- Struct `TagRulesDeleteResponse` has been removed
- Struct `TagRulesGetOptions` has been removed
- Struct `TagRulesGetResponse` has been removed
- Struct `TagRulesGetResult` has been removed
- Struct `TagRulesListOptions` has been removed
- Struct `TagRulesListPager` has been removed
- Struct `TagRulesListResponse` has been removed
- Struct `TagRulesListResult` has been removed
- Struct `VMCollectionUpdateOptions` has been removed
- Struct `VMCollectionUpdateResponse` has been removed
- Struct `VMHostListOptions` has been removed
- Struct `VMHostListPager` has been removed
- Struct `VMHostListResponseEnvelope` has been removed
- Struct `VMHostListResult` has been removed
- Struct `VMIngestionDetailsOptions` has been removed
- Struct `VMIngestionDetailsResponseEnvelope` has been removed
- Struct `VMIngestionDetailsResult` has been removed
- Field `InnerError` of struct `ResourceProviderDefaultErrorResponse` has been removed

### Features Added

- New function `*MonitorsClientListByResourceGroupPager.NextPage(context.Context) bool`
- New function `MonitorResourceListResponse.MarshalJSON() ([]byte, error)`
- New function `*MonitorsClientDeletePoller.ResumeToken() (string, error)`
- New function `*MonitorsClientCreatePollerResponse.Resume(context.Context, *MonitorsClient, string) error`
- New function `MonitorResource.MarshalJSON() ([]byte, error)`
- New function `*TagRulesClientListPager.NextPage(context.Context) bool`
- New function `*MonitorsClientCreatePoller.FinalResponse(context.Context) (MonitorsClientCreateResponse, error)`
- New function `MonitorsClientCreatePollerResponse.PollUntilDone(context.Context, time.Duration) (MonitorsClientCreateResponse, error)`
- New function `*MonitorsClientCreatePoller.ResumeToken() (string, error)`
- New function `*TagRulesClientListPager.PageResponse() TagRulesClientListResponse`
- New function `*MonitoredResourcesClientListPager.PageResponse() MonitoredResourcesClientListResponse`
- New function `*MonitorsClientListPager.NextPage(context.Context) bool`
- New function `*MonitorsClientListByResourceGroupPager.PageResponse() MonitorsClientListByResourceGroupResponse`
- New function `*TagRulesClientDeletePoller.FinalResponse(context.Context) (TagRulesClientDeleteResponse, error)`
- New function `*TagRulesClientDeletePoller.Done() bool`
- New function `*MonitoredResourcesClientListPager.Err() error`
- New function `*MonitorsClientListByResourceGroupPager.Err() error`
- New function `*MonitorsClientDeletePoller.FinalResponse(context.Context) (MonitorsClientDeleteResponse, error)`
- New function `*TagRulesClientDeletePoller.ResumeToken() (string, error)`
- New function `*MonitorsClientDeletePoller.Done() bool`
- New function `*OperationsClientListPager.NextPage(context.Context) bool`
- New function `*VMHostClientListPager.Err() error`
- New function `*TagRulesClientDeletePollerResponse.Resume(context.Context, *TagRulesClient, string) error`
- New function `*OperationsClientListPager.Err() error`
- New function `*VMHostClientListPager.NextPage(context.Context) bool`
- New function `*MonitorsClientListPager.PageResponse() MonitorsClientListResponse`
- New function `MonitorResourceUpdateParameters.MarshalJSON() ([]byte, error)`
- New function `TagRulesClientDeletePollerResponse.PollUntilDone(context.Context, time.Duration) (TagRulesClientDeleteResponse, error)`
- New function `*MonitoredResourcesClientListPager.NextPage(context.Context) bool`
- New function `*VMHostClientListPager.PageResponse() VMHostClientListResponse`
- New function `*MonitorsClientDeletePoller.Poll(context.Context) (*http.Response, error)`
- New function `MonitorsClientDeletePollerResponse.PollUntilDone(context.Context, time.Duration) (MonitorsClientDeleteResponse, error)`
- New function `*OperationsClientListPager.PageResponse() OperationsClientListResponse`
- New function `*MonitorsClientDeletePollerResponse.Resume(context.Context, *MonitorsClient, string) error`
- New function `*MonitorsClientListPager.Err() error`
- New function `*TagRulesClientDeletePoller.Poll(context.Context) (*http.Response, error)`
- New function `*MonitorsClientCreatePoller.Poll(context.Context) (*http.Response, error)`
- New function `*MonitorsClientCreatePoller.Done() bool`
- New function `*TagRulesClientListPager.Err() error`
- New struct `CloudDeployment`
- New struct `CloudUser`
- New struct `DeploymentInfoClientListOptions`
- New struct `DeploymentInfoClientListResponse`
- New struct `DeploymentInfoClientListResult`
- New struct `MonitorResource`
- New struct `MonitorResourceListResponse`
- New struct `MonitorResourceUpdateParameters`
- New struct `MonitoredResourcesClientListOptions`
- New struct `MonitoredResourcesClientListPager`
- New struct `MonitoredResourcesClientListResponse`
- New struct `MonitoredResourcesClientListResult`
- New struct `MonitorsClientBeginCreateOptions`
- New struct `MonitorsClientBeginDeleteOptions`
- New struct `MonitorsClientCreatePoller`
- New struct `MonitorsClientCreatePollerResponse`
- New struct `MonitorsClientCreateResponse`
- New struct `MonitorsClientCreateResult`
- New struct `MonitorsClientDeletePoller`
- New struct `MonitorsClientDeletePollerResponse`
- New struct `MonitorsClientDeleteResponse`
- New struct `MonitorsClientGetOptions`
- New struct `MonitorsClientGetResponse`
- New struct `MonitorsClientGetResult`
- New struct `MonitorsClientListByResourceGroupOptions`
- New struct `MonitorsClientListByResourceGroupPager`
- New struct `MonitorsClientListByResourceGroupResponse`
- New struct `MonitorsClientListByResourceGroupResult`
- New struct `MonitorsClientListOptions`
- New struct `MonitorsClientListPager`
- New struct `MonitorsClientListResponse`
- New struct `MonitorsClientListResult`
- New struct `MonitorsClientUpdateOptions`
- New struct `MonitorsClientUpdateResponse`
- New struct `MonitorsClientUpdateResult`
- New struct `OperationsClientListOptions`
- New struct `OperationsClientListPager`
- New struct `OperationsClientListResponse`
- New struct `OperationsClientListResult`
- New struct `Properties`
- New struct `TagRulesClientBeginDeleteOptions`
- New struct `TagRulesClientCreateOrUpdateOptions`
- New struct `TagRulesClientCreateOrUpdateResponse`
- New struct `TagRulesClientCreateOrUpdateResult`
- New struct `TagRulesClientDeletePoller`
- New struct `TagRulesClientDeletePollerResponse`
- New struct `TagRulesClientDeleteResponse`
- New struct `TagRulesClientGetOptions`
- New struct `TagRulesClientGetResponse`
- New struct `TagRulesClientGetResult`
- New struct `TagRulesClientListOptions`
- New struct `TagRulesClientListPager`
- New struct `TagRulesClientListResponse`
- New struct `TagRulesClientListResult`
- New struct `VMCollectionClientUpdateOptions`
- New struct `VMCollectionClientUpdateResponse`
- New struct `VMHostClientListOptions`
- New struct `VMHostClientListPager`
- New struct `VMHostClientListResponse`
- New struct `VMHostClientListResult`
- New struct `VMIngestionClientDetailsOptions`
- New struct `VMIngestionClientDetailsResponse`
- New struct `VMIngestionClientDetailsResult`
- New field `Error` in struct `ResourceProviderDefaultErrorResponse`


## 0.1.0 (2021-12-07)

- Initial preview release.
