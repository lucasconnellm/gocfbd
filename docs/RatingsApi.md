# {{classname}}

All URIs are relative to *https://api.collegefootballdata.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetConferenceSPRatings**](RatingsApi.md#GetConferenceSPRatings) | **Get** /ratings/sp/conferences | Historical SP+ ratings by conference
[**GetEloRatings**](RatingsApi.md#GetEloRatings) | **Get** /ratings/elo | Historical Elo ratings
[**GetSPRatings**](RatingsApi.md#GetSPRatings) | **Get** /ratings/sp | Historical SP+ ratings
[**GetSRSRatings**](RatingsApi.md#GetSRSRatings) | **Get** /ratings/srs | Historical SRS ratings

# **GetConferenceSPRatings**
> []ConferenceSpRating GetConferenceSPRatings(ctx, optional)
Historical SP+ ratings by conference

Get average SP+ historical rating data by conference

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RatingsApiGetConferenceSPRatingsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RatingsApiGetConferenceSPRatingsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **year** | **optional.Int32**| Season filter | 
 **conference** | **optional.String**| Conference abbreviation filter | 

### Return type

[**[]ConferenceSpRating**](ConferenceSPRating.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEloRatings**
> []TeamEloRating GetEloRatings(ctx, optional)
Historical Elo ratings

Elo rating data

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RatingsApiGetEloRatingsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RatingsApiGetEloRatingsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **year** | **optional.Int32**| Season filter | 
 **week** | **optional.Int32**| Maximum week filter | 
 **team** | **optional.String**| Team filter | 
 **conference** | **optional.String**| Conference filter | 

### Return type

[**[]TeamEloRating**](TeamEloRating.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSPRatings**
> []TeamSpRating GetSPRatings(ctx, optional)
Historical SP+ ratings

SP+ rating data

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RatingsApiGetSPRatingsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RatingsApiGetSPRatingsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **year** | **optional.Int32**| Season filter (required if team not specified) | 
 **team** | **optional.String**| Team filter (required if year not specified) | 

### Return type

[**[]TeamSpRating**](TeamSPRating.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSRSRatings**
> []TeamSrsRating GetSRSRatings(ctx, optional)
Historical SRS ratings

SRS rating data (requires either a year or team specified)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RatingsApiGetSRSRatingsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RatingsApiGetSRSRatingsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **year** | **optional.Int32**| Season filter (required if team not specified) | 
 **team** | **optional.String**| Team filter (required if year not specified) | 
 **conference** | **optional.String**| Conference filter | 

### Return type

[**[]TeamSrsRating**](TeamSRSRating.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

