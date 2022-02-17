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

// EnvironmentApplicationsSupportedLanguageList struct for EnvironmentApplicationsSupportedLanguageList
type EnvironmentApplicationsSupportedLanguageList struct {
	Results *[]EnvironmentApplicationsSupportedLanguage `json:"results,omitempty"`
}

// NewEnvironmentApplicationsSupportedLanguageList instantiates a new EnvironmentApplicationsSupportedLanguageList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentApplicationsSupportedLanguageList() *EnvironmentApplicationsSupportedLanguageList {
	this := EnvironmentApplicationsSupportedLanguageList{}
	return &this
}

// NewEnvironmentApplicationsSupportedLanguageListWithDefaults instantiates a new EnvironmentApplicationsSupportedLanguageList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentApplicationsSupportedLanguageListWithDefaults() *EnvironmentApplicationsSupportedLanguageList {
	this := EnvironmentApplicationsSupportedLanguageList{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *EnvironmentApplicationsSupportedLanguageList) GetResults() []EnvironmentApplicationsSupportedLanguage {
	if o == nil || o.Results == nil {
		var ret []EnvironmentApplicationsSupportedLanguage
		return ret
	}
	return *o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentApplicationsSupportedLanguageList) GetResultsOk() (*[]EnvironmentApplicationsSupportedLanguage, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *EnvironmentApplicationsSupportedLanguageList) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []EnvironmentApplicationsSupportedLanguage and assigns it to the Results field.
func (o *EnvironmentApplicationsSupportedLanguageList) SetResults(v []EnvironmentApplicationsSupportedLanguage) {
	o.Results = &v
}

func (o EnvironmentApplicationsSupportedLanguageList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullableEnvironmentApplicationsSupportedLanguageList struct {
	value *EnvironmentApplicationsSupportedLanguageList
	isSet bool
}

func (v NullableEnvironmentApplicationsSupportedLanguageList) Get() *EnvironmentApplicationsSupportedLanguageList {
	return v.value
}

func (v *NullableEnvironmentApplicationsSupportedLanguageList) Set(val *EnvironmentApplicationsSupportedLanguageList) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentApplicationsSupportedLanguageList) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentApplicationsSupportedLanguageList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentApplicationsSupportedLanguageList(val *EnvironmentApplicationsSupportedLanguageList) *NullableEnvironmentApplicationsSupportedLanguageList {
	return &NullableEnvironmentApplicationsSupportedLanguageList{value: val, isSet: true}
}

func (v NullableEnvironmentApplicationsSupportedLanguageList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentApplicationsSupportedLanguageList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


