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

// ApplicationEditRequest struct for ApplicationEditRequest
type ApplicationEditRequest struct {
	Storage []ServiceStorageRequestStorageInner `json:"storage,omitempty"`
	// name is case insensitive
	Name *string `json:"name,omitempty"`
	// give a description to this application
	Description   *string                          `json:"description,omitempty"`
	GitRepository *ApplicationGitRepositoryRequest `json:"git_repository,omitempty"`
	BuildMode     *BuildModeEnum                   `json:"build_mode,omitempty"`
	// The path of the associated Dockerfile
	DockerfilePath    *string                       `json:"dockerfile_path,omitempty"`
	BuildpackLanguage NullableBuildPackLanguageEnum `json:"buildpack_language,omitempty"`
	// unit is millicores (m). 1000m = 1 cpu
	Cpu *int32 `json:"cpu,omitempty"`
	// unit is MB. 1024 MB = 1GB
	Memory *int32 `json:"memory,omitempty"`
	// Minimum number of instances running. This resource auto-scale based on the CPU and Memory consumption. Note: 0 means that there is no application running.
	MinRunningInstances *int32 `json:"min_running_instances,omitempty"`
	// Maximum number of instances running. This resource auto-scale based on the CPU and Memory consumption. Note: -1 means that there is no limit.
	MaxRunningInstances *int32       `json:"max_running_instances,omitempty"`
	Healthchecks        *Healthcheck `json:"healthchecks,omitempty"`
	// Specify if the environment preview option is activated or not for this application.   If activated, a preview environment will be automatically cloned at each pull request.   If not specified, it takes the value of the `auto_preview` property from the associated environment.
	AutoPreview *bool         `json:"auto_preview,omitempty"`
	Ports       []ServicePort `json:"ports,omitempty"`
	Arguments   []string      `json:"arguments,omitempty"`
	// optional entrypoint when launching container
	Entrypoint *string `json:"entrypoint,omitempty"`
	// Specify if the application will be automatically updated after receiving a new commit.
	AutoDeploy NullableBool `json:"auto_deploy,omitempty"`
}

// NewApplicationEditRequest instantiates a new ApplicationEditRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationEditRequest() *ApplicationEditRequest {
	this := ApplicationEditRequest{}
	var buildMode BuildModeEnum = BUILDMODEENUM_BUILDPACKS
	this.BuildMode = &buildMode
	var cpu int32 = 500
	this.Cpu = &cpu
	var memory int32 = 512
	this.Memory = &memory
	var minRunningInstances int32 = 1
	this.MinRunningInstances = &minRunningInstances
	var maxRunningInstances int32 = 1
	this.MaxRunningInstances = &maxRunningInstances
	var autoPreview bool = true
	this.AutoPreview = &autoPreview
	return &this
}

// NewApplicationEditRequestWithDefaults instantiates a new ApplicationEditRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationEditRequestWithDefaults() *ApplicationEditRequest {
	this := ApplicationEditRequest{}
	var buildMode BuildModeEnum = BUILDMODEENUM_BUILDPACKS
	this.BuildMode = &buildMode
	var cpu int32 = 500
	this.Cpu = &cpu
	var memory int32 = 512
	this.Memory = &memory
	var minRunningInstances int32 = 1
	this.MinRunningInstances = &minRunningInstances
	var maxRunningInstances int32 = 1
	this.MaxRunningInstances = &maxRunningInstances
	var autoPreview bool = true
	this.AutoPreview = &autoPreview
	return &this
}

// GetStorage returns the Storage field value if set, zero value otherwise.
func (o *ApplicationEditRequest) GetStorage() []ServiceStorageRequestStorageInner {
	if o == nil || o.Storage == nil {
		var ret []ServiceStorageRequestStorageInner
		return ret
	}
	return o.Storage
}

// GetStorageOk returns a tuple with the Storage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationEditRequest) GetStorageOk() ([]ServiceStorageRequestStorageInner, bool) {
	if o == nil || o.Storage == nil {
		return nil, false
	}
	return o.Storage, true
}

// HasStorage returns a boolean if a field has been set.
func (o *ApplicationEditRequest) HasStorage() bool {
	if o != nil && o.Storage != nil {
		return true
	}

	return false
}

