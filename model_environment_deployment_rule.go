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

// EnvironmentDeploymentRule struct for EnvironmentDeploymentRule
type EnvironmentDeploymentRule struct {
	Id              string        `json:"id"`
	CreatedAt       time.Time     `json:"created_at"`
	UpdatedAt       *time.Time    `json:"updated_at,omitempty"`
	AutoDeploy      *bool         `json:"auto_deploy,omitempty"`
	OnDemandPreview *bool         `json:"on_demand_preview,omitempty"`
	AutoStop        *bool         `json:"auto_stop,omitempty"`
	AutoDelete      *bool         `json:"auto_delete,omitempty"`
	AutoPreview     *bool         `json:"auto_preview,omitempty"`
	Timezone        string        `json:"timezone"`
	StartTime       time.Time     `json:"start_time"`
	StopTime        time.Time     `json:"stop_time"`
	Weekdays        []WeekdayEnum `json:"weekdays"`
}

// NewEnvironmentDeploymentRule instantiates a new EnvironmentDeploymentRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentDeploymentRule(id string, createdAt time.Time, timezone string, startTime time.Time, stopTime time.Time, weekdays []WeekdayEnum) *EnvironmentDeploymentRule {
	this := EnvironmentDeploymentRule{}
	this.Id = id
	this.CreatedAt = createdAt
	var autoDeploy bool = true
	this.AutoDeploy = &autoDeploy
	var onDemandPreview bool = false
	this.OnDemandPreview = &onDemandPreview
	var autoStop bool = false
	this.AutoStop = &autoStop
	var autoDelete bool = false
	this.AutoDelete = &autoDelete
	var autoPreview bool = false
	this.AutoPreview = &autoPreview
	this.Timezone = timezone
	this.StartTime = startTime
	this.StopTime = stopTime
	this.Weekdays = weekdays
	return &this
}

// NewEnvironmentDeploymentRuleWithDefaults instantiates a new EnvironmentDeploymentRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentDeploymentRuleWithDefaults() *EnvironmentDeploymentRule {
	this := EnvironmentDeploymentRule{}
	var autoDeploy bool = true
	this.AutoDeploy = &autoDeploy
	var onDemandPreview bool = false
	this.OnDemandPreview = &onDemandPreview
	var autoStop bool = false
	this.AutoStop = &autoStop
	var autoDelete bool = false
	this.AutoDelete = &autoDelete
	var autoPreview bool = false
	this.AutoPreview = &autoPreview
	return &this
}

// GetId returns the Id field value
func (o *EnvironmentDeploymentRule) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EnvironmentDeploymentRule) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EnvironmentDeploymentRule) SetId(v string) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *EnvironmentDeploymentRule) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *EnvironmentDeploymentRule) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *EnvironmentDeploymentRule) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *EnvironmentDeploymentRule) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentDeploymentRule) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *EnvironmentDeploymentRule) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *EnvironmentDeploymentRule) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetAutoDeploy returns the AutoDeploy field value if set, zero value otherwise.
func (o *EnvironmentDeploymentRule) GetAutoDeploy() bool {
	if o == nil || o.AutoDeploy == nil {
		var ret bool
		return ret
	}
	return *o.AutoDeploy
}

// GetAutoDeployOk returns a tuple with the AutoDeploy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentDeploymentRule) GetAutoDeployOk() (*bool, bool) {
	if o == nil || o.AutoDeploy == nil {
		return nil, false
	}
	return o.AutoDeploy, true
}

// HasAutoDeploy returns a boolean if a field has been set.
func (o *EnvironmentDeploymentRule) HasAutoDeploy() bool {
	if o != nil && o.AutoDeploy != nil {
		return true
	}

	return false
}

// SetAutoDeploy gets a reference to the given bool and assigns it to the AutoDeploy field.
func (o *EnvironmentDeploymentRule) SetAutoDeploy(v bool) {
	o.AutoDeploy = &v
}

// GetOnDemandPreview returns the OnDemandPreview field value if set, zero value otherwise.
func (o *EnvironmentDeploymentRule) GetOnDemandPreview() bool {
	if o == nil || o.OnDemandPreview == nil {
		var ret bool
		return ret
	}
	return *o.OnDemandPreview
}

// GetOnDemandPreviewOk returns a tuple with the OnDemandPreview field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentDeploymentRule) GetOnDemandPreviewOk() (*bool, bool) {
	if o == nil || o.OnDemandPreview == nil {
		return nil, false
	}
	return o.OnDemandPreview, true
}

// HasOnDemandPreview returns a boolean if a field has been set.
func (o *EnvironmentDeploymentRule) HasOnDemandPreview() bool {
	if o != nil && o.OnDemandPreview != nil {
		return true
	}

	return false
}

// SetOnDemandPreview gets a reference to the given bool and assigns it to the OnDemandPreview field.
func (o *EnvironmentDeploymentRule) SetOnDemandPreview(v bool) {
	o.OnDemandPreview = &v
}

// GetAutoStop returns the AutoStop field value if set, zero value otherwise.
func (o *EnvironmentDeploymentRule) GetAutoStop() bool {
	if o == nil || o.AutoStop == nil {
		var ret bool
		return ret
	}
	return *o.AutoStop
}

// GetAutoStopOk returns a tuple with the AutoStop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentDeploymentRule) GetAutoStopOk() (*bool, bool) {
	if o == nil || o.AutoStop == nil {
		return nil, false
	}
	return o.AutoStop, true
}

// HasAutoStop returns a boolean if a field has been set.
func (o *EnvironmentDeploymentRule) HasAutoStop() bool {
	if o != nil && o.AutoStop != nil {
		return true
	}

	return false
}

