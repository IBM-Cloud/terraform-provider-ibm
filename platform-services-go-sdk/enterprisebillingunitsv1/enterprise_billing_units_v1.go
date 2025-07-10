/**
 * (C) Copyright IBM Corp. 2023.
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
 * IBM OpenAPI SDK Code Generator Version: 3.64.1-cee95189-20230124-211647
 */

// Package enterprisebillingunitsv1 : Operations and models for the EnterpriseBillingUnitsV1 service
package enterprisebillingunitsv1

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

// EnterpriseBillingUnitsV1 : Billing units for IBM Cloud Enterprise
//
// API Version: 1.0.0
type EnterpriseBillingUnitsV1 struct {
	Service *core.BaseService
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://billing.cloud.ibm.com"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "enterprise_billing_units"

// EnterpriseBillingUnitsV1Options : Service options
type EnterpriseBillingUnitsV1Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
}

// NewEnterpriseBillingUnitsV1UsingExternalConfig : constructs an instance of EnterpriseBillingUnitsV1 with passed in options and external configuration.
func NewEnterpriseBillingUnitsV1UsingExternalConfig(options *EnterpriseBillingUnitsV1Options) (enterpriseBillingUnits *EnterpriseBillingUnitsV1, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	enterpriseBillingUnits, err = NewEnterpriseBillingUnitsV1(options)
	if err != nil {
		return
	}

	err = enterpriseBillingUnits.Service.ConfigureService(options.ServiceName)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = enterpriseBillingUnits.Service.SetServiceURL(options.URL)
	}
	return
}

// NewEnterpriseBillingUnitsV1 : constructs an instance of EnterpriseBillingUnitsV1 with passed in options.
func NewEnterpriseBillingUnitsV1(options *EnterpriseBillingUnitsV1Options) (service *EnterpriseBillingUnitsV1, err error) {
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

	service = &EnterpriseBillingUnitsV1{
		Service: baseService,
	}

	return
}

// GetServiceURLForRegion returns the service URL to be used for the specified region
func GetServiceURLForRegion(region string) (string, error) {
	return "", fmt.Errorf("service does not support regional URLs")
}

// Clone makes a copy of "enterpriseBillingUnits" suitable for processing requests.
func (enterpriseBillingUnits *EnterpriseBillingUnitsV1) Clone() *EnterpriseBillingUnitsV1 {
	if core.IsNil(enterpriseBillingUnits) {
		return nil
	}
	clone := *enterpriseBillingUnits
	clone.Service = enterpriseBillingUnits.Service.Clone()
	return &clone
}

// SetServiceURL sets the service URL
func (enterpriseBillingUnits *EnterpriseBillingUnitsV1) SetServiceURL(url string) error {
	return enterpriseBillingUnits.Service.SetServiceURL(url)
}

// GetServiceURL returns the service URL
func (enterpriseBillingUnits *EnterpriseBillingUnitsV1) GetServiceURL() string {
	return enterpriseBillingUnits.Service.GetServiceURL()
}

// SetDefaultHeaders sets HTTP headers to be sent in every request
func (enterpriseBillingUnits *EnterpriseBillingUnitsV1) SetDefaultHeaders(headers http.Header) {
	enterpriseBillingUnits.Service.SetDefaultHeaders(headers)
}

// SetEnableGzipCompression sets the service's EnableGzipCompression field
func (enterpriseBillingUnits *EnterpriseBillingUnitsV1) SetEnableGzipCompression(enableGzip bool) {
	enterpriseBillingUnits.Service.SetEnableGzipCompression(enableGzip)
}

// GetEnableGzipCompression returns the service's EnableGzipCompression field
func (enterpriseBillingUnits *EnterpriseBillingUnitsV1) GetEnableGzipCompression() bool {
	return enterpriseBillingUnits.Service.GetEnableGzipCompression()
}

// EnableRetries enables automatic retries for requests invoked for this service instance.
// If either parameter is specified as 0, then a default value is used instead.
func (enterpriseBillingUnits *EnterpriseBillingUnitsV1) EnableRetries(maxRetries int, maxRetryInterval time.Duration) {
	enterpriseBillingUnits.Service.EnableRetries(maxRetries, maxRetryInterval)
}

// DisableRetries disables automatic retries for requests invoked for this service instance.
func (enterpriseBillingUnits *EnterpriseBillingUnitsV1) DisableRetries() {
	enterpriseBillingUnits.Service.DisableRetries()
}

