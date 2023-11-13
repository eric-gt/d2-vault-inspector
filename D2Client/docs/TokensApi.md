# {{classname}}

All URIs are relative to *https://www.bungie.net/Platform*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TokensApplyMissingPartnerOffersWithoutClaim**](TokensApi.md#TokensApplyMissingPartnerOffersWithoutClaim) | **Post** /Tokens/Partner/ApplyMissingOffers/{partnerApplicationId}/{targetBnetMembershipId}/ | 
[**TokensClaimPartnerOffer**](TokensApi.md#TokensClaimPartnerOffer) | **Post** /Tokens/Partner/ClaimOffer/ | 
[**TokensForceDropsRepair**](TokensApi.md#TokensForceDropsRepair) | **Post** /Tokens/Partner/ForceDropsRepair/ | 
[**TokensGetBungieRewardsForPlatformUser**](TokensApi.md#TokensGetBungieRewardsForPlatformUser) | **Get** /Tokens/Rewards/GetRewardsForPlatformUser/{membershipId}/{membershipType}/ | 
[**TokensGetBungieRewardsForUser**](TokensApi.md#TokensGetBungieRewardsForUser) | **Get** /Tokens/Rewards/GetRewardsForUser/{membershipId}/ | 
[**TokensGetBungieRewardsList**](TokensApi.md#TokensGetBungieRewardsList) | **Get** /Tokens/Rewards/BungieRewards/ | 
[**TokensGetPartnerOfferSkuHistory**](TokensApi.md#TokensGetPartnerOfferSkuHistory) | **Get** /Tokens/Partner/History/{partnerApplicationId}/{targetBnetMembershipId}/ | 
[**TokensGetPartnerRewardHistory**](TokensApi.md#TokensGetPartnerRewardHistory) | **Get** /Tokens/Partner/History/{targetBnetMembershipId}/Application/{partnerApplicationId}/ | 

# **TokensApplyMissingPartnerOffersWithoutClaim**
> TokensApplyMissingPartnerOffersWithoutClaim(ctx, partnerApplicationId, targetBnetMembershipId)


Apply a partner offer to the targeted user. This endpoint does not claim a new offer, but any already claimed offers will be applied to the game if not already.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **partnerApplicationId** | **int32**| The partner application identifier. | 
  **targetBnetMembershipId** | **int64**| The bungie.net user to apply missing offers to. If not self, elevated permissions are required. | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TokensClaimPartnerOffer**
> TokensClaimPartnerOffer(ctx, body)


Claim a partner offer as the authenticated user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TokensPartnerOfferClaimRequest**](TokensPartnerOfferClaimRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TokensForceDropsRepair**
> TokensForceDropsRepair(ctx, )


Twitch Drops self-repair function - scans twitch for drops not marked as fulfilled and resyncs them.

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

# **TokensGetBungieRewardsForPlatformUser**
> InlineResponse2002 TokensGetBungieRewardsForPlatformUser(ctx, membershipId, membershipType)


Returns the bungie rewards for the targeted user when a platform membership Id and Type are used.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **membershipId** | **int64**| users platform membershipId for requested user rewards. If not self, elevated permissions are required. | 
  **membershipType** | **int32**| The target Destiny 2 membership type. | 

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TokensGetBungieRewardsForUser**
> InlineResponse2002 TokensGetBungieRewardsForUser(ctx, membershipId)


Returns the bungie rewards for the targeted user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **membershipId** | **int64**| bungie.net user membershipId for requested user rewards. If not self, elevated permissions are required. | 

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TokensGetBungieRewardsList**
> InlineResponse2002 TokensGetBungieRewardsList(ctx, )


Returns a list of the current bungie rewards

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TokensGetPartnerOfferSkuHistory**
> TokensGetPartnerOfferSkuHistory(ctx, partnerApplicationId, targetBnetMembershipId)


Returns the partner sku and offer history of the targeted user. Elevated permissions are required to see users that are not yourself.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **partnerApplicationId** | **int32**| The partner application identifier. | 
  **targetBnetMembershipId** | **int64**| The bungie.net user to apply missing offers to. If not self, elevated permissions are required. | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TokensGetPartnerRewardHistory**
> TokensGetPartnerRewardHistory(ctx, partnerApplicationId, targetBnetMembershipId)


Returns the partner rewards history of the targeted user, both partner offers and Twitch drops.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **partnerApplicationId** | **int32**| The partner application identifier. | 
  **targetBnetMembershipId** | **int64**| The bungie.net user to return reward history for. | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

