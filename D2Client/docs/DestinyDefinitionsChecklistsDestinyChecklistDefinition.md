# DestinyDefinitionsChecklistsDestinyChecklistDefinition

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayProperties** | [***DestinyDefinitionsCommonDestinyDisplayPropertiesDefinition**](Destiny.Definitions.Common.DestinyDisplayPropertiesDefinition.md) |  | [optional] [default to null]
**ViewActionString** | **string** | A localized string prompting you to view the checklist. | [optional] [default to null]
**Scope** | **int32** | Indicates whether you will find this checklist on the Profile or Character components. | [optional] [default to null]
**Entries** | [**[]DestinyDefinitionsChecklistsDestinyChecklistEntryDefinition**](Destiny.Definitions.Checklists.DestinyChecklistEntryDefinition.md) | The individual checklist items. Gotta catch &#x27;em all. | [optional] [default to null]
**Hash** | **int32** | The unique identifier for this entity. Guaranteed to be unique for the type of entity, but not globally.  When entities refer to each other in Destiny content, it is this hash that they are referring to. | [optional] [default to null]
**Index** | **int32** | The index of the entity as it was found in the investment tables. | [optional] [default to null]
**Redacted** | **bool** | If this is true, then there is an entity with this identifier/type combination, but BNet is not yet allowed to show it. Sorry! | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

