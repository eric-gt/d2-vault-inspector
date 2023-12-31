# DestinyDefinitionsMilestonesDestinyMilestoneRewardCategoryDefinition

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CategoryHash** | **int32** | Identifies the reward category. Only guaranteed unique within this specific component! | [optional] [default to null]
**CategoryIdentifier** | **string** | The string identifier for the category, if you want to use it for some end. Guaranteed unique within the specific component. | [optional] [default to null]
**DisplayProperties** | [***AllOfDestinyDefinitionsMilestonesDestinyMilestoneRewardCategoryDefinitionDisplayProperties**](AllOfDestinyDefinitionsMilestonesDestinyMilestoneRewardCategoryDefinitionDisplayProperties.md) | Hopefully this is obvious by now. | [optional] [default to null]
**RewardEntries** | [**map[string]DestinyDefinitionsMilestonesDestinyMilestoneRewardEntryDefinition**](Destiny.Definitions.Milestones.DestinyMilestoneRewardEntryDefinition.md) | If this milestone can provide rewards, this will define the sets of rewards that can be earned, the conditions under which they can be acquired, internal data that we&#x27;ll use at runtime to determine whether you&#x27;ve already earned or redeemed this set of rewards, and the category that this reward should be placed under. | [optional] [default to null]
**Order** | **int32** | If you want to use BNet&#x27;s recommended order for rendering categories programmatically, use this value and compare it to other categories to determine the order in which they should be rendered. I don&#x27;t feel great about putting this here, I won&#x27;t lie. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

