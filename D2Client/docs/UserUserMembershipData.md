# UserUserMembershipData

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinyMemberships** | [**[]GroupsV2GroupUserInfoCard**](GroupsV2.GroupUserInfoCard.md) | this allows you to see destiny memberships that are visible and linked to this account (regardless of whether or not they have characters on the world server) | [optional] [default to null]
**PrimaryMembershipId** | **int64** | If this property is populated, it will have the membership ID of the account considered to be \&quot;primary\&quot; in this user&#x27;s cross save relationship.   If null, this user has no cross save relationship, nor primary account. | [optional] [default to null]
**BungieNetUser** | [***UserGeneralUser**](User.GeneralUser.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

