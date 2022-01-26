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

// ApplicationsApiService ApplicationsApi service
type ApplicationsApiService service

type ApiCreateApplicationRequest struct {
	ctx _context.Context
	ApiService *ApplicationsApiService
	environmentId string
	applicationRequest *ApplicationRequest
}

func (r ApiCreateApplicationRequest) ApplicationRequest(applicationRequest ApplicationRequest) ApiCreateApplicationRequest {
	r.applicationRequest = &applicationRequest
	return r
}

func (r ApiCreateApplicationRequest) Execute() (ApplicationResponse, *_nethttp.Response, error) {
	return r.ApiService.CreateApplicationExecute(r)
}

/*
CreateApplication Create an application

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param environmentId Environment ID
 @return ApiCreateApplicationRequest
*/
func (a *ApplicationsApiService) CreateApplication(ctx _context.Context, environmentId string) ApiCreateApplicationRequest {
	return ApiCreateApplicationRequest{
		ApiService: a,
		ctx: ctx,
		environmentId: environmentId,
	}
}

// Execute executes the request
//  @return ApplicationResponse
func (a *ApplicationsApiService) CreateApplicationExecute(r ApiCreateApplicationRequest) (ApplicationResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  ApplicationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApplicationsApiService.CreateApplication")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/environment/{environmentId}/application"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentId"+"}", _neturl.PathEscape(parameterToString(r.environmentId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	localVarPostBody = r.applicationRequest
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

type ApiGetEnvironmentApplicationCurrentInstanceRequest struct {
	ctx _context.Context
	ApiService *ApplicationsApiService
	environmentId string
}


func (r ApiGetEnvironmentApplicationCurrentInstanceRequest) Execute() (EnvironmentApplicationsInstanceResponseList, *_nethttp.Response, error) {
	return r.ApiService.GetEnvironmentApplicationCurrentInstanceExecute(r)
}

/*
GetEnvironmentApplicationCurrentInstance List running instances with CPU and RAM usage for each application

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param environmentId Environment ID
 @return ApiGetEnvironmentApplicationCurrentInstanceRequest
*/
func (a *ApplicationsApiService) GetEnvironmentApplicationCurrentInstance(ctx _context.Context, environmentId string) ApiGetEnvironmentApplicationCurrentInstanceRequest {
	return ApiGetEnvironmentApplicationCurrentInstanceRequest{
		ApiService: a,
		ctx: ctx,
		environmentId: environmentId,
	}
}

// Execute executes the request
//  @return EnvironmentApplicationsInstanceResponseList
func (a *ApplicationsApiService) GetEnvironmentApplicationCurrentInstanceExecute(r ApiGetEnvironmentApplicationCurrentInstanceRequest) (EnvironmentApplicationsInstanceResponseList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  EnvironmentApplicationsInstanceResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApplicationsApiService.GetEnvironmentApplicationCurrentInstance")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/environment/{environmentId}/application/instance"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentId"+"}", _neturl.PathEscape(parameterToString(r.environmentId, "")), -1)

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

type ApiGetEnvironmentApplicationCurrentScaleRequest struct {
	ctx _context.Context
	ApiService *ApplicationsApiService
	environmentId string
}


func (r ApiGetEnvironmentApplicationCurrentScaleRequest) Execute() (EnvironmentApplicationsCurrentScaleResponseList, *_nethttp.Response, error) {
	return r.ApiService.GetEnvironmentApplicationCurrentScaleExecute(r)
}

/*
GetEnvironmentApplicationCurrentScale List current scaling information for each application

Returns min, max, and running number of instances for each application

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param environmentId Environment ID
 @return ApiGetEnvironmentApplicationCurrentScaleRequest
*/
func (a *ApplicationsApiService) GetEnvironmentApplicationCurrentScale(ctx _context.Context, environmentId string) ApiGetEnvironmentApplicationCurrentScaleRequest {
	return ApiGetEnvironmentApplicationCurrentScaleRequest{
		ApiService: a,
		ctx: ctx,
		environmentId: environmentId,
	}
}

// Execute executes the request
//  @return EnvironmentApplicationsCurrentScaleResponseList
func (a *ApplicationsApiService) GetEnvironmentApplicationCurrentScaleExecute(r ApiGetEnvironmentApplicationCurrentScaleRequest) (EnvironmentApplicationsCurrentScaleResponseList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  EnvironmentApplicationsCurrentScaleResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApplicationsApiService.GetEnvironmentApplicationCurrentScale")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/environment/{environmentId}/application/currentScale"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentId"+"}", _neturl.PathEscape(parameterToString(r.environmentId, "")), -1)

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

type ApiGetEnvironmentApplicationCurrentStorageRequest struct {
	ctx _context.Context
	ApiService *ApplicationsApiService
	environmentId string
}


func (r ApiGetEnvironmentApplicationCurrentStorageRequest) Execute() (EnvironmentApplicationsStorageResponseList, *_nethttp.Response, error) {
	return r.ApiService.GetEnvironmentApplicationCurrentStorageExecute(r)
}

