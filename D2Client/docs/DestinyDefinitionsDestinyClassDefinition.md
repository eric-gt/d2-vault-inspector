# DestinyDefinitionsDestinyClassDefinition

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassType** | **int32** | In Destiny 1, we added a convenience Enumeration for referring to classes. We&#x27;ve kept it, though mostly for posterity. This is the enum value for this definition&#x27;s class. | [optional] [default to null]
**DisplayProperties** | [***DestinyDefinitionsCommonDestinyDisplayPropertiesDefinition**](Destiny.Definitions.Common.DestinyDisplayPropertiesDefinition.md) |  | [optional] [default to null]
**GenderedClassNames** | **map[string]string** | A localized string referring to the singular form of the Class&#x27;s name when referred to in gendered form. Keyed by the DestinyGender. | [optional] [default to null]
**GenderedClassNamesByGenderHash** | **map[string]string** |  | [optional] [default to null]
**MentorVendorHash** | **int32** | Mentors don&#x27;t really mean anything anymore. Don&#x27;t expect this to be populated. | [optional] [default to null]
**Hash** | **int32** | The unique identifier for this entity. Guaranteed to be unique for the type of entity, but not globally.  When entities refer to each other in Destiny content, it is this hash that they are referring to. | [optional] [default to null]
**Index** | **int32** | The index of the entity as it was found in the investment tables. | [optional] [default to null]
**Redacted** | **bool** | If this is true, then there is an entity with this identifier/type combination, but BNet is not yet allowed to show it. Sorry! | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

