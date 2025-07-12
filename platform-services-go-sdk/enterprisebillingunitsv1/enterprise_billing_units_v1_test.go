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

package enterprisebillingunitsv1_test

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
	"github.com/IBM/platform-services-go-sdk/enterprisebillingunitsv1"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`EnterpriseBillingUnitsV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			enterpriseBillingUnitsService, serviceErr := enterprisebillingunitsv1.NewEnterpriseBillingUnitsV1(&enterprisebillingunitsv1.EnterpriseBillingUnitsV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(enterpriseBillingUnitsService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			enterpriseBillingUnitsService, serviceErr := enterprisebillingunitsv1.NewEnterpriseBillingUnitsV1(&enterprisebillingunitsv1.EnterpriseBillingUnitsV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(enterpriseBillingUnitsService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			enterpriseBillingUnitsService, serviceErr := enterprisebillingunitsv1.NewEnterpriseBillingUnitsV1(&enterprisebillingunitsv1.EnterpriseBillingUnitsV1Options{
				URL: "https://enterprisebillingunitsv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(enterpriseBillingUnitsService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"ENTERPRISE_BILLING_UNITS_URL":       "https://enterprisebillingunitsv1/api",
				"ENTERPRISE_BILLING_UNITS_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				enterpriseBillingUnitsService, serviceErr := enterprisebillingunitsv1.NewEnterpriseBillingUnitsV1UsingExternalConfig(&enterprisebillingunitsv1.EnterpriseBillingUnitsV1Options{})
				Expect(enterpriseBillingUnitsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := enterpriseBillingUnitsService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != enterpriseBillingUnitsService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(enterpriseBillingUnitsService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(enterpriseBillingUnitsService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				enterpriseBillingUnitsService, serviceErr := enterprisebillingunitsv1.NewEnterpriseBillingUnitsV1UsingExternalConfig(&enterprisebillingunitsv1.EnterpriseBillingUnitsV1Options{
					URL: "https://testService/api",
				})
				Expect(enterpriseBillingUnitsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(enterpriseBillingUnitsService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := enterpriseBillingUnitsService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != enterpriseBillingUnitsService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(enterpriseBillingUnitsService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(enterpriseBillingUnitsService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				enterpriseBillingUnitsService, serviceErr := enterprisebillingunitsv1.NewEnterpriseBillingUnitsV1UsingExternalConfig(&enterprisebillingunitsv1.EnterpriseBillingUnitsV1Options{})
				err := enterpriseBillingUnitsService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(enterpriseBillingUnitsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(enterpriseBillingUnitsService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := enterpriseBillingUnitsService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != enterpriseBillingUnitsService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(enterpriseBillingUnitsService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(enterpriseBillingUnitsService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"ENTERPRISE_BILLING_UNITS_URL":       "https://enterprisebillingunitsv1/api",
				"ENTERPRISE_BILLING_UNITS_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			enterpriseBillingUnitsService, serviceErr := enterprisebillingunitsv1.NewEnterpriseBillingUnitsV1UsingExternalConfig(&enterprisebillingunitsv1.EnterpriseBillingUnitsV1Options{})

			It(`Instantiate service client with error`, func() {
				Expect(enterpriseBillingUnitsService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"ENTERPRISE_BILLING_UNITS_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			enterpriseBillingUnitsService, serviceErr := enterprisebillingunitsv1.NewEnterpriseBillingUnitsV1UsingExternalConfig(&enterprisebillingunitsv1.EnterpriseBillingUnitsV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(enterpriseBillingUnitsService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = enterprisebillingunitsv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`GetBillingUnit(getBillingUnitOptions *GetBillingUnitOptions) - Operation response error`, func() {
		getBillingUnitPath := "/v1/billing-units/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getBillingUnitPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetBillingUnit with error: Operation response processing error`, func() {
				enterpriseBillingUnitsService, serviceErr := enterprisebillingunitsv1.NewEnterpriseBillingUnitsV1(&enterprisebillingunitsv1.EnterpriseBillingUnitsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(enterpriseBillingUnitsService).ToNot(BeNil())

				// Construct an instance of the GetBillingUnitOptions model
				getBillingUnitOptionsModel := new(enterprisebillingunitsv1.GetBillingUnitOptions)
				getBillingUnitOptionsModel.BillingUnitID = core.StringPtr("testString")
				getBillingUnitOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := enterpriseBillingUnitsService.GetBillingUnit(getBillingUnitOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				enterpriseBillingUnitsService.EnableRetries(0, 0)
				result, response, operationErr = enterpriseBillingUnitsService.GetBillingUnit(getBillingUnitOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetBillingUnit(getBillingUnitOptions *GetBillingUnitOptions)`, func() {
		getBillingUnitPath := "/v1/billing-units/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getBillingUnitPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "crn": "crn:v1:bluemix:public:billing::a/<<enterprise_account_id>>::billing-unit:<<billing_unit_id>>", "name": "Name", "enterprise_id": "EnterpriseID", "currency_code": "USD", "country_code": "USA", "master": true, "created_at": "2019-05-01T00:00:00.000Z"}`)
				}))
			})
			It(`Invoke GetBillingUnit successfully with retries`, func() {
				enterpriseBillingUnitsService, serviceErr := enterprisebillingunitsv1.NewEnterpriseBillingUnitsV1(&enterprisebillingunitsv1.EnterpriseBillingUnitsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(enterpriseBillingUnitsService).ToNot(BeNil())
				enterpriseBillingUnitsService.EnableRetries(0, 0)

				// Construct an instance of the GetBillingUnitOptions model
				getBillingUnitOptionsModel := new(enterprisebillingunitsv1.GetBillingUnitOptions)
				getBillingUnitOptionsModel.BillingUnitID = core.StringPtr("testString")
				getBillingUnitOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := enterpriseBillingUnitsService.GetBillingUnitWithContext(ctx, getBillingUnitOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				enterpriseBillingUnitsService.DisableRetries()
				result, response, operationErr := enterpriseBillingUnitsService.GetBillingUnit(getBillingUnitOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = enterpriseBillingUnitsService.GetBillingUnitWithContext(ctx, getBillingUnitOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getBillingUnitPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "crn": "crn:v1:bluemix:public:billing::a/<<enterprise_account_id>>::billing-unit:<<billing_unit_id>>", "name": "Name", "enterprise_id": "EnterpriseID", "currency_code": "USD", "country_code": "USA", "master": true, "created_at": "2019-05-01T00:00:00.000Z"}`)
				}))
			})
			It(`Invoke GetBillingUnit successfully`, func() {
				enterpriseBillingUnitsService, serviceErr := enterprisebillingunitsv1.NewEnterpriseBillingUnitsV1(&enterprisebillingunitsv1.EnterpriseBillingUnitsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(enterpriseBillingUnitsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := enterpriseBillingUnitsService.GetBillingUnit(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetBillingUnitOptions model
				getBillingUnitOptionsModel := new(enterprisebillingunitsv1.GetBillingUnitOptions)
				getBillingUnitOptionsModel.BillingUnitID = core.StringPtr("testString")
				getBillingUnitOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = enterpriseBillingUnitsService.GetBillingUnit(getBillingUnitOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetBillingUnit with error: Operation validation and request error`, func() {
				enterpriseBillingUnitsService, serviceErr := enterprisebillingunitsv1.NewEnterpriseBillingUnitsV1(&enterprisebillingunitsv1.EnterpriseBillingUnitsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(enterpriseBillingUnitsService).ToNot(BeNil())

				// Construct an instance of the GetBillingUnitOptions model
				getBillingUnitOptionsModel := new(enterprisebillingunitsv1.GetBillingUnitOptions)
				getBillingUnitOptionsModel.BillingUnitID = core.StringPtr("testString")
				getBillingUnitOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := enterpriseBillingUnitsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := enterpriseBillingUnitsService.GetBillingUnit(getBillingUnitOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetBillingUnitOptions model with no property values
				getBillingUnitOptionsModelNew := new(enterprisebillingunitsv1.GetBillingUnitOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = enterpriseBillingUnitsService.GetBillingUnit(getBillingUnitOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
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
			It(`Invoke GetBillingUnit successfully`, func() {
				enterpriseBillingUnitsService, serviceErr := enterprisebillingunitsv1.NewEnterpriseBillingUnitsV1(&enterprisebillingunitsv1.EnterpriseBillingUnitsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(enterpriseBillingUnitsService).ToNot(BeNil())

				// Construct an instance of the GetBillingUnitOptions model
				getBillingUnitOptionsModel := new(enterprisebillingunitsv1.GetBillingUnitOptions)
				getBillingUnitOptionsModel.BillingUnitID = core.StringPtr("testString")
				getBillingUnitOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := enterpriseBillingUnitsService.GetBillingUnit(getBillingUnitOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListBillingUnits(listBillingUnitsOptions *ListBillingUnitsOptions) - Operation response error`, func() {
		listBillingUnitsPath := "/v1/billing-units"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listBillingUnitsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["enterprise_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["account_group_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListBillingUnits with error: Operation response processing error`, func() {
				enterpriseBillingUnitsService, serviceErr := enterprisebillingunitsv1.NewEnterpriseBillingUnitsV1(&enterprisebillingunitsv1.EnterpriseBillingUnitsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(enterpriseBillingUnitsService).ToNot(BeNil())

				// Construct an instance of the ListBillingUnitsOptions model
				listBillingUnitsOptionsModel := new(enterprisebillingunitsv1.ListBillingUnitsOptions)
				listBillingUnitsOptionsModel.AccountID = core.StringPtr("testString")
				listBillingUnitsOptionsModel.EnterpriseID = core.StringPtr("testString")
				listBillingUnitsOptionsModel.AccountGroupID = core.StringPtr("testString")
				listBillingUnitsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listBillingUnitsOptionsModel.Start = core.StringPtr("testString")
				listBillingUnitsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := enterpriseBillingUnitsService.ListBillingUnits(listBillingUnitsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				enterpriseBillingUnitsService.EnableRetries(0, 0)
				result, response, operationErr = enterpriseBillingUnitsService.ListBillingUnits(listBillingUnitsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListBillingUnits(listBillingUnitsOptions *ListBillingUnitsOptions)`, func() {
		listBillingUnitsPath := "/v1/billing-units"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listBillingUnitsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["enterprise_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["account_group_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"rows_count": 9, "next_url": "NextURL", "resources": [{"id": "ID", "crn": "crn:v1:bluemix:public:billing::a/<<enterprise_account_id>>::billing-unit:<<billing_unit_id>>", "name": "Name", "enterprise_id": "EnterpriseID", "currency_code": "USD", "country_code": "USA", "master": true, "created_at": "2019-05-01T00:00:00.000Z"}]}`)
				}))
			})
			It(`Invoke ListBillingUnits successfully with retries`, func() {
				enterpriseBillingUnitsService, serviceErr := enterprisebillingunitsv1.NewEnterpriseBillingUnitsV1(&enterprisebillingunitsv1.EnterpriseBillingUnitsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(enterpriseBillingUnitsService).ToNot(BeNil())
				enterpriseBillingUnitsService.EnableRetries(0, 0)

				// Construct an instance of the ListBillingUnitsOptions model
				listBillingUnitsOptionsModel := new(enterprisebillingunitsv1.ListBillingUnitsOptions)
				listBillingUnitsOptionsModel.AccountID = core.StringPtr("testString")
				listBillingUnitsOptionsModel.EnterpriseID = core.StringPtr("testString")
				listBillingUnitsOptionsModel.AccountGroupID = core.StringPtr("testString")
				listBillingUnitsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listBillingUnitsOptionsModel.Start = core.StringPtr("testString")
				listBillingUnitsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := enterpriseBillingUnitsService.ListBillingUnitsWithContext(ctx, listBillingUnitsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				enterpriseBillingUnitsService.DisableRetries()
				result, response, operationErr := enterpriseBillingUnitsService.ListBillingUnits(listBillingUnitsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = enterpriseBillingUnitsService.ListBillingUnitsWithContext(ctx, listBillingUnitsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listBillingUnitsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["enterprise_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["account_group_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"rows_count": 9, "next_url": "NextURL", "resources": [{"id": "ID", "crn": "crn:v1:bluemix:public:billing::a/<<enterprise_account_id>>::billing-unit:<<billing_unit_id>>", "name": "Name", "enterprise_id": "EnterpriseID", "currency_code": "USD", "country_code": "USA", "master": true, "created_at": "2019-05-01T00:00:00.000Z"}]}`)
				}))
			})
			It(`Invoke ListBillingUnits successfully`, func() {
				enterpriseBillingUnitsService, serviceErr := enterprisebillingunitsv1.NewEnterpriseBillingUnitsV1(&enterprisebillingunitsv1.EnterpriseBillingUnitsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(enterpriseBillingUnitsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := enterpriseBillingUnitsService.ListBillingUnits(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListBillingUnitsOptions model
				listBillingUnitsOptionsModel := new(enterprisebillingunitsv1.ListBillingUnitsOptions)
				listBillingUnitsOptionsModel.AccountID = core.StringPtr("testString")
				listBillingUnitsOptionsModel.EnterpriseID = core.StringPtr("testString")
				listBillingUnitsOptionsModel.AccountGroupID = core.StringPtr("testString")
				listBillingUnitsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listBillingUnitsOptionsModel.Start = core.StringPtr("testString")
				listBillingUnitsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = enterpriseBillingUnitsService.ListBillingUnits(listBillingUnitsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListBillingUnits with error: Operation request error`, func() {
				enterpriseBillingUnitsService, serviceErr := enterprisebillingunitsv1.NewEnterpriseBillingUnitsV1(&enterprisebillingunitsv1.EnterpriseBillingUnitsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(enterpriseBillingUnitsService).ToNot(BeNil())

				// Construct an instance of the ListBillingUnitsOptions model
				listBillingUnitsOptionsModel := new(enterprisebillingunitsv1.ListBillingUnitsOptions)
				listBillingUnitsOptionsModel.AccountID = core.StringPtr("testString")
				listBillingUnitsOptionsModel.EnterpriseID = core.StringPtr("testString")
				listBillingUnitsOptionsModel.AccountGroupID = core.StringPtr("testString")
				listBillingUnitsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listBillingUnitsOptionsModel.Start = core.StringPtr("testString")
				listBillingUnitsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := enterpriseBillingUnitsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := enterpriseBillingUnitsService.ListBillingUnits(listBillingUnitsOptionsModel)
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
			It(`Invoke ListBillingUnits successfully`, func() {
				enterpriseBillingUnitsService, serviceErr := enterprisebillingunitsv1.NewEnterpriseBillingUnitsV1(&enterprisebillingunitsv1.EnterpriseBillingUnitsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(enterpriseBillingUnitsService).ToNot(BeNil())

				// Construct an instance of the ListBillingUnitsOptions model
				listBillingUnitsOptionsModel := new(enterprisebillingunitsv1.ListBillingUnitsOptions)
				listBillingUnitsOptionsModel.AccountID = core.StringPtr("testString")
				listBillingUnitsOptionsModel.EnterpriseID = core.StringPtr("testString")
				listBillingUnitsOptionsModel.AccountGroupID = core.StringPtr("testString")
				listBillingUnitsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listBillingUnitsOptionsModel.Start = core.StringPtr("testString")
				listBillingUnitsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := enterpriseBillingUnitsService.ListBillingUnits(listBillingUnitsOptionsModel)
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
			It(`Invoke GetNextStart successfully`, func() {
				responseObject := new(enterprisebillingunitsv1.BillingUnitsList)
				responseObject.NextURL = core.StringPtr("ibm.com?start=abc-123")

				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.StringPtr("abc-123")))
			})
			It(`Invoke GetNextStart without a "NextURL" property in the response`, func() {
				responseObject := new(enterprisebillingunitsv1.BillingUnitsList)

				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
			It(`Invoke GetNextStart without any query params in the "NextURL" URL`, func() {
				responseObject := new(enterprisebillingunitsv1.BillingUnitsList)
				responseObject.NextURL = core.StringPtr("ibm.com")

				value, err := responseObject.GetNextStart()
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
					Expect(req.URL.EscapedPath()).To(Equal(listBillingUnitsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					requestNumber++
					if requestNumber == 1 {
						fmt.Fprintf(res, "%s", `{"total_count":2,"limit":1,"next_url":"https://myhost.com/somePath?start=1","resources":[{"id":"ID","crn":"crn:v1:bluemix:public:billing::a/<<enterprise_account_id>>::billing-unit:<<billing_unit_id>>","name":"Name","enterprise_id":"EnterpriseID","currency_code":"USD","country_code":"USA","master":true,"created_at":"2019-05-01T00:00:00.000Z"}]}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"total_count":2,"limit":1,"resources":[{"id":"ID","crn":"crn:v1:bluemix:public:billing::a/<<enterprise_account_id>>::billing-unit:<<billing_unit_id>>","name":"Name","enterprise_id":"EnterpriseID","currency_code":"USD","country_code":"USA","master":true,"created_at":"2019-05-01T00:00:00.000Z"}]}`)
					} else {
						res.WriteHeader(400)
					}
				}))
			})
			It(`Use BillingUnitsPager.GetNext successfully`, func() {
				enterpriseBillingUnitsService, serviceErr := enterprisebillingunitsv1.NewEnterpriseBillingUnitsV1(&enterprisebillingunitsv1.EnterpriseBillingUnitsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(enterpriseBillingUnitsService).ToNot(BeNil())

				listBillingUnitsOptionsModel := &enterprisebillingunitsv1.ListBillingUnitsOptions{
					AccountID:      core.StringPtr("testString"),
					EnterpriseID:   core.StringPtr("testString"),
					AccountGroupID: core.StringPtr("testString"),
					Limit:          core.Int64Ptr(int64(10)),
				}

				pager, err := enterpriseBillingUnitsService.NewBillingUnitsPager(listBillingUnitsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				var allResults []enterprisebillingunitsv1.BillingUnit
				for pager.HasNext() {
					nextPage, err := pager.GetNext()
					Expect(err).To(BeNil())
					Expect(nextPage).ToNot(BeNil())
					allResults = append(allResults, nextPage...)
				}
				Expect(len(allResults)).To(Equal(2))
			})
			It(`Use BillingUnitsPager.GetAll successfully`, func() {
				enterpriseBillingUnitsService, serviceErr := enterprisebillingunitsv1.NewEnterpriseBillingUnitsV1(&enterprisebillingunitsv1.EnterpriseBillingUnitsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(enterpriseBillingUnitsService).ToNot(BeNil())

				listBillingUnitsOptionsModel := &enterprisebillingunitsv1.ListBillingUnitsOptions{
					AccountID:      core.StringPtr("testString"),
					EnterpriseID:   core.StringPtr("testString"),
					AccountGroupID: core.StringPtr("testString"),
					Limit:          core.Int64Ptr(int64(10)),
				}

				pager, err := enterpriseBillingUnitsService.NewBillingUnitsPager(listBillingUnitsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				allResults, err := pager.GetAll()
				Expect(err).To(BeNil())
				Expect(allResults).ToNot(BeNil())
				Expect(len(allResults)).To(Equal(2))
			})
		})
	})
	Describe(`ListBillingOptions(listBillingOptionsOptions *ListBillingOptionsOptions) - Operation response error`, func() {
		listBillingOptionsPath := "/v1/billing-options"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listBillingOptionsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["billing_unit_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListBillingOptions with error: Operation response processing error`, func() {
				enterpriseBillingUnitsService, serviceErr := enterprisebillingunitsv1.NewEnterpriseBillingUnitsV1(&enterprisebillingunitsv1.EnterpriseBillingUnitsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(enterpriseBillingUnitsService).ToNot(BeNil())

				// Construct an instance of the ListBillingOptionsOptions model
				listBillingOptionsOptionsModel := new(enterprisebillingunitsv1.ListBillingOptionsOptions)
				listBillingOptionsOptionsModel.BillingUnitID = core.StringPtr("testString")
				listBillingOptionsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listBillingOptionsOptionsModel.Start = core.StringPtr("testString")
				listBillingOptionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := enterpriseBillingUnitsService.ListBillingOptions(listBillingOptionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				enterpriseBillingUnitsService.EnableRetries(0, 0)
				result, response, operationErr = enterpriseBillingUnitsService.ListBillingOptions(listBillingOptionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListBillingOptions(listBillingOptionsOptions *ListBillingOptionsOptions)`, func() {
		listBillingOptionsPath := "/v1/billing-options"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listBillingOptionsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["billing_unit_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"rows_count": 9, "next_url": "NextURL", "resources": [{"id": "ID", "billing_unit_id": "BillingUnitID", "start_date": "2019-05-01T00:00:00.000Z", "end_date": "2020-05-01T00:00:00.000Z", "state": "ACTIVE", "type": "SUBSCRIPTION", "category": "PLATFORM", "payment_instrument": {"anyKey": "anyValue"}, "duration_in_months": 11, "line_item_id": 10, "billing_system": {"anyKey": "anyValue"}, "renewal_mode_code": "RenewalModeCode", "updated_at": "2019-06-01T00:00:00.000Z"}]}`)
				}))
			})
			It(`Invoke ListBillingOptions successfully with retries`, func() {
				enterpriseBillingUnitsService, serviceErr := enterprisebillingunitsv1.NewEnterpriseBillingUnitsV1(&enterprisebillingunitsv1.EnterpriseBillingUnitsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(enterpriseBillingUnitsService).ToNot(BeNil())
				enterpriseBillingUnitsService.EnableRetries(0, 0)

				// Construct an instance of the ListBillingOptionsOptions model
				listBillingOptionsOptionsModel := new(enterprisebillingunitsv1.ListBillingOptionsOptions)
				listBillingOptionsOptionsModel.BillingUnitID = core.StringPtr("testString")
				listBillingOptionsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listBillingOptionsOptionsModel.Start = core.StringPtr("testString")
				listBillingOptionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := enterpriseBillingUnitsService.ListBillingOptionsWithContext(ctx, listBillingOptionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				enterpriseBillingUnitsService.DisableRetries()
				result, response, operationErr := enterpriseBillingUnitsService.ListBillingOptions(listBillingOptionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = enterpriseBillingUnitsService.ListBillingOptionsWithContext(ctx, listBillingOptionsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listBillingOptionsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["billing_unit_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"rows_count": 9, "next_url": "NextURL", "resources": [{"id": "ID", "billing_unit_id": "BillingUnitID", "start_date": "2019-05-01T00:00:00.000Z", "end_date": "2020-05-01T00:00:00.000Z", "state": "ACTIVE", "type": "SUBSCRIPTION", "category": "PLATFORM", "payment_instrument": {"anyKey": "anyValue"}, "duration_in_months": 11, "line_item_id": 10, "billing_system": {"anyKey": "anyValue"}, "renewal_mode_code": "RenewalModeCode", "updated_at": "2019-06-01T00:00:00.000Z"}]}`)
				}))
			})
			It(`Invoke ListBillingOptions successfully`, func() {
				enterpriseBillingUnitsService, serviceErr := enterprisebillingunitsv1.NewEnterpriseBillingUnitsV1(&enterprisebillingunitsv1.EnterpriseBillingUnitsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(enterpriseBillingUnitsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := enterpriseBillingUnitsService.ListBillingOptions(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListBillingOptionsOptions model
				listBillingOptionsOptionsModel := new(enterprisebillingunitsv1.ListBillingOptionsOptions)
				listBillingOptionsOptionsModel.BillingUnitID = core.StringPtr("testString")
				listBillingOptionsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listBillingOptionsOptionsModel.Start = core.StringPtr("testString")
				listBillingOptionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = enterpriseBillingUnitsService.ListBillingOptions(listBillingOptionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListBillingOptions with error: Operation validation and request error`, func() {
				enterpriseBillingUnitsService, serviceErr := enterprisebillingunitsv1.NewEnterpriseBillingUnitsV1(&enterprisebillingunitsv1.EnterpriseBillingUnitsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(enterpriseBillingUnitsService).ToNot(BeNil())

				// Construct an instance of the ListBillingOptionsOptions model
				listBillingOptionsOptionsModel := new(enterprisebillingunitsv1.ListBillingOptionsOptions)
				listBillingOptionsOptionsModel.BillingUnitID = core.StringPtr("testString")
				listBillingOptionsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listBillingOptionsOptionsModel.Start = core.StringPtr("testString")
				listBillingOptionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := enterpriseBillingUnitsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := enterpriseBillingUnitsService.ListBillingOptions(listBillingOptionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListBillingOptionsOptions model with no property values
				listBillingOptionsOptionsModelNew := new(enterprisebillingunitsv1.ListBillingOptionsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = enterpriseBillingUnitsService.ListBillingOptions(listBillingOptionsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
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
			It(`Invoke ListBillingOptions successfully`, func() {
				enterpriseBillingUnitsService, serviceErr := enterprisebillingunitsv1.NewEnterpriseBillingUnitsV1(&enterprisebillingunitsv1.EnterpriseBillingUnitsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(enterpriseBillingUnitsService).ToNot(BeNil())

				// Construct an instance of the ListBillingOptionsOptions model
				listBillingOptionsOptionsModel := new(enterprisebillingunitsv1.ListBillingOptionsOptions)
				listBillingOptionsOptionsModel.BillingUnitID = core.StringPtr("testString")
				listBillingOptionsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listBillingOptionsOptionsModel.Start = core.StringPtr("testString")
				listBillingOptionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := enterpriseBillingUnitsService.ListBillingOptions(listBillingOptionsOptionsModel)
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
			It(`Invoke GetNextStart successfully`, func() {
				responseObject := new(enterprisebillingunitsv1.BillingOptionsList)
				responseObject.NextURL = core.StringPtr("ibm.com?start=abc-123")

				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.StringPtr("abc-123")))
			})
			It(`Invoke GetNextStart without a "NextURL" property in the response`, func() {
				responseObject := new(enterprisebillingunitsv1.BillingOptionsList)

				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
			It(`Invoke GetNextStart without any query params in the "NextURL" URL`, func() {
				responseObject := new(enterprisebillingunitsv1.BillingOptionsList)
				responseObject.NextURL = core.StringPtr("ibm.com")

				value, err := responseObject.GetNextStart()
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
					Expect(req.URL.EscapedPath()).To(Equal(listBillingOptionsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					requestNumber++
					if requestNumber == 1 {
						fmt.Fprintf(res, "%s", `{"total_count":2,"limit":1,"next_url":"https://myhost.com/somePath?start=1","resources":[{"id":"ID","billing_unit_id":"BillingUnitID","start_date":"2019-05-01T00:00:00.000Z","end_date":"2020-05-01T00:00:00.000Z","state":"ACTIVE","type":"SUBSCRIPTION","category":"PLATFORM","payment_instrument":{"anyKey":"anyValue"},"duration_in_months":11,"line_item_id":10,"billing_system":{"anyKey":"anyValue"},"renewal_mode_code":"RenewalModeCode","updated_at":"2019-06-01T00:00:00.000Z"}]}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"total_count":2,"limit":1,"resources":[{"id":"ID","billing_unit_id":"BillingUnitID","start_date":"2019-05-01T00:00:00.000Z","end_date":"2020-05-01T00:00:00.000Z","state":"ACTIVE","type":"SUBSCRIPTION","category":"PLATFORM","payment_instrument":{"anyKey":"anyValue"},"duration_in_months":11,"line_item_id":10,"billing_system":{"anyKey":"anyValue"},"renewal_mode_code":"RenewalModeCode","updated_at":"2019-06-01T00:00:00.000Z"}]}`)
					} else {
						res.WriteHeader(400)
					}
				}))
			})
			It(`Use BillingOptionsPager.GetNext successfully`, func() {
				enterpriseBillingUnitsService, serviceErr := enterprisebillingunitsv1.NewEnterpriseBillingUnitsV1(&enterprisebillingunitsv1.EnterpriseBillingUnitsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(enterpriseBillingUnitsService).ToNot(BeNil())

				listBillingOptionsOptionsModel := &enterprisebillingunitsv1.ListBillingOptionsOptions{
					BillingUnitID: core.StringPtr("testString"),
					Limit:         core.Int64Ptr(int64(10)),
				}

				pager, err := enterpriseBillingUnitsService.NewBillingOptionsPager(listBillingOptionsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				var allResults []enterprisebillingunitsv1.BillingOption
				for pager.HasNext() {
					nextPage, err := pager.GetNext()
					Expect(err).To(BeNil())
					Expect(nextPage).ToNot(BeNil())
					allResults = append(allResults, nextPage...)
				}
				Expect(len(allResults)).To(Equal(2))
			})
			It(`Use BillingOptionsPager.GetAll successfully`, func() {
				enterpriseBillingUnitsService, serviceErr := enterprisebillingunitsv1.NewEnterpriseBillingUnitsV1(&enterprisebillingunitsv1.EnterpriseBillingUnitsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(enterpriseBillingUnitsService).ToNot(BeNil())

				listBillingOptionsOptionsModel := &enterprisebillingunitsv1.ListBillingOptionsOptions{
					BillingUnitID: core.StringPtr("testString"),
					Limit:         core.Int64Ptr(int64(10)),
				}

				pager, err := enterpriseBillingUnitsService.NewBillingOptionsPager(listBillingOptionsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				allResults, err := pager.GetAll()
				Expect(err).To(BeNil())
				Expect(allResults).ToNot(BeNil())
				Expect(len(allResults)).To(Equal(2))
			})
		})
	})
	Describe(`GetCreditPools(getCreditPoolsOptions *GetCreditPoolsOptions) - Operation response error`, func() {
		getCreditPoolsPath := "/v1/credit-pools"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCreditPoolsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["billing_unit_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["date"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["type"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetCreditPools with error: Operation response processing error`, func() {
				enterpriseBillingUnitsService, serviceErr := enterprisebillingunitsv1.NewEnterpriseBillingUnitsV1(&enterprisebillingunitsv1.EnterpriseBillingUnitsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(enterpriseBillingUnitsService).ToNot(BeNil())

				// Construct an instance of the GetCreditPoolsOptions model
				getCreditPoolsOptionsModel := new(enterprisebillingunitsv1.GetCreditPoolsOptions)
				getCreditPoolsOptionsModel.BillingUnitID = core.StringPtr("testString")
				getCreditPoolsOptionsModel.Date = core.StringPtr("testString")
				getCreditPoolsOptionsModel.Type = core.StringPtr("testString")
				getCreditPoolsOptionsModel.Limit = core.Int64Ptr(int64(10))
				getCreditPoolsOptionsModel.Start = core.StringPtr("testString")
				getCreditPoolsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := enterpriseBillingUnitsService.GetCreditPools(getCreditPoolsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				enterpriseBillingUnitsService.EnableRetries(0, 0)
				result, response, operationErr = enterpriseBillingUnitsService.GetCreditPools(getCreditPoolsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetCreditPools(getCreditPoolsOptions *GetCreditPoolsOptions)`, func() {
		getCreditPoolsPath := "/v1/credit-pools"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCreditPoolsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["billing_unit_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["date"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["type"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"rows_count": 2, "next_url": "NextURL", "resources": [{"type": "PLATFORM", "currency_code": "USD", "billing_unit_id": "BillingUnitID", "term_credits": [{"billing_option_id": "JWX986YRGFSHACQUEFOI", "category": "PLATFORM", "start_date": "2019-05-01T00:00:00.000Z", "end_date": "2020-04-30T23:59:29.999Z", "total_credits": 10000, "starting_balance": 9000, "used_credits": 9500, "current_balance": 0, "resources": [{"anyKey": "anyValue"}]}], "overage": {"cost": 500, "resources": [{"anyKey": "anyValue"}]}}]}`)
				}))
			})
			It(`Invoke GetCreditPools successfully with retries`, func() {
				enterpriseBillingUnitsService, serviceErr := enterprisebillingunitsv1.NewEnterpriseBillingUnitsV1(&enterprisebillingunitsv1.EnterpriseBillingUnitsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(enterpriseBillingUnitsService).ToNot(BeNil())
				enterpriseBillingUnitsService.EnableRetries(0, 0)

				// Construct an instance of the GetCreditPoolsOptions model
				getCreditPoolsOptionsModel := new(enterprisebillingunitsv1.GetCreditPoolsOptions)
				getCreditPoolsOptionsModel.BillingUnitID = core.StringPtr("testString")
				getCreditPoolsOptionsModel.Date = core.StringPtr("testString")
				getCreditPoolsOptionsModel.Type = core.StringPtr("testString")
				getCreditPoolsOptionsModel.Limit = core.Int64Ptr(int64(10))
				getCreditPoolsOptionsModel.Start = core.StringPtr("testString")
				getCreditPoolsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := enterpriseBillingUnitsService.GetCreditPoolsWithContext(ctx, getCreditPoolsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				enterpriseBillingUnitsService.DisableRetries()
				result, response, operationErr := enterpriseBillingUnitsService.GetCreditPools(getCreditPoolsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = enterpriseBillingUnitsService.GetCreditPoolsWithContext(ctx, getCreditPoolsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getCreditPoolsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["billing_unit_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["date"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["type"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"rows_count": 2, "next_url": "NextURL", "resources": [{"type": "PLATFORM", "currency_code": "USD", "billing_unit_id": "BillingUnitID", "term_credits": [{"billing_option_id": "JWX986YRGFSHACQUEFOI", "category": "PLATFORM", "start_date": "2019-05-01T00:00:00.000Z", "end_date": "2020-04-30T23:59:29.999Z", "total_credits": 10000, "starting_balance": 9000, "used_credits": 9500, "current_balance": 0, "resources": [{"anyKey": "anyValue"}]}], "overage": {"cost": 500, "resources": [{"anyKey": "anyValue"}]}}]}`)
				}))
			})
			It(`Invoke GetCreditPools successfully`, func() {
				enterpriseBillingUnitsService, serviceErr := enterprisebillingunitsv1.NewEnterpriseBillingUnitsV1(&enterprisebillingunitsv1.EnterpriseBillingUnitsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(enterpriseBillingUnitsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := enterpriseBillingUnitsService.GetCreditPools(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetCreditPoolsOptions model
				getCreditPoolsOptionsModel := new(enterprisebillingunitsv1.GetCreditPoolsOptions)
				getCreditPoolsOptionsModel.BillingUnitID = core.StringPtr("testString")
				getCreditPoolsOptionsModel.Date = core.StringPtr("testString")
				getCreditPoolsOptionsModel.Type = core.StringPtr("testString")
				getCreditPoolsOptionsModel.Limit = core.Int64Ptr(int64(10))
				getCreditPoolsOptionsModel.Start = core.StringPtr("testString")
				getCreditPoolsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = enterpriseBillingUnitsService.GetCreditPools(getCreditPoolsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetCreditPools with error: Operation validation and request error`, func() {
				enterpriseBillingUnitsService, serviceErr := enterprisebillingunitsv1.NewEnterpriseBillingUnitsV1(&enterprisebillingunitsv1.EnterpriseBillingUnitsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(enterpriseBillingUnitsService).ToNot(BeNil())

				// Construct an instance of the GetCreditPoolsOptions model
				getCreditPoolsOptionsModel := new(enterprisebillingunitsv1.GetCreditPoolsOptions)
				getCreditPoolsOptionsModel.BillingUnitID = core.StringPtr("testString")
				getCreditPoolsOptionsModel.Date = core.StringPtr("testString")
				getCreditPoolsOptionsModel.Type = core.StringPtr("testString")
				getCreditPoolsOptionsModel.Limit = core.Int64Ptr(int64(10))
				getCreditPoolsOptionsModel.Start = core.StringPtr("testString")
				getCreditPoolsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := enterpriseBillingUnitsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := enterpriseBillingUnitsService.GetCreditPools(getCreditPoolsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetCreditPoolsOptions model with no property values
				getCreditPoolsOptionsModelNew := new(enterprisebillingunitsv1.GetCreditPoolsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = enterpriseBillingUnitsService.GetCreditPools(getCreditPoolsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
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
			It(`Invoke GetCreditPools successfully`, func() {
				enterpriseBillingUnitsService, serviceErr := enterprisebillingunitsv1.NewEnterpriseBillingUnitsV1(&enterprisebillingunitsv1.EnterpriseBillingUnitsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(enterpriseBillingUnitsService).ToNot(BeNil())

				// Construct an instance of the GetCreditPoolsOptions model
				getCreditPoolsOptionsModel := new(enterprisebillingunitsv1.GetCreditPoolsOptions)
				getCreditPoolsOptionsModel.BillingUnitID = core.StringPtr("testString")
				getCreditPoolsOptionsModel.Date = core.StringPtr("testString")
				getCreditPoolsOptionsModel.Type = core.StringPtr("testString")
				getCreditPoolsOptionsModel.Limit = core.Int64Ptr(int64(10))
				getCreditPoolsOptionsModel.Start = core.StringPtr("testString")
				getCreditPoolsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := enterpriseBillingUnitsService.GetCreditPools(getCreditPoolsOptionsModel)
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
			It(`Invoke GetNextStart successfully`, func() {
				responseObject := new(enterprisebillingunitsv1.CreditPoolsList)
				responseObject.NextURL = core.StringPtr("ibm.com?start=abc-123")

				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.StringPtr("abc-123")))
			})
			It(`Invoke GetNextStart without a "NextURL" property in the response`, func() {
				responseObject := new(enterprisebillingunitsv1.CreditPoolsList)

				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
			It(`Invoke GetNextStart without any query params in the "NextURL" URL`, func() {
				responseObject := new(enterprisebillingunitsv1.CreditPoolsList)
				responseObject.NextURL = core.StringPtr("ibm.com")

				value, err := responseObject.GetNextStart()
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
					Expect(req.URL.EscapedPath()).To(Equal(getCreditPoolsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					requestNumber++
					if requestNumber == 1 {
						fmt.Fprintf(res, "%s", `{"total_count":2,"limit":1,"next_url":"https://myhost.com/somePath?start=1","resources":[{"type":"PLATFORM","currency_code":"USD","billing_unit_id":"BillingUnitID","term_credits":[{"billing_option_id":"JWX986YRGFSHACQUEFOI","category":"PLATFORM","start_date":"2019-05-01T00:00:00.000Z","end_date":"2020-04-30T23:59:29.999Z","total_credits":10000,"starting_balance":9000,"used_credits":9500,"current_balance":0,"resources":[{"anyKey":"anyValue"}]}],"overage":{"cost":500,"resources":[{"anyKey":"anyValue"}]}}]}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"total_count":2,"limit":1,"resources":[{"type":"PLATFORM","currency_code":"USD","billing_unit_id":"BillingUnitID","term_credits":[{"billing_option_id":"JWX986YRGFSHACQUEFOI","category":"PLATFORM","start_date":"2019-05-01T00:00:00.000Z","end_date":"2020-04-30T23:59:29.999Z","total_credits":10000,"starting_balance":9000,"used_credits":9500,"current_balance":0,"resources":[{"anyKey":"anyValue"}]}],"overage":{"cost":500,"resources":[{"anyKey":"anyValue"}]}}]}`)
					} else {
						res.WriteHeader(400)
					}
				}))
			})
		})
	})
	Describe(`Model constructor tests`, func() {
		Context(`Using a service client instance`, func() {
			enterpriseBillingUnitsService, _ := enterprisebillingunitsv1.NewEnterpriseBillingUnitsV1(&enterprisebillingunitsv1.EnterpriseBillingUnitsV1Options{
				URL:           "http://enterprisebillingunitsv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewGetBillingUnitOptions successfully`, func() {
				// Construct an instance of the GetBillingUnitOptions model
				billingUnitID := "testString"
				getBillingUnitOptionsModel := enterpriseBillingUnitsService.NewGetBillingUnitOptions(billingUnitID)
				getBillingUnitOptionsModel.SetBillingUnitID("testString")
				getBillingUnitOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getBillingUnitOptionsModel).ToNot(BeNil())
				Expect(getBillingUnitOptionsModel.BillingUnitID).To(Equal(core.StringPtr("testString")))
				Expect(getBillingUnitOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetCreditPoolsOptions successfully`, func() {
				// Construct an instance of the GetCreditPoolsOptions model
				billingUnitID := "testString"
				getCreditPoolsOptionsModel := enterpriseBillingUnitsService.NewGetCreditPoolsOptions(billingUnitID)
				getCreditPoolsOptionsModel.SetBillingUnitID("testString")
				getCreditPoolsOptionsModel.SetDate("testString")
				getCreditPoolsOptionsModel.SetType("testString")
				getCreditPoolsOptionsModel.SetLimit(int64(10))
				getCreditPoolsOptionsModel.SetStart("testString")
				getCreditPoolsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getCreditPoolsOptionsModel).ToNot(BeNil())
				Expect(getCreditPoolsOptionsModel.BillingUnitID).To(Equal(core.StringPtr("testString")))
				Expect(getCreditPoolsOptionsModel.Date).To(Equal(core.StringPtr("testString")))
				Expect(getCreditPoolsOptionsModel.Type).To(Equal(core.StringPtr("testString")))
				Expect(getCreditPoolsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(10))))
				Expect(getCreditPoolsOptionsModel.Start).To(Equal(core.StringPtr("testString")))
				Expect(getCreditPoolsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListBillingOptionsOptions successfully`, func() {
				// Construct an instance of the ListBillingOptionsOptions model
				billingUnitID := "testString"
				listBillingOptionsOptionsModel := enterpriseBillingUnitsService.NewListBillingOptionsOptions(billingUnitID)
				listBillingOptionsOptionsModel.SetBillingUnitID("testString")
				listBillingOptionsOptionsModel.SetLimit(int64(10))
				listBillingOptionsOptionsModel.SetStart("testString")
				listBillingOptionsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listBillingOptionsOptionsModel).ToNot(BeNil())
				Expect(listBillingOptionsOptionsModel.BillingUnitID).To(Equal(core.StringPtr("testString")))
				Expect(listBillingOptionsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(10))))
				Expect(listBillingOptionsOptionsModel.Start).To(Equal(core.StringPtr("testString")))
				Expect(listBillingOptionsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListBillingUnitsOptions successfully`, func() {
				// Construct an instance of the ListBillingUnitsOptions model
				listBillingUnitsOptionsModel := enterpriseBillingUnitsService.NewListBillingUnitsOptions()
				listBillingUnitsOptionsModel.SetAccountID("testString")
				listBillingUnitsOptionsModel.SetEnterpriseID("testString")
				listBillingUnitsOptionsModel.SetAccountGroupID("testString")
				listBillingUnitsOptionsModel.SetLimit(int64(10))
				listBillingUnitsOptionsModel.SetStart("testString")
				listBillingUnitsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listBillingUnitsOptionsModel).ToNot(BeNil())
				Expect(listBillingUnitsOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(listBillingUnitsOptionsModel.EnterpriseID).To(Equal(core.StringPtr("testString")))
				Expect(listBillingUnitsOptionsModel.AccountGroupID).To(Equal(core.StringPtr("testString")))
				Expect(listBillingUnitsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(10))))
				Expect(listBillingUnitsOptionsModel.Start).To(Equal(core.StringPtr("testString")))
				Expect(listBillingUnitsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
