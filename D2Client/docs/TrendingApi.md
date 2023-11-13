# {{classname}}

All URIs are relative to *https://www.bungie.net/Platform*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TrendingGetTrendingCategories**](TrendingApi.md#TrendingGetTrendingCategories) | **Get** /Trending/Categories/ | 
[**TrendingGetTrendingCategory**](TrendingApi.md#TrendingGetTrendingCategory) | **Get** /Trending/Categories/{categoryId}/{pageNumber}/ | 
[**TrendingGetTrendingEntryDetail**](TrendingApi.md#TrendingGetTrendingEntryDetail) | **Get** /Trending/Details/{trendingEntryType}/{identifier}/ | 

# **TrendingGetTrendingCategories**
> TrendingGetTrendingCategories(ctx, )


Returns trending items for Bungie.net, collapsed into the first page of items per category. For pagination within a category, call GetTrendingCategory.

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

# **TrendingGetTrendingCategory**
> TrendingGetTrendingCategory(ctx, categoryId, pageNumber)


Returns paginated lists of trending items for a category.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **categoryId** | **string**| The ID of the category for whom you want additional results. | 
  **pageNumber** | **int32**| The page # of results to return. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TrendingGetTrendingEntryDetail**
> TrendingGetTrendingEntryDetail(ctx, identifier, trendingEntryType)


Returns the detailed results for a specific trending entry. Note that trending entries are uniquely identified by a combination of *both* the TrendingEntryType *and* the identifier: the identifier alone is not guaranteed to be globally unique.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identifier** | **string**| The identifier for the entity to be returned. | 
  **trendingEntryType** | **int32**| The type of entity to be returned. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

