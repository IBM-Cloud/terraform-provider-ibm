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

package rangeapplicationsv1_test

import (
	"bytes"
	"fmt"
	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.ibm.com/ibmcloud/networking-go-sdk/rangeapplicationsv1"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"time"
)

var _ = Describe(`RangeApplicationsV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		It(`Instantiate service client`, func() {
			testService, testServiceErr := rangeapplicationsv1.NewRangeApplicationsV1(&rangeapplicationsv1.RangeApplicationsV1Options{
				Authenticator:  &core.NoAuthAuthenticator{},
				Crn:            core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})
			Expect(testService).ToNot(BeNil())
			Expect(testServiceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			testService, testServiceErr := rangeapplicationsv1.NewRangeApplicationsV1(&rangeapplicationsv1.RangeApplicationsV1Options{
				URL:            "{BAD_URL_STRING",
				Crn:            core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			testService, testServiceErr := rangeapplicationsv1.NewRangeApplicationsV1(&rangeapplicationsv1.RangeApplicationsV1Options{
				URL:            "https://rangeapplicationsv1/api",
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
			testService, testServiceErr := rangeapplicationsv1.NewRangeApplicationsV1(&rangeapplicationsv1.RangeApplicationsV1Options{})
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
				"RANGE_APPLICATIONS_URL":       "https://rangeapplicationsv1/api",
				"RANGE_APPLICATIONS_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := rangeapplicationsv1.NewRangeApplicationsV1UsingExternalConfig(&rangeapplicationsv1.RangeApplicationsV1Options{
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := rangeapplicationsv1.NewRangeApplicationsV1UsingExternalConfig(&rangeapplicationsv1.RangeApplicationsV1Options{
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
				testService, testServiceErr := rangeapplicationsv1.NewRangeApplicationsV1UsingExternalConfig(&rangeapplicationsv1.RangeApplicationsV1Options{
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
				"RANGE_APPLICATIONS_URL":       "https://rangeapplicationsv1/api",
				"RANGE_APPLICATIONS_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := rangeapplicationsv1.NewRangeApplicationsV1UsingExternalConfig(&rangeapplicationsv1.RangeApplicationsV1Options{
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
				"RANGE_APPLICATIONS_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := rangeapplicationsv1.NewRangeApplicationsV1UsingExternalConfig(&rangeapplicationsv1.RangeApplicationsV1Options{
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
	Describe(`ListRangeApps(listRangeAppsOptions *ListRangeAppsOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		listRangeAppsPath := "/v1/testString/zones/testString/range/apps"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listRangeAppsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["page"]).To(Equal([]string{fmt.Sprint(int64(38))}))

					Expect(req.URL.Query()["per_page"]).To(Equal([]string{fmt.Sprint(int64(1))}))

					Expect(req.URL.Query()["order"]).To(Equal([]string{"protocol"}))

					Expect(req.URL.Query()["direction"]).To(Equal([]string{"asc"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListRangeApps with error: Operation response processing error`, func() {
				testService, testServiceErr := rangeapplicationsv1.NewRangeApplicationsV1(&rangeapplicationsv1.RangeApplicationsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ListRangeAppsOptions model
				listRangeAppsOptionsModel := new(rangeapplicationsv1.ListRangeAppsOptions)
				listRangeAppsOptionsModel.Page = core.Int64Ptr(int64(38))
				listRangeAppsOptionsModel.PerPage = core.Int64Ptr(int64(1))
				listRangeAppsOptionsModel.Order = core.StringPtr("protocol")
				listRangeAppsOptionsModel.Direction = core.StringPtr("asc")
				listRangeAppsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.ListRangeApps(listRangeAppsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListRangeApps(listRangeAppsOptions *ListRangeAppsOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		listRangeAppsPath := "/v1/testString/zones/testString/range/apps"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listRangeAppsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["page"]).To(Equal([]string{fmt.Sprint(int64(38))}))

					Expect(req.URL.Query()["per_page"]).To(Equal([]string{fmt.Sprint(int64(1))}))

					Expect(req.URL.Query()["order"]).To(Equal([]string{"protocol"}))

					Expect(req.URL.Query()["direction"]).To(Equal([]string{"asc"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": [{"id": "ea95132c15732412d22c1476fa83f27a", "protocol": "tcp/22", "dns": {"type": "CNAME", "name": "ssh.example.com"}, "origin_direct": ["OriginDirect"], "ip_firewall": true, "proxy_protocol": "v1", "edge_ips": {"type": "dynamic", "connectivity": "ipv4"}, "tls": "flexible", "traffic_type": "direct", "created_on": "2019-01-01T12:00:00", "modified_on": "2019-01-01T12:00:00"}]}`)
				}))
			})
			It(`Invoke ListRangeApps successfully`, func() {
				testService, testServiceErr := rangeapplicationsv1.NewRangeApplicationsV1(&rangeapplicationsv1.RangeApplicationsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListRangeApps(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListRangeAppsOptions model
				listRangeAppsOptionsModel := new(rangeapplicationsv1.ListRangeAppsOptions)
				listRangeAppsOptionsModel.Page = core.Int64Ptr(int64(38))
				listRangeAppsOptionsModel.PerPage = core.Int64Ptr(int64(1))
				listRangeAppsOptionsModel.Order = core.StringPtr("protocol")
				listRangeAppsOptionsModel.Direction = core.StringPtr("asc")
				listRangeAppsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListRangeApps(listRangeAppsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke ListRangeApps with error: Operation request error`, func() {
				testService, testServiceErr := rangeapplicationsv1.NewRangeApplicationsV1(&rangeapplicationsv1.RangeApplicationsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ListRangeAppsOptions model
				listRangeAppsOptionsModel := new(rangeapplicationsv1.ListRangeAppsOptions)
				listRangeAppsOptionsModel.Page = core.Int64Ptr(int64(38))
				listRangeAppsOptionsModel.PerPage = core.Int64Ptr(int64(1))
				listRangeAppsOptionsModel.Order = core.StringPtr("protocol")
				listRangeAppsOptionsModel.Direction = core.StringPtr("asc")
				listRangeAppsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.ListRangeApps(listRangeAppsOptionsModel)
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
	Describe(`CreateRangeApp(createRangeAppOptions *CreateRangeAppOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		createRangeAppPath := "/v1/testString/zones/testString/range/apps"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(createRangeAppPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateRangeApp with error: Operation response processing error`, func() {
				testService, testServiceErr := rangeapplicationsv1.NewRangeApplicationsV1(&rangeapplicationsv1.RangeApplicationsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the RangeAppReqDns model
				rangeAppReqDnsModel := new(rangeapplicationsv1.RangeAppReqDns)
				rangeAppReqDnsModel.Type = core.StringPtr("CNAME")
				rangeAppReqDnsModel.Name = core.StringPtr("ssh.example.com")

				// Construct an instance of the RangeAppReqEdgeIps model
				rangeAppReqEdgeIpsModel := new(rangeapplicationsv1.RangeAppReqEdgeIps)
				rangeAppReqEdgeIpsModel.Type = core.StringPtr("dynamic")
				rangeAppReqEdgeIpsModel.Connectivity = core.StringPtr("all")

				// Construct an instance of the RangeAppReqOriginDns model
				rangeAppReqOriginDnsModel := new(rangeapplicationsv1.RangeAppReqOriginDns)
				rangeAppReqOriginDnsModel.Name = core.StringPtr("origin.net")

				// Construct an instance of the CreateRangeAppOptions model
				createRangeAppOptionsModel := new(rangeapplicationsv1.CreateRangeAppOptions)
				createRangeAppOptionsModel.Protocol = core.StringPtr("tcp/22")
				createRangeAppOptionsModel.Dns = rangeAppReqDnsModel
				createRangeAppOptionsModel.OriginDirect = []string{"testString"}
				createRangeAppOptionsModel.OriginDns = rangeAppReqOriginDnsModel
				createRangeAppOptionsModel.OriginPort = core.Int64Ptr(int64(22))
				createRangeAppOptionsModel.IpFirewall = core.BoolPtr(true)
				createRangeAppOptionsModel.ProxyProtocol = core.StringPtr("off")
				createRangeAppOptionsModel.EdgeIps = rangeAppReqEdgeIpsModel
				createRangeAppOptionsModel.TrafficType = core.StringPtr("direct")
				createRangeAppOptionsModel.Tls = core.StringPtr("off")
				createRangeAppOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.CreateRangeApp(createRangeAppOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CreateRangeApp(createRangeAppOptions *CreateRangeAppOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		createRangeAppPath := "/v1/testString/zones/testString/range/apps"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(createRangeAppPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "ea95132c15732412d22c1476fa83f27a", "protocol": "tcp/22", "dns": {"type": "CNAME", "name": "ssh.example.com"}, "origin_direct": ["OriginDirect"], "ip_firewall": true, "proxy_protocol": "v1", "edge_ips": {"type": "dynamic", "connectivity": "ipv4"}, "tls": "flexible", "traffic_type": "direct", "created_on": "2019-01-01T12:00:00", "modified_on": "2019-01-01T12:00:00"}}`)
				}))
			})
			It(`Invoke CreateRangeApp successfully`, func() {
				testService, testServiceErr := rangeapplicationsv1.NewRangeApplicationsV1(&rangeapplicationsv1.RangeApplicationsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.CreateRangeApp(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the RangeAppReqDns model
				rangeAppReqDnsModel := new(rangeapplicationsv1.RangeAppReqDns)
				rangeAppReqDnsModel.Type = core.StringPtr("CNAME")
				rangeAppReqDnsModel.Name = core.StringPtr("ssh.example.com")

				// Construct an instance of the RangeAppReqEdgeIps model
				rangeAppReqEdgeIpsModel := new(rangeapplicationsv1.RangeAppReqEdgeIps)
				rangeAppReqEdgeIpsModel.Type = core.StringPtr("dynamic")
				rangeAppReqEdgeIpsModel.Connectivity = core.StringPtr("all")

				// Construct an instance of the RangeAppReqOriginDns model
				rangeAppReqOriginDnsModel := new(rangeapplicationsv1.RangeAppReqOriginDns)
				rangeAppReqOriginDnsModel.Name = core.StringPtr("origin.net")

				// Construct an instance of the CreateRangeAppOptions model
				createRangeAppOptionsModel := new(rangeapplicationsv1.CreateRangeAppOptions)
				createRangeAppOptionsModel.Protocol = core.StringPtr("tcp/22")
				createRangeAppOptionsModel.Dns = rangeAppReqDnsModel
				createRangeAppOptionsModel.OriginDirect = []string{"testString"}
				createRangeAppOptionsModel.OriginDns = rangeAppReqOriginDnsModel
				createRangeAppOptionsModel.OriginPort = core.Int64Ptr(int64(22))
				createRangeAppOptionsModel.IpFirewall = core.BoolPtr(true)
				createRangeAppOptionsModel.ProxyProtocol = core.StringPtr("off")
				createRangeAppOptionsModel.EdgeIps = rangeAppReqEdgeIpsModel
				createRangeAppOptionsModel.TrafficType = core.StringPtr("direct")
				createRangeAppOptionsModel.Tls = core.StringPtr("off")
				createRangeAppOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.CreateRangeApp(createRangeAppOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke CreateRangeApp with error: Operation validation and request error`, func() {
				testService, testServiceErr := rangeapplicationsv1.NewRangeApplicationsV1(&rangeapplicationsv1.RangeApplicationsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the RangeAppReqDns model
				rangeAppReqDnsModel := new(rangeapplicationsv1.RangeAppReqDns)
				rangeAppReqDnsModel.Type = core.StringPtr("CNAME")
				rangeAppReqDnsModel.Name = core.StringPtr("ssh.example.com")

				// Construct an instance of the RangeAppReqEdgeIps model
				rangeAppReqEdgeIpsModel := new(rangeapplicationsv1.RangeAppReqEdgeIps)
				rangeAppReqEdgeIpsModel.Type = core.StringPtr("dynamic")
				rangeAppReqEdgeIpsModel.Connectivity = core.StringPtr("all")

				// Construct an instance of the RangeAppReqOriginDns model
				rangeAppReqOriginDnsModel := new(rangeapplicationsv1.RangeAppReqOriginDns)
				rangeAppReqOriginDnsModel.Name = core.StringPtr("origin.net")

				// Construct an instance of the CreateRangeAppOptions model
				createRangeAppOptionsModel := new(rangeapplicationsv1.CreateRangeAppOptions)
				createRangeAppOptionsModel.Protocol = core.StringPtr("tcp/22")
				createRangeAppOptionsModel.Dns = rangeAppReqDnsModel
				createRangeAppOptionsModel.OriginDirect = []string{"testString"}
				createRangeAppOptionsModel.OriginDns = rangeAppReqOriginDnsModel
				createRangeAppOptionsModel.OriginPort = core.Int64Ptr(int64(22))
				createRangeAppOptionsModel.IpFirewall = core.BoolPtr(true)
				createRangeAppOptionsModel.ProxyProtocol = core.StringPtr("off")
				createRangeAppOptionsModel.EdgeIps = rangeAppReqEdgeIpsModel
				createRangeAppOptionsModel.TrafficType = core.StringPtr("direct")
				createRangeAppOptionsModel.Tls = core.StringPtr("off")
				createRangeAppOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.CreateRangeApp(createRangeAppOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateRangeAppOptions model with no property values
				createRangeAppOptionsModelNew := new(rangeapplicationsv1.CreateRangeAppOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.CreateRangeApp(createRangeAppOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetRangeApp(getRangeAppOptions *GetRangeAppOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getRangeAppPath := "/v1/testString/zones/testString/range/apps/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getRangeAppPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetRangeApp with error: Operation response processing error`, func() {
				testService, testServiceErr := rangeapplicationsv1.NewRangeApplicationsV1(&rangeapplicationsv1.RangeApplicationsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetRangeAppOptions model
				getRangeAppOptionsModel := new(rangeapplicationsv1.GetRangeAppOptions)
				getRangeAppOptionsModel.AppIdentifier = core.StringPtr("testString")
				getRangeAppOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.GetRangeApp(getRangeAppOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetRangeApp(getRangeAppOptions *GetRangeAppOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getRangeAppPath := "/v1/testString/zones/testString/range/apps/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getRangeAppPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "ea95132c15732412d22c1476fa83f27a", "protocol": "tcp/22", "dns": {"type": "CNAME", "name": "ssh.example.com"}, "origin_direct": ["OriginDirect"], "ip_firewall": true, "proxy_protocol": "v1", "edge_ips": {"type": "dynamic", "connectivity": "ipv4"}, "tls": "flexible", "traffic_type": "direct", "created_on": "2019-01-01T12:00:00", "modified_on": "2019-01-01T12:00:00"}}`)
				}))
			})
			It(`Invoke GetRangeApp successfully`, func() {
				testService, testServiceErr := rangeapplicationsv1.NewRangeApplicationsV1(&rangeapplicationsv1.RangeApplicationsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetRangeApp(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetRangeAppOptions model
				getRangeAppOptionsModel := new(rangeapplicationsv1.GetRangeAppOptions)
				getRangeAppOptionsModel.AppIdentifier = core.StringPtr("testString")
				getRangeAppOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetRangeApp(getRangeAppOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetRangeApp with error: Operation validation and request error`, func() {
				testService, testServiceErr := rangeapplicationsv1.NewRangeApplicationsV1(&rangeapplicationsv1.RangeApplicationsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetRangeAppOptions model
				getRangeAppOptionsModel := new(rangeapplicationsv1.GetRangeAppOptions)
				getRangeAppOptionsModel.AppIdentifier = core.StringPtr("testString")
				getRangeAppOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.GetRangeApp(getRangeAppOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetRangeAppOptions model with no property values
				getRangeAppOptionsModelNew := new(rangeapplicationsv1.GetRangeAppOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.GetRangeApp(getRangeAppOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateRangeApp(updateRangeAppOptions *UpdateRangeAppOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateRangeAppPath := "/v1/testString/zones/testString/range/apps/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(updateRangeAppPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateRangeApp with error: Operation response processing error`, func() {
				testService, testServiceErr := rangeapplicationsv1.NewRangeApplicationsV1(&rangeapplicationsv1.RangeApplicationsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the RangeAppReqDns model
				rangeAppReqDnsModel := new(rangeapplicationsv1.RangeAppReqDns)
				rangeAppReqDnsModel.Type = core.StringPtr("CNAME")
				rangeAppReqDnsModel.Name = core.StringPtr("ssh.example.com")

				// Construct an instance of the RangeAppReqEdgeIps model
				rangeAppReqEdgeIpsModel := new(rangeapplicationsv1.RangeAppReqEdgeIps)
				rangeAppReqEdgeIpsModel.Type = core.StringPtr("dynamic")
				rangeAppReqEdgeIpsModel.Connectivity = core.StringPtr("all")

				// Construct an instance of the RangeAppReqOriginDns model
				rangeAppReqOriginDnsModel := new(rangeapplicationsv1.RangeAppReqOriginDns)
				rangeAppReqOriginDnsModel.Name = core.StringPtr("origin.net")

				// Construct an instance of the UpdateRangeAppOptions model
				updateRangeAppOptionsModel := new(rangeapplicationsv1.UpdateRangeAppOptions)
				updateRangeAppOptionsModel.AppIdentifier = core.StringPtr("testString")
				updateRangeAppOptionsModel.Protocol = core.StringPtr("tcp/22")
				updateRangeAppOptionsModel.Dns = rangeAppReqDnsModel
				updateRangeAppOptionsModel.OriginDirect = []string{"testString"}
				updateRangeAppOptionsModel.OriginDns = rangeAppReqOriginDnsModel
				updateRangeAppOptionsModel.OriginPort = core.Int64Ptr(int64(22))
				updateRangeAppOptionsModel.IpFirewall = core.BoolPtr(true)
				updateRangeAppOptionsModel.ProxyProtocol = core.StringPtr("off")
				updateRangeAppOptionsModel.EdgeIps = rangeAppReqEdgeIpsModel
				updateRangeAppOptionsModel.TrafficType = core.StringPtr("direct")
				updateRangeAppOptionsModel.Tls = core.StringPtr("off")
				updateRangeAppOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.UpdateRangeApp(updateRangeAppOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateRangeApp(updateRangeAppOptions *UpdateRangeAppOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateRangeAppPath := "/v1/testString/zones/testString/range/apps/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(updateRangeAppPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "ea95132c15732412d22c1476fa83f27a", "protocol": "tcp/22", "dns": {"type": "CNAME", "name": "ssh.example.com"}, "origin_direct": ["OriginDirect"], "ip_firewall": true, "proxy_protocol": "v1", "edge_ips": {"type": "dynamic", "connectivity": "ipv4"}, "tls": "flexible", "traffic_type": "direct", "created_on": "2019-01-01T12:00:00", "modified_on": "2019-01-01T12:00:00"}}`)
				}))
			})
			It(`Invoke UpdateRangeApp successfully`, func() {
				testService, testServiceErr := rangeapplicationsv1.NewRangeApplicationsV1(&rangeapplicationsv1.RangeApplicationsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.UpdateRangeApp(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the RangeAppReqDns model
				rangeAppReqDnsModel := new(rangeapplicationsv1.RangeAppReqDns)
				rangeAppReqDnsModel.Type = core.StringPtr("CNAME")
				rangeAppReqDnsModel.Name = core.StringPtr("ssh.example.com")

				// Construct an instance of the RangeAppReqEdgeIps model
				rangeAppReqEdgeIpsModel := new(rangeapplicationsv1.RangeAppReqEdgeIps)
				rangeAppReqEdgeIpsModel.Type = core.StringPtr("dynamic")
				rangeAppReqEdgeIpsModel.Connectivity = core.StringPtr("all")

				// Construct an instance of the RangeAppReqOriginDns model
				rangeAppReqOriginDnsModel := new(rangeapplicationsv1.RangeAppReqOriginDns)
				rangeAppReqOriginDnsModel.Name = core.StringPtr("origin.net")

				// Construct an instance of the UpdateRangeAppOptions model
				updateRangeAppOptionsModel := new(rangeapplicationsv1.UpdateRangeAppOptions)
				updateRangeAppOptionsModel.AppIdentifier = core.StringPtr("testString")
				updateRangeAppOptionsModel.Protocol = core.StringPtr("tcp/22")
				updateRangeAppOptionsModel.Dns = rangeAppReqDnsModel
				updateRangeAppOptionsModel.OriginDirect = []string{"testString"}
				updateRangeAppOptionsModel.OriginDns = rangeAppReqOriginDnsModel
				updateRangeAppOptionsModel.OriginPort = core.Int64Ptr(int64(22))
				updateRangeAppOptionsModel.IpFirewall = core.BoolPtr(true)
				updateRangeAppOptionsModel.ProxyProtocol = core.StringPtr("off")
				updateRangeAppOptionsModel.EdgeIps = rangeAppReqEdgeIpsModel
				updateRangeAppOptionsModel.TrafficType = core.StringPtr("direct")
				updateRangeAppOptionsModel.Tls = core.StringPtr("off")
				updateRangeAppOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.UpdateRangeApp(updateRangeAppOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke UpdateRangeApp with error: Operation validation and request error`, func() {
				testService, testServiceErr := rangeapplicationsv1.NewRangeApplicationsV1(&rangeapplicationsv1.RangeApplicationsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the RangeAppReqDns model
				rangeAppReqDnsModel := new(rangeapplicationsv1.RangeAppReqDns)
				rangeAppReqDnsModel.Type = core.StringPtr("CNAME")
				rangeAppReqDnsModel.Name = core.StringPtr("ssh.example.com")

				// Construct an instance of the RangeAppReqEdgeIps model
				rangeAppReqEdgeIpsModel := new(rangeapplicationsv1.RangeAppReqEdgeIps)
				rangeAppReqEdgeIpsModel.Type = core.StringPtr("dynamic")
				rangeAppReqEdgeIpsModel.Connectivity = core.StringPtr("all")

				// Construct an instance of the RangeAppReqOriginDns model
				rangeAppReqOriginDnsModel := new(rangeapplicationsv1.RangeAppReqOriginDns)
				rangeAppReqOriginDnsModel.Name = core.StringPtr("origin.net")

				// Construct an instance of the UpdateRangeAppOptions model
				updateRangeAppOptionsModel := new(rangeapplicationsv1.UpdateRangeAppOptions)
				updateRangeAppOptionsModel.AppIdentifier = core.StringPtr("testString")
				updateRangeAppOptionsModel.Protocol = core.StringPtr("tcp/22")
				updateRangeAppOptionsModel.Dns = rangeAppReqDnsModel
				updateRangeAppOptionsModel.OriginDirect = []string{"testString"}
				updateRangeAppOptionsModel.OriginDns = rangeAppReqOriginDnsModel
				updateRangeAppOptionsModel.OriginPort = core.Int64Ptr(int64(22))
				updateRangeAppOptionsModel.IpFirewall = core.BoolPtr(true)
				updateRangeAppOptionsModel.ProxyProtocol = core.StringPtr("off")
				updateRangeAppOptionsModel.EdgeIps = rangeAppReqEdgeIpsModel
				updateRangeAppOptionsModel.TrafficType = core.StringPtr("direct")
				updateRangeAppOptionsModel.Tls = core.StringPtr("off")
				updateRangeAppOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.UpdateRangeApp(updateRangeAppOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateRangeAppOptions model with no property values
				updateRangeAppOptionsModelNew := new(rangeapplicationsv1.UpdateRangeAppOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.UpdateRangeApp(updateRangeAppOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteRangeApp(deleteRangeAppOptions *DeleteRangeAppOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		deleteRangeAppPath := "/v1/testString/zones/testString/range/apps/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(deleteRangeAppPath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteRangeApp with error: Operation response processing error`, func() {
				testService, testServiceErr := rangeapplicationsv1.NewRangeApplicationsV1(&rangeapplicationsv1.RangeApplicationsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the DeleteRangeAppOptions model
				deleteRangeAppOptionsModel := new(rangeapplicationsv1.DeleteRangeAppOptions)
				deleteRangeAppOptionsModel.AppIdentifier = core.StringPtr("testString")
				deleteRangeAppOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.DeleteRangeApp(deleteRangeAppOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteRangeApp(deleteRangeAppOptions *DeleteRangeAppOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		deleteRangeAppPath := "/v1/testString/zones/testString/range/apps/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(deleteRangeAppPath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "ea95132c15732412d22c1476fa83f27a", "protocol": "tcp/22", "dns": {"type": "CNAME", "name": "ssh.example.com"}, "origin_direct": ["OriginDirect"], "ip_firewall": true, "proxy_protocol": "v1", "edge_ips": {"type": "dynamic", "connectivity": "ipv4"}, "tls": "flexible", "traffic_type": "direct", "created_on": "2019-01-01T12:00:00", "modified_on": "2019-01-01T12:00:00"}}`)
				}))
			})
			It(`Invoke DeleteRangeApp successfully`, func() {
				testService, testServiceErr := rangeapplicationsv1.NewRangeApplicationsV1(&rangeapplicationsv1.RangeApplicationsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.DeleteRangeApp(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteRangeAppOptions model
				deleteRangeAppOptionsModel := new(rangeapplicationsv1.DeleteRangeAppOptions)
				deleteRangeAppOptionsModel.AppIdentifier = core.StringPtr("testString")
				deleteRangeAppOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.DeleteRangeApp(deleteRangeAppOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke DeleteRangeApp with error: Operation validation and request error`, func() {
				testService, testServiceErr := rangeapplicationsv1.NewRangeApplicationsV1(&rangeapplicationsv1.RangeApplicationsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the DeleteRangeAppOptions model
				deleteRangeAppOptionsModel := new(rangeapplicationsv1.DeleteRangeAppOptions)
				deleteRangeAppOptionsModel.AppIdentifier = core.StringPtr("testString")
				deleteRangeAppOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.DeleteRangeApp(deleteRangeAppOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteRangeAppOptions model with no property values
				deleteRangeAppOptionsModelNew := new(rangeapplicationsv1.DeleteRangeAppOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.DeleteRangeApp(deleteRangeAppOptionsModelNew)
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
			testService, _ := rangeapplicationsv1.NewRangeApplicationsV1(&rangeapplicationsv1.RangeApplicationsV1Options{
				URL:            "http://rangeapplicationsv1modelgenerator.com",
				Authenticator:  &core.NoAuthAuthenticator{},
				Crn:            core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})
			It(`Invoke NewCreateRangeAppOptions successfully`, func() {
				// Construct an instance of the RangeAppReqDns model
				rangeAppReqDnsModel := new(rangeapplicationsv1.RangeAppReqDns)
				Expect(rangeAppReqDnsModel).ToNot(BeNil())
				rangeAppReqDnsModel.Type = core.StringPtr("CNAME")
				rangeAppReqDnsModel.Name = core.StringPtr("ssh.example.com")
				Expect(rangeAppReqDnsModel.Type).To(Equal(core.StringPtr("CNAME")))
				Expect(rangeAppReqDnsModel.Name).To(Equal(core.StringPtr("ssh.example.com")))

				// Construct an instance of the RangeAppReqEdgeIps model
				rangeAppReqEdgeIpsModel := new(rangeapplicationsv1.RangeAppReqEdgeIps)
				Expect(rangeAppReqEdgeIpsModel).ToNot(BeNil())
				rangeAppReqEdgeIpsModel.Type = core.StringPtr("dynamic")
				rangeAppReqEdgeIpsModel.Connectivity = core.StringPtr("all")
				Expect(rangeAppReqEdgeIpsModel.Type).To(Equal(core.StringPtr("dynamic")))
				Expect(rangeAppReqEdgeIpsModel.Connectivity).To(Equal(core.StringPtr("all")))

				// Construct an instance of the RangeAppReqOriginDns model
				rangeAppReqOriginDnsModel := new(rangeapplicationsv1.RangeAppReqOriginDns)
				Expect(rangeAppReqOriginDnsModel).ToNot(BeNil())
				rangeAppReqOriginDnsModel.Name = core.StringPtr("origin.net")
				Expect(rangeAppReqOriginDnsModel.Name).To(Equal(core.StringPtr("origin.net")))

				// Construct an instance of the CreateRangeAppOptions model
				createRangeAppOptionsProtocol := "tcp/22"
				var createRangeAppOptionsDns *rangeapplicationsv1.RangeAppReqDns = nil
				createRangeAppOptionsModel := testService.NewCreateRangeAppOptions(createRangeAppOptionsProtocol, createRangeAppOptionsDns)
				createRangeAppOptionsModel.SetProtocol("tcp/22")
				createRangeAppOptionsModel.SetDns(rangeAppReqDnsModel)
				createRangeAppOptionsModel.SetOriginDirect([]string{"testString"})
				createRangeAppOptionsModel.SetOriginDns(rangeAppReqOriginDnsModel)
				createRangeAppOptionsModel.SetOriginPort(int64(22))
				createRangeAppOptionsModel.SetIpFirewall(true)
				createRangeAppOptionsModel.SetProxyProtocol("off")
				createRangeAppOptionsModel.SetEdgeIps(rangeAppReqEdgeIpsModel)
				createRangeAppOptionsModel.SetTrafficType("direct")
				createRangeAppOptionsModel.SetTls("off")
				createRangeAppOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createRangeAppOptionsModel).ToNot(BeNil())
				Expect(createRangeAppOptionsModel.Protocol).To(Equal(core.StringPtr("tcp/22")))
				Expect(createRangeAppOptionsModel.Dns).To(Equal(rangeAppReqDnsModel))
				Expect(createRangeAppOptionsModel.OriginDirect).To(Equal([]string{"testString"}))
				Expect(createRangeAppOptionsModel.OriginDns).To(Equal(rangeAppReqOriginDnsModel))
				Expect(createRangeAppOptionsModel.OriginPort).To(Equal(core.Int64Ptr(int64(22))))
				Expect(createRangeAppOptionsModel.IpFirewall).To(Equal(core.BoolPtr(true)))
				Expect(createRangeAppOptionsModel.ProxyProtocol).To(Equal(core.StringPtr("off")))
				Expect(createRangeAppOptionsModel.EdgeIps).To(Equal(rangeAppReqEdgeIpsModel))
				Expect(createRangeAppOptionsModel.TrafficType).To(Equal(core.StringPtr("direct")))
				Expect(createRangeAppOptionsModel.Tls).To(Equal(core.StringPtr("off")))
				Expect(createRangeAppOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteRangeAppOptions successfully`, func() {
				// Construct an instance of the DeleteRangeAppOptions model
				appIdentifier := "testString"
				deleteRangeAppOptionsModel := testService.NewDeleteRangeAppOptions(appIdentifier)
				deleteRangeAppOptionsModel.SetAppIdentifier("testString")
				deleteRangeAppOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteRangeAppOptionsModel).ToNot(BeNil())
				Expect(deleteRangeAppOptionsModel.AppIdentifier).To(Equal(core.StringPtr("testString")))
				Expect(deleteRangeAppOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetRangeAppOptions successfully`, func() {
				// Construct an instance of the GetRangeAppOptions model
				appIdentifier := "testString"
				getRangeAppOptionsModel := testService.NewGetRangeAppOptions(appIdentifier)
				getRangeAppOptionsModel.SetAppIdentifier("testString")
				getRangeAppOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getRangeAppOptionsModel).ToNot(BeNil())
				Expect(getRangeAppOptionsModel.AppIdentifier).To(Equal(core.StringPtr("testString")))
				Expect(getRangeAppOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListRangeAppsOptions successfully`, func() {
				// Construct an instance of the ListRangeAppsOptions model
				listRangeAppsOptionsModel := testService.NewListRangeAppsOptions()
				listRangeAppsOptionsModel.SetPage(int64(38))
				listRangeAppsOptionsModel.SetPerPage(int64(1))
				listRangeAppsOptionsModel.SetOrder("protocol")
				listRangeAppsOptionsModel.SetDirection("asc")
				listRangeAppsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listRangeAppsOptionsModel).ToNot(BeNil())
				Expect(listRangeAppsOptionsModel.Page).To(Equal(core.Int64Ptr(int64(38))))
				Expect(listRangeAppsOptionsModel.PerPage).To(Equal(core.Int64Ptr(int64(1))))
				Expect(listRangeAppsOptionsModel.Order).To(Equal(core.StringPtr("protocol")))
				Expect(listRangeAppsOptionsModel.Direction).To(Equal(core.StringPtr("asc")))
				Expect(listRangeAppsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewRangeAppReqOriginDns successfully`, func() {
				name := "origin.net"
				model, err := testService.NewRangeAppReqOriginDns(name)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewUpdateRangeAppOptions successfully`, func() {
				// Construct an instance of the RangeAppReqDns model
				rangeAppReqDnsModel := new(rangeapplicationsv1.RangeAppReqDns)
				Expect(rangeAppReqDnsModel).ToNot(BeNil())
				rangeAppReqDnsModel.Type = core.StringPtr("CNAME")
				rangeAppReqDnsModel.Name = core.StringPtr("ssh.example.com")
				Expect(rangeAppReqDnsModel.Type).To(Equal(core.StringPtr("CNAME")))
				Expect(rangeAppReqDnsModel.Name).To(Equal(core.StringPtr("ssh.example.com")))

				// Construct an instance of the RangeAppReqEdgeIps model
				rangeAppReqEdgeIpsModel := new(rangeapplicationsv1.RangeAppReqEdgeIps)
				Expect(rangeAppReqEdgeIpsModel).ToNot(BeNil())
				rangeAppReqEdgeIpsModel.Type = core.StringPtr("dynamic")
				rangeAppReqEdgeIpsModel.Connectivity = core.StringPtr("all")
				Expect(rangeAppReqEdgeIpsModel.Type).To(Equal(core.StringPtr("dynamic")))
				Expect(rangeAppReqEdgeIpsModel.Connectivity).To(Equal(core.StringPtr("all")))

				// Construct an instance of the RangeAppReqOriginDns model
				rangeAppReqOriginDnsModel := new(rangeapplicationsv1.RangeAppReqOriginDns)
				Expect(rangeAppReqOriginDnsModel).ToNot(BeNil())
				rangeAppReqOriginDnsModel.Name = core.StringPtr("origin.net")
				Expect(rangeAppReqOriginDnsModel.Name).To(Equal(core.StringPtr("origin.net")))

				// Construct an instance of the UpdateRangeAppOptions model
				appIdentifier := "testString"
				updateRangeAppOptionsProtocol := "tcp/22"
				var updateRangeAppOptionsDns *rangeapplicationsv1.RangeAppReqDns = nil
				updateRangeAppOptionsModel := testService.NewUpdateRangeAppOptions(appIdentifier, updateRangeAppOptionsProtocol, updateRangeAppOptionsDns)
				updateRangeAppOptionsModel.SetAppIdentifier("testString")
				updateRangeAppOptionsModel.SetProtocol("tcp/22")
				updateRangeAppOptionsModel.SetDns(rangeAppReqDnsModel)
				updateRangeAppOptionsModel.SetOriginDirect([]string{"testString"})
				updateRangeAppOptionsModel.SetOriginDns(rangeAppReqOriginDnsModel)
				updateRangeAppOptionsModel.SetOriginPort(int64(22))
				updateRangeAppOptionsModel.SetIpFirewall(true)
				updateRangeAppOptionsModel.SetProxyProtocol("off")
				updateRangeAppOptionsModel.SetEdgeIps(rangeAppReqEdgeIpsModel)
				updateRangeAppOptionsModel.SetTrafficType("direct")
				updateRangeAppOptionsModel.SetTls("off")
				updateRangeAppOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateRangeAppOptionsModel).ToNot(BeNil())
				Expect(updateRangeAppOptionsModel.AppIdentifier).To(Equal(core.StringPtr("testString")))
				Expect(updateRangeAppOptionsModel.Protocol).To(Equal(core.StringPtr("tcp/22")))
				Expect(updateRangeAppOptionsModel.Dns).To(Equal(rangeAppReqDnsModel))
				Expect(updateRangeAppOptionsModel.OriginDirect).To(Equal([]string{"testString"}))
				Expect(updateRangeAppOptionsModel.OriginDns).To(Equal(rangeAppReqOriginDnsModel))
				Expect(updateRangeAppOptionsModel.OriginPort).To(Equal(core.Int64Ptr(int64(22))))
				Expect(updateRangeAppOptionsModel.IpFirewall).To(Equal(core.BoolPtr(true)))
				Expect(updateRangeAppOptionsModel.ProxyProtocol).To(Equal(core.StringPtr("off")))
				Expect(updateRangeAppOptionsModel.EdgeIps).To(Equal(rangeAppReqEdgeIpsModel))
				Expect(updateRangeAppOptionsModel.TrafficType).To(Equal(core.StringPtr("direct")))
				Expect(updateRangeAppOptionsModel.Tls).To(Equal(core.StringPtr("off")))
				Expect(updateRangeAppOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
