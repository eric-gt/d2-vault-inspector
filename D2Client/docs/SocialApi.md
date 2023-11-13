# {{classname}}

All URIs are relative to *https://www.bungie.net/Platform*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SocialAcceptFriendRequest**](SocialApi.md#SocialAcceptFriendRequest) | **Post** /Social/Friends/Requests/Accept/{membershipId}/ | 
[**SocialDeclineFriendRequest**](SocialApi.md#SocialDeclineFriendRequest) | **Post** /Social/Friends/Requests/Decline/{membershipId}/ | 
[**SocialGetFriendList**](SocialApi.md#SocialGetFriendList) | **Get** /Social/Friends/ | 
[**SocialGetFriendRequestList**](SocialApi.md#SocialGetFriendRequestList) | **Get** /Social/Friends/Requests/ | 
[**SocialGetPlatformFriendList**](SocialApi.md#SocialGetPlatformFriendList) | **Get** /Social/PlatformFriends/{friendPlatform}/{page}/ | 
[**SocialIssueFriendRequest**](SocialApi.md#SocialIssueFriendRequest) | **Post** /Social/Friends/Add/{membershipId}/ | 
[**SocialRemoveFriend**](SocialApi.md#SocialRemoveFriend) | **Post** /Social/Friends/Remove/{membershipId}/ | 
[**SocialRemoveFriendRequest**](SocialApi.md#SocialRemoveFriendRequest) | **Post** /Social/Friends/Requests/Remove/{membershipId}/ | 

# **SocialAcceptFriendRequest**
> SocialAcceptFriendRequest(ctx, membershipId)


Accepts a friend relationship with the target user. The user must be on your incoming friend request list, though no error will occur if they are not.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **membershipId** | **string**| The membership id of the user you wish to accept. | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SocialDeclineFriendRequest**
> SocialDeclineFriendRequest(ctx, membershipId)


Declines a friend relationship with the target user. The user must be on your incoming friend request list, though no error will occur if they are not.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **membershipId** | **string**| The membership id of the user you wish to decline. | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SocialGetFriendList**
> SocialGetFriendList(ctx, )


Returns your Bungie Friend list

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SocialGetFriendRequestList**
> SocialGetFriendRequestList(ctx, )


Returns your friend request queue.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SocialGetPlatformFriendList**
> SocialGetPlatformFriendList(ctx, friendPlatform, page)


Gets the platform friend of the requested type, with additional information if they have Bungie accounts. Must have a recent login session with said platform.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **friendPlatform** | **int32**| The platform friend type. | 
  **page** | **string**| The zero based page to return. Page size is 100. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SocialIssueFriendRequest**
> SocialIssueFriendRequest(ctx, membershipId)


Requests a friend relationship with the target user. Any of the target user's linked membership ids are valid inputs.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **membershipId** | **string**| The membership id of the user you wish to add. | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SocialRemoveFriend**
> SocialRemoveFriend(ctx, membershipId)


Remove a friend relationship with the target user. The user must be on your friend list, though no error will occur if they are not.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **membershipId** | **string**| The membership id of the user you wish to remove. | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SocialRemoveFriendRequest**
> SocialRemoveFriendRequest(ctx, membershipId)


Remove a friend relationship with the target user. The user must be on your outgoing request friend list, though no error will occur if they are not.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **membershipId** | **string**| The membership id of the user you wish to remove. | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

