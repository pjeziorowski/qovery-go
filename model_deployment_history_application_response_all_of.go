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

// DeploymentHistoryApplicationResponseAllOf struct for DeploymentHistoryApplicationResponseAllOf
type DeploymentHistoryApplicationResponseAllOf struct {
	Name   *string         `json:"name,omitempty"`
	Commit *CommitResponse `json:"commit,omitempty"`
	Status *string         `json:"status,omitempty"`
}

// NewDeploymentHistoryApplicationResponseAllOf instantiates a new DeploymentHistoryApplicationResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploymentHistoryApplicationResponseAllOf() *DeploymentHistoryApplicationResponseAllOf {
	this := DeploymentHistoryApplicationResponseAllOf{}
	return &this
}

// NewDeploymentHistoryApplicationResponseAllOfWithDefaults instantiates a new DeploymentHistoryApplicationResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentHistoryApplicationResponseAllOfWithDefaults() *DeploymentHistoryApplicationResponseAllOf {
	this := DeploymentHistoryApplicationResponseAllOf{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DeploymentHistoryApplicationResponseAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentHistoryApplicationResponseAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DeploymentHistoryApplicationResponseAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DeploymentHistoryApplicationResponseAllOf) SetName(v string) {
	o.Name = &v
}

// GetCommit returns the Commit field value if set, zero value otherwise.
func (o *DeploymentHistoryApplicationResponseAllOf) GetCommit() CommitResponse {
	if o == nil || o.Commit == nil {
		var ret CommitResponse
		return ret
	}
	return *o.Commit
}

// GetCommitOk returns a tuple with the Commit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentHistoryApplicationResponseAllOf) GetCommitOk() (*CommitResponse, bool) {
	if o == nil || o.Commit == nil {
		return nil, false
	}
	return o.Commit, true
}

// HasCommit returns a boolean if a field has been set.
func (o *DeploymentHistoryApplicationResponseAllOf) HasCommit() bool {
	if o != nil && o.Commit != nil {
		return true
	}

	return false
}

// SetCommit gets a reference to the given CommitResponse and assigns it to the Commit field.
func (o *DeploymentHistoryApplicationResponseAllOf) SetCommit(v CommitResponse) {
	o.Commit = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DeploymentHistoryApplicationResponseAllOf) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentHistoryApplicationResponseAllOf) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DeploymentHistoryApplicationResponseAllOf) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *DeploymentHistoryApplicationResponseAllOf) SetStatus(v string) {
	o.Status = &v
}

func (o DeploymentHistoryApplicationResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Commit != nil {
		toSerialize["commit"] = o.Commit
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableDeploymentHistoryApplicationResponseAllOf struct {
	value *DeploymentHistoryApplicationResponseAllOf
	isSet bool
}

func (v NullableDeploymentHistoryApplicationResponseAllOf) Get() *DeploymentHistoryApplicationResponseAllOf {
	return v.value
}

func (v *NullableDeploymentHistoryApplicationResponseAllOf) Set(val *DeploymentHistoryApplicationResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploymentHistoryApplicationResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploymentHistoryApplicationResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploymentHistoryApplicationResponseAllOf(val *DeploymentHistoryApplicationResponseAllOf) *NullableDeploymentHistoryApplicationResponseAllOf {
	return &NullableDeploymentHistoryApplicationResponseAllOf{value: val, isSet: true}
}

func (v NullableDeploymentHistoryApplicationResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploymentHistoryApplicationResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
