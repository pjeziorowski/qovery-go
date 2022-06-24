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

// OrganizationWebhookCreateRequest struct for OrganizationWebhookCreateRequest
type OrganizationWebhookCreateRequest struct {
	Kind OrganizationWebhookKindEnum `json:"kind"`
	// Set the public HTTP or HTTPS endpoint that will receive the specified events. The target URL must starts with `http://` or `https://`
	TargetUrl string `json:"target_url"`
	// Make sure you receive a payload to sign the Qovery request with your secret. Qovery will add a HTTP header `Qovery-Signature: <Your Secret>` to every webhook requests sent to your target URL.
	TargetSecret *string `json:"target_secret,omitempty"`
	Description  *string `json:"description,omitempty"`
	// Turn on or off your endpoint.
	Enabled *bool                          `json:"enabled,omitempty"`
	Events  []OrganizationWebhookEventEnum `json:"events"`
	// Specify the project names you want to filter to.  This webhook will be triggered only if the event is coming from the specified Project IDs. Notes: 1. Wildcard is accepted E.g. `product*`. 2. Name is case insensitive.
	ProjectNamesFilter []string `json:"project_names_filter,omitempty"`
	// Specify the environment modes you want to filter to. This webhook will be triggered only if the event is coming from an environment with the specified mode.
	EnvironmentTypesFilter []EnvironmentModeEnum `json:"environment_types_filter,omitempty"`
}

// NewOrganizationWebhookCreateRequest instantiates a new OrganizationWebhookCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationWebhookCreateRequest(kind OrganizationWebhookKindEnum, targetUrl string, events []OrganizationWebhookEventEnum) *OrganizationWebhookCreateRequest {
	this := OrganizationWebhookCreateRequest{}
	this.Kind = kind
	this.TargetUrl = targetUrl
	this.Events = events
	return &this
}

// NewOrganizationWebhookCreateRequestWithDefaults instantiates a new OrganizationWebhookCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationWebhookCreateRequestWithDefaults() *OrganizationWebhookCreateRequest {
	this := OrganizationWebhookCreateRequest{}
	return &this
}

// GetKind returns the Kind field value
func (o *OrganizationWebhookCreateRequest) GetKind() OrganizationWebhookKindEnum {
	if o == nil {
		var ret OrganizationWebhookKindEnum
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *OrganizationWebhookCreateRequest) GetKindOk() (*OrganizationWebhookKindEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *OrganizationWebhookCreateRequest) SetKind(v OrganizationWebhookKindEnum) {
	o.Kind = v
}

// GetTargetUrl returns the TargetUrl field value
func (o *OrganizationWebhookCreateRequest) GetTargetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetUrl
}

// GetTargetUrlOk returns a tuple with the TargetUrl field value
// and a boolean to check if the value has been set.
func (o *OrganizationWebhookCreateRequest) GetTargetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetUrl, true
}

// SetTargetUrl sets field value
func (o *OrganizationWebhookCreateRequest) SetTargetUrl(v string) {
	o.TargetUrl = v
}

// GetTargetSecret returns the TargetSecret field value if set, zero value otherwise.
func (o *OrganizationWebhookCreateRequest) GetTargetSecret() string {
	if o == nil || o.TargetSecret == nil {
		var ret string
		return ret
	}
	return *o.TargetSecret
}

// GetTargetSecretOk returns a tuple with the TargetSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationWebhookCreateRequest) GetTargetSecretOk() (*string, bool) {
	if o == nil || o.TargetSecret == nil {
		return nil, false
	}
	return o.TargetSecret, true
}

// HasTargetSecret returns a boolean if a field has been set.
func (o *OrganizationWebhookCreateRequest) HasTargetSecret() bool {
	if o != nil && o.TargetSecret != nil {
		return true
	}

	return false
}

// SetTargetSecret gets a reference to the given string and assigns it to the TargetSecret field.
func (o *OrganizationWebhookCreateRequest) SetTargetSecret(v string) {
	o.TargetSecret = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *OrganizationWebhookCreateRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationWebhookCreateRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *OrganizationWebhookCreateRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *OrganizationWebhookCreateRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *OrganizationWebhookCreateRequest) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationWebhookCreateRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *OrganizationWebhookCreateRequest) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *OrganizationWebhookCreateRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetEvents returns the Events field value
