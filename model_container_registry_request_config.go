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

// ContainerRegistryRequestConfig This field is dependent of the container registry kind: * `ECR` needs in the config: region, access_key_id, secret_access_key * `SCALEWAY_CR` needs in the config: region, scaleway_access_key, scaleway_secret_key * `DOCKER_HUB` needs in the config: username, password * `PUBLIC_ECR` needs in the config: access_key_id, secret_access_key * `DOCR` is not supported anymore
type ContainerRegistryRequestConfig struct {
	// Required if kind is `ECR` or `PUBLIC_ECR`
	AccessKeyId *string `json:"access_key_id,omitempty"`
	// Required if kind is `ECR` or `PUBLIC_ECR`
	SecretAccessKey *string `json:"secret_access_key,omitempty"`
	// Required if kind is `ECR` or `SCALEWAY_CR`
	Region *string `json:"region,omitempty"`
	// Required if kind is `SCALEWAY_CR`
	ScalewayAccessKey *string `json:"scaleway_access_key,omitempty"`
	// Required if kind is `SCALEWAY_CR`
	ScalewaySecretKey *string `json:"scaleway_secret_key,omitempty"`
	// Required if kind is `DOCKER_HUB`
	Username *string `json:"username,omitempty"`
	// Required if kind is `DOCKER_HUB`
	Password *string `json:"password,omitempty"`
}

// NewContainerRegistryRequestConfig instantiates a new ContainerRegistryRequestConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerRegistryRequestConfig() *ContainerRegistryRequestConfig {
	this := ContainerRegistryRequestConfig{}
	return &this
}

// NewContainerRegistryRequestConfigWithDefaults instantiates a new ContainerRegistryRequestConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerRegistryRequestConfigWithDefaults() *ContainerRegistryRequestConfig {
	this := ContainerRegistryRequestConfig{}
	return &this
}

// GetAccessKeyId returns the AccessKeyId field value if set, zero value otherwise.
func (o *ContainerRegistryRequestConfig) GetAccessKeyId() string {
	if o == nil || o.AccessKeyId == nil {
		var ret string
		return ret
	}
	return *o.AccessKeyId
}

// GetAccessKeyIdOk returns a tuple with the AccessKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerRegistryRequestConfig) GetAccessKeyIdOk() (*string, bool) {
	if o == nil || o.AccessKeyId == nil {
		return nil, false
	}
	return o.AccessKeyId, true
}

// HasAccessKeyId returns a boolean if a field has been set.
func (o *ContainerRegistryRequestConfig) HasAccessKeyId() bool {
	if o != nil && o.AccessKeyId != nil {
		return true
	}

	return false
}

// SetAccessKeyId gets a reference to the given string and assigns it to the AccessKeyId field.
func (o *ContainerRegistryRequestConfig) SetAccessKeyId(v string) {
	o.AccessKeyId = &v
}

// GetSecretAccessKey returns the SecretAccessKey field value if set, zero value otherwise.
func (o *ContainerRegistryRequestConfig) GetSecretAccessKey() string {
	if o == nil || o.SecretAccessKey == nil {
		var ret string
		return ret
	}
	return *o.SecretAccessKey
}

// GetSecretAccessKeyOk returns a tuple with the SecretAccessKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerRegistryRequestConfig) GetSecretAccessKeyOk() (*string, bool) {
	if o == nil || o.SecretAccessKey == nil {
		return nil, false
	}
	return o.SecretAccessKey, true
}

// HasSecretAccessKey returns a boolean if a field has been set.
func (o *ContainerRegistryRequestConfig) HasSecretAccessKey() bool {
	if o != nil && o.SecretAccessKey != nil {
		return true
	}

	return false
}

