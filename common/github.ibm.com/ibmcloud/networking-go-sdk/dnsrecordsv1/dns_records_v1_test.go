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

package dnsrecordsv1_test

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.ibm.com/ibmcloud/networking-go-sdk/dnsrecordsv1"
)

var _ = Describe(`DnsRecordsV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		It(`Instantiate service client`, func() {
			testService, testServiceErr := dnsrecordsv1.NewDnsRecordsV1(&dnsrecordsv1.DnsRecordsV1Options{
				Authenticator:  &core.NoAuthAuthenticator{},
				Crn:            core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})
			Expect(testService).ToNot(BeNil())
			Expect(testServiceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			testService, testServiceErr := dnsrecordsv1.NewDnsRecordsV1(&dnsrecordsv1.DnsRecordsV1Options{
				URL:            "{BAD_URL_STRING",
				Crn:            core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			testService, testServiceErr := dnsrecordsv1.NewDnsRecordsV1(&dnsrecordsv1.DnsRecordsV1Options{
				URL:            "https://dnsrecordsv1/api",
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
			testService, testServiceErr := dnsrecordsv1.NewDnsRecordsV1(&dnsrecordsv1.DnsRecordsV1Options{})
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
				"DNS_RECORDS_URL":       "https://dnsrecordsv1/api",
				"DNS_RECORDS_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := dnsrecordsv1.NewDnsRecordsV1UsingExternalConfig(&dnsrecordsv1.DnsRecordsV1Options{
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := dnsrecordsv1.NewDnsRecordsV1UsingExternalConfig(&dnsrecordsv1.DnsRecordsV1Options{
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
				testService, testServiceErr := dnsrecordsv1.NewDnsRecordsV1UsingExternalConfig(&dnsrecordsv1.DnsRecordsV1Options{
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
				"DNS_RECORDS_URL":       "https://dnsrecordsv1/api",
				"DNS_RECORDS_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := dnsrecordsv1.NewDnsRecordsV1UsingExternalConfig(&dnsrecordsv1.DnsRecordsV1Options{
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
				"DNS_RECORDS_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := dnsrecordsv1.NewDnsRecordsV1UsingExternalConfig(&dnsrecordsv1.DnsRecordsV1Options{
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
	Describe(`ListAllDnsRecords(listAllDnsRecordsOptions *ListAllDnsRecordsOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		listAllDnsRecordsPath := "/v1/testString/zones/testString/dns_records"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listAllDnsRecordsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["type"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["name"]).To(Equal([]string{"host1.test-example.com"}))

					Expect(req.URL.Query()["content"]).To(Equal([]string{"1.2.3.4"}))

					Expect(req.URL.Query()["page"]).To(Equal([]string{fmt.Sprint(int64(38))}))

					Expect(req.URL.Query()["per_page"]).To(Equal([]string{fmt.Sprint(int64(5))}))

					Expect(req.URL.Query()["order"]).To(Equal([]string{"type"}))

					Expect(req.URL.Query()["direction"]).To(Equal([]string{"asc"}))

					Expect(req.URL.Query()["match"]).To(Equal([]string{"any"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListAllDnsRecords with error: Operation response processing error`, func() {
				testService, testServiceErr := dnsrecordsv1.NewDnsRecordsV1(&dnsrecordsv1.DnsRecordsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ListAllDnsRecordsOptions model
				listAllDnsRecordsOptionsModel := new(dnsrecordsv1.ListAllDnsRecordsOptions)
				listAllDnsRecordsOptionsModel.Type = core.StringPtr("testString")
				listAllDnsRecordsOptionsModel.Name = core.StringPtr("host1.test-example.com")
				listAllDnsRecordsOptionsModel.Content = core.StringPtr("1.2.3.4")
				listAllDnsRecordsOptionsModel.Page = core.Int64Ptr(int64(38))
				listAllDnsRecordsOptionsModel.PerPage = core.Int64Ptr(int64(5))
				listAllDnsRecordsOptionsModel.Order = core.StringPtr("type")
				listAllDnsRecordsOptionsModel.Direction = core.StringPtr("asc")
				listAllDnsRecordsOptionsModel.Match = core.StringPtr("any")
				listAllDnsRecordsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.ListAllDnsRecords(listAllDnsRecordsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListAllDnsRecords(listAllDnsRecordsOptions *ListAllDnsRecordsOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		listAllDnsRecordsPath := "/v1/testString/zones/testString/dns_records"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listAllDnsRecordsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["type"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["name"]).To(Equal([]string{"host1.test-example.com"}))

					Expect(req.URL.Query()["content"]).To(Equal([]string{"1.2.3.4"}))

					Expect(req.URL.Query()["page"]).To(Equal([]string{fmt.Sprint(int64(38))}))

					Expect(req.URL.Query()["per_page"]).To(Equal([]string{fmt.Sprint(int64(5))}))

					Expect(req.URL.Query()["order"]).To(Equal([]string{"type"}))

					Expect(req.URL.Query()["direction"]).To(Equal([]string{"asc"}))

					Expect(req.URL.Query()["match"]).To(Equal([]string{"any"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": [{"id": "f1aba936b94213e5b8dca0c0dbf1f9cc", "created_on": "2014-01-01T05:20:00.12345Z", "modified_on": "2014-01-01T05:20:00.12345Z", "name": "host-1.test-example.com", "type": "A", "content": "169.154.10.10", "zone_id": "023e105f4ecef8ad9ca31a8372d0c353", "zone_name": "test-example.com", "proxiable": true, "proxied": false, "ttl": 120, "priority": 5, "data": {"anyKey": "anyValue"}}], "result_info": {"page": 1, "per_page": 20, "count": 1, "total_count": 2000}}`)
				}))
			})
			It(`Invoke ListAllDnsRecords successfully`, func() {
				testService, testServiceErr := dnsrecordsv1.NewDnsRecordsV1(&dnsrecordsv1.DnsRecordsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListAllDnsRecords(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListAllDnsRecordsOptions model
				listAllDnsRecordsOptionsModel := new(dnsrecordsv1.ListAllDnsRecordsOptions)
				listAllDnsRecordsOptionsModel.Type = core.StringPtr("testString")
				listAllDnsRecordsOptionsModel.Name = core.StringPtr("host1.test-example.com")
				listAllDnsRecordsOptionsModel.Content = core.StringPtr("1.2.3.4")
				listAllDnsRecordsOptionsModel.Page = core.Int64Ptr(int64(38))
				listAllDnsRecordsOptionsModel.PerPage = core.Int64Ptr(int64(5))
				listAllDnsRecordsOptionsModel.Order = core.StringPtr("type")
				listAllDnsRecordsOptionsModel.Direction = core.StringPtr("asc")
				listAllDnsRecordsOptionsModel.Match = core.StringPtr("any")
				listAllDnsRecordsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListAllDnsRecords(listAllDnsRecordsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke ListAllDnsRecords with error: Operation request error`, func() {
				testService, testServiceErr := dnsrecordsv1.NewDnsRecordsV1(&dnsrecordsv1.DnsRecordsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ListAllDnsRecordsOptions model
				listAllDnsRecordsOptionsModel := new(dnsrecordsv1.ListAllDnsRecordsOptions)
				listAllDnsRecordsOptionsModel.Type = core.StringPtr("testString")
				listAllDnsRecordsOptionsModel.Name = core.StringPtr("host1.test-example.com")
				listAllDnsRecordsOptionsModel.Content = core.StringPtr("1.2.3.4")
				listAllDnsRecordsOptionsModel.Page = core.Int64Ptr(int64(38))
				listAllDnsRecordsOptionsModel.PerPage = core.Int64Ptr(int64(5))
				listAllDnsRecordsOptionsModel.Order = core.StringPtr("type")
				listAllDnsRecordsOptionsModel.Direction = core.StringPtr("asc")
				listAllDnsRecordsOptionsModel.Match = core.StringPtr("any")
				listAllDnsRecordsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.ListAllDnsRecords(listAllDnsRecordsOptionsModel)
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
	Describe(`CreateDnsRecord(createDnsRecordOptions *CreateDnsRecordOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		createDnsRecordPath := "/v1/testString/zones/testString/dns_records"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(createDnsRecordPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateDnsRecord with error: Operation response processing error`, func() {
				testService, testServiceErr := dnsrecordsv1.NewDnsRecordsV1(&dnsrecordsv1.DnsRecordsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the CreateDnsRecordOptions model
				createDnsRecordOptionsModel := new(dnsrecordsv1.CreateDnsRecordOptions)
				createDnsRecordOptionsModel.Name = core.StringPtr("host-1.test-example.com")
				createDnsRecordOptionsModel.Type = core.StringPtr("A")
				createDnsRecordOptionsModel.Content = core.StringPtr("1.2.3.4")
				createDnsRecordOptionsModel.Priority = core.Int64Ptr(int64(5))
				createDnsRecordOptionsModel.Data = map[string]interface{}{"anyKey": "anyValue"}
				createDnsRecordOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.CreateDnsRecord(createDnsRecordOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CreateDnsRecord(createDnsRecordOptions *CreateDnsRecordOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		createDnsRecordPath := "/v1/testString/zones/testString/dns_records"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(createDnsRecordPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "f1aba936b94213e5b8dca0c0dbf1f9cc", "created_on": "2014-01-01T05:20:00.12345Z", "modified_on": "2014-01-01T05:20:00.12345Z", "name": "host-1.test-example.com", "type": "A", "content": "169.154.10.10", "zone_id": "023e105f4ecef8ad9ca31a8372d0c353", "zone_name": "test-example.com", "proxiable": true, "proxied": false, "ttl": 120, "priority": 5, "data": {"anyKey": "anyValue"}}}`)
				}))
			})
			It(`Invoke CreateDnsRecord successfully`, func() {
				testService, testServiceErr := dnsrecordsv1.NewDnsRecordsV1(&dnsrecordsv1.DnsRecordsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.CreateDnsRecord(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateDnsRecordOptions model
				createDnsRecordOptionsModel := new(dnsrecordsv1.CreateDnsRecordOptions)
				createDnsRecordOptionsModel.Name = core.StringPtr("host-1.test-example.com")
				createDnsRecordOptionsModel.Type = core.StringPtr("A")
				createDnsRecordOptionsModel.Content = core.StringPtr("1.2.3.4")
				createDnsRecordOptionsModel.Priority = core.Int64Ptr(int64(5))
				createDnsRecordOptionsModel.Data = map[string]interface{}{"anyKey": "anyValue"}
				createDnsRecordOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.CreateDnsRecord(createDnsRecordOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke CreateDnsRecord with error: Operation request error`, func() {
				testService, testServiceErr := dnsrecordsv1.NewDnsRecordsV1(&dnsrecordsv1.DnsRecordsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the CreateDnsRecordOptions model
				createDnsRecordOptionsModel := new(dnsrecordsv1.CreateDnsRecordOptions)
				createDnsRecordOptionsModel.Name = core.StringPtr("host-1.test-example.com")
				createDnsRecordOptionsModel.Type = core.StringPtr("A")
				createDnsRecordOptionsModel.Content = core.StringPtr("1.2.3.4")
				createDnsRecordOptionsModel.Priority = core.Int64Ptr(int64(5))
				createDnsRecordOptionsModel.Data = map[string]interface{}{"anyKey": "anyValue"}
				createDnsRecordOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.CreateDnsRecord(createDnsRecordOptionsModel)
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
	Describe(`DeleteDnsRecord(deleteDnsRecordOptions *DeleteDnsRecordOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		deleteDnsRecordPath := "/v1/testString/zones/testString/dns_records/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(deleteDnsRecordPath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteDnsRecord with error: Operation response processing error`, func() {
				testService, testServiceErr := dnsrecordsv1.NewDnsRecordsV1(&dnsrecordsv1.DnsRecordsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the DeleteDnsRecordOptions model
				deleteDnsRecordOptionsModel := new(dnsrecordsv1.DeleteDnsRecordOptions)
				deleteDnsRecordOptionsModel.DnsrecordIdentifier = core.StringPtr("testString")
				deleteDnsRecordOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.DeleteDnsRecord(deleteDnsRecordOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteDnsRecord(deleteDnsRecordOptions *DeleteDnsRecordOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		deleteDnsRecordPath := "/v1/testString/zones/testString/dns_records/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(deleteDnsRecordPath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "f1aba936b94213e5b8dca0c0dbf1f9cc"}}`)
				}))
			})
			It(`Invoke DeleteDnsRecord successfully`, func() {
				testService, testServiceErr := dnsrecordsv1.NewDnsRecordsV1(&dnsrecordsv1.DnsRecordsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.DeleteDnsRecord(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteDnsRecordOptions model
				deleteDnsRecordOptionsModel := new(dnsrecordsv1.DeleteDnsRecordOptions)
				deleteDnsRecordOptionsModel.DnsrecordIdentifier = core.StringPtr("testString")
				deleteDnsRecordOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.DeleteDnsRecord(deleteDnsRecordOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke DeleteDnsRecord with error: Operation validation and request error`, func() {
				testService, testServiceErr := dnsrecordsv1.NewDnsRecordsV1(&dnsrecordsv1.DnsRecordsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the DeleteDnsRecordOptions model
				deleteDnsRecordOptionsModel := new(dnsrecordsv1.DeleteDnsRecordOptions)
				deleteDnsRecordOptionsModel.DnsrecordIdentifier = core.StringPtr("testString")
				deleteDnsRecordOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.DeleteDnsRecord(deleteDnsRecordOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteDnsRecordOptions model with no property values
				deleteDnsRecordOptionsModelNew := new(dnsrecordsv1.DeleteDnsRecordOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.DeleteDnsRecord(deleteDnsRecordOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetDnsRecord(getDnsRecordOptions *GetDnsRecordOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getDnsRecordPath := "/v1/testString/zones/testString/dns_records/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getDnsRecordPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetDnsRecord with error: Operation response processing error`, func() {
				testService, testServiceErr := dnsrecordsv1.NewDnsRecordsV1(&dnsrecordsv1.DnsRecordsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetDnsRecordOptions model
				getDnsRecordOptionsModel := new(dnsrecordsv1.GetDnsRecordOptions)
				getDnsRecordOptionsModel.DnsrecordIdentifier = core.StringPtr("testString")
				getDnsRecordOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.GetDnsRecord(getDnsRecordOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetDnsRecord(getDnsRecordOptions *GetDnsRecordOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getDnsRecordPath := "/v1/testString/zones/testString/dns_records/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getDnsRecordPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "f1aba936b94213e5b8dca0c0dbf1f9cc", "created_on": "2014-01-01T05:20:00.12345Z", "modified_on": "2014-01-01T05:20:00.12345Z", "name": "host-1.test-example.com", "type": "A", "content": "169.154.10.10", "zone_id": "023e105f4ecef8ad9ca31a8372d0c353", "zone_name": "test-example.com", "proxiable": true, "proxied": false, "ttl": 120, "priority": 5, "data": {"anyKey": "anyValue"}}}`)
				}))
			})
			It(`Invoke GetDnsRecord successfully`, func() {
				testService, testServiceErr := dnsrecordsv1.NewDnsRecordsV1(&dnsrecordsv1.DnsRecordsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetDnsRecord(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetDnsRecordOptions model
				getDnsRecordOptionsModel := new(dnsrecordsv1.GetDnsRecordOptions)
				getDnsRecordOptionsModel.DnsrecordIdentifier = core.StringPtr("testString")
				getDnsRecordOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetDnsRecord(getDnsRecordOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetDnsRecord with error: Operation validation and request error`, func() {
				testService, testServiceErr := dnsrecordsv1.NewDnsRecordsV1(&dnsrecordsv1.DnsRecordsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetDnsRecordOptions model
				getDnsRecordOptionsModel := new(dnsrecordsv1.GetDnsRecordOptions)
				getDnsRecordOptionsModel.DnsrecordIdentifier = core.StringPtr("testString")
				getDnsRecordOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.GetDnsRecord(getDnsRecordOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetDnsRecordOptions model with no property values
				getDnsRecordOptionsModelNew := new(dnsrecordsv1.GetDnsRecordOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.GetDnsRecord(getDnsRecordOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateDnsRecord(updateDnsRecordOptions *UpdateDnsRecordOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateDnsRecordPath := "/v1/testString/zones/testString/dns_records/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(updateDnsRecordPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateDnsRecord with error: Operation response processing error`, func() {
				testService, testServiceErr := dnsrecordsv1.NewDnsRecordsV1(&dnsrecordsv1.DnsRecordsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the UpdateDnsRecordOptions model
				updateDnsRecordOptionsModel := new(dnsrecordsv1.UpdateDnsRecordOptions)
				updateDnsRecordOptionsModel.DnsrecordIdentifier = core.StringPtr("testString")
				updateDnsRecordOptionsModel.Name = core.StringPtr("host-1.test-example.com")
				updateDnsRecordOptionsModel.Type = core.StringPtr("A")
				updateDnsRecordOptionsModel.Content = core.StringPtr("1.2.3.4")
				updateDnsRecordOptionsModel.Priority = core.Int64Ptr(int64(5))
				updateDnsRecordOptionsModel.Proxied = core.BoolPtr(false)
				updateDnsRecordOptionsModel.Data = map[string]interface{}{"anyKey": "anyValue"}
				updateDnsRecordOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.UpdateDnsRecord(updateDnsRecordOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateDnsRecord(updateDnsRecordOptions *UpdateDnsRecordOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateDnsRecordPath := "/v1/testString/zones/testString/dns_records/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(updateDnsRecordPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "f1aba936b94213e5b8dca0c0dbf1f9cc", "created_on": "2014-01-01T05:20:00.12345Z", "modified_on": "2014-01-01T05:20:00.12345Z", "name": "host-1.test-example.com", "type": "A", "content": "169.154.10.10", "zone_id": "023e105f4ecef8ad9ca31a8372d0c353", "zone_name": "test-example.com", "proxiable": true, "proxied": false, "ttl": 120, "priority": 5, "data": {"anyKey": "anyValue"}}}`)
				}))
			})
			It(`Invoke UpdateDnsRecord successfully`, func() {
				testService, testServiceErr := dnsrecordsv1.NewDnsRecordsV1(&dnsrecordsv1.DnsRecordsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.UpdateDnsRecord(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateDnsRecordOptions model
				updateDnsRecordOptionsModel := new(dnsrecordsv1.UpdateDnsRecordOptions)
				updateDnsRecordOptionsModel.DnsrecordIdentifier = core.StringPtr("testString")
				updateDnsRecordOptionsModel.Name = core.StringPtr("host-1.test-example.com")
				updateDnsRecordOptionsModel.Type = core.StringPtr("A")
				updateDnsRecordOptionsModel.Content = core.StringPtr("1.2.3.4")
				updateDnsRecordOptionsModel.Priority = core.Int64Ptr(int64(5))
				updateDnsRecordOptionsModel.Proxied = core.BoolPtr(false)
				updateDnsRecordOptionsModel.Data = map[string]interface{}{"anyKey": "anyValue"}
				updateDnsRecordOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.UpdateDnsRecord(updateDnsRecordOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke UpdateDnsRecord with error: Operation validation and request error`, func() {
				testService, testServiceErr := dnsrecordsv1.NewDnsRecordsV1(&dnsrecordsv1.DnsRecordsV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the UpdateDnsRecordOptions model
				updateDnsRecordOptionsModel := new(dnsrecordsv1.UpdateDnsRecordOptions)
				updateDnsRecordOptionsModel.DnsrecordIdentifier = core.StringPtr("testString")
				updateDnsRecordOptionsModel.Name = core.StringPtr("host-1.test-example.com")
				updateDnsRecordOptionsModel.Type = core.StringPtr("A")
				updateDnsRecordOptionsModel.Content = core.StringPtr("1.2.3.4")
				updateDnsRecordOptionsModel.Priority = core.Int64Ptr(int64(5))
				updateDnsRecordOptionsModel.Proxied = core.BoolPtr(false)
				updateDnsRecordOptionsModel.Data = map[string]interface{}{"anyKey": "anyValue"}
				updateDnsRecordOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.UpdateDnsRecord(updateDnsRecordOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateDnsRecordOptions model with no property values
				updateDnsRecordOptionsModelNew := new(dnsrecordsv1.UpdateDnsRecordOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.UpdateDnsRecord(updateDnsRecordOptionsModelNew)
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
			testService, _ := dnsrecordsv1.NewDnsRecordsV1(&dnsrecordsv1.DnsRecordsV1Options{
				URL:            "http://dnsrecordsv1modelgenerator.com",
				Authenticator:  &core.NoAuthAuthenticator{},
				Crn:            core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})
			It(`Invoke NewCreateDnsRecordOptions successfully`, func() {
				// Construct an instance of the CreateDnsRecordOptions model
				createDnsRecordOptionsModel := testService.NewCreateDnsRecordOptions()
				createDnsRecordOptionsModel.SetName("host-1.test-example.com")
				createDnsRecordOptionsModel.SetType("A")
				createDnsRecordOptionsModel.SetContent("1.2.3.4")
				createDnsRecordOptionsModel.SetPriority(int64(5))
				createDnsRecordOptionsModel.SetData(map[string]interface{}{"anyKey": "anyValue"})
				createDnsRecordOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createDnsRecordOptionsModel).ToNot(BeNil())
				Expect(createDnsRecordOptionsModel.Name).To(Equal(core.StringPtr("host-1.test-example.com")))
				Expect(createDnsRecordOptionsModel.Type).To(Equal(core.StringPtr("A")))
				Expect(createDnsRecordOptionsModel.Content).To(Equal(core.StringPtr("1.2.3.4")))
				Expect(createDnsRecordOptionsModel.Priority).To(Equal(core.Int64Ptr(int64(5))))
				Expect(createDnsRecordOptionsModel.Data).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(createDnsRecordOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteDnsRecordOptions successfully`, func() {
				// Construct an instance of the DeleteDnsRecordOptions model
				dnsrecordIdentifier := "testString"
				deleteDnsRecordOptionsModel := testService.NewDeleteDnsRecordOptions(dnsrecordIdentifier)
				deleteDnsRecordOptionsModel.SetDnsrecordIdentifier("testString")
				deleteDnsRecordOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteDnsRecordOptionsModel).ToNot(BeNil())
				Expect(deleteDnsRecordOptionsModel.DnsrecordIdentifier).To(Equal(core.StringPtr("testString")))
				Expect(deleteDnsRecordOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetDnsRecordOptions successfully`, func() {
				// Construct an instance of the GetDnsRecordOptions model
				dnsrecordIdentifier := "testString"
				getDnsRecordOptionsModel := testService.NewGetDnsRecordOptions(dnsrecordIdentifier)
				getDnsRecordOptionsModel.SetDnsrecordIdentifier("testString")
				getDnsRecordOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getDnsRecordOptionsModel).ToNot(BeNil())
				Expect(getDnsRecordOptionsModel.DnsrecordIdentifier).To(Equal(core.StringPtr("testString")))
				Expect(getDnsRecordOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListAllDnsRecordsOptions successfully`, func() {
				// Construct an instance of the ListAllDnsRecordsOptions model
				listAllDnsRecordsOptionsModel := testService.NewListAllDnsRecordsOptions()
				listAllDnsRecordsOptionsModel.SetType("testString")
				listAllDnsRecordsOptionsModel.SetName("host1.test-example.com")
				listAllDnsRecordsOptionsModel.SetContent("1.2.3.4")
				listAllDnsRecordsOptionsModel.SetPage(int64(38))
				listAllDnsRecordsOptionsModel.SetPerPage(int64(5))
				listAllDnsRecordsOptionsModel.SetOrder("type")
				listAllDnsRecordsOptionsModel.SetDirection("asc")
				listAllDnsRecordsOptionsModel.SetMatch("any")
				listAllDnsRecordsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listAllDnsRecordsOptionsModel).ToNot(BeNil())
				Expect(listAllDnsRecordsOptionsModel.Type).To(Equal(core.StringPtr("testString")))
				Expect(listAllDnsRecordsOptionsModel.Name).To(Equal(core.StringPtr("host1.test-example.com")))
				Expect(listAllDnsRecordsOptionsModel.Content).To(Equal(core.StringPtr("1.2.3.4")))
				Expect(listAllDnsRecordsOptionsModel.Page).To(Equal(core.Int64Ptr(int64(38))))
				Expect(listAllDnsRecordsOptionsModel.PerPage).To(Equal(core.Int64Ptr(int64(5))))
				Expect(listAllDnsRecordsOptionsModel.Order).To(Equal(core.StringPtr("type")))
				Expect(listAllDnsRecordsOptionsModel.Direction).To(Equal(core.StringPtr("asc")))
				Expect(listAllDnsRecordsOptionsModel.Match).To(Equal(core.StringPtr("any")))
				Expect(listAllDnsRecordsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateDnsRecordOptions successfully`, func() {
				// Construct an instance of the UpdateDnsRecordOptions model
				dnsrecordIdentifier := "testString"
				updateDnsRecordOptionsModel := testService.NewUpdateDnsRecordOptions(dnsrecordIdentifier)
				updateDnsRecordOptionsModel.SetDnsrecordIdentifier("testString")
				updateDnsRecordOptionsModel.SetName("host-1.test-example.com")
				updateDnsRecordOptionsModel.SetType("A")
				updateDnsRecordOptionsModel.SetContent("1.2.3.4")
				updateDnsRecordOptionsModel.SetPriority(int64(5))
				updateDnsRecordOptionsModel.SetProxied(false)
				updateDnsRecordOptionsModel.SetData(map[string]interface{}{"anyKey": "anyValue"})
				updateDnsRecordOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateDnsRecordOptionsModel).ToNot(BeNil())
				Expect(updateDnsRecordOptionsModel.DnsrecordIdentifier).To(Equal(core.StringPtr("testString")))
				Expect(updateDnsRecordOptionsModel.Name).To(Equal(core.StringPtr("host-1.test-example.com")))
				Expect(updateDnsRecordOptionsModel.Type).To(Equal(core.StringPtr("A")))
				Expect(updateDnsRecordOptionsModel.Content).To(Equal(core.StringPtr("1.2.3.4")))
				Expect(updateDnsRecordOptionsModel.Priority).To(Equal(core.Int64Ptr(int64(5))))
				Expect(updateDnsRecordOptionsModel.Proxied).To(Equal(core.BoolPtr(false)))
				Expect(updateDnsRecordOptionsModel.Data).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(updateDnsRecordOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
