/*
[BETA] Qovery API

- Qovery is the fastest way to deploy your full-stack apps on any Cloud provider. - ℹ️ The API is in Beta and still in progress. Some endpoints are not available yet.  # Authentication  <!-- ReDoc-Inject: <security-definitions> -->

API version: 1.0.0
Contact: support+api+documentation@qovery.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package qovery

import (
	"encoding/json"
	"time"
)

// CustomDomainResponse struct for CustomDomainResponse
type CustomDomainResponse struct {
	// URL provided by Qovery. You must create a CNAME on your DNS provider using that URL
	ValidationDomain *string `json:"validation_domain,omitempty"`
	Status *string `json:"status,omitempty"`
	Id string `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// your custom domain
	Domain string `json:"domain"`
}

// NewCustomDomainResponse instantiates a new CustomDomainResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomDomainResponse(id string, createdAt time.Time, domain string) *CustomDomainResponse {
	this := CustomDomainResponse{}
	this.Id = id
	this.CreatedAt = createdAt
	this.Domain = domain
	return &this
}

// NewCustomDomainResponseWithDefaults instantiates a new CustomDomainResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomDomainResponseWithDefaults() *CustomDomainResponse {
	this := CustomDomainResponse{}
	return &this
}

// GetValidationDomain returns the ValidationDomain field value if set, zero value otherwise.
func (o *CustomDomainResponse) GetValidationDomain() string {
	if o == nil || o.ValidationDomain == nil {
		var ret string
		return ret
	}
	return *o.ValidationDomain
}

// GetValidationDomainOk returns a tuple with the ValidationDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomDomainResponse) GetValidationDomainOk() (*string, bool) {
	if o == nil || o.ValidationDomain == nil {
		return nil, false
	}
	return o.ValidationDomain, true
}

// HasValidationDomain returns a boolean if a field has been set.
func (o *CustomDomainResponse) HasValidationDomain() bool {
	if o != nil && o.ValidationDomain != nil {
		return true
	}

	return false
}

// SetValidationDomain gets a reference to the given string and assigns it to the ValidationDomain field.
func (o *CustomDomainResponse) SetValidationDomain(v string) {
	o.ValidationDomain = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CustomDomainResponse) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomDomainResponse) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CustomDomainResponse) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *CustomDomainResponse) SetStatus(v string) {
	o.Status = &v
}

// GetId returns the Id field value
func (o *CustomDomainResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CustomDomainResponse) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CustomDomainResponse) SetId(v string) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *CustomDomainResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *CustomDomainResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *CustomDomainResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *CustomDomainResponse) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomDomainResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *CustomDomainResponse) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *CustomDomainResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetDomain returns the Domain field value
func (o *CustomDomainResponse) GetDomain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Domain
}

// GetDomainOk returns a tuple with the Domain field value
// and a boolean to check if the value has been set.
func (o *CustomDomainResponse) GetDomainOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Domain, true
}

// SetDomain sets field value
func (o *CustomDomainResponse) SetDomain(v string) {
	o.Domain = v
}

func (o CustomDomainResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ValidationDomain != nil {
		toSerialize["validation_domain"] = o.ValidationDomain
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if true {
		toSerialize["domain"] = o.Domain
	}
	return json.Marshal(toSerialize)
}

type NullableCustomDomainResponse struct {
	value *CustomDomainResponse
	isSet bool
}

func (v NullableCustomDomainResponse) Get() *CustomDomainResponse {
	return v.value
}

func (v *NullableCustomDomainResponse) Set(val *CustomDomainResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomDomainResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomDomainResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomDomainResponse(val *CustomDomainResponse) *NullableCustomDomainResponse {
	return &NullableCustomDomainResponse{value: val, isSet: true}
}

func (v NullableCustomDomainResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomDomainResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


