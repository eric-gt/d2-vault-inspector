# DestinyComponentsProfilesDestinyProfileTransitoryCurrentActivity

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTime** | [**time.Time**](time.Time.md) | When the activity started. | [optional] [default to null]
**EndTime** | [**time.Time**](time.Time.md) | If you&#x27;re still in it but it \&quot;ended\&quot; (like when folks are dancing around the loot after they beat a boss), this is when the activity ended. | [optional] [default to null]
**Score** | **float32** | This is what our non-authoritative source thought the score was. | [optional] [default to null]
**HighestOpposingFactionScore** | **float32** | If you have human opponents, this is the highest opposing team&#x27;s score. | [optional] [default to null]
**NumberOfOpponents** | **int32** | This is how many human or poorly crafted aimbot opponents you have. | [optional] [default to null]
**NumberOfPlayers** | **int32** | This is how many human or poorly crafted aimbots are on your team. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

