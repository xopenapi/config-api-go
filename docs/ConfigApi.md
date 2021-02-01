# \ConfigApi

All URIs are relative to *https://api.lucfish.com/configuration/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchDeleteConfigs**](ConfigApi.md#BatchDeleteConfigs) | **Post** /config/batchDelete | 批量删除配置
[**BatchRetrieveByKeysConfigs**](ConfigApi.md#BatchRetrieveByKeysConfigs) | **Post** /config/batchRetrieveByKeys | 查询配置
[**BatchRetrieveByResourcesConfigs**](ConfigApi.md#BatchRetrieveByResourcesConfigs) | **Post** /config/batchRetrieveByResources | 查询配置
[**BatchRetrieveConfigs**](ConfigApi.md#BatchRetrieveConfigs) | **Post** /config/batchRetrieve | 批量查询配置
[**CreateConfig**](ConfigApi.md#CreateConfig) | **Post** /config | 创建配置
[**CursorConfigs**](ConfigApi.md#CursorConfigs) | **Post** /config/cursor | Cursor查询配置
[**DeleteConfig**](ConfigApi.md#DeleteConfig) | **Delete** /config/{id} | 删除配置
[**GetConfig**](ConfigApi.md#GetConfig) | **Get** /config/{id} | 查询配置
[**PageConfigs**](ConfigApi.md#PageConfigs) | **Post** /config/page | Page查询配置
[**RetrieveByKeyConfig**](ConfigApi.md#RetrieveByKeyConfig) | **Post** /config/retrieveByKey | 查询配置
[**UpdateConfig**](ConfigApi.md#UpdateConfig) | **Put** /config | 更新配置
[**UpdateConfigTemplate**](ConfigApi.md#UpdateConfigTemplate) | **Put** /configTemplate | 更新配置模版



## BatchDeleteConfigs

> DeleteResponse BatchDeleteConfigs(ctx, optional)

批量删除配置

批量删除配置通过配置Id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BatchDeleteConfigsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BatchDeleteConfigsOpts struct


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


## BatchRetrieveByKeysConfigs

> BatchRetrieveConfigsByKeysResponse BatchRetrieveByKeysConfigs(ctx, optional)

查询配置

查询配置通过配置Keys和资源ID

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BatchRetrieveByKeysConfigsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BatchRetrieveByKeysConfigsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchRetrieveByKeysReq** | [**optional.Interface of BatchRetrieveByKeysReq**](BatchRetrieveByKeysReq.md)|  | 

### Return type

[**BatchRetrieveConfigsByKeysResponse**](BatchRetrieveConfigsByKeysResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchRetrieveByResourcesConfigs

> BatchRetrieveConfigsByResourcesResponse BatchRetrieveByResourcesConfigs(ctx, optional)

查询配置

查询配置通过配置GroupID

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BatchRetrieveByResourcesConfigsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BatchRetrieveByResourcesConfigsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchRetrieveByResourcesReq** | [**optional.Interface of BatchRetrieveByResourcesReq**](BatchRetrieveByResourcesReq.md)|  | 

### Return type

[**BatchRetrieveConfigsByResourcesResponse**](BatchRetrieveConfigsByResourcesResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchRetrieveConfigs

> BatchRetrieveConfigsResponse BatchRetrieveConfigs(ctx, optional)

批量查询配置

批量查询配置通过配置Id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BatchRetrieveConfigsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BatchRetrieveConfigsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **idsReq** | [**optional.Interface of IdsReq**](IdsReq.md)|  | 

### Return type

[**BatchRetrieveConfigsResponse**](BatchRetrieveConfigsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateConfig

> CreateConfigResponse CreateConfig(ctx, optional)

创建配置

创建配置

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateConfigOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createConfigReq** | [**optional.Interface of CreateConfigReq**](CreateConfigReq.md)|  | 

### Return type

[**CreateConfigResponse**](CreateConfigResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CursorConfigs

> CursorConfigsResponse CursorConfigs(ctx, optional)

Cursor查询配置

Cursor查询配置

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CursorConfigsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CursorConfigsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursorQuery** | [**optional.Interface of CursorQuery**](CursorQuery.md)|  | 

### Return type

[**CursorConfigsResponse**](CursorConfigsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteConfig

> DeleteResponse DeleteConfig(ctx, id)

删除配置

删除配置

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| 删除配置 | 

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


## GetConfig

> GetConfigResponse GetConfig(ctx, id)

查询配置

查询配置通过配置ID

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| 查询配置通过配置ID | 

### Return type

[**GetConfigResponse**](GetConfigResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PageConfigs

> PageConfigsResponse PageConfigs(ctx, optional)

Page查询配置

Page查询配置

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PageConfigsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PageConfigsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageQuery** | [**optional.Interface of PageQuery**](PageQuery.md)|  | 

### Return type

[**PageConfigsResponse**](PageConfigsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveByKeyConfig

> RetrieveConfigByKeyResponse RetrieveByKeyConfig(ctx, optional)

查询配置

查询配置通过配置Key和资源ID

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RetrieveByKeyConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RetrieveByKeyConfigOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **retrieveConfigByKeyReq** | [**optional.Interface of RetrieveConfigByKeyReq**](RetrieveConfigByKeyReq.md)|  | 

### Return type

[**RetrieveConfigByKeyResponse**](RetrieveConfigByKeyResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConfig

> UpdateConfigResponse UpdateConfig(ctx, optional)

更新配置

更新配置，需要全部的配置信息

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UpdateConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateConfigOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateConfigReq** | [**optional.Interface of UpdateConfigReq**](UpdateConfigReq.md)|  | 

### Return type

[**UpdateConfigResponse**](UpdateConfigResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConfigTemplate

> UpdateConfigTemplateResponse UpdateConfigTemplate(ctx, optional)

更新配置模版

更新配置模版，需要全部的配置模版信息

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UpdateConfigTemplateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateConfigTemplateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateConfigTemplateReq** | [**optional.Interface of UpdateConfigTemplateReq**](UpdateConfigTemplateReq.md)|  | 

### Return type

[**UpdateConfigTemplateResponse**](UpdateConfigTemplateResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

