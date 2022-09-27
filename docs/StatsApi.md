# {{classname}}

All URIs are relative to *https://api.collegefootballdata.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAdvancedTeamGameStats**](StatsApi.md#GetAdvancedTeamGameStats) | **Get** /stats/game/advanced | Advanced team metrics by game
[**GetAdvancedTeamSeasonStats**](StatsApi.md#GetAdvancedTeamSeasonStats) | **Get** /stats/season/advanced | Advanced team metrics by season
[**GetStatCategories**](StatsApi.md#GetStatCategories) | **Get** /stats/categories | Team stat categories
[**GetTeamSeasonStats**](StatsApi.md#GetTeamSeasonStats) | **Get** /stats/season | Team statistics by season

# **GetAdvancedTeamGameStats**
> []AdvancedGameStat GetAdvancedTeamGameStats(ctx, optional)
Advanced team metrics by game

Advanced team game stats

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***StatsApiGetAdvancedTeamGameStatsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StatsApiGetAdvancedTeamGameStatsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **year** | **optional.Int32**| Year filter (required if no team specified) | 
 **week** | **optional.Int32**| Week filter | 
 **team** | **optional.String**| Team filter (required if no year specified) | 
 **opponent** | **optional.String**| Opponent filter | 
 **excludeGarbageTime** | **optional.Bool**| Filter to remove garbage time plays from calculations | 
 **seasonType** | **optional.String**| Season type filter (regular, postseason, or both) | 

### Return type

[**[]AdvancedGameStat**](AdvancedGameStat.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAdvancedTeamSeasonStats**
> []AdvancedSeasonStat GetAdvancedTeamSeasonStats(ctx, optional)
Advanced team metrics by season

Advanced team season stats

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***StatsApiGetAdvancedTeamSeasonStatsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StatsApiGetAdvancedTeamSeasonStatsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **year** | **optional.Int32**| Year filter (required if no team specified) | 
 **team** | **optional.String**| Team filter (required if no year specified) | 
 **excludeGarbageTime** | **optional.Bool**| Filter to remove garbage time plays from calculations | 
 **startWeek** | **optional.Int32**| Starting week filter | 
 **endWeek** | **optional.Int32**| Starting week filter | 

### Return type

[**[]AdvancedSeasonStat**](AdvancedSeasonStat.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStatCategories**
> []string GetStatCategories(ctx, )
Team stat categories

Stat category list

### Required Parameters
This endpoint does not need any parameter.

### Return type

**[]string**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTeamSeasonStats**
> []TeamSeasonStat GetTeamSeasonStats(ctx, optional)
Team statistics by season

Team season stats

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***StatsApiGetTeamSeasonStatsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StatsApiGetTeamSeasonStatsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **year** | **optional.Int32**| Year filter (required if no team specified) | 
 **team** | **optional.String**| Team filter (required if no year specified) | 
 **conference** | **optional.String**| Conference abbreviation filter | 
 **startWeek** | **optional.Int32**| Starting week filter | 
 **endWeek** | **optional.Int32**| Starting week filter | 

### Return type

[**[]TeamSeasonStat**](TeamSeasonStat.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

