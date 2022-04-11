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

// ClusterResponseList struct for ClusterResponseList
type ClusterResponseList struct {
	Results []Cluster `json:"results,omitempty"`
}

// NewClusterResponseList instantiates a new ClusterResponseList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterResponseList() *ClusterResponseList {
	this := ClusterResponseList{}
	return &this
}

// NewClusterResponseListWithDefaults instantiates a new ClusterResponseList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterResponseListWithDefaults() *ClusterResponseList {
	this := ClusterResponseList{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *ClusterResponseList) GetResults() []Cluster {
	if o == nil || o.Results == nil {
		var ret []Cluster
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterResponseList) GetResultsOk() ([]Cluster, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *ClusterResponseList) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []Cluster and assigns it to the Results field.
func (o *ClusterResponseList) SetResults(v []Cluster) {
	o.Results = v
}

func (o ClusterResponseList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullableClusterResponseList struct {
	value *ClusterResponseList
	isSet bool
}

func (v NullableClusterResponseList) Get() *ClusterResponseList {
	return v.value
}

func (v *NullableClusterResponseList) Set(val *ClusterResponseList) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterResponseList) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterResponseList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterResponseList(val *ClusterResponseList) *NullableClusterResponseList {
	return &NullableClusterResponseList{value: val, isSet: true}
}

func (v NullableClusterResponseList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterResponseList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
