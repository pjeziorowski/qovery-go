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

// VariableImportRequestVars struct for VariableImportRequestVars
type VariableImportRequestVars struct {
	Name string `json:"name"`
	Value string `json:"value"`
	Scope string `json:"scope"`
	IsSecret bool `json:"is_secret"`
}

// NewVariableImportRequestVars instantiates a new VariableImportRequestVars object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVariableImportRequestVars(name string, value string, scope string, isSecret bool) *VariableImportRequestVars {
	this := VariableImportRequestVars{}
	this.Name = name
	this.Value = value
	this.Scope = scope
	this.IsSecret = isSecret
	return &this
}

// NewVariableImportRequestVarsWithDefaults instantiates a new VariableImportRequestVars object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVariableImportRequestVarsWithDefaults() *VariableImportRequestVars {
	this := VariableImportRequestVars{}
	return &this
}

// GetName returns the Name field value
func (o *VariableImportRequestVars) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *VariableImportRequestVars) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *VariableImportRequestVars) SetName(v string) {
	o.Name = v
}

// GetValue returns the Value field value
func (o *VariableImportRequestVars) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *VariableImportRequestVars) GetValueOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *VariableImportRequestVars) SetValue(v string) {
	o.Value = v
}

// GetScope returns the Scope field value
func (o *VariableImportRequestVars) GetScope() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *VariableImportRequestVars) GetScopeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Scope, true
}

// SetScope sets field value
func (o *VariableImportRequestVars) SetScope(v string) {
	o.Scope = v
}

// GetIsSecret returns the IsSecret field value
func (o *VariableImportRequestVars) GetIsSecret() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsSecret
}

// GetIsSecretOk returns a tuple with the IsSecret field value
// and a boolean to check if the value has been set.
func (o *VariableImportRequestVars) GetIsSecretOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsSecret, true
}

// SetIsSecret sets field value
func (o *VariableImportRequestVars) SetIsSecret(v bool) {
	o.IsSecret = v
}

func (o VariableImportRequestVars) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["value"] = o.Value
	}
	if true {
		toSerialize["scope"] = o.Scope
	}
	if true {
		toSerialize["is_secret"] = o.IsSecret
	}
	return json.Marshal(toSerialize)
}

type NullableVariableImportRequestVars struct {
	value *VariableImportRequestVars
	isSet bool
}

func (v NullableVariableImportRequestVars) Get() *VariableImportRequestVars {
	return v.value
}

func (v *NullableVariableImportRequestVars) Set(val *VariableImportRequestVars) {
	v.value = val
	v.isSet = true
}

func (v NullableVariableImportRequestVars) IsSet() bool {
	return v.isSet
}

func (v *NullableVariableImportRequestVars) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVariableImportRequestVars(val *VariableImportRequestVars) *NullableVariableImportRequestVars {
	return &NullableVariableImportRequestVars{value: val, isSet: true}
}

func (v NullableVariableImportRequestVars) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVariableImportRequestVars) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


