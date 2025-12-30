/**
 * (C) Copyright IBM Corp. 2025.
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
 * IBM OpenAPI SDK Code Generator Version: 3.108.0-56772134-20251111-102802
 */

// Package distributionlistapiv1 : Operations and models for the DistributionListApiV1 service
package distributionlistapiv1

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/go-openapi/strfmt"
)

// DistributionListApiV1 : API for managing distribution lists for IBM Cloud accounts.
//
// API Version: 1.0.0
type DistributionListApiV1 struct {
	Service *core.BaseService
}

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "distribution_list_api"

// DistributionListApiV1Options : Service options
type DistributionListApiV1Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
}

// NewDistributionListApiV1UsingExternalConfig : constructs an instance of DistributionListApiV1 with passed in options and external configuration.
func NewDistributionListApiV1UsingExternalConfig(options *DistributionListApiV1Options) (distributionListApi *DistributionListApiV1, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			err = core.SDKErrorf(err, "", "env-auth-error", getServiceComponentInfo())
			return
		}
	}

	distributionListApi, err = NewDistributionListApiV1(options)
	err = core.RepurposeSDKProblem(err, "new-client-error")
	if err != nil {
		return
	}

	err = distributionListApi.Service.ConfigureService(options.ServiceName)
	if err != nil {
		err = core.SDKErrorf(err, "", "client-config-error", getServiceComponentInfo())
		return
	}

	if options.URL != "" {
		err = distributionListApi.Service.SetServiceURL(options.URL)
		err = core.RepurposeSDKProblem(err, "url-set-error")
	}
	return
}

// NewDistributionListApiV1 : constructs an instance of DistributionListApiV1 with passed in options.
func NewDistributionListApiV1(options *DistributionListApiV1Options) (service *DistributionListApiV1, err error) {
	serviceOptions := &core.ServiceOptions{
		Authenticator: options.Authenticator,
	}

	baseService, err := core.NewBaseService(serviceOptions)
	if err != nil {
		err = core.SDKErrorf(err, "", "new-base-error", getServiceComponentInfo())
		return
	}

	if options.URL != "" {
		err = baseService.SetServiceURL(options.URL)
		if err != nil {
			err = core.SDKErrorf(err, "", "set-url-error", getServiceComponentInfo())
			return
		}
	}

	service = &DistributionListApiV1{
		Service: baseService,
	}

	return
}

// GetServiceURLForRegion returns the service URL to be used for the specified region
func GetServiceURLForRegion(region string) (string, error) {
	return "", core.SDKErrorf(nil, "service does not support regional URLs", "no-regional-support", getServiceComponentInfo())
}

// Clone makes a copy of "distributionListApi" suitable for processing requests.
func (distributionListApi *DistributionListApiV1) Clone() *DistributionListApiV1 {
	if core.IsNil(distributionListApi) {
		return nil
	}
	clone := *distributionListApi
	clone.Service = distributionListApi.Service.Clone()
	return &clone
}

// SetServiceURL sets the service URL
func (distributionListApi *DistributionListApiV1) SetServiceURL(url string) error {
	err := distributionListApi.Service.SetServiceURL(url)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-set-error", getServiceComponentInfo())
	}
	return err
}

// GetServiceURL returns the service URL
func (distributionListApi *DistributionListApiV1) GetServiceURL() string {
	return distributionListApi.Service.GetServiceURL()
}

// SetDefaultHeaders sets HTTP headers to be sent in every request
func (distributionListApi *DistributionListApiV1) SetDefaultHeaders(headers http.Header) {
	distributionListApi.Service.SetDefaultHeaders(headers)
}

// SetEnableGzipCompression sets the service's EnableGzipCompression field
func (distributionListApi *DistributionListApiV1) SetEnableGzipCompression(enableGzip bool) {
	distributionListApi.Service.SetEnableGzipCompression(enableGzip)
}

// GetEnableGzipCompression returns the service's EnableGzipCompression field
func (distributionListApi *DistributionListApiV1) GetEnableGzipCompression() bool {
	return distributionListApi.Service.GetEnableGzipCompression()
}

// EnableRetries enables automatic retries for requests invoked for this service instance.
// If either parameter is specified as 0, then a default value is used instead.
func (distributionListApi *DistributionListApiV1) EnableRetries(maxRetries int, maxRetryInterval time.Duration) {
	distributionListApi.Service.EnableRetries(maxRetries, maxRetryInterval)
}

// DisableRetries disables automatic retries for requests invoked for this service instance.
func (distributionListApi *DistributionListApiV1) DisableRetries() {
	distributionListApi.Service.DisableRetries()
}

