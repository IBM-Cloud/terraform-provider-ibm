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
	"github.com/IBM/go-sdk-core/core"
)

// DnsSvcsV1 : Manage DNS Permitted Network
//

// ListPermittedNetworks : List permitted networks
// List the permitted networks for a given DNS zone.
func (dnsSvcs *DnsSvcsV1) ListPermittedNetworks(listPermittedNetworksOptions *ListPermittedNetworksOptions) (result *ListPermittedNetworks, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listPermittedNetworksOptions, "listPermittedNetworksOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listPermittedNetworksOptions, "listPermittedNetworksOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"instances", "dnszones", "permitted_networks"}
	pathParameters := []string{*listPermittedNetworksOptions.InstanceID, *listPermittedNetworksOptions.DnszoneID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(dnsSvcs.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range listPermittedNetworksOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("dns_svcs", "V1", "ListPermittedNetworks")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddHeader("Accept", "application/json")
	if listPermittedNetworksOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*listPermittedNetworksOptions.XCorrelationID))
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
		result, err = UnmarshalListPermittedNetworks(m)
		response.Result = result
	}

	return
}

// CreatePermittedNetwork : Create a permitted network
// Create a permitted network for a given DNS zone.
func (dnsSvcs *DnsSvcsV1) CreatePermittedNetwork(createPermittedNetworkOptions *CreatePermittedNetworkOptions) (result *PermittedNetwork, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createPermittedNetworkOptions, "createPermittedNetworkOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createPermittedNetworkOptions, "createPermittedNetworkOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"instances", "dnszones", "permitted_networks"}
	pathParameters := []string{*createPermittedNetworkOptions.InstanceID, *createPermittedNetworkOptions.DnszoneID}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(dnsSvcs.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range createPermittedNetworkOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("dns_svcs", "V1", "CreatePermittedNetwork")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if createPermittedNetworkOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*createPermittedNetworkOptions.XCorrelationID))
	}

	body := make(map[string]interface{})
	if createPermittedNetworkOptions.PermittedNetwork != nil {
		body["permitted_network"] = createPermittedNetworkOptions.PermittedNetwork
	}
	if createPermittedNetworkOptions.Type != nil {
		body["type"] = createPermittedNetworkOptions.Type
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
		result, err = UnmarshalPermittedNetwork(m)
		response.Result = result
	}

	return
}

// DeletePermittedNetwork : Delete a permitted network
// Delete a permitted network.
func (dnsSvcs *DnsSvcsV1) DeletePermittedNetwork(deletePermittedNetworkOptions *DeletePermittedNetworkOptions) (result *PermittedNetwork, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deletePermittedNetworkOptions, "deletePermittedNetworkOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deletePermittedNetworkOptions, "deletePermittedNetworkOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"instances", "dnszones", "permitted_networks"}
	pathParameters := []string{*deletePermittedNetworkOptions.InstanceID, *deletePermittedNetworkOptions.DnszoneID, *deletePermittedNetworkOptions.PermittedNetworkID}

	builder := core.NewRequestBuilder(core.DELETE)
	_, err = builder.ConstructHTTPURL(dnsSvcs.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range deletePermittedNetworkOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("dns_svcs", "V1", "DeletePermittedNetwork")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	if deletePermittedNetworkOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*deletePermittedNetworkOptions.XCorrelationID))
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
		result, err = UnmarshalPermittedNetwork(m)
		response.Result = result
	}

	return
}

// GetPermittedNetwork : Get a permitted network
// Get details of a permitted network.
func (dnsSvcs *DnsSvcsV1) GetPermittedNetwork(getPermittedNetworkOptions *GetPermittedNetworkOptions) (result *PermittedNetwork, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getPermittedNetworkOptions, "getPermittedNetworkOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getPermittedNetworkOptions, "getPermittedNetworkOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"instances", "dnszones", "permitted_networks"}
	pathParameters := []string{*getPermittedNetworkOptions.InstanceID, *getPermittedNetworkOptions.DnszoneID, *getPermittedNetworkOptions.PermittedNetworkID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(dnsSvcs.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getPermittedNetworkOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("dns_svcs", "V1", "GetPermittedNetwork")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddHeader("Accept", "application/json")
	if getPermittedNetworkOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*getPermittedNetworkOptions.XCorrelationID))
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
		result, err = UnmarshalPermittedNetwork(m)
		response.Result = result
	}

	return
}

