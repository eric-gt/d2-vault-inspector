# DestinyResponsesDestinyCollectibleNodeDetailResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collectibles** | [***AllOfDestinyResponsesDestinyCollectibleNodeDetailResponseCollectibles**](AllOfDestinyResponsesDestinyCollectibleNodeDetailResponseCollectibles.md) | COMPONENT TYPE: Collectibles | [optional] [default to null]
**CollectibleItemComponents** | [***AllOfDestinyResponsesDestinyCollectibleNodeDetailResponseCollectibleItemComponents**](AllOfDestinyResponsesDestinyCollectibleNodeDetailResponseCollectibleItemComponents.md) | Item components, keyed by the item hash of the items pointed at collectibles found under the requested Presentation Node.  NOTE: I had a lot of hemming and hawing about whether these should be keyed by collectible hash or item hash... but ultimately having it be keyed by item hash meant that UI that already uses DestinyItemComponentSet data wouldn&#x27;t have to have a special override to do the collectible -&gt; item lookup once you delve into an item&#x27;s details, and it also meant that you didn&#x27;t have to remember that the Hash being used as the key for plugSets was different from the Hash being used for the other Dictionaries. As a result, using the Item Hash felt like the least crappy solution.  We may all come to regret this decision. We will see.  COMPONENT TYPE: [See inside the DestinyItemComponentSet contract for component types.] | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

