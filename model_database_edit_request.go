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

// DatabaseEditRequest struct for DatabaseEditRequest
type DatabaseEditRequest struct {
	// name is case-insensitive
	Name *string `json:"name,omitempty"`
	Version *string `json:"version,omitempty"`
	Accessibility *string `json:"accessibility,omitempty"`
	// unit is millicores (m). 1000m = 1 cpu
	Cpu *int32 `json:"cpu,omitempty"`
	// unit is MB. 1024 MB = 1GB
	Memory *int32 `json:"memory,omitempty"`
	// unit is GB
	Storage *int32 `json:"storage,omitempty"`
}

// NewDatabaseEditRequest instantiates a new DatabaseEditRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatabaseEditRequest() *DatabaseEditRequest {
	this := DatabaseEditRequest{}
	var cpu int32 = 250
	this.Cpu = &cpu
	var memory int32 = 256
	this.Memory = &memory
	return &this
}

// NewDatabaseEditRequestWithDefaults instantiates a new DatabaseEditRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatabaseEditRequestWithDefaults() *DatabaseEditRequest {
	this := DatabaseEditRequest{}
	var cpu int32 = 250
	this.Cpu = &cpu
	var memory int32 = 256
	this.Memory = &memory
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DatabaseEditRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseEditRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DatabaseEditRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DatabaseEditRequest) SetName(v string) {
	o.Name = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *DatabaseEditRequest) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseEditRequest) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *DatabaseEditRequest) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *DatabaseEditRequest) SetVersion(v string) {
	o.Version = &v
}

// GetAccessibility returns the Accessibility field value if set, zero value otherwise.
func (o *DatabaseEditRequest) GetAccessibility() string {
	if o == nil || o.Accessibility == nil {
		var ret string
		return ret
	}
	return *o.Accessibility
}

// GetAccessibilityOk returns a tuple with the Accessibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseEditRequest) GetAccessibilityOk() (*string, bool) {
	if o == nil || o.Accessibility == nil {
		return nil, false
	}
	return o.Accessibility, true
}

// HasAccessibility returns a boolean if a field has been set.
func (o *DatabaseEditRequest) HasAccessibility() bool {
	if o != nil && o.Accessibility != nil {
		return true
	}

	return false
}

// SetAccessibility gets a reference to the given string and assigns it to the Accessibility field.
func (o *DatabaseEditRequest) SetAccessibility(v string) {
	o.Accessibility = &v
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *DatabaseEditRequest) GetCpu() int32 {
	if o == nil || o.Cpu == nil {
		var ret int32
		return ret
	}
	return *o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseEditRequest) GetCpuOk() (*int32, bool) {
	if o == nil || o.Cpu == nil {
		return nil, false
	}
	return o.Cpu, true
}

// HasCpu returns a boolean if a field has been set.
func (o *DatabaseEditRequest) HasCpu() bool {
	if o != nil && o.Cpu != nil {
		return true
	}

	return false
}

// SetCpu gets a reference to the given int32 and assigns it to the Cpu field.
func (o *DatabaseEditRequest) SetCpu(v int32) {
	o.Cpu = &v
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *DatabaseEditRequest) GetMemory() int32 {
	if o == nil || o.Memory == nil {
		var ret int32
		return ret
	}
	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseEditRequest) GetMemoryOk() (*int32, bool) {
	if o == nil || o.Memory == nil {
		return nil, false
	}
	return o.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *DatabaseEditRequest) HasMemory() bool {
	if o != nil && o.Memory != nil {
		return true
	}

	return false
}

// SetMemory gets a reference to the given int32 and assigns it to the Memory field.
func (o *DatabaseEditRequest) SetMemory(v int32) {
	o.Memory = &v
}

// GetStorage returns the Storage field value if set, zero value otherwise.
func (o *DatabaseEditRequest) GetStorage() int32 {
	if o == nil || o.Storage == nil {
		var ret int32
		return ret
	}
	return *o.Storage
}

// GetStorageOk returns a tuple with the Storage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseEditRequest) GetStorageOk() (*int32, bool) {
	if o == nil || o.Storage == nil {
		return nil, false
	}
	return o.Storage, true
}

// HasStorage returns a boolean if a field has been set.
func (o *DatabaseEditRequest) HasStorage() bool {
	if o != nil && o.Storage != nil {
		return true
	}

	return false
}

// SetStorage gets a reference to the given int32 and assigns it to the Storage field.
func (o *DatabaseEditRequest) SetStorage(v int32) {
	o.Storage = &v
}

func (o DatabaseEditRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.Accessibility != nil {
		toSerialize["accessibility"] = o.Accessibility
	}
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

type NullableDatabaseEditRequest struct {
	value *DatabaseEditRequest
	isSet bool
}

func (v NullableDatabaseEditRequest) Get() *DatabaseEditRequest {
	return v.value
}

func (v *NullableDatabaseEditRequest) Set(val *DatabaseEditRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDatabaseEditRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDatabaseEditRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatabaseEditRequest(val *DatabaseEditRequest) *NullableDatabaseEditRequest {
	return &NullableDatabaseEditRequest{value: val, isSet: true}
}

func (v NullableDatabaseEditRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatabaseEditRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


