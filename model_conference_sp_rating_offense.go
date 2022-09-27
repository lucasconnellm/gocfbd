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

type ConferenceSpRatingOffense struct {
	Rating float64 `json:"rating,omitempty"`
	Success float64 `json:"success,omitempty"`
	Explosiveness float64 `json:"explosiveness,omitempty"`
	Rushing float64 `json:"rushing,omitempty"`
	Passing float64 `json:"passing,omitempty"`
	StandardDowns float64 `json:"standardDowns,omitempty"`
	PassingDowns float64 `json:"passingDowns,omitempty"`
	RunRate float64 `json:"runRate,omitempty"`
	Pace float64 `json:"pace,omitempty"`
}