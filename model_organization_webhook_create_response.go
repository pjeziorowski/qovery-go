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

// OrganizationWebhookCreateResponse struct for OrganizationWebhookCreateResponse
type OrganizationWebhookCreateResponse struct {
	Id        string     `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	Kind      *Kind      `json:"kind,omitempty"`
	// Set the public HTTP or HTTPS endpoint that will receive the specified events. The target URL must starts with `http://` or `https://`
	TargetUrl       *string `json:"target_url,omitempty"`
	TargetSecretSet *bool   `json:"target_secret_set,omitempty"`
	Description     *string `json:"description,omitempty"`
	// Turn on or off your endpoint.
	Enabled                *bool                 `json:"enabled,omitempty"`
	Events                 []Items               `json:"events,omitempty"`
	ProjectIdFilter        []string              `json:"project_id_filter,omitempty"`
	EnvironmentTypesFilter []EnvironmentModeEnum `json:"environment_types_filter,omitempty"`
}

// NewOrganizationWebhookCreateResponse instantiates a new OrganizationWebhookCreateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationWebhookCreateResponse(id string, createdAt time.Time) *OrganizationWebhookCreateResponse {
	this := OrganizationWebhookCreateResponse{}
	this.Id = id
	this.CreatedAt = createdAt
	return &this
}

// NewOrganizationWebhookCreateResponseWithDefaults instantiates a new OrganizationWebhookCreateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationWebhookCreateResponseWithDefaults() *OrganizationWebhookCreateResponse {
	this := OrganizationWebhookCreateResponse{}
	return &this
}

// GetId returns the Id field value
func (o *OrganizationWebhookCreateResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *OrganizationWebhookCreateResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *OrganizationWebhookCreateResponse) SetId(v string) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *OrganizationWebhookCreateResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *OrganizationWebhookCreateResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *OrganizationWebhookCreateResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *OrganizationWebhookCreateResponse) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationWebhookCreateResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *OrganizationWebhookCreateResponse) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *OrganizationWebhookCreateResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *OrganizationWebhookCreateResponse) GetKind() Kind {
	if o == nil || o.Kind == nil {
		var ret Kind
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationWebhookCreateResponse) GetKindOk() (*Kind, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *OrganizationWebhookCreateResponse) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given Kind and assigns it to the Kind field.
func (o *OrganizationWebhookCreateResponse) SetKind(v Kind) {
	o.Kind = &v
}

// GetTargetUrl returns the TargetUrl field value if set, zero value otherwise.
func (o *OrganizationWebhookCreateResponse) GetTargetUrl() string {
	if o == nil || o.TargetUrl == nil {
		var ret string
		return ret
	}
	return *o.TargetUrl
}

// GetTargetUrlOk returns a tuple with the TargetUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationWebhookCreateResponse) GetTargetUrlOk() (*string, bool) {
	if o == nil || o.TargetUrl == nil {
		return nil, false
	}
	return o.TargetUrl, true
}

// HasTargetUrl returns a boolean if a field has been set.
func (o *OrganizationWebhookCreateResponse) HasTargetUrl() bool {
	if o != nil && o.TargetUrl != nil {
		return true
	}

	return false
}

// SetTargetUrl gets a reference to the given string and assigns it to the TargetUrl field.
func (o *OrganizationWebhookCreateResponse) SetTargetUrl(v string) {
	o.TargetUrl = &v
}

// GetTargetSecretSet returns the TargetSecretSet field value if set, zero value otherwise.
func (o *OrganizationWebhookCreateResponse) GetTargetSecretSet() bool {
	if o == nil || o.TargetSecretSet == nil {
		var ret bool
		return ret
	}
	return *o.TargetSecretSet
}

// GetTargetSecretSetOk returns a tuple with the TargetSecretSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationWebhookCreateResponse) GetTargetSecretSetOk() (*bool, bool) {
	if o == nil || o.TargetSecretSet == nil {
		return nil, false
	}
	return o.TargetSecretSet, true
}

// HasTargetSecretSet returns a boolean if a field has been set.
func (o *OrganizationWebhookCreateResponse) HasTargetSecretSet() bool {
	if o != nil && o.TargetSecretSet != nil {
		return true
	}

	return false
}