/*
GetEnvironmentApplicationCurrentStorage List current storage disk usage for each application

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param environmentId Environment ID
 @return ApiGetEnvironmentApplicationCurrentStorageRequest
*/
func (a *ApplicationsApiService) GetEnvironmentApplicationCurrentStorage(ctx _context.Context, environmentId string) ApiGetEnvironmentApplicationCurrentStorageRequest {
	return ApiGetEnvironmentApplicationCurrentStorageRequest{
		ApiService: a,
		ctx: ctx,
		environmentId: environmentId,
	}
}

// Execute executes the request
//  @return EnvironmentApplicationsStorageResponseList
func (a *ApplicationsApiService) GetEnvironmentApplicationCurrentStorageExecute(r ApiGetEnvironmentApplicationCurrentStorageRequest) (EnvironmentApplicationsStorageResponseList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  EnvironmentApplicationsStorageResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApplicationsApiService.GetEnvironmentApplicationCurrentStorage")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/environment/{environmentId}/application/currentStorage"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentId"+"}", _neturl.PathEscape(parameterToString(r.environmentId, "")), -1)

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

type ApiGetEnvironmentApplicationStatusRequest struct {
	ctx _context.Context
	ApiService *ApplicationsApiService
	environmentId string
}


func (r ApiGetEnvironmentApplicationStatusRequest) Execute() (ReferenceObjectStatusResponseList, *_nethttp.Response, error) {
	return r.ApiService.GetEnvironmentApplicationStatusExecute(r)
}

/*
GetEnvironmentApplicationStatus List all environment applications statuses

Returns a list of applications with only their id and status.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param environmentId Environment ID
 @return ApiGetEnvironmentApplicationStatusRequest
*/
func (a *ApplicationsApiService) GetEnvironmentApplicationStatus(ctx _context.Context, environmentId string) ApiGetEnvironmentApplicationStatusRequest {
	return ApiGetEnvironmentApplicationStatusRequest{
		ApiService: a,
		ctx: ctx,
		environmentId: environmentId,
	}
}

// Execute executes the request
//  @return ReferenceObjectStatusResponseList
func (a *ApplicationsApiService) GetEnvironmentApplicationStatusExecute(r ApiGetEnvironmentApplicationStatusRequest) (ReferenceObjectStatusResponseList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  ReferenceObjectStatusResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApplicationsApiService.GetEnvironmentApplicationStatus")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/environment/{environmentId}/application/status"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentId"+"}", _neturl.PathEscape(parameterToString(r.environmentId, "")), -1)

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

type ApiGetEnvironmentApplicationSupportedLanguagesRequest struct {
	ctx _context.Context
	ApiService *ApplicationsApiService
	environmentId string
}


func (r ApiGetEnvironmentApplicationSupportedLanguagesRequest) Execute() (EnvironmentApplicationsSupportedLanguageList, *_nethttp.Response, error) {
	return r.ApiService.GetEnvironmentApplicationSupportedLanguagesExecute(r)
}

/*
GetEnvironmentApplicationSupportedLanguages List supported languages

Returns list of languages supported by Buildpacks.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param environmentId Environment ID
 @return ApiGetEnvironmentApplicationSupportedLanguagesRequest
*/
func (a *ApplicationsApiService) GetEnvironmentApplicationSupportedLanguages(ctx _context.Context, environmentId string) ApiGetEnvironmentApplicationSupportedLanguagesRequest {
	return ApiGetEnvironmentApplicationSupportedLanguagesRequest{
		ApiService: a,
		ctx: ctx,
		environmentId: environmentId,
	}
}

// Execute executes the request
//  @return EnvironmentApplicationsSupportedLanguageList
func (a *ApplicationsApiService) GetEnvironmentApplicationSupportedLanguagesExecute(r ApiGetEnvironmentApplicationSupportedLanguagesRequest) (EnvironmentApplicationsSupportedLanguageList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  EnvironmentApplicationsSupportedLanguageList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApplicationsApiService.GetEnvironmentApplicationSupportedLanguages")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/environment/{environmentId}/application/supportedLanguage"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentId"+"}", _neturl.PathEscape(parameterToString(r.environmentId, "")), -1)

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

type ApiListApplicationRequest struct {
	ctx _context.Context
	ApiService *ApplicationsApiService
	environmentId string
	toUpdate *bool
}

// return (or not) results that must be updated
func (r ApiListApplicationRequest) ToUpdate(toUpdate bool) ApiListApplicationRequest {
	r.toUpdate = &toUpdate
	return r
}

func (r ApiListApplicationRequest) Execute() (ApplicationResponseList, *_nethttp.Response, error) {
	return r.ApiService.ListApplicationExecute(r)
}

/*
ListApplication List applications

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param environmentId Environment ID
 @return ApiListApplicationRequest
*/
func (a *ApplicationsApiService) ListApplication(ctx _context.Context, environmentId string) ApiListApplicationRequest {
	return ApiListApplicationRequest{
		ApiService: a,
		ctx: ctx,
		environmentId: environmentId,
	}
}

// Execute executes the request
//  @return ApplicationResponseList
func (a *ApplicationsApiService) ListApplicationExecute(r ApiListApplicationRequest) (ApplicationResponseList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  ApplicationResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApplicationsApiService.ListApplication")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/environment/{environmentId}/application"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentId"+"}", _neturl.PathEscape(parameterToString(r.environmentId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.toUpdate != nil {
		localVarQueryParams.Add("toUpdate", parameterToString(*r.toUpdate, ""))
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
