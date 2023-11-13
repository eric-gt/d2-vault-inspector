# {{classname}}

All URIs are relative to *https://www.bungie.net/Platform*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FireteamGetActivePrivateClanFireteamCount**](FireteamApi.md#FireteamGetActivePrivateClanFireteamCount) | **Get** /Fireteam/Clan/{groupId}/ActiveCount/ | 
[**FireteamGetAvailableClanFireteams**](FireteamApi.md#FireteamGetAvailableClanFireteams) | **Get** /Fireteam/Clan/{groupId}/Available/{platform}/{activityType}/{dateRange}/{slotFilter}/{publicOnly}/{page}/ | 
[**FireteamGetClanFireteam**](FireteamApi.md#FireteamGetClanFireteam) | **Get** /Fireteam/Clan/{groupId}/Summary/{fireteamId}/ | 
[**FireteamGetMyClanFireteams**](FireteamApi.md#FireteamGetMyClanFireteams) | **Get** /Fireteam/Clan/{groupId}/My/{platform}/{includeClosed}/{page}/ | 
[**FireteamSearchPublicAvailableClanFireteams**](FireteamApi.md#FireteamSearchPublicAvailableClanFireteams) | **Get** /Fireteam/Search/Available/{platform}/{activityType}/{dateRange}/{slotFilter}/{page}/ | 

# **FireteamGetActivePrivateClanFireteamCount**
> FireteamGetActivePrivateClanFireteamCount(ctx, groupId)


Gets a count of all active non-public fireteams for the specified clan. Maximum value returned is 25.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **int64**| The group id of the clan. | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FireteamGetAvailableClanFireteams**
> FireteamGetAvailableClanFireteams(ctx, activityType, dateRange, groupId, page, platform, publicOnly, slotFilter, optional)


Gets a listing of all of this clan's fireteams that are have available slots. Caller is not checked for join criteria so caching is maximized.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **activityType** | **int32**| The activity type to filter by. | 
  **dateRange** | **int32**| The date range to grab available fireteams. | 
  **groupId** | **int64**| The group id of the clan. | 
  **page** | **int32**| Zero based page | 
  **platform** | **int32**| The platform filter. | 
  **publicOnly** | **int32**| Determines public/private filtering. | 
  **slotFilter** | **int32**| Filters based on available slots | 
 **optional** | ***FireteamApiFireteamGetAvailableClanFireteamsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FireteamApiFireteamGetAvailableClanFireteamsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------







 **excludeImmediate** | **optional.Bool**| If you wish the result to exclude immediate fireteams, set this to true. Immediate-only can be forced using the dateRange enum. | 
 **langFilter** | **optional.String**| An optional language filter. | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FireteamGetClanFireteam**
> FireteamGetClanFireteam(ctx, fireteamId, groupId)


Gets a specific fireteam.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fireteamId** | **int64**| The unique id of the fireteam. | 
  **groupId** | **int64**| The group id of the clan. | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FireteamGetMyClanFireteams**
> FireteamGetMyClanFireteams(ctx, groupId, includeClosed, page, platform, optional)


Gets a listing of all fireteams that caller is an applicant, a member, or an alternate of.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **int64**| The group id of the clan. (This parameter is ignored unless the optional query parameter groupFilter is true). | 
  **includeClosed** | **bool**| If true, return fireteams that have been closed. | 
  **page** | **int32**| Deprecated parameter, ignored. | 
  **platform** | **int32**| The platform filter. | 
 **optional** | ***FireteamApiFireteamGetMyClanFireteamsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FireteamApiFireteamGetMyClanFireteamsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **groupFilter** | **optional.Bool**| If true, filter by clan. Otherwise, ignore the clan and show all of the user&#x27;s fireteams. | 
 **langFilter** | **optional.String**| An optional language filter. | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FireteamSearchPublicAvailableClanFireteams**
> FireteamSearchPublicAvailableClanFireteams(ctx, activityType, dateRange, page, platform, slotFilter, optional)


Gets a listing of all public fireteams starting now with open slots. Caller is not checked for join criteria so caching is maximized.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **activityType** | **int32**| The activity type to filter by. | 
  **dateRange** | **int32**| The date range to grab available fireteams. | 
  **page** | **int32**| Zero based page | 
  **platform** | **int32**| The platform filter. | 
  **slotFilter** | **int32**| Filters based on available slots | 
 **optional** | ***FireteamApiFireteamSearchPublicAvailableClanFireteamsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FireteamApiFireteamSearchPublicAvailableClanFireteamsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **excludeImmediate** | **optional.Bool**| If you wish the result to exclude immediate fireteams, set this to true. Immediate-only can be forced using the dateRange enum. | 
 **langFilter** | **optional.String**| An optional language filter. | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