// GetAllDestinationEntries : Get all destination entries
// Retrieve all destinations in the distribution list for the specified account.
func (distributionListApi *DistributionListApiV1) GetAllDestinationEntries(getAllDestinationEntriesOptions *GetAllDestinationEntriesOptions) (result []DestinationListItemIntf, response *core.DetailedResponse, err error) {
	result, response, err = distributionListApi.GetAllDestinationEntriesWithContext(context.Background(), getAllDestinationEntriesOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetAllDestinationEntriesWithContext is an alternate form of the GetAllDestinationEntries method which supports a Context parameter
func (distributionListApi *DistributionListApiV1) GetAllDestinationEntriesWithContext(ctx context.Context, getAllDestinationEntriesOptions *GetAllDestinationEntriesOptions) (result []DestinationListItemIntf, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getAllDestinationEntriesOptions, "getAllDestinationEntriesOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", getServiceComponentInfo())
		return
	}
	err = core.ValidateStruct(getAllDestinationEntriesOptions, "getAllDestinationEntriesOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", getServiceComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"account_id": *getAllDestinationEntriesOptions.AccountID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = distributionListApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(distributionListApi.Service.Options.URL, `/notification-api/v1/distribution_lists/{account_id}/destinations`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", getServiceComponentInfo())
		return
	}

	sdkHeaders := map[string]string{
		"User-Agent": "terraform-provider-ibm/distribution-list-api/1.0.0",
	}
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range getAllDestinationEntriesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", getServiceComponentInfo())
		return
	}

	var rawResponse []json.RawMessage
	response, err = distributionListApi.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "getAllDestinationEntries", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", getServiceComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDestinationListItem)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", getServiceComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// AddDestinationEntry : Add a destination entry
// Add a destination entry to the distribution list. Maximum of 10 destination entries per destination type. In case of
// enterprise accounts, you can provide an Event Notifications destination that is from a different account than the
// distribution list account, provided these two accounts are from the same enterprise.
func (distributionListApi *DistributionListApiV1) AddDestinationEntry(addDestinationEntryOptions *AddDestinationEntryOptions) (result AddDestinationEntryResponseIntf, response *core.DetailedResponse, err error) {
	result, response, err = distributionListApi.AddDestinationEntryWithContext(context.Background(), addDestinationEntryOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// AddDestinationEntryWithContext is an alternate form of the AddDestinationEntry method which supports a Context parameter
func (distributionListApi *DistributionListApiV1) AddDestinationEntryWithContext(ctx context.Context, addDestinationEntryOptions *AddDestinationEntryOptions) (result AddDestinationEntryResponseIntf, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(addDestinationEntryOptions, "addDestinationEntryOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", getServiceComponentInfo())
		return
	}
	err = core.ValidateStruct(addDestinationEntryOptions, "addDestinationEntryOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", getServiceComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"account_id": *addDestinationEntryOptions.AccountID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = distributionListApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(distributionListApi.Service.Options.URL, `/notification-api/v1/distribution_lists/{account_id}/destinations`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", getServiceComponentInfo())
		return
	}

	sdkHeaders := map[string]string{
		"User-Agent": "terraform-provider-ibm/distribution-list-api/1.0.0",
	}
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range addDestinationEntryOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	_, err = builder.SetBodyContentJSON(addDestinationEntryOptions.AddDestinationEntryRequest)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", getServiceComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", getServiceComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = distributionListApi.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "addDestinationEntry", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", getServiceComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalAddDestinationEntryResponse)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", getServiceComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// GetDestinationEntry : Get a destination entry
// Retrieve a specific destination from the distribution list of the given account.
func (distributionListApi *DistributionListApiV1) GetDestinationEntry(getDestinationEntryOptions *GetDestinationEntryOptions) (result GetDestinationEntryResponseIntf, response *core.DetailedResponse, err error) {
	result, response, err = distributionListApi.GetDestinationEntryWithContext(context.Background(), getDestinationEntryOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetDestinationEntryWithContext is an alternate form of the GetDestinationEntry method which supports a Context parameter
func (distributionListApi *DistributionListApiV1) GetDestinationEntryWithContext(ctx context.Context, getDestinationEntryOptions *GetDestinationEntryOptions) (result GetDestinationEntryResponseIntf, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getDestinationEntryOptions, "getDestinationEntryOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", getServiceComponentInfo())
		return
	}
	err = core.ValidateStruct(getDestinationEntryOptions, "getDestinationEntryOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", getServiceComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"account_id":     *getDestinationEntryOptions.AccountID,
		"destination_id": fmt.Sprint(*getDestinationEntryOptions.DestinationID),
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = distributionListApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(distributionListApi.Service.Options.URL, `/notification-api/v1/distribution_lists/{account_id}/destinations/{destination_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", getServiceComponentInfo())
		return
	}

	sdkHeaders := map[string]string{
		"User-Agent": "terraform-provider-ibm/distribution-list-api/1.0.0",
	}
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range getDestinationEntryOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", getServiceComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = distributionListApi.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "getDestinationEntry", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", getServiceComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetDestinationEntryResponse)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", getServiceComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// DeleteDestinationEntry : Delete destination entry
// Remove a destination entry.
func (distributionListApi *DistributionListApiV1) DeleteDestinationEntry(deleteDestinationEntryOptions *DeleteDestinationEntryOptions) (response *core.DetailedResponse, err error) {
	response, err = distributionListApi.DeleteDestinationEntryWithContext(context.Background(), deleteDestinationEntryOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// DeleteDestinationEntryWithContext is an alternate form of the DeleteDestinationEntry method which supports a Context parameter
func (distributionListApi *DistributionListApiV1) DeleteDestinationEntryWithContext(ctx context.Context, deleteDestinationEntryOptions *DeleteDestinationEntryOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteDestinationEntryOptions, "deleteDestinationEntryOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", getServiceComponentInfo())
		return
	}
	err = core.ValidateStruct(deleteDestinationEntryOptions, "deleteDestinationEntryOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", getServiceComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"account_id":     *deleteDestinationEntryOptions.AccountID,
		"destination_id": fmt.Sprint(*deleteDestinationEntryOptions.DestinationID),
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = distributionListApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(distributionListApi.Service.Options.URL, `/notification-api/v1/distribution_lists/{account_id}/destinations/{destination_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", getServiceComponentInfo())
		return
	}

	sdkHeaders := map[string]string{
		"User-Agent": "terraform-provider-ibm/distribution-list-api/1.0.0",
	}
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range deleteDestinationEntryOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", getServiceComponentInfo())
		return
	}

	response, err = distributionListApi.Service.Request(request, nil)
	if err != nil {
		core.EnrichHTTPProblem(err, "deleteDestinationEntry", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", getServiceComponentInfo())
		return
	}

	return
}

// TestDestinationEntry : Test destination entry
// Send a test notification to a destination in the distribution list. This allows you to verify that the destination is
// properly configured and can receive notifications.
func (distributionListApi *DistributionListApiV1) TestDestinationEntry(testDestinationEntryOptions *TestDestinationEntryOptions) (result *TestDestinationEntryResponse, response *core.DetailedResponse, err error) {
	result, response, err = distributionListApi.TestDestinationEntryWithContext(context.Background(), testDestinationEntryOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// TestDestinationEntryWithContext is an alternate form of the TestDestinationEntry method which supports a Context parameter
func (distributionListApi *DistributionListApiV1) TestDestinationEntryWithContext(ctx context.Context, testDestinationEntryOptions *TestDestinationEntryOptions) (result *TestDestinationEntryResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(testDestinationEntryOptions, "testDestinationEntryOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", getServiceComponentInfo())
		return
	}
	err = core.ValidateStruct(testDestinationEntryOptions, "testDestinationEntryOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", getServiceComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"account_id":     *testDestinationEntryOptions.AccountID,
		"destination_id": fmt.Sprint(*testDestinationEntryOptions.DestinationID),
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = distributionListApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(distributionListApi.Service.Options.URL, `/notification-api/v1/distribution_lists/{account_id}/destinations/{destination_id}/test`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", getServiceComponentInfo())
		return
	}

	sdkHeaders := map[string]string{
		"User-Agent": "terraform-provider-ibm/distribution-list-api/1.0.0",
	}
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range testDestinationEntryOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	_, err = builder.SetBodyContentJSON(testDestinationEntryOptions.TestDestinationEntryRequest)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", getServiceComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", getServiceComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = distributionListApi.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "testDestinationEntry", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", getServiceComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTestDestinationEntryResponse)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", getServiceComponentInfo())
			return
		}
		response.Result = result
	}

	return
}
func getServiceComponentInfo() *core.ProblemComponent {
	return core.NewProblemComponent(DefaultServiceName, "1.0.0")
}

// AddDestinationEntryOptions : The AddDestinationEntry options.
type AddDestinationEntryOptions struct {
	// The IBM Cloud account ID.
	AccountID *string `json:"account_id" validate:"required,ne="`

	AddDestinationEntryRequest AddDestinationEntryRequestIntf `json:"AddDestinationEntryRequest" validate:"required"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewAddDestinationEntryOptions : Instantiate AddDestinationEntryOptions
func (*DistributionListApiV1) NewAddDestinationEntryOptions(accountID string, addDestinationEntryRequest AddDestinationEntryRequestIntf) *AddDestinationEntryOptions {
	return &AddDestinationEntryOptions{
		AccountID:                  core.StringPtr(accountID),
		AddDestinationEntryRequest: addDestinationEntryRequest,
	}
}

// SetAccountID : Allow user to set AccountID
func (_options *AddDestinationEntryOptions) SetAccountID(accountID string) *AddDestinationEntryOptions {
	_options.AccountID = core.StringPtr(accountID)
	return _options
}

// SetAddDestinationEntryRequest : Allow user to set AddDestinationEntryRequest
func (_options *AddDestinationEntryOptions) SetAddDestinationEntryRequest(addDestinationEntryRequest AddDestinationEntryRequestIntf) *AddDestinationEntryOptions {
	_options.AddDestinationEntryRequest = addDestinationEntryRequest
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *AddDestinationEntryOptions) SetHeaders(param map[string]string) *AddDestinationEntryOptions {
	options.Headers = param
	return options
}

// AddDestinationEntryRequest : AddDestinationEntryRequest struct
// Models which "extend" this model:
// - AddDestinationEntryRequestEventNotificationDestination
type AddDestinationEntryRequest struct {
	// The GUID of the Event Notifications instance.
	ID *strfmt.UUID `json:"id,omitempty"`

	// The type of the destination.
	DestinationType *string `json:"destination_type,omitempty"`
}

// Constants associated with the AddDestinationEntryRequest.DestinationType property.
// The type of the destination.
const (
	AddDestinationEntryRequest_DestinationType_EventNotifications = "event_notifications"
)

func (*AddDestinationEntryRequest) isaAddDestinationEntryRequest() bool {
	return true
}

type AddDestinationEntryRequestIntf interface {
	isaAddDestinationEntryRequest() bool
}

// UnmarshalAddDestinationEntryRequest unmarshals an instance of AddDestinationEntryRequest from the specified map of raw messages.
func UnmarshalAddDestinationEntryRequest(m map[string]json.RawMessage, result interface{}) (err error) {
	// Retrieve discriminator value to determine correct "subclass".
	var discValue string
	err = core.UnmarshalPrimitive(m, "destination_type", &discValue)
	if err != nil {
		errMsg := fmt.Sprintf("error unmarshalling discriminator property 'destination_type': %s", err.Error())
		err = core.SDKErrorf(err, errMsg, "discriminator-unmarshal-error", getServiceComponentInfo())
		return
	}
	if discValue == "" {
		err = core.SDKErrorf(err, "required discriminator property 'destination_type' not found in JSON object", "missing-discriminator", getServiceComponentInfo())
		return
	}
	if discValue == "event_notifications" {
		err = core.UnmarshalModel(m, "", result, UnmarshalAddDestinationEntryRequestEventNotificationDestination)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-AddDestinationEntryRequestEventNotificationDestination-error", getServiceComponentInfo())
		}
	} else {
		errMsg := fmt.Sprintf("unrecognized value for discriminator property 'destination_type': %s", discValue)
		err = core.SDKErrorf(err, errMsg, "invalid-discriminator", getServiceComponentInfo())
	}
	return
}

// AddDestinationEntryResponse : AddDestinationEntryResponse struct
// Models which "extend" this model:
// - AddDestinationEntryResponseEventNotificationDestination
type AddDestinationEntryResponse struct {
	// The GUID of the Event Notifications instance.
	ID *strfmt.UUID `json:"id,omitempty"`

	// The type of the destination.
	DestinationType *string `json:"destination_type,omitempty"`
}

// Constants associated with the AddDestinationEntryResponse.DestinationType property.
// The type of the destination.
const (
	AddDestinationEntryResponse_DestinationType_EventNotifications = "event_notifications"
)

func (*AddDestinationEntryResponse) isaAddDestinationEntryResponse() bool {
	return true
}

type AddDestinationEntryResponseIntf interface {
	isaAddDestinationEntryResponse() bool
}

// UnmarshalAddDestinationEntryResponse unmarshals an instance of AddDestinationEntryResponse from the specified map of raw messages.
func UnmarshalAddDestinationEntryResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	// Retrieve discriminator value to determine correct "subclass".
	var discValue string
	err = core.UnmarshalPrimitive(m, "destination_type", &discValue)
	if err != nil {
		errMsg := fmt.Sprintf("error unmarshalling discriminator property 'destination_type': %s", err.Error())
		err = core.SDKErrorf(err, errMsg, "discriminator-unmarshal-error", getServiceComponentInfo())
		return
	}
	if discValue == "" {
		err = core.SDKErrorf(err, "required discriminator property 'destination_type' not found in JSON object", "missing-discriminator", getServiceComponentInfo())
		return
	}
	if discValue == "event_notifications" {
		err = core.UnmarshalModel(m, "", result, UnmarshalAddDestinationEntryResponseEventNotificationDestination)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-AddDestinationEntryResponseEventNotificationDestination-error", getServiceComponentInfo())
		}
	} else {
		errMsg := fmt.Sprintf("unrecognized value for discriminator property 'destination_type': %s", discValue)
		err = core.SDKErrorf(err, errMsg, "invalid-discriminator", getServiceComponentInfo())
	}
	return
}

// DeleteDestinationEntryOptions : The DeleteDestinationEntry options.
type DeleteDestinationEntryOptions struct {
	// The IBM Cloud account ID.
	AccountID *string `json:"account_id" validate:"required,ne="`

	// The GUID of the destination.
	DestinationID *strfmt.UUID `json:"destination_id" validate:"required"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewDeleteDestinationEntryOptions : Instantiate DeleteDestinationEntryOptions
func (*DistributionListApiV1) NewDeleteDestinationEntryOptions(accountID string, destinationID *strfmt.UUID) *DeleteDestinationEntryOptions {
	return &DeleteDestinationEntryOptions{
		AccountID:     core.StringPtr(accountID),
		DestinationID: destinationID,
	}
}

// SetAccountID : Allow user to set AccountID
func (_options *DeleteDestinationEntryOptions) SetAccountID(accountID string) *DeleteDestinationEntryOptions {
	_options.AccountID = core.StringPtr(accountID)
	return _options
}

// SetDestinationID : Allow user to set DestinationID
func (_options *DeleteDestinationEntryOptions) SetDestinationID(destinationID *strfmt.UUID) *DeleteDestinationEntryOptions {
	_options.DestinationID = destinationID
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteDestinationEntryOptions) SetHeaders(param map[string]string) *DeleteDestinationEntryOptions {
	options.Headers = param
	return options
}

// DestinationListItem : DestinationListItem struct
// Models which "extend" this model:
// - DestinationListItemEventNotificationDestination
type DestinationListItem struct {
	// The GUID of the Event Notifications instance.
	ID *strfmt.UUID `json:"id,omitempty"`

	// The type of the destination.
	DestinationType *string `json:"destination_type,omitempty"`
}

// Constants associated with the DestinationListItem.DestinationType property.
// The type of the destination.
const (
	DestinationListItem_DestinationType_EventNotifications = "event_notifications"
)

func (*DestinationListItem) isaDestinationListItem() bool {
	return true
}

type DestinationListItemIntf interface {
	isaDestinationListItem() bool
}

// UnmarshalDestinationListItem unmarshals an instance of DestinationListItem from the specified map of raw messages.
func UnmarshalDestinationListItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DestinationListItem)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", getServiceComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "destination_type", &obj.DestinationType)
	if err != nil {
		err = core.SDKErrorf(err, "", "destination_type-error", getServiceComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetAllDestinationEntriesOptions : The GetAllDestinationEntries options.
type GetAllDestinationEntriesOptions struct {
	// The IBM Cloud account ID.
	AccountID *string `json:"account_id" validate:"required,ne="`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetAllDestinationEntriesOptions : Instantiate GetAllDestinationEntriesOptions
func (*DistributionListApiV1) NewGetAllDestinationEntriesOptions(accountID string) *GetAllDestinationEntriesOptions {
	return &GetAllDestinationEntriesOptions{
		AccountID: core.StringPtr(accountID),
	}
}

// SetAccountID : Allow user to set AccountID
func (_options *GetAllDestinationEntriesOptions) SetAccountID(accountID string) *GetAllDestinationEntriesOptions {
	_options.AccountID = core.StringPtr(accountID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetAllDestinationEntriesOptions) SetHeaders(param map[string]string) *GetAllDestinationEntriesOptions {
	options.Headers = param
	return options
}

// GetDestinationEntryOptions : The GetDestinationEntry options.
type GetDestinationEntryOptions struct {
	// The IBM Cloud account ID.
	AccountID *string `json:"account_id" validate:"required,ne="`

	// The GUID of the destination.
	DestinationID *strfmt.UUID `json:"destination_id" validate:"required"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetDestinationEntryOptions : Instantiate GetDestinationEntryOptions
func (*DistributionListApiV1) NewGetDestinationEntryOptions(accountID string, destinationID *strfmt.UUID) *GetDestinationEntryOptions {
	return &GetDestinationEntryOptions{
		AccountID:     core.StringPtr(accountID),
		DestinationID: destinationID,
	}
}

// SetAccountID : Allow user to set AccountID
func (_options *GetDestinationEntryOptions) SetAccountID(accountID string) *GetDestinationEntryOptions {
	_options.AccountID = core.StringPtr(accountID)
	return _options
}

// SetDestinationID : Allow user to set DestinationID
func (_options *GetDestinationEntryOptions) SetDestinationID(destinationID *strfmt.UUID) *GetDestinationEntryOptions {
	_options.DestinationID = destinationID
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetDestinationEntryOptions) SetHeaders(param map[string]string) *GetDestinationEntryOptions {
	options.Headers = param
	return options
}

// GetDestinationEntryResponse : GetDestinationEntryResponse struct
// Models which "extend" this model:
// - GetDestinationEntryResponseEventNotificationDestination
type GetDestinationEntryResponse struct {
	// The GUID of the Event Notifications instance.
	ID *strfmt.UUID `json:"id,omitempty"`

	// The type of the destination.
	DestinationType *string `json:"destination_type,omitempty"`
}

// Constants associated with the GetDestinationEntryResponse.DestinationType property.
// The type of the destination.
const (
	GetDestinationEntryResponse_DestinationType_EventNotifications = "event_notifications"
)

func (*GetDestinationEntryResponse) isaGetDestinationEntryResponse() bool {
	return true
}

type GetDestinationEntryResponseIntf interface {
	isaGetDestinationEntryResponse() bool
}

// UnmarshalGetDestinationEntryResponse unmarshals an instance of GetDestinationEntryResponse from the specified map of raw messages.
func UnmarshalGetDestinationEntryResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	// Retrieve discriminator value to determine correct "subclass".
	var discValue string
	err = core.UnmarshalPrimitive(m, "destination_type", &discValue)
	if err != nil {
		errMsg := fmt.Sprintf("error unmarshalling discriminator property 'destination_type': %s", err.Error())
		err = core.SDKErrorf(err, errMsg, "discriminator-unmarshal-error", getServiceComponentInfo())
		return
	}
	if discValue == "" {
		err = core.SDKErrorf(err, "required discriminator property 'destination_type' not found in JSON object", "missing-discriminator", getServiceComponentInfo())
		return
	}
	if discValue == "event_notifications" {
		err = core.UnmarshalModel(m, "", result, UnmarshalGetDestinationEntryResponseEventNotificationDestination)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-GetDestinationEntryResponseEventNotificationDestination-error", getServiceComponentInfo())
		}
	} else {
		errMsg := fmt.Sprintf("unrecognized value for discriminator property 'destination_type': %s", discValue)
		err = core.SDKErrorf(err, errMsg, "invalid-discriminator", getServiceComponentInfo())
	}
	return
}

// TestDestinationEntryOptions : The TestDestinationEntry options.
type TestDestinationEntryOptions struct {
	// The IBM Cloud account ID.
	AccountID *string `json:"account_id" validate:"required,ne="`

	// The GUID of the destination.
	DestinationID *strfmt.UUID `json:"destination_id" validate:"required"`

	TestDestinationEntryRequest TestDestinationEntryRequestIntf `json:"TestDestinationEntryRequest" validate:"required"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewTestDestinationEntryOptions : Instantiate TestDestinationEntryOptions
func (*DistributionListApiV1) NewTestDestinationEntryOptions(accountID string, destinationID *strfmt.UUID, testDestinationEntryRequest TestDestinationEntryRequestIntf) *TestDestinationEntryOptions {
	return &TestDestinationEntryOptions{
		AccountID:                   core.StringPtr(accountID),
		DestinationID:               destinationID,
		TestDestinationEntryRequest: testDestinationEntryRequest,
	}
}

// SetAccountID : Allow user to set AccountID
func (_options *TestDestinationEntryOptions) SetAccountID(accountID string) *TestDestinationEntryOptions {
	_options.AccountID = core.StringPtr(accountID)
	return _options
}

// SetDestinationID : Allow user to set DestinationID
func (_options *TestDestinationEntryOptions) SetDestinationID(destinationID *strfmt.UUID) *TestDestinationEntryOptions {
	_options.DestinationID = destinationID
	return _options
}

// SetTestDestinationEntryRequest : Allow user to set TestDestinationEntryRequest
func (_options *TestDestinationEntryOptions) SetTestDestinationEntryRequest(testDestinationEntryRequest TestDestinationEntryRequestIntf) *TestDestinationEntryOptions {
	_options.TestDestinationEntryRequest = testDestinationEntryRequest
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *TestDestinationEntryOptions) SetHeaders(param map[string]string) *TestDestinationEntryOptions {
	options.Headers = param
	return options
}

// TestDestinationEntryRequest : TestDestinationEntryRequest struct
// Models which "extend" this model:
// - TestDestinationEntryRequestTestEventNotificationDestination
type TestDestinationEntryRequest struct {
	// The type of the destination.
	DestinationType *string `json:"destination_type,omitempty"`

	// Type of notification to test.
	NotificationType *string `json:"notification_type,omitempty"`
}

// Constants associated with the TestDestinationEntryRequest.DestinationType property.
// The type of the destination.
const (
	TestDestinationEntryRequest_DestinationType_EventNotifications = "event_notifications"
)

// Constants associated with the TestDestinationEntryRequest.NotificationType property.
// Type of notification to test.
const (
	TestDestinationEntryRequest_NotificationType_Announcements     = "announcements"
	TestDestinationEntryRequest_NotificationType_BillingAndUsage   = "billing_and_usage"
	TestDestinationEntryRequest_NotificationType_Incident          = "incident"
	TestDestinationEntryRequest_NotificationType_Maintenance       = "maintenance"
	TestDestinationEntryRequest_NotificationType_Resource          = "resource"
	TestDestinationEntryRequest_NotificationType_SecurityBulletins = "security_bulletins"
)

func (*TestDestinationEntryRequest) isaTestDestinationEntryRequest() bool {
	return true
}

type TestDestinationEntryRequestIntf interface {
	isaTestDestinationEntryRequest() bool
}

// UnmarshalTestDestinationEntryRequest unmarshals an instance of TestDestinationEntryRequest from the specified map of raw messages.
func UnmarshalTestDestinationEntryRequest(m map[string]json.RawMessage, result interface{}) (err error) {
	// Retrieve discriminator value to determine correct "subclass".
	var discValue string
	err = core.UnmarshalPrimitive(m, "destination_type", &discValue)
	if err != nil {
		errMsg := fmt.Sprintf("error unmarshalling discriminator property 'destination_type': %s", err.Error())
		err = core.SDKErrorf(err, errMsg, "discriminator-unmarshal-error", getServiceComponentInfo())
		return
	}
	if discValue == "" {
		err = core.SDKErrorf(err, "required discriminator property 'destination_type' not found in JSON object", "missing-discriminator", getServiceComponentInfo())
		return
	}
	if discValue == "event_notifications" {
		err = core.UnmarshalModel(m, "", result, UnmarshalTestDestinationEntryRequestTestEventNotificationDestination)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-TestDestinationEntryRequestTestEventNotificationDestination-error", getServiceComponentInfo())
		}
	} else {
		errMsg := fmt.Sprintf("unrecognized value for discriminator property 'destination_type': %s", discValue)
		err = core.SDKErrorf(err, errMsg, "invalid-discriminator", getServiceComponentInfo())
	}
	return
}

// TestDestinationEntryResponse : TestDestinationEntryResponse struct
type TestDestinationEntryResponse struct {
	Message *string `json:"message,omitempty"`
}

// UnmarshalTestDestinationEntryResponse unmarshals an instance of TestDestinationEntryResponse from the specified map of raw messages.
func UnmarshalTestDestinationEntryResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TestDestinationEntryResponse)
	err = core.UnmarshalPrimitive(m, "message", &obj.Message)
	if err != nil {
		err = core.SDKErrorf(err, "", "message-error", getServiceComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AddDestinationEntryRequestEventNotificationDestination : An Event Notifications destination entry in the distribution list.
// This model "extends" AddDestinationEntryRequest
type AddDestinationEntryRequestEventNotificationDestination struct {
	// The GUID of the Event Notifications instance.
	ID *strfmt.UUID `json:"id" validate:"required"`

	// The type of the destination.
	DestinationType *string `json:"destination_type" validate:"required"`
}

// Constants associated with the AddDestinationEntryRequestEventNotificationDestination.DestinationType property.
// The type of the destination.
const (
	AddDestinationEntryRequestEventNotificationDestination_DestinationType_EventNotifications = "event_notifications"
)

// NewAddDestinationEntryRequestEventNotificationDestination : Instantiate AddDestinationEntryRequestEventNotificationDestination (Generic Model Constructor)
func (*DistributionListApiV1) NewAddDestinationEntryRequestEventNotificationDestination(id *strfmt.UUID, destinationType string) (_model *AddDestinationEntryRequestEventNotificationDestination, err error) {
	_model = &AddDestinationEntryRequestEventNotificationDestination{
		ID:              id,
		DestinationType: core.StringPtr(destinationType),
	}
	err = core.ValidateStruct(_model, "required parameters")
	if err != nil {
		err = core.SDKErrorf(err, "", "model-missing-required", getServiceComponentInfo())
	}
	return
}

func (*AddDestinationEntryRequestEventNotificationDestination) isaAddDestinationEntryRequest() bool {
	return true
}

// UnmarshalAddDestinationEntryRequestEventNotificationDestination unmarshals an instance of AddDestinationEntryRequestEventNotificationDestination from the specified map of raw messages.
func UnmarshalAddDestinationEntryRequestEventNotificationDestination(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AddDestinationEntryRequestEventNotificationDestination)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", getServiceComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "destination_type", &obj.DestinationType)
	if err != nil {
		err = core.SDKErrorf(err, "", "destination_type-error", getServiceComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AddDestinationEntryResponseEventNotificationDestination : An Event Notifications destination entry in the distribution list.
// This model "extends" AddDestinationEntryResponse
type AddDestinationEntryResponseEventNotificationDestination struct {
	// The GUID of the Event Notifications instance.
	ID *strfmt.UUID `json:"id" validate:"required"`

	// The type of the destination.
	DestinationType *string `json:"destination_type" validate:"required"`
}

// Constants associated with the AddDestinationEntryResponseEventNotificationDestination.DestinationType property.
// The type of the destination.
const (
	AddDestinationEntryResponseEventNotificationDestination_DestinationType_EventNotifications = "event_notifications"
)

func (*AddDestinationEntryResponseEventNotificationDestination) isaAddDestinationEntryResponse() bool {
	return true
}

// UnmarshalAddDestinationEntryResponseEventNotificationDestination unmarshals an instance of AddDestinationEntryResponseEventNotificationDestination from the specified map of raw messages.
func UnmarshalAddDestinationEntryResponseEventNotificationDestination(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AddDestinationEntryResponseEventNotificationDestination)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", getServiceComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "destination_type", &obj.DestinationType)
	if err != nil {
		err = core.SDKErrorf(err, "", "destination_type-error", getServiceComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DestinationListItemEventNotificationDestination : An Event Notifications destination entry in the distribution list.
// This model "extends" DestinationListItem
type DestinationListItemEventNotificationDestination struct {
	// The GUID of the Event Notifications instance.
	ID *strfmt.UUID `json:"id" validate:"required"`

	// The type of the destination.
	DestinationType *string `json:"destination_type" validate:"required"`
}

// Constants associated with the DestinationListItemEventNotificationDestination.DestinationType property.
// The type of the destination.
const (
	DestinationListItemEventNotificationDestination_DestinationType_EventNotifications = "event_notifications"
)

func (*DestinationListItemEventNotificationDestination) isaDestinationListItem() bool {
	return true
}

// UnmarshalDestinationListItemEventNotificationDestination unmarshals an instance of DestinationListItemEventNotificationDestination from the specified map of raw messages.
func UnmarshalDestinationListItemEventNotificationDestination(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DestinationListItemEventNotificationDestination)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", getServiceComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "destination_type", &obj.DestinationType)
	if err != nil {
		err = core.SDKErrorf(err, "", "destination_type-error", getServiceComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetDestinationEntryResponseEventNotificationDestination : An Event Notifications destination entry in the distribution list.
// This model "extends" GetDestinationEntryResponse
type GetDestinationEntryResponseEventNotificationDestination struct {
	// The GUID of the Event Notifications instance.
	ID *strfmt.UUID `json:"id" validate:"required"`

	// The type of the destination.
	DestinationType *string `json:"destination_type" validate:"required"`
}

// Constants associated with the GetDestinationEntryResponseEventNotificationDestination.DestinationType property.
// The type of the destination.
const (
	GetDestinationEntryResponseEventNotificationDestination_DestinationType_EventNotifications = "event_notifications"
)

func (*GetDestinationEntryResponseEventNotificationDestination) isaGetDestinationEntryResponse() bool {
	return true
}

// UnmarshalGetDestinationEntryResponseEventNotificationDestination unmarshals an instance of GetDestinationEntryResponseEventNotificationDestination from the specified map of raw messages.
func UnmarshalGetDestinationEntryResponseEventNotificationDestination(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetDestinationEntryResponseEventNotificationDestination)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", getServiceComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "destination_type", &obj.DestinationType)
	if err != nil {
		err = core.SDKErrorf(err, "", "destination_type-error", getServiceComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TestDestinationEntryRequestTestEventNotificationDestination : TestDestinationEntryRequestTestEventNotificationDestination struct
// This model "extends" TestDestinationEntryRequest
type TestDestinationEntryRequestTestEventNotificationDestination struct {
	// The type of the destination.
	DestinationType *string `json:"destination_type" validate:"required"`

	// Type of notification to test.
	NotificationType *string `json:"notification_type" validate:"required"`
}

// Constants associated with the TestDestinationEntryRequestTestEventNotificationDestination.DestinationType property.
// The type of the destination.
const (
	TestDestinationEntryRequestTestEventNotificationDestination_DestinationType_EventNotifications = "event_notifications"
)

// Constants associated with the TestDestinationEntryRequestTestEventNotificationDestination.NotificationType property.
// Type of notification to test.
const (
	TestDestinationEntryRequestTestEventNotificationDestination_NotificationType_Announcements     = "announcements"
	TestDestinationEntryRequestTestEventNotificationDestination_NotificationType_BillingAndUsage   = "billing_and_usage"
	TestDestinationEntryRequestTestEventNotificationDestination_NotificationType_Incident          = "incident"
	TestDestinationEntryRequestTestEventNotificationDestination_NotificationType_Maintenance       = "maintenance"
	TestDestinationEntryRequestTestEventNotificationDestination_NotificationType_Resource          = "resource"
	TestDestinationEntryRequestTestEventNotificationDestination_NotificationType_SecurityBulletins = "security_bulletins"
)

// NewTestDestinationEntryRequestTestEventNotificationDestination : Instantiate TestDestinationEntryRequestTestEventNotificationDestination (Generic Model Constructor)
func (*DistributionListApiV1) NewTestDestinationEntryRequestTestEventNotificationDestination(destinationType string, notificationType string) (_model *TestDestinationEntryRequestTestEventNotificationDestination, err error) {
	_model = &TestDestinationEntryRequestTestEventNotificationDestination{
		DestinationType:  core.StringPtr(destinationType),
		NotificationType: core.StringPtr(notificationType),
	}
	err = core.ValidateStruct(_model, "required parameters")
	if err != nil {
		err = core.SDKErrorf(err, "", "model-missing-required", getServiceComponentInfo())
	}
	return
}

func (*TestDestinationEntryRequestTestEventNotificationDestination) isaTestDestinationEntryRequest() bool {
	return true
}

// UnmarshalTestDestinationEntryRequestTestEventNotificationDestination unmarshals an instance of TestDestinationEntryRequestTestEventNotificationDestination from the specified map of raw messages.
func UnmarshalTestDestinationEntryRequestTestEventNotificationDestination(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TestDestinationEntryRequestTestEventNotificationDestination)
	err = core.UnmarshalPrimitive(m, "destination_type", &obj.DestinationType)
	if err != nil {
		err = core.SDKErrorf(err, "", "destination_type-error", getServiceComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "notification_type", &obj.NotificationType)
	if err != nil {
		err = core.SDKErrorf(err, "", "notification_type-error", getServiceComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}
