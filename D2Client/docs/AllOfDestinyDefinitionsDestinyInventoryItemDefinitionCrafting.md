# AllOfDestinyDefinitionsDestinyInventoryItemDefinitionCrafting

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OutputItemHash** | **int32** | A reference to the item definition that is created when crafting with this &#x27;recipe&#x27; item. | [optional] [default to null]
**RequiredSocketTypeHashes** | **[]int32** | A list of socket type hashes that describes which sockets are required for crafting with this recipe. | [optional] [default to null]
**FailedRequirementStrings** | **[]string** |  | [optional] [default to null]
**BaseMaterialRequirements** | **int32** | A reference to the base material requirements for crafting with this recipe. | [optional] [default to null]
**BonusPlugs** | [**[]DestinyDefinitionsDestinyItemCraftingBlockBonusPlugDefinition**](Destiny.Definitions.DestinyItemCraftingBlockBonusPlugDefinition.md) | A list of &#x27;bonus&#x27; socket plugs that may be available if certain requirements are met. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

