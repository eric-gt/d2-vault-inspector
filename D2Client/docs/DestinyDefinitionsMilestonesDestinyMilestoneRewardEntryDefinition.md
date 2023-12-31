# DestinyDefinitionsMilestonesDestinyMilestoneRewardEntryDefinition

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RewardEntryHash** | **int32** | The identifier for this reward entry. Runtime data will refer to reward entries by this hash. Only guaranteed unique within the specific Milestone. | [optional] [default to null]
**RewardEntryIdentifier** | **string** | The string identifier, if you care about it. Only guaranteed unique within the specific Milestone. | [optional] [default to null]
**Items** | [**[]DestinyDestinyItemQuantity**](Destiny.DestinyItemQuantity.md) | The items you will get as rewards, and how much of it you&#x27;ll get. | [optional] [default to null]
**VendorHash** | **int32** | If this reward is redeemed at a Vendor, this is the hash of the Vendor to go to in order to redeem the reward. Use this hash to look up the DestinyVendorDefinition. | [optional] [default to null]
**DisplayProperties** | [***AllOfDestinyDefinitionsMilestonesDestinyMilestoneRewardEntryDefinitionDisplayProperties**](AllOfDestinyDefinitionsMilestonesDestinyMilestoneRewardEntryDefinitionDisplayProperties.md) | For us to bother returning this info, we should be able to return some kind of information about why these rewards are grouped together. This is ideally that information. Look at how confident I am that this will always remain true. | [optional] [default to null]
**Order** | **int32** | If you want to follow BNet&#x27;s ordering of these rewards, use this number within a given category to order the rewards. Yeah, I know. I feel dirty too. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

