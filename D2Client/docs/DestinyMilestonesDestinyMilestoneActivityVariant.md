# DestinyMilestonesDestinyMilestoneActivityVariant

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivityHash** | **int32** | The hash for the specific variant of the activity related to this milestone. You can pull more detailed static info from the DestinyActivityDefinition, such as difficulty level. | [optional] [default to null]
**CompletionStatus** | [***AllOfDestinyMilestonesDestinyMilestoneActivityVariantCompletionStatus**](AllOfDestinyMilestonesDestinyMilestoneActivityVariantCompletionStatus.md) | An OPTIONAL component: if it makes sense to talk about this activity variant in terms of whether or not it has been completed or what progress you have made in it, this will be returned. Otherwise, this will be NULL. | [optional] [default to null]
**ActivityModeHash** | **int32** | The hash identifier of the most specific Activity Mode under which this activity is played. This is useful for situations where the activity in question is - for instance - a PVP map, but it&#x27;s not clear what mode the PVP map is being played under. If it&#x27;s a playlist, this will be less specific: but hopefully useful in some way. | [optional] [default to null]
**ActivityModeType** | **int32** | The enumeration equivalent of the most specific Activity Mode under which this activity is played. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