// SetStorage gets a reference to the given []ServiceStorageRequestStorageInner and assigns it to the Storage field.
func (o *ApplicationEditRequest) SetStorage(v []ServiceStorageRequestStorageInner) {
	o.Storage = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ApplicationEditRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationEditRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApplicationEditRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApplicationEditRequest) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ApplicationEditRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationEditRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ApplicationEditRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ApplicationEditRequest) SetDescription(v string) {
	o.Description = &v
}

// GetGitRepository returns the GitRepository field value if set, zero value otherwise.
func (o *ApplicationEditRequest) GetGitRepository() ApplicationGitRepositoryRequest {
	if o == nil || o.GitRepository == nil {
		var ret ApplicationGitRepositoryRequest
		return ret
	}
	return *o.GitRepository
}

// GetGitRepositoryOk returns a tuple with the GitRepository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationEditRequest) GetGitRepositoryOk() (*ApplicationGitRepositoryRequest, bool) {
	if o == nil || o.GitRepository == nil {
		return nil, false
	}
	return o.GitRepository, true
}

// HasGitRepository returns a boolean if a field has been set.
func (o *ApplicationEditRequest) HasGitRepository() bool {
	if o != nil && o.GitRepository != nil {
		return true
	}

	return false
}

// SetGitRepository gets a reference to the given ApplicationGitRepositoryRequest and assigns it to the GitRepository field.
func (o *ApplicationEditRequest) SetGitRepository(v ApplicationGitRepositoryRequest) {
	o.GitRepository = &v
}

// GetBuildMode returns the BuildMode field value if set, zero value otherwise.
func (o *ApplicationEditRequest) GetBuildMode() BuildModeEnum {
	if o == nil || o.BuildMode == nil {
		var ret BuildModeEnum
		return ret
	}
	return *o.BuildMode
}

// GetBuildModeOk returns a tuple with the BuildMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationEditRequest) GetBuildModeOk() (*BuildModeEnum, bool) {
	if o == nil || o.BuildMode == nil {
		return nil, false
	}
	return o.BuildMode, true
}

// HasBuildMode returns a boolean if a field has been set.
func (o *ApplicationEditRequest) HasBuildMode() bool {
	if o != nil && o.BuildMode != nil {
		return true
	}

	return false
}

// SetBuildMode gets a reference to the given BuildModeEnum and assigns it to the BuildMode field.
func (o *ApplicationEditRequest) SetBuildMode(v BuildModeEnum) {
	o.BuildMode = &v
}

// GetDockerfilePath returns the DockerfilePath field value if set, zero value otherwise.
func (o *ApplicationEditRequest) GetDockerfilePath() string {
	if o == nil || o.DockerfilePath == nil {
		var ret string
		return ret
	}
	return *o.DockerfilePath
}

// GetDockerfilePathOk returns a tuple with the DockerfilePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationEditRequest) GetDockerfilePathOk() (*string, bool) {
	if o == nil || o.DockerfilePath == nil {
		return nil, false
	}
	return o.DockerfilePath, true
}

// HasDockerfilePath returns a boolean if a field has been set.
func (o *ApplicationEditRequest) HasDockerfilePath() bool {
	if o != nil && o.DockerfilePath != nil {
		return true
	}

	return false
}

// SetDockerfilePath gets a reference to the given string and assigns it to the DockerfilePath field.
func (o *ApplicationEditRequest) SetDockerfilePath(v string) {
	o.DockerfilePath = &v
}

// GetBuildpackLanguage returns the BuildpackLanguage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationEditRequest) GetBuildpackLanguage() BuildPackLanguageEnum {
	if o == nil || o.BuildpackLanguage.Get() == nil {
		var ret BuildPackLanguageEnum
		return ret
	}
	return *o.BuildpackLanguage.Get()
}

// GetBuildpackLanguageOk returns a tuple with the BuildpackLanguage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationEditRequest) GetBuildpackLanguageOk() (*BuildPackLanguageEnum, bool) {
	if o == nil {
		return nil, false
	}
	return o.BuildpackLanguage.Get(), o.BuildpackLanguage.IsSet()
}

