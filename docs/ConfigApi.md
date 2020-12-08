# \ConfigApi

All URIs are relative to *https://api.lucfish.com/configuration/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConfigBatchDeletePost**](ConfigApi.md#ConfigBatchDeletePost) | **Post** /config/batchDelete | 批量删除配置
[**ConfigBatchRetrieveByKeysPost**](ConfigApi.md#ConfigBatchRetrieveByKeysPost) | **Post** /config/batchRetrieveByKeys | 查询配置
[**ConfigBatchRetrieveByResourcesPost**](ConfigApi.md#ConfigBatchRetrieveByResourcesPost) | **Post** /config/batchRetrieveByResources | 查询配置
[**ConfigBatchRetrievePost**](ConfigApi.md#ConfigBatchRetrievePost) | **Post** /config/batchRetrieve | 批量查询配置
[**ConfigCursorPost**](ConfigApi.md#ConfigCursorPost) | **Post** /config/cursor | Cursor查询配置
[**ConfigIdDelete**](ConfigApi.md#ConfigIdDelete) | **Delete** /config/{id} | 删除配置
[**ConfigIdGet**](ConfigApi.md#ConfigIdGet) | **Get** /config/{id} | 查询配置
[**ConfigPagePost**](ConfigApi.md#ConfigPagePost) | **Post** /config/page | Page查询配置
[**ConfigPost**](ConfigApi.md#ConfigPost) | **Post** /config | 创建配置
[**ConfigPut**](ConfigApi.md#ConfigPut) | **Put** /config | 更新配置
[**ConfigRetrieveByKeyPost**](ConfigApi.md#ConfigRetrieveByKeyPost) | **Post** /config/retrieveByKey | 查询配置
[**ConfigTemplatePut**](ConfigApi.md#ConfigTemplatePut) | **Put** /configTemplate | 更新配置模版



## ConfigBatchDeletePost

> InlineResponse2001 ConfigBatchDeletePost(ctx, optional)

批量删除配置

批量删除配置通过配置Id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConfigBatchDeletePostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ConfigBatchDeletePostOpts struct


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


## ConfigBatchRetrieveByKeysPost

> InlineResponse2002 ConfigBatchRetrieveByKeysPost(ctx, optional)

查询配置

查询配置通过配置Keys和资源ID

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConfigBatchRetrieveByKeysPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ConfigBatchRetrieveByKeysPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchRetrieveByKeysReq** | [**optional.Interface of BatchRetrieveByKeysReq**](BatchRetrieveByKeysReq.md)|  | 

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfigBatchRetrieveByResourcesPost

> InlineResponse2002 ConfigBatchRetrieveByResourcesPost(ctx, optional)

查询配置

查询配置通过配置GroupID

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConfigBatchRetrieveByResourcesPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ConfigBatchRetrieveByResourcesPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchRetrieveByResourcesReq** | [**optional.Interface of BatchRetrieveByResourcesReq**](BatchRetrieveByResourcesReq.md)|  | 

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfigBatchRetrievePost

> InlineResponse2002 ConfigBatchRetrievePost(ctx, optional)

批量查询配置

批量查询配置通过配置Id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConfigBatchRetrievePostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ConfigBatchRetrievePostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **idsReq** | [**optional.Interface of IdsReq**](IdsReq.md)|  | 

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfigCursorPost

> InlineResponse2004 ConfigCursorPost(ctx, optional)

Cursor查询配置

Cursor查询配置

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConfigCursorPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ConfigCursorPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursorQuery** | [**optional.Interface of CursorQuery**](CursorQuery.md)|  | 

### Return type

[**InlineResponse2004**](inline_response_200_4.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfigIdDelete

> InlineResponse2001 ConfigIdDelete(ctx, id)

删除配置

删除配置

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| 删除配置 | 

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


## ConfigIdGet

> InlineResponse200 ConfigIdGet(ctx, id)

查询配置

查询配置通过配置ID

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| 查询配置通过配置ID | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfigPagePost

> InlineResponse2003 ConfigPagePost(ctx, optional)

Page查询配置

Page查询配置

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConfigPagePostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ConfigPagePostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageQuery** | [**optional.Interface of PageQuery**](PageQuery.md)|  | 

### Return type

[**InlineResponse2003**](inline_response_200_3.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfigPost

> InlineResponse200 ConfigPost(ctx, optional)

创建配置

创建配置

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConfigPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ConfigPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createConfigReq** | [**optional.Interface of CreateConfigReq**](CreateConfigReq.md)|  | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfigPut

> InlineResponse200 ConfigPut(ctx, optional)

更新配置

更新配置，需要全部的配置信息

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConfigPutOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ConfigPutOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateConfigReq** | [**optional.Interface of UpdateConfigReq**](UpdateConfigReq.md)|  | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfigRetrieveByKeyPost

> InlineResponse200 ConfigRetrieveByKeyPost(ctx, optional)

查询配置

查询配置通过配置Key和资源ID

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConfigRetrieveByKeyPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ConfigRetrieveByKeyPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **retrieveConfigByKeyReq** | [**optional.Interface of RetrieveConfigByKeyReq**](RetrieveConfigByKeyReq.md)|  | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfigTemplatePut

> InlineResponse2005 ConfigTemplatePut(ctx, optional)

更新配置模版

更新配置模版，需要全部的配置模版信息

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConfigTemplatePutOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ConfigTemplatePutOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateConfigTemplateReq** | [**optional.Interface of UpdateConfigTemplateReq**](UpdateConfigTemplateReq.md)|  | 

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

