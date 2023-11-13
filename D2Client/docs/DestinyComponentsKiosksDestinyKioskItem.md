# DestinyComponentsKiosksDestinyKioskItem

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | **int32** | The index of the item in the related DestinyVendorDefintion&#x27;s itemList property, representing the sale. | [optional] [default to null]
**CanAcquire** | **bool** | If true, the user can not only see the item, but they can acquire it. It is possible that a user can see a kiosk item and not be able to acquire it. | [optional] [default to null]
**FailureIndexes** | **[]int32** | Indexes into failureStrings for the Vendor, indicating the reasons why it failed if any. | [optional] [default to null]
**FlavorObjective** | [***AllOfDestinyComponentsKiosksDestinyKioskItemFlavorObjective**](AllOfDestinyComponentsKiosksDestinyKioskItemFlavorObjective.md) | I may regret naming it this way - but this represents when an item has an objective that doesn&#x27;t serve a beneficial purpose, but rather is used for \&quot;flavor\&quot; or additional information. For instance, when Emblems track specific stats, those stats are represented as Objectives on the item. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

