# \ConfigGroupApi

All URIs are relative to *https://api.lucfish.com/configuration/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchDeleteGroups**](ConfigGroupApi.md#BatchDeleteGroups) | **Post** /configGroup/batchDelete | 批量删除组
[**BatchRetrieveGroups**](ConfigGroupApi.md#BatchRetrieveGroups) | **Post** /configGroup/batchRetrieve | 批量查询组
[**CreateGroup**](ConfigGroupApi.md#CreateGroup) | **Post** /configGroup | 创建组
[**CursorGroups**](ConfigGroupApi.md#CursorGroups) | **Post** /configGroup/cursor | Cursor查询组
[**DeleteGroup**](ConfigGroupApi.md#DeleteGroup) | **Delete** /configGroup/{id} | 删除组
[**GetGroup**](ConfigGroupApi.md#GetGroup) | **Get** /configGroup/{id} | 查询组
[**PageGroups**](ConfigGroupApi.md#PageGroups) | **Post** /configGroup/page | Page查询组
[**RetrieveByNameGroup**](ConfigGroupApi.md#RetrieveByNameGroup) | **Post** /configGroup/retrieveByName | 查询组
[**UpdateGroup**](ConfigGroupApi.md#UpdateGroup) | **Put** /configGroup | 更新组



## BatchDeleteGroups

> DeleteResponse BatchDeleteGroups(ctx, optional)

批量删除组

批量删除组通过组Id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BatchDeleteGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BatchDeleteGroupsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **idsReq** | [**optional.Interface of IdsReq**](IdsReq.md)|  | 

### Return type

[**DeleteResponse**](DeleteResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchRetrieveGroups

> BatchRetrieveConfigGroupsResponse BatchRetrieveGroups(ctx, optional)

批量查询组

批量查询组通过组Id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BatchRetrieveGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BatchRetrieveGroupsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **idsReq** | [**optional.Interface of IdsReq**](IdsReq.md)|  | 

### Return type

[**BatchRetrieveConfigGroupsResponse**](BatchRetrieveConfigGroupsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateGroup

> CreateConfigGroupResponse CreateGroup(ctx, optional)

创建组

创建组

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateGroupOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createConfigGroupReq** | [**optional.Interface of CreateConfigGroupReq**](CreateConfigGroupReq.md)|  | 

### Return type

[**CreateConfigGroupResponse**](CreateConfigGroupResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CursorGroups

> CursorConfigGroupsResponse CursorGroups(ctx, optional)

Cursor查询组

Cursor查询组

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CursorGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CursorGroupsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursorQuery** | [**optional.Interface of CursorQuery**](CursorQuery.md)|  | 

### Return type

[**CursorConfigGroupsResponse**](CursorConfigGroupsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGroup

> DeleteResponse DeleteGroup(ctx, id)

删除组

删除组

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| 删除组 | 

### Return type

[**DeleteResponse**](DeleteResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroup

> GetConfigGroupResponse GetGroup(ctx, id)

查询组

查询组通过组ID

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| 查询组通过组ID | 

### Return type

[**GetConfigGroupResponse**](GetConfigGroupResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PageGroups

> PageConfigGroupsResponse PageGroups(ctx, optional)

Page查询组

Page查询组

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PageGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PageGroupsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageQuery** | [**optional.Interface of PageQuery**](PageQuery.md)|  | 

### Return type

[**PageConfigGroupsResponse**](PageConfigGroupsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveByNameGroup

> RetrieveConfigGroupByNameResponse RetrieveByNameGroup(ctx, name)

查询组

查询组通过Name

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**| 查询组通过Name | 

### Return type

[**RetrieveConfigGroupByNameResponse**](RetrieveConfigGroupByNameResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGroup

> UpdateConfigGroupResponse UpdateGroup(ctx, optional)

更新组

更新组，需要全部的组信息

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UpdateGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateGroupOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateConfigGroupReq** | [**optional.Interface of UpdateConfigGroupReq**](UpdateConfigGroupReq.md)|  | 

### Return type

[**UpdateConfigGroupResponse**](UpdateConfigGroupResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

