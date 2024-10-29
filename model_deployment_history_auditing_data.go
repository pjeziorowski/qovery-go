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
	"fmt"
	"time"
)

// checks if the DeploymentHistoryAuditingData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeploymentHistoryAuditingData{}

// DeploymentHistoryAuditingData struct for DeploymentHistoryAuditingData
type DeploymentHistoryAuditingData struct {
	CreatedAt            time.Time                `json:"created_at"`
	UpdatedAt            time.Time                `json:"updated_at"`
	TriggeredBy          string                   `json:"triggered_by"`
	Origin               *OrganizationEventOrigin `json:"origin,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DeploymentHistoryAuditingData DeploymentHistoryAuditingData

// NewDeploymentHistoryAuditingData instantiates a new DeploymentHistoryAuditingData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploymentHistoryAuditingData(createdAt time.Time, updatedAt time.Time, triggeredBy string) *DeploymentHistoryAuditingData {
	this := DeploymentHistoryAuditingData{}
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.TriggeredBy = triggeredBy
	return &this
}

// NewDeploymentHistoryAuditingDataWithDefaults instantiates a new DeploymentHistoryAuditingData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentHistoryAuditingDataWithDefaults() *DeploymentHistoryAuditingData {
	this := DeploymentHistoryAuditingData{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value
func (o *DeploymentHistoryAuditingData) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *DeploymentHistoryAuditingData) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *DeploymentHistoryAuditingData) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *DeploymentHistoryAuditingData) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *DeploymentHistoryAuditingData) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *DeploymentHistoryAuditingData) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetTriggeredBy returns the TriggeredBy field value
func (o *DeploymentHistoryAuditingData) GetTriggeredBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TriggeredBy
}

// GetTriggeredByOk returns a tuple with the TriggeredBy field value
// and a boolean to check if the value has been set.
func (o *DeploymentHistoryAuditingData) GetTriggeredByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TriggeredBy, true
}

// SetTriggeredBy sets field value
func (o *DeploymentHistoryAuditingData) SetTriggeredBy(v string) {
	o.TriggeredBy = v
}

// GetOrigin returns the Origin field value if set, zero value otherwise.
func (o *DeploymentHistoryAuditingData) GetOrigin() OrganizationEventOrigin {
	if o == nil || IsNil(o.Origin) {
		var ret OrganizationEventOrigin
		return ret
	}
	return *o.Origin
}

// GetOriginOk returns a tuple with the Origin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentHistoryAuditingData) GetOriginOk() (*OrganizationEventOrigin, bool) {
	if o == nil || IsNil(o.Origin) {
		return nil, false
	}
	return o.Origin, true
}

// HasOrigin returns a boolean if a field has been set.
func (o *DeploymentHistoryAuditingData) HasOrigin() bool {
	if o != nil && !IsNil(o.Origin) {
		return true
	}

	return false
}

// SetOrigin gets a reference to the given OrganizationEventOrigin and assigns it to the Origin field.
func (o *DeploymentHistoryAuditingData) SetOrigin(v OrganizationEventOrigin) {
	o.Origin = &v
}

func (o DeploymentHistoryAuditingData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeploymentHistoryAuditingData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt
	toSerialize["triggered_by"] = o.TriggeredBy
	if !IsNil(o.Origin) {
		toSerialize["origin"] = o.Origin
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DeploymentHistoryAuditingData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"created_at",
		"updated_at",
		"triggered_by",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varDeploymentHistoryAuditingData := _DeploymentHistoryAuditingData{}

	err = json.Unmarshal(data, &varDeploymentHistoryAuditingData)

	if err != nil {
		return err
	}

	*o = DeploymentHistoryAuditingData(varDeploymentHistoryAuditingData)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "triggered_by")
		delete(additionalProperties, "origin")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeploymentHistoryAuditingData struct {
	value *DeploymentHistoryAuditingData
	isSet bool
}

func (v NullableDeploymentHistoryAuditingData) Get() *DeploymentHistoryAuditingData {
	return v.value
}

func (v *NullableDeploymentHistoryAuditingData) Set(val *DeploymentHistoryAuditingData) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploymentHistoryAuditingData) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploymentHistoryAuditingData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploymentHistoryAuditingData(val *DeploymentHistoryAuditingData) *NullableDeploymentHistoryAuditingData {
	return &NullableDeploymentHistoryAuditingData{value: val, isSet: true}
}

func (v NullableDeploymentHistoryAuditingData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploymentHistoryAuditingData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
