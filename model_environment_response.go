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
	"time"
)

// EnvironmentResponse struct for EnvironmentResponse
type EnvironmentResponse struct {
	// name is case insensitive
	Name    string           `json:"name"`
	Project *ReferenceObject `json:"project,omitempty"`
	// uuid of the user that made the last update
	LastUpdatedBy *string                `json:"last_updated_by,omitempty"`
	CloudProvider map[string]interface{} `json:"cloud_provider"`
	Mode          string                 `json:"mode"`
	Id            string                 `json:"id"`
	CreatedAt     time.Time              `json:"created_at"`
	UpdatedAt     *time.Time             `json:"updated_at,omitempty"`
}

// NewEnvironmentResponse instantiates a new EnvironmentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentResponse(name string, cloudProvider map[string]interface{}, mode string, id string, createdAt time.Time) *EnvironmentResponse {
	this := EnvironmentResponse{}
	this.Id = id
	this.CreatedAt = createdAt
	return &this
}

// NewEnvironmentResponseWithDefaults instantiates a new EnvironmentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentResponseWithDefaults() *EnvironmentResponse {
	this := EnvironmentResponse{}
	return &this
}

// GetName returns the Name field value
func (o *EnvironmentResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EnvironmentResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EnvironmentResponse) SetName(v string) {
	o.Name = v
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *EnvironmentResponse) GetProject() ReferenceObject {
	if o == nil || o.Project == nil {
		var ret ReferenceObject
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentResponse) GetProjectOk() (*ReferenceObject, bool) {
	if o == nil || o.Project == nil {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *EnvironmentResponse) HasProject() bool {
	if o != nil && o.Project != nil {
		return true
	}

	return false
}

// SetProject gets a reference to the given ReferenceObject and assigns it to the Project field.
func (o *EnvironmentResponse) SetProject(v ReferenceObject) {
	o.Project = &v
}

// GetLastUpdatedBy returns the LastUpdatedBy field value if set, zero value otherwise.
func (o *EnvironmentResponse) GetLastUpdatedBy() string {
	if o == nil || o.LastUpdatedBy == nil {
		var ret string
		return ret
	}
	return *o.LastUpdatedBy
}

// GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentResponse) GetLastUpdatedByOk() (*string, bool) {
	if o == nil || o.LastUpdatedBy == nil {
		return nil, false
	}
	return o.LastUpdatedBy, true
}

// HasLastUpdatedBy returns a boolean if a field has been set.
func (o *EnvironmentResponse) HasLastUpdatedBy() bool {
	if o != nil && o.LastUpdatedBy != nil {
		return true
	}

	return false
}

// SetLastUpdatedBy gets a reference to the given string and assigns it to the LastUpdatedBy field.
func (o *EnvironmentResponse) SetLastUpdatedBy(v string) {
	o.LastUpdatedBy = &v
}

// GetCloudProvider returns the CloudProvider field value
func (o *EnvironmentResponse) GetCloudProvider() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value
// and a boolean to check if the value has been set.
func (o *EnvironmentResponse) GetCloudProviderOk() (*map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudProvider, true
}

// SetCloudProvider sets field value
func (o *EnvironmentResponse) SetCloudProvider(v map[string]interface{}) {
	o.CloudProvider = v
}

// GetMode returns the Mode field value
func (o *EnvironmentResponse) GetMode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *EnvironmentResponse) GetModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value
func (o *EnvironmentResponse) SetMode(v string) {
	o.Mode = v
}

// GetId returns the Id field value
func (o *EnvironmentResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EnvironmentResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EnvironmentResponse) SetId(v string) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *EnvironmentResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *EnvironmentResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *EnvironmentResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *EnvironmentResponse) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *EnvironmentResponse) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *EnvironmentResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o EnvironmentResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Project != nil {
		toSerialize["project"] = o.Project
	}
	if o.LastUpdatedBy != nil {
		toSerialize["last_updated_by"] = o.LastUpdatedBy
	}
	if true {
		toSerialize["cloud_provider"] = o.CloudProvider
	}
	if true {
		toSerialize["mode"] = o.Mode
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableEnvironmentResponse struct {
	value *EnvironmentResponse
	isSet bool
}

func (v NullableEnvironmentResponse) Get() *EnvironmentResponse {
	return v.value
}

func (v *NullableEnvironmentResponse) Set(val *EnvironmentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentResponse(val *EnvironmentResponse) *NullableEnvironmentResponse {
	return &NullableEnvironmentResponse{value: val, isSet: true}
}

func (v NullableEnvironmentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
