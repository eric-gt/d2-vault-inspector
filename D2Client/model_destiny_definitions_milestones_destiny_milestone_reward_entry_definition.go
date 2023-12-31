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

// The definition of a specific reward, which may be contained in a category of rewards and that has optional information about how it is obtained.
type DestinyDefinitionsMilestonesDestinyMilestoneRewardEntryDefinition struct {
	// The identifier for this reward entry. Runtime data will refer to reward entries by this hash. Only guaranteed unique within the specific Milestone.
	RewardEntryHash int32 `json:"rewardEntryHash,omitempty"`
	// The string identifier, if you care about it. Only guaranteed unique within the specific Milestone.
	RewardEntryIdentifier string `json:"rewardEntryIdentifier,omitempty"`
	// The items you will get as rewards, and how much of it you'll get.
	Items []DestinyDestinyItemQuantity `json:"items,omitempty"`
	// If this reward is redeemed at a Vendor, this is the hash of the Vendor to go to in order to redeem the reward. Use this hash to look up the DestinyVendorDefinition.
	VendorHash int32 `json:"vendorHash,omitempty"`
	// For us to bother returning this info, we should be able to return some kind of information about why these rewards are grouped together. This is ideally that information. Look at how confident I am that this will always remain true.
	DisplayProperties *AllOfDestinyDefinitionsMilestonesDestinyMilestoneRewardEntryDefinitionDisplayProperties `json:"displayProperties,omitempty"`
	// If you want to follow BNet's ordering of these rewards, use this number within a given category to order the rewards. Yeah, I know. I feel dirty too.
	Order int32 `json:"order,omitempty"`
}
