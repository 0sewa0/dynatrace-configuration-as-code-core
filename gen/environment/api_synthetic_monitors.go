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

// SyntheticMonitorsAPIService SyntheticMonitorsAPI service
type SyntheticMonitorsAPIService service

type ApiAddMonitorRequest struct {
	ctx                    context.Context
	ApiService             *SyntheticMonitorsAPIService
	syntheticMonitorUpdate *SyntheticMonitorUpdate
}

// The JSON body of the request, containing parameters of the new synthetic monitor.
func (r ApiAddMonitorRequest) SyntheticMonitorUpdate(syntheticMonitorUpdate SyntheticMonitorUpdate) ApiAddMonitorRequest {
	r.syntheticMonitorUpdate = &syntheticMonitorUpdate
	return r
}

func (r ApiAddMonitorRequest) Execute() (*EntityIdDto, *http.Response, error) {
	return r.ApiService.AddMonitorExecute(r)
}

/*
AddMonitor Creates a new synthetic monitor

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAddMonitorRequest
*/
func (a *SyntheticMonitorsAPIService) AddMonitor(ctx context.Context) ApiAddMonitorRequest {
	return ApiAddMonitorRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return EntityIdDto
func (a *SyntheticMonitorsAPIService) AddMonitorExecute(r ApiAddMonitorRequest) (*EntityIdDto, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *EntityIdDto
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SyntheticMonitorsAPIService.AddMonitor")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/synthetic/monitors"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json; charset=utf-8"}

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
	// body params
	localVarPostBody = r.syntheticMonitorUpdate
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

type ApiDeleteMonitorRequest struct {
	ctx        context.Context
	ApiService *SyntheticMonitorsAPIService
	monitorId  string
}

func (r ApiDeleteMonitorRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteMonitorExecute(r)
}

/*
DeleteMonitor Deletes the specified synthetic monitor

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param monitorId The ID of the synthetic monitor to be deleted.
	@return ApiDeleteMonitorRequest
*/
func (a *SyntheticMonitorsAPIService) DeleteMonitor(ctx context.Context, monitorId string) ApiDeleteMonitorRequest {
	return ApiDeleteMonitorRequest{
		ApiService: a,
		ctx:        ctx,
		monitorId:  monitorId,
	}
}

// Execute executes the request
func (a *SyntheticMonitorsAPIService) DeleteMonitorExecute(r ApiDeleteMonitorRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SyntheticMonitorsAPIService.DeleteMonitor")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/synthetic/monitors/{monitorId}"
	localVarPath = strings.Replace(localVarPath, "{"+"monitorId"+"}", url.PathEscape(parameterValueToString(r.monitorId, "monitorId")), -1)

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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
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
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode >= 500 {
			var v ErrorEnvelope
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetMonitorRequest struct {
	ctx        context.Context
	ApiService *SyntheticMonitorsAPIService
	monitorId  string
}

func (r ApiGetMonitorRequest) Execute() (*SyntheticMonitor, *http.Response, error) {
	return r.ApiService.GetMonitorExecute(r)
}

/*
GetMonitor Gets parameters of the specified synthetic monitor

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param monitorId The ID of the required synthetic monitor
	@return ApiGetMonitorRequest
*/
func (a *SyntheticMonitorsAPIService) GetMonitor(ctx context.Context, monitorId string) ApiGetMonitorRequest {
	return ApiGetMonitorRequest{
		ApiService: a,
		ctx:        ctx,
		monitorId:  monitorId,
	}
}

// Execute executes the request
//
//	@return SyntheticMonitor
func (a *SyntheticMonitorsAPIService) GetMonitorExecute(r ApiGetMonitorRequest) (*SyntheticMonitor, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SyntheticMonitor
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SyntheticMonitorsAPIService.GetMonitor")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/synthetic/monitors/{monitorId}"
	localVarPath = strings.Replace(localVarPath, "{"+"monitorId"+"}", url.PathEscape(parameterValueToString(r.monitorId, "monitorId")), -1)

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

type ApiGetMonitorsCollectionRequest struct {
	ctx             context.Context
	ApiService      *SyntheticMonitorsAPIService
	managementZone  *int64
	tag             *[]string
	location        *string
	assignedApps    *[]string
	type_           *string
	enabled         *bool
	credentialId    *string
	credentialOwner *string
}

// Filters the resulting set of monitors to those which are part of the specified management zone.    Specify the ID of the management zone here.
func (r ApiGetMonitorsCollectionRequest) ManagementZone(managementZone int64) ApiGetMonitorsCollectionRequest {
	r.managementZone = &managementZone
	return r
}

// Filters the resulting set of monitors by specified tags.   You can specify several tags in the following format: &#x60;tag&#x3D;tag1&amp;tag&#x3D;tag2&#x60;. The monitor has to match **all** the specified tags.   In case of key-value tags, such as imported AWS or CloudFoundry tags use following format: &#x60;[context]key:value&#x60;.
func (r ApiGetMonitorsCollectionRequest) Tag(tag []string) ApiGetMonitorsCollectionRequest {
	r.tag = &tag
	return r
}

// Filters the resulting set of monitors to those assigned to a specified Synthetic location.     Specify the ID of the location here.
func (r ApiGetMonitorsCollectionRequest) Location(location string) ApiGetMonitorsCollectionRequest {
	r.location = &location
	return r
}

// Filters the resulting set of monitors to those assigned to the specified applications.   You can specify several applications in the following format: &#x60;assignedApps&#x3D;app1&amp;assignedApps&#x3D;app2&#x60;. The monitor has to have **all** the specified applications assigned.   Specify Dynatrace entity IDs of applications here.
func (r ApiGetMonitorsCollectionRequest) AssignedApps(assignedApps []string) ApiGetMonitorsCollectionRequest {
	r.assignedApps = &assignedApps
	return r
}

// Filters the resulting set of monitors to those of the specified type: &#x60;BROWSER&#x60; or &#x60;HTTP&#x60;.
func (r ApiGetMonitorsCollectionRequest) Type_(type_ string) ApiGetMonitorsCollectionRequest {
	r.type_ = &type_
	return r
}

// Filters the resulting set of monitors to those which are enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;).
func (r ApiGetMonitorsCollectionRequest) Enabled(enabled bool) ApiGetMonitorsCollectionRequest {
	r.enabled = &enabled
	return r
}

// Filters the resulting set of monitors to those using the specified credential set.   Specify the ID of the credentials set here.
func (r ApiGetMonitorsCollectionRequest) CredentialId(credentialId string) ApiGetMonitorsCollectionRequest {
	r.credentialId = &credentialId
	return r
}

// Filters the resulting set of monitors to those using a credential owned by the specified user.
func (r ApiGetMonitorsCollectionRequest) CredentialOwner(credentialOwner string) ApiGetMonitorsCollectionRequest {
	r.credentialOwner = &credentialOwner
	return r
}

func (r ApiGetMonitorsCollectionRequest) Execute() (*Monitors, *http.Response, error) {
	return r.ApiService.GetMonitorsCollectionExecute(r)
}

/*
GetMonitorsCollection Lists all synthetic monitors in your Dynatrace environment

The full list can be lengthy, but you can narrow it down by specifying filter query parameters.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetMonitorsCollectionRequest
*/
func (a *SyntheticMonitorsAPIService) GetMonitorsCollection(ctx context.Context) ApiGetMonitorsCollectionRequest {
	return ApiGetMonitorsCollectionRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return Monitors
func (a *SyntheticMonitorsAPIService) GetMonitorsCollectionExecute(r ApiGetMonitorsCollectionRequest) (*Monitors, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Monitors
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SyntheticMonitorsAPIService.GetMonitorsCollection")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/synthetic/monitors"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.managementZone != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "managementZone", r.managementZone, "")
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
	if r.location != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "location", r.location, "")
	}
	if r.assignedApps != nil {
		t := *r.assignedApps
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "assignedApps", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "assignedApps", t, "multi")
		}
	}
	if r.type_ != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "type", r.type_, "")
	}
	if r.enabled != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "enabled", r.enabled, "")
	}
	if r.credentialId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "credentialId", r.credentialId, "")
	}
	if r.credentialOwner != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "credentialOwner", r.credentialOwner, "")
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

