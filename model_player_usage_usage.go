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

type PlayerUsageUsage struct {
	Overall float64 `json:"overall,omitempty"`
	Pass float64 `json:"pass,omitempty"`
	Rush float64 `json:"rush,omitempty"`
	FirstDown float64 `json:"firstDown,omitempty"`
	SecondDown float64 `json:"secondDown,omitempty"`
	ThirdDown float64 `json:"thirdDown,omitempty"`
	StandardDowns float64 `json:"standardDowns,omitempty"`
	PassingDowns float64 `json:"passingDowns,omitempty"`
}
