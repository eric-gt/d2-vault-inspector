# DestinyDefinitionsCollectiblesDestinyCollectibleDefinition

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayProperties** | [***DestinyDefinitionsCommonDestinyDisplayPropertiesDefinition**](Destiny.Definitions.Common.DestinyDisplayPropertiesDefinition.md) |  | [optional] [default to null]
**Scope** | **int32** | Indicates whether the state of this Collectible is determined on a per-character or on an account-wide basis. | [optional] [default to null]
**SourceString** | **string** | A human readable string for a hint about how to acquire the item. | [optional] [default to null]
**SourceHash** | **int32** | This is a hash identifier we are building on the BNet side in an attempt to let people group collectibles by similar sources.  I can&#x27;t promise that it&#x27;s going to be 100% accurate, but if the designers were consistent in assigning the same source strings to items with the same sources, it *ought to* be. No promises though.  This hash also doesn&#x27;t relate to an actual definition, just to note: we&#x27;ve got nothing useful other than the source string for this data. | [optional] [default to null]
**ItemHash** | **int32** |  | [optional] [default to null]
**AcquisitionInfo** | [***DestinyDefinitionsCollectiblesDestinyCollectibleAcquisitionBlock**](Destiny.Definitions.Collectibles.DestinyCollectibleAcquisitionBlock.md) |  | [optional] [default to null]
**StateInfo** | [***DestinyDefinitionsCollectiblesDestinyCollectibleStateBlock**](Destiny.Definitions.Collectibles.DestinyCollectibleStateBlock.md) |  | [optional] [default to null]
**PresentationInfo** | [***DestinyDefinitionsPresentationDestinyPresentationChildBlock**](Destiny.Definitions.Presentation.DestinyPresentationChildBlock.md) |  | [optional] [default to null]
**PresentationNodeType** | **int32** |  | [optional] [default to null]
**TraitIds** | **[]string** |  | [optional] [default to null]
**TraitHashes** | **[]int32** |  | [optional] [default to null]
**ParentNodeHashes** | **[]int32** | A quick reference to presentation nodes that have this node as a child. Presentation nodes can be parented under multiple parents. | [optional] [default to null]
**Hash** | **int32** | The unique identifier for this entity. Guaranteed to be unique for the type of entity, but not globally.  When entities refer to each other in Destiny content, it is this hash that they are referring to. | [optional] [default to null]
**Index** | **int32** | The index of the entity as it was found in the investment tables. | [optional] [default to null]
**Redacted** | **bool** | If this is true, then there is an entity with this identifier/type combination, but BNet is not yet allowed to show it. Sorry! | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

