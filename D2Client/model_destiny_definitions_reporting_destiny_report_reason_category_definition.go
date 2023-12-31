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

// If you're going to report someone for a Terms of Service violation, you need to choose a category and reason for the report. This definition holds both the categories and the reasons within those categories, for simplicity and my own laziness' sake.  Note tha this means that, to refer to a Reason by reasonHash, you need a combination of the reasonHash *and* the associated ReasonCategory's hash: there are some reasons defined under multiple categories.
type DestinyDefinitionsReportingDestinyReportReasonCategoryDefinition struct {
	DisplayProperties *DestinyDefinitionsCommonDestinyDisplayPropertiesDefinition `json:"displayProperties,omitempty"`
	// The specific reasons for the report under this category.
	Reasons map[string]DestinyDefinitionsReportingDestinyReportReasonDefinition `json:"reasons,omitempty"`
	// The unique identifier for this entity. Guaranteed to be unique for the type of entity, but not globally.  When entities refer to each other in Destiny content, it is this hash that they are referring to.
	Hash int32 `json:"hash,omitempty"`
	// The index of the entity as it was found in the investment tables.
	Index int32 `json:"index,omitempty"`
	// If this is true, then there is an entity with this identifier/type combination, but BNet is not yet allowed to show it. Sorry!
	Redacted bool `json:"redacted,omitempty"`
}
