# {{classname}}

All URIs are relative to *https://api.collegefootballdata.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLines**](BettingApi.md#GetLines) | **Get** /lines | Betting lines

# **GetLines**
> []GameLines GetLines(ctx, optional)
Betting lines

Closing betting lines

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BettingApiGetLinesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BettingApiGetLinesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gameId** | **optional.Int32**| Game id filter | 
 **year** | **optional.Int32**| Year/season filter for games | 
 **week** | **optional.Int32**| Week filter | 
 **seasonType** | **optional.String**| Season type filter (regular or postseason) | [default to regular]
 **team** | **optional.String**| Team | 
 **home** | **optional.String**| Home team filter | 
 **away** | **optional.String**| Away team filter | 
 **conference** | **optional.String**| Conference abbreviation filter | 

### Return type

[**[]GameLines**](GameLines.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

