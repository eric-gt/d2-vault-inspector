# {{classname}}

All URIs are relative to *https://www.bungie.net/Platform*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GroupV2AbdicateFoundership**](GroupV2Api.md#GroupV2AbdicateFoundership) | **Post** /GroupV2/{groupId}/Admin/AbdicateFoundership/{membershipType}/{founderIdNew}/ | 
[**GroupV2AddOptionalConversation**](GroupV2Api.md#GroupV2AddOptionalConversation) | **Post** /GroupV2/{groupId}/OptionalConversations/Add/ | 
[**GroupV2ApproveAllPending**](GroupV2Api.md#GroupV2ApproveAllPending) | **Post** /GroupV2/{groupId}/Members/ApproveAll/ | 
[**GroupV2ApprovePending**](GroupV2Api.md#GroupV2ApprovePending) | **Post** /GroupV2/{groupId}/Members/Approve/{membershipType}/{membershipId}/ | 
[**GroupV2ApprovePendingForList**](GroupV2Api.md#GroupV2ApprovePendingForList) | **Post** /GroupV2/{groupId}/Members/ApproveList/ | 
[**GroupV2BanMember**](GroupV2Api.md#GroupV2BanMember) | **Post** /GroupV2/{groupId}/Members/{membershipType}/{membershipId}/Ban/ | 
[**GroupV2DenyAllPending**](GroupV2Api.md#GroupV2DenyAllPending) | **Post** /GroupV2/{groupId}/Members/DenyAll/ | 
[**GroupV2DenyPendingForList**](GroupV2Api.md#GroupV2DenyPendingForList) | **Post** /GroupV2/{groupId}/Members/DenyList/ | 
[**GroupV2EditClanBanner**](GroupV2Api.md#GroupV2EditClanBanner) | **Post** /GroupV2/{groupId}/EditClanBanner/ | 
[**GroupV2EditFounderOptions**](GroupV2Api.md#GroupV2EditFounderOptions) | **Post** /GroupV2/{groupId}/EditFounderOptions/ | 
[**GroupV2EditGroup**](GroupV2Api.md#GroupV2EditGroup) | **Post** /GroupV2/{groupId}/Edit/ | 
[**GroupV2EditGroupMembership**](GroupV2Api.md#GroupV2EditGroupMembership) | **Post** /GroupV2/{groupId}/Members/{membershipType}/{membershipId}/SetMembershipType/{memberType}/ | 
[**GroupV2EditOptionalConversation**](GroupV2Api.md#GroupV2EditOptionalConversation) | **Post** /GroupV2/{groupId}/OptionalConversations/Edit/{conversationId}/ | 
[**GroupV2GetAdminsAndFounderOfGroup**](GroupV2Api.md#GroupV2GetAdminsAndFounderOfGroup) | **Get** /GroupV2/{groupId}/AdminsAndFounder/{currentpage} | 
[**GroupV2GetAvailableAvatars**](GroupV2Api.md#GroupV2GetAvailableAvatars) | **Get** /GroupV2/GetAvailableAvatars/ | 
[**GroupV2GetAvailableThemes**](GroupV2Api.md#GroupV2GetAvailableThemes) | **Get** /GroupV2/GetAvailableThemes/ | 
[**GroupV2GetBannedMembersOfGroup**](GroupV2Api.md#GroupV2GetBannedMembersOfGroup) | **Get** /GroupV2/{groupId}/Banned/{currentpage} | 
[**GroupV2GetGroup**](GroupV2Api.md#GroupV2GetGroup) | **Get** /GroupV2/{groupId}/ | 
[**GroupV2GetGroupByName**](GroupV2Api.md#GroupV2GetGroupByName) | **Get** /GroupV2/Name/{groupName}/{groupType}/ | 
[**GroupV2GetGroupByNameV2**](GroupV2Api.md#GroupV2GetGroupByNameV2) | **Post** /GroupV2/NameV2/ | 
[**GroupV2GetGroupOptionalConversations**](GroupV2Api.md#GroupV2GetGroupOptionalConversations) | **Get** /GroupV2/{groupId}/OptionalConversations/ | 
[**GroupV2GetGroupsForMember**](GroupV2Api.md#GroupV2GetGroupsForMember) | **Get** /GroupV2/User/{membershipType}/{membershipId}/{filter}/{groupType}/ | 
[**GroupV2GetInvitedIndividuals**](GroupV2Api.md#GroupV2GetInvitedIndividuals) | **Get** /GroupV2/{groupId}/Members/InvitedIndividuals/{currentpage} | 
[**GroupV2GetMembersOfGroupcurrentPage**](GroupV2Api.md#GroupV2GetMembersOfGroupcurrentPage) | **Get** /GroupV2/{groupId}/Members/{currentpage} | 
[**GroupV2GetPendingMemberships**](GroupV2Api.md#GroupV2GetPendingMemberships) | **Get** /GroupV2/{groupId}/Members/Pending/{currentpage} | 
[**GroupV2GetPotentialGroupsForMember**](GroupV2Api.md#GroupV2GetPotentialGroupsForMember) | **Get** /GroupV2/User/Potential/{membershipType}/{membershipId}/{filter}/{groupType}/ | 
[**GroupV2GetRecommendedGroups**](GroupV2Api.md#GroupV2GetRecommendedGroups) | **Post** /GroupV2/Recommended/{groupType}/{createDateRange}/ | 
[**GroupV2GetUserClanInviteSetting**](GroupV2Api.md#GroupV2GetUserClanInviteSetting) | **Get** /GroupV2/GetUserClanInviteSetting/{mType}/ | 
[**GroupV2GroupSearch**](GroupV2Api.md#GroupV2GroupSearch) | **Post** /GroupV2/Search/ | 
[**GroupV2IndividualGroupInvite**](GroupV2Api.md#GroupV2IndividualGroupInvite) | **Post** /GroupV2/{groupId}/Members/IndividualInvite/{membershipType}/{membershipId}/ | 
[**GroupV2IndividualGroupInviteCancel**](GroupV2Api.md#GroupV2IndividualGroupInviteCancel) | **Post** /GroupV2/{groupId}/Members/IndividualInviteCancel/{membershipType}/{membershipId}/ | 
[**GroupV2KickMember**](GroupV2Api.md#GroupV2KickMember) | **Post** /GroupV2/{groupId}/Members/{membershipType}/{membershipId}/Kick/ | 
[**GroupV2RecoverGroupForFounder**](GroupV2Api.md#GroupV2RecoverGroupForFounder) | **Get** /GroupV2/Recover/{membershipType}/{membershipId}/{groupType}/ | 
[**GroupV2UnbanMember**](GroupV2Api.md#GroupV2UnbanMember) | **Post** /GroupV2/{groupId}/Members/{membershipType}/{membershipId}/Unban/ | 

# **GroupV2AbdicateFoundership**
> GroupV2AbdicateFoundership(ctx, founderIdNew, groupId, membershipType)


An administrative method to allow the founder of a group or clan to give up their position to another admin permanently.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **founderIdNew** | **int64**| The new founder for this group. Must already be a group admin. | 
  **groupId** | **int64**| The target group id. | 
  **membershipType** | **int32**| Membership type of the provided founderIdNew. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupV2AddOptionalConversation**
> GroupV2AddOptionalConversation(ctx, body, groupId)


Add a new optional conversation/chat channel. Requires admin permissions to the group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GroupsV2GroupOptionalConversationAddRequest**](GroupsV2GroupOptionalConversationAddRequest.md)|  | 
  **groupId** | **int64**| Group ID of the group to edit. | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupV2ApproveAllPending**
> InlineResponse2001 GroupV2ApproveAllPending(ctx, body, groupId)


Approve all of the pending users for the given group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GroupsV2GroupApplicationRequest**](GroupsV2GroupApplicationRequest.md)|  | 
  **groupId** | **int64**| ID of the group. | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupV2ApprovePending**
> GroupV2ApprovePending(ctx, body, groupId, membershipId, membershipType)


Approve the given membershipId to join the group/clan as long as they have applied.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GroupsV2GroupApplicationRequest**](GroupsV2GroupApplicationRequest.md)|  | 
  **groupId** | **int64**| ID of the group. | 
  **membershipId** | **int64**| The membership id being approved. | 
  **membershipType** | **int32**| Membership type of the supplied membership ID. | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupV2ApprovePendingForList**
> InlineResponse2001 GroupV2ApprovePendingForList(ctx, body, groupId)


Approve all of the pending users for the given group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GroupsV2GroupApplicationListRequest**](GroupsV2GroupApplicationListRequest.md)|  | 
  **groupId** | **int64**| ID of the group. | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupV2BanMember**
> GroupV2BanMember(ctx, body, groupId, membershipId, membershipType)


Bans the requested member from the requested group for the specified period of time.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GroupsV2GroupBanRequest**](GroupsV2GroupBanRequest.md)|  | 
  **groupId** | **int64**| Group ID that has the member to ban. | 
  **membershipId** | **int64**| Membership ID of the member to ban from the group. | 
  **membershipType** | **int32**| Membership type of the provided membership ID. | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupV2DenyAllPending**
> InlineResponse2001 GroupV2DenyAllPending(ctx, body, groupId)


Deny all of the pending users for the given group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GroupsV2GroupApplicationRequest**](GroupsV2GroupApplicationRequest.md)|  | 
  **groupId** | **int64**| ID of the group. | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupV2DenyPendingForList**
> InlineResponse2001 GroupV2DenyPendingForList(ctx, body, groupId)


Deny all of the pending users for the given group that match the passed-in .

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GroupsV2GroupApplicationListRequest**](GroupsV2GroupApplicationListRequest.md)|  | 
  **groupId** | **int64**| ID of the group. | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupV2EditClanBanner**
> GroupV2EditClanBanner(ctx, body, groupId)


Edit an existing group's clan banner. You must have suitable permissions in the group to perform this operation. All fields are required.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GroupsV2ClanBanner**](GroupsV2ClanBanner.md)|  | 
  **groupId** | **int64**| Group ID of the group to edit. | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupV2EditFounderOptions**
> GroupV2EditFounderOptions(ctx, body, groupId)


Edit group options only available to a founder. You must have suitable permissions in the group to perform this operation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GroupsV2GroupOptionsEditAction**](GroupsV2GroupOptionsEditAction.md)|  | 
  **groupId** | **int64**| Group ID of the group to edit. | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupV2EditGroup**
> GroupV2EditGroup(ctx, body, groupId)


Edit an existing group. You must have suitable permissions in the group to perform this operation. This latest revision will only edit the fields you pass in - pass null for properties you want to leave unaltered.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GroupsV2GroupEditAction**](GroupsV2GroupEditAction.md)|  | 
  **groupId** | **int64**| Group ID of the group to edit. | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupV2EditGroupMembership**
> GroupV2EditGroupMembership(ctx, groupId, membershipId, membershipType, memberType)


Edit the membership type of a given member. You must have suitable permissions in the group to perform this operation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **int64**| ID of the group to which the member belongs. | 
  **membershipId** | **int64**| Membership ID to modify. | 
  **membershipType** | **int32**| Membership type of the provide membership ID. | 
  **memberType** | **int32**| New membertype for the specified member. | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupV2EditOptionalConversation**
> GroupV2EditOptionalConversation(ctx, body, conversationId, groupId)


Edit the settings of an optional conversation/chat channel. Requires admin permissions to the group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GroupsV2GroupOptionalConversationEditRequest**](GroupsV2GroupOptionalConversationEditRequest.md)|  | 
  **conversationId** | **int64**| Conversation Id of the channel being edited. | 
  **groupId** | **int64**| Group ID of the group to edit. | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupV2GetAdminsAndFounderOfGroup**
> GroupV2GetAdminsAndFounderOfGroup(ctx, currentpage, groupId)


Get the list of members in a given group who are of admin level or higher.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **currentpage** | **int32**| Page number (starting with 1). Each page has a fixed size of 50 items per page. | 
  **groupId** | **int64**| The ID of the group. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupV2GetAvailableAvatars**
> InlineResponse200 GroupV2GetAvailableAvatars(ctx, )


Returns a list of all available group avatars for the signed-in user.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupV2GetAvailableThemes**
> GroupV2GetAvailableThemes(ctx, )


Returns a list of all available group themes.

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

# **GroupV2GetBannedMembersOfGroup**
> GroupV2GetBannedMembersOfGroup(ctx, currentpage, groupId)


Get the list of banned members in a given group. Only accessible to group Admins and above. Not applicable to all groups. Check group features.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **currentpage** | **int32**| Page number (starting with 1). Each page has a fixed size of 50 entries. | 
  **groupId** | **int64**| Group ID whose banned members you are fetching | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupV2GetGroup**
> GroupV2GetGroup(ctx, groupId)


Get information about a specific group of the given ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **int64**| Requested group&#x27;s id. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupV2GetGroupByName**
> GroupV2GetGroupByName(ctx, groupName, groupType)


Get information about a specific group with the given name and type.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupName** | **string**| Exact name of the group to find. | 
  **groupType** | **int32**| Type of group to find. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupV2GetGroupByNameV2**
> GroupV2GetGroupByNameV2(ctx, body)


Get information about a specific group with the given name and type. The POST version.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GroupsV2GroupNameSearchRequest**](GroupsV2GroupNameSearchRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupV2GetGroupOptionalConversations**
> GroupV2GetGroupOptionalConversations(ctx, groupId)


Gets a list of available optional conversation channels and their settings.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **int64**| Requested group&#x27;s id. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupV2GetGroupsForMember**
> GroupV2GetGroupsForMember(ctx, filter, groupType, membershipId, membershipType)


Get information about the groups that a given member has joined.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **filter** | **int32**| Filter apply to list of joined groups. | 
  **groupType** | **int32**| Type of group the supplied member founded. | 
  **membershipId** | **int64**| Membership ID to for which to find founded groups. | 
  **membershipType** | **int32**| Membership type of the supplied membership ID. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupV2GetInvitedIndividuals**
> GroupV2GetInvitedIndividuals(ctx, currentpage, groupId)


Get the list of users who have been invited into the group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **currentpage** | **int32**| Page number (starting with 1). Each page has a fixed size of 50 items per page. | 
  **groupId** | **int64**| ID of the group. | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupV2GetMembersOfGroupcurrentPage**
> GroupV2GetMembersOfGroupcurrentPage(ctx, currentpage, groupId, optional)


Get the list of members in a given group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **currentpage** | **int32**| Page number (starting with 1). Each page has a fixed size of 50 items per page. | 
  **groupId** | **int64**| The ID of the group. | 
 **optional** | ***GroupV2ApiGroupV2GetMembersOfGroupcurrentPageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupV2ApiGroupV2GetMembersOfGroupcurrentPageOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **memberType** | **optional.Int32**| Filter out other member types. Use None for all members. | 
 **nameSearch** | **optional.String**| The name fragment upon which a search should be executed for members with matching display or unique names. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupV2GetPendingMemberships**
> GroupV2GetPendingMemberships(ctx, currentpage, groupId)


Get the list of users who are awaiting a decision on their application to join a given group. Modified to include application info.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **currentpage** | **int32**| Page number (starting with 1). Each page has a fixed size of 50 items per page. | 
  **groupId** | **int64**| ID of the group. | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupV2GetPotentialGroupsForMember**
> GroupV2GetPotentialGroupsForMember(ctx, filter, groupType, membershipId, membershipType)


Get information about the groups that a given member has applied to or been invited to.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **filter** | **int32**| Filter apply to list of potential joined groups. | 
  **groupType** | **int32**| Type of group the supplied member applied. | 
  **membershipId** | **int64**| Membership ID to for which to find applied groups. | 
  **membershipType** | **int32**| Membership type of the supplied membership ID. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupV2GetRecommendedGroups**
> GroupV2GetRecommendedGroups(ctx, createDateRange, groupType)


Gets groups recommended for you based on the groups to whom those you follow belong.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createDateRange** | **int32**| Requested range in which to pull recommended groups | 
  **groupType** | **int32**| Type of groups requested | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupV2GetUserClanInviteSetting**
> GroupV2GetUserClanInviteSetting(ctx, mType)


Gets the state of the user's clan invite preferences for a particular membership type - true if they wish to be invited to clans, false otherwise.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mType** | **int32**| The Destiny membership type of the account we wish to access settings. | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupV2GroupSearch**
> GroupV2GroupSearch(ctx, body)


Search for Groups.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GroupsV2GroupQuery**](GroupsV2GroupQuery.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupV2IndividualGroupInvite**
> GroupV2IndividualGroupInvite(ctx, body, groupId, membershipId, membershipType)


Invite a user to join this group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GroupsV2GroupApplicationRequest**](GroupsV2GroupApplicationRequest.md)|  | 
  **groupId** | **int64**| ID of the group you would like to join. | 
  **membershipId** | **int64**| Membership id of the account being invited. | 
  **membershipType** | **int32**| MembershipType of the account being invited. | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupV2IndividualGroupInviteCancel**
> GroupV2IndividualGroupInviteCancel(ctx, groupId, membershipId, membershipType)


Cancels a pending invitation to join a group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **int64**| ID of the group you would like to join. | 
  **membershipId** | **int64**| Membership id of the account being cancelled. | 
  **membershipType** | **int32**| MembershipType of the account being cancelled. | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupV2KickMember**
> GroupV2KickMember(ctx, groupId, membershipId, membershipType)


Kick a member from the given group, forcing them to reapply if they wish to re-join the group. You must have suitable permissions in the group to perform this operation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **int64**| Group ID to kick the user from. | 
  **membershipId** | **int64**| Membership ID to kick. | 
  **membershipType** | **int32**| Membership type of the provided membership ID. | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupV2RecoverGroupForFounder**
> GroupV2RecoverGroupForFounder(ctx, groupType, membershipId, membershipType)


Allows a founder to manually recover a group they can see in game but not on bungie.net

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupType** | **int32**| Type of group the supplied member founded. | 
  **membershipId** | **int64**| Membership ID to for which to find founded groups. | 
  **membershipType** | **int32**| Membership type of the supplied membership ID. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupV2UnbanMember**
> GroupV2UnbanMember(ctx, groupId, membershipId, membershipType)


Unbans the requested member, allowing them to re-apply for membership.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **int64**|  | 
  **membershipId** | **int64**| Membership ID of the member to unban from the group | 
  **membershipType** | **int32**| Membership type of the provided membership ID. | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

