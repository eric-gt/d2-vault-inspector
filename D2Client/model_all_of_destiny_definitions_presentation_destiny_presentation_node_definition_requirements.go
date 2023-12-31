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

// The requirements for being able to interact with this presentation node and its children.
type AllOfDestinyDefinitionsPresentationDestinyPresentationNodeDefinitionRequirements struct {
	// If this node is not accessible due to Entitlements (for instance, you don't own the required game expansion), this is the message to show.
	EntitlementUnavailableMessage string `json:"entitlementUnavailableMessage,omitempty"`
}