// GetBillingUnit : Get billing unit by ID
// Return the billing unit information if it exists.
func (enterpriseBillingUnits *EnterpriseBillingUnitsV1) GetBillingUnit(getBillingUnitOptions *GetBillingUnitOptions) (result *BillingUnit, response *core.DetailedResponse, err error) {
	return enterpriseBillingUnits.GetBillingUnitWithContext(context.Background(), getBillingUnitOptions)
}

// GetBillingUnitWithContext is an alternate form of the GetBillingUnit method which supports a Context parameter
func (enterpriseBillingUnits *EnterpriseBillingUnitsV1) GetBillingUnitWithContext(ctx context.Context, getBillingUnitOptions *GetBillingUnitOptions) (result *BillingUnit, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getBillingUnitOptions, "getBillingUnitOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getBillingUnitOptions, "getBillingUnitOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"billing_unit_id": *getBillingUnitOptions.BillingUnitID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = enterpriseBillingUnits.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(enterpriseBillingUnits.Service.Options.URL, `/v1/billing-units/{billing_unit_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getBillingUnitOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("enterprise_billing_units", "V1", "GetBillingUnit")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = enterpriseBillingUnits.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalBillingUnit)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListBillingUnits : List billing units
// Return matching billing unit information if any exists. Omits internal properties and enterprise account ID from the
// billing unit.
func (enterpriseBillingUnits *EnterpriseBillingUnitsV1) ListBillingUnits(listBillingUnitsOptions *ListBillingUnitsOptions) (result *BillingUnitsList, response *core.DetailedResponse, err error) {
	return enterpriseBillingUnits.ListBillingUnitsWithContext(context.Background(), listBillingUnitsOptions)
}

// ListBillingUnitsWithContext is an alternate form of the ListBillingUnits method which supports a Context parameter
func (enterpriseBillingUnits *EnterpriseBillingUnitsV1) ListBillingUnitsWithContext(ctx context.Context, listBillingUnitsOptions *ListBillingUnitsOptions) (result *BillingUnitsList, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listBillingUnitsOptions, "listBillingUnitsOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = enterpriseBillingUnits.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(enterpriseBillingUnits.Service.Options.URL, `/v1/billing-units`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range listBillingUnitsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("enterprise_billing_units", "V1", "ListBillingUnits")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if listBillingUnitsOptions.AccountID != nil {
		builder.AddQuery("account_id", fmt.Sprint(*listBillingUnitsOptions.AccountID))
	}
	if listBillingUnitsOptions.EnterpriseID != nil {
		builder.AddQuery("enterprise_id", fmt.Sprint(*listBillingUnitsOptions.EnterpriseID))
	}
	if listBillingUnitsOptions.AccountGroupID != nil {
		builder.AddQuery("account_group_id", fmt.Sprint(*listBillingUnitsOptions.AccountGroupID))
	}
	if listBillingUnitsOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*listBillingUnitsOptions.Limit))
	}
	if listBillingUnitsOptions.Start != nil {
		builder.AddQuery("start", fmt.Sprint(*listBillingUnitsOptions.Start))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = enterpriseBillingUnits.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalBillingUnitsList)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListBillingOptions : List billing options
// Return matching billing options if any exist. Show subscriptions and promotional offers that are available to a
// billing unit.
func (enterpriseBillingUnits *EnterpriseBillingUnitsV1) ListBillingOptions(listBillingOptionsOptions *ListBillingOptionsOptions) (result *BillingOptionsList, response *core.DetailedResponse, err error) {
	return enterpriseBillingUnits.ListBillingOptionsWithContext(context.Background(), listBillingOptionsOptions)
}

// ListBillingOptionsWithContext is an alternate form of the ListBillingOptions method which supports a Context parameter
func (enterpriseBillingUnits *EnterpriseBillingUnitsV1) ListBillingOptionsWithContext(ctx context.Context, listBillingOptionsOptions *ListBillingOptionsOptions) (result *BillingOptionsList, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listBillingOptionsOptions, "listBillingOptionsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listBillingOptionsOptions, "listBillingOptionsOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = enterpriseBillingUnits.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(enterpriseBillingUnits.Service.Options.URL, `/v1/billing-options`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range listBillingOptionsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("enterprise_billing_units", "V1", "ListBillingOptions")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("billing_unit_id", fmt.Sprint(*listBillingOptionsOptions.BillingUnitID))
	if listBillingOptionsOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*listBillingOptionsOptions.Limit))
	}
	if listBillingOptionsOptions.Start != nil {
		builder.AddQuery("start", fmt.Sprint(*listBillingOptionsOptions.Start))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = enterpriseBillingUnits.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalBillingOptionsList)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetCreditPools : Get credit pools
// Get credit pools for a billing unit. Credit pools can be either platform or support credit pools. The platform credit
// pool contains credit from platform subscriptions and promotional offers. The support credit pool contains credit from
// support subscriptions.
func (enterpriseBillingUnits *EnterpriseBillingUnitsV1) GetCreditPools(getCreditPoolsOptions *GetCreditPoolsOptions) (result *CreditPoolsList, response *core.DetailedResponse, err error) {
	return enterpriseBillingUnits.GetCreditPoolsWithContext(context.Background(), getCreditPoolsOptions)
}

// GetCreditPoolsWithContext is an alternate form of the GetCreditPools method which supports a Context parameter
func (enterpriseBillingUnits *EnterpriseBillingUnitsV1) GetCreditPoolsWithContext(ctx context.Context, getCreditPoolsOptions *GetCreditPoolsOptions) (result *CreditPoolsList, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getCreditPoolsOptions, "getCreditPoolsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getCreditPoolsOptions, "getCreditPoolsOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = enterpriseBillingUnits.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(enterpriseBillingUnits.Service.Options.URL, `/v1/credit-pools`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range getCreditPoolsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("enterprise_billing_units", "V1", "GetCreditPools")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("billing_unit_id", fmt.Sprint(*getCreditPoolsOptions.BillingUnitID))
	if getCreditPoolsOptions.Date != nil {
		builder.AddQuery("date", fmt.Sprint(*getCreditPoolsOptions.Date))
	}
	if getCreditPoolsOptions.Type != nil {
		builder.AddQuery("type", fmt.Sprint(*getCreditPoolsOptions.Type))
	}
	if getCreditPoolsOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*getCreditPoolsOptions.Limit))
	}
	if getCreditPoolsOptions.Start != nil {
		builder.AddQuery("start", fmt.Sprint(*getCreditPoolsOptions.Start))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = enterpriseBillingUnits.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCreditPoolsList)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// BillingOption : Information about a billing option.
type BillingOption struct {
	// The ID of the billing option.
	ID *string `json:"id,omitempty"`

	// The ID of the billing unit that's associated with the billing option.
	BillingUnitID *string `json:"billing_unit_id,omitempty"`

	// The start date of billing option.
	StartDate *strfmt.DateTime `json:"start_date,omitempty"`

	// The end date of billing option.
	EndDate *strfmt.DateTime `json:"end_date,omitempty"`

	// The state of the billing option. The valid values include `ACTIVE, `SUSPENDED`, and `CANCELED`.
	State *string `json:"state,omitempty"`

	// The type of billing option. The valid values are `SUBSCRIPTION` and `OFFER`.
	Type *string `json:"type,omitempty"`

	// The category of the billing option. The valid values are `PLATFORM`, `SERVICE`, and `SUPPORT`.
	Category *string `json:"category,omitempty"`

	// The payment method for support.
	PaymentInstrument map[string]interface{} `json:"payment_instrument,omitempty"`

	// The duration of the billing options in months.
	DurationInMonths *int64 `json:"duration_in_months,omitempty"`

	// The line item ID for support.
	LineItemID *int64 `json:"line_item_id,omitempty"`

	// The support billing system.
	BillingSystem map[string]interface{} `json:"billing_system,omitempty"`

	// The renewal code for support. This code denotes whether the subscription automatically renews, is assessed monthly,
	// and so on.
	RenewalModeCode *string `json:"renewal_mode_code,omitempty"`

	// The date when the billing option was updated.
	UpdatedAt *strfmt.DateTime `json:"updated_at,omitempty"`
}

// Constants associated with the BillingOption.State property.
// The state of the billing option. The valid values include `ACTIVE, `SUSPENDED`, and `CANCELED`.
const (
	BillingOptionStateActiveConst    = "ACTIVE"
	BillingOptionStateCanceledConst  = "CANCELED"
	BillingOptionStateSuspendedConst = "SUSPENDED"
)

// Constants associated with the BillingOption.Type property.
// The type of billing option. The valid values are `SUBSCRIPTION` and `OFFER`.
const (
	BillingOptionTypeOfferConst        = "OFFER"
	BillingOptionTypeSubscriptionConst = "SUBSCRIPTION"
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
		return
	}
	err = core.UnmarshalPrimitive(m, "billing_unit_id", &obj.BillingUnitID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "start_date", &obj.StartDate)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "end_date", &obj.EndDate)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "state", &obj.State)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "category", &obj.Category)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "payment_instrument", &obj.PaymentInstrument)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "duration_in_months", &obj.DurationInMonths)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "line_item_id", &obj.LineItemID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "billing_system", &obj.BillingSystem)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "renewal_mode_code", &obj.RenewalModeCode)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "updated_at", &obj.UpdatedAt)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// BillingOptionsList : A search result containing zero or more billing options.
