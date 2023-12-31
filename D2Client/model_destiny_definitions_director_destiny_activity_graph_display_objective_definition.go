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

// When a Graph needs to show active Objectives, this defines those objectives as well as an identifier.
type DestinyDefinitionsDirectorDestinyActivityGraphDisplayObjectiveDefinition struct {
	// $NOTE $amola 2017-01-19 This field is apparently something that CUI uses to manually wire up objectives to display info. I am unsure how it works.
	Id int32 `json:"id,omitempty"`
	// The objective being shown on the map.
	ObjectiveHash int32 `json:"objectiveHash,omitempty"`
}
