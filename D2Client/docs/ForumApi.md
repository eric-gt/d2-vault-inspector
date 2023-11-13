# {{classname}}

All URIs are relative to *https://www.bungie.net/Platform*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ForumGetCoreTopicsPaged**](ForumApi.md#ForumGetCoreTopicsPaged) | **Get** /Forum/GetCoreTopicsPaged/{page}/{sort}/{quickDate}/{categoryFilter}/ | 
[**ForumGetForumTagSuggestions**](ForumApi.md#ForumGetForumTagSuggestions) | **Get** /Forum/GetForumTagSuggestions/ | 
[**ForumGetPoll**](ForumApi.md#ForumGetPoll) | **Get** /Forum/Poll/{topicId}/ | 
[**ForumGetPostAndParent**](ForumApi.md#ForumGetPostAndParent) | **Get** /Forum/GetPostAndParent/{childPostId}/ | 
[**ForumGetPostAndParentAwaitingApproval**](ForumApi.md#ForumGetPostAndParentAwaitingApproval) | **Get** /Forum/GetPostAndParentAwaitingApproval/{childPostId}/ | 
[**ForumGetPostsThreadedPaged**](ForumApi.md#ForumGetPostsThreadedPaged) | **Get** /Forum/GetPostsThreadedPaged/{parentPostId}/{page}/{pageSize}/{replySize}/{getParentPost}/{rootThreadMode}/{sortMode}/ | 
[**ForumGetPostsThreadedPagedFromChild**](ForumApi.md#ForumGetPostsThreadedPagedFromChild) | **Get** /Forum/GetPostsThreadedPagedFromChild/{childPostId}/{page}/{pageSize}/{replySize}/{rootThreadMode}/{sortMode}/ | 
[**ForumGetRecruitmentThreadSummaries**](ForumApi.md#ForumGetRecruitmentThreadSummaries) | **Post** /Forum/Recruit/Summaries/ | 
[**ForumGetTopicForContent**](ForumApi.md#ForumGetTopicForContent) | **Get** /Forum/GetTopicForContent/{contentId}/ | 
[**ForumGetTopicsPaged**](ForumApi.md#ForumGetTopicsPaged) | **Get** /Forum/GetTopicsPaged/{page}/{pageSize}/{group}/{sort}/{quickDate}/{categoryFilter}/ | 

# **ForumGetCoreTopicsPaged**
> ForumGetCoreTopicsPaged(ctx, categoryFilter, page, quickDate, sort, optional)


Gets a listing of all topics marked as part of the core group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **categoryFilter** | **int32**| The category filter. | 
  **page** | **int32**| Zero base page | 
  **quickDate** | **int32**| The date filter. | 
  **sort** | **int32**| The sort mode. | 
 **optional** | ***ForumApiForumGetCoreTopicsPagedOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ForumApiForumGetCoreTopicsPagedOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **locales** | **optional.String**| Comma seperated list of locales posts must match to return in the result list. Default &#x27;en&#x27; | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ForumGetForumTagSuggestions**
> ForumGetForumTagSuggestions(ctx, optional)


Gets tag suggestions based on partial text entry, matching them with other tags previously used in the forums.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ForumApiForumGetForumTagSuggestionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ForumApiForumGetForumTagSuggestionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **partialtag** | **optional.String**| The partial tag input to generate suggestions from. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ForumGetPoll**
> ForumGetPoll(ctx, topicId)


Gets the specified forum poll.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **topicId** | **int64**| The post id of the topic that has the poll. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ForumGetPostAndParent**
> ForumGetPostAndParent(ctx, childPostId, optional)


Returns the post specified and its immediate parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **childPostId** | **int64**|  | 
 **optional** | ***ForumApiForumGetPostAndParentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ForumApiForumGetPostAndParentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **showbanned** | **optional.String**| If this value is not null or empty, banned posts are requested to be returned | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ForumGetPostAndParentAwaitingApproval**
> ForumGetPostAndParentAwaitingApproval(ctx, childPostId, optional)


Returns the post specified and its immediate parent of posts that are awaiting approval.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **childPostId** | **int64**|  | 
 **optional** | ***ForumApiForumGetPostAndParentAwaitingApprovalOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ForumApiForumGetPostAndParentAwaitingApprovalOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **showbanned** | **optional.String**| If this value is not null or empty, banned posts are requested to be returned | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ForumGetPostsThreadedPaged**
> ForumGetPostsThreadedPaged(ctx, getParentPost, page, pageSize, parentPostId, replySize, rootThreadMode, sortMode, optional)


Returns a thread of posts at the given parent, optionally returning replies to those posts as well as the original parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **getParentPost** | **bool**|  | 
  **page** | **int32**|  | 
  **pageSize** | **int32**|  | 
  **parentPostId** | **int64**|  | 
  **replySize** | **int32**|  | 
  **rootThreadMode** | **bool**|  | 
  **sortMode** | **int32**|  | 
 **optional** | ***ForumApiForumGetPostsThreadedPagedOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ForumApiForumGetPostsThreadedPagedOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------







 **showbanned** | **optional.String**| If this value is not null or empty, banned posts are requested to be returned | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ForumGetPostsThreadedPagedFromChild**
> ForumGetPostsThreadedPagedFromChild(ctx, childPostId, page, pageSize, replySize, rootThreadMode, sortMode, optional)


Returns a thread of posts starting at the topicId of the input childPostId, optionally returning replies to those posts as well as the original parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **childPostId** | **int64**|  | 
  **page** | **int32**|  | 
  **pageSize** | **int32**|  | 
  **replySize** | **int32**|  | 
  **rootThreadMode** | **bool**|  | 
  **sortMode** | **int32**|  | 
 **optional** | ***ForumApiForumGetPostsThreadedPagedFromChildOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ForumApiForumGetPostsThreadedPagedFromChildOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **showbanned** | **optional.String**| If this value is not null or empty, banned posts are requested to be returned | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ForumGetRecruitmentThreadSummaries**
> ForumGetRecruitmentThreadSummaries(ctx, body)


Allows the caller to get a list of to 25 recruitment thread summary information objects.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]int64**](int64.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ForumGetTopicForContent**
> ForumGetTopicForContent(ctx, contentId)


Gets the post Id for the given content item's comments, if it exists.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contentId** | **int64**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ForumGetTopicsPaged**
> ForumGetTopicsPaged(ctx, categoryFilter, group, page, pageSize, quickDate, sort, optional)


Get topics from any forum.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **categoryFilter** | **int32**| A category filter | 
  **group** | **int64**| The group, if any. | 
  **page** | **int32**| Zero paged page number | 
  **pageSize** | **int32**| Unused | 
  **quickDate** | **int32**| A date filter. | 
  **sort** | **int32**| The sort mode. | 
 **optional** | ***ForumApiForumGetTopicsPagedOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ForumApiForumGetTopicsPagedOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **locales** | **optional.String**| Comma seperated list of locales posts must match to return in the result list. Default &#x27;en&#x27; | 
 **tagstring** | **optional.String**| The tags to search, if any. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

