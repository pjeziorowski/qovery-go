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

// EnvironmentStatusEventOriginEnum Origin of the organization event
type EnvironmentStatusEventOriginEnum string

// List of EnvironmentStatusEventOriginEnum
const (
	ENVIRONMENTSTATUSEVENTORIGINENUM_API                EnvironmentStatusEventOriginEnum = "API"
	ENVIRONMENTSTATUSEVENTORIGINENUM_CLI                EnvironmentStatusEventOriginEnum = "CLI"
	ENVIRONMENTSTATUSEVENTORIGINENUM_CONSOLE            EnvironmentStatusEventOriginEnum = "CONSOLE"
	ENVIRONMENTSTATUSEVENTORIGINENUM_GIT                EnvironmentStatusEventOriginEnum = "GIT"
	ENVIRONMENTSTATUSEVENTORIGINENUM_QOVERY_INTERNAL    EnvironmentStatusEventOriginEnum = "QOVERY_INTERNAL"
	ENVIRONMENTSTATUSEVENTORIGINENUM_TERRAFORM_PROVIDER EnvironmentStatusEventOriginEnum = "TERRAFORM_PROVIDER"
)

// All allowed values of EnvironmentStatusEventOriginEnum enum
var AllowedEnvironmentStatusEventOriginEnumEnumValues = []EnvironmentStatusEventOriginEnum{
	"API",
	"CLI",
	"CONSOLE",
	"GIT",
	"QOVERY_INTERNAL",
	"TERRAFORM_PROVIDER",
}

func (v *EnvironmentStatusEventOriginEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnvironmentStatusEventOriginEnum(value)
	for _, existing := range AllowedEnvironmentStatusEventOriginEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnvironmentStatusEventOriginEnum", value)
}

// NewEnvironmentStatusEventOriginEnumFromValue returns a pointer to a valid EnvironmentStatusEventOriginEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnvironmentStatusEventOriginEnumFromValue(v string) (*EnvironmentStatusEventOriginEnum, error) {
	ev := EnvironmentStatusEventOriginEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnvironmentStatusEventOriginEnum: valid values are %v", v, AllowedEnvironmentStatusEventOriginEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnvironmentStatusEventOriginEnum) IsValid() bool {
	for _, existing := range AllowedEnvironmentStatusEventOriginEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnvironmentStatusEventOriginEnum value
func (v EnvironmentStatusEventOriginEnum) Ptr() *EnvironmentStatusEventOriginEnum {
	return &v
}

type NullableEnvironmentStatusEventOriginEnum struct {
	value *EnvironmentStatusEventOriginEnum
	isSet bool
}

func (v NullableEnvironmentStatusEventOriginEnum) Get() *EnvironmentStatusEventOriginEnum {
	return v.value
}

func (v *NullableEnvironmentStatusEventOriginEnum) Set(val *EnvironmentStatusEventOriginEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentStatusEventOriginEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentStatusEventOriginEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentStatusEventOriginEnum(val *EnvironmentStatusEventOriginEnum) *NullableEnvironmentStatusEventOriginEnum {
	return &NullableEnvironmentStatusEventOriginEnum{value: val, isSet: true}
}

func (v NullableEnvironmentStatusEventOriginEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentStatusEventOriginEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
