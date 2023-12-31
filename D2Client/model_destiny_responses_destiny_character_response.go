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

// The response contract for GetDestinyCharacter, with components that can be returned for character and item-level data.
type DestinyResponsesDestinyCharacterResponse struct {
	// The character-level non-equipped inventory items.  COMPONENT TYPE: CharacterInventories
	Inventory *AllOfDestinyResponsesDestinyCharacterResponseInventory `json:"inventory,omitempty"`
	// Base information about the character in question.  COMPONENT TYPE: Characters
	Character *AllOfDestinyResponsesDestinyCharacterResponseCharacter `json:"character,omitempty"`
	// Character progression data, including Milestones.  COMPONENT TYPE: CharacterProgressions
	Progressions *AllOfDestinyResponsesDestinyCharacterResponseProgressions `json:"progressions,omitempty"`
	// Character rendering data - a minimal set of information about equipment and dyes used for rendering.  COMPONENT TYPE: CharacterRenderData
	RenderData *AllOfDestinyResponsesDestinyCharacterResponseRenderData `json:"renderData,omitempty"`
	// Activity data - info about current activities available to the player.  COMPONENT TYPE: CharacterActivities
	Activities *AllOfDestinyResponsesDestinyCharacterResponseActivities `json:"activities,omitempty"`
	// Equipped items on the character.  COMPONENT TYPE: CharacterEquipment
	Equipment *AllOfDestinyResponsesDestinyCharacterResponseEquipment `json:"equipment,omitempty"`
	// The loadouts available to the character.  COMPONENT TYPE: CharacterLoadouts
	Loadouts *AllOfDestinyResponsesDestinyCharacterResponseLoadouts `json:"loadouts,omitempty"`
	// Items available from Kiosks that are available to this specific character.   COMPONENT TYPE: Kiosks
	Kiosks *AllOfDestinyResponsesDestinyCharacterResponseKiosks `json:"kiosks,omitempty"`
	// When sockets refer to reusable Plug Sets (see DestinyPlugSetDefinition for more info), this is the set of plugs and their states that are scoped to this character.  This comes back with ItemSockets, as it is needed for a complete picture of the sockets on requested items.  COMPONENT TYPE: ItemSockets
	PlugSets *AllOfDestinyResponsesDestinyCharacterResponsePlugSets `json:"plugSets,omitempty"`
	// COMPONENT TYPE: PresentationNodes
	PresentationNodes *AllOfDestinyResponsesDestinyCharacterResponsePresentationNodes `json:"presentationNodes,omitempty"`
	// COMPONENT TYPE: Records
	Records *AllOfDestinyResponsesDestinyCharacterResponseRecords `json:"records,omitempty"`
	// COMPONENT TYPE: Collectibles
	Collectibles *AllOfDestinyResponsesDestinyCharacterResponseCollectibles `json:"collectibles,omitempty"`
	// The set of components belonging to the player's instanced items.  COMPONENT TYPE: [See inside the DestinyItemComponentSet contract for component types.]
	ItemComponents *AllOfDestinyResponsesDestinyCharacterResponseItemComponents `json:"itemComponents,omitempty"`
	// The set of components belonging to the player's UNinstanced items. Because apparently now those too can have information relevant to the character's state.  COMPONENT TYPE: [See inside the DestinyItemComponentSet contract for component types.]
	UninstancedItemComponents *AllOfDestinyResponsesDestinyCharacterResponseUninstancedItemComponents `json:"uninstancedItemComponents,omitempty"`
	// A \"lookup\" convenience component that can be used to quickly check if the character has access to items that can be used for purchasing.  COMPONENT TYPE: CurrencyLookups
	CurrencyLookups *AllOfDestinyResponsesDestinyCharacterResponseCurrencyLookups `json:"currencyLookups,omitempty"`
}
