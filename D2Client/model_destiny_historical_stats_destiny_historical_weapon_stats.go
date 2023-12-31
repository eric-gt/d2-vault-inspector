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

type DestinyHistoricalStatsDestinyHistoricalWeaponStats struct {
	// The hash ID of the item definition that describes the weapon.
	ReferenceId int32 `json:"referenceId,omitempty"`
	// Collection of stats for the period.
	Values map[string]DestinyHistoricalStatsDestinyHistoricalStatsValue `json:"values,omitempty"`
}
