/*
[BETA] Qovery API

- Qovery is the fastest way to deploy your full-stack apps on any Cloud provider. - ℹ️ The API is in Beta and still in progress. Some endpoints are not available yet. 

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

// DependencyApiService DependencyApi service
type DependencyApiService service

type ApiCreateApplicationDependencyRequest struct {
	ctx _context.Context
	ApiService *DependencyApiService
	applicationId string
	targetApplicationId string
}


func (r ApiCreateApplicationDependencyRequest) Execute() (ApplicationResponse, *_nethttp.Response, error) {
	return r.ApiService.CreateApplicationDependencyExecute(r)
}

/*
CreateApplicationDependency Add application dependency to this application.

Add application dependency to this application to prevent this application starting before the linked dependencies

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param applicationId Application ID
 @param targetApplicationId Target application ID
 @return ApiCreateApplicationDependencyRequest
*/
func (a *DependencyApiService) CreateApplicationDependency(ctx _context.Context, applicationId string, targetApplicationId string) ApiCreateApplicationDependencyRequest {
	return ApiCreateApplicationDependencyRequest{
		ApiService: a,
		ctx: ctx,
		applicationId: applicationId,
		targetApplicationId: targetApplicationId,
	}
}

// Execute executes the request
//  @return ApplicationResponse
func (a *DependencyApiService) CreateApplicationDependencyExecute(r ApiCreateApplicationDependencyRequest) (ApplicationResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  ApplicationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DependencyApiService.CreateApplicationDependency")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/application/{applicationId}/dependency/{targetApplicationId}"
	localVarPath = strings.Replace(localVarPath, "{"+"applicationId"+"}", _neturl.PathEscape(parameterToString(r.applicationId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"targetApplicationId"+"}", _neturl.PathEscape(parameterToString(r.targetApplicationId, "")), -1)

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

type ApiListApplicationDependencyRequest struct {
	ctx _context.Context
	ApiService *DependencyApiService
	applicationId string
}


func (r ApiListApplicationDependencyRequest) Execute() (ApplicationResponseList, *_nethttp.Response, error) {
	return r.ApiService.ListApplicationDependencyExecute(r)
}

/*
ListApplicationDependency List application dependencies

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param applicationId Application ID
 @return ApiListApplicationDependencyRequest
*/
func (a *DependencyApiService) ListApplicationDependency(ctx _context.Context, applicationId string) ApiListApplicationDependencyRequest {
	return ApiListApplicationDependencyRequest{
		ApiService: a,
		ctx: ctx,
		applicationId: applicationId,
	}
}

// Execute executes the request
//  @return ApplicationResponseList
func (a *DependencyApiService) ListApplicationDependencyExecute(r ApiListApplicationDependencyRequest) (ApplicationResponseList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  ApplicationResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DependencyApiService.ListApplicationDependency")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/application/{applicationId}/dependency"
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

type ApiRemoveApplicationDependencyRequest struct {
	ctx _context.Context
	ApiService *DependencyApiService
	applicationId string
	targetApplicationId string
}


func (r ApiRemoveApplicationDependencyRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.RemoveApplicationDependencyExecute(r)
}

/*
RemoveApplicationDependency Remove application dependency to this application.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param applicationId Application ID
 @param targetApplicationId Target application ID
 @return ApiRemoveApplicationDependencyRequest
*/
func (a *DependencyApiService) RemoveApplicationDependency(ctx _context.Context, applicationId string, targetApplicationId string) ApiRemoveApplicationDependencyRequest {
	return ApiRemoveApplicationDependencyRequest{
		ApiService: a,
		ctx: ctx,
		applicationId: applicationId,
		targetApplicationId: targetApplicationId,
	}
}

// Execute executes the request
func (a *DependencyApiService) RemoveApplicationDependencyExecute(r ApiRemoveApplicationDependencyRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DependencyApiService.RemoveApplicationDependency")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/application/{applicationId}/dependency/{targetApplicationId}"
	localVarPath = strings.Replace(localVarPath, "{"+"applicationId"+"}", _neturl.PathEscape(parameterToString(r.applicationId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"targetApplicationId"+"}", _neturl.PathEscape(parameterToString(r.targetApplicationId, "")), -1)

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
