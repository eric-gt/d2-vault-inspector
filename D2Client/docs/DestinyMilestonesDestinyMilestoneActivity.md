# DestinyMilestonesDestinyMilestoneActivity

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivityHash** | **int32** | The hash of an arbitrarily chosen variant of this activity. We&#x27;ll go ahead and call that the \&quot;canonical\&quot; activity, because if you&#x27;re using this value you should only use it for properties that are common across the variants: things like the name of the activity, it&#x27;s location, etc... Use this hash to look up the DestinyActivityDefinition of this activity for rendering data. | [optional] [default to null]
**ActivityModeHash** | **int32** | The hash identifier of the most specific Activity Mode under which this activity is played. This is useful for situations where the activity in question is - for instance - a PVP map, but it&#x27;s not clear what mode the PVP map is being played under. If it&#x27;s a playlist, this will be less specific: but hopefully useful in some way. | [optional] [default to null]
**ActivityModeType** | **int32** | The enumeration equivalent of the most specific Activity Mode under which this activity is played. | [optional] [default to null]
**ModifierHashes** | **[]int32** | If the activity has modifiers, this will be the list of modifiers that all variants have in common. Perform lookups against DestinyActivityModifierDefinition which defines the modifier being applied to get at the modifier data. Note that, in the DestiyActivityDefinition, you will see many more modifiers than this being referred to: those are all *possible* modifiers for the activity, not the active ones. Use only the active ones to match what&#x27;s really live. | [optional] [default to null]
**Variants** | [**[]DestinyMilestonesDestinyMilestoneActivityVariant**](Destiny.Milestones.DestinyMilestoneActivityVariant.md) | If you want more than just name/location/etc... you&#x27;re going to have to dig into and show the variants of the conceptual activity. These will differ in seemingly arbitrary ways, like difficulty level and modifiers applied. Show it in whatever way tickles your fancy. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

