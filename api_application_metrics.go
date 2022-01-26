/*
[BETA] Qovery API

- Qovery is the fastest way to deploy your full-stack apps on any Cloud provider. - ℹ️ The API is in Beta and still in progress. Some endpoints are not available yet.  # Authentication  <!-- ReDoc-Inject: <security-definitions> -->

API version: 1.0.0
Contact: support+api+documentation@qovery.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package qovery

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// ApplicationMetricsApiService ApplicationMetricsApi service
type ApplicationMetricsApiService service

type ApiGetApplicationCurrentInstanceRequest struct {
	ctx _context.Context
	ApiService *ApplicationMetricsApiService
	applicationId string
}


func (r ApiGetApplicationCurrentInstanceRequest) Execute() (InstanceResponseList, *_nethttp.Response, error) {
	return r.ApiService.GetApplicationCurrentInstanceExecute(r)
}

/*
GetApplicationCurrentInstance List currently running instances of the application with their CPU and RAM metrics

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param applicationId Application ID
 @return ApiGetApplicationCurrentInstanceRequest
*/
func (a *ApplicationMetricsApiService) GetApplicationCurrentInstance(ctx _context.Context, applicationId string) ApiGetApplicationCurrentInstanceRequest {
	return ApiGetApplicationCurrentInstanceRequest{
		ApiService: a,
		ctx: ctx,
		applicationId: applicationId,
	}
}

