/**
 * (C) Copyright IBM Corp. 2020.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Package dnssvcsv1 : Operations and models for the DnsSvcsV1 service
package dnssvcsv1

import (
	"fmt"

	common "github.com/IBM/dns-svcs-go-sdk/common"
	"github.com/IBM/go-sdk-core/v3/core"
)

// DnsSvcsV1 : Manage DNS Zones
//

// ListDnszones : List DNS zones
// List the DNS zones for a given service instance.
func (dnsSvcs *DnsSvcsV1) ListDnszones(listDnszonesOptions *ListDnszonesOptions) (result *ListDnszones, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listDnszonesOptions, "listDnszonesOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listDnszonesOptions, "listDnszonesOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"instances", "dnszones"}
	pathParameters := []string{*listDnszonesOptions.InstanceID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(dnsSvcs.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range listDnszonesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("dns_svcs", "V1", "ListDnszones")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddHeader("Accept", "application/json")
	if listDnszonesOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*listDnszonesOptions.XCorrelationID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = dnsSvcs.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalListDnszones(m)
		response.Result = result
	}

	return
}

// CreateDnszone : Create a DNS zone
// Create a DNS zone for a given service instance.
func (dnsSvcs *DnsSvcsV1) CreateDnszone(createDnszoneOptions *CreateDnszoneOptions) (result *Dnszone, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createDnszoneOptions, "createDnszoneOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createDnszoneOptions, "createDnszoneOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"instances", "dnszones"}
	pathParameters := []string{*createDnszoneOptions.InstanceID}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(dnsSvcs.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range createDnszoneOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("dns_svcs", "V1", "CreateDnszone")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if createDnszoneOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*createDnszoneOptions.XCorrelationID))
	}

	body := make(map[string]interface{})
	if createDnszoneOptions.Description != nil {
		body["description"] = createDnszoneOptions.Description
	}
	if createDnszoneOptions.Label != nil {
		body["label"] = createDnszoneOptions.Label
	}
	if createDnszoneOptions.Name != nil {
		body["name"] = createDnszoneOptions.Name
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = dnsSvcs.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalDnszone(m)
		response.Result = result
	}

	return
}

// DeleteDnszone : Delete a DNS zone
// Delete a DNS zone.
func (dnsSvcs *DnsSvcsV1) DeleteDnszone(deleteDnszoneOptions *DeleteDnszoneOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteDnszoneOptions, "deleteDnszoneOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteDnszoneOptions, "deleteDnszoneOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"instances", "dnszones"}
	pathParameters := []string{*deleteDnszoneOptions.InstanceID, *deleteDnszoneOptions.DnszoneID}

	builder := core.NewRequestBuilder(core.DELETE)
	_, err = builder.ConstructHTTPURL(dnsSvcs.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteDnszoneOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("dns_svcs", "V1", "DeleteDnszone")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	if deleteDnszoneOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*deleteDnszoneOptions.XCorrelationID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = dnsSvcs.Service.Request(request, nil)

	return
}

// GetDnszone : Get a DNS zone
// Get details of a DNS zone.
func (dnsSvcs *DnsSvcsV1) GetDnszone(getDnszoneOptions *GetDnszoneOptions) (result *Dnszone, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getDnszoneOptions, "getDnszoneOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getDnszoneOptions, "getDnszoneOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"instances", "dnszones"}
	pathParameters := []string{*getDnszoneOptions.InstanceID, *getDnszoneOptions.DnszoneID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(dnsSvcs.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getDnszoneOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("dns_svcs", "V1", "GetDnszone")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddHeader("Accept", "application/json")
	if getDnszoneOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*getDnszoneOptions.XCorrelationID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = dnsSvcs.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalDnszone(m)
		response.Result = result
	}

	return
}

// UpdateDnszone : Update the properties of a DNS zone
// Update the properties of a DNS zone.
func (dnsSvcs *DnsSvcsV1) UpdateDnszone(updateDnszoneOptions *UpdateDnszoneOptions) (result *Dnszone, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateDnszoneOptions, "updateDnszoneOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateDnszoneOptions, "updateDnszoneOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"instances", "dnszones"}
	pathParameters := []string{*updateDnszoneOptions.InstanceID, *updateDnszoneOptions.DnszoneID}

	builder := core.NewRequestBuilder(core.PATCH)
	_, err = builder.ConstructHTTPURL(dnsSvcs.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateDnszoneOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("dns_svcs", "V1", "UpdateDnszone")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if updateDnszoneOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*updateDnszoneOptions.XCorrelationID))
	}

	body := make(map[string]interface{})
	if updateDnszoneOptions.Description != nil {
		body["description"] = updateDnszoneOptions.Description
	}
	if updateDnszoneOptions.Label != nil {
		body["label"] = updateDnszoneOptions.Label
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = dnsSvcs.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalDnszone(m)
		response.Result = result
	}

	return
}

// CreateDnszoneOptions : The CreateDnszone options.
type CreateDnszoneOptions struct {
	// The unique identifier of a service instance.
	InstanceID *string `json:"instance_id" validate:"required"`

	// The text describing the purpose of a DNS zone.
	Description *string `json:"description,omitempty"`

	// The label of a DNS zone.
	Label *string `json:"label,omitempty"`

	// Name of DNS zone.
	Name *string `json:"name" validate:"required"`

	// Uniquely identifying a request.
	XCorrelationID *string `json:"X-Correlation-ID,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateDnszoneOptions : Instantiate CreateDnszoneOptions
func (*DnsSvcsV1) NewCreateDnszoneOptions(instanceID, zoneName string) *CreateDnszoneOptions {
	return &CreateDnszoneOptions{
		InstanceID: core.StringPtr(instanceID),
		Name:       core.StringPtr(zoneName),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (options *CreateDnszoneOptions) SetInstanceID(instanceID string) *CreateDnszoneOptions {
	options.InstanceID = core.StringPtr(instanceID)
	return options
}

// SetDescription : Allow user to set Description
func (options *CreateDnszoneOptions) SetDescription(description string) *CreateDnszoneOptions {
	options.Description = core.StringPtr(description)
	return options
}

// SetLabel : Allow user to set Label
func (options *CreateDnszoneOptions) SetLabel(label string) *CreateDnszoneOptions {
	options.Label = core.StringPtr(label)
	return options
}

// SetName : Allow user to set Name
func (options *CreateDnszoneOptions) SetName(name string) *CreateDnszoneOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (options *CreateDnszoneOptions) SetXCorrelationID(xCorrelationID string) *CreateDnszoneOptions {
	options.XCorrelationID = core.StringPtr(xCorrelationID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateDnszoneOptions) SetHeaders(param map[string]string) *CreateDnszoneOptions {
	options.Headers = param
	return options
}

// DeleteDnszoneOptions : The DeleteDnszone options.
type DeleteDnszoneOptions struct {
	// The unique identifier of a service instance.
	InstanceID *string `json:"instance_id" validate:"required"`

	// The unique identifier of a DNS zone.
	DnszoneID *string `json:"dnszone_id" validate:"required"`

	// Uniquely identifying a request.
	XCorrelationID *string `json:"X-Correlation-ID,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteDnszoneOptions : Instantiate DeleteDnszoneOptions
func (*DnsSvcsV1) NewDeleteDnszoneOptions(instanceID string, dnszoneID string) *DeleteDnszoneOptions {
	return &DeleteDnszoneOptions{
		InstanceID: core.StringPtr(instanceID),
		DnszoneID:  core.StringPtr(dnszoneID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (options *DeleteDnszoneOptions) SetInstanceID(instanceID string) *DeleteDnszoneOptions {
	options.InstanceID = core.StringPtr(instanceID)
	return options
}

// SetDnszoneID : Allow user to set DnszoneID
func (options *DeleteDnszoneOptions) SetDnszoneID(dnszoneID string) *DeleteDnszoneOptions {
	options.DnszoneID = core.StringPtr(dnszoneID)
	return options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (options *DeleteDnszoneOptions) SetXCorrelationID(xCorrelationID string) *DeleteDnszoneOptions {
	options.XCorrelationID = core.StringPtr(xCorrelationID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteDnszoneOptions) SetHeaders(param map[string]string) *DeleteDnszoneOptions {
	options.Headers = param
	return options
}

// GetDnszoneOptions : The GetDnszone options.
type GetDnszoneOptions struct {
	// The unique identifier of a service instance.
	InstanceID *string `json:"instance_id" validate:"required"`

	// The unique identifier of a DNS zone.
	DnszoneID *string `json:"dnszone_id" validate:"required"`

	// Uniquely identifying a request.
	XCorrelationID *string `json:"X-Correlation-ID,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetDnszoneOptions : Instantiate GetDnszoneOptions
func (*DnsSvcsV1) NewGetDnszoneOptions(instanceID string, dnszoneID string) *GetDnszoneOptions {
	return &GetDnszoneOptions{
		InstanceID: core.StringPtr(instanceID),
		DnszoneID:  core.StringPtr(dnszoneID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (options *GetDnszoneOptions) SetInstanceID(instanceID string) *GetDnszoneOptions {
	options.InstanceID = core.StringPtr(instanceID)
	return options
}

// SetDnszoneID : Allow user to set DnszoneID
func (options *GetDnszoneOptions) SetDnszoneID(dnszoneID string) *GetDnszoneOptions {
	options.DnszoneID = core.StringPtr(dnszoneID)
	return options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (options *GetDnszoneOptions) SetXCorrelationID(xCorrelationID string) *GetDnszoneOptions {
	options.XCorrelationID = core.StringPtr(xCorrelationID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetDnszoneOptions) SetHeaders(param map[string]string) *GetDnszoneOptions {
	options.Headers = param
	return options
}

// UpdateDnszoneOptions : The UpdateDnszone options.
type UpdateDnszoneOptions struct {
	// The unique identifier of a service instance.
	InstanceID *string `json:"instance_id" validate:"required"`

	// The unique identifier of a DNS zone.
	DnszoneID *string `json:"dnszone_id" validate:"required"`

	// The text describing the purpose of a DNS zone.
	Description *string `json:"description,omitempty"`

	// The label of a DNS zone.
	Label *string `json:"label,omitempty"`

	// Uniquely identifying a request.
	XCorrelationID *string `json:"X-Correlation-ID,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateDnszoneOptions : Instantiate UpdateDnszoneOptions
func (*DnsSvcsV1) NewUpdateDnszoneOptions(instanceID string, dnszoneID string) *UpdateDnszoneOptions {
	return &UpdateDnszoneOptions{
		InstanceID: core.StringPtr(instanceID),
		DnszoneID:  core.StringPtr(dnszoneID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (options *UpdateDnszoneOptions) SetInstanceID(instanceID string) *UpdateDnszoneOptions {
	options.InstanceID = core.StringPtr(instanceID)
	return options
}

// SetDnszoneID : Allow user to set DnszoneID
func (options *UpdateDnszoneOptions) SetDnszoneID(dnszoneID string) *UpdateDnszoneOptions {
	options.DnszoneID = core.StringPtr(dnszoneID)
	return options
}

// SetDescription : Allow user to set Description
func (options *UpdateDnszoneOptions) SetDescription(description string) *UpdateDnszoneOptions {
	options.Description = core.StringPtr(description)
	return options
}

// SetLabel : Allow user to set Label
func (options *UpdateDnszoneOptions) SetLabel(label string) *UpdateDnszoneOptions {
	options.Label = core.StringPtr(label)
	return options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (options *UpdateDnszoneOptions) SetXCorrelationID(xCorrelationID string) *UpdateDnszoneOptions {
	options.XCorrelationID = core.StringPtr(xCorrelationID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateDnszoneOptions) SetHeaders(param map[string]string) *UpdateDnszoneOptions {
	options.Headers = param
	return options
}

// ListDnszonesOptions : The ListDnszones options.
type ListDnszonesOptions struct {
	// The unique identifier of a service instance.
	InstanceID *string `json:"instance_id" validate:"required"`

	// Uniquely identifying a request.
	XCorrelationID *string `json:"X-Correlation-ID,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListDnszonesOptions : Instantiate ListDnszonesOptions
func (*DnsSvcsV1) NewListDnszonesOptions(instanceID string) *ListDnszonesOptions {
	return &ListDnszonesOptions{
		InstanceID: core.StringPtr(instanceID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (options *ListDnszonesOptions) SetInstanceID(instanceID string) *ListDnszonesOptions {
	options.InstanceID = core.StringPtr(instanceID)
	return options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (options *ListDnszonesOptions) SetXCorrelationID(xCorrelationID string) *ListDnszonesOptions {
	options.XCorrelationID = core.StringPtr(xCorrelationID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListDnszonesOptions) SetHeaders(param map[string]string) *ListDnszonesOptions {
	options.Headers = param
	return options
}

// Dnszone : DNS zone details.
type Dnszone struct {
	// the time when a DNS zone is created.
	CreatedOn *string `json:"created_on,omitempty"`

	// The text describing the purpose of a DNS zone.
	Description *string `json:"description,omitempty"`

	// Unique identifier of a DNS zone.
	ID *string `json:"id,omitempty"`

	// Unique identifier of a service instance.
	InstanceID *string `json:"instance_id,omitempty"`

	// The label of a DNS zone.
	Label *string `json:"label,omitempty"`

	// the recent time when a DNS zone is modified.
	ModifiedOn *string `json:"modified_on,omitempty"`

	// Name of DNS zone.
	Name *string `json:"name,omitempty"`

	// State of DNS zone.
	State *string `json:"state,omitempty"`
}

// Constants associated with the Dnszone.State property.
// State of DNS zone.
const (
	Dnszone_State_Active            = "active"
	Dnszone_State_Deleted           = "deleted"
	Dnszone_State_Disabled          = "disabled"
	Dnszone_State_PendingDelete     = "pending_delete"
	Dnszone_State_PendingNetworkAdd = "pending_network_add"
)

// UnmarshalDnszone constructs an instance of Dnszone from the specified map.
func UnmarshalDnszone(m map[string]interface{}) (result *Dnszone, err error) {
	obj := new(Dnszone)
	obj.CreatedOn, err = core.UnmarshalString(m, "created_on")
	if err != nil {
		return
	}
	obj.Description, err = core.UnmarshalString(m, "description")
	if err != nil {
		return
	}
	obj.ID, err = core.UnmarshalString(m, "id")
	if err != nil {
		return
	}
	obj.InstanceID, err = core.UnmarshalString(m, "instance_id")
	if err != nil {
		return
	}
	obj.Label, err = core.UnmarshalString(m, "label")
	if err != nil {
		return
	}
	obj.ModifiedOn, err = core.UnmarshalString(m, "modified_on")
	if err != nil {
		return
	}
	obj.Name, err = core.UnmarshalString(m, "name")
	if err != nil {
		return
	}
	obj.State, err = core.UnmarshalString(m, "state")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalDnszoneSlice unmarshals a slice of Dnszone instances from the specified list of maps.
func UnmarshalDnszoneSlice(s []interface{}) (slice []Dnszone, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'Dnszone'")
			return
		}
		obj, e := UnmarshalDnszone(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalDnszoneSliceAsProperty unmarshals a slice of Dnszone instances that are stored as a property
// within the specified map.
func UnmarshalDnszoneSliceAsProperty(m map[string]interface{}, propertyName string) (slice []Dnszone, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'Dnszone'", propertyName)
			return
		}
		slice, err = UnmarshalDnszoneSlice(vSlice)
	}
	return
}

// FirstHref : href.
type FirstHref struct {
	// href.
	Href *string `json:"href,omitempty"`
}

// UnmarshalFirstHref constructs an instance of FirstHref from the specified map.
func UnmarshalFirstHref(m map[string]interface{}) (result *FirstHref, err error) {
	obj := new(FirstHref)
	obj.Href, err = core.UnmarshalString(m, "href")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalFirstHrefAsProperty unmarshals an instance of FirstHref that is stored as a property
// within the specified map.
func UnmarshalFirstHrefAsProperty(m map[string]interface{}, propertyName string) (result *FirstHref, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'FirstHref'", propertyName)
			return
		}
		result, err = UnmarshalFirstHref(objMap)
	}
	return
}

// ListDnszones : List DNS zones response.
type ListDnszones struct {
	// Number of DNS zones.
	Count *int64 `json:"count" validate:"required"`

	// An array of DNS zones.
	Dnszones []Dnszone `json:"dnszones" validate:"required"`

	// href.
	First *FirstHref `json:"first" validate:"required"`

	// Number of DNS zones per page.
	Limit *int64 `json:"limit" validate:"required"`

	// href.
	Next *NextHref `json:"next" validate:"required"`

	// Page number.
	Offset *int64 `json:"offset" validate:"required"`

	// Total number of DNS zones.
	TotalCount *int64 `json:"total_count" validate:"required"`
}

// UnmarshalListDnszones constructs an instance of ListDnszones from the specified map.
func UnmarshalListDnszones(m map[string]interface{}) (result *ListDnszones, err error) {
	obj := new(ListDnszones)
	obj.Count, err = core.UnmarshalInt64(m, "count")
	if err != nil {
		return
	}
	obj.Dnszones, err = UnmarshalDnszoneSliceAsProperty(m, "dnszones")
	if err != nil {
		return
	}
	obj.First, err = UnmarshalFirstHrefAsProperty(m, "first")
	if err != nil {
		return
	}
	obj.Limit, err = core.UnmarshalInt64(m, "limit")
	if err != nil {
		return
	}
	obj.Next, err = UnmarshalNextHrefAsProperty(m, "next")
	if err != nil {
		return
	}
	obj.Offset, err = core.UnmarshalInt64(m, "offset")
	if err != nil {
		return
	}
	obj.TotalCount, err = core.UnmarshalInt64(m, "total_count")
	if err != nil {
		return
	}
	result = obj
	return
}