type BillingOptionsList struct {
	// A count of the billing units that were found by the query.
	RowsCount *int64 `json:"rows_count,omitempty"`

	// Bookmark URL to query for next batch of billing units. This returns `null` if no additional pages are required.
	NextURL *string `json:"next_url,omitempty"`

	// A list of billing units found.
	Resources []BillingOption `json:"resources,omitempty"`
}

// UnmarshalBillingOptionsList unmarshals an instance of BillingOptionsList from the specified map of raw messages.
func UnmarshalBillingOptionsList(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(BillingOptionsList)
	err = core.UnmarshalPrimitive(m, "rows_count", &obj.RowsCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "next_url", &obj.NextURL)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "resources", &obj.Resources, UnmarshalBillingOption)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Retrieve the value to be passed to a request to access the next page of results
func (resp *BillingOptionsList) GetNextStart() (*string, error) {
	if core.IsNil(resp.NextURL) {
		return nil, nil
	}
	start, err := core.GetQueryParam(resp.NextURL, "start")
	if err != nil || start == nil {
		return nil, err
	}
	return start, nil
}

// BillingUnit : Information about a billing unit.
type BillingUnit struct {
	// The ID of the billing unit, which is a globally unique identifier (GUID).
	ID *string `json:"id,omitempty"`

	// The Cloud Resource Name (CRN) of the billing unit, scoped to the enterprise account ID.
	CRN *string `json:"crn,omitempty"`

	// The name of the billing unit.
	Name *string `json:"name,omitempty"`

	// The ID of the enterprise to which the billing unit is associated.
	EnterpriseID *string `json:"enterprise_id,omitempty"`

	// The currency code for the billing unit.
	CurrencyCode *string `json:"currency_code,omitempty"`

	// The country code for the billing unit.
	CountryCode *string `json:"country_code,omitempty"`

	// A flag that indicates whether this billing unit is the primary billing mechanism for the enterprise.
	Master *bool `json:"master,omitempty"`

	// The creation date of the billing unit.
	CreatedAt *strfmt.DateTime `json:"created_at,omitempty"`
}

