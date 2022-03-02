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

// OrganizationApiTokenCreateResponse struct for OrganizationApiTokenCreateResponse
type OrganizationApiTokenCreateResponse struct {
	Id          string     `json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
	Name        *string    `json:"name,omitempty"`
	Description *string    `json:"description,omitempty"`
	// the generated token to send in 'Authorization' header prefixed by 'Token '
	Token *string                    `json:"token,omitempty"`
	Scope *OrganizationApiTokenScope `json:"scope,omitempty"`
}

// NewOrganizationApiTokenCreateResponse instantiates a new OrganizationApiTokenCreateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationApiTokenCreateResponse(id string, createdAt time.Time) *OrganizationApiTokenCreateResponse {
	this := OrganizationApiTokenCreateResponse{}
	this.Id = id
	this.CreatedAt = createdAt
	return &this
}

// NewOrganizationApiTokenCreateResponseWithDefaults instantiates a new OrganizationApiTokenCreateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationApiTokenCreateResponseWithDefaults() *OrganizationApiTokenCreateResponse {
	this := OrganizationApiTokenCreateResponse{}
	return &this
}

// GetId returns the Id field value
func (o *OrganizationApiTokenCreateResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *OrganizationApiTokenCreateResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *OrganizationApiTokenCreateResponse) SetId(v string) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *OrganizationApiTokenCreateResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *OrganizationApiTokenCreateResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *OrganizationApiTokenCreateResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *OrganizationApiTokenCreateResponse) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationApiTokenCreateResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *OrganizationApiTokenCreateResponse) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *OrganizationApiTokenCreateResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OrganizationApiTokenCreateResponse) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationApiTokenCreateResponse) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OrganizationApiTokenCreateResponse) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OrganizationApiTokenCreateResponse) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *OrganizationApiTokenCreateResponse) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationApiTokenCreateResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *OrganizationApiTokenCreateResponse) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *OrganizationApiTokenCreateResponse) SetDescription(v string) {
	o.Description = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *OrganizationApiTokenCreateResponse) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationApiTokenCreateResponse) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *OrganizationApiTokenCreateResponse) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *OrganizationApiTokenCreateResponse) SetToken(v string) {
	o.Token = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *OrganizationApiTokenCreateResponse) GetScope() OrganizationApiTokenScope {
	if o == nil || o.Scope == nil {
		var ret OrganizationApiTokenScope
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationApiTokenCreateResponse) GetScopeOk() (*OrganizationApiTokenScope, bool) {
	if o == nil || o.Scope == nil {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *OrganizationApiTokenCreateResponse) HasScope() bool {
	if o != nil && o.Scope != nil {
		return true
	}

	return false
}

// SetScope gets a reference to the given OrganizationApiTokenScope and assigns it to the Scope field.
func (o *OrganizationApiTokenCreateResponse) SetScope(v OrganizationApiTokenScope) {
	o.Scope = &v
}

func (o OrganizationApiTokenCreateResponse) MarshalJSON() ([]byte, error) {
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
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.Scope != nil {
		toSerialize["scope"] = o.Scope
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationApiTokenCreateResponse struct {
	value *OrganizationApiTokenCreateResponse
	isSet bool
}

func (v NullableOrganizationApiTokenCreateResponse) Get() *OrganizationApiTokenCreateResponse {
	return v.value
}

func (v *NullableOrganizationApiTokenCreateResponse) Set(val *OrganizationApiTokenCreateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationApiTokenCreateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationApiTokenCreateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationApiTokenCreateResponse(val *OrganizationApiTokenCreateResponse) *NullableOrganizationApiTokenCreateResponse {
	return &NullableOrganizationApiTokenCreateResponse{value: val, isSet: true}
}

func (v NullableOrganizationApiTokenCreateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationApiTokenCreateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
