/*
 * [BETA] Qovery API
 *
 * - Qovery is the fastest way to deploy your full-stack apps on any Cloud provider. - ℹ️ The API is in Beta and still in progress. Some endpoints are not available yet.
 *
 * API version: 1.0.0
 * Contact: support+api+documentation@qovery.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package qovery

import (
	"encoding/json"
)

// ClusterRequest struct for ClusterRequest
type ClusterRequest struct {
	// name is case insensitive
	Name                 string  `json:"name"`
	Description          *string `json:"description,omitempty"`
	AutoUpdate           *bool   `json:"auto_update,omitempty"`
	EnableHistoricMetric *bool   `json:"enable_historic_metric,omitempty"`
	EnableHistoricLog    *bool   `json:"enable_historic_log,omitempty"`
	// unit is millicores (m). 1000m = 1 cpu
	Cpu *float32 `json:"cpu,omitempty"`
	// unit is MB. 1024 MB = 1GB
	Memory          *float32 `json:"memory,omitempty"`
	MinRunningNodes *int32   `json:"min_running_nodes,omitempty"`
	MaxRunningNodes *int32   `json:"max_running_nodes,omitempty"`
}

// NewClusterRequest instantiates a new ClusterRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterRequest(name string) *ClusterRequest {
	this := ClusterRequest{}
	this.Name = name
	var cpu float32 = 250
	this.Cpu = &cpu
	var memory float32 = 256
	this.Memory = &memory
	var minRunningNodes int32 = 1
	this.MinRunningNodes = &minRunningNodes
	var maxRunningNodes int32 = 1
	this.MaxRunningNodes = &maxRunningNodes
	return &this
}

// NewClusterRequestWithDefaults instantiates a new ClusterRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterRequestWithDefaults() *ClusterRequest {
	this := ClusterRequest{}
	var cpu float32 = 250
	this.Cpu = &cpu
	var memory float32 = 256
	this.Memory = &memory
	var minRunningNodes int32 = 1
	this.MinRunningNodes = &minRunningNodes
	var maxRunningNodes int32 = 1
	this.MaxRunningNodes = &maxRunningNodes
	return &this
}

// GetName returns the Name field value
func (o *ClusterRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ClusterRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ClusterRequest) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ClusterRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ClusterRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ClusterRequest) SetDescription(v string) {
	o.Description = &v
}

// GetAutoUpdate returns the AutoUpdate field value if set, zero value otherwise.
func (o *ClusterRequest) GetAutoUpdate() bool {
	if o == nil || o.AutoUpdate == nil {
		var ret bool
		return ret
	}
	return *o.AutoUpdate
}

// GetAutoUpdateOk returns a tuple with the AutoUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterRequest) GetAutoUpdateOk() (*bool, bool) {
	if o == nil || o.AutoUpdate == nil {
		return nil, false
	}
	return o.AutoUpdate, true
}

// HasAutoUpdate returns a boolean if a field has been set.
func (o *ClusterRequest) HasAutoUpdate() bool {
	if o != nil && o.AutoUpdate != nil {
		return true
	}

	return false
}

// SetAutoUpdate gets a reference to the given bool and assigns it to the AutoUpdate field.
func (o *ClusterRequest) SetAutoUpdate(v bool) {
	o.AutoUpdate = &v
}

// GetEnableHistoricMetric returns the EnableHistoricMetric field value if set, zero value otherwise.
func (o *ClusterRequest) GetEnableHistoricMetric() bool {
	if o == nil || o.EnableHistoricMetric == nil {
		var ret bool
		return ret
	}
	return *o.EnableHistoricMetric
}

// GetEnableHistoricMetricOk returns a tuple with the EnableHistoricMetric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterRequest) GetEnableHistoricMetricOk() (*bool, bool) {
	if o == nil || o.EnableHistoricMetric == nil {
		return nil, false
	}
	return o.EnableHistoricMetric, true
}

// HasEnableHistoricMetric returns a boolean if a field has been set.
func (o *ClusterRequest) HasEnableHistoricMetric() bool {
	if o != nil && o.EnableHistoricMetric != nil {
		return true
	}

	return false
}

// SetEnableHistoricMetric gets a reference to the given bool and assigns it to the EnableHistoricMetric field.
func (o *ClusterRequest) SetEnableHistoricMetric(v bool) {
	o.EnableHistoricMetric = &v
}

// GetEnableHistoricLog returns the EnableHistoricLog field value if set, zero value otherwise.
func (o *ClusterRequest) GetEnableHistoricLog() bool {
	if o == nil || o.EnableHistoricLog == nil {
		var ret bool
		return ret
	}
	return *o.EnableHistoricLog
}

// GetEnableHistoricLogOk returns a tuple with the EnableHistoricLog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterRequest) GetEnableHistoricLogOk() (*bool, bool) {
	if o == nil || o.EnableHistoricLog == nil {
		return nil, false
	}
	return o.EnableHistoricLog, true
}

// HasEnableHistoricLog returns a boolean if a field has been set.
func (o *ClusterRequest) HasEnableHistoricLog() bool {
	if o != nil && o.EnableHistoricLog != nil {
		return true
	}

	return false
}

// SetEnableHistoricLog gets a reference to the given bool and assigns it to the EnableHistoricLog field.
func (o *ClusterRequest) SetEnableHistoricLog(v bool) {
	o.EnableHistoricLog = &v
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *ClusterRequest) GetCpu() float32 {
	if o == nil || o.Cpu == nil {
		var ret float32
		return ret
	}
	return *o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterRequest) GetCpuOk() (*float32, bool) {
	if o == nil || o.Cpu == nil {
		return nil, false
	}
	return o.Cpu, true
}

// HasCpu returns a boolean if a field has been set.
func (o *ClusterRequest) HasCpu() bool {
	if o != nil && o.Cpu != nil {
		return true
	}

	return false
}

// SetCpu gets a reference to the given float32 and assigns it to the Cpu field.
func (o *ClusterRequest) SetCpu(v float32) {
	o.Cpu = &v
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *ClusterRequest) GetMemory() float32 {
	if o == nil || o.Memory == nil {
		var ret float32
		return ret
	}
	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterRequest) GetMemoryOk() (*float32, bool) {
	if o == nil || o.Memory == nil {
		return nil, false
	}
	return o.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *ClusterRequest) HasMemory() bool {
	if o != nil && o.Memory != nil {
		return true
	}

	return false
}

// SetMemory gets a reference to the given float32 and assigns it to the Memory field.
func (o *ClusterRequest) SetMemory(v float32) {
	o.Memory = &v
}

// GetMinRunningNodes returns the MinRunningNodes field value if set, zero value otherwise.
func (o *ClusterRequest) GetMinRunningNodes() int32 {
	if o == nil || o.MinRunningNodes == nil {
		var ret int32
		return ret
	}
	return *o.MinRunningNodes
}

// GetMinRunningNodesOk returns a tuple with the MinRunningNodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterRequest) GetMinRunningNodesOk() (*int32, bool) {
	if o == nil || o.MinRunningNodes == nil {
		return nil, false
	}
	return o.MinRunningNodes, true
}

// HasMinRunningNodes returns a boolean if a field has been set.
func (o *ClusterRequest) HasMinRunningNodes() bool {
	if o != nil && o.MinRunningNodes != nil {
		return true
	}

	return false
}

// SetMinRunningNodes gets a reference to the given int32 and assigns it to the MinRunningNodes field.
func (o *ClusterRequest) SetMinRunningNodes(v int32) {
	o.MinRunningNodes = &v
}

// GetMaxRunningNodes returns the MaxRunningNodes field value if set, zero value otherwise.
func (o *ClusterRequest) GetMaxRunningNodes() int32 {
	if o == nil || o.MaxRunningNodes == nil {
		var ret int32
		return ret
	}
	return *o.MaxRunningNodes
}

// GetMaxRunningNodesOk returns a tuple with the MaxRunningNodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterRequest) GetMaxRunningNodesOk() (*int32, bool) {
	if o == nil || o.MaxRunningNodes == nil {
		return nil, false
	}
	return o.MaxRunningNodes, true
}

// HasMaxRunningNodes returns a boolean if a field has been set.
func (o *ClusterRequest) HasMaxRunningNodes() bool {
	if o != nil && o.MaxRunningNodes != nil {
		return true
	}

	return false
}

// SetMaxRunningNodes gets a reference to the given int32 and assigns it to the MaxRunningNodes field.
func (o *ClusterRequest) SetMaxRunningNodes(v int32) {
	o.MaxRunningNodes = &v
}

func (o ClusterRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.AutoUpdate != nil {
		toSerialize["auto_update"] = o.AutoUpdate
	}
	if o.EnableHistoricMetric != nil {
		toSerialize["enable_historic_metric"] = o.EnableHistoricMetric
	}
	if o.EnableHistoricLog != nil {
		toSerialize["enable_historic_log"] = o.EnableHistoricLog
	}
	if o.Cpu != nil {
		toSerialize["cpu"] = o.Cpu
	}
	if o.Memory != nil {
		toSerialize["memory"] = o.Memory
	}
	if o.MinRunningNodes != nil {
		toSerialize["min_running_nodes"] = o.MinRunningNodes
	}
	if o.MaxRunningNodes != nil {
		toSerialize["max_running_nodes"] = o.MaxRunningNodes
	}
	return json.Marshal(toSerialize)
}

type NullableClusterRequest struct {
	value *ClusterRequest
	isSet bool
}

func (v NullableClusterRequest) Get() *ClusterRequest {
	return v.value
}

func (v *NullableClusterRequest) Set(val *ClusterRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterRequest(val *ClusterRequest) *NullableClusterRequest {
	return &NullableClusterRequest{value: val, isSet: true}
}

func (v NullableClusterRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
