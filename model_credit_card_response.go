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

// CreditCardResponse struct for CreditCardResponse
type CreditCardResponse struct {
	Id string `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	ExpiryMonth int32 `json:"expiry_month"`
	ExpiryYear int32 `json:"expiry_year"`
	LastDigit string `json:"last_digit"`
	IsExpired bool `json:"is_expired"`
}

// NewCreditCardResponse instantiates a new CreditCardResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditCardResponse(id string, createdAt time.Time, expiryMonth int32, expiryYear int32, lastDigit string, isExpired bool) *CreditCardResponse {
	this := CreditCardResponse{}
	this.Id = id
	this.CreatedAt = createdAt
	this.ExpiryMonth = expiryMonth
	this.ExpiryYear = expiryYear
	this.LastDigit = lastDigit
	this.IsExpired = isExpired
	return &this
}

// NewCreditCardResponseWithDefaults instantiates a new CreditCardResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditCardResponseWithDefaults() *CreditCardResponse {
	this := CreditCardResponse{}
	return &this
}

// GetId returns the Id field value
func (o *CreditCardResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CreditCardResponse) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CreditCardResponse) SetId(v string) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *CreditCardResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *CreditCardResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *CreditCardResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetExpiryMonth returns the ExpiryMonth field value
func (o *CreditCardResponse) GetExpiryMonth() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ExpiryMonth
}

// GetExpiryMonthOk returns a tuple with the ExpiryMonth field value
// and a boolean to check if the value has been set.
func (o *CreditCardResponse) GetExpiryMonthOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ExpiryMonth, true
}

// SetExpiryMonth sets field value
func (o *CreditCardResponse) SetExpiryMonth(v int32) {
	o.ExpiryMonth = v
}

// GetExpiryYear returns the ExpiryYear field value
func (o *CreditCardResponse) GetExpiryYear() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ExpiryYear
}

// GetExpiryYearOk returns a tuple with the ExpiryYear field value
// and a boolean to check if the value has been set.
func (o *CreditCardResponse) GetExpiryYearOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ExpiryYear, true
}

// SetExpiryYear sets field value
func (o *CreditCardResponse) SetExpiryYear(v int32) {
	o.ExpiryYear = v
}

// GetLastDigit returns the LastDigit field value
func (o *CreditCardResponse) GetLastDigit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastDigit
}

// GetLastDigitOk returns a tuple with the LastDigit field value
// and a boolean to check if the value has been set.
func (o *CreditCardResponse) GetLastDigitOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LastDigit, true
}

// SetLastDigit sets field value
func (o *CreditCardResponse) SetLastDigit(v string) {
	o.LastDigit = v
}

// GetIsExpired returns the IsExpired field value
func (o *CreditCardResponse) GetIsExpired() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsExpired
}

// GetIsExpiredOk returns a tuple with the IsExpired field value
// and a boolean to check if the value has been set.
func (o *CreditCardResponse) GetIsExpiredOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsExpired, true
}

// SetIsExpired sets field value
func (o *CreditCardResponse) SetIsExpired(v bool) {
	o.IsExpired = v
}

func (o CreditCardResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["expiry_month"] = o.ExpiryMonth
	}
	if true {
		toSerialize["expiry_year"] = o.ExpiryYear
	}
	if true {
		toSerialize["last_digit"] = o.LastDigit
	}
	if true {
		toSerialize["is_expired"] = o.IsExpired
	}
	return json.Marshal(toSerialize)
}

type NullableCreditCardResponse struct {
	value *CreditCardResponse
	isSet bool
}

func (v NullableCreditCardResponse) Get() *CreditCardResponse {
	return v.value
}

func (v *NullableCreditCardResponse) Set(val *CreditCardResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditCardResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditCardResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditCardResponse(val *CreditCardResponse) *NullableCreditCardResponse {
	return &NullableCreditCardResponse{value: val, isSet: true}
}

func (v NullableCreditCardResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditCardResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


