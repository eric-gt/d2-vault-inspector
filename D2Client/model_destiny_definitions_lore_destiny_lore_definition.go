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

// These are definitions for in-game \"Lore,\" meant to be narrative enhancements of the game experience.  DestinyInventoryItemDefinitions for interesting items point to these definitions, but nothing's stopping you from scraping all of these and doing something cool with them. If they end up having cool data.
type DestinyDefinitionsLoreDestinyLoreDefinition struct {
	DisplayProperties *DestinyDefinitionsCommonDestinyDisplayPropertiesDefinition `json:"displayProperties,omitempty"`
	Subtitle string `json:"subtitle,omitempty"`
	// The unique identifier for this entity. Guaranteed to be unique for the type of entity, but not globally.  When entities refer to each other in Destiny content, it is this hash that they are referring to.
	Hash int32 `json:"hash,omitempty"`
	// The index of the entity as it was found in the investment tables.
	Index int32 `json:"index,omitempty"`
	// If this is true, then there is an entity with this identifier/type combination, but BNet is not yet allowed to show it. Sorry!
	Redacted bool `json:"redacted,omitempty"`
}
