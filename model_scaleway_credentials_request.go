/*
[BETA] Qovery API

- Qovery is the fastest way to deploy your full-stack apps on any Cloud provider. - ℹ️ The API is in Beta and still in progress. Some endpoints are not available yet.  # Authentication  <!-- ReDoc-Inject: <security-definitions> -->

API version: 1.0.0
Contact: support+api+documentation@qovery.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package qovery

import (
	"encoding/json"
)

// ScalewayCredentialsRequest struct for ScalewayCredentialsRequest
type ScalewayCredentialsRequest struct {
	Name string `json:"name"`
	ScalewayAccessKey *string `json:"scaleway_access_key,omitempty"`
	ScalewaySecretKey *string `json:"scaleway_secret_key,omitempty"`
	ScalewayProjectId *string `json:"scaleway_project_id,omitempty"`
}

// NewScalewayCredentialsRequest instantiates a new ScalewayCredentialsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScalewayCredentialsRequest(name string) *ScalewayCredentialsRequest {
	this := ScalewayCredentialsRequest{}
	this.Name = name
	return &this
}

// NewScalewayCredentialsRequestWithDefaults instantiates a new ScalewayCredentialsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScalewayCredentialsRequestWithDefaults() *ScalewayCredentialsRequest {
	this := ScalewayCredentialsRequest{}
	return &this
}

// GetName returns the Name field value
func (o *ScalewayCredentialsRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ScalewayCredentialsRequest) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ScalewayCredentialsRequest) SetName(v string) {
	o.Name = v
}

// GetScalewayAccessKey returns the ScalewayAccessKey field value if set, zero value otherwise.
func (o *ScalewayCredentialsRequest) GetScalewayAccessKey() string {
	if o == nil || o.ScalewayAccessKey == nil {
		var ret string
		return ret
	}
	return *o.ScalewayAccessKey
}

// GetScalewayAccessKeyOk returns a tuple with the ScalewayAccessKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScalewayCredentialsRequest) GetScalewayAccessKeyOk() (*string, bool) {
	if o == nil || o.ScalewayAccessKey == nil {
		return nil, false
	}
	return o.ScalewayAccessKey, true
}

// HasScalewayAccessKey returns a boolean if a field has been set.
func (o *ScalewayCredentialsRequest) HasScalewayAccessKey() bool {
	if o != nil && o.ScalewayAccessKey != nil {
		return true
	}

	return false
}

// SetScalewayAccessKey gets a reference to the given string and assigns it to the ScalewayAccessKey field.
func (o *ScalewayCredentialsRequest) SetScalewayAccessKey(v string) {
	o.ScalewayAccessKey = &v
}

// GetScalewaySecretKey returns the ScalewaySecretKey field value if set, zero value otherwise.
func (o *ScalewayCredentialsRequest) GetScalewaySecretKey() string {
	if o == nil || o.ScalewaySecretKey == nil {
		var ret string
		return ret
	}
	return *o.ScalewaySecretKey
}

// GetScalewaySecretKeyOk returns a tuple with the ScalewaySecretKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScalewayCredentialsRequest) GetScalewaySecretKeyOk() (*string, bool) {
	if o == nil || o.ScalewaySecretKey == nil {
		return nil, false
	}
	return o.ScalewaySecretKey, true
}

// HasScalewaySecretKey returns a boolean if a field has been set.
func (o *ScalewayCredentialsRequest) HasScalewaySecretKey() bool {
	if o != nil && o.ScalewaySecretKey != nil {
		return true
	}

	return false
}

// SetScalewaySecretKey gets a reference to the given string and assigns it to the ScalewaySecretKey field.
func (o *ScalewayCredentialsRequest) SetScalewaySecretKey(v string) {
	o.ScalewaySecretKey = &v
}

// GetScalewayProjectId returns the ScalewayProjectId field value if set, zero value otherwise.
func (o *ScalewayCredentialsRequest) GetScalewayProjectId() string {
	if o == nil || o.ScalewayProjectId == nil {
		var ret string
		return ret
	}
	return *o.ScalewayProjectId
}

// GetScalewayProjectIdOk returns a tuple with the ScalewayProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScalewayCredentialsRequest) GetScalewayProjectIdOk() (*string, bool) {
	if o == nil || o.ScalewayProjectId == nil {
		return nil, false
	}
	return o.ScalewayProjectId, true
}

// HasScalewayProjectId returns a boolean if a field has been set.
func (o *ScalewayCredentialsRequest) HasScalewayProjectId() bool {
	if o != nil && o.ScalewayProjectId != nil {
		return true
	}

	return false
}

// SetScalewayProjectId gets a reference to the given string and assigns it to the ScalewayProjectId field.
func (o *ScalewayCredentialsRequest) SetScalewayProjectId(v string) {
	o.ScalewayProjectId = &v
}

func (o ScalewayCredentialsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.ScalewayAccessKey != nil {
		toSerialize["scaleway_access_key"] = o.ScalewayAccessKey
	}
	if o.ScalewaySecretKey != nil {
		toSerialize["scaleway_secret_key"] = o.ScalewaySecretKey
	}
	if o.ScalewayProjectId != nil {
		toSerialize["scaleway_project_id"] = o.ScalewayProjectId
	}
	return json.Marshal(toSerialize)
}

type NullableScalewayCredentialsRequest struct {
	value *ScalewayCredentialsRequest
	isSet bool
}

func (v NullableScalewayCredentialsRequest) Get() *ScalewayCredentialsRequest {
	return v.value
}

func (v *NullableScalewayCredentialsRequest) Set(val *ScalewayCredentialsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableScalewayCredentialsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableScalewayCredentialsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScalewayCredentialsRequest(val *ScalewayCredentialsRequest) *NullableScalewayCredentialsRequest {
	return &NullableScalewayCredentialsRequest{value: val, isSet: true}
}

func (v NullableScalewayCredentialsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScalewayCredentialsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


