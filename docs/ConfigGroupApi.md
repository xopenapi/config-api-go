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

> InlineResponse2001 ConfigGroupBatchDeletePost(ctx, optional)

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

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfigGroupBatchRetrievePost

> InlineResponse20010 ConfigGroupBatchRetrievePost(ctx, optional)

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

[**InlineResponse20010**](inline_response_200_10.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfigGroupCursorPost

> InlineResponse20012 ConfigGroupCursorPost(ctx, optional)

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

[**InlineResponse20012**](inline_response_200_12.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfigGroupIdDelete

> InlineResponse2001 ConfigGroupIdDelete(ctx, id)

删除组

删除组

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| 删除组 | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfigGroupIdGet

> InlineResponse2009 ConfigGroupIdGet(ctx, id)

查询组

查询组通过组ID

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| 查询组通过组ID | 

### Return type

[**InlineResponse2009**](inline_response_200_9.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfigGroupPagePost

> InlineResponse20011 ConfigGroupPagePost(ctx, optional)

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

[**InlineResponse20011**](inline_response_200_11.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfigGroupPost

> InlineResponse2009 ConfigGroupPost(ctx, optional)

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

[**InlineResponse2009**](inline_response_200_9.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfigGroupPut

> InlineResponse2009 ConfigGroupPut(ctx, optional)

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

[**InlineResponse2009**](inline_response_200_9.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfigGroupRetrieveByNamePost

> InlineResponse2005 ConfigGroupRetrieveByNamePost(ctx, name)

查询组

查询组通过Name

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**| 查询组通过Name | 

### Return type

[**InlineResponse2005**](inline_response_200_5.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

