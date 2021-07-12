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

// OverriddenSecret struct for OverriddenSecret
type OverriddenSecret struct {
	Id    *string `json:"id,omitempty"`
	Key   *string `json:"key,omitempty"`
	Scope *string `json:"scope,omitempty"`
}

// NewOverriddenSecret instantiates a new OverriddenSecret object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOverriddenSecret() *OverriddenSecret {
	this := OverriddenSecret{}
	return &this
}

// NewOverriddenSecretWithDefaults instantiates a new OverriddenSecret object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOverriddenSecretWithDefaults() *OverriddenSecret {
	this := OverriddenSecret{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OverriddenSecret) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OverriddenSecret) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OverriddenSecret) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OverriddenSecret) SetId(v string) {
	o.Id = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *OverriddenSecret) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OverriddenSecret) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *OverriddenSecret) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *OverriddenSecret) SetKey(v string) {
	o.Key = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *OverriddenSecret) GetScope() string {
	if o == nil || o.Scope == nil {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OverriddenSecret) GetScopeOk() (*string, bool) {
	if o == nil || o.Scope == nil {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *OverriddenSecret) HasScope() bool {
	if o != nil && o.Scope != nil {
		return true
	}

	return false
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *OverriddenSecret) SetScope(v string) {
	o.Scope = &v
}

func (o OverriddenSecret) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if o.Scope != nil {
		toSerialize["scope"] = o.Scope
	}
	return json.Marshal(toSerialize)
}

type NullableOverriddenSecret struct {
	value *OverriddenSecret
	isSet bool
}

func (v NullableOverriddenSecret) Get() *OverriddenSecret {
	return v.value
}

func (v *NullableOverriddenSecret) Set(val *OverriddenSecret) {
	v.value = val
	v.isSet = true
}

func (v NullableOverriddenSecret) IsSet() bool {
	return v.isSet
}

func (v *NullableOverriddenSecret) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOverriddenSecret(val *OverriddenSecret) *NullableOverriddenSecret {
	return &NullableOverriddenSecret{value: val, isSet: true}
}

func (v NullableOverriddenSecret) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOverriddenSecret) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
