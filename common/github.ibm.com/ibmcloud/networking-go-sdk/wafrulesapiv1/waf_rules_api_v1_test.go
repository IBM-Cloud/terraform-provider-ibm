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

package wafrulesapiv1_test

import (
	"bytes"
	"fmt"
	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.ibm.com/ibmcloud/networking-go-sdk/wafrulesapiv1"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"time"
)

var _ = Describe(`WafRulesApiV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		crn := "testString"
		zoneID := "testString"
		It(`Instantiate service client`, func() {
			testService, testServiceErr := wafrulesapiv1.NewWafRulesApiV1(&wafrulesapiv1.WafRulesApiV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
				Crn:           core.StringPtr(crn),
				ZoneID:        core.StringPtr(zoneID),
			})
			Expect(testService).ToNot(BeNil())
			Expect(testServiceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			testService, testServiceErr := wafrulesapiv1.NewWafRulesApiV1(&wafrulesapiv1.WafRulesApiV1Options{
				URL:    "{BAD_URL_STRING",
				Crn:    core.StringPtr(crn),
				ZoneID: core.StringPtr(zoneID),
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			testService, testServiceErr := wafrulesapiv1.NewWafRulesApiV1(&wafrulesapiv1.WafRulesApiV1Options{
				URL:    "https://wafrulesapiv1/api",
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
			testService, testServiceErr := wafrulesapiv1.NewWafRulesApiV1(&wafrulesapiv1.WafRulesApiV1Options{})
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
				"WAF_RULES_API_URL":       "https://wafrulesapiv1/api",
				"WAF_RULES_API_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := wafrulesapiv1.NewWafRulesApiV1UsingExternalConfig(&wafrulesapiv1.WafRulesApiV1Options{
					Crn:    core.StringPtr(crn),
					ZoneID: core.StringPtr(zoneID),
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := wafrulesapiv1.NewWafRulesApiV1UsingExternalConfig(&wafrulesapiv1.WafRulesApiV1Options{
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
				testService, testServiceErr := wafrulesapiv1.NewWafRulesApiV1UsingExternalConfig(&wafrulesapiv1.WafRulesApiV1Options{
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
				"WAF_RULES_API_URL":       "https://wafrulesapiv1/api",
				"WAF_RULES_API_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := wafrulesapiv1.NewWafRulesApiV1UsingExternalConfig(&wafrulesapiv1.WafRulesApiV1Options{
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
				"WAF_RULES_API_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := wafrulesapiv1.NewWafRulesApiV1UsingExternalConfig(&wafrulesapiv1.WafRulesApiV1Options{
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
	Describe(`ListWafRules(listWafRulesOptions *ListWafRulesOptions) - Operation response error`, func() {
		crn := "testString"
		zoneID := "testString"
		listWafRulesPath := "/v1/testString/zones/testString/firewall/waf/packages/testString/rules"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listWafRulesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["mode"]).To(Equal([]string{"on"}))

					Expect(req.URL.Query()["priority"]).To(Equal([]string{"5"}))

					Expect(req.URL.Query()["match"]).To(Equal([]string{"all"}))

					Expect(req.URL.Query()["order"]).To(Equal([]string{"status"}))

					Expect(req.URL.Query()["group_id"]).To(Equal([]string{"de677e5818985db1285d0e80225f06e5"}))

					Expect(req.URL.Query()["description"]).To(Equal([]string{"SQL injection prevention for SELECT statements"}))

					Expect(req.URL.Query()["direction"]).To(Equal([]string{"desc"}))

					Expect(req.URL.Query()["page"]).To(Equal([]string{fmt.Sprint(int64(1))}))

					Expect(req.URL.Query()["per_page"]).To(Equal([]string{fmt.Sprint(int64(50))}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListWafRules with error: Operation response processing error`, func() {
				testService, testServiceErr := wafrulesapiv1.NewWafRulesApiV1(&wafrulesapiv1.WafRulesApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ListWafRulesOptions model
				listWafRulesOptionsModel := new(wafrulesapiv1.ListWafRulesOptions)
				listWafRulesOptionsModel.PackageID = core.StringPtr("testString")
				listWafRulesOptionsModel.Mode = core.StringPtr("on")
				listWafRulesOptionsModel.Priority = core.StringPtr("5")
				listWafRulesOptionsModel.Match = core.StringPtr("all")
				listWafRulesOptionsModel.Order = core.StringPtr("status")
				listWafRulesOptionsModel.GroupID = core.StringPtr("de677e5818985db1285d0e80225f06e5")
				listWafRulesOptionsModel.Description = core.StringPtr("SQL injection prevention for SELECT statements")
				listWafRulesOptionsModel.Direction = core.StringPtr("desc")
				listWafRulesOptionsModel.Page = core.Int64Ptr(int64(1))
				listWafRulesOptionsModel.PerPage = core.Int64Ptr(int64(50))
				listWafRulesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.ListWafRules(listWafRulesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListWafRules(listWafRulesOptions *ListWafRulesOptions)`, func() {
		crn := "testString"
		zoneID := "testString"
		listWafRulesPath := "/v1/testString/zones/testString/firewall/waf/packages/testString/rules"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listWafRulesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["mode"]).To(Equal([]string{"on"}))

					Expect(req.URL.Query()["priority"]).To(Equal([]string{"5"}))

					Expect(req.URL.Query()["match"]).To(Equal([]string{"all"}))

					Expect(req.URL.Query()["order"]).To(Equal([]string{"status"}))

					Expect(req.URL.Query()["group_id"]).To(Equal([]string{"de677e5818985db1285d0e80225f06e5"}))

					Expect(req.URL.Query()["description"]).To(Equal([]string{"SQL injection prevention for SELECT statements"}))

					Expect(req.URL.Query()["direction"]).To(Equal([]string{"desc"}))

					Expect(req.URL.Query()["page"]).To(Equal([]string{fmt.Sprint(int64(1))}))

					Expect(req.URL.Query()["per_page"]).To(Equal([]string{fmt.Sprint(int64(50))}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": [{"id": "f939de3be84e66e757adcdcb87908023", "description": "SQL injection prevention for SELECT statements", "priority": "5", "group": {"id": "de677e5818985db1285d0e80225f06e5", "name": "Project abc"}, "package_id": "a25a9a7e9c00afc1fb2e0245519d725b", "allowed_modes": ["AllowedModes"], "mode": "on"}], "result_info": {"page": 1, "per_page": 20, "count": 1, "total_count": 2000}}`)
				}))
			})
			It(`Invoke ListWafRules successfully`, func() {
				testService, testServiceErr := wafrulesapiv1.NewWafRulesApiV1(&wafrulesapiv1.WafRulesApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListWafRules(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListWafRulesOptions model
				listWafRulesOptionsModel := new(wafrulesapiv1.ListWafRulesOptions)
				listWafRulesOptionsModel.PackageID = core.StringPtr("testString")
				listWafRulesOptionsModel.Mode = core.StringPtr("on")
				listWafRulesOptionsModel.Priority = core.StringPtr("5")
				listWafRulesOptionsModel.Match = core.StringPtr("all")
				listWafRulesOptionsModel.Order = core.StringPtr("status")
				listWafRulesOptionsModel.GroupID = core.StringPtr("de677e5818985db1285d0e80225f06e5")
				listWafRulesOptionsModel.Description = core.StringPtr("SQL injection prevention for SELECT statements")
				listWafRulesOptionsModel.Direction = core.StringPtr("desc")
				listWafRulesOptionsModel.Page = core.Int64Ptr(int64(1))
				listWafRulesOptionsModel.PerPage = core.Int64Ptr(int64(50))
				listWafRulesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListWafRules(listWafRulesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke ListWafRules with error: Operation validation and request error`, func() {
				testService, testServiceErr := wafrulesapiv1.NewWafRulesApiV1(&wafrulesapiv1.WafRulesApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ListWafRulesOptions model
				listWafRulesOptionsModel := new(wafrulesapiv1.ListWafRulesOptions)
				listWafRulesOptionsModel.PackageID = core.StringPtr("testString")
				listWafRulesOptionsModel.Mode = core.StringPtr("on")
				listWafRulesOptionsModel.Priority = core.StringPtr("5")
				listWafRulesOptionsModel.Match = core.StringPtr("all")
				listWafRulesOptionsModel.Order = core.StringPtr("status")
				listWafRulesOptionsModel.GroupID = core.StringPtr("de677e5818985db1285d0e80225f06e5")
				listWafRulesOptionsModel.Description = core.StringPtr("SQL injection prevention for SELECT statements")
				listWafRulesOptionsModel.Direction = core.StringPtr("desc")
				listWafRulesOptionsModel.Page = core.Int64Ptr(int64(1))
				listWafRulesOptionsModel.PerPage = core.Int64Ptr(int64(50))
				listWafRulesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.ListWafRules(listWafRulesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListWafRulesOptions model with no property values
				listWafRulesOptionsModelNew := new(wafrulesapiv1.ListWafRulesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.ListWafRules(listWafRulesOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetWafRule(getWafRuleOptions *GetWafRuleOptions) - Operation response error`, func() {
		crn := "testString"
		zoneID := "testString"
		getWafRulePath := "/v1/testString/zones/testString/firewall/waf/packages/testString/rules/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getWafRulePath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetWafRule with error: Operation response processing error`, func() {
				testService, testServiceErr := wafrulesapiv1.NewWafRulesApiV1(&wafrulesapiv1.WafRulesApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetWafRuleOptions model
				getWafRuleOptionsModel := new(wafrulesapiv1.GetWafRuleOptions)
				getWafRuleOptionsModel.PackageID = core.StringPtr("testString")
				getWafRuleOptionsModel.Identifier = core.StringPtr("testString")
				getWafRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.GetWafRule(getWafRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetWafRule(getWafRuleOptions *GetWafRuleOptions)`, func() {
		crn := "testString"
		zoneID := "testString"
		getWafRulePath := "/v1/testString/zones/testString/firewall/waf/packages/testString/rules/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getWafRulePath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "f939de3be84e66e757adcdcb87908023", "description": "SQL injection prevention for SELECT statements", "priority": "5", "group": {"id": "de677e5818985db1285d0e80225f06e5", "name": "Project abc"}, "package_id": "a25a9a7e9c00afc1fb2e0245519d725b", "allowed_modes": ["AllowedModes"], "mode": "on"}}`)
				}))
			})
			It(`Invoke GetWafRule successfully`, func() {
				testService, testServiceErr := wafrulesapiv1.NewWafRulesApiV1(&wafrulesapiv1.WafRulesApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetWafRule(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetWafRuleOptions model
				getWafRuleOptionsModel := new(wafrulesapiv1.GetWafRuleOptions)
				getWafRuleOptionsModel.PackageID = core.StringPtr("testString")
				getWafRuleOptionsModel.Identifier = core.StringPtr("testString")
				getWafRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetWafRule(getWafRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetWafRule with error: Operation validation and request error`, func() {
				testService, testServiceErr := wafrulesapiv1.NewWafRulesApiV1(&wafrulesapiv1.WafRulesApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetWafRuleOptions model
				getWafRuleOptionsModel := new(wafrulesapiv1.GetWafRuleOptions)
				getWafRuleOptionsModel.PackageID = core.StringPtr("testString")
				getWafRuleOptionsModel.Identifier = core.StringPtr("testString")
				getWafRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.GetWafRule(getWafRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetWafRuleOptions model with no property values
				getWafRuleOptionsModelNew := new(wafrulesapiv1.GetWafRuleOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.GetWafRule(getWafRuleOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateWafRule(updateWafRuleOptions *UpdateWafRuleOptions) - Operation response error`, func() {
		crn := "testString"
		zoneID := "testString"
		updateWafRulePath := "/v1/testString/zones/testString/firewall/waf/packages/testString/rules/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(updateWafRulePath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateWafRule with error: Operation response processing error`, func() {
				testService, testServiceErr := wafrulesapiv1.NewWafRulesApiV1(&wafrulesapiv1.WafRulesApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the WafRuleBodyCis model
				wafRuleBodyCisModel := new(wafrulesapiv1.WafRuleBodyCis)
				wafRuleBodyCisModel.Mode = core.StringPtr("default")

				// Construct an instance of the WafRuleBodyOwasp model
				wafRuleBodyOwaspModel := new(wafrulesapiv1.WafRuleBodyOwasp)
				wafRuleBodyOwaspModel.Mode = core.StringPtr("on")

				// Construct an instance of the UpdateWafRuleOptions model
				updateWafRuleOptionsModel := new(wafrulesapiv1.UpdateWafRuleOptions)
				updateWafRuleOptionsModel.PackageID = core.StringPtr("testString")
				updateWafRuleOptionsModel.Identifier = core.StringPtr("testString")
				updateWafRuleOptionsModel.Cis = wafRuleBodyCisModel
				updateWafRuleOptionsModel.Owasp = wafRuleBodyOwaspModel
				updateWafRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.UpdateWafRule(updateWafRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateWafRule(updateWafRuleOptions *UpdateWafRuleOptions)`, func() {
		crn := "testString"
		zoneID := "testString"
		updateWafRulePath := "/v1/testString/zones/testString/firewall/waf/packages/testString/rules/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(updateWafRulePath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "f939de3be84e66e757adcdcb87908023", "description": "SQL injection prevention for SELECT statements", "priority": "5", "group": {"id": "de677e5818985db1285d0e80225f06e5", "name": "Project abc"}, "package_id": "a25a9a7e9c00afc1fb2e0245519d725b", "allowed_modes": ["AllowedModes"], "mode": "on"}}`)
				}))
			})
			It(`Invoke UpdateWafRule successfully`, func() {
				testService, testServiceErr := wafrulesapiv1.NewWafRulesApiV1(&wafrulesapiv1.WafRulesApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.UpdateWafRule(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the WafRuleBodyCis model
				wafRuleBodyCisModel := new(wafrulesapiv1.WafRuleBodyCis)
				wafRuleBodyCisModel.Mode = core.StringPtr("default")

				// Construct an instance of the WafRuleBodyOwasp model
				wafRuleBodyOwaspModel := new(wafrulesapiv1.WafRuleBodyOwasp)
				wafRuleBodyOwaspModel.Mode = core.StringPtr("on")

				// Construct an instance of the UpdateWafRuleOptions model
				updateWafRuleOptionsModel := new(wafrulesapiv1.UpdateWafRuleOptions)
				updateWafRuleOptionsModel.PackageID = core.StringPtr("testString")
				updateWafRuleOptionsModel.Identifier = core.StringPtr("testString")
				updateWafRuleOptionsModel.Cis = wafRuleBodyCisModel
				updateWafRuleOptionsModel.Owasp = wafRuleBodyOwaspModel
				updateWafRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.UpdateWafRule(updateWafRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke UpdateWafRule with error: Operation validation and request error`, func() {
				testService, testServiceErr := wafrulesapiv1.NewWafRulesApiV1(&wafrulesapiv1.WafRulesApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the WafRuleBodyCis model
				wafRuleBodyCisModel := new(wafrulesapiv1.WafRuleBodyCis)
				wafRuleBodyCisModel.Mode = core.StringPtr("default")

				// Construct an instance of the WafRuleBodyOwasp model
				wafRuleBodyOwaspModel := new(wafrulesapiv1.WafRuleBodyOwasp)
				wafRuleBodyOwaspModel.Mode = core.StringPtr("on")

				// Construct an instance of the UpdateWafRuleOptions model
				updateWafRuleOptionsModel := new(wafrulesapiv1.UpdateWafRuleOptions)
				updateWafRuleOptionsModel.PackageID = core.StringPtr("testString")
				updateWafRuleOptionsModel.Identifier = core.StringPtr("testString")
				updateWafRuleOptionsModel.Cis = wafRuleBodyCisModel
				updateWafRuleOptionsModel.Owasp = wafRuleBodyOwaspModel
				updateWafRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.UpdateWafRule(updateWafRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateWafRuleOptions model with no property values
				updateWafRuleOptionsModelNew := new(wafrulesapiv1.UpdateWafRuleOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.UpdateWafRule(updateWafRuleOptionsModelNew)
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
			testService, _ := wafrulesapiv1.NewWafRulesApiV1(&wafrulesapiv1.WafRulesApiV1Options{
				URL:           "http://wafrulesapiv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
				Crn:           core.StringPtr(crn),
				ZoneID:        core.StringPtr(zoneID),
			})
			It(`Invoke NewGetWafRuleOptions successfully`, func() {
				// Construct an instance of the GetWafRuleOptions model
				packageID := "testString"
				identifier := "testString"
				getWafRuleOptionsModel := testService.NewGetWafRuleOptions(packageID, identifier)
				getWafRuleOptionsModel.SetPackageID("testString")
				getWafRuleOptionsModel.SetIdentifier("testString")
				getWafRuleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getWafRuleOptionsModel).ToNot(BeNil())
				Expect(getWafRuleOptionsModel.PackageID).To(Equal(core.StringPtr("testString")))
				Expect(getWafRuleOptionsModel.Identifier).To(Equal(core.StringPtr("testString")))
				Expect(getWafRuleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListWafRulesOptions successfully`, func() {
				// Construct an instance of the ListWafRulesOptions model
				packageID := "testString"
				listWafRulesOptionsModel := testService.NewListWafRulesOptions(packageID)
				listWafRulesOptionsModel.SetPackageID("testString")
				listWafRulesOptionsModel.SetMode("on")
				listWafRulesOptionsModel.SetPriority("5")
				listWafRulesOptionsModel.SetMatch("all")
				listWafRulesOptionsModel.SetOrder("status")
				listWafRulesOptionsModel.SetGroupID("de677e5818985db1285d0e80225f06e5")
				listWafRulesOptionsModel.SetDescription("SQL injection prevention for SELECT statements")
				listWafRulesOptionsModel.SetDirection("desc")
				listWafRulesOptionsModel.SetPage(int64(1))
				listWafRulesOptionsModel.SetPerPage(int64(50))
				listWafRulesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listWafRulesOptionsModel).ToNot(BeNil())
				Expect(listWafRulesOptionsModel.PackageID).To(Equal(core.StringPtr("testString")))
				Expect(listWafRulesOptionsModel.Mode).To(Equal(core.StringPtr("on")))
				Expect(listWafRulesOptionsModel.Priority).To(Equal(core.StringPtr("5")))
				Expect(listWafRulesOptionsModel.Match).To(Equal(core.StringPtr("all")))
				Expect(listWafRulesOptionsModel.Order).To(Equal(core.StringPtr("status")))
				Expect(listWafRulesOptionsModel.GroupID).To(Equal(core.StringPtr("de677e5818985db1285d0e80225f06e5")))
				Expect(listWafRulesOptionsModel.Description).To(Equal(core.StringPtr("SQL injection prevention for SELECT statements")))
				Expect(listWafRulesOptionsModel.Direction).To(Equal(core.StringPtr("desc")))
				Expect(listWafRulesOptionsModel.Page).To(Equal(core.Int64Ptr(int64(1))))
				Expect(listWafRulesOptionsModel.PerPage).To(Equal(core.Int64Ptr(int64(50))))
				Expect(listWafRulesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateWafRuleOptions successfully`, func() {
				// Construct an instance of the WafRuleBodyCis model
				wafRuleBodyCisModel := new(wafrulesapiv1.WafRuleBodyCis)
				Expect(wafRuleBodyCisModel).ToNot(BeNil())
				wafRuleBodyCisModel.Mode = core.StringPtr("default")
				Expect(wafRuleBodyCisModel.Mode).To(Equal(core.StringPtr("default")))

				// Construct an instance of the WafRuleBodyOwasp model
				wafRuleBodyOwaspModel := new(wafrulesapiv1.WafRuleBodyOwasp)
				Expect(wafRuleBodyOwaspModel).ToNot(BeNil())
				wafRuleBodyOwaspModel.Mode = core.StringPtr("on")
				Expect(wafRuleBodyOwaspModel.Mode).To(Equal(core.StringPtr("on")))

				// Construct an instance of the UpdateWafRuleOptions model
				packageID := "testString"
				identifier := "testString"
				updateWafRuleOptionsModel := testService.NewUpdateWafRuleOptions(packageID, identifier)
				updateWafRuleOptionsModel.SetPackageID("testString")
				updateWafRuleOptionsModel.SetIdentifier("testString")
				updateWafRuleOptionsModel.SetCis(wafRuleBodyCisModel)
				updateWafRuleOptionsModel.SetOwasp(wafRuleBodyOwaspModel)
				updateWafRuleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateWafRuleOptionsModel).ToNot(BeNil())
				Expect(updateWafRuleOptionsModel.PackageID).To(Equal(core.StringPtr("testString")))
				Expect(updateWafRuleOptionsModel.Identifier).To(Equal(core.StringPtr("testString")))
				Expect(updateWafRuleOptionsModel.Cis).To(Equal(wafRuleBodyCisModel))
				Expect(updateWafRuleOptionsModel.Owasp).To(Equal(wafRuleBodyOwaspModel))
				Expect(updateWafRuleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewWafRuleBodyCis successfully`, func() {
				mode := "default"
				model, err := testService.NewWafRuleBodyCis(mode)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewWafRuleBodyOwasp successfully`, func() {
				mode := "on"
				model, err := testService.NewWafRuleBodyOwasp(mode)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
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