type ApiReplaceMonitorRequest struct {
	ctx                    context.Context
	ApiService             *SyntheticMonitorsAPIService
	monitorId              string
	syntheticMonitorUpdate *SyntheticMonitorUpdate
}

func (r ApiReplaceMonitorRequest) SyntheticMonitorUpdate(syntheticMonitorUpdate SyntheticMonitorUpdate) ApiReplaceMonitorRequest {
	r.syntheticMonitorUpdate = &syntheticMonitorUpdate
	return r
}

func (r ApiReplaceMonitorRequest) Execute() (*http.Response, error) {
	return r.ApiService.ReplaceMonitorExecute(r)
}

/*
ReplaceMonitor Updates parameters of the specified synthetic monitor

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param monitorId The ID of the synthetic monitor to be updated.
	@return ApiReplaceMonitorRequest
*/
func (a *SyntheticMonitorsAPIService) ReplaceMonitor(ctx context.Context, monitorId string) ApiReplaceMonitorRequest {
	return ApiReplaceMonitorRequest{
		ApiService: a,
		ctx:        ctx,
		monitorId:  monitorId,
	}
}

// Execute executes the request
func (a *SyntheticMonitorsAPIService) ReplaceMonitorExecute(r ApiReplaceMonitorRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPut
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SyntheticMonitorsAPIService.ReplaceMonitor")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/synthetic/monitors/{monitorId}"
	localVarPath = strings.Replace(localVarPath, "{"+"monitorId"+"}", url.PathEscape(parameterValueToString(r.monitorId, "monitorId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json; charset=utf-8"}

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
	// body params
	localVarPostBody = r.syntheticMonitorUpdate
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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
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
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode >= 500 {
			var v ErrorEnvelope
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
