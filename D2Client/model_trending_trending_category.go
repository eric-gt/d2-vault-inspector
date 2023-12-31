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

type TrendingTrendingCategory struct {
	CategoryName string `json:"categoryName,omitempty"`
	Entries *SearchResultOfTrendingEntry `json:"entries,omitempty"`
	CategoryId string `json:"categoryId,omitempty"`
}
