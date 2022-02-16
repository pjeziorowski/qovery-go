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

// ApplicationDeploymentHistoryApiService ApplicationDeploymentHistoryApi service
type ApplicationDeploymentHistoryApiService service

type ApiListApplicationDeploymentHistoryRequest struct {
	ctx _context.Context
	ApiService *ApplicationDeploymentHistoryApiService
	applicationId string
	startId *string
}

// Starting point after which to return results
func (r ApiListApplicationDeploymentHistoryRequest) StartId(startId string) ApiListApplicationDeploymentHistoryRequest {
	r.startId = &startId
	return r
}

func (r ApiListApplicationDeploymentHistoryRequest) Execute() (DeploymentHistoryPaginatedResponseList, *_nethttp.Response, error) {
	return r.ApiService.ListApplicationDeploymentHistoryExecute(r)
}

/*
ListApplicationDeploymentHistory List application deploys

By default it returns the 20 last results. The response is paginated. In order to request the next page, you can use the startId query parameter. You can also filter by status (FAILED or SUCCESS), and git_commit_id

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param applicationId Application ID
 @return ApiListApplicationDeploymentHistoryRequest
*/
func (a *ApplicationDeploymentHistoryApiService) ListApplicationDeploymentHistory(ctx _context.Context, applicationId string) ApiListApplicationDeploymentHistoryRequest {
	return ApiListApplicationDeploymentHistoryRequest{
		ApiService: a,
		ctx: ctx,
		applicationId: applicationId,
	}
}

// Execute executes the request
//  @return DeploymentHistoryPaginatedResponseList
func (a *ApplicationDeploymentHistoryApiService) ListApplicationDeploymentHistoryExecute(r ApiListApplicationDeploymentHistoryRequest) (DeploymentHistoryPaginatedResponseList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  DeploymentHistoryPaginatedResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApplicationDeploymentHistoryApiService.ListApplicationDeploymentHistory")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/application/{applicationId}/deploymentHistory"
	localVarPath = strings.Replace(localVarPath, "{"+"applicationId"+"}", _neturl.PathEscape(parameterToString(r.applicationId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.startId != nil {
		localVarQueryParams.Add("startId", parameterToString(*r.startId, ""))
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
