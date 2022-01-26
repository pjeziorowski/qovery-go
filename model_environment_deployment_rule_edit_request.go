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
	"time"
)

// EnvironmentDeploymentRuleEditRequest struct for EnvironmentDeploymentRuleEditRequest
type EnvironmentDeploymentRuleEditRequest struct {
	AutoDeploy *bool `json:"auto_deploy,omitempty"`
	AutoDelete *bool `json:"auto_delete,omitempty"`
	AutoStop *bool `json:"auto_stop,omitempty"`
	// specify value only if auto_stop = false
	Timezone *string `json:"timezone,omitempty"`
	// specify value only if auto_stop = false
	StartTime NullableTime `json:"start_time,omitempty"`
	// specify value only if auto_stop = false
	StopTime NullableTime `json:"stop_time,omitempty"`
	// specify value only if auto_stop = false
	Weekdays []string `json:"weekdays,omitempty"`
}

// NewEnvironmentDeploymentRuleEditRequest instantiates a new EnvironmentDeploymentRuleEditRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentDeploymentRuleEditRequest() *EnvironmentDeploymentRuleEditRequest {
	this := EnvironmentDeploymentRuleEditRequest{}
	var autoDeploy bool = true
	this.AutoDeploy = &autoDeploy
	var autoDelete bool = false
	this.AutoDelete = &autoDelete
	var autoStop bool = false
	this.AutoStop = &autoStop
	var timezone string = "Europe/London"
	this.Timezone = &timezone
	return &this
}

// NewEnvironmentDeploymentRuleEditRequestWithDefaults instantiates a new EnvironmentDeploymentRuleEditRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentDeploymentRuleEditRequestWithDefaults() *EnvironmentDeploymentRuleEditRequest {
	this := EnvironmentDeploymentRuleEditRequest{}
	var autoDeploy bool = true
	this.AutoDeploy = &autoDeploy
	var autoDelete bool = false
	this.AutoDelete = &autoDelete
	var autoStop bool = false
	this.AutoStop = &autoStop
	var timezone string = "Europe/London"
	this.Timezone = &timezone
	return &this
}

// GetAutoDeploy returns the AutoDeploy field value if set, zero value otherwise.
func (o *EnvironmentDeploymentRuleEditRequest) GetAutoDeploy() bool {
	if o == nil || o.AutoDeploy == nil {
		var ret bool
		return ret
	}
	return *o.AutoDeploy
}

// GetAutoDeployOk returns a tuple with the AutoDeploy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentDeploymentRuleEditRequest) GetAutoDeployOk() (*bool, bool) {
	if o == nil || o.AutoDeploy == nil {
		return nil, false
	}
	return o.AutoDeploy, true
}

// HasAutoDeploy returns a boolean if a field has been set.
func (o *EnvironmentDeploymentRuleEditRequest) HasAutoDeploy() bool {
	if o != nil && o.AutoDeploy != nil {
		return true
	}

	return false
}

// SetAutoDeploy gets a reference to the given bool and assigns it to the AutoDeploy field.
func (o *EnvironmentDeploymentRuleEditRequest) SetAutoDeploy(v bool) {
	o.AutoDeploy = &v
}

// GetAutoDelete returns the AutoDelete field value if set, zero value otherwise.
func (o *EnvironmentDeploymentRuleEditRequest) GetAutoDelete() bool {
	if o == nil || o.AutoDelete == nil {
		var ret bool
		return ret
	}
	return *o.AutoDelete
}

// GetAutoDeleteOk returns a tuple with the AutoDelete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentDeploymentRuleEditRequest) GetAutoDeleteOk() (*bool, bool) {
	if o == nil || o.AutoDelete == nil {
		return nil, false
	}
	return o.AutoDelete, true
}

// HasAutoDelete returns a boolean if a field has been set.
func (o *EnvironmentDeploymentRuleEditRequest) HasAutoDelete() bool {
	if o != nil && o.AutoDelete != nil {
		return true
	}

	return false
}

// SetAutoDelete gets a reference to the given bool and assigns it to the AutoDelete field.
func (o *EnvironmentDeploymentRuleEditRequest) SetAutoDelete(v bool) {
	o.AutoDelete = &v
}

// GetAutoStop returns the AutoStop field value if set, zero value otherwise.
func (o *EnvironmentDeploymentRuleEditRequest) GetAutoStop() bool {
	if o == nil || o.AutoStop == nil {
		var ret bool
		return ret
	}
	return *o.AutoStop
}

// GetAutoStopOk returns a tuple with the AutoStop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentDeploymentRuleEditRequest) GetAutoStopOk() (*bool, bool) {
	if o == nil || o.AutoStop == nil {
		return nil, false
	}
	return o.AutoStop, true
}

// HasAutoStop returns a boolean if a field has been set.
func (o *EnvironmentDeploymentRuleEditRequest) HasAutoStop() bool {
	if o != nil && o.AutoStop != nil {
		return true
	}

	return false
}

