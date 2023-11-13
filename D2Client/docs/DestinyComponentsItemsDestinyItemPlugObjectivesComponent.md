# DestinyComponentsItemsDestinyItemPlugObjectivesComponent

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectivesPerPlug** | [**map[string][]DestinyQuestsDestinyObjectiveProgress**](array.md) | This set of data is keyed by the Item Hash (DestinyInventoryItemDefinition) of the plug whose objectives are being returned, with the value being the list of those objectives.   What if two plugs with the same hash are returned for an item, you ask?   Good question! They share the same item-scoped state, and as such would have identical objective state as a result. How&#x27;s that for convenient.   Sometimes, Plugs may have objectives: generally, these are used for flavor and display purposes. For instance, a Plug might be tracking the number of PVP kills you have made. It will use the parent item&#x27;s data about that tracking status to determine what to show, and will generally show it using the DestinyObjectiveDefinition&#x27;s progressDescription property. Refer to the plug&#x27;s itemHash and objective property for more information if you would like to display even more data. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

