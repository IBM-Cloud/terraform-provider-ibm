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

package useragentblockingrulesv1_test

import (
	"bytes"
	"fmt"
	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.ibm.com/ibmcloud/networking-go-sdk/useragentblockingrulesv1"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"time"
)

var _ = Describe(`UserAgentBlockingRulesV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		It(`Instantiate service client`, func() {
			testService, testServiceErr := useragentblockingrulesv1.NewUserAgentBlockingRulesV1(&useragentblockingrulesv1.UserAgentBlockingRulesV1Options{
				Authenticator:  &core.NoAuthAuthenticator{},
				Crn:            core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})
			Expect(testService).ToNot(BeNil())
			Expect(testServiceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			testService, testServiceErr := useragentblockingrulesv1.NewUserAgentBlockingRulesV1(&useragentblockingrulesv1.UserAgentBlockingRulesV1Options{
				URL:            "{BAD_URL_STRING",
				Crn:            core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			testService, testServiceErr := useragentblockingrulesv1.NewUserAgentBlockingRulesV1(&useragentblockingrulesv1.UserAgentBlockingRulesV1Options{
				URL:            "https://useragentblockingrulesv1/api",
				Crn:            core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Validation Error`, func() {
			testService, testServiceErr := useragentblockingrulesv1.NewUserAgentBlockingRulesV1(&useragentblockingrulesv1.UserAgentBlockingRulesV1Options{})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"USER_AGENT_BLOCKING_RULES_URL":       "https://useragentblockingrulesv1/api",
				"USER_AGENT_BLOCKING_RULES_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := useragentblockingrulesv1.NewUserAgentBlockingRulesV1UsingExternalConfig(&useragentblockingrulesv1.UserAgentBlockingRulesV1Options{
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := useragentblockingrulesv1.NewUserAgentBlockingRulesV1UsingExternalConfig(&useragentblockingrulesv1.UserAgentBlockingRulesV1Options{
					URL:            "https://testService/api",
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				Expect(testService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := useragentblockingrulesv1.NewUserAgentBlockingRulesV1UsingExternalConfig(&useragentblockingrulesv1.UserAgentBlockingRulesV1Options{
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
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
				"USER_AGENT_BLOCKING_RULES_URL":       "https://useragentblockingrulesv1/api",
				"USER_AGENT_BLOCKING_RULES_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := useragentblockingrulesv1.NewUserAgentBlockingRulesV1UsingExternalConfig(&useragentblockingrulesv1.UserAgentBlockingRulesV1Options{
				Crn:            core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
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
				"USER_AGENT_BLOCKING_RULES_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := useragentblockingrulesv1.NewUserAgentBlockingRulesV1UsingExternalConfig(&useragentblockingrulesv1.UserAgentBlockingRulesV1Options{
				URL:            "{BAD_URL_STRING",
				Crn:            core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})

			It(`Instantiate service client with error`, func() {
				Expect(testService).To(BeNil())
				Expect(testServiceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`ListAllZoneUserAgentRules(listAllZoneUserAgentRulesOptions *ListAllZoneUserAgentRulesOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		listAllZoneUserAgentRulesPath := "/v1/testString/zones/testString/firewall/ua_rules"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listAllZoneUserAgentRulesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["page"]).To(Equal([]string{fmt.Sprint(int64(38))}))

					Expect(req.URL.Query()["per_page"]).To(Equal([]string{fmt.Sprint(int64(5))}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListAllZoneUserAgentRules with error: Operation response processing error`, func() {
				testService, testServiceErr := useragentblockingrulesv1.NewUserAgentBlockingRulesV1(&useragentblockingrulesv1.UserAgentBlockingRulesV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ListAllZoneUserAgentRulesOptions model
				listAllZoneUserAgentRulesOptionsModel := new(useragentblockingrulesv1.ListAllZoneUserAgentRulesOptions)
				listAllZoneUserAgentRulesOptionsModel.Page = core.Int64Ptr(int64(38))
				listAllZoneUserAgentRulesOptionsModel.PerPage = core.Int64Ptr(int64(5))
				listAllZoneUserAgentRulesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.ListAllZoneUserAgentRules(listAllZoneUserAgentRulesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListAllZoneUserAgentRules(listAllZoneUserAgentRulesOptions *ListAllZoneUserAgentRulesOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		listAllZoneUserAgentRulesPath := "/v1/testString/zones/testString/firewall/ua_rules"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listAllZoneUserAgentRulesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["page"]).To(Equal([]string{fmt.Sprint(int64(38))}))

					Expect(req.URL.Query()["per_page"]).To(Equal([]string{fmt.Sprint(int64(5))}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": [{"id": "92f17202ed8bd63d69a66b86a49a8f6b", "paused": true, "description": "Prevent access from abusive clients identified by this UserAgent to mitigate DDoS attack", "mode": "block", "configuration": {"target": "ua", "value": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_5) AppleWebKit/603.2.4 (KHTML, like Gecko) Version/10.1.1 Safari/603.2.4"}}], "result_info": {"page": 1, "per_page": 2, "count": 1, "total_count": 200}}`)
				}))
			})
			It(`Invoke ListAllZoneUserAgentRules successfully`, func() {
				testService, testServiceErr := useragentblockingrulesv1.NewUserAgentBlockingRulesV1(&useragentblockingrulesv1.UserAgentBlockingRulesV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListAllZoneUserAgentRules(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListAllZoneUserAgentRulesOptions model
				listAllZoneUserAgentRulesOptionsModel := new(useragentblockingrulesv1.ListAllZoneUserAgentRulesOptions)
				listAllZoneUserAgentRulesOptionsModel.Page = core.Int64Ptr(int64(38))
				listAllZoneUserAgentRulesOptionsModel.PerPage = core.Int64Ptr(int64(5))
				listAllZoneUserAgentRulesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListAllZoneUserAgentRules(listAllZoneUserAgentRulesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke ListAllZoneUserAgentRules with error: Operation request error`, func() {
				testService, testServiceErr := useragentblockingrulesv1.NewUserAgentBlockingRulesV1(&useragentblockingrulesv1.UserAgentBlockingRulesV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ListAllZoneUserAgentRulesOptions model
				listAllZoneUserAgentRulesOptionsModel := new(useragentblockingrulesv1.ListAllZoneUserAgentRulesOptions)
				listAllZoneUserAgentRulesOptionsModel.Page = core.Int64Ptr(int64(38))
				listAllZoneUserAgentRulesOptionsModel.PerPage = core.Int64Ptr(int64(5))
				listAllZoneUserAgentRulesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.ListAllZoneUserAgentRules(listAllZoneUserAgentRulesOptionsModel)
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
	Describe(`CreateZoneUserAgentRule(createZoneUserAgentRuleOptions *CreateZoneUserAgentRuleOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		createZoneUserAgentRulePath := "/v1/testString/zones/testString/firewall/ua_rules"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(createZoneUserAgentRulePath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateZoneUserAgentRule with error: Operation response processing error`, func() {
				testService, testServiceErr := useragentblockingrulesv1.NewUserAgentBlockingRulesV1(&useragentblockingrulesv1.UserAgentBlockingRulesV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the UseragentRuleInputConfiguration model
				useragentRuleInputConfigurationModel := new(useragentblockingrulesv1.UseragentRuleInputConfiguration)
				useragentRuleInputConfigurationModel.Target = core.StringPtr("ua")
				useragentRuleInputConfigurationModel.Value = core.StringPtr("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_5) AppleWebKit/603.2.4 (KHTML, like Gecko) Version/10.1.1 Safari/603.2.4")

				// Construct an instance of the CreateZoneUserAgentRuleOptions model
				createZoneUserAgentRuleOptionsModel := new(useragentblockingrulesv1.CreateZoneUserAgentRuleOptions)
				createZoneUserAgentRuleOptionsModel.Paused = core.BoolPtr(true)
				createZoneUserAgentRuleOptionsModel.Description = core.StringPtr("Prevent access from abusive clients identified by this UserAgent to mitigate DDoS attack")
				createZoneUserAgentRuleOptionsModel.Mode = core.StringPtr("block")
				createZoneUserAgentRuleOptionsModel.Configuration = useragentRuleInputConfigurationModel
				createZoneUserAgentRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.CreateZoneUserAgentRule(createZoneUserAgentRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CreateZoneUserAgentRule(createZoneUserAgentRuleOptions *CreateZoneUserAgentRuleOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		createZoneUserAgentRulePath := "/v1/testString/zones/testString/firewall/ua_rules"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(createZoneUserAgentRulePath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "92f17202ed8bd63d69a66b86a49a8f6b", "paused": true, "description": "Prevent access from abusive clients identified by this UserAgent to mitigate DDoS attack", "mode": "block", "configuration": {"target": "ua", "value": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_5) AppleWebKit/603.2.4 (KHTML, like Gecko) Version/10.1.1 Safari/603.2.4"}}}`)
				}))
			})
			It(`Invoke CreateZoneUserAgentRule successfully`, func() {
				testService, testServiceErr := useragentblockingrulesv1.NewUserAgentBlockingRulesV1(&useragentblockingrulesv1.UserAgentBlockingRulesV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.CreateZoneUserAgentRule(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UseragentRuleInputConfiguration model
				useragentRuleInputConfigurationModel := new(useragentblockingrulesv1.UseragentRuleInputConfiguration)
				useragentRuleInputConfigurationModel.Target = core.StringPtr("ua")
				useragentRuleInputConfigurationModel.Value = core.StringPtr("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_5) AppleWebKit/603.2.4 (KHTML, like Gecko) Version/10.1.1 Safari/603.2.4")

				// Construct an instance of the CreateZoneUserAgentRuleOptions model
				createZoneUserAgentRuleOptionsModel := new(useragentblockingrulesv1.CreateZoneUserAgentRuleOptions)
				createZoneUserAgentRuleOptionsModel.Paused = core.BoolPtr(true)
				createZoneUserAgentRuleOptionsModel.Description = core.StringPtr("Prevent access from abusive clients identified by this UserAgent to mitigate DDoS attack")
				createZoneUserAgentRuleOptionsModel.Mode = core.StringPtr("block")
				createZoneUserAgentRuleOptionsModel.Configuration = useragentRuleInputConfigurationModel
				createZoneUserAgentRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.CreateZoneUserAgentRule(createZoneUserAgentRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke CreateZoneUserAgentRule with error: Operation request error`, func() {
				testService, testServiceErr := useragentblockingrulesv1.NewUserAgentBlockingRulesV1(&useragentblockingrulesv1.UserAgentBlockingRulesV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the UseragentRuleInputConfiguration model
				useragentRuleInputConfigurationModel := new(useragentblockingrulesv1.UseragentRuleInputConfiguration)
				useragentRuleInputConfigurationModel.Target = core.StringPtr("ua")
				useragentRuleInputConfigurationModel.Value = core.StringPtr("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_5) AppleWebKit/603.2.4 (KHTML, like Gecko) Version/10.1.1 Safari/603.2.4")

				// Construct an instance of the CreateZoneUserAgentRuleOptions model
				createZoneUserAgentRuleOptionsModel := new(useragentblockingrulesv1.CreateZoneUserAgentRuleOptions)
				createZoneUserAgentRuleOptionsModel.Paused = core.BoolPtr(true)
				createZoneUserAgentRuleOptionsModel.Description = core.StringPtr("Prevent access from abusive clients identified by this UserAgent to mitigate DDoS attack")
				createZoneUserAgentRuleOptionsModel.Mode = core.StringPtr("block")
				createZoneUserAgentRuleOptionsModel.Configuration = useragentRuleInputConfigurationModel
				createZoneUserAgentRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.CreateZoneUserAgentRule(createZoneUserAgentRuleOptionsModel)
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
	Describe(`DeleteZoneUserAgentRule(deleteZoneUserAgentRuleOptions *DeleteZoneUserAgentRuleOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		deleteZoneUserAgentRulePath := "/v1/testString/zones/testString/firewall/ua_rules/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(deleteZoneUserAgentRulePath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteZoneUserAgentRule with error: Operation response processing error`, func() {
				testService, testServiceErr := useragentblockingrulesv1.NewUserAgentBlockingRulesV1(&useragentblockingrulesv1.UserAgentBlockingRulesV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the DeleteZoneUserAgentRuleOptions model
				deleteZoneUserAgentRuleOptionsModel := new(useragentblockingrulesv1.DeleteZoneUserAgentRuleOptions)
				deleteZoneUserAgentRuleOptionsModel.UseragentRuleIdentifier = core.StringPtr("testString")
				deleteZoneUserAgentRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.DeleteZoneUserAgentRule(deleteZoneUserAgentRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteZoneUserAgentRule(deleteZoneUserAgentRuleOptions *DeleteZoneUserAgentRuleOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		deleteZoneUserAgentRulePath := "/v1/testString/zones/testString/firewall/ua_rules/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(deleteZoneUserAgentRulePath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "f1aba936b94213e5b8dca0c0dbf1f9cc"}}`)
				}))
			})
			It(`Invoke DeleteZoneUserAgentRule successfully`, func() {
				testService, testServiceErr := useragentblockingrulesv1.NewUserAgentBlockingRulesV1(&useragentblockingrulesv1.UserAgentBlockingRulesV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.DeleteZoneUserAgentRule(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteZoneUserAgentRuleOptions model
				deleteZoneUserAgentRuleOptionsModel := new(useragentblockingrulesv1.DeleteZoneUserAgentRuleOptions)
				deleteZoneUserAgentRuleOptionsModel.UseragentRuleIdentifier = core.StringPtr("testString")
				deleteZoneUserAgentRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.DeleteZoneUserAgentRule(deleteZoneUserAgentRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke DeleteZoneUserAgentRule with error: Operation validation and request error`, func() {
				testService, testServiceErr := useragentblockingrulesv1.NewUserAgentBlockingRulesV1(&useragentblockingrulesv1.UserAgentBlockingRulesV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the DeleteZoneUserAgentRuleOptions model
				deleteZoneUserAgentRuleOptionsModel := new(useragentblockingrulesv1.DeleteZoneUserAgentRuleOptions)
				deleteZoneUserAgentRuleOptionsModel.UseragentRuleIdentifier = core.StringPtr("testString")
				deleteZoneUserAgentRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.DeleteZoneUserAgentRule(deleteZoneUserAgentRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteZoneUserAgentRuleOptions model with no property values
				deleteZoneUserAgentRuleOptionsModelNew := new(useragentblockingrulesv1.DeleteZoneUserAgentRuleOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.DeleteZoneUserAgentRule(deleteZoneUserAgentRuleOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetUserAgentRule(getUserAgentRuleOptions *GetUserAgentRuleOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getUserAgentRulePath := "/v1/testString/zones/testString/firewall/ua_rules/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getUserAgentRulePath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetUserAgentRule with error: Operation response processing error`, func() {
				testService, testServiceErr := useragentblockingrulesv1.NewUserAgentBlockingRulesV1(&useragentblockingrulesv1.UserAgentBlockingRulesV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetUserAgentRuleOptions model
				getUserAgentRuleOptionsModel := new(useragentblockingrulesv1.GetUserAgentRuleOptions)
				getUserAgentRuleOptionsModel.UseragentRuleIdentifier = core.StringPtr("testString")
				getUserAgentRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.GetUserAgentRule(getUserAgentRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetUserAgentRule(getUserAgentRuleOptions *GetUserAgentRuleOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getUserAgentRulePath := "/v1/testString/zones/testString/firewall/ua_rules/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getUserAgentRulePath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "92f17202ed8bd63d69a66b86a49a8f6b", "paused": true, "description": "Prevent access from abusive clients identified by this UserAgent to mitigate DDoS attack", "mode": "block", "configuration": {"target": "ua", "value": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_5) AppleWebKit/603.2.4 (KHTML, like Gecko) Version/10.1.1 Safari/603.2.4"}}}`)
				}))
			})
			It(`Invoke GetUserAgentRule successfully`, func() {
				testService, testServiceErr := useragentblockingrulesv1.NewUserAgentBlockingRulesV1(&useragentblockingrulesv1.UserAgentBlockingRulesV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetUserAgentRule(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetUserAgentRuleOptions model
				getUserAgentRuleOptionsModel := new(useragentblockingrulesv1.GetUserAgentRuleOptions)
				getUserAgentRuleOptionsModel.UseragentRuleIdentifier = core.StringPtr("testString")
				getUserAgentRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetUserAgentRule(getUserAgentRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetUserAgentRule with error: Operation validation and request error`, func() {
				testService, testServiceErr := useragentblockingrulesv1.NewUserAgentBlockingRulesV1(&useragentblockingrulesv1.UserAgentBlockingRulesV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetUserAgentRuleOptions model
				getUserAgentRuleOptionsModel := new(useragentblockingrulesv1.GetUserAgentRuleOptions)
				getUserAgentRuleOptionsModel.UseragentRuleIdentifier = core.StringPtr("testString")
				getUserAgentRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.GetUserAgentRule(getUserAgentRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetUserAgentRuleOptions model with no property values
				getUserAgentRuleOptionsModelNew := new(useragentblockingrulesv1.GetUserAgentRuleOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.GetUserAgentRule(getUserAgentRuleOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateUserAgentRule(updateUserAgentRuleOptions *UpdateUserAgentRuleOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateUserAgentRulePath := "/v1/testString/zones/testString/firewall/ua_rules/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(updateUserAgentRulePath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateUserAgentRule with error: Operation response processing error`, func() {
				testService, testServiceErr := useragentblockingrulesv1.NewUserAgentBlockingRulesV1(&useragentblockingrulesv1.UserAgentBlockingRulesV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the UseragentRuleInputConfiguration model
				useragentRuleInputConfigurationModel := new(useragentblockingrulesv1.UseragentRuleInputConfiguration)
				useragentRuleInputConfigurationModel.Target = core.StringPtr("ua")
				useragentRuleInputConfigurationModel.Value = core.StringPtr("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_5) AppleWebKit/603.2.4 (KHTML, like Gecko) Version/10.1.1 Safari/603.2.4")

				// Construct an instance of the UpdateUserAgentRuleOptions model
				updateUserAgentRuleOptionsModel := new(useragentblockingrulesv1.UpdateUserAgentRuleOptions)
				updateUserAgentRuleOptionsModel.UseragentRuleIdentifier = core.StringPtr("testString")
				updateUserAgentRuleOptionsModel.Paused = core.BoolPtr(true)
				updateUserAgentRuleOptionsModel.Description = core.StringPtr("Prevent access from abusive clients identified by this UserAgent to mitigate DDoS attack")
				updateUserAgentRuleOptionsModel.Mode = core.StringPtr("block")
				updateUserAgentRuleOptionsModel.Configuration = useragentRuleInputConfigurationModel
				updateUserAgentRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.UpdateUserAgentRule(updateUserAgentRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateUserAgentRule(updateUserAgentRuleOptions *UpdateUserAgentRuleOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateUserAgentRulePath := "/v1/testString/zones/testString/firewall/ua_rules/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(updateUserAgentRulePath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "92f17202ed8bd63d69a66b86a49a8f6b", "paused": true, "description": "Prevent access from abusive clients identified by this UserAgent to mitigate DDoS attack", "mode": "block", "configuration": {"target": "ua", "value": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_5) AppleWebKit/603.2.4 (KHTML, like Gecko) Version/10.1.1 Safari/603.2.4"}}}`)
				}))
			})
			It(`Invoke UpdateUserAgentRule successfully`, func() {
				testService, testServiceErr := useragentblockingrulesv1.NewUserAgentBlockingRulesV1(&useragentblockingrulesv1.UserAgentBlockingRulesV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.UpdateUserAgentRule(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UseragentRuleInputConfiguration model
				useragentRuleInputConfigurationModel := new(useragentblockingrulesv1.UseragentRuleInputConfiguration)
				useragentRuleInputConfigurationModel.Target = core.StringPtr("ua")
				useragentRuleInputConfigurationModel.Value = core.StringPtr("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_5) AppleWebKit/603.2.4 (KHTML, like Gecko) Version/10.1.1 Safari/603.2.4")

				// Construct an instance of the UpdateUserAgentRuleOptions model
				updateUserAgentRuleOptionsModel := new(useragentblockingrulesv1.UpdateUserAgentRuleOptions)
				updateUserAgentRuleOptionsModel.UseragentRuleIdentifier = core.StringPtr("testString")
				updateUserAgentRuleOptionsModel.Paused = core.BoolPtr(true)
				updateUserAgentRuleOptionsModel.Description = core.StringPtr("Prevent access from abusive clients identified by this UserAgent to mitigate DDoS attack")
				updateUserAgentRuleOptionsModel.Mode = core.StringPtr("block")
				updateUserAgentRuleOptionsModel.Configuration = useragentRuleInputConfigurationModel
				updateUserAgentRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.UpdateUserAgentRule(updateUserAgentRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke UpdateUserAgentRule with error: Operation validation and request error`, func() {
				testService, testServiceErr := useragentblockingrulesv1.NewUserAgentBlockingRulesV1(&useragentblockingrulesv1.UserAgentBlockingRulesV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the UseragentRuleInputConfiguration model
				useragentRuleInputConfigurationModel := new(useragentblockingrulesv1.UseragentRuleInputConfiguration)
				useragentRuleInputConfigurationModel.Target = core.StringPtr("ua")
				useragentRuleInputConfigurationModel.Value = core.StringPtr("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_5) AppleWebKit/603.2.4 (KHTML, like Gecko) Version/10.1.1 Safari/603.2.4")

				// Construct an instance of the UpdateUserAgentRuleOptions model
				updateUserAgentRuleOptionsModel := new(useragentblockingrulesv1.UpdateUserAgentRuleOptions)
				updateUserAgentRuleOptionsModel.UseragentRuleIdentifier = core.StringPtr("testString")
				updateUserAgentRuleOptionsModel.Paused = core.BoolPtr(true)
				updateUserAgentRuleOptionsModel.Description = core.StringPtr("Prevent access from abusive clients identified by this UserAgent to mitigate DDoS attack")
				updateUserAgentRuleOptionsModel.Mode = core.StringPtr("block")
				updateUserAgentRuleOptionsModel.Configuration = useragentRuleInputConfigurationModel
				updateUserAgentRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.UpdateUserAgentRule(updateUserAgentRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateUserAgentRuleOptions model with no property values
				updateUserAgentRuleOptionsModelNew := new(useragentblockingrulesv1.UpdateUserAgentRuleOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.UpdateUserAgentRule(updateUserAgentRuleOptionsModelNew)
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
			zoneIdentifier := "testString"
			testService, _ := useragentblockingrulesv1.NewUserAgentBlockingRulesV1(&useragentblockingrulesv1.UserAgentBlockingRulesV1Options{
				URL:            "http://useragentblockingrulesv1modelgenerator.com",
				Authenticator:  &core.NoAuthAuthenticator{},
				Crn:            core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})
			It(`Invoke NewCreateZoneUserAgentRuleOptions successfully`, func() {
				// Construct an instance of the UseragentRuleInputConfiguration model
				useragentRuleInputConfigurationModel := new(useragentblockingrulesv1.UseragentRuleInputConfiguration)
				Expect(useragentRuleInputConfigurationModel).ToNot(BeNil())
				useragentRuleInputConfigurationModel.Target = core.StringPtr("ua")
				useragentRuleInputConfigurationModel.Value = core.StringPtr("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_5) AppleWebKit/603.2.4 (KHTML, like Gecko) Version/10.1.1 Safari/603.2.4")
				Expect(useragentRuleInputConfigurationModel.Target).To(Equal(core.StringPtr("ua")))
				Expect(useragentRuleInputConfigurationModel.Value).To(Equal(core.StringPtr("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_5) AppleWebKit/603.2.4 (KHTML, like Gecko) Version/10.1.1 Safari/603.2.4")))

				// Construct an instance of the CreateZoneUserAgentRuleOptions model
				createZoneUserAgentRuleOptionsModel := testService.NewCreateZoneUserAgentRuleOptions()
				createZoneUserAgentRuleOptionsModel.SetPaused(true)
				createZoneUserAgentRuleOptionsModel.SetDescription("Prevent access from abusive clients identified by this UserAgent to mitigate DDoS attack")
				createZoneUserAgentRuleOptionsModel.SetMode("block")
				createZoneUserAgentRuleOptionsModel.SetConfiguration(useragentRuleInputConfigurationModel)
				createZoneUserAgentRuleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createZoneUserAgentRuleOptionsModel).ToNot(BeNil())
				Expect(createZoneUserAgentRuleOptionsModel.Paused).To(Equal(core.BoolPtr(true)))
				Expect(createZoneUserAgentRuleOptionsModel.Description).To(Equal(core.StringPtr("Prevent access from abusive clients identified by this UserAgent to mitigate DDoS attack")))
				Expect(createZoneUserAgentRuleOptionsModel.Mode).To(Equal(core.StringPtr("block")))
				Expect(createZoneUserAgentRuleOptionsModel.Configuration).To(Equal(useragentRuleInputConfigurationModel))
				Expect(createZoneUserAgentRuleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteZoneUserAgentRuleOptions successfully`, func() {
				// Construct an instance of the DeleteZoneUserAgentRuleOptions model
				useragentRuleIdentifier := "testString"
				deleteZoneUserAgentRuleOptionsModel := testService.NewDeleteZoneUserAgentRuleOptions(useragentRuleIdentifier)
				deleteZoneUserAgentRuleOptionsModel.SetUseragentRuleIdentifier("testString")
				deleteZoneUserAgentRuleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteZoneUserAgentRuleOptionsModel).ToNot(BeNil())
				Expect(deleteZoneUserAgentRuleOptionsModel.UseragentRuleIdentifier).To(Equal(core.StringPtr("testString")))
				Expect(deleteZoneUserAgentRuleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetUserAgentRuleOptions successfully`, func() {
				// Construct an instance of the GetUserAgentRuleOptions model
				useragentRuleIdentifier := "testString"
				getUserAgentRuleOptionsModel := testService.NewGetUserAgentRuleOptions(useragentRuleIdentifier)
				getUserAgentRuleOptionsModel.SetUseragentRuleIdentifier("testString")
				getUserAgentRuleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getUserAgentRuleOptionsModel).ToNot(BeNil())
				Expect(getUserAgentRuleOptionsModel.UseragentRuleIdentifier).To(Equal(core.StringPtr("testString")))
				Expect(getUserAgentRuleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListAllZoneUserAgentRulesOptions successfully`, func() {
				// Construct an instance of the ListAllZoneUserAgentRulesOptions model
				listAllZoneUserAgentRulesOptionsModel := testService.NewListAllZoneUserAgentRulesOptions()
				listAllZoneUserAgentRulesOptionsModel.SetPage(int64(38))
				listAllZoneUserAgentRulesOptionsModel.SetPerPage(int64(5))
				listAllZoneUserAgentRulesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listAllZoneUserAgentRulesOptionsModel).ToNot(BeNil())
				Expect(listAllZoneUserAgentRulesOptionsModel.Page).To(Equal(core.Int64Ptr(int64(38))))
				Expect(listAllZoneUserAgentRulesOptionsModel.PerPage).To(Equal(core.Int64Ptr(int64(5))))
				Expect(listAllZoneUserAgentRulesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateUserAgentRuleOptions successfully`, func() {
				// Construct an instance of the UseragentRuleInputConfiguration model
				useragentRuleInputConfigurationModel := new(useragentblockingrulesv1.UseragentRuleInputConfiguration)
				Expect(useragentRuleInputConfigurationModel).ToNot(BeNil())
				useragentRuleInputConfigurationModel.Target = core.StringPtr("ua")
				useragentRuleInputConfigurationModel.Value = core.StringPtr("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_5) AppleWebKit/603.2.4 (KHTML, like Gecko) Version/10.1.1 Safari/603.2.4")
				Expect(useragentRuleInputConfigurationModel.Target).To(Equal(core.StringPtr("ua")))
				Expect(useragentRuleInputConfigurationModel.Value).To(Equal(core.StringPtr("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_5) AppleWebKit/603.2.4 (KHTML, like Gecko) Version/10.1.1 Safari/603.2.4")))

				// Construct an instance of the UpdateUserAgentRuleOptions model
				useragentRuleIdentifier := "testString"
				updateUserAgentRuleOptionsModel := testService.NewUpdateUserAgentRuleOptions(useragentRuleIdentifier)
				updateUserAgentRuleOptionsModel.SetUseragentRuleIdentifier("testString")
				updateUserAgentRuleOptionsModel.SetPaused(true)
				updateUserAgentRuleOptionsModel.SetDescription("Prevent access from abusive clients identified by this UserAgent to mitigate DDoS attack")
				updateUserAgentRuleOptionsModel.SetMode("block")
				updateUserAgentRuleOptionsModel.SetConfiguration(useragentRuleInputConfigurationModel)
				updateUserAgentRuleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateUserAgentRuleOptionsModel).ToNot(BeNil())
				Expect(updateUserAgentRuleOptionsModel.UseragentRuleIdentifier).To(Equal(core.StringPtr("testString")))
				Expect(updateUserAgentRuleOptionsModel.Paused).To(Equal(core.BoolPtr(true)))
				Expect(updateUserAgentRuleOptionsModel.Description).To(Equal(core.StringPtr("Prevent access from abusive clients identified by this UserAgent to mitigate DDoS attack")))
				Expect(updateUserAgentRuleOptionsModel.Mode).To(Equal(core.StringPtr("block")))
				Expect(updateUserAgentRuleOptionsModel.Configuration).To(Equal(useragentRuleInputConfigurationModel))
				Expect(updateUserAgentRuleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUseragentRuleInputConfiguration successfully`, func() {
				target := "ua"
				value := "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_5) AppleWebKit/603.2.4 (KHTML, like Gecko) Version/10.1.1 Safari/603.2.4"
				model, err := testService.NewUseragentRuleInputConfiguration(target, value)
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
