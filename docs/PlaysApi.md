# {{classname}}

All URIs are relative to *https://api.collegefootballdata.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLivePlays**](PlaysApi.md#GetLivePlays) | **Get** /live/plays | Live metrics and PBP (Patreon only)
[**GetPlayStatTypes**](PlaysApi.md#GetPlayStatTypes) | **Get** /play/stat/types | Types of player play stats
[**GetPlayStats**](PlaysApi.md#GetPlayStats) | **Get** /play/stats | Play stats by play
[**GetPlayTypes**](PlaysApi.md#GetPlayTypes) | **Get** /play/types | Play types
[**GetPlays**](PlaysApi.md#GetPlays) | **Get** /plays | Play by play data

# **GetLivePlays**
> LivePlayByPlay GetLivePlays(ctx, id)
Live metrics and PBP (Patreon only)

Get live metrics and PBP

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Game id | 

### Return type

[**LivePlayByPlay**](LivePlayByPlay.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPlayStatTypes**
> []PlayStatType GetPlayStatTypes(ctx, )
Types of player play stats

Type of play stats

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]PlayStatType**](PlayStatType.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPlayStats**
> []PlayStat GetPlayStats(ctx, optional)
Play stats by play

Gets player stats associated by play (limit 1000)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PlaysApiGetPlayStatsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PlaysApiGetPlayStatsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **year** | **optional.Int32**| Year filter | 
 **week** | **optional.Int32**| Week filter | 
 **team** | **optional.String**| Team filter | 
 **gameId** | **optional.Int32**| gameId filter (from /games endpoint) | 
 **athleteId** | **optional.Int32**| athleteId filter (from /roster endpoint) | 
 **statTypeId** | **optional.Int32**| statTypeId filter (from /play/stat/types endpoint) | 
 **seasonType** | **optional.String**| regular, postseason, or both | 
 **conference** | **optional.String**| conference abbreviation filter | 

### Return type

[**[]PlayStat**](PlayStat.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPlayTypes**
> []PlayType GetPlayTypes(ctx, )
Play types

Types of plays

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]PlayType**](PlayType.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPlays**
> []Play GetPlays(ctx, year, week, optional)
Play by play data

Get play data and results

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **year** | **int32**| Year filter | 
  **week** | **int32**| Week filter (required if team, offense, or defense, not specified) | 
 **optional** | ***PlaysApiGetPlaysOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PlaysApiGetPlaysOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **seasonType** | **optional.String**| Season type filter | [default to regular]
 **team** | **optional.String**| Team filter | 
 **offense** | **optional.String**| Offensive team filter | 
 **defense** | **optional.String**| Defensive team filter | 
 **conference** | **optional.String**| Conference filter | 
 **offenseConference** | **optional.String**| Offensive conference filter | 
 **defenseConference** | **optional.String**| Defensive conference filter | 
 **playType** | **optional.Int32**| Play type filter | 
 **classification** | **optional.String**| Division classification filter (fbs/fcs/ii/iii) | 

### Return type

[**[]Play**](Play.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

