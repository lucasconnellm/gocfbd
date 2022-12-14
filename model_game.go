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

type Game struct {
	Id int32 `json:"id,omitempty"`
	Season int32 `json:"season,omitempty"`
	Week int32 `json:"week,omitempty"`
	SeasonType string `json:"season_type,omitempty"`
	StartDate string `json:"start_date,omitempty"`
	StartTimeTbd bool `json:"start_time_tbd,omitempty"`
	NeutralSite bool `json:"neutral_site,omitempty"`
	ConferenceGame bool `json:"conference_game,omitempty"`
	Attendance int32 `json:"attendance,omitempty"`
	VenueId int32 `json:"venue_id,omitempty"`
	Venue string `json:"venue,omitempty"`
	HomeId int32 `json:"home_id,omitempty"`
	HomeTeam string `json:"home_team,omitempty"`
	HomeConference string `json:"home_conference,omitempty"`
	HomeDivision string `json:"home_division,omitempty"`
	HomePoints int32 `json:"home_points,omitempty"`
	HomeLineScores []int32 `json:"home_line_scores,omitempty"`
	HomePostWinProb float64 `json:"home_post_win_prob,omitempty"`
	HomePregameElo int32 `json:"home_pregame_elo,omitempty"`
	HomePostgameElo int32 `json:"home_postgame_elo,omitempty"`
	AwayId int32 `json:"away_id,omitempty"`
	AwayTeam string `json:"away_team,omitempty"`
	AwayConference string `json:"away_conference,omitempty"`
	AwayDivision string `json:"away_division,omitempty"`
	AwayPoints int32 `json:"away_points,omitempty"`
	AwayLineScores []int32 `json:"away_line_scores,omitempty"`
	AwayPostWinProb float64 `json:"away_post_win_prob,omitempty"`
	AwayPregameElo int32 `json:"away_pregame_elo,omitempty"`
	AwayPostgameElo int32 `json:"away_postgame_elo,omitempty"`
	ExcitementIndex float64 `json:"excitement_index,omitempty"`
	Highlights string `json:"highlights,omitempty"`
	Notes string `json:"notes,omitempty"`
}
