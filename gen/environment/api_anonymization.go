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

// AnonymizationAPIService AnonymizationAPI service
type AnonymizationAPIService service

type ApiAnonymizeRequest struct {
	ctx             context.Context
	ApiService      *AnonymizationAPIService
	startTimestamp  *int64
	endTimestamp    *int64
	userIds         *[]string
	ips             *[]string
	additionalField *[]string
}

// The start timestamp of the user session to anonymize, in UTC milliseconds.   If not set the earliest available time is used.
func (r ApiAnonymizeRequest) StartTimestamp(startTimestamp int64) ApiAnonymizeRequest {
	r.startTimestamp = &startTimestamp
	return r
}

// The end timestamp of the user session to anonymize, in UTC milliseconds.   If not set the current time is used.
func (r ApiAnonymizeRequest) EndTimestamp(endTimestamp int64) ApiAnonymizeRequest {
	r.endTimestamp = &endTimestamp
	return r
}

// The UserID of the user to anonymize.    You can specify several IDs, in the following format: &#x60;userIds&#x3D;user1&amp;userIds&#x3D;user2&#x60;.
func (r ApiAnonymizeRequest) UserIds(userIds []string) ApiAnonymizeRequest {
	r.userIds = &userIds
	return r
}

// The IP address of the user to anonymize. All user sessions from this IP will be anonymized.   You can specify several IPs, in the following format: &#x60;ips&#x3D;ip1&amp;ips&#x3D;ip2&#x60;.
func (r ApiAnonymizeRequest) Ips(ips []string) ApiAnonymizeRequest {
	r.ips = &ips
	return r
}

// A list of fields to be anonymized.   You can specify several fields, in the following format: &#x60;additionalField&#x3D;field1&amp;additionalField&#x3D;field2&#x60;.
func (r ApiAnonymizeRequest) AdditionalField(additionalField []string) ApiAnonymizeRequest {
	r.additionalField = &additionalField
	return r
}

func (r ApiAnonymizeRequest) Execute() (*AnonymizationIdResult, *http.Response, error) {
	return r.ApiService.AnonymizeExecute(r)
}

/*
Anonymize Creates user session anonymization job

The job anonymizes all user sessions in the specified timeframe by masking the specified fields.

To identify user sessions to be anonymized you can specify either userID, or IP address, or both. If you specify both the **OR** logic applies.

You can't undo the anonymization.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAnonymizeRequest
*/
func (a *AnonymizationAPIService) Anonymize(ctx context.Context) ApiAnonymizeRequest {
	return ApiAnonymizeRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AnonymizationIdResult
func (a *AnonymizationAPIService) AnonymizeExecute(r ApiAnonymizeRequest) (*AnonymizationIdResult, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AnonymizationIdResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AnonymizationAPIService.Anonymize")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/anonymize/anonymizationJobs"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.startTimestamp != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "startTimestamp", r.startTimestamp, "")
	} else {
		var defaultValue int64 = 0
		r.startTimestamp = &defaultValue
	}
	if r.endTimestamp != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "endTimestamp", r.endTimestamp, "")
	} else {
		var defaultValue int64 = 0
		r.endTimestamp = &defaultValue
	}
	if r.userIds != nil {
		t := *r.userIds
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "userIds", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "userIds", t, "multi")
		}
	}
	if r.ips != nil {
		t := *r.ips
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "ips", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "ips", t, "multi")
		}
	}
	if r.additionalField != nil {
		t := *r.additionalField
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "additionalField", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "additionalField", t, "multi")
		}
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

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

type ApiGetStatusRequest struct {
	ctx        context.Context
	ApiService *AnonymizationAPIService
	requestId  string
}

func (r ApiGetStatusRequest) Execute() (*AnonymizationProgressResult, *http.Response, error) {
	return r.ApiService.GetStatusExecute(r)
}

/*
GetStatus Shows the progress of the specified anonymization job

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param requestId The ID of the required anonymization job.
	@return ApiGetStatusRequest
*/
func (a *AnonymizationAPIService) GetStatus(ctx context.Context, requestId string) ApiGetStatusRequest {
	return ApiGetStatusRequest{
		ApiService: a,
		ctx:        ctx,
		requestId:  requestId,
	}
}

// Execute executes the request
//
//	@return AnonymizationProgressResult
func (a *AnonymizationAPIService) GetStatusExecute(r ApiGetStatusRequest) (*AnonymizationProgressResult, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AnonymizationProgressResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AnonymizationAPIService.GetStatus")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/anonymize/anonymizationJobs/{requestId}"
	localVarPath = strings.Replace(localVarPath, "{"+"requestId"+"}", url.PathEscape(parameterValueToString(r.requestId, "requestId")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/json"}

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
