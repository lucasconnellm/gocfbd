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

type Team struct {
	Id int32 `json:"id,omitempty"`
	School string `json:"school,omitempty"`
	Mascot string `json:"mascot,omitempty"`
	Abbreviation string `json:"abbreviation,omitempty"`
	AltName1 string `json:"alt_name_1,omitempty"`
	AltName2 string `json:"alt_name_2,omitempty"`
	AltName3 string `json:"alt_name_3,omitempty"`
	Classification string `json:"classification,omitempty"`
	Conference string `json:"conference,omitempty"`
	Division string `json:"division,omitempty"`
	Color string `json:"color,omitempty"`
	AltColor string `json:"alt_color,omitempty"`
	Logos []string `json:"logos,omitempty"`
	Twitter string `json:"twitter,omitempty"`
	Location *TeamLocation `json:"location,omitempty"`
}