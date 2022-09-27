# {{classname}}

All URIs are relative to *https://api.collegefootballdata.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFbsTeams**](TeamsApi.md#GetFbsTeams) | **Get** /teams/fbs | FBS team list
[**GetRoster**](TeamsApi.md#GetRoster) | **Get** /roster | Team rosters
[**GetTalent**](TeamsApi.md#GetTalent) | **Get** /talent | Team talent composite rankings
[**GetTeamMatchup**](TeamsApi.md#GetTeamMatchup) | **Get** /teams/matchup | Team matchup history
[**GetTeams**](TeamsApi.md#GetTeams) | **Get** /teams | Team information

# **GetFbsTeams**
> []Team GetFbsTeams(ctx, optional)
FBS team list

Information on major division teams

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TeamsApiGetFbsTeamsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TeamsApiGetFbsTeamsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **year** | **optional.Int32**| Year filter | 

### Return type

[**[]Team**](Team.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRoster**
> []Player GetRoster(ctx, optional)
Team rosters

Roster data

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TeamsApiGetRosterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TeamsApiGetRosterOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **team** | **optional.String**| Team name | 
 **year** | **optional.Int32**| Season year | 

### Return type

[**[]Player**](Player.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTalent**
> []TeamTalent GetTalent(ctx, optional)
Team talent composite rankings

Team talent composite

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TeamsApiGetTalentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TeamsApiGetTalentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **year** | **optional.Int32**| Year filter | 

### Return type

[**[]TeamTalent**](TeamTalent.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTeamMatchup**
> TeamMatchup GetTeamMatchup(ctx, team1, team2, optional)
Team matchup history

Matchup history

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **team1** | **string**| First team | 
  **team2** | **string**| Second team | 
 **optional** | ***TeamsApiGetTeamMatchupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TeamsApiGetTeamMatchupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **minYear** | **optional.Int32**| Minimum year | 
 **maxYear** | **optional.Int32**| Maximum year | 

### Return type

[**TeamMatchup**](TeamMatchup.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTeams**
> []Team GetTeams(ctx, optional)
Team information

Get team information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TeamsApiGetTeamsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TeamsApiGetTeamsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **conference** | **optional.String**| Conference abbreviation filter | 

### Return type

[**[]Team**](Team.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

