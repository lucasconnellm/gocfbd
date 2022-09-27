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

type BoxScoreTeamsOverall struct {
	Total float64 `json:"total,omitempty"`
	Quarter1 float64 `json:"quarter1,omitempty"`
	Quarter2 float64 `json:"quarter2,omitempty"`
	Quarter3 float64 `json:"quarter3,omitempty"`
	Quarter4 float64 `json:"quarter4,omitempty"`
}
