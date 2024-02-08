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

// ClusterRequestFeaturesInnerValue - struct for ClusterRequestFeaturesInnerValue
type ClusterRequestFeaturesInnerValue struct {
	ClusterFeatureAwsExistingVpc *ClusterFeatureAwsExistingVpc
	Bool                         *bool
	String                       *string
}

// ClusterFeatureAwsExistingVpcAsClusterRequestFeaturesInnerValue is a convenience function that returns ClusterFeatureAwsExistingVpc wrapped in ClusterRequestFeaturesInnerValue
func ClusterFeatureAwsExistingVpcAsClusterRequestFeaturesInnerValue(v *ClusterFeatureAwsExistingVpc) ClusterRequestFeaturesInnerValue {
	return ClusterRequestFeaturesInnerValue{
		ClusterFeatureAwsExistingVpc: v,
	}
}

// boolAsClusterRequestFeaturesInnerValue is a convenience function that returns bool wrapped in ClusterRequestFeaturesInnerValue
func BoolAsClusterRequestFeaturesInnerValue(v *bool) ClusterRequestFeaturesInnerValue {
	return ClusterRequestFeaturesInnerValue{
		Bool: v,
	}
}

// stringAsClusterRequestFeaturesInnerValue is a convenience function that returns string wrapped in ClusterRequestFeaturesInnerValue
func StringAsClusterRequestFeaturesInnerValue(v *string) ClusterRequestFeaturesInnerValue {
	return ClusterRequestFeaturesInnerValue{
		String: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ClusterRequestFeaturesInnerValue) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	match := 0
	// try to unmarshal data into ClusterFeatureAwsExistingVpc
	err = newStrictDecoder(data).Decode(&dst.ClusterFeatureAwsExistingVpc)
	if err == nil {
		jsonClusterFeatureAwsExistingVpc, _ := json.Marshal(dst.ClusterFeatureAwsExistingVpc)
		if string(jsonClusterFeatureAwsExistingVpc) == "{}" { // empty struct
			dst.ClusterFeatureAwsExistingVpc = nil
		} else {
			match++
		}
	} else {
		dst.ClusterFeatureAwsExistingVpc = nil
	}

	// try to unmarshal data into Bool
	err = newStrictDecoder(data).Decode(&dst.Bool)
	if err == nil {
		jsonBool, _ := json.Marshal(dst.Bool)
		if string(jsonBool) == "{}" { // empty struct
			dst.Bool = nil
		} else {
			match++
		}
	} else {
		dst.Bool = nil
	}

	// try to unmarshal data into String
	err = newStrictDecoder(data).Decode(&dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			match++
		}
	} else {
		dst.String = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ClusterFeatureAwsExistingVpc = nil
		dst.Bool = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ClusterRequestFeaturesInnerValue)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ClusterRequestFeaturesInnerValue)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ClusterRequestFeaturesInnerValue) MarshalJSON() ([]byte, error) {
	if src.ClusterFeatureAwsExistingVpc != nil {
		return json.Marshal(&src.ClusterFeatureAwsExistingVpc)
	}

	if src.Bool != nil {
		return json.Marshal(&src.Bool)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ClusterRequestFeaturesInnerValue) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.ClusterFeatureAwsExistingVpc != nil {
		return obj.ClusterFeatureAwsExistingVpc
	}

	if obj.Bool != nil {
		return obj.Bool
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableClusterRequestFeaturesInnerValue struct {
	value *ClusterRequestFeaturesInnerValue
	isSet bool
}

func (v NullableClusterRequestFeaturesInnerValue) Get() *ClusterRequestFeaturesInnerValue {
	return v.value
}

func (v *NullableClusterRequestFeaturesInnerValue) Set(val *ClusterRequestFeaturesInnerValue) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterRequestFeaturesInnerValue) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterRequestFeaturesInnerValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterRequestFeaturesInnerValue(val *ClusterRequestFeaturesInnerValue) *NullableClusterRequestFeaturesInnerValue {
	return &NullableClusterRequestFeaturesInnerValue{value: val, isSet: true}
}

func (v NullableClusterRequestFeaturesInnerValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterRequestFeaturesInnerValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
