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

// RemainingCredits struct for RemainingCredits
type RemainingCredits struct {
	TotalInCents int32 `json:"total_in_cents"`
	Total float32 `json:"total"`
	CurrencyCode string `json:"currency_code"`
}

// NewRemainingCredits instantiates a new RemainingCredits object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemainingCredits(totalInCents int32, total float32, currencyCode string) *RemainingCredits {
	this := RemainingCredits{}
	this.TotalInCents = totalInCents
	this.Total = total
	this.CurrencyCode = currencyCode
	return &this
}

// NewRemainingCreditsWithDefaults instantiates a new RemainingCredits object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemainingCreditsWithDefaults() *RemainingCredits {
	this := RemainingCredits{}
	return &this
}

// GetTotalInCents returns the TotalInCents field value
func (o *RemainingCredits) GetTotalInCents() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalInCents
}

// GetTotalInCentsOk returns a tuple with the TotalInCents field value
// and a boolean to check if the value has been set.
func (o *RemainingCredits) GetTotalInCentsOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TotalInCents, true
}

// SetTotalInCents sets field value
func (o *RemainingCredits) SetTotalInCents(v int32) {
	o.TotalInCents = v
}

// GetTotal returns the Total field value
func (o *RemainingCredits) GetTotal() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *RemainingCredits) GetTotalOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *RemainingCredits) SetTotal(v float32) {
	o.Total = v
}

// GetCurrencyCode returns the CurrencyCode field value
func (o *RemainingCredits) GetCurrencyCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value
// and a boolean to check if the value has been set.
func (o *RemainingCredits) GetCurrencyCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CurrencyCode, true
}

// SetCurrencyCode sets field value
func (o *RemainingCredits) SetCurrencyCode(v string) {
	o.CurrencyCode = v
}

func (o RemainingCredits) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["total_in_cents"] = o.TotalInCents
	}
	if true {
		toSerialize["total"] = o.Total
	}
	if true {
		toSerialize["currency_code"] = o.CurrencyCode
	}
	return json.Marshal(toSerialize)
}

type NullableRemainingCredits struct {
	value *RemainingCredits
	isSet bool
}

func (v NullableRemainingCredits) Get() *RemainingCredits {
	return v.value
}

func (v *NullableRemainingCredits) Set(val *RemainingCredits) {
	v.value = val
	v.isSet = true
}

func (v NullableRemainingCredits) IsSet() bool {
	return v.isSet
}

func (v *NullableRemainingCredits) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemainingCredits(val *RemainingCredits) *NullableRemainingCredits {
	return &NullableRemainingCredits{value: val, isSet: true}
}

func (v NullableRemainingCredits) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemainingCredits) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


