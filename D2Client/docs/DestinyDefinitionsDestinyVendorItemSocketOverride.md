# DestinyDefinitionsDestinyVendorItemSocketOverride

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SingleItemHash** | **int32** | If this is populated, the socket will be overridden with a specific plug.  If this isn&#x27;t populated, it&#x27;s being overridden by something more complicated that is only known by the Game Server and God, which means we can&#x27;t tell you in advance what it&#x27;ll be. | [optional] [default to null]
**RandomizedOptionsCount** | **int32** | If this is greater than -1, the number of randomized plugs on this socket will be set to this quantity instead of whatever it&#x27;s set to by default. | [optional] [default to null]
**SocketTypeHash** | **int32** | This appears to be used to select which socket ultimately gets the override defined here. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

