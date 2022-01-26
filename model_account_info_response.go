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

// AccountInfoResponse struct for AccountInfoResponse
type AccountInfoResponse struct {
	Id *string `json:"id,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	Nickname *string `json:"nickname,omitempty"`
	FirstName *string `json:"first_name,omitempty"`
	LastName *string `json:"last_name,omitempty"`
	ProfilePictureUrl *string `json:"profile_picture_url,omitempty"`
}

// NewAccountInfoResponse instantiates a new AccountInfoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountInfoResponse() *AccountInfoResponse {
	this := AccountInfoResponse{}
	return &this
}

// NewAccountInfoResponseWithDefaults instantiates a new AccountInfoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountInfoResponseWithDefaults() *AccountInfoResponse {
	this := AccountInfoResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AccountInfoResponse) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInfoResponse) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AccountInfoResponse) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AccountInfoResponse) SetId(v string) {
	o.Id = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AccountInfoResponse) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInfoResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AccountInfoResponse) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *AccountInfoResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetNickname returns the Nickname field value if set, zero value otherwise.
func (o *AccountInfoResponse) GetNickname() string {
	if o == nil || o.Nickname == nil {
		var ret string
		return ret
	}
	return *o.Nickname
}

// GetNicknameOk returns a tuple with the Nickname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInfoResponse) GetNicknameOk() (*string, bool) {
	if o == nil || o.Nickname == nil {
		return nil, false
	}
	return o.Nickname, true
}

// HasNickname returns a boolean if a field has been set.
func (o *AccountInfoResponse) HasNickname() bool {
	if o != nil && o.Nickname != nil {
		return true
	}

	return false
}

// SetNickname gets a reference to the given string and assigns it to the Nickname field.
func (o *AccountInfoResponse) SetNickname(v string) {
	o.Nickname = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *AccountInfoResponse) GetFirstName() string {
	if o == nil || o.FirstName == nil {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInfoResponse) GetFirstNameOk() (*string, bool) {
	if o == nil || o.FirstName == nil {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *AccountInfoResponse) HasFirstName() bool {
	if o != nil && o.FirstName != nil {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *AccountInfoResponse) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *AccountInfoResponse) GetLastName() string {
	if o == nil || o.LastName == nil {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInfoResponse) GetLastNameOk() (*string, bool) {
	if o == nil || o.LastName == nil {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *AccountInfoResponse) HasLastName() bool {
	if o != nil && o.LastName != nil {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *AccountInfoResponse) SetLastName(v string) {
	o.LastName = &v
}

// GetProfilePictureUrl returns the ProfilePictureUrl field value if set, zero value otherwise.
func (o *AccountInfoResponse) GetProfilePictureUrl() string {
	if o == nil || o.ProfilePictureUrl == nil {
		var ret string
		return ret
	}
	return *o.ProfilePictureUrl
}

// GetProfilePictureUrlOk returns a tuple with the ProfilePictureUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInfoResponse) GetProfilePictureUrlOk() (*string, bool) {
	if o == nil || o.ProfilePictureUrl == nil {
		return nil, false
	}
	return o.ProfilePictureUrl, true
}

// HasProfilePictureUrl returns a boolean if a field has been set.
func (o *AccountInfoResponse) HasProfilePictureUrl() bool {
	if o != nil && o.ProfilePictureUrl != nil {
		return true
	}

	return false
}

// SetProfilePictureUrl gets a reference to the given string and assigns it to the ProfilePictureUrl field.
func (o *AccountInfoResponse) SetProfilePictureUrl(v string) {
	o.ProfilePictureUrl = &v
}

func (o AccountInfoResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.Nickname != nil {
		toSerialize["nickname"] = o.Nickname
	}
	if o.FirstName != nil {
		toSerialize["first_name"] = o.FirstName
	}
	if o.LastName != nil {
		toSerialize["last_name"] = o.LastName
	}
	if o.ProfilePictureUrl != nil {
		toSerialize["profile_picture_url"] = o.ProfilePictureUrl
	}
	return json.Marshal(toSerialize)
}

type NullableAccountInfoResponse struct {
	value *AccountInfoResponse
	isSet bool
}

func (v NullableAccountInfoResponse) Get() *AccountInfoResponse {
	return v.value
}

func (v *NullableAccountInfoResponse) Set(val *AccountInfoResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountInfoResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountInfoResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountInfoResponse(val *AccountInfoResponse) *NullableAccountInfoResponse {
	return &NullableAccountInfoResponse{value: val, isSet: true}
}

func (v NullableAccountInfoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountInfoResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


