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

// WebdavUpdate0 struct for WebdavUpdate0
type WebdavUpdate0 struct {
	Protocol *string `json:"protocol,omitempty"`
	Tcpport *int32 `json:"tcpport,omitempty"`
	Tcpportssl *int32 `json:"tcpportssl,omitempty"`
	Password *string `json:"password,omitempty"`
	Htauth *string `json:"htauth,omitempty"`
	Certssl NullableInt32 `json:"certssl,omitempty"`
}

// NewWebdavUpdate0 instantiates a new WebdavUpdate0 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebdavUpdate0() *WebdavUpdate0 {
	this := WebdavUpdate0{}
	return &this
}

// NewWebdavUpdate0WithDefaults instantiates a new WebdavUpdate0 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebdavUpdate0WithDefaults() *WebdavUpdate0 {
	this := WebdavUpdate0{}
	return &this
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *WebdavUpdate0) GetProtocol() string {
	if o == nil || o.Protocol == nil {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebdavUpdate0) GetProtocolOk() (*string, bool) {
	if o == nil || o.Protocol == nil {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *WebdavUpdate0) HasProtocol() bool {
	if o != nil && o.Protocol != nil {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *WebdavUpdate0) SetProtocol(v string) {
	o.Protocol = &v
}

// GetTcpport returns the Tcpport field value if set, zero value otherwise.
func (o *WebdavUpdate0) GetTcpport() int32 {
	if o == nil || o.Tcpport == nil {
		var ret int32
		return ret
	}
	return *o.Tcpport
}

// GetTcpportOk returns a tuple with the Tcpport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebdavUpdate0) GetTcpportOk() (*int32, bool) {
	if o == nil || o.Tcpport == nil {
		return nil, false
	}
	return o.Tcpport, true
}

// HasTcpport returns a boolean if a field has been set.
func (o *WebdavUpdate0) HasTcpport() bool {
	if o != nil && o.Tcpport != nil {
		return true
	}

	return false
}

// SetTcpport gets a reference to the given int32 and assigns it to the Tcpport field.
func (o *WebdavUpdate0) SetTcpport(v int32) {
	o.Tcpport = &v
}

// GetTcpportssl returns the Tcpportssl field value if set, zero value otherwise.
func (o *WebdavUpdate0) GetTcpportssl() int32 {
	if o == nil || o.Tcpportssl == nil {
		var ret int32
		return ret
	}
	return *o.Tcpportssl
}

// GetTcpportsslOk returns a tuple with the Tcpportssl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebdavUpdate0) GetTcpportsslOk() (*int32, bool) {
	if o == nil || o.Tcpportssl == nil {
		return nil, false
	}
	return o.Tcpportssl, true
}

// HasTcpportssl returns a boolean if a field has been set.
func (o *WebdavUpdate0) HasTcpportssl() bool {
	if o != nil && o.Tcpportssl != nil {
		return true
	}

	return false
}

// SetTcpportssl gets a reference to the given int32 and assigns it to the Tcpportssl field.
func (o *WebdavUpdate0) SetTcpportssl(v int32) {
	o.Tcpportssl = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *WebdavUpdate0) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebdavUpdate0) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *WebdavUpdate0) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *WebdavUpdate0) SetPassword(v string) {
	o.Password = &v
}

// GetHtauth returns the Htauth field value if set, zero value otherwise.
func (o *WebdavUpdate0) GetHtauth() string {
	if o == nil || o.Htauth == nil {
		var ret string
		return ret
	}
	return *o.Htauth
}

// GetHtauthOk returns a tuple with the Htauth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebdavUpdate0) GetHtauthOk() (*string, bool) {
	if o == nil || o.Htauth == nil {
		return nil, false
	}
	return o.Htauth, true
}

// HasHtauth returns a boolean if a field has been set.
func (o *WebdavUpdate0) HasHtauth() bool {
	if o != nil && o.Htauth != nil {
		return true
	}

	return false
}

// SetHtauth gets a reference to the given string and assigns it to the Htauth field.
func (o *WebdavUpdate0) SetHtauth(v string) {
	o.Htauth = &v
}

// GetCertssl returns the Certssl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebdavUpdate0) GetCertssl() int32 {
	if o == nil || o.Certssl.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Certssl.Get()
}

// GetCertsslOk returns a tuple with the Certssl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebdavUpdate0) GetCertsslOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Certssl.Get(), o.Certssl.IsSet()
}

// HasCertssl returns a boolean if a field has been set.
func (o *WebdavUpdate0) HasCertssl() bool {
	if o != nil && o.Certssl.IsSet() {
		return true
	}

	return false
}

// SetCertssl gets a reference to the given NullableInt32 and assigns it to the Certssl field.
func (o *WebdavUpdate0) SetCertssl(v int32) {
	o.Certssl.Set(&v)
}
// SetCertsslNil sets the value for Certssl to be an explicit nil
func (o *WebdavUpdate0) SetCertsslNil() {
	o.Certssl.Set(nil)
}

// UnsetCertssl ensures that no value is present for Certssl, not even an explicit nil
func (o *WebdavUpdate0) UnsetCertssl() {
	o.Certssl.Unset()
}

func (o WebdavUpdate0) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Protocol != nil {
		toSerialize["protocol"] = o.Protocol
	}
	if o.Tcpport != nil {
		toSerialize["tcpport"] = o.Tcpport
	}
	if o.Tcpportssl != nil {
		toSerialize["tcpportssl"] = o.Tcpportssl
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	if o.Htauth != nil {
		toSerialize["htauth"] = o.Htauth
	}
	if o.Certssl.IsSet() {
		toSerialize["certssl"] = o.Certssl.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableWebdavUpdate0 struct {
	value *WebdavUpdate0
	isSet bool
}

func (v NullableWebdavUpdate0) Get() *WebdavUpdate0 {
	return v.value
}

func (v *NullableWebdavUpdate0) Set(val *WebdavUpdate0) {
	v.value = val
	v.isSet = true
}

func (v NullableWebdavUpdate0) IsSet() bool {
	return v.isSet
}

func (v *NullableWebdavUpdate0) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebdavUpdate0(val *WebdavUpdate0) *NullableWebdavUpdate0 {
	return &NullableWebdavUpdate0{value: val, isSet: true}
}

func (v NullableWebdavUpdate0) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebdavUpdate0) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

