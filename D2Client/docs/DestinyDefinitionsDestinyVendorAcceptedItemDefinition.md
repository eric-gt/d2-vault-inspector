# DestinyDefinitionsDestinyVendorAcceptedItemDefinition

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceptedInventoryBucketHash** | **int32** | The \&quot;source\&quot; bucket for a transfer. When a user wants to transfer an item, the appropriate DestinyVendorDefinition&#x27;s acceptedItems property is evaluated, looking for an entry where acceptedInventoryBucketHash matches the bucket that the item being transferred is currently located. If it exists, the item will be transferred into whatever bucket is defined by destinationInventoryBucketHash. | [optional] [default to null]
**DestinationInventoryBucketHash** | **int32** | This is the bucket where the item being transferred will be put, given that it was being transferred *from* the bucket defined in acceptedInventoryBucketHash. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

