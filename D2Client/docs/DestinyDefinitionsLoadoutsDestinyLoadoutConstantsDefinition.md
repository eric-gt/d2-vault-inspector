# DestinyDefinitionsLoadoutsDestinyLoadoutConstantsDefinition

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayProperties** | [***DestinyDefinitionsCommonDestinyDisplayPropertiesDefinition**](Destiny.Definitions.Common.DestinyDisplayPropertiesDefinition.md) |  | [optional] [default to null]
**WhiteIconImagePath** | **string** | This is the same icon as the one in the display properties, offered here as well with a more descriptive name. | [optional] [default to null]
**BlackIconImagePath** | **string** | This is a color-inverted version of the whiteIconImagePath. | [optional] [default to null]
**LoadoutCountPerCharacter** | **int32** | The maximum number of loadouts available to each character. The loadouts component API response can return fewer loadouts than this, as more loadouts are unlocked by reaching higher Guardian Ranks. | [optional] [default to null]
**LoadoutPreviewFilterOutSocketCategoryHashes** | **[]int32** | A list of the socket category hashes to be filtered out of loadout item preview displays. | [optional] [default to null]
**LoadoutPreviewFilterOutSocketTypeHashes** | **[]int32** | A list of the socket type hashes to be filtered out of loadout item preview displays. | [optional] [default to null]
**LoadoutNameHashes** | **[]int32** | A list of the loadout name hashes in index order, for convenience. | [optional] [default to null]
**LoadoutIconHashes** | **[]int32** | A list of the loadout icon hashes in index order, for convenience. | [optional] [default to null]
**LoadoutColorHashes** | **[]int32** | A list of the loadout color hashes in index order, for convenience. | [optional] [default to null]
**Hash** | **int32** | The unique identifier for this entity. Guaranteed to be unique for the type of entity, but not globally.  When entities refer to each other in Destiny content, it is this hash that they are referring to. | [optional] [default to null]
**Index** | **int32** | The index of the entity as it was found in the investment tables. | [optional] [default to null]
**Redacted** | **bool** | If this is true, then there is an entity with this identifier/type combination, but BNet is not yet allowed to show it. Sorry! | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

