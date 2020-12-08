# \ConfigTemplateApi

All URIs are relative to *https://api.lucfish.com/configuration/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConfigTemplateBatchDeletePost**](ConfigTemplateApi.md#ConfigTemplateBatchDeletePost) | **Post** /configTemplate/batchDelete | 批量删除配置模版
[**ConfigTemplateBatchRetrieveByGroupPost**](ConfigTemplateApi.md#ConfigTemplateBatchRetrieveByGroupPost) | **Post** /configTemplate/batchRetrieveByGroup | 查询配置模版
[**ConfigTemplateBatchRetrieveByKeysPost**](ConfigTemplateApi.md#ConfigTemplateBatchRetrieveByKeysPost) | **Post** /configTemplate/batchRetrieveByKeys | 查询配置模版
[**ConfigTemplateBatchRetrievePost**](ConfigTemplateApi.md#ConfigTemplateBatchRetrievePost) | **Post** /configTemplate/batchRetrieve | 批量查询配置模版
[**ConfigTemplateCursorPost**](ConfigTemplateApi.md#ConfigTemplateCursorPost) | **Post** /configTemplate/cursor | Cursor查询配置模版
[**ConfigTemplateIdDelete**](ConfigTemplateApi.md#ConfigTemplateIdDelete) | **Delete** /configTemplate/{id} | 删除配置模版
[**ConfigTemplateIdGet**](ConfigTemplateApi.md#ConfigTemplateIdGet) | **Get** /configTemplate/{id} | 查询配置模版
[**ConfigTemplatePagePost**](ConfigTemplateApi.md#ConfigTemplatePagePost) | **Post** /configTemplate/page | Page查询配置模版
[**ConfigTemplatePost**](ConfigTemplateApi.md#ConfigTemplatePost) | **Post** /configTemplate | 创建配置模版
[**ConfigTemplateRetrieveByKeyPost**](ConfigTemplateApi.md#ConfigTemplateRetrieveByKeyPost) | **Post** /configTemplate/retrieveByKey | 查询配置模版



## ConfigTemplateBatchDeletePost

> InlineResponse2001 ConfigTemplateBatchDeletePost(ctx, optional)

批量删除配置模版

批量删除配置模版通过配置Id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConfigTemplateBatchDeletePostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ConfigTemplateBatchDeletePostOpts struct


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


## ConfigTemplateBatchRetrieveByGroupPost

> InlineResponse2006 ConfigTemplateBatchRetrieveByGroupPost(ctx, groupName)

查询配置模版

查询配置模版通过配置模版GroupID

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupName** | **string**| 分组标签名 | 

### Return type

[**InlineResponse2006**](inline_response_200_6.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfigTemplateBatchRetrieveByKeysPost

> InlineResponse2006 ConfigTemplateBatchRetrieveByKeysPost(ctx, optional)

查询配置模版

查询配置通过配置Keys

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConfigTemplateBatchRetrieveByKeysPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ConfigTemplateBatchRetrieveByKeysPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchRetrieveConfigTemplateByKeysReq** | [**optional.Interface of BatchRetrieveConfigTemplateByKeysReq**](BatchRetrieveConfigTemplateByKeysReq.md)|  | 

### Return type

[**InlineResponse2006**](inline_response_200_6.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfigTemplateBatchRetrievePost

> InlineResponse2006 ConfigTemplateBatchRetrievePost(ctx, optional)

批量查询配置模版

批量查询配置模版通过配置Id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConfigTemplateBatchRetrievePostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ConfigTemplateBatchRetrievePostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **idsReq** | [**optional.Interface of IdsReq**](IdsReq.md)|  | 

### Return type

[**InlineResponse2006**](inline_response_200_6.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfigTemplateCursorPost

> InlineResponse2008 ConfigTemplateCursorPost(ctx, optional)

Cursor查询配置模版

Cursor查询配置模版

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConfigTemplateCursorPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ConfigTemplateCursorPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursorQuery** | [**optional.Interface of CursorQuery**](CursorQuery.md)|  | 

### Return type

[**InlineResponse2008**](inline_response_200_8.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfigTemplateIdDelete

> InlineResponse2001 ConfigTemplateIdDelete(ctx, id)

删除配置模版

删除配置模版

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| 删除配置模版 | 

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


## ConfigTemplateIdGet

> InlineResponse2005 ConfigTemplateIdGet(ctx, id)

查询配置模版

查询配置通过配置模版ID

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| 查询配置通过配置ID | 

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


## ConfigTemplatePagePost

> InlineResponse2007 ConfigTemplatePagePost(ctx, optional)

Page查询配置模版

Page查询配置模版

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConfigTemplatePagePostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ConfigTemplatePagePostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageQuery** | [**optional.Interface of PageQuery**](PageQuery.md)|  | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfigTemplatePost

> InlineResponse2005 ConfigTemplatePost(ctx, optional)

创建配置模版

创建配置模版

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConfigTemplatePostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ConfigTemplatePostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createConfigTemplateReq** | [**optional.Interface of CreateConfigTemplateReq**](CreateConfigTemplateReq.md)|  | 

### Return type

[**InlineResponse2005**](inline_response_200_5.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfigTemplateRetrieveByKeyPost

> InlineResponse2005 ConfigTemplateRetrieveByKeyPost(ctx, key)

查询配置模版

查询配置模版通过配置Key

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string**| 主键Key | 

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

