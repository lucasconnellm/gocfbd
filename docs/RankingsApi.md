# {{classname}}

All URIs are relative to *https://api.collegefootballdata.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRankings**](RankingsApi.md#GetRankings) | **Get** /rankings | Historical polls and rankings

# **GetRankings**
> []RankingWeek GetRankings(ctx, year, optional)
Historical polls and rankings

Poll rankings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **year** | **int32**| Year/season filter for games | 
 **optional** | ***RankingsApiGetRankingsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RankingsApiGetRankingsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **week** | **optional.Int32**| Week filter | 
 **seasonType** | **optional.String**| Season type filter (regular or postseason) | [default to regular]

### Return type

[**[]RankingWeek**](RankingWeek.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