// UnmarshalBillingUnit unmarshals an instance of BillingUnit from the specified map of raw messages.
func UnmarshalBillingUnit(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(BillingUnit)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "crn", &obj.CRN)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "enterprise_id", &obj.EnterpriseID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "currency_code", &obj.CurrencyCode)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "country_code", &obj.CountryCode)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "master", &obj.Master)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_at", &obj.CreatedAt)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// BillingUnitsList : A search result contining zero or more billing units.
type BillingUnitsList struct {
	// A count of the billing units that were found by the query.
	RowsCount *int64 `json:"rows_count,omitempty"`

	// Bookmark URL to query for next batch of billing units. This returns `null` if no additional pages are required.
	NextURL *string `json:"next_url,omitempty"`

	// A list of billing units found.
	Resources []BillingUnit `json:"resources,omitempty"`
}

// UnmarshalBillingUnitsList unmarshals an instance of BillingUnitsList from the specified map of raw messages.
func UnmarshalBillingUnitsList(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(BillingUnitsList)
	err = core.UnmarshalPrimitive(m, "rows_count", &obj.RowsCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "next_url", &obj.NextURL)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "resources", &obj.Resources, UnmarshalBillingUnit)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Retrieve the value to be passed to a request to access the next page of results
func (resp *BillingUnitsList) GetNextStart() (*string, error) {
	if core.IsNil(resp.NextURL) {
		return nil, nil
	}
	start, err := core.GetQueryParam(resp.NextURL, "start")
	if err != nil || start == nil {
		return nil, err
	}
	return start, nil
}

