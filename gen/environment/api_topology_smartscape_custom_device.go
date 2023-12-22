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
	"strings"
)

type TopologySmartscapeCustomDeviceAPI interface {

	/*
		CreateCustomDataPoints Creates or updates a custom device, or reports metric data points to the custom device.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param customDeviceId The ID of the custom device.    If you use the ID of an existing device, the respective parameters will be updated.   Don't use Dynatrace entity ID here.
		@return ApiCreateCustomDataPointsRequest

		Deprecated
	*/
	CreateCustomDataPoints(ctx context.Context, customDeviceId string) ApiCreateCustomDataPointsRequest

	// CreateCustomDataPointsExecute executes the request
	//  @return CustomDevicePushResult
	// Deprecated
	CreateCustomDataPointsExecute(r ApiCreateCustomDataPointsRequest) (*CustomDevicePushResult, *http.Response, error)
}

// TopologySmartscapeCustomDeviceAPIService TopologySmartscapeCustomDeviceAPI service
type TopologySmartscapeCustomDeviceAPIService service

type ApiCreateCustomDataPointsRequest struct {
	ctx                     context.Context
	ApiService              TopologySmartscapeCustomDeviceAPI
	customDeviceId          string
	customDevicePushMessage *CustomDevicePushMessage
}

// The JSON body of the request. Contains parameters of a custom device.
func (r ApiCreateCustomDataPointsRequest) CustomDevicePushMessage(customDevicePushMessage CustomDevicePushMessage) ApiCreateCustomDataPointsRequest {
	r.customDevicePushMessage = &customDevicePushMessage
	return r
}

func (r ApiCreateCustomDataPointsRequest) Execute() (*CustomDevicePushResult, *http.Response, error) {
	return r.ApiService.CreateCustomDataPointsExecute(r)
}

/*
CreateCustomDataPoints Creates or updates a custom device, or reports metric data points to the custom device.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param customDeviceId The ID of the custom device.    If you use the ID of an existing device, the respective parameters will be updated.   Don't use Dynatrace entity ID here.
	@return ApiCreateCustomDataPointsRequest

Deprecated
*/
func (a *TopologySmartscapeCustomDeviceAPIService) CreateCustomDataPoints(ctx context.Context, customDeviceId string) ApiCreateCustomDataPointsRequest {
	return ApiCreateCustomDataPointsRequest{
		ApiService:     a,
		ctx:            ctx,
		customDeviceId: customDeviceId,
	}
}

// Execute executes the request
//
//	@return CustomDevicePushResult
//
// Deprecated
func (a *TopologySmartscapeCustomDeviceAPIService) CreateCustomDataPointsExecute(r ApiCreateCustomDataPointsRequest) (*CustomDevicePushResult, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CustomDevicePushResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TopologySmartscapeCustomDeviceAPIService.CreateCustomDataPoints")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/entity/infrastructure/custom/{customDeviceId}"
	localVarPath = strings.Replace(localVarPath, "{"+"customDeviceId"+"}", url.PathEscape(parameterValueToString(r.customDeviceId, "customDeviceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.customDeviceId) < 0 {
		return localVarReturnValue, nil, reportError("customDeviceId must have at least 0 elements")
	}
	if strlen(r.customDeviceId) > 512 {
		return localVarReturnValue, nil, reportError("customDeviceId must have less than 512 elements")
	}

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
	localVarPostBody = r.customDevicePushMessage
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
