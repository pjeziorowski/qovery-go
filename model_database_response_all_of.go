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

// DatabaseResponseAllOf struct for DatabaseResponseAllOf
type DatabaseResponseAllOf struct {
	Environment *ReferenceObject `json:"environment,omitempty"`
	Host        *string          `json:"host,omitempty"`
	Port        *int32           `json:"port,omitempty"`
	// Maximum cpu that can be allocated to the database based on organization cluster configuration. unit is millicores (m). 1000m = 1 cpu
	MaximumCpu *int32 `json:"maximum_cpu,omitempty"`
	// Maximum memory that can be allocated to the database based on organization cluster configuration. unit is MB. 1024 MB = 1GB
	MaximumMemory *int32 `json:"maximum_memory,omitempty"`
	// indicates if the database disk is encrypted or not
	DiskEncrypted *bool `json:"disk_encrypted,omitempty"`
}

// NewDatabaseResponseAllOf instantiates a new DatabaseResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatabaseResponseAllOf() *DatabaseResponseAllOf {
	this := DatabaseResponseAllOf{}
	var maximumCpu int32 = 250
	this.MaximumCpu = &maximumCpu
	var maximumMemory int32 = 256
	this.MaximumMemory = &maximumMemory
	return &this
}

// NewDatabaseResponseAllOfWithDefaults instantiates a new DatabaseResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatabaseResponseAllOfWithDefaults() *DatabaseResponseAllOf {
	this := DatabaseResponseAllOf{}
	var maximumCpu int32 = 250
	this.MaximumCpu = &maximumCpu
	var maximumMemory int32 = 256
	this.MaximumMemory = &maximumMemory
	return &this
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *DatabaseResponseAllOf) GetEnvironment() ReferenceObject {
	if o == nil || o.Environment == nil {
		var ret ReferenceObject
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseResponseAllOf) GetEnvironmentOk() (*ReferenceObject, bool) {
	if o == nil || o.Environment == nil {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *DatabaseResponseAllOf) HasEnvironment() bool {
	if o != nil && o.Environment != nil {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given ReferenceObject and assigns it to the Environment field.
func (o *DatabaseResponseAllOf) SetEnvironment(v ReferenceObject) {
	o.Environment = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *DatabaseResponseAllOf) GetHost() string {
	if o == nil || o.Host == nil {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseResponseAllOf) GetHostOk() (*string, bool) {
	if o == nil || o.Host == nil {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *DatabaseResponseAllOf) HasHost() bool {
	if o != nil && o.Host != nil {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *DatabaseResponseAllOf) SetHost(v string) {
	o.Host = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *DatabaseResponseAllOf) GetPort() int32 {
	if o == nil || o.Port == nil {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseResponseAllOf) GetPortOk() (*int32, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *DatabaseResponseAllOf) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *DatabaseResponseAllOf) SetPort(v int32) {
	o.Port = &v
}

// GetMaximumCpu returns the MaximumCpu field value if set, zero value otherwise.
func (o *DatabaseResponseAllOf) GetMaximumCpu() int32 {
	if o == nil || o.MaximumCpu == nil {
		var ret int32
		return ret
	}
	return *o.MaximumCpu
}

// GetMaximumCpuOk returns a tuple with the MaximumCpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseResponseAllOf) GetMaximumCpuOk() (*int32, bool) {
	if o == nil || o.MaximumCpu == nil {
		return nil, false
	}
	return o.MaximumCpu, true
}

// HasMaximumCpu returns a boolean if a field has been set.
func (o *DatabaseResponseAllOf) HasMaximumCpu() bool {
	if o != nil && o.MaximumCpu != nil {
		return true
	}

	return false
}

// SetMaximumCpu gets a reference to the given int32 and assigns it to the MaximumCpu field.
func (o *DatabaseResponseAllOf) SetMaximumCpu(v int32) {
	o.MaximumCpu = &v
}

// GetMaximumMemory returns the MaximumMemory field value if set, zero value otherwise.
func (o *DatabaseResponseAllOf) GetMaximumMemory() int32 {
	if o == nil || o.MaximumMemory == nil {
		var ret int32
		return ret
	}
	return *o.MaximumMemory
}

// GetMaximumMemoryOk returns a tuple with the MaximumMemory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseResponseAllOf) GetMaximumMemoryOk() (*int32, bool) {
	if o == nil || o.MaximumMemory == nil {
		return nil, false
	}
	return o.MaximumMemory, true
}

// HasMaximumMemory returns a boolean if a field has been set.
func (o *DatabaseResponseAllOf) HasMaximumMemory() bool {
	if o != nil && o.MaximumMemory != nil {
		return true
	}

	return false
}

// SetMaximumMemory gets a reference to the given int32 and assigns it to the MaximumMemory field.
func (o *DatabaseResponseAllOf) SetMaximumMemory(v int32) {
	o.MaximumMemory = &v
}

// GetDiskEncrypted returns the DiskEncrypted field value if set, zero value otherwise.
func (o *DatabaseResponseAllOf) GetDiskEncrypted() bool {
	if o == nil || o.DiskEncrypted == nil {
		var ret bool
		return ret
	}
	return *o.DiskEncrypted
}

// GetDiskEncryptedOk returns a tuple with the DiskEncrypted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseResponseAllOf) GetDiskEncryptedOk() (*bool, bool) {
	if o == nil || o.DiskEncrypted == nil {
		return nil, false
	}
	return o.DiskEncrypted, true
}

// HasDiskEncrypted returns a boolean if a field has been set.
func (o *DatabaseResponseAllOf) HasDiskEncrypted() bool {
	if o != nil && o.DiskEncrypted != nil {
		return true
	}

	return false
}

// SetDiskEncrypted gets a reference to the given bool and assigns it to the DiskEncrypted field.
func (o *DatabaseResponseAllOf) SetDiskEncrypted(v bool) {
	o.DiskEncrypted = &v
}

func (o DatabaseResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Environment != nil {
		toSerialize["environment"] = o.Environment
	}
	if o.Host != nil {
		toSerialize["host"] = o.Host
	}
	if o.Port != nil {
		toSerialize["port"] = o.Port
	}
	if o.MaximumCpu != nil {
		toSerialize["maximum_cpu"] = o.MaximumCpu
	}
	if o.MaximumMemory != nil {
		toSerialize["maximum_memory"] = o.MaximumMemory
	}
	if o.DiskEncrypted != nil {
		toSerialize["disk_encrypted"] = o.DiskEncrypted
	}
	return json.Marshal(toSerialize)
}

type NullableDatabaseResponseAllOf struct {
	value *DatabaseResponseAllOf
	isSet bool
}

func (v NullableDatabaseResponseAllOf) Get() *DatabaseResponseAllOf {
	return v.value
}

func (v *NullableDatabaseResponseAllOf) Set(val *DatabaseResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableDatabaseResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableDatabaseResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatabaseResponseAllOf(val *DatabaseResponseAllOf) *NullableDatabaseResponseAllOf {
	return &NullableDatabaseResponseAllOf{value: val, isSet: true}
}

func (v NullableDatabaseResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatabaseResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
