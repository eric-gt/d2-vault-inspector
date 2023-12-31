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

// Summary data about the item.
type AllOfDestinyDefinitionsDestinyInventoryItemDefinitionSummary struct {
	// Apparently when rendering an item in a reward, this should be used as a sort priority. We're not doing it presently.
	SortPriority int32 `json:"sortPriority,omitempty"`
}
