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
// DestinyDestinyScope : There's a lot of places where we need to know scope on more than just a profile or character level. For everything else, there's this more generic sense of scope.
type DestinyDestinyScope int32

// List of Destiny.DestinyScope
const (
	_0__DestinyDestinyScope DestinyDestinyScope = 0
	_1__DestinyDestinyScope DestinyDestinyScope = 1
)
