# DestinyDefinitionsChecklistsDestinyChecklistEntryDefinition

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hash** | **int32** | The identifier for this Checklist entry. Guaranteed unique only within this Checklist Definition, and not globally/for all checklists. | [optional] [default to null]
**DisplayProperties** | [***AllOfDestinyDefinitionsChecklistsDestinyChecklistEntryDefinitionDisplayProperties**](AllOfDestinyDefinitionsChecklistsDestinyChecklistEntryDefinitionDisplayProperties.md) | Even if no other associations exist, we will give you *something* for display properties. In cases where we have no associated entities, it may be as simple as a numerical identifier. | [optional] [default to null]
**DestinationHash** | **int32** |  | [optional] [default to null]
**LocationHash** | **int32** |  | [optional] [default to null]
**BubbleHash** | **int32** | Note that a Bubble&#x27;s hash doesn&#x27;t uniquely identify a \&quot;top level\&quot; entity in Destiny. Only the combination of location and bubble can uniquely identify a place in the world of Destiny: so if bubbleHash is populated, locationHash must too be populated for it to have any meaning.  You can use this property if it is populated to look up the DestinyLocationDefinition&#x27;s associated .locationReleases[].activityBubbleName property. | [optional] [default to null]
**ActivityHash** | **int32** |  | [optional] [default to null]
**ItemHash** | **int32** |  | [optional] [default to null]
**VendorHash** | **int32** |  | [optional] [default to null]
**VendorInteractionIndex** | **int32** |  | [optional] [default to null]
**Scope** | **int32** | The scope at which this specific entry can be computed. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

