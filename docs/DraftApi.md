# {{classname}}

All URIs are relative to *https://api.collegefootballdata.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDraftPicks**](DraftApi.md#GetDraftPicks) | **Get** /draft/picks | List of NFL Draft picks
[**GetNFLPositions**](DraftApi.md#GetNFLPositions) | **Get** /draft/positions | List of NFL positions
[**GetNFLTeams**](DraftApi.md#GetNFLTeams) | **Get** /draft/teams | List of NFL teams

# **GetDraftPicks**
> []DraftPick GetDraftPicks(ctx, optional)
List of NFL Draft picks

List of NFL Draft picks

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DraftApiGetDraftPicksOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DraftApiGetDraftPicksOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **year** | **optional.Int32**| Year filter | 
 **nflTeam** | **optional.String**| NFL team filter | 
 **college** | **optional.String**| Player college filter | 
 **conference** | **optional.String**| College confrence abbreviation filter | 
 **position** | **optional.String**| NFL position filter | 

### Return type

[**[]DraftPick**](DraftPick.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNFLPositions**
> []DraftPosition GetNFLPositions(ctx, )
List of NFL positions

List of NFL positions

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]DraftPosition**](DraftPosition.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNFLTeams**
> []DraftTeam GetNFLTeams(ctx, )
List of NFL teams

List of NFL teams

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]DraftTeam**](DraftTeam.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

