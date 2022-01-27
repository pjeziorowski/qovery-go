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

// EnvironmentLogPaginatedResponseList struct for EnvironmentLogPaginatedResponseList
type EnvironmentLogPaginatedResponseList struct {
	Results  *[]EnvironmentLogResponse `json:"results,omitempty"`
	Page     float32                   `json:"page"`
	PageSize float32                   `json:"page_size"`
}

// NewEnvironmentLogPaginatedResponseList instantiates a new EnvironmentLogPaginatedResponseList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentLogPaginatedResponseList(page float32, pageSize float32) *EnvironmentLogPaginatedResponseList {
	this := EnvironmentLogPaginatedResponseList{}
	this.Page = page
	this.PageSize = pageSize
	return &this
}

// NewEnvironmentLogPaginatedResponseListWithDefaults instantiates a new EnvironmentLogPaginatedResponseList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentLogPaginatedResponseListWithDefaults() *EnvironmentLogPaginatedResponseList {
	this := EnvironmentLogPaginatedResponseList{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *EnvironmentLogPaginatedResponseList) GetResults() []EnvironmentLogResponse {
	if o == nil || o.Results == nil {
		var ret []EnvironmentLogResponse
		return ret
	}
	return *o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentLogPaginatedResponseList) GetResultsOk() (*[]EnvironmentLogResponse, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *EnvironmentLogPaginatedResponseList) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []EnvironmentLogResponse and assigns it to the Results field.
func (o *EnvironmentLogPaginatedResponseList) SetResults(v []EnvironmentLogResponse) {
	o.Results = &v
}

// GetPage returns the Page field value
func (o *EnvironmentLogPaginatedResponseList) GetPage() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Page
}

// GetPageOk returns a tuple with the Page field value
// and a boolean to check if the value has been set.
func (o *EnvironmentLogPaginatedResponseList) GetPageOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Page, true
}

// SetPage sets field value
func (o *EnvironmentLogPaginatedResponseList) SetPage(v float32) {
	o.Page = v
}

// GetPageSize returns the PageSize field value
func (o *EnvironmentLogPaginatedResponseList) GetPageSize() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value
// and a boolean to check if the value has been set.
func (o *EnvironmentLogPaginatedResponseList) GetPageSizeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PageSize, true
}

// SetPageSize sets field value
func (o *EnvironmentLogPaginatedResponseList) SetPageSize(v float32) {
	o.PageSize = v
}

func (o EnvironmentLogPaginatedResponseList) MarshalJSON() ([]byte, error) {
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

type NullableEnvironmentLogPaginatedResponseList struct {
	value *EnvironmentLogPaginatedResponseList
	isSet bool
}

func (v NullableEnvironmentLogPaginatedResponseList) Get() *EnvironmentLogPaginatedResponseList {
	return v.value
}

func (v *NullableEnvironmentLogPaginatedResponseList) Set(val *EnvironmentLogPaginatedResponseList) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentLogPaginatedResponseList) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentLogPaginatedResponseList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentLogPaginatedResponseList(val *EnvironmentLogPaginatedResponseList) *NullableEnvironmentLogPaginatedResponseList {
	return &NullableEnvironmentLogPaginatedResponseList{value: val, isSet: true}
}

func (v NullableEnvironmentLogPaginatedResponseList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentLogPaginatedResponseList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
