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

// Package ibmcloudfunctionsnamespaceapiv1 : Operations and models for the ibmcloudfunctionsnamespaceapiv1 service
package ibmcloudfunctionsnamespaceapiv1

import (
	"fmt"

	"github.com/IBM/go-sdk-core/v3/core"
	"github.com/go-openapi/strfmt"
	common "github.ibm.com/ibmcloud/namespace-go-sdk/common"
)

// const ..
const (
	NamespaceTypeCFBased     = 1
	NamespaceTypeIamMigrated = 2
	NamespaceTypeIamBased    = 3
)

// IbmCloudFunctionsNamespaceAPIV1 : The purpose is to provide an API to manage IBM Cloud Functions namespaces
//
// Version: 1.0
type IbmCloudFunctionsNamespaceAPIV1 struct {
	Service *core.BaseService
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://gateway.watsonplatform.net/servicebroker/API/v1"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "ibm_cloud_functions_namespace_API"

// IbmCloudFunctionsNamespaceOptions : Service options
type IbmCloudFunctionsNamespaceOptions struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
}

// NewIbmCloudFunctionsNamespaceAPIV1UsingExternalConfig : constructs an instance of IbmCloudFunctionsNamespaceAPIV1 with passed in options and external configuration.
func NewIbmCloudFunctionsNamespaceAPIV1UsingExternalConfig(options *IbmCloudFunctionsNamespaceOptions) (ibmCloudFunctionsNamespaceAPI *IbmCloudFunctionsNamespaceAPIV1, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	ibmCloudFunctionsNamespaceAPI, err = NewIbmCloudFunctionsNamespaceAPIV1(options)
	if err != nil {
		return
	}

	err = ibmCloudFunctionsNamespaceAPI.Service.ConfigureService(options.ServiceName)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = ibmCloudFunctionsNamespaceAPI.Service.SetServiceURL(options.URL)
	}
	return
}

// NewIbmCloudFunctionsNamespaceAPIV1 : constructs an instance of IbmCloudFunctionsNamespaceAPIV1 with passed in options.
func NewIbmCloudFunctionsNamespaceAPIV1(options *IbmCloudFunctionsNamespaceOptions) (service *IbmCloudFunctionsNamespaceAPIV1, err error) {
	serviceOptions := &core.ServiceOptions{
		URL:           DefaultServiceURL,
		Authenticator: options.Authenticator,
	}

	baseService, err := core.NewBaseService(serviceOptions)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = baseService.SetServiceURL(options.URL)
		if err != nil {
			return
		}
	}

	service = &IbmCloudFunctionsNamespaceAPIV1{
		Service: baseService,
	}

	return
}

// SetServiceURL sets the service URL
func (ibmCloudFunctionsNamespaceAPI *IbmCloudFunctionsNamespaceAPIV1) SetServiceURL(url string) error {
	return ibmCloudFunctionsNamespaceAPI.Service.SetServiceURL(url)
}

// GetNamespaces : Retrieve all IBM Cloud Functions namespaces (classic and IAM)
// Compatibility: If passing basic authorization instead of an IAM access token the classic namespace associated with
// these authorization credentials is returned.
func (ibmCloudFunctionsNamespaceAPI *IbmCloudFunctionsNamespaceAPIV1) GetNamespaces(getNamespacesOptions *GetNamespacesOptions) (result *NamespaceResponseList, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getNamespacesOptions, "getNamespacesOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"api", "v1", "namespaces"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(ibmCloudFunctionsNamespaceAPI.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}
	fmt.Println(builder.URL)
	for headerName, headerValue := range getNamespacesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_cloud_functions_namespace_API", "V1", "GetNamespaces")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddHeader("Accept", "application/json")

	if getNamespacesOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*getNamespacesOptions.Limit))
	}
	if getNamespacesOptions.Offset != nil {
		builder.AddQuery("offset", fmt.Sprint(*getNamespacesOptions.Offset))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = ibmCloudFunctionsNamespaceAPI.Service.Request(request, new(NamespaceResponseList))
	if err == nil {
		var ok bool
		result, ok = response.Result.(*NamespaceResponseList)
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
		}
	}

	return
}