// CreditPool : The credit pool for a billing unit.
type CreditPool struct {
	// The type of credit, either `PLATFORM` or `SUPPORT`.
	Type *string `json:"type,omitempty"`

	// The currency code of the associated billing unit.
	CurrencyCode *string `json:"currency_code,omitempty"`

	// The ID of the billing unit that's associated with the credit pool. This value is a globally unique identifier
	// (GUID).
	BillingUnitID *string `json:"billing_unit_id,omitempty"`

	// A list of active subscription terms available within a credit pool.
	TermCredits []TermCredits `json:"term_credits,omitempty"`

	// Overage that was generated on the credit pool.
	Overage *CreditPoolOverage `json:"overage,omitempty"`
}

// Constants associated with the CreditPool.Type property.
// The type of credit, either `PLATFORM` or `SUPPORT`.
const (
	CreditPoolTypePlatformConst = "PLATFORM"
	CreditPoolTypeSupportConst  = "SUPPORT"
)

// UnmarshalCreditPool unmarshals an instance of CreditPool from the specified map of raw messages.
func UnmarshalCreditPool(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CreditPool)
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "currency_code", &obj.CurrencyCode)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "billing_unit_id", &obj.BillingUnitID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "term_credits", &obj.TermCredits, UnmarshalTermCredits)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "overage", &obj.Overage, UnmarshalCreditPoolOverage)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CreditPoolOverage : Overage that was generated on the credit pool.
type CreditPoolOverage struct {
	// The number of credits used as overage.
	Cost *float64 `json:"cost,omitempty"`

	// A list of resources that generated overage.
	Resources []map[string]interface{} `json:"resources,omitempty"`
}

// UnmarshalCreditPoolOverage unmarshals an instance of CreditPoolOverage from the specified map of raw messages.
func UnmarshalCreditPoolOverage(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CreditPoolOverage)
	err = core.UnmarshalPrimitive(m, "cost", &obj.Cost)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "resources", &obj.Resources)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CreditPoolsList : A search result containing zero or more credit pools.
type CreditPoolsList struct {
	// The number of credit pools that were found by the query.
	RowsCount *int64 `json:"rows_count,omitempty"`

	// A bookmark URL to the query for the next batch of billing units. Use a value of `null` if no additional pages are
	// required.
	NextURL *string `json:"next_url,omitempty"`

	// A list of credit pools found by the query.
	Resources []CreditPool `json:"resources,omitempty"`
}

// UnmarshalCreditPoolsList unmarshals an instance of CreditPoolsList from the specified map of raw messages.
func UnmarshalCreditPoolsList(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CreditPoolsList)
	err = core.UnmarshalPrimitive(m, "rows_count", &obj.RowsCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "next_url", &obj.NextURL)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "resources", &obj.Resources, UnmarshalCreditPool)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Retrieve the value to be passed to a request to access the next page of results
func (resp *CreditPoolsList) GetNextStart() (*string, error) {
	if core.IsNil(resp.NextURL) {
		return nil, nil
	}
	start, err := core.GetQueryParam(resp.NextURL, "start")
	if err != nil || start == nil {
		return nil, err
	}
	return start, nil
}

