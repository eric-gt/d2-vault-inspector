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

// If you are playing in an activity, this is some information about it.  Note that we cannot guarantee any of this resembles what ends up in the PGCR in any way. They are sourced by two entirely separate systems with their own logic, and the one we source this data from should be considered non-authoritative in comparison.
type DestinyComponentsProfilesDestinyProfileTransitoryCurrentActivity struct {
	// When the activity started.
	StartTime time.Time `json:"startTime,omitempty"`
	// If you're still in it but it \"ended\" (like when folks are dancing around the loot after they beat a boss), this is when the activity ended.
	EndTime time.Time `json:"endTime,omitempty"`
	// This is what our non-authoritative source thought the score was.
	Score float32 `json:"score,omitempty"`
	// If you have human opponents, this is the highest opposing team's score.
	HighestOpposingFactionScore float32 `json:"highestOpposingFactionScore,omitempty"`
	// This is how many human or poorly crafted aimbot opponents you have.
	NumberOfOpponents int32 `json:"numberOfOpponents,omitempty"`
	// This is how many human or poorly crafted aimbots are on your team.
	NumberOfPlayers int32 `json:"numberOfPlayers,omitempty"`
}
