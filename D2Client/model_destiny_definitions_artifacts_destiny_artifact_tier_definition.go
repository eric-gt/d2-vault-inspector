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

type DestinyDefinitionsArtifactsDestinyArtifactTierDefinition struct {
	// An identifier, unique within the Artifact, for this specific tier.
	TierHash int32 `json:"tierHash,omitempty"`
	// The human readable title of this tier, if any.
	DisplayTitle string `json:"displayTitle,omitempty"`
	// A string representing the localized minimum requirement text for this Tier, if any.
	ProgressRequirementMessage string `json:"progressRequirementMessage,omitempty"`
	// The items that can be earned within this tier.
	Items []DestinyDefinitionsArtifactsDestinyArtifactTierItemDefinition `json:"items,omitempty"`
	// The minimum number of \"unlock points\" that you must have used before you can unlock items from this tier.
	MinimumUnlockPointsUsedRequirement int32 `json:"minimumUnlockPointsUsedRequirement,omitempty"`
}
