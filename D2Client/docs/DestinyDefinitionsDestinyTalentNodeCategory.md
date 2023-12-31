# DestinyDefinitionsDestinyTalentNodeCategory

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | **string** | Mostly just for debug purposes, but if you find it useful you can have it. This is BNet&#x27;s manually created identifier for this category. | [optional] [default to null]
**IsLoreDriven** | **bool** | If true, we found the localized content in a related DestinyLoreDefinition instead of local BNet localization files. This is mostly for ease of my own future investigations. | [optional] [default to null]
**DisplayProperties** | [***AllOfDestinyDefinitionsDestinyTalentNodeCategoryDisplayProperties**](AllOfDestinyDefinitionsDestinyTalentNodeCategoryDisplayProperties.md) | Will contain at least the \&quot;name\&quot;, which will be the title of the category. We will likely not have description and an icon yet, but I&#x27;m going to keep my options open. | [optional] [default to null]
**NodeHashes** | **[]int32** | The set of all hash identifiers for Talent Nodes (DestinyTalentNodeDefinition) in this Talent Grid that are part of this Category. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

