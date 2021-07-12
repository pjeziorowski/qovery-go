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
	"time"
)

// ApplicationGitRepositoryResponse struct for ApplicationGitRepositoryResponse
type ApplicationGitRepositoryResponse struct {
	HasAccess *bool   `json:"has_access,omitempty"`
	Provider  *string `json:"provider,omitempty"`
	Owner     *string `json:"owner,omitempty"`
	// repository name
	Name *string `json:"name,omitempty"`
	// Git commit ID corresponding to the deployed version of the app
	DeployedCommitId *string `json:"deployed_commit_id,omitempty"`
	// Git commit date corresponding to the deployed version of the app
	DeployedCommitDate *time.Time `json:"deployed_commit_date,omitempty"`
	// Git commit user corresponding to the deployed version of the app
	DeployedCommitContributor *string `json:"deployed_commit_contributor,omitempty"`
	DeployedCommitTag         *string `json:"deployed_commit_tag,omitempty"`
}

// NewApplicationGitRepositoryResponse instantiates a new ApplicationGitRepositoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationGitRepositoryResponse() *ApplicationGitRepositoryResponse {
	this := ApplicationGitRepositoryResponse{}
	return &this
}

// NewApplicationGitRepositoryResponseWithDefaults instantiates a new ApplicationGitRepositoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationGitRepositoryResponseWithDefaults() *ApplicationGitRepositoryResponse {
	this := ApplicationGitRepositoryResponse{}
	return &this
}

// GetHasAccess returns the HasAccess field value if set, zero value otherwise.
func (o *ApplicationGitRepositoryResponse) GetHasAccess() bool {
	if o == nil || o.HasAccess == nil {
		var ret bool
		return ret
	}
	return *o.HasAccess
}

// GetHasAccessOk returns a tuple with the HasAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationGitRepositoryResponse) GetHasAccessOk() (*bool, bool) {
	if o == nil || o.HasAccess == nil {
		return nil, false
	}
	return o.HasAccess, true
}

// HasHasAccess returns a boolean if a field has been set.
func (o *ApplicationGitRepositoryResponse) HasHasAccess() bool {
	if o != nil && o.HasAccess != nil {
		return true
	}

	return false
}

// SetHasAccess gets a reference to the given bool and assigns it to the HasAccess field.
func (o *ApplicationGitRepositoryResponse) SetHasAccess(v bool) {
	o.HasAccess = &v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *ApplicationGitRepositoryResponse) GetProvider() string {
	if o == nil || o.Provider == nil {
		var ret string
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationGitRepositoryResponse) GetProviderOk() (*string, bool) {
	if o == nil || o.Provider == nil {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *ApplicationGitRepositoryResponse) HasProvider() bool {
	if o != nil && o.Provider != nil {
		return true
	}

	return false
}

// SetProvider gets a reference to the given string and assigns it to the Provider field.
func (o *ApplicationGitRepositoryResponse) SetProvider(v string) {
	o.Provider = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *ApplicationGitRepositoryResponse) GetOwner() string {
	if o == nil || o.Owner == nil {
		var ret string
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationGitRepositoryResponse) GetOwnerOk() (*string, bool) {
	if o == nil || o.Owner == nil {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *ApplicationGitRepositoryResponse) HasOwner() bool {
	if o != nil && o.Owner != nil {
		return true
	}

	return false
}

// SetOwner gets a reference to the given string and assigns it to the Owner field.
func (o *ApplicationGitRepositoryResponse) SetOwner(v string) {
	o.Owner = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ApplicationGitRepositoryResponse) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationGitRepositoryResponse) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApplicationGitRepositoryResponse) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApplicationGitRepositoryResponse) SetName(v string) {
	o.Name = &v
}

// GetDeployedCommitId returns the DeployedCommitId field value if set, zero value otherwise.
func (o *ApplicationGitRepositoryResponse) GetDeployedCommitId() string {
	if o == nil || o.DeployedCommitId == nil {
		var ret string
		return ret
	}
	return *o.DeployedCommitId
}

// GetDeployedCommitIdOk returns a tuple with the DeployedCommitId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationGitRepositoryResponse) GetDeployedCommitIdOk() (*string, bool) {
	if o == nil || o.DeployedCommitId == nil {
		return nil, false
	}
	return o.DeployedCommitId, true
}

// HasDeployedCommitId returns a boolean if a field has been set.
func (o *ApplicationGitRepositoryResponse) HasDeployedCommitId() bool {
	if o != nil && o.DeployedCommitId != nil {
		return true
	}

	return false
}

// SetDeployedCommitId gets a reference to the given string and assigns it to the DeployedCommitId field.
func (o *ApplicationGitRepositoryResponse) SetDeployedCommitId(v string) {
	o.DeployedCommitId = &v
}