// CreateNamespace : Create an IBM Cloud Functions namespace
// The IBM Cloud artefacts created for the namespace can be inspected with following commands: 1. IBM Cloud Namespace
// service instance <br/>ibmcloud resource service-instances<br/><br/> In addition the following IBM Cloud artefacts
// will be created and MUST NOT BE DELETED 2. IBM Cloud Namespace ServiceID. All actions defined in this namespace will
// run with this ServiceID. <br/>ibmcloud iam service-ids 3. API key for the serviceID: Used to authenticate the
// ServiceID <br/>ibmcloud iam service-API-keys SERVICEID.
func (ibmCloudFunctionsNamespaceAPI *IbmCloudFunctionsNamespaceAPIV1) CreateNamespace(createNamespaceOptions *CreateNamespaceOptions) (result *NamespaceCreateResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createNamespaceOptions, "createNamespaceOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createNamespaceOptions, "createNamespaceOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"namespaces"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(ibmCloudFunctionsNamespaceAPI.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range createNamespaceOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_cloud_functions_namespace_API", "V1", "CreateNamespace")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if createNamespaceOptions.Name != nil {
		body["name"] = createNamespaceOptions.Name
	}
	if createNamespaceOptions.ResourceGroupID != nil {
		body["resource_group_id"] = createNamespaceOptions.ResourceGroupID
	}
	if createNamespaceOptions.ResourcePlanID != nil {
		body["resource_plan_id"] = createNamespaceOptions.ResourcePlanID
	}
	if createNamespaceOptions.Description != nil {
		body["description"] = createNamespaceOptions.Description
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = ibmCloudFunctionsNamespaceAPI.Service.Request(request, new(NamespaceCreateResponse))
	if err == nil {
		var ok bool
		result, ok = response.Result.(*NamespaceCreateResponse)
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
		}
	}

	return
}

// GetNamespace : Retrieve an IBM Cloud Functions namespace
// Retrieve an IBM Cloud Functions namespace.
func (ibmCloudFunctionsNamespaceAPI *IbmCloudFunctionsNamespaceAPIV1) GetNamespace(getNamespaceOptions *GetNamespaceOptions) (result NamespaceResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getNamespaceOptions, "getNamespaceOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getNamespaceOptions, "getNamespaceOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"api", "v1", "namespaces"}
	pathParameters := []string{*getNamespaceOptions.ID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(ibmCloudFunctionsNamespaceAPI.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getNamespaceOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_cloud_functions_namespace_API", "V1", "GetNamespace")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = ibmCloudFunctionsNamespaceAPI.Service.Request(request, new(NamespaceResponse))
	if err == nil {
		var ok bool
		result, ok = response.Result.(NamespaceResponse)
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
		}
	}

	return
}

// DeleteNamespace : Delete an IBM Cloud Functions namespace
// Delete an IBM Cloud Functions namespace.
func (ibmCloudFunctionsNamespaceAPI *IbmCloudFunctionsNamespaceAPIV1) DeleteNamespace(deleteNamespaceOptions *DeleteNamespaceOptions) (result *NamespaceDeleteResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteNamespaceOptions, "deleteNamespaceOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteNamespaceOptions, "deleteNamespaceOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"namespaces"}
	pathParameters := []string{*deleteNamespaceOptions.ID}

	builder := core.NewRequestBuilder(core.DELETE)
	_, err = builder.ConstructHTTPURL(ibmCloudFunctionsNamespaceAPI.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteNamespaceOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_cloud_functions_namespace_API", "V1", "DeleteNamespace")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = ibmCloudFunctionsNamespaceAPI.Service.Request(request, new(NamespaceDeleteResponse))
	if err == nil {
		var ok bool
		result, ok = response.Result.(*NamespaceDeleteResponse)
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
		}
	}

	return
}

