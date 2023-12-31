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

type DestinyDefinitionsDestinyPlugItemCraftingRequirements struct {
	UnlockRequirements []DestinyDefinitionsDestinyPlugItemCraftingUnlockRequirement `json:"unlockRequirements,omitempty"`
	// If the plug has a known level requirement, it'll be available here.
	RequiredLevel int32 `json:"requiredLevel,omitempty"`
	MaterialRequirementHashes []int32 `json:"materialRequirementHashes,omitempty"`
}
