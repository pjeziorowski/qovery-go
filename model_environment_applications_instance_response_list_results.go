/*
[BETA] Qovery API

- Qovery is the fastest way to deploy your full-stack apps on any Cloud provider. - ℹ️ The API is in Beta and still in progress. Some endpoints are not available yet.  # Authentication  <!-- ReDoc-Inject: <security-definitions> -->

API version: 1.0.0
Contact: support+api+documentation@qovery.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package qovery

import (
	"encoding/json"
)

// EnvironmentApplicationsInstanceResponseListResults struct for EnvironmentApplicationsInstanceResponseListResults
type EnvironmentApplicationsInstanceResponseListResults struct {
	Application string `json:"application"`
	Instances []InstanceResponse `json:"instances"`
}

// NewEnvironmentApplicationsInstanceResponseListResults instantiates a new EnvironmentApplicationsInstanceResponseListResults object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentApplicationsInstanceResponseListResults(application string, instances []InstanceResponse) *EnvironmentApplicationsInstanceResponseListResults {
	this := EnvironmentApplicationsInstanceResponseListResults{}
	this.Application = application
	this.Instances = instances
	return &this
}

// NewEnvironmentApplicationsInstanceResponseListResultsWithDefaults instantiates a new EnvironmentApplicationsInstanceResponseListResults object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentApplicationsInstanceResponseListResultsWithDefaults() *EnvironmentApplicationsInstanceResponseListResults {
	this := EnvironmentApplicationsInstanceResponseListResults{}
	return &this
}

// GetApplication returns the Application field value
func (o *EnvironmentApplicationsInstanceResponseListResults) GetApplication() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Application
}

// GetApplicationOk returns a tuple with the Application field value
// and a boolean to check if the value has been set.
func (o *EnvironmentApplicationsInstanceResponseListResults) GetApplicationOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Application, true
}

// SetApplication sets field value
func (o *EnvironmentApplicationsInstanceResponseListResults) SetApplication(v string) {
	o.Application = v
}

// GetInstances returns the Instances field value
func (o *EnvironmentApplicationsInstanceResponseListResults) GetInstances() []InstanceResponse {
	if o == nil {
		var ret []InstanceResponse
		return ret
	}

	return o.Instances
}

// GetInstancesOk returns a tuple with the Instances field value
// and a boolean to check if the value has been set.
func (o *EnvironmentApplicationsInstanceResponseListResults) GetInstancesOk() (*[]InstanceResponse, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Instances, true
}

// SetInstances sets field value
func (o *EnvironmentApplicationsInstanceResponseListResults) SetInstances(v []InstanceResponse) {
	o.Instances = v
}

func (o EnvironmentApplicationsInstanceResponseListResults) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["application"] = o.Application
	}
	if true {
		toSerialize["instances"] = o.Instances
	}
	return json.Marshal(toSerialize)
}

type NullableEnvironmentApplicationsInstanceResponseListResults struct {
	value *EnvironmentApplicationsInstanceResponseListResults
	isSet bool
}

func (v NullableEnvironmentApplicationsInstanceResponseListResults) Get() *EnvironmentApplicationsInstanceResponseListResults {
	return v.value
}

func (v *NullableEnvironmentApplicationsInstanceResponseListResults) Set(val *EnvironmentApplicationsInstanceResponseListResults) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentApplicationsInstanceResponseListResults) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentApplicationsInstanceResponseListResults) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentApplicationsInstanceResponseListResults(val *EnvironmentApplicationsInstanceResponseListResults) *NullableEnvironmentApplicationsInstanceResponseListResults {
	return &NullableEnvironmentApplicationsInstanceResponseListResults{value: val, isSet: true}
}

func (v NullableEnvironmentApplicationsInstanceResponseListResults) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentApplicationsInstanceResponseListResults) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


