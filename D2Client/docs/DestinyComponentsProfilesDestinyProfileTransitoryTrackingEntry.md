# DestinyComponentsProfilesDestinyProfileTransitoryTrackingEntry

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocationHash** | **int32** | OPTIONAL - If this is tracking a DestinyLocationDefinition, this is the identifier for that location. | [optional] [default to null]
**ItemHash** | **int32** | OPTIONAL - If this is tracking the status of a DestinyInventoryItemDefinition, this is the identifier for that item. | [optional] [default to null]
**ObjectiveHash** | **int32** | OPTIONAL - If this is tracking the status of a DestinyObjectiveDefinition, this is the identifier for that objective. | [optional] [default to null]
**ActivityHash** | **int32** | OPTIONAL - If this is tracking the status of a DestinyActivityDefinition, this is the identifier for that activity. | [optional] [default to null]
**QuestlineItemHash** | **int32** | OPTIONAL - If this is tracking the status of a quest, this is the identifier for the DestinyInventoryItemDefinition that containst that questline data. | [optional] [default to null]
**TrackedDate** | [**time.Time**](time.Time.md) | OPTIONAL - I&#x27;ve got to level with you, I don&#x27;t really know what this is. Is it when you started tracking it? Is it only populated for tracked items that have time limits?  I don&#x27;t know, but we can get at it - when I get time to actually test what it is, I&#x27;ll update this. In the meantime, bask in the mysterious data. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

