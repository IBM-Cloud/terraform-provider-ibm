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

package wafrulegroupsapiv1_test

import (
	"bytes"
	"fmt"
	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.ibm.com/ibmcloud/networking-go-sdk/wafrulegroupsapiv1"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"time"
)

var _ = Describe(`WafRuleGroupsApiV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		crn := "testString"
		zoneID := "testString"
		It(`Instantiate service client`, func() {
			testService, testServiceErr := wafrulegroupsapiv1.NewWafRuleGroupsApiV1(&wafrulegroupsapiv1.WafRuleGroupsApiV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
				Crn:           core.StringPtr(crn),
				ZoneID:        core.StringPtr(zoneID),
			})
			Expect(testService).ToNot(BeNil())
			Expect(testServiceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			testService, testServiceErr := wafrulegroupsapiv1.NewWafRuleGroupsApiV1(&wafrulegroupsapiv1.WafRuleGroupsApiV1Options{
				URL:    "{BAD_URL_STRING",
				Crn:    core.StringPtr(crn),
				ZoneID: core.StringPtr(zoneID),
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			testService, testServiceErr := wafrulegroupsapiv1.NewWafRuleGroupsApiV1(&wafrulegroupsapiv1.WafRuleGroupsApiV1Options{
				URL:    "https://wafrulegroupsapiv1/api",
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
			testService, testServiceErr := wafrulegroupsapiv1.NewWafRuleGroupsApiV1(&wafrulegroupsapiv1.WafRuleGroupsApiV1Options{})
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
				"WAF_RULE_GROUPS_API_URL":       "https://wafrulegroupsapiv1/api",
				"WAF_RULE_GROUPS_API_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := wafrulegroupsapiv1.NewWafRuleGroupsApiV1UsingExternalConfig(&wafrulegroupsapiv1.WafRuleGroupsApiV1Options{
					Crn:    core.StringPtr(crn),
					ZoneID: core.StringPtr(zoneID),
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := wafrulegroupsapiv1.NewWafRuleGroupsApiV1UsingExternalConfig(&wafrulegroupsapiv1.WafRuleGroupsApiV1Options{
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
				testService, testServiceErr := wafrulegroupsapiv1.NewWafRuleGroupsApiV1UsingExternalConfig(&wafrulegroupsapiv1.WafRuleGroupsApiV1Options{
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
				"WAF_RULE_GROUPS_API_URL":       "https://wafrulegroupsapiv1/api",
				"WAF_RULE_GROUPS_API_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := wafrulegroupsapiv1.NewWafRuleGroupsApiV1UsingExternalConfig(&wafrulegroupsapiv1.WafRuleGroupsApiV1Options{
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
				"WAF_RULE_GROUPS_API_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := wafrulegroupsapiv1.NewWafRuleGroupsApiV1UsingExternalConfig(&wafrulegroupsapiv1.WafRuleGroupsApiV1Options{
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
	Describe(`ListWafRuleGroups(listWafRuleGroupsOptions *ListWafRuleGroupsOptions) - Operation response error`, func() {
		crn := "testString"
		zoneID := "testString"
		listWafRuleGroupsPath := "/v1/testString/zones/testString/firewall/waf/packages/testString/groups"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listWafRuleGroupsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"Wordpress rules"}))

					Expect(req.URL.Query()["mode"]).To(Equal([]string{"true"}))

					Expect(req.URL.Query()["rules_count"]).To(Equal([]string{"10"}))

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
			It(`Invoke ListWafRuleGroups with error: Operation response processing error`, func() {
				testService, testServiceErr := wafrulegroupsapiv1.NewWafRuleGroupsApiV1(&wafrulegroupsapiv1.WafRuleGroupsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ListWafRuleGroupsOptions model
				listWafRuleGroupsOptionsModel := new(wafrulegroupsapiv1.ListWafRuleGroupsOptions)
				listWafRuleGroupsOptionsModel.PkgID = core.StringPtr("testString")
				listWafRuleGroupsOptionsModel.Name = core.StringPtr("Wordpress rules")
				listWafRuleGroupsOptionsModel.Mode = core.StringPtr("true")
				listWafRuleGroupsOptionsModel.RulesCount = core.StringPtr("10")
				listWafRuleGroupsOptionsModel.Page = core.Int64Ptr(int64(1))
				listWafRuleGroupsOptionsModel.PerPage = core.Int64Ptr(int64(50))
				listWafRuleGroupsOptionsModel.Order = core.StringPtr("status")
				listWafRuleGroupsOptionsModel.Direction = core.StringPtr("desc")
				listWafRuleGroupsOptionsModel.Match = core.StringPtr("all")
				listWafRuleGroupsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.ListWafRuleGroups(listWafRuleGroupsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListWafRuleGroups(listWafRuleGroupsOptions *ListWafRuleGroupsOptions)`, func() {
		crn := "testString"
		zoneID := "testString"
		listWafRuleGroupsPath := "/v1/testString/zones/testString/firewall/waf/packages/testString/groups"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listWafRuleGroupsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"Wordpress rules"}))

					Expect(req.URL.Query()["mode"]).To(Equal([]string{"true"}))

					Expect(req.URL.Query()["rules_count"]).To(Equal([]string{"10"}))

					Expect(req.URL.Query()["page"]).To(Equal([]string{fmt.Sprint(int64(1))}))

					Expect(req.URL.Query()["per_page"]).To(Equal([]string{fmt.Sprint(int64(50))}))

					Expect(req.URL.Query()["order"]).To(Equal([]string{"status"}))

					Expect(req.URL.Query()["direction"]).To(Equal([]string{"desc"}))

					Expect(req.URL.Query()["match"]).To(Equal([]string{"all"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": [{"id": "a25a9a7e9c00afc1fb2e0245519d725b", "name": "Project Honey Pot", "description": "Group designed to protect against IP addresses that are a threat and typically used to launch DDoS attacks", "rules_count": 10, "modified_rules_count": 10, "package_id": "a25a9a7e9c00afc1fb2e0245519d725b", "mode": "on", "allowed_modes": ["AllowedModes"]}], "result_info": {"page": 1, "per_page": 2, "count": 1, "total_count": 200}}`)
				}))
			})
			It(`Invoke ListWafRuleGroups successfully`, func() {
				testService, testServiceErr := wafrulegroupsapiv1.NewWafRuleGroupsApiV1(&wafrulegroupsapiv1.WafRuleGroupsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListWafRuleGroups(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListWafRuleGroupsOptions model
				listWafRuleGroupsOptionsModel := new(wafrulegroupsapiv1.ListWafRuleGroupsOptions)
				listWafRuleGroupsOptionsModel.PkgID = core.StringPtr("testString")
				listWafRuleGroupsOptionsModel.Name = core.StringPtr("Wordpress rules")
				listWafRuleGroupsOptionsModel.Mode = core.StringPtr("true")
				listWafRuleGroupsOptionsModel.RulesCount = core.StringPtr("10")
				listWafRuleGroupsOptionsModel.Page = core.Int64Ptr(int64(1))
				listWafRuleGroupsOptionsModel.PerPage = core.Int64Ptr(int64(50))
				listWafRuleGroupsOptionsModel.Order = core.StringPtr("status")
				listWafRuleGroupsOptionsModel.Direction = core.StringPtr("desc")
				listWafRuleGroupsOptionsModel.Match = core.StringPtr("all")
				listWafRuleGroupsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListWafRuleGroups(listWafRuleGroupsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke ListWafRuleGroups with error: Operation validation and request error`, func() {
				testService, testServiceErr := wafrulegroupsapiv1.NewWafRuleGroupsApiV1(&wafrulegroupsapiv1.WafRuleGroupsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ListWafRuleGroupsOptions model
				listWafRuleGroupsOptionsModel := new(wafrulegroupsapiv1.ListWafRuleGroupsOptions)
				listWafRuleGroupsOptionsModel.PkgID = core.StringPtr("testString")
				listWafRuleGroupsOptionsModel.Name = core.StringPtr("Wordpress rules")
				listWafRuleGroupsOptionsModel.Mode = core.StringPtr("true")
				listWafRuleGroupsOptionsModel.RulesCount = core.StringPtr("10")
				listWafRuleGroupsOptionsModel.Page = core.Int64Ptr(int64(1))
				listWafRuleGroupsOptionsModel.PerPage = core.Int64Ptr(int64(50))
				listWafRuleGroupsOptionsModel.Order = core.StringPtr("status")
				listWafRuleGroupsOptionsModel.Direction = core.StringPtr("desc")
				listWafRuleGroupsOptionsModel.Match = core.StringPtr("all")
				listWafRuleGroupsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.ListWafRuleGroups(listWafRuleGroupsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListWafRuleGroupsOptions model with no property values
				listWafRuleGroupsOptionsModelNew := new(wafrulegroupsapiv1.ListWafRuleGroupsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.ListWafRuleGroups(listWafRuleGroupsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetWafRuleGroup(getWafRuleGroupOptions *GetWafRuleGroupOptions) - Operation response error`, func() {
		crn := "testString"
		zoneID := "testString"
		getWafRuleGroupPath := "/v1/testString/zones/testString/firewall/waf/packages/testString/groups/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getWafRuleGroupPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetWafRuleGroup with error: Operation response processing error`, func() {
				testService, testServiceErr := wafrulegroupsapiv1.NewWafRuleGroupsApiV1(&wafrulegroupsapiv1.WafRuleGroupsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetWafRuleGroupOptions model
				getWafRuleGroupOptionsModel := new(wafrulegroupsapiv1.GetWafRuleGroupOptions)
				getWafRuleGroupOptionsModel.PkgID = core.StringPtr("testString")
				getWafRuleGroupOptionsModel.GroupID = core.StringPtr("testString")
				getWafRuleGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.GetWafRuleGroup(getWafRuleGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetWafRuleGroup(getWafRuleGroupOptions *GetWafRuleGroupOptions)`, func() {
		crn := "testString"
		zoneID := "testString"
		getWafRuleGroupPath := "/v1/testString/zones/testString/firewall/waf/packages/testString/groups/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getWafRuleGroupPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "a25a9a7e9c00afc1fb2e0245519d725b", "name": "Project Honey Pot", "description": "Group designed to protect against IP addresses that are a threat and typically used to launch DDoS attacks", "rules_count": 10, "modified_rules_count": 10, "package_id": "a25a9a7e9c00afc1fb2e0245519d725b", "mode": "on", "allowed_modes": ["AllowedModes"]}, "result_info": {"page": 1, "per_page": 2, "count": 1, "total_count": 200}}`)
				}))
			})
			It(`Invoke GetWafRuleGroup successfully`, func() {
				testService, testServiceErr := wafrulegroupsapiv1.NewWafRuleGroupsApiV1(&wafrulegroupsapiv1.WafRuleGroupsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetWafRuleGroup(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetWafRuleGroupOptions model
				getWafRuleGroupOptionsModel := new(wafrulegroupsapiv1.GetWafRuleGroupOptions)
				getWafRuleGroupOptionsModel.PkgID = core.StringPtr("testString")
				getWafRuleGroupOptionsModel.GroupID = core.StringPtr("testString")
				getWafRuleGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetWafRuleGroup(getWafRuleGroupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetWafRuleGroup with error: Operation validation and request error`, func() {
				testService, testServiceErr := wafrulegroupsapiv1.NewWafRuleGroupsApiV1(&wafrulegroupsapiv1.WafRuleGroupsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetWafRuleGroupOptions model
				getWafRuleGroupOptionsModel := new(wafrulegroupsapiv1.GetWafRuleGroupOptions)
				getWafRuleGroupOptionsModel.PkgID = core.StringPtr("testString")
				getWafRuleGroupOptionsModel.GroupID = core.StringPtr("testString")
				getWafRuleGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.GetWafRuleGroup(getWafRuleGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetWafRuleGroupOptions model with no property values
				getWafRuleGroupOptionsModelNew := new(wafrulegroupsapiv1.GetWafRuleGroupOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.GetWafRuleGroup(getWafRuleGroupOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateWafRuleGroup(updateWafRuleGroupOptions *UpdateWafRuleGroupOptions) - Operation response error`, func() {
		crn := "testString"
		zoneID := "testString"
		updateWafRuleGroupPath := "/v1/testString/zones/testString/firewall/waf/packages/testString/groups/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(updateWafRuleGroupPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateWafRuleGroup with error: Operation response processing error`, func() {
				testService, testServiceErr := wafrulegroupsapiv1.NewWafRuleGroupsApiV1(&wafrulegroupsapiv1.WafRuleGroupsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the UpdateWafRuleGroupOptions model
				updateWafRuleGroupOptionsModel := new(wafrulegroupsapiv1.UpdateWafRuleGroupOptions)
				updateWafRuleGroupOptionsModel.PkgID = core.StringPtr("testString")
				updateWafRuleGroupOptionsModel.GroupID = core.StringPtr("testString")
				updateWafRuleGroupOptionsModel.Mode = core.StringPtr("on")
				updateWafRuleGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.UpdateWafRuleGroup(updateWafRuleGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateWafRuleGroup(updateWafRuleGroupOptions *UpdateWafRuleGroupOptions)`, func() {
		crn := "testString"
		zoneID := "testString"
		updateWafRuleGroupPath := "/v1/testString/zones/testString/firewall/waf/packages/testString/groups/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(updateWafRuleGroupPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "a25a9a7e9c00afc1fb2e0245519d725b", "name": "Project Honey Pot", "description": "Group designed to protect against IP addresses that are a threat and typically used to launch DDoS attacks", "rules_count": 10, "modified_rules_count": 10, "package_id": "a25a9a7e9c00afc1fb2e0245519d725b", "mode": "on", "allowed_modes": ["AllowedModes"]}, "result_info": {"page": 1, "per_page": 2, "count": 1, "total_count": 200}}`)
				}))
			})
			It(`Invoke UpdateWafRuleGroup successfully`, func() {
				testService, testServiceErr := wafrulegroupsapiv1.NewWafRuleGroupsApiV1(&wafrulegroupsapiv1.WafRuleGroupsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.UpdateWafRuleGroup(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateWafRuleGroupOptions model
				updateWafRuleGroupOptionsModel := new(wafrulegroupsapiv1.UpdateWafRuleGroupOptions)
				updateWafRuleGroupOptionsModel.PkgID = core.StringPtr("testString")
				updateWafRuleGroupOptionsModel.GroupID = core.StringPtr("testString")
				updateWafRuleGroupOptionsModel.Mode = core.StringPtr("on")
				updateWafRuleGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.UpdateWafRuleGroup(updateWafRuleGroupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke UpdateWafRuleGroup with error: Operation validation and request error`, func() {
				testService, testServiceErr := wafrulegroupsapiv1.NewWafRuleGroupsApiV1(&wafrulegroupsapiv1.WafRuleGroupsApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the UpdateWafRuleGroupOptions model
				updateWafRuleGroupOptionsModel := new(wafrulegroupsapiv1.UpdateWafRuleGroupOptions)
				updateWafRuleGroupOptionsModel.PkgID = core.StringPtr("testString")
				updateWafRuleGroupOptionsModel.GroupID = core.StringPtr("testString")
				updateWafRuleGroupOptionsModel.Mode = core.StringPtr("on")
				updateWafRuleGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.UpdateWafRuleGroup(updateWafRuleGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateWafRuleGroupOptions model with no property values
				updateWafRuleGroupOptionsModelNew := new(wafrulegroupsapiv1.UpdateWafRuleGroupOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.UpdateWafRuleGroup(updateWafRuleGroupOptionsModelNew)
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
			testService, _ := wafrulegroupsapiv1.NewWafRuleGroupsApiV1(&wafrulegroupsapiv1.WafRuleGroupsApiV1Options{
				URL:           "http://wafrulegroupsapiv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
				Crn:           core.StringPtr(crn),
				ZoneID:        core.StringPtr(zoneID),
			})
			It(`Invoke NewGetWafRuleGroupOptions successfully`, func() {
				// Construct an instance of the GetWafRuleGroupOptions model
				pkgID := "testString"
				groupID := "testString"
				getWafRuleGroupOptionsModel := testService.NewGetWafRuleGroupOptions(pkgID, groupID)
				getWafRuleGroupOptionsModel.SetPkgID("testString")
				getWafRuleGroupOptionsModel.SetGroupID("testString")
				getWafRuleGroupOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getWafRuleGroupOptionsModel).ToNot(BeNil())
				Expect(getWafRuleGroupOptionsModel.PkgID).To(Equal(core.StringPtr("testString")))
				Expect(getWafRuleGroupOptionsModel.GroupID).To(Equal(core.StringPtr("testString")))
				Expect(getWafRuleGroupOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListWafRuleGroupsOptions successfully`, func() {
				// Construct an instance of the ListWafRuleGroupsOptions model
				pkgID := "testString"
				listWafRuleGroupsOptionsModel := testService.NewListWafRuleGroupsOptions(pkgID)
				listWafRuleGroupsOptionsModel.SetPkgID("testString")
				listWafRuleGroupsOptionsModel.SetName("Wordpress rules")
				listWafRuleGroupsOptionsModel.SetMode("true")
				listWafRuleGroupsOptionsModel.SetRulesCount("10")
				listWafRuleGroupsOptionsModel.SetPage(int64(1))
				listWafRuleGroupsOptionsModel.SetPerPage(int64(50))
				listWafRuleGroupsOptionsModel.SetOrder("status")
				listWafRuleGroupsOptionsModel.SetDirection("desc")
				listWafRuleGroupsOptionsModel.SetMatch("all")
				listWafRuleGroupsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listWafRuleGroupsOptionsModel).ToNot(BeNil())
				Expect(listWafRuleGroupsOptionsModel.PkgID).To(Equal(core.StringPtr("testString")))
				Expect(listWafRuleGroupsOptionsModel.Name).To(Equal(core.StringPtr("Wordpress rules")))
				Expect(listWafRuleGroupsOptionsModel.Mode).To(Equal(core.StringPtr("true")))
				Expect(listWafRuleGroupsOptionsModel.RulesCount).To(Equal(core.StringPtr("10")))
				Expect(listWafRuleGroupsOptionsModel.Page).To(Equal(core.Int64Ptr(int64(1))))
				Expect(listWafRuleGroupsOptionsModel.PerPage).To(Equal(core.Int64Ptr(int64(50))))
				Expect(listWafRuleGroupsOptionsModel.Order).To(Equal(core.StringPtr("status")))
				Expect(listWafRuleGroupsOptionsModel.Direction).To(Equal(core.StringPtr("desc")))
				Expect(listWafRuleGroupsOptionsModel.Match).To(Equal(core.StringPtr("all")))
				Expect(listWafRuleGroupsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateWafRuleGroupOptions successfully`, func() {
				// Construct an instance of the UpdateWafRuleGroupOptions model
				pkgID := "testString"
				groupID := "testString"
				updateWafRuleGroupOptionsModel := testService.NewUpdateWafRuleGroupOptions(pkgID, groupID)
				updateWafRuleGroupOptionsModel.SetPkgID("testString")
				updateWafRuleGroupOptionsModel.SetGroupID("testString")
				updateWafRuleGroupOptionsModel.SetMode("on")
				updateWafRuleGroupOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateWafRuleGroupOptionsModel).ToNot(BeNil())
				Expect(updateWafRuleGroupOptionsModel.PkgID).To(Equal(core.StringPtr("testString")))
				Expect(updateWafRuleGroupOptionsModel.GroupID).To(Equal(core.StringPtr("testString")))
				Expect(updateWafRuleGroupOptionsModel.Mode).To(Equal(core.StringPtr("on")))
				Expect(updateWafRuleGroupOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
