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

// ContainerRegistryRequest struct for ContainerRegistryRequest
type ContainerRegistryRequest struct {
	Name        string                    `json:"name"`
	Kind        ContainerRegistryKindEnum `json:"kind"`
	Description *string                   `json:"description,omitempty"`
	// URL of the container registry: * For `DOCKER_HUB`: it must be `https://docker.io` (default with 'https://docker.io' if no url provided for DOCKER_HUB) * For others: it's required and must start by `https://`
	Url    *string                        `json:"url,omitempty"`
	Config ContainerRegistryRequestConfig `json:"config"`
}

// NewContainerRegistryRequest instantiates a new ContainerRegistryRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerRegistryRequest(name string, kind ContainerRegistryKindEnum, config ContainerRegistryRequestConfig) *ContainerRegistryRequest {
	this := ContainerRegistryRequest{}
	this.Name = name
	this.Kind = kind
	this.Config = config
	return &this
}

// NewContainerRegistryRequestWithDefaults instantiates a new ContainerRegistryRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerRegistryRequestWithDefaults() *ContainerRegistryRequest {
	this := ContainerRegistryRequest{}
	return &this
}

// GetName returns the Name field value
func (o *ContainerRegistryRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ContainerRegistryRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ContainerRegistryRequest) SetName(v string) {
	o.Name = v
}

// GetKind returns the Kind field value
func (o *ContainerRegistryRequest) GetKind() ContainerRegistryKindEnum {
	if o == nil {
		var ret ContainerRegistryKindEnum
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *ContainerRegistryRequest) GetKindOk() (*ContainerRegistryKindEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *ContainerRegistryRequest) SetKind(v ContainerRegistryKindEnum) {
	o.Kind = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ContainerRegistryRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerRegistryRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ContainerRegistryRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ContainerRegistryRequest) SetDescription(v string) {
	o.Description = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *ContainerRegistryRequest) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerRegistryRequest) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *ContainerRegistryRequest) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *ContainerRegistryRequest) SetUrl(v string) {
	o.Url = &v
}

// GetConfig returns the Config field value
func (o *ContainerRegistryRequest) GetConfig() ContainerRegistryRequestConfig {
	if o == nil {
		var ret ContainerRegistryRequestConfig
		return ret
	}

	return o.Config
}

// GetConfigOk returns a tuple with the Config field value
// and a boolean to check if the value has been set.
func (o *ContainerRegistryRequest) GetConfigOk() (*ContainerRegistryRequestConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Config, true
}

// SetConfig sets field value
func (o *ContainerRegistryRequest) SetConfig(v ContainerRegistryRequestConfig) {
	o.Config = v
}

func (o ContainerRegistryRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["kind"] = o.Kind
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if true {
		toSerialize["config"] = o.Config
	}
	return json.Marshal(toSerialize)
}

type NullableContainerRegistryRequest struct {
	value *ContainerRegistryRequest
	isSet bool
}

func (v NullableContainerRegistryRequest) Get() *ContainerRegistryRequest {
	return v.value
}

func (v *NullableContainerRegistryRequest) Set(val *ContainerRegistryRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerRegistryRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerRegistryRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerRegistryRequest(val *ContainerRegistryRequest) *NullableContainerRegistryRequest {
	return &NullableContainerRegistryRequest{value: val, isSet: true}
}

func (v NullableContainerRegistryRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerRegistryRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
