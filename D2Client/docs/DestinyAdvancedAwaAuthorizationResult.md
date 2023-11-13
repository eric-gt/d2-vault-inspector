# DestinyAdvancedAwaAuthorizationResult

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserSelection** | **int32** | Indication of how the user responded to the request. If the value is \&quot;Approved\&quot; the actionToken will contain the token that can be presented when performing the advanced write action. | [optional] [default to null]
**ResponseReason** | **int32** |  | [optional] [default to null]
**DeveloperNote** | **string** | Message to the app developer to help understand the response. | [optional] [default to null]
**ActionToken** | **string** | Credential used to prove the user authorized an advanced write action. | [optional] [default to null]
**MaximumNumberOfUses** | **int32** | This token may be used to perform the requested action this number of times, at a maximum. If this value is 0, then there is no limit. | [optional] [default to null]
**ValidUntil** | [**time.Time**](time.Time.md) | Time, UTC, when token expires. | [optional] [default to null]
**Type_** | **int32** | Advanced Write Action Type from the permission request. | [optional] [default to null]
**MembershipType** | **int32** | MembershipType from the permission request. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

