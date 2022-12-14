# Go API client for swagger

This is an API for accessing all sorts of college football data.  Please note that API keys should be supplied with \"Bearer \" prepended (e.g. \"Bearer your_key\"). API keys can be acquired from the CollegeFootballData.com website.

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 4.4.8
- Package version: 1.0.0
- Build package: io.swagger.codegen.v3.generators.go.GoClientCodegen

## Installation
Put the package under your project folder and add the following in import:
```golang
import "./swagger"
```

## Documentation for API Endpoints

All URIs are relative to *https://api.collegefootballdata.com/*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*BettingApi* | [**GetLines**](docs/BettingApi.md#getlines) | **Get** /lines | Betting lines
*CoachesApi* | [**GetCoaches**](docs/CoachesApi.md#getcoaches) | **Get** /coaches | Coaching records and history
*ConferencesApi* | [**GetConferences**](docs/ConferencesApi.md#getconferences) | **Get** /conferences | Conferences
*DraftApi* | [**GetDraftPicks**](docs/DraftApi.md#getdraftpicks) | **Get** /draft/picks | List of NFL Draft picks
*DraftApi* | [**GetNFLPositions**](docs/DraftApi.md#getnflpositions) | **Get** /draft/positions | List of NFL positions
*DraftApi* | [**GetNFLTeams**](docs/DraftApi.md#getnflteams) | **Get** /draft/teams | List of NFL teams
*DrivesApi* | [**GetDrives**](docs/DrivesApi.md#getdrives) | **Get** /drives | Drive data and results
*GamesApi* | [**GetAdvancedBoxScore**](docs/GamesApi.md#getadvancedboxscore) | **Get** /game/box/advanced | Advanced box scores
*GamesApi* | [**GetCalendar**](docs/GamesApi.md#getcalendar) | **Get** /calendar | Season calendar
*GamesApi* | [**GetGameMedia**](docs/GamesApi.md#getgamemedia) | **Get** /games/media | Game media information and schedules
*GamesApi* | [**GetGameWeather**](docs/GamesApi.md#getgameweather) | **Get** /games/weather | Game weather information (Patreon only)
*GamesApi* | [**GetGames**](docs/GamesApi.md#getgames) | **Get** /games | Games and results
*GamesApi* | [**GetPlayerGameStats**](docs/GamesApi.md#getplayergamestats) | **Get** /games/players | Player game stats
*GamesApi* | [**GetScoreboard**](docs/GamesApi.md#getscoreboard) | **Get** /scoreboard | Live game results (Patreon only)
*GamesApi* | [**GetTeamGameStats**](docs/GamesApi.md#getteamgamestats) | **Get** /games/teams | Team game stats
*GamesApi* | [**GetTeamRecords**](docs/GamesApi.md#getteamrecords) | **Get** /records | Team records
*MetricsApi* | [**GetGamePPA**](docs/MetricsApi.md#getgameppa) | **Get** /ppa/games | Team Predicated Points Added (PPA/EPA) by game
*MetricsApi* | [**GetPlayerGamePPA**](docs/MetricsApi.md#getplayergameppa) | **Get** /ppa/players/games | Player Predicated Points Added (PPA/EPA) broken down by game
*MetricsApi* | [**GetPlayerSeasonPPA**](docs/MetricsApi.md#getplayerseasonppa) | **Get** /ppa/players/season | Player Predicated Points Added (PPA/EPA) broken down by season
*MetricsApi* | [**GetPredictedPoints**](docs/MetricsApi.md#getpredictedpoints) | **Get** /ppa/predicted | Predicted Points (i.e. Expected Points or EP)
*MetricsApi* | [**GetPregameWinProbabilities**](docs/MetricsApi.md#getpregamewinprobabilities) | **Get** /metrics/wp/pregame | Pregame win probability data
*MetricsApi* | [**GetTeamPPA**](docs/MetricsApi.md#getteamppa) | **Get** /ppa/teams | Predicted Points Added (PPA/EPA) data by team
*MetricsApi* | [**GetWinProbabilityData**](docs/MetricsApi.md#getwinprobabilitydata) | **Get** /metrics/wp | Win probability chart data
*PlayersApi* | [**GetPlayerSeasonStats**](docs/PlayersApi.md#getplayerseasonstats) | **Get** /stats/player/season | Player stats by season
*PlayersApi* | [**GetPlayerUsage**](docs/PlayersApi.md#getplayerusage) | **Get** /player/usage | Player usage metrics broken down by season
*PlayersApi* | [**GetReturningProduction**](docs/PlayersApi.md#getreturningproduction) | **Get** /player/returning | Team returning production metrics
*PlayersApi* | [**GetTransferPortal**](docs/PlayersApi.md#gettransferportal) | **Get** /player/portal | Transfer portal by season
*PlayersApi* | [**PlayerSearch**](docs/PlayersApi.md#playersearch) | **Get** /player/search | Search for player information
*PlaysApi* | [**GetLivePlays**](docs/PlaysApi.md#getliveplays) | **Get** /live/plays | Live metrics and PBP (Patreon only)
*PlaysApi* | [**GetPlayStatTypes**](docs/PlaysApi.md#getplaystattypes) | **Get** /play/stat/types | Types of player play stats
*PlaysApi* | [**GetPlayStats**](docs/PlaysApi.md#getplaystats) | **Get** /play/stats | Play stats by play
*PlaysApi* | [**GetPlayTypes**](docs/PlaysApi.md#getplaytypes) | **Get** /play/types | Play types
*PlaysApi* | [**GetPlays**](docs/PlaysApi.md#getplays) | **Get** /plays | Play by play data
*RankingsApi* | [**GetRankings**](docs/RankingsApi.md#getrankings) | **Get** /rankings | Historical polls and rankings
*RatingsApi* | [**GetConferenceSPRatings**](docs/RatingsApi.md#getconferencespratings) | **Get** /ratings/sp/conferences | Historical SP+ ratings by conference
*RatingsApi* | [**GetEloRatings**](docs/RatingsApi.md#geteloratings) | **Get** /ratings/elo | Historical Elo ratings
*RatingsApi* | [**GetSPRatings**](docs/RatingsApi.md#getspratings) | **Get** /ratings/sp | Historical SP+ ratings
*RatingsApi* | [**GetSRSRatings**](docs/RatingsApi.md#getsrsratings) | **Get** /ratings/srs | Historical SRS ratings
*RecruitingApi* | [**GetRecruitingGroups**](docs/RecruitingApi.md#getrecruitinggroups) | **Get** /recruiting/groups | Recruit position group ratings
*RecruitingApi* | [**GetRecruitingPlayers**](docs/RecruitingApi.md#getrecruitingplayers) | **Get** /recruiting/players | Player recruiting ratings and rankings
*RecruitingApi* | [**GetRecruitingTeams**](docs/RecruitingApi.md#getrecruitingteams) | **Get** /recruiting/teams | Team recruiting rankings and ratings
*StatsApi* | [**GetAdvancedTeamGameStats**](docs/StatsApi.md#getadvancedteamgamestats) | **Get** /stats/game/advanced | Advanced team metrics by game
*StatsApi* | [**GetAdvancedTeamSeasonStats**](docs/StatsApi.md#getadvancedteamseasonstats) | **Get** /stats/season/advanced | Advanced team metrics by season
*StatsApi* | [**GetStatCategories**](docs/StatsApi.md#getstatcategories) | **Get** /stats/categories | Team stat categories
*StatsApi* | [**GetTeamSeasonStats**](docs/StatsApi.md#getteamseasonstats) | **Get** /stats/season | Team statistics by season
*TeamsApi* | [**GetFbsTeams**](docs/TeamsApi.md#getfbsteams) | **Get** /teams/fbs | FBS team list
*TeamsApi* | [**GetRoster**](docs/TeamsApi.md#getroster) | **Get** /roster | Team rosters
*TeamsApi* | [**GetTalent**](docs/TeamsApi.md#gettalent) | **Get** /talent | Team talent composite rankings
*TeamsApi* | [**GetTeamMatchup**](docs/TeamsApi.md#getteammatchup) | **Get** /teams/matchup | Team matchup history
*TeamsApi* | [**GetTeams**](docs/TeamsApi.md#getteams) | **Get** /teams | Team information
*VenuesApi* | [**GetVenues**](docs/VenuesApi.md#getvenues) | **Get** /venues | Arena and venue information

## Documentation For Models

 - [AdvancedGameStat](docs/AdvancedGameStat.md)
 - [AdvancedGameStatOffense](docs/AdvancedGameStatOffense.md)
 - [AdvancedGameStatOffenseRushingPlays](docs/AdvancedGameStatOffenseRushingPlays.md)
 - [AdvancedGameStatOffenseStandardDowns](docs/AdvancedGameStatOffenseStandardDowns.md)
 - [AdvancedSeasonStat](docs/AdvancedSeasonStat.md)
 - [AdvancedSeasonStatOffense](docs/AdvancedSeasonStatOffense.md)
 - [AdvancedSeasonStatOffenseFieldPosition](docs/AdvancedSeasonStatOffenseFieldPosition.md)
 - [AdvancedSeasonStatOffenseRushingPlays](docs/AdvancedSeasonStatOffenseRushingPlays.md)
 - [AdvancedSeasonStatOffenseStandardDowns](docs/AdvancedSeasonStatOffenseStandardDowns.md)
 - [BoxScore](docs/BoxScore.md)
 - [BoxScorePlayers](docs/BoxScorePlayers.md)
 - [BoxScorePlayersAverage](docs/BoxScorePlayersAverage.md)
 - [BoxScorePlayersPpa](docs/BoxScorePlayersPpa.md)
 - [BoxScorePlayersUsage](docs/BoxScorePlayersUsage.md)
 - [BoxScoreTeams](docs/BoxScoreTeams.md)
 - [BoxScoreTeamsExplosiveness](docs/BoxScoreTeamsExplosiveness.md)
 - [BoxScoreTeamsFieldPosition](docs/BoxScoreTeamsFieldPosition.md)
 - [BoxScoreTeamsHavoc](docs/BoxScoreTeamsHavoc.md)
 - [BoxScoreTeamsOverall](docs/BoxScoreTeamsOverall.md)
 - [BoxScoreTeamsPpa](docs/BoxScoreTeamsPpa.md)
 - [BoxScoreTeamsRushing](docs/BoxScoreTeamsRushing.md)
 - [BoxScoreTeamsScoringOpportunities](docs/BoxScoreTeamsScoringOpportunities.md)
 - [BoxScoreTeamsSuccessRates](docs/BoxScoreTeamsSuccessRates.md)
 - [Coach](docs/Coach.md)
 - [CoachSeasons](docs/CoachSeasons.md)
 - [Conference](docs/Conference.md)
 - [ConferenceSpRating](docs/ConferenceSpRating.md)
 - [ConferenceSpRatingDefense](docs/ConferenceSpRatingDefense.md)
 - [ConferenceSpRatingOffense](docs/ConferenceSpRatingOffense.md)
 - [DraftPick](docs/DraftPick.md)
 - [DraftPickHometownInfo](docs/DraftPickHometownInfo.md)
 - [DraftPosition](docs/DraftPosition.md)
 - [DraftTeam](docs/DraftTeam.md)
 - [Drive](docs/Drive.md)
 - [DriveStartTime](docs/DriveStartTime.md)
 - [Game](docs/Game.md)
 - [GameLines](docs/GameLines.md)
 - [GameLinesLines](docs/GameLinesLines.md)
 - [GameMedia](docs/GameMedia.md)
 - [GamePpa](docs/GamePpa.md)
 - [GamePpaOffense](docs/GamePpaOffense.md)
 - [GameWeather](docs/GameWeather.md)
 - [LivePlayByPlay](docs/LivePlayByPlay.md)
 - [LivePlayByPlayDrives](docs/LivePlayByPlayDrives.md)
 - [LivePlayByPlayPlays](docs/LivePlayByPlayPlays.md)
 - [LivePlayByPlayTeams](docs/LivePlayByPlayTeams.md)
 - [Play](docs/Play.md)
 - [PlayStat](docs/PlayStat.md)
 - [PlayStatType](docs/PlayStatType.md)
 - [PlayType](docs/PlayType.md)
 - [PlayWp](docs/PlayWp.md)
 - [Player](docs/Player.md)
 - [PlayerGame](docs/PlayerGame.md)
 - [PlayerGameAthletes](docs/PlayerGameAthletes.md)
 - [PlayerGameCategories](docs/PlayerGameCategories.md)
 - [PlayerGamePpa](docs/PlayerGamePpa.md)
 - [PlayerGamePpaAveragePpa](docs/PlayerGamePpaAveragePpa.md)
 - [PlayerGameSchool](docs/PlayerGameSchool.md)
 - [PlayerGameTeams](docs/PlayerGameTeams.md)
 - [PlayerGameTypes](docs/PlayerGameTypes.md)
 - [PlayerSearchResult](docs/PlayerSearchResult.md)
 - [PlayerSeasonPpa](docs/PlayerSeasonPpa.md)
 - [PlayerSeasonPpaAveragePpa](docs/PlayerSeasonPpaAveragePpa.md)
 - [PlayerSeasonStat](docs/PlayerSeasonStat.md)
 - [PlayerUsage](docs/PlayerUsage.md)
 - [PlayerUsageUsage](docs/PlayerUsageUsage.md)
 - [PortalPlayer](docs/PortalPlayer.md)
 - [PositionGroupRecruitingRating](docs/PositionGroupRecruitingRating.md)
 - [PredictedPoints](docs/PredictedPoints.md)
 - [PregameWp](docs/PregameWp.md)
 - [RankingWeek](docs/RankingWeek.md)
 - [RankingWeekPolls](docs/RankingWeekPolls.md)
 - [RankingWeekRanks](docs/RankingWeekRanks.md)
 - [Recruit](docs/Recruit.md)
 - [RecruitHometownInfo](docs/RecruitHometownInfo.md)
 - [ReturningProduction](docs/ReturningProduction.md)
 - [ScoreboardGame](docs/ScoreboardGame.md)
 - [ScoreboardGameBetting](docs/ScoreboardGameBetting.md)
 - [ScoreboardGameHomeTeam](docs/ScoreboardGameHomeTeam.md)
 - [ScoreboardGameVenue](docs/ScoreboardGameVenue.md)
 - [ScoreboardGameWeather](docs/ScoreboardGameWeather.md)
 - [Team](docs/Team.md)
 - [TeamEloRating](docs/TeamEloRating.md)
 - [TeamGame](docs/TeamGame.md)
 - [TeamGameStats](docs/TeamGameStats.md)
 - [TeamGameTeams](docs/TeamGameTeams.md)
 - [TeamLocation](docs/TeamLocation.md)
 - [TeamMatchup](docs/TeamMatchup.md)
 - [TeamMatchupGames](docs/TeamMatchupGames.md)
 - [TeamPpa](docs/TeamPpa.md)
 - [TeamPpaOffense](docs/TeamPpaOffense.md)
 - [TeamPpaOffenseCumulative](docs/TeamPpaOffenseCumulative.md)
 - [TeamRecord](docs/TeamRecord.md)
 - [TeamRecordTotal](docs/TeamRecordTotal.md)
 - [TeamRecruitingRank](docs/TeamRecruitingRank.md)
 - [TeamSeason](docs/TeamSeason.md)
 - [TeamSeasonStat](docs/TeamSeasonStat.md)
 - [TeamSpRating](docs/TeamSpRating.md)
 - [TeamSpRatingDefense](docs/TeamSpRatingDefense.md)
 - [TeamSpRatingDefenseHavoc](docs/TeamSpRatingDefenseHavoc.md)
 - [TeamSpRatingOffense](docs/TeamSpRatingOffense.md)
 - [TeamSpRatingSpecialTeams](docs/TeamSpRatingSpecialTeams.md)
 - [TeamSrsRating](docs/TeamSrsRating.md)
 - [TeamTalent](docs/TeamTalent.md)
 - [Venue](docs/Venue.md)
 - [VenueLocation](docs/VenueLocation.md)
 - [Week](docs/Week.md)

## Documentation For Authorization

## ApiKeyAuth
- **Type**: API key 

Example
```golang
auth := context.WithValue(context.Background(), sw.ContextAPIKey, sw.APIKey{
	Key: "APIKEY",
	Prefix: "Bearer", // Omit if not necessary.
})
r, err := client.Service.Operation(auth, args)
```

## Author

admin@collegefootballdata.com
