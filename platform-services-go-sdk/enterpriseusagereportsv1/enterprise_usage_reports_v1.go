/**
 * (C) Copyright IBM Corp. 2020, 2022.
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
 * IBM OpenAPI SDK Code Generator Version: 3.60.0-13f6e1ba-20221019-164457
 */

// Package enterpriseusagereportsv1 : Operations and models for the EnterpriseUsageReportsV1 service
package enterpriseusagereportsv1

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

// EnterpriseUsageReportsV1 : Usage reports for IBM Cloud enterprise entities
//
// API Version: 1.0.0-beta.1
type EnterpriseUsageReportsV1 struct {
	Service *core.BaseService
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://enterprise.cloud.ibm.com"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "enterprise_usage_reports"

// EnterpriseUsageReportsV1Options : Service options
type EnterpriseUsageReportsV1Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
}

// NewEnterpriseUsageReportsV1UsingExternalConfig : constructs an instance of EnterpriseUsageReportsV1 with passed in options and external configuration.
func NewEnterpriseUsageReportsV1UsingExternalConfig(options *EnterpriseUsageReportsV1Options) (enterpriseUsageReports *EnterpriseUsageReportsV1, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	enterpriseUsageReports, err = NewEnterpriseUsageReportsV1(options)
	if err != nil {
		return
	}

	err = enterpriseUsageReports.Service.ConfigureService(options.ServiceName)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = enterpriseUsageReports.Service.SetServiceURL(options.URL)
	}
	return
}

// NewEnterpriseUsageReportsV1 : constructs an instance of EnterpriseUsageReportsV1 with passed in options.
func NewEnterpriseUsageReportsV1(options *EnterpriseUsageReportsV1Options) (service *EnterpriseUsageReportsV1, err error) {
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

	service = &EnterpriseUsageReportsV1{
		Service: baseService,
	}

	return
}

// GetServiceURLForRegion returns the service URL to be used for the specified region
func GetServiceURLForRegion(region string) (string, error) {
	return "", fmt.Errorf("service does not support regional URLs")
}

// Clone makes a copy of "enterpriseUsageReports" suitable for processing requests.
func (enterpriseUsageReports *EnterpriseUsageReportsV1) Clone() *EnterpriseUsageReportsV1 {
	if core.IsNil(enterpriseUsageReports) {
		return nil
	}
	clone := *enterpriseUsageReports
	clone.Service = enterpriseUsageReports.Service.Clone()
	return &clone
}

// SetServiceURL sets the service URL
func (enterpriseUsageReports *EnterpriseUsageReportsV1) SetServiceURL(url string) error {
	return enterpriseUsageReports.Service.SetServiceURL(url)
}

// GetServiceURL returns the service URL
func (enterpriseUsageReports *EnterpriseUsageReportsV1) GetServiceURL() string {
	return enterpriseUsageReports.Service.GetServiceURL()
}

// SetDefaultHeaders sets HTTP headers to be sent in every request
func (enterpriseUsageReports *EnterpriseUsageReportsV1) SetDefaultHeaders(headers http.Header) {
	enterpriseUsageReports.Service.SetDefaultHeaders(headers)
}

// SetEnableGzipCompression sets the service's EnableGzipCompression field
func (enterpriseUsageReports *EnterpriseUsageReportsV1) SetEnableGzipCompression(enableGzip bool) {
	enterpriseUsageReports.Service.SetEnableGzipCompression(enableGzip)
}

// GetEnableGzipCompression returns the service's EnableGzipCompression field
func (enterpriseUsageReports *EnterpriseUsageReportsV1) GetEnableGzipCompression() bool {
	return enterpriseUsageReports.Service.GetEnableGzipCompression()
}

// EnableRetries enables automatic retries for requests invoked for this service instance.
// If either parameter is specified as 0, then a default value is used instead.
func (enterpriseUsageReports *EnterpriseUsageReportsV1) EnableRetries(maxRetries int, maxRetryInterval time.Duration) {
	enterpriseUsageReports.Service.EnableRetries(maxRetries, maxRetryInterval)
}

// DisableRetries disables automatic retries for requests invoked for this service instance.
func (enterpriseUsageReports *EnterpriseUsageReportsV1) DisableRetries() {
	enterpriseUsageReports.Service.DisableRetries()
}