// CreatePermittedNetworkOptions : The CreatePermittedNetwork options.
type CreatePermittedNetworkOptions struct {
	// The unique identifier of a service instance.
	InstanceID *string `json:"instance_id" validate:"required"`

	// The unique identifier of a DNS zone.
	DnszoneID *string `json:"dnszone_id" validate:"required"`

	// Permitted network data.
	PermittedNetwork *PermittedNetworkVpc `json:"permitted_network" validate:"required"`

	// The type of a permitted network.
	Type *string `json:"type,omitempty"`

	// Uniquely identifying a request.
	XCorrelationID *string `json:"X-Correlation-ID,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the CreatePermittedNetworkOptions.Type property.
// The type of a permitted network.
const (
	CreatePermittedNetworkOptions_Type_Vpc = "vpc"
)

// NewCreatePermittedNetworkOptions : Instantiate CreatePermittedNetworkOptions
func (*DnsSvcsV1) NewCreatePermittedNetworkOptions(instanceID string, dnszoneID string) *CreatePermittedNetworkOptions {
	return &CreatePermittedNetworkOptions{
		InstanceID: core.StringPtr(instanceID),
		DnszoneID:  core.StringPtr(dnszoneID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (options *CreatePermittedNetworkOptions) SetInstanceID(instanceID string) *CreatePermittedNetworkOptions {
	options.InstanceID = core.StringPtr(instanceID)
	return options
}

// SetDnszoneID : Allow user to set DnszoneID
func (options *CreatePermittedNetworkOptions) SetDnszoneID(dnszoneID string) *CreatePermittedNetworkOptions {
	options.DnszoneID = core.StringPtr(dnszoneID)
	return options
}

// SetPermittedNetwork : Allow user to set PermittedNetwork
func (options *CreatePermittedNetworkOptions) SetPermittedNetwork(permittedNetwork *PermittedNetworkVpc) *CreatePermittedNetworkOptions {
	options.PermittedNetwork = permittedNetwork
	return options
}

// SetType : Allow user to set Type
func (options *CreatePermittedNetworkOptions) SetType(typeVar string) *CreatePermittedNetworkOptions {
	options.Type = core.StringPtr(typeVar)
	return options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (options *CreatePermittedNetworkOptions) SetXCorrelationID(xCorrelationID string) *CreatePermittedNetworkOptions {
	options.XCorrelationID = core.StringPtr(xCorrelationID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CreatePermittedNetworkOptions) SetHeaders(param map[string]string) *CreatePermittedNetworkOptions {
	options.Headers = param
	return options
}

// DeletePermittedNetworkOptions : The DeletePermittedNetwork options.
type DeletePermittedNetworkOptions struct {
	// The unique identifier of a service instance.
	InstanceID *string `json:"instance_id" validate:"required"`

	// The unique identifier of a DNS zone.
	DnszoneID *string `json:"dnszone_id" validate:"required"`

	// The unique identifier of a permitted network.
	PermittedNetworkID *string `json:"permitted_network_id" validate:"required"`

	// Uniquely identifying a request.
	XCorrelationID *string `json:"X-Correlation-ID,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeletePermittedNetworkOptions : Instantiate DeletePermittedNetworkOptions
func (*DnsSvcsV1) NewDeletePermittedNetworkOptions(instanceID string, dnszoneID string, permittedNetworkID string) *DeletePermittedNetworkOptions {
	return &DeletePermittedNetworkOptions{
		InstanceID:         core.StringPtr(instanceID),
		DnszoneID:          core.StringPtr(dnszoneID),
		PermittedNetworkID: core.StringPtr(permittedNetworkID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (options *DeletePermittedNetworkOptions) SetInstanceID(instanceID string) *DeletePermittedNetworkOptions {
	options.InstanceID = core.StringPtr(instanceID)
	return options
}

// SetDnszoneID : Allow user to set DnszoneID
func (options *DeletePermittedNetworkOptions) SetDnszoneID(dnszoneID string) *DeletePermittedNetworkOptions {
	options.DnszoneID = core.StringPtr(dnszoneID)
	return options
}

// SetPermittedNetworkID : Allow user to set PermittedNetworkID
func (options *DeletePermittedNetworkOptions) SetPermittedNetworkID(permittedNetworkID string) *DeletePermittedNetworkOptions {
	options.PermittedNetworkID = core.StringPtr(permittedNetworkID)
	return options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (options *DeletePermittedNetworkOptions) SetXCorrelationID(xCorrelationID string) *DeletePermittedNetworkOptions {
	options.XCorrelationID = core.StringPtr(xCorrelationID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeletePermittedNetworkOptions) SetHeaders(param map[string]string) *DeletePermittedNetworkOptions {
	options.Headers = param
	return options
}

// GetPermittedNetworkOptions : The GetPermittedNetwork options.
type GetPermittedNetworkOptions struct {
	// The unique identifier of a service instance.
	InstanceID *string `json:"instance_id" validate:"required"`

	// The unique identifier of a DNS zone.
	DnszoneID *string `json:"dnszone_id" validate:"required"`

	// The unique identifier of a permitted network.
	PermittedNetworkID *string `json:"permitted_network_id" validate:"required"`

	// Uniquely identifying a request.
	XCorrelationID *string `json:"X-Correlation-ID,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetPermittedNetworkOptions : Instantiate GetPermittedNetworkOptions
func (*DnsSvcsV1) NewGetPermittedNetworkOptions(instanceID string, dnszoneID string, permittedNetworkID string) *GetPermittedNetworkOptions {
	return &GetPermittedNetworkOptions{
		InstanceID:         core.StringPtr(instanceID),
		DnszoneID:          core.StringPtr(dnszoneID),
		PermittedNetworkID: core.StringPtr(permittedNetworkID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (options *GetPermittedNetworkOptions) SetInstanceID(instanceID string) *GetPermittedNetworkOptions {
	options.InstanceID = core.StringPtr(instanceID)
	return options
}

// SetDnszoneID : Allow user to set DnszoneID
func (options *GetPermittedNetworkOptions) SetDnszoneID(dnszoneID string) *GetPermittedNetworkOptions {
	options.DnszoneID = core.StringPtr(dnszoneID)
	return options
}

// SetPermittedNetworkID : Allow user to set PermittedNetworkID
func (options *GetPermittedNetworkOptions) SetPermittedNetworkID(permittedNetworkID string) *GetPermittedNetworkOptions {
	options.PermittedNetworkID = core.StringPtr(permittedNetworkID)
	return options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (options *GetPermittedNetworkOptions) SetXCorrelationID(xCorrelationID string) *GetPermittedNetworkOptions {
	options.XCorrelationID = core.StringPtr(xCorrelationID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetPermittedNetworkOptions) SetHeaders(param map[string]string) *GetPermittedNetworkOptions {
	options.Headers = param
	return options
}

// ListPermittedNetworksOptions : The ListPermittedNetworks options.
type ListPermittedNetworksOptions struct {
	// The unique identifier of a service instance.
	InstanceID *string `json:"instance_id" validate:"required"`

	// The unique identifier of a DNS zone.
	DnszoneID *string `json:"dnszone_id" validate:"required"`

	// Uniquely identifying a request.
	XCorrelationID *string `json:"X-Correlation-ID,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListPermittedNetworksOptions : Instantiate ListPermittedNetworksOptions
func (*DnsSvcsV1) NewListPermittedNetworksOptions(instanceID string, dnszoneID string) *ListPermittedNetworksOptions {
	return &ListPermittedNetworksOptions{
		InstanceID: core.StringPtr(instanceID),
		DnszoneID:  core.StringPtr(dnszoneID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (options *ListPermittedNetworksOptions) SetInstanceID(instanceID string) *ListPermittedNetworksOptions {
	options.InstanceID = core.StringPtr(instanceID)
	return options
}

// SetDnszoneID : Allow user to set DnszoneID
func (options *ListPermittedNetworksOptions) SetDnszoneID(dnszoneID string) *ListPermittedNetworksOptions {
	options.DnszoneID = core.StringPtr(dnszoneID)
	return options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (options *ListPermittedNetworksOptions) SetXCorrelationID(xCorrelationID string) *ListPermittedNetworksOptions {
	options.XCorrelationID = core.StringPtr(xCorrelationID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListPermittedNetworksOptions) SetHeaders(param map[string]string) *ListPermittedNetworksOptions {
	options.Headers = param
	return options
}

// ListPermittedNetworks : List permitted networks response.
type ListPermittedNetworks struct {
	// An array of permitted networks.
	PermittedNetworks []PermittedNetwork `json:"permitted_networks" validate:"required"`

	// Number of permitted networks.
	Count *int64 `json:"count" validate:"required"`

	// href.
	First *FirstHref `json:"first" validate:"required"`

	// Number of permitted networks per page.
	Limit *int64 `json:"limit" validate:"required"`

	// href.
	Next *NextHref `json:"next" validate:"required"`

	// Page number.
	Offset *int64 `json:"offset" validate:"required"`

	// Total number of permitted networks.
	TotalCount *int64 `json:"total_count" validate:"required"`
}

// UnmarshalListPermittedNetworks constructs an instance of ListPermittedNetworks from the specified map.
func UnmarshalListPermittedNetworks(m map[string]interface{}) (result *ListPermittedNetworks, err error) {
	obj := new(ListPermittedNetworks)
	obj.PermittedNetworks, err = UnmarshalPermittedNetworkSliceAsProperty(m, "permitted_networks")
	if err != nil {
		return
	}
	obj.Count, err = core.UnmarshalInt64(m, "count")
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

// PermittedNetwork : Permitted network details.
type PermittedNetwork struct {
	// Permitted network data.
	PermittedNetwork *PermittedNetworkVpc `json:"permitted_network,omitempty"`

	// The time when a permitted network is created.
	CreatedOn *string `json:"created_on,omitempty"`

	// Unique identifier of a permitted network.
	ID *string `json:"id,omitempty"`

	// The recent time when a permitted network is modified.
	ModifiedOn *string `json:"modified_on,omitempty"`

	// The type of a permitted network.
	Type *string `json:"type,omitempty"`

	// The state of a permitted network.
	State *string `json:"state,omitempty"`
}

// Constants associated with the PermittedNetwork.Type property.
// The type of a permitted network.
const (
	PermittedNetwork_Type_Vpc = "vpc"
)

// Constants associated with the PermittedNetwork.State property.
// The state of a permitted network.
const (
	PermittedNetwork_State_Active        = "ACTIVE"
	PermittedNetwork_State_PendingRemove = "REMOVAL_IN_PROGRESS"
)

// UnmarshalPermittedNetwork constructs an instance of PermittedNetwork from the specified map.
func UnmarshalPermittedNetwork(m map[string]interface{}) (result *PermittedNetwork, err error) {
	obj := new(PermittedNetwork)
	obj.PermittedNetwork, err = UnmarshalPermittedNetworkVpcAsProperty(m, "permitted_network")
	if err != nil {
		return
	}
	obj.CreatedOn, err = core.UnmarshalString(m, "created_on")
	if err != nil {
		return
	}
	obj.ID, err = core.UnmarshalString(m, "id")
	if err != nil {
		return
	}
	obj.ModifiedOn, err = core.UnmarshalString(m, "modified_on")
	if err != nil {
		return
	}
	obj.Type, err = core.UnmarshalString(m, "type")
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

// UnmarshalPermittedNetworkSlice unmarshals a slice of PermittedNetwork instances from the specified list of maps.
func UnmarshalPermittedNetworkSlice(s []interface{}) (slice []PermittedNetwork, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'PermittedNetwork'")
			return
		}
		obj, e := UnmarshalPermittedNetwork(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalPermittedNetworkSliceAsProperty unmarshals a slice of PermittedNetwork instances that are stored as a property
// within the specified map.
func UnmarshalPermittedNetworkSliceAsProperty(m map[string]interface{}, propertyName string) (slice []PermittedNetwork, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'PermittedNetwork'", propertyName)
			return
		}
		slice, err = UnmarshalPermittedNetworkSlice(vSlice)
	}
	return
}

// PermittedNetworkVpc : Permitted network data for VPC.
type PermittedNetworkVpc struct {
	// CRN string uniquely identifies a VPC.
	VpcCrn *string `json:"vpc_crn" validate:"required"`
}

// NewPermittedNetworkVpc : Instantiate PermittedNetworkVpc (Generic Model Constructor)
func (*DnsSvcsV1) NewPermittedNetworkVpc(vpcCrn string) (model *PermittedNetworkVpc, err error) {
	model = &PermittedNetworkVpc{
		VpcCrn: core.StringPtr(vpcCrn),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalPermittedNetworkVpc constructs an instance of PermittedNetworkVpc from the specified map.
func UnmarshalPermittedNetworkVpc(m map[string]interface{}) (result *PermittedNetworkVpc, err error) {
	obj := new(PermittedNetworkVpc)
	obj.VpcCrn, err = core.UnmarshalString(m, "vpc_crn")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalPermittedNetworkVpcAsProperty unmarshals an instance of PermittedNetworkVpc that is stored as a property
// within the specified map.
func UnmarshalPermittedNetworkVpcAsProperty(m map[string]interface{}, propertyName string) (result *PermittedNetworkVpc, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'PermittedNetworkVpc'", propertyName)
			return
		}
		result, err = UnmarshalPermittedNetworkVpc(objMap)
	}
	return
}
