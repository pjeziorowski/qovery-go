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

// CustomDomainResponseAllOf struct for CustomDomainResponseAllOf
type CustomDomainResponseAllOf struct {
	// URL provided by Qovery. You must create a CNAME on your DNS provider using that URL
	ValidationDomain *string `json:"validation_domain,omitempty"`
	Status           *string `json:"status,omitempty"`
}

// NewCustomDomainResponseAllOf instantiates a new CustomDomainResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomDomainResponseAllOf() *CustomDomainResponseAllOf {
	this := CustomDomainResponseAllOf{}
	return &this
}

// NewCustomDomainResponseAllOfWithDefaults instantiates a new CustomDomainResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomDomainResponseAllOfWithDefaults() *CustomDomainResponseAllOf {
	this := CustomDomainResponseAllOf{}
	return &this
}

// GetValidationDomain returns the ValidationDomain field value if set, zero value otherwise.
func (o *CustomDomainResponseAllOf) GetValidationDomain() string {
	if o == nil || o.ValidationDomain == nil {
		var ret string
		return ret
	}
	return *o.ValidationDomain
}

// GetValidationDomainOk returns a tuple with the ValidationDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomDomainResponseAllOf) GetValidationDomainOk() (*string, bool) {
	if o == nil || o.ValidationDomain == nil {
		return nil, false
	}
	return o.ValidationDomain, true
}

// HasValidationDomain returns a boolean if a field has been set.
func (o *CustomDomainResponseAllOf) HasValidationDomain() bool {
	if o != nil && o.ValidationDomain != nil {
		return true
	}

	return false
}

// SetValidationDomain gets a reference to the given string and assigns it to the ValidationDomain field.
func (o *CustomDomainResponseAllOf) SetValidationDomain(v string) {
	o.ValidationDomain = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CustomDomainResponseAllOf) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomDomainResponseAllOf) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CustomDomainResponseAllOf) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *CustomDomainResponseAllOf) SetStatus(v string) {
	o.Status = &v
}

func (o CustomDomainResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ValidationDomain != nil {
		toSerialize["validation_domain"] = o.ValidationDomain
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableCustomDomainResponseAllOf struct {
	value *CustomDomainResponseAllOf
	isSet bool
}

func (v NullableCustomDomainResponseAllOf) Get() *CustomDomainResponseAllOf {
	return v.value
}

func (v *NullableCustomDomainResponseAllOf) Set(val *CustomDomainResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomDomainResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomDomainResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomDomainResponseAllOf(val *CustomDomainResponseAllOf) *NullableCustomDomainResponseAllOf {
	return &NullableCustomDomainResponseAllOf{value: val, isSet: true}
}

func (v NullableCustomDomainResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomDomainResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
