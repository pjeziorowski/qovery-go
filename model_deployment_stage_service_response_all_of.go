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

// DeploymentStageServiceResponseAllOf struct for DeploymentStageServiceResponseAllOf
type DeploymentStageServiceResponseAllOf struct {
	// id of the service attached to the stage
	ServiceId *string `json:"service_id,omitempty"`
}

// NewDeploymentStageServiceResponseAllOf instantiates a new DeploymentStageServiceResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploymentStageServiceResponseAllOf() *DeploymentStageServiceResponseAllOf {
	this := DeploymentStageServiceResponseAllOf{}
	return &this
}

// NewDeploymentStageServiceResponseAllOfWithDefaults instantiates a new DeploymentStageServiceResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentStageServiceResponseAllOfWithDefaults() *DeploymentStageServiceResponseAllOf {
	this := DeploymentStageServiceResponseAllOf{}
	return &this
}

// GetServiceId returns the ServiceId field value if set, zero value otherwise.
func (o *DeploymentStageServiceResponseAllOf) GetServiceId() string {
	if o == nil || o.ServiceId == nil {
		var ret string
		return ret
	}
	return *o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentStageServiceResponseAllOf) GetServiceIdOk() (*string, bool) {
	if o == nil || o.ServiceId == nil {
		return nil, false
	}
	return o.ServiceId, true
}

// HasServiceId returns a boolean if a field has been set.
func (o *DeploymentStageServiceResponseAllOf) HasServiceId() bool {
	if o != nil && o.ServiceId != nil {
		return true
	}

	return false
}

// SetServiceId gets a reference to the given string and assigns it to the ServiceId field.
func (o *DeploymentStageServiceResponseAllOf) SetServiceId(v string) {
	o.ServiceId = &v
}

func (o DeploymentStageServiceResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ServiceId != nil {
		toSerialize["service_id"] = o.ServiceId
	}
	return json.Marshal(toSerialize)
}

type NullableDeploymentStageServiceResponseAllOf struct {
	value *DeploymentStageServiceResponseAllOf
	isSet bool
}

func (v NullableDeploymentStageServiceResponseAllOf) Get() *DeploymentStageServiceResponseAllOf {
	return v.value
}

func (v *NullableDeploymentStageServiceResponseAllOf) Set(val *DeploymentStageServiceResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploymentStageServiceResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploymentStageServiceResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploymentStageServiceResponseAllOf(val *DeploymentStageServiceResponseAllOf) *NullableDeploymentStageServiceResponseAllOf {
	return &NullableDeploymentStageServiceResponseAllOf{value: val, isSet: true}
}

func (v NullableDeploymentStageServiceResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploymentStageServiceResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
