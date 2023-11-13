# DestinyHistoricalStatsDestinyHistoricalStatsActivity

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReferenceId** | **int32** | The unique hash identifier of the DestinyActivityDefinition that was played. If I had this to do over, it&#x27;d be named activityHash. Too late now. | [optional] [default to null]
**DirectorActivityHash** | **int32** | The unique hash identifier of the DestinyActivityDefinition that was played. | [optional] [default to null]
**InstanceId** | **int64** | The unique identifier for this *specific* match that was played.  This value can be used to get additional data about this activity such as who else was playing via the GetPostGameCarnageReport endpoint. | [optional] [default to null]
**Mode** | **int32** | Indicates the most specific game mode of the activity that we could find. | [optional] [default to null]
**Modes** | **[]int32** | The list of all Activity Modes to which this activity applies, including aggregates. This will let you see, for example, whether the activity was both Clash and part of the Trials of the Nine event. | [optional] [default to null]
**IsPrivate** | **bool** | Whether or not the match was a private match. | [optional] [default to null]
**MembershipType** | **int32** | The Membership Type indicating the platform on which this match was played. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

