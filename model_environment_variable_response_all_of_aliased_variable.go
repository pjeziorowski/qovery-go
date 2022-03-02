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

// EnvironmentVariableResponseAllOfAliasedVariable struct for EnvironmentVariableResponseAllOfAliasedVariable
type EnvironmentVariableResponseAllOfAliasedVariable struct {
	Id    *string `json:"id,omitempty"`
	Key   *string `json:"key,omitempty"`
	Value *string `json:"value,omitempty"`
	Scope *string `json:"scope,omitempty"`
}

// NewEnvironmentVariableResponseAllOfAliasedVariable instantiates a new EnvironmentVariableResponseAllOfAliasedVariable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentVariableResponseAllOfAliasedVariable() *EnvironmentVariableResponseAllOfAliasedVariable {
	this := EnvironmentVariableResponseAllOfAliasedVariable{}
	return &this
}

// NewEnvironmentVariableResponseAllOfAliasedVariableWithDefaults instantiates a new EnvironmentVariableResponseAllOfAliasedVariable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentVariableResponseAllOfAliasedVariableWithDefaults() *EnvironmentVariableResponseAllOfAliasedVariable {
	this := EnvironmentVariableResponseAllOfAliasedVariable{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EnvironmentVariableResponseAllOfAliasedVariable) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentVariableResponseAllOfAliasedVariable) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EnvironmentVariableResponseAllOfAliasedVariable) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *EnvironmentVariableResponseAllOfAliasedVariable) SetId(v string) {
	o.Id = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *EnvironmentVariableResponseAllOfAliasedVariable) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentVariableResponseAllOfAliasedVariable) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *EnvironmentVariableResponseAllOfAliasedVariable) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *EnvironmentVariableResponseAllOfAliasedVariable) SetKey(v string) {
	o.Key = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *EnvironmentVariableResponseAllOfAliasedVariable) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentVariableResponseAllOfAliasedVariable) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *EnvironmentVariableResponseAllOfAliasedVariable) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *EnvironmentVariableResponseAllOfAliasedVariable) SetValue(v string) {
	o.Value = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *EnvironmentVariableResponseAllOfAliasedVariable) GetScope() string {
	if o == nil || o.Scope == nil {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentVariableResponseAllOfAliasedVariable) GetScopeOk() (*string, bool) {
	if o == nil || o.Scope == nil {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *EnvironmentVariableResponseAllOfAliasedVariable) HasScope() bool {
	if o != nil && o.Scope != nil {
		return true
	}

	return false
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *EnvironmentVariableResponseAllOfAliasedVariable) SetScope(v string) {
	o.Scope = &v
}

func (o EnvironmentVariableResponseAllOfAliasedVariable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.Scope != nil {
		toSerialize["scope"] = o.Scope
	}
	return json.Marshal(toSerialize)
}

type NullableEnvironmentVariableResponseAllOfAliasedVariable struct {
	value *EnvironmentVariableResponseAllOfAliasedVariable
	isSet bool
}

func (v NullableEnvironmentVariableResponseAllOfAliasedVariable) Get() *EnvironmentVariableResponseAllOfAliasedVariable {
	return v.value
}

func (v *NullableEnvironmentVariableResponseAllOfAliasedVariable) Set(val *EnvironmentVariableResponseAllOfAliasedVariable) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentVariableResponseAllOfAliasedVariable) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentVariableResponseAllOfAliasedVariable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentVariableResponseAllOfAliasedVariable(val *EnvironmentVariableResponseAllOfAliasedVariable) *NullableEnvironmentVariableResponseAllOfAliasedVariable {
	return &NullableEnvironmentVariableResponseAllOfAliasedVariable{value: val, isSet: true}
}

func (v NullableEnvironmentVariableResponseAllOfAliasedVariable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentVariableResponseAllOfAliasedVariable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
