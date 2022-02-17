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

// GitRepositoryResponseList struct for GitRepositoryResponseList
type GitRepositoryResponseList struct {
	Results *[]GitRepositoryResponse `json:"results,omitempty"`
}

// NewGitRepositoryResponseList instantiates a new GitRepositoryResponseList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGitRepositoryResponseList() *GitRepositoryResponseList {
	this := GitRepositoryResponseList{}
	return &this
}

// NewGitRepositoryResponseListWithDefaults instantiates a new GitRepositoryResponseList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGitRepositoryResponseListWithDefaults() *GitRepositoryResponseList {
	this := GitRepositoryResponseList{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *GitRepositoryResponseList) GetResults() []GitRepositoryResponse {
	if o == nil || o.Results == nil {
		var ret []GitRepositoryResponse
		return ret
	}
	return *o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitRepositoryResponseList) GetResultsOk() (*[]GitRepositoryResponse, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *GitRepositoryResponseList) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []GitRepositoryResponse and assigns it to the Results field.
func (o *GitRepositoryResponseList) SetResults(v []GitRepositoryResponse) {
	o.Results = &v
}

func (o GitRepositoryResponseList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullableGitRepositoryResponseList struct {
	value *GitRepositoryResponseList
	isSet bool
}

func (v NullableGitRepositoryResponseList) Get() *GitRepositoryResponseList {
	return v.value
}

func (v *NullableGitRepositoryResponseList) Set(val *GitRepositoryResponseList) {
	v.value = val
	v.isSet = true
}

func (v NullableGitRepositoryResponseList) IsSet() bool {
	return v.isSet
}

func (v *NullableGitRepositoryResponseList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGitRepositoryResponseList(val *GitRepositoryResponseList) *NullableGitRepositoryResponseList {
	return &NullableGitRepositoryResponseList{value: val, isSet: true}
}

func (v NullableGitRepositoryResponseList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGitRepositoryResponseList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


