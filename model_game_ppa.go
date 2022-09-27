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

type GamePpa struct {
	GameId int32 `json:"gameId,omitempty"`
	Season int32 `json:"season,omitempty"`
	Week int32 `json:"week,omitempty"`
	Team string `json:"team,omitempty"`
	Conference string `json:"conference,omitempty"`
	Opponent string `json:"opponent,omitempty"`
	Offense *GamePpaOffense `json:"offense,omitempty"`
	Defense *GamePpaOffense `json:"defense,omitempty"`
}
