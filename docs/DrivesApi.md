# {{classname}}

All URIs are relative to *https://api.collegefootballdata.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDrives**](DrivesApi.md#GetDrives) | **Get** /drives | Drive data and results

# **GetDrives**
> []Drive GetDrives(ctx, year, optional)
Drive data and results

Get game drives

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **year** | **int32**| Year filter | 
 **optional** | ***DrivesApiGetDrivesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DrivesApiGetDrivesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **seasonType** | **optional.String**| Season type filter | [default to regular]
 **week** | **optional.Int32**| Week filter | 
 **team** | **optional.String**| Team filter | 
 **offense** | **optional.String**| Offensive team filter | 
 **defense** | **optional.String**| Defensive team filter | 
 **conference** | **optional.String**| Conference filter | 
 **offenseConference** | **optional.String**| Offensive conference filter | 
 **defenseConference** | **optional.String**| Defensive conference filter | 
 **classification** | **optional.String**| Division classification filter (fbs/fcs/ii/iii) | 

### Return type

[**[]Drive**](Drive.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

