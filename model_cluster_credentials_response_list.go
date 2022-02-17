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

// ClusterCredentialsResponseList struct for ClusterCredentialsResponseList
type ClusterCredentialsResponseList struct {
	Results *[]ClusterCredentialsResponse `json:"results,omitempty"`
}

// NewClusterCredentialsResponseList instantiates a new ClusterCredentialsResponseList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterCredentialsResponseList() *ClusterCredentialsResponseList {
	this := ClusterCredentialsResponseList{}
	return &this
}

// NewClusterCredentialsResponseListWithDefaults instantiates a new ClusterCredentialsResponseList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterCredentialsResponseListWithDefaults() *ClusterCredentialsResponseList {
	this := ClusterCredentialsResponseList{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *ClusterCredentialsResponseList) GetResults() []ClusterCredentialsResponse {
	if o == nil || o.Results == nil {
		var ret []ClusterCredentialsResponse
		return ret
	}
	return *o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterCredentialsResponseList) GetResultsOk() (*[]ClusterCredentialsResponse, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *ClusterCredentialsResponseList) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []ClusterCredentialsResponse and assigns it to the Results field.
func (o *ClusterCredentialsResponseList) SetResults(v []ClusterCredentialsResponse) {
	o.Results = &v
}

func (o ClusterCredentialsResponseList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullableClusterCredentialsResponseList struct {
	value *ClusterCredentialsResponseList
	isSet bool
}

func (v NullableClusterCredentialsResponseList) Get() *ClusterCredentialsResponseList {
	return v.value
}

func (v *NullableClusterCredentialsResponseList) Set(val *ClusterCredentialsResponseList) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterCredentialsResponseList) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterCredentialsResponseList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterCredentialsResponseList(val *ClusterCredentialsResponseList) *NullableClusterCredentialsResponseList {
	return &NullableClusterCredentialsResponseList{value: val, isSet: true}
}

func (v NullableClusterCredentialsResponseList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterCredentialsResponseList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


