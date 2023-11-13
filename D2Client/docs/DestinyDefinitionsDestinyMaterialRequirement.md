# DestinyDefinitionsDestinyMaterialRequirement

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ItemHash** | **int32** | The hash identifier of the material required. Use it to look up the material&#x27;s DestinyInventoryItemDefinition. | [optional] [default to null]
**DeleteOnAction** | **bool** | If True, the material will be removed from the character&#x27;s inventory when the action is performed. | [optional] [default to null]
**Count** | **int32** | The amount of the material required. | [optional] [default to null]
**CountIsConstant** | **bool** | If true, the material requirement count value is constant. Since The Witch Queen expansion, some material requirement counts can be dynamic and will need to be returned with an API call. | [optional] [default to null]
**OmitFromRequirements** | **bool** | If True, this requirement is \&quot;silent\&quot;: don&#x27;t bother showing it in a material requirements display. I mean, I&#x27;m not your mom: I&#x27;m not going to tell you you *can&#x27;t* show it. But we won&#x27;t show it in our UI. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

