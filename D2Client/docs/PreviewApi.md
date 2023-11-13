# {{classname}}

All URIs are relative to *https://www.bungie.net/Platform*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Destiny2GetClanAggregateStats**](PreviewApi.md#Destiny2GetClanAggregateStats) | **Get** /Destiny2/Stats/AggregateClanStats/{groupId}/ | 
[**Destiny2GetClanLeaderboards**](PreviewApi.md#Destiny2GetClanLeaderboards) | **Get** /Destiny2/Stats/Leaderboards/Clans/{groupId}/ | 
[**Destiny2GetLeaderboards**](PreviewApi.md#Destiny2GetLeaderboards) | **Get** /Destiny2/{membershipType}/Account/{destinyMembershipId}/Stats/Leaderboards/ | 
[**Destiny2GetLeaderboardsForCharacter**](PreviewApi.md#Destiny2GetLeaderboardsForCharacter) | **Get** /Destiny2/Stats/Leaderboards/{membershipType}/{destinyMembershipId}/{characterId}/ | 
[**Destiny2GetPublicVendors**](PreviewApi.md#Destiny2GetPublicVendors) | **Get** /Destiny2/Vendors/ | 
[**Destiny2InsertSocketPlug**](PreviewApi.md#Destiny2InsertSocketPlug) | **Post** /Destiny2/Actions/Items/InsertSocketPlug/ | 
[**Destiny2InsertSocketPlugFree**](PreviewApi.md#Destiny2InsertSocketPlugFree) | **Post** /Destiny2/Actions/Items/InsertSocketPlugFree/ | 

# **Destiny2GetClanAggregateStats**
> Destiny2GetClanAggregateStats(ctx, groupId, optional)


Gets aggregated stats for a clan using the same categories as the clan leaderboards. PREVIEW: This endpoint is still in beta, and may experience rough edges. The schema is in final form, but there may be bugs that prevent desirable operation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **int64**| Group ID of the clan whose leaderboards you wish to fetch. | 
 **optional** | ***PreviewApiDestiny2GetClanAggregateStatsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PreviewApiDestiny2GetClanAggregateStatsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **modes** | **optional.String**| List of game modes for which to get leaderboards. See the documentation for DestinyActivityModeType for valid values, and pass in string representation, comma delimited. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Destiny2GetClanLeaderboards**
> Destiny2GetClanLeaderboards(ctx, groupId, optional)


Gets leaderboards with the signed in user's friends and the supplied destinyMembershipId as the focus. PREVIEW: This endpoint is still in beta, and may experience rough edges. The schema is in final form, but there may be bugs that prevent desirable operation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **int64**| Group ID of the clan whose leaderboards you wish to fetch. | 
 **optional** | ***PreviewApiDestiny2GetClanLeaderboardsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PreviewApiDestiny2GetClanLeaderboardsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **maxtop** | **optional.Int32**| Maximum number of top players to return. Use a large number to get entire leaderboard. | 
 **modes** | **optional.String**| List of game modes for which to get leaderboards. See the documentation for DestinyActivityModeType for valid values, and pass in string representation, comma delimited. | 
 **statid** | **optional.String**| ID of stat to return rather than returning all Leaderboard stats. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Destiny2GetLeaderboards**
> Destiny2GetLeaderboards(ctx, destinyMembershipId, membershipType, optional)


Gets leaderboards with the signed in user's friends and the supplied destinyMembershipId as the focus. PREVIEW: This endpoint has not yet been implemented. It is being returned for a preview of future functionality, and for public comment/suggestion/preparation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **destinyMembershipId** | **int64**| The Destiny membershipId of the user to retrieve. | 
  **membershipType** | **int32**| A valid non-BungieNet membership type. | 
 **optional** | ***PreviewApiDestiny2GetLeaderboardsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PreviewApiDestiny2GetLeaderboardsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **maxtop** | **optional.Int32**| Maximum number of top players to return. Use a large number to get entire leaderboard. | 
 **modes** | **optional.String**| List of game modes for which to get leaderboards. See the documentation for DestinyActivityModeType for valid values, and pass in string representation, comma delimited. | 
 **statid** | **optional.String**| ID of stat to return rather than returning all Leaderboard stats. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Destiny2GetLeaderboardsForCharacter**
> Destiny2GetLeaderboardsForCharacter(ctx, characterId, destinyMembershipId, membershipType, optional)


Gets leaderboards with the signed in user's friends and the supplied destinyMembershipId as the focus. PREVIEW: This endpoint is still in beta, and may experience rough edges. The schema is in final form, but there may be bugs that prevent desirable operation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **characterId** | **int64**| The specific character to build the leaderboard around for the provided Destiny Membership. | 
  **destinyMembershipId** | **int64**| The Destiny membershipId of the user to retrieve. | 
  **membershipType** | **int32**| A valid non-BungieNet membership type. | 
 **optional** | ***PreviewApiDestiny2GetLeaderboardsForCharacterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PreviewApiDestiny2GetLeaderboardsForCharacterOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **maxtop** | **optional.Int32**| Maximum number of top players to return. Use a large number to get entire leaderboard. | 
 **modes** | **optional.String**| List of game modes for which to get leaderboards. See the documentation for DestinyActivityModeType for valid values, and pass in string representation, comma delimited. | 
 **statid** | **optional.String**| ID of stat to return rather than returning all Leaderboard stats. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Destiny2GetPublicVendors**
> Destiny2GetPublicVendors(ctx, optional)


Get items available from vendors where the vendors have items for sale that are common for everyone. If any portion of the Vendor's available inventory is character or account specific, we will be unable to return their data from this endpoint due to the way that available inventory is computed. As I am often guilty of saying: 'It's a long story...'

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PreviewApiDestiny2GetPublicVendorsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PreviewApiDestiny2GetPublicVendorsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **components** | [**optional.Interface of []int32**](int32.md)| A comma separated list of components to return (as strings or numeric values). See the DestinyComponentType enum for valid components to request. You must request at least one component to receive results. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Destiny2InsertSocketPlug**
> Destiny2InsertSocketPlug(ctx, body)


Insert a plug into a socketed item. I know how it sounds, but I assure you it's much more G-rated than you might be guessing. We haven't decided yet whether this will be able to insert plugs that have side effects, but if we do it will require special scope permission for an application attempting to do so. You must have a valid Destiny Account, and either be in a social space, in orbit, or offline. Request must include proof of permission for 'InsertPlugs' from the account owner.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DestinyRequestsActionsDestinyInsertPlugsActionRequest**](DestinyRequestsActionsDestinyInsertPlugsActionRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Destiny2InsertSocketPlugFree**
> Destiny2InsertSocketPlugFree(ctx, body)


Insert a 'free' plug into an item's socket. This does not require 'Advanced Write Action' authorization and is available to 3rd-party apps, but will only work on 'free and reversible' socket actions (Perks, Armor Mods, Shaders, Ornaments, etc.). You must have a valid Destiny Account, and the character must either be in a social space, in orbit, or offline.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DestinyRequestsActionsDestinyInsertPlugsFreeActionRequest**](DestinyRequestsActionsDestinyInsertPlugsFreeActionRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