// GetBillingUnitOptions : The GetBillingUnit options.
type GetBillingUnitOptions struct {
	// The ID of the requested billing unit.
	BillingUnitID *string `json:"billing_unit_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetBillingUnitOptions : Instantiate GetBillingUnitOptions
func (*EnterpriseBillingUnitsV1) NewGetBillingUnitOptions(billingUnitID string) *GetBillingUnitOptions {
	return &GetBillingUnitOptions{
		BillingUnitID: core.StringPtr(billingUnitID),
	}
}

// SetBillingUnitID : Allow user to set BillingUnitID
func (_options *GetBillingUnitOptions) SetBillingUnitID(billingUnitID string) *GetBillingUnitOptions {
	_options.BillingUnitID = core.StringPtr(billingUnitID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetBillingUnitOptions) SetHeaders(param map[string]string) *GetBillingUnitOptions {
	options.Headers = param
	return options
}

// GetCreditPoolsOptions : The GetCreditPools options.
type GetCreditPoolsOptions struct {
	// The ID of the billing unit.
	BillingUnitID *string `json:"billing_unit_id" validate:"required"`

	// The date in the format of YYYY-MM.
	Date *string `json:"date,omitempty"`

	// Filters the credit pool by type, either `PLATFORM` or `SUPPORT`.
	Type *string `json:"type,omitempty"`

	// Return results up to this limit. Valid values are between 0 and 100.
	Limit *int64 `json:"limit,omitempty"`

	// The pagination offset. This represents the index of the first returned result.
	Start *string `json:"start,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetCreditPoolsOptions : Instantiate GetCreditPoolsOptions
func (*EnterpriseBillingUnitsV1) NewGetCreditPoolsOptions(billingUnitID string) *GetCreditPoolsOptions {
	return &GetCreditPoolsOptions{
		BillingUnitID: core.StringPtr(billingUnitID),
	}
}

// SetBillingUnitID : Allow user to set BillingUnitID
func (_options *GetCreditPoolsOptions) SetBillingUnitID(billingUnitID string) *GetCreditPoolsOptions {
	_options.BillingUnitID = core.StringPtr(billingUnitID)
	return _options
}

// SetDate : Allow user to set Date
func (_options *GetCreditPoolsOptions) SetDate(date string) *GetCreditPoolsOptions {
	_options.Date = core.StringPtr(date)
	return _options
}

// SetType : Allow user to set Type
func (_options *GetCreditPoolsOptions) SetType(typeVar string) *GetCreditPoolsOptions {
	_options.Type = core.StringPtr(typeVar)
	return _options
}

// SetLimit : Allow user to set Limit
func (_options *GetCreditPoolsOptions) SetLimit(limit int64) *GetCreditPoolsOptions {
	_options.Limit = core.Int64Ptr(limit)
	return _options
}

// SetStart : Allow user to set Start
func (_options *GetCreditPoolsOptions) SetStart(start string) *GetCreditPoolsOptions {
	_options.Start = core.StringPtr(start)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetCreditPoolsOptions) SetHeaders(param map[string]string) *GetCreditPoolsOptions {
	options.Headers = param
	return options
}

// ListBillingOptionsOptions : The ListBillingOptions options.
type ListBillingOptionsOptions struct {
	// The billing unit ID.
	BillingUnitID *string `json:"billing_unit_id" validate:"required"`

	// Return results up to this limit. Valid values are between 0 and 100.
	Limit *int64 `json:"limit,omitempty"`

	// The pagination offset. This represents the index of the first returned result.
	Start *string `json:"start,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListBillingOptionsOptions : Instantiate ListBillingOptionsOptions
func (*EnterpriseBillingUnitsV1) NewListBillingOptionsOptions(billingUnitID string) *ListBillingOptionsOptions {
	return &ListBillingOptionsOptions{
		BillingUnitID: core.StringPtr(billingUnitID),
	}
}

// SetBillingUnitID : Allow user to set BillingUnitID
func (_options *ListBillingOptionsOptions) SetBillingUnitID(billingUnitID string) *ListBillingOptionsOptions {
	_options.BillingUnitID = core.StringPtr(billingUnitID)
	return _options
}

// SetLimit : Allow user to set Limit
func (_options *ListBillingOptionsOptions) SetLimit(limit int64) *ListBillingOptionsOptions {
	_options.Limit = core.Int64Ptr(limit)
	return _options
}

// SetStart : Allow user to set Start
func (_options *ListBillingOptionsOptions) SetStart(start string) *ListBillingOptionsOptions {
	_options.Start = core.StringPtr(start)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListBillingOptionsOptions) SetHeaders(param map[string]string) *ListBillingOptionsOptions {
	options.Headers = param
	return options
}

// ListBillingUnitsOptions : The ListBillingUnits options.
type ListBillingUnitsOptions struct {
	// The enterprise account ID.
	AccountID *string `json:"account_id,omitempty"`

	// The enterprise ID.
	EnterpriseID *string `json:"enterprise_id,omitempty"`

	// The account group ID.
	AccountGroupID *string `json:"account_group_id,omitempty"`

	// Return results up to this limit. Valid values are between 0 and 100.
	Limit *int64 `json:"limit,omitempty"`

	// The pagination offset. This represents the index of the first returned result.
	Start *string `json:"start,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListBillingUnitsOptions : Instantiate ListBillingUnitsOptions
func (*EnterpriseBillingUnitsV1) NewListBillingUnitsOptions() *ListBillingUnitsOptions {
	return &ListBillingUnitsOptions{}
}

// SetAccountID : Allow user to set AccountID
func (_options *ListBillingUnitsOptions) SetAccountID(accountID string) *ListBillingUnitsOptions {
	_options.AccountID = core.StringPtr(accountID)
	return _options
}

// SetEnterpriseID : Allow user to set EnterpriseID
func (_options *ListBillingUnitsOptions) SetEnterpriseID(enterpriseID string) *ListBillingUnitsOptions {
	_options.EnterpriseID = core.StringPtr(enterpriseID)
	return _options
}

// SetAccountGroupID : Allow user to set AccountGroupID
func (_options *ListBillingUnitsOptions) SetAccountGroupID(accountGroupID string) *ListBillingUnitsOptions {
	_options.AccountGroupID = core.StringPtr(accountGroupID)
	return _options
}

// SetLimit : Allow user to set Limit
func (_options *ListBillingUnitsOptions) SetLimit(limit int64) *ListBillingUnitsOptions {
	_options.Limit = core.Int64Ptr(limit)
	return _options
}

// SetStart : Allow user to set Start
func (_options *ListBillingUnitsOptions) SetStart(start string) *ListBillingUnitsOptions {
	_options.Start = core.StringPtr(start)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListBillingUnitsOptions) SetHeaders(param map[string]string) *ListBillingUnitsOptions {
	options.Headers = param
	return options
}

// TermCredits : The subscription term that is active in the current month.
type TermCredits struct {
	// The ID of the billing option from which the subscription term is derived.
	BillingOptionID *string `json:"billing_option_id,omitempty"`

	// The category of the credit pool. The valid values are `PLATFORM`, `OFFER`, or `SERVICE` for platform credit and
	// `SUPPORT` for support credit.
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
// The category of the credit pool. The valid values are `PLATFORM`, `OFFER`, or `SERVICE` for platform credit and
// `SUPPORT` for support credit.
const (
	TermCreditsCategoryOfferConst    = "OFFER"
	TermCreditsCategoryPlatformConst = "PLATFORM"
	TermCreditsCategoryServiceConst  = "SERVICE"
	TermCreditsCategorySupportConst  = "SUPPORT"
)

// UnmarshalTermCredits unmarshals an instance of TermCredits from the specified map of raw messages.
func UnmarshalTermCredits(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TermCredits)
	err = core.UnmarshalPrimitive(m, "billing_option_id", &obj.BillingOptionID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "category", &obj.Category)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "start_date", &obj.StartDate)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "end_date", &obj.EndDate)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "total_credits", &obj.TotalCredits)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "starting_balance", &obj.StartingBalance)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "used_credits", &obj.UsedCredits)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "current_balance", &obj.CurrentBalance)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "resources", &obj.Resources)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// BillingUnitsPager can be used to simplify the use of the "ListBillingUnits" method.
