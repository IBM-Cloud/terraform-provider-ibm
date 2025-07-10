/**
 * (C) Copyright IBM Corp. 2024.
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
 * IBM OpenAPI SDK Code Generator Version: 3.96.0-d6dec9d7-20241008-212902
 */

// Package partnermanagementv1 : Operations and models for the PartnerManagementV1 service
package partnermanagementv1

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	common "github.com/IBM/platform-services-go-sdk/common"
	"github.com/go-openapi/strfmt"
)

// PartnerManagementV1 : The Partner Management APIs enable you to manage the IBM Cloud partner entities and fetch
// multiple reports in different formats.
//
// API Version: 1.0.0
type PartnerManagementV1 struct {
	Service *core.BaseService
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://partner.cloud.ibm.com"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "partner_management"

// PartnerManagementV1Options : Service options
type PartnerManagementV1Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
}

// NewPartnerManagementV1UsingExternalConfig : constructs an instance of PartnerManagementV1 with passed in options and external configuration.
func NewPartnerManagementV1UsingExternalConfig(options *PartnerManagementV1Options) (partnerManagement *PartnerManagementV1, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			err = core.SDKErrorf(err, "", "env-auth-error", common.GetComponentInfo())
			return
		}
	}

	partnerManagement, err = NewPartnerManagementV1(options)
	err = core.RepurposeSDKProblem(err, "new-client-error")
	if err != nil {
		return
	}

	err = partnerManagement.Service.ConfigureService(options.ServiceName)
	if err != nil {
		err = core.SDKErrorf(err, "", "client-config-error", common.GetComponentInfo())
		return
	}

	if options.URL != "" {
		err = partnerManagement.Service.SetServiceURL(options.URL)
		err = core.RepurposeSDKProblem(err, "url-set-error")
	}
	return
}

// NewPartnerManagementV1 : constructs an instance of PartnerManagementV1 with passed in options.
func NewPartnerManagementV1(options *PartnerManagementV1Options) (service *PartnerManagementV1, err error) {
	serviceOptions := &core.ServiceOptions{
		URL:           DefaultServiceURL,
		Authenticator: options.Authenticator,
	}

	baseService, err := core.NewBaseService(serviceOptions)
	if err != nil {
		err = core.SDKErrorf(err, "", "new-base-error", common.GetComponentInfo())
		return
	}

	if options.URL != "" {
		err = baseService.SetServiceURL(options.URL)
		if err != nil {
			err = core.SDKErrorf(err, "", "set-url-error", common.GetComponentInfo())
			return
		}
	}

	service = &PartnerManagementV1{
		Service: baseService,
	}

	return
}

// GetServiceURLForRegion returns the service URL to be used for the specified region
func GetServiceURLForRegion(region string) (string, error) {
	return "", core.SDKErrorf(nil, "service does not support regional URLs", "no-regional-support", common.GetComponentInfo())
}

// Clone makes a copy of "partnerManagement" suitable for processing requests.
func (partnerManagement *PartnerManagementV1) Clone() *PartnerManagementV1 {
	if core.IsNil(partnerManagement) {
		return nil
	}
	clone := *partnerManagement
	clone.Service = partnerManagement.Service.Clone()
	return &clone
}

// SetServiceURL sets the service URL
func (partnerManagement *PartnerManagementV1) SetServiceURL(url string) error {
	err := partnerManagement.Service.SetServiceURL(url)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-set-error", common.GetComponentInfo())
	}
	return err
}

// GetServiceURL returns the service URL
func (partnerManagement *PartnerManagementV1) GetServiceURL() string {
	return partnerManagement.Service.GetServiceURL()
}

// SetDefaultHeaders sets HTTP headers to be sent in every request
func (partnerManagement *PartnerManagementV1) SetDefaultHeaders(headers http.Header) {
	partnerManagement.Service.SetDefaultHeaders(headers)
}

// SetEnableGzipCompression sets the service's EnableGzipCompression field
func (partnerManagement *PartnerManagementV1) SetEnableGzipCompression(enableGzip bool) {
	partnerManagement.Service.SetEnableGzipCompression(enableGzip)
}

// GetEnableGzipCompression returns the service's EnableGzipCompression field
func (partnerManagement *PartnerManagementV1) GetEnableGzipCompression() bool {
	return partnerManagement.Service.GetEnableGzipCompression()
}

// EnableRetries enables automatic retries for requests invoked for this service instance.
// If either parameter is specified as 0, then a default value is used instead.
func (partnerManagement *PartnerManagementV1) EnableRetries(maxRetries int, maxRetryInterval time.Duration) {
	partnerManagement.Service.EnableRetries(maxRetries, maxRetryInterval)
}

// DisableRetries disables automatic retries for requests invoked for this service instance.
func (partnerManagement *PartnerManagementV1) DisableRetries() {
	partnerManagement.Service.DisableRetries()
}

