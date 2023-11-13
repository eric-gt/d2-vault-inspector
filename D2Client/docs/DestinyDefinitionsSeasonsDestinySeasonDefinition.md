# DestinyDefinitionsSeasonsDestinySeasonDefinition

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayProperties** | [***DestinyDefinitionsCommonDestinyDisplayPropertiesDefinition**](Destiny.Definitions.Common.DestinyDisplayPropertiesDefinition.md) |  | [optional] [default to null]
**BackgroundImagePath** | **string** |  | [optional] [default to null]
**SeasonNumber** | **int32** |  | [optional] [default to null]
**StartDate** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**EndDate** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**SeasonPassHash** | **int32** |  | [optional] [default to null]
**SeasonPassProgressionHash** | **int32** |  | [optional] [default to null]
**ArtifactItemHash** | **int32** |  | [optional] [default to null]
**SealPresentationNodeHash** | **int32** |  | [optional] [default to null]
**SeasonalChallengesPresentationNodeHash** | **int32** |  | [optional] [default to null]
**Preview** | [***AllOfDestinyDefinitionsSeasonsDestinySeasonDefinitionPreview**](AllOfDestinyDefinitionsSeasonsDestinySeasonDefinitionPreview.md) | Optional - Defines the promotional text, images, and links to preview this season. | [optional] [default to null]
**Hash** | **int32** | The unique identifier for this entity. Guaranteed to be unique for the type of entity, but not globally.  When entities refer to each other in Destiny content, it is this hash that they are referring to. | [optional] [default to null]
**Index** | **int32** | The index of the entity as it was found in the investment tables. | [optional] [default to null]
**Redacted** | **bool** | If this is true, then there is an entity with this identifier/type combination, but BNet is not yet allowed to show it. Sorry! | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

