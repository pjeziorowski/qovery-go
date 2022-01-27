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

// LinkResponseList struct for LinkResponseList
type LinkResponseList struct {
	Results *[]LinkResponse `json:"results,omitempty"`
}

// NewLinkResponseList instantiates a new LinkResponseList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkResponseList() *LinkResponseList {
	this := LinkResponseList{}
	return &this
}

// NewLinkResponseListWithDefaults instantiates a new LinkResponseList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkResponseListWithDefaults() *LinkResponseList {
	this := LinkResponseList{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *LinkResponseList) GetResults() []LinkResponse {
	if o == nil || o.Results == nil {
		var ret []LinkResponse
		return ret
	}
	return *o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkResponseList) GetResultsOk() (*[]LinkResponse, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *LinkResponseList) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []LinkResponse and assigns it to the Results field.
func (o *LinkResponseList) SetResults(v []LinkResponse) {
	o.Results = &v
}

func (o LinkResponseList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullableLinkResponseList struct {
	value *LinkResponseList
	isSet bool
}

func (v NullableLinkResponseList) Get() *LinkResponseList {
	return v.value
}

func (v *NullableLinkResponseList) Set(val *LinkResponseList) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkResponseList) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkResponseList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkResponseList(val *LinkResponseList) *NullableLinkResponseList {
	return &NullableLinkResponseList{value: val, isSet: true}
}

func (v NullableLinkResponseList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkResponseList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
