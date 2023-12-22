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

// checks if the HelmResponseAllOfValuesOverride type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HelmResponseAllOfValuesOverride{}

// HelmResponseAllOfValuesOverride Specify helm values you want to set or override
type HelmResponseAllOfValuesOverride struct {
	// The input is in json array format: [ [$KEY,$VALUE], [...] ]
	Set [][]string `json:"set,omitempty"`
	// The input is in json array format: [ [$KEY,$VALUE], [...] ]
	SetString [][]string `json:"set_string,omitempty"`
	// The input is in json array format: [ [$KEY,$VALUE], [...] ]
	SetJson [][]string                                  `json:"set_json,omitempty"`
	File    NullableHelmResponseAllOfValuesOverrideFile `json:"file,omitempty"`
}

// NewHelmResponseAllOfValuesOverride instantiates a new HelmResponseAllOfValuesOverride object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHelmResponseAllOfValuesOverride() *HelmResponseAllOfValuesOverride {
	this := HelmResponseAllOfValuesOverride{}
	return &this
}

// NewHelmResponseAllOfValuesOverrideWithDefaults instantiates a new HelmResponseAllOfValuesOverride object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHelmResponseAllOfValuesOverrideWithDefaults() *HelmResponseAllOfValuesOverride {
	this := HelmResponseAllOfValuesOverride{}
	return &this
}

// GetSet returns the Set field value if set, zero value otherwise.
func (o *HelmResponseAllOfValuesOverride) GetSet() [][]string {
	if o == nil || IsNil(o.Set) {
		var ret [][]string
		return ret
	}
	return o.Set
}

// GetSetOk returns a tuple with the Set field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HelmResponseAllOfValuesOverride) GetSetOk() ([][]string, bool) {
	if o == nil || IsNil(o.Set) {
		return nil, false
	}
	return o.Set, true
}

// HasSet returns a boolean if a field has been set.
func (o *HelmResponseAllOfValuesOverride) HasSet() bool {
	if o != nil && !IsNil(o.Set) {
		return true
	}

	return false
}

// SetSet gets a reference to the given [][]string and assigns it to the Set field.
func (o *HelmResponseAllOfValuesOverride) SetSet(v [][]string) {
	o.Set = v
}

// GetSetString returns the SetString field value if set, zero value otherwise.
func (o *HelmResponseAllOfValuesOverride) GetSetString() [][]string {
	if o == nil || IsNil(o.SetString) {
		var ret [][]string
		return ret
	}
	return o.SetString
}

// GetSetStringOk returns a tuple with the SetString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HelmResponseAllOfValuesOverride) GetSetStringOk() ([][]string, bool) {
	if o == nil || IsNil(o.SetString) {
		return nil, false
	}
	return o.SetString, true
}

// HasSetString returns a boolean if a field has been set.
func (o *HelmResponseAllOfValuesOverride) HasSetString() bool {
	if o != nil && !IsNil(o.SetString) {
		return true
	}

	return false
}

// SetSetString gets a reference to the given [][]string and assigns it to the SetString field.
func (o *HelmResponseAllOfValuesOverride) SetSetString(v [][]string) {
	o.SetString = v
}

// GetSetJson returns the SetJson field value if set, zero value otherwise.
func (o *HelmResponseAllOfValuesOverride) GetSetJson() [][]string {
	if o == nil || IsNil(o.SetJson) {
		var ret [][]string
		return ret
	}
	return o.SetJson
}

// GetSetJsonOk returns a tuple with the SetJson field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HelmResponseAllOfValuesOverride) GetSetJsonOk() ([][]string, bool) {
	if o == nil || IsNil(o.SetJson) {
		return nil, false
	}
	return o.SetJson, true
}

// HasSetJson returns a boolean if a field has been set.
func (o *HelmResponseAllOfValuesOverride) HasSetJson() bool {
	if o != nil && !IsNil(o.SetJson) {
		return true
	}

	return false
}

// SetSetJson gets a reference to the given [][]string and assigns it to the SetJson field.
func (o *HelmResponseAllOfValuesOverride) SetSetJson(v [][]string) {
	o.SetJson = v
}

// GetFile returns the File field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HelmResponseAllOfValuesOverride) GetFile() HelmResponseAllOfValuesOverrideFile {
	if o == nil || IsNil(o.File.Get()) {
		var ret HelmResponseAllOfValuesOverrideFile
		return ret
	}
	return *o.File.Get()
}

// GetFileOk returns a tuple with the File field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HelmResponseAllOfValuesOverride) GetFileOk() (*HelmResponseAllOfValuesOverrideFile, bool) {
	if o == nil {
		return nil, false
	}
	return o.File.Get(), o.File.IsSet()
}

// HasFile returns a boolean if a field has been set.
func (o *HelmResponseAllOfValuesOverride) HasFile() bool {
	if o != nil && o.File.IsSet() {
		return true
	}

	return false
}

// SetFile gets a reference to the given NullableHelmResponseAllOfValuesOverrideFile and assigns it to the File field.
func (o *HelmResponseAllOfValuesOverride) SetFile(v HelmResponseAllOfValuesOverrideFile) {
	o.File.Set(&v)
}

// SetFileNil sets the value for File to be an explicit nil
func (o *HelmResponseAllOfValuesOverride) SetFileNil() {
	o.File.Set(nil)
}

// UnsetFile ensures that no value is present for File, not even an explicit nil
func (o *HelmResponseAllOfValuesOverride) UnsetFile() {
	o.File.Unset()
}

func (o HelmResponseAllOfValuesOverride) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HelmResponseAllOfValuesOverride) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Set) {
		toSerialize["set"] = o.Set
	}
	if !IsNil(o.SetString) {
		toSerialize["set_string"] = o.SetString
	}
	if !IsNil(o.SetJson) {
		toSerialize["set_json"] = o.SetJson
	}
	if o.File.IsSet() {
		toSerialize["file"] = o.File.Get()
	}
	return toSerialize, nil
}

type NullableHelmResponseAllOfValuesOverride struct {
	value *HelmResponseAllOfValuesOverride
	isSet bool
}

func (v NullableHelmResponseAllOfValuesOverride) Get() *HelmResponseAllOfValuesOverride {
	return v.value
}

func (v *NullableHelmResponseAllOfValuesOverride) Set(val *HelmResponseAllOfValuesOverride) {
	v.value = val
	v.isSet = true
}

func (v NullableHelmResponseAllOfValuesOverride) IsSet() bool {
	return v.isSet
}

func (v *NullableHelmResponseAllOfValuesOverride) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHelmResponseAllOfValuesOverride(val *HelmResponseAllOfValuesOverride) *NullableHelmResponseAllOfValuesOverride {
	return &NullableHelmResponseAllOfValuesOverride{value: val, isSet: true}
}

func (v NullableHelmResponseAllOfValuesOverride) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHelmResponseAllOfValuesOverride) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
