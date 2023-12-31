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

// If a Destiny Profile can't be returned, but we're pretty certain it's a valid Destiny account, this will contain as much info as we can get about the profile for your use.  Assume that the most you'll get is the Error Code, the Membership Type and the Membership ID.
type DestinyResponsesDestinyErrorProfile struct {
	// The error that we encountered. You should be able to look up localized text to show to the user for these failures.
	ErrorCode int32 `json:"errorCode,omitempty"`
	// Basic info about the account that failed. Don't expect anything other than membership ID, Membership Type, and displayName to be populated.
	InfoCard *AllOfDestinyResponsesDestinyErrorProfileInfoCard `json:"infoCard,omitempty"`
}
