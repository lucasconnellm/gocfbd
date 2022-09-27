# {{classname}}

All URIs are relative to *https://api.collegefootballdata.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetGamePPA**](MetricsApi.md#GetGamePPA) | **Get** /ppa/games | Team Predicated Points Added (PPA/EPA) by game
[**GetPlayerGamePPA**](MetricsApi.md#GetPlayerGamePPA) | **Get** /ppa/players/games | Player Predicated Points Added (PPA/EPA) broken down by game
[**GetPlayerSeasonPPA**](MetricsApi.md#GetPlayerSeasonPPA) | **Get** /ppa/players/season | Player Predicated Points Added (PPA/EPA) broken down by season
[**GetPredictedPoints**](MetricsApi.md#GetPredictedPoints) | **Get** /ppa/predicted | Predicted Points (i.e. Expected Points or EP)
[**GetPregameWinProbabilities**](MetricsApi.md#GetPregameWinProbabilities) | **Get** /metrics/wp/pregame | Pregame win probability data
[**GetTeamPPA**](MetricsApi.md#GetTeamPPA) | **Get** /ppa/teams | Predicted Points Added (PPA/EPA) data by team
[**GetWinProbabilityData**](MetricsApi.md#GetWinProbabilityData) | **Get** /metrics/wp | Win probability chart data

# **GetGamePPA**
> []GamePpa GetGamePPA(ctx, year, optional)
Team Predicated Points Added (PPA/EPA) by game

Predicted Points Added (PPA) by game

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **year** | **int32**| Year filter | 
 **optional** | ***MetricsApiGetGamePPAOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MetricsApiGetGamePPAOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **week** | **optional.Int32**| Week filter | 
 **team** | **optional.String**| Team filter | 
 **conference** | **optional.String**| Conference filter | 
 **excludeGarbageTime** | **optional.Bool**| Filter to remove garbage time plays from calculations | 
 **seasonType** | **optional.String**| Season type filter (regular or postseason) | [default to regular]

### Return type

[**[]GamePpa**](GamePPA.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPlayerGamePPA**
> []PlayerGamePpa GetPlayerGamePPA(ctx, optional)
Player Predicated Points Added (PPA/EPA) broken down by game

Predicted Points Added (PPA) by player game

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MetricsApiGetPlayerGamePPAOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MetricsApiGetPlayerGamePPAOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **year** | **optional.Int32**| Year filter | 
 **week** | **optional.Int32**| Week filter | 
 **team** | **optional.String**| Team filter | 
 **position** | **optional.String**| Position abbreviation filter | 
 **playerId** | **optional.Int32**| Player id filter | 
 **threshold** | **optional.String**| Minimum play threshold filter | 
 **excludeGarbageTime** | **optional.Bool**| Filter to remove garbage time plays from calculations | 
 **seasonType** | **optional.String**| Season type filter (regular or postseason) | [default to regular]

### Return type

[**[]PlayerGamePpa**](PlayerGamePPA.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPlayerSeasonPPA**
> []PlayerSeasonPpa GetPlayerSeasonPPA(ctx, optional)
Player Predicated Points Added (PPA/EPA) broken down by season

Predicted Points Added (PPA) by player season

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MetricsApiGetPlayerSeasonPPAOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MetricsApiGetPlayerSeasonPPAOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **year** | **optional.Int32**| Year filter | 
 **team** | **optional.String**| Team filter | 
 **conference** | **optional.String**| Conference abbreviation filter | 
 **position** | **optional.String**| Position abbreviation filter | 
 **playerId** | **optional.Int32**| Player id filter | 
 **threshold** | **optional.String**| Minimum play threshold filter | 
 **excludeGarbageTime** | **optional.Bool**| Filter to remove garbage time plays from calculations | 

### Return type

[**[]PlayerSeasonPpa**](PlayerSeasonPPA.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPredictedPoints**
> []PredictedPoints GetPredictedPoints(ctx, down, distance)
Predicted Points (i.e. Expected Points or EP)

Predicted Points

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **down** | **int32**| Down filter | 
  **distance** | **int32**| Distance filter | 

### Return type

[**[]PredictedPoints**](PredictedPoints.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPregameWinProbabilities**
> []PregameWp GetPregameWinProbabilities(ctx, optional)
Pregame win probability data

Pregame win probabilities

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MetricsApiGetPregameWinProbabilitiesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MetricsApiGetPregameWinProbabilitiesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **year** | **optional.Int32**| Year filter | 
 **week** | **optional.Int32**| Week filter | 
 **team** | **optional.String**| Team filter | 
 **seasonType** | **optional.String**| regular or postseason | 

### Return type

[**[]PregameWp**](PregameWP.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTeamPPA**
> []TeamPpa GetTeamPPA(ctx, optional)
Predicted Points Added (PPA/EPA) data by team

Predicted Points Added (PPA)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MetricsApiGetTeamPPAOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MetricsApiGetTeamPPAOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **year** | **optional.Int32**| Year filter (required if team not specified) | 
 **team** | **optional.String**| Team filter (required if year not specified) | 
 **conference** | **optional.String**| Conference filter | 
 **excludeGarbageTime** | **optional.Bool**| Filter to remove garbage time plays from calculations | 

### Return type

[**[]TeamPpa**](TeamPPA.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetWinProbabilityData**
> []PlayWp GetWinProbabilityData(ctx, gameId)
Win probability chart data

Win probability data

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gameId** | **int32**| Game id filter | 

### Return type

[**[]PlayWp**](PlayWP.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