// UpdateNamespace : Update an IBM Cloud Functions namespace
// Update an IBM Cloud Functions namespace.
func (ibmCloudFunctionsNamespaceAPI *IbmCloudFunctionsNamespaceAPIV1) UpdateNamespace(updateNamespaceOptions *UpdateNamespaceOptions) (result *NamespaceResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateNamespaceOptions, "updateNamespaceOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateNamespaceOptions, "updateNamespaceOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"namespaces"}
	pathParameters := []string{*updateNamespaceOptions.ID}

	builder := core.NewRequestBuilder(core.PATCH)
	_, err = builder.ConstructHTTPURL(ibmCloudFunctionsNamespaceAPI.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateNamespaceOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_cloud_functions_namespace_API", "V1", "UpdateNamespace")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateNamespaceOptions.Description != nil {
		body["description"] = updateNamespaceOptions.Description
	}
	if updateNamespaceOptions.Name != nil {
		body["name"] = updateNamespaceOptions.Name
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = ibmCloudFunctionsNamespaceAPI.Service.Request(request, new(NamespaceResponse))
	if err == nil {
		var ok bool
		result, ok = response.Result.(*NamespaceResponse)
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
		}
	}

	return
}

// UpdateNamespaceAPIKey : Update an IBM Cloud Functions namespace API key
// If a serviceid is passed this serviceid will be used to generate the API key. The serviceid will get the name of the
// namespace and the description will get the namespace id. Finally the serviceid will be locked.
func (ibmCloudFunctionsNamespaceAPI *IbmCloudFunctionsNamespaceAPIV1) UpdateNamespaceAPIKey(updateNamespaceAPIKeyOptions *UpdateNamespaceAPIKeyOptions) (result *NamespaceResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateNamespaceAPIKeyOptions, "updateNamespaceAPIKeyOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateNamespaceAPIKeyOptions, "updateNamespaceAPIKeyOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"namespaces", "APIkey"}
	pathParameters := []string{*updateNamespaceAPIKeyOptions.ID}

	builder := core.NewRequestBuilder(core.PATCH)
	_, err = builder.ConstructHTTPURL(ibmCloudFunctionsNamespaceAPI.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateNamespaceAPIKeyOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_cloud_functions_namespace_API", "V1", "UpdateNamespaceAPIKey")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateNamespaceAPIKeyOptions.ServiceID != nil {
		body["service_id"] = updateNamespaceAPIKeyOptions.ServiceID
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = ibmCloudFunctionsNamespaceAPI.Service.Request(request, new(NamespaceResponse))
	if err == nil {
		var ok bool
		result, ok = response.Result.(*NamespaceResponse)
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
		}
	}

	return
}

// MigrateNamespace : Migrate a classic namespace and create an IAM enabled IBM Cloud Functions namespace
// A classic namespace can be migrated only once.
func (ibmCloudFunctionsNamespaceAPI *IbmCloudFunctionsNamespaceAPIV1) MigrateNamespace(migrateNamespaceOptions *MigrateNamespaceOptions) (result *NamespaceCreateResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(migrateNamespaceOptions, "migrateNamespaceOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(migrateNamespaceOptions, "migrateNamespaceOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"namespaces", "migrate"}
	pathParameters := []string{*migrateNamespaceOptions.ID}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(ibmCloudFunctionsNamespaceAPI.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range migrateNamespaceOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_cloud_functions_namespace_API", "V1", "MigrateNamespace")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if migrateNamespaceOptions.Name != nil {
		body["name"] = migrateNamespaceOptions.Name
	}
	if migrateNamespaceOptions.ResourceGroupID != nil {
		body["resource_group_id"] = migrateNamespaceOptions.ResourceGroupID
	}
	if migrateNamespaceOptions.ResourcePlanID != nil {
		body["resource_plan_id"] = migrateNamespaceOptions.ResourcePlanID
	}
	if migrateNamespaceOptions.Description != nil {
		body["description"] = migrateNamespaceOptions.Description
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = ibmCloudFunctionsNamespaceAPI.Service.Request(request, new(NamespaceCreateResponse))
	if err == nil {
		var ok bool
		result, ok = response.Result.(*NamespaceCreateResponse)
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
		}
	}

	return
}