// SetAutoStop gets a reference to the given bool and assigns it to the AutoStop field.
func (o *EnvironmentDeploymentRuleEditRequest) SetAutoStop(v bool) {
	o.AutoStop = &v
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *EnvironmentDeploymentRuleEditRequest) GetTimezone() string {
	if o == nil || o.Timezone == nil {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentDeploymentRuleEditRequest) GetTimezoneOk() (*string, bool) {
	if o == nil || o.Timezone == nil {
		return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *EnvironmentDeploymentRuleEditRequest) HasTimezone() bool {
	if o != nil && o.Timezone != nil {
		return true
	}

	return false
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *EnvironmentDeploymentRuleEditRequest) SetTimezone(v string) {
	o.Timezone = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentDeploymentRuleEditRequest) GetStartTime() time.Time {
	if o == nil || o.StartTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.StartTime.Get()
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentDeploymentRuleEditRequest) GetStartTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StartTime.Get(), o.StartTime.IsSet()
}

// HasStartTime returns a boolean if a field has been set.
func (o *EnvironmentDeploymentRuleEditRequest) HasStartTime() bool {
	if o != nil && o.StartTime.IsSet() {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given NullableTime and assigns it to the StartTime field.
func (o *EnvironmentDeploymentRuleEditRequest) SetStartTime(v time.Time) {
	o.StartTime.Set(&v)
}
// SetStartTimeNil sets the value for StartTime to be an explicit nil
func (o *EnvironmentDeploymentRuleEditRequest) SetStartTimeNil() {
	o.StartTime.Set(nil)
}

// UnsetStartTime ensures that no value is present for StartTime, not even an explicit nil
func (o *EnvironmentDeploymentRuleEditRequest) UnsetStartTime() {
	o.StartTime.Unset()
}

// GetStopTime returns the StopTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentDeploymentRuleEditRequest) GetStopTime() time.Time {
	if o == nil || o.StopTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.StopTime.Get()
}

// GetStopTimeOk returns a tuple with the StopTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentDeploymentRuleEditRequest) GetStopTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StopTime.Get(), o.StopTime.IsSet()
}

// HasStopTime returns a boolean if a field has been set.
func (o *EnvironmentDeploymentRuleEditRequest) HasStopTime() bool {
	if o != nil && o.StopTime.IsSet() {
		return true
	}

	return false
}

// SetStopTime gets a reference to the given NullableTime and assigns it to the StopTime field.
func (o *EnvironmentDeploymentRuleEditRequest) SetStopTime(v time.Time) {
	o.StopTime.Set(&v)
}
// SetStopTimeNil sets the value for StopTime to be an explicit nil
func (o *EnvironmentDeploymentRuleEditRequest) SetStopTimeNil() {
	o.StopTime.Set(nil)
}

// UnsetStopTime ensures that no value is present for StopTime, not even an explicit nil
func (o *EnvironmentDeploymentRuleEditRequest) UnsetStopTime() {
	o.StopTime.Unset()
}

// GetWeekdays returns the Weekdays field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentDeploymentRuleEditRequest) GetWeekdays() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.Weekdays
}

// GetWeekdaysOk returns a tuple with the Weekdays field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentDeploymentRuleEditRequest) GetWeekdaysOk() (*[]string, bool) {
	if o == nil || o.Weekdays == nil {
		return nil, false
	}
	return &o.Weekdays, true
}

// HasWeekdays returns a boolean if a field has been set.
func (o *EnvironmentDeploymentRuleEditRequest) HasWeekdays() bool {
	if o != nil && o.Weekdays != nil {
		return true
	}

	return false
}

// SetWeekdays gets a reference to the given []string and assigns it to the Weekdays field.
func (o *EnvironmentDeploymentRuleEditRequest) SetWeekdays(v []string) {
	o.Weekdays = v
}

func (o EnvironmentDeploymentRuleEditRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AutoDeploy != nil {
		toSerialize["auto_deploy"] = o.AutoDeploy
	}
	if o.AutoDelete != nil {
		toSerialize["auto_delete"] = o.AutoDelete
	}
	if o.AutoStop != nil {
		toSerialize["auto_stop"] = o.AutoStop
	}
	if o.Timezone != nil {
		toSerialize["timezone"] = o.Timezone
	}
	if o.StartTime.IsSet() {
		toSerialize["start_time"] = o.StartTime.Get()
	}
	if o.StopTime.IsSet() {
		toSerialize["stop_time"] = o.StopTime.Get()
	}
	if o.Weekdays != nil {
		toSerialize["weekdays"] = o.Weekdays
	}
	return json.Marshal(toSerialize)
}

type NullableEnvironmentDeploymentRuleEditRequest struct {
	value *EnvironmentDeploymentRuleEditRequest
	isSet bool
}

func (v NullableEnvironmentDeploymentRuleEditRequest) Get() *EnvironmentDeploymentRuleEditRequest {
	return v.value
}

func (v *NullableEnvironmentDeploymentRuleEditRequest) Set(val *EnvironmentDeploymentRuleEditRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentDeploymentRuleEditRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentDeploymentRuleEditRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentDeploymentRuleEditRequest(val *EnvironmentDeploymentRuleEditRequest) *NullableEnvironmentDeploymentRuleEditRequest {
	return &NullableEnvironmentDeploymentRuleEditRequest{value: val, isSet: true}
}

func (v NullableEnvironmentDeploymentRuleEditRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentDeploymentRuleEditRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


