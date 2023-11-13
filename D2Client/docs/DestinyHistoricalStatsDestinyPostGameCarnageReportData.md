# DestinyHistoricalStatsDestinyPostGameCarnageReportData

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Period** | [**time.Time**](time.Time.md) | Date and time for the activity. | [optional] [default to null]
**StartingPhaseIndex** | **int32** | If this activity has \&quot;phases\&quot;, this is the phase at which the activity was started. This value is only valid for activities before the Beyond Light expansion shipped. Subsequent activities will not have a valid value here. | [optional] [default to null]
**ActivityWasStartedFromBeginning** | **bool** | True if the activity was started from the beginning, if that information is available and the activity was played post Witch Queen release. | [optional] [default to null]
**ActivityDetails** | [***AllOfDestinyHistoricalStatsDestinyPostGameCarnageReportDataActivityDetails**](AllOfDestinyHistoricalStatsDestinyPostGameCarnageReportDataActivityDetails.md) | Details about the activity. | [optional] [default to null]
**Entries** | [**[]DestinyHistoricalStatsDestinyPostGameCarnageReportEntry**](Destiny.HistoricalStats.DestinyPostGameCarnageReportEntry.md) | Collection of players and their data for this activity. | [optional] [default to null]
**Teams** | [**[]DestinyHistoricalStatsDestinyPostGameCarnageReportTeamEntry**](Destiny.HistoricalStats.DestinyPostGameCarnageReportTeamEntry.md) | Collection of stats for the player in this activity. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

