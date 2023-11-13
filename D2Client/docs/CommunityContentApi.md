# {{classname}}

All URIs are relative to *https://www.bungie.net/Platform*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CommunityContentGetCommunityContent**](CommunityContentApi.md#CommunityContentGetCommunityContent) | **Get** /CommunityContent/Get/{sort}/{mediaFilter}/{page}/ | 

# **CommunityContentGetCommunityContent**
> CommunityContentGetCommunityContent(ctx, mediaFilter, page, sort)


Returns community content.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mediaFilter** | **int32**| The type of media to get | 
  **page** | **int32**| Zero based page | 
  **sort** | **int32**| The sort mode. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

