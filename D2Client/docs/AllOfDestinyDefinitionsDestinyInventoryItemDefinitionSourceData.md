# AllOfDestinyDefinitionsDestinyInventoryItemDefinitionSourceData

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceHashes** | **[]int32** | The list of hash identifiers for Reward Sources that hint where the item can be found (DestinyRewardSourceDefinition). | [optional] [default to null]
**Sources** | [**[]DestinyDefinitionsSourcesDestinyItemSourceDefinition**](Destiny.Definitions.Sources.DestinyItemSourceDefinition.md) | A collection of details about the stats that were computed for the ways we found that the item could be spawned. | [optional] [default to null]
**Exclusive** | **int32** | If we found that this item is exclusive to a specific platform, this will be set to the BungieMembershipType enumeration that matches that platform. | [optional] [default to null]
**VendorSources** | [**[]DestinyDefinitionsDestinyItemVendorSourceReference**](Destiny.Definitions.DestinyItemVendorSourceReference.md) | A denormalized reference back to vendors that potentially sell this item. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