type BillingUnitsPager struct {
	hasNext     bool
	options     *ListBillingUnitsOptions
	client      *EnterpriseBillingUnitsV1
	pageContext struct {
		next *string
	}
}

// NewBillingUnitsPager returns a new BillingUnitsPager instance.
func (enterpriseBillingUnits *EnterpriseBillingUnitsV1) NewBillingUnitsPager(options *ListBillingUnitsOptions) (pager *BillingUnitsPager, err error) {
	if options.Start != nil && *options.Start != "" {
		err = fmt.Errorf("the 'options.Start' field should not be set")
		return
	}

	var optionsCopy ListBillingUnitsOptions = *options
	pager = &BillingUnitsPager{
		hasNext: true,
		options: &optionsCopy,
		client:  enterpriseBillingUnits,
	}
	return
}

// HasNext returns true if there are potentially more results to be retrieved.
func (pager *BillingUnitsPager) HasNext() bool {
	return pager.hasNext
}

// GetNextWithContext returns the next page of results using the specified Context.
func (pager *BillingUnitsPager) GetNextWithContext(ctx context.Context) (page []BillingUnit, err error) {
	if !pager.HasNext() {
		return nil, fmt.Errorf("no more results available")
	}

	pager.options.Start = pager.pageContext.next

	result, _, err := pager.client.ListBillingUnitsWithContext(ctx, pager.options)
	if err != nil {
		return
	}

	var next *string
	if result.NextURL != nil {
		var start *string
		start, err = core.GetQueryParam(result.NextURL, "start")
		if err != nil {
			err = fmt.Errorf("error retrieving 'start' query parameter from URL '%s': %s", *result.NextURL, err.Error())
			return
		}
		next = start
	}
	pager.pageContext.next = next
	pager.hasNext = (pager.pageContext.next != nil)
	page = result.Resources

	return
}

