# AllOfTrendingTrendingEntryDestinyRitualMilestoneDetails

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MilestoneHash** | **int32** | The hash identifier for the milestone. Use it to look up the DestinyMilestoneDefinition for static data about the Milestone. | [optional] [default to null]
**AvailableQuests** | [**[]DestinyMilestonesDestinyPublicMilestoneQuest**](Destiny.Milestones.DestinyPublicMilestoneQuest.md) | A milestone not need have even a single quest, but if there are active quests they will be returned here. | [optional] [default to null]
**Activities** | [**[]DestinyMilestonesDestinyPublicMilestoneChallengeActivity**](Destiny.Milestones.DestinyPublicMilestoneChallengeActivity.md) |  | [optional] [default to null]
**VendorHashes** | **[]int32** | Sometimes milestones - or activities active in milestones - will have relevant vendors. These are the vendors that are currently relevant.  Deprecated, already, for the sake of the new \&quot;vendors\&quot; property that has more data. What was I thinking. | [optional] [default to null]
**Vendors** | [**[]DestinyMilestonesDestinyPublicMilestoneVendor**](Destiny.Milestones.DestinyPublicMilestoneVendor.md) | This is why we can&#x27;t have nice things. This is the ordered list of vendors to be shown that relate to this milestone, potentially along with other interesting data. | [optional] [default to null]
**StartDate** | [**time.Time**](time.Time.md) | If known, this is the date when the Milestone started/became active. | [optional] [default to null]
**EndDate** | [**time.Time**](time.Time.md) | If known, this is the date when the Milestone will expire/recycle/end. | [optional] [default to null]
**Order** | **int32** | Used for ordering milestones in a display to match how we order them in BNet. May pull from static data, or possibly in the future from dynamic information. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

