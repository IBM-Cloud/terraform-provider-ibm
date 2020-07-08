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

package pageruleapiv1_test

import (
	"bytes"
	"fmt"
	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.ibm.com/ibmcloud/networking-go-sdk/pageruleapiv1"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"time"
)

var _ = Describe(`PageRuleApiV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		crn := "testString"
		zoneID := "testString"
		It(`Instantiate service client`, func() {
			testService, testServiceErr := pageruleapiv1.NewPageRuleApiV1(&pageruleapiv1.PageRuleApiV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
				Crn:           core.StringPtr(crn),
				ZoneID:        core.StringPtr(zoneID),
			})
			Expect(testService).ToNot(BeNil())
			Expect(testServiceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			testService, testServiceErr := pageruleapiv1.NewPageRuleApiV1(&pageruleapiv1.PageRuleApiV1Options{
				URL:    "{BAD_URL_STRING",
				Crn:    core.StringPtr(crn),
				ZoneID: core.StringPtr(zoneID),
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			testService, testServiceErr := pageruleapiv1.NewPageRuleApiV1(&pageruleapiv1.PageRuleApiV1Options{
				URL:    "https://pageruleapiv1/api",
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
			testService, testServiceErr := pageruleapiv1.NewPageRuleApiV1(&pageruleapiv1.PageRuleApiV1Options{})
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
				"PAGE_RULE_API_URL":       "https://pageruleapiv1/api",
				"PAGE_RULE_API_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := pageruleapiv1.NewPageRuleApiV1UsingExternalConfig(&pageruleapiv1.PageRuleApiV1Options{
					Crn:    core.StringPtr(crn),
					ZoneID: core.StringPtr(zoneID),
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := pageruleapiv1.NewPageRuleApiV1UsingExternalConfig(&pageruleapiv1.PageRuleApiV1Options{
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
				testService, testServiceErr := pageruleapiv1.NewPageRuleApiV1UsingExternalConfig(&pageruleapiv1.PageRuleApiV1Options{
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
				"PAGE_RULE_API_URL":       "https://pageruleapiv1/api",
				"PAGE_RULE_API_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := pageruleapiv1.NewPageRuleApiV1UsingExternalConfig(&pageruleapiv1.PageRuleApiV1Options{
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
				"PAGE_RULE_API_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := pageruleapiv1.NewPageRuleApiV1UsingExternalConfig(&pageruleapiv1.PageRuleApiV1Options{
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
	Describe(`GetPageRule(getPageRuleOptions *GetPageRuleOptions) - Operation response error`, func() {
		crn := "testString"
		zoneID := "testString"
		getPageRulePath := "/v1/testString/zones/testString/pagerules/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getPageRulePath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetPageRule with error: Operation response processing error`, func() {
				testService, testServiceErr := pageruleapiv1.NewPageRuleApiV1(&pageruleapiv1.PageRuleApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetPageRuleOptions model
				getPageRuleOptionsModel := new(pageruleapiv1.GetPageRuleOptions)
				getPageRuleOptionsModel.RuleID = core.StringPtr("testString")
				getPageRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.GetPageRule(getPageRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetPageRule(getPageRuleOptions *GetPageRuleOptions)`, func() {
		crn := "testString"
		zoneID := "testString"
		getPageRulePath := "/v1/testString/zones/testString/pagerules/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getPageRulePath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "9a7806061c88ada191ed06f989cc3dac", "targets": [{"target": "url", "constraint": {"operator": "matches", "value": "*example.com/images/*"}}], "actions": [{"value": {"anyKey": "anyValue"}, "id": "disable_security"}], "priority": 1, "status": "active", "modified_on": "2014-01-01T05:20:00.12345Z", "created_on": "2014-01-01T05:20:00.12345Z"}}`)
				}))
			})
			It(`Invoke GetPageRule successfully`, func() {
				testService, testServiceErr := pageruleapiv1.NewPageRuleApiV1(&pageruleapiv1.PageRuleApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetPageRule(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetPageRuleOptions model
				getPageRuleOptionsModel := new(pageruleapiv1.GetPageRuleOptions)
				getPageRuleOptionsModel.RuleID = core.StringPtr("testString")
				getPageRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetPageRule(getPageRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetPageRule with error: Operation validation and request error`, func() {
				testService, testServiceErr := pageruleapiv1.NewPageRuleApiV1(&pageruleapiv1.PageRuleApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetPageRuleOptions model
				getPageRuleOptionsModel := new(pageruleapiv1.GetPageRuleOptions)
				getPageRuleOptionsModel.RuleID = core.StringPtr("testString")
				getPageRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.GetPageRule(getPageRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetPageRuleOptions model with no property values
				getPageRuleOptionsModelNew := new(pageruleapiv1.GetPageRuleOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.GetPageRule(getPageRuleOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ChangePageRule(changePageRuleOptions *ChangePageRuleOptions) - Operation response error`, func() {
		crn := "testString"
		zoneID := "testString"
		changePageRulePath := "/v1/testString/zones/testString/pagerules/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(changePageRulePath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ChangePageRule with error: Operation response processing error`, func() {
				testService, testServiceErr := pageruleapiv1.NewPageRuleApiV1(&pageruleapiv1.PageRuleApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the TargetsItemConstraint model
				targetsItemConstraintModel := new(pageruleapiv1.TargetsItemConstraint)
				targetsItemConstraintModel.Operator = core.StringPtr("matches")
				targetsItemConstraintModel.Value = core.StringPtr("*example.com/images/*")

				// Construct an instance of the PageRulesBodyActionsItemActionsSecurity model
				pageRulesBodyActionsItemModel := new(pageruleapiv1.PageRulesBodyActionsItemActionsSecurity)
				pageRulesBodyActionsItemModel.Value = map[string]interface{}{"anyKey": "anyValue"}
				pageRulesBodyActionsItemModel.ID = core.StringPtr("disable_security")

				// Construct an instance of the TargetsItem model
				targetsItemModel := new(pageruleapiv1.TargetsItem)
				targetsItemModel.Target = core.StringPtr("url")
				targetsItemModel.Constraint = targetsItemConstraintModel

				// Construct an instance of the ChangePageRuleOptions model
				changePageRuleOptionsModel := new(pageruleapiv1.ChangePageRuleOptions)
				changePageRuleOptionsModel.RuleID = core.StringPtr("testString")
				changePageRuleOptionsModel.Targets = []pageruleapiv1.TargetsItem{*targetsItemModel}
				changePageRuleOptionsModel.Actions = []pageruleapiv1.PageRulesBodyActionsItemIntf{pageRulesBodyActionsItemModel}
				changePageRuleOptionsModel.Priority = core.Int64Ptr(int64(1))
				changePageRuleOptionsModel.Status = core.StringPtr("active")
				changePageRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.ChangePageRule(changePageRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ChangePageRule(changePageRuleOptions *ChangePageRuleOptions)`, func() {
		crn := "testString"
		zoneID := "testString"
		changePageRulePath := "/v1/testString/zones/testString/pagerules/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(changePageRulePath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "9a7806061c88ada191ed06f989cc3dac", "targets": [{"target": "url", "constraint": {"operator": "matches", "value": "*example.com/images/*"}}], "actions": [{"value": {"anyKey": "anyValue"}, "id": "disable_security"}], "priority": 1, "status": "active", "modified_on": "2014-01-01T05:20:00.12345Z", "created_on": "2014-01-01T05:20:00.12345Z"}}`)
				}))
			})
			It(`Invoke ChangePageRule successfully`, func() {
				testService, testServiceErr := pageruleapiv1.NewPageRuleApiV1(&pageruleapiv1.PageRuleApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ChangePageRule(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the TargetsItemConstraint model
				targetsItemConstraintModel := new(pageruleapiv1.TargetsItemConstraint)
				targetsItemConstraintModel.Operator = core.StringPtr("matches")
				targetsItemConstraintModel.Value = core.StringPtr("*example.com/images/*")

				// Construct an instance of the PageRulesBodyActionsItemActionsSecurity model
				pageRulesBodyActionsItemModel := new(pageruleapiv1.PageRulesBodyActionsItemActionsSecurity)
				pageRulesBodyActionsItemModel.Value = map[string]interface{}{"anyKey": "anyValue"}
				pageRulesBodyActionsItemModel.ID = core.StringPtr("disable_security")

				// Construct an instance of the TargetsItem model
				targetsItemModel := new(pageruleapiv1.TargetsItem)
				targetsItemModel.Target = core.StringPtr("url")
				targetsItemModel.Constraint = targetsItemConstraintModel

				// Construct an instance of the ChangePageRuleOptions model
				changePageRuleOptionsModel := new(pageruleapiv1.ChangePageRuleOptions)
				changePageRuleOptionsModel.RuleID = core.StringPtr("testString")
				changePageRuleOptionsModel.Targets = []pageruleapiv1.TargetsItem{*targetsItemModel}
				changePageRuleOptionsModel.Actions = []pageruleapiv1.PageRulesBodyActionsItemIntf{pageRulesBodyActionsItemModel}
				changePageRuleOptionsModel.Priority = core.Int64Ptr(int64(1))
				changePageRuleOptionsModel.Status = core.StringPtr("active")
				changePageRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ChangePageRule(changePageRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke ChangePageRule with error: Operation validation and request error`, func() {
				testService, testServiceErr := pageruleapiv1.NewPageRuleApiV1(&pageruleapiv1.PageRuleApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the TargetsItemConstraint model
				targetsItemConstraintModel := new(pageruleapiv1.TargetsItemConstraint)
				targetsItemConstraintModel.Operator = core.StringPtr("matches")
				targetsItemConstraintModel.Value = core.StringPtr("*example.com/images/*")

				// Construct an instance of the PageRulesBodyActionsItemActionsSecurity model
				pageRulesBodyActionsItemModel := new(pageruleapiv1.PageRulesBodyActionsItemActionsSecurity)
				pageRulesBodyActionsItemModel.Value = map[string]interface{}{"anyKey": "anyValue"}
				pageRulesBodyActionsItemModel.ID = core.StringPtr("disable_security")

				// Construct an instance of the TargetsItem model
				targetsItemModel := new(pageruleapiv1.TargetsItem)
				targetsItemModel.Target = core.StringPtr("url")
				targetsItemModel.Constraint = targetsItemConstraintModel

				// Construct an instance of the ChangePageRuleOptions model
				changePageRuleOptionsModel := new(pageruleapiv1.ChangePageRuleOptions)
				changePageRuleOptionsModel.RuleID = core.StringPtr("testString")
				changePageRuleOptionsModel.Targets = []pageruleapiv1.TargetsItem{*targetsItemModel}
				changePageRuleOptionsModel.Actions = []pageruleapiv1.PageRulesBodyActionsItemIntf{pageRulesBodyActionsItemModel}
				changePageRuleOptionsModel.Priority = core.Int64Ptr(int64(1))
				changePageRuleOptionsModel.Status = core.StringPtr("active")
				changePageRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.ChangePageRule(changePageRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ChangePageRuleOptions model with no property values
				changePageRuleOptionsModelNew := new(pageruleapiv1.ChangePageRuleOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.ChangePageRule(changePageRuleOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdatePageRule(updatePageRuleOptions *UpdatePageRuleOptions) - Operation response error`, func() {
		crn := "testString"
		zoneID := "testString"
		updatePageRulePath := "/v1/testString/zones/testString/pagerules/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(updatePageRulePath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdatePageRule with error: Operation response processing error`, func() {
				testService, testServiceErr := pageruleapiv1.NewPageRuleApiV1(&pageruleapiv1.PageRuleApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the TargetsItemConstraint model
				targetsItemConstraintModel := new(pageruleapiv1.TargetsItemConstraint)
				targetsItemConstraintModel.Operator = core.StringPtr("matches")
				targetsItemConstraintModel.Value = core.StringPtr("*example.com/images/*")

				// Construct an instance of the PageRulesBodyActionsItemActionsSecurity model
				pageRulesBodyActionsItemModel := new(pageruleapiv1.PageRulesBodyActionsItemActionsSecurity)
				pageRulesBodyActionsItemModel.Value = map[string]interface{}{"anyKey": "anyValue"}
				pageRulesBodyActionsItemModel.ID = core.StringPtr("disable_security")

				// Construct an instance of the TargetsItem model
				targetsItemModel := new(pageruleapiv1.TargetsItem)
				targetsItemModel.Target = core.StringPtr("url")
				targetsItemModel.Constraint = targetsItemConstraintModel

				// Construct an instance of the UpdatePageRuleOptions model
				updatePageRuleOptionsModel := new(pageruleapiv1.UpdatePageRuleOptions)
				updatePageRuleOptionsModel.RuleID = core.StringPtr("testString")
				updatePageRuleOptionsModel.Targets = []pageruleapiv1.TargetsItem{*targetsItemModel}
				updatePageRuleOptionsModel.Actions = []pageruleapiv1.PageRulesBodyActionsItemIntf{pageRulesBodyActionsItemModel}
				updatePageRuleOptionsModel.Priority = core.Int64Ptr(int64(1))
				updatePageRuleOptionsModel.Status = core.StringPtr("active")
				updatePageRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.UpdatePageRule(updatePageRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdatePageRule(updatePageRuleOptions *UpdatePageRuleOptions)`, func() {
		crn := "testString"
		zoneID := "testString"
		updatePageRulePath := "/v1/testString/zones/testString/pagerules/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(updatePageRulePath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "9a7806061c88ada191ed06f989cc3dac", "targets": [{"target": "url", "constraint": {"operator": "matches", "value": "*example.com/images/*"}}], "actions": [{"value": {"anyKey": "anyValue"}, "id": "disable_security"}], "priority": 1, "status": "active", "modified_on": "2014-01-01T05:20:00.12345Z", "created_on": "2014-01-01T05:20:00.12345Z"}}`)
				}))
			})
			It(`Invoke UpdatePageRule successfully`, func() {
				testService, testServiceErr := pageruleapiv1.NewPageRuleApiV1(&pageruleapiv1.PageRuleApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.UpdatePageRule(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the TargetsItemConstraint model
				targetsItemConstraintModel := new(pageruleapiv1.TargetsItemConstraint)
				targetsItemConstraintModel.Operator = core.StringPtr("matches")
				targetsItemConstraintModel.Value = core.StringPtr("*example.com/images/*")

				// Construct an instance of the PageRulesBodyActionsItemActionsSecurity model
				pageRulesBodyActionsItemModel := new(pageruleapiv1.PageRulesBodyActionsItemActionsSecurity)
				pageRulesBodyActionsItemModel.Value = map[string]interface{}{"anyKey": "anyValue"}
				pageRulesBodyActionsItemModel.ID = core.StringPtr("disable_security")

				// Construct an instance of the TargetsItem model
				targetsItemModel := new(pageruleapiv1.TargetsItem)
				targetsItemModel.Target = core.StringPtr("url")
				targetsItemModel.Constraint = targetsItemConstraintModel

				// Construct an instance of the UpdatePageRuleOptions model
				updatePageRuleOptionsModel := new(pageruleapiv1.UpdatePageRuleOptions)
				updatePageRuleOptionsModel.RuleID = core.StringPtr("testString")
				updatePageRuleOptionsModel.Targets = []pageruleapiv1.TargetsItem{*targetsItemModel}
				updatePageRuleOptionsModel.Actions = []pageruleapiv1.PageRulesBodyActionsItemIntf{pageRulesBodyActionsItemModel}
				updatePageRuleOptionsModel.Priority = core.Int64Ptr(int64(1))
				updatePageRuleOptionsModel.Status = core.StringPtr("active")
				updatePageRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.UpdatePageRule(updatePageRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke UpdatePageRule with error: Operation validation and request error`, func() {
				testService, testServiceErr := pageruleapiv1.NewPageRuleApiV1(&pageruleapiv1.PageRuleApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the TargetsItemConstraint model
				targetsItemConstraintModel := new(pageruleapiv1.TargetsItemConstraint)
				targetsItemConstraintModel.Operator = core.StringPtr("matches")
				targetsItemConstraintModel.Value = core.StringPtr("*example.com/images/*")

				// Construct an instance of the PageRulesBodyActionsItemActionsSecurity model
				pageRulesBodyActionsItemModel := new(pageruleapiv1.PageRulesBodyActionsItemActionsSecurity)
				pageRulesBodyActionsItemModel.Value = map[string]interface{}{"anyKey": "anyValue"}
				pageRulesBodyActionsItemModel.ID = core.StringPtr("disable_security")

				// Construct an instance of the TargetsItem model
				targetsItemModel := new(pageruleapiv1.TargetsItem)
				targetsItemModel.Target = core.StringPtr("url")
				targetsItemModel.Constraint = targetsItemConstraintModel

				// Construct an instance of the UpdatePageRuleOptions model
				updatePageRuleOptionsModel := new(pageruleapiv1.UpdatePageRuleOptions)
				updatePageRuleOptionsModel.RuleID = core.StringPtr("testString")
				updatePageRuleOptionsModel.Targets = []pageruleapiv1.TargetsItem{*targetsItemModel}
				updatePageRuleOptionsModel.Actions = []pageruleapiv1.PageRulesBodyActionsItemIntf{pageRulesBodyActionsItemModel}
				updatePageRuleOptionsModel.Priority = core.Int64Ptr(int64(1))
				updatePageRuleOptionsModel.Status = core.StringPtr("active")
				updatePageRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.UpdatePageRule(updatePageRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdatePageRuleOptions model with no property values
				updatePageRuleOptionsModelNew := new(pageruleapiv1.UpdatePageRuleOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.UpdatePageRule(updatePageRuleOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeletePageRule(deletePageRuleOptions *DeletePageRuleOptions) - Operation response error`, func() {
		crn := "testString"
		zoneID := "testString"
		deletePageRulePath := "/v1/testString/zones/testString/pagerules/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(deletePageRulePath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeletePageRule with error: Operation response processing error`, func() {
				testService, testServiceErr := pageruleapiv1.NewPageRuleApiV1(&pageruleapiv1.PageRuleApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the DeletePageRuleOptions model
				deletePageRuleOptionsModel := new(pageruleapiv1.DeletePageRuleOptions)
				deletePageRuleOptionsModel.RuleID = core.StringPtr("testString")
				deletePageRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.DeletePageRule(deletePageRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeletePageRule(deletePageRuleOptions *DeletePageRuleOptions)`, func() {
		crn := "testString"
		zoneID := "testString"
		deletePageRulePath := "/v1/testString/zones/testString/pagerules/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(deletePageRulePath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "9a7806061c88ada191ed06f989cc3dac"}}`)
				}))
			})
			It(`Invoke DeletePageRule successfully`, func() {
				testService, testServiceErr := pageruleapiv1.NewPageRuleApiV1(&pageruleapiv1.PageRuleApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.DeletePageRule(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeletePageRuleOptions model
				deletePageRuleOptionsModel := new(pageruleapiv1.DeletePageRuleOptions)
				deletePageRuleOptionsModel.RuleID = core.StringPtr("testString")
				deletePageRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.DeletePageRule(deletePageRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke DeletePageRule with error: Operation validation and request error`, func() {
				testService, testServiceErr := pageruleapiv1.NewPageRuleApiV1(&pageruleapiv1.PageRuleApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the DeletePageRuleOptions model
				deletePageRuleOptionsModel := new(pageruleapiv1.DeletePageRuleOptions)
				deletePageRuleOptionsModel.RuleID = core.StringPtr("testString")
				deletePageRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.DeletePageRule(deletePageRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeletePageRuleOptions model with no property values
				deletePageRuleOptionsModelNew := new(pageruleapiv1.DeletePageRuleOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.DeletePageRule(deletePageRuleOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListPageRules(listPageRulesOptions *ListPageRulesOptions) - Operation response error`, func() {
		crn := "testString"
		zoneID := "testString"
		listPageRulesPath := "/v1/testString/zones/testString/pagerules"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listPageRulesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["status"]).To(Equal([]string{"active"}))

					Expect(req.URL.Query()["order"]).To(Equal([]string{"status"}))

					Expect(req.URL.Query()["direction"]).To(Equal([]string{"desc"}))

					Expect(req.URL.Query()["match"]).To(Equal([]string{"all"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListPageRules with error: Operation response processing error`, func() {
				testService, testServiceErr := pageruleapiv1.NewPageRuleApiV1(&pageruleapiv1.PageRuleApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ListPageRulesOptions model
				listPageRulesOptionsModel := new(pageruleapiv1.ListPageRulesOptions)
				listPageRulesOptionsModel.Status = core.StringPtr("active")
				listPageRulesOptionsModel.Order = core.StringPtr("status")
				listPageRulesOptionsModel.Direction = core.StringPtr("desc")
				listPageRulesOptionsModel.Match = core.StringPtr("all")
				listPageRulesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.ListPageRules(listPageRulesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListPageRules(listPageRulesOptions *ListPageRulesOptions)`, func() {
		crn := "testString"
		zoneID := "testString"
		listPageRulesPath := "/v1/testString/zones/testString/pagerules"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listPageRulesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["status"]).To(Equal([]string{"active"}))

					Expect(req.URL.Query()["order"]).To(Equal([]string{"status"}))

					Expect(req.URL.Query()["direction"]).To(Equal([]string{"desc"}))

					Expect(req.URL.Query()["match"]).To(Equal([]string{"all"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": [{"id": "9a7806061c88ada191ed06f989cc3dac", "targets": [{"target": "url", "constraint": {"operator": "matches", "value": "*example.com/images/*"}}], "actions": [{"value": {"anyKey": "anyValue"}, "id": "disable_security"}], "priority": 1, "status": "active", "modified_on": "2014-01-01T05:20:00.12345Z", "created_on": "2014-01-01T05:20:00.12345Z"}]}`)
				}))
			})
			It(`Invoke ListPageRules successfully`, func() {
				testService, testServiceErr := pageruleapiv1.NewPageRuleApiV1(&pageruleapiv1.PageRuleApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListPageRules(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListPageRulesOptions model
				listPageRulesOptionsModel := new(pageruleapiv1.ListPageRulesOptions)
				listPageRulesOptionsModel.Status = core.StringPtr("active")
				listPageRulesOptionsModel.Order = core.StringPtr("status")
				listPageRulesOptionsModel.Direction = core.StringPtr("desc")
				listPageRulesOptionsModel.Match = core.StringPtr("all")
				listPageRulesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListPageRules(listPageRulesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke ListPageRules with error: Operation request error`, func() {
				testService, testServiceErr := pageruleapiv1.NewPageRuleApiV1(&pageruleapiv1.PageRuleApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ListPageRulesOptions model
				listPageRulesOptionsModel := new(pageruleapiv1.ListPageRulesOptions)
				listPageRulesOptionsModel.Status = core.StringPtr("active")
				listPageRulesOptionsModel.Order = core.StringPtr("status")
				listPageRulesOptionsModel.Direction = core.StringPtr("desc")
				listPageRulesOptionsModel.Match = core.StringPtr("all")
				listPageRulesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.ListPageRules(listPageRulesOptionsModel)
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
	Describe(`CreatePageRule(createPageRuleOptions *CreatePageRuleOptions) - Operation response error`, func() {
		crn := "testString"
		zoneID := "testString"
		createPageRulePath := "/v1/testString/zones/testString/pagerules"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(createPageRulePath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreatePageRule with error: Operation response processing error`, func() {
				testService, testServiceErr := pageruleapiv1.NewPageRuleApiV1(&pageruleapiv1.PageRuleApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the TargetsItemConstraint model
				targetsItemConstraintModel := new(pageruleapiv1.TargetsItemConstraint)
				targetsItemConstraintModel.Operator = core.StringPtr("matches")
				targetsItemConstraintModel.Value = core.StringPtr("*example.com/images/*")

				// Construct an instance of the PageRulesBodyActionsItemActionsSecurity model
				pageRulesBodyActionsItemModel := new(pageruleapiv1.PageRulesBodyActionsItemActionsSecurity)
				pageRulesBodyActionsItemModel.Value = map[string]interface{}{"anyKey": "anyValue"}
				pageRulesBodyActionsItemModel.ID = core.StringPtr("disable_security")

				// Construct an instance of the TargetsItem model
				targetsItemModel := new(pageruleapiv1.TargetsItem)
				targetsItemModel.Target = core.StringPtr("url")
				targetsItemModel.Constraint = targetsItemConstraintModel

				// Construct an instance of the CreatePageRuleOptions model
				createPageRuleOptionsModel := new(pageruleapiv1.CreatePageRuleOptions)
				createPageRuleOptionsModel.Targets = []pageruleapiv1.TargetsItem{*targetsItemModel}
				createPageRuleOptionsModel.Actions = []pageruleapiv1.PageRulesBodyActionsItemIntf{pageRulesBodyActionsItemModel}
				createPageRuleOptionsModel.Priority = core.Int64Ptr(int64(1))
				createPageRuleOptionsModel.Status = core.StringPtr("active")
				createPageRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.CreatePageRule(createPageRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CreatePageRule(createPageRuleOptions *CreatePageRuleOptions)`, func() {
		crn := "testString"
		zoneID := "testString"
		createPageRulePath := "/v1/testString/zones/testString/pagerules"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(createPageRulePath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "9a7806061c88ada191ed06f989cc3dac", "targets": [{"target": "url", "constraint": {"operator": "matches", "value": "*example.com/images/*"}}], "actions": [{"value": {"anyKey": "anyValue"}, "id": "disable_security"}], "priority": 1, "status": "active", "modified_on": "2014-01-01T05:20:00.12345Z", "created_on": "2014-01-01T05:20:00.12345Z"}}`)
				}))
			})
			It(`Invoke CreatePageRule successfully`, func() {
				testService, testServiceErr := pageruleapiv1.NewPageRuleApiV1(&pageruleapiv1.PageRuleApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.CreatePageRule(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the TargetsItemConstraint model
				targetsItemConstraintModel := new(pageruleapiv1.TargetsItemConstraint)
				targetsItemConstraintModel.Operator = core.StringPtr("matches")
				targetsItemConstraintModel.Value = core.StringPtr("*example.com/images/*")

				// Construct an instance of the PageRulesBodyActionsItemActionsSecurity model
				pageRulesBodyActionsItemModel := new(pageruleapiv1.PageRulesBodyActionsItemActionsSecurity)
				pageRulesBodyActionsItemModel.Value = map[string]interface{}{"anyKey": "anyValue"}
				pageRulesBodyActionsItemModel.ID = core.StringPtr("disable_security")

				// Construct an instance of the TargetsItem model
				targetsItemModel := new(pageruleapiv1.TargetsItem)
				targetsItemModel.Target = core.StringPtr("url")
				targetsItemModel.Constraint = targetsItemConstraintModel

				// Construct an instance of the CreatePageRuleOptions model
				createPageRuleOptionsModel := new(pageruleapiv1.CreatePageRuleOptions)
				createPageRuleOptionsModel.Targets = []pageruleapiv1.TargetsItem{*targetsItemModel}
				createPageRuleOptionsModel.Actions = []pageruleapiv1.PageRulesBodyActionsItemIntf{pageRulesBodyActionsItemModel}
				createPageRuleOptionsModel.Priority = core.Int64Ptr(int64(1))
				createPageRuleOptionsModel.Status = core.StringPtr("active")
				createPageRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.CreatePageRule(createPageRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke CreatePageRule with error: Operation request error`, func() {
				testService, testServiceErr := pageruleapiv1.NewPageRuleApiV1(&pageruleapiv1.PageRuleApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the TargetsItemConstraint model
				targetsItemConstraintModel := new(pageruleapiv1.TargetsItemConstraint)
				targetsItemConstraintModel.Operator = core.StringPtr("matches")
				targetsItemConstraintModel.Value = core.StringPtr("*example.com/images/*")

				// Construct an instance of the PageRulesBodyActionsItemActionsSecurity model
				pageRulesBodyActionsItemModel := new(pageruleapiv1.PageRulesBodyActionsItemActionsSecurity)
				pageRulesBodyActionsItemModel.Value = map[string]interface{}{"anyKey": "anyValue"}
				pageRulesBodyActionsItemModel.ID = core.StringPtr("disable_security")

				// Construct an instance of the TargetsItem model
				targetsItemModel := new(pageruleapiv1.TargetsItem)
				targetsItemModel.Target = core.StringPtr("url")
				targetsItemModel.Constraint = targetsItemConstraintModel

				// Construct an instance of the CreatePageRuleOptions model
				createPageRuleOptionsModel := new(pageruleapiv1.CreatePageRuleOptions)
				createPageRuleOptionsModel.Targets = []pageruleapiv1.TargetsItem{*targetsItemModel}
				createPageRuleOptionsModel.Actions = []pageruleapiv1.PageRulesBodyActionsItemIntf{pageRulesBodyActionsItemModel}
				createPageRuleOptionsModel.Priority = core.Int64Ptr(int64(1))
				createPageRuleOptionsModel.Status = core.StringPtr("active")
				createPageRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.CreatePageRule(createPageRuleOptionsModel)
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
	Describe(`Model constructor tests`, func() {
		Context(`Using a service client instance`, func() {
			crn := "testString"
			zoneID := "testString"
			testService, _ := pageruleapiv1.NewPageRuleApiV1(&pageruleapiv1.PageRuleApiV1Options{
				URL:           "http://pageruleapiv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
				Crn:           core.StringPtr(crn),
				ZoneID:        core.StringPtr(zoneID),
			})
			It(`Invoke NewChangePageRuleOptions successfully`, func() {
				// Construct an instance of the TargetsItemConstraint model
				targetsItemConstraintModel := new(pageruleapiv1.TargetsItemConstraint)
				Expect(targetsItemConstraintModel).ToNot(BeNil())
				targetsItemConstraintModel.Operator = core.StringPtr("matches")
				targetsItemConstraintModel.Value = core.StringPtr("*example.com/images/*")
				Expect(targetsItemConstraintModel.Operator).To(Equal(core.StringPtr("matches")))
				Expect(targetsItemConstraintModel.Value).To(Equal(core.StringPtr("*example.com/images/*")))

				// Construct an instance of the PageRulesBodyActionsItemActionsSecurity model
				pageRulesBodyActionsItemModel := new(pageruleapiv1.PageRulesBodyActionsItemActionsSecurity)
				Expect(pageRulesBodyActionsItemModel).ToNot(BeNil())
				pageRulesBodyActionsItemModel.Value = map[string]interface{}{"anyKey": "anyValue"}
				pageRulesBodyActionsItemModel.ID = core.StringPtr("disable_security")
				Expect(pageRulesBodyActionsItemModel.Value).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(pageRulesBodyActionsItemModel.ID).To(Equal(core.StringPtr("disable_security")))

				// Construct an instance of the TargetsItem model
				targetsItemModel := new(pageruleapiv1.TargetsItem)
				Expect(targetsItemModel).ToNot(BeNil())
				targetsItemModel.Target = core.StringPtr("url")
				targetsItemModel.Constraint = targetsItemConstraintModel
				Expect(targetsItemModel.Target).To(Equal(core.StringPtr("url")))
				Expect(targetsItemModel.Constraint).To(Equal(targetsItemConstraintModel))

				// Construct an instance of the ChangePageRuleOptions model
				ruleID := "testString"
				changePageRuleOptionsModel := testService.NewChangePageRuleOptions(ruleID)
				changePageRuleOptionsModel.SetRuleID("testString")
				changePageRuleOptionsModel.SetTargets([]pageruleapiv1.TargetsItem{*targetsItemModel})
				changePageRuleOptionsModel.SetActions([]pageruleapiv1.PageRulesBodyActionsItemIntf{pageRulesBodyActionsItemModel})
				changePageRuleOptionsModel.SetPriority(int64(1))
				changePageRuleOptionsModel.SetStatus("active")
				changePageRuleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(changePageRuleOptionsModel).ToNot(BeNil())
				Expect(changePageRuleOptionsModel.RuleID).To(Equal(core.StringPtr("testString")))
				Expect(changePageRuleOptionsModel.Targets).To(Equal([]pageruleapiv1.TargetsItem{*targetsItemModel}))
				Expect(changePageRuleOptionsModel.Actions).To(Equal([]pageruleapiv1.PageRulesBodyActionsItemIntf{pageRulesBodyActionsItemModel}))
				Expect(changePageRuleOptionsModel.Priority).To(Equal(core.Int64Ptr(int64(1))))
				Expect(changePageRuleOptionsModel.Status).To(Equal(core.StringPtr("active")))
				Expect(changePageRuleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreatePageRuleOptions successfully`, func() {
				// Construct an instance of the TargetsItemConstraint model
				targetsItemConstraintModel := new(pageruleapiv1.TargetsItemConstraint)
				Expect(targetsItemConstraintModel).ToNot(BeNil())
				targetsItemConstraintModel.Operator = core.StringPtr("matches")
				targetsItemConstraintModel.Value = core.StringPtr("*example.com/images/*")
				Expect(targetsItemConstraintModel.Operator).To(Equal(core.StringPtr("matches")))
				Expect(targetsItemConstraintModel.Value).To(Equal(core.StringPtr("*example.com/images/*")))

				// Construct an instance of the PageRulesBodyActionsItemActionsSecurity model
				pageRulesBodyActionsItemModel := new(pageruleapiv1.PageRulesBodyActionsItemActionsSecurity)
				Expect(pageRulesBodyActionsItemModel).ToNot(BeNil())
				pageRulesBodyActionsItemModel.Value = map[string]interface{}{"anyKey": "anyValue"}
				pageRulesBodyActionsItemModel.ID = core.StringPtr("disable_security")
				Expect(pageRulesBodyActionsItemModel.Value).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(pageRulesBodyActionsItemModel.ID).To(Equal(core.StringPtr("disable_security")))

				// Construct an instance of the TargetsItem model
				targetsItemModel := new(pageruleapiv1.TargetsItem)
				Expect(targetsItemModel).ToNot(BeNil())
				targetsItemModel.Target = core.StringPtr("url")
				targetsItemModel.Constraint = targetsItemConstraintModel
				Expect(targetsItemModel.Target).To(Equal(core.StringPtr("url")))
				Expect(targetsItemModel.Constraint).To(Equal(targetsItemConstraintModel))

				// Construct an instance of the CreatePageRuleOptions model
				createPageRuleOptionsModel := testService.NewCreatePageRuleOptions()
				createPageRuleOptionsModel.SetTargets([]pageruleapiv1.TargetsItem{*targetsItemModel})
				createPageRuleOptionsModel.SetActions([]pageruleapiv1.PageRulesBodyActionsItemIntf{pageRulesBodyActionsItemModel})
				createPageRuleOptionsModel.SetPriority(int64(1))
				createPageRuleOptionsModel.SetStatus("active")
				createPageRuleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createPageRuleOptionsModel).ToNot(BeNil())
				Expect(createPageRuleOptionsModel.Targets).To(Equal([]pageruleapiv1.TargetsItem{*targetsItemModel}))
				Expect(createPageRuleOptionsModel.Actions).To(Equal([]pageruleapiv1.PageRulesBodyActionsItemIntf{pageRulesBodyActionsItemModel}))
				Expect(createPageRuleOptionsModel.Priority).To(Equal(core.Int64Ptr(int64(1))))
				Expect(createPageRuleOptionsModel.Status).To(Equal(core.StringPtr("active")))
				Expect(createPageRuleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeletePageRuleOptions successfully`, func() {
				// Construct an instance of the DeletePageRuleOptions model
				ruleID := "testString"
				deletePageRuleOptionsModel := testService.NewDeletePageRuleOptions(ruleID)
				deletePageRuleOptionsModel.SetRuleID("testString")
				deletePageRuleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deletePageRuleOptionsModel).ToNot(BeNil())
				Expect(deletePageRuleOptionsModel.RuleID).To(Equal(core.StringPtr("testString")))
				Expect(deletePageRuleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetPageRuleOptions successfully`, func() {
				// Construct an instance of the GetPageRuleOptions model
				ruleID := "testString"
				getPageRuleOptionsModel := testService.NewGetPageRuleOptions(ruleID)
				getPageRuleOptionsModel.SetRuleID("testString")
				getPageRuleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getPageRuleOptionsModel).ToNot(BeNil())
				Expect(getPageRuleOptionsModel.RuleID).To(Equal(core.StringPtr("testString")))
				Expect(getPageRuleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListPageRulesOptions successfully`, func() {
				// Construct an instance of the ListPageRulesOptions model
				listPageRulesOptionsModel := testService.NewListPageRulesOptions()
				listPageRulesOptionsModel.SetStatus("active")
				listPageRulesOptionsModel.SetOrder("status")
				listPageRulesOptionsModel.SetDirection("desc")
				listPageRulesOptionsModel.SetMatch("all")
				listPageRulesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listPageRulesOptionsModel).ToNot(BeNil())
				Expect(listPageRulesOptionsModel.Status).To(Equal(core.StringPtr("active")))
				Expect(listPageRulesOptionsModel.Order).To(Equal(core.StringPtr("status")))
				Expect(listPageRulesOptionsModel.Direction).To(Equal(core.StringPtr("desc")))
				Expect(listPageRulesOptionsModel.Match).To(Equal(core.StringPtr("all")))
				Expect(listPageRulesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewTargetsItem successfully`, func() {
				target := "url"
				var constraint *pageruleapiv1.TargetsItemConstraint = nil
				_, err := testService.NewTargetsItem(target, constraint)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewTargetsItemConstraint successfully`, func() {
				operator := "matches"
				value := "*example.com/images/*"
				model, err := testService.NewTargetsItemConstraint(operator, value)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewUpdatePageRuleOptions successfully`, func() {
				// Construct an instance of the TargetsItemConstraint model
				targetsItemConstraintModel := new(pageruleapiv1.TargetsItemConstraint)
				Expect(targetsItemConstraintModel).ToNot(BeNil())
				targetsItemConstraintModel.Operator = core.StringPtr("matches")
				targetsItemConstraintModel.Value = core.StringPtr("*example.com/images/*")
				Expect(targetsItemConstraintModel.Operator).To(Equal(core.StringPtr("matches")))
				Expect(targetsItemConstraintModel.Value).To(Equal(core.StringPtr("*example.com/images/*")))

				// Construct an instance of the PageRulesBodyActionsItemActionsSecurity model
				pageRulesBodyActionsItemModel := new(pageruleapiv1.PageRulesBodyActionsItemActionsSecurity)
				Expect(pageRulesBodyActionsItemModel).ToNot(BeNil())
				pageRulesBodyActionsItemModel.Value = map[string]interface{}{"anyKey": "anyValue"}
				pageRulesBodyActionsItemModel.ID = core.StringPtr("disable_security")
				Expect(pageRulesBodyActionsItemModel.Value).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(pageRulesBodyActionsItemModel.ID).To(Equal(core.StringPtr("disable_security")))

				// Construct an instance of the TargetsItem model
				targetsItemModel := new(pageruleapiv1.TargetsItem)
				Expect(targetsItemModel).ToNot(BeNil())
				targetsItemModel.Target = core.StringPtr("url")
				targetsItemModel.Constraint = targetsItemConstraintModel
				Expect(targetsItemModel.Target).To(Equal(core.StringPtr("url")))
				Expect(targetsItemModel.Constraint).To(Equal(targetsItemConstraintModel))

				// Construct an instance of the UpdatePageRuleOptions model
				ruleID := "testString"
				updatePageRuleOptionsModel := testService.NewUpdatePageRuleOptions(ruleID)
				updatePageRuleOptionsModel.SetRuleID("testString")
				updatePageRuleOptionsModel.SetTargets([]pageruleapiv1.TargetsItem{*targetsItemModel})
				updatePageRuleOptionsModel.SetActions([]pageruleapiv1.PageRulesBodyActionsItemIntf{pageRulesBodyActionsItemModel})
				updatePageRuleOptionsModel.SetPriority(int64(1))
				updatePageRuleOptionsModel.SetStatus("active")
				updatePageRuleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updatePageRuleOptionsModel).ToNot(BeNil())
				Expect(updatePageRuleOptionsModel.RuleID).To(Equal(core.StringPtr("testString")))
				Expect(updatePageRuleOptionsModel.Targets).To(Equal([]pageruleapiv1.TargetsItem{*targetsItemModel}))
				Expect(updatePageRuleOptionsModel.Actions).To(Equal([]pageruleapiv1.PageRulesBodyActionsItemIntf{pageRulesBodyActionsItemModel}))
				Expect(updatePageRuleOptionsModel.Priority).To(Equal(core.Int64Ptr(int64(1))))
				Expect(updatePageRuleOptionsModel.Status).To(Equal(core.StringPtr("active")))
				Expect(updatePageRuleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPageRulesBodyActionsItemActionsBypassCacheOnCookie successfully`, func() {
				id := "bypass_cache_on_cookie"
				model, err := testService.NewPageRulesBodyActionsItemActionsBypassCacheOnCookie(id)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewPageRulesBodyActionsItemActionsCacheLevel successfully`, func() {
				id := "cache_level"
				model, err := testService.NewPageRulesBodyActionsItemActionsCacheLevel(id)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewPageRulesBodyActionsItemActionsEdgeCacheTTL successfully`, func() {
				id := "edge_cache_ttl"
				model, err := testService.NewPageRulesBodyActionsItemActionsEdgeCacheTTL(id)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewPageRulesBodyActionsItemActionsForwardingURL successfully`, func() {
				id := "forwarding_url"
				model, err := testService.NewPageRulesBodyActionsItemActionsForwardingURL(id)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewPageRulesBodyActionsItemActionsSecurity successfully`, func() {
				id := "disable_security"
				model, err := testService.NewPageRulesBodyActionsItemActionsSecurity(id)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewPageRulesBodyActionsItemActionsSecurityLevel successfully`, func() {
				id := "security_level"
				model, err := testService.NewPageRulesBodyActionsItemActionsSecurityLevel(id)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewPageRulesBodyActionsItemActionsSecurityOptions successfully`, func() {
				id := "browser_check"
				model, err := testService.NewPageRulesBodyActionsItemActionsSecurityOptions(id)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewPageRulesBodyActionsItemActionsSsl successfully`, func() {
				id := "ssl"
				model, err := testService.NewPageRulesBodyActionsItemActionsSsl(id)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewPageRulesBodyActionsItemActionsTTL successfully`, func() {
				id := "browser_cache_ttl"
				model, err := testService.NewPageRulesBodyActionsItemActionsTTL(id)
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
