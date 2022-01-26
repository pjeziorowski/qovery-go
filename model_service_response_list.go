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

// ServiceResponseList struct for ServiceResponseList
type ServiceResponseList struct {
	Results *[]ServiceResponse `json:"results,omitempty"`
}

// NewServiceResponseList instantiates a new ServiceResponseList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceResponseList() *ServiceResponseList {
	this := ServiceResponseList{}
	return &this
}

// NewServiceResponseListWithDefaults instantiates a new ServiceResponseList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceResponseListWithDefaults() *ServiceResponseList {
	this := ServiceResponseList{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *ServiceResponseList) GetResults() []ServiceResponse {
	if o == nil || o.Results == nil {
		var ret []ServiceResponse
		return ret
	}
	return *o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceResponseList) GetResultsOk() (*[]ServiceResponse, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *ServiceResponseList) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []ServiceResponse and assigns it to the Results field.
func (o *ServiceResponseList) SetResults(v []ServiceResponse) {
	o.Results = &v
}

func (o ServiceResponseList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullableServiceResponseList struct {
	value *ServiceResponseList
	isSet bool
}

func (v NullableServiceResponseList) Get() *ServiceResponseList {
	return v.value
}

func (v *NullableServiceResponseList) Set(val *ServiceResponseList) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceResponseList) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceResponseList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceResponseList(val *ServiceResponseList) *NullableServiceResponseList {
	return &NullableServiceResponseList{value: val, isSet: true}
}

func (v NullableServiceResponseList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceResponseList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


