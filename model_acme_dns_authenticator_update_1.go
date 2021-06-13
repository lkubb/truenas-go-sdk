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

// AcmeDnsAuthenticatorUpdate1 struct for AcmeDnsAuthenticatorUpdate1
type AcmeDnsAuthenticatorUpdate1 struct {
	Name *string `json:"name,omitempty"`
	Attributes *map[string]map[string]interface{} `json:"attributes,omitempty"`
}

// NewAcmeDnsAuthenticatorUpdate1 instantiates a new AcmeDnsAuthenticatorUpdate1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAcmeDnsAuthenticatorUpdate1() *AcmeDnsAuthenticatorUpdate1 {
	this := AcmeDnsAuthenticatorUpdate1{}
	return &this
}

// NewAcmeDnsAuthenticatorUpdate1WithDefaults instantiates a new AcmeDnsAuthenticatorUpdate1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAcmeDnsAuthenticatorUpdate1WithDefaults() *AcmeDnsAuthenticatorUpdate1 {
	this := AcmeDnsAuthenticatorUpdate1{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AcmeDnsAuthenticatorUpdate1) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcmeDnsAuthenticatorUpdate1) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AcmeDnsAuthenticatorUpdate1) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AcmeDnsAuthenticatorUpdate1) SetName(v string) {
	o.Name = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *AcmeDnsAuthenticatorUpdate1) GetAttributes() map[string]map[string]interface{} {
	if o == nil || o.Attributes == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcmeDnsAuthenticatorUpdate1) GetAttributesOk() (*map[string]map[string]interface{}, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *AcmeDnsAuthenticatorUpdate1) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]map[string]interface{} and assigns it to the Attributes field.
func (o *AcmeDnsAuthenticatorUpdate1) SetAttributes(v map[string]map[string]interface{}) {
	o.Attributes = &v
}

func (o AcmeDnsAuthenticatorUpdate1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	return json.Marshal(toSerialize)
}

type NullableAcmeDnsAuthenticatorUpdate1 struct {
	value *AcmeDnsAuthenticatorUpdate1
	isSet bool
}

func (v NullableAcmeDnsAuthenticatorUpdate1) Get() *AcmeDnsAuthenticatorUpdate1 {
	return v.value
}

func (v *NullableAcmeDnsAuthenticatorUpdate1) Set(val *AcmeDnsAuthenticatorUpdate1) {
	v.value = val
	v.isSet = true
}

func (v NullableAcmeDnsAuthenticatorUpdate1) IsSet() bool {
	return v.isSet
}

func (v *NullableAcmeDnsAuthenticatorUpdate1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAcmeDnsAuthenticatorUpdate1(val *AcmeDnsAuthenticatorUpdate1) *NullableAcmeDnsAuthenticatorUpdate1 {
	return &NullableAcmeDnsAuthenticatorUpdate1{value: val, isSet: true}
}

func (v NullableAcmeDnsAuthenticatorUpdate1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAcmeDnsAuthenticatorUpdate1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

