# AllOfDestinyDefinitionsDestinyTalentNodeDefinitionRandomActivationRequirement

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GridLevel** | **int32** | The Progression level on the Talent Grid required to activate this node.  See DestinyTalentGridDefinition.progressionHash for the related Progression, and read DestinyProgressionDefinition&#x27;s documentation to learn more about Progressions. | [optional] [default to null]
**MaterialRequirementHashes** | **[]int32** | The list of hash identifiers for material requirement sets: materials that are required for the node to be activated. See DestinyMaterialRequirementSetDefinition for more information about material requirements.  In this case, only a single DestinyMaterialRequirementSetDefinition will be chosen from this list, and we won&#x27;t know which one will be chosen until an instance of the item is created. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

