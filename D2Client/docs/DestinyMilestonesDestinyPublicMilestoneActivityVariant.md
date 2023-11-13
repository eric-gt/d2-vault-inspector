# DestinyMilestonesDestinyPublicMilestoneActivityVariant

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivityHash** | **int32** | The hash identifier of this activity variant. Examine the activity&#x27;s definition in the Manifest database to determine what makes it a distinct variant. Usually it will be difficulty level or whether or not it is a guided game variant of the activity, but theoretically it could be distinguished in any arbitrary way. | [optional] [default to null]
**ActivityModeHash** | **int32** | The hash identifier of the most specific Activity Mode under which this activity is played. This is useful for situations where the activity in question is - for instance - a PVP map, but it&#x27;s not clear what mode the PVP map is being played under. If it&#x27;s a playlist, this will be less specific: but hopefully useful in some way. | [optional] [default to null]
**ActivityModeType** | **int32** | The enumeration equivalent of the most specific Activity Mode under which this activity is played. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

