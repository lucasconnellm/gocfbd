/*
 * College Football Data API
 *
 * This is an API for accessing all sorts of college football data.  Please note that API keys should be supplied with \"Bearer \" prepended (e.g. \"Bearer your_key\"). API keys can be acquired from the CollegeFootballData.com website.
 *
 * API version: 4.4.8
 * Contact: admin@collegefootballdata.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type LivePlayByPlayDrives struct {
	Id int32 `json:"id,omitempty"`
	OffenseId int32 `json:"offenseId,omitempty"`
	Offense string `json:"offense,omitempty"`
	DefenseId int32 `json:"defenseId,omitempty"`
	Defense string `json:"defense,omitempty"`
	PlayCount int32 `json:"playCount,omitempty"`
	Yards int32 `json:"yards,omitempty"`
	StartPeriod int32 `json:"startPeriod,omitempty"`
	StartClock string `json:"startClock,omitempty"`
	StartYardsToGoal int32 `json:"startYardsToGoal,omitempty"`
	EndPeriod int32 `json:"endPeriod,omitempty"`
	EndClock string `json:"endClock,omitempty"`
	EndYardsToGoal int32 `json:"endYardsToGoal,omitempty"`
	Duration string `json:"duration,omitempty"`
	ScoringOpportunity bool `json:"scoringOpportunity,omitempty"`
	Plays []LivePlayByPlayPlays `json:"plays,omitempty"`
}
