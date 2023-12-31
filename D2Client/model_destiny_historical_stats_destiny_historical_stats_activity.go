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

// Summary information about the activity that was played.
type DestinyHistoricalStatsDestinyHistoricalStatsActivity struct {
	// The unique hash identifier of the DestinyActivityDefinition that was played. If I had this to do over, it'd be named activityHash. Too late now.
	ReferenceId int32 `json:"referenceId,omitempty"`
	// The unique hash identifier of the DestinyActivityDefinition that was played.
	DirectorActivityHash int32 `json:"directorActivityHash,omitempty"`
	// The unique identifier for this *specific* match that was played.  This value can be used to get additional data about this activity such as who else was playing via the GetPostGameCarnageReport endpoint.
	InstanceId int64 `json:"instanceId,omitempty"`
	// Indicates the most specific game mode of the activity that we could find.
	Mode int32 `json:"mode,omitempty"`
	// The list of all Activity Modes to which this activity applies, including aggregates. This will let you see, for example, whether the activity was both Clash and part of the Trials of the Nine event.
	Modes []int32 `json:"modes,omitempty"`
	// Whether or not the match was a private match.
	IsPrivate bool `json:"isPrivate,omitempty"`
	// The Membership Type indicating the platform on which this match was played.
	MembershipType int32 `json:"membershipType,omitempty"`
}
