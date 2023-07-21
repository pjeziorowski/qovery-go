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

// VariableMainCallsApiService VariableMainCallsApi service
type VariableMainCallsApiService service

type ApiCreateVariableAliasRequest struct {
	ctx                  context.Context
	ApiService           *VariableMainCallsApiService
	variableId           string
	variableAliasRequest *VariableAliasRequest
}

func (r ApiCreateVariableAliasRequest) VariableAliasRequest(variableAliasRequest VariableAliasRequest) ApiCreateVariableAliasRequest {
	r.variableAliasRequest = &variableAliasRequest
	return r
}

func (r ApiCreateVariableAliasRequest) Execute() (*VariableResponse, *http.Response, error) {
	return r.ApiService.CreateVariableAliasExecute(r)
}

/*
CreateVariableAlias Create a variable alias

- Allows you to add an alias at the level defined in the request body on an existing variable having a higher scope, in order to customize its key.
- You have to specify a key in the request body and the scope and the parent id of the alias
- The system will create a new variable at the requested level with the same value as the one corresponding to the variable id in the path
- The response body will contain the newly created variable
- Information regarding the aliased_variable will be exposed in the "aliased_variable" or in the "aliased_secret" field of the newly created variable
- Only 1 alias level is allowed. You can't create an alias on an alias


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param variableId Variable ID
 @return ApiCreateVariableAliasRequest
*/
func (a *VariableMainCallsApiService) CreateVariableAlias(ctx context.Context, variableId string) ApiCreateVariableAliasRequest {
	return ApiCreateVariableAliasRequest{
		ApiService: a,
		ctx:        ctx,
		variableId: variableId,
	}
}

// Execute executes the request
//  @return VariableResponse
func (a *VariableMainCallsApiService) CreateVariableAliasExecute(r ApiCreateVariableAliasRequest) (*VariableResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *VariableResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VariableMainCallsApiService.CreateVariableAlias")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/variable/{variableId}/alias"
	localVarPath = strings.Replace(localVarPath, "{"+"variableId"+"}", url.PathEscape(parameterToString(r.variableId, "")), -1)

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
	localVarPostBody = r.variableAliasRequest
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

type ApiCreateVariableOverrideRequest struct {
	ctx                     context.Context
	ApiService              *VariableMainCallsApiService
	variableId              string
	variableOverrideRequest *VariableOverrideRequest
}

func (r ApiCreateVariableOverrideRequest) VariableOverrideRequest(variableOverrideRequest VariableOverrideRequest) ApiCreateVariableOverrideRequest {
	r.variableOverrideRequest = &variableOverrideRequest
	return r
}

func (r ApiCreateVariableOverrideRequest) Execute() (*VariableResponse, *http.Response, error) {
	return r.ApiService.CreateVariableOverrideExecute(r)
}

/*
CreateVariableOverride Create a variable override

- Allows you to override a variable that has a higher scope.
- You have to specify a value in the request body and the scope and the parent id of the variable to alias
- The system will create a new environment variable at project level with the same key as the one corresponding to the variable id in the path
- The response body will contain the newly created variable
- Information regarding the overridden_variable will be exposed in the "overridden_variable" or in the "overridden_secret" field of the newly created variable


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param variableId Variable ID
 @return ApiCreateVariableOverrideRequest
*/
func (a *VariableMainCallsApiService) CreateVariableOverride(ctx context.Context, variableId string) ApiCreateVariableOverrideRequest {
	return ApiCreateVariableOverrideRequest{
		ApiService: a,
		ctx:        ctx,
		variableId: variableId,
	}
}

// Execute executes the request
//  @return VariableResponse
func (a *VariableMainCallsApiService) CreateVariableOverrideExecute(r ApiCreateVariableOverrideRequest) (*VariableResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *VariableResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VariableMainCallsApiService.CreateVariableOverride")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/variable/{variableId}/override"
	localVarPath = strings.Replace(localVarPath, "{"+"variableId"+"}", url.PathEscape(parameterToString(r.variableId, "")), -1)

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
	localVarPostBody = r.variableOverrideRequest
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
