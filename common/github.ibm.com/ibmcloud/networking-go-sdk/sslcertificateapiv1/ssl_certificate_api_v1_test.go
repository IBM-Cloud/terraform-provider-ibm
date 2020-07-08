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

package sslcertificateapiv1_test

import (
	"bytes"
	"fmt"
	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.ibm.com/ibmcloud/networking-go-sdk/sslcertificateapiv1"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"time"
)

var _ = Describe(`SslCertificateApiV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		It(`Instantiate service client`, func() {
			testService, testServiceErr := sslcertificateapiv1.NewSslCertificateApiV1(&sslcertificateapiv1.SslCertificateApiV1Options{
				Authenticator:  &core.NoAuthAuthenticator{},
				Crn:            core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})
			Expect(testService).ToNot(BeNil())
			Expect(testServiceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			testService, testServiceErr := sslcertificateapiv1.NewSslCertificateApiV1(&sslcertificateapiv1.SslCertificateApiV1Options{
				URL:            "{BAD_URL_STRING",
				Crn:            core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			testService, testServiceErr := sslcertificateapiv1.NewSslCertificateApiV1(&sslcertificateapiv1.SslCertificateApiV1Options{
				URL:            "https://sslcertificateapiv1/api",
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
			testService, testServiceErr := sslcertificateapiv1.NewSslCertificateApiV1(&sslcertificateapiv1.SslCertificateApiV1Options{})
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
				"SSL_CERTIFICATE_API_URL":       "https://sslcertificateapiv1/api",
				"SSL_CERTIFICATE_API_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := sslcertificateapiv1.NewSslCertificateApiV1UsingExternalConfig(&sslcertificateapiv1.SslCertificateApiV1Options{
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := sslcertificateapiv1.NewSslCertificateApiV1UsingExternalConfig(&sslcertificateapiv1.SslCertificateApiV1Options{
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
				testService, testServiceErr := sslcertificateapiv1.NewSslCertificateApiV1UsingExternalConfig(&sslcertificateapiv1.SslCertificateApiV1Options{
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
				"SSL_CERTIFICATE_API_URL":       "https://sslcertificateapiv1/api",
				"SSL_CERTIFICATE_API_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := sslcertificateapiv1.NewSslCertificateApiV1UsingExternalConfig(&sslcertificateapiv1.SslCertificateApiV1Options{
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
				"SSL_CERTIFICATE_API_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := sslcertificateapiv1.NewSslCertificateApiV1UsingExternalConfig(&sslcertificateapiv1.SslCertificateApiV1Options{
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
	Describe(`ListCertificates(listCertificatesOptions *ListCertificatesOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		listCertificatesPath := "/v1/testString/zones/testString/ssl/certificate_packs"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listCertificatesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListCertificates with error: Operation response processing error`, func() {
				testService, testServiceErr := sslcertificateapiv1.NewSslCertificateApiV1(&sslcertificateapiv1.SslCertificateApiV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ListCertificatesOptions model
				listCertificatesOptionsModel := new(sslcertificateapiv1.ListCertificatesOptions)
				listCertificatesOptionsModel.XCorrelationID = core.StringPtr("testString")
				listCertificatesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.ListCertificates(listCertificatesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListCertificates(listCertificatesOptions *ListCertificatesOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		listCertificatesPath := "/v1/testString/zones/testString/ssl/certificate_packs"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listCertificatesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"result": [{"id": "0f405ba2-8c18-49eb-a30b-28b85427780f", "type": "dedicated", "hosts": ["example.com"], "certificates": [{"id": 436627, "hosts": ["example.com"], "status": "active"}], "primary_certificate": 0, "status": "active"}], "result_info": {"page": 1, "per_page": 2, "count": 1, "total_count": 200}, "success": true, "errors": [["Errors"]], "messages": [{"status": "OK"}]}`)
				}))
			})
			It(`Invoke ListCertificates successfully`, func() {
				testService, testServiceErr := sslcertificateapiv1.NewSslCertificateApiV1(&sslcertificateapiv1.SslCertificateApiV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListCertificates(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListCertificatesOptions model
				listCertificatesOptionsModel := new(sslcertificateapiv1.ListCertificatesOptions)
				listCertificatesOptionsModel.XCorrelationID = core.StringPtr("testString")
				listCertificatesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListCertificates(listCertificatesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke ListCertificates with error: Operation request error`, func() {
				testService, testServiceErr := sslcertificateapiv1.NewSslCertificateApiV1(&sslcertificateapiv1.SslCertificateApiV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ListCertificatesOptions model
				listCertificatesOptionsModel := new(sslcertificateapiv1.ListCertificatesOptions)
				listCertificatesOptionsModel.XCorrelationID = core.StringPtr("testString")
				listCertificatesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.ListCertificates(listCertificatesOptionsModel)
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
	Describe(`OrderCertificate(orderCertificateOptions *OrderCertificateOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		orderCertificatePath := "/v1/testString/zones/testString/ssl/certificate_packs"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(orderCertificatePath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke OrderCertificate with error: Operation response processing error`, func() {
				testService, testServiceErr := sslcertificateapiv1.NewSslCertificateApiV1(&sslcertificateapiv1.SslCertificateApiV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the OrderCertificateOptions model
				orderCertificateOptionsModel := new(sslcertificateapiv1.OrderCertificateOptions)
				orderCertificateOptionsModel.Type = core.StringPtr("dedicated")
				orderCertificateOptionsModel.Hosts = []string{"example.com"}
				orderCertificateOptionsModel.XCorrelationID = core.StringPtr("testString")
				orderCertificateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.OrderCertificate(orderCertificateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`OrderCertificate(orderCertificateOptions *OrderCertificateOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		orderCertificatePath := "/v1/testString/zones/testString/ssl/certificate_packs"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(orderCertificatePath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"id": "0f405ba2-8c18-49eb-a30b-28b85427780f", "type": "dedicated", "hosts": ["example.com"], "certificates": [{"id": 436627, "hosts": ["example.com"], "status": "active"}], "primary_certificate": 0, "status": "active"}`)
				}))
			})
			It(`Invoke OrderCertificate successfully`, func() {
				testService, testServiceErr := sslcertificateapiv1.NewSslCertificateApiV1(&sslcertificateapiv1.SslCertificateApiV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.OrderCertificate(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the OrderCertificateOptions model
				orderCertificateOptionsModel := new(sslcertificateapiv1.OrderCertificateOptions)
				orderCertificateOptionsModel.Type = core.StringPtr("dedicated")
				orderCertificateOptionsModel.Hosts = []string{"example.com"}
				orderCertificateOptionsModel.XCorrelationID = core.StringPtr("testString")
				orderCertificateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.OrderCertificate(orderCertificateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke OrderCertificate with error: Operation request error`, func() {
				testService, testServiceErr := sslcertificateapiv1.NewSslCertificateApiV1(&sslcertificateapiv1.SslCertificateApiV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the OrderCertificateOptions model
				orderCertificateOptionsModel := new(sslcertificateapiv1.OrderCertificateOptions)
				orderCertificateOptionsModel.Type = core.StringPtr("dedicated")
				orderCertificateOptionsModel.Hosts = []string{"example.com"}
				orderCertificateOptionsModel.XCorrelationID = core.StringPtr("testString")
				orderCertificateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.OrderCertificate(orderCertificateOptionsModel)
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

	Describe(`DeleteCertificate(deleteCertificateOptions *DeleteCertificateOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		deleteCertificatePath := "/v1/testString/zones/testString/ssl/certificate_packs/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(deleteCertificatePath))
					Expect(req.Method).To(Equal("DELETE"))
					Expect(req.Header["X-Correlation-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Correlation-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.WriteHeader(200)
				}))
			})
			It(`Invoke DeleteCertificate successfully`, func() {
				testService, testServiceErr := sslcertificateapiv1.NewSslCertificateApiV1(&sslcertificateapiv1.SslCertificateApiV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.DeleteCertificate(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteCertificateOptions model
				deleteCertificateOptionsModel := new(sslcertificateapiv1.DeleteCertificateOptions)
				deleteCertificateOptionsModel.CertIdentifier = core.StringPtr("testString")
				deleteCertificateOptionsModel.XCorrelationID = core.StringPtr("testString")
				deleteCertificateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.DeleteCertificate(deleteCertificateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteCertificate with error: Operation validation and request error`, func() {
				testService, testServiceErr := sslcertificateapiv1.NewSslCertificateApiV1(&sslcertificateapiv1.SslCertificateApiV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the DeleteCertificateOptions model
				deleteCertificateOptionsModel := new(sslcertificateapiv1.DeleteCertificateOptions)
				deleteCertificateOptionsModel.CertIdentifier = core.StringPtr("testString")
				deleteCertificateOptionsModel.XCorrelationID = core.StringPtr("testString")
				deleteCertificateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := testService.DeleteCertificate(deleteCertificateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteCertificateOptions model with no property values
				deleteCertificateOptionsModelNew := new(sslcertificateapiv1.DeleteCertificateOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = testService.DeleteCertificate(deleteCertificateOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetSslSetting(getSslSettingOptions *GetSslSettingOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getSslSettingPath := "/v1/testString/zones/testString/settings/ssl"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getSslSettingPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetSslSetting with error: Operation response processing error`, func() {
				testService, testServiceErr := sslcertificateapiv1.NewSslCertificateApiV1(&sslcertificateapiv1.SslCertificateApiV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetSslSettingOptions model
				getSslSettingOptionsModel := new(sslcertificateapiv1.GetSslSettingOptions)
				getSslSettingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.GetSslSetting(getSslSettingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetSslSetting(getSslSettingOptions *GetSslSettingOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getSslSettingPath := "/v1/testString/zones/testString/settings/ssl"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getSslSettingPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"success": true, "result": {"id": "ssl", "value": "off", "editable": true, "modified_on": "2017-01-01T05:20:00.12345Z"}, "errors": [["Errors"]], "messages": [{"status": "OK"}]}`)
				}))
			})
			It(`Invoke GetSslSetting successfully`, func() {
				testService, testServiceErr := sslcertificateapiv1.NewSslCertificateApiV1(&sslcertificateapiv1.SslCertificateApiV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetSslSetting(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetSslSettingOptions model
				getSslSettingOptionsModel := new(sslcertificateapiv1.GetSslSettingOptions)
				getSslSettingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetSslSetting(getSslSettingOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetSslSetting with error: Operation request error`, func() {
				testService, testServiceErr := sslcertificateapiv1.NewSslCertificateApiV1(&sslcertificateapiv1.SslCertificateApiV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetSslSettingOptions model
				getSslSettingOptionsModel := new(sslcertificateapiv1.GetSslSettingOptions)
				getSslSettingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.GetSslSetting(getSslSettingOptionsModel)
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
	Describe(`ChangeSslSetting(changeSslSettingOptions *ChangeSslSettingOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		changeSslSettingPath := "/v1/testString/zones/testString/settings/ssl"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(changeSslSettingPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ChangeSslSetting with error: Operation response processing error`, func() {
				testService, testServiceErr := sslcertificateapiv1.NewSslCertificateApiV1(&sslcertificateapiv1.SslCertificateApiV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ChangeSslSettingOptions model
				changeSslSettingOptionsModel := new(sslcertificateapiv1.ChangeSslSettingOptions)
				changeSslSettingOptionsModel.Value = core.StringPtr("off")
				changeSslSettingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.ChangeSslSetting(changeSslSettingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ChangeSslSetting(changeSslSettingOptions *ChangeSslSettingOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		changeSslSettingPath := "/v1/testString/zones/testString/settings/ssl"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(changeSslSettingPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"success": true, "result": {"id": "ssl", "value": "off", "editable": true, "modified_on": "2017-01-01T05:20:00.12345Z"}, "errors": [["Errors"]], "messages": [{"status": "OK"}]}`)
				}))
			})
			It(`Invoke ChangeSslSetting successfully`, func() {
				testService, testServiceErr := sslcertificateapiv1.NewSslCertificateApiV1(&sslcertificateapiv1.SslCertificateApiV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ChangeSslSetting(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ChangeSslSettingOptions model
				changeSslSettingOptionsModel := new(sslcertificateapiv1.ChangeSslSettingOptions)
				changeSslSettingOptionsModel.Value = core.StringPtr("off")
				changeSslSettingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ChangeSslSetting(changeSslSettingOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke ChangeSslSetting with error: Operation request error`, func() {
				testService, testServiceErr := sslcertificateapiv1.NewSslCertificateApiV1(&sslcertificateapiv1.SslCertificateApiV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ChangeSslSettingOptions model
				changeSslSettingOptionsModel := new(sslcertificateapiv1.ChangeSslSettingOptions)
				changeSslSettingOptionsModel.Value = core.StringPtr("off")
				changeSslSettingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.ChangeSslSetting(changeSslSettingOptionsModel)
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
	Describe(`ListCustomCertificates(listCustomCertificatesOptions *ListCustomCertificatesOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		listCustomCertificatesPath := "/v1/testString/zones/testString/custom_certificates"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listCustomCertificatesPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListCustomCertificates with error: Operation response processing error`, func() {
				testService, testServiceErr := sslcertificateapiv1.NewSslCertificateApiV1(&sslcertificateapiv1.SslCertificateApiV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ListCustomCertificatesOptions model
				listCustomCertificatesOptionsModel := new(sslcertificateapiv1.ListCustomCertificatesOptions)
				listCustomCertificatesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.ListCustomCertificates(listCustomCertificatesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListCustomCertificates(listCustomCertificatesOptions *ListCustomCertificatesOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		listCustomCertificatesPath := "/v1/testString/zones/testString/custom_certificates"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listCustomCertificatesPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"result": [{"id": "0f405ba2-8c18-49eb-a30b-28b85427780f", "hosts": ["example.com"], "issuer": "/Country=US/Organization=Lets Encrypt/CommonName=Lets Encrypt Authority X3", "signature": "SHA256WithRSA", "status": "active", "bundle_method": "BundleMethod", "zone_id": "ZoneID", "uploaded_on": "UploadedOn", "modified_on": "ModifiedOn", "expires_on": "ExpiresOn", "priority": 8}], "result_info": {"page": 1, "per_page": 2, "count": 1, "total_count": 200}, "success": true, "errors": [["Errors"]], "messages": [{"status": "OK"}]}`)
				}))
			})
			It(`Invoke ListCustomCertificates successfully`, func() {
				testService, testServiceErr := sslcertificateapiv1.NewSslCertificateApiV1(&sslcertificateapiv1.SslCertificateApiV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListCustomCertificates(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListCustomCertificatesOptions model
				listCustomCertificatesOptionsModel := new(sslcertificateapiv1.ListCustomCertificatesOptions)
				listCustomCertificatesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListCustomCertificates(listCustomCertificatesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke ListCustomCertificates with error: Operation request error`, func() {
				testService, testServiceErr := sslcertificateapiv1.NewSslCertificateApiV1(&sslcertificateapiv1.SslCertificateApiV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ListCustomCertificatesOptions model
				listCustomCertificatesOptionsModel := new(sslcertificateapiv1.ListCustomCertificatesOptions)
				listCustomCertificatesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.ListCustomCertificates(listCustomCertificatesOptionsModel)
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
	Describe(`UploadCustomCertificate(uploadCustomCertificateOptions *UploadCustomCertificateOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		uploadCustomCertificatePath := "/v1/testString/zones/testString/custom_certificates"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(uploadCustomCertificatePath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UploadCustomCertificate with error: Operation response processing error`, func() {
				testService, testServiceErr := sslcertificateapiv1.NewSslCertificateApiV1(&sslcertificateapiv1.SslCertificateApiV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the CustomCertReqGeoRestrictions model
				customCertReqGeoRestrictionsModel := new(sslcertificateapiv1.CustomCertReqGeoRestrictions)
				customCertReqGeoRestrictionsModel.Label = core.StringPtr("us")

				// Construct an instance of the UploadCustomCertificateOptions model
				uploadCustomCertificateOptionsModel := new(sslcertificateapiv1.UploadCustomCertificateOptions)
				uploadCustomCertificateOptionsModel.Certificate = core.StringPtr("testString")
				uploadCustomCertificateOptionsModel.PrivateKey = core.StringPtr("testString")
				uploadCustomCertificateOptionsModel.BundleMethod = core.StringPtr("ubiquitous")
				uploadCustomCertificateOptionsModel.GeoRestrictions = customCertReqGeoRestrictionsModel
				uploadCustomCertificateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.UploadCustomCertificate(uploadCustomCertificateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UploadCustomCertificate(uploadCustomCertificateOptions *UploadCustomCertificateOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		uploadCustomCertificatePath := "/v1/testString/zones/testString/custom_certificates"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(uploadCustomCertificatePath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"result": {"id": "0f405ba2-8c18-49eb-a30b-28b85427780f", "hosts": ["example.com"], "issuer": "/Country=US/Organization=Lets Encrypt/CommonName=Lets Encrypt Authority X3", "signature": "SHA256WithRSA", "status": "active", "bundle_method": "BundleMethod", "zone_id": "ZoneID", "uploaded_on": "UploadedOn", "modified_on": "ModifiedOn", "expires_on": "ExpiresOn", "priority": 8}, "success": true, "errors": [["Errors"]], "messages": [{"status": "OK"}]}`)
				}))
			})
			It(`Invoke UploadCustomCertificate successfully`, func() {
				testService, testServiceErr := sslcertificateapiv1.NewSslCertificateApiV1(&sslcertificateapiv1.SslCertificateApiV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.UploadCustomCertificate(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CustomCertReqGeoRestrictions model
				customCertReqGeoRestrictionsModel := new(sslcertificateapiv1.CustomCertReqGeoRestrictions)
				customCertReqGeoRestrictionsModel.Label = core.StringPtr("us")

				// Construct an instance of the UploadCustomCertificateOptions model
				uploadCustomCertificateOptionsModel := new(sslcertificateapiv1.UploadCustomCertificateOptions)
				uploadCustomCertificateOptionsModel.Certificate = core.StringPtr("testString")
				uploadCustomCertificateOptionsModel.PrivateKey = core.StringPtr("testString")
				uploadCustomCertificateOptionsModel.BundleMethod = core.StringPtr("ubiquitous")
				uploadCustomCertificateOptionsModel.GeoRestrictions = customCertReqGeoRestrictionsModel
				uploadCustomCertificateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.UploadCustomCertificate(uploadCustomCertificateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke UploadCustomCertificate with error: Operation request error`, func() {
				testService, testServiceErr := sslcertificateapiv1.NewSslCertificateApiV1(&sslcertificateapiv1.SslCertificateApiV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the CustomCertReqGeoRestrictions model
				customCertReqGeoRestrictionsModel := new(sslcertificateapiv1.CustomCertReqGeoRestrictions)
				customCertReqGeoRestrictionsModel.Label = core.StringPtr("us")

				// Construct an instance of the UploadCustomCertificateOptions model
				uploadCustomCertificateOptionsModel := new(sslcertificateapiv1.UploadCustomCertificateOptions)
				uploadCustomCertificateOptionsModel.Certificate = core.StringPtr("testString")
				uploadCustomCertificateOptionsModel.PrivateKey = core.StringPtr("testString")
				uploadCustomCertificateOptionsModel.BundleMethod = core.StringPtr("ubiquitous")
				uploadCustomCertificateOptionsModel.GeoRestrictions = customCertReqGeoRestrictionsModel
				uploadCustomCertificateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.UploadCustomCertificate(uploadCustomCertificateOptionsModel)
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
	Describe(`GetCustomCertificate(getCustomCertificateOptions *GetCustomCertificateOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getCustomCertificatePath := "/v1/testString/zones/testString/custom_certificates/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getCustomCertificatePath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetCustomCertificate with error: Operation response processing error`, func() {
				testService, testServiceErr := sslcertificateapiv1.NewSslCertificateApiV1(&sslcertificateapiv1.SslCertificateApiV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetCustomCertificateOptions model
				getCustomCertificateOptionsModel := new(sslcertificateapiv1.GetCustomCertificateOptions)
				getCustomCertificateOptionsModel.CustomCertID = core.StringPtr("testString")
				getCustomCertificateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.GetCustomCertificate(getCustomCertificateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetCustomCertificate(getCustomCertificateOptions *GetCustomCertificateOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getCustomCertificatePath := "/v1/testString/zones/testString/custom_certificates/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getCustomCertificatePath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"result": {"id": "0f405ba2-8c18-49eb-a30b-28b85427780f", "hosts": ["example.com"], "issuer": "/Country=US/Organization=Lets Encrypt/CommonName=Lets Encrypt Authority X3", "signature": "SHA256WithRSA", "status": "active", "bundle_method": "BundleMethod", "zone_id": "ZoneID", "uploaded_on": "UploadedOn", "modified_on": "ModifiedOn", "expires_on": "ExpiresOn", "priority": 8}, "success": true, "errors": [["Errors"]], "messages": [{"status": "OK"}]}`)
				}))
			})
			It(`Invoke GetCustomCertificate successfully`, func() {
				testService, testServiceErr := sslcertificateapiv1.NewSslCertificateApiV1(&sslcertificateapiv1.SslCertificateApiV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetCustomCertificate(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetCustomCertificateOptions model
				getCustomCertificateOptionsModel := new(sslcertificateapiv1.GetCustomCertificateOptions)
				getCustomCertificateOptionsModel.CustomCertID = core.StringPtr("testString")
				getCustomCertificateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetCustomCertificate(getCustomCertificateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetCustomCertificate with error: Operation validation and request error`, func() {
				testService, testServiceErr := sslcertificateapiv1.NewSslCertificateApiV1(&sslcertificateapiv1.SslCertificateApiV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetCustomCertificateOptions model
				getCustomCertificateOptionsModel := new(sslcertificateapiv1.GetCustomCertificateOptions)
				getCustomCertificateOptionsModel.CustomCertID = core.StringPtr("testString")
				getCustomCertificateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.GetCustomCertificate(getCustomCertificateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetCustomCertificateOptions model with no property values
				getCustomCertificateOptionsModelNew := new(sslcertificateapiv1.GetCustomCertificateOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.GetCustomCertificate(getCustomCertificateOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateCustomCertificate(updateCustomCertificateOptions *UpdateCustomCertificateOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateCustomCertificatePath := "/v1/testString/zones/testString/custom_certificates/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(updateCustomCertificatePath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateCustomCertificate with error: Operation response processing error`, func() {
				testService, testServiceErr := sslcertificateapiv1.NewSslCertificateApiV1(&sslcertificateapiv1.SslCertificateApiV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the CustomCertReqGeoRestrictions model
				customCertReqGeoRestrictionsModel := new(sslcertificateapiv1.CustomCertReqGeoRestrictions)
				customCertReqGeoRestrictionsModel.Label = core.StringPtr("us")

				// Construct an instance of the UpdateCustomCertificateOptions model
				updateCustomCertificateOptionsModel := new(sslcertificateapiv1.UpdateCustomCertificateOptions)
				updateCustomCertificateOptionsModel.CustomCertID = core.StringPtr("testString")
				updateCustomCertificateOptionsModel.Certificate = core.StringPtr("testString")
				updateCustomCertificateOptionsModel.PrivateKey = core.StringPtr("testString")
				updateCustomCertificateOptionsModel.BundleMethod = core.StringPtr("ubiquitous")
				updateCustomCertificateOptionsModel.GeoRestrictions = customCertReqGeoRestrictionsModel
				updateCustomCertificateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.UpdateCustomCertificate(updateCustomCertificateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateCustomCertificate(updateCustomCertificateOptions *UpdateCustomCertificateOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateCustomCertificatePath := "/v1/testString/zones/testString/custom_certificates/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(updateCustomCertificatePath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"result": {"id": "0f405ba2-8c18-49eb-a30b-28b85427780f", "hosts": ["example.com"], "issuer": "/Country=US/Organization=Lets Encrypt/CommonName=Lets Encrypt Authority X3", "signature": "SHA256WithRSA", "status": "active", "bundle_method": "BundleMethod", "zone_id": "ZoneID", "uploaded_on": "UploadedOn", "modified_on": "ModifiedOn", "expires_on": "ExpiresOn", "priority": 8}, "success": true, "errors": [["Errors"]], "messages": [{"status": "OK"}]}`)
				}))
			})
			It(`Invoke UpdateCustomCertificate successfully`, func() {
				testService, testServiceErr := sslcertificateapiv1.NewSslCertificateApiV1(&sslcertificateapiv1.SslCertificateApiV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.UpdateCustomCertificate(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CustomCertReqGeoRestrictions model
				customCertReqGeoRestrictionsModel := new(sslcertificateapiv1.CustomCertReqGeoRestrictions)
				customCertReqGeoRestrictionsModel.Label = core.StringPtr("us")

				// Construct an instance of the UpdateCustomCertificateOptions model
				updateCustomCertificateOptionsModel := new(sslcertificateapiv1.UpdateCustomCertificateOptions)
				updateCustomCertificateOptionsModel.CustomCertID = core.StringPtr("testString")
				updateCustomCertificateOptionsModel.Certificate = core.StringPtr("testString")
				updateCustomCertificateOptionsModel.PrivateKey = core.StringPtr("testString")
				updateCustomCertificateOptionsModel.BundleMethod = core.StringPtr("ubiquitous")
				updateCustomCertificateOptionsModel.GeoRestrictions = customCertReqGeoRestrictionsModel
				updateCustomCertificateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.UpdateCustomCertificate(updateCustomCertificateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke UpdateCustomCertificate with error: Operation validation and request error`, func() {
				testService, testServiceErr := sslcertificateapiv1.NewSslCertificateApiV1(&sslcertificateapiv1.SslCertificateApiV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the CustomCertReqGeoRestrictions model
				customCertReqGeoRestrictionsModel := new(sslcertificateapiv1.CustomCertReqGeoRestrictions)
				customCertReqGeoRestrictionsModel.Label = core.StringPtr("us")

				// Construct an instance of the UpdateCustomCertificateOptions model
				updateCustomCertificateOptionsModel := new(sslcertificateapiv1.UpdateCustomCertificateOptions)
				updateCustomCertificateOptionsModel.CustomCertID = core.StringPtr("testString")
				updateCustomCertificateOptionsModel.Certificate = core.StringPtr("testString")
				updateCustomCertificateOptionsModel.PrivateKey = core.StringPtr("testString")
				updateCustomCertificateOptionsModel.BundleMethod = core.StringPtr("ubiquitous")
				updateCustomCertificateOptionsModel.GeoRestrictions = customCertReqGeoRestrictionsModel
				updateCustomCertificateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.UpdateCustomCertificate(updateCustomCertificateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateCustomCertificateOptions model with no property values
				updateCustomCertificateOptionsModelNew := new(sslcertificateapiv1.UpdateCustomCertificateOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.UpdateCustomCertificate(updateCustomCertificateOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteCustomCertificate(deleteCustomCertificateOptions *DeleteCustomCertificateOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		deleteCustomCertificatePath := "/v1/testString/zones/testString/custom_certificates/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(deleteCustomCertificatePath))
					Expect(req.Method).To(Equal("DELETE"))
					res.WriteHeader(200)
				}))
			})
			It(`Invoke DeleteCustomCertificate successfully`, func() {
				testService, testServiceErr := sslcertificateapiv1.NewSslCertificateApiV1(&sslcertificateapiv1.SslCertificateApiV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.DeleteCustomCertificate(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteCustomCertificateOptions model
				deleteCustomCertificateOptionsModel := new(sslcertificateapiv1.DeleteCustomCertificateOptions)
				deleteCustomCertificateOptionsModel.CustomCertID = core.StringPtr("testString")
				deleteCustomCertificateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.DeleteCustomCertificate(deleteCustomCertificateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteCustomCertificate with error: Operation validation and request error`, func() {
				testService, testServiceErr := sslcertificateapiv1.NewSslCertificateApiV1(&sslcertificateapiv1.SslCertificateApiV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the DeleteCustomCertificateOptions model
				deleteCustomCertificateOptionsModel := new(sslcertificateapiv1.DeleteCustomCertificateOptions)
				deleteCustomCertificateOptionsModel.CustomCertID = core.StringPtr("testString")
				deleteCustomCertificateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := testService.DeleteCustomCertificate(deleteCustomCertificateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteCustomCertificateOptions model with no property values
				deleteCustomCertificateOptionsModelNew := new(sslcertificateapiv1.DeleteCustomCertificateOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = testService.DeleteCustomCertificate(deleteCustomCertificateOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ChangeCertificatePriority(changeCertificatePriorityOptions *ChangeCertificatePriorityOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		changeCertificatePriorityPath := "/v1/testString/zones/testString/custom_certificates/prioritize"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(changeCertificatePriorityPath))
					Expect(req.Method).To(Equal("PUT"))
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ChangeCertificatePriority successfully`, func() {
				testService, testServiceErr := sslcertificateapiv1.NewSslCertificateApiV1(&sslcertificateapiv1.SslCertificateApiV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.ChangeCertificatePriority(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the CertPriorityReqCertificatesItem model
				certPriorityReqCertificatesItemModel := new(sslcertificateapiv1.CertPriorityReqCertificatesItem)
				certPriorityReqCertificatesItemModel.ID = core.StringPtr("5a7805061c76ada191ed06f989cc3dac")
				certPriorityReqCertificatesItemModel.Priority = core.Int64Ptr(int64(1))

				// Construct an instance of the ChangeCertificatePriorityOptions model
				changeCertificatePriorityOptionsModel := new(sslcertificateapiv1.ChangeCertificatePriorityOptions)
				changeCertificatePriorityOptionsModel.Certificates = []sslcertificateapiv1.CertPriorityReqCertificatesItem{*certPriorityReqCertificatesItemModel}
				changeCertificatePriorityOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.ChangeCertificatePriority(changeCertificatePriorityOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke ChangeCertificatePriority with error: Operation request error`, func() {
				testService, testServiceErr := sslcertificateapiv1.NewSslCertificateApiV1(&sslcertificateapiv1.SslCertificateApiV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the CertPriorityReqCertificatesItem model
				certPriorityReqCertificatesItemModel := new(sslcertificateapiv1.CertPriorityReqCertificatesItem)
				certPriorityReqCertificatesItemModel.ID = core.StringPtr("5a7805061c76ada191ed06f989cc3dac")
				certPriorityReqCertificatesItemModel.Priority = core.Int64Ptr(int64(1))

				// Construct an instance of the ChangeCertificatePriorityOptions model
				changeCertificatePriorityOptionsModel := new(sslcertificateapiv1.ChangeCertificatePriorityOptions)
				changeCertificatePriorityOptionsModel.Certificates = []sslcertificateapiv1.CertPriorityReqCertificatesItem{*certPriorityReqCertificatesItemModel}
				changeCertificatePriorityOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := testService.ChangeCertificatePriority(changeCertificatePriorityOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetUniversalCertificateSetting(getUniversalCertificateSettingOptions *GetUniversalCertificateSettingOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getUniversalCertificateSettingPath := "/v1/testString/zones/testString/ssl/universal/settings"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getUniversalCertificateSettingPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetUniversalCertificateSetting with error: Operation response processing error`, func() {
				testService, testServiceErr := sslcertificateapiv1.NewSslCertificateApiV1(&sslcertificateapiv1.SslCertificateApiV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetUniversalCertificateSettingOptions model
				getUniversalCertificateSettingOptionsModel := new(sslcertificateapiv1.GetUniversalCertificateSettingOptions)
				getUniversalCertificateSettingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.GetUniversalCertificateSetting(getUniversalCertificateSettingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetUniversalCertificateSetting(getUniversalCertificateSettingOptions *GetUniversalCertificateSettingOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getUniversalCertificateSettingPath := "/v1/testString/zones/testString/ssl/universal/settings"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getUniversalCertificateSettingPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"result": {"enabled": true}, "success": true, "errors": [["Errors"]], "messages": [{"status": "OK"}]}`)
				}))
			})
			It(`Invoke GetUniversalCertificateSetting successfully`, func() {
				testService, testServiceErr := sslcertificateapiv1.NewSslCertificateApiV1(&sslcertificateapiv1.SslCertificateApiV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetUniversalCertificateSetting(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetUniversalCertificateSettingOptions model
				getUniversalCertificateSettingOptionsModel := new(sslcertificateapiv1.GetUniversalCertificateSettingOptions)
				getUniversalCertificateSettingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetUniversalCertificateSetting(getUniversalCertificateSettingOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetUniversalCertificateSetting with error: Operation request error`, func() {
				testService, testServiceErr := sslcertificateapiv1.NewSslCertificateApiV1(&sslcertificateapiv1.SslCertificateApiV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetUniversalCertificateSettingOptions model
				getUniversalCertificateSettingOptionsModel := new(sslcertificateapiv1.GetUniversalCertificateSettingOptions)
				getUniversalCertificateSettingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.GetUniversalCertificateSetting(getUniversalCertificateSettingOptionsModel)
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

	Describe(`ChangeUniversalCertificateSetting(changeUniversalCertificateSettingOptions *ChangeUniversalCertificateSettingOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		changeUniversalCertificateSettingPath := "/v1/testString/zones/testString/ssl/universal/settings"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(changeUniversalCertificateSettingPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ChangeUniversalCertificateSetting successfully`, func() {
				testService, testServiceErr := sslcertificateapiv1.NewSslCertificateApiV1(&sslcertificateapiv1.SslCertificateApiV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.ChangeUniversalCertificateSetting(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the ChangeUniversalCertificateSettingOptions model
				changeUniversalCertificateSettingOptionsModel := new(sslcertificateapiv1.ChangeUniversalCertificateSettingOptions)
				changeUniversalCertificateSettingOptionsModel.Enabled = core.BoolPtr(true)
				changeUniversalCertificateSettingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.ChangeUniversalCertificateSetting(changeUniversalCertificateSettingOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke ChangeUniversalCertificateSetting with error: Operation request error`, func() {
				testService, testServiceErr := sslcertificateapiv1.NewSslCertificateApiV1(&sslcertificateapiv1.SslCertificateApiV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ChangeUniversalCertificateSettingOptions model
				changeUniversalCertificateSettingOptionsModel := new(sslcertificateapiv1.ChangeUniversalCertificateSettingOptions)
				changeUniversalCertificateSettingOptionsModel.Enabled = core.BoolPtr(true)
				changeUniversalCertificateSettingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := testService.ChangeUniversalCertificateSetting(changeUniversalCertificateSettingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetTls12Setting(getTls12SettingOptions *GetTls12SettingOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getTls12SettingPath := "/v1/testString/zones/testString/settings/tls_1_2_only"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getTls12SettingPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetTls12Setting with error: Operation response processing error`, func() {
				testService, testServiceErr := sslcertificateapiv1.NewSslCertificateApiV1(&sslcertificateapiv1.SslCertificateApiV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetTls12SettingOptions model
				getTls12SettingOptionsModel := new(sslcertificateapiv1.GetTls12SettingOptions)
				getTls12SettingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.GetTls12Setting(getTls12SettingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetTls12Setting(getTls12SettingOptions *GetTls12SettingOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getTls12SettingPath := "/v1/testString/zones/testString/settings/tls_1_2_only"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getTls12SettingPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"result": {"id": "tls_1_2_only", "value": "on", "editable": true, "modified_on": "2019-01-01T12:00:00"}, "success": true, "errors": [["Errors"]], "messages": [{"status": "OK"}]}`)
				}))
			})
			It(`Invoke GetTls12Setting successfully`, func() {
				testService, testServiceErr := sslcertificateapiv1.NewSslCertificateApiV1(&sslcertificateapiv1.SslCertificateApiV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetTls12Setting(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetTls12SettingOptions model
				getTls12SettingOptionsModel := new(sslcertificateapiv1.GetTls12SettingOptions)
				getTls12SettingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetTls12Setting(getTls12SettingOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetTls12Setting with error: Operation request error`, func() {
				testService, testServiceErr := sslcertificateapiv1.NewSslCertificateApiV1(&sslcertificateapiv1.SslCertificateApiV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetTls12SettingOptions model
				getTls12SettingOptionsModel := new(sslcertificateapiv1.GetTls12SettingOptions)
				getTls12SettingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.GetTls12Setting(getTls12SettingOptionsModel)
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
	Describe(`ChangeTls12Setting(changeTls12SettingOptions *ChangeTls12SettingOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		changeTls12SettingPath := "/v1/testString/zones/testString/settings/tls_1_2_only"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(changeTls12SettingPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ChangeTls12Setting with error: Operation response processing error`, func() {
				testService, testServiceErr := sslcertificateapiv1.NewSslCertificateApiV1(&sslcertificateapiv1.SslCertificateApiV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ChangeTls12SettingOptions model
				changeTls12SettingOptionsModel := new(sslcertificateapiv1.ChangeTls12SettingOptions)
				changeTls12SettingOptionsModel.Value = core.StringPtr("on")
				changeTls12SettingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.ChangeTls12Setting(changeTls12SettingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ChangeTls12Setting(changeTls12SettingOptions *ChangeTls12SettingOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		changeTls12SettingPath := "/v1/testString/zones/testString/settings/tls_1_2_only"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(changeTls12SettingPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"result": {"id": "tls_1_2_only", "value": "on", "editable": true, "modified_on": "2019-01-01T12:00:00"}, "success": true, "errors": [["Errors"]], "messages": [{"status": "OK"}]}`)
				}))
			})
			It(`Invoke ChangeTls12Setting successfully`, func() {
				testService, testServiceErr := sslcertificateapiv1.NewSslCertificateApiV1(&sslcertificateapiv1.SslCertificateApiV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ChangeTls12Setting(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ChangeTls12SettingOptions model
				changeTls12SettingOptionsModel := new(sslcertificateapiv1.ChangeTls12SettingOptions)
				changeTls12SettingOptionsModel.Value = core.StringPtr("on")
				changeTls12SettingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ChangeTls12Setting(changeTls12SettingOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke ChangeTls12Setting with error: Operation request error`, func() {
				testService, testServiceErr := sslcertificateapiv1.NewSslCertificateApiV1(&sslcertificateapiv1.SslCertificateApiV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ChangeTls12SettingOptions model
				changeTls12SettingOptionsModel := new(sslcertificateapiv1.ChangeTls12SettingOptions)
				changeTls12SettingOptionsModel.Value = core.StringPtr("on")
				changeTls12SettingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.ChangeTls12Setting(changeTls12SettingOptionsModel)
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
	Describe(`GetTls13Setting(getTls13SettingOptions *GetTls13SettingOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getTls13SettingPath := "/v1/testString/zones/testString/settings/tls_1_3"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getTls13SettingPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetTls13Setting with error: Operation response processing error`, func() {
				testService, testServiceErr := sslcertificateapiv1.NewSslCertificateApiV1(&sslcertificateapiv1.SslCertificateApiV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetTls13SettingOptions model
				getTls13SettingOptionsModel := new(sslcertificateapiv1.GetTls13SettingOptions)
				getTls13SettingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.GetTls13Setting(getTls13SettingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetTls13Setting(getTls13SettingOptions *GetTls13SettingOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getTls13SettingPath := "/v1/testString/zones/testString/settings/tls_1_3"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getTls13SettingPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"result": {"id": "tls_1_3", "value": "on", "editable": true, "modified_on": "2019-01-01T12:00:00"}, "success": true, "errors": [["Errors"]], "messages": [{"status": "OK"}]}`)
				}))
			})
			It(`Invoke GetTls13Setting successfully`, func() {
				testService, testServiceErr := sslcertificateapiv1.NewSslCertificateApiV1(&sslcertificateapiv1.SslCertificateApiV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetTls13Setting(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetTls13SettingOptions model
				getTls13SettingOptionsModel := new(sslcertificateapiv1.GetTls13SettingOptions)
				getTls13SettingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetTls13Setting(getTls13SettingOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetTls13Setting with error: Operation request error`, func() {
				testService, testServiceErr := sslcertificateapiv1.NewSslCertificateApiV1(&sslcertificateapiv1.SslCertificateApiV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetTls13SettingOptions model
				getTls13SettingOptionsModel := new(sslcertificateapiv1.GetTls13SettingOptions)
				getTls13SettingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.GetTls13Setting(getTls13SettingOptionsModel)
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
	Describe(`ChangeTls13Setting(changeTls13SettingOptions *ChangeTls13SettingOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		changeTls13SettingPath := "/v1/testString/zones/testString/settings/tls_1_3"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(changeTls13SettingPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ChangeTls13Setting with error: Operation response processing error`, func() {
				testService, testServiceErr := sslcertificateapiv1.NewSslCertificateApiV1(&sslcertificateapiv1.SslCertificateApiV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ChangeTls13SettingOptions model
				changeTls13SettingOptionsModel := new(sslcertificateapiv1.ChangeTls13SettingOptions)
				changeTls13SettingOptionsModel.Value = core.StringPtr("on")
				changeTls13SettingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.ChangeTls13Setting(changeTls13SettingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ChangeTls13Setting(changeTls13SettingOptions *ChangeTls13SettingOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		changeTls13SettingPath := "/v1/testString/zones/testString/settings/tls_1_3"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(changeTls13SettingPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"result": {"id": "tls_1_3", "value": "on", "editable": true, "modified_on": "2019-01-01T12:00:00"}, "success": true, "errors": [["Errors"]], "messages": [{"status": "OK"}]}`)
				}))
			})
			It(`Invoke ChangeTls13Setting successfully`, func() {
				testService, testServiceErr := sslcertificateapiv1.NewSslCertificateApiV1(&sslcertificateapiv1.SslCertificateApiV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ChangeTls13Setting(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ChangeTls13SettingOptions model
				changeTls13SettingOptionsModel := new(sslcertificateapiv1.ChangeTls13SettingOptions)
				changeTls13SettingOptionsModel.Value = core.StringPtr("on")
				changeTls13SettingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ChangeTls13Setting(changeTls13SettingOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke ChangeTls13Setting with error: Operation request error`, func() {
				testService, testServiceErr := sslcertificateapiv1.NewSslCertificateApiV1(&sslcertificateapiv1.SslCertificateApiV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ChangeTls13SettingOptions model
				changeTls13SettingOptionsModel := new(sslcertificateapiv1.ChangeTls13SettingOptions)
				changeTls13SettingOptionsModel.Value = core.StringPtr("on")
				changeTls13SettingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.ChangeTls13Setting(changeTls13SettingOptionsModel)
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
			zoneIdentifier := "testString"
			testService, _ := sslcertificateapiv1.NewSslCertificateApiV1(&sslcertificateapiv1.SslCertificateApiV1Options{
				URL:            "http://sslcertificateapiv1modelgenerator.com",
				Authenticator:  &core.NoAuthAuthenticator{},
				Crn:            core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})
			It(`Invoke NewCertPriorityReqCertificatesItem successfully`, func() {
				id := "5a7805061c76ada191ed06f989cc3dac"
				priority := int64(1)
				model, err := testService.NewCertPriorityReqCertificatesItem(id, priority)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewChangeCertificatePriorityOptions successfully`, func() {
				// Construct an instance of the CertPriorityReqCertificatesItem model
				certPriorityReqCertificatesItemModel := new(sslcertificateapiv1.CertPriorityReqCertificatesItem)
				Expect(certPriorityReqCertificatesItemModel).ToNot(BeNil())
				certPriorityReqCertificatesItemModel.ID = core.StringPtr("5a7805061c76ada191ed06f989cc3dac")
				certPriorityReqCertificatesItemModel.Priority = core.Int64Ptr(int64(1))
				Expect(certPriorityReqCertificatesItemModel.ID).To(Equal(core.StringPtr("5a7805061c76ada191ed06f989cc3dac")))
				Expect(certPriorityReqCertificatesItemModel.Priority).To(Equal(core.Int64Ptr(int64(1))))

				// Construct an instance of the ChangeCertificatePriorityOptions model
				changeCertificatePriorityOptionsModel := testService.NewChangeCertificatePriorityOptions()
				changeCertificatePriorityOptionsModel.SetCertificates([]sslcertificateapiv1.CertPriorityReqCertificatesItem{*certPriorityReqCertificatesItemModel})
				changeCertificatePriorityOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(changeCertificatePriorityOptionsModel).ToNot(BeNil())
				Expect(changeCertificatePriorityOptionsModel.Certificates).To(Equal([]sslcertificateapiv1.CertPriorityReqCertificatesItem{*certPriorityReqCertificatesItemModel}))
				Expect(changeCertificatePriorityOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewChangeSslSettingOptions successfully`, func() {
				// Construct an instance of the ChangeSslSettingOptions model
				changeSslSettingOptionsModel := testService.NewChangeSslSettingOptions()
				changeSslSettingOptionsModel.SetValue("off")
				changeSslSettingOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(changeSslSettingOptionsModel).ToNot(BeNil())
				Expect(changeSslSettingOptionsModel.Value).To(Equal(core.StringPtr("off")))
				Expect(changeSslSettingOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewChangeTls12SettingOptions successfully`, func() {
				// Construct an instance of the ChangeTls12SettingOptions model
				changeTls12SettingOptionsModel := testService.NewChangeTls12SettingOptions()
				changeTls12SettingOptionsModel.SetValue("on")
				changeTls12SettingOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(changeTls12SettingOptionsModel).ToNot(BeNil())
				Expect(changeTls12SettingOptionsModel.Value).To(Equal(core.StringPtr("on")))
				Expect(changeTls12SettingOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewChangeTls13SettingOptions successfully`, func() {
				// Construct an instance of the ChangeTls13SettingOptions model
				changeTls13SettingOptionsModel := testService.NewChangeTls13SettingOptions()
				changeTls13SettingOptionsModel.SetValue("on")
				changeTls13SettingOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(changeTls13SettingOptionsModel).ToNot(BeNil())
				Expect(changeTls13SettingOptionsModel.Value).To(Equal(core.StringPtr("on")))
				Expect(changeTls13SettingOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewChangeUniversalCertificateSettingOptions successfully`, func() {
				// Construct an instance of the ChangeUniversalCertificateSettingOptions model
				changeUniversalCertificateSettingOptionsModel := testService.NewChangeUniversalCertificateSettingOptions()
				changeUniversalCertificateSettingOptionsModel.SetEnabled(true)
				changeUniversalCertificateSettingOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(changeUniversalCertificateSettingOptionsModel).ToNot(BeNil())
				Expect(changeUniversalCertificateSettingOptionsModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(changeUniversalCertificateSettingOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCustomCertReqGeoRestrictions successfully`, func() {
				label := "us"
				model, err := testService.NewCustomCertReqGeoRestrictions(label)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewDeleteCertificateOptions successfully`, func() {
				// Construct an instance of the DeleteCertificateOptions model
				certIdentifier := "testString"
				deleteCertificateOptionsModel := testService.NewDeleteCertificateOptions(certIdentifier)
				deleteCertificateOptionsModel.SetCertIdentifier("testString")
				deleteCertificateOptionsModel.SetXCorrelationID("testString")
				deleteCertificateOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteCertificateOptionsModel).ToNot(BeNil())
				Expect(deleteCertificateOptionsModel.CertIdentifier).To(Equal(core.StringPtr("testString")))
				Expect(deleteCertificateOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(deleteCertificateOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteCustomCertificateOptions successfully`, func() {
				// Construct an instance of the DeleteCustomCertificateOptions model
				customCertID := "testString"
				deleteCustomCertificateOptionsModel := testService.NewDeleteCustomCertificateOptions(customCertID)
				deleteCustomCertificateOptionsModel.SetCustomCertID("testString")
				deleteCustomCertificateOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteCustomCertificateOptionsModel).ToNot(BeNil())
				Expect(deleteCustomCertificateOptionsModel.CustomCertID).To(Equal(core.StringPtr("testString")))
				Expect(deleteCustomCertificateOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetCustomCertificateOptions successfully`, func() {
				// Construct an instance of the GetCustomCertificateOptions model
				customCertID := "testString"
				getCustomCertificateOptionsModel := testService.NewGetCustomCertificateOptions(customCertID)
				getCustomCertificateOptionsModel.SetCustomCertID("testString")
				getCustomCertificateOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getCustomCertificateOptionsModel).ToNot(BeNil())
				Expect(getCustomCertificateOptionsModel.CustomCertID).To(Equal(core.StringPtr("testString")))
				Expect(getCustomCertificateOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetSslSettingOptions successfully`, func() {
				// Construct an instance of the GetSslSettingOptions model
				getSslSettingOptionsModel := testService.NewGetSslSettingOptions()
				getSslSettingOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getSslSettingOptionsModel).ToNot(BeNil())
				Expect(getSslSettingOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetTls12SettingOptions successfully`, func() {
				// Construct an instance of the GetTls12SettingOptions model
				getTls12SettingOptionsModel := testService.NewGetTls12SettingOptions()
				getTls12SettingOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getTls12SettingOptionsModel).ToNot(BeNil())
				Expect(getTls12SettingOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetTls13SettingOptions successfully`, func() {
				// Construct an instance of the GetTls13SettingOptions model
				getTls13SettingOptionsModel := testService.NewGetTls13SettingOptions()
				getTls13SettingOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getTls13SettingOptionsModel).ToNot(BeNil())
				Expect(getTls13SettingOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetUniversalCertificateSettingOptions successfully`, func() {
				// Construct an instance of the GetUniversalCertificateSettingOptions model
				getUniversalCertificateSettingOptionsModel := testService.NewGetUniversalCertificateSettingOptions()
				getUniversalCertificateSettingOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getUniversalCertificateSettingOptionsModel).ToNot(BeNil())
				Expect(getUniversalCertificateSettingOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListCertificatesOptions successfully`, func() {
				// Construct an instance of the ListCertificatesOptions model
				listCertificatesOptionsModel := testService.NewListCertificatesOptions()
				listCertificatesOptionsModel.SetXCorrelationID("testString")
				listCertificatesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listCertificatesOptionsModel).ToNot(BeNil())
				Expect(listCertificatesOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(listCertificatesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListCustomCertificatesOptions successfully`, func() {
				// Construct an instance of the ListCustomCertificatesOptions model
				listCustomCertificatesOptionsModel := testService.NewListCustomCertificatesOptions()
				listCustomCertificatesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listCustomCertificatesOptionsModel).ToNot(BeNil())
				Expect(listCustomCertificatesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewOrderCertificateOptions successfully`, func() {
				// Construct an instance of the OrderCertificateOptions model
				orderCertificateOptionsModel := testService.NewOrderCertificateOptions()
				orderCertificateOptionsModel.SetType("dedicated")
				orderCertificateOptionsModel.SetHosts([]string{"example.com"})
				orderCertificateOptionsModel.SetXCorrelationID("testString")
				orderCertificateOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(orderCertificateOptionsModel).ToNot(BeNil())
				Expect(orderCertificateOptionsModel.Type).To(Equal(core.StringPtr("dedicated")))
				Expect(orderCertificateOptionsModel.Hosts).To(Equal([]string{"example.com"}))
				Expect(orderCertificateOptionsModel.XCorrelationID).To(Equal(core.StringPtr("testString")))
				Expect(orderCertificateOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateCustomCertificateOptions successfully`, func() {
				// Construct an instance of the CustomCertReqGeoRestrictions model
				customCertReqGeoRestrictionsModel := new(sslcertificateapiv1.CustomCertReqGeoRestrictions)
				Expect(customCertReqGeoRestrictionsModel).ToNot(BeNil())
				customCertReqGeoRestrictionsModel.Label = core.StringPtr("us")
				Expect(customCertReqGeoRestrictionsModel.Label).To(Equal(core.StringPtr("us")))

				// Construct an instance of the UpdateCustomCertificateOptions model
				customCertID := "testString"
				updateCustomCertificateOptionsModel := testService.NewUpdateCustomCertificateOptions(customCertID)
				updateCustomCertificateOptionsModel.SetCustomCertID("testString")
				updateCustomCertificateOptionsModel.SetCertificate("testString")
				updateCustomCertificateOptionsModel.SetPrivateKey("testString")
				updateCustomCertificateOptionsModel.SetBundleMethod("ubiquitous")
				updateCustomCertificateOptionsModel.SetGeoRestrictions(customCertReqGeoRestrictionsModel)
				updateCustomCertificateOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateCustomCertificateOptionsModel).ToNot(BeNil())
				Expect(updateCustomCertificateOptionsModel.CustomCertID).To(Equal(core.StringPtr("testString")))
				Expect(updateCustomCertificateOptionsModel.Certificate).To(Equal(core.StringPtr("testString")))
				Expect(updateCustomCertificateOptionsModel.PrivateKey).To(Equal(core.StringPtr("testString")))
				Expect(updateCustomCertificateOptionsModel.BundleMethod).To(Equal(core.StringPtr("ubiquitous")))
				Expect(updateCustomCertificateOptionsModel.GeoRestrictions).To(Equal(customCertReqGeoRestrictionsModel))
				Expect(updateCustomCertificateOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUploadCustomCertificateOptions successfully`, func() {
				// Construct an instance of the CustomCertReqGeoRestrictions model
				customCertReqGeoRestrictionsModel := new(sslcertificateapiv1.CustomCertReqGeoRestrictions)
				Expect(customCertReqGeoRestrictionsModel).ToNot(BeNil())
				customCertReqGeoRestrictionsModel.Label = core.StringPtr("us")
				Expect(customCertReqGeoRestrictionsModel.Label).To(Equal(core.StringPtr("us")))

				// Construct an instance of the UploadCustomCertificateOptions model
				uploadCustomCertificateOptionsModel := testService.NewUploadCustomCertificateOptions()
				uploadCustomCertificateOptionsModel.SetCertificate("testString")
				uploadCustomCertificateOptionsModel.SetPrivateKey("testString")
				uploadCustomCertificateOptionsModel.SetBundleMethod("ubiquitous")
				uploadCustomCertificateOptionsModel.SetGeoRestrictions(customCertReqGeoRestrictionsModel)
				uploadCustomCertificateOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(uploadCustomCertificateOptionsModel).ToNot(BeNil())
				Expect(uploadCustomCertificateOptionsModel.Certificate).To(Equal(core.StringPtr("testString")))
				Expect(uploadCustomCertificateOptionsModel.PrivateKey).To(Equal(core.StringPtr("testString")))
				Expect(uploadCustomCertificateOptionsModel.BundleMethod).To(Equal(core.StringPtr("ubiquitous")))
				Expect(uploadCustomCertificateOptionsModel.GeoRestrictions).To(Equal(customCertReqGeoRestrictionsModel))
				Expect(uploadCustomCertificateOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
