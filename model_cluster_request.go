/*
[BETA] Qovery API

- Qovery is the fastest way to deploy your full-stack apps on any Cloud provider. - ℹ️ The API is in Beta and still in progress. Some endpoints are not available yet. 

API version: 1.0.0
Contact: support+api+documentation@qovery.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package qovery

import (
	"encoding/json"
)

// ClusterRequest struct for ClusterRequest
type ClusterRequest struct {
	// name is case-insensitive
	Name string `json:"name"`
	Description *string `json:"description,omitempty"`
	CloudProvider string `json:"cloud_provider"`
	Region string `json:"region"`
	AutoUpdate *bool `json:"auto_update,omitempty"`
	// unit is millicores (m). 1000m = 1 cpu
	Cpu *int32 `json:"cpu,omitempty"`
	// unit is MB. 1024 MB = 1GB
	Memory *int32 `json:"memory,omitempty"`
	MinRunningNodes *int32 `json:"min_running_nodes,omitempty"`
	MaxRunningNodes *int32 `json:"max_running_nodes,omitempty"`
	Features *[]ClusterFeatureRequestFeatures `json:"features,omitempty"`
}

// NewClusterRequest instantiates a new ClusterRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterRequest(name string, cloudProvider string, region string) *ClusterRequest {
	this := ClusterRequest{}
	this.Name = name
	this.CloudProvider = cloudProvider
	this.Region = region
	var cpu int32 = 250
	this.Cpu = &cpu
	var memory int32 = 256
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
	var cpu int32 = 250
	this.Cpu = &cpu
	var memory int32 = 256
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
	if o == nil  {
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

// GetCloudProvider returns the CloudProvider field value
func (o *ClusterRequest) GetCloudProvider() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value
// and a boolean to check if the value has been set.
func (o *ClusterRequest) GetCloudProviderOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CloudProvider, true
}

// SetCloudProvider sets field value
func (o *ClusterRequest) SetCloudProvider(v string) {
	o.CloudProvider = v
}

// GetRegion returns the Region field value
func (o *ClusterRequest) GetRegion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *ClusterRequest) GetRegionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value
func (o *ClusterRequest) SetRegion(v string) {
	o.Region = v
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

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *ClusterRequest) GetCpu() int32 {
	if o == nil || o.Cpu == nil {
		var ret int32
		return ret
	}
	return *o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterRequest) GetCpuOk() (*int32, bool) {
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

// SetCpu gets a reference to the given int32 and assigns it to the Cpu field.
func (o *ClusterRequest) SetCpu(v int32) {
	o.Cpu = &v
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *ClusterRequest) GetMemory() int32 {
	if o == nil || o.Memory == nil {
		var ret int32
		return ret
	}
	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterRequest) GetMemoryOk() (*int32, bool) {
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

// SetMemory gets a reference to the given int32 and assigns it to the Memory field.
func (o *ClusterRequest) SetMemory(v int32) {
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

// GetFeatures returns the Features field value if set, zero value otherwise.
func (o *ClusterRequest) GetFeatures() []ClusterFeatureRequestFeatures {
	if o == nil || o.Features == nil {
		var ret []ClusterFeatureRequestFeatures
		return ret
	}
	return *o.Features
}

// GetFeaturesOk returns a tuple with the Features field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterRequest) GetFeaturesOk() (*[]ClusterFeatureRequestFeatures, bool) {
	if o == nil || o.Features == nil {
		return nil, false
	}
	return o.Features, true
}

// HasFeatures returns a boolean if a field has been set.
func (o *ClusterRequest) HasFeatures() bool {
	if o != nil && o.Features != nil {
		return true
	}

	return false
}

// SetFeatures gets a reference to the given []ClusterFeatureRequestFeatures and assigns it to the Features field.
func (o *ClusterRequest) SetFeatures(v []ClusterFeatureRequestFeatures) {
	o.Features = &v
}

func (o ClusterRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["cloud_provider"] = o.CloudProvider
	}
	if true {
		toSerialize["region"] = o.Region
	}
	if o.AutoUpdate != nil {
		toSerialize["auto_update"] = o.AutoUpdate
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
	if o.Features != nil {
		toSerialize["features"] = o.Features
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