// SetAutoStop gets a reference to the given bool and assigns it to the AutoStop field.
func (o *EnvironmentDeploymentRule) SetAutoStop(v bool) {
	o.AutoStop = &v
}

// GetAutoDelete returns the AutoDelete field value if set, zero value otherwise.
func (o *EnvironmentDeploymentRule) GetAutoDelete() bool {
	if o == nil || o.AutoDelete == nil {
		var ret bool
		return ret
	}
	return *o.AutoDelete
}

// GetAutoDeleteOk returns a tuple with the AutoDelete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentDeploymentRule) GetAutoDeleteOk() (*bool, bool) {
	if o == nil || o.AutoDelete == nil {
		return nil, false
	}
	return o.AutoDelete, true
}

// HasAutoDelete returns a boolean if a field has been set.
func (o *EnvironmentDeploymentRule) HasAutoDelete() bool {
	if o != nil && o.AutoDelete != nil {
		return true
	}

	return false
}

// SetAutoDelete gets a reference to the given bool and assigns it to the AutoDelete field.
func (o *EnvironmentDeploymentRule) SetAutoDelete(v bool) {
	o.AutoDelete = &v
}

// GetAutoPreview returns the AutoPreview field value if set, zero value otherwise.
func (o *EnvironmentDeploymentRule) GetAutoPreview() bool {
	if o == nil || o.AutoPreview == nil {
		var ret bool
		return ret
	}
	return *o.AutoPreview
}

// GetAutoPreviewOk returns a tuple with the AutoPreview field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentDeploymentRule) GetAutoPreviewOk() (*bool, bool) {
	if o == nil || o.AutoPreview == nil {
		return nil, false
	}
	return o.AutoPreview, true
}

// HasAutoPreview returns a boolean if a field has been set.
func (o *EnvironmentDeploymentRule) HasAutoPreview() bool {
	if o != nil && o.AutoPreview != nil {
		return true
	}

	return false
}

// SetAutoPreview gets a reference to the given bool and assigns it to the AutoPreview field.
func (o *EnvironmentDeploymentRule) SetAutoPreview(v bool) {
	o.AutoPreview = &v
}

// GetTimezone returns the Timezone field value
func (o *EnvironmentDeploymentRule) GetTimezone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value
// and a boolean to check if the value has been set.
func (o *EnvironmentDeploymentRule) GetTimezoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timezone, true
}

// SetTimezone sets field value
func (o *EnvironmentDeploymentRule) SetTimezone(v string) {
	o.Timezone = v
}

// GetStartTime returns the StartTime field value
func (o *EnvironmentDeploymentRule) GetStartTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value
// and a boolean to check if the value has been set.
func (o *EnvironmentDeploymentRule) GetStartTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartTime, true
}

// SetStartTime sets field value
func (o *EnvironmentDeploymentRule) SetStartTime(v time.Time) {
	o.StartTime = v
}

// GetStopTime returns the StopTime field value
func (o *EnvironmentDeploymentRule) GetStopTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StopTime
}

// GetStopTimeOk returns a tuple with the StopTime field value
// and a boolean to check if the value has been set.
func (o *EnvironmentDeploymentRule) GetStopTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StopTime, true
}

// SetStopTime sets field value
func (o *EnvironmentDeploymentRule) SetStopTime(v time.Time) {
	o.StopTime = v
}

// GetWeekdays returns the Weekdays field value
func (o *EnvironmentDeploymentRule) GetWeekdays() []WeekdayEnum {
	if o == nil {
		var ret []WeekdayEnum
		return ret
	}

	return o.Weekdays
}

// GetWeekdaysOk returns a tuple with the Weekdays field value
// and a boolean to check if the value has been set.
func (o *EnvironmentDeploymentRule) GetWeekdaysOk() ([]WeekdayEnum, bool) {
	if o == nil {
		return nil, false
	}
	return o.Weekdays, true
}

// SetWeekdays sets field value
func (o *EnvironmentDeploymentRule) SetWeekdays(v []WeekdayEnum) {
	o.Weekdays = v
}

func (o EnvironmentDeploymentRule) MarshalJSON() ([]byte, error) {
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
	if o.AutoDeploy != nil {
		toSerialize["auto_deploy"] = o.AutoDeploy
	}
	if o.OnDemandPreview != nil {
		toSerialize["on_demand_preview"] = o.OnDemandPreview
	}
	if o.AutoStop != nil {
		toSerialize["auto_stop"] = o.AutoStop
	}
	if o.AutoDelete != nil {
		toSerialize["auto_delete"] = o.AutoDelete
	}
	if o.AutoPreview != nil {
		toSerialize["auto_preview"] = o.AutoPreview
	}
	if true {
		toSerialize["timezone"] = o.Timezone
	}
	if true {
		toSerialize["start_time"] = o.StartTime
	}
	if true {
		toSerialize["stop_time"] = o.StopTime
	}
	if true {
		toSerialize["weekdays"] = o.Weekdays
	}
	return json.Marshal(toSerialize)
}

type NullableEnvironmentDeploymentRule struct {
	value *EnvironmentDeploymentRule
	isSet bool
}

func (v NullableEnvironmentDeploymentRule) Get() *EnvironmentDeploymentRule {
	return v.value
}

func (v *NullableEnvironmentDeploymentRule) Set(val *EnvironmentDeploymentRule) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentDeploymentRule) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentDeploymentRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentDeploymentRule(val *EnvironmentDeploymentRule) *NullableEnvironmentDeploymentRule {
	return &NullableEnvironmentDeploymentRule{value: val, isSet: true}
}

func (v NullableEnvironmentDeploymentRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentDeploymentRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