// SetSecretAccessKey gets a reference to the given string and assigns it to the SecretAccessKey field.
func (o *ContainerRegistryRequestConfig) SetSecretAccessKey(v string) {
	o.SecretAccessKey = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *ContainerRegistryRequestConfig) GetRegion() string {
	if o == nil || o.Region == nil {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerRegistryRequestConfig) GetRegionOk() (*string, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *ContainerRegistryRequestConfig) HasRegion() bool {
	if o != nil && o.Region != nil {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *ContainerRegistryRequestConfig) SetRegion(v string) {
	o.Region = &v
}

// GetScalewayAccessKey returns the ScalewayAccessKey field value if set, zero value otherwise.
func (o *ContainerRegistryRequestConfig) GetScalewayAccessKey() string {
	if o == nil || o.ScalewayAccessKey == nil {
		var ret string
		return ret
	}
	return *o.ScalewayAccessKey
}

// GetScalewayAccessKeyOk returns a tuple with the ScalewayAccessKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerRegistryRequestConfig) GetScalewayAccessKeyOk() (*string, bool) {
	if o == nil || o.ScalewayAccessKey == nil {
		return nil, false
	}
	return o.ScalewayAccessKey, true
}

// HasScalewayAccessKey returns a boolean if a field has been set.
func (o *ContainerRegistryRequestConfig) HasScalewayAccessKey() bool {
	if o != nil && o.ScalewayAccessKey != nil {
		return true
	}

	return false
}

// SetScalewayAccessKey gets a reference to the given string and assigns it to the ScalewayAccessKey field.
func (o *ContainerRegistryRequestConfig) SetScalewayAccessKey(v string) {
	o.ScalewayAccessKey = &v
}

// GetScalewaySecretKey returns the ScalewaySecretKey field value if set, zero value otherwise.
func (o *ContainerRegistryRequestConfig) GetScalewaySecretKey() string {
	if o == nil || o.ScalewaySecretKey == nil {
		var ret string
		return ret
	}
	return *o.ScalewaySecretKey
}

// GetScalewaySecretKeyOk returns a tuple with the ScalewaySecretKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerRegistryRequestConfig) GetScalewaySecretKeyOk() (*string, bool) {
	if o == nil || o.ScalewaySecretKey == nil {
		return nil, false
	}
	return o.ScalewaySecretKey, true
}

// HasScalewaySecretKey returns a boolean if a field has been set.
func (o *ContainerRegistryRequestConfig) HasScalewaySecretKey() bool {
	if o != nil && o.ScalewaySecretKey != nil {
		return true
	}

	return false
}

// SetScalewaySecretKey gets a reference to the given string and assigns it to the ScalewaySecretKey field.
func (o *ContainerRegistryRequestConfig) SetScalewaySecretKey(v string) {
	o.ScalewaySecretKey = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *ContainerRegistryRequestConfig) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerRegistryRequestConfig) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *ContainerRegistryRequestConfig) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *ContainerRegistryRequestConfig) SetUsername(v string) {
	o.Username = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *ContainerRegistryRequestConfig) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerRegistryRequestConfig) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *ContainerRegistryRequestConfig) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *ContainerRegistryRequestConfig) SetPassword(v string) {
	o.Password = &v
}

func (o ContainerRegistryRequestConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccessKeyId != nil {
		toSerialize["access_key_id"] = o.AccessKeyId
	}
	if o.SecretAccessKey != nil {
		toSerialize["secret_access_key"] = o.SecretAccessKey
	}
	if o.Region != nil {
		toSerialize["region"] = o.Region
	}
	if o.ScalewayAccessKey != nil {
		toSerialize["scaleway_access_key"] = o.ScalewayAccessKey
	}
	if o.ScalewaySecretKey != nil {
		toSerialize["scaleway_secret_key"] = o.ScalewaySecretKey
	}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	return json.Marshal(toSerialize)
}

type NullableContainerRegistryRequestConfig struct {
	value *ContainerRegistryRequestConfig
	isSet bool
}

func (v NullableContainerRegistryRequestConfig) Get() *ContainerRegistryRequestConfig {
	return v.value
}

func (v *NullableContainerRegistryRequestConfig) Set(val *ContainerRegistryRequestConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerRegistryRequestConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerRegistryRequestConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerRegistryRequestConfig(val *ContainerRegistryRequestConfig) *NullableContainerRegistryRequestConfig {
	return &NullableContainerRegistryRequestConfig{value: val, isSet: true}
}

func (v NullableContainerRegistryRequestConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerRegistryRequestConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
