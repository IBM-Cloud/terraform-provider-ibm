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

/*
 * IBM OpenAPI SDK Code Generator Version: 99-SNAPSHOT-d753183b-20201209-163011
 */

// Package openservicebrokerv1 : Operations and models for the OpenServiceBrokerV1 service
package openservicebrokerv1

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	common "github.com/IBM/platform-services-go-sdk/common"
)

// OpenServiceBrokerV1 : Contribute resources to the IBM Cloud catalog by implementing a `service broker` that conforms
// to the [Open Service Broker API](https://github.com/openservicebrokerapi/servicebroker/blob/master/spec.md) version
// 2.12  specification and provides enablement extensions for integration with IBM Cloud and the Resource Controller
// provisioning model.
//
// Version: 1.4
type OpenServiceBrokerV1 struct {
	Service *core.BaseService
}

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "open_service_broker"

// OpenServiceBrokerV1Options : Service options
type OpenServiceBrokerV1Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
}

// NewOpenServiceBrokerV1UsingExternalConfig : constructs an instance of OpenServiceBrokerV1 with passed in options and external configuration.
func NewOpenServiceBrokerV1UsingExternalConfig(options *OpenServiceBrokerV1Options) (openServiceBroker *OpenServiceBrokerV1, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	openServiceBroker, err = NewOpenServiceBrokerV1(options)
	if err != nil {
		return
	}

	err = openServiceBroker.Service.ConfigureService(options.ServiceName)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = openServiceBroker.Service.SetServiceURL(options.URL)
	}
	return
}

