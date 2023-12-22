/*
Dynatrace Environment API

Documentation of the Dynatrace Environment API v1. To read about use cases and examples, see [Dynatrace Documentation](https://dt-url.net/xc03k3c).  Notes about compatibility: * Operations marked as early adopter or preview may be changed in non-compatible ways, although we try to avoid this. * We may add new enum constants without incrementing the API version; thus, clients need to handle unknown enum constants gracefully.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package environment

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"
)

type TopologySmartscapeProcessAPI interface {

	/*
			GetProcesses Lists all monitored processes along with their parameters

			You can narrow down the output by specifying filtering parameters for the request.

		You can additionally limit the output by using pagination:
		1. Specify the number of results per page in the **pageSize** query parameter.
		2. Then use the URL-encoded cursor from the **Next-Page-Key** response header in the **nextPageKey** query parameter to obtain subsequent pages.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return ApiGetProcessesRequest

			Deprecated
	*/
	GetProcesses(ctx context.Context) ApiGetProcessesRequest

	// GetProcessesExecute executes the request
	//  @return []ProcessGroupInstance
	// Deprecated
	GetProcessesExecute(r ApiGetProcessesRequest) ([]ProcessGroupInstance, *http.Response, error)

	/*
		GetSingleProcess List properties of the specified process

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param meIdentifier The Dynatrace entity ID of the required process.
		@return ApiGetSingleProcessRequest

		Deprecated
	*/
	GetSingleProcess(ctx context.Context, meIdentifier string) ApiGetSingleProcessRequest

	// GetSingleProcessExecute executes the request
	//  @return ProcessGroupInstance
	// Deprecated
	GetSingleProcessExecute(r ApiGetSingleProcessRequest) (*ProcessGroupInstance, *http.Response, error)
}

// TopologySmartscapeProcessAPIService TopologySmartscapeProcessAPI service
type TopologySmartscapeProcessAPIService service

type ApiGetProcessesRequest struct {
	ctx                     context.Context
	ApiService              TopologySmartscapeProcessAPI
	startTimestamp          *int64
	endTimestamp            *int64
	relativeTime            *string
	tag                     *[]string
	entity                  *[]string
	hostTag                 *[]string
	host                    *[]string
	actualMonitoringState   *string
	expectedMonitoringState *string
	managementZone          *int64
	includeDetails          *bool
	pageSize                *int32
	nextPageKey             *string
}

// The start timestamp of the requested timeframe, in milliseconds (UTC).   If not set, then 72 hours behind from now is used.
func (r ApiGetProcessesRequest) StartTimestamp(startTimestamp int64) ApiGetProcessesRequest {
	r.startTimestamp = &startTimestamp
	return r
}

// The end timestamp of the requested timeframe, in milliseconds (UTC).   If not set, then the current timestamp is used.   The timeframe must not exceed 3 days.
func (r ApiGetProcessesRequest) EndTimestamp(endTimestamp int64) ApiGetProcessesRequest {
	r.endTimestamp = &endTimestamp
	return r
}

// The relative timeframe, back from now.
func (r ApiGetProcessesRequest) RelativeTime(relativeTime string) ApiGetProcessesRequest {
	r.relativeTime = &relativeTime
	return r
}

// Filters the resulting set of processes by the specified tag. You can specify several tags in the following format: &#x60;tag&#x3D;tag1&amp;tag&#x3D;tag2&#x60;. The process has to match **all** the specified tags.   In case of key-value tags, such as imported AWS or CloudFoundry tags, use the following format: &#x60;tag&#x3D;[context]key:value&#x60;. For custom key-value tags, omit the context: &#x60;tag&#x3D;key:value&#x60;.
func (r ApiGetProcessesRequest) Tag(tag []string) ApiGetProcessesRequest {
	r.tag = &tag
	return r
}

// Filters result to the specified processes only.    To specify several processes use the following format: &#x60;entity&#x3D;ID1&amp;entity&#x3D;ID2&#x60;.
func (r ApiGetProcessesRequest) Entity(entity []string) ApiGetProcessesRequest {
	r.entity = &entity
	return r
}

// Filters processes by the host they&#39;re running at.    Specify tags of the host you&#39;re interested in.
func (r ApiGetProcessesRequest) HostTag(hostTag []string) ApiGetProcessesRequest {
	r.hostTag = &hostTag
	return r
}

// Filters processes by the host they&#39;re running at.    Specify Dynatrace IDs of the host you&#39;re interested in.    To specify several hosts use the following format: &#x60;host&#x3D;hostID1&amp;host&#x3D;hostID2&#x60;.    The **OR** logic applies.
func (r ApiGetProcessesRequest) Host(host []string) ApiGetProcessesRequest {
	r.host = &host
	return r
}

// Filters processes by the actual monitoring state of the process.
func (r ApiGetProcessesRequest) ActualMonitoringState(actualMonitoringState string) ApiGetProcessesRequest {
	r.actualMonitoringState = &actualMonitoringState
	return r
}

// Filters processes by the expected monitoring state of the process.
func (r ApiGetProcessesRequest) ExpectedMonitoringState(expectedMonitoringState string) ApiGetProcessesRequest {
	r.expectedMonitoringState = &expectedMonitoringState
	return r
}

