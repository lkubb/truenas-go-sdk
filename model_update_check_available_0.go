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

// UpdateCheckAvailable0 struct for UpdateCheckAvailable0
type UpdateCheckAvailable0 struct {
	Train *string `json:"train,omitempty"`
}

// NewUpdateCheckAvailable0 instantiates a new UpdateCheckAvailable0 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateCheckAvailable0() *UpdateCheckAvailable0 {
	this := UpdateCheckAvailable0{}
	return &this
}

// NewUpdateCheckAvailable0WithDefaults instantiates a new UpdateCheckAvailable0 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateCheckAvailable0WithDefaults() *UpdateCheckAvailable0 {
	this := UpdateCheckAvailable0{}
	return &this
}

// GetTrain returns the Train field value if set, zero value otherwise.
func (o *UpdateCheckAvailable0) GetTrain() string {
	if o == nil || o.Train == nil {
		var ret string
		return ret
	}
	return *o.Train
}

// GetTrainOk returns a tuple with the Train field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCheckAvailable0) GetTrainOk() (*string, bool) {
	if o == nil || o.Train == nil {
		return nil, false
	}
	return o.Train, true
}

// HasTrain returns a boolean if a field has been set.
func (o *UpdateCheckAvailable0) HasTrain() bool {
	if o != nil && o.Train != nil {
		return true
	}

	return false
}

// SetTrain gets a reference to the given string and assigns it to the Train field.
func (o *UpdateCheckAvailable0) SetTrain(v string) {
	o.Train = &v
}

func (o UpdateCheckAvailable0) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Train != nil {
		toSerialize["train"] = o.Train
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateCheckAvailable0 struct {
	value *UpdateCheckAvailable0
	isSet bool
}

func (v NullableUpdateCheckAvailable0) Get() *UpdateCheckAvailable0 {
	return v.value
}

func (v *NullableUpdateCheckAvailable0) Set(val *UpdateCheckAvailable0) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateCheckAvailable0) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateCheckAvailable0) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateCheckAvailable0(val *UpdateCheckAvailable0) *NullableUpdateCheckAvailable0 {
	return &NullableUpdateCheckAvailable0{value: val, isSet: true}
}

func (v NullableUpdateCheckAvailable0) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateCheckAvailable0) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

