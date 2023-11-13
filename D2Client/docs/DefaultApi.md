# {{classname}}

All URIs are relative to *https://www.bungie.net/Platform*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAvailableLocales**](DefaultApi.md#GetAvailableLocales) | **Get** /GetAvailableLocales/ | 
[**GetCommonSettings**](DefaultApi.md#GetCommonSettings) | **Get** /Settings/ | 
[**GetGlobalAlerts**](DefaultApi.md#GetGlobalAlerts) | **Get** /GlobalAlerts/ | 
[**GetUserSystemOverrides**](DefaultApi.md#GetUserSystemOverrides) | **Get** /UserSystemOverrides/ | 

# **GetAvailableLocales**
> GetAvailableLocales(ctx, )


List of available localization cultures

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCommonSettings**
> GetCommonSettings(ctx, )


Get the common settings used by the Bungie.Net environment.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGlobalAlerts**
> GetGlobalAlerts(ctx, optional)


Gets any active global alert for display in the forum banners, help pages, etc. Usually used for DOC alerts.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DefaultApiGetGlobalAlertsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiGetGlobalAlertsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includestreaming** | **optional.Bool**| Determines whether Streaming Alerts are included in results | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserSystemOverrides**
> GetUserSystemOverrides(ctx, )


Get the user-specific system overrides that should be respected alongside common systems.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

