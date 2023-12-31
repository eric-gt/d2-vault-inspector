# AllOfDestinyMilestonesDestinyMilestoneQuestStatus

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QuestHash** | **int32** | The hash identifier for the Quest Item. (Note: Quests are defined as Items, and thus you would use this to look up the quest&#x27;s DestinyInventoryItemDefinition). For information on all steps in the quest, you can then examine its DestinyInventoryItemDefinition.setData property for Quest Steps (which are *also* items). You can use the Item Definition to display human readable data about the overall quest. | [optional] [default to null]
**StepHash** | **int32** | The hash identifier of the current Quest Step, which is also a DestinyInventoryItemDefinition. You can use this to get human readable data about the current step and what to do in that step. | [optional] [default to null]
**StepObjectives** | [**[]DestinyQuestsDestinyObjectiveProgress**](Destiny.Quests.DestinyObjectiveProgress.md) | A step can have multiple objectives. This will give you the progress for each objective in the current step, in the order in which they are rendered in-game. | [optional] [default to null]
**Tracked** | **bool** | Whether or not the quest is tracked | [optional] [default to null]
**ItemInstanceId** | **int64** | The current Quest Step will be an instanced item in the player&#x27;s inventory. If you care about that, this is the instance ID of that item. | [optional] [default to null]
**Completed** | **bool** | Whether or not the whole quest has been completed, regardless of whether or not you have redeemed the rewards for the quest. | [optional] [default to null]
**Redeemed** | **bool** | Whether or not you have redeemed rewards for this quest. | [optional] [default to null]
**Started** | **bool** | Whether or not you have started this quest. | [optional] [default to null]
**VendorHash** | **int32** | If the quest has a related Vendor that you should talk to in order to initiate the quest/earn rewards/continue the quest, this will be the hash identifier of that Vendor. Look it up its DestinyVendorDefinition. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

