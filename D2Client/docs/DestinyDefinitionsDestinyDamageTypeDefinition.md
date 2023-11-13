# DestinyDefinitionsDestinyDamageTypeDefinition

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayProperties** | [***AllOfDestinyDefinitionsDestinyDamageTypeDefinitionDisplayProperties**](AllOfDestinyDefinitionsDestinyDamageTypeDefinitionDisplayProperties.md) | The description of the damage type, icon etc... | [optional] [default to null]
**TransparentIconPath** | **string** | A variant of the icon that is transparent and colorless. | [optional] [default to null]
**ShowIcon** | **bool** | If TRUE, the game shows this damage type&#x27;s icon. Otherwise, it doesn&#x27;t. Whether you show it or not is up to you. | [optional] [default to null]
**EnumValue** | **int32** | We have an enumeration for damage types for quick reference. This is the current definition&#x27;s damage type enum value. | [optional] [default to null]
**Color** | [***AllOfDestinyDefinitionsDestinyDamageTypeDefinitionColor**](AllOfDestinyDefinitionsDestinyDamageTypeDefinitionColor.md) | A color associated with the damage type. The displayProperties icon is tinted with a color close to this. | [optional] [default to null]
**Hash** | **int32** | The unique identifier for this entity. Guaranteed to be unique for the type of entity, but not globally.  When entities refer to each other in Destiny content, it is this hash that they are referring to. | [optional] [default to null]
**Index** | **int32** | The index of the entity as it was found in the investment tables. | [optional] [default to null]
**Redacted** | **bool** | If this is true, then there is an entity with this identifier/type combination, but BNet is not yet allowed to show it. Sorry! | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

