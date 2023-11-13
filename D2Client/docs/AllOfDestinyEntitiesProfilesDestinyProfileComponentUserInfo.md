# AllOfDestinyEntitiesProfilesDestinyProfileComponentUserInfo

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SupplementalDisplayName** | **string** | A platform specific additional display name - ex: psn Real Name, bnet Unique Name, etc. | [optional] [default to null]
**IconPath** | **string** | URL the Icon if available. | [optional] [default to null]
**CrossSaveOverride** | **int32** | If there is a cross save override in effect, this value will tell you the type that is overridding this one. | [optional] [default to null]
**ApplicableMembershipTypes** | **[]int32** | The list of Membership Types indicating the platforms on which this Membership can be used.   Not in Cross Save &#x3D; its original membership type. Cross Save Primary &#x3D; Any membership types it is overridding, and its original membership type Cross Save Overridden &#x3D; Empty list | [optional] [default to null]
**IsPublic** | **bool** | If True, this is a public user membership. | [optional] [default to null]
**MembershipType** | **int32** | Type of the membership. Not necessarily the native type. | [optional] [default to null]
**MembershipId** | **int64** | Membership ID as they user is known in the Accounts service | [optional] [default to null]
**DisplayName** | **string** | Display Name the player has chosen for themselves. The display name is optional when the data type is used as input to a platform API. | [optional] [default to null]
**BungieGlobalDisplayName** | **string** | The bungie global display name, if set. | [optional] [default to null]
**BungieGlobalDisplayNameCode** | **int32** | The bungie global display name code, if set. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

