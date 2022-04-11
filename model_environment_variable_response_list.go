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

// EnvironmentVariableResponseList struct for EnvironmentVariableResponseList
type EnvironmentVariableResponseList struct {
	Results []EnvironmentVariable `json:"results,omitempty"`
}

// NewEnvironmentVariableResponseList instantiates a new EnvironmentVariableResponseList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentVariableResponseList() *EnvironmentVariableResponseList {
	this := EnvironmentVariableResponseList{}
	return &this
}

// NewEnvironmentVariableResponseListWithDefaults instantiates a new EnvironmentVariableResponseList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentVariableResponseListWithDefaults() *EnvironmentVariableResponseList {
	this := EnvironmentVariableResponseList{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *EnvironmentVariableResponseList) GetResults() []EnvironmentVariable {
	if o == nil || o.Results == nil {
		var ret []EnvironmentVariable
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentVariableResponseList) GetResultsOk() ([]EnvironmentVariable, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *EnvironmentVariableResponseList) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []EnvironmentVariable and assigns it to the Results field.
func (o *EnvironmentVariableResponseList) SetResults(v []EnvironmentVariable) {
	o.Results = v
}

func (o EnvironmentVariableResponseList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullableEnvironmentVariableResponseList struct {
	value *EnvironmentVariableResponseList
	isSet bool
}

func (v NullableEnvironmentVariableResponseList) Get() *EnvironmentVariableResponseList {
	return v.value
}

func (v *NullableEnvironmentVariableResponseList) Set(val *EnvironmentVariableResponseList) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentVariableResponseList) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentVariableResponseList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentVariableResponseList(val *EnvironmentVariableResponseList) *NullableEnvironmentVariableResponseList {
	return &NullableEnvironmentVariableResponseList{value: val, isSet: true}
}

func (v NullableEnvironmentVariableResponseList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentVariableResponseList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
