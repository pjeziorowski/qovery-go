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

// checks if the HelmSourceGitResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HelmSourceGitResponse{}

// HelmSourceGitResponse struct for HelmSourceGitResponse
type HelmSourceGitResponse struct {
	GitRepository        ApplicationGitRepository `json:"git_repository"`
	AdditionalProperties map[string]interface{}
}

type _HelmSourceGitResponse HelmSourceGitResponse

// NewHelmSourceGitResponse instantiates a new HelmSourceGitResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHelmSourceGitResponse(gitRepository ApplicationGitRepository) *HelmSourceGitResponse {
	this := HelmSourceGitResponse{}
	this.GitRepository = gitRepository
	return &this
}

// NewHelmSourceGitResponseWithDefaults instantiates a new HelmSourceGitResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHelmSourceGitResponseWithDefaults() *HelmSourceGitResponse {
	this := HelmSourceGitResponse{}
	return &this
}

// GetGitRepository returns the GitRepository field value
func (o *HelmSourceGitResponse) GetGitRepository() ApplicationGitRepository {
	if o == nil {
		var ret ApplicationGitRepository
		return ret
	}

	return o.GitRepository
}

// GetGitRepositoryOk returns a tuple with the GitRepository field value
// and a boolean to check if the value has been set.
func (o *HelmSourceGitResponse) GetGitRepositoryOk() (*ApplicationGitRepository, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GitRepository, true
}

// SetGitRepository sets field value
func (o *HelmSourceGitResponse) SetGitRepository(v ApplicationGitRepository) {
	o.GitRepository = v
}

func (o HelmSourceGitResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HelmSourceGitResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["git_repository"] = o.GitRepository

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *HelmSourceGitResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"git_repository",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varHelmSourceGitResponse := _HelmSourceGitResponse{}

	err = json.Unmarshal(data, &varHelmSourceGitResponse)

	if err != nil {
		return err
	}

	*o = HelmSourceGitResponse(varHelmSourceGitResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "git_repository")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHelmSourceGitResponse struct {
	value *HelmSourceGitResponse
	isSet bool
}

func (v NullableHelmSourceGitResponse) Get() *HelmSourceGitResponse {
	return v.value
}

func (v *NullableHelmSourceGitResponse) Set(val *HelmSourceGitResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableHelmSourceGitResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableHelmSourceGitResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHelmSourceGitResponse(val *HelmSourceGitResponse) *NullableHelmSourceGitResponse {
	return &NullableHelmSourceGitResponse{value: val, isSet: true}
}

func (v NullableHelmSourceGitResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHelmSourceGitResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
