/**
 * (C) Copyright IBM Corp. 2022.
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

package enterpriseusagereportsv1_test

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/enterpriseusagereportsv1"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`EnterpriseUsageReportsV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			enterpriseUsageReportsService, serviceErr := enterpriseusagereportsv1.NewEnterpriseUsageReportsV1(&enterpriseusagereportsv1.EnterpriseUsageReportsV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(enterpriseUsageReportsService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			enterpriseUsageReportsService, serviceErr := enterpriseusagereportsv1.NewEnterpriseUsageReportsV1(&enterpriseusagereportsv1.EnterpriseUsageReportsV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(enterpriseUsageReportsService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			enterpriseUsageReportsService, serviceErr := enterpriseusagereportsv1.NewEnterpriseUsageReportsV1(&enterpriseusagereportsv1.EnterpriseUsageReportsV1Options{
				URL: "https://enterpriseusagereportsv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(enterpriseUsageReportsService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"ENTERPRISE_USAGE_REPORTS_URL":       "https://enterpriseusagereportsv1/api",
				"ENTERPRISE_USAGE_REPORTS_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				enterpriseUsageReportsService, serviceErr := enterpriseusagereportsv1.NewEnterpriseUsageReportsV1UsingExternalConfig(&enterpriseusagereportsv1.EnterpriseUsageReportsV1Options{})
				Expect(enterpriseUsageReportsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := enterpriseUsageReportsService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != enterpriseUsageReportsService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(enterpriseUsageReportsService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(enterpriseUsageReportsService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				enterpriseUsageReportsService, serviceErr := enterpriseusagereportsv1.NewEnterpriseUsageReportsV1UsingExternalConfig(&enterpriseusagereportsv1.EnterpriseUsageReportsV1Options{
					URL: "https://testService/api",
				})
				Expect(enterpriseUsageReportsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(enterpriseUsageReportsService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := enterpriseUsageReportsService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != enterpriseUsageReportsService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(enterpriseUsageReportsService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(enterpriseUsageReportsService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				enterpriseUsageReportsService, serviceErr := enterpriseusagereportsv1.NewEnterpriseUsageReportsV1UsingExternalConfig(&enterpriseusagereportsv1.EnterpriseUsageReportsV1Options{})
				err := enterpriseUsageReportsService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(enterpriseUsageReportsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(enterpriseUsageReportsService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := enterpriseUsageReportsService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != enterpriseUsageReportsService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(enterpriseUsageReportsService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(enterpriseUsageReportsService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"ENTERPRISE_USAGE_REPORTS_URL":       "https://enterpriseusagereportsv1/api",
				"ENTERPRISE_USAGE_REPORTS_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			enterpriseUsageReportsService, serviceErr := enterpriseusagereportsv1.NewEnterpriseUsageReportsV1UsingExternalConfig(&enterpriseusagereportsv1.EnterpriseUsageReportsV1Options{})

			It(`Instantiate service client with error`, func() {
				Expect(enterpriseUsageReportsService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"ENTERPRISE_USAGE_REPORTS_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			enterpriseUsageReportsService, serviceErr := enterpriseusagereportsv1.NewEnterpriseUsageReportsV1UsingExternalConfig(&enterpriseusagereportsv1.EnterpriseUsageReportsV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(enterpriseUsageReportsService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = enterpriseusagereportsv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`GetResourceUsageReport(getResourceUsageReportOptions *GetResourceUsageReportOptions) - Operation response error`, func() {
		getResourceUsageReportPath := "/v1/resource-usage-reports"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getResourceUsageReportPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["enterprise_id"]).To(Equal([]string{"abc12340d4bf4e36b0423d209b286f24"}))
					Expect(req.URL.Query()["account_group_id"]).To(Equal([]string{"def456a237b94b9a9238ef024e204c9f"}))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"987abcba31834216b8c726a7dd9eb8d6"}))
					// TODO: Add check for children query parameter
					Expect(req.URL.Query()["month"]).To(Equal([]string{"2019-06"}))
					Expect(req.URL.Query()["billing_unit_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetResourceUsageReport with error: Operation response processing error`, func() {
				enterpriseUsageReportsService, serviceErr := enterpriseusagereportsv1.NewEnterpriseUsageReportsV1(&enterpriseusagereportsv1.EnterpriseUsageReportsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(enterpriseUsageReportsService).ToNot(BeNil())

				// Construct an instance of the GetResourceUsageReportOptions model
				getResourceUsageReportOptionsModel := new(enterpriseusagereportsv1.GetResourceUsageReportOptions)
				getResourceUsageReportOptionsModel.EnterpriseID = core.StringPtr("abc12340d4bf4e36b0423d209b286f24")
				getResourceUsageReportOptionsModel.AccountGroupID = core.StringPtr("def456a237b94b9a9238ef024e204c9f")
				getResourceUsageReportOptionsModel.AccountID = core.StringPtr("987abcba31834216b8c726a7dd9eb8d6")
				getResourceUsageReportOptionsModel.Children = core.BoolPtr(true)
				getResourceUsageReportOptionsModel.Month = core.StringPtr("2019-06")
				getResourceUsageReportOptionsModel.BillingUnitID = core.StringPtr("testString")
				getResourceUsageReportOptionsModel.Limit = core.Int64Ptr(int64(10))
				getResourceUsageReportOptionsModel.Offset = core.StringPtr("testString")
				getResourceUsageReportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := enterpriseUsageReportsService.GetResourceUsageReport(getResourceUsageReportOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				enterpriseUsageReportsService.EnableRetries(0, 0)
				result, response, operationErr = enterpriseUsageReportsService.GetResourceUsageReport(getResourceUsageReportOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetResourceUsageReport(getResourceUsageReportOptions *GetResourceUsageReportOptions)`, func() {
		getResourceUsageReportPath := "/v1/resource-usage-reports"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getResourceUsageReportPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["enterprise_id"]).To(Equal([]string{"abc12340d4bf4e36b0423d209b286f24"}))
					Expect(req.URL.Query()["account_group_id"]).To(Equal([]string{"def456a237b94b9a9238ef024e204c9f"}))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"987abcba31834216b8c726a7dd9eb8d6"}))
					// TODO: Add check for children query parameter
					Expect(req.URL.Query()["month"]).To(Equal([]string{"2019-06"}))
					Expect(req.URL.Query()["billing_unit_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"limit": 5, "first": {"href": "Href"}, "next": {"href": "Href"}, "reports": [{"entity_id": "de129b787b86403db7d3a14be2ae5f76", "entity_type": "enterprise", "entity_crn": "crn:v1:bluemix:public:enterprise::a/e9a57260546c4b4aa9ebfa316a82e56e::enterprise:de129b787b86403db7d3a14be2ae5f76", "entity_name": "Platform-Services", "billing_unit_id": "65719a07280a4022a9efa2f6ff4c3369", "billing_unit_crn": "crn:v1:bluemix:public:billing::a/3f99f8accbc848ea96f3c61a0ae22c44::billing-unit:65719a07280a4022a9efa2f6ff4c3369", "billing_unit_name": "Operations", "country_code": "USA", "currency_code": "USD", "month": "2017-08", "billable_cost": 12, "non_billable_cost": 15, "billable_rated_cost": 17, "non_billable_rated_cost": 20, "resources": [{"resource_id": "ResourceID", "billable_cost": 12, "billable_rated_cost": 17, "non_billable_cost": 15, "non_billable_rated_cost": 20, "plans": [{"plan_id": "PlanID", "pricing_region": "PricingRegion", "pricing_plan_id": "PricingPlanID", "billable": true, "cost": 4, "rated_cost": 9, "usage": [{"metric": "UP-TIME", "unit": "HOURS", "quantity": 711.11, "rateable_quantity": 700, "cost": 123.45, "rated_cost": 130, "price": [{"anyKey": "anyValue"}]}]}]}]}]}`)
				}))
			})
			It(`Invoke GetResourceUsageReport successfully with retries`, func() {
				enterpriseUsageReportsService, serviceErr := enterpriseusagereportsv1.NewEnterpriseUsageReportsV1(&enterpriseusagereportsv1.EnterpriseUsageReportsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(enterpriseUsageReportsService).ToNot(BeNil())
				enterpriseUsageReportsService.EnableRetries(0, 0)

				// Construct an instance of the GetResourceUsageReportOptions model
				getResourceUsageReportOptionsModel := new(enterpriseusagereportsv1.GetResourceUsageReportOptions)
				getResourceUsageReportOptionsModel.EnterpriseID = core.StringPtr("abc12340d4bf4e36b0423d209b286f24")
				getResourceUsageReportOptionsModel.AccountGroupID = core.StringPtr("def456a237b94b9a9238ef024e204c9f")
				getResourceUsageReportOptionsModel.AccountID = core.StringPtr("987abcba31834216b8c726a7dd9eb8d6")
				getResourceUsageReportOptionsModel.Children = core.BoolPtr(true)
				getResourceUsageReportOptionsModel.Month = core.StringPtr("2019-06")
				getResourceUsageReportOptionsModel.BillingUnitID = core.StringPtr("testString")
				getResourceUsageReportOptionsModel.Limit = core.Int64Ptr(int64(10))
				getResourceUsageReportOptionsModel.Offset = core.StringPtr("testString")
				getResourceUsageReportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := enterpriseUsageReportsService.GetResourceUsageReportWithContext(ctx, getResourceUsageReportOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				enterpriseUsageReportsService.DisableRetries()
				result, response, operationErr := enterpriseUsageReportsService.GetResourceUsageReport(getResourceUsageReportOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = enterpriseUsageReportsService.GetResourceUsageReportWithContext(ctx, getResourceUsageReportOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getResourceUsageReportPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["enterprise_id"]).To(Equal([]string{"abc12340d4bf4e36b0423d209b286f24"}))
					Expect(req.URL.Query()["account_group_id"]).To(Equal([]string{"def456a237b94b9a9238ef024e204c9f"}))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"987abcba31834216b8c726a7dd9eb8d6"}))
					// TODO: Add check for children query parameter
					Expect(req.URL.Query()["month"]).To(Equal([]string{"2019-06"}))
					Expect(req.URL.Query()["billing_unit_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"limit": 5, "first": {"href": "Href"}, "next": {"href": "Href"}, "reports": [{"entity_id": "de129b787b86403db7d3a14be2ae5f76", "entity_type": "enterprise", "entity_crn": "crn:v1:bluemix:public:enterprise::a/e9a57260546c4b4aa9ebfa316a82e56e::enterprise:de129b787b86403db7d3a14be2ae5f76", "entity_name": "Platform-Services", "billing_unit_id": "65719a07280a4022a9efa2f6ff4c3369", "billing_unit_crn": "crn:v1:bluemix:public:billing::a/3f99f8accbc848ea96f3c61a0ae22c44::billing-unit:65719a07280a4022a9efa2f6ff4c3369", "billing_unit_name": "Operations", "country_code": "USA", "currency_code": "USD", "month": "2017-08", "billable_cost": 12, "non_billable_cost": 15, "billable_rated_cost": 17, "non_billable_rated_cost": 20, "resources": [{"resource_id": "ResourceID", "billable_cost": 12, "billable_rated_cost": 17, "non_billable_cost": 15, "non_billable_rated_cost": 20, "plans": [{"plan_id": "PlanID", "pricing_region": "PricingRegion", "pricing_plan_id": "PricingPlanID", "billable": true, "cost": 4, "rated_cost": 9, "usage": [{"metric": "UP-TIME", "unit": "HOURS", "quantity": 711.11, "rateable_quantity": 700, "cost": 123.45, "rated_cost": 130, "price": [{"anyKey": "anyValue"}]}]}]}]}]}`)
				}))
			})
			It(`Invoke GetResourceUsageReport successfully`, func() {
				enterpriseUsageReportsService, serviceErr := enterpriseusagereportsv1.NewEnterpriseUsageReportsV1(&enterpriseusagereportsv1.EnterpriseUsageReportsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(enterpriseUsageReportsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := enterpriseUsageReportsService.GetResourceUsageReport(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetResourceUsageReportOptions model
				getResourceUsageReportOptionsModel := new(enterpriseusagereportsv1.GetResourceUsageReportOptions)
				getResourceUsageReportOptionsModel.EnterpriseID = core.StringPtr("abc12340d4bf4e36b0423d209b286f24")
				getResourceUsageReportOptionsModel.AccountGroupID = core.StringPtr("def456a237b94b9a9238ef024e204c9f")
				getResourceUsageReportOptionsModel.AccountID = core.StringPtr("987abcba31834216b8c726a7dd9eb8d6")
				getResourceUsageReportOptionsModel.Children = core.BoolPtr(true)
				getResourceUsageReportOptionsModel.Month = core.StringPtr("2019-06")
				getResourceUsageReportOptionsModel.BillingUnitID = core.StringPtr("testString")
				getResourceUsageReportOptionsModel.Limit = core.Int64Ptr(int64(10))
				getResourceUsageReportOptionsModel.Offset = core.StringPtr("testString")
				getResourceUsageReportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = enterpriseUsageReportsService.GetResourceUsageReport(getResourceUsageReportOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetResourceUsageReport with error: Operation request error`, func() {
				enterpriseUsageReportsService, serviceErr := enterpriseusagereportsv1.NewEnterpriseUsageReportsV1(&enterpriseusagereportsv1.EnterpriseUsageReportsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(enterpriseUsageReportsService).ToNot(BeNil())

				// Construct an instance of the GetResourceUsageReportOptions model
				getResourceUsageReportOptionsModel := new(enterpriseusagereportsv1.GetResourceUsageReportOptions)
				getResourceUsageReportOptionsModel.EnterpriseID = core.StringPtr("abc12340d4bf4e36b0423d209b286f24")
				getResourceUsageReportOptionsModel.AccountGroupID = core.StringPtr("def456a237b94b9a9238ef024e204c9f")
				getResourceUsageReportOptionsModel.AccountID = core.StringPtr("987abcba31834216b8c726a7dd9eb8d6")
				getResourceUsageReportOptionsModel.Children = core.BoolPtr(true)
				getResourceUsageReportOptionsModel.Month = core.StringPtr("2019-06")
				getResourceUsageReportOptionsModel.BillingUnitID = core.StringPtr("testString")
				getResourceUsageReportOptionsModel.Limit = core.Int64Ptr(int64(10))
				getResourceUsageReportOptionsModel.Offset = core.StringPtr("testString")
				getResourceUsageReportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := enterpriseUsageReportsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := enterpriseUsageReportsService.GetResourceUsageReport(getResourceUsageReportOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetResourceUsageReport successfully`, func() {
				enterpriseUsageReportsService, serviceErr := enterpriseusagereportsv1.NewEnterpriseUsageReportsV1(&enterpriseusagereportsv1.EnterpriseUsageReportsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(enterpriseUsageReportsService).ToNot(BeNil())

				// Construct an instance of the GetResourceUsageReportOptions model
				getResourceUsageReportOptionsModel := new(enterpriseusagereportsv1.GetResourceUsageReportOptions)
				getResourceUsageReportOptionsModel.EnterpriseID = core.StringPtr("abc12340d4bf4e36b0423d209b286f24")
				getResourceUsageReportOptionsModel.AccountGroupID = core.StringPtr("def456a237b94b9a9238ef024e204c9f")
				getResourceUsageReportOptionsModel.AccountID = core.StringPtr("987abcba31834216b8c726a7dd9eb8d6")
				getResourceUsageReportOptionsModel.Children = core.BoolPtr(true)
				getResourceUsageReportOptionsModel.Month = core.StringPtr("2019-06")
				getResourceUsageReportOptionsModel.BillingUnitID = core.StringPtr("testString")
				getResourceUsageReportOptionsModel.Limit = core.Int64Ptr(int64(10))
				getResourceUsageReportOptionsModel.Offset = core.StringPtr("testString")
				getResourceUsageReportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := enterpriseUsageReportsService.GetResourceUsageReport(getResourceUsageReportOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Test pagination helper method on response`, func() {
			It(`Invoke GetNextOffset successfully`, func() {
				responseObject := new(enterpriseusagereportsv1.Reports)
				nextObject := new(enterpriseusagereportsv1.Link)
				nextObject.Href = core.StringPtr("ibm.com?offset=abc-123")
				responseObject.Next = nextObject

				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.StringPtr("abc-123")))
			})
			It(`Invoke GetNextOffset without a "Next" property in the response`, func() {
				responseObject := new(enterpriseusagereportsv1.Reports)

				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
			It(`Invoke GetNextOffset without any query params in the "Next" URL`, func() {
				responseObject := new(enterpriseusagereportsv1.Reports)
				nextObject := new(enterpriseusagereportsv1.Link)
				nextObject.Href = core.StringPtr("ibm.com")
				responseObject.Next = nextObject

				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
		})
		Context(`Using mock server endpoint - paginated response`, func() {
			BeforeEach(func() {
				var requestNumber int = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getResourceUsageReportPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					requestNumber++
					if requestNumber == 1 {
						fmt.Fprintf(res, "%s", `{"next":{"href":"https://myhost.com/somePath?offset=1"},"reports":[{"entity_id":"de129b787b86403db7d3a14be2ae5f76","entity_type":"enterprise","entity_crn":"crn:v1:bluemix:public:enterprise::a/e9a57260546c4b4aa9ebfa316a82e56e::enterprise:de129b787b86403db7d3a14be2ae5f76","entity_name":"Platform-Services","billing_unit_id":"65719a07280a4022a9efa2f6ff4c3369","billing_unit_crn":"crn:v1:bluemix:public:billing::a/3f99f8accbc848ea96f3c61a0ae22c44::billing-unit:65719a07280a4022a9efa2f6ff4c3369","billing_unit_name":"Operations","country_code":"USA","currency_code":"USD","month":"2017-08","billable_cost":12,"non_billable_cost":15,"billable_rated_cost":17,"non_billable_rated_cost":20,"resources":[{"resource_id":"ResourceID","billable_cost":12,"billable_rated_cost":17,"non_billable_cost":15,"non_billable_rated_cost":20,"plans":[{"plan_id":"PlanID","pricing_region":"PricingRegion","pricing_plan_id":"PricingPlanID","billable":true,"cost":4,"rated_cost":9,"usage":[{"metric":"UP-TIME","unit":"HOURS","quantity":711.11,"rateable_quantity":700,"cost":123.45,"rated_cost":130,"price":[{"anyKey":"anyValue"}]}]}]}]}],"total_count":2,"limit":1}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"reports":[{"entity_id":"de129b787b86403db7d3a14be2ae5f76","entity_type":"enterprise","entity_crn":"crn:v1:bluemix:public:enterprise::a/e9a57260546c4b4aa9ebfa316a82e56e::enterprise:de129b787b86403db7d3a14be2ae5f76","entity_name":"Platform-Services","billing_unit_id":"65719a07280a4022a9efa2f6ff4c3369","billing_unit_crn":"crn:v1:bluemix:public:billing::a/3f99f8accbc848ea96f3c61a0ae22c44::billing-unit:65719a07280a4022a9efa2f6ff4c3369","billing_unit_name":"Operations","country_code":"USA","currency_code":"USD","month":"2017-08","billable_cost":12,"non_billable_cost":15,"billable_rated_cost":17,"non_billable_rated_cost":20,"resources":[{"resource_id":"ResourceID","billable_cost":12,"billable_rated_cost":17,"non_billable_cost":15,"non_billable_rated_cost":20,"plans":[{"plan_id":"PlanID","pricing_region":"PricingRegion","pricing_plan_id":"PricingPlanID","billable":true,"cost":4,"rated_cost":9,"usage":[{"metric":"UP-TIME","unit":"HOURS","quantity":711.11,"rateable_quantity":700,"cost":123.45,"rated_cost":130,"price":[{"anyKey":"anyValue"}]}]}]}]}],"total_count":2,"limit":1}`)
					} else {
						res.WriteHeader(400)
					}
				}))
			})
			It(`Use GetResourceUsageReportPager.GetNext successfully`, func() {
				enterpriseUsageReportsService, serviceErr := enterpriseusagereportsv1.NewEnterpriseUsageReportsV1(&enterpriseusagereportsv1.EnterpriseUsageReportsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(enterpriseUsageReportsService).ToNot(BeNil())

				getResourceUsageReportOptionsModel := &enterpriseusagereportsv1.GetResourceUsageReportOptions{
					EnterpriseID:   core.StringPtr("abc12340d4bf4e36b0423d209b286f24"),
					AccountGroupID: core.StringPtr("def456a237b94b9a9238ef024e204c9f"),
					AccountID:      core.StringPtr("987abcba31834216b8c726a7dd9eb8d6"),
					Children:       core.BoolPtr(true),
					Month:          core.StringPtr("2019-06"),
					BillingUnitID:  core.StringPtr("testString"),
					Limit:          core.Int64Ptr(int64(10)),
				}

				pager, err := enterpriseUsageReportsService.NewGetResourceUsageReportPager(getResourceUsageReportOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				var allResults []enterpriseusagereportsv1.ResourceUsageReport
				for pager.HasNext() {
					nextPage, err := pager.GetNext()
					Expect(err).To(BeNil())
					Expect(nextPage).ToNot(BeNil())
					allResults = append(allResults, nextPage...)
				}
				Expect(len(allResults)).To(Equal(2))
			})
			It(`Use GetResourceUsageReportPager.GetAll successfully`, func() {
				enterpriseUsageReportsService, serviceErr := enterpriseusagereportsv1.NewEnterpriseUsageReportsV1(&enterpriseusagereportsv1.EnterpriseUsageReportsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(enterpriseUsageReportsService).ToNot(BeNil())

				getResourceUsageReportOptionsModel := &enterpriseusagereportsv1.GetResourceUsageReportOptions{
					EnterpriseID:   core.StringPtr("abc12340d4bf4e36b0423d209b286f24"),
					AccountGroupID: core.StringPtr("def456a237b94b9a9238ef024e204c9f"),
					AccountID:      core.StringPtr("987abcba31834216b8c726a7dd9eb8d6"),
					Children:       core.BoolPtr(true),
					Month:          core.StringPtr("2019-06"),
					BillingUnitID:  core.StringPtr("testString"),
					Limit:          core.Int64Ptr(int64(10)),
				}

				pager, err := enterpriseUsageReportsService.NewGetResourceUsageReportPager(getResourceUsageReportOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				allResults, err := pager.GetAll()
				Expect(err).To(BeNil())
				Expect(allResults).ToNot(BeNil())
				Expect(len(allResults)).To(Equal(2))
			})
		})
	})
	Describe(`Model constructor tests`, func() {
		Context(`Using a service client instance`, func() {
			enterpriseUsageReportsService, _ := enterpriseusagereportsv1.NewEnterpriseUsageReportsV1(&enterpriseusagereportsv1.EnterpriseUsageReportsV1Options{
				URL:           "http://enterpriseusagereportsv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewGetResourceUsageReportOptions successfully`, func() {
				// Construct an instance of the GetResourceUsageReportOptions model
				getResourceUsageReportOptionsModel := enterpriseUsageReportsService.NewGetResourceUsageReportOptions()
				getResourceUsageReportOptionsModel.SetEnterpriseID("abc12340d4bf4e36b0423d209b286f24")
				getResourceUsageReportOptionsModel.SetAccountGroupID("def456a237b94b9a9238ef024e204c9f")
				getResourceUsageReportOptionsModel.SetAccountID("987abcba31834216b8c726a7dd9eb8d6")
				getResourceUsageReportOptionsModel.SetChildren(true)
				getResourceUsageReportOptionsModel.SetMonth("2019-06")
				getResourceUsageReportOptionsModel.SetBillingUnitID("testString")
				getResourceUsageReportOptionsModel.SetLimit(int64(10))
				getResourceUsageReportOptionsModel.SetOffset("testString")
				getResourceUsageReportOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getResourceUsageReportOptionsModel).ToNot(BeNil())
				Expect(getResourceUsageReportOptionsModel.EnterpriseID).To(Equal(core.StringPtr("abc12340d4bf4e36b0423d209b286f24")))
				Expect(getResourceUsageReportOptionsModel.AccountGroupID).To(Equal(core.StringPtr("def456a237b94b9a9238ef024e204c9f")))
				Expect(getResourceUsageReportOptionsModel.AccountID).To(Equal(core.StringPtr("987abcba31834216b8c726a7dd9eb8d6")))
				Expect(getResourceUsageReportOptionsModel.Children).To(Equal(core.BoolPtr(true)))
				Expect(getResourceUsageReportOptionsModel.Month).To(Equal(core.StringPtr("2019-06")))
				Expect(getResourceUsageReportOptionsModel.BillingUnitID).To(Equal(core.StringPtr("testString")))
				Expect(getResourceUsageReportOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(10))))
				Expect(getResourceUsageReportOptionsModel.Offset).To(Equal(core.StringPtr("testString")))
				Expect(getResourceUsageReportOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
		})
	})
	Describe(`Utility function tests`, func() {
		It(`Invoke CreateMockByteArray() successfully`, func() {
			mockByteArray := CreateMockByteArray("This is a test")
			Expect(mockByteArray).ToNot(BeNil())
		})
		It(`Invoke CreateMockUUID() successfully`, func() {
			mockUUID := CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
			Expect(mockUUID).ToNot(BeNil())
		})
		It(`Invoke CreateMockReader() successfully`, func() {
			mockReader := CreateMockReader("This is a test.")
			Expect(mockReader).ToNot(BeNil())
		})
		It(`Invoke CreateMockDate() successfully`, func() {
			mockDate := CreateMockDate("2019-01-01")
			Expect(mockDate).ToNot(BeNil())
		})
		It(`Invoke CreateMockDateTime() successfully`, func() {
			mockDateTime := CreateMockDateTime("2019-01-01T12:00:00.000Z")
			Expect(mockDateTime).ToNot(BeNil())
		})
	})
})

//
// Utility functions used by the generated test code
//

func CreateMockByteArray(mockData string) *[]byte {
	ba := make([]byte, 0)
	ba = append(ba, mockData...)
	return &ba
}

func CreateMockUUID(mockData string) *strfmt.UUID {
	uuid := strfmt.UUID(mockData)
	return &uuid
}

func CreateMockReader(mockData string) io.ReadCloser {
	return io.NopCloser(bytes.NewReader([]byte(mockData)))
}

func CreateMockDate(mockData string) *strfmt.Date {
	d, err := core.ParseDate(mockData)
	if err != nil {
		return nil
	}
	return &d
}

func CreateMockDateTime(mockData string) *strfmt.DateTime {
	d, err := core.ParseDateTime(mockData)
	if err != nil {
		return nil
	}
	return &d
}

func SetTestEnvironment(testEnvironment map[string]string) {
	for key, value := range testEnvironment {
		os.Setenv(key, value)
	}
}

func ClearTestEnvironment(testEnvironment map[string]string) {
	for key := range testEnvironment {
		os.Unsetenv(key)
	}
}
