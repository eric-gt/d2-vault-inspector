# DestinyDefinitionsDestinyVendorCategoryEntryDefinition

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CategoryIndex** | **int32** | The index of the category in the original category definitions for the vendor. | [optional] [default to null]
**SortValue** | **int32** | Used in sorting items in vendors... but there&#x27;s a lot more to it. Just go with the order provided in the itemIndexes property on the DestinyVendorCategoryComponent instead, it should be more reliable than trying to recalculate it yourself. | [optional] [default to null]
**CategoryHash** | **int32** | The hashed identifier for the category. | [optional] [default to null]
**QuantityAvailable** | **int32** | The amount of items that will be available when this category is shown. | [optional] [default to null]
**ShowUnavailableItems** | **bool** | If items aren&#x27;t up for sale in this category, should we still show them (greyed out)? | [optional] [default to null]
**HideIfNoCurrency** | **bool** | If you don&#x27;t have the currency required to buy items from this category, should the items be hidden? | [optional] [default to null]
**HideFromRegularPurchase** | **bool** | True if this category doesn&#x27;t allow purchases. | [optional] [default to null]
**BuyStringOverride** | **string** | The localized string for making purchases from this category, if it is different from the vendor&#x27;s string for purchasing. | [optional] [default to null]
**DisabledDescription** | **string** | If the category is disabled, this is the localized description to show. | [optional] [default to null]
**DisplayTitle** | **string** | The localized title of the category. | [optional] [default to null]
**Overlay** | [***AllOfDestinyDefinitionsDestinyVendorCategoryEntryDefinitionOverlay**](AllOfDestinyDefinitionsDestinyVendorCategoryEntryDefinitionOverlay.md) | If this category has an overlay prompt that should appear, this contains the details of that prompt. | [optional] [default to null]
**VendorItemIndexes** | **[]int32** | A shortcut for the vendor item indexes sold under this category. Saves us from some expensive reorganization at runtime. | [optional] [default to null]
**IsPreview** | **bool** | Sometimes a category isn&#x27;t actually used to sell items, but rather to preview them. This implies different UI (and manual placement of the category in the UI) in the game, and special treatment. | [optional] [default to null]
**IsDisplayOnly** | **bool** | If true, this category only displays items: you can&#x27;t purchase anything in them. | [optional] [default to null]
**ResetIntervalMinutesOverride** | **int32** |  | [optional] [default to null]
**ResetOffsetMinutesOverride** | **int32** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

