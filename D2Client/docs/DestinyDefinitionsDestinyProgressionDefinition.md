# DestinyDefinitionsDestinyProgressionDefinition

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayProperties** | [***DestinyDefinitionsDestinyProgressionDisplayPropertiesDefinition**](Destiny.Definitions.DestinyProgressionDisplayPropertiesDefinition.md) |  | [optional] [default to null]
**Scope** | **int32** | The \&quot;Scope\&quot; of the progression indicates the source of the progression&#x27;s live data.  See the DestinyProgressionScope enum for more info: but essentially, a Progression can either be backed by a stored value, or it can be a calculated derivative of other values. | [optional] [default to null]
**RepeatLastStep** | **bool** | If this is True, then the progression doesn&#x27;t have a maximum level. | [optional] [default to null]
**Source** | **string** | If there&#x27;s a description of how to earn this progression in the local config, this will be that localized description. | [optional] [default to null]
**Steps** | [**[]DestinyDefinitionsDestinyProgressionStepDefinition**](Destiny.Definitions.DestinyProgressionStepDefinition.md) | Progressions are divided into Steps, which roughly equate to \&quot;Levels\&quot; in the traditional sense of a Progression. Notably, the last step can be repeated indefinitely if repeatLastStep is true, meaning that the calculation for your level is not as simple as comparing your current progress to the max progress of the steps.   These and more calculations are done for you if you grab live character progression data, such as in the DestinyCharacterProgressionComponent. | [optional] [default to null]
**Visible** | **bool** | If true, the Progression is something worth showing to users.  If false, BNet isn&#x27;t going to show it. But that doesn&#x27;t mean you can&#x27;t. We&#x27;re all friends here. | [optional] [default to null]
**FactionHash** | **int32** | If the value exists, this is the hash identifier for the Faction that owns this Progression.  This is purely for convenience, if you&#x27;re looking at a progression and want to know if and who it&#x27;s related to in terms of Faction Reputation. | [optional] [default to null]
**Color** | [***AllOfDestinyDefinitionsDestinyProgressionDefinitionColor**](AllOfDestinyDefinitionsDestinyProgressionDefinitionColor.md) | The #RGB string value for the color related to this progression, if there is one. | [optional] [default to null]
**RankIcon** | **string** | For progressions that have it, this is the rank icon we use in the Companion, displayed above the progressions&#x27; rank value. | [optional] [default to null]
**RewardItems** | [**[]DestinyDefinitionsDestinyProgressionRewardItemQuantity**](Destiny.Definitions.DestinyProgressionRewardItemQuantity.md) |  | [optional] [default to null]
**Hash** | **int32** | The unique identifier for this entity. Guaranteed to be unique for the type of entity, but not globally.  When entities refer to each other in Destiny content, it is this hash that they are referring to. | [optional] [default to null]
**Index** | **int32** | The index of the entity as it was found in the investment tables. | [optional] [default to null]
**Redacted** | **bool** | If this is true, then there is an entity with this identifier/type combination, but BNet is not yet allowed to show it. Sorry! | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

