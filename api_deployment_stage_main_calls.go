/*
Qovery API

- Qovery is the fastest way to deploy your full-stack apps on any Cloud provider. - ℹ️ The API is stable and still in development.

API version: 1.0.3
Contact: support+api+documentation@qovery.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package qovery

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// DeploymentStageMainCallsApiService DeploymentStageMainCallsApi service
type DeploymentStageMainCallsApiService service

type ApiAttachServiceToDeploymentStageRequest struct {
	ctx        context.Context
	ApiService *DeploymentStageMainCallsApiService
	serviceId  string
}

func (r ApiAttachServiceToDeploymentStageRequest) Execute() (*http.Response, error) {
	return r.ApiService.AttachServiceToDeploymentStageExecute(r)
}

/*
AttachServiceToDeploymentStage Attach service to deployment stage

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceId Service ID of an application/job/container/database
 @return ApiAttachServiceToDeploymentStageRequest
*/
func (a *DeploymentStageMainCallsApiService) AttachServiceToDeploymentStage(ctx context.Context, serviceId string) ApiAttachServiceToDeploymentStageRequest {
	return ApiAttachServiceToDeploymentStageRequest{
		ApiService: a,
		ctx:        ctx,
		serviceId:  serviceId,
	}
}

// Execute executes the request
func (a *DeploymentStageMainCallsApiService) AttachServiceToDeploymentStageExecute(r ApiAttachServiceToDeploymentStageRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPut
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeploymentStageMainCallsApiService.AttachServiceToDeploymentStage")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/deploymentStage/{deploymentStageId}/service/{serviceId}"
	localVarPath = strings.Replace(localVarPath, "{"+"serviceId"+"}", url.PathEscape(parameterToString(r.serviceId, "")), -1)

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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiCreateEnvironmentDeploymentStageRequest struct {
	ctx                    context.Context
	ApiService             *DeploymentStageMainCallsApiService
	environmentId          string
	deploymentStageRequest *DeploymentStageRequest
}

func (r ApiCreateEnvironmentDeploymentStageRequest) DeploymentStageRequest(deploymentStageRequest DeploymentStageRequest) ApiCreateEnvironmentDeploymentStageRequest {
	r.deploymentStageRequest = &deploymentStageRequest
	return r
}

func (r ApiCreateEnvironmentDeploymentStageRequest) Execute() (*DeploymentStageResponse, *http.Response, error) {
	return r.ApiService.CreateEnvironmentDeploymentStageExecute(r)
}

/*
CreateEnvironmentDeploymentStage Create environment deployment stage

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param environmentId Environment ID
 @return ApiCreateEnvironmentDeploymentStageRequest
*/
func (a *DeploymentStageMainCallsApiService) CreateEnvironmentDeploymentStage(ctx context.Context, environmentId string) ApiCreateEnvironmentDeploymentStageRequest {
	return ApiCreateEnvironmentDeploymentStageRequest{
		ApiService:    a,
		ctx:           ctx,
		environmentId: environmentId,
	}
}

// Execute executes the request
//  @return DeploymentStageResponse
func (a *DeploymentStageMainCallsApiService) CreateEnvironmentDeploymentStageExecute(r ApiCreateEnvironmentDeploymentStageRequest) (*DeploymentStageResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DeploymentStageResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeploymentStageMainCallsApiService.CreateEnvironmentDeploymentStage")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/environment/{environmentId}/deploymentStage"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentId"+"}", url.PathEscape(parameterToString(r.environmentId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	localVarPostBody = r.deploymentStageRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
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

type ApiDeleteDeploymentStageRequest struct {
	ctx        context.Context
	ApiService *DeploymentStageMainCallsApiService
}

func (r ApiDeleteDeploymentStageRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteDeploymentStageExecute(r)
}

/*
DeleteDeploymentStage Delete deployment stage

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiDeleteDeploymentStageRequest
*/
func (a *DeploymentStageMainCallsApiService) DeleteDeploymentStage(ctx context.Context) ApiDeleteDeploymentStageRequest {
	return ApiDeleteDeploymentStageRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *DeploymentStageMainCallsApiService) DeleteDeploymentStageExecute(r ApiDeleteDeploymentStageRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeploymentStageMainCallsApiService.DeleteDeploymentStage")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/deploymentStage/{deploymentStageId}"

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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiEditDeploymentStageRequest struct {
	ctx                    context.Context
	ApiService             *DeploymentStageMainCallsApiService
	deploymentStageId      string
	deploymentStageRequest *DeploymentStageRequest
}

func (r ApiEditDeploymentStageRequest) DeploymentStageRequest(deploymentStageRequest DeploymentStageRequest) ApiEditDeploymentStageRequest {
	r.deploymentStageRequest = &deploymentStageRequest
	return r
}

func (r ApiEditDeploymentStageRequest) Execute() (*DeploymentStageResponse, *http.Response, error) {
	return r.ApiService.EditDeploymentStageExecute(r)
}

/*
EditDeploymentStage Edit deployment stage

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param deploymentStageId Deployment Stage ID
 @return ApiEditDeploymentStageRequest
*/
func (a *DeploymentStageMainCallsApiService) EditDeploymentStage(ctx context.Context, deploymentStageId string) ApiEditDeploymentStageRequest {
	return ApiEditDeploymentStageRequest{
		ApiService:        a,
		ctx:               ctx,
		deploymentStageId: deploymentStageId,
	}
}

// Execute executes the request
//  @return DeploymentStageResponse
func (a *DeploymentStageMainCallsApiService) EditDeploymentStageExecute(r ApiEditDeploymentStageRequest) (*DeploymentStageResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DeploymentStageResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeploymentStageMainCallsApiService.EditDeploymentStage")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/deploymentStage/{deploymentStageId}"
	localVarPath = strings.Replace(localVarPath, "{"+"deploymentStageId"+"}", url.PathEscape(parameterToString(r.deploymentStageId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	localVarPostBody = r.deploymentStageRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
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

type ApiListEnvironmentDeploymentStageRequest struct {
	ctx           context.Context
	ApiService    *DeploymentStageMainCallsApiService
	environmentId string
}

func (r ApiListEnvironmentDeploymentStageRequest) Execute() (*DeploymentStageResponseList, *http.Response, error) {
	return r.ApiService.ListEnvironmentDeploymentStageExecute(r)
}

/*
ListEnvironmentDeploymentStage List environment deployment stage

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param environmentId Environment ID
 @return ApiListEnvironmentDeploymentStageRequest
*/
func (a *DeploymentStageMainCallsApiService) ListEnvironmentDeploymentStage(ctx context.Context, environmentId string) ApiListEnvironmentDeploymentStageRequest {
	return ApiListEnvironmentDeploymentStageRequest{
		ApiService:    a,
		ctx:           ctx,
		environmentId: environmentId,
	}
}

// Execute executes the request
//  @return DeploymentStageResponseList
func (a *DeploymentStageMainCallsApiService) ListEnvironmentDeploymentStageExecute(r ApiListEnvironmentDeploymentStageRequest) (*DeploymentStageResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DeploymentStageResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeploymentStageMainCallsApiService.ListEnvironmentDeploymentStage")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/environment/{environmentId}/deploymentStage"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentId"+"}", url.PathEscape(parameterToString(r.environmentId, "")), -1)

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
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

type ApiMoveDeploymentStageRequest struct {
	ctx        context.Context
	ApiService *DeploymentStageMainCallsApiService
	stageId    string
}

func (r ApiMoveDeploymentStageRequest) Execute() (*DeploymentStageResponseList, *http.Response, error) {
	return r.ApiService.MoveDeploymentStageExecute(r)
}

/*
MoveDeploymentStage Move deployment stage before requested stage

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param stageId Deployment Stage ID
 @return ApiMoveDeploymentStageRequest
*/
func (a *DeploymentStageMainCallsApiService) MoveDeploymentStage(ctx context.Context, stageId string) ApiMoveDeploymentStageRequest {
	return ApiMoveDeploymentStageRequest{
		ApiService: a,
		ctx:        ctx,
		stageId:    stageId,
	}
}

// Execute executes the request
//  @return DeploymentStageResponseList
func (a *DeploymentStageMainCallsApiService) MoveDeploymentStageExecute(r ApiMoveDeploymentStageRequest) (*DeploymentStageResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DeploymentStageResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeploymentStageMainCallsApiService.MoveDeploymentStage")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/deploymentStage/{deploymentStageId}/moveBefore/{stageId}"
	localVarPath = strings.Replace(localVarPath, "{"+"stageId"+"}", url.PathEscape(parameterToString(r.stageId, "")), -1)

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
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