func (o *OrganizationWebhookCreateRequest) GetEvents() []OrganizationWebhookEventEnum {
	if o == nil {
		var ret []OrganizationWebhookEventEnum
		return ret
	}

	return o.Events
}

// GetEventsOk returns a tuple with the Events field value
// and a boolean to check if the value has been set.
func (o *OrganizationWebhookCreateRequest) GetEventsOk() ([]OrganizationWebhookEventEnum, bool) {
	if o == nil {
		return nil, false
	}
	return o.Events, true
}

// SetEvents sets field value
func (o *OrganizationWebhookCreateRequest) SetEvents(v []OrganizationWebhookEventEnum) {
	o.Events = v
}

// GetProjectNamesFilter returns the ProjectNamesFilter field value if set, zero value otherwise.
func (o *OrganizationWebhookCreateRequest) GetProjectNamesFilter() []string {
	if o == nil || o.ProjectNamesFilter == nil {
		var ret []string
		return ret
	}
	return o.ProjectNamesFilter
}

// GetProjectNamesFilterOk returns a tuple with the ProjectNamesFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationWebhookCreateRequest) GetProjectNamesFilterOk() ([]string, bool) {
	if o == nil || o.ProjectNamesFilter == nil {
		return nil, false
	}
	return o.ProjectNamesFilter, true
}

// HasProjectNamesFilter returns a boolean if a field has been set.
func (o *OrganizationWebhookCreateRequest) HasProjectNamesFilter() bool {
	if o != nil && o.ProjectNamesFilter != nil {
		return true
	}

	return false
}

// SetProjectNamesFilter gets a reference to the given []string and assigns it to the ProjectNamesFilter field.
func (o *OrganizationWebhookCreateRequest) SetProjectNamesFilter(v []string) {
	o.ProjectNamesFilter = v
}

// GetEnvironmentTypesFilter returns the EnvironmentTypesFilter field value if set, zero value otherwise.
func (o *OrganizationWebhookCreateRequest) GetEnvironmentTypesFilter() []EnvironmentModeEnum {
	if o == nil || o.EnvironmentTypesFilter == nil {
		var ret []EnvironmentModeEnum
		return ret
	}
	return o.EnvironmentTypesFilter
}

// GetEnvironmentTypesFilterOk returns a tuple with the EnvironmentTypesFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationWebhookCreateRequest) GetEnvironmentTypesFilterOk() ([]EnvironmentModeEnum, bool) {
	if o == nil || o.EnvironmentTypesFilter == nil {
		return nil, false
	}
	return o.EnvironmentTypesFilter, true
}

// HasEnvironmentTypesFilter returns a boolean if a field has been set.
func (o *OrganizationWebhookCreateRequest) HasEnvironmentTypesFilter() bool {
	if o != nil && o.EnvironmentTypesFilter != nil {
		return true
	}

	return false
}

// SetEnvironmentTypesFilter gets a reference to the given []EnvironmentModeEnum and assigns it to the EnvironmentTypesFilter field.
func (o *OrganizationWebhookCreateRequest) SetEnvironmentTypesFilter(v []EnvironmentModeEnum) {
	o.EnvironmentTypesFilter = v
}

func (o OrganizationWebhookCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["kind"] = o.Kind
	}
	if true {
		toSerialize["target_url"] = o.TargetUrl
	}
	if o.TargetSecret != nil {
		toSerialize["target_secret"] = o.TargetSecret
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if true {
		toSerialize["events"] = o.Events
	}
	if o.ProjectNamesFilter != nil {
		toSerialize["project_names_filter"] = o.ProjectNamesFilter
	}
	if o.EnvironmentTypesFilter != nil {
		toSerialize["environment_types_filter"] = o.EnvironmentTypesFilter
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationWebhookCreateRequest struct {
	value *OrganizationWebhookCreateRequest
	isSet bool
}

func (v NullableOrganizationWebhookCreateRequest) Get() *OrganizationWebhookCreateRequest {
	return v.value
}

func (v *NullableOrganizationWebhookCreateRequest) Set(val *OrganizationWebhookCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationWebhookCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationWebhookCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationWebhookCreateRequest(val *OrganizationWebhookCreateRequest) *NullableOrganizationWebhookCreateRequest {
	return &NullableOrganizationWebhookCreateRequest{value: val, isSet: true}
}

func (v NullableOrganizationWebhookCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationWebhookCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
