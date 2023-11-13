# DestinyEntitiesCharactersDestinyCharacterActivitiesComponent

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DateActivityStarted** | [**time.Time**](time.Time.md) | The last date that the user started playing an activity. | [optional] [default to null]
**AvailableActivities** | [**[]DestinyDestinyActivity**](Destiny.DestinyActivity.md) | The list of activities that the user can play. | [optional] [default to null]
**CurrentActivityHash** | **int32** | If the user is in an activity, this will be the hash of the Activity being played. Note that you must combine this info with currentActivityModeHash to get a real picture of what the user is doing right now. For instance, PVP \&quot;Activities\&quot; are just maps: it&#x27;s the ActivityMode that determines what type of PVP game they&#x27;re playing. | [optional] [default to null]
**CurrentActivityModeHash** | **int32** | If the user is in an activity, this will be the hash of the activity mode being played. Combine with currentActivityHash to give a person a full picture of what they&#x27;re doing right now. | [optional] [default to null]
**CurrentActivityModeType** | **int32** | And the current activity&#x27;s most specific mode type, if it can be found. | [optional] [default to null]
**CurrentActivityModeHashes** | **[]int32** | If the user is in an activity, this will be the hashes of the DestinyActivityModeDefinition being played. Combine with currentActivityHash to give a person a full picture of what they&#x27;re doing right now. | [optional] [default to null]
**CurrentActivityModeTypes** | **[]int32** | All Activity Modes that apply to the current activity being played, in enum form. | [optional] [default to null]
**CurrentPlaylistActivityHash** | **int32** | If the user is in a playlist, this is the hash identifier for the playlist that they chose. | [optional] [default to null]
**LastCompletedStoryHash** | **int32** | This will have the activity hash of the last completed story/campaign mission, in case you care about that. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