// HasBuildpackLanguage returns a boolean if a field has been set.
func (o *ApplicationEditRequest) HasBuildpackLanguage() bool {
	if o != nil && o.BuildpackLanguage.IsSet() {
		return true
	}

	return false
}

// SetBuildpackLanguage gets a reference to the given NullableBuildPackLanguageEnum and assigns it to the BuildpackLanguage field.
func (o *ApplicationEditRequest) SetBuildpackLanguage(v BuildPackLanguageEnum) {
	o.BuildpackLanguage.Set(&v)
}

// SetBuildpackLanguageNil sets the value for BuildpackLanguage to be an explicit nil
func (o *ApplicationEditRequest) SetBuildpackLanguageNil() {
	o.BuildpackLanguage.Set(nil)
}

// UnsetBuildpackLanguage ensures that no value is present for BuildpackLanguage, not even an explicit nil
func (o *ApplicationEditRequest) UnsetBuildpackLanguage() {
	o.BuildpackLanguage.Unset()
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *ApplicationEditRequest) GetCpu() int32 {
	if o == nil || o.Cpu == nil {
		var ret int32
		return ret
	}
	return *o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationEditRequest) GetCpuOk() (*int32, bool) {
	if o == nil || o.Cpu == nil {
		return nil, false
	}
	return o.Cpu, true
}

// HasCpu returns a boolean if a field has been set.
func (o *ApplicationEditRequest) HasCpu() bool {
	if o != nil && o.Cpu != nil {
		return true
	}

	return false
}

// SetCpu gets a reference to the given int32 and assigns it to the Cpu field.
func (o *ApplicationEditRequest) SetCpu(v int32) {
	o.Cpu = &v
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *ApplicationEditRequest) GetMemory() int32 {
	if o == nil || o.Memory == nil {
		var ret int32
		return ret
	}
	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationEditRequest) GetMemoryOk() (*int32, bool) {
	if o == nil || o.Memory == nil {
		return nil, false
	}
	return o.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *ApplicationEditRequest) HasMemory() bool {
	if o != nil && o.Memory != nil {
		return true
	}

	return false
}

// SetMemory gets a reference to the given int32 and assigns it to the Memory field.
func (o *ApplicationEditRequest) SetMemory(v int32) {
	o.Memory = &v
}

// GetMinRunningInstances returns the MinRunningInstances field value if set, zero value otherwise.
func (o *ApplicationEditRequest) GetMinRunningInstances() int32 {
	if o == nil || o.MinRunningInstances == nil {
		var ret int32
		return ret
	}
	return *o.MinRunningInstances
}

// GetMinRunningInstancesOk returns a tuple with the MinRunningInstances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationEditRequest) GetMinRunningInstancesOk() (*int32, bool) {
	if o == nil || o.MinRunningInstances == nil {
		return nil, false
	}
	return o.MinRunningInstances, true
}

// HasMinRunningInstances returns a boolean if a field has been set.
func (o *ApplicationEditRequest) HasMinRunningInstances() bool {
	if o != nil && o.MinRunningInstances != nil {
		return true
	}

	return false
}

// SetMinRunningInstances gets a reference to the given int32 and assigns it to the MinRunningInstances field.
func (o *ApplicationEditRequest) SetMinRunningInstances(v int32) {
	o.MinRunningInstances = &v
}

// GetMaxRunningInstances returns the MaxRunningInstances field value if set, zero value otherwise.
func (o *ApplicationEditRequest) GetMaxRunningInstances() int32 {
	if o == nil || o.MaxRunningInstances == nil {
		var ret int32
		return ret
	}
	return *o.MaxRunningInstances
}

// GetMaxRunningInstancesOk returns a tuple with the MaxRunningInstances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationEditRequest) GetMaxRunningInstancesOk() (*int32, bool) {
	if o == nil || o.MaxRunningInstances == nil {
		return nil, false
	}
	return o.MaxRunningInstances, true
}

// HasMaxRunningInstances returns a boolean if a field has been set.
func (o *ApplicationEditRequest) HasMaxRunningInstances() bool {
	if o != nil && o.MaxRunningInstances != nil {
		return true
	}

	return false
}

