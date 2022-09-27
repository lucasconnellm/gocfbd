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

type DraftPickHometownInfo struct {
	City string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
	Country string `json:"country,omitempty"`
	Latitude float64 `json:"latitude,omitempty"`
	Longitude float64 `json:"longitude,omitempty"`
	CountryFips int32 `json:"countryFips,omitempty"`
}
