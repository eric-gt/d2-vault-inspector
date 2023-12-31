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
// DestinyActivityGraphNodeHighlightType : The various known UI styles in which an item can be highlighted. It'll be up to you to determine what you want to show based on this highlighting, BNet doesn't have any assets that correspond to these states. And yeah, RiseOfIron and Comet have their own special highlight states. Don't ask me, I can't imagine they're still used.
type DestinyActivityGraphNodeHighlightType int32

// List of Destiny.ActivityGraphNodeHighlightType
const (
	_0__DestinyActivityGraphNodeHighlightType DestinyActivityGraphNodeHighlightType = 0
	_1__DestinyActivityGraphNodeHighlightType DestinyActivityGraphNodeHighlightType = 1
	_2__DestinyActivityGraphNodeHighlightType DestinyActivityGraphNodeHighlightType = 2
	_3__DestinyActivityGraphNodeHighlightType DestinyActivityGraphNodeHighlightType = 3
	_4__DestinyActivityGraphNodeHighlightType DestinyActivityGraphNodeHighlightType = 4
)
