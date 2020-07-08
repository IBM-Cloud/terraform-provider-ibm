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

package wafrulepackagesapiv1_test

import (
	"bytes"
	"fmt"
	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.ibm.com/ibmcloud/networking-go-sdk/wafrulepackagesapiv1"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"time"
)

var _ = Describe(`WafRulePackagesApiV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		crn := "testString"
		zoneID := "testString"
		It(`Instantiate service client`, func() {
			testService, testServiceErr := wafrulepackagesapiv1.NewWafRulePackagesApiV1(&wafrulepackagesapiv1.WafRulePackagesApiV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
				Crn:           core.StringPtr(crn),
				ZoneID:        core.StringPtr(zoneID),
			})
			Expect(testService).ToNot(BeNil())
			Expect(testServiceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			testService, testServiceErr := wafrulepackagesapiv1.NewWafRulePackagesApiV1(&wafrulepackagesapiv1.WafRulePackagesApiV1Options{
				URL:    "{BAD_URL_STRING",
				Crn:    core.StringPtr(crn),
				ZoneID: core.StringPtr(zoneID),
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			testService, testServiceErr := wafrulepackagesapiv1.NewWafRulePackagesApiV1(&wafrulepackagesapiv1.WafRulePackagesApiV1Options{
				URL:    "https://wafrulepackagesapiv1/api",
				Crn:    core.StringPtr(crn),
				ZoneID: core.StringPtr(zoneID),
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Validation Error`, func() {
			testService, testServiceErr := wafrulepackagesapiv1.NewWafRulePackagesApiV1(&wafrulepackagesapiv1.WafRulePackagesApiV1Options{})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		crn := "testString"
		zoneID := "testString"
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"WAF_RULE_PACKAGES_API_URL":       "https://wafrulepackagesapiv1/api",
				"WAF_RULE_PACKAGES_API_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := wafrulepackagesapiv1.NewWafRulePackagesApiV1UsingExternalConfig(&wafrulepackagesapiv1.WafRulePackagesApiV1Options{
					Crn:    core.StringPtr(crn),
					ZoneID: core.StringPtr(zoneID),
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := wafrulepackagesapiv1.NewWafRulePackagesApiV1UsingExternalConfig(&wafrulepackagesapiv1.WafRulePackagesApiV1Options{
					URL:    "https://testService/api",
					Crn:    core.StringPtr(crn),
					ZoneID: core.StringPtr(zoneID),
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				Expect(testService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := wafrulepackagesapiv1.NewWafRulePackagesApiV1UsingExternalConfig(&wafrulepackagesapiv1.WafRulePackagesApiV1Options{
					Crn:    core.StringPtr(crn),
					ZoneID: core.StringPtr(zoneID),
				})
				err := testService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				Expect(testService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"WAF_RULE_PACKAGES_API_URL":       "https://wafrulepackagesapiv1/api",
				"WAF_RULE_PACKAGES_API_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := wafrulepackagesapiv1.NewWafRulePackagesApiV1UsingExternalConfig(&wafrulepackagesapiv1.WafRulePackagesApiV1Options{
				Crn:    core.StringPtr(crn),
				ZoneID: core.StringPtr(zoneID),
			})

			It(`Instantiate service client with error`, func() {
				Expect(testService).To(BeNil())
				Expect(testServiceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"WAF_RULE_PACKAGES_API_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := wafrulepackagesapiv1.NewWafRulePackagesApiV1UsingExternalConfig(&wafrulepackagesapiv1.WafRulePackagesApiV1Options{
				URL:    "{BAD_URL_STRING",
				Crn:    core.StringPtr(crn),
				ZoneID: core.StringPtr(zoneID),
			})

			It(`Instantiate service client with error`, func() {
				Expect(testService).To(BeNil())
				Expect(testServiceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`ListWafPackages(listWafPackagesOptions *ListWafPackagesOptions) - Operation response error`, func() {
		crn := "testString"
		zoneID := "testString"
		listWafPackagesPath := "/v1/testString/zones/testString/firewall/waf/packages"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listWafPackagesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"Wordpress rules"}))

					Expect(req.URL.Query()["page"]).To(Equal([]string{fmt.Sprint(int64(1))}))

					Expect(req.URL.Query()["per_page"]).To(Equal([]string{fmt.Sprint(int64(50))}))

					Expect(req.URL.Query()["order"]).To(Equal([]string{"status"}))

					Expect(req.URL.Query()["direction"]).To(Equal([]string{"desc"}))

					Expect(req.URL.Query()["match"]).To(Equal([]string{"all"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListWafPackages with error: Operation response processing error`, func() {
				testService, testServiceErr := wafrulepackagesapiv1.NewWafRulePackagesApiV1(&wafrulepackagesapiv1.WafRulePackagesApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ListWafPackagesOptions model
				listWafPackagesOptionsModel := new(wafrulepackagesapiv1.ListWafPackagesOptions)
				listWafPackagesOptionsModel.Name = core.StringPtr("Wordpress rules")
				listWafPackagesOptionsModel.Page = core.Int64Ptr(int64(1))
				listWafPackagesOptionsModel.PerPage = core.Int64Ptr(int64(50))
				listWafPackagesOptionsModel.Order = core.StringPtr("status")
				listWafPackagesOptionsModel.Direction = core.StringPtr("desc")
				listWafPackagesOptionsModel.Match = core.StringPtr("all")
				listWafPackagesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.ListWafPackages(listWafPackagesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListWafPackages(listWafPackagesOptions *ListWafPackagesOptions)`, func() {
		crn := "testString"
		zoneID := "testString"
		listWafPackagesPath := "/v1/testString/zones/testString/firewall/waf/packages"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listWafPackagesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"Wordpress rules"}))

					Expect(req.URL.Query()["page"]).To(Equal([]string{fmt.Sprint(int64(1))}))

					Expect(req.URL.Query()["per_page"]).To(Equal([]string{fmt.Sprint(int64(50))}))

					Expect(req.URL.Query()["order"]).To(Equal([]string{"status"}))

					Expect(req.URL.Query()["direction"]).To(Equal([]string{"desc"}))

					Expect(req.URL.Query()["match"]).To(Equal([]string{"all"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": [{"id": "a25a9a7e9c00afc1fb2e0245519d725b", "name": "WordPress rules", "description": "Common WordPress exploit protections", "detection_mode": "traditional", "zone_id": "023e105f4ecef8ad9ca31a8372d0c353", "status": "active"}], "result_info": {"page": 1, "per_page": 2, "count": 1, "total_count": 200}}`)
				}))
			})
			It(`Invoke ListWafPackages successfully`, func() {
				testService, testServiceErr := wafrulepackagesapiv1.NewWafRulePackagesApiV1(&wafrulepackagesapiv1.WafRulePackagesApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListWafPackages(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListWafPackagesOptions model
				listWafPackagesOptionsModel := new(wafrulepackagesapiv1.ListWafPackagesOptions)
				listWafPackagesOptionsModel.Name = core.StringPtr("Wordpress rules")
				listWafPackagesOptionsModel.Page = core.Int64Ptr(int64(1))
				listWafPackagesOptionsModel.PerPage = core.Int64Ptr(int64(50))
				listWafPackagesOptionsModel.Order = core.StringPtr("status")
				listWafPackagesOptionsModel.Direction = core.StringPtr("desc")
				listWafPackagesOptionsModel.Match = core.StringPtr("all")
				listWafPackagesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListWafPackages(listWafPackagesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke ListWafPackages with error: Operation request error`, func() {
				testService, testServiceErr := wafrulepackagesapiv1.NewWafRulePackagesApiV1(&wafrulepackagesapiv1.WafRulePackagesApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ListWafPackagesOptions model
				listWafPackagesOptionsModel := new(wafrulepackagesapiv1.ListWafPackagesOptions)
				listWafPackagesOptionsModel.Name = core.StringPtr("Wordpress rules")
				listWafPackagesOptionsModel.Page = core.Int64Ptr(int64(1))
				listWafPackagesOptionsModel.PerPage = core.Int64Ptr(int64(50))
				listWafPackagesOptionsModel.Order = core.StringPtr("status")
				listWafPackagesOptionsModel.Direction = core.StringPtr("desc")
				listWafPackagesOptionsModel.Match = core.StringPtr("all")
				listWafPackagesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.ListWafPackages(listWafPackagesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetWafPackage(getWafPackageOptions *GetWafPackageOptions) - Operation response error`, func() {
		crn := "testString"
		zoneID := "testString"
		getWafPackagePath := "/v1/testString/zones/testString/firewall/waf/packages/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getWafPackagePath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetWafPackage with error: Operation response processing error`, func() {
				testService, testServiceErr := wafrulepackagesapiv1.NewWafRulePackagesApiV1(&wafrulepackagesapiv1.WafRulePackagesApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetWafPackageOptions model
				getWafPackageOptionsModel := new(wafrulepackagesapiv1.GetWafPackageOptions)
				getWafPackageOptionsModel.PackageID = core.StringPtr("testString")
				getWafPackageOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.GetWafPackage(getWafPackageOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetWafPackage(getWafPackageOptions *GetWafPackageOptions)`, func() {
		crn := "testString"
		zoneID := "testString"
		getWafPackagePath := "/v1/testString/zones/testString/firewall/waf/packages/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getWafPackagePath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "a25a9a7e9c00afc1fb2e0245519d725b", "name": "WordPress rules", "description": "Common WordPress exploit protections", "detection_mode": "traditional", "zone_id": "023e105f4ecef8ad9ca31a8372d0c353", "status": "active", "sensitivity": "high", "action_mode": "challenge"}}`)
				}))
			})
			It(`Invoke GetWafPackage successfully`, func() {
				testService, testServiceErr := wafrulepackagesapiv1.NewWafRulePackagesApiV1(&wafrulepackagesapiv1.WafRulePackagesApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetWafPackage(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetWafPackageOptions model
				getWafPackageOptionsModel := new(wafrulepackagesapiv1.GetWafPackageOptions)
				getWafPackageOptionsModel.PackageID = core.StringPtr("testString")
				getWafPackageOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetWafPackage(getWafPackageOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetWafPackage with error: Operation validation and request error`, func() {
				testService, testServiceErr := wafrulepackagesapiv1.NewWafRulePackagesApiV1(&wafrulepackagesapiv1.WafRulePackagesApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetWafPackageOptions model
				getWafPackageOptionsModel := new(wafrulepackagesapiv1.GetWafPackageOptions)
				getWafPackageOptionsModel.PackageID = core.StringPtr("testString")
				getWafPackageOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.GetWafPackage(getWafPackageOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetWafPackageOptions model with no property values
				getWafPackageOptionsModelNew := new(wafrulepackagesapiv1.GetWafPackageOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.GetWafPackage(getWafPackageOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateWafPackage(updateWafPackageOptions *UpdateWafPackageOptions) - Operation response error`, func() {
		crn := "testString"
		zoneID := "testString"
		updateWafPackagePath := "/v1/testString/zones/testString/firewall/waf/packages/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(updateWafPackagePath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateWafPackage with error: Operation response processing error`, func() {
				testService, testServiceErr := wafrulepackagesapiv1.NewWafRulePackagesApiV1(&wafrulepackagesapiv1.WafRulePackagesApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the UpdateWafPackageOptions model
				updateWafPackageOptionsModel := new(wafrulepackagesapiv1.UpdateWafPackageOptions)
				updateWafPackageOptionsModel.PackageID = core.StringPtr("testString")
				updateWafPackageOptionsModel.Sensitivity = core.StringPtr("high")
				updateWafPackageOptionsModel.ActionMode = core.StringPtr("simulate")
				updateWafPackageOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.UpdateWafPackage(updateWafPackageOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateWafPackage(updateWafPackageOptions *UpdateWafPackageOptions)`, func() {
		crn := "testString"
		zoneID := "testString"
		updateWafPackagePath := "/v1/testString/zones/testString/firewall/waf/packages/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(updateWafPackagePath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "a25a9a7e9c00afc1fb2e0245519d725b", "name": "WordPress rules", "description": "Common WordPress exploit protections", "detection_mode": "traditional", "zone_id": "023e105f4ecef8ad9ca31a8372d0c353", "status": "active", "sensitivity": "high", "action_mode": "challenge"}}`)
				}))
			})
			It(`Invoke UpdateWafPackage successfully`, func() {
				testService, testServiceErr := wafrulepackagesapiv1.NewWafRulePackagesApiV1(&wafrulepackagesapiv1.WafRulePackagesApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.UpdateWafPackage(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateWafPackageOptions model
				updateWafPackageOptionsModel := new(wafrulepackagesapiv1.UpdateWafPackageOptions)
				updateWafPackageOptionsModel.PackageID = core.StringPtr("testString")
				updateWafPackageOptionsModel.Sensitivity = core.StringPtr("high")
				updateWafPackageOptionsModel.ActionMode = core.StringPtr("simulate")
				updateWafPackageOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.UpdateWafPackage(updateWafPackageOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke UpdateWafPackage with error: Operation validation and request error`, func() {
				testService, testServiceErr := wafrulepackagesapiv1.NewWafRulePackagesApiV1(&wafrulepackagesapiv1.WafRulePackagesApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the UpdateWafPackageOptions model
				updateWafPackageOptionsModel := new(wafrulepackagesapiv1.UpdateWafPackageOptions)
				updateWafPackageOptionsModel.PackageID = core.StringPtr("testString")
				updateWafPackageOptionsModel.Sensitivity = core.StringPtr("high")
				updateWafPackageOptionsModel.ActionMode = core.StringPtr("simulate")
				updateWafPackageOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.UpdateWafPackage(updateWafPackageOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateWafPackageOptions model with no property values
				updateWafPackageOptionsModelNew := new(wafrulepackagesapiv1.UpdateWafPackageOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.UpdateWafPackage(updateWafPackageOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Model constructor tests`, func() {
		Context(`Using a service client instance`, func() {
			crn := "testString"
			zoneID := "testString"
			testService, _ := wafrulepackagesapiv1.NewWafRulePackagesApiV1(&wafrulepackagesapiv1.WafRulePackagesApiV1Options{
				URL:           "http://wafrulepackagesapiv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
				Crn:           core.StringPtr(crn),
				ZoneID:        core.StringPtr(zoneID),
			})
			It(`Invoke NewGetWafPackageOptions successfully`, func() {
				// Construct an instance of the GetWafPackageOptions model
				packageID := "testString"
				getWafPackageOptionsModel := testService.NewGetWafPackageOptions(packageID)
				getWafPackageOptionsModel.SetPackageID("testString")
				getWafPackageOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getWafPackageOptionsModel).ToNot(BeNil())
				Expect(getWafPackageOptionsModel.PackageID).To(Equal(core.StringPtr("testString")))
				Expect(getWafPackageOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListWafPackagesOptions successfully`, func() {
				// Construct an instance of the ListWafPackagesOptions model
				listWafPackagesOptionsModel := testService.NewListWafPackagesOptions()
				listWafPackagesOptionsModel.SetName("Wordpress rules")
				listWafPackagesOptionsModel.SetPage(int64(1))
				listWafPackagesOptionsModel.SetPerPage(int64(50))
				listWafPackagesOptionsModel.SetOrder("status")
				listWafPackagesOptionsModel.SetDirection("desc")
				listWafPackagesOptionsModel.SetMatch("all")
				listWafPackagesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listWafPackagesOptionsModel).ToNot(BeNil())
				Expect(listWafPackagesOptionsModel.Name).To(Equal(core.StringPtr("Wordpress rules")))
				Expect(listWafPackagesOptionsModel.Page).To(Equal(core.Int64Ptr(int64(1))))
				Expect(listWafPackagesOptionsModel.PerPage).To(Equal(core.Int64Ptr(int64(50))))
				Expect(listWafPackagesOptionsModel.Order).To(Equal(core.StringPtr("status")))
				Expect(listWafPackagesOptionsModel.Direction).To(Equal(core.StringPtr("desc")))
				Expect(listWafPackagesOptionsModel.Match).To(Equal(core.StringPtr("all")))
				Expect(listWafPackagesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateWafPackageOptions successfully`, func() {
				// Construct an instance of the UpdateWafPackageOptions model
				packageID := "testString"
				updateWafPackageOptionsModel := testService.NewUpdateWafPackageOptions(packageID)
				updateWafPackageOptionsModel.SetPackageID("testString")
				updateWafPackageOptionsModel.SetSensitivity("high")
				updateWafPackageOptionsModel.SetActionMode("simulate")
				updateWafPackageOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateWafPackageOptionsModel).ToNot(BeNil())
				Expect(updateWafPackageOptionsModel.PackageID).To(Equal(core.StringPtr("testString")))
				Expect(updateWafPackageOptionsModel.Sensitivity).To(Equal(core.StringPtr("high")))
				Expect(updateWafPackageOptionsModel.ActionMode).To(Equal(core.StringPtr("simulate")))
				Expect(updateWafPackageOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
			mockDate := CreateMockDate()
			Expect(mockDate).ToNot(BeNil())
		})
		It(`Invoke CreateMockDateTime() successfully`, func() {
			mockDateTime := CreateMockDateTime()
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
	return ioutil.NopCloser(bytes.NewReader([]byte(mockData)))
}

func CreateMockDate() *strfmt.Date {
	d := strfmt.Date(time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))
	return &d
}

func CreateMockDateTime() *strfmt.DateTime {
	d := strfmt.DateTime(time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))
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
