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

// DeploymentHistoryEnvironmentAllOf struct for DeploymentHistoryEnvironmentAllOf
type DeploymentHistoryEnvironmentAllOf struct {
	Status       *GlobalDeploymentStatus        `json:"status,omitempty"`
	Applications []DeploymentHistoryApplication `json:"applications,omitempty"`
	Databases    []DeploymentHistoryDatabase    `json:"databases,omitempty"`
}

// NewDeploymentHistoryEnvironmentAllOf instantiates a new DeploymentHistoryEnvironmentAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploymentHistoryEnvironmentAllOf() *DeploymentHistoryEnvironmentAllOf {
	this := DeploymentHistoryEnvironmentAllOf{}
	return &this
}

// NewDeploymentHistoryEnvironmentAllOfWithDefaults instantiates a new DeploymentHistoryEnvironmentAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentHistoryEnvironmentAllOfWithDefaults() *DeploymentHistoryEnvironmentAllOf {
	this := DeploymentHistoryEnvironmentAllOf{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DeploymentHistoryEnvironmentAllOf) GetStatus() GlobalDeploymentStatus {
	if o == nil || o.Status == nil {
		var ret GlobalDeploymentStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentHistoryEnvironmentAllOf) GetStatusOk() (*GlobalDeploymentStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DeploymentHistoryEnvironmentAllOf) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given GlobalDeploymentStatus and assigns it to the Status field.
func (o *DeploymentHistoryEnvironmentAllOf) SetStatus(v GlobalDeploymentStatus) {
	o.Status = &v
}

// GetApplications returns the Applications field value if set, zero value otherwise.
func (o *DeploymentHistoryEnvironmentAllOf) GetApplications() []DeploymentHistoryApplication {
	if o == nil || o.Applications == nil {
		var ret []DeploymentHistoryApplication
		return ret
	}
	return o.Applications
}

// GetApplicationsOk returns a tuple with the Applications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentHistoryEnvironmentAllOf) GetApplicationsOk() ([]DeploymentHistoryApplication, bool) {
	if o == nil || o.Applications == nil {
		return nil, false
	}
	return o.Applications, true
}

// HasApplications returns a boolean if a field has been set.
func (o *DeploymentHistoryEnvironmentAllOf) HasApplications() bool {
	if o != nil && o.Applications != nil {
		return true
	}

	return false
}

// SetApplications gets a reference to the given []DeploymentHistoryApplication and assigns it to the Applications field.
func (o *DeploymentHistoryEnvironmentAllOf) SetApplications(v []DeploymentHistoryApplication) {
	o.Applications = v
}

// GetDatabases returns the Databases field value if set, zero value otherwise.
func (o *DeploymentHistoryEnvironmentAllOf) GetDatabases() []DeploymentHistoryDatabase {
	if o == nil || o.Databases == nil {
		var ret []DeploymentHistoryDatabase
		return ret
	}
	return o.Databases
}

// GetDatabasesOk returns a tuple with the Databases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentHistoryEnvironmentAllOf) GetDatabasesOk() ([]DeploymentHistoryDatabase, bool) {
	if o == nil || o.Databases == nil {
		return nil, false
	}
	return o.Databases, true
}

// HasDatabases returns a boolean if a field has been set.
func (o *DeploymentHistoryEnvironmentAllOf) HasDatabases() bool {
	if o != nil && o.Databases != nil {
		return true
	}

	return false
}

// SetDatabases gets a reference to the given []DeploymentHistoryDatabase and assigns it to the Databases field.
func (o *DeploymentHistoryEnvironmentAllOf) SetDatabases(v []DeploymentHistoryDatabase) {
	o.Databases = v
}

func (o DeploymentHistoryEnvironmentAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Applications != nil {
		toSerialize["applications"] = o.Applications
	}
	if o.Databases != nil {
		toSerialize["databases"] = o.Databases
	}
	return json.Marshal(toSerialize)
}

type NullableDeploymentHistoryEnvironmentAllOf struct {
	value *DeploymentHistoryEnvironmentAllOf
	isSet bool
}

func (v NullableDeploymentHistoryEnvironmentAllOf) Get() *DeploymentHistoryEnvironmentAllOf {
	return v.value
}

func (v *NullableDeploymentHistoryEnvironmentAllOf) Set(val *DeploymentHistoryEnvironmentAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploymentHistoryEnvironmentAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploymentHistoryEnvironmentAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploymentHistoryEnvironmentAllOf(val *DeploymentHistoryEnvironmentAllOf) *NullableDeploymentHistoryEnvironmentAllOf {
	return &NullableDeploymentHistoryEnvironmentAllOf{value: val, isSet: true}
}

func (v NullableDeploymentHistoryEnvironmentAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploymentHistoryEnvironmentAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
