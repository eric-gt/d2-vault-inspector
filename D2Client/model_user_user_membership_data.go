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

type UserUserMembershipData struct {
	// this allows you to see destiny memberships that are visible and linked to this account (regardless of whether or not they have characters on the world server)
	DestinyMemberships []GroupsV2GroupUserInfoCard `json:"destinyMemberships,omitempty"`
	// If this property is populated, it will have the membership ID of the account considered to be \"primary\" in this user's cross save relationship.   If null, this user has no cross save relationship, nor primary account.
	PrimaryMembershipId int64 `json:"primaryMembershipId,omitempty"`
	BungieNetUser *UserGeneralUser `json:"bungieNetUser,omitempty"`
}
