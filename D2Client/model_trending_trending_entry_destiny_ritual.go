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

type TrendingTrendingEntryDestinyRitual struct {
	Image string `json:"image,omitempty"`
	Icon string `json:"icon,omitempty"`
	Title string `json:"title,omitempty"`
	Subtitle string `json:"subtitle,omitempty"`
	DateStart time.Time `json:"dateStart,omitempty"`
	DateEnd time.Time `json:"dateEnd,omitempty"`
	// A destiny event does not necessarily have a related Milestone, but if it does the details will be returned here.
	MilestoneDetails *AllOfTrendingTrendingEntryDestinyRitualMilestoneDetails `json:"milestoneDetails,omitempty"`
	// A destiny event will not necessarily have milestone \"custom content\", but if it does the details will be here.
	EventContent *AllOfTrendingTrendingEntryDestinyRitualEventContent `json:"eventContent,omitempty"`
}
