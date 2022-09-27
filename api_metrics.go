
/*
 * College Football Data API
 *
 * This is an API for accessing all sorts of college football data.  Please note that API keys should be supplied with \"Bearer \" prepended (e.g. \"Bearer your_key\"). API keys can be acquired from the CollegeFootballData.com website.
 *
 * API version: 4.4.8
 * Contact: admin@collegefootballdata.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"github.com/antihax/optional"
)

// Linger please
var (
	_ context.Context
)

type MetricsApiService service
/*
MetricsApiService Team Predicated Points Added (PPA/EPA) by game
Predicted Points Added (PPA) by game
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param year Year filter
 * @param optional nil or *MetricsApiGetGamePPAOpts - Optional Parameters:
     * @param "Week" (optional.Int32) -  Week filter
     * @param "Team" (optional.String) -  Team filter
     * @param "Conference" (optional.String) -  Conference filter
     * @param "ExcludeGarbageTime" (optional.Bool) -  Filter to remove garbage time plays from calculations
     * @param "SeasonType" (optional.String) -  Season type filter (regular or postseason)
@return []GamePpa
*/

type MetricsApiGetGamePPAOpts struct {
    Week optional.Int32
    Team optional.String
    Conference optional.String
    ExcludeGarbageTime optional.Bool
    SeasonType optional.String
}

