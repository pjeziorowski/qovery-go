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

// DeploymentHistoryApplicationResponse struct for DeploymentHistoryApplicationResponse
type DeploymentHistoryApplicationResponse struct {
	Id        string          `json:"id"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt *time.Time      `json:"updated_at,omitempty"`
	Name      *string         `json:"name,omitempty"`
	Commit    *CommitResponse `json:"commit,omitempty"`
	Status    *string         `json:"status,omitempty"`
}

// NewDeploymentHistoryApplicationResponse instantiates a new DeploymentHistoryApplicationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploymentHistoryApplicationResponse(id string, createdAt time.Time) *DeploymentHistoryApplicationResponse {
	this := DeploymentHistoryApplicationResponse{}
	this.Id = id
	this.CreatedAt = createdAt
	return &this
}

// NewDeploymentHistoryApplicationResponseWithDefaults instantiates a new DeploymentHistoryApplicationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentHistoryApplicationResponseWithDefaults() *DeploymentHistoryApplicationResponse {
	this := DeploymentHistoryApplicationResponse{}
	return &this
}

// GetId returns the Id field value
func (o *DeploymentHistoryApplicationResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DeploymentHistoryApplicationResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DeploymentHistoryApplicationResponse) SetId(v string) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *DeploymentHistoryApplicationResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *DeploymentHistoryApplicationResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *DeploymentHistoryApplicationResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *DeploymentHistoryApplicationResponse) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentHistoryApplicationResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *DeploymentHistoryApplicationResponse) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *DeploymentHistoryApplicationResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DeploymentHistoryApplicationResponse) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentHistoryApplicationResponse) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DeploymentHistoryApplicationResponse) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DeploymentHistoryApplicationResponse) SetName(v string) {
	o.Name = &v
}

// GetCommit returns the Commit field value if set, zero value otherwise.
func (o *DeploymentHistoryApplicationResponse) GetCommit() CommitResponse {
	if o == nil || o.Commit == nil {
		var ret CommitResponse
		return ret
	}
	return *o.Commit
}

// GetCommitOk returns a tuple with the Commit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentHistoryApplicationResponse) GetCommitOk() (*CommitResponse, bool) {
	if o == nil || o.Commit == nil {
		return nil, false
	}
	return o.Commit, true
}

// HasCommit returns a boolean if a field has been set.
func (o *DeploymentHistoryApplicationResponse) HasCommit() bool {
	if o != nil && o.Commit != nil {
		return true
	}

	return false
}

// SetCommit gets a reference to the given CommitResponse and assigns it to the Commit field.
func (o *DeploymentHistoryApplicationResponse) SetCommit(v CommitResponse) {
	o.Commit = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DeploymentHistoryApplicationResponse) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentHistoryApplicationResponse) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DeploymentHistoryApplicationResponse) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *DeploymentHistoryApplicationResponse) SetStatus(v string) {
	o.Status = &v
}

func (o DeploymentHistoryApplicationResponse) MarshalJSON() ([]byte, error) {
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

type NullableDeploymentHistoryApplicationResponse struct {
	value *DeploymentHistoryApplicationResponse
	isSet bool
}

func (v NullableDeploymentHistoryApplicationResponse) Get() *DeploymentHistoryApplicationResponse {
	return v.value
}

func (v *NullableDeploymentHistoryApplicationResponse) Set(val *DeploymentHistoryApplicationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploymentHistoryApplicationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploymentHistoryApplicationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploymentHistoryApplicationResponse(val *DeploymentHistoryApplicationResponse) *NullableDeploymentHistoryApplicationResponse {
	return &NullableDeploymentHistoryApplicationResponse{value: val, isSet: true}
}

func (v NullableDeploymentHistoryApplicationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploymentHistoryApplicationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
