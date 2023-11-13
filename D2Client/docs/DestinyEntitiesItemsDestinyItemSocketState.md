# DestinyEntitiesItemsDestinyItemSocketState

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlugHash** | **int32** | The currently active plug, if any.  Note that, because all plugs are statically defined, its effect on stats and perks can be statically determined using the plug item&#x27;s definition. The stats and perks can be taken at face value on the plug item as the stats and perks it will provide to the user/item. | [optional] [default to null]
**IsEnabled** | **bool** | Even if a plug is inserted, it doesn&#x27;t mean it&#x27;s enabled.  This flag indicates whether the plug is active and providing its benefits. | [optional] [default to null]
**IsVisible** | **bool** | A plug may theoretically provide benefits but not be visible - for instance, some older items use a plug&#x27;s damage type perk to modify their own damage type. These, though they are not visible, still affect the item. This field indicates that state.  An invisible plug, while it provides benefits if it is Enabled, cannot be directly modified by the user. | [optional] [default to null]
**EnableFailIndexes** | **[]int32** | If a plug is inserted but not enabled, this will be populated with indexes into the plug item definition&#x27;s plug.enabledRules property, so that you can show the reasons why it is not enabled. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

