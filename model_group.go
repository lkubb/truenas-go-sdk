/*
 * TrueNAS RESTful API
 *
 * Go SDK for interacting with TrueNAS APIs (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package truenas

import (
	"encoding/json"
)

// Group struct for Group
type Group struct {
	Id                   int32     `json:"id"`
	Gid                  *int32    `json:"gid,omitempty"`
	Group                string    `json:"group"`
	Builtin              *bool     `json:"builtin,omitempty"`
	Sudo                 *bool     `json:"sudo,omitempty"`
	SudoNopasswd         *bool     `json:"sudo_nopasswd,omitempty"`
	SudoCommands         *[]string `json:"sudo_commands,omitempty"`
	Smb                  *bool     `json:"smb,omitempty"`
	Users                *[]int32  `json:"users,omitempty"`
	Local                *bool     `json:"local,omitempty"`
	IdTypeBoth           *bool     `json:"id_type_both,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Group Group

// NewGroup instantiates a new Group object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroup(id int32, group string) *Group {
	this := Group{}
	this.Id = id
	this.Group = group
	return &this
}

// NewGroupWithDefaults instantiates a new Group object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupWithDefaults() *Group {
	this := Group{}
	return &this
}

// GetId returns the Id field value
func (o *Group) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Group) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Group) SetId(v int32) {
	o.Id = v
}

// GetGid returns the Gid field value if set, zero value otherwise.
func (o *Group) GetGid() int32 {
	if o == nil || o.Gid == nil {
		var ret int32
		return ret
	}
	return *o.Gid
}

// GetGidOk returns a tuple with the Gid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetGidOk() (*int32, bool) {
	if o == nil || o.Gid == nil {
		return nil, false
	}
	return o.Gid, true
}

// HasGid returns a boolean if a field has been set.
func (o *Group) HasGid() bool {
	if o != nil && o.Gid != nil {
		return true
	}

	return false
}

// SetGid gets a reference to the given int32 and assigns it to the Gid field.
func (o *Group) SetGid(v int32) {
	o.Gid = &v
}

// GetGroup returns the Group field value
func (o *Group) GetGroup() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Group
}

// GetGroupOk returns a tuple with the Group field value
// and a boolean to check if the value has been set.
func (o *Group) GetGroupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Group, true
}

// SetGroup sets field value
func (o *Group) SetGroup(v string) {
	o.Group = v
}

// GetBuiltin returns the Builtin field value if set, zero value otherwise.
func (o *Group) GetBuiltin() bool {
	if o == nil || o.Builtin == nil {
		var ret bool
		return ret
	}
	return *o.Builtin
}

// GetBuiltinOk returns a tuple with the Builtin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetBuiltinOk() (*bool, bool) {
	if o == nil || o.Builtin == nil {
		return nil, false
	}
	return o.Builtin, true
}

// HasBuiltin returns a boolean if a field has been set.
func (o *Group) HasBuiltin() bool {
	if o != nil && o.Builtin != nil {
		return true
	}

	return false
}

// SetBuiltin gets a reference to the given bool and assigns it to the Builtin field.
func (o *Group) SetBuiltin(v bool) {
	o.Builtin = &v
}

// GetSudo returns the Sudo field value if set, zero value otherwise.
func (o *Group) GetSudo() bool {
	if o == nil || o.Sudo == nil {
		var ret bool
		return ret
	}
	return *o.Sudo
}

// GetSudoOk returns a tuple with the Sudo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetSudoOk() (*bool, bool) {
	if o == nil || o.Sudo == nil {
		return nil, false
	}
	return o.Sudo, true
}

// HasSudo returns a boolean if a field has been set.
func (o *Group) HasSudo() bool {
	if o != nil && o.Sudo != nil {
		return true
	}

	return false
}

// SetSudo gets a reference to the given bool and assigns it to the Sudo field.
func (o *Group) SetSudo(v bool) {
	o.Sudo = &v
}

// GetSudoNopasswd returns the SudoNopasswd field value if set, zero value otherwise.
func (o *Group) GetSudoNopasswd() bool {
	if o == nil || o.SudoNopasswd == nil {
		var ret bool
		return ret
	}
	return *o.SudoNopasswd
}

// GetSudoNopasswdOk returns a tuple with the SudoNopasswd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetSudoNopasswdOk() (*bool, bool) {
	if o == nil || o.SudoNopasswd == nil {
		return nil, false
	}
	return o.SudoNopasswd, true
}

// HasSudoNopasswd returns a boolean if a field has been set.
func (o *Group) HasSudoNopasswd() bool {
	if o != nil && o.SudoNopasswd != nil {
		return true
	}

	return false
}

// SetSudoNopasswd gets a reference to the given bool and assigns it to the SudoNopasswd field.
func (o *Group) SetSudoNopasswd(v bool) {
	o.SudoNopasswd = &v
}

// GetSudoCommands returns the SudoCommands field value if set, zero value otherwise.
func (o *Group) GetSudoCommands() []string {
	if o == nil || o.SudoCommands == nil {
		var ret []string
		return ret
	}
	return *o.SudoCommands
}

// GetSudoCommandsOk returns a tuple with the SudoCommands field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetSudoCommandsOk() (*[]string, bool) {
	if o == nil || o.SudoCommands == nil {
		return nil, false
	}
	return o.SudoCommands, true
}

// HasSudoCommands returns a boolean if a field has been set.
func (o *Group) HasSudoCommands() bool {
	if o != nil && o.SudoCommands != nil {
		return true
	}

	return false
}

// SetSudoCommands gets a reference to the given []string and assigns it to the SudoCommands field.
func (o *Group) SetSudoCommands(v []string) {
	o.SudoCommands = &v
}

// GetSmb returns the Smb field value if set, zero value otherwise.
func (o *Group) GetSmb() bool {
	if o == nil || o.Smb == nil {
		var ret bool
		return ret
	}
	return *o.Smb
}

// GetSmbOk returns a tuple with the Smb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetSmbOk() (*bool, bool) {
	if o == nil || o.Smb == nil {
		return nil, false
	}
	return o.Smb, true
}

// HasSmb returns a boolean if a field has been set.
func (o *Group) HasSmb() bool {
	if o != nil && o.Smb != nil {
		return true
	}

	return false
}

// SetSmb gets a reference to the given bool and assigns it to the Smb field.
func (o *Group) SetSmb(v bool) {
	o.Smb = &v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *Group) GetUsers() []int32 {
	if o == nil || o.Users == nil {
		var ret []int32
		return ret
	}
	return *o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetUsersOk() (*[]int32, bool) {
	if o == nil || o.Users == nil {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *Group) HasUsers() bool {
	if o != nil && o.Users != nil {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []int32 and assigns it to the Users field.
func (o *Group) SetUsers(v []int32) {
	o.Users = &v
}

// GetLocal returns the Local field value if set, zero value otherwise.
func (o *Group) GetLocal() bool {
	if o == nil || o.Local == nil {
		var ret bool
		return ret
	}
	return *o.Local
}

// GetLocalOk returns a tuple with the Local field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetLocalOk() (*bool, bool) {
	if o == nil || o.Local == nil {
		return nil, false
	}
	return o.Local, true
}

// HasLocal returns a boolean if a field has been set.
func (o *Group) HasLocal() bool {
	if o != nil && o.Local != nil {
		return true
	}

	return false
}

// SetLocal gets a reference to the given bool and assigns it to the Local field.
func (o *Group) SetLocal(v bool) {
	o.Local = &v
}

// GetIdTypeBoth returns the IdTypeBoth field value if set, zero value otherwise.
func (o *Group) GetIdTypeBoth() bool {
	if o == nil || o.IdTypeBoth == nil {
		var ret bool
		return ret
	}
	return *o.IdTypeBoth
}

// GetIdTypeBothOk returns a tuple with the IdTypeBoth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetIdTypeBothOk() (*bool, bool) {
	if o == nil || o.IdTypeBoth == nil {
		return nil, false
	}
	return o.IdTypeBoth, true
}

// HasIdTypeBoth returns a boolean if a field has been set.
func (o *Group) HasIdTypeBoth() bool {
	if o != nil && o.IdTypeBoth != nil {
		return true
	}

	return false
}

// SetIdTypeBoth gets a reference to the given bool and assigns it to the IdTypeBoth field.
func (o *Group) SetIdTypeBoth(v bool) {
	o.IdTypeBoth = &v
}

func (o Group) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.Gid != nil {
		toSerialize["gid"] = o.Gid
	}
	if true {
		toSerialize["group"] = o.Group
	}
	if o.Builtin != nil {
		toSerialize["builtin"] = o.Builtin
	}
	if o.Sudo != nil {
		toSerialize["sudo"] = o.Sudo
	}
	if o.SudoNopasswd != nil {
		toSerialize["sudo_nopasswd"] = o.SudoNopasswd
	}
	if o.SudoCommands != nil {
		toSerialize["sudo_commands"] = o.SudoCommands
	}
	if o.Smb != nil {
		toSerialize["smb"] = o.Smb
	}
	if o.Users != nil {
		toSerialize["users"] = o.Users
	}
	if o.Local != nil {
		toSerialize["local"] = o.Local
	}
	if o.IdTypeBoth != nil {
		toSerialize["id_type_both"] = o.IdTypeBoth
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Group) UnmarshalJSON(bytes []byte) (err error) {
	varGroup := _Group{}

	if err = json.Unmarshal(bytes, &varGroup); err == nil {
		*o = Group(varGroup)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "gid")
		delete(additionalProperties, "group")
		delete(additionalProperties, "builtin")
		delete(additionalProperties, "sudo")
		delete(additionalProperties, "sudo_nopasswd")
		delete(additionalProperties, "sudo_commands")
		delete(additionalProperties, "smb")
		delete(additionalProperties, "users")
		delete(additionalProperties, "local")
		delete(additionalProperties, "id_type_both")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGroup struct {
	value *Group
	isSet bool
}

func (v NullableGroup) Get() *Group {
	return v.value
}

func (v *NullableGroup) Set(val *Group) {
	v.value = val
	v.isSet = true
}

func (v NullableGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroup(val *Group) *NullableGroup {
	return &NullableGroup{value: val, isSet: true}
}

func (v NullableGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}