// GetDeployedCommitDate returns the DeployedCommitDate field value if set, zero value otherwise.
func (o *ApplicationGitRepositoryResponse) GetDeployedCommitDate() time.Time {
	if o == nil || o.DeployedCommitDate == nil {
		var ret time.Time
		return ret
	}
	return *o.DeployedCommitDate
}

// GetDeployedCommitDateOk returns a tuple with the DeployedCommitDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationGitRepositoryResponse) GetDeployedCommitDateOk() (*time.Time, bool) {
	if o == nil || o.DeployedCommitDate == nil {
		return nil, false
	}
	return o.DeployedCommitDate, true
}

// HasDeployedCommitDate returns a boolean if a field has been set.
func (o *ApplicationGitRepositoryResponse) HasDeployedCommitDate() bool {
	if o != nil && o.DeployedCommitDate != nil {
		return true
	}

	return false
}

// SetDeployedCommitDate gets a reference to the given time.Time and assigns it to the DeployedCommitDate field.
func (o *ApplicationGitRepositoryResponse) SetDeployedCommitDate(v time.Time) {
	o.DeployedCommitDate = &v
}

// GetDeployedCommitContributor returns the DeployedCommitContributor field value if set, zero value otherwise.
func (o *ApplicationGitRepositoryResponse) GetDeployedCommitContributor() string {
	if o == nil || o.DeployedCommitContributor == nil {
		var ret string
		return ret
	}
	return *o.DeployedCommitContributor
}

// GetDeployedCommitContributorOk returns a tuple with the DeployedCommitContributor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationGitRepositoryResponse) GetDeployedCommitContributorOk() (*string, bool) {
	if o == nil || o.DeployedCommitContributor == nil {
		return nil, false
	}
	return o.DeployedCommitContributor, true
}

// HasDeployedCommitContributor returns a boolean if a field has been set.
func (o *ApplicationGitRepositoryResponse) HasDeployedCommitContributor() bool {
	if o != nil && o.DeployedCommitContributor != nil {
		return true
	}

	return false
}

// SetDeployedCommitContributor gets a reference to the given string and assigns it to the DeployedCommitContributor field.
func (o *ApplicationGitRepositoryResponse) SetDeployedCommitContributor(v string) {
	o.DeployedCommitContributor = &v
}

// GetDeployedCommitTag returns the DeployedCommitTag field value if set, zero value otherwise.
func (o *ApplicationGitRepositoryResponse) GetDeployedCommitTag() string {
	if o == nil || o.DeployedCommitTag == nil {
		var ret string
		return ret
	}
	return *o.DeployedCommitTag
}

// GetDeployedCommitTagOk returns a tuple with the DeployedCommitTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationGitRepositoryResponse) GetDeployedCommitTagOk() (*string, bool) {
	if o == nil || o.DeployedCommitTag == nil {
		return nil, false
	}
	return o.DeployedCommitTag, true
}

// HasDeployedCommitTag returns a boolean if a field has been set.
func (o *ApplicationGitRepositoryResponse) HasDeployedCommitTag() bool {
	if o != nil && o.DeployedCommitTag != nil {
		return true
	}

	return false
}

// SetDeployedCommitTag gets a reference to the given string and assigns it to the DeployedCommitTag field.
func (o *ApplicationGitRepositoryResponse) SetDeployedCommitTag(v string) {
	o.DeployedCommitTag = &v
}

func (o ApplicationGitRepositoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.HasAccess != nil {
		toSerialize["has_access"] = o.HasAccess
	}
	if o.Provider != nil {
		toSerialize["provider"] = o.Provider
	}
	if o.Owner != nil {
		toSerialize["owner"] = o.Owner
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.DeployedCommitId != nil {
		toSerialize["deployed_commit_id"] = o.DeployedCommitId
	}
	if o.DeployedCommitDate != nil {
		toSerialize["deployed_commit_date"] = o.DeployedCommitDate
	}
	if o.DeployedCommitContributor != nil {
		toSerialize["deployed_commit_contributor"] = o.DeployedCommitContributor
	}
	if o.DeployedCommitTag != nil {
		toSerialize["deployed_commit_tag"] = o.DeployedCommitTag
	}
	return json.Marshal(toSerialize)
}

type NullableApplicationGitRepositoryResponse struct {
	value *ApplicationGitRepositoryResponse
	isSet bool
}

func (v NullableApplicationGitRepositoryResponse) Get() *ApplicationGitRepositoryResponse {
	return v.value
}

func (v *NullableApplicationGitRepositoryResponse) Set(val *ApplicationGitRepositoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationGitRepositoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationGitRepositoryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationGitRepositoryResponse(val *ApplicationGitRepositoryResponse) *NullableApplicationGitRepositoryResponse {
	return &NullableApplicationGitRepositoryResponse{value: val, isSet: true}
}

func (v NullableApplicationGitRepositoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationGitRepositoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
