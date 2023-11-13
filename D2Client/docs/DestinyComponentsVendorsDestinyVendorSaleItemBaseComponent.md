# DestinyComponentsVendorsDestinyVendorSaleItemBaseComponent

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VendorItemIndex** | **int32** | The index into the DestinyVendorDefinition.itemList property. Note that this means Vendor data *is* Content Version dependent: make sure you have the latest content before you use Vendor data, or these indexes may mismatch.   Most systems avoid this problem, but Vendors is one area where we are unable to reasonably avoid content dependency at the moment. | [optional] [default to null]
**ItemHash** | **int32** | The hash of the item being sold, as a quick shortcut for looking up the DestinyInventoryItemDefinition of the sale item. | [optional] [default to null]
**OverrideStyleItemHash** | **int32** | If populated, this is the hash of the item whose icon (and other secondary styles, but *not* the human readable strings) should override whatever icons/styles are on the item being sold.  If you don&#x27;t do this, certain items whose styles are being overridden by socketed items - such as the \&quot;Recycle Shader\&quot; item - would show whatever their default icon/style is, and it wouldn&#x27;t be pretty or look accurate. | [optional] [default to null]
**Quantity** | **int32** | How much of the item you&#x27;ll be getting. | [optional] [default to null]
**Costs** | [**[]DestinyDestinyItemQuantity**](Destiny.DestinyItemQuantity.md) | A summary of the current costs of the item. | [optional] [default to null]
**OverrideNextRefreshDate** | [**time.Time**](time.Time.md) | If this item has its own custom date where it may be removed from the Vendor&#x27;s rotation, this is that date.  Note that there&#x27;s not actually any guarantee that it will go away: it could be chosen again and end up still being in the Vendor&#x27;s sale items! But this is the next date where that test will occur, and is also the date that the game shows for availability on things like Bounties being sold. So it&#x27;s the best we can give. | [optional] [default to null]
**ApiPurchasable** | **bool** | If true, this item can be purchased through the Bungie.net API. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

