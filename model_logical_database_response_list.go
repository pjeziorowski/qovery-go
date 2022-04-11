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

// LogicalDatabaseResponseList struct for LogicalDatabaseResponseList
type LogicalDatabaseResponseList struct {
	Results []LogicalDatabase `json:"results,omitempty"`
}

// NewLogicalDatabaseResponseList instantiates a new LogicalDatabaseResponseList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogicalDatabaseResponseList() *LogicalDatabaseResponseList {
	this := LogicalDatabaseResponseList{}
	return &this
}

// NewLogicalDatabaseResponseListWithDefaults instantiates a new LogicalDatabaseResponseList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogicalDatabaseResponseListWithDefaults() *LogicalDatabaseResponseList {
	this := LogicalDatabaseResponseList{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *LogicalDatabaseResponseList) GetResults() []LogicalDatabase {
	if o == nil || o.Results == nil {
		var ret []LogicalDatabase
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogicalDatabaseResponseList) GetResultsOk() ([]LogicalDatabase, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *LogicalDatabaseResponseList) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []LogicalDatabase and assigns it to the Results field.
func (o *LogicalDatabaseResponseList) SetResults(v []LogicalDatabase) {
	o.Results = v
}

func (o LogicalDatabaseResponseList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullableLogicalDatabaseResponseList struct {
	value *LogicalDatabaseResponseList
	isSet bool
}

func (v NullableLogicalDatabaseResponseList) Get() *LogicalDatabaseResponseList {
	return v.value
}

func (v *NullableLogicalDatabaseResponseList) Set(val *LogicalDatabaseResponseList) {
	v.value = val
	v.isSet = true
}

func (v NullableLogicalDatabaseResponseList) IsSet() bool {
	return v.isSet
}

func (v *NullableLogicalDatabaseResponseList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogicalDatabaseResponseList(val *LogicalDatabaseResponseList) *NullableLogicalDatabaseResponseList {
	return &NullableLogicalDatabaseResponseList{value: val, isSet: true}
}

func (v NullableLogicalDatabaseResponseList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogicalDatabaseResponseList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