// Execute executes the request
//  @return InstanceResponseList
func (a *ApplicationMetricsApiService) GetApplicationCurrentInstanceExecute(r ApiGetApplicationCurrentInstanceRequest) (InstanceResponseList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  InstanceResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApplicationMetricsApiService.GetApplicationCurrentInstance")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/application/{applicationId}/instance"
	localVarPath = strings.Replace(localVarPath, "{"+"applicationId"+"}", _neturl.PathEscape(parameterToString(r.applicationId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetApplicationCurrentScaleRequest struct {
	ctx _context.Context
	ApiService *ApplicationMetricsApiService
	applicationId string
}


func (r ApiGetApplicationCurrentScaleRequest) Execute() (ApplicationCurrentScaleResponse, *_nethttp.Response, error) {
	return r.ApiService.GetApplicationCurrentScaleExecute(r)
}

/*
GetApplicationCurrentScale Get current scaling of the application

Returns min, max, and running number of instances of the application

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param applicationId Application ID
 @return ApiGetApplicationCurrentScaleRequest
*/
func (a *ApplicationMetricsApiService) GetApplicationCurrentScale(ctx _context.Context, applicationId string) ApiGetApplicationCurrentScaleRequest {
	return ApiGetApplicationCurrentScaleRequest{
		ApiService: a,
		ctx: ctx,
		applicationId: applicationId,
	}
}

// Execute executes the request
//  @return ApplicationCurrentScaleResponse
func (a *ApplicationMetricsApiService) GetApplicationCurrentScaleExecute(r ApiGetApplicationCurrentScaleRequest) (ApplicationCurrentScaleResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  ApplicationCurrentScaleResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApplicationMetricsApiService.GetApplicationCurrentScale")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/application/{applicationId}/currentScale"
	localVarPath = strings.Replace(localVarPath, "{"+"applicationId"+"}", _neturl.PathEscape(parameterToString(r.applicationId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetApplicationCurrentStorageDiskRequest struct {
	ctx _context.Context
	ApiService *ApplicationMetricsApiService
	applicationId string
}


func (r ApiGetApplicationCurrentStorageDiskRequest) Execute() (StorageDiskResponseList, *_nethttp.Response, error) {
	return r.ApiService.GetApplicationCurrentStorageDiskExecute(r)
}

/*
GetApplicationCurrentStorageDisk List current storage disk usage

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param applicationId Application ID
 @return ApiGetApplicationCurrentStorageDiskRequest
*/
func (a *ApplicationMetricsApiService) GetApplicationCurrentStorageDisk(ctx _context.Context, applicationId string) ApiGetApplicationCurrentStorageDiskRequest {
	return ApiGetApplicationCurrentStorageDiskRequest{
		ApiService: a,
		ctx: ctx,
		applicationId: applicationId,
	}
}

// Execute executes the request
//  @return StorageDiskResponseList
func (a *ApplicationMetricsApiService) GetApplicationCurrentStorageDiskExecute(r ApiGetApplicationCurrentStorageDiskRequest) (StorageDiskResponseList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  StorageDiskResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApplicationMetricsApiService.GetApplicationCurrentStorageDisk")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/application/{applicationId}/currentStorage"
	localVarPath = strings.Replace(localVarPath, "{"+"applicationId"+"}", _neturl.PathEscape(parameterToString(r.applicationId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetApplicationMetricCpuRequest struct {
	ctx _context.Context
	ApiService *ApplicationMetricsApiService
	applicationId string
	lastSeconds *float32
}

// Up to how many seconds in the past to ask analytics results
func (r ApiGetApplicationMetricCpuRequest) LastSeconds(lastSeconds float32) ApiGetApplicationMetricCpuRequest {
	r.lastSeconds = &lastSeconds
	return r
}

func (r ApiGetApplicationMetricCpuRequest) Execute() (MetricCPUResponseList, *_nethttp.Response, error) {
	return r.ApiService.GetApplicationMetricCpuExecute(r)
}

/*
GetApplicationMetricCpu Get CPU consumption metric over time for the application

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param applicationId Application ID
 @return ApiGetApplicationMetricCpuRequest
*/
func (a *ApplicationMetricsApiService) GetApplicationMetricCpu(ctx _context.Context, applicationId string) ApiGetApplicationMetricCpuRequest {
	return ApiGetApplicationMetricCpuRequest{
		ApiService: a,
		ctx: ctx,
		applicationId: applicationId,
	}
}

// Execute executes the request
//  @return MetricCPUResponseList
func (a *ApplicationMetricsApiService) GetApplicationMetricCpuExecute(r ApiGetApplicationMetricCpuRequest) (MetricCPUResponseList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  MetricCPUResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApplicationMetricsApiService.GetApplicationMetricCpu")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/application/{applicationId}/metric/cpu"
	localVarPath = strings.Replace(localVarPath, "{"+"applicationId"+"}", _neturl.PathEscape(parameterToString(r.applicationId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.lastSeconds == nil {
		return localVarReturnValue, nil, reportError("lastSeconds is required and must be specified")
	}

	localVarQueryParams.Add("lastSeconds", parameterToString(*r.lastSeconds, ""))
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetApplicationMetricHealthCheckRequest struct {
	ctx _context.Context
	ApiService *ApplicationMetricsApiService
	applicationId string
	lastSeconds *float32
}

// Up to how many seconds in the past to ask analytics results
func (r ApiGetApplicationMetricHealthCheckRequest) LastSeconds(lastSeconds float32) ApiGetApplicationMetricHealthCheckRequest {
	r.lastSeconds = &lastSeconds
	return r
}

func (r ApiGetApplicationMetricHealthCheckRequest) Execute() (MetricGenericResponseList, *_nethttp.Response, error) {
	return r.ApiService.GetApplicationMetricHealthCheckExecute(r)
}

/*
GetApplicationMetricHealthCheck Get Health Check latency  metric over time for the application

The value returned corresponds to the 95th centile

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param applicationId Application ID
 @return ApiGetApplicationMetricHealthCheckRequest
*/
func (a *ApplicationMetricsApiService) GetApplicationMetricHealthCheck(ctx _context.Context, applicationId string) ApiGetApplicationMetricHealthCheckRequest {
	return ApiGetApplicationMetricHealthCheckRequest{
		ApiService: a,
		ctx: ctx,
		applicationId: applicationId,
	}
}

// Execute executes the request
//  @return MetricGenericResponseList
func (a *ApplicationMetricsApiService) GetApplicationMetricHealthCheckExecute(r ApiGetApplicationMetricHealthCheckRequest) (MetricGenericResponseList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  MetricGenericResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApplicationMetricsApiService.GetApplicationMetricHealthCheck")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/application/{applicationId}/metric/healthCheck"
	localVarPath = strings.Replace(localVarPath, "{"+"applicationId"+"}", _neturl.PathEscape(parameterToString(r.applicationId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.lastSeconds == nil {
		return localVarReturnValue, nil, reportError("lastSeconds is required and must be specified")
	}

	localVarQueryParams.Add("lastSeconds", parameterToString(*r.lastSeconds, ""))
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetApplicationMetricMemoryRequest struct {
	ctx _context.Context
	ApiService *ApplicationMetricsApiService
	applicationId string
	lastSeconds *float32
}

// Up to how many seconds in the past to ask analytics results
func (r ApiGetApplicationMetricMemoryRequest) LastSeconds(lastSeconds float32) ApiGetApplicationMetricMemoryRequest {
	r.lastSeconds = &lastSeconds
	return r
}

func (r ApiGetApplicationMetricMemoryRequest) Execute() (MetricMemoryResponseList, *_nethttp.Response, error) {
	return r.ApiService.GetApplicationMetricMemoryExecute(r)
}

/*
GetApplicationMetricMemory Get Memory consumption metric over time for the application

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param applicationId Application ID
 @return ApiGetApplicationMetricMemoryRequest
*/
func (a *ApplicationMetricsApiService) GetApplicationMetricMemory(ctx _context.Context, applicationId string) ApiGetApplicationMetricMemoryRequest {
	return ApiGetApplicationMetricMemoryRequest{
		ApiService: a,
		ctx: ctx,
		applicationId: applicationId,
	}
}

// Execute executes the request
//  @return MetricMemoryResponseList
func (a *ApplicationMetricsApiService) GetApplicationMetricMemoryExecute(r ApiGetApplicationMetricMemoryRequest) (MetricMemoryResponseList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  MetricMemoryResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApplicationMetricsApiService.GetApplicationMetricMemory")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/application/{applicationId}/metric/memory"
	localVarPath = strings.Replace(localVarPath, "{"+"applicationId"+"}", _neturl.PathEscape(parameterToString(r.applicationId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.lastSeconds == nil {
		return localVarReturnValue, nil, reportError("lastSeconds is required and must be specified")
	}

	localVarQueryParams.Add("lastSeconds", parameterToString(*r.lastSeconds, ""))
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetApplicationMetricRestartRequest struct {
	ctx _context.Context
	ApiService *ApplicationMetricsApiService
	applicationId string
	lastSeconds *float32
}

// Up to how many seconds in the past to ask analytics results
func (r ApiGetApplicationMetricRestartRequest) LastSeconds(lastSeconds float32) ApiGetApplicationMetricRestartRequest {
	r.lastSeconds = &lastSeconds
	return r
}

func (r ApiGetApplicationMetricRestartRequest) Execute() (MetricRestartResponse, *_nethttp.Response, error) {
	return r.ApiService.GetApplicationMetricRestartExecute(r)
}

/*
GetApplicationMetricRestart List application restarts

Get application restart message and timestamp.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param applicationId Application ID
 @return ApiGetApplicationMetricRestartRequest
*/
func (a *ApplicationMetricsApiService) GetApplicationMetricRestart(ctx _context.Context, applicationId string) ApiGetApplicationMetricRestartRequest {
	return ApiGetApplicationMetricRestartRequest{
		ApiService: a,
		ctx: ctx,
		applicationId: applicationId,
	}
}

// Execute executes the request
//  @return MetricRestartResponse
func (a *ApplicationMetricsApiService) GetApplicationMetricRestartExecute(r ApiGetApplicationMetricRestartRequest) (MetricRestartResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  MetricRestartResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApplicationMetricsApiService.GetApplicationMetricRestart")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/application/{applicationId}/metric/restart"
	localVarPath = strings.Replace(localVarPath, "{"+"applicationId"+"}", _neturl.PathEscape(parameterToString(r.applicationId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.lastSeconds == nil {
		return localVarReturnValue, nil, reportError("lastSeconds is required and must be specified")
	}

	localVarQueryParams.Add("lastSeconds", parameterToString(*r.lastSeconds, ""))
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetApplicationMetricStorageRequest struct {
	ctx _context.Context
	ApiService *ApplicationMetricsApiService
	applicationId string
	lastSeconds *float32
}

// Up to how many seconds in the past to ask analytics results
func (r ApiGetApplicationMetricStorageRequest) LastSeconds(lastSeconds float32) ApiGetApplicationMetricStorageRequest {
	r.lastSeconds = &lastSeconds
	return r
}

func (r ApiGetApplicationMetricStorageRequest) Execute() (MetricStorageResponseList, *_nethttp.Response, error) {
	return r.ApiService.GetApplicationMetricStorageExecute(r)
}

/*
GetApplicationMetricStorage Get Storage consumption metric over time for the application

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param applicationId Application ID
 @return ApiGetApplicationMetricStorageRequest
*/
func (a *ApplicationMetricsApiService) GetApplicationMetricStorage(ctx _context.Context, applicationId string) ApiGetApplicationMetricStorageRequest {
	return ApiGetApplicationMetricStorageRequest{
		ApiService: a,
		ctx: ctx,
		applicationId: applicationId,
	}
}

// Execute executes the request
//  @return MetricStorageResponseList
func (a *ApplicationMetricsApiService) GetApplicationMetricStorageExecute(r ApiGetApplicationMetricStorageRequest) (MetricStorageResponseList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  MetricStorageResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApplicationMetricsApiService.GetApplicationMetricStorage")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/application/{applicationId}/metric/storage"
	localVarPath = strings.Replace(localVarPath, "{"+"applicationId"+"}", _neturl.PathEscape(parameterToString(r.applicationId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.lastSeconds == nil {
		return localVarReturnValue, nil, reportError("lastSeconds is required and must be specified")
	}

	localVarQueryParams.Add("lastSeconds", parameterToString(*r.lastSeconds, ""))
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
