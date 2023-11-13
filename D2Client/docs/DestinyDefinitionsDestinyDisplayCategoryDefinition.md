# DestinyDefinitionsDestinyDisplayCategoryDefinition

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | **int32** |  | [optional] [default to null]
**Identifier** | **string** | A string identifier for the display category. | [optional] [default to null]
**DisplayCategoryHash** | **int32** |  | [optional] [default to null]
**DisplayProperties** | [***DestinyDefinitionsCommonDestinyDisplayPropertiesDefinition**](Destiny.Definitions.Common.DestinyDisplayPropertiesDefinition.md) |  | [optional] [default to null]
**DisplayInBanner** | **bool** | If true, this category should be displayed in the \&quot;Banner\&quot; section of the vendor&#x27;s UI. | [optional] [default to null]
**ProgressionHash** | **int32** | If it exists, this is the hash identifier of a DestinyProgressionDefinition that represents the progression to show on this display category.  Specific categories can now have thier own distinct progression, apparently. So that&#x27;s cool. | [optional] [default to null]
**SortOrder** | **int32** | If this category sorts items in a nonstandard way, this will be the way we sort. | [optional] [default to null]
**DisplayStyleHash** | **int32** | An indicator of how the category will be displayed in the UI. It&#x27;s up to you to do something cool or interesting in response to this, or just to treat it as a normal category. | [optional] [default to null]
**DisplayStyleIdentifier** | **string** | An indicator of how the category will be displayed in the UI. It&#x27;s up to you to do something cool or interesting in response to this, or just to treat it as a normal category. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

