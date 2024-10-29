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

// checks if the DeploymentHistoryServiceDetailsOneOf3Repository type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeploymentHistoryServiceDetailsOneOf3Repository{}

// DeploymentHistoryServiceDetailsOneOf3Repository struct for DeploymentHistoryServiceDetailsOneOf3Repository
type DeploymentHistoryServiceDetailsOneOf3Repository struct {
	ChartName            *string `json:"chart_name,omitempty"`
	ChartVersion         *string `json:"chart_version,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DeploymentHistoryServiceDetailsOneOf3Repository DeploymentHistoryServiceDetailsOneOf3Repository

// NewDeploymentHistoryServiceDetailsOneOf3Repository instantiates a new DeploymentHistoryServiceDetailsOneOf3Repository object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploymentHistoryServiceDetailsOneOf3Repository() *DeploymentHistoryServiceDetailsOneOf3Repository {
	this := DeploymentHistoryServiceDetailsOneOf3Repository{}
	return &this
}

// NewDeploymentHistoryServiceDetailsOneOf3RepositoryWithDefaults instantiates a new DeploymentHistoryServiceDetailsOneOf3Repository object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentHistoryServiceDetailsOneOf3RepositoryWithDefaults() *DeploymentHistoryServiceDetailsOneOf3Repository {
	this := DeploymentHistoryServiceDetailsOneOf3Repository{}
	return &this
}

// GetChartName returns the ChartName field value if set, zero value otherwise.
func (o *DeploymentHistoryServiceDetailsOneOf3Repository) GetChartName() string {
	if o == nil || IsNil(o.ChartName) {
		var ret string
		return ret
	}
	return *o.ChartName
}

// GetChartNameOk returns a tuple with the ChartName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentHistoryServiceDetailsOneOf3Repository) GetChartNameOk() (*string, bool) {
	if o == nil || IsNil(o.ChartName) {
		return nil, false
	}
	return o.ChartName, true
}

// HasChartName returns a boolean if a field has been set.
func (o *DeploymentHistoryServiceDetailsOneOf3Repository) HasChartName() bool {
	if o != nil && !IsNil(o.ChartName) {
		return true
	}

	return false
}

// SetChartName gets a reference to the given string and assigns it to the ChartName field.
func (o *DeploymentHistoryServiceDetailsOneOf3Repository) SetChartName(v string) {
	o.ChartName = &v
}

// GetChartVersion returns the ChartVersion field value if set, zero value otherwise.
func (o *DeploymentHistoryServiceDetailsOneOf3Repository) GetChartVersion() string {
	if o == nil || IsNil(o.ChartVersion) {
		var ret string
		return ret
	}
	return *o.ChartVersion
}

// GetChartVersionOk returns a tuple with the ChartVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentHistoryServiceDetailsOneOf3Repository) GetChartVersionOk() (*string, bool) {
	if o == nil || IsNil(o.ChartVersion) {
		return nil, false
	}
	return o.ChartVersion, true
}

// HasChartVersion returns a boolean if a field has been set.
func (o *DeploymentHistoryServiceDetailsOneOf3Repository) HasChartVersion() bool {
	if o != nil && !IsNil(o.ChartVersion) {
		return true
	}

	return false
}

// SetChartVersion gets a reference to the given string and assigns it to the ChartVersion field.
func (o *DeploymentHistoryServiceDetailsOneOf3Repository) SetChartVersion(v string) {
	o.ChartVersion = &v
}

func (o DeploymentHistoryServiceDetailsOneOf3Repository) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeploymentHistoryServiceDetailsOneOf3Repository) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ChartName) {
		toSerialize["chart_name"] = o.ChartName
	}
	if !IsNil(o.ChartVersion) {
		toSerialize["chart_version"] = o.ChartVersion
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DeploymentHistoryServiceDetailsOneOf3Repository) UnmarshalJSON(data []byte) (err error) {
	varDeploymentHistoryServiceDetailsOneOf3Repository := _DeploymentHistoryServiceDetailsOneOf3Repository{}

	err = json.Unmarshal(data, &varDeploymentHistoryServiceDetailsOneOf3Repository)

	if err != nil {
		return err
	}

	*o = DeploymentHistoryServiceDetailsOneOf3Repository(varDeploymentHistoryServiceDetailsOneOf3Repository)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "chart_name")
		delete(additionalProperties, "chart_version")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeploymentHistoryServiceDetailsOneOf3Repository struct {
	value *DeploymentHistoryServiceDetailsOneOf3Repository
	isSet bool
}

func (v NullableDeploymentHistoryServiceDetailsOneOf3Repository) Get() *DeploymentHistoryServiceDetailsOneOf3Repository {
	return v.value
}

func (v *NullableDeploymentHistoryServiceDetailsOneOf3Repository) Set(val *DeploymentHistoryServiceDetailsOneOf3Repository) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploymentHistoryServiceDetailsOneOf3Repository) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploymentHistoryServiceDetailsOneOf3Repository) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploymentHistoryServiceDetailsOneOf3Repository(val *DeploymentHistoryServiceDetailsOneOf3Repository) *NullableDeploymentHistoryServiceDetailsOneOf3Repository {
	return &NullableDeploymentHistoryServiceDetailsOneOf3Repository{value: val, isSet: true}
}

func (v NullableDeploymentHistoryServiceDetailsOneOf3Repository) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploymentHistoryServiceDetailsOneOf3Repository) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
