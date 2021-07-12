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

// ClusterReadinessStatus struct for ClusterReadinessStatus
type ClusterReadinessStatus struct {
	IsReady *bool `json:"is_ready,omitempty"`
}

// NewClusterReadinessStatus instantiates a new ClusterReadinessStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterReadinessStatus() *ClusterReadinessStatus {
	this := ClusterReadinessStatus{}
	return &this
}

// NewClusterReadinessStatusWithDefaults instantiates a new ClusterReadinessStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterReadinessStatusWithDefaults() *ClusterReadinessStatus {
	this := ClusterReadinessStatus{}
	return &this
}

// GetIsReady returns the IsReady field value if set, zero value otherwise.
func (o *ClusterReadinessStatus) GetIsReady() bool {
	if o == nil || o.IsReady == nil {
		var ret bool
		return ret
	}
	return *o.IsReady
}

// GetIsReadyOk returns a tuple with the IsReady field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterReadinessStatus) GetIsReadyOk() (*bool, bool) {
	if o == nil || o.IsReady == nil {
		return nil, false
	}
	return o.IsReady, true
}

// HasIsReady returns a boolean if a field has been set.
func (o *ClusterReadinessStatus) HasIsReady() bool {
	if o != nil && o.IsReady != nil {
		return true
	}

	return false
}

// SetIsReady gets a reference to the given bool and assigns it to the IsReady field.
func (o *ClusterReadinessStatus) SetIsReady(v bool) {
	o.IsReady = &v
}

func (o ClusterReadinessStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IsReady != nil {
		toSerialize["is_ready"] = o.IsReady
	}
	return json.Marshal(toSerialize)
}

type NullableClusterReadinessStatus struct {
	value *ClusterReadinessStatus
	isSet bool
}

func (v NullableClusterReadinessStatus) Get() *ClusterReadinessStatus {
	return v.value
}

func (v *NullableClusterReadinessStatus) Set(val *ClusterReadinessStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterReadinessStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterReadinessStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterReadinessStatus(val *ClusterReadinessStatus) *NullableClusterReadinessStatus {
	return &NullableClusterReadinessStatus{value: val, isSet: true}
}

func (v NullableClusterReadinessStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterReadinessStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
