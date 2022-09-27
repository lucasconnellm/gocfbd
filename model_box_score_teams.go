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

type BoxScoreTeams struct {
	Ppa []BoxScoreTeamsPpa `json:"ppa,omitempty"`
	CumulativePpa []BoxScoreTeamsPpa `json:"cumulativePpa,omitempty"`
	SuccessRates []BoxScoreTeamsSuccessRates `json:"successRates,omitempty"`
	Explosiveness []BoxScoreTeamsExplosiveness `json:"explosiveness,omitempty"`
	Rushing []BoxScoreTeamsRushing `json:"rushing,omitempty"`
	Havoc []BoxScoreTeamsHavoc `json:"havoc,omitempty"`
	ScoringOpportunities []BoxScoreTeamsScoringOpportunities `json:"scoringOpportunities,omitempty"`
	FieldPosition []BoxScoreTeamsFieldPosition `json:"fieldPosition,omitempty"`
}