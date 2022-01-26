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

// ApplicationStorageRequest struct for ApplicationStorageRequest
type ApplicationStorageRequest struct {
	Storage *[]ApplicationStorageRequestStorage `json:"storage,omitempty"`
}

// NewApplicationStorageRequest instantiates a new ApplicationStorageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationStorageRequest() *ApplicationStorageRequest {
	this := ApplicationStorageRequest{}
	return &this
}

// NewApplicationStorageRequestWithDefaults instantiates a new ApplicationStorageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationStorageRequestWithDefaults() *ApplicationStorageRequest {
	this := ApplicationStorageRequest{}
	return &this
}

// GetStorage returns the Storage field value if set, zero value otherwise.
func (o *ApplicationStorageRequest) GetStorage() []ApplicationStorageRequestStorage {
	if o == nil || o.Storage == nil {
		var ret []ApplicationStorageRequestStorage
		return ret
	}
	return *o.Storage
}

// GetStorageOk returns a tuple with the Storage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationStorageRequest) GetStorageOk() (*[]ApplicationStorageRequestStorage, bool) {
	if o == nil || o.Storage == nil {
		return nil, false
	}
	return o.Storage, true
}

// HasStorage returns a boolean if a field has been set.
func (o *ApplicationStorageRequest) HasStorage() bool {
	if o != nil && o.Storage != nil {
		return true
	}

	return false
}

// SetStorage gets a reference to the given []ApplicationStorageRequestStorage and assigns it to the Storage field.
func (o *ApplicationStorageRequest) SetStorage(v []ApplicationStorageRequestStorage) {
	o.Storage = &v
}

func (o ApplicationStorageRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Storage != nil {
		toSerialize["storage"] = o.Storage
	}
	return json.Marshal(toSerialize)
}

type NullableApplicationStorageRequest struct {
	value *ApplicationStorageRequest
	isSet bool
}

func (v NullableApplicationStorageRequest) Get() *ApplicationStorageRequest {
	return v.value
}

func (v *NullableApplicationStorageRequest) Set(val *ApplicationStorageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationStorageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationStorageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationStorageRequest(val *ApplicationStorageRequest) *NullableApplicationStorageRequest {
	return &NullableApplicationStorageRequest{value: val, isSet: true}
}

func (v NullableApplicationStorageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationStorageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