// NewOpenServiceBrokerV1 : constructs an instance of OpenServiceBrokerV1 with passed in options.
func NewOpenServiceBrokerV1(options *OpenServiceBrokerV1Options) (service *OpenServiceBrokerV1, err error) {
	serviceOptions := &core.ServiceOptions{
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

	service = &OpenServiceBrokerV1{
		Service: baseService,
	}

	return
}

// GetServiceURLForRegion returns the service URL to be used for the specified region
func GetServiceURLForRegion(region string) (string, error) {
	return "", fmt.Errorf("service does not support regional URLs")
}

// Clone makes a copy of "openServiceBroker" suitable for processing requests.
func (openServiceBroker *OpenServiceBrokerV1) Clone() *OpenServiceBrokerV1 {
	if core.IsNil(openServiceBroker) {
		return nil
	}
	clone := *openServiceBroker
	clone.Service = openServiceBroker.Service.Clone()
	return &clone
}

// SetServiceURL sets the service URL
func (openServiceBroker *OpenServiceBrokerV1) SetServiceURL(url string) error {
	return openServiceBroker.Service.SetServiceURL(url)
}

// GetServiceURL returns the service URL
func (openServiceBroker *OpenServiceBrokerV1) GetServiceURL() string {
	return openServiceBroker.Service.GetServiceURL()
}

// SetDefaultHeaders sets HTTP headers to be sent in every request
func (openServiceBroker *OpenServiceBrokerV1) SetDefaultHeaders(headers http.Header) {
	openServiceBroker.Service.SetDefaultHeaders(headers)
}

// SetEnableGzipCompression sets the service's EnableGzipCompression field
func (openServiceBroker *OpenServiceBrokerV1) SetEnableGzipCompression(enableGzip bool) {
	openServiceBroker.Service.SetEnableGzipCompression(enableGzip)
}

// GetEnableGzipCompression returns the service's EnableGzipCompression field
func (openServiceBroker *OpenServiceBrokerV1) GetEnableGzipCompression() bool {
	return openServiceBroker.Service.GetEnableGzipCompression()
}

// EnableRetries enables automatic retries for requests invoked for this service instance.
// If either parameter is specified as 0, then a default value is used instead.
func (openServiceBroker *OpenServiceBrokerV1) EnableRetries(maxRetries int, maxRetryInterval time.Duration) {
	openServiceBroker.Service.EnableRetries(maxRetries, maxRetryInterval)
}

// DisableRetries disables automatic retries for requests invoked for this service instance.
func (openServiceBroker *OpenServiceBrokerV1) DisableRetries() {
	openServiceBroker.Service.DisableRetries()
}

// GetServiceInstanceState : Get the current state of the service instance
// Get the current state information associated with the service instance.
//
// As a service provider you need a way to manage provisioned service instances.  If an account comes past due, you may
// need a to disable the service (without deleting it), and when the account is settled re-enable the service.
//
// This endpoint allows both the provider and IBM Cloud to query for the state of a provisioned service instance.  For
// example, IBM Cloud may query the provider to figure out if a given service is disabled or not and present that state
// to the user.
func (openServiceBroker *OpenServiceBrokerV1) GetServiceInstanceState(getServiceInstanceStateOptions *GetServiceInstanceStateOptions) (result *Resp1874644Root, response *core.DetailedResponse, err error) {
	return openServiceBroker.GetServiceInstanceStateWithContext(context.Background(), getServiceInstanceStateOptions)
}

// GetServiceInstanceStateWithContext is an alternate form of the GetServiceInstanceState method which supports a Context parameter
func (openServiceBroker *OpenServiceBrokerV1) GetServiceInstanceStateWithContext(ctx context.Context, getServiceInstanceStateOptions *GetServiceInstanceStateOptions) (result *Resp1874644Root, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getServiceInstanceStateOptions, "getServiceInstanceStateOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getServiceInstanceStateOptions, "getServiceInstanceStateOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *getServiceInstanceStateOptions.InstanceID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = openServiceBroker.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(openServiceBroker.Service.Options.URL, `/bluemix_v1/service_instances/{instance_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getServiceInstanceStateOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("open_service_broker", "V1", "GetServiceInstanceState")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = openServiceBroker.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalResp1874644Root)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// ReplaceServiceInstanceState : Update the state of a provisioned service instance
// Update (disable or enable) the state of a provisioned service instance. As a service provider you need a way to
// manage provisioned service instances. If an account comes past due, you may need a to disable the service (without
// deleting it), and when the account is settled re-enable the service. This endpoint allows the provider to enable or
// disable the state of a provisioned service instance. It is the service provider's responsibility to disable access to
// the service instance when the disable endpoint is invoked and to re-enable that access when the enable endpoint is
// invoked. When your service broker receives an enable / disable request, it should take whatever action is necessary
// to enable / disable (respectively) the service.  Additionally, If a bind request comes in for a disabled service, the
// broker should reject that request with any code other than `204`, and provide a user-facing message in the
// description.
func (openServiceBroker *OpenServiceBrokerV1) ReplaceServiceInstanceState(replaceServiceInstanceStateOptions *ReplaceServiceInstanceStateOptions) (result *Resp2448145Root, response *core.DetailedResponse, err error) {
	return openServiceBroker.ReplaceServiceInstanceStateWithContext(context.Background(), replaceServiceInstanceStateOptions)
}

// ReplaceServiceInstanceStateWithContext is an alternate form of the ReplaceServiceInstanceState method which supports a Context parameter
func (openServiceBroker *OpenServiceBrokerV1) ReplaceServiceInstanceStateWithContext(ctx context.Context, replaceServiceInstanceStateOptions *ReplaceServiceInstanceStateOptions) (result *Resp2448145Root, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(replaceServiceInstanceStateOptions, "replaceServiceInstanceStateOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(replaceServiceInstanceStateOptions, "replaceServiceInstanceStateOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *replaceServiceInstanceStateOptions.InstanceID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = openServiceBroker.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(openServiceBroker.Service.Options.URL, `/bluemix_v1/service_instances/{instance_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range replaceServiceInstanceStateOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("open_service_broker", "V1", "ReplaceServiceInstanceState")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if replaceServiceInstanceStateOptions.Enabled != nil {
		body["enabled"] = replaceServiceInstanceStateOptions.Enabled
	}
	if replaceServiceInstanceStateOptions.InitiatorID != nil {
		body["initiator_id"] = replaceServiceInstanceStateOptions.InitiatorID
	}
	if replaceServiceInstanceStateOptions.ReasonCode != nil {
		body["reason_code"] = replaceServiceInstanceStateOptions.ReasonCode
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = openServiceBroker.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalResp2448145Root)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// ReplaceServiceInstance : Create (provision) a service instance
// Create a service instance with GUID. When your service broker receives a provision request from the IBM Cloud
// platform, it MUST take whatever action is necessary to create a new resource.
//
// When a user creates a service instance from the IBM Cloud console or the IBM Cloud CLI, the IBM Cloud platform
// validates that the user has permission to create the service instance using IBM Cloud IAM. After this validation
// occurs, your service broker's provision endpoint (PUT /v2/resource_instances/:instance_id) will be invoked. When
// provisioning occurs, the IBM Cloud platform provides the following values:
//
// - The IBM Cloud context is included in the context variable
// - The X-Broker-API-Originating-Identity will have the IBM IAM ID of the user that initiated the request
// - The parameters section will include the requested location (and additional parameters required by your service).
func (openServiceBroker *OpenServiceBrokerV1) ReplaceServiceInstance(replaceServiceInstanceOptions *ReplaceServiceInstanceOptions) (result *Resp2079872Root, response *core.DetailedResponse, err error) {
	return openServiceBroker.ReplaceServiceInstanceWithContext(context.Background(), replaceServiceInstanceOptions)
}

// ReplaceServiceInstanceWithContext is an alternate form of the ReplaceServiceInstance method which supports a Context parameter
func (openServiceBroker *OpenServiceBrokerV1) ReplaceServiceInstanceWithContext(ctx context.Context, replaceServiceInstanceOptions *ReplaceServiceInstanceOptions) (result *Resp2079872Root, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(replaceServiceInstanceOptions, "replaceServiceInstanceOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(replaceServiceInstanceOptions, "replaceServiceInstanceOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *replaceServiceInstanceOptions.InstanceID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = openServiceBroker.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(openServiceBroker.Service.Options.URL, `/v2/service_instances/{instance_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range replaceServiceInstanceOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("open_service_broker", "V1", "ReplaceServiceInstance")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	if replaceServiceInstanceOptions.AcceptsIncomplete != nil {
		builder.AddQuery("accepts_incomplete", fmt.Sprint(*replaceServiceInstanceOptions.AcceptsIncomplete))
	}

	body := make(map[string]interface{})
	if replaceServiceInstanceOptions.Context != nil {
		body["context"] = replaceServiceInstanceOptions.Context
	}
	if replaceServiceInstanceOptions.OrganizationGUID != nil {
		body["organization_guid"] = replaceServiceInstanceOptions.OrganizationGUID
	}
	if replaceServiceInstanceOptions.Parameters != nil {
		body["parameters"] = replaceServiceInstanceOptions.Parameters
	}
	if replaceServiceInstanceOptions.PlanID != nil {
		body["plan_id"] = replaceServiceInstanceOptions.PlanID
	}
	if replaceServiceInstanceOptions.ServiceID != nil {
		body["service_id"] = replaceServiceInstanceOptions.ServiceID
	}
	if replaceServiceInstanceOptions.SpaceGUID != nil {
		body["space_guid"] = replaceServiceInstanceOptions.SpaceGUID
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = openServiceBroker.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalResp2079872Root)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// UpdateServiceInstance : Update a service instance
// Patch an instance by GUID. Enabling this endpoint allows your user to change plans and service parameters in a
// provisioned service instance. If your offering supports multiple plans, and you want users to be able to change plans
// for a provisioned instance, you will need to enable the ability for users to update their service instance.
//
// To enable support for the update of the plan, a broker MUST declare support per service by specifying
// `"plan_updateable": true` in your brokers' catalog.json.
func (openServiceBroker *OpenServiceBrokerV1) UpdateServiceInstance(updateServiceInstanceOptions *UpdateServiceInstanceOptions) (result *Resp2079874Root, response *core.DetailedResponse, err error) {
	return openServiceBroker.UpdateServiceInstanceWithContext(context.Background(), updateServiceInstanceOptions)
}

// UpdateServiceInstanceWithContext is an alternate form of the UpdateServiceInstance method which supports a Context parameter
func (openServiceBroker *OpenServiceBrokerV1) UpdateServiceInstanceWithContext(ctx context.Context, updateServiceInstanceOptions *UpdateServiceInstanceOptions) (result *Resp2079874Root, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateServiceInstanceOptions, "updateServiceInstanceOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateServiceInstanceOptions, "updateServiceInstanceOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *updateServiceInstanceOptions.InstanceID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = openServiceBroker.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(openServiceBroker.Service.Options.URL, `/v2/service_instances/{instance_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateServiceInstanceOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("open_service_broker", "V1", "UpdateServiceInstance")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	if updateServiceInstanceOptions.AcceptsIncomplete != nil {
		builder.AddQuery("accepts_incomplete", fmt.Sprint(*updateServiceInstanceOptions.AcceptsIncomplete))
	}

	body := make(map[string]interface{})
	if updateServiceInstanceOptions.Context != nil {
		body["context"] = updateServiceInstanceOptions.Context
	}
	if updateServiceInstanceOptions.Parameters != nil {
		body["parameters"] = updateServiceInstanceOptions.Parameters
	}
	if updateServiceInstanceOptions.PlanID != nil {
		body["plan_id"] = updateServiceInstanceOptions.PlanID
	}
	if updateServiceInstanceOptions.PreviousValues != nil {
		body["previous_values"] = updateServiceInstanceOptions.PreviousValues
	}
	if updateServiceInstanceOptions.ServiceID != nil {
		body["service_id"] = updateServiceInstanceOptions.ServiceID
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = openServiceBroker.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalResp2079874Root)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// DeleteServiceInstance : Delete (deprovision) a service instance
// Delete (deprovision) a service instance by GUID. When a service broker receives a deprovision request from the IBM
// Cloud platform, it MUST delete any resources it created during the provision. Usually this means that all resources
// are immediately reclaimed for future provisions.
func (openServiceBroker *OpenServiceBrokerV1) DeleteServiceInstance(deleteServiceInstanceOptions *DeleteServiceInstanceOptions) (result *Resp2079874Root, response *core.DetailedResponse, err error) {
	return openServiceBroker.DeleteServiceInstanceWithContext(context.Background(), deleteServiceInstanceOptions)
}

// DeleteServiceInstanceWithContext is an alternate form of the DeleteServiceInstance method which supports a Context parameter
func (openServiceBroker *OpenServiceBrokerV1) DeleteServiceInstanceWithContext(ctx context.Context, deleteServiceInstanceOptions *DeleteServiceInstanceOptions) (result *Resp2079874Root, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteServiceInstanceOptions, "deleteServiceInstanceOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteServiceInstanceOptions, "deleteServiceInstanceOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *deleteServiceInstanceOptions.InstanceID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = openServiceBroker.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(openServiceBroker.Service.Options.URL, `/v2/service_instances/{instance_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteServiceInstanceOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("open_service_broker", "V1", "DeleteServiceInstance")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("service_id", fmt.Sprint(*deleteServiceInstanceOptions.ServiceID))
	builder.AddQuery("plan_id", fmt.Sprint(*deleteServiceInstanceOptions.PlanID))
	if deleteServiceInstanceOptions.AcceptsIncomplete != nil {
		builder.AddQuery("accepts_incomplete", fmt.Sprint(*deleteServiceInstanceOptions.AcceptsIncomplete))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = openServiceBroker.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalResp2079874Root)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// ListCatalog : Get the catalog metadata stored within the broker
// This endpoints defines the contract between the broker and the IBM Cloud platform for the services and plans that the
// broker supports. This endpoint returns the catalog metadata stored within your broker. These values define the
// minimal provisioning contract between your service and the IBM Cloud platform. All additional catalog metadata that
// is not required for provisioning is stored within the IBM Cloud catalog, and any updates to catalog display values
// that are used to render your dashboard like links, icons, and i18n translated metadata should be updated in the
// Resource Management Console (RMC), and not housed in your broker. None of metadata stored in your broker is displayed
// in the IBM Cloud console or the IBM Cloud CLI; the console and CLI will return what was set withn RMC and stored in
// the IBM Cloud catalog.
func (openServiceBroker *OpenServiceBrokerV1) ListCatalog(listCatalogOptions *ListCatalogOptions) (result *Resp1874650Root, response *core.DetailedResponse, err error) {
	return openServiceBroker.ListCatalogWithContext(context.Background(), listCatalogOptions)
}

// ListCatalogWithContext is an alternate form of the ListCatalog method which supports a Context parameter
func (openServiceBroker *OpenServiceBrokerV1) ListCatalogWithContext(ctx context.Context, listCatalogOptions *ListCatalogOptions) (result *Resp1874650Root, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listCatalogOptions, "listCatalogOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = openServiceBroker.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(openServiceBroker.Service.Options.URL, `/v2/catalog`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range listCatalogOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("open_service_broker", "V1", "ListCatalog")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = openServiceBroker.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalResp1874650Root)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetLastOperation : Get the current status of a provision in-progress for a service instance
// Get `last_operation` for instance by GUID (for asynchronous provision calls). When a broker returns status code `202
// Accepted` during a provision, update, or deprovision call, the IBM Cloud platform will begin polling the
// `last_operation` endpoint to obtain the state of the last requested operation. The broker response MUST contain the
// field `state` and MAY contain the field `description`.
//
// Valid values for `state` are `in progress`, `succeeded`, and `failed`. The platform will poll the `last_operation
// `endpoint as long as the broker returns "state": "in progress". Returning "state": "succeeded" or "state": "failed"
// will cause the platform to cease polling. The value provided for description will be passed through to the platform
// API client and can be used to provide additional detail for users about the progress of the operation.
func (openServiceBroker *OpenServiceBrokerV1) GetLastOperation(getLastOperationOptions *GetLastOperationOptions) (result *Resp2079894Root, response *core.DetailedResponse, err error) {
	return openServiceBroker.GetLastOperationWithContext(context.Background(), getLastOperationOptions)
}

// GetLastOperationWithContext is an alternate form of the GetLastOperation method which supports a Context parameter
func (openServiceBroker *OpenServiceBrokerV1) GetLastOperationWithContext(ctx context.Context, getLastOperationOptions *GetLastOperationOptions) (result *Resp2079894Root, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getLastOperationOptions, "getLastOperationOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getLastOperationOptions, "getLastOperationOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *getLastOperationOptions.InstanceID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = openServiceBroker.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(openServiceBroker.Service.Options.URL, `/v2/service_instances/{instance_id}/last_operation`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getLastOperationOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("open_service_broker", "V1", "GetLastOperation")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if getLastOperationOptions.Operation != nil {
		builder.AddQuery("operation", fmt.Sprint(*getLastOperationOptions.Operation))
	}
	if getLastOperationOptions.PlanID != nil {
		builder.AddQuery("plan_id", fmt.Sprint(*getLastOperationOptions.PlanID))
	}
	if getLastOperationOptions.ServiceID != nil {
		builder.AddQuery("service_id", fmt.Sprint(*getLastOperationOptions.ServiceID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = openServiceBroker.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalResp2079894Root)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// ReplaceServiceBinding : Bind a service instance to another resource
// Create binding by GUID on service instance.
//
// If your service can be bound to applications in IBM Cloud, `bindable:true` must be specified in the catalog.json of
// your service broker. If bindable, it must be able to return API endpoints and credentials to your service consumers.
//
// **Note:** Brokers that do not offer any bindable services do not need to implement the endpoint for bind requests.
//
// See the OSB 2.12 spec for more details on
// [binding](https://github.com/openservicebrokerapi/servicebroker/blob/v2.12/spec.md#binding).
func (openServiceBroker *OpenServiceBrokerV1) ReplaceServiceBinding(replaceServiceBindingOptions *ReplaceServiceBindingOptions) (result *Resp2079876Root, response *core.DetailedResponse, err error) {
	return openServiceBroker.ReplaceServiceBindingWithContext(context.Background(), replaceServiceBindingOptions)
}

// ReplaceServiceBindingWithContext is an alternate form of the ReplaceServiceBinding method which supports a Context parameter
func (openServiceBroker *OpenServiceBrokerV1) ReplaceServiceBindingWithContext(ctx context.Context, replaceServiceBindingOptions *ReplaceServiceBindingOptions) (result *Resp2079876Root, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(replaceServiceBindingOptions, "replaceServiceBindingOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(replaceServiceBindingOptions, "replaceServiceBindingOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"binding_id":  *replaceServiceBindingOptions.BindingID,
		"instance_id": *replaceServiceBindingOptions.InstanceID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = openServiceBroker.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(openServiceBroker.Service.Options.URL, `/v2/service_instances/{instance_id}/service_bindings/{binding_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range replaceServiceBindingOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("open_service_broker", "V1", "ReplaceServiceBinding")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if replaceServiceBindingOptions.BindResource != nil {
		body["bind_resource"] = replaceServiceBindingOptions.BindResource
	}
	if replaceServiceBindingOptions.Parameters != nil {
		body["parameters"] = replaceServiceBindingOptions.Parameters
	}
	if replaceServiceBindingOptions.PlanID != nil {
		body["plan_id"] = replaceServiceBindingOptions.PlanID
	}
	if replaceServiceBindingOptions.ServiceID != nil {
		body["service_id"] = replaceServiceBindingOptions.ServiceID
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = openServiceBroker.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalResp2079876Root)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// DeleteServiceBinding : Delete (unbind) the credentials bound to a resource
// Delete instance binding by GUID.
//
// When a broker receives an unbind request from the IBM Cloud platform, it MUST delete any resources associated with
// the binding. In the case where credentials were generated, this might result in requests to the service instance
// failing to authenticate.
//
// **Note**: Brokers that do not provide any bindable services or plans do not need to implement this endpoint.
func (openServiceBroker *OpenServiceBrokerV1) DeleteServiceBinding(deleteServiceBindingOptions *DeleteServiceBindingOptions) (response *core.DetailedResponse, err error) {
	return openServiceBroker.DeleteServiceBindingWithContext(context.Background(), deleteServiceBindingOptions)
}

// DeleteServiceBindingWithContext is an alternate form of the DeleteServiceBinding method which supports a Context parameter
func (openServiceBroker *OpenServiceBrokerV1) DeleteServiceBindingWithContext(ctx context.Context, deleteServiceBindingOptions *DeleteServiceBindingOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteServiceBindingOptions, "deleteServiceBindingOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteServiceBindingOptions, "deleteServiceBindingOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"binding_id":  *deleteServiceBindingOptions.BindingID,
		"instance_id": *deleteServiceBindingOptions.InstanceID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = openServiceBroker.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(openServiceBroker.Service.Options.URL, `/v2/service_instances/{instance_id}/service_bindings/{binding_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteServiceBindingOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("open_service_broker", "V1", "DeleteServiceBinding")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("plan_id", fmt.Sprint(*deleteServiceBindingOptions.PlanID))
	builder.AddQuery("service_id", fmt.Sprint(*deleteServiceBindingOptions.ServiceID))

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = openServiceBroker.Service.Request(request, nil)

	return
}

// DeleteServiceBindingOptions : The DeleteServiceBinding options.
type DeleteServiceBindingOptions struct {
	// The `binding_id` is the ID of a previously provisioned binding for that service instance.
	BindingID *string `json:"binding_id" validate:"required,ne="`

	// The `instance_id` is the ID of a previously provisioned service instance.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// The ID of the plan from the catalog.json in the broker. It MUST be a non-empty string and should be a GUID.
	PlanID *string `json:"plan_id" validate:"required"`

	// The ID of the service from the catalog.json in the broker. It MUST be a non-empty string and should be a GUID.
	ServiceID *string `json:"service_id" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteServiceBindingOptions : Instantiate DeleteServiceBindingOptions
func (*OpenServiceBrokerV1) NewDeleteServiceBindingOptions(bindingID string, instanceID string, planID string, serviceID string) *DeleteServiceBindingOptions {
	return &DeleteServiceBindingOptions{
		BindingID:  core.StringPtr(bindingID),
		InstanceID: core.StringPtr(instanceID),
		PlanID:     core.StringPtr(planID),
		ServiceID:  core.StringPtr(serviceID),
	}
}

// SetBindingID : Allow user to set BindingID
func (options *DeleteServiceBindingOptions) SetBindingID(bindingID string) *DeleteServiceBindingOptions {
	options.BindingID = core.StringPtr(bindingID)
	return options
}

// SetInstanceID : Allow user to set InstanceID
func (options *DeleteServiceBindingOptions) SetInstanceID(instanceID string) *DeleteServiceBindingOptions {
	options.InstanceID = core.StringPtr(instanceID)
	return options
}

// SetPlanID : Allow user to set PlanID
func (options *DeleteServiceBindingOptions) SetPlanID(planID string) *DeleteServiceBindingOptions {
	options.PlanID = core.StringPtr(planID)
	return options
}

// SetServiceID : Allow user to set ServiceID
func (options *DeleteServiceBindingOptions) SetServiceID(serviceID string) *DeleteServiceBindingOptions {
	options.ServiceID = core.StringPtr(serviceID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteServiceBindingOptions) SetHeaders(param map[string]string) *DeleteServiceBindingOptions {
	options.Headers = param
	return options
}

// DeleteServiceInstanceOptions : The DeleteServiceInstance options.
type DeleteServiceInstanceOptions struct {
	// The ID of the service stored in the catalog.json of your broker. This value should be a GUID. MUST be a non-empty
	// string.
	ServiceID *string `json:"service_id" validate:"required"`

	// The ID of the plan for which the service instance has been requested, which is stored in the catalog.json of your
	// broker. This value should be a GUID. MUST be a non-empty string.
	PlanID *string `json:"plan_id" validate:"required"`

	// The ID of a previously provisioned service instance.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// A value of true indicates that both the IBM Cloud platform and the requesting client support asynchronous
	// deprovisioning. If this parameter is not included in the request, and the broker can only deprovision a service
	// instance of the requested plan asynchronously, the broker MUST reject the request with a `422` Unprocessable Entity.
	AcceptsIncomplete *bool `json:"accepts_incomplete,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteServiceInstanceOptions : Instantiate DeleteServiceInstanceOptions
func (*OpenServiceBrokerV1) NewDeleteServiceInstanceOptions(serviceID string, planID string, instanceID string) *DeleteServiceInstanceOptions {
	return &DeleteServiceInstanceOptions{
		ServiceID:  core.StringPtr(serviceID),
		PlanID:     core.StringPtr(planID),
		InstanceID: core.StringPtr(instanceID),
	}
}

// SetServiceID : Allow user to set ServiceID
func (options *DeleteServiceInstanceOptions) SetServiceID(serviceID string) *DeleteServiceInstanceOptions {
	options.ServiceID = core.StringPtr(serviceID)
	return options
}

// SetPlanID : Allow user to set PlanID
func (options *DeleteServiceInstanceOptions) SetPlanID(planID string) *DeleteServiceInstanceOptions {
	options.PlanID = core.StringPtr(planID)
	return options
}

// SetInstanceID : Allow user to set InstanceID
func (options *DeleteServiceInstanceOptions) SetInstanceID(instanceID string) *DeleteServiceInstanceOptions {
	options.InstanceID = core.StringPtr(instanceID)
	return options
}

// SetAcceptsIncomplete : Allow user to set AcceptsIncomplete
func (options *DeleteServiceInstanceOptions) SetAcceptsIncomplete(acceptsIncomplete bool) *DeleteServiceInstanceOptions {
	options.AcceptsIncomplete = core.BoolPtr(acceptsIncomplete)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteServiceInstanceOptions) SetHeaders(param map[string]string) *DeleteServiceInstanceOptions {
	options.Headers = param
	return options
}

// GetLastOperationOptions : The GetLastOperation options.
type GetLastOperationOptions struct {
	// The unique instance ID generated during provisioning by the IBM Cloud platform.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// A broker-provided identifier for the operation. When a value for operation is included with asynchronous responses
	// for provision and update, and deprovision requests, the IBM Cloud platform will provide the same value using this
	// query parameter as a URL-encoded string. If present, MUST be a non-empty string.
	Operation *string `json:"operation,omitempty"`

	// ID of the plan from the catalog.json in your broker. If present, MUST be a non-empty string.
	PlanID *string `json:"plan_id,omitempty"`

	// ID of the service from the catalog.json in your service broker. If present, MUST be a non-empty string.
	ServiceID *string `json:"service_id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetLastOperationOptions : Instantiate GetLastOperationOptions
func (*OpenServiceBrokerV1) NewGetLastOperationOptions(instanceID string) *GetLastOperationOptions {
	return &GetLastOperationOptions{
		InstanceID: core.StringPtr(instanceID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (options *GetLastOperationOptions) SetInstanceID(instanceID string) *GetLastOperationOptions {
	options.InstanceID = core.StringPtr(instanceID)
	return options
}

// SetOperation : Allow user to set Operation
func (options *GetLastOperationOptions) SetOperation(operation string) *GetLastOperationOptions {
	options.Operation = core.StringPtr(operation)
	return options
}

// SetPlanID : Allow user to set PlanID
func (options *GetLastOperationOptions) SetPlanID(planID string) *GetLastOperationOptions {
	options.PlanID = core.StringPtr(planID)
	return options
}

// SetServiceID : Allow user to set ServiceID
func (options *GetLastOperationOptions) SetServiceID(serviceID string) *GetLastOperationOptions {
	options.ServiceID = core.StringPtr(serviceID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetLastOperationOptions) SetHeaders(param map[string]string) *GetLastOperationOptions {
	options.Headers = param
	return options
}

// GetServiceInstanceStateOptions : The GetServiceInstanceState options.
type GetServiceInstanceStateOptions struct {
	// The `instance_id` of a service instance is provided by the IBM Cloud platform. This ID will be used for future
	// requests to bind and deprovision, so the broker can use it to correlate the resource it creates.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetServiceInstanceStateOptions : Instantiate GetServiceInstanceStateOptions
func (*OpenServiceBrokerV1) NewGetServiceInstanceStateOptions(instanceID string) *GetServiceInstanceStateOptions {
	return &GetServiceInstanceStateOptions{
		InstanceID: core.StringPtr(instanceID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (options *GetServiceInstanceStateOptions) SetInstanceID(instanceID string) *GetServiceInstanceStateOptions {
	options.InstanceID = core.StringPtr(instanceID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetServiceInstanceStateOptions) SetHeaders(param map[string]string) *GetServiceInstanceStateOptions {
	options.Headers = param
	return options
}

// ListCatalogOptions : The ListCatalog options.
type ListCatalogOptions struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListCatalogOptions : Instantiate ListCatalogOptions
func (*OpenServiceBrokerV1) NewListCatalogOptions() *ListCatalogOptions {
	return &ListCatalogOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *ListCatalogOptions) SetHeaders(param map[string]string) *ListCatalogOptions {
	options.Headers = param
	return options
}

// ReplaceServiceBindingOptions : The ReplaceServiceBinding options.
type ReplaceServiceBindingOptions struct {
	// The `binding_id` is provided by the IBM Cloud platform. This ID will be used for future unbind requests, so the
	// broker can use it to correlate the resource it creates.
	BindingID *string `json:"binding_id" validate:"required,ne="`

	// The :`instance_id` is the ID of a previously provisioned service instance.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// A JSON object that contains data for platform resources associated with the binding to be created.
	BindResource *BindResource `json:"bind_resource,omitempty"`

	// Configuration options for the service instance. An opaque object, controller treats this as a blob. Brokers should
	// ensure that the client has provided valid configuration parameters and values for the operation. If this field is
	// not present in the request message, then the broker MUST NOT change the parameters of the instance as a result of
	// this request.
	Parameters map[string]string `json:"parameters,omitempty"`

	// The ID of the plan from the catalog.json in your broker. If present, it MUST be a non-empty string.
	PlanID *string `json:"plan_id,omitempty"`

	// The ID of the service from the catalog.json in your broker. If present, it MUST be a non-empty string.
	ServiceID *string `json:"service_id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewReplaceServiceBindingOptions : Instantiate ReplaceServiceBindingOptions
func (*OpenServiceBrokerV1) NewReplaceServiceBindingOptions(bindingID string, instanceID string) *ReplaceServiceBindingOptions {
	return &ReplaceServiceBindingOptions{
		BindingID:  core.StringPtr(bindingID),
		InstanceID: core.StringPtr(instanceID),
	}
}

// SetBindingID : Allow user to set BindingID
func (options *ReplaceServiceBindingOptions) SetBindingID(bindingID string) *ReplaceServiceBindingOptions {
	options.BindingID = core.StringPtr(bindingID)
	return options
}

// SetInstanceID : Allow user to set InstanceID
func (options *ReplaceServiceBindingOptions) SetInstanceID(instanceID string) *ReplaceServiceBindingOptions {
	options.InstanceID = core.StringPtr(instanceID)
	return options
}

// SetBindResource : Allow user to set BindResource
func (options *ReplaceServiceBindingOptions) SetBindResource(bindResource *BindResource) *ReplaceServiceBindingOptions {
	options.BindResource = bindResource
	return options
}

// SetParameters : Allow user to set Parameters
func (options *ReplaceServiceBindingOptions) SetParameters(parameters map[string]string) *ReplaceServiceBindingOptions {
	options.Parameters = parameters
	return options
}

// SetPlanID : Allow user to set PlanID
func (options *ReplaceServiceBindingOptions) SetPlanID(planID string) *ReplaceServiceBindingOptions {
	options.PlanID = core.StringPtr(planID)
	return options
}

// SetServiceID : Allow user to set ServiceID
func (options *ReplaceServiceBindingOptions) SetServiceID(serviceID string) *ReplaceServiceBindingOptions {
	options.ServiceID = core.StringPtr(serviceID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ReplaceServiceBindingOptions) SetHeaders(param map[string]string) *ReplaceServiceBindingOptions {
	options.Headers = param
	return options
}

// ReplaceServiceInstanceOptions : The ReplaceServiceInstance options.
type ReplaceServiceInstanceOptions struct {
	// The `instance_id` of a service instance is provided by the IBM Cloud platform. This ID will be used for future
	// requests to bind and deprovision, so the broker can use it to correlate the resource it creates.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// Platform specific contextual information under which the service instance is to be provisioned.
	Context *Context `json:"context,omitempty"`

	// Deprecated in favor of `context`. The IBM Cloud platform GUID for the organization under which the service instance
	// is to be provisioned. Although most brokers will not use this field, it might be helpful for executing operations on
	// a user's behalf. It MUST be a non-empty string.
	OrganizationGUID *string `json:"organization_guid,omitempty"`

	// Configuration options for the service instance. An opaque object, controller treats this as a blob. Brokers should
	// ensure that the client has provided valid configuration parameters and values for the operation. If this field is
	// not present in the request message, then the broker MUST NOT change the parameters of the instance as a result of
	// this request.
	Parameters map[string]string `json:"parameters,omitempty"`

	// The ID of the plan for which the service instance has been requested, which is stored in the catalog.json of your
	// broker. This value should be a GUID and it MUST be unique to a service.
	PlanID *string `json:"plan_id,omitempty"`

	// The ID of the service stored in the catalog.json of your broker. This value should be a GUID and it MUST be a
	// non-empty string.
	ServiceID *string `json:"service_id,omitempty"`

	// Deprecated in favor of `context`. The identifier for the project space within the IBM Cloud platform organization.
	// Although most brokers will not use this field, it might be helpful for executing operations on a user's behalf. It
	// MUST be a non-empty string.
	SpaceGUID *string `json:"space_guid,omitempty"`

	// A value of true indicates that both the IBM Cloud platform and the requesting client support asynchronous
	// deprovisioning. If this parameter is not included in the request, and the broker can only deprovision a service
	// instance of the requested plan asynchronously, the broker MUST reject the request with a `422` Unprocessable Entity.
	AcceptsIncomplete *bool `json:"accepts_incomplete,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewReplaceServiceInstanceOptions : Instantiate ReplaceServiceInstanceOptions
func (*OpenServiceBrokerV1) NewReplaceServiceInstanceOptions(instanceID string) *ReplaceServiceInstanceOptions {
	return &ReplaceServiceInstanceOptions{
		InstanceID: core.StringPtr(instanceID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (options *ReplaceServiceInstanceOptions) SetInstanceID(instanceID string) *ReplaceServiceInstanceOptions {
	options.InstanceID = core.StringPtr(instanceID)
	return options
}

// SetContext : Allow user to set Context
func (options *ReplaceServiceInstanceOptions) SetContext(context *Context) *ReplaceServiceInstanceOptions {
	options.Context = context
	return options
}

// SetOrganizationGUID : Allow user to set OrganizationGUID
func (options *ReplaceServiceInstanceOptions) SetOrganizationGUID(organizationGUID string) *ReplaceServiceInstanceOptions {
	options.OrganizationGUID = core.StringPtr(organizationGUID)
	return options
}

// SetParameters : Allow user to set Parameters
func (options *ReplaceServiceInstanceOptions) SetParameters(parameters map[string]string) *ReplaceServiceInstanceOptions {
	options.Parameters = parameters
	return options
}

// SetPlanID : Allow user to set PlanID
func (options *ReplaceServiceInstanceOptions) SetPlanID(planID string) *ReplaceServiceInstanceOptions {
	options.PlanID = core.StringPtr(planID)
	return options
}

// SetServiceID : Allow user to set ServiceID
func (options *ReplaceServiceInstanceOptions) SetServiceID(serviceID string) *ReplaceServiceInstanceOptions {
	options.ServiceID = core.StringPtr(serviceID)
	return options
}

// SetSpaceGUID : Allow user to set SpaceGUID
func (options *ReplaceServiceInstanceOptions) SetSpaceGUID(spaceGUID string) *ReplaceServiceInstanceOptions {
	options.SpaceGUID = core.StringPtr(spaceGUID)
	return options
}

// SetAcceptsIncomplete : Allow user to set AcceptsIncomplete
func (options *ReplaceServiceInstanceOptions) SetAcceptsIncomplete(acceptsIncomplete bool) *ReplaceServiceInstanceOptions {
	options.AcceptsIncomplete = core.BoolPtr(acceptsIncomplete)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ReplaceServiceInstanceOptions) SetHeaders(param map[string]string) *ReplaceServiceInstanceOptions {
	options.Headers = param
	return options
}

// ReplaceServiceInstanceStateOptions : The ReplaceServiceInstanceState options.
type ReplaceServiceInstanceStateOptions struct {
	// The `instance_id` of a service instance is provided by the IBM Cloud platform. This ID will be used for future
	// requests to bind and deprovision, so the broker can use it to correlate the resource it creates.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// Indicates the current state of the service instance.
	Enabled *bool `json:"enabled,omitempty"`

	// Optional string that shows the user ID that is initiating the call.
	InitiatorID *string `json:"initiator_id,omitempty"`

	// Optional string that states the reason code for the service instance state change. Valid values are
	// `IBMCLOUD_ACCT_ACTIVATE`, `IBMCLOUD_RECLAMATION_RESTORE`, or `IBMCLOUD_SERVICE_INSTANCE_BELOW_CAP` for enable calls;
	// `IBMCLOUD_ACCT_SUSPEND`, `IBMCLOUD_RECLAMATION_SCHEDULE`, or `IBMCLOUD_SERVICE_INSTANCE_ABOVE_CAP` for disable
	// calls; and `IBMCLOUD_ADMIN_REQUEST` for enable and disable calls.<br/><br/>Previously accepted values had a `BMX_`
	// prefix, such as `BMX_ACCT_ACTIVATE`. These values are deprecated.
	ReasonCode *string `json:"reason_code,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewReplaceServiceInstanceStateOptions : Instantiate ReplaceServiceInstanceStateOptions
func (*OpenServiceBrokerV1) NewReplaceServiceInstanceStateOptions(instanceID string) *ReplaceServiceInstanceStateOptions {
	return &ReplaceServiceInstanceStateOptions{
		InstanceID: core.StringPtr(instanceID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (options *ReplaceServiceInstanceStateOptions) SetInstanceID(instanceID string) *ReplaceServiceInstanceStateOptions {
	options.InstanceID = core.StringPtr(instanceID)
	return options
}

// SetEnabled : Allow user to set Enabled
func (options *ReplaceServiceInstanceStateOptions) SetEnabled(enabled bool) *ReplaceServiceInstanceStateOptions {
	options.Enabled = core.BoolPtr(enabled)
	return options
}

// SetInitiatorID : Allow user to set InitiatorID
func (options *ReplaceServiceInstanceStateOptions) SetInitiatorID(initiatorID string) *ReplaceServiceInstanceStateOptions {
	options.InitiatorID = core.StringPtr(initiatorID)
	return options
}

// SetReasonCode : Allow user to set ReasonCode
func (options *ReplaceServiceInstanceStateOptions) SetReasonCode(reasonCode string) *ReplaceServiceInstanceStateOptions {
	options.ReasonCode = core.StringPtr(reasonCode)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ReplaceServiceInstanceStateOptions) SetHeaders(param map[string]string) *ReplaceServiceInstanceStateOptions {
	options.Headers = param
	return options
}

// Resp1874644Root : Check the active status of an enabled service.
type Resp1874644Root struct {
	// Indicates (from the viewpoint of the provider) whether the service instance is active and is meaningful if enabled
	// is true. The default value is true if not specified.
	Active *bool `json:"active,omitempty"`

	// Indicates the current state of the service instance.
	Enabled *bool `json:"enabled,omitempty"`

	// Indicates when the service instance was last accessed/modified/etc., and is meaningful if enabled is true AND active
	// is false. Represented as milliseconds since the epoch, but does not need to be accurate to the second/hour.
	LastActive *float64 `json:"last_active,omitempty"`
}

// UnmarshalResp1874644Root unmarshals an instance of Resp1874644Root from the specified map of raw messages.
func UnmarshalResp1874644Root(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Resp1874644Root)
	err = core.UnmarshalPrimitive(m, "active", &obj.Active)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "enabled", &obj.Enabled)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "last_active", &obj.LastActive)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Resp1874650Root : Resp1874650Root struct
type Resp1874650Root struct {
	// List of services.
	Services []Services `json:"services,omitempty"`
}

// UnmarshalResp1874650Root unmarshals an instance of Resp1874650Root from the specified map of raw messages.
func UnmarshalResp1874650Root(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Resp1874650Root)
	err = core.UnmarshalModel(m, "services", &obj.Services, UnmarshalServices)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Resp2079872Root : OK - MUST be returned if the service instance already exists, is fully provisioned, and the requested parameters are
// identical to the existing service instance.
type Resp2079872Root struct {
	// The URL of a web-based management user interface for the service instance; we refer to this as a service dashboard.
	// The URL MUST contain enough information for the dashboard to identify the resource being accessed. Note: a broker
	// that wishes to return `dashboard_url` for a service instance MUST return it with the initial response to the
	// provision request, even if the service is provisioned asynchronously. If present, it MUST be a non-empty string.
	DashboardURL *string `json:"dashboard_url,omitempty"`

	// For asynchronous responses, service brokers MAY return an identifier representing the operation. The value of this
	// field MUST be provided by the platform with requests to the `last_operation` endpoint in a URL encoded query
	// parameter. If present, MUST be a non-empty string.
	Operation *string `json:"operation,omitempty"`
}

// UnmarshalResp2079872Root unmarshals an instance of Resp2079872Root from the specified map of raw messages.
func UnmarshalResp2079872Root(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Resp2079872Root)
	err = core.UnmarshalPrimitive(m, "dashboard_url", &obj.DashboardURL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "operation", &obj.Operation)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Resp2079874Root : Accepted - MUST be returned if the service instance provisioning is in progress. This triggers the IBM Cloud platform
// to poll the Service Instance `last_operation` Endpoint for operation status. Note that a re-sent `PUT` request MUST
// return a `202 Accepted`, not a `200 OK`, if the service instance is not yet fully provisioned.
type Resp2079874Root struct {
	// For asynchronous responses, service brokers MAY return an identifier representing the operation. The value of this
	// field MUST be provided by the platform with requests to the Last Operation endpoint in a URL encoded query
	// parameter. If present, MUST be a non-empty string.
	Operation *string `json:"operation,omitempty"`
}

// UnmarshalResp2079874Root unmarshals an instance of Resp2079874Root from the specified map of raw messages.
func UnmarshalResp2079874Root(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Resp2079874Root)
	err = core.UnmarshalPrimitive(m, "operation", &obj.Operation)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Resp2079876Root : Resp2079876Root struct
type Resp2079876Root struct {
	// A free-form hash of credentials that can be used by applications or users to access the service.
	Credentials interface{} `json:"credentials,omitempty"`

	// A URL to which logs MUST be streamed. 'requires':['syslog_drain'] MUST be declared in the Catalog endpoint or the
	// platform MUST consider the response invalid.
	SyslogDrainURL *string `json:"syslog_drain_url,omitempty"`

	// A URL to which the platform MUST proxy requests for the address sent with bind_resource.route in the request body.
	// 'requires':['route_forwarding'] MUST be declared in the Catalog endpoint or the platform can consider the response
	// invalid.
	RouteServiceURL *string `json:"route_service_url,omitempty"`

	// An array of configuration for remote storage devices to be mounted into an application container filesystem.
	// 'requires':['volume_mount'] MUST be declared in the Catalog endpoint or the platform can consider the response
	// invalid.
	VolumeMounts []VolumeMount `json:"volume_mounts,omitempty"`
}

// UnmarshalResp2079876Root unmarshals an instance of Resp2079876Root from the specified map of raw messages.
func UnmarshalResp2079876Root(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Resp2079876Root)
	err = core.UnmarshalPrimitive(m, "credentials", &obj.Credentials)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "syslog_drain_url", &obj.SyslogDrainURL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "route_service_url", &obj.RouteServiceURL)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "volume_mounts", &obj.VolumeMounts, UnmarshalVolumeMount)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Resp2079894Root : OK - MUST be returned upon successful processing of this request.
type Resp2079894Root struct {
	// A user-facing message displayed to the platform API client. Can be used to tell the user details about the status of
	// the operation. If present, MUST be a non-empty string.
	Description *string `json:"description,omitempty"`

	// Valid values are `in progress`, `succeeded`, and `failed`. While `â€ state": "in progressâ€ `, the platform SHOULD
	// continue polling. A response with `â€ state": "succeededâ€ ` or `â€ state": "failedâ€ ` MUST cause the platform to
	// cease polling.
	State *string `json:"state" validate:"required"`
}

// UnmarshalResp2079894Root unmarshals an instance of Resp2079894Root from the specified map of raw messages.
func UnmarshalResp2079894Root(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Resp2079894Root)
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "state", &obj.State)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Resp2448145Root : Check the enabled status of active service.
type Resp2448145Root struct {
	// Indicates (from the viewpoint of the provider) whether the service instance is active and is meaningful if `enabled`
	// is true.  The default value is true if not specified.
	Active *bool `json:"active,omitempty"`

	// Indicates the current state of the service instance.
	Enabled *bool `json:"enabled" validate:"required"`

	// Indicates when the service instance was last accessed or modified, and is meaningful if `enabled` is true AND
	// `active` is false.  Represented as milliseconds since the epoch, but does not need to be accurate to the
	// second/hour.
	LastActive *int64 `json:"last_active,omitempty"`
}

// UnmarshalResp2448145Root unmarshals an instance of Resp2448145Root from the specified map of raw messages.
func UnmarshalResp2448145Root(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Resp2448145Root)
	err = core.UnmarshalPrimitive(m, "active", &obj.Active)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "enabled", &obj.Enabled)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "last_active", &obj.LastActive)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// UpdateServiceInstanceOptions : The UpdateServiceInstance options.
type UpdateServiceInstanceOptions struct {
	// The ID of a previously provisioned service instance.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// Platform specific contextual information under which the service instance is to be provisioned.
	Context *Context `json:"context,omitempty"`

	// Configuration options for the service instance. An opaque object, controller treats this as a blob. Brokers should
	// ensure that the client has provided valid configuration parameters and values for the operation. If this field is
	// not present in the request message, then the broker MUST NOT change the parameters of the instance as a result of
	// this request.
	Parameters map[string]string `json:"parameters,omitempty"`

	// The ID of the plan for which the service instance has been requested, which is stored in the catalog.json of your
	// broker. This value should be a GUID. MUST be unique to a service. If present, MUST be a non-empty string. If this
	// field is not present in the request message, then the broker MUST NOT change the plan of the instance as a result of
	// this request.
	PlanID *string `json:"plan_id,omitempty"`

	// Information about the service instance prior to the update.
	PreviousValues map[string]string `json:"previous_values,omitempty"`

	// The ID of the service stored in the catalog.json of your broker. This value should be a GUID. It MUST be a non-empty
	// string.
	ServiceID *string `json:"service_id,omitempty"`

	// A value of true indicates that both the IBM Cloud platform and the requesting client support asynchronous
	// deprovisioning. If this parameter is not included in the request, and the broker can only deprovision a service
	// instance of the requested plan asynchronously, the broker MUST reject the request with a `422` Unprocessable Entity.
	AcceptsIncomplete *bool `json:"accepts_incomplete,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateServiceInstanceOptions : Instantiate UpdateServiceInstanceOptions
func (*OpenServiceBrokerV1) NewUpdateServiceInstanceOptions(instanceID string) *UpdateServiceInstanceOptions {
	return &UpdateServiceInstanceOptions{
		InstanceID: core.StringPtr(instanceID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (options *UpdateServiceInstanceOptions) SetInstanceID(instanceID string) *UpdateServiceInstanceOptions {
	options.InstanceID = core.StringPtr(instanceID)
	return options
}

// SetContext : Allow user to set Context
func (options *UpdateServiceInstanceOptions) SetContext(context *Context) *UpdateServiceInstanceOptions {
	options.Context = context
	return options
}

// SetParameters : Allow user to set Parameters
func (options *UpdateServiceInstanceOptions) SetParameters(parameters map[string]string) *UpdateServiceInstanceOptions {
	options.Parameters = parameters
	return options
}

// SetPlanID : Allow user to set PlanID
func (options *UpdateServiceInstanceOptions) SetPlanID(planID string) *UpdateServiceInstanceOptions {
	options.PlanID = core.StringPtr(planID)
	return options
}

// SetPreviousValues : Allow user to set PreviousValues
func (options *UpdateServiceInstanceOptions) SetPreviousValues(previousValues map[string]string) *UpdateServiceInstanceOptions {
	options.PreviousValues = previousValues
	return options
}

// SetServiceID : Allow user to set ServiceID
func (options *UpdateServiceInstanceOptions) SetServiceID(serviceID string) *UpdateServiceInstanceOptions {
	options.ServiceID = core.StringPtr(serviceID)
	return options
}

// SetAcceptsIncomplete : Allow user to set AcceptsIncomplete
func (options *UpdateServiceInstanceOptions) SetAcceptsIncomplete(acceptsIncomplete bool) *UpdateServiceInstanceOptions {
	options.AcceptsIncomplete = core.BoolPtr(acceptsIncomplete)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateServiceInstanceOptions) SetHeaders(param map[string]string) *UpdateServiceInstanceOptions {
	options.Headers = param
	return options
}

// BindResource : A JSON object that contains data for platform resources associated with the binding to be created.
type BindResource struct {
	// Account owner of resource to bind.
	AccountID *string `json:"account_id,omitempty"`

	// Service ID of resource to bind.
	ServiceidCRN *string `json:"serviceid_crn,omitempty"`

	// Target ID of resource to bind.
	TargetCRN *string `json:"target_crn,omitempty"`

	// GUID of an application associated with the binding. For credentials bindings.
	AppGUID *string `json:"app_guid,omitempty"`

	// URL of the application to be intermediated. For route services bindings.
	Route *string `json:"route,omitempty"`
}

// UnmarshalBindResource unmarshals an instance of BindResource from the specified map of raw messages.
func UnmarshalBindResource(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(BindResource)
	err = core.UnmarshalPrimitive(m, "account_id", &obj.AccountID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "serviceid_crn", &obj.ServiceidCRN)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "target_crn", &obj.TargetCRN)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "app_guid", &obj.AppGUID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "route", &obj.Route)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Context : Platform specific contextual information under which the service instance is to be provisioned.
type Context struct {
	// Returns the ID of the account in IBM Cloud that is provisioning the service instance.
	AccountID *string `json:"account_id,omitempty"`

	// When a customer provisions your service in IBM Cloud, a service instance is created and this instance is identified
	// by its IBM Cloud Resource Name (CRN). The CRN is utilized in all aspects of the interaction with IBM Cloud including
	// provisioning, binding (creating credentials and endpoints), metering, dashboard display, and access control. From a
	// service provider perspective, the CRN can largely be treated as an opaque string to be utilized with the IBM Cloud
	// APIs, but it can also be decomposed via the following structure:
	// `crn:version:cname:ctype:service-name:location:scope:service-instance:resource-type:resource`.
	CRN *string `json:"crn,omitempty"`

	// Identifies the platform as "ibmcloud".
	Platform *string `json:"platform,omitempty"`
}

// UnmarshalContext unmarshals an instance of Context from the specified map of raw messages.
func UnmarshalContext(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Context)
	err = core.UnmarshalPrimitive(m, "account_id", &obj.AccountID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "crn", &obj.CRN)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "platform", &obj.Platform)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Plans : Where is this in the source?.
type Plans struct {
	// A short description of the plan. It MUST be a non-empty string. The description is NOT displayed in the IBM Cloud
	// catalog or IBM Cloud CLI.
	Description *string `json:"description" validate:"required"`

	// When false, service instances of this plan have a cost. The default is true.
	Free *bool `json:"free,omitempty"`

	// An identifier used to correlate this plan in future requests to the broker.  This MUST be globally unique within a
	// platform marketplace. It MUST be a non-empty string and using a GUID is RECOMMENDED. If you define your service in
	// the RMC, it will create a unique GUID for you to use. It is recommended to use the RMC to define and generate these
	// values and then use them in your catalog.json metadata in your broker. This value is NOT displayed in the IBM Cloud
	// catalog or IBM Cloud CLI.
	ID *string `json:"id" validate:"required"`

	// The programmatic name of the plan. It MUST be unique within the service. All lowercase, no spaces. It MUST be a
	// non-empty string, and it's NOT displayed in the IBM Cloud catalog or IBM Cloud CLI.
	Name *string `json:"name" validate:"required"`
}

// UnmarshalPlans unmarshals an instance of Plans from the specified map of raw messages.
func UnmarshalPlans(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Plans)
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "free", &obj.Free)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Services : The service object that describes the properties of your service.
type Services struct {
	// Specifies whether or not your service can be bound to applications in IBM Cloud. If bindable, it must be able to
	// return API endpoints and credentials to your service consumers.
	Bindable *bool `json:"bindable" validate:"required"`

	// A short description of the service. It MUST be a non-empty string. Note that this description is not displayed by
	// the the IBM Cloud console or IBM Cloud CLI.
	Description *string `json:"description" validate:"required"`

	// An identifier used to correlate this service in future requests to the broker. This MUST be globally unique within
	// the IBM Cloud platform. It MUST be a non-empty string, and using a GUID is recommended. Recommended: If you define
	// your service in the RMC, the RMC will generate a globally unique GUID service ID that you can use in your service
	// broker.
	ID *string `json:"id" validate:"required"`

	// The service name is not your display name. Your service name must follow the follow these rules:
	//  - It must be all lowercase.
	//  - It can't include spaces but may include hyphens (`-`).
	//  - It must be less than 32 characters.
	//  Your service name should include your company name. If your company has more then one offering your service name
	// should include both company and offering as part of the name. For example, the Compose company has offerings for
	// Redis and Elasticsearch. Sample service names on IBM Cloud for these offerings would be `compose-redis` and
	// `compose-elasticsearch`.  Each of these service names have associated display names that are shown in the IBM Cloud
	// catalog: *Compose Redis* and *Compose Elasticsearch*. Another company (e.g. FastJetMail) may only have the single
	// JetMail offering, in which case the service name should be `fastjetmail`. Recommended: If you define your service in
	// RMC, you can export a catalog.json that will include the service name you defined within the RMC.
	Name *string `json:"name" validate:"required"`

	// The Default is false. This specifices whether or not you support plan changes for provisioned instances. If your
	// offering supports multiple plans, and you want users to be able to change plans for a provisioned instance, you will
	// need to enable the ability for users to update their service instance by using /v2/service_instances/{instance_id}
	// PATCH.
	PlanUpdateable *bool `json:"plan_updateable,omitempty"`

	// A list of plans for this service that must contain at least one plan.
	Plans []Plans `json:"plans" validate:"required"`
}

// UnmarshalServices unmarshals an instance of Services from the specified map of raw messages.
func UnmarshalServices(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Services)
	err = core.UnmarshalPrimitive(m, "bindable", &obj.Bindable)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "plan_updateable", &obj.PlanUpdateable)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "plans", &obj.Plans, UnmarshalPlans)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// VolumeMount : VolumeMount struct
type VolumeMount struct {
	// A free-form hash of credentials that can be used by applications or users to access the service.
	Driver *string `json:"driver" validate:"required"`

	// The path in the application container onto which the volume will be mounted. This specification does not mandate
	// what action the platform is to take if the path specified already exists in the container.
	ContainerDir *string `json:"container_dir" validate:"required"`

	// 'r' to mount the volume read-only or 'rw' to mount it read-write.
	Mode *string `json:"mode" validate:"required"`

	// A string specifying the type of device to mount. Currently the only supported value is 'shared'.
	DeviceType *string `json:"device_type" validate:"required"`

	// Device object containing device_type specific details. Currently only shared devices are supported.
	Device *string `json:"device" validate:"required"`
}

// UnmarshalVolumeMount unmarshals an instance of VolumeMount from the specified map of raw messages.
func UnmarshalVolumeMount(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(VolumeMount)
	err = core.UnmarshalPrimitive(m, "driver", &obj.Driver)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "container_dir", &obj.ContainerDir)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "mode", &obj.Mode)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "device_type", &obj.DeviceType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "device", &obj.Device)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}
