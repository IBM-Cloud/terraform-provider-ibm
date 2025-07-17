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

// Package usagemeteringv4 : Operations and models for the UsageMeteringV4 service
package usagemeteringv4

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

// UsageMeteringV4 : IBM Cloud Usage Metering is a platform service that enables service providers to submit metrics
// collected for  resource instances provisioned by IBM Cloud users. IBM and third-party service providers that are
// delivering  an integrated billing service in IBM Cloud are required to submit usage for all active service instances
// every hour.  This is important because inability to report usage can lead to loss of revenue collection for IBM,  in
// turn causing loss of revenue share for the service providers.
//
// Version: 4.0.8
type UsageMeteringV4 struct {
	Service *core.BaseService
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://billing.cloud.ibm.com"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "usage_metering"

// UsageMeteringV4Options : Service options
type UsageMeteringV4Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
}

// NewUsageMeteringV4UsingExternalConfig : constructs an instance of UsageMeteringV4 with passed in options and external configuration.
func NewUsageMeteringV4UsingExternalConfig(options *UsageMeteringV4Options) (usageMetering *UsageMeteringV4, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	usageMetering, err = NewUsageMeteringV4(options)
	if err != nil {
		return
	}

	err = usageMetering.Service.ConfigureService(options.ServiceName)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = usageMetering.Service.SetServiceURL(options.URL)
	}
	return
}

// NewUsageMeteringV4 : constructs an instance of UsageMeteringV4 with passed in options.
func NewUsageMeteringV4(options *UsageMeteringV4Options) (service *UsageMeteringV4, err error) {
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

	service = &UsageMeteringV4{
		Service: baseService,
	}

	return
}

// GetServiceURLForRegion returns the service URL to be used for the specified region
func GetServiceURLForRegion(region string) (string, error) {
	return "", fmt.Errorf("service does not support regional URLs")
}

// Clone makes a copy of "usageMetering" suitable for processing requests.
func (usageMetering *UsageMeteringV4) Clone() *UsageMeteringV4 {
	if core.IsNil(usageMetering) {
		return nil
	}
	clone := *usageMetering
	clone.Service = usageMetering.Service.Clone()
	return &clone
}

// SetServiceURL sets the service URL
func (usageMetering *UsageMeteringV4) SetServiceURL(url string) error {
	return usageMetering.Service.SetServiceURL(url)
}

// GetServiceURL returns the service URL
func (usageMetering *UsageMeteringV4) GetServiceURL() string {
	return usageMetering.Service.GetServiceURL()
}

// SetDefaultHeaders sets HTTP headers to be sent in every request
func (usageMetering *UsageMeteringV4) SetDefaultHeaders(headers http.Header) {
	usageMetering.Service.SetDefaultHeaders(headers)
}

// SetEnableGzipCompression sets the service's EnableGzipCompression field
func (usageMetering *UsageMeteringV4) SetEnableGzipCompression(enableGzip bool) {
	usageMetering.Service.SetEnableGzipCompression(enableGzip)
}

// GetEnableGzipCompression returns the service's EnableGzipCompression field
func (usageMetering *UsageMeteringV4) GetEnableGzipCompression() bool {
	return usageMetering.Service.GetEnableGzipCompression()
}

// EnableRetries enables automatic retries for requests invoked for this service instance.
// If either parameter is specified as 0, then a default value is used instead.
func (usageMetering *UsageMeteringV4) EnableRetries(maxRetries int, maxRetryInterval time.Duration) {
	usageMetering.Service.EnableRetries(maxRetries, maxRetryInterval)
}

// DisableRetries disables automatic retries for requests invoked for this service instance.
func (usageMetering *UsageMeteringV4) DisableRetries() {
	usageMetering.Service.DisableRetries()
}

// ReportResourceUsage : Report Resource Controller resource usage
// Report usage for resource instances that were provisioned through the resource controller.
func (usageMetering *UsageMeteringV4) ReportResourceUsage(reportResourceUsageOptions *ReportResourceUsageOptions) (result *ResponseAccepted, response *core.DetailedResponse, err error) {
	return usageMetering.ReportResourceUsageWithContext(context.Background(), reportResourceUsageOptions)
}

