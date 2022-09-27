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

type TeamRecordTotal struct {
	Games int32 `json:"games,omitempty"`
	Wins int32 `json:"wins,omitempty"`
	Losses int32 `json:"losses,omitempty"`
	Ties int32 `json:"ties,omitempty"`
}
