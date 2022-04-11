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

// CommitResponseList struct for CommitResponseList
type CommitResponseList struct {
	Results []Commit `json:"results,omitempty"`
}

// NewCommitResponseList instantiates a new CommitResponseList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommitResponseList() *CommitResponseList {
	this := CommitResponseList{}
	return &this
}

// NewCommitResponseListWithDefaults instantiates a new CommitResponseList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommitResponseListWithDefaults() *CommitResponseList {
	this := CommitResponseList{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *CommitResponseList) GetResults() []Commit {
	if o == nil || o.Results == nil {
		var ret []Commit
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitResponseList) GetResultsOk() ([]Commit, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *CommitResponseList) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []Commit and assigns it to the Results field.
func (o *CommitResponseList) SetResults(v []Commit) {
	o.Results = v
}

func (o CommitResponseList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullableCommitResponseList struct {
	value *CommitResponseList
	isSet bool
}

func (v NullableCommitResponseList) Get() *CommitResponseList {
	return v.value
}

func (v *NullableCommitResponseList) Set(val *CommitResponseList) {
	v.value = val
	v.isSet = true
}

func (v NullableCommitResponseList) IsSet() bool {
	return v.isSet
}

func (v *NullableCommitResponseList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommitResponseList(val *CommitResponseList) *NullableCommitResponseList {
	return &NullableCommitResponseList{value: val, isSet: true}
}

func (v NullableCommitResponseList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommitResponseList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
