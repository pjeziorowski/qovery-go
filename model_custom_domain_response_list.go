/*
 * [BETA] Qovery API
 *
 * - Qovery is the fastest way to deploy your full-stack apps on any Cloud provider. - ℹ️ The API is in Beta and still in progress. Some endpoints are not available yet.
 *
 * API version: 1.0.0
 * Contact: support+api+documentation@qovery.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package qovery

import (
	"encoding/json"
)

// CustomDomainResponseList struct for CustomDomainResponseList
type CustomDomainResponseList struct {
	Results *[]CustomDomainResponse `json:"results,omitempty"`
}

// NewCustomDomainResponseList instantiates a new CustomDomainResponseList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomDomainResponseList() *CustomDomainResponseList {
	this := CustomDomainResponseList{}
	return &this
}

// NewCustomDomainResponseListWithDefaults instantiates a new CustomDomainResponseList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomDomainResponseListWithDefaults() *CustomDomainResponseList {
	this := CustomDomainResponseList{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *CustomDomainResponseList) GetResults() []CustomDomainResponse {
	if o == nil || o.Results == nil {
		var ret []CustomDomainResponse
		return ret
	}
	return *o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomDomainResponseList) GetResultsOk() (*[]CustomDomainResponse, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *CustomDomainResponseList) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []CustomDomainResponse and assigns it to the Results field.
func (o *CustomDomainResponseList) SetResults(v []CustomDomainResponse) {
	o.Results = &v
}

func (o CustomDomainResponseList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullableCustomDomainResponseList struct {
	value *CustomDomainResponseList
	isSet bool
}

func (v NullableCustomDomainResponseList) Get() *CustomDomainResponseList {
	return v.value
}

func (v *NullableCustomDomainResponseList) Set(val *CustomDomainResponseList) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomDomainResponseList) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomDomainResponseList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomDomainResponseList(val *CustomDomainResponseList) *NullableCustomDomainResponseList {
	return &NullableCustomDomainResponseList{value: val, isSet: true}
}

func (v NullableCustomDomainResponseList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomDomainResponseList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
