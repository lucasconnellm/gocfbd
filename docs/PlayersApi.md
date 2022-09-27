# {{classname}}

All URIs are relative to *https://api.collegefootballdata.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPlayerSeasonStats**](PlayersApi.md#GetPlayerSeasonStats) | **Get** /stats/player/season | Player stats by season
[**GetPlayerUsage**](PlayersApi.md#GetPlayerUsage) | **Get** /player/usage | Player usage metrics broken down by season
[**GetReturningProduction**](PlayersApi.md#GetReturningProduction) | **Get** /player/returning | Team returning production metrics
[**GetTransferPortal**](PlayersApi.md#GetTransferPortal) | **Get** /player/portal | Transfer portal by season
[**PlayerSearch**](PlayersApi.md#PlayerSearch) | **Get** /player/search | Search for player information

# **GetPlayerSeasonStats**
> []PlayerSeasonStat GetPlayerSeasonStats(ctx, year, optional)
Player stats by season

Season player stats

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **year** | **int32**| Year filter | 
 **optional** | ***PlayersApiGetPlayerSeasonStatsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PlayersApiGetPlayerSeasonStatsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **team** | **optional.String**| Team filter | 
 **conference** | **optional.String**| Conference abbreviation filter | 
 **startWeek** | **optional.Int32**| Start week filter | 
 **endWeek** | **optional.Int32**| Start week filter | 
 **seasonType** | **optional.String**| Season type filter (regular, postseason, or both) | 
 **category** | **optional.String**| Stat category filter (e.g. passing) | 

### Return type

[**[]PlayerSeasonStat**](PlayerSeasonStat.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPlayerUsage**
> []PlayerUsage GetPlayerUsage(ctx, year, optional)
Player usage metrics broken down by season

Player usage metrics by season

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **year** | **int32**| Year filter | [default to 2020]
 **optional** | ***PlayersApiGetPlayerUsageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PlayersApiGetPlayerUsageOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **team** | **optional.String**| Team filter | 
 **conference** | **optional.String**| Conference abbreviation filter | 
 **position** | **optional.String**| Position abbreviation filter | 
 **playerId** | **optional.Int32**| Player id filter | 
 **excludeGarbageTime** | **optional.Bool**| Filter to remove garbage time plays from calculations | 

### Return type

[**[]PlayerUsage**](PlayerUsage.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReturningProduction**
> []ReturningProduction GetReturningProduction(ctx, optional)
Team returning production metrics

Returning production metrics

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PlayersApiGetReturningProductionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PlayersApiGetReturningProductionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **year** | **optional.Int32**| Year filter | 
 **team** | **optional.String**| Team filter | 
 **conference** | **optional.String**| Conference abbreviation filter | 

### Return type

[**[]ReturningProduction**](ReturningProduction.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTransferPortal**
> []PortalPlayer GetTransferPortal(ctx, year)
Transfer portal by season

Transfer portal by season

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **year** | **int32**| Year filter | 

### Return type

[**[]PortalPlayer**](PortalPlayer.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PlayerSearch**
> []PlayerSearchResult PlayerSearch(ctx, searchTerm, optional)
Search for player information

Search for players

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **searchTerm** | **string**| Term to search on | 
 **optional** | ***PlayersApiPlayerSearchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PlayersApiPlayerSearchOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **position** | **optional.String**| Position abbreviation filter | 
 **team** | **optional.String**| Team filter | 
 **year** | **optional.Int32**| Year filter | 

### Return type

[**[]PlayerSearchResult**](PlayerSearchResult.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

