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
import (
	"time"
)

// Information about milestones, presented in a character state-agnostic manner. Combine this data with DestinyMilestoneDefinition to get a full picture of the milestone, which is basically a checklist of things to do in the game. Think of this as GetPublicAdvisors 3.0, for those who used the Destiny 1 API.
type DestinyMilestonesDestinyPublicMilestone struct {
	// The hash identifier for the milestone. Use it to look up the DestinyMilestoneDefinition for static data about the Milestone.
	MilestoneHash int32 `json:"milestoneHash,omitempty"`
	// A milestone not need have even a single quest, but if there are active quests they will be returned here.
	AvailableQuests []DestinyMilestonesDestinyPublicMilestoneQuest `json:"availableQuests,omitempty"`
	Activities []DestinyMilestonesDestinyPublicMilestoneChallengeActivity `json:"activities,omitempty"`
	// Sometimes milestones - or activities active in milestones - will have relevant vendors. These are the vendors that are currently relevant.  Deprecated, already, for the sake of the new \"vendors\" property that has more data. What was I thinking.
	VendorHashes []int32 `json:"vendorHashes,omitempty"`
	// This is why we can't have nice things. This is the ordered list of vendors to be shown that relate to this milestone, potentially along with other interesting data.
	Vendors []DestinyMilestonesDestinyPublicMilestoneVendor `json:"vendors,omitempty"`
	// If known, this is the date when the Milestone started/became active.
	StartDate time.Time `json:"startDate,omitempty"`
	// If known, this is the date when the Milestone will expire/recycle/end.
	EndDate time.Time `json:"endDate,omitempty"`
	// Used for ordering milestones in a display to match how we order them in BNet. May pull from static data, or possibly in the future from dynamic information.
	Order int32 `json:"order,omitempty"`
}