// GetResourceUsageReport : Get usage reports for enterprise entities
// Usage reports for entities in the IBM Cloud enterprise. These entities can be the enterprise, an account group, or an
// account.
func (enterpriseUsageReports *EnterpriseUsageReportsV1) GetResourceUsageReport(getResourceUsageReportOptions *GetResourceUsageReportOptions) (result *Reports, response *core.DetailedResponse, err error) {
	return enterpriseUsageReports.GetResourceUsageReportWithContext(context.Background(), getResourceUsageReportOptions)
}

// GetResourceUsageReportWithContext is an alternate form of the GetResourceUsageReport method which supports a Context parameter
func (enterpriseUsageReports *EnterpriseUsageReportsV1) GetResourceUsageReportWithContext(ctx context.Context, getResourceUsageReportOptions *GetResourceUsageReportOptions) (result *Reports, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getResourceUsageReportOptions, "getResourceUsageReportOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = enterpriseUsageReports.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(enterpriseUsageReports.Service.Options.URL, `/v1/resource-usage-reports`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range getResourceUsageReportOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("enterprise_usage_reports", "V1", "GetResourceUsageReport")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if getResourceUsageReportOptions.EnterpriseID != nil {
		builder.AddQuery("enterprise_id", fmt.Sprint(*getResourceUsageReportOptions.EnterpriseID))
	}
	if getResourceUsageReportOptions.AccountGroupID != nil {
		builder.AddQuery("account_group_id", fmt.Sprint(*getResourceUsageReportOptions.AccountGroupID))
	}
	if getResourceUsageReportOptions.AccountID != nil {
		builder.AddQuery("account_id", fmt.Sprint(*getResourceUsageReportOptions.AccountID))
	}
	if getResourceUsageReportOptions.Children != nil {
		builder.AddQuery("children", fmt.Sprint(*getResourceUsageReportOptions.Children))
	}
	if getResourceUsageReportOptions.Month != nil {
		builder.AddQuery("month", fmt.Sprint(*getResourceUsageReportOptions.Month))
	}
	if getResourceUsageReportOptions.BillingUnitID != nil {
		builder.AddQuery("billing_unit_id", fmt.Sprint(*getResourceUsageReportOptions.BillingUnitID))
	}
	if getResourceUsageReportOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*getResourceUsageReportOptions.Limit))
	}
	if getResourceUsageReportOptions.Offset != nil {
		builder.AddQuery("offset", fmt.Sprint(*getResourceUsageReportOptions.Offset))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = enterpriseUsageReports.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalReports)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetResourceUsageReportOptions : The GetResourceUsageReport options.