// GetResourceUsageReport : Get partner resource usage report
// Returns the summary for the partner for a given month. Partner billing managers are authorized to access this report.
func (partnerManagement *PartnerManagementV1) GetResourceUsageReport(getResourceUsageReportOptions *GetResourceUsageReportOptions) (result *PartnerUsageReportSummary, response *core.DetailedResponse, err error) {
	result, response, err = partnerManagement.GetResourceUsageReportWithContext(context.Background(), getResourceUsageReportOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetResourceUsageReportWithContext is an alternate form of the GetResourceUsageReport method which supports a Context parameter
func (partnerManagement *PartnerManagementV1) GetResourceUsageReportWithContext(ctx context.Context, getResourceUsageReportOptions *GetResourceUsageReportOptions) (result *PartnerUsageReportSummary, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getResourceUsageReportOptions, "getResourceUsageReportOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(getResourceUsageReportOptions, "getResourceUsageReportOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = partnerManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(partnerManagement.Service.Options.URL, `/v1/resource-usage-reports`, nil)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range getResourceUsageReportOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("partner_management", "V1", "GetResourceUsageReport")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("partner_id", fmt.Sprint(*getResourceUsageReportOptions.PartnerID))
	if getResourceUsageReportOptions.ResellerID != nil {
		builder.AddQuery("reseller_id", fmt.Sprint(*getResourceUsageReportOptions.ResellerID))
	}
	if getResourceUsageReportOptions.CustomerID != nil {
		builder.AddQuery("customer_id", fmt.Sprint(*getResourceUsageReportOptions.CustomerID))
	}
	if getResourceUsageReportOptions.Children != nil {
		builder.AddQuery("children", fmt.Sprint(*getResourceUsageReportOptions.Children))
	}
	if getResourceUsageReportOptions.Month != nil {
		builder.AddQuery("month", fmt.Sprint(*getResourceUsageReportOptions.Month))
	}
	if getResourceUsageReportOptions.Viewpoint != nil {
		builder.AddQuery("viewpoint", fmt.Sprint(*getResourceUsageReportOptions.Viewpoint))
	}
	if getResourceUsageReportOptions.Recurse != nil {
		builder.AddQuery("recurse", fmt.Sprint(*getResourceUsageReportOptions.Recurse))
	}
	if getResourceUsageReportOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*getResourceUsageReportOptions.Limit))
	}
	if getResourceUsageReportOptions.Offset != nil {
		builder.AddQuery("offset", fmt.Sprint(*getResourceUsageReportOptions.Offset))
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = partnerManagement.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_resource_usage_report", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalPartnerUsageReportSummary)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// GetBillingOptions : Get customers billing options
// Returns the billing options for the requested customer for a given month.
func (partnerManagement *PartnerManagementV1) GetBillingOptions(getBillingOptionsOptions *GetBillingOptionsOptions) (result *BillingOptionsSummary, response *core.DetailedResponse, err error) {
	result, response, err = partnerManagement.GetBillingOptionsWithContext(context.Background(), getBillingOptionsOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetBillingOptionsWithContext is an alternate form of the GetBillingOptions method which supports a Context parameter
func (partnerManagement *PartnerManagementV1) GetBillingOptionsWithContext(ctx context.Context, getBillingOptionsOptions *GetBillingOptionsOptions) (result *BillingOptionsSummary, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getBillingOptionsOptions, "getBillingOptionsOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(getBillingOptionsOptions, "getBillingOptionsOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = partnerManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(partnerManagement.Service.Options.URL, `/v1/billing-options`, nil)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range getBillingOptionsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("partner_management", "V1", "GetBillingOptions")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("partner_id", fmt.Sprint(*getBillingOptionsOptions.PartnerID))
	if getBillingOptionsOptions.CustomerID != nil {
		builder.AddQuery("customer_id", fmt.Sprint(*getBillingOptionsOptions.CustomerID))
	}
	if getBillingOptionsOptions.ResellerID != nil {
		builder.AddQuery("reseller_id", fmt.Sprint(*getBillingOptionsOptions.ResellerID))
	}
	if getBillingOptionsOptions.Date != nil {
		builder.AddQuery("date", fmt.Sprint(*getBillingOptionsOptions.Date))
	}
	if getBillingOptionsOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*getBillingOptionsOptions.Limit))
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = partnerManagement.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_billing_options", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalBillingOptionsSummary)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// GetCreditPoolsReport : Get credit pools report
// Returns the subscription or commitment burn-down reports for the end customers for a given month.
func (partnerManagement *PartnerManagementV1) GetCreditPoolsReport(getCreditPoolsReportOptions *GetCreditPoolsReportOptions) (result *CreditPoolsReportSummary, response *core.DetailedResponse, err error) {
	result, response, err = partnerManagement.GetCreditPoolsReportWithContext(context.Background(), getCreditPoolsReportOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetCreditPoolsReportWithContext is an alternate form of the GetCreditPoolsReport method which supports a Context parameter
func (partnerManagement *PartnerManagementV1) GetCreditPoolsReportWithContext(ctx context.Context, getCreditPoolsReportOptions *GetCreditPoolsReportOptions) (result *CreditPoolsReportSummary, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getCreditPoolsReportOptions, "getCreditPoolsReportOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(getCreditPoolsReportOptions, "getCreditPoolsReportOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = partnerManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(partnerManagement.Service.Options.URL, `/v1/credit-pools`, nil)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range getCreditPoolsReportOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("partner_management", "V1", "GetCreditPoolsReport")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("partner_id", fmt.Sprint(*getCreditPoolsReportOptions.PartnerID))
	if getCreditPoolsReportOptions.CustomerID != nil {
		builder.AddQuery("customer_id", fmt.Sprint(*getCreditPoolsReportOptions.CustomerID))
	}
	if getCreditPoolsReportOptions.ResellerID != nil {
		builder.AddQuery("reseller_id", fmt.Sprint(*getCreditPoolsReportOptions.ResellerID))
	}
	if getCreditPoolsReportOptions.Date != nil {
		builder.AddQuery("date", fmt.Sprint(*getCreditPoolsReportOptions.Date))
	}
	if getCreditPoolsReportOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*getCreditPoolsReportOptions.Limit))
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = partnerManagement.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_credit_pools_report", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCreditPoolsReportSummary)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}
func getServiceComponentInfo() *core.ProblemComponent {
	return core.NewProblemComponent(DefaultServiceName, "1.0.0")
}

// BillingOption : Billing options report for the end customers.
type BillingOption struct {
	// The ID of the billing option.
	ID *string `json:"id,omitempty"`

	// The ID of the billing unit that's associated with the billing option.
	BillingUnitID *string `json:"billing_unit_id,omitempty"`

	// Account ID of the customer.
	CustomerID *string `json:"customer_id,omitempty"`

	// The customer type. The valid values are `ENTERPRISE`, `ACCOUNT`, and `ACCOUNT_GROUP`.
	CustomerType *string `json:"customer_type,omitempty"`

	// A user-defined name for the customer.
	CustomerName *string `json:"customer_name,omitempty"`

	// ID of the reseller in the heirarchy of the requested customer.
	ResellerID *string `json:"reseller_id,omitempty"`

	// Name of the reseller in the heirarchy of the requested customer.
	ResellerName *string `json:"reseller_name,omitempty"`

	// The billing month for which the burn-down report is requested. Format is yyyy-mm. Defaults to current month.
	Month *string `json:"month,omitempty"`

	// Errors in the billing.
	Errors []map[string]interface{} `json:"errors,omitempty"`

	// The type of billing option. The valid values are `SUBSCRIPTION` and `OFFER`.
	Type *string `json:"type,omitempty"`

	// The start date of billing option.
	StartDate *strfmt.DateTime `json:"start_date,omitempty"`

	// The end date of billing option.
	EndDate *strfmt.DateTime `json:"end_date,omitempty"`

	// The state of the billing option. The valid values include `ACTIVE, `SUSPENDED`, and `CANCELED`.
	State *string `json:"state,omitempty"`

	// The category of the billing option. The valid values are `PLATFORM`, `SERVICE`, and `SUPPORT`.
	Category *string `json:"category,omitempty"`

	// The payment method for support.
	PaymentInstrument map[string]interface{} `json:"payment_instrument,omitempty"`

	// Part number of the offering.
	PartNumber *string `json:"part_number,omitempty"`

	// ID of the catalog containing this offering.
	CatalogID *string `json:"catalog_id,omitempty"`

	// ID of the order containing this offering.
	OrderID *string `json:"order_id,omitempty"`

	// PO Number of the offering.
	PoNumber *string `json:"po_number,omitempty"`

	// Subscription model.
	SubscriptionModel *string `json:"subscription_model,omitempty"`

	// The duration of the billing options in months.
	DurationInMonths *int64 `json:"duration_in_months,omitempty"`

	// Amount billed monthly for this offering.
	MonthlyAmount *float64 `json:"monthly_amount,omitempty"`

	// The support billing system.
	BillingSystem map[string]interface{} `json:"billing_system,omitempty"`

	// The country code for the billing unit.
	CountryCode *string `json:"country_code,omitempty"`

	// The currency code of the billing unit.
	CurrencyCode *string `json:"currency_code,omitempty"`
}

// Constants associated with the BillingOption.CustomerType property.
// The customer type. The valid values are `ENTERPRISE`, `ACCOUNT`, and `ACCOUNT_GROUP`.
const (
	BillingOptionCustomerTypeAccountConst      = "ACCOUNT"
	BillingOptionCustomerTypeAccountGroupConst = "ACCOUNT_GROUP"
	BillingOptionCustomerTypeEnterpriseConst   = "ENTERPRISE"
)

// Constants associated with the BillingOption.Type property.
// The type of billing option. The valid values are `SUBSCRIPTION` and `OFFER`.
const (
	BillingOptionTypeOfferConst        = "OFFER"
	BillingOptionTypeSubscriptionConst = "SUBSCRIPTION"
)

// Constants associated with the BillingOption.State property.
// The state of the billing option. The valid values include `ACTIVE, `SUSPENDED`, and `CANCELED`.
const (
	BillingOptionStateActiveConst    = "ACTIVE"
	BillingOptionStateCanceledConst  = "CANCELED"
	BillingOptionStateSuspendedConst = "SUSPENDED"
)

// Constants associated with the BillingOption.Category property.
// The category of the billing option. The valid values are `PLATFORM`, `SERVICE`, and `SUPPORT`.
const (
	BillingOptionCategoryPlatformConst = "PLATFORM"
	BillingOptionCategoryServiceConst  = "SERVICE"
	BillingOptionCategorySupportConst  = "SUPPORT"
)

// UnmarshalBillingOption unmarshals an instance of BillingOption from the specified map of raw messages.
func UnmarshalBillingOption(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(BillingOption)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "billing_unit_id", &obj.BillingUnitID)
	if err != nil {
		err = core.SDKErrorf(err, "", "billing_unit_id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "customer_id", &obj.CustomerID)
	if err != nil {
		err = core.SDKErrorf(err, "", "customer_id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "customer_type", &obj.CustomerType)
	if err != nil {
		err = core.SDKErrorf(err, "", "customer_type-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "customer_name", &obj.CustomerName)
	if err != nil {
		err = core.SDKErrorf(err, "", "customer_name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "reseller_id", &obj.ResellerID)
	if err != nil {
		err = core.SDKErrorf(err, "", "reseller_id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "reseller_name", &obj.ResellerName)
	if err != nil {
		err = core.SDKErrorf(err, "", "reseller_name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "month", &obj.Month)
	if err != nil {
		err = core.SDKErrorf(err, "", "month-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "errors", &obj.Errors)
	if err != nil {
		err = core.SDKErrorf(err, "", "errors-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		err = core.SDKErrorf(err, "", "type-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "start_date", &obj.StartDate)
	if err != nil {
		err = core.SDKErrorf(err, "", "start_date-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "end_date", &obj.EndDate)
	if err != nil {
		err = core.SDKErrorf(err, "", "end_date-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "state", &obj.State)
	if err != nil {
		err = core.SDKErrorf(err, "", "state-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "category", &obj.Category)
	if err != nil {
		err = core.SDKErrorf(err, "", "category-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "payment_instrument", &obj.PaymentInstrument)
	if err != nil {
		err = core.SDKErrorf(err, "", "payment_instrument-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "part_number", &obj.PartNumber)
	if err != nil {
		err = core.SDKErrorf(err, "", "part_number-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "catalog_id", &obj.CatalogID)
	if err != nil {
		err = core.SDKErrorf(err, "", "catalog_id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "order_id", &obj.OrderID)
	if err != nil {
		err = core.SDKErrorf(err, "", "order_id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "po_number", &obj.PoNumber)
	if err != nil {
		err = core.SDKErrorf(err, "", "po_number-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "subscription_model", &obj.SubscriptionModel)
	if err != nil {
		err = core.SDKErrorf(err, "", "subscription_model-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "duration_in_months", &obj.DurationInMonths)
	if err != nil {
		err = core.SDKErrorf(err, "", "duration_in_months-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "monthly_amount", &obj.MonthlyAmount)
	if err != nil {
		err = core.SDKErrorf(err, "", "monthly_amount-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "billing_system", &obj.BillingSystem)
	if err != nil {
		err = core.SDKErrorf(err, "", "billing_system-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "country_code", &obj.CountryCode)
	if err != nil {
		err = core.SDKErrorf(err, "", "country_code-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "currency_code", &obj.CurrencyCode)
	if err != nil {
		err = core.SDKErrorf(err, "", "currency_code-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// BillingOptionsSummary : The billing options report for the customer.
type BillingOptionsSummary struct {
	// The max number of reports in the response.
	RowsCount *int64 `json:"rows_count,omitempty"`

	// The link to the next page of the search query.
	NextURL *string `json:"next_url,omitempty"`

	// Aggregated usage report of all requested partners.
	Resources []BillingOption `json:"resources,omitempty"`
}

// UnmarshalBillingOptionsSummary unmarshals an instance of BillingOptionsSummary from the specified map of raw messages.
func UnmarshalBillingOptionsSummary(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(BillingOptionsSummary)
	err = core.UnmarshalPrimitive(m, "rows_count", &obj.RowsCount)
	if err != nil {
		err = core.SDKErrorf(err, "", "rows_count-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "next_url", &obj.NextURL)
	if err != nil {
		err = core.SDKErrorf(err, "", "next_url-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "resources", &obj.Resources, UnmarshalBillingOption)
	if err != nil {
		err = core.SDKErrorf(err, "", "resources-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CreditPoolsReport : Aggregated subscription burn-down report for the end customers.
type CreditPoolsReport struct {
	// The category of the billing option. The valid values are `PLATFORM`, `SERVICE` and `SUPPORT`.
	Type *string `json:"type,omitempty"`

	// The ID of the billing unit that's associated with the billing option.
	BillingUnitID *string `json:"billing_unit_id,omitempty"`

	// Account ID of the customer.
	CustomerID *string `json:"customer_id,omitempty"`

	// The customer type. The valid values are `ENTERPRISE`, `ACCOUNT`, and `ACCOUNT_GROUP`.
	CustomerType *string `json:"customer_type,omitempty"`

	// A user-defined name for the customer.
	CustomerName *string `json:"customer_name,omitempty"`

	// ID of the reseller in the heirarchy of the requested customer.
	ResellerID *string `json:"reseller_id,omitempty"`

	// Name of the reseller in the heirarchy of the requested customer.
	ResellerName *string `json:"reseller_name,omitempty"`

	// The billing month for which the burn-down report is requested. Format is yyyy-mm. Defaults to current month.
	Month *string `json:"month,omitempty"`

	// The currency code of the billing unit.
	CurrencyCode *string `json:"currency_code,omitempty"`

	// A list of active subscription terms available within a credit.
	TermCredits []TermCredits `json:"term_credits,omitempty"`

	// Overage that was generated on the credit pool.
	Overage *Overage `json:"overage,omitempty"`
}

// Constants associated with the CreditPoolsReport.Type property.
// The category of the billing option. The valid values are `PLATFORM`, `SERVICE` and `SUPPORT`.
const (
	CreditPoolsReportTypePlatformConst = "PLATFORM"
	CreditPoolsReportTypeServiceConst  = "SERVICE"
	CreditPoolsReportTypeSupportConst  = "SUPPORT"
)

// Constants associated with the CreditPoolsReport.CustomerType property.
// The customer type. The valid values are `ENTERPRISE`, `ACCOUNT`, and `ACCOUNT_GROUP`.
const (
	CreditPoolsReportCustomerTypeAccountConst      = "ACCOUNT"
	CreditPoolsReportCustomerTypeAccountGroupConst = "ACCOUNT_GROUP"
	CreditPoolsReportCustomerTypeEnterpriseConst   = "ENTERPRISE"
)

// UnmarshalCreditPoolsReport unmarshals an instance of CreditPoolsReport from the specified map of raw messages.
func UnmarshalCreditPoolsReport(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CreditPoolsReport)
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		err = core.SDKErrorf(err, "", "type-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "billing_unit_id", &obj.BillingUnitID)
	if err != nil {
		err = core.SDKErrorf(err, "", "billing_unit_id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "customer_id", &obj.CustomerID)
	if err != nil {
		err = core.SDKErrorf(err, "", "customer_id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "customer_type", &obj.CustomerType)
	if err != nil {
		err = core.SDKErrorf(err, "", "customer_type-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "customer_name", &obj.CustomerName)
	if err != nil {
		err = core.SDKErrorf(err, "", "customer_name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "reseller_id", &obj.ResellerID)
	if err != nil {
		err = core.SDKErrorf(err, "", "reseller_id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "reseller_name", &obj.ResellerName)
	if err != nil {
		err = core.SDKErrorf(err, "", "reseller_name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "month", &obj.Month)
	if err != nil {
		err = core.SDKErrorf(err, "", "month-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "currency_code", &obj.CurrencyCode)
	if err != nil {
		err = core.SDKErrorf(err, "", "currency_code-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "term_credits", &obj.TermCredits, UnmarshalTermCredits)
	if err != nil {
		err = core.SDKErrorf(err, "", "term_credits-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "overage", &obj.Overage, UnmarshalOverage)
	if err != nil {
		err = core.SDKErrorf(err, "", "overage-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CreditPoolsReportSummary : The aggregated credit pools report.
type CreditPoolsReportSummary struct {
	// The max number of reports in the response.
	RowsCount *int64 `json:"rows_count,omitempty"`

	// The link to the next page of the search query.
	NextURL *string `json:"next_url,omitempty"`

	// Aggregated usage report of all requested partners.
	Resources []CreditPoolsReport `json:"resources,omitempty"`
}

// UnmarshalCreditPoolsReportSummary unmarshals an instance of CreditPoolsReportSummary from the specified map of raw messages.
func UnmarshalCreditPoolsReportSummary(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CreditPoolsReportSummary)
	err = core.UnmarshalPrimitive(m, "rows_count", &obj.RowsCount)
	if err != nil {
		err = core.SDKErrorf(err, "", "rows_count-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "next_url", &obj.NextURL)
	if err != nil {
		err = core.SDKErrorf(err, "", "next_url-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "resources", &obj.Resources, UnmarshalCreditPoolsReport)
	if err != nil {
		err = core.SDKErrorf(err, "", "resources-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetBillingOptionsOptions : The GetBillingOptions options.
type GetBillingOptionsOptions struct {
	// Enterprise ID of the distributor or reseller for which the report is requested.
	PartnerID *string `json:"partner_id" validate:"required"`

	// Account ID/Enterprise ID of the end customer for which the report is requested. This parameter cannot be used along
	// with `reseller_id` query parameter.
	CustomerID *string `json:"customer_id,omitempty"`

	// Enterprise ID of the reseller for which the report is requested. This parameter cannot be used along with
	// `customer_id` query parameter.
	ResellerID *string `json:"reseller_id,omitempty"`

	// The billing month for which the report is requested. Format is yyyy-mm. Defaults to current month.
	Date *string `json:"date,omitempty"`

	// Number of billing option reports returned. The default value is 200. Maximum value is 200.
	Limit *int64 `json:"limit,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetBillingOptionsOptions : Instantiate GetBillingOptionsOptions
func (*PartnerManagementV1) NewGetBillingOptionsOptions(partnerID string, billingMonth string) *GetBillingOptionsOptions {
	return &GetBillingOptionsOptions{
		PartnerID: core.StringPtr(partnerID),
		Date:      core.StringPtr(billingMonth),
	}
}

// SetPartnerID : Allow user to set PartnerID
func (_options *GetBillingOptionsOptions) SetPartnerID(partnerID string) *GetBillingOptionsOptions {
	_options.PartnerID = core.StringPtr(partnerID)
	return _options
}

// SetCustomerID : Allow user to set CustomerID
func (_options *GetBillingOptionsOptions) SetCustomerID(customerID string) *GetBillingOptionsOptions {
	_options.CustomerID = core.StringPtr(customerID)
	return _options
}

// SetResellerID : Allow user to set ResellerID
func (_options *GetBillingOptionsOptions) SetResellerID(resellerID string) *GetBillingOptionsOptions {
	_options.ResellerID = core.StringPtr(resellerID)
	return _options
}

// SetDate : Allow user to set Date
func (_options *GetBillingOptionsOptions) SetDate(date string) *GetBillingOptionsOptions {
	_options.Date = core.StringPtr(date)
	return _options
}

// SetLimit : Allow user to set Limit
func (_options *GetBillingOptionsOptions) SetLimit(limit int64) *GetBillingOptionsOptions {
	_options.Limit = core.Int64Ptr(limit)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetBillingOptionsOptions) SetHeaders(param map[string]string) *GetBillingOptionsOptions {
	options.Headers = param
	return options
}

// GetCreditPoolsReportOptions : The GetCreditPoolsReport options.
type GetCreditPoolsReportOptions struct {
	// Enterprise ID of the distributor or reseller for which the report is requested.
	PartnerID *string `json:"partner_id" validate:"required"`

	// Account ID/Enterprise ID of the end customer for which the report is requested. This parameter cannot be used along
	// with `reseller_id` query parameter.
	CustomerID *string `json:"customer_id,omitempty"`

	// Enterprise ID of the reseller for which the report is requested. This parameter cannot be used along with
	// `customer_id` query parameter.
	ResellerID *string `json:"reseller_id,omitempty"`

	// The billing month for which the report is requested. Format is yyyy-mm. Defaults to current month.
	Date *string `json:"date,omitempty"`

	// Number of billing units fetched to get the credit pools report. The default value is 30. Maximum value is 30.
	Limit *int64 `json:"limit,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetCreditPoolsReportOptions : Instantiate GetCreditPoolsReportOptions
func (*PartnerManagementV1) NewGetCreditPoolsReportOptions(partnerID string, billingMonth string) *GetCreditPoolsReportOptions {
	return &GetCreditPoolsReportOptions{
		PartnerID: core.StringPtr(partnerID),
		Date:      core.StringPtr(billingMonth),
	}
}

// SetPartnerID : Allow user to set PartnerID
func (_options *GetCreditPoolsReportOptions) SetPartnerID(partnerID string) *GetCreditPoolsReportOptions {
	_options.PartnerID = core.StringPtr(partnerID)
	return _options
}

// SetCustomerID : Allow user to set CustomerID
func (_options *GetCreditPoolsReportOptions) SetCustomerID(customerID string) *GetCreditPoolsReportOptions {
	_options.CustomerID = core.StringPtr(customerID)
	return _options
}

// SetResellerID : Allow user to set ResellerID
func (_options *GetCreditPoolsReportOptions) SetResellerID(resellerID string) *GetCreditPoolsReportOptions {
	_options.ResellerID = core.StringPtr(resellerID)
	return _options
}

// SetDate : Allow user to set Date
func (_options *GetCreditPoolsReportOptions) SetDate(date string) *GetCreditPoolsReportOptions {
	_options.Date = core.StringPtr(date)
	return _options
}

// SetLimit : Allow user to set Limit
func (_options *GetCreditPoolsReportOptions) SetLimit(limit int64) *GetCreditPoolsReportOptions {
	_options.Limit = core.Int64Ptr(limit)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetCreditPoolsReportOptions) SetHeaders(param map[string]string) *GetCreditPoolsReportOptions {
	options.Headers = param
	return options
}

// GetResourceUsageReportOptions : The GetResourceUsageReport options.
type GetResourceUsageReportOptions struct {
	// Enterprise ID of the distributor or reseller for which the report is requested.
	PartnerID *string `json:"partner_id" validate:"required"`

	// Enterprise ID of the reseller for which the report is requested. This parameter cannot be used along with
	// `customer_id` query parameter.
	ResellerID *string `json:"reseller_id,omitempty"`

	// Account ID/Enterprise ID of the end customer for which the report is requested. This parameter cannot be used along
	// with `reseller_id` query parameter.
	CustomerID *string `json:"customer_id,omitempty"`

	// Get report rolled-up to the direct children of the requested entity. Defaults to false. This parameter cannot be
	// used along with `customer_id` query parameter.
	Children *bool `json:"children,omitempty"`

	// The billing month for which the usage report is requested. Format is `yyyy-mm`. Defaults to current month.
	Month *string `json:"month,omitempty"`

	// Enables partner to view the cost of provisioned services as applicable at the given level. Defaults to the type of
	// the calling partner. The valid values are `DISTRIBUTOR`, `RESELLER` and `END_CUSTOMER`.
	Viewpoint *string `json:"viewpoint,omitempty"`

	// Get usage report rolled-up to the end customers of the requesting partner. Defaults to false. This parameter cannot
	// be used along with `reseller_id` query parameter or `customer_id` query parameter.
	Recurse *bool `json:"recurse,omitempty"`

	// Number of usage records to be returned. The default value is 30. Maximum value is 100.
	Limit *int64 `json:"limit,omitempty"`

	// An opaque value representing the offset of the first item to be returned by a search query. If not specified, then
	// the first page of results is returned. To retrieve the next page of search results, use the 'offset' query parameter
	// value within the 'next.href' URL found within a prior search query response.
	Offset *string `json:"offset,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// Constants associated with the GetResourceUsageReportOptions.Viewpoint property.
// Enables partner to view the cost of provisioned services as applicable at the given level. Defaults to the type of
// the calling partner. The valid values are `DISTRIBUTOR`, `RESELLER` and `END_CUSTOMER`.
const (
	GetResourceUsageReportOptionsViewpointDistributorConst = "DISTRIBUTOR"
	GetResourceUsageReportOptionsViewpointEndCustomerConst = "END_CUSTOMER"
	GetResourceUsageReportOptionsViewpointResellerConst    = "RESELLER"
)

// NewGetResourceUsageReportOptions : Instantiate GetResourceUsageReportOptions
func (*PartnerManagementV1) NewGetResourceUsageReportOptions(partnerID string) *GetResourceUsageReportOptions {
	return &GetResourceUsageReportOptions{
		PartnerID: core.StringPtr(partnerID),
	}
}

// SetPartnerID : Allow user to set PartnerID
func (_options *GetResourceUsageReportOptions) SetPartnerID(partnerID string) *GetResourceUsageReportOptions {
	_options.PartnerID = core.StringPtr(partnerID)
	return _options
}

// SetResellerID : Allow user to set ResellerID
func (_options *GetResourceUsageReportOptions) SetResellerID(resellerID string) *GetResourceUsageReportOptions {
	_options.ResellerID = core.StringPtr(resellerID)
	return _options
}

// SetCustomerID : Allow user to set CustomerID
func (_options *GetResourceUsageReportOptions) SetCustomerID(customerID string) *GetResourceUsageReportOptions {
	_options.CustomerID = core.StringPtr(customerID)
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

// SetViewpoint : Allow user to set Viewpoint
func (_options *GetResourceUsageReportOptions) SetViewpoint(viewpoint string) *GetResourceUsageReportOptions {
	_options.Viewpoint = core.StringPtr(viewpoint)
	return _options
}

// SetRecurse : Allow user to set Recurse
func (_options *GetResourceUsageReportOptions) SetRecurse(recurse bool) *GetResourceUsageReportOptions {
	_options.Recurse = core.BoolPtr(recurse)
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
		err = core.SDKErrorf(err, "", "metric-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "unit", &obj.Unit)
	if err != nil {
		err = core.SDKErrorf(err, "", "unit-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "quantity", &obj.Quantity)
	if err != nil {
		err = core.SDKErrorf(err, "", "quantity-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "rateable_quantity", &obj.RateableQuantity)
	if err != nil {
		err = core.SDKErrorf(err, "", "rateable_quantity-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "cost", &obj.Cost)
	if err != nil {
		err = core.SDKErrorf(err, "", "cost-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "rated_cost", &obj.RatedCost)
	if err != nil {
		err = core.SDKErrorf(err, "", "rated_cost-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "price", &obj.Price)
	if err != nil {
		err = core.SDKErrorf(err, "", "price-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Overage : Overage that was generated on the credit pool.
type Overage struct {
	// The number of credits used as overage.
	Cost *float64 `json:"cost,omitempty"`

	// A list of resources that generated overage.
	Resources []map[string]interface{} `json:"resources,omitempty"`
}

// UnmarshalOverage unmarshals an instance of Overage from the specified map of raw messages.
func UnmarshalOverage(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Overage)
	err = core.UnmarshalPrimitive(m, "cost", &obj.Cost)
	if err != nil {
		err = core.SDKErrorf(err, "", "cost-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "resources", &obj.Resources)
	if err != nil {
		err = core.SDKErrorf(err, "", "resources-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PartnerUsageReportSummaryFirst : The link to the first page of the search query.
type PartnerUsageReportSummaryFirst struct {
	// A link to a page of query results.
	Href *string `json:"href,omitempty"`
}

// UnmarshalPartnerUsageReportSummaryFirst unmarshals an instance of PartnerUsageReportSummaryFirst from the specified map of raw messages.
func UnmarshalPartnerUsageReportSummaryFirst(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PartnerUsageReportSummaryFirst)
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		err = core.SDKErrorf(err, "", "href-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PartnerUsageReportSummaryNext : The link to the next page of the search query.
type PartnerUsageReportSummaryNext struct {
	// A link to a page of query results.
	Href *string `json:"href,omitempty"`

	// The value of the `_start` query parameter to fetch the next page.
	Offset *string `json:"offset,omitempty"`
}

// UnmarshalPartnerUsageReportSummaryNext unmarshals an instance of PartnerUsageReportSummaryNext from the specified map of raw messages.
func UnmarshalPartnerUsageReportSummaryNext(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PartnerUsageReportSummaryNext)
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		err = core.SDKErrorf(err, "", "href-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "offset", &obj.Offset)
	if err != nil {
		err = core.SDKErrorf(err, "", "offset-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PartnerUsageReport : Aggregated usage report of a partner.
type PartnerUsageReport struct {
	// The ID of the entity.
	EntityID *string `json:"entity_id,omitempty"`

	// The entity type.
	EntityType *string `json:"entity_type,omitempty"`

	// The Cloud Resource Name (CRN) of the entity towards which the resource usages were rolled up.
	EntityCRN *string `json:"entity_crn,omitempty"`

	// A user-defined name for the entity, such as the enterprise name or account name.
	EntityName *string `json:"entity_name,omitempty"`

	// Role of the `entity_id` for which the usage report is fetched.
	EntityPartnerType *string `json:"entity_partner_type,omitempty"`

	// Enables partner to view the cost of provisioned services as applicable at the given level.
	Viewpoint *string `json:"viewpoint,omitempty"`

	// The billing month for which the usage report is requested. Format is yyyy-mm.
	Month *string `json:"month,omitempty"`

	// The currency code of the billing unit.
	CurrencyCode *string `json:"currency_code,omitempty"`

	// The country code of the billing unit.
	CountryCode *string `json:"country_code,omitempty"`

	// Billable charges that are aggregated from all entities in the report.
	BillableCost *float64 `json:"billable_cost,omitempty"`

	// Aggregated billable charges before discounts.
	BillableRatedCost *float64 `json:"billable_rated_cost,omitempty"`

	// Non-billable charges that are aggregated from all entities in the report.
	NonBillableCost *float64 `json:"non_billable_cost,omitempty"`

	// Aggregated non-billable charges before discounts.
	NonBillableRatedCost *float64 `json:"non_billable_rated_cost,omitempty"`

	Resources []ResourceUsage `json:"resources,omitempty"`
}

// UnmarshalPartnerUsageReport unmarshals an instance of PartnerUsageReport from the specified map of raw messages.
func UnmarshalPartnerUsageReport(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PartnerUsageReport)
	err = core.UnmarshalPrimitive(m, "entity_id", &obj.EntityID)
	if err != nil {
		err = core.SDKErrorf(err, "", "entity_id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "entity_type", &obj.EntityType)
	if err != nil {
		err = core.SDKErrorf(err, "", "entity_type-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "entity_crn", &obj.EntityCRN)
	if err != nil {
		err = core.SDKErrorf(err, "", "entity_crn-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "entity_name", &obj.EntityName)
	if err != nil {
		err = core.SDKErrorf(err, "", "entity_name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "entity_partner_type", &obj.EntityPartnerType)
	if err != nil {
		err = core.SDKErrorf(err, "", "entity_partner_type-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "viewpoint", &obj.Viewpoint)
	if err != nil {
		err = core.SDKErrorf(err, "", "viewpoint-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "month", &obj.Month)
	if err != nil {
		err = core.SDKErrorf(err, "", "month-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "currency_code", &obj.CurrencyCode)
	if err != nil {
		err = core.SDKErrorf(err, "", "currency_code-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "country_code", &obj.CountryCode)
	if err != nil {
		err = core.SDKErrorf(err, "", "country_code-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "billable_cost", &obj.BillableCost)
	if err != nil {
		err = core.SDKErrorf(err, "", "billable_cost-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "billable_rated_cost", &obj.BillableRatedCost)
	if err != nil {
		err = core.SDKErrorf(err, "", "billable_rated_cost-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "non_billable_cost", &obj.NonBillableCost)
	if err != nil {
		err = core.SDKErrorf(err, "", "non_billable_cost-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "non_billable_rated_cost", &obj.NonBillableRatedCost)
	if err != nil {
		err = core.SDKErrorf(err, "", "non_billable_rated_cost-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "resources", &obj.Resources, UnmarshalResourceUsage)
	if err != nil {
		err = core.SDKErrorf(err, "", "resources-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PartnerUsageReportSummary : The aggregated partner usage report.
type PartnerUsageReportSummary struct {
	// The maximum number of usage records in the response.
	Limit *int64 `json:"limit,omitempty"`

	// The link to the first page of the search query.
	First *PartnerUsageReportSummaryFirst `json:"first,omitempty"`

	// The link to the next page of the search query.
	Next *PartnerUsageReportSummaryNext `json:"next,omitempty"`

	// Aggregated usage report of all requested partners.
	Reports []PartnerUsageReport `json:"reports,omitempty"`
}

// UnmarshalPartnerUsageReportSummary unmarshals an instance of PartnerUsageReportSummary from the specified map of raw messages.
func UnmarshalPartnerUsageReportSummary(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PartnerUsageReportSummary)
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		err = core.SDKErrorf(err, "", "limit-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "first", &obj.First, UnmarshalPartnerUsageReportSummaryFirst)
	if err != nil {
		err = core.SDKErrorf(err, "", "first-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "next", &obj.Next, UnmarshalPartnerUsageReportSummaryNext)
	if err != nil {
		err = core.SDKErrorf(err, "", "next-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "reports", &obj.Reports, UnmarshalPartnerUsageReport)
	if err != nil {
		err = core.SDKErrorf(err, "", "reports-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Retrieve the value to be passed to a request to access the next page of results
func (resp *PartnerUsageReportSummary) GetNextOffset() (*string, error) {
	if core.IsNil(resp.Next) {
		return nil, nil
	}
	return resp.Next.Offset, nil
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
		err = core.SDKErrorf(err, "", "plan_id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "pricing_region", &obj.PricingRegion)
	if err != nil {
		err = core.SDKErrorf(err, "", "pricing_region-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "pricing_plan_id", &obj.PricingPlanID)
	if err != nil {
		err = core.SDKErrorf(err, "", "pricing_plan_id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "billable", &obj.Billable)
	if err != nil {
		err = core.SDKErrorf(err, "", "billable-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "cost", &obj.Cost)
	if err != nil {
		err = core.SDKErrorf(err, "", "cost-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "rated_cost", &obj.RatedCost)
	if err != nil {
		err = core.SDKErrorf(err, "", "rated_cost-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "usage", &obj.Usage, UnmarshalMetricUsage)
	if err != nil {
		err = core.SDKErrorf(err, "", "usage-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ResourceUsage : A container for all the plans in the resource.
type ResourceUsage struct {
	// The ID of the resource.
	ResourceID *string `json:"resource_id" validate:"required"`

	// The name of the resource.
	ResourceName *string `json:"resource_name,omitempty"`

	// The billable charges for the partner.
	BillableCost *float64 `json:"billable_cost" validate:"required"`

	// The pre-discounted billable charges for the partner.
	BillableRatedCost *float64 `json:"billable_rated_cost" validate:"required"`

	// The non-billable charges for the partner.
	NonBillableCost *float64 `json:"non_billable_cost" validate:"required"`

	// The pre-discounted, non-billable charges for the partner.
	NonBillableRatedCost *float64 `json:"non_billable_rated_cost" validate:"required"`

	// All of the plans in the resource.
	Plans []PlanUsage `json:"plans" validate:"required"`
}

// UnmarshalResourceUsage unmarshals an instance of ResourceUsage from the specified map of raw messages.
func UnmarshalResourceUsage(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ResourceUsage)
	err = core.UnmarshalPrimitive(m, "resource_id", &obj.ResourceID)
	if err != nil {
		err = core.SDKErrorf(err, "", "resource_id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "resource_name", &obj.ResourceName)
	if err != nil {
		err = core.SDKErrorf(err, "", "resource_name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "billable_cost", &obj.BillableCost)
	if err != nil {
		err = core.SDKErrorf(err, "", "billable_cost-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "billable_rated_cost", &obj.BillableRatedCost)
	if err != nil {
		err = core.SDKErrorf(err, "", "billable_rated_cost-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "non_billable_cost", &obj.NonBillableCost)
	if err != nil {
		err = core.SDKErrorf(err, "", "non_billable_cost-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "non_billable_rated_cost", &obj.NonBillableRatedCost)
	if err != nil {
		err = core.SDKErrorf(err, "", "non_billable_rated_cost-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "plans", &obj.Plans, UnmarshalPlanUsage)
	if err != nil {
		err = core.SDKErrorf(err, "", "plans-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TermCredits : The subscription term that is active in the requested month.
type TermCredits struct {
	// The ID of the billing option from which the subscription term is derived.
	BillingOptionID *string `json:"billing_option_id,omitempty"`

	// Billing option model.
	BillingOptionModel *string `json:"billing_option_model,omitempty"`

	// The category of the billing option. The valid values are `PLATFORM`, `SERVICE`, and `SUPPORT`.
	Category *string `json:"category,omitempty"`

	// The start date of the term in ISO format.
	StartDate *strfmt.DateTime `json:"start_date,omitempty"`

	// The end date of the term in ISO format.
	EndDate *strfmt.DateTime `json:"end_date,omitempty"`

	// The total credit available in this term.
	TotalCredits *float64 `json:"total_credits,omitempty"`

	// The balance of available credit at the start of the current month.
	StartingBalance *float64 `json:"starting_balance,omitempty"`

	// The amount of credit used during the current month.
	UsedCredits *float64 `json:"used_credits,omitempty"`

	// The balance of remaining credit in the subscription term.
	CurrentBalance *float64 `json:"current_balance,omitempty"`

	// A list of resources that used credit during the month.
	Resources []map[string]interface{} `json:"resources,omitempty"`
}

// Constants associated with the TermCredits.Category property.
// The category of the billing option. The valid values are `PLATFORM`, `SERVICE`, and `SUPPORT`.
const (
	TermCreditsCategoryPlatformConst = "PLATFORM"
	TermCreditsCategoryServiceConst  = "SERVICE"
	TermCreditsCategorySupportConst  = "SUPPORT"
)

// UnmarshalTermCredits unmarshals an instance of TermCredits from the specified map of raw messages.
func UnmarshalTermCredits(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TermCredits)
	err = core.UnmarshalPrimitive(m, "billing_option_id", &obj.BillingOptionID)
	if err != nil {
		err = core.SDKErrorf(err, "", "billing_option_id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "billing_option_model", &obj.BillingOptionModel)
	if err != nil {
		err = core.SDKErrorf(err, "", "billing_option_model-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "category", &obj.Category)
	if err != nil {
		err = core.SDKErrorf(err, "", "category-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "start_date", &obj.StartDate)
	if err != nil {
		err = core.SDKErrorf(err, "", "start_date-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "end_date", &obj.EndDate)
	if err != nil {
		err = core.SDKErrorf(err, "", "end_date-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "total_credits", &obj.TotalCredits)
	if err != nil {
		err = core.SDKErrorf(err, "", "total_credits-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "starting_balance", &obj.StartingBalance)
	if err != nil {
		err = core.SDKErrorf(err, "", "starting_balance-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "used_credits", &obj.UsedCredits)
	if err != nil {
		err = core.SDKErrorf(err, "", "used_credits-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "current_balance", &obj.CurrentBalance)
	if err != nil {
		err = core.SDKErrorf(err, "", "current_balance-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "resources", &obj.Resources)
	if err != nil {
		err = core.SDKErrorf(err, "", "resources-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetResourceUsageReportPager can be used to simplify the use of the "GetResourceUsageReport" method.
type GetResourceUsageReportPager struct {
	hasNext     bool
	options     *GetResourceUsageReportOptions
	client      *PartnerManagementV1
	pageContext struct {
		next *string
	}
}

// NewGetResourceUsageReportPager returns a new GetResourceUsageReportPager instance.
func (partnerManagement *PartnerManagementV1) NewGetResourceUsageReportPager(options *GetResourceUsageReportOptions) (pager *GetResourceUsageReportPager, err error) {
	if options.Offset != nil && *options.Offset != "" {
		err = core.SDKErrorf(nil, "the 'options.Offset' field should not be set", "no-query-setting", common.GetComponentInfo())
		return
	}

	var optionsCopy GetResourceUsageReportOptions = *options
	pager = &GetResourceUsageReportPager{
		hasNext: true,
		options: &optionsCopy,
		client:  partnerManagement,
	}
	return
}

// HasNext returns true if there are potentially more results to be retrieved.
func (pager *GetResourceUsageReportPager) HasNext() bool {
	return pager.hasNext
}

// GetNextWithContext returns the next page of results using the specified Context.
func (pager *GetResourceUsageReportPager) GetNextWithContext(ctx context.Context) (page []PartnerUsageReport, err error) {
	if !pager.HasNext() {
		return nil, fmt.Errorf("no more results available")
	}

	pager.options.Offset = pager.pageContext.next

	result, _, err := pager.client.GetResourceUsageReportWithContext(ctx, pager.options)
	if err != nil {
		err = core.RepurposeSDKProblem(err, "error-getting-next-page")
		return
	}

	var next *string
	if result.Next != nil {
		next = result.Next.Offset
	}
	pager.pageContext.next = next
	pager.hasNext = (pager.pageContext.next != nil)
	page = result.Reports

	return
}

// GetAllWithContext returns all results by invoking GetNextWithContext() repeatedly
// until all pages of results have been retrieved.
func (pager *GetResourceUsageReportPager) GetAllWithContext(ctx context.Context) (allItems []PartnerUsageReport, err error) {
	for pager.HasNext() {
		var nextPage []PartnerUsageReport
		nextPage, err = pager.GetNextWithContext(ctx)
		if err != nil {
			err = core.RepurposeSDKProblem(err, "error-getting-next-page")
			return
		}
		allItems = append(allItems, nextPage...)
	}
	return
}

// GetNext invokes GetNextWithContext() using context.Background() as the Context parameter.
func (pager *GetResourceUsageReportPager) GetNext() (page []PartnerUsageReport, err error) {
	page, err = pager.GetNextWithContext(context.Background())
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetAll invokes GetAllWithContext() using context.Background() as the Context parameter.
func (pager *GetResourceUsageReportPager) GetAll() (allItems []PartnerUsageReport, err error) {
	allItems, err = pager.GetAllWithContext(context.Background())
	err = core.RepurposeSDKProblem(err, "")
	return
}
