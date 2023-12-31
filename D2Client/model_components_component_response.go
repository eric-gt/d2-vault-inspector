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

// The base class for any component-returning object that may need to indicate information about the state of the component being returned.
type ComponentsComponentResponse struct {
	Privacy int32 `json:"privacy,omitempty"`
	// If true, this component is disabled.
	Disabled bool `json:"disabled,omitempty"`
}
