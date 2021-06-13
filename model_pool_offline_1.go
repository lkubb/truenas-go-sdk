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

// PoolOffline1 struct for PoolOffline1
type PoolOffline1 struct {
	Label *string `json:"label,omitempty"`
}

// NewPoolOffline1 instantiates a new PoolOffline1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPoolOffline1() *PoolOffline1 {
	this := PoolOffline1{}
	return &this
}

// NewPoolOffline1WithDefaults instantiates a new PoolOffline1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPoolOffline1WithDefaults() *PoolOffline1 {
	this := PoolOffline1{}
	return &this
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *PoolOffline1) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolOffline1) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *PoolOffline1) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *PoolOffline1) SetLabel(v string) {
	o.Label = &v
}

func (o PoolOffline1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Label != nil {
		toSerialize["label"] = o.Label
	}
	return json.Marshal(toSerialize)
}

type NullablePoolOffline1 struct {
	value *PoolOffline1
	isSet bool
}

func (v NullablePoolOffline1) Get() *PoolOffline1 {
	return v.value
}

func (v *NullablePoolOffline1) Set(val *PoolOffline1) {
	v.value = val
	v.isSet = true
}

func (v NullablePoolOffline1) IsSet() bool {
	return v.isSet
}

func (v *NullablePoolOffline1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePoolOffline1(val *PoolOffline1) *NullablePoolOffline1 {
	return &NullablePoolOffline1{value: val, isSet: true}
}

func (v NullablePoolOffline1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePoolOffline1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