// Only return processes that are part of the specified management zone.
func (r ApiGetProcessesRequest) ManagementZone(managementZone int64) ApiGetProcessesRequest {
	r.managementZone = &managementZone
	return r
}

// Includes (&#x60;true&#x60;) or excludes (&#x60;false&#x60;) details which are queried from related entities.  Excluding details may make queries faster.   If not set, then &#x60;true&#x60; is used.
func (r ApiGetProcessesRequest) IncludeDetails(includeDetails bool) ApiGetProcessesRequest {
	r.includeDetails = &includeDetails
	return r
}

// The number of processes per result page.    If not set, pagination is not used and the result contains all processes fitting the specified filtering criteria.
func (r ApiGetProcessesRequest) PageSize(pageSize int32) ApiGetProcessesRequest {
	r.pageSize = &pageSize
	return r
}

// The cursor for the next page of results. You can find it in the **Next-Page-Key** header of the previous response.   If you&#39;re using pagination, the first page is always returned without this cursor.   You must keep all other query parameters as they were in the first request to obtain subsequent pages.
func (r ApiGetProcessesRequest) NextPageKey(nextPageKey string) ApiGetProcessesRequest {
	r.nextPageKey = &nextPageKey
	return r
}

func (r ApiGetProcessesRequest) Execute() ([]ProcessGroupInstance, *http.Response, error) {
	return r.ApiService.GetProcessesExecute(r)
}

/*
GetProcesses Lists all monitored processes along with their parameters

You can narrow down the output by specifying filtering parameters for the request.

You can additionally limit the output by using pagination:
1. Specify the number of results per page in the **pageSize** query parameter.
2. Then use the URL-encoded cursor from the **Next-Page-Key** response header in the **nextPageKey** query parameter to obtain subsequent pages.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetProcessesRequest

Deprecated
*/
func (a *TopologySmartscapeProcessAPIService) GetProcesses(ctx context.Context) ApiGetProcessesRequest {
	return ApiGetProcessesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return []ProcessGroupInstance
//
// Deprecated
func (a *TopologySmartscapeProcessAPIService) GetProcessesExecute(r ApiGetProcessesRequest) ([]ProcessGroupInstance, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []ProcessGroupInstance
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TopologySmartscapeProcessAPIService.GetProcesses")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/entity/infrastructure/processes"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.startTimestamp != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "startTimestamp", r.startTimestamp, "")
	}
	if r.endTimestamp != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "endTimestamp", r.endTimestamp, "")
	}
	if r.relativeTime != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "relativeTime", r.relativeTime, "")
	}
	if r.tag != nil {
		t := *r.tag
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "tag", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "tag", t, "multi")
		}
	}
	if r.entity != nil {
		t := *r.entity
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "entity", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "entity", t, "multi")
		}
	}
	if r.hostTag != nil {
		t := *r.hostTag
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "hostTag", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "hostTag", t, "multi")
		}
	}
	if r.host != nil {
		t := *r.host
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "host", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "host", t, "multi")
		}
	}
	if r.actualMonitoringState != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "actualMonitoringState", r.actualMonitoringState, "")
	}
	if r.expectedMonitoringState != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "expectedMonitoringState", r.expectedMonitoringState, "")
	}
	if r.managementZone != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "managementZone", r.managementZone, "")
	}
	if r.includeDetails != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeDetails", r.includeDetails, "")
	} else {
		var defaultValue bool = true
		r.includeDetails = &defaultValue
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageSize", r.pageSize, "")
	}
	if r.nextPageKey != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "nextPageKey", r.nextPageKey, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json; charset=utf-8", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Api-Token"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorEnvelope
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode >= 400 && localVarHTTPResponse.StatusCode < 500 {
			var v ErrorEnvelope
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode >= 500 {
			var v ErrorEnvelope
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetSingleProcessRequest struct {
	ctx          context.Context
	ApiService   TopologySmartscapeProcessAPI
	meIdentifier string
}

func (r ApiGetSingleProcessRequest) Execute() (*ProcessGroupInstance, *http.Response, error) {
	return r.ApiService.GetSingleProcessExecute(r)
}

/*
GetSingleProcess List properties of the specified process

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param meIdentifier The Dynatrace entity ID of the required process.
	@return ApiGetSingleProcessRequest

Deprecated
*/
func (a *TopologySmartscapeProcessAPIService) GetSingleProcess(ctx context.Context, meIdentifier string) ApiGetSingleProcessRequest {
	return ApiGetSingleProcessRequest{
		ApiService:   a,
		ctx:          ctx,
		meIdentifier: meIdentifier,
	}
}

// Execute executes the request
//
//	@return ProcessGroupInstance
//
// Deprecated
func (a *TopologySmartscapeProcessAPIService) GetSingleProcessExecute(r ApiGetSingleProcessRequest) (*ProcessGroupInstance, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ProcessGroupInstance
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TopologySmartscapeProcessAPIService.GetSingleProcess")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/entity/infrastructure/processes/{meIdentifier}"
	localVarPath = strings.Replace(localVarPath, "{"+"meIdentifier"+"}", url.PathEscape(parameterValueToString(r.meIdentifier, "meIdentifier")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json; charset=utf-8", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Api-Token"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode >= 400 && localVarHTTPResponse.StatusCode < 500 {
			var v ErrorEnvelope
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode >= 500 {
			var v ErrorEnvelope
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
