# \ConfigGroupApi

All URIs are relative to *https://api.lucfish.com/configuration/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConfigGroupBatchDeletePost**](ConfigGroupApi.md#ConfigGroupBatchDeletePost) | **Post** /configGroup/batchDelete | 批量删除组
[**ConfigGroupBatchRetrievePost**](ConfigGroupApi.md#ConfigGroupBatchRetrievePost) | **Post** /configGroup/batchRetrieve | 批量查询组
[**ConfigGroupCursorPost**](ConfigGroupApi.md#ConfigGroupCursorPost) | **Post** /configGroup/cursor | Cursor查询组
[**ConfigGroupIdDelete**](ConfigGroupApi.md#ConfigGroupIdDelete) | **Delete** /configGroup/{id} | 删除组
[**ConfigGroupIdGet**](ConfigGroupApi.md#ConfigGroupIdGet) | **Get** /configGroup/{id} | 查询组
[**ConfigGroupPagePost**](ConfigGroupApi.md#ConfigGroupPagePost) | **Post** /configGroup/page | Page查询组
[**ConfigGroupPost**](ConfigGroupApi.md#ConfigGroupPost) | **Post** /configGroup | 创建组
[**ConfigGroupPut**](ConfigGroupApi.md#ConfigGroupPut) | **Put** /configGroup | 更新组
[**ConfigGroupRetrieveByNamePost**](ConfigGroupApi.md#ConfigGroupRetrieveByNamePost) | **Post** /configGroup/retrieveByName | 查询组



## ConfigGroupBatchDeletePost

> DeleteRsp ConfigGroupBatchDeletePost(ctx, optional)

批量删除组

批量删除组通过组Id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConfigGroupBatchDeletePostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ConfigGroupBatchDeletePostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **idsReq** | [**optional.Interface of IdsReq**](IdsReq.md)|  | 

### Return type

[**DeleteRsp**](DeleteRsp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfigGroupBatchRetrievePost

> GetConfigGroupsRsp ConfigGroupBatchRetrievePost(ctx, optional)

批量查询组

批量查询组通过组Id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConfigGroupBatchRetrievePostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ConfigGroupBatchRetrievePostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **idsReq** | [**optional.Interface of IdsReq**](IdsReq.md)|  | 

### Return type

[**GetConfigGroupsRsp**](GetConfigGroupsRsp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfigGroupCursorPost

> CursorConfigGroupsRsp ConfigGroupCursorPost(ctx, optional)

Cursor查询组

Cursor查询组

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConfigGroupCursorPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ConfigGroupCursorPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursorQuery** | [**optional.Interface of CursorQuery**](CursorQuery.md)|  | 

### Return type

[**CursorConfigGroupsRsp**](CursorConfigGroupsRsp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfigGroupIdDelete

> DeleteRsp ConfigGroupIdDelete(ctx, id)

删除组

删除组

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| 删除组 | 

### Return type

[**DeleteRsp**](DeleteRsp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfigGroupIdGet

> GetConfigGroupRsp ConfigGroupIdGet(ctx, id)

查询组

查询组通过组ID

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| 查询组通过组ID | 

### Return type

[**GetConfigGroupRsp**](GetConfigGroupRsp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfigGroupPagePost

> PageConfigGroupsRsp ConfigGroupPagePost(ctx, optional)

Page查询组

Page查询组

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConfigGroupPagePostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ConfigGroupPagePostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageQuery** | [**optional.Interface of PageQuery**](PageQuery.md)|  | 

### Return type

[**PageConfigGroupsRsp**](PageConfigGroupsRsp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfigGroupPost

> CreateConfigGroupRsp ConfigGroupPost(ctx, optional)

创建组

创建组

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConfigGroupPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ConfigGroupPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createConfigGroupReq** | [**optional.Interface of CreateConfigGroupReq**](CreateConfigGroupReq.md)|  | 

### Return type

[**CreateConfigGroupRsp**](CreateConfigGroupRsp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfigGroupPut

> UpdateConfigGroupRsp ConfigGroupPut(ctx, optional)

更新组

更新组，需要全部的组信息

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConfigGroupPutOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ConfigGroupPutOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateConfigGroupReq** | [**optional.Interface of UpdateConfigGroupReq**](UpdateConfigGroupReq.md)|  | 

### Return type

[**UpdateConfigGroupRsp**](UpdateConfigGroupRsp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfigGroupRetrieveByNamePost

> RetrieveConfigGroupByNameRsp ConfigGroupRetrieveByNamePost(ctx, name)

查询组

查询组通过Name

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**| 查询组通过Name | 

### Return type

[**RetrieveConfigGroupByNameRsp**](RetrieveConfigGroupByNameRsp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

