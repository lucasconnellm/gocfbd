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

type PlayerUsage struct {
	Season int32 `json:"season,omitempty"`
	Id int32 `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Position string `json:"position,omitempty"`
	Team string `json:"team,omitempty"`
	Conference string `json:"conference,omitempty"`
	Usage *PlayerUsageUsage `json:"usage,omitempty"`
}
