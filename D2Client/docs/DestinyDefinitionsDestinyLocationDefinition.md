# DestinyDefinitionsDestinyLocationDefinition

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VendorHash** | **int32** | If the location has a Vendor on it, this is the hash identifier for that Vendor. Look them up with DestinyVendorDefinition. | [optional] [default to null]
**LocationReleases** | [**[]DestinyDefinitionsDestinyLocationReleaseDefinition**](Destiny.Definitions.DestinyLocationReleaseDefinition.md) | A Location may refer to different specific spots in the world based on the world&#x27;s current state. This is a list of those potential spots, and the data we can use at runtime to determine which one of the spots is the currently valid one. | [optional] [default to null]
**Hash** | **int32** | The unique identifier for this entity. Guaranteed to be unique for the type of entity, but not globally.  When entities refer to each other in Destiny content, it is this hash that they are referring to. | [optional] [default to null]
**Index** | **int32** | The index of the entity as it was found in the investment tables. | [optional] [default to null]
**Redacted** | **bool** | If this is true, then there is an entity with this identifier/type combination, but BNet is not yet allowed to show it. Sorry! | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

