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

type Player struct {
	Id int32 `json:"id,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName string `json:"last_name,omitempty"`
	Team string `json:"team,omitempty"`
	Height int32 `json:"height,omitempty"`
	Weight int32 `json:"weight,omitempty"`
	Jersey int32 `json:"jersey,omitempty"`
	Year int32 `json:"year,omitempty"`
	Position string `json:"position,omitempty"`
	HomeCity string `json:"home_city,omitempty"`
	HomeState string `json:"home_state,omitempty"`
	HomeCountry string `json:"home_country,omitempty"`
	HomeLatitude float64 `json:"home_latitude,omitempty"`
	HomeLongitude float64 `json:"home_longitude,omitempty"`
	HomeCountyFips string `json:"home_county_fips,omitempty"`
	RecruitIds []int32 `json:"recruit_ids,omitempty"`
}
