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

type DestinySocketsDestinyItemPlugBase struct {
	// The hash identifier of the DestinyInventoryItemDefinition that represents this plug.
	PlugItemHash int32 `json:"plugItemHash,omitempty"`
	// If true, this plug has met all of its insertion requirements. Big if true.
	CanInsert bool `json:"canInsert,omitempty"`
	// If true, this plug will provide its benefits while inserted.
	Enabled bool `json:"enabled,omitempty"`
	// If the plug cannot be inserted for some reason, this will have the indexes into the plug item definition's plug.insertionRules property, so you can show the reasons why it can't be inserted.  This list will be empty if the plug can be inserted.
	InsertFailIndexes []int32 `json:"insertFailIndexes,omitempty"`
	// If a plug is not enabled, this will be populated with indexes into the plug item definition's plug.enabledRules property, so that you can show the reasons why it is not enabled.  This list will be empty if the plug is enabled.
	EnableFailIndexes []int32 `json:"enableFailIndexes,omitempty"`
}
