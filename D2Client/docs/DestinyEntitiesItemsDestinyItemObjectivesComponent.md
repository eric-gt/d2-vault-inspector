# DestinyEntitiesItemsDestinyItemObjectivesComponent

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Objectives** | [**[]DestinyQuestsDestinyObjectiveProgress**](Destiny.Quests.DestinyObjectiveProgress.md) | If the item has a hard association with objectives, your progress on them will be defined here.   Objectives are our standard way to describe a series of tasks that have to be completed for a reward. | [optional] [default to null]
**FlavorObjective** | [***AllOfDestinyEntitiesItemsDestinyItemObjectivesComponentFlavorObjective**](AllOfDestinyEntitiesItemsDestinyItemObjectivesComponentFlavorObjective.md) | I may regret naming it this way - but this represents when an item has an objective that doesn&#x27;t serve a beneficial purpose, but rather is used for \&quot;flavor\&quot; or additional information. For instance, when Emblems track specific stats, those stats are represented as Objectives on the item. | [optional] [default to null]
**DateCompleted** | [**time.Time**](time.Time.md) | If we have any information on when these objectives were completed, this will be the date of that completion. This won&#x27;t be on many items, but could be interesting for some items that do store this information. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

