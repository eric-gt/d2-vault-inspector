# DestinyComponentsPresentationDestinyPresentationNodeComponent

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | **int32** |  | [optional] [default to null]
**Objective** | [***AllOfDestinyComponentsPresentationDestinyPresentationNodeComponentObjective**](AllOfDestinyComponentsPresentationDestinyPresentationNodeComponentObjective.md) | An optional property: presentation nodes MAY have objectives, which can be used to infer more human readable data about the progress. However, progressValue and completionValue ought to be considered the canonical values for progress on Progression Nodes. | [optional] [default to null]
**ProgressValue** | **int32** | How much of the presentation node is considered to be completed so far by the given character/profile. | [optional] [default to null]
**CompletionValue** | **int32** | The value at which the presentation node is considered to be completed. | [optional] [default to null]
**RecordCategoryScore** | **int32** | If available, this is the current score for the record category that this node represents. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

