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

type TeamRecruitingRank struct {
	Year int32 `json:"year,omitempty"`
	Rank int32 `json:"rank,omitempty"`
	Team string `json:"team,omitempty"`
	Points float64 `json:"points,omitempty"`
}
