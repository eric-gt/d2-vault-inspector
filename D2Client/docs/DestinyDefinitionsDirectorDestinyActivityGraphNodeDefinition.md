# DestinyDefinitionsDirectorDestinyActivityGraphNodeDefinition

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeId** | **int32** | An identifier for the Activity Graph Node, only guaranteed to be unique within its parent Activity Graph. | [optional] [default to null]
**OverrideDisplay** | [***AllOfDestinyDefinitionsDirectorDestinyActivityGraphNodeDefinitionOverrideDisplay**](AllOfDestinyDefinitionsDirectorDestinyActivityGraphNodeDefinitionOverrideDisplay.md) | The node *may* have display properties that override the active Activity&#x27;s display properties. | [optional] [default to null]
**Position** | [***AllOfDestinyDefinitionsDirectorDestinyActivityGraphNodeDefinitionPosition**](AllOfDestinyDefinitionsDirectorDestinyActivityGraphNodeDefinitionPosition.md) | The position on the map for this node. | [optional] [default to null]
**FeaturingStates** | [**[]DestinyDefinitionsDirectorDestinyActivityGraphNodeFeaturingStateDefinition**](Destiny.Definitions.Director.DestinyActivityGraphNodeFeaturingStateDefinition.md) | The node may have various visual accents placed on it, or styles applied. These are the list of possible styles that the Node can have. The game iterates through each, looking for the first one that passes a check of the required game/character/account state in order to show that style, and then renders the node in that style. | [optional] [default to null]
**Activities** | [**[]DestinyDefinitionsDirectorDestinyActivityGraphNodeActivityDefinition**](Destiny.Definitions.Director.DestinyActivityGraphNodeActivityDefinition.md) | The node may have various possible activities that could be active for it, however only one may be active at a time. See the DestinyActivityGraphNodeActivityDefinition for details. | [optional] [default to null]
**States** | [**[]DestinyDefinitionsDirectorDestinyActivityGraphNodeStateEntry**](Destiny.Definitions.Director.DestinyActivityGraphNodeStateEntry.md) | Represents possible states that the graph node can be in. These are combined with some checking that happens in the game client and server to determine which state is actually active at any given time. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