// SetMaxRunningInstances gets a reference to the given int32 and assigns it to the MaxRunningInstances field.
func (o *ApplicationEditRequest) SetMaxRunningInstances(v int32) {
	o.MaxRunningInstances = &v
}

// GetHealthchecks returns the Healthchecks field value if set, zero value otherwise.
func (o *ApplicationEditRequest) GetHealthchecks() Healthcheck {
	if o == nil || o.Healthchecks == nil {
		var ret Healthcheck
		return ret
	}
	return *o.Healthchecks
}

// GetHealthchecksOk returns a tuple with the Healthchecks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationEditRequest) GetHealthchecksOk() (*Healthcheck, bool) {
	if o == nil || o.Healthchecks == nil {
		return nil, false
	}
	return o.Healthchecks, true
}

// HasHealthchecks returns a boolean if a field has been set.
func (o *ApplicationEditRequest) HasHealthchecks() bool {
	if o != nil && o.Healthchecks != nil {
		return true
	}

	return false
}

// SetHealthchecks gets a reference to the given Healthcheck and assigns it to the Healthchecks field.
func (o *ApplicationEditRequest) SetHealthchecks(v Healthcheck) {
	o.Healthchecks = &v
}

// GetAutoPreview returns the AutoPreview field value if set, zero value otherwise.
func (o *ApplicationEditRequest) GetAutoPreview() bool {
	if o == nil || o.AutoPreview == nil {
		var ret bool
		return ret
	}
	return *o.AutoPreview
}

// GetAutoPreviewOk returns a tuple with the AutoPreview field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationEditRequest) GetAutoPreviewOk() (*bool, bool) {
	if o == nil || o.AutoPreview == nil {
		return nil, false
	}
	return o.AutoPreview, true
}

// HasAutoPreview returns a boolean if a field has been set.
func (o *ApplicationEditRequest) HasAutoPreview() bool {
	if o != nil && o.AutoPreview != nil {
		return true
	}

	return false
}

// SetAutoPreview gets a reference to the given bool and assigns it to the AutoPreview field.
func (o *ApplicationEditRequest) SetAutoPreview(v bool) {
	o.AutoPreview = &v
}

// GetPorts returns the Ports field value if set, zero value otherwise.
func (o *ApplicationEditRequest) GetPorts() []ServicePort {
	if o == nil || o.Ports == nil {
		var ret []ServicePort
		return ret
	}
	return o.Ports
}

// GetPortsOk returns a tuple with the Ports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationEditRequest) GetPortsOk() ([]ServicePort, bool) {
	if o == nil || o.Ports == nil {
		return nil, false
	}
	return o.Ports, true
}

// HasPorts returns a boolean if a field has been set.
func (o *ApplicationEditRequest) HasPorts() bool {
	if o != nil && o.Ports != nil {
		return true
	}

	return false
}

// SetPorts gets a reference to the given []ServicePort and assigns it to the Ports field.
func (o *ApplicationEditRequest) SetPorts(v []ServicePort) {
	o.Ports = v
}

// GetArguments returns the Arguments field value if set, zero value otherwise.
func (o *ApplicationEditRequest) GetArguments() []string {
	if o == nil || o.Arguments == nil {
		var ret []string
		return ret
	}
	return o.Arguments
}

// GetArgumentsOk returns a tuple with the Arguments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationEditRequest) GetArgumentsOk() ([]string, bool) {
	if o == nil || o.Arguments == nil {
		return nil, false
	}
	return o.Arguments, true
}

// HasArguments returns a boolean if a field has been set.
func (o *ApplicationEditRequest) HasArguments() bool {
	if o != nil && o.Arguments != nil {
		return true
	}

	return false
}

// SetArguments gets a reference to the given []string and assigns it to the Arguments field.
func (o *ApplicationEditRequest) SetArguments(v []string) {
	o.Arguments = v
}

// GetEntrypoint returns the Entrypoint field value if set, zero value otherwise.
func (o *ApplicationEditRequest) GetEntrypoint() string {
	if o == nil || o.Entrypoint == nil {
		var ret string
		return ret
	}
	return *o.Entrypoint
}