// CreateNamespaceOptions : The CreateNamespace options.
type CreateNamespaceOptions struct {

	// Name.
	Name *string `json:"name" validate:"required"`

	// Resourcegroupid of resource group the namespace resource should be placed in. Use 'ibmcloud resource groups' to
	// query your resources groups and their ids.
	ResourceGroupID *string `json:"resource_group_id" validate:"required"`

	// Resourceplanid to use, e.g. 'functions-base-plan'.
	ResourcePlanID *string `json:"resource_plan_id" validate:"required"`

	// Description.
	Description *string `json:"description,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewCreateNamespaceOptions : Instantiate CreateNamespaceOptions
func (ibmCloudFunctionsNamespaceAPI *IbmCloudFunctionsNamespaceAPIV1) NewCreateNamespaceOptions(name string, resourceGroupID string, resourcePlanID string) *CreateNamespaceOptions {
	return &CreateNamespaceOptions{
		Name:            core.StringPtr(name),
		ResourceGroupID: core.StringPtr(resourceGroupID),
		ResourcePlanID:  core.StringPtr(resourcePlanID),
	}
}

// SetName : Allow user to set Name
func (options *CreateNamespaceOptions) SetName(name string) *CreateNamespaceOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetResourceGroupID : Allow user to set ResourceGroupID
func (options *CreateNamespaceOptions) SetResourceGroupID(resourceGroupID string) *CreateNamespaceOptions {
	options.ResourceGroupID = core.StringPtr(resourceGroupID)
	return options
}

// SetResourcePlanID : Allow user to set ResourcePlanID
func (options *CreateNamespaceOptions) SetResourcePlanID(resourcePlanID string) *CreateNamespaceOptions {
	options.ResourcePlanID = core.StringPtr(resourcePlanID)
	return options
}

// SetDescription : Allow user to set Description
func (options *CreateNamespaceOptions) SetDescription(description string) *CreateNamespaceOptions {
	options.Description = core.StringPtr(description)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateNamespaceOptions) SetHeaders(param map[string]string) *CreateNamespaceOptions {
	options.Headers = param
	return options
}

// DeleteNamespaceOptions : The DeleteNamespace options.
type DeleteNamespaceOptions struct {

	// The id of the namespace to delete.
	ID *string `json:"id" validate:"required"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewDeleteNamespaceOptions : Instantiate DeleteNamespaceOptions
func (ibmCloudFunctionsNamespaceAPI *IbmCloudFunctionsNamespaceAPIV1) NewDeleteNamespaceOptions(ID string) *DeleteNamespaceOptions {
	return &DeleteNamespaceOptions{
		ID: core.StringPtr(ID),
	}
}

// SetID : Allow user to set ID
func (options *DeleteNamespaceOptions) SetID(ID string) *DeleteNamespaceOptions {
	options.ID = core.StringPtr(ID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteNamespaceOptions) SetHeaders(param map[string]string) *DeleteNamespaceOptions {
	options.Headers = param
	return options
}

// GetNamespaceOptions : The GetNamespace options.
type GetNamespaceOptions struct {

	// The id of the namespace to retrieve.
	ID *string `json:"id" validate:"required"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewGetNamespaceOptions : Instantiate GetNamespaceOptions
func (ibmCloudFunctionsNamespaceAPI *IbmCloudFunctionsNamespaceAPIV1) NewGetNamespaceOptions(ID string) *GetNamespaceOptions {
	return &GetNamespaceOptions{
		ID: core.StringPtr(ID),
	}
}

// SetID : Allow user to set ID
func (options *GetNamespaceOptions) SetID(ID string) *GetNamespaceOptions {
	options.ID = core.StringPtr(ID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetNamespaceOptions) SetHeaders(param map[string]string) *GetNamespaceOptions {
	options.Headers = param
	return options
}

// GetNamespacesOptions : The GetNamespaces options.
type GetNamespacesOptions struct {

	// The maximum number of namespaces to return. Default 100. Maximum 200.
	Limit *int64 `json:"limit,omitempty"`

	// The number of namespaces to skip. Default 0.
	Offset *int64 `json:"offset,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewGetNamespacesOptions : Instantiate GetNamespacesOptions
func (ibmCloudFunctionsNamespaceAPI *IbmCloudFunctionsNamespaceAPIV1) NewGetNamespacesOptions() *GetNamespacesOptions {
	return &GetNamespacesOptions{}
}

// SetLimit : Allow user to set Limit
func (options *GetNamespacesOptions) SetLimit(limit int64) *GetNamespacesOptions {
	options.Limit = core.Int64Ptr(limit)
	return options
}

// SetOffset : Allow user to set Offset
func (options *GetNamespacesOptions) SetOffset(offset int64) *GetNamespacesOptions {
	options.Offset = core.Int64Ptr(offset)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetNamespacesOptions) SetHeaders(param map[string]string) *GetNamespacesOptions {
	options.Headers = param
	return options
}

// MigrateNamespaceOptions : The MigrateNamespace options.
type MigrateNamespaceOptions struct {

	// The id of the classic namespace to migrate.
	ID *string `json:"id" validate:"required"`

	// Name.
	Name *string `json:"name" validate:"required"`

	// Resourcegroupid of resource group the namespace resource should be placed in. Use 'ibmcloud resource groups' to
	// query your resources groups and their ids.
	ResourceGroupID *string `json:"resource_group_id" validate:"required"`

	// Resourceplanid to use, e.g. 'functions-base-plan'.
	ResourcePlanID *string `json:"resource_plan_id" validate:"required"`

	// Description.
	Description *string `json:"description,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewMigrateNamespaceOptions : Instantiate MigrateNamespaceOptions
func (ibmCloudFunctionsNamespaceAPI *IbmCloudFunctionsNamespaceAPIV1) NewMigrateNamespaceOptions(ID string, name string, resourceGroupID string, resourcePlanID string) *MigrateNamespaceOptions {
	return &MigrateNamespaceOptions{
		ID:              core.StringPtr(ID),
		Name:            core.StringPtr(name),
		ResourceGroupID: core.StringPtr(resourceGroupID),
		ResourcePlanID:  core.StringPtr(resourcePlanID),
	}
}

// SetID : Allow user to set ID
func (options *MigrateNamespaceOptions) SetID(ID string) *MigrateNamespaceOptions {
	options.ID = core.StringPtr(ID)
	return options
}

// SetName : Allow user to set Name
func (options *MigrateNamespaceOptions) SetName(name string) *MigrateNamespaceOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetResourceGroupID : Allow user to set ResourceGroupID
func (options *MigrateNamespaceOptions) SetResourceGroupID(resourceGroupID string) *MigrateNamespaceOptions {
	options.ResourceGroupID = core.StringPtr(resourceGroupID)
	return options
}

// SetResourcePlanID : Allow user to set ResourcePlanID
func (options *MigrateNamespaceOptions) SetResourcePlanID(resourcePlanID string) *MigrateNamespaceOptions {
	options.ResourcePlanID = core.StringPtr(resourcePlanID)
	return options
}

// SetDescription : Allow user to set Description
func (options *MigrateNamespaceOptions) SetDescription(description string) *MigrateNamespaceOptions {
	options.Description = core.StringPtr(description)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *MigrateNamespaceOptions) SetHeaders(param map[string]string) *MigrateNamespaceOptions {
	options.Headers = param
	return options
}

// NamespaceCreateResponse : NamespaceCreateResponse - create/get response.
type NamespaceCreateResponse struct {

	// Time the API key was created.
	APIKeyCreated *strfmt.DateTime `json:"API_key_created" validate:"required"`

	// ID of API key used by the namespace.
	APIKeyID *string `json:"API_key_id" validate:"required"`

	// Cloud Foundry space GUID - present if it was a classic namespace.
	ClassicSpaceguid *string `json:"classic_spaceguid,omitempty"`

	// ClassicType <br/> This attribute will be absent for an IAM namespace, a namespace which is IAM-enabled and not
	// associated with any CF space. <br/> 1 : Classic - A namespace which is associated with a CF space.  <br/> Such
	// namespace is NOT IAM-enabled and can only be used by using the legacy API key ('entitlement key'). <br/> 2 : Classic
	// IAM enabled - A namespace which is associated with a CF space and which is IAM-enabled.  <br/> It accepts IMA token
	// and legacy API key ('entitlement key') for authorization.<br/> 3 : IAM migration complete - A namespace which was/is
	// associated with a CF space, which is IAM-enabled.  <br/> It accepts only an IAM token for authorization.<br/>.
	ClassicType *int64 `json:"classic_type,omitempty"`

	// CRN of namespace.
	Crn *string `json:"crn" validate:"required"`

	// Description.
	Description *string `json:"description" validate:"required"`

	// UUID of namespace.
	ID *string `json:"id" validate:"required"`

	// Location of the resource.
	Location *string `json:"location" validate:"required"`

	// Name.
	Name *string `json:"name" validate:"required"`

	// Resourcegroupid of resource group the namespace resource was placed in.
	ResourceGroupID *string `json:"resource_group_id" validate:"required"`

	// Resourceplanid used.
	ResourcePlanID *string `json:"resource_plan_id" validate:"required"`

	// Serviceid used by the namespace.
	ServiceID *string `json:"service_id" validate:"required"`
}

// NamespaceDeleteResponse : NamespaceDeleteResponse - delete response.
type NamespaceDeleteResponse struct {

	// UUID of deleted namespace.
	ID *string `json:"id" validate:"required"`
}

// NamespaceResponse : NamespaceResponse - create/get response.
type NamespaceResponse struct {

	// Time the API key was activated.
	APIKeyCreated *strfmt.DateTime `json:"API_key_created,omitempty"`

	// ID of API key used by the namespace.
	APIKeyID *string `json:"API_key_id,omitempty"`

	// CF space GUID of classic namespace - present if it is or was a classic namespace.
	ClassicSpaceguid *string `json:"classic_spaceguid,omitempty"`

	// ClassicType <br/> This attribute will be absent for an IAM namespace, a namespace which is IAM-enabled and not
	// associated with any CF space. <br/> 1 : Classic - A namespace which is associated with a CF space.  <br/> Such
	// namespace is NOT IAM-enabled and can only be used by using the legacy API key ('entitlement key'). <br/> 2 : Classic
	// IAM enabled - A namespace which is associated with a CF space and which is IAM-enabled.  <br/> It accepts IMA token
	// and legacy API key ('entitlement key') for authorization.<br/> 3 : IAM migration complete - A namespace which was/is
	// associated with a CF space, which is IAM-enabled.  <br/> It accepts only an IAM token for authorization.<br/>.
	ClassicType *int64 `json:"classic_type,omitempty"`

	// CRN of namespace - absent if namespace is NOT IAM-enabled.
	Crn *string `json:"crn,omitempty"`

	// Description - absent if namespace is NOT IAM-enabled.
	Description *string `json:"description,omitempty"`

	// UUID of namespace.
	ID *string `json:"id" validate:"required"`

	// Location of the resource.
	Location *string `json:"location" validate:"required"`

	// Name - absent if namespace is NOT IAM-enabled.
	Name *string `json:"name,omitempty"`

	// Resourceplanid used - absent if namespace is NOT IAM-enabled.
	ResourcePlanID *string `json:"resource_plan_id,omitempty"`

	// Serviceid used by the namespace - absent if namespace is NOT IAM-enabled.
	ServiceID *string `json:"service_id,omitempty"`

	// Key used by the cf based namespace.
	Key string `json:"key,omitempty"`

	// UUID used by the cf based namespace.
	UUID string `json:"uuid,omitempty"`
}

// NamespaceResponseList : NamespaceResponseList -.
type NamespaceResponseList struct {

	// Maximum number of namespaces to return.
	Limit *int64 `json:"limit" validate:"required"`

	// List of namespaces.
	Namespaces []NamespaceResponse `json:"namespaces" validate:"required"`

	// Number of namespaces to skip.
	Offset *int64 `json:"offset" validate:"required"`

	// Total number of namespaces available.
	TotalCount *int64 `json:"total_count" validate:"required"`
}

//NamespaceResource ..
type NamespaceResource interface {
	GetID() string
	GetLocation() string
	GetName() string
	GetUUID() string
	GetKey() string
	IsIamEnabled() bool
	IsCf() bool
}

//FunctionsNamespaceListResource ..
type FunctionsNamespaceListResource interface {
	GetNamespaces() []NamespaceResource
}

//FunctionsNamespaceResourceList ..
type FunctionsNamespaceResourceList []NamespaceResource

//GetNamespaces ..
func (nsl *NamespaceResponseList) GetNamespaces() []NamespaceResource {
	namespacesList := make([]NamespaceResource, len(nsl.Namespaces))
	for i := range nsl.Namespaces {
		namespacesList[i] = &nsl.Namespaces[i]
	}

	return namespacesList
}

//GetID ..
func (ns *NamespaceResponse) GetID() string {
	return *ns.ID
}

//GetName ..
func (ns *NamespaceResponse) GetName() string {
	// Classic support - if no name included in namespace obj return the ID (classic namespace name)
	if ns.Name != nil {
		return *ns.Name
	}
	return *ns.ID
}

//GetKey ..
func (ns *NamespaceResponse) GetKey() string {
	return ns.Key
}

//GetUUID ..
func (ns *NamespaceResponse) GetUUID() string {
	return ns.UUID
}

//GetLocation ..
func (ns *NamespaceResponse) GetLocation() string {
	return *ns.Location
}

//IsCf ..
func (ns *NamespaceResponse) IsCf() bool {
	var iscf bool = false
	if ns.ClassicType != nil {
		iscf = (*ns.ClassicType == NamespaceTypeCFBased)
	}
	return iscf
}

//IsIamEnabled ..
func (ns *NamespaceResponse) IsIamEnabled() bool {
	// IAM support - classic_type field is not included for new IAM namespaces so always return true if nil
	if ns.ClassicType != nil {
		return (*ns.ClassicType == NamespaceTypeIamMigrated)
	}
	return true
}

// UpdateNamespaceAPIKeyOptions : The UpdateNamespaceAPIKey options.
type UpdateNamespaceAPIKeyOptions struct {

	// The id of the namespace to update the API key.
	ID *string `json:"id" validate:"required"`

	// If passed this serviceid will replace the current used serviceid Note: Serviceid will get the name of the namespace.
	// Description will get the namespace id. Finally the serviceid will locked.
	ServiceID *string `json:"service_id,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewUpdateNamespaceAPIKeyOptions : Instantiate UpdateNamespaceAPIKeyOptions
func (ibmCloudFunctionsNamespaceAPI *IbmCloudFunctionsNamespaceAPIV1) NewUpdateNamespaceAPIKeyOptions(ID string) *UpdateNamespaceAPIKeyOptions {
	return &UpdateNamespaceAPIKeyOptions{
		ID: core.StringPtr(ID),
	}
}

// SetID : Allow user to set ID
func (options *UpdateNamespaceAPIKeyOptions) SetID(ID string) *UpdateNamespaceAPIKeyOptions {
	options.ID = core.StringPtr(ID)
	return options
}

// SetServiceID : Allow user to set ServiceID
func (options *UpdateNamespaceAPIKeyOptions) SetServiceID(serviceID string) *UpdateNamespaceAPIKeyOptions {
	options.ServiceID = core.StringPtr(serviceID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateNamespaceAPIKeyOptions) SetHeaders(param map[string]string) *UpdateNamespaceAPIKeyOptions {
	options.Headers = param
	return options
}

// UpdateNamespaceOptions : The UpdateNamespace options.
type UpdateNamespaceOptions struct {

	// The id of the namespace to update.
	ID *string `json:"id" validate:"required"`

	// New description.
	Description *string `json:"description,omitempty"`

	// New name.
	Name *string `json:"name,omitempty"`

	// Allows users to set headers to be GDPR compliant
	Headers map[string]string
}

// NewUpdateNamespaceOptions : Instantiate UpdateNamespaceOptions
func (ibmCloudFunctionsNamespaceAPI *IbmCloudFunctionsNamespaceAPIV1) NewUpdateNamespaceOptions(ID string) *UpdateNamespaceOptions {
	return &UpdateNamespaceOptions{
		ID: core.StringPtr(ID),
	}
}

// SetID : Allow user to set ID
func (options *UpdateNamespaceOptions) SetID(ID string) *UpdateNamespaceOptions {
	options.ID = core.StringPtr(ID)
	return options
}

// SetDescription : Allow user to set Description
func (options *UpdateNamespaceOptions) SetDescription(description string) *UpdateNamespaceOptions {
	options.Description = core.StringPtr(description)
	return options
}

// SetName : Allow user to set Name
func (options *UpdateNamespaceOptions) SetName(name string) *UpdateNamespaceOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateNamespaceOptions) SetHeaders(param map[string]string) *UpdateNamespaceOptions {
	options.Headers = param
	return options
}
