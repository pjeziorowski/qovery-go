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
	"io"
	"net/http"
	"net/url"
	"strings"
)

// EnvironmentDeploymentHistoryAPIService EnvironmentDeploymentHistoryAPI service
type EnvironmentDeploymentHistoryAPIService service

type ApiListEnvironmentDeploymentHistoryRequest struct {
	ctx           context.Context
	ApiService    *EnvironmentDeploymentHistoryAPIService
	environmentId string
	startId       *string
}

// Starting point after which to return results
func (r ApiListEnvironmentDeploymentHistoryRequest) StartId(startId string) ApiListEnvironmentDeploymentHistoryRequest {
	r.startId = &startId
	return r
}

func (r ApiListEnvironmentDeploymentHistoryRequest) Execute() (*DeploymentHistoryEnvironmentPaginatedResponseList, *http.Response, error) {
	return r.ApiService.ListEnvironmentDeploymentHistoryExecute(r)
}

/*
ListEnvironmentDeploymentHistory List environment deployments

List previous and current environment deployments with the status deployment and the related services. By default it returns the 20 last results. The response is paginated. In order to request the next page, you can use the startId query parameter

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param environmentId Environment ID
 @return ApiListEnvironmentDeploymentHistoryRequest
*/
func (a *EnvironmentDeploymentHistoryAPIService) ListEnvironmentDeploymentHistory(ctx context.Context, environmentId string) ApiListEnvironmentDeploymentHistoryRequest {
	return ApiListEnvironmentDeploymentHistoryRequest{
		ApiService:    a,
		ctx:           ctx,
		environmentId: environmentId,
	}
}

// Execute executes the request
//  @return DeploymentHistoryEnvironmentPaginatedResponseList
func (a *EnvironmentDeploymentHistoryAPIService) ListEnvironmentDeploymentHistoryExecute(r ApiListEnvironmentDeploymentHistoryRequest) (*DeploymentHistoryEnvironmentPaginatedResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DeploymentHistoryEnvironmentPaginatedResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EnvironmentDeploymentHistoryAPIService.ListEnvironmentDeploymentHistory")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/environment/{environmentId}/deploymentHistory"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentId"+"}", url.PathEscape(parameterValueToString(r.environmentId, "environmentId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.startId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "startId", r.startId, "")
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
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
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

type ApiListEnvironmentDeploymentHistoryV2Request struct {
	ctx           context.Context
	ApiService    *EnvironmentDeploymentHistoryAPIService
	environmentId string
	startId       *string
}

// Starting point after which to return results
func (r ApiListEnvironmentDeploymentHistoryV2Request) StartId(startId string) ApiListEnvironmentDeploymentHistoryV2Request {
	r.startId = &startId
	return r
}

func (r ApiListEnvironmentDeploymentHistoryV2Request) Execute() (*DeploymentHistoryEnvironmentPaginatedResponseListV2, *http.Response, error) {
	return r.ApiService.ListEnvironmentDeploymentHistoryV2Execute(r)
}

/*
ListEnvironmentDeploymentHistoryV2 List environment deployments

List previous and current environment deployments with the status deployment and the related services. By default it returns the 20 last results. The response is paginated. In order to request the next page, you can use the startId query parameter

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param environmentId Environment ID
 @return ApiListEnvironmentDeploymentHistoryV2Request
*/
func (a *EnvironmentDeploymentHistoryAPIService) ListEnvironmentDeploymentHistoryV2(ctx context.Context, environmentId string) ApiListEnvironmentDeploymentHistoryV2Request {
	return ApiListEnvironmentDeploymentHistoryV2Request{
		ApiService:    a,
		ctx:           ctx,
		environmentId: environmentId,
	}
}

// Execute executes the request
//  @return DeploymentHistoryEnvironmentPaginatedResponseListV2
func (a *EnvironmentDeploymentHistoryAPIService) ListEnvironmentDeploymentHistoryV2Execute(r ApiListEnvironmentDeploymentHistoryV2Request) (*DeploymentHistoryEnvironmentPaginatedResponseListV2, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DeploymentHistoryEnvironmentPaginatedResponseListV2
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EnvironmentDeploymentHistoryAPIService.ListEnvironmentDeploymentHistoryV2")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/environment/{environmentId}/deploymentHistoryV2"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentId"+"}", url.PathEscape(parameterValueToString(r.environmentId, "environmentId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.startId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "startId", r.startId, "")
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
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
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
