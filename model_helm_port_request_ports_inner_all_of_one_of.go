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

// checks if the HelmPortRequestPortsInnerAllOfOneOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HelmPortRequestPortsInnerAllOfOneOf{}

// HelmPortRequestPortsInnerAllOfOneOf struct for HelmPortRequestPortsInnerAllOfOneOf
type HelmPortRequestPortsInnerAllOfOneOf struct {
	ServiceSelectors     []KubernetesSelector `json:"service_selectors,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HelmPortRequestPortsInnerAllOfOneOf HelmPortRequestPortsInnerAllOfOneOf

// NewHelmPortRequestPortsInnerAllOfOneOf instantiates a new HelmPortRequestPortsInnerAllOfOneOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHelmPortRequestPortsInnerAllOfOneOf() *HelmPortRequestPortsInnerAllOfOneOf {
	this := HelmPortRequestPortsInnerAllOfOneOf{}
	return &this
}

// NewHelmPortRequestPortsInnerAllOfOneOfWithDefaults instantiates a new HelmPortRequestPortsInnerAllOfOneOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHelmPortRequestPortsInnerAllOfOneOfWithDefaults() *HelmPortRequestPortsInnerAllOfOneOf {
	this := HelmPortRequestPortsInnerAllOfOneOf{}
	return &this
}

// GetServiceSelectors returns the ServiceSelectors field value if set, zero value otherwise.
func (o *HelmPortRequestPortsInnerAllOfOneOf) GetServiceSelectors() []KubernetesSelector {
	if o == nil || IsNil(o.ServiceSelectors) {
		var ret []KubernetesSelector
		return ret
	}
	return o.ServiceSelectors
}

// GetServiceSelectorsOk returns a tuple with the ServiceSelectors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HelmPortRequestPortsInnerAllOfOneOf) GetServiceSelectorsOk() ([]KubernetesSelector, bool) {
	if o == nil || IsNil(o.ServiceSelectors) {
		return nil, false
	}
	return o.ServiceSelectors, true
}

// HasServiceSelectors returns a boolean if a field has been set.
func (o *HelmPortRequestPortsInnerAllOfOneOf) HasServiceSelectors() bool {
	if o != nil && !IsNil(o.ServiceSelectors) {
		return true
	}

	return false
}

// SetServiceSelectors gets a reference to the given []KubernetesSelector and assigns it to the ServiceSelectors field.
func (o *HelmPortRequestPortsInnerAllOfOneOf) SetServiceSelectors(v []KubernetesSelector) {
	o.ServiceSelectors = v
}

func (o HelmPortRequestPortsInnerAllOfOneOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HelmPortRequestPortsInnerAllOfOneOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ServiceSelectors) {
		toSerialize["service_selectors"] = o.ServiceSelectors
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *HelmPortRequestPortsInnerAllOfOneOf) UnmarshalJSON(data []byte) (err error) {
	varHelmPortRequestPortsInnerAllOfOneOf := _HelmPortRequestPortsInnerAllOfOneOf{}

	err = json.Unmarshal(data, &varHelmPortRequestPortsInnerAllOfOneOf)

	if err != nil {
		return err
	}

	*o = HelmPortRequestPortsInnerAllOfOneOf(varHelmPortRequestPortsInnerAllOfOneOf)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "service_selectors")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHelmPortRequestPortsInnerAllOfOneOf struct {
	value *HelmPortRequestPortsInnerAllOfOneOf
	isSet bool
}

func (v NullableHelmPortRequestPortsInnerAllOfOneOf) Get() *HelmPortRequestPortsInnerAllOfOneOf {
	return v.value
}

func (v *NullableHelmPortRequestPortsInnerAllOfOneOf) Set(val *HelmPortRequestPortsInnerAllOfOneOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHelmPortRequestPortsInnerAllOfOneOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHelmPortRequestPortsInnerAllOfOneOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHelmPortRequestPortsInnerAllOfOneOf(val *HelmPortRequestPortsInnerAllOfOneOf) *NullableHelmPortRequestPortsInnerAllOfOneOf {
	return &NullableHelmPortRequestPortsInnerAllOfOneOf{value: val, isSet: true}
}

func (v NullableHelmPortRequestPortsInnerAllOfOneOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHelmPortRequestPortsInnerAllOfOneOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
