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

// checks if the KubernetesMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubernetesMetadata{}

// KubernetesMetadata struct for KubernetesMetadata
type KubernetesMetadata struct {
	Name                 *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesMetadata KubernetesMetadata

// NewKubernetesMetadata instantiates a new KubernetesMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesMetadata() *KubernetesMetadata {
	this := KubernetesMetadata{}
	return &this
}

// NewKubernetesMetadataWithDefaults instantiates a new KubernetesMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesMetadataWithDefaults() *KubernetesMetadata {
	this := KubernetesMetadata{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *KubernetesMetadata) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesMetadata) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *KubernetesMetadata) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *KubernetesMetadata) SetName(v string) {
	o.Name = &v
}

func (o KubernetesMetadata) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubernetesMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *KubernetesMetadata) UnmarshalJSON(data []byte) (err error) {
	varKubernetesMetadata := _KubernetesMetadata{}

	err = json.Unmarshal(data, &varKubernetesMetadata)

	if err != nil {
		return err
	}

	*o = KubernetesMetadata(varKubernetesMetadata)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableKubernetesMetadata struct {
	value *KubernetesMetadata
	isSet bool
}

func (v NullableKubernetesMetadata) Get() *KubernetesMetadata {
	return v.value
}

func (v *NullableKubernetesMetadata) Set(val *KubernetesMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesMetadata(val *KubernetesMetadata) *NullableKubernetesMetadata {
	return &NullableKubernetesMetadata{value: val, isSet: true}
}

func (v NullableKubernetesMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
