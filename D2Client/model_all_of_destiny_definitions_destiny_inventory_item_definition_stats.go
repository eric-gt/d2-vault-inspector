/*
 * Bungie.Net API
 *
 * These endpoints constitute the functionality exposed by Bungie.net, both for more traditional website functionality and for connectivity to Bungie video games and their related functionality.
 *
 * API version: 2.18.0
 * Contact: support@bungie.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package D2Client

// If this item can have stats (such as a weapon, armor, or vehicle), this block will be non-null and populated with the stats found on the item.
type AllOfDestinyDefinitionsDestinyInventoryItemDefinitionStats struct {
	// If true, the game won't show the \"primary\" stat on this item when you inspect it.  NOTE: This is being manually mapped, because I happen to want it in a block that isn't going to directly create this derivative block.
	DisablePrimaryStatDisplay bool `json:"disablePrimaryStatDisplay,omitempty"`
	// If the item's stats are meant to be modified by a DestinyStatGroupDefinition, this will be the identifier for that definition.  If you are using live data or precomputed stats data on the DestinyInventoryItemDefinition.stats.stats property, you don't have to worry about statGroupHash and how it alters stats: the already altered stats are provided to you. But if you want to see how the sausage gets made, or perform computations yourself, this is valuable information.
	StatGroupHash int32 `json:"statGroupHash,omitempty"`
	// If you are looking for precomputed values for the stats on a weapon, this is where they are stored. Technically these are the \"Display\" stat values. Please see DestinyStatsDefinition for what Display Stat Values means, it's a very long story... but essentially these are the closest values BNet can get to the item stats that you see in-game.  These stats are keyed by the DestinyStatDefinition's hash identifier for the stat that's found on the item.
	Stats map[string]DestinyDefinitionsDestinyInventoryItemStatDefinition `json:"stats,omitempty"`
	// A quick and lazy way to determine whether any stat other than the \"primary\" stat is actually visible on the item. Items often have stats that we return in case people find them useful, but they're not part of the \"Stat Group\" and thus we wouldn't display them in our UI. If this is False, then we're not going to display any of these stats other than the primary one.
	HasDisplayableStats bool `json:"hasDisplayableStats,omitempty"`
	// This stat is determined to be the \"primary\" stat, and can be looked up in the stats or any other stat collection related to the item.  Use this hash to look up the stat's value using DestinyInventoryItemDefinition.stats.stats, and the renderable data for the primary stat in the related DestinyStatDefinition.
	PrimaryBaseStatHash int32 `json:"primaryBaseStatHash,omitempty"`
}
