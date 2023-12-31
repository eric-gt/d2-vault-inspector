# DestinyEntitiesProfilesDestinyProfileComponent

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserInfo** | [***AllOfDestinyEntitiesProfilesDestinyProfileComponentUserInfo**](AllOfDestinyEntitiesProfilesDestinyProfileComponentUserInfo.md) | If you need to render the Profile (their platform name, icon, etc...) somewhere, this property contains that information. | [optional] [default to null]
**DateLastPlayed** | [**time.Time**](time.Time.md) | The last time the user played with any character on this Profile. | [optional] [default to null]
**VersionsOwned** | **int32** | If you want to know what expansions they own, this will contain that data.   IMPORTANT: This field may not return the data you&#x27;re interested in for Cross-Saved users. It returns the last ownership data we saw for this account - which is to say, what they&#x27;ve purchased on the platform on which they last played, which now could be a different platform.   If you don&#x27;t care about per-platform ownership and only care about whatever platform it seems they are playing on most recently, then this should be \&quot;good enough.\&quot; Otherwise, this should be considered deprecated. We do not have a good alternative to provide at this time with platform specific ownership data for DLC. | [optional] [default to null]
**CharacterIds** | **[]int64** | A list of the character IDs, for further querying on your part. | [optional] [default to null]
**SeasonHashes** | **[]int32** | A list of seasons that this profile owns. Unlike versionsOwned, these stay with the profile across Platforms, and thus will be valid.   It turns out that Stadia Pro subscriptions will give access to seasons but only while playing on Stadia and with an active subscription. So some users (users who have Stadia Pro but choose to play on some other platform) won&#x27;t see these as available: it will be whatever seasons are available for the platform on which they last played. | [optional] [default to null]
**EventCardHashesOwned** | **[]int32** | A list of hashes for event cards that a profile owns. Unlike most values in versionsOwned, these stay with the profile across all platforms. | [optional] [default to null]
**CurrentSeasonHash** | **int32** | If populated, this is a reference to the season that is currently active. | [optional] [default to null]
**CurrentSeasonRewardPowerCap** | **int32** | If populated, this is the reward power cap for the current season. | [optional] [default to null]
**ActiveEventCardHash** | **int32** | If populated, this is a reference to the event card that is currently active. | [optional] [default to null]
**CurrentGuardianRank** | **int32** | The &#x27;current&#x27; Guardian Rank value, which starts at rank 1. | [optional] [default to null]
**LifetimeHighestGuardianRank** | **int32** | The &#x27;lifetime highest&#x27; Guardian Rank value, which starts at rank 1. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

