/* Bungie.Net API
 *
 * These endpoints constitute the functionality exposed by Bungie.net, both for more traditional website functionality and for connectivity to Bungie video games and their related functionality.
 *
 * API version: 2.18.0
 * Contact: support@bungie.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package D2Client

// Inventory Items can reward progression when actions are performed on them. A common example of this in Destiny 1 was Bounties, which would reward Experience on your Character and the like when you completed the bounty.  Note that this maps to a DestinyProgressionMappingDefinition, and *not* a DestinyProgressionDefinition directly. This is apparently so that multiple progressions can be granted progression points/experience at the same time.
type DestinyDefinitionsDestinyProgressionRewardDefinition struct {
	// The hash identifier of the DestinyProgressionMappingDefinition that contains the progressions for which experience should be applied.
	ProgressionMappingHash int32 `json:"progressionMappingHash,omitempty"`
	// The amount of experience to give to each of the mapped progressions.
	Amount int32 `json:"amount,omitempty"`
	// If true, the game's internal mechanisms to throttle progression should be applied.
	ApplyThrottles bool `json:"applyThrottles,omitempty"`
}
