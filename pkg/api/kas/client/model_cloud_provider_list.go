/*
 * Kafka Service Fleet Manager
 *
 * Kafka Service Fleet Manager is a Rest API to manage kafka instances and connectors.
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kasclient

import (
	"encoding/json"
)

// CloudProviderList struct for CloudProviderList
type CloudProviderList struct {
	Kind  string          `json:"kind"`
	Page  int32           `json:"page"`
	Size  int32           `json:"size"`
	Total int32           `json:"total"`
	Items []CloudProvider `json:"items"`
}

// NewCloudProviderList instantiates a new CloudProviderList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudProviderList(kind string, page int32, size int32, total int32, items []CloudProvider) *CloudProviderList {
	this := CloudProviderList{}
	this.Kind = kind
	this.Page = page
	this.Size = size
	this.Total = total
	this.Items = items
	return &this
}

// NewCloudProviderListWithDefaults instantiates a new CloudProviderList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudProviderListWithDefaults() *CloudProviderList {
	this := CloudProviderList{}
	return &this
}

// GetKind returns the Kind field value
func (o *CloudProviderList) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *CloudProviderList) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *CloudProviderList) SetKind(v string) {
	o.Kind = v
}

// GetPage returns the Page field value
func (o *CloudProviderList) GetPage() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Page
}

// GetPageOk returns a tuple with the Page field value
// and a boolean to check if the value has been set.
func (o *CloudProviderList) GetPageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Page, true
}

// SetPage sets field value
func (o *CloudProviderList) SetPage(v int32) {
	o.Page = v
}

// GetSize returns the Size field value
func (o *CloudProviderList) GetSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *CloudProviderList) GetSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *CloudProviderList) SetSize(v int32) {
	o.Size = v
}

// GetTotal returns the Total field value
func (o *CloudProviderList) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *CloudProviderList) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *CloudProviderList) SetTotal(v int32) {
	o.Total = v
}

// GetItems returns the Items field value
func (o *CloudProviderList) GetItems() []CloudProvider {
	if o == nil {
		var ret []CloudProvider
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *CloudProviderList) GetItemsOk() (*[]CloudProvider, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Items, true
}

// SetItems sets field value
func (o *CloudProviderList) SetItems(v []CloudProvider) {
	o.Items = v
}

func (o CloudProviderList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["kind"] = o.Kind
	}
	if true {
		toSerialize["page"] = o.Page
	}
	if true {
		toSerialize["size"] = o.Size
	}
	if true {
		toSerialize["total"] = o.Total
	}
	if true {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableCloudProviderList struct {
	value *CloudProviderList
	isSet bool
}

func (v NullableCloudProviderList) Get() *CloudProviderList {
	return v.value
}

func (v *NullableCloudProviderList) Set(val *CloudProviderList) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudProviderList) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudProviderList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudProviderList(val *CloudProviderList) *NullableCloudProviderList {
	return &NullableCloudProviderList{value: val, isSet: true}
}

func (v NullableCloudProviderList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudProviderList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
