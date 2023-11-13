# {{classname}}

All URIs are relative to *https://www.bungie.net/Platform*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContentGetContentById**](ContentApi.md#ContentGetContentById) | **Get** /Content/GetContentById/{id}/{locale}/ | 
[**ContentGetContentByTagAndType**](ContentApi.md#ContentGetContentByTagAndType) | **Get** /Content/GetContentByTagAndType/{tag}/{type}/{locale}/ | 
[**ContentGetContentType**](ContentApi.md#ContentGetContentType) | **Get** /Content/GetContentType/{type}/ | 
[**ContentRssNewsArticles**](ContentApi.md#ContentRssNewsArticles) | **Get** /Content/Rss/NewsArticles/{pageToken}/ | 
[**ContentSearchContentByTagAndType**](ContentApi.md#ContentSearchContentByTagAndType) | **Get** /Content/SearchContentByTagAndType/{tag}/{type}/{locale}/ | 
[**ContentSearchContentWithText**](ContentApi.md#ContentSearchContentWithText) | **Get** /Content/Search/{locale}/ | 
[**ContentSearchHelpArticles**](ContentApi.md#ContentSearchHelpArticles) | **Get** /Content/SearchHelpArticles/{searchtext}/{size}/ | 

# **ContentGetContentById**
> ContentGetContentById(ctx, id, locale, optional)


Returns a content item referenced by id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**|  | 
  **locale** | **string**|  | 
 **optional** | ***ContentApiContentGetContentByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ContentApiContentGetContentByIdOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **head** | **optional.Bool**| false | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContentGetContentByTagAndType**
> ContentGetContentByTagAndType(ctx, locale, tag, type_, optional)


Returns the newest item that matches a given tag and Content Type.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **locale** | **string**|  | 
  **tag** | **string**|  | 
  **type_** | **string**|  | 
 **optional** | ***ContentApiContentGetContentByTagAndTypeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ContentApiContentGetContentByTagAndTypeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **head** | **optional.Bool**| Not used. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContentGetContentType**
> ContentGetContentType(ctx, type_)


Gets an object describing a particular variant of content.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **type_** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContentRssNewsArticles**
> ContentRssNewsArticles(ctx, pageToken, optional)


Returns a JSON string response that is the RSS feed for news articles.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageToken** | **string**| Zero-based pagination token for paging through result sets. | 
 **optional** | ***ContentApiContentRssNewsArticlesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ContentApiContentRssNewsArticlesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **categoryfilter** | **optional.String**| Optionally filter response to only include news items in a certain category. | 
 **includebody** | **optional.Bool**| Optionally include full content body for each news item. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContentSearchContentByTagAndType**
> ContentSearchContentByTagAndType(ctx, locale, tag, type_, optional)


Searches for Content Items that match the given Tag and Content Type.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **locale** | **string**|  | 
  **tag** | **string**|  | 
  **type_** | **string**|  | 
 **optional** | ***ContentApiContentSearchContentByTagAndTypeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ContentApiContentSearchContentByTagAndTypeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **currentpage** | **optional.Int32**| Page number for the search results starting with page 1. | 
 **head** | **optional.Bool**| Not used. | 
 **itemsperpage** | **optional.Int32**| Not used. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContentSearchContentWithText**
> ContentSearchContentWithText(ctx, locale, optional)


Gets content based on querystring information passed in. Provides basic search and text search capabilities.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **locale** | **string**|  | 
 **optional** | ***ContentApiContentSearchContentWithTextOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ContentApiContentSearchContentWithTextOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ctype** | **optional.String**| Content type tag: Help, News, etc. Supply multiple ctypes separated by space. | 
 **currentpage** | **optional.Int32**| Page number for the search results, starting with page 1. | 
 **head** | **optional.Bool**| Not used. | 
 **searchtext** | **optional.String**| Word or phrase for the search. | 
 **source** | **optional.String**| For analytics, hint at the part of the app that triggered the search. Optional. | 
 **tag** | **optional.String**| Tag used on the content to be searched. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContentSearchHelpArticles**
> ContentSearchHelpArticles(ctx, searchtext, size)


Search for Help Articles.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **searchtext** | **string**|  | 
  **size** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