// ReportResourceUsageWithContext is an alternate form of the ReportResourceUsage method which supports a Context parameter
func (usageMetering *UsageMeteringV4) ReportResourceUsageWithContext(ctx context.Context, reportResourceUsageOptions *ReportResourceUsageOptions) (result *ResponseAccepted, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(reportResourceUsageOptions, "reportResourceUsageOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(reportResourceUsageOptions, "reportResourceUsageOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"resource_id": *reportResourceUsageOptions.ResourceID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = usageMetering.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(usageMetering.Service.Options.URL, `/v4/metering/resources/{resource_id}/usage`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range reportResourceUsageOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("usage_metering", "V4", "ReportResourceUsage")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	_, err = builder.SetBodyContentJSON(reportResourceUsageOptions.ResourceUsage)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = usageMetering.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalResponseAccepted)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// ReportResourceUsageOptions : The ReportResourceUsage options.
type ReportResourceUsageOptions struct {
	// The resource for which the usage is submitted.
	ResourceID *string `json:"resource_id" validate:"required,ne="`

	// Array of usage records.
	ResourceUsage []ResourceInstanceUsage `json:"resource_usage" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewReportResourceUsageOptions : Instantiate ReportResourceUsageOptions
func (*UsageMeteringV4) NewReportResourceUsageOptions(resourceID string, resourceUsage []ResourceInstanceUsage) *ReportResourceUsageOptions {
	return &ReportResourceUsageOptions{
		ResourceID:    core.StringPtr(resourceID),
		ResourceUsage: resourceUsage,
	}
}

// SetResourceID : Allow user to set ResourceID
func (options *ReportResourceUsageOptions) SetResourceID(resourceID string) *ReportResourceUsageOptions {
	options.ResourceID = core.StringPtr(resourceID)
	return options
}

// SetResourceUsage : Allow user to set ResourceUsage
func (options *ReportResourceUsageOptions) SetResourceUsage(resourceUsage []ResourceInstanceUsage) *ReportResourceUsageOptions {
	options.ResourceUsage = resourceUsage
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ReportResourceUsageOptions) SetHeaders(param map[string]string) *ReportResourceUsageOptions {
	options.Headers = param
	return options
}

// MeasureAndQuantity : A usage measurement.
type MeasureAndQuantity struct {
	// The name of the measure.
	Measure *string `json:"measure" validate:"required"`

	// For consumption-based submissions, `quantity` can be a double or integer value. For event-based submissions that do
	// not have binary states, previous and current values are required, such as `{ "previous": 1, "current": 2 }`.
	Quantity interface{} `json:"quantity" validate:"required"`
}

// NewMeasureAndQuantity : Instantiate MeasureAndQuantity (Generic Model Constructor)
func (*UsageMeteringV4) NewMeasureAndQuantity(measure string, quantity interface{}) (model *MeasureAndQuantity, err error) {
	model = &MeasureAndQuantity{
		Measure:  core.StringPtr(measure),
		Quantity: quantity,
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalMeasureAndQuantity unmarshals an instance of MeasureAndQuantity from the specified map of raw messages.
func UnmarshalMeasureAndQuantity(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MeasureAndQuantity)
	err = core.UnmarshalPrimitive(m, "measure", &obj.Measure)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "quantity", &obj.Quantity)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ResourceInstanceUsage : Usage information for a resource instance.
type ResourceInstanceUsage struct {
	// The ID of the instance that incurred the usage. The ID is a CRN for instances that are provisioned with the resource
	// controller.
	ResourceInstanceID *string `json:"resource_instance_id" validate:"required"`

	// The plan with which the instance's usage should be metered.
	PlanID *string `json:"plan_id" validate:"required"`

	// The pricing region to which the usage must be aggregated. This field is required if the ID is not a CRN or if the
	// CRN does not have a region.
	Region *string `json:"region,omitempty"`

	// The time from which the resource instance was metered in the format milliseconds since epoch.
	Start *int64 `json:"start" validate:"required"`

	// The time until which the resource instance was metered in the format milliseconds since epoch. This value is the
	// same as start value for event-based submissions.
	End *int64 `json:"end" validate:"required"`

	// Usage measurements for the resource instance.
	MeasuredUsage []MeasureAndQuantity `json:"measured_usage" validate:"required"`

	// If an instance's usage should be aggregated at the consumer level, specify the ID of the consumer. Usage is
	// accumulated to the instance-consumer combination.
	ConsumerID *string `json:"consumer_id,omitempty"`
}

// NewResourceInstanceUsage : Instantiate ResourceInstanceUsage (Generic Model Constructor)
func (*UsageMeteringV4) NewResourceInstanceUsage(resourceInstanceID string, planID string, start int64, end int64, measuredUsage []MeasureAndQuantity) (model *ResourceInstanceUsage, err error) {
	model = &ResourceInstanceUsage{
		ResourceInstanceID: core.StringPtr(resourceInstanceID),
		PlanID:             core.StringPtr(planID),
		Start:              core.Int64Ptr(start),
		End:                core.Int64Ptr(end),
		MeasuredUsage:      measuredUsage,
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalResourceInstanceUsage unmarshals an instance of ResourceInstanceUsage from the specified map of raw messages.
func UnmarshalResourceInstanceUsage(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ResourceInstanceUsage)
	err = core.UnmarshalPrimitive(m, "resource_instance_id", &obj.ResourceInstanceID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "plan_id", &obj.PlanID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "region", &obj.Region)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "start", &obj.Start)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "end", &obj.End)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "measured_usage", &obj.MeasuredUsage, UnmarshalMeasureAndQuantity)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "consumer_id", &obj.ConsumerID)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ResourceUsageDetails : Resource usage details.
type ResourceUsageDetails struct {
	// A response code similar to HTTP status codes.
	Status *int64 `json:"status" validate:"required"`

	// The location of the usage.
	Location *string `json:"location" validate:"required"`

	// The error code that was encountered.
	Code *string `json:"code,omitempty"`

	// A description of the error.
	Message *string `json:"message,omitempty"`
}

// UnmarshalResourceUsageDetails unmarshals an instance of ResourceUsageDetails from the specified map of raw messages.
func UnmarshalResourceUsageDetails(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ResourceUsageDetails)
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "location", &obj.Location)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "code", &obj.Code)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "message", &obj.Message)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ResponseAccepted : Response when usage submitted is accepted.
type ResponseAccepted struct {
	// Response body that contains the status of each submitted usage record.
	Resources []ResourceUsageDetails `json:"resources" validate:"required"`
}

// UnmarshalResponseAccepted unmarshals an instance of ResponseAccepted from the specified map of raw messages.
func UnmarshalResponseAccepted(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ResponseAccepted)
	err = core.UnmarshalModel(m, "resources", &obj.Resources, UnmarshalResourceUsageDetails)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}
