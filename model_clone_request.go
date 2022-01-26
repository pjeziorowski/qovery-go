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

// CloneRequest struct for CloneRequest
type CloneRequest struct {
	// name is case insensitive
	Name string `json:"name"`
	Cluster string `json:"cluster"`
}

// NewCloneRequest instantiates a new CloneRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloneRequest(name string, cluster string) *CloneRequest {
	this := CloneRequest{}
	this.Name = name
	this.Cluster = cluster
	return &this
}

// NewCloneRequestWithDefaults instantiates a new CloneRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloneRequestWithDefaults() *CloneRequest {
	this := CloneRequest{}
	return &this
}

// GetName returns the Name field value
func (o *CloneRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CloneRequest) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CloneRequest) SetName(v string) {
	o.Name = v
}

// GetCluster returns the Cluster field value
func (o *CloneRequest) GetCluster() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value
// and a boolean to check if the value has been set.
func (o *CloneRequest) GetClusterOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Cluster, true
}

// SetCluster sets field value
func (o *CloneRequest) SetCluster(v string) {
	o.Cluster = v
}

func (o CloneRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["cluster"] = o.Cluster
	}
	return json.Marshal(toSerialize)
}

type NullableCloneRequest struct {
	value *CloneRequest
	isSet bool
}

func (v NullableCloneRequest) Get() *CloneRequest {
	return v.value
}

func (v *NullableCloneRequest) Set(val *CloneRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCloneRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCloneRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloneRequest(val *CloneRequest) *NullableCloneRequest {
	return &NullableCloneRequest{value: val, isSet: true}
}

func (v NullableCloneRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloneRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


