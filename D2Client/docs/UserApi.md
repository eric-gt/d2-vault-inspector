# {{classname}}

All URIs are relative to *https://www.bungie.net/Platform*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UserGetAvailableThemes**](UserApi.md#UserGetAvailableThemes) | **Get** /User/GetAvailableThemes/ | 
[**UserGetBungieNetUserById**](UserApi.md#UserGetBungieNetUserById) | **Get** /User/GetBungieNetUserById/{id}/ | 
[**UserGetCredentialTypesForTargetAccount**](UserApi.md#UserGetCredentialTypesForTargetAccount) | **Get** /User/GetCredentialTypesForTargetAccount/{membershipId}/ | 
[**UserGetMembershipDataById**](UserApi.md#UserGetMembershipDataById) | **Get** /User/GetMembershipsById/{membershipId}/{membershipType}/ | 
[**UserGetMembershipDataForCurrentUser**](UserApi.md#UserGetMembershipDataForCurrentUser) | **Get** /User/GetMembershipsForCurrentUser/ | 
[**UserGetMembershipFromHardLinkedCredential**](UserApi.md#UserGetMembershipFromHardLinkedCredential) | **Get** /User/GetMembershipFromHardLinkedCredential/{crType}/{credential}/ | 
[**UserGetSanitizedPlatformDisplayNames**](UserApi.md#UserGetSanitizedPlatformDisplayNames) | **Get** /User/GetSanitizedPlatformDisplayNames/{membershipId}/ | 
[**UserSearchByGlobalNamePost**](UserApi.md#UserSearchByGlobalNamePost) | **Post** /User/Search/GlobalName/{page}/ | 
[**UserSearchByGlobalNamePrefix**](UserApi.md#UserSearchByGlobalNamePrefix) | **Get** /User/Search/Prefix/{displayNamePrefix}/{page}/ | 

# **UserGetAvailableThemes**
> UserGetAvailableThemes(ctx, )


Returns a list of all available user themes.

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

# **UserGetBungieNetUserById**
> UserGetBungieNetUserById(ctx, id)


Loads a bungienet user by membership id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| The requested Bungie.net membership id. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserGetCredentialTypesForTargetAccount**
> UserGetCredentialTypesForTargetAccount(ctx, membershipId)


Returns a list of credential types attached to the requested account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **membershipId** | **int64**| The user&#x27;s membership id | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserGetMembershipDataById**
> UserGetMembershipDataById(ctx, membershipId, membershipType)


Returns a list of accounts associated with the supplied membership ID and membership type. This will include all linked accounts (even when hidden) if supplied credentials permit it.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **membershipId** | **int64**| The membership ID of the target user. | 
  **membershipType** | **int32**| Type of the supplied membership ID. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserGetMembershipDataForCurrentUser**
> UserGetMembershipDataForCurrentUser(ctx, )


Returns a list of accounts associated with signed in user. This is useful for OAuth implementations that do not give you access to the token response.

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

# **UserGetMembershipFromHardLinkedCredential**
> UserGetMembershipFromHardLinkedCredential(ctx, credential, crType)


Gets any hard linked membership given a credential. Only works for credentials that are public (just SteamID64 right now). Cross Save aware.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **credential** | **string**| The credential to look up. Must be a valid SteamID64. | 
  **crType** | **int32**| The credential type. &#x27;SteamId&#x27; is the only valid value at present. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserGetSanitizedPlatformDisplayNames**
> UserGetSanitizedPlatformDisplayNames(ctx, membershipId)


Gets a list of all display names linked to this membership id but sanitized (profanity filtered). Obeys all visibility rules of calling user and is heavily cached.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **membershipId** | **int64**| The requested membership id to load. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserSearchByGlobalNamePost**
> UserSearchByGlobalNamePost(ctx, body, page)


Given the prefix of a global display name, returns all users who share that name.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UserUserSearchPrefixRequest**](UserUserSearchPrefixRequest.md)|  | 
  **page** | **int32**| The zero-based page of results you desire. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserSearchByGlobalNamePrefix**
> UserSearchByGlobalNamePrefix(ctx, displayNamePrefix, page)


[OBSOLETE] Do not use this to search users, use SearchByGlobalNamePost instead.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **displayNamePrefix** | **string**| The display name prefix you&#x27;re looking for. | 
  **page** | **int32**| The zero-based page of results you desire. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