type GetResourceUsageReportOptions struct {
	// The ID of the enterprise for which the reports are queried. This parameter cannot be used with the `account_id` or
	// `account_group_id` query parameters.
	EnterpriseID *string `json:"enterprise_id,omitempty"`

	// The ID of the account group for which the reports are queried. This parameter cannot be used with the `account_id`
	// or `enterprise_id` query parameters.
	AccountGroupID *string `json:"account_group_id,omitempty"`

	// The ID of the account for which the reports are queried. This parameter cannot be used with the `account_group_id`
	// or `enterprise_id` query parameters.
	AccountID *string `json:"account_id,omitempty"`

	// Returns the reports for the immediate child entities under the current account group or enterprise. This parameter
	// cannot be used with the `account_id` query parameter.
	Children *bool `json:"children,omitempty"`

	// The billing month for which the usage report is requested. The format is in yyyy-mm. Defaults to the month in which
	// the report is queried.
	Month *string `json:"month,omitempty"`

	// The ID of the billing unit by which to filter the reports.
	BillingUnitID *string `json:"billing_unit_id,omitempty"`

	// The maximum number of search results to be returned.
	Limit *int64 `json:"limit,omitempty"`

	// An opaque value representing the offset of the first item to be returned by a search query. If not specified, then
	// the first page of results is returned. To retrieve the next page of search results, use the 'offset' query parameter
	// value within the 'next.href' URL found within a prior search query response.
	Offset *string `json:"offset,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetResourceUsageReportOptions : Instantiate GetResourceUsageReportOptions
func (*EnterpriseUsageReportsV1) NewGetResourceUsageReportOptions() *GetResourceUsageReportOptions {
	return &GetResourceUsageReportOptions{}
}

// SetEnterpriseID : Allow user to set EnterpriseID
func (_options *GetResourceUsageReportOptions) SetEnterpriseID(enterpriseID string) *GetResourceUsageReportOptions {
	_options.EnterpriseID = core.StringPtr(enterpriseID)
	return _options
}

// SetAccountGroupID : Allow user to set AccountGroupID
func (_options *GetResourceUsageReportOptions) SetAccountGroupID(accountGroupID string) *GetResourceUsageReportOptions {
	_options.AccountGroupID = core.StringPtr(accountGroupID)
	return _options
}

// SetAccountID : Allow user to set AccountID
func (_options *GetResourceUsageReportOptions) SetAccountID(accountID string) *GetResourceUsageReportOptions {
	_options.AccountID = core.StringPtr(accountID)
	return _options
}

// SetChildren : Allow user to set Children
func (_options *GetResourceUsageReportOptions) SetChildren(children bool) *GetResourceUsageReportOptions {
	_options.Children = core.BoolPtr(children)
	return _options
}

// SetMonth : Allow user to set Month
func (_options *GetResourceUsageReportOptions) SetMonth(month string) *GetResourceUsageReportOptions {
	_options.Month = core.StringPtr(month)
	return _options
}

// SetBillingUnitID : Allow user to set BillingUnitID
func (_options *GetResourceUsageReportOptions) SetBillingUnitID(billingUnitID string) *GetResourceUsageReportOptions {
	_options.BillingUnitID = core.StringPtr(billingUnitID)
	return _options
}

// SetLimit : Allow user to set Limit
func (_options *GetResourceUsageReportOptions) SetLimit(limit int64) *GetResourceUsageReportOptions {
	_options.Limit = core.Int64Ptr(limit)
	return _options
}

// SetOffset : Allow user to set Offset
func (_options *GetResourceUsageReportOptions) SetOffset(offset string) *GetResourceUsageReportOptions {
	_options.Offset = core.StringPtr(offset)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetResourceUsageReportOptions) SetHeaders(param map[string]string) *GetResourceUsageReportOptions {
	options.Headers = param
	return options
}

// Link : An object that contains a link to a page of search results.
type Link struct {
	// A link to a page of search results.
	Href *string `json:"href,omitempty"`
}

// UnmarshalLink unmarshals an instance of Link from the specified map of raw messages.
func UnmarshalLink(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Link)
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MetricUsage : An object that represents a metric.
type MetricUsage struct {
	// The name of the metric.
	Metric *string `json:"metric" validate:"required"`

	// A unit to qualify the quantity.
	Unit *string `json:"unit" validate:"required"`

	// The aggregated value for the metric.
	Quantity *float64 `json:"quantity" validate:"required"`

	// The quantity that is used for calculating charges.
	RateableQuantity *float64 `json:"rateable_quantity" validate:"required"`

	// The cost that was incurred by the metric.
	Cost *float64 `json:"cost" validate:"required"`

	// The pre-discounted cost that was incurred by the metric.
	RatedCost *float64 `json:"rated_cost" validate:"required"`

	// The price with which cost was calculated.
	Price []map[string]interface{} `json:"price,omitempty"`
}

// UnmarshalMetricUsage unmarshals an instance of MetricUsage from the specified map of raw messages.
func UnmarshalMetricUsage(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MetricUsage)
	err = core.UnmarshalPrimitive(m, "metric", &obj.Metric)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "unit", &obj.Unit)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "quantity", &obj.Quantity)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "rateable_quantity", &obj.RateableQuantity)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "cost", &obj.Cost)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "rated_cost", &obj.RatedCost)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "price", &obj.Price)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PlanUsage : Aggregated values for the plan.
type PlanUsage struct {
	// The ID of the plan.
	PlanID *string `json:"plan_id" validate:"required"`

	// The pricing region for the plan.
	PricingRegion *string `json:"pricing_region,omitempty"`

	// The pricing plan with which the usage was rated.
	PricingPlanID *string `json:"pricing_plan_id,omitempty"`

	// Whether the plan charges are billed to the customer.
	Billable *bool `json:"billable" validate:"required"`

	// The total cost that was incurred by the plan.
	Cost *float64 `json:"cost" validate:"required"`

	// The total pre-discounted cost that was incurred by the plan.
	RatedCost *float64 `json:"rated_cost" validate:"required"`

	// All of the metrics in the plan.
	Usage []MetricUsage `json:"usage" validate:"required"`
}

// UnmarshalPlanUsage unmarshals an instance of PlanUsage from the specified map of raw messages.
func UnmarshalPlanUsage(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PlanUsage)
	err = core.UnmarshalPrimitive(m, "plan_id", &obj.PlanID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "pricing_region", &obj.PricingRegion)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "pricing_plan_id", &obj.PricingPlanID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "billable", &obj.Billable)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "cost", &obj.Cost)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "rated_cost", &obj.RatedCost)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "usage", &obj.Usage, UnmarshalMetricUsage)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Reports : Resource Usage Reports API response.
type Reports struct {
	// The maximum number of reports in the response.
	Limit *int64 `json:"limit,omitempty"`

	// An object that contains the link to the first page of the search query.
	First *Link `json:"first,omitempty"`

	// An object that contains the link to the next page of the search query.
	Next *Link `json:"next,omitempty"`

	// The list of usage reports.
	Reports []ResourceUsageReport `json:"reports,omitempty"`
}

// UnmarshalReports unmarshals an instance of Reports from the specified map of raw messages.
func UnmarshalReports(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Reports)
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "first", &obj.First, UnmarshalLink)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "next", &obj.Next, UnmarshalLink)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "reports", &obj.Reports, UnmarshalResourceUsageReport)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Retrieve the value to be passed to a request to access the next page of results
func (resp *Reports) GetNextOffset() (*string, error) {
	if core.IsNil(resp.Next) {
		return nil, nil
	}
	offset, err := core.GetQueryParam(resp.Next.Href, "offset")
	if err != nil || offset == nil {
		return nil, err
	}
	return offset, nil
}

// ResourceUsage : A container for all the plans in the resource.
type ResourceUsage struct {
	// The ID of the resource.
	ResourceID *string `json:"resource_id" validate:"required"`

	// The billable charges for the account.
	BillableCost *float64 `json:"billable_cost" validate:"required"`

	// The pre-discounted billable charges for the account.
	BillableRatedCost *float64 `json:"billable_rated_cost" validate:"required"`

	// The non-billable charges for the account.
	NonBillableCost *float64 `json:"non_billable_cost" validate:"required"`

	// The pre-discounted, non-billable charges for the account.
	NonBillableRatedCost *float64 `json:"non_billable_rated_cost" validate:"required"`

	// All of the plans in the resource.
	Plans []PlanUsage `json:"plans" validate:"required"`
}

// UnmarshalResourceUsage unmarshals an instance of ResourceUsage from the specified map of raw messages.
func UnmarshalResourceUsage(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ResourceUsage)
	err = core.UnmarshalPrimitive(m, "resource_id", &obj.ResourceID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "billable_cost", &obj.BillableCost)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "billable_rated_cost", &obj.BillableRatedCost)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "non_billable_cost", &obj.NonBillableCost)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "non_billable_rated_cost", &obj.NonBillableRatedCost)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "plans", &obj.Plans, UnmarshalPlanUsage)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ResourceUsageReport : An object that represents a usage report.
type ResourceUsageReport struct {
	// The ID of the entity.
	EntityID *string `json:"entity_id" validate:"required"`

	// The entity type.
	EntityType *string `json:"entity_type" validate:"required"`

	// The Cloud Resource Name (CRN) of the entity towards which the resource usages were rolled up.
	EntityCRN *string `json:"entity_crn" validate:"required"`

	// A user-defined name for the entity, such as the enterprise name or account group name.
	EntityName *string `json:"entity_name" validate:"required"`

	// The ID of the billing unit.
	BillingUnitID *string `json:"billing_unit_id" validate:"required"`

	// The CRN of the billing unit.
	BillingUnitCRN *string `json:"billing_unit_crn" validate:"required"`

	// The name of the billing unit.
	BillingUnitName *string `json:"billing_unit_name" validate:"required"`

	// The country code of the billing unit.
	CountryCode *string `json:"country_code" validate:"required"`

	// The currency code of the billing unit.
	CurrencyCode *string `json:"currency_code" validate:"required"`

	// Billing month.
	Month *string `json:"month" validate:"required"`

	// Billable charges that are aggregated from all entities in the report.
	BillableCost *float64 `json:"billable_cost" validate:"required"`

	// Non-billable charges that are aggregated from all entities in the report.
	NonBillableCost *float64 `json:"non_billable_cost" validate:"required"`

	// Aggregated billable charges before discounts.
	BillableRatedCost *float64 `json:"billable_rated_cost" validate:"required"`

	// Aggregated non-billable charges before discounts.
	NonBillableRatedCost *float64 `json:"non_billable_rated_cost" validate:"required"`

	// Details about all the resources that are included in the aggregated charges.
	Resources []ResourceUsage `json:"resources" validate:"required"`
}

// Constants associated with the ResourceUsageReport.EntityType property.
// The entity type.
const (
	ResourceUsageReportEntityTypeAccountConst      = "account"
	ResourceUsageReportEntityTypeAccountGroupConst = "account-group"
	ResourceUsageReportEntityTypeEnterpriseConst   = "enterprise"
)

// UnmarshalResourceUsageReport unmarshals an instance of ResourceUsageReport from the specified map of raw messages.
func UnmarshalResourceUsageReport(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ResourceUsageReport)
	err = core.UnmarshalPrimitive(m, "entity_id", &obj.EntityID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "entity_type", &obj.EntityType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "entity_crn", &obj.EntityCRN)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "entity_name", &obj.EntityName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "billing_unit_id", &obj.BillingUnitID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "billing_unit_crn", &obj.BillingUnitCRN)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "billing_unit_name", &obj.BillingUnitName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "country_code", &obj.CountryCode)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "currency_code", &obj.CurrencyCode)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "month", &obj.Month)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "billable_cost", &obj.BillableCost)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "non_billable_cost", &obj.NonBillableCost)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "billable_rated_cost", &obj.BillableRatedCost)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "non_billable_rated_cost", &obj.NonBillableRatedCost)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "resources", &obj.Resources, UnmarshalResourceUsage)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetResourceUsageReportPager can be used to simplify the use of the "GetResourceUsageReport" method.
type GetResourceUsageReportPager struct {
	hasNext     bool
	options     *GetResourceUsageReportOptions
	client      *EnterpriseUsageReportsV1
	pageContext struct {
		next *string
	}
}

// NewGetResourceUsageReportPager returns a new GetResourceUsageReportPager instance.
func (enterpriseUsageReports *EnterpriseUsageReportsV1) NewGetResourceUsageReportPager(options *GetResourceUsageReportOptions) (pager *GetResourceUsageReportPager, err error) {
	if options.Offset != nil && *options.Offset != "" {
		err = fmt.Errorf("the 'options.Offset' field should not be set")
		return
	}

	var optionsCopy GetResourceUsageReportOptions = *options
	pager = &GetResourceUsageReportPager{
		hasNext: true,
		options: &optionsCopy,
		client:  enterpriseUsageReports,
	}
	return
}

// HasNext returns true if there are potentially more results to be retrieved.
func (pager *GetResourceUsageReportPager) HasNext() bool {
	return pager.hasNext
}

// GetNextWithContext returns the next page of results using the specified Context.
func (pager *GetResourceUsageReportPager) GetNextWithContext(ctx context.Context) (page []ResourceUsageReport, err error) {
	if !pager.HasNext() {
		return nil, fmt.Errorf("no more results available")
	}

	pager.options.Offset = pager.pageContext.next

	result, _, err := pager.client.GetResourceUsageReportWithContext(ctx, pager.options)
	if err != nil {
		return
	}

	var next *string
	if result.Next != nil {
		var offset *string
		offset, err = core.GetQueryParam(result.Next.Href, "offset")
		if err != nil {
			err = fmt.Errorf("error retrieving 'offset' query parameter from URL '%s': %s", *result.Next.Href, err.Error())
			return
		}
		next = offset
	}
	pager.pageContext.next = next
	pager.hasNext = (pager.pageContext.next != nil)
	page = result.Reports

	return
}

// GetAllWithContext returns all results by invoking GetNextWithContext() repeatedly
// until all pages of results have been retrieved.
func (pager *GetResourceUsageReportPager) GetAllWithContext(ctx context.Context) (allItems []ResourceUsageReport, err error) {
	for pager.HasNext() {
		var nextPage []ResourceUsageReport
		nextPage, err = pager.GetNextWithContext(ctx)
		if err != nil {
			return
		}
		allItems = append(allItems, nextPage...)
	}
	return
}

// GetNext invokes GetNextWithContext() using context.Background() as the Context parameter.
func (pager *GetResourceUsageReportPager) GetNext() (page []ResourceUsageReport, err error) {
	return pager.GetNextWithContext(context.Background())
}

// GetAll invokes GetAllWithContext() using context.Background() as the Context parameter.
func (pager *GetResourceUsageReportPager) GetAll() (allItems []ResourceUsageReport, err error) {
	return pager.GetAllWithContext(context.Background())
}
