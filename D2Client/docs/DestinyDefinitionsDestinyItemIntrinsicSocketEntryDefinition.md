# DestinyDefinitionsDestinyItemIntrinsicSocketEntryDefinition

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlugItemHash** | **int32** | Indicates the plug that is intrinsically inserted into this socket. | [optional] [default to null]
**SocketTypeHash** | **int32** | Indicates the type of this intrinsic socket. | [optional] [default to null]
**DefaultVisible** | **bool** | If true, then this socket is visible in the item&#x27;s \&quot;default\&quot; state. If you have an instance, you should always check the runtime state, as that can override this visibility setting: but if you&#x27;re looking at the item on a conceptual level, this property can be useful for hiding data such as legacy sockets - which remain defined on items for infrastructure purposes, but can be confusing for users to see. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

