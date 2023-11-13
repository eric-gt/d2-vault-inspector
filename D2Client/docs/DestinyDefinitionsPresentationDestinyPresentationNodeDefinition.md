# DestinyDefinitionsPresentationDestinyPresentationNodeDefinition

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayProperties** | [***DestinyDefinitionsCommonDestinyDisplayPropertiesDefinition**](Destiny.Definitions.Common.DestinyDisplayPropertiesDefinition.md) |  | [optional] [default to null]
**OriginalIcon** | **string** | The original icon for this presentation node, before we futzed with it. | [optional] [default to null]
**RootViewIcon** | **string** | Some presentation nodes are meant to be explicitly shown on the \&quot;root\&quot; or \&quot;entry\&quot; screens for the feature to which they are related. You should use this icon when showing them on such a view, if you have a similar \&quot;entry point\&quot; view in your UI. If you don&#x27;t have a UI, then I guess it doesn&#x27;t matter either way does it? | [optional] [default to null]
**NodeType** | **int32** |  | [optional] [default to null]
**Scope** | **int32** | Indicates whether this presentation node&#x27;s state is determined on a per-character or on an account-wide basis. | [optional] [default to null]
**ObjectiveHash** | **int32** | If this presentation node shows a related objective (for instance, if it tracks the progress of its children), the objective being tracked is indicated here. | [optional] [default to null]
**CompletionRecordHash** | **int32** | If this presentation node has an associated \&quot;Record\&quot; that you can accomplish for completing its children, this is the identifier of that Record. | [optional] [default to null]
**Children** | [***AllOfDestinyDefinitionsPresentationDestinyPresentationNodeDefinitionChildren**](AllOfDestinyDefinitionsPresentationDestinyPresentationNodeDefinitionChildren.md) | The child entities contained by this presentation node. | [optional] [default to null]
**DisplayStyle** | **int32** | A hint for how to display this presentation node when it&#x27;s shown in a list. | [optional] [default to null]
**ScreenStyle** | **int32** | A hint for how to display this presentation node when it&#x27;s shown in its own detail screen. | [optional] [default to null]
**Requirements** | [***AllOfDestinyDefinitionsPresentationDestinyPresentationNodeDefinitionRequirements**](AllOfDestinyDefinitionsPresentationDestinyPresentationNodeDefinitionRequirements.md) | The requirements for being able to interact with this presentation node and its children. | [optional] [default to null]
**DisableChildSubscreenNavigation** | **bool** | If this presentation node has children, but the game doesn&#x27;t let you inspect the details of those children, that is indicated here. | [optional] [default to null]
**MaxCategoryRecordScore** | **int32** |  | [optional] [default to null]
**PresentationNodeType** | **int32** |  | [optional] [default to null]
**TraitIds** | **[]string** |  | [optional] [default to null]
**TraitHashes** | **[]int32** |  | [optional] [default to null]
**ParentNodeHashes** | **[]int32** | A quick reference to presentation nodes that have this node as a child. Presentation nodes can be parented under multiple parents. | [optional] [default to null]
**Hash** | **int32** | The unique identifier for this entity. Guaranteed to be unique for the type of entity, but not globally.  When entities refer to each other in Destiny content, it is this hash that they are referring to. | [optional] [default to null]
**Index** | **int32** | The index of the entity as it was found in the investment tables. | [optional] [default to null]
**Redacted** | **bool** | If this is true, then there is an entity with this identifier/type combination, but BNet is not yet allowed to show it. Sorry! | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

