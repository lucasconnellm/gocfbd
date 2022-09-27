# {{classname}}

All URIs are relative to *https://api.collegefootballdata.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRecruitingGroups**](RecruitingApi.md#GetRecruitingGroups) | **Get** /recruiting/groups | Recruit position group ratings
[**GetRecruitingPlayers**](RecruitingApi.md#GetRecruitingPlayers) | **Get** /recruiting/players | Player recruiting ratings and rankings
[**GetRecruitingTeams**](RecruitingApi.md#GetRecruitingTeams) | **Get** /recruiting/teams | Team recruiting rankings and ratings

# **GetRecruitingGroups**
> []PositionGroupRecruitingRating GetRecruitingGroups(ctx, optional)
Recruit position group ratings

Gets a list of aggregated statistics by team and position grouping

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RecruitingApiGetRecruitingGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RecruitingApiGetRecruitingGroupsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startYear** | **optional.Int32**| Starting year | 
 **endYear** | **optional.Int32**| Ending year | 
 **team** | **optional.String**| Team filter | 
 **conference** | **optional.String**| conference filter | 

### Return type

[**[]PositionGroupRecruitingRating**](PositionGroupRecruitingRating.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRecruitingPlayers**
> []Recruit GetRecruitingPlayers(ctx, optional)
Player recruiting ratings and rankings

Get player recruiting rankings and data. Requires either a year or team to be specified.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RecruitingApiGetRecruitingPlayersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RecruitingApiGetRecruitingPlayersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **year** | **optional.Int32**| Recruiting class year (required if team no specified) | 
 **classification** | **optional.String**| Type of recruit (HighSchool, JUCO, PrepSchool) | [default to HighSchool]
 **position** | **optional.String**| Position abbreviation filter | 
 **state** | **optional.String**| State or province abbreviation filter | 
 **team** | **optional.String**| Committed team filter (required if year not specified) | 

### Return type

[**[]Recruit**](Recruit.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRecruitingTeams**
> []TeamRecruitingRank GetRecruitingTeams(ctx, optional)
Team recruiting rankings and ratings

Team recruiting rankings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RecruitingApiGetRecruitingTeamsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RecruitingApiGetRecruitingTeamsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **year** | **optional.Int32**| Recruiting class year | 
 **team** | **optional.String**| Team filter | 

### Return type

[**[]TeamRecruitingRank**](TeamRecruitingRank.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

