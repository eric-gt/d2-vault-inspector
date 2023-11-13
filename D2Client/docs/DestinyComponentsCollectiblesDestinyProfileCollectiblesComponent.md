# DestinyComponentsCollectiblesDestinyProfileCollectiblesComponent

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecentCollectibleHashes** | **[]int32** | The list of collectibles determined by the game as having been \&quot;recently\&quot; acquired. | [optional] [default to null]
**NewnessFlaggedCollectibleHashes** | **[]int32** | The list of collectibles determined by the game as having been \&quot;recently\&quot; acquired.  The game client itself actually controls this data, so I personally question whether anyone will get much use out of this: because we can&#x27;t edit this value through the API. But in case anyone finds it useful, here it is. | [optional] [default to null]
**Collectibles** | [**map[string]DestinyComponentsCollectiblesDestinyCollectibleComponent**](Destiny.Components.Collectibles.DestinyCollectibleComponent.md) |  | [optional] [default to null]
**CollectionCategoriesRootNodeHash** | **int32** | The hash for the root presentation node definition of Collection categories. | [optional] [default to null]
**CollectionBadgesRootNodeHash** | **int32** | The hash for the root presentation node definition of Collection Badges. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

