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

// checks if the HelmPortRequestPortsInnerAllOfOneOf1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HelmPortRequestPortsInnerAllOfOneOf1{}

// HelmPortRequestPortsInnerAllOfOneOf1 struct for HelmPortRequestPortsInnerAllOfOneOf1
type HelmPortRequestPortsInnerAllOfOneOf1 struct {
	ServiceName          *string `json:"service_name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HelmPortRequestPortsInnerAllOfOneOf1 HelmPortRequestPortsInnerAllOfOneOf1

// NewHelmPortRequestPortsInnerAllOfOneOf1 instantiates a new HelmPortRequestPortsInnerAllOfOneOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHelmPortRequestPortsInnerAllOfOneOf1() *HelmPortRequestPortsInnerAllOfOneOf1 {
	this := HelmPortRequestPortsInnerAllOfOneOf1{}
	return &this
}

// NewHelmPortRequestPortsInnerAllOfOneOf1WithDefaults instantiates a new HelmPortRequestPortsInnerAllOfOneOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHelmPortRequestPortsInnerAllOfOneOf1WithDefaults() *HelmPortRequestPortsInnerAllOfOneOf1 {
	this := HelmPortRequestPortsInnerAllOfOneOf1{}
	return &this
}

// GetServiceName returns the ServiceName field value if set, zero value otherwise.
func (o *HelmPortRequestPortsInnerAllOfOneOf1) GetServiceName() string {
	if o == nil || IsNil(o.ServiceName) {
		var ret string
		return ret
	}
	return *o.ServiceName
}

// GetServiceNameOk returns a tuple with the ServiceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HelmPortRequestPortsInnerAllOfOneOf1) GetServiceNameOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceName) {
		return nil, false
	}
	return o.ServiceName, true
}

// HasServiceName returns a boolean if a field has been set.
func (o *HelmPortRequestPortsInnerAllOfOneOf1) HasServiceName() bool {
	if o != nil && !IsNil(o.ServiceName) {
		return true
	}

	return false
}

// SetServiceName gets a reference to the given string and assigns it to the ServiceName field.
func (o *HelmPortRequestPortsInnerAllOfOneOf1) SetServiceName(v string) {
	o.ServiceName = &v
}

func (o HelmPortRequestPortsInnerAllOfOneOf1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HelmPortRequestPortsInnerAllOfOneOf1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ServiceName) {
		toSerialize["service_name"] = o.ServiceName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *HelmPortRequestPortsInnerAllOfOneOf1) UnmarshalJSON(data []byte) (err error) {
	varHelmPortRequestPortsInnerAllOfOneOf1 := _HelmPortRequestPortsInnerAllOfOneOf1{}

	err = json.Unmarshal(data, &varHelmPortRequestPortsInnerAllOfOneOf1)

	if err != nil {
		return err
	}

	*o = HelmPortRequestPortsInnerAllOfOneOf1(varHelmPortRequestPortsInnerAllOfOneOf1)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "service_name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHelmPortRequestPortsInnerAllOfOneOf1 struct {
	value *HelmPortRequestPortsInnerAllOfOneOf1
	isSet bool
}

func (v NullableHelmPortRequestPortsInnerAllOfOneOf1) Get() *HelmPortRequestPortsInnerAllOfOneOf1 {
	return v.value
}

func (v *NullableHelmPortRequestPortsInnerAllOfOneOf1) Set(val *HelmPortRequestPortsInnerAllOfOneOf1) {
	v.value = val
	v.isSet = true
}

func (v NullableHelmPortRequestPortsInnerAllOfOneOf1) IsSet() bool {
	return v.isSet
}

func (v *NullableHelmPortRequestPortsInnerAllOfOneOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHelmPortRequestPortsInnerAllOfOneOf1(val *HelmPortRequestPortsInnerAllOfOneOf1) *NullableHelmPortRequestPortsInnerAllOfOneOf1 {
	return &NullableHelmPortRequestPortsInnerAllOfOneOf1{value: val, isSet: true}
}

func (v NullableHelmPortRequestPortsInnerAllOfOneOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHelmPortRequestPortsInnerAllOfOneOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
