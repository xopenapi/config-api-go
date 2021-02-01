# \ConfigTemplateApi

All URIs are relative to *https://api.lucfish.com/configuration/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchDeleteConfigTemplates**](ConfigTemplateApi.md#BatchDeleteConfigTemplates) | **Post** /configTemplate/batchDelete | 批量删除配置模版
[**BatchRetrieveByGroupConfigTemplates**](ConfigTemplateApi.md#BatchRetrieveByGroupConfigTemplates) | **Post** /configTemplate/batchRetrieveByGroup | 查询配置模版
[**BatchRetrieveByKeysConfigTemplates**](ConfigTemplateApi.md#BatchRetrieveByKeysConfigTemplates) | **Post** /configTemplate/batchRetrieveByKeys | 查询配置模版
[**BatchRetrieveConfigTemplates**](ConfigTemplateApi.md#BatchRetrieveConfigTemplates) | **Post** /configTemplate/batchRetrieve | 批量查询配置模版
[**CreateConfigTemplate**](ConfigTemplateApi.md#CreateConfigTemplate) | **Post** /configTemplate | 创建配置模版
[**CursorConfigTemplates**](ConfigTemplateApi.md#CursorConfigTemplates) | **Post** /configTemplate/cursor | Cursor查询配置模版
[**DeleteConfigTemplate**](ConfigTemplateApi.md#DeleteConfigTemplate) | **Delete** /configTemplate/{id} | 删除配置模版
[**GetConfigTemplate**](ConfigTemplateApi.md#GetConfigTemplate) | **Get** /configTemplate/{id} | 查询配置模版
[**PageConfigTemplate**](ConfigTemplateApi.md#PageConfigTemplate) | **Post** /configTemplate/page | Page查询配置模版
[**RetrieveByKeyConfigTemplate**](ConfigTemplateApi.md#RetrieveByKeyConfigTemplate) | **Post** /configTemplate/retrieveByKey | 查询配置模版



## BatchDeleteConfigTemplates

> DeleteResponse BatchDeleteConfigTemplates(ctx, optional)

批量删除配置模版

批量删除配置模版通过配置Id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BatchDeleteConfigTemplatesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BatchDeleteConfigTemplatesOpts struct


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


## BatchRetrieveByGroupConfigTemplates

> BatchRetrieveConfigTemplatesByGroupResponse BatchRetrieveByGroupConfigTemplates(ctx, groupName)

查询配置模版

查询配置模版通过配置模版GroupID

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupName** | **string**| 分组标签名 | 

### Return type

[**BatchRetrieveConfigTemplatesByGroupResponse**](BatchRetrieveConfigTemplatesByGroupResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchRetrieveByKeysConfigTemplates

> BatchRetrieveConfigTemplatesByKeysResponse BatchRetrieveByKeysConfigTemplates(ctx, optional)

查询配置模版

查询配置通过配置Keys

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BatchRetrieveByKeysConfigTemplatesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BatchRetrieveByKeysConfigTemplatesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchRetrieveConfigTemplateByKeysReq** | [**optional.Interface of BatchRetrieveConfigTemplateByKeysReq**](BatchRetrieveConfigTemplateByKeysReq.md)|  | 

### Return type

[**BatchRetrieveConfigTemplatesByKeysResponse**](BatchRetrieveConfigTemplatesByKeysResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchRetrieveConfigTemplates

> BatchRetrieveConfigTemplatesResponse BatchRetrieveConfigTemplates(ctx, optional)

批量查询配置模版

批量查询配置模版通过配置Id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BatchRetrieveConfigTemplatesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BatchRetrieveConfigTemplatesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **idsReq** | [**optional.Interface of IdsReq**](IdsReq.md)|  | 

### Return type

[**BatchRetrieveConfigTemplatesResponse**](BatchRetrieveConfigTemplatesResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateConfigTemplate

> CreateConfigTemplateResponse CreateConfigTemplate(ctx, optional)

创建配置模版

创建配置模版

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateConfigTemplateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateConfigTemplateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createConfigTemplateReq** | [**optional.Interface of CreateConfigTemplateReq**](CreateConfigTemplateReq.md)|  | 

### Return type

[**CreateConfigTemplateResponse**](CreateConfigTemplateResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CursorConfigTemplates

> CursorConfigTemplatesResponse CursorConfigTemplates(ctx, optional)

Cursor查询配置模版

Cursor查询配置模版

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CursorConfigTemplatesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CursorConfigTemplatesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursorQuery** | [**optional.Interface of CursorQuery**](CursorQuery.md)|  | 

### Return type

[**CursorConfigTemplatesResponse**](CursorConfigTemplatesResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteConfigTemplate

> DeleteResponse DeleteConfigTemplate(ctx, id)

删除配置模版

删除配置模版

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| 删除配置模版 | 

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


## GetConfigTemplate

> GetConfigTemplateResponse GetConfigTemplate(ctx, id)

查询配置模版

查询配置通过配置模版ID

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| 查询配置通过配置ID | 

### Return type

[**GetConfigTemplateResponse**](GetConfigTemplateResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PageConfigTemplate

> PageConfigTemplatesResponse PageConfigTemplate(ctx, optional)

Page查询配置模版

Page查询配置模版

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PageConfigTemplateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PageConfigTemplateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageQuery** | [**optional.Interface of PageQuery**](PageQuery.md)|  | 

### Return type

[**PageConfigTemplatesResponse**](PageConfigTemplatesResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveByKeyConfigTemplate

> RetrieveConfigTemplateByKeyResponse RetrieveByKeyConfigTemplate(ctx, key)

查询配置模版

查询配置模版通过配置Key

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string**| 主键Key | 

### Return type

[**RetrieveConfigTemplateByKeyResponse**](RetrieveConfigTemplateByKeyResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

