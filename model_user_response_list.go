/*
[BETA] Qovery API

- Qovery is the fastest way to deploy your full-stack apps on any Cloud provider. - ℹ️ The API is in Beta and still in progress. Some endpoints are not available yet.  # Authentication  <!-- ReDoc-Inject: <security-definitions> -->

API version: 1.0.0
Contact: support+api+documentation@qovery.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package qovery

import (
	"encoding/json"
)

// UserResponseList struct for UserResponseList
type UserResponseList struct {
	Results *[]UserResponse `json:"results,omitempty"`
}

// NewUserResponseList instantiates a new UserResponseList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserResponseList() *UserResponseList {
	this := UserResponseList{}
	return &this
}

// NewUserResponseListWithDefaults instantiates a new UserResponseList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserResponseListWithDefaults() *UserResponseList {
	this := UserResponseList{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *UserResponseList) GetResults() []UserResponse {
	if o == nil || o.Results == nil {
		var ret []UserResponse
		return ret
	}
	return *o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResponseList) GetResultsOk() (*[]UserResponse, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *UserResponseList) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []UserResponse and assigns it to the Results field.
func (o *UserResponseList) SetResults(v []UserResponse) {
	o.Results = &v
}

func (o UserResponseList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullableUserResponseList struct {
	value *UserResponseList
	isSet bool
}

func (v NullableUserResponseList) Get() *UserResponseList {
	return v.value
}

func (v *NullableUserResponseList) Set(val *UserResponseList) {
	v.value = val
	v.isSet = true
}

func (v NullableUserResponseList) IsSet() bool {
	return v.isSet
}

func (v *NullableUserResponseList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserResponseList(val *UserResponseList) *NullableUserResponseList {
	return &NullableUserResponseList{value: val, isSet: true}
}

func (v NullableUserResponseList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserResponseList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


