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
)

// checks if the StatusDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StatusDetails{}

// StatusDetails struct for StatusDetails
type StatusDetails struct {
	Action               ServiceActionEnum       `json:"action"`
	Status               ServiceActionStatusEnum `json:"status"`
	AdditionalProperties map[string]interface{}
}

type _StatusDetails StatusDetails

// NewStatusDetails instantiates a new StatusDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatusDetails(action ServiceActionEnum, status ServiceActionStatusEnum) *StatusDetails {
	this := StatusDetails{}
	this.Action = action
	this.Status = status
	return &this
}

// NewStatusDetailsWithDefaults instantiates a new StatusDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatusDetailsWithDefaults() *StatusDetails {
	this := StatusDetails{}
	return &this
}

// GetAction returns the Action field value
func (o *StatusDetails) GetAction() ServiceActionEnum {
	if o == nil {
		var ret ServiceActionEnum
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *StatusDetails) GetActionOk() (*ServiceActionEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *StatusDetails) SetAction(v ServiceActionEnum) {
	o.Action = v
}

// GetStatus returns the Status field value
func (o *StatusDetails) GetStatus() ServiceActionStatusEnum {
	if o == nil {
		var ret ServiceActionStatusEnum
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *StatusDetails) GetStatusOk() (*ServiceActionStatusEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *StatusDetails) SetStatus(v ServiceActionStatusEnum) {
	o.Status = v
}

func (o StatusDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StatusDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["action"] = o.Action
	toSerialize["status"] = o.Status

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StatusDetails) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"action",
		"status",
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

	varStatusDetails := _StatusDetails{}

	err = json.Unmarshal(data, &varStatusDetails)

	if err != nil {
		return err
	}

	*o = StatusDetails(varStatusDetails)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "action")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStatusDetails struct {
	value *StatusDetails
	isSet bool
}

func (v NullableStatusDetails) Get() *StatusDetails {
	return v.value
}

func (v *NullableStatusDetails) Set(val *StatusDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableStatusDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableStatusDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatusDetails(val *StatusDetails) *NullableStatusDetails {
	return &NullableStatusDetails{value: val, isSet: true}
}

func (v NullableStatusDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatusDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