func (a *MetricsApiService) GetGamePPA(ctx context.Context, year int32, localVarOptionals *MetricsApiGetGamePPAOpts) ([]GamePpa, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue []GamePpa
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/ppa/games"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if year < 2001 {
		return localVarReturnValue, nil, reportError("year must be greater than 2001")
	}

	localVarQueryParams.Add("year", parameterToString(year, ""))
	if localVarOptionals != nil && localVarOptionals.Week.IsSet() {
		localVarQueryParams.Add("week", parameterToString(localVarOptionals.Week.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Team.IsSet() {
		localVarQueryParams.Add("team", parameterToString(localVarOptionals.Team.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Conference.IsSet() {
		localVarQueryParams.Add("conference", parameterToString(localVarOptionals.Conference.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ExcludeGarbageTime.IsSet() {
		localVarQueryParams.Add("excludeGarbageTime", parameterToString(localVarOptionals.ExcludeGarbageTime.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SeasonType.IsSet() {
		localVarQueryParams.Add("seasonType", parameterToString(localVarOptionals.SeasonType.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorization"] = key
			
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v []GamePpa
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
MetricsApiService Player Predicated Points Added (PPA/EPA) broken down by game
Predicted Points Added (PPA) by player game
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *MetricsApiGetPlayerGamePPAOpts - Optional Parameters:
     * @param "Year" (optional.Int32) -  Year filter
     * @param "Week" (optional.Int32) -  Week filter
     * @param "Team" (optional.String) -  Team filter
     * @param "Position" (optional.String) -  Position abbreviation filter
     * @param "PlayerId" (optional.Int32) -  Player id filter
     * @param "Threshold" (optional.String) -  Minimum play threshold filter
     * @param "ExcludeGarbageTime" (optional.Bool) -  Filter to remove garbage time plays from calculations
     * @param "SeasonType" (optional.String) -  Season type filter (regular or postseason)
@return []PlayerGamePpa
*/

type MetricsApiGetPlayerGamePPAOpts struct {
    Year optional.Int32
    Week optional.Int32
    Team optional.String
    Position optional.String
    PlayerId optional.Int32
    Threshold optional.String
    ExcludeGarbageTime optional.Bool
    SeasonType optional.String
}

func (a *MetricsApiService) GetPlayerGamePPA(ctx context.Context, localVarOptionals *MetricsApiGetPlayerGamePPAOpts) ([]PlayerGamePpa, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue []PlayerGamePpa
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/ppa/players/games"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Year.IsSet() {
		localVarQueryParams.Add("year", parameterToString(localVarOptionals.Year.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Week.IsSet() {
		localVarQueryParams.Add("week", parameterToString(localVarOptionals.Week.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Team.IsSet() {
		localVarQueryParams.Add("team", parameterToString(localVarOptionals.Team.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Position.IsSet() {
		localVarQueryParams.Add("position", parameterToString(localVarOptionals.Position.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PlayerId.IsSet() {
		localVarQueryParams.Add("playerId", parameterToString(localVarOptionals.PlayerId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Threshold.IsSet() {
		localVarQueryParams.Add("threshold", parameterToString(localVarOptionals.Threshold.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ExcludeGarbageTime.IsSet() {
		localVarQueryParams.Add("excludeGarbageTime", parameterToString(localVarOptionals.ExcludeGarbageTime.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SeasonType.IsSet() {
		localVarQueryParams.Add("seasonType", parameterToString(localVarOptionals.SeasonType.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorization"] = key
			
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v []PlayerGamePpa
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
MetricsApiService Player Predicated Points Added (PPA/EPA) broken down by season
Predicted Points Added (PPA) by player season
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *MetricsApiGetPlayerSeasonPPAOpts - Optional Parameters:
     * @param "Year" (optional.Int32) -  Year filter
     * @param "Team" (optional.String) -  Team filter
     * @param "Conference" (optional.String) -  Conference abbreviation filter
     * @param "Position" (optional.String) -  Position abbreviation filter
     * @param "PlayerId" (optional.Int32) -  Player id filter
     * @param "Threshold" (optional.String) -  Minimum play threshold filter
     * @param "ExcludeGarbageTime" (optional.Bool) -  Filter to remove garbage time plays from calculations
@return []PlayerSeasonPpa
*/

type MetricsApiGetPlayerSeasonPPAOpts struct {
    Year optional.Int32
    Team optional.String
    Conference optional.String
    Position optional.String
    PlayerId optional.Int32
    Threshold optional.String
    ExcludeGarbageTime optional.Bool
}

func (a *MetricsApiService) GetPlayerSeasonPPA(ctx context.Context, localVarOptionals *MetricsApiGetPlayerSeasonPPAOpts) ([]PlayerSeasonPpa, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue []PlayerSeasonPpa
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/ppa/players/season"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Year.IsSet() {
		localVarQueryParams.Add("year", parameterToString(localVarOptionals.Year.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Team.IsSet() {
		localVarQueryParams.Add("team", parameterToString(localVarOptionals.Team.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Conference.IsSet() {
		localVarQueryParams.Add("conference", parameterToString(localVarOptionals.Conference.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Position.IsSet() {
		localVarQueryParams.Add("position", parameterToString(localVarOptionals.Position.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PlayerId.IsSet() {
		localVarQueryParams.Add("playerId", parameterToString(localVarOptionals.PlayerId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Threshold.IsSet() {
		localVarQueryParams.Add("threshold", parameterToString(localVarOptionals.Threshold.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ExcludeGarbageTime.IsSet() {
		localVarQueryParams.Add("excludeGarbageTime", parameterToString(localVarOptionals.ExcludeGarbageTime.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorization"] = key
			
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v []PlayerSeasonPpa
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
MetricsApiService Predicted Points (i.e. Expected Points or EP)
Predicted Points
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param down Down filter
 * @param distance Distance filter
@return []PredictedPoints
*/
func (a *MetricsApiService) GetPredictedPoints(ctx context.Context, down int32, distance int32) ([]PredictedPoints, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue []PredictedPoints
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/ppa/predicted"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if down < 1 {
		return localVarReturnValue, nil, reportError("down must be greater than 1")
	}
	if down > 4 {
		return localVarReturnValue, nil, reportError("down must be less than 4")
	}
	if distance < 1 {
		return localVarReturnValue, nil, reportError("distance must be greater than 1")
	}
	if distance > 99 {
		return localVarReturnValue, nil, reportError("distance must be less than 99")
	}

	localVarQueryParams.Add("down", parameterToString(down, ""))
	localVarQueryParams.Add("distance", parameterToString(distance, ""))
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorization"] = key
			
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v []PredictedPoints
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
MetricsApiService Pregame win probability data
Pregame win probabilities
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *MetricsApiGetPregameWinProbabilitiesOpts - Optional Parameters:
     * @param "Year" (optional.Int32) -  Year filter
     * @param "Week" (optional.Int32) -  Week filter
     * @param "Team" (optional.String) -  Team filter
     * @param "SeasonType" (optional.String) -  regular or postseason
@return []PregameWp
*/

type MetricsApiGetPregameWinProbabilitiesOpts struct {
    Year optional.Int32
    Week optional.Int32
    Team optional.String
    SeasonType optional.String
}

func (a *MetricsApiService) GetPregameWinProbabilities(ctx context.Context, localVarOptionals *MetricsApiGetPregameWinProbabilitiesOpts) ([]PregameWp, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue []PregameWp
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/metrics/wp/pregame"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Year.IsSet() {
		localVarQueryParams.Add("year", parameterToString(localVarOptionals.Year.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Week.IsSet() {
		localVarQueryParams.Add("week", parameterToString(localVarOptionals.Week.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Team.IsSet() {
		localVarQueryParams.Add("team", parameterToString(localVarOptionals.Team.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SeasonType.IsSet() {
		localVarQueryParams.Add("seasonType", parameterToString(localVarOptionals.SeasonType.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorization"] = key
			
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v []PregameWp
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
MetricsApiService Predicted Points Added (PPA/EPA) data by team
Predicted Points Added (PPA)
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *MetricsApiGetTeamPPAOpts - Optional Parameters:
     * @param "Year" (optional.Int32) -  Year filter (required if team not specified)
     * @param "Team" (optional.String) -  Team filter (required if year not specified)
     * @param "Conference" (optional.String) -  Conference filter
     * @param "ExcludeGarbageTime" (optional.Bool) -  Filter to remove garbage time plays from calculations
@return []TeamPpa
*/

type MetricsApiGetTeamPPAOpts struct {
    Year optional.Int32
    Team optional.String
    Conference optional.String
    ExcludeGarbageTime optional.Bool
}

func (a *MetricsApiService) GetTeamPPA(ctx context.Context, localVarOptionals *MetricsApiGetTeamPPAOpts) ([]TeamPpa, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue []TeamPpa
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/ppa/teams"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Year.IsSet() {
		localVarQueryParams.Add("year", parameterToString(localVarOptionals.Year.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Team.IsSet() {
		localVarQueryParams.Add("team", parameterToString(localVarOptionals.Team.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Conference.IsSet() {
		localVarQueryParams.Add("conference", parameterToString(localVarOptionals.Conference.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ExcludeGarbageTime.IsSet() {
		localVarQueryParams.Add("excludeGarbageTime", parameterToString(localVarOptionals.ExcludeGarbageTime.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorization"] = key
			
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v []TeamPpa
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
MetricsApiService Win probability chart data
Win probability data
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param gameId Game id filter
@return []PlayWp
*/
func (a *MetricsApiService) GetWinProbabilityData(ctx context.Context, gameId int32) ([]PlayWp, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue []PlayWp
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/metrics/wp"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("gameId", parameterToString(gameId, ""))
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorization"] = key
			
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v []PlayWp
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
