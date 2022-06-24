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
	"time"
)

// DeploymentHistory struct for DeploymentHistory
type DeploymentHistory struct {
	Id        string     `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// name of the service
	Name   *string                      `json:"name,omitempty"`
	Commit *Commit                      `json:"commit,omitempty"`
	Status *DeploymentHistoryStatusEnum `json:"status,omitempty"`
}

// NewDeploymentHistory instantiates a new DeploymentHistory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploymentHistory(id string, createdAt time.Time) *DeploymentHistory {
	this := DeploymentHistory{}
	this.Id = id
	this.CreatedAt = createdAt
	return &this
}

// NewDeploymentHistoryWithDefaults instantiates a new DeploymentHistory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentHistoryWithDefaults() *DeploymentHistory {
	this := DeploymentHistory{}
	return &this
}

// GetId returns the Id field value
func (o *DeploymentHistory) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DeploymentHistory) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DeploymentHistory) SetId(v string) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *DeploymentHistory) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *DeploymentHistory) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *DeploymentHistory) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *DeploymentHistory) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentHistory) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *DeploymentHistory) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *DeploymentHistory) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DeploymentHistory) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentHistory) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DeploymentHistory) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DeploymentHistory) SetName(v string) {
	o.Name = &v
}

// GetCommit returns the Commit field value if set, zero value otherwise.
func (o *DeploymentHistory) GetCommit() Commit {
	if o == nil || o.Commit == nil {
		var ret Commit
		return ret
	}
	return *o.Commit
}

// GetCommitOk returns a tuple with the Commit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentHistory) GetCommitOk() (*Commit, bool) {
	if o == nil || o.Commit == nil {
		return nil, false
	}
	return o.Commit, true
}

// HasCommit returns a boolean if a field has been set.
func (o *DeploymentHistory) HasCommit() bool {
	if o != nil && o.Commit != nil {
		return true
	}

	return false
}

// SetCommit gets a reference to the given Commit and assigns it to the Commit field.
func (o *DeploymentHistory) SetCommit(v Commit) {
	o.Commit = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DeploymentHistory) GetStatus() DeploymentHistoryStatusEnum {
	if o == nil || o.Status == nil {
		var ret DeploymentHistoryStatusEnum
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentHistory) GetStatusOk() (*DeploymentHistoryStatusEnum, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DeploymentHistory) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given DeploymentHistoryStatusEnum and assigns it to the Status field.
func (o *DeploymentHistory) SetStatus(v DeploymentHistoryStatusEnum) {
	o.Status = &v
}

func (o DeploymentHistory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
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

type NullableDeploymentHistory struct {
	value *DeploymentHistory
	isSet bool
}

func (v NullableDeploymentHistory) Get() *DeploymentHistory {
	return v.value
}

func (v *NullableDeploymentHistory) Set(val *DeploymentHistory) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploymentHistory) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploymentHistory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploymentHistory(val *DeploymentHistory) *NullableDeploymentHistory {
	return &NullableDeploymentHistory{value: val, isSet: true}
}

func (v NullableDeploymentHistory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploymentHistory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
