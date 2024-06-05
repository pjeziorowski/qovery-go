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

// checks if the JobSourceDockerResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JobSourceDockerResponse{}

// JobSourceDockerResponse struct for JobSourceDockerResponse
type JobSourceDockerResponse struct {
	GitRepository *ApplicationGitRepository `json:"git_repository,omitempty"`
	// The path of the associated Dockerfile. Only if you are using build_mode = DOCKER
	DockerfilePath NullableString `json:"dockerfile_path,omitempty"`
	// The content of your dockerfile if it is not stored inside your git repository
	DockerfileRaw NullableString `json:"dockerfile_raw,omitempty"`
}

// NewJobSourceDockerResponse instantiates a new JobSourceDockerResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJobSourceDockerResponse() *JobSourceDockerResponse {
	this := JobSourceDockerResponse{}
	return &this
}

// NewJobSourceDockerResponseWithDefaults instantiates a new JobSourceDockerResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJobSourceDockerResponseWithDefaults() *JobSourceDockerResponse {
	this := JobSourceDockerResponse{}
	return &this
}

// GetGitRepository returns the GitRepository field value if set, zero value otherwise.
func (o *JobSourceDockerResponse) GetGitRepository() ApplicationGitRepository {
	if o == nil || IsNil(o.GitRepository) {
		var ret ApplicationGitRepository
		return ret
	}
	return *o.GitRepository
}

// GetGitRepositoryOk returns a tuple with the GitRepository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobSourceDockerResponse) GetGitRepositoryOk() (*ApplicationGitRepository, bool) {
	if o == nil || IsNil(o.GitRepository) {
		return nil, false
	}
	return o.GitRepository, true
}

// HasGitRepository returns a boolean if a field has been set.
func (o *JobSourceDockerResponse) HasGitRepository() bool {
	if o != nil && !IsNil(o.GitRepository) {
		return true
	}

	return false
}

// SetGitRepository gets a reference to the given ApplicationGitRepository and assigns it to the GitRepository field.
func (o *JobSourceDockerResponse) SetGitRepository(v ApplicationGitRepository) {
	o.GitRepository = &v
}

// GetDockerfilePath returns the DockerfilePath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JobSourceDockerResponse) GetDockerfilePath() string {
	if o == nil || IsNil(o.DockerfilePath.Get()) {
		var ret string
		return ret
	}
	return *o.DockerfilePath.Get()
}

// GetDockerfilePathOk returns a tuple with the DockerfilePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JobSourceDockerResponse) GetDockerfilePathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DockerfilePath.Get(), o.DockerfilePath.IsSet()
}

// HasDockerfilePath returns a boolean if a field has been set.
func (o *JobSourceDockerResponse) HasDockerfilePath() bool {
	if o != nil && o.DockerfilePath.IsSet() {
		return true
	}

	return false
}

// SetDockerfilePath gets a reference to the given NullableString and assigns it to the DockerfilePath field.
func (o *JobSourceDockerResponse) SetDockerfilePath(v string) {
	o.DockerfilePath.Set(&v)
}

// SetDockerfilePathNil sets the value for DockerfilePath to be an explicit nil
func (o *JobSourceDockerResponse) SetDockerfilePathNil() {
	o.DockerfilePath.Set(nil)
}

// UnsetDockerfilePath ensures that no value is present for DockerfilePath, not even an explicit nil
func (o *JobSourceDockerResponse) UnsetDockerfilePath() {
	o.DockerfilePath.Unset()
}

// GetDockerfileRaw returns the DockerfileRaw field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JobSourceDockerResponse) GetDockerfileRaw() string {
	if o == nil || IsNil(o.DockerfileRaw.Get()) {
		var ret string
		return ret
	}
	return *o.DockerfileRaw.Get()
}

// GetDockerfileRawOk returns a tuple with the DockerfileRaw field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JobSourceDockerResponse) GetDockerfileRawOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DockerfileRaw.Get(), o.DockerfileRaw.IsSet()
}

// HasDockerfileRaw returns a boolean if a field has been set.
func (o *JobSourceDockerResponse) HasDockerfileRaw() bool {
	if o != nil && o.DockerfileRaw.IsSet() {
		return true
	}

	return false
}

// SetDockerfileRaw gets a reference to the given NullableString and assigns it to the DockerfileRaw field.
func (o *JobSourceDockerResponse) SetDockerfileRaw(v string) {
	o.DockerfileRaw.Set(&v)
}

// SetDockerfileRawNil sets the value for DockerfileRaw to be an explicit nil
func (o *JobSourceDockerResponse) SetDockerfileRawNil() {
	o.DockerfileRaw.Set(nil)
}

// UnsetDockerfileRaw ensures that no value is present for DockerfileRaw, not even an explicit nil
func (o *JobSourceDockerResponse) UnsetDockerfileRaw() {
	o.DockerfileRaw.Unset()
}

func (o JobSourceDockerResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JobSourceDockerResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GitRepository) {
		toSerialize["git_repository"] = o.GitRepository
	}
	if o.DockerfilePath.IsSet() {
		toSerialize["dockerfile_path"] = o.DockerfilePath.Get()
	}
	if o.DockerfileRaw.IsSet() {
		toSerialize["dockerfile_raw"] = o.DockerfileRaw.Get()
	}
	return toSerialize, nil
}

type NullableJobSourceDockerResponse struct {
	value *JobSourceDockerResponse
	isSet bool
}

func (v NullableJobSourceDockerResponse) Get() *JobSourceDockerResponse {
	return v.value
}

func (v *NullableJobSourceDockerResponse) Set(val *JobSourceDockerResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableJobSourceDockerResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableJobSourceDockerResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobSourceDockerResponse(val *JobSourceDockerResponse) *NullableJobSourceDockerResponse {
	return &NullableJobSourceDockerResponse{value: val, isSet: true}
}

func (v NullableJobSourceDockerResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobSourceDockerResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
