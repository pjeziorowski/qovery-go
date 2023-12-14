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

// CloudProviderEnum the model 'CloudProviderEnum'
type CloudProviderEnum string

// List of CloudProviderEnum
const (
	CLOUDPROVIDERENUM_AWS CloudProviderEnum = "AWS"
	CLOUDPROVIDERENUM_DO  CloudProviderEnum = "DO"
	CLOUDPROVIDERENUM_SCW CloudProviderEnum = "SCW"
	CLOUDPROVIDERENUM_GCP CloudProviderEnum = "GCP"
)

// All allowed values of CloudProviderEnum enum
var AllowedCloudProviderEnumEnumValues = []CloudProviderEnum{
	"AWS",
	"DO",
	"SCW",
	"GCP",
}

func (v *CloudProviderEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CloudProviderEnum(value)
	for _, existing := range AllowedCloudProviderEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CloudProviderEnum", value)
}

// NewCloudProviderEnumFromValue returns a pointer to a valid CloudProviderEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCloudProviderEnumFromValue(v string) (*CloudProviderEnum, error) {
	ev := CloudProviderEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CloudProviderEnum: valid values are %v", v, AllowedCloudProviderEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CloudProviderEnum) IsValid() bool {
	for _, existing := range AllowedCloudProviderEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CloudProviderEnum value
func (v CloudProviderEnum) Ptr() *CloudProviderEnum {
	return &v
}

type NullableCloudProviderEnum struct {
	value *CloudProviderEnum
	isSet bool
}

func (v NullableCloudProviderEnum) Get() *CloudProviderEnum {
	return v.value
}

func (v *NullableCloudProviderEnum) Set(val *CloudProviderEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudProviderEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudProviderEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudProviderEnum(val *CloudProviderEnum) *NullableCloudProviderEnum {
	return &NullableCloudProviderEnum{value: val, isSet: true}
}

func (v NullableCloudProviderEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudProviderEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
