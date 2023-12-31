# DestinyDefinitionsSourcesDestinyItemSourceDefinition

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Level** | **int32** | The level at which the item spawns. Essentially the Primary Key for this source data: there will be multiple of these source entries per item that has source data, grouped by the level at which the item spawns. | [optional] [default to null]
**MinQuality** | **int32** | The minimum Quality at which the item spawns for this level. Examine DestinyInventoryItemDefinition for more information about what Quality means. Just don&#x27;t ask Phaedrus about it, he&#x27;ll never stop talking and you&#x27;ll have to write a book about it. | [optional] [default to null]
**MaxQuality** | **int32** | The maximum quality at which the item spawns for this level. | [optional] [default to null]
**MinLevelRequired** | **int32** | The minimum Character Level required for equipping the item when the item spawns at the item level defined on this DestinyItemSourceDefinition, as far as we saw in our processing. | [optional] [default to null]
**MaxLevelRequired** | **int32** | The maximum Character Level required for equipping the item when the item spawns at the item level defined on this DestinyItemSourceDefinition, as far as we saw in our processing. | [optional] [default to null]
**ComputedStats** | [**map[string]DestinyDefinitionsDestinyInventoryItemStatDefinition**](Destiny.Definitions.DestinyInventoryItemStatDefinition.md) | The stats computed for this level/quality range. | [optional] [default to null]
**SourceHashes** | **[]int32** | The DestinyRewardSourceDefinitions found that can spawn the item at this level. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

