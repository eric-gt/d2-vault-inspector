# DestinyMilestonesDestinyPublicMilestoneChallengeActivity

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivityHash** | **int32** |  | [optional] [default to null]
**ChallengeObjectiveHashes** | **[]int32** |  | [optional] [default to null]
**ModifierHashes** | **[]int32** | If the activity has modifiers, this will be the list of modifiers that all variants have in common. Perform lookups against DestinyActivityModifierDefinition which defines the modifier being applied to get at the modifier data.  Note that, in the DestiyActivityDefinition, you will see many more modifiers than this being referred to: those are all *possible* modifiers for the activity, not the active ones. Use only the active ones to match what&#x27;s really live. | [optional] [default to null]
**LoadoutRequirementIndex** | **int32** | If returned, this is the index into the DestinyActivityDefinition&#x27;s \&quot;loadouts\&quot; property, indicating the currently active loadout requirements. | [optional] [default to null]
**PhaseHashes** | **[]int32** | The ordered list of phases for this activity, if any. Note that we have no human readable info for phases, nor any entities to relate them to: relating these hashes to something human readable is up to you unfortunately. | [optional] [default to null]
**BooleanActivityOptions** | **map[string]bool** | The set of activity options for this activity, keyed by an identifier that&#x27;s unique for this activity (not guaranteed to be unique between or across all activities, though should be unique for every *variant* of a given *conceptual* activity: for instance, the original D2 Raid has many variant DestinyActivityDefinitions. While other activities could potentially have the same option hashes, for any given D2 base Raid variant the hash will be unique).  As a concrete example of this data, the hashes you get for Raids will correspond to the currently active \&quot;Challenge Mode\&quot;.  We have no human readable information for this data, so it&#x27;s up to you if you want to associate it with such info to show it. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

