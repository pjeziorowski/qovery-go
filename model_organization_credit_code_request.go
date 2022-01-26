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
)

// OrganizationCreditCodeRequest struct for OrganizationCreditCodeRequest
type OrganizationCreditCodeRequest struct {
	Code *string `json:"code,omitempty"`
}

// NewOrganizationCreditCodeRequest instantiates a new OrganizationCreditCodeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationCreditCodeRequest() *OrganizationCreditCodeRequest {
	this := OrganizationCreditCodeRequest{}
	return &this
}

// NewOrganizationCreditCodeRequestWithDefaults instantiates a new OrganizationCreditCodeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationCreditCodeRequestWithDefaults() *OrganizationCreditCodeRequest {
	this := OrganizationCreditCodeRequest{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *OrganizationCreditCodeRequest) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationCreditCodeRequest) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *OrganizationCreditCodeRequest) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *OrganizationCreditCodeRequest) SetCode(v string) {
	o.Code = &v
}

func (o OrganizationCreditCodeRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationCreditCodeRequest struct {
	value *OrganizationCreditCodeRequest
	isSet bool
}

func (v NullableOrganizationCreditCodeRequest) Get() *OrganizationCreditCodeRequest {
	return v.value
}

func (v *NullableOrganizationCreditCodeRequest) Set(val *OrganizationCreditCodeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationCreditCodeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationCreditCodeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationCreditCodeRequest(val *OrganizationCreditCodeRequest) *NullableOrganizationCreditCodeRequest {
	return &NullableOrganizationCreditCodeRequest{value: val, isSet: true}
}

func (v NullableOrganizationCreditCodeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationCreditCodeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


