# AllOfDestinyEntitiesItemsDestinyItemInstanceComponentEnergy

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnergyTypeHash** | **int32** | The type of energy for this item. Plugs that require Energy can only be inserted if they have the \&quot;Any\&quot; Energy Type or the matching energy type of this item. This is a reference to the DestinyEnergyTypeDefinition for the energy type, where you can find extended info about it. | [optional] [default to null]
**EnergyType** | **int32** | This is the enum version of the Energy Type value, for convenience. | [optional] [default to null]
**EnergyCapacity** | **int32** | The total capacity of Energy that the item currently has, regardless of if it is currently being used. | [optional] [default to null]
**EnergyUsed** | **int32** | The amount of Energy currently in use by inserted plugs. | [optional] [default to null]
**EnergyUnused** | **int32** | The amount of energy still available for inserting new plugs. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

