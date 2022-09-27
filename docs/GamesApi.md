# {{classname}}

All URIs are relative to *https://api.collegefootballdata.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAdvancedBoxScore**](GamesApi.md#GetAdvancedBoxScore) | **Get** /game/box/advanced | Advanced box scores
[**GetCalendar**](GamesApi.md#GetCalendar) | **Get** /calendar | Season calendar
[**GetGameMedia**](GamesApi.md#GetGameMedia) | **Get** /games/media | Game media information and schedules
[**GetGameWeather**](GamesApi.md#GetGameWeather) | **Get** /games/weather | Game weather information (Patreon only)
[**GetGames**](GamesApi.md#GetGames) | **Get** /games | Games and results
[**GetPlayerGameStats**](GamesApi.md#GetPlayerGameStats) | **Get** /games/players | Player game stats
[**GetScoreboard**](GamesApi.md#GetScoreboard) | **Get** /scoreboard | Live game results (Patreon only)
[**GetTeamGameStats**](GamesApi.md#GetTeamGameStats) | **Get** /games/teams | Team game stats
[**GetTeamRecords**](GamesApi.md#GetTeamRecords) | **Get** /records | Team records

# **GetAdvancedBoxScore**
> BoxScore GetAdvancedBoxScore(ctx, gameId)
Advanced box scores

Get advanced box score data

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gameId** | **int32**| Game id parameters | 

### Return type

[**BoxScore**](BoxScore.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCalendar**
> []Week GetCalendar(ctx, year)
Season calendar

Get calendar of weeks by season

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **year** | **int32**| Year filter | 

### Return type

[**[]Week**](Week.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGameMedia**
> []GameMedia GetGameMedia(ctx, year, optional)
Game media information and schedules

Game media information (TV, radio, etc)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **year** | **int32**| Year filter | 
 **optional** | ***GamesApiGetGameMediaOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GamesApiGetGameMediaOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **week** | **optional.Int32**| Week filter | 
 **seasonType** | **optional.String**| Season type filter (regular, postseason, or both) | 
 **team** | **optional.String**| Team filter | 
 **conference** | **optional.String**| Conference filter | 
 **mediaType** | **optional.String**| Media type filter (tv, radio, web, ppv, or mobile) | 
 **classification** | **optional.String**| Division classification filter (fbs/fcs/ii/iii) | 

### Return type

[**[]GameMedia**](GameMedia.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGameWeather**
> []GameWeather GetGameWeather(ctx, optional)
Game weather information (Patreon only)

Weather information for the hour of kickoff

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GamesApiGetGameWeatherOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GamesApiGetGameWeatherOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gameId** | **optional.Int32**| Game id filter (required if no year) | 
 **year** | **optional.Int32**| Year filter (required if no game id) | 
 **week** | **optional.Int32**| Week filter | 
 **seasonType** | **optional.String**| Season type filter (regular, postseason, or both) | 
 **team** | **optional.String**| Team filter | 
 **conference** | **optional.String**| Conference filter | 
 **classification** | **optional.String**| Division classification filter (fbs/fcs/ii/iii) | 

### Return type

[**[]GameWeather**](GameWeather.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGames**
> []Game GetGames(ctx, year, optional)
Games and results

Get game results

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **year** | **int32**| Year/season filter for games | 
 **optional** | ***GamesApiGetGamesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GamesApiGetGamesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **week** | **optional.Int32**| Week filter | 
 **seasonType** | **optional.String**| Season type filter (regular or postseason) | [default to regular]
 **team** | **optional.String**| Team | 
 **home** | **optional.String**| Home team filter | 
 **away** | **optional.String**| Away team filter | 
 **conference** | **optional.String**| Conference abbreviation filter | 
 **division** | **optional.String**| Division classification filter (fbs/fcs/ii/iii) | 
 **id** | **optional.Int32**| id filter for querying a single game | 

### Return type

[**[]Game**](Game.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPlayerGameStats**
> []PlayerGame GetPlayerGameStats(ctx, year, optional)
Player game stats

Player stats broken down by game

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **year** | **int32**| Year/season filter for games | 
 **optional** | ***GamesApiGetPlayerGameStatsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GamesApiGetPlayerGameStatsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **week** | **optional.Int32**| Week filter | 
 **seasonType** | **optional.String**| Season type filter (regular or postseason) | [default to regular]
 **team** | **optional.String**| Team filter | 
 **conference** | **optional.String**| Conference abbreviation filter | 
 **category** | **optional.String**| Category filter (e.g defensive) | 
 **gameId** | **optional.Int32**| Game id filter | 

### Return type

[**[]PlayerGame**](PlayerGame.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetScoreboard**
> []ScoreboardGame GetScoreboard(ctx, optional)
Live game results (Patreon only)

Get live game results

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GamesApiGetScoreboardOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GamesApiGetScoreboardOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **classification** | **optional.String**| Classification filter (fbs, fcs, ii, or iii). Defaults to fbs. | 
 **conference** | **optional.String**| Conference abbreviation filter. | 

### Return type

[**[]ScoreboardGame**](ScoreboardGame.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTeamGameStats**
> []TeamGame GetTeamGameStats(ctx, year, optional)
Team game stats

Team stats broken down by game

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **year** | **int32**| Year/season filter for games | 
 **optional** | ***GamesApiGetTeamGameStatsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GamesApiGetTeamGameStatsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **week** | **optional.Int32**| Week filter | 
 **seasonType** | **optional.String**| Season type filter (regular or postseason) | [default to regular]
 **team** | **optional.String**| Team filter | 
 **conference** | **optional.String**| Conference abbreviation filter | 
 **gameId** | **optional.Int32**| Game id filter | 
 **classification** | **optional.String**| Division classification filter (fbs/fcs/ii/iii) | 

### Return type

[**[]TeamGame**](TeamGame.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTeamRecords**
> []TeamRecord GetTeamRecords(ctx, optional)
Team records

Get team records by year

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GamesApiGetTeamRecordsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GamesApiGetTeamRecordsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **year** | **optional.Int32**| Year filter | 
 **team** | **optional.String**| Team filter | 
 **conference** | **optional.String**| Conference filter | 

### Return type

[**[]TeamRecord**](TeamRecord.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

