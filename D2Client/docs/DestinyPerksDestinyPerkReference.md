# DestinyPerksDestinyPerkReference

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PerkHash** | **int32** | The hash identifier for the perk, which can be used to look up DestinySandboxPerkDefinition if it exists. Be warned, perks frequently do not have user-viewable information. You should examine whether you actually found a name/description in the perk&#x27;s definition before you show it to the user. | [optional] [default to null]
**IconPath** | **string** | The icon for the perk. | [optional] [default to null]
**IsActive** | **bool** | Whether this perk is currently active. (We may return perks that you have not actually activated yet: these represent perks that you should show in the item&#x27;s tooltip, but that the user has not yet activated.) | [optional] [default to null]
**Visible** | **bool** | Some perks provide benefits, but aren&#x27;t visible in the UI. This value will let you know if this is perk should be shown in your UI. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

