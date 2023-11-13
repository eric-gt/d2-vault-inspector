# DestinyResponsesDestinyVendorResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Vendor** | [***AllOfDestinyResponsesDestinyVendorResponseVendor**](AllOfDestinyResponsesDestinyVendorResponseVendor.md) | The base properties of the vendor.  COMPONENT TYPE: Vendors | [optional] [default to null]
**Categories** | [***AllOfDestinyResponsesDestinyVendorResponseCategories**](AllOfDestinyResponsesDestinyVendorResponseCategories.md) | Categories that the vendor has available, and references to the sales therein.  COMPONENT TYPE: VendorCategories | [optional] [default to null]
**Sales** | [***AllOfDestinyResponsesDestinyVendorResponseSales**](AllOfDestinyResponsesDestinyVendorResponseSales.md) | Sales, keyed by the vendorItemIndex of the item being sold.  COMPONENT TYPE: VendorSales | [optional] [default to null]
**ItemComponents** | [***AllOfDestinyResponsesDestinyVendorResponseItemComponents**](AllOfDestinyResponsesDestinyVendorResponseItemComponents.md) | Item components, keyed by the vendorItemIndex of the active sale items.  COMPONENT TYPE: [See inside the DestinyItemComponentSet contract for component types.] | [optional] [default to null]
**CurrencyLookups** | [***AllOfDestinyResponsesDestinyVendorResponseCurrencyLookups**](AllOfDestinyResponsesDestinyVendorResponseCurrencyLookups.md) | A \&quot;lookup\&quot; convenience component that can be used to quickly check if the character has access to items that can be used for purchasing.  COMPONENT TYPE: CurrencyLookups | [optional] [default to null]
**StringVariables** | [***AllOfDestinyResponsesDestinyVendorResponseStringVariables**](AllOfDestinyResponsesDestinyVendorResponseStringVariables.md) | A map of string variable values by hash for this character context.  COMPONENT TYPE: StringVariables | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

