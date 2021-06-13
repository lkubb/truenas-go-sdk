/*
 * TrueNAS RESTful API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package truenas

import (
	"encoding/json"
)

// SmbSharesecGetacl1 struct for SmbSharesecGetacl1
type SmbSharesecGetacl1 struct {
	ResolveSids *bool `json:"resolve_sids,omitempty"`
}

// NewSmbSharesecGetacl1 instantiates a new SmbSharesecGetacl1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmbSharesecGetacl1() *SmbSharesecGetacl1 {
	this := SmbSharesecGetacl1{}
	return &this
}

// NewSmbSharesecGetacl1WithDefaults instantiates a new SmbSharesecGetacl1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmbSharesecGetacl1WithDefaults() *SmbSharesecGetacl1 {
	this := SmbSharesecGetacl1{}
	return &this
}

// GetResolveSids returns the ResolveSids field value if set, zero value otherwise.
func (o *SmbSharesecGetacl1) GetResolveSids() bool {
	if o == nil || o.ResolveSids == nil {
		var ret bool
		return ret
	}
	return *o.ResolveSids
}

// GetResolveSidsOk returns a tuple with the ResolveSids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmbSharesecGetacl1) GetResolveSidsOk() (*bool, bool) {
	if o == nil || o.ResolveSids == nil {
		return nil, false
	}
	return o.ResolveSids, true
}

// HasResolveSids returns a boolean if a field has been set.
func (o *SmbSharesecGetacl1) HasResolveSids() bool {
	if o != nil && o.ResolveSids != nil {
		return true
	}

	return false
}

// SetResolveSids gets a reference to the given bool and assigns it to the ResolveSids field.
func (o *SmbSharesecGetacl1) SetResolveSids(v bool) {
	o.ResolveSids = &v
}

func (o SmbSharesecGetacl1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ResolveSids != nil {
		toSerialize["resolve_sids"] = o.ResolveSids
	}
	return json.Marshal(toSerialize)
}

type NullableSmbSharesecGetacl1 struct {
	value *SmbSharesecGetacl1
	isSet bool
}

func (v NullableSmbSharesecGetacl1) Get() *SmbSharesecGetacl1 {
	return v.value
}

func (v *NullableSmbSharesecGetacl1) Set(val *SmbSharesecGetacl1) {
	v.value = val
	v.isSet = true
}

func (v NullableSmbSharesecGetacl1) IsSet() bool {
	return v.isSet
}

func (v *NullableSmbSharesecGetacl1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmbSharesecGetacl1(val *SmbSharesecGetacl1) *NullableSmbSharesecGetacl1 {
	return &NullableSmbSharesecGetacl1{value: val, isSet: true}
}

func (v NullableSmbSharesecGetacl1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmbSharesecGetacl1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

