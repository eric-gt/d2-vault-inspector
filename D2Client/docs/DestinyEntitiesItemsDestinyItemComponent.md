# DestinyEntitiesItemsDestinyItemComponent

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ItemHash** | **int32** | The identifier for the item&#x27;s definition, which is where most of the useful static information for the item can be found. | [optional] [default to null]
**ItemInstanceId** | **int64** | If the item is instanced, it will have an instance ID. Lack of an instance ID implies that the item has no distinct local qualities aside from stack size. | [optional] [default to null]
**Quantity** | **int32** | The quantity of the item in this stack. Note that Instanced items cannot stack. If an instanced item, this value will always be 1 (as the stack has exactly one item in it) | [optional] [default to null]
**BindStatus** | **int32** | If the item is bound to a location, it will be specified in this enum. | [optional] [default to null]
**Location** | **int32** | An easy reference for where the item is located. Redundant if you got the item from an Inventory, but useful when making detail calls on specific items. | [optional] [default to null]
**BucketHash** | **int32** | The hash identifier for the specific inventory bucket in which the item is located. | [optional] [default to null]
**TransferStatus** | **int32** | If there is a known error state that would cause this item to not be transferable, this Flags enum will indicate all of those error states. Otherwise, it will be 0 (CanTransfer). | [optional] [default to null]
**Lockable** | **bool** | If the item can be locked, this will indicate that state. | [optional] [default to null]
**State** | **int32** | A flags enumeration indicating the transient/custom states of the item that affect how it is rendered: whether it&#x27;s tracked or locked for example, or whether it has a masterwork plug inserted. | [optional] [default to null]
**OverrideStyleItemHash** | **int32** | If populated, this is the hash of the item whose icon (and other secondary styles, but *not* the human readable strings) should override whatever icons/styles are on the item being sold.  If you don&#x27;t do this, certain items whose styles are being overridden by socketed items - such as the \&quot;Recycle Shader\&quot; item - would show whatever their default icon/style is, and it wouldn&#x27;t be pretty or look accurate. | [optional] [default to null]
**ExpirationDate** | [**time.Time**](time.Time.md) | If the item can expire, this is the date at which it will/did expire. | [optional] [default to null]
**IsWrapper** | **bool** | If this is true, the object is actually a \&quot;wrapper\&quot; of the object it&#x27;s representing. This means that it&#x27;s not the actual item itself, but rather an item that must be \&quot;opened\&quot; in game before you have and can use the item.   Wrappers are an evolution of \&quot;bundles\&quot;, which give an easy way to let you preview the contents of what you purchased while still letting you get a refund before you \&quot;open\&quot; it. | [optional] [default to null]
**TooltipNotificationIndexes** | **[]int32** | If this is populated, it is a list of indexes into DestinyInventoryItemDefinition.tooltipNotifications for any special tooltip messages that need to be shown for this item. | [optional] [default to null]
**MetricHash** | **int32** | The identifier for the currently-selected metric definition, to be displayed on the emblem nameplate. | [optional] [default to null]
**MetricObjective** | [***AllOfDestinyEntitiesItemsDestinyItemComponentMetricObjective**](AllOfDestinyEntitiesItemsDestinyItemComponentMetricObjective.md) | The objective progress for the currently-selected metric definition, to be displayed on the emblem nameplate. | [optional] [default to null]
**VersionNumber** | **int32** | The version of this item, used to index into the versions list in the item definition quality block. | [optional] [default to null]
**ItemValueVisibility** | **[]bool** | If available, a list that describes which item values (rewards) should be shown (true) or hidden (false). | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

