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

package zonefirewallaccessrulesv1_test

import (
	"bytes"
	"fmt"
	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.ibm.com/ibmcloud/networking-go-sdk/zonefirewallaccessrulesv1"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"time"
)

var _ = Describe(`ZoneFirewallAccessRulesV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		It(`Instantiate service client`, func() {
			testService, testServiceErr := zonefirewallaccessrulesv1.NewZoneFirewallAccessRulesV1(&zonefirewallaccessrulesv1.ZoneFirewallAccessRulesV1Options{
				Authenticator:  &core.NoAuthAuthenticator{},
				Crn:            core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})
			Expect(testService).ToNot(BeNil())
			Expect(testServiceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			testService, testServiceErr := zonefirewallaccessrulesv1.NewZoneFirewallAccessRulesV1(&zonefirewallaccessrulesv1.ZoneFirewallAccessRulesV1Options{
				URL:            "{BAD_URL_STRING",
				Crn:            core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			testService, testServiceErr := zonefirewallaccessrulesv1.NewZoneFirewallAccessRulesV1(&zonefirewallaccessrulesv1.ZoneFirewallAccessRulesV1Options{
				URL:            "https://zonefirewallaccessrulesv1/api",
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
			testService, testServiceErr := zonefirewallaccessrulesv1.NewZoneFirewallAccessRulesV1(&zonefirewallaccessrulesv1.ZoneFirewallAccessRulesV1Options{})
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
				"ZONE_FIREWALL_ACCESS_RULES_URL":       "https://zonefirewallaccessrulesv1/api",
				"ZONE_FIREWALL_ACCESS_RULES_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := zonefirewallaccessrulesv1.NewZoneFirewallAccessRulesV1UsingExternalConfig(&zonefirewallaccessrulesv1.ZoneFirewallAccessRulesV1Options{
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := zonefirewallaccessrulesv1.NewZoneFirewallAccessRulesV1UsingExternalConfig(&zonefirewallaccessrulesv1.ZoneFirewallAccessRulesV1Options{
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
				testService, testServiceErr := zonefirewallaccessrulesv1.NewZoneFirewallAccessRulesV1UsingExternalConfig(&zonefirewallaccessrulesv1.ZoneFirewallAccessRulesV1Options{
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
				"ZONE_FIREWALL_ACCESS_RULES_URL":       "https://zonefirewallaccessrulesv1/api",
				"ZONE_FIREWALL_ACCESS_RULES_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := zonefirewallaccessrulesv1.NewZoneFirewallAccessRulesV1UsingExternalConfig(&zonefirewallaccessrulesv1.ZoneFirewallAccessRulesV1Options{
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
				"ZONE_FIREWALL_ACCESS_RULES_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := zonefirewallaccessrulesv1.NewZoneFirewallAccessRulesV1UsingExternalConfig(&zonefirewallaccessrulesv1.ZoneFirewallAccessRulesV1Options{
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
	Describe(`ListAllZoneAccessRules(listAllZoneAccessRulesOptions *ListAllZoneAccessRulesOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		listAllZoneAccessRulesPath := "/v1/testString/zones/testString/firewall/access_rules/rules"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listAllZoneAccessRulesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["notes"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["mode"]).To(Equal([]string{"block"}))

					Expect(req.URL.Query()["configuration.target"]).To(Equal([]string{"ip"}))

					Expect(req.URL.Query()["configuration.value"]).To(Equal([]string{"1.2.3.4"}))

					Expect(req.URL.Query()["page"]).To(Equal([]string{fmt.Sprint(int64(38))}))

					Expect(req.URL.Query()["per_page"]).To(Equal([]string{fmt.Sprint(int64(5))}))

					Expect(req.URL.Query()["order"]).To(Equal([]string{"configuration.target"}))

					Expect(req.URL.Query()["direction"]).To(Equal([]string{"asc"}))

					Expect(req.URL.Query()["match"]).To(Equal([]string{"any"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListAllZoneAccessRules with error: Operation response processing error`, func() {
				testService, testServiceErr := zonefirewallaccessrulesv1.NewZoneFirewallAccessRulesV1(&zonefirewallaccessrulesv1.ZoneFirewallAccessRulesV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ListAllZoneAccessRulesOptions model
				listAllZoneAccessRulesOptionsModel := new(zonefirewallaccessrulesv1.ListAllZoneAccessRulesOptions)
				listAllZoneAccessRulesOptionsModel.Notes = core.StringPtr("testString")
				listAllZoneAccessRulesOptionsModel.Mode = core.StringPtr("block")
				listAllZoneAccessRulesOptionsModel.ConfigurationTarget = core.StringPtr("ip")
				listAllZoneAccessRulesOptionsModel.ConfigurationValue = core.StringPtr("1.2.3.4")
				listAllZoneAccessRulesOptionsModel.Page = core.Int64Ptr(int64(38))
				listAllZoneAccessRulesOptionsModel.PerPage = core.Int64Ptr(int64(5))
				listAllZoneAccessRulesOptionsModel.Order = core.StringPtr("configuration.target")
				listAllZoneAccessRulesOptionsModel.Direction = core.StringPtr("asc")
				listAllZoneAccessRulesOptionsModel.Match = core.StringPtr("any")
				listAllZoneAccessRulesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.ListAllZoneAccessRules(listAllZoneAccessRulesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListAllZoneAccessRules(listAllZoneAccessRulesOptions *ListAllZoneAccessRulesOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		listAllZoneAccessRulesPath := "/v1/testString/zones/testString/firewall/access_rules/rules"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listAllZoneAccessRulesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["notes"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["mode"]).To(Equal([]string{"block"}))

					Expect(req.URL.Query()["configuration.target"]).To(Equal([]string{"ip"}))

					Expect(req.URL.Query()["configuration.value"]).To(Equal([]string{"1.2.3.4"}))

					Expect(req.URL.Query()["page"]).To(Equal([]string{fmt.Sprint(int64(38))}))

					Expect(req.URL.Query()["per_page"]).To(Equal([]string{fmt.Sprint(int64(5))}))

					Expect(req.URL.Query()["order"]).To(Equal([]string{"configuration.target"}))

					Expect(req.URL.Query()["direction"]).To(Equal([]string{"asc"}))

					Expect(req.URL.Query()["match"]).To(Equal([]string{"any"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": [{"id": "92f17202ed8bd63d69a66b86a49a8f6b", "notes": "This rule is set because of an event that occurred and caused X.", "allowed_modes": ["block"], "mode": "block", "scope": {"type": "account"}, "created_on": "2014-01-01T05:20:00.12345Z", "modified_on": "2014-01-01T05:20:00.12345Z", "configuration": {"target": "ip", "value": "ip example 198.51.100.4; ip_range example 198.51.100.4/16 ; asn example AS12345; country example AZ"}}], "result_info": {"page": 1, "per_page": 2, "count": 1, "total_count": 200}}`)
				}))
			})
			It(`Invoke ListAllZoneAccessRules successfully`, func() {
				testService, testServiceErr := zonefirewallaccessrulesv1.NewZoneFirewallAccessRulesV1(&zonefirewallaccessrulesv1.ZoneFirewallAccessRulesV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListAllZoneAccessRules(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListAllZoneAccessRulesOptions model
				listAllZoneAccessRulesOptionsModel := new(zonefirewallaccessrulesv1.ListAllZoneAccessRulesOptions)
				listAllZoneAccessRulesOptionsModel.Notes = core.StringPtr("testString")
				listAllZoneAccessRulesOptionsModel.Mode = core.StringPtr("block")
				listAllZoneAccessRulesOptionsModel.ConfigurationTarget = core.StringPtr("ip")
				listAllZoneAccessRulesOptionsModel.ConfigurationValue = core.StringPtr("1.2.3.4")
				listAllZoneAccessRulesOptionsModel.Page = core.Int64Ptr(int64(38))
				listAllZoneAccessRulesOptionsModel.PerPage = core.Int64Ptr(int64(5))
				listAllZoneAccessRulesOptionsModel.Order = core.StringPtr("configuration.target")
				listAllZoneAccessRulesOptionsModel.Direction = core.StringPtr("asc")
				listAllZoneAccessRulesOptionsModel.Match = core.StringPtr("any")
				listAllZoneAccessRulesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListAllZoneAccessRules(listAllZoneAccessRulesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke ListAllZoneAccessRules with error: Operation request error`, func() {
				testService, testServiceErr := zonefirewallaccessrulesv1.NewZoneFirewallAccessRulesV1(&zonefirewallaccessrulesv1.ZoneFirewallAccessRulesV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ListAllZoneAccessRulesOptions model
				listAllZoneAccessRulesOptionsModel := new(zonefirewallaccessrulesv1.ListAllZoneAccessRulesOptions)
				listAllZoneAccessRulesOptionsModel.Notes = core.StringPtr("testString")
				listAllZoneAccessRulesOptionsModel.Mode = core.StringPtr("block")
				listAllZoneAccessRulesOptionsModel.ConfigurationTarget = core.StringPtr("ip")
				listAllZoneAccessRulesOptionsModel.ConfigurationValue = core.StringPtr("1.2.3.4")
				listAllZoneAccessRulesOptionsModel.Page = core.Int64Ptr(int64(38))
				listAllZoneAccessRulesOptionsModel.PerPage = core.Int64Ptr(int64(5))
				listAllZoneAccessRulesOptionsModel.Order = core.StringPtr("configuration.target")
				listAllZoneAccessRulesOptionsModel.Direction = core.StringPtr("asc")
				listAllZoneAccessRulesOptionsModel.Match = core.StringPtr("any")
				listAllZoneAccessRulesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.ListAllZoneAccessRules(listAllZoneAccessRulesOptionsModel)
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
	Describe(`CreateZoneAccessRule(createZoneAccessRuleOptions *CreateZoneAccessRuleOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		createZoneAccessRulePath := "/v1/testString/zones/testString/firewall/access_rules/rules"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(createZoneAccessRulePath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateZoneAccessRule with error: Operation response processing error`, func() {
				testService, testServiceErr := zonefirewallaccessrulesv1.NewZoneFirewallAccessRulesV1(&zonefirewallaccessrulesv1.ZoneFirewallAccessRulesV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ZoneAccessRuleInputConfiguration model
				zoneAccessRuleInputConfigurationModel := new(zonefirewallaccessrulesv1.ZoneAccessRuleInputConfiguration)
				zoneAccessRuleInputConfigurationModel.Target = core.StringPtr("ip")
				zoneAccessRuleInputConfigurationModel.Value = core.StringPtr("ip example 198.51.100.4; ip_range example 198.51.100.4/16 ; asn example AS12345; country example AZ")

				// Construct an instance of the CreateZoneAccessRuleOptions model
				createZoneAccessRuleOptionsModel := new(zonefirewallaccessrulesv1.CreateZoneAccessRuleOptions)
				createZoneAccessRuleOptionsModel.Mode = core.StringPtr("block")
				createZoneAccessRuleOptionsModel.Notes = core.StringPtr("This rule is added because of event X that occurred on date xyz")
				createZoneAccessRuleOptionsModel.Configuration = zoneAccessRuleInputConfigurationModel
				createZoneAccessRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.CreateZoneAccessRule(createZoneAccessRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CreateZoneAccessRule(createZoneAccessRuleOptions *CreateZoneAccessRuleOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		createZoneAccessRulePath := "/v1/testString/zones/testString/firewall/access_rules/rules"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(createZoneAccessRulePath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "92f17202ed8bd63d69a66b86a49a8f6b", "notes": "This rule is set because of an event that occurred and caused X.", "allowed_modes": ["block"], "mode": "block", "scope": {"type": "account"}, "created_on": "2014-01-01T05:20:00.12345Z", "modified_on": "2014-01-01T05:20:00.12345Z", "configuration": {"target": "ip", "value": "ip example 198.51.100.4; ip_range example 198.51.100.4/16 ; asn example AS12345; country example AZ"}}}`)
				}))
			})
			It(`Invoke CreateZoneAccessRule successfully`, func() {
				testService, testServiceErr := zonefirewallaccessrulesv1.NewZoneFirewallAccessRulesV1(&zonefirewallaccessrulesv1.ZoneFirewallAccessRulesV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.CreateZoneAccessRule(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ZoneAccessRuleInputConfiguration model
				zoneAccessRuleInputConfigurationModel := new(zonefirewallaccessrulesv1.ZoneAccessRuleInputConfiguration)
				zoneAccessRuleInputConfigurationModel.Target = core.StringPtr("ip")
				zoneAccessRuleInputConfigurationModel.Value = core.StringPtr("ip example 198.51.100.4; ip_range example 198.51.100.4/16 ; asn example AS12345; country example AZ")

				// Construct an instance of the CreateZoneAccessRuleOptions model
				createZoneAccessRuleOptionsModel := new(zonefirewallaccessrulesv1.CreateZoneAccessRuleOptions)
				createZoneAccessRuleOptionsModel.Mode = core.StringPtr("block")
				createZoneAccessRuleOptionsModel.Notes = core.StringPtr("This rule is added because of event X that occurred on date xyz")
				createZoneAccessRuleOptionsModel.Configuration = zoneAccessRuleInputConfigurationModel
				createZoneAccessRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.CreateZoneAccessRule(createZoneAccessRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke CreateZoneAccessRule with error: Operation request error`, func() {
				testService, testServiceErr := zonefirewallaccessrulesv1.NewZoneFirewallAccessRulesV1(&zonefirewallaccessrulesv1.ZoneFirewallAccessRulesV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ZoneAccessRuleInputConfiguration model
				zoneAccessRuleInputConfigurationModel := new(zonefirewallaccessrulesv1.ZoneAccessRuleInputConfiguration)
				zoneAccessRuleInputConfigurationModel.Target = core.StringPtr("ip")
				zoneAccessRuleInputConfigurationModel.Value = core.StringPtr("ip example 198.51.100.4; ip_range example 198.51.100.4/16 ; asn example AS12345; country example AZ")

				// Construct an instance of the CreateZoneAccessRuleOptions model
				createZoneAccessRuleOptionsModel := new(zonefirewallaccessrulesv1.CreateZoneAccessRuleOptions)
				createZoneAccessRuleOptionsModel.Mode = core.StringPtr("block")
				createZoneAccessRuleOptionsModel.Notes = core.StringPtr("This rule is added because of event X that occurred on date xyz")
				createZoneAccessRuleOptionsModel.Configuration = zoneAccessRuleInputConfigurationModel
				createZoneAccessRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.CreateZoneAccessRule(createZoneAccessRuleOptionsModel)
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
	Describe(`DeleteZoneAccessRule(deleteZoneAccessRuleOptions *DeleteZoneAccessRuleOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		deleteZoneAccessRulePath := "/v1/testString/zones/testString/firewall/access_rules/rules/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(deleteZoneAccessRulePath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteZoneAccessRule with error: Operation response processing error`, func() {
				testService, testServiceErr := zonefirewallaccessrulesv1.NewZoneFirewallAccessRulesV1(&zonefirewallaccessrulesv1.ZoneFirewallAccessRulesV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the DeleteZoneAccessRuleOptions model
				deleteZoneAccessRuleOptionsModel := new(zonefirewallaccessrulesv1.DeleteZoneAccessRuleOptions)
				deleteZoneAccessRuleOptionsModel.AccessruleIdentifier = core.StringPtr("testString")
				deleteZoneAccessRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.DeleteZoneAccessRule(deleteZoneAccessRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteZoneAccessRule(deleteZoneAccessRuleOptions *DeleteZoneAccessRuleOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		deleteZoneAccessRulePath := "/v1/testString/zones/testString/firewall/access_rules/rules/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(deleteZoneAccessRulePath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "f1aba936b94213e5b8dca0c0dbf1f9cc"}}`)
				}))
			})
			It(`Invoke DeleteZoneAccessRule successfully`, func() {
				testService, testServiceErr := zonefirewallaccessrulesv1.NewZoneFirewallAccessRulesV1(&zonefirewallaccessrulesv1.ZoneFirewallAccessRulesV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.DeleteZoneAccessRule(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteZoneAccessRuleOptions model
				deleteZoneAccessRuleOptionsModel := new(zonefirewallaccessrulesv1.DeleteZoneAccessRuleOptions)
				deleteZoneAccessRuleOptionsModel.AccessruleIdentifier = core.StringPtr("testString")
				deleteZoneAccessRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.DeleteZoneAccessRule(deleteZoneAccessRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke DeleteZoneAccessRule with error: Operation validation and request error`, func() {
				testService, testServiceErr := zonefirewallaccessrulesv1.NewZoneFirewallAccessRulesV1(&zonefirewallaccessrulesv1.ZoneFirewallAccessRulesV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the DeleteZoneAccessRuleOptions model
				deleteZoneAccessRuleOptionsModel := new(zonefirewallaccessrulesv1.DeleteZoneAccessRuleOptions)
				deleteZoneAccessRuleOptionsModel.AccessruleIdentifier = core.StringPtr("testString")
				deleteZoneAccessRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.DeleteZoneAccessRule(deleteZoneAccessRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteZoneAccessRuleOptions model with no property values
				deleteZoneAccessRuleOptionsModelNew := new(zonefirewallaccessrulesv1.DeleteZoneAccessRuleOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.DeleteZoneAccessRule(deleteZoneAccessRuleOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetZoneAccessRule(getZoneAccessRuleOptions *GetZoneAccessRuleOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getZoneAccessRulePath := "/v1/testString/zones/testString/firewall/access_rules/rules/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getZoneAccessRulePath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetZoneAccessRule with error: Operation response processing error`, func() {
				testService, testServiceErr := zonefirewallaccessrulesv1.NewZoneFirewallAccessRulesV1(&zonefirewallaccessrulesv1.ZoneFirewallAccessRulesV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetZoneAccessRuleOptions model
				getZoneAccessRuleOptionsModel := new(zonefirewallaccessrulesv1.GetZoneAccessRuleOptions)
				getZoneAccessRuleOptionsModel.AccessruleIdentifier = core.StringPtr("testString")
				getZoneAccessRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.GetZoneAccessRule(getZoneAccessRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetZoneAccessRule(getZoneAccessRuleOptions *GetZoneAccessRuleOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getZoneAccessRulePath := "/v1/testString/zones/testString/firewall/access_rules/rules/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getZoneAccessRulePath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "92f17202ed8bd63d69a66b86a49a8f6b", "notes": "This rule is set because of an event that occurred and caused X.", "allowed_modes": ["block"], "mode": "block", "scope": {"type": "account"}, "created_on": "2014-01-01T05:20:00.12345Z", "modified_on": "2014-01-01T05:20:00.12345Z", "configuration": {"target": "ip", "value": "ip example 198.51.100.4; ip_range example 198.51.100.4/16 ; asn example AS12345; country example AZ"}}}`)
				}))
			})
			It(`Invoke GetZoneAccessRule successfully`, func() {
				testService, testServiceErr := zonefirewallaccessrulesv1.NewZoneFirewallAccessRulesV1(&zonefirewallaccessrulesv1.ZoneFirewallAccessRulesV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetZoneAccessRule(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetZoneAccessRuleOptions model
				getZoneAccessRuleOptionsModel := new(zonefirewallaccessrulesv1.GetZoneAccessRuleOptions)
				getZoneAccessRuleOptionsModel.AccessruleIdentifier = core.StringPtr("testString")
				getZoneAccessRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetZoneAccessRule(getZoneAccessRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetZoneAccessRule with error: Operation validation and request error`, func() {
				testService, testServiceErr := zonefirewallaccessrulesv1.NewZoneFirewallAccessRulesV1(&zonefirewallaccessrulesv1.ZoneFirewallAccessRulesV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetZoneAccessRuleOptions model
				getZoneAccessRuleOptionsModel := new(zonefirewallaccessrulesv1.GetZoneAccessRuleOptions)
				getZoneAccessRuleOptionsModel.AccessruleIdentifier = core.StringPtr("testString")
				getZoneAccessRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.GetZoneAccessRule(getZoneAccessRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetZoneAccessRuleOptions model with no property values
				getZoneAccessRuleOptionsModelNew := new(zonefirewallaccessrulesv1.GetZoneAccessRuleOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.GetZoneAccessRule(getZoneAccessRuleOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateZoneAccessRule(updateZoneAccessRuleOptions *UpdateZoneAccessRuleOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateZoneAccessRulePath := "/v1/testString/zones/testString/firewall/access_rules/rules/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(updateZoneAccessRulePath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateZoneAccessRule with error: Operation response processing error`, func() {
				testService, testServiceErr := zonefirewallaccessrulesv1.NewZoneFirewallAccessRulesV1(&zonefirewallaccessrulesv1.ZoneFirewallAccessRulesV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the UpdateZoneAccessRuleOptions model
				updateZoneAccessRuleOptionsModel := new(zonefirewallaccessrulesv1.UpdateZoneAccessRuleOptions)
				updateZoneAccessRuleOptionsModel.AccessruleIdentifier = core.StringPtr("testString")
				updateZoneAccessRuleOptionsModel.Mode = core.StringPtr("block")
				updateZoneAccessRuleOptionsModel.Notes = core.StringPtr("This rule is added because of event X that occurred on date xyz")
				updateZoneAccessRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.UpdateZoneAccessRule(updateZoneAccessRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateZoneAccessRule(updateZoneAccessRuleOptions *UpdateZoneAccessRuleOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateZoneAccessRulePath := "/v1/testString/zones/testString/firewall/access_rules/rules/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(updateZoneAccessRulePath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "92f17202ed8bd63d69a66b86a49a8f6b", "notes": "This rule is set because of an event that occurred and caused X.", "allowed_modes": ["block"], "mode": "block", "scope": {"type": "account"}, "created_on": "2014-01-01T05:20:00.12345Z", "modified_on": "2014-01-01T05:20:00.12345Z", "configuration": {"target": "ip", "value": "ip example 198.51.100.4; ip_range example 198.51.100.4/16 ; asn example AS12345; country example AZ"}}}`)
				}))
			})
			It(`Invoke UpdateZoneAccessRule successfully`, func() {
				testService, testServiceErr := zonefirewallaccessrulesv1.NewZoneFirewallAccessRulesV1(&zonefirewallaccessrulesv1.ZoneFirewallAccessRulesV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.UpdateZoneAccessRule(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateZoneAccessRuleOptions model
				updateZoneAccessRuleOptionsModel := new(zonefirewallaccessrulesv1.UpdateZoneAccessRuleOptions)
				updateZoneAccessRuleOptionsModel.AccessruleIdentifier = core.StringPtr("testString")
				updateZoneAccessRuleOptionsModel.Mode = core.StringPtr("block")
				updateZoneAccessRuleOptionsModel.Notes = core.StringPtr("This rule is added because of event X that occurred on date xyz")
				updateZoneAccessRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.UpdateZoneAccessRule(updateZoneAccessRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke UpdateZoneAccessRule with error: Operation validation and request error`, func() {
				testService, testServiceErr := zonefirewallaccessrulesv1.NewZoneFirewallAccessRulesV1(&zonefirewallaccessrulesv1.ZoneFirewallAccessRulesV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the UpdateZoneAccessRuleOptions model
				updateZoneAccessRuleOptionsModel := new(zonefirewallaccessrulesv1.UpdateZoneAccessRuleOptions)
				updateZoneAccessRuleOptionsModel.AccessruleIdentifier = core.StringPtr("testString")
				updateZoneAccessRuleOptionsModel.Mode = core.StringPtr("block")
				updateZoneAccessRuleOptionsModel.Notes = core.StringPtr("This rule is added because of event X that occurred on date xyz")
				updateZoneAccessRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.UpdateZoneAccessRule(updateZoneAccessRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateZoneAccessRuleOptions model with no property values
				updateZoneAccessRuleOptionsModelNew := new(zonefirewallaccessrulesv1.UpdateZoneAccessRuleOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.UpdateZoneAccessRule(updateZoneAccessRuleOptionsModelNew)
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
			testService, _ := zonefirewallaccessrulesv1.NewZoneFirewallAccessRulesV1(&zonefirewallaccessrulesv1.ZoneFirewallAccessRulesV1Options{
				URL:            "http://zonefirewallaccessrulesv1modelgenerator.com",
				Authenticator:  &core.NoAuthAuthenticator{},
				Crn:            core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})
			It(`Invoke NewCreateZoneAccessRuleOptions successfully`, func() {
				// Construct an instance of the ZoneAccessRuleInputConfiguration model
				zoneAccessRuleInputConfigurationModel := new(zonefirewallaccessrulesv1.ZoneAccessRuleInputConfiguration)
				Expect(zoneAccessRuleInputConfigurationModel).ToNot(BeNil())
				zoneAccessRuleInputConfigurationModel.Target = core.StringPtr("ip")
				zoneAccessRuleInputConfigurationModel.Value = core.StringPtr("ip example 198.51.100.4; ip_range example 198.51.100.4/16 ; asn example AS12345; country example AZ")
				Expect(zoneAccessRuleInputConfigurationModel.Target).To(Equal(core.StringPtr("ip")))
				Expect(zoneAccessRuleInputConfigurationModel.Value).To(Equal(core.StringPtr("ip example 198.51.100.4; ip_range example 198.51.100.4/16 ; asn example AS12345; country example AZ")))

				// Construct an instance of the CreateZoneAccessRuleOptions model
				createZoneAccessRuleOptionsModel := testService.NewCreateZoneAccessRuleOptions()
				createZoneAccessRuleOptionsModel.SetMode("block")
				createZoneAccessRuleOptionsModel.SetNotes("This rule is added because of event X that occurred on date xyz")
				createZoneAccessRuleOptionsModel.SetConfiguration(zoneAccessRuleInputConfigurationModel)
				createZoneAccessRuleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createZoneAccessRuleOptionsModel).ToNot(BeNil())
				Expect(createZoneAccessRuleOptionsModel.Mode).To(Equal(core.StringPtr("block")))
				Expect(createZoneAccessRuleOptionsModel.Notes).To(Equal(core.StringPtr("This rule is added because of event X that occurred on date xyz")))
				Expect(createZoneAccessRuleOptionsModel.Configuration).To(Equal(zoneAccessRuleInputConfigurationModel))
				Expect(createZoneAccessRuleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteZoneAccessRuleOptions successfully`, func() {
				// Construct an instance of the DeleteZoneAccessRuleOptions model
				accessruleIdentifier := "testString"
				deleteZoneAccessRuleOptionsModel := testService.NewDeleteZoneAccessRuleOptions(accessruleIdentifier)
				deleteZoneAccessRuleOptionsModel.SetAccessruleIdentifier("testString")
				deleteZoneAccessRuleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteZoneAccessRuleOptionsModel).ToNot(BeNil())
				Expect(deleteZoneAccessRuleOptionsModel.AccessruleIdentifier).To(Equal(core.StringPtr("testString")))
				Expect(deleteZoneAccessRuleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetZoneAccessRuleOptions successfully`, func() {
				// Construct an instance of the GetZoneAccessRuleOptions model
				accessruleIdentifier := "testString"
				getZoneAccessRuleOptionsModel := testService.NewGetZoneAccessRuleOptions(accessruleIdentifier)
				getZoneAccessRuleOptionsModel.SetAccessruleIdentifier("testString")
				getZoneAccessRuleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getZoneAccessRuleOptionsModel).ToNot(BeNil())
				Expect(getZoneAccessRuleOptionsModel.AccessruleIdentifier).To(Equal(core.StringPtr("testString")))
				Expect(getZoneAccessRuleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListAllZoneAccessRulesOptions successfully`, func() {
				// Construct an instance of the ListAllZoneAccessRulesOptions model
				listAllZoneAccessRulesOptionsModel := testService.NewListAllZoneAccessRulesOptions()
				listAllZoneAccessRulesOptionsModel.SetNotes("testString")
				listAllZoneAccessRulesOptionsModel.SetMode("block")
				listAllZoneAccessRulesOptionsModel.SetConfigurationTarget("ip")
				listAllZoneAccessRulesOptionsModel.SetConfigurationValue("1.2.3.4")
				listAllZoneAccessRulesOptionsModel.SetPage(int64(38))
				listAllZoneAccessRulesOptionsModel.SetPerPage(int64(5))
				listAllZoneAccessRulesOptionsModel.SetOrder("configuration.target")
				listAllZoneAccessRulesOptionsModel.SetDirection("asc")
				listAllZoneAccessRulesOptionsModel.SetMatch("any")
				listAllZoneAccessRulesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listAllZoneAccessRulesOptionsModel).ToNot(BeNil())
				Expect(listAllZoneAccessRulesOptionsModel.Notes).To(Equal(core.StringPtr("testString")))
				Expect(listAllZoneAccessRulesOptionsModel.Mode).To(Equal(core.StringPtr("block")))
				Expect(listAllZoneAccessRulesOptionsModel.ConfigurationTarget).To(Equal(core.StringPtr("ip")))
				Expect(listAllZoneAccessRulesOptionsModel.ConfigurationValue).To(Equal(core.StringPtr("1.2.3.4")))
				Expect(listAllZoneAccessRulesOptionsModel.Page).To(Equal(core.Int64Ptr(int64(38))))
				Expect(listAllZoneAccessRulesOptionsModel.PerPage).To(Equal(core.Int64Ptr(int64(5))))
				Expect(listAllZoneAccessRulesOptionsModel.Order).To(Equal(core.StringPtr("configuration.target")))
				Expect(listAllZoneAccessRulesOptionsModel.Direction).To(Equal(core.StringPtr("asc")))
				Expect(listAllZoneAccessRulesOptionsModel.Match).To(Equal(core.StringPtr("any")))
				Expect(listAllZoneAccessRulesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateZoneAccessRuleOptions successfully`, func() {
				// Construct an instance of the UpdateZoneAccessRuleOptions model
				accessruleIdentifier := "testString"
				updateZoneAccessRuleOptionsModel := testService.NewUpdateZoneAccessRuleOptions(accessruleIdentifier)
				updateZoneAccessRuleOptionsModel.SetAccessruleIdentifier("testString")
				updateZoneAccessRuleOptionsModel.SetMode("block")
				updateZoneAccessRuleOptionsModel.SetNotes("This rule is added because of event X that occurred on date xyz")
				updateZoneAccessRuleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateZoneAccessRuleOptionsModel).ToNot(BeNil())
				Expect(updateZoneAccessRuleOptionsModel.AccessruleIdentifier).To(Equal(core.StringPtr("testString")))
				Expect(updateZoneAccessRuleOptionsModel.Mode).To(Equal(core.StringPtr("block")))
				Expect(updateZoneAccessRuleOptionsModel.Notes).To(Equal(core.StringPtr("This rule is added because of event X that occurred on date xyz")))
				Expect(updateZoneAccessRuleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewZoneAccessRuleInputConfiguration successfully`, func() {
				target := "ip"
				value := "ip example 198.51.100.4; ip_range example 198.51.100.4/16 ; asn example AS12345; country example AZ"
				model, err := testService.NewZoneAccessRuleInputConfiguration(target, value)
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
