# {{classname}}

All URIs are relative to *https://api.collegefootballdata.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCoaches**](CoachesApi.md#GetCoaches) | **Get** /coaches | Coaching records and history

# **GetCoaches**
> []Coach GetCoaches(ctx, optional)
Coaching records and history

Coaching history

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CoachesApiGetCoachesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CoachesApiGetCoachesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **firstName** | **optional.String**| First name filter | 
 **lastName** | **optional.String**| Last name filter | 
 **team** | **optional.String**| Team name filter | 
 **year** | **optional.Int32**| Year filter | 
 **minYear** | **optional.Int32**| Minimum year filter (inclusive) | 
 **maxYear** | **optional.Int32**| Maximum year filter (inclusive) | 

### Return type

[**[]Coach**](Coach.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

