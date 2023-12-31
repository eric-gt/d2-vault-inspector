# DestinyEntitiesVendorsDestinyVendorComponent

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CanPurchase** | **bool** | If True, you can purchase from the Vendor. | [optional] [default to null]
**Progression** | [***AllOfDestinyEntitiesVendorsDestinyVendorComponentProgression**](AllOfDestinyEntitiesVendorsDestinyVendorComponentProgression.md) | If the Vendor has a related Reputation, this is the Progression data that represents the character&#x27;s Reputation level with this Vendor. | [optional] [default to null]
**VendorLocationIndex** | **int32** | An index into the vendor definition&#x27;s \&quot;locations\&quot; property array, indicating which location they are at currently. If -1, then the vendor has no known location (and you may choose not to show them in your UI as a result. I mean, it&#x27;s your bag honey) | [optional] [default to null]
**SeasonalRank** | **int32** | If this vendor has a seasonal rank, this will be the calculated value of that rank. How nice is that? I mean, that&#x27;s pretty sweeet. It&#x27;s a whole 32 bit integer. | [optional] [default to null]
**VendorHash** | **int32** | The unique identifier for the vendor. Use it to look up their DestinyVendorDefinition. | [optional] [default to null]
**NextRefreshDate** | [**time.Time**](time.Time.md) | The date when this vendor&#x27;s inventory will next rotate/refresh.  Note that this is distinct from the date ranges that the vendor is visible/available in-game: this field indicates the specific time when the vendor&#x27;s available items refresh and rotate, regardless of whether the vendor is actually available at that time. Unfortunately, these two values may be (and are, for the case of important vendors like Xur) different.  Issue https://github.com/Bungie-net/api/issues/353 is tracking a fix to start providing visibility date ranges where possible in addition to this refresh date, so that all important dates for vendors are available for use. | [optional] [default to null]
**Enabled** | **bool** | If True, the Vendor is currently accessible.   If False, they may not actually be visible in the world at the moment. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

