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

// A destiny event will not necessarily have milestone \"custom content\", but if it does the details will be here.
type AllOfTrendingTrendingEntryDestinyRitualEventContent struct {
	// The \"About this Milestone\" text from the Firehose.
	About string `json:"about,omitempty"`
	// The Current Status of the Milestone, as driven by the Firehose.
	Status string `json:"status,omitempty"`
	// A list of tips, provided by the Firehose.
	Tips []string `json:"tips,omitempty"`
	// If DPS has defined items related to this Milestone, they can categorize those items in the Firehose. That data will then be returned as item categories here.
	ItemCategories []DestinyMilestonesDestinyMilestoneContentItemCategory `json:"itemCategories,omitempty"`
}
