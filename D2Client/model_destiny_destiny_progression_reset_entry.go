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

// Represents a season and the number of resets you had in that season.   We do not necessarily - even for progressions with resets - track it over all seasons. So be careful and check the season numbers being returned.
type DestinyDestinyProgressionResetEntry struct {
	Season int32 `json:"season,omitempty"`
	Resets int32 `json:"resets,omitempty"`
}