// SetTargetSecretSet gets a reference to the given bool and assigns it to the TargetSecretSet field.
func (o *OrganizationWebhookCreateResponse) SetTargetSecretSet(v bool) {
	o.TargetSecretSet = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *OrganizationWebhookCreateResponse) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationWebhookCreateResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *OrganizationWebhookCreateResponse) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *OrganizationWebhookCreateResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *OrganizationWebhookCreateResponse) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationWebhookCreateResponse) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *OrganizationWebhookCreateResponse) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *OrganizationWebhookCreateResponse) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *OrganizationWebhookCreateResponse) GetEvents() []Items {
	if o == nil || o.Events == nil {
		var ret []Items
		return ret
	}
	return o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationWebhookCreateResponse) GetEventsOk() ([]Items, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *OrganizationWebhookCreateResponse) HasEvents() bool {
	if o != nil && o.Events != nil {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []Items and assigns it to the Events field.
func (o *OrganizationWebhookCreateResponse) SetEvents(v []Items) {
	o.Events = v
}

// GetProjectIdFilter returns the ProjectIdFilter field value if set, zero value otherwise.
func (o *OrganizationWebhookCreateResponse) GetProjectIdFilter() []string {
	if o == nil || o.ProjectIdFilter == nil {
		var ret []string
		return ret
	}
	return o.ProjectIdFilter
}

// GetProjectIdFilterOk returns a tuple with the ProjectIdFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationWebhookCreateResponse) GetProjectIdFilterOk() ([]string, bool) {
	if o == nil || o.ProjectIdFilter == nil {
		return nil, false
	}
	return o.ProjectIdFilter, true
}

// HasProjectIdFilter returns a boolean if a field has been set.
func (o *OrganizationWebhookCreateResponse) HasProjectIdFilter() bool {
	if o != nil && o.ProjectIdFilter != nil {
		return true
	}

	return false
}

// SetProjectIdFilter gets a reference to the given []string and assigns it to the ProjectIdFilter field.
func (o *OrganizationWebhookCreateResponse) SetProjectIdFilter(v []string) {
	o.ProjectIdFilter = v
}

// GetEnvironmentTypesFilter returns the EnvironmentTypesFilter field value if set, zero value otherwise.
func (o *OrganizationWebhookCreateResponse) GetEnvironmentTypesFilter() []EnvironmentModeEnum {
	if o == nil || o.EnvironmentTypesFilter == nil {
		var ret []EnvironmentModeEnum
		return ret
	}
	return o.EnvironmentTypesFilter
}

// GetEnvironmentTypesFilterOk returns a tuple with the EnvironmentTypesFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationWebhookCreateResponse) GetEnvironmentTypesFilterOk() ([]EnvironmentModeEnum, bool) {
	if o == nil || o.EnvironmentTypesFilter == nil {
		return nil, false
	}
	return o.EnvironmentTypesFilter, true
}

// HasEnvironmentTypesFilter returns a boolean if a field has been set.
func (o *OrganizationWebhookCreateResponse) HasEnvironmentTypesFilter() bool {
	if o != nil && o.EnvironmentTypesFilter != nil {
		return true
	}

	return false
}

// SetEnvironmentTypesFilter gets a reference to the given []EnvironmentModeEnum and assigns it to the EnvironmentTypesFilter field.
func (o *OrganizationWebhookCreateResponse) SetEnvironmentTypesFilter(v []EnvironmentModeEnum) {
	o.EnvironmentTypesFilter = v
}

func (o OrganizationWebhookCreateResponse) MarshalJSON() ([]byte, error) {
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
	if o.Kind != nil {
		toSerialize["kind"] = o.Kind
	}
	if o.TargetUrl != nil {
		toSerialize["target_url"] = o.TargetUrl
	}
	if o.TargetSecretSet != nil {
		toSerialize["target_secret_set"] = o.TargetSecretSet
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Events != nil {
		toSerialize["events"] = o.Events
	}
	if o.ProjectIdFilter != nil {
		toSerialize["project_id_filter"] = o.ProjectIdFilter
	}
	if o.EnvironmentTypesFilter != nil {
		toSerialize["environment_types_filter"] = o.EnvironmentTypesFilter
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationWebhookCreateResponse struct {
	value *OrganizationWebhookCreateResponse
	isSet bool
}

func (v NullableOrganizationWebhookCreateResponse) Get() *OrganizationWebhookCreateResponse {
	return v.value
}

func (v *NullableOrganizationWebhookCreateResponse) Set(val *OrganizationWebhookCreateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationWebhookCreateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationWebhookCreateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationWebhookCreateResponse(val *OrganizationWebhookCreateResponse) *NullableOrganizationWebhookCreateResponse {
	return &NullableOrganizationWebhookCreateResponse{value: val, isSet: true}
}

func (v NullableOrganizationWebhookCreateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationWebhookCreateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
