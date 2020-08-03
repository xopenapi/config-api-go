# \ConfigApi

All URIs are relative to *https://config.api.lucfish.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](ConfigApi.md#Create) | **Post** /config/create | 创建配置
[**Update**](ConfigApi.md#Update) | **Post** /config/update | 更新配置



## Create

> ApiResponseConfig Create(ctx, body)

创建配置

创建配置

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**CreateConfigReq**](CreateConfigReq.md)|  | 

### Return type

[**ApiResponseConfig**](APIResponse_Config.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> ApiResponseConfig Update(ctx, body)

更新配置

更新配置

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**UpdateConfigReq**](UpdateConfigReq.md)|  | 

### Return type

[**ApiResponseConfig**](APIResponse_Config.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

