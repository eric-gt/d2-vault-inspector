# DestinyComponentsVendorsDestinyPublicVendorComponent

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VendorHash** | **int32** | The unique identifier for the vendor. Use it to look up their DestinyVendorDefinition. | [optional] [default to null]
**NextRefreshDate** | [**time.Time**](time.Time.md) | The date when this vendor&#x27;s inventory will next rotate/refresh.  Note that this is distinct from the date ranges that the vendor is visible/available in-game: this field indicates the specific time when the vendor&#x27;s available items refresh and rotate, regardless of whether the vendor is actually available at that time. Unfortunately, these two values may be (and are, for the case of important vendors like Xur) different.  Issue https://github.com/Bungie-net/api/issues/353 is tracking a fix to start providing visibility date ranges where possible in addition to this refresh date, so that all important dates for vendors are available for use. | [optional] [default to null]
**Enabled** | **bool** | If True, the Vendor is currently accessible.   If False, they may not actually be visible in the world at the moment. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

