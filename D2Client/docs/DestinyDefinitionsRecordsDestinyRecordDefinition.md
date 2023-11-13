# DestinyDefinitionsRecordsDestinyRecordDefinition

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayProperties** | [***DestinyDefinitionsCommonDestinyDisplayPropertiesDefinition**](Destiny.Definitions.Common.DestinyDisplayPropertiesDefinition.md) |  | [optional] [default to null]
**Scope** | **int32** | Indicates whether this Record&#x27;s state is determined on a per-character or on an account-wide basis. | [optional] [default to null]
**PresentationInfo** | [***DestinyDefinitionsPresentationDestinyPresentationChildBlock**](Destiny.Definitions.Presentation.DestinyPresentationChildBlock.md) |  | [optional] [default to null]
**LoreHash** | **int32** |  | [optional] [default to null]
**ObjectiveHashes** | **[]int32** |  | [optional] [default to null]
**RecordValueStyle** | **int32** |  | [optional] [default to null]
**ForTitleGilding** | **bool** |  | [optional] [default to null]
**ShouldShowLargeIcons** | **bool** | A hint to show a large icon for a reward | [optional] [default to null]
**TitleInfo** | [***DestinyDefinitionsRecordsDestinyRecordTitleBlock**](Destiny.Definitions.Records.DestinyRecordTitleBlock.md) |  | [optional] [default to null]
**CompletionInfo** | [***DestinyDefinitionsRecordsDestinyRecordCompletionBlock**](Destiny.Definitions.Records.DestinyRecordCompletionBlock.md) |  | [optional] [default to null]
**StateInfo** | [***DestinyDefinitionsRecordsSchemaRecordStateBlock**](Destiny.Definitions.Records.SchemaRecordStateBlock.md) |  | [optional] [default to null]
**Requirements** | [***DestinyDefinitionsPresentationDestinyPresentationNodeRequirementsBlock**](Destiny.Definitions.Presentation.DestinyPresentationNodeRequirementsBlock.md) |  | [optional] [default to null]
**ExpirationInfo** | [***DestinyDefinitionsRecordsDestinyRecordExpirationBlock**](Destiny.Definitions.Records.DestinyRecordExpirationBlock.md) |  | [optional] [default to null]
**IntervalInfo** | [***AllOfDestinyDefinitionsRecordsDestinyRecordDefinitionIntervalInfo**](AllOfDestinyDefinitionsRecordsDestinyRecordDefinitionIntervalInfo.md) | Some records have multiple &#x27;interval&#x27; objectives, and the record may be claimed at each completed interval | [optional] [default to null]
**RewardItems** | [**[]DestinyDestinyItemQuantity**](Destiny.DestinyItemQuantity.md) | If there is any publicly available information about rewards earned for achieving this record, this is the list of those items.   However, note that some records intentionally have \&quot;hidden\&quot; rewards. These will not be returned in this list. | [optional] [default to null]
**PresentationNodeType** | **int32** |  | [optional] [default to null]
**TraitIds** | **[]string** |  | [optional] [default to null]
**TraitHashes** | **[]int32** |  | [optional] [default to null]
**ParentNodeHashes** | **[]int32** | A quick reference to presentation nodes that have this node as a child. Presentation nodes can be parented under multiple parents. | [optional] [default to null]
**Hash** | **int32** | The unique identifier for this entity. Guaranteed to be unique for the type of entity, but not globally.  When entities refer to each other in Destiny content, it is this hash that they are referring to. | [optional] [default to null]
**Index** | **int32** | The index of the entity as it was found in the investment tables. | [optional] [default to null]
**Redacted** | **bool** | If this is true, then there is an entity with this identifier/type combination, but BNet is not yet allowed to show it. Sorry! | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

