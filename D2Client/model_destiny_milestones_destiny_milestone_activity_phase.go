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

// Represents whatever information we can return about an explicit phase in an activity. In the future, I hope we'll have more than just \"guh, you done gone and did something,\" but for the forseeable future that's all we've got. I'm making it more than just a list of booleans out of that overly-optimistic hope.
type DestinyMilestonesDestinyMilestoneActivityPhase struct {
	// Indicates if the phase has been completed.
	Complete bool `json:"complete,omitempty"`
	// In DestinyActivityDefinition, if the activity has phases, there will be a set of phases defined in the \"insertionPoints\" property. This is the hash that maps to that phase.
	PhaseHash int32 `json:"phaseHash,omitempty"`
}
