# DestinyComponentsItemsDestinyItemPlugComponent

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlugObjectives** | [**[]DestinyQuestsDestinyObjectiveProgress**](Destiny.Quests.DestinyObjectiveProgress.md) | Sometimes, Plugs may have objectives: these are often used for flavor and display purposes, but they can be used for any arbitrary purpose (both fortunately and unfortunately). Recently (with Season 2) they were expanded in use to be used as the \&quot;gating\&quot; for whether the plug can be inserted at all. For instance, a Plug might be tracking the number of PVP kills you have made. It will use the parent item&#x27;s data about that tracking status to determine what to show, and will generally show it using the DestinyObjectiveDefinition&#x27;s progressDescription property. Refer to the plug&#x27;s itemHash and objective property for more information if you would like to display even more data. | [optional] [default to null]
**PlugItemHash** | **int32** | The hash identifier of the DestinyInventoryItemDefinition that represents this plug. | [optional] [default to null]
**CanInsert** | **bool** | If true, this plug has met all of its insertion requirements. Big if true. | [optional] [default to null]
**Enabled** | **bool** | If true, this plug will provide its benefits while inserted. | [optional] [default to null]
**InsertFailIndexes** | **[]int32** | If the plug cannot be inserted for some reason, this will have the indexes into the plug item definition&#x27;s plug.insertionRules property, so you can show the reasons why it can&#x27;t be inserted.  This list will be empty if the plug can be inserted. | [optional] [default to null]
**EnableFailIndexes** | **[]int32** | If a plug is not enabled, this will be populated with indexes into the plug item definition&#x27;s plug.enabledRules property, so that you can show the reasons why it is not enabled.  This list will be empty if the plug is enabled. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