// GetEntrypointOk returns a tuple with the Entrypoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationEditRequest) GetEntrypointOk() (*string, bool) {
	if o == nil || o.Entrypoint == nil {
		return nil, false
	}
	return o.Entrypoint, true
}

// HasEntrypoint returns a boolean if a field has been set.
func (o *ApplicationEditRequest) HasEntrypoint() bool {
	if o != nil && o.Entrypoint != nil {
		return true
	}

	return false
}

// SetEntrypoint gets a reference to the given string and assigns it to the Entrypoint field.
func (o *ApplicationEditRequest) SetEntrypoint(v string) {
	o.Entrypoint = &v
}

// GetAutoDeploy returns the AutoDeploy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationEditRequest) GetAutoDeploy() bool {
	if o == nil || o.AutoDeploy.Get() == nil {
		var ret bool
		return ret
	}
	return *o.AutoDeploy.Get()
}

// GetAutoDeployOk returns a tuple with the AutoDeploy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationEditRequest) GetAutoDeployOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.AutoDeploy.Get(), o.AutoDeploy.IsSet()
}

// HasAutoDeploy returns a boolean if a field has been set.
func (o *ApplicationEditRequest) HasAutoDeploy() bool {
	if o != nil && o.AutoDeploy.IsSet() {
		return true
	}

	return false
}

// SetAutoDeploy gets a reference to the given NullableBool and assigns it to the AutoDeploy field.
func (o *ApplicationEditRequest) SetAutoDeploy(v bool) {
	o.AutoDeploy.Set(&v)
}

// SetAutoDeployNil sets the value for AutoDeploy to be an explicit nil
func (o *ApplicationEditRequest) SetAutoDeployNil() {
	o.AutoDeploy.Set(nil)
}

// UnsetAutoDeploy ensures that no value is present for AutoDeploy, not even an explicit nil
func (o *ApplicationEditRequest) UnsetAutoDeploy() {
	o.AutoDeploy.Unset()
}

func (o ApplicationEditRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Storage != nil {
		toSerialize["storage"] = o.Storage
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.GitRepository != nil {
		toSerialize["git_repository"] = o.GitRepository
	}
	if o.BuildMode != nil {
		toSerialize["build_mode"] = o.BuildMode
	}
	if o.DockerfilePath != nil {
		toSerialize["dockerfile_path"] = o.DockerfilePath
	}
	if o.BuildpackLanguage.IsSet() {
		toSerialize["buildpack_language"] = o.BuildpackLanguage.Get()
	}
	if o.Cpu != nil {
		toSerialize["cpu"] = o.Cpu
	}
	if o.Memory != nil {
		toSerialize["memory"] = o.Memory
	}
	if o.MinRunningInstances != nil {
		toSerialize["min_running_instances"] = o.MinRunningInstances
	}
	if o.MaxRunningInstances != nil {
		toSerialize["max_running_instances"] = o.MaxRunningInstances
	}
	if o.Healthchecks != nil {
		toSerialize["healthchecks"] = o.Healthchecks
	}
	if o.AutoPreview != nil {
		toSerialize["auto_preview"] = o.AutoPreview
	}
	if o.Ports != nil {
		toSerialize["ports"] = o.Ports
	}
	if o.Arguments != nil {
		toSerialize["arguments"] = o.Arguments
	}
	if o.Entrypoint != nil {
		toSerialize["entrypoint"] = o.Entrypoint
	}
	if o.AutoDeploy.IsSet() {
		toSerialize["auto_deploy"] = o.AutoDeploy.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableApplicationEditRequest struct {
	value *ApplicationEditRequest
	isSet bool
}

func (v NullableApplicationEditRequest) Get() *ApplicationEditRequest {
	return v.value
}

func (v *NullableApplicationEditRequest) Set(val *ApplicationEditRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationEditRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationEditRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationEditRequest(val *ApplicationEditRequest) *NullableApplicationEditRequest {
	return &NullableApplicationEditRequest{value: val, isSet: true}
}

func (v NullableApplicationEditRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationEditRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
