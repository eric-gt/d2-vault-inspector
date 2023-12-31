# AllOfDestinyDefinitionsDestinyInventoryItemDefinitionSetData

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ItemList** | [**[]DestinyDefinitionsDestinyItemSetBlockEntryDefinition**](Destiny.Definitions.DestinyItemSetBlockEntryDefinition.md) | A collection of hashes of set items, for items such as Quest Metadata items that possess this data. | [optional] [default to null]
**RequireOrderedSetItemAdd** | **bool** | If true, items in the set can only be added in increasing order, and adding an item will remove any previous item. For Quests, this is by necessity true. Only one quest step is present at a time, and previous steps are removed as you advance in the quest. | [optional] [default to null]
**SetIsFeatured** | **bool** | If true, the UI should treat this quest as \&quot;featured\&quot; | [optional] [default to null]
**SetType** | **string** | A string identifier we can use to attempt to identify the category of the Quest. | [optional] [default to null]
**QuestLineName** | **string** | The name of the quest line that this quest step is a part of. | [optional] [default to null]
**QuestLineDescription** | **string** | The description of the quest line that this quest step is a part of. | [optional] [default to null]
**QuestStepSummary** | **string** | An additional summary of this step in the quest line. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

