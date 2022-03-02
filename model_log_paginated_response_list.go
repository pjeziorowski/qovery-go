/*
Qovery API

- Qovery is the fastest way to deploy your full-stack apps on any Cloud provider. - ℹ️ The API is stable and still in development.

API version: 1.0.3
Contact: support+api+documentation@qovery.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package qovery

import (
	"encoding/json"
)

// LogPaginatedResponseList struct for LogPaginatedResponseList
type LogPaginatedResponseList struct {
	Page     float32       `json:"page"`
	PageSize float32       `json:"page_size"`
	Results  []LogResponse `json:"results,omitempty"`
}

// NewLogPaginatedResponseList instantiates a new LogPaginatedResponseList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogPaginatedResponseList(page float32, pageSize float32) *LogPaginatedResponseList {
	this := LogPaginatedResponseList{}
	this.Page = page
	this.PageSize = pageSize
	return &this
}

// NewLogPaginatedResponseListWithDefaults instantiates a new LogPaginatedResponseList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogPaginatedResponseListWithDefaults() *LogPaginatedResponseList {
	this := LogPaginatedResponseList{}
	return &this
}

// GetPage returns the Page field value
func (o *LogPaginatedResponseList) GetPage() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Page
}

// GetPageOk returns a tuple with the Page field value
// and a boolean to check if the value has been set.
func (o *LogPaginatedResponseList) GetPageOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Page, true
}

// SetPage sets field value
func (o *LogPaginatedResponseList) SetPage(v float32) {
	o.Page = v
}

// GetPageSize returns the PageSize field value
func (o *LogPaginatedResponseList) GetPageSize() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value
// and a boolean to check if the value has been set.
func (o *LogPaginatedResponseList) GetPageSizeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PageSize, true
}

// SetPageSize sets field value
func (o *LogPaginatedResponseList) SetPageSize(v float32) {
	o.PageSize = v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *LogPaginatedResponseList) GetResults() []LogResponse {
	if o == nil || o.Results == nil {
		var ret []LogResponse
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogPaginatedResponseList) GetResultsOk() ([]LogResponse, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *LogPaginatedResponseList) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []LogResponse and assigns it to the Results field.
func (o *LogPaginatedResponseList) SetResults(v []LogResponse) {
	o.Results = v
}

func (o LogPaginatedResponseList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["page"] = o.Page
	}
	if true {
		toSerialize["page_size"] = o.PageSize
	}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullableLogPaginatedResponseList struct {
	value *LogPaginatedResponseList
	isSet bool
}

func (v NullableLogPaginatedResponseList) Get() *LogPaginatedResponseList {
	return v.value
}

func (v *NullableLogPaginatedResponseList) Set(val *LogPaginatedResponseList) {
	v.value = val
	v.isSet = true
}

func (v NullableLogPaginatedResponseList) IsSet() bool {
	return v.isSet
}

func (v *NullableLogPaginatedResponseList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogPaginatedResponseList(val *LogPaginatedResponseList) *NullableLogPaginatedResponseList {
	return &NullableLogPaginatedResponseList{value: val, isSet: true}
}

func (v NullableLogPaginatedResponseList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogPaginatedResponseList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