// GetAllWithContext returns all results by invoking GetNextWithContext() repeatedly
// until all pages of results have been retrieved.
func (pager *BillingUnitsPager) GetAllWithContext(ctx context.Context) (allItems []BillingUnit, err error) {
	for pager.HasNext() {
		var nextPage []BillingUnit
		nextPage, err = pager.GetNextWithContext(ctx)
		if err != nil {
			return
		}
		allItems = append(allItems, nextPage...)
	}
	return
}

// GetNext invokes GetNextWithContext() using context.Background() as the Context parameter.
func (pager *BillingUnitsPager) GetNext() (page []BillingUnit, err error) {
	return pager.GetNextWithContext(context.Background())
}

// GetAll invokes GetAllWithContext() using context.Background() as the Context parameter.
func (pager *BillingUnitsPager) GetAll() (allItems []BillingUnit, err error) {
	return pager.GetAllWithContext(context.Background())
}

// BillingOptionsPager can be used to simplify the use of the "ListBillingOptions" method.
type BillingOptionsPager struct {
	hasNext     bool
	options     *ListBillingOptionsOptions
	client      *EnterpriseBillingUnitsV1
	pageContext struct {
		next *string
	}
}

// NewBillingOptionsPager returns a new BillingOptionsPager instance.
func (enterpriseBillingUnits *EnterpriseBillingUnitsV1) NewBillingOptionsPager(options *ListBillingOptionsOptions) (pager *BillingOptionsPager, err error) {
	if options.Start != nil && *options.Start != "" {
		err = fmt.Errorf("the 'options.Start' field should not be set")
		return
	}

	var optionsCopy ListBillingOptionsOptions = *options
	pager = &BillingOptionsPager{
		hasNext: true,
		options: &optionsCopy,
		client:  enterpriseBillingUnits,
	}
	return
}

// HasNext returns true if there are potentially more results to be retrieved.
func (pager *BillingOptionsPager) HasNext() bool {
	return pager.hasNext
}

// GetNextWithContext returns the next page of results using the specified Context.
func (pager *BillingOptionsPager) GetNextWithContext(ctx context.Context) (page []BillingOption, err error) {
	if !pager.HasNext() {
		return nil, fmt.Errorf("no more results available")
	}

	pager.options.Start = pager.pageContext.next

	result, _, err := pager.client.ListBillingOptionsWithContext(ctx, pager.options)
	if err != nil {
		return
	}

	var next *string
	if result.NextURL != nil {
		var start *string
		start, err = core.GetQueryParam(result.NextURL, "start")
		if err != nil {
			err = fmt.Errorf("error retrieving 'start' query parameter from URL '%s': %s", *result.NextURL, err.Error())
			return
		}
		next = start
	}
	pager.pageContext.next = next
	pager.hasNext = (pager.pageContext.next != nil)
	page = result.Resources

	return
}

// GetAllWithContext returns all results by invoking GetNextWithContext() repeatedly
// until all pages of results have been retrieved.
func (pager *BillingOptionsPager) GetAllWithContext(ctx context.Context) (allItems []BillingOption, err error) {
	for pager.HasNext() {
		var nextPage []BillingOption
		nextPage, err = pager.GetNextWithContext(ctx)
		if err != nil {
			return
		}
		allItems = append(allItems, nextPage...)
	}
	return
}

// GetNext invokes GetNextWithContext() using context.Background() as the Context parameter.
func (pager *BillingOptionsPager) GetNext() (page []BillingOption, err error) {
	return pager.GetNextWithContext(context.Background())
}

// GetAll invokes GetAllWithContext() using context.Background() as the Context parameter.
func (pager *BillingOptionsPager) GetAll() (allItems []BillingOption, err error) {
	return pager.GetAllWithContext(context.Background())
}
