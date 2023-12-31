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

// The stat being modified, and the value used.
type AllOfDestinyDefinitionsDestinyObjectiveStatEntryDefinitionStat struct {
	// The hash identifier for the DestinyStatDefinition defining this stat.
	StatTypeHash int32 `json:"statTypeHash,omitempty"`
	// The raw \"Investment\" value for the stat, before transformations are performed to turn this raw stat into stats that are displayed in the game UI.
	Value int32 `json:"value,omitempty"`
	// If this is true, the stat will only be applied on the item in certain game state conditions, and we can't know statically whether or not this stat will be applied. Check the \"live\" API data instead for whether this value is being applied on a specific instance of the item in question, and you can use this to decide whether you want to show the stat on the generic view of the item, or whether you want to show some kind of caveat or warning about the stat value being conditional on game state.
	IsConditionallyActive bool `json:"isConditionallyActive,omitempty"`
}
