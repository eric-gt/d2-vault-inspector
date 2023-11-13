# DestinyComponentsRecordsDestinyProfileRecordsComponent

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Score** | **int32** | Your &#x27;active&#x27; Triumphs score, maintained for backwards compatibility. | [optional] [default to null]
**ActiveScore** | **int32** | Your &#x27;active&#x27; Triumphs score. | [optional] [default to null]
**LegacyScore** | **int32** | Your &#x27;legacy&#x27; Triumphs score. | [optional] [default to null]
**LifetimeScore** | **int32** | Your &#x27;lifetime&#x27; Triumphs score. | [optional] [default to null]
**TrackedRecordHash** | **int32** | If this profile is tracking a record, this is the hash identifier of the record it is tracking. | [optional] [default to null]
**Records** | [**map[string]DestinyComponentsRecordsDestinyRecordComponent**](Destiny.Components.Records.DestinyRecordComponent.md) |  | [optional] [default to null]
**RecordCategoriesRootNodeHash** | **int32** | The hash for the root presentation node definition of Triumph categories. | [optional] [default to null]
**RecordSealsRootNodeHash** | **int32** | The hash for the root presentation node definition of Triumph Seals. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

