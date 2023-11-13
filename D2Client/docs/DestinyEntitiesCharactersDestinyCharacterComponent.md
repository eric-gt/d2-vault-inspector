# DestinyEntitiesCharactersDestinyCharacterComponent

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MembershipId** | **int64** | Every Destiny Profile has a membershipId. This is provided on the character as well for convenience. | [optional] [default to null]
**MembershipType** | **int32** | membershipType tells you the platform on which the character plays. Examine the BungieMembershipType enumeration for possible values. | [optional] [default to null]
**CharacterId** | **int64** | The unique identifier for the character. | [optional] [default to null]
**DateLastPlayed** | [**time.Time**](time.Time.md) | The last date that the user played Destiny. | [optional] [default to null]
**MinutesPlayedThisSession** | **int64** | If the user is currently playing, this is how long they&#x27;ve been playing. | [optional] [default to null]
**MinutesPlayedTotal** | **int64** | If this value is 525,600, then they played Destiny for a year. Or they&#x27;re a very dedicated Rent fan. Note that this includes idle time, not just time spent actually in activities shooting things. | [optional] [default to null]
**Light** | **int32** | The user&#x27;s calculated \&quot;Light Level\&quot;. Light level is an indicator of your power that mostly matters in the end game, once you&#x27;ve reached the maximum character level: it&#x27;s a level that&#x27;s dependent on the average Attack/Defense power of your items. | [optional] [default to null]
**Stats** | **map[string]int32** | Your character&#x27;s stats, such as Agility, Resilience, etc... *not* historical stats.  You&#x27;ll have to call a different endpoint for those. | [optional] [default to null]
**RaceHash** | **int32** | Use this hash to look up the character&#x27;s DestinyRaceDefinition. | [optional] [default to null]
**GenderHash** | **int32** | Use this hash to look up the character&#x27;s DestinyGenderDefinition. | [optional] [default to null]
**ClassHash** | **int32** | Use this hash to look up the character&#x27;s DestinyClassDefinition. | [optional] [default to null]
**RaceType** | **int32** | Mostly for historical purposes at this point, this is an enumeration for the character&#x27;s race.  It&#x27;ll be preferable in the general case to look up the related definition: but for some people this was too convenient to remove. | [optional] [default to null]
**ClassType** | **int32** | Mostly for historical purposes at this point, this is an enumeration for the character&#x27;s class.  It&#x27;ll be preferable in the general case to look up the related definition: but for some people this was too convenient to remove. | [optional] [default to null]
**GenderType** | **int32** | Mostly for historical purposes at this point, this is an enumeration for the character&#x27;s Gender.  It&#x27;ll be preferable in the general case to look up the related definition: but for some people this was too convenient to remove. And yeah, it&#x27;s an enumeration and not a boolean. Fight me. | [optional] [default to null]
**EmblemPath** | **string** | A shortcut path to the user&#x27;s currently equipped emblem image. If you&#x27;re just showing summary info for a user, this is more convenient than examining their equipped emblem and looking up the definition. | [optional] [default to null]
**EmblemBackgroundPath** | **string** | A shortcut path to the user&#x27;s currently equipped emblem background image. If you&#x27;re just showing summary info for a user, this is more convenient than examining their equipped emblem and looking up the definition. | [optional] [default to null]
**EmblemHash** | **int32** | The hash of the currently equipped emblem for the user. Can be used to look up the DestinyInventoryItemDefinition. | [optional] [default to null]
**EmblemColor** | [***AllOfDestinyEntitiesCharactersDestinyCharacterComponentEmblemColor**](AllOfDestinyEntitiesCharactersDestinyCharacterComponentEmblemColor.md) | A shortcut for getting the background color of the user&#x27;s currently equipped emblem without having to do a DestinyInventoryItemDefinition lookup. | [optional] [default to null]
**LevelProgression** | [***AllOfDestinyEntitiesCharactersDestinyCharacterComponentLevelProgression**](AllOfDestinyEntitiesCharactersDestinyCharacterComponentLevelProgression.md) | The progression that indicates your character&#x27;s level. Not their light level, but their character level: you know, the thing you max out a couple hours in and then ignore for the sake of light level. | [optional] [default to null]
**BaseCharacterLevel** | **int32** | The \&quot;base\&quot; level of your character, not accounting for any light level. | [optional] [default to null]
**PercentToNextLevel** | **float32** | A number between 0 and 100, indicating the whole and fractional % remaining to get to the next character level. | [optional] [default to null]
**TitleRecordHash** | **int32** | If this Character has a title assigned to it, this is the identifier of the DestinyRecordDefinition that has that title information. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

