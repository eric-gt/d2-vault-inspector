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

type GroupsV2ClanBanner struct {
	DecalId int32 `json:"decalId,omitempty"`
	DecalColorId int32 `json:"decalColorId,omitempty"`
	DecalBackgroundColorId int32 `json:"decalBackgroundColorId,omitempty"`
	GonfalonId int32 `json:"gonfalonId,omitempty"`
	GonfalonColorId int32 `json:"gonfalonColorId,omitempty"`
	GonfalonDetailId int32 `json:"gonfalonDetailId,omitempty"`
	GonfalonDetailColorId int32 `json:"gonfalonDetailColorId,omitempty"`
}
