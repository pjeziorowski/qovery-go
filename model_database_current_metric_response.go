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

// DatabaseCurrentMetricResponse struct for DatabaseCurrentMetricResponse
type DatabaseCurrentMetricResponse struct {
	Cpu     *EnvironmentDatabasesCurrentMetricResponseCpu     `json:"cpu,omitempty"`
	Memory  *EnvironmentDatabasesCurrentMetricResponseMemory  `json:"memory,omitempty"`
	Storage *EnvironmentDatabasesCurrentMetricResponseStorage `json:"storage,omitempty"`
}

// NewDatabaseCurrentMetricResponse instantiates a new DatabaseCurrentMetricResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatabaseCurrentMetricResponse() *DatabaseCurrentMetricResponse {
	this := DatabaseCurrentMetricResponse{}
	return &this
}

// NewDatabaseCurrentMetricResponseWithDefaults instantiates a new DatabaseCurrentMetricResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatabaseCurrentMetricResponseWithDefaults() *DatabaseCurrentMetricResponse {
	this := DatabaseCurrentMetricResponse{}
	return &this
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *DatabaseCurrentMetricResponse) GetCpu() EnvironmentDatabasesCurrentMetricResponseCpu {
	if o == nil || o.Cpu == nil {
		var ret EnvironmentDatabasesCurrentMetricResponseCpu
		return ret
	}
	return *o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseCurrentMetricResponse) GetCpuOk() (*EnvironmentDatabasesCurrentMetricResponseCpu, bool) {
	if o == nil || o.Cpu == nil {
		return nil, false
	}
	return o.Cpu, true
}

// HasCpu returns a boolean if a field has been set.
func (o *DatabaseCurrentMetricResponse) HasCpu() bool {
	if o != nil && o.Cpu != nil {
		return true
	}

	return false
}

// SetCpu gets a reference to the given EnvironmentDatabasesCurrentMetricResponseCpu and assigns it to the Cpu field.
func (o *DatabaseCurrentMetricResponse) SetCpu(v EnvironmentDatabasesCurrentMetricResponseCpu) {
	o.Cpu = &v
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *DatabaseCurrentMetricResponse) GetMemory() EnvironmentDatabasesCurrentMetricResponseMemory {
	if o == nil || o.Memory == nil {
		var ret EnvironmentDatabasesCurrentMetricResponseMemory
		return ret
	}
	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseCurrentMetricResponse) GetMemoryOk() (*EnvironmentDatabasesCurrentMetricResponseMemory, bool) {
	if o == nil || o.Memory == nil {
		return nil, false
	}
	return o.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *DatabaseCurrentMetricResponse) HasMemory() bool {
	if o != nil && o.Memory != nil {
		return true
	}

	return false
}

// SetMemory gets a reference to the given EnvironmentDatabasesCurrentMetricResponseMemory and assigns it to the Memory field.
func (o *DatabaseCurrentMetricResponse) SetMemory(v EnvironmentDatabasesCurrentMetricResponseMemory) {
	o.Memory = &v
}

// GetStorage returns the Storage field value if set, zero value otherwise.
func (o *DatabaseCurrentMetricResponse) GetStorage() EnvironmentDatabasesCurrentMetricResponseStorage {
	if o == nil || o.Storage == nil {
		var ret EnvironmentDatabasesCurrentMetricResponseStorage
		return ret
	}
	return *o.Storage
}

// GetStorageOk returns a tuple with the Storage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseCurrentMetricResponse) GetStorageOk() (*EnvironmentDatabasesCurrentMetricResponseStorage, bool) {
	if o == nil || o.Storage == nil {
		return nil, false
	}
	return o.Storage, true
}

// HasStorage returns a boolean if a field has been set.
func (o *DatabaseCurrentMetricResponse) HasStorage() bool {
	if o != nil && o.Storage != nil {
		return true
	}

	return false
}

// SetStorage gets a reference to the given EnvironmentDatabasesCurrentMetricResponseStorage and assigns it to the Storage field.
func (o *DatabaseCurrentMetricResponse) SetStorage(v EnvironmentDatabasesCurrentMetricResponseStorage) {
	o.Storage = &v
}

func (o DatabaseCurrentMetricResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Cpu != nil {
		toSerialize["cpu"] = o.Cpu
	}
	if o.Memory != nil {
		toSerialize["memory"] = o.Memory
	}
	if o.Storage != nil {
		toSerialize["storage"] = o.Storage
	}
	return json.Marshal(toSerialize)
}

type NullableDatabaseCurrentMetricResponse struct {
	value *DatabaseCurrentMetricResponse
	isSet bool
}

func (v NullableDatabaseCurrentMetricResponse) Get() *DatabaseCurrentMetricResponse {
	return v.value
}

func (v *NullableDatabaseCurrentMetricResponse) Set(val *DatabaseCurrentMetricResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDatabaseCurrentMetricResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDatabaseCurrentMetricResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatabaseCurrentMetricResponse(val *DatabaseCurrentMetricResponse) *NullableDatabaseCurrentMetricResponse {
	return &NullableDatabaseCurrentMetricResponse{value: val, isSet: true}
}

func (v NullableDatabaseCurrentMetricResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatabaseCurrentMetricResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
