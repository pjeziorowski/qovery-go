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

// EnvironmentVariable struct for EnvironmentVariable
type EnvironmentVariable struct {
	Id        string     `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// key is case sensitive
	Key string `json:"key"`
	// value of the env variable.
	Value              string                                      `json:"value"`
	OverriddenVariable *EnvironmentVariableAllOfOverriddenVariable `json:"overridden_variable,omitempty"`
	AliasedVariable    *EnvironmentVariableAllOfAliasedVariable    `json:"aliased_variable,omitempty"`
	Scope              EnvironmentVariableScopeEnum                `json:"scope"`
	// present only for `BUILT_IN` variable
	ServiceId *string `json:"service_id,omitempty"`
	// present only for `BUILT_IN` variable
	ServiceName *string      `json:"service_name,omitempty"`
	ServiceType *ServiceType `json:"service_type,omitempty"`
}

// NewEnvironmentVariable instantiates a new EnvironmentVariable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentVariable(id string, createdAt time.Time, key string, value string, scope EnvironmentVariableScopeEnum) *EnvironmentVariable {
	this := EnvironmentVariable{}
	this.Id = id
	this.CreatedAt = createdAt
	this.Key = key
	this.Value = value
	this.Scope = scope
	return &this
}

// NewEnvironmentVariableWithDefaults instantiates a new EnvironmentVariable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentVariableWithDefaults() *EnvironmentVariable {
	this := EnvironmentVariable{}
	return &this
}

// GetId returns the Id field value
func (o *EnvironmentVariable) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EnvironmentVariable) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EnvironmentVariable) SetId(v string) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *EnvironmentVariable) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *EnvironmentVariable) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *EnvironmentVariable) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *EnvironmentVariable) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentVariable) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *EnvironmentVariable) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *EnvironmentVariable) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetKey returns the Key field value
func (o *EnvironmentVariable) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *EnvironmentVariable) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *EnvironmentVariable) SetKey(v string) {
	o.Key = v
}

// GetValue returns the Value field value
func (o *EnvironmentVariable) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *EnvironmentVariable) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *EnvironmentVariable) SetValue(v string) {
	o.Value = v
}

// GetOverriddenVariable returns the OverriddenVariable field value if set, zero value otherwise.
func (o *EnvironmentVariable) GetOverriddenVariable() EnvironmentVariableAllOfOverriddenVariable {
	if o == nil || o.OverriddenVariable == nil {
		var ret EnvironmentVariableAllOfOverriddenVariable
		return ret
	}
	return *o.OverriddenVariable
}

// GetOverriddenVariableOk returns a tuple with the OverriddenVariable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentVariable) GetOverriddenVariableOk() (*EnvironmentVariableAllOfOverriddenVariable, bool) {
	if o == nil || o.OverriddenVariable == nil {
		return nil, false
	}
	return o.OverriddenVariable, true
}

// HasOverriddenVariable returns a boolean if a field has been set.
func (o *EnvironmentVariable) HasOverriddenVariable() bool {
	if o != nil && o.OverriddenVariable != nil {
		return true
	}

	return false
}

// SetOverriddenVariable gets a reference to the given EnvironmentVariableAllOfOverriddenVariable and assigns it to the OverriddenVariable field.
func (o *EnvironmentVariable) SetOverriddenVariable(v EnvironmentVariableAllOfOverriddenVariable) {
	o.OverriddenVariable = &v
}

// GetAliasedVariable returns the AliasedVariable field value if set, zero value otherwise.
func (o *EnvironmentVariable) GetAliasedVariable() EnvironmentVariableAllOfAliasedVariable {
	if o == nil || o.AliasedVariable == nil {
		var ret EnvironmentVariableAllOfAliasedVariable
		return ret
	}
	return *o.AliasedVariable
}

// GetAliasedVariableOk returns a tuple with the AliasedVariable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentVariable) GetAliasedVariableOk() (*EnvironmentVariableAllOfAliasedVariable, bool) {
	if o == nil || o.AliasedVariable == nil {
		return nil, false
	}
	return o.AliasedVariable, true
}

// HasAliasedVariable returns a boolean if a field has been set.
func (o *EnvironmentVariable) HasAliasedVariable() bool {
	if o != nil && o.AliasedVariable != nil {
		return true
	}

	return false
}

// SetAliasedVariable gets a reference to the given EnvironmentVariableAllOfAliasedVariable and assigns it to the AliasedVariable field.
func (o *EnvironmentVariable) SetAliasedVariable(v EnvironmentVariableAllOfAliasedVariable) {
	o.AliasedVariable = &v
}

// GetScope returns the Scope field value
func (o *EnvironmentVariable) GetScope() EnvironmentVariableScopeEnum {
	if o == nil {
		var ret EnvironmentVariableScopeEnum
		return ret
	}

	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *EnvironmentVariable) GetScopeOk() (*EnvironmentVariableScopeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scope, true
}

// SetScope sets field value
func (o *EnvironmentVariable) SetScope(v EnvironmentVariableScopeEnum) {
	o.Scope = v
}

// GetServiceId returns the ServiceId field value if set, zero value otherwise.
func (o *EnvironmentVariable) GetServiceId() string {
	if o == nil || o.ServiceId == nil {
		var ret string
		return ret
	}
	return *o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentVariable) GetServiceIdOk() (*string, bool) {
	if o == nil || o.ServiceId == nil {
		return nil, false
	}
	return o.ServiceId, true
}

// HasServiceId returns a boolean if a field has been set.
func (o *EnvironmentVariable) HasServiceId() bool {
	if o != nil && o.ServiceId != nil {
		return true
	}

	return false
}

// SetServiceId gets a reference to the given string and assigns it to the ServiceId field.
func (o *EnvironmentVariable) SetServiceId(v string) {
	o.ServiceId = &v
}

// GetServiceName returns the ServiceName field value if set, zero value otherwise.
func (o *EnvironmentVariable) GetServiceName() string {
	if o == nil || o.ServiceName == nil {
		var ret string
		return ret
	}
	return *o.ServiceName
}

// GetServiceNameOk returns a tuple with the ServiceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentVariable) GetServiceNameOk() (*string, bool) {
	if o == nil || o.ServiceName == nil {
		return nil, false
	}
	return o.ServiceName, true
}

// HasServiceName returns a boolean if a field has been set.
func (o *EnvironmentVariable) HasServiceName() bool {
	if o != nil && o.ServiceName != nil {
		return true
	}

	return false
}

// SetServiceName gets a reference to the given string and assigns it to the ServiceName field.
func (o *EnvironmentVariable) SetServiceName(v string) {
	o.ServiceName = &v
}

// GetServiceType returns the ServiceType field value if set, zero value otherwise.
func (o *EnvironmentVariable) GetServiceType() ServiceType {
	if o == nil || o.ServiceType == nil {
		var ret ServiceType
		return ret
	}
	return *o.ServiceType
}

// GetServiceTypeOk returns a tuple with the ServiceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentVariable) GetServiceTypeOk() (*ServiceType, bool) {
	if o == nil || o.ServiceType == nil {
		return nil, false
	}
	return o.ServiceType, true
}

// HasServiceType returns a boolean if a field has been set.
func (o *EnvironmentVariable) HasServiceType() bool {
	if o != nil && o.ServiceType != nil {
		return true
	}

	return false
}

// SetServiceType gets a reference to the given ServiceType and assigns it to the ServiceType field.
func (o *EnvironmentVariable) SetServiceType(v ServiceType) {
	o.ServiceType = &v
}

func (o EnvironmentVariable) MarshalJSON() ([]byte, error) {
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
	if true {
		toSerialize["key"] = o.Key
	}
	if true {
		toSerialize["value"] = o.Value
	}
	if o.OverriddenVariable != nil {
		toSerialize["overridden_variable"] = o.OverriddenVariable
	}
	if o.AliasedVariable != nil {
		toSerialize["aliased_variable"] = o.AliasedVariable
	}
	if true {
		toSerialize["scope"] = o.Scope
	}
	if o.ServiceId != nil {
		toSerialize["service_id"] = o.ServiceId
	}
	if o.ServiceName != nil {
		toSerialize["service_name"] = o.ServiceName
	}
	if o.ServiceType != nil {
		toSerialize["service_type"] = o.ServiceType
	}
	return json.Marshal(toSerialize)
}

type NullableEnvironmentVariable struct {
	value *EnvironmentVariable
	isSet bool
}

func (v NullableEnvironmentVariable) Get() *EnvironmentVariable {
	return v.value
}

func (v *NullableEnvironmentVariable) Set(val *EnvironmentVariable) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentVariable) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentVariable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentVariable(val *EnvironmentVariable) *NullableEnvironmentVariable {
	return &NullableEnvironmentVariable{value: val, isSet: true}
}

func (v NullableEnvironmentVariable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentVariable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
