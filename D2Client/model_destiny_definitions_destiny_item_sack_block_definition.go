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

// Some items are \"sacks\" - they can be \"opened\" to produce other items. This is information related to its sack status, mostly UI strings. Engrams are an example of items that are considered to be \"Sacks\".
type DestinyDefinitionsDestinyItemSackBlockDefinition struct {
	// A description of what will happen when you open the sack. As far as I can tell, this is blank currently. Unknown whether it will eventually be populated with useful info.
	DetailAction string `json:"detailAction,omitempty"`
	// The localized name of the action being performed when you open the sack.
	OpenAction string `json:"openAction,omitempty"`
	SelectItemCount int32 `json:"selectItemCount,omitempty"`
	VendorSackType string `json:"vendorSackType,omitempty"`
	OpenOnAcquire bool `json:"openOnAcquire,omitempty"`
}
