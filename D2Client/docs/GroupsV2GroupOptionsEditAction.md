# GroupsV2GroupOptionsEditAction

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InvitePermissionOverride** | **bool** | Minimum Member Level allowed to invite new members to group  Always Allowed: Founder, Acting Founder  True means admins have this power, false means they don&#x27;t  Default is false for clans, true for groups. | [optional] [default to null]
**UpdateCulturePermissionOverride** | **bool** | Minimum Member Level allowed to update group culture  Always Allowed: Founder, Acting Founder  True means admins have this power, false means they don&#x27;t  Default is false for clans, true for groups. | [optional] [default to null]
**HostGuidedGamePermissionOverride** | **int32** | Minimum Member Level allowed to host guided games  Always Allowed: Founder, Acting Founder, Admin  Allowed Overrides: None, Member, Beginner  Default is Member for clans, None for groups, although this means nothing for groups. | [optional] [default to null]
**UpdateBannerPermissionOverride** | **bool** | Minimum Member Level allowed to update banner  Always Allowed: Founder, Acting Founder  True means admins have this power, false means they don&#x27;t  Default is false for clans, true for groups. | [optional] [default to null]
**JoinLevel** | **int32** | Level to join a member at when accepting an invite, application, or joining an open clan  Default is Beginner. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

