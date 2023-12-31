# DestinyDefinitionsDestinyEntitySearchResultItem

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hash** | **int32** | The hash identifier of the entity. You will use this to look up the DestinyDefinition relevant for the entity found. | [optional] [default to null]
**EntityType** | **string** | The type of entity, returned as a string matching the DestinyDefinition&#x27;s contract class name. You&#x27;ll have to have your own mapping from class names to actually looking up those definitions in the manifest databases. | [optional] [default to null]
**DisplayProperties** | [***AllOfDestinyDefinitionsDestinyEntitySearchResultItemDisplayProperties**](AllOfDestinyDefinitionsDestinyEntitySearchResultItemDisplayProperties.md) | Basic display properties on the entity, so you don&#x27;t have to look up the definition to show basic results for the item. | [optional] [default to null]
**Weight** | **float64** | The ranking value for sorting that we calculated using our relevance formula. This will hopefully get better with time and iteration. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

