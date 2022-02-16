/*
[BETA] Qovery API

- Qovery is the fastest way to deploy your full-stack apps on any Cloud provider. - ℹ️ The API is in Beta and still in progress. Some endpoints are not available yet. 

API version: 1.0.0
Contact: support+api+documentation@qovery.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package qovery

import (
	"encoding/json"
)

// DeploymentHistoryPaginatedResponseList struct for DeploymentHistoryPaginatedResponseList
type DeploymentHistoryPaginatedResponseList struct {
	Results *[]DeploymentHistoryResponse `json:"results,omitempty"`
	Page float32 `json:"page"`
	PageSize float32 `json:"page_size"`
}

// NewDeploymentHistoryPaginatedResponseList instantiates a new DeploymentHistoryPaginatedResponseList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploymentHistoryPaginatedResponseList(page float32, pageSize float32) *DeploymentHistoryPaginatedResponseList {
	this := DeploymentHistoryPaginatedResponseList{}
	this.Page = page
	this.PageSize = pageSize
	return &this
}

// NewDeploymentHistoryPaginatedResponseListWithDefaults instantiates a new DeploymentHistoryPaginatedResponseList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentHistoryPaginatedResponseListWithDefaults() *DeploymentHistoryPaginatedResponseList {
	this := DeploymentHistoryPaginatedResponseList{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *DeploymentHistoryPaginatedResponseList) GetResults() []DeploymentHistoryResponse {
	if o == nil || o.Results == nil {
		var ret []DeploymentHistoryResponse
		return ret
	}
	return *o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentHistoryPaginatedResponseList) GetResultsOk() (*[]DeploymentHistoryResponse, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *DeploymentHistoryPaginatedResponseList) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []DeploymentHistoryResponse and assigns it to the Results field.
func (o *DeploymentHistoryPaginatedResponseList) SetResults(v []DeploymentHistoryResponse) {
	o.Results = &v
}

// GetPage returns the Page field value
func (o *DeploymentHistoryPaginatedResponseList) GetPage() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Page
}

// GetPageOk returns a tuple with the Page field value
// and a boolean to check if the value has been set.
func (o *DeploymentHistoryPaginatedResponseList) GetPageOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Page, true
}

// SetPage sets field value
func (o *DeploymentHistoryPaginatedResponseList) SetPage(v float32) {
	o.Page = v
}

// GetPageSize returns the PageSize field value
func (o *DeploymentHistoryPaginatedResponseList) GetPageSize() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value
// and a boolean to check if the value has been set.
func (o *DeploymentHistoryPaginatedResponseList) GetPageSizeOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PageSize, true
}

// SetPageSize sets field value
func (o *DeploymentHistoryPaginatedResponseList) SetPageSize(v float32) {
	o.PageSize = v
}

func (o DeploymentHistoryPaginatedResponseList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}
	if true {
		toSerialize["page"] = o.Page
	}
	if true {
		toSerialize["page_size"] = o.PageSize
	}
	return json.Marshal(toSerialize)
}

type NullableDeploymentHistoryPaginatedResponseList struct {
	value *DeploymentHistoryPaginatedResponseList
	isSet bool
}

func (v NullableDeploymentHistoryPaginatedResponseList) Get() *DeploymentHistoryPaginatedResponseList {
	return v.value
}

func (v *NullableDeploymentHistoryPaginatedResponseList) Set(val *DeploymentHistoryPaginatedResponseList) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploymentHistoryPaginatedResponseList) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploymentHistoryPaginatedResponseList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploymentHistoryPaginatedResponseList(val *DeploymentHistoryPaginatedResponseList) *NullableDeploymentHistoryPaginatedResponseList {
	return &NullableDeploymentHistoryPaginatedResponseList{value: val, isSet: true}
}

func (v NullableDeploymentHistoryPaginatedResponseList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploymentHistoryPaginatedResponseList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


