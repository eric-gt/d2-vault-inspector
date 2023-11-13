# DestinyDefinitionsSocketsDestinySocketTypeDefinition

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayProperties** | [***AllOfDestinyDefinitionsSocketsDestinySocketTypeDefinitionDisplayProperties**](AllOfDestinyDefinitionsSocketsDestinySocketTypeDefinitionDisplayProperties.md) | There are fields for this display data, but they appear to be unpopulated as of now. I am not sure where in the UI these would show if they even were populated, but I will continue to return this data in case it becomes useful. | [optional] [default to null]
**InsertAction** | [***AllOfDestinyDefinitionsSocketsDestinySocketTypeDefinitionInsertAction**](AllOfDestinyDefinitionsSocketsDestinySocketTypeDefinitionInsertAction.md) | Defines what happens when a plug is inserted into sockets of this type. | [optional] [default to null]
**PlugWhitelist** | [**[]DestinyDefinitionsSocketsDestinyPlugWhitelistEntryDefinition**](Destiny.Definitions.Sockets.DestinyPlugWhitelistEntryDefinition.md) | A list of Plug \&quot;Categories\&quot; that are allowed to be plugged into sockets of this type.  These should be compared against a given plug item&#x27;s DestinyInventoryItemDefinition.plug.plugCategoryHash, which indicates the plug item&#x27;s category.  If the plug&#x27;s category matches any whitelisted plug, or if the whitelist is empty, it is allowed to be inserted. | [optional] [default to null]
**SocketCategoryHash** | **int32** |  | [optional] [default to null]
**Visibility** | **int32** | Sometimes a socket isn&#x27;t visible. These are some of the conditions under which sockets of this type are not visible. Unfortunately, the truth of visibility is much, much more complex. Best to rely on the live data for whether the socket is visible and enabled. | [optional] [default to null]
**AlwaysRandomizeSockets** | **bool** |  | [optional] [default to null]
**IsPreviewEnabled** | **bool** |  | [optional] [default to null]
**HideDuplicateReusablePlugs** | **bool** |  | [optional] [default to null]
**OverridesUiAppearance** | **bool** | This property indicates if the socket type determines whether Emblem icons and nameplates should be overridden by the inserted plug item&#x27;s icon and nameplate. | [optional] [default to null]
**AvoidDuplicatesOnInitialization** | **bool** |  | [optional] [default to null]
**CurrencyScalars** | [**[]DestinyDefinitionsSocketsDestinySocketTypeScalarMaterialRequirementEntry**](Destiny.Definitions.Sockets.DestinySocketTypeScalarMaterialRequirementEntry.md) |  | [optional] [default to null]
**Hash** | **int32** | The unique identifier for this entity. Guaranteed to be unique for the type of entity, but not globally.  When entities refer to each other in Destiny content, it is this hash that they are referring to. | [optional] [default to null]
**Index** | **int32** | The index of the entity as it was found in the investment tables. | [optional] [default to null]
**Redacted** | **bool** | If this is true, then there is an entity with this identifier/type combination, but BNet is not yet allowed to show it. Sorry! | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

