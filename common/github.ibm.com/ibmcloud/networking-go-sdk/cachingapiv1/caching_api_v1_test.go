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

package cachingapiv1_test

import (
	"bytes"
	"fmt"
	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.ibm.com/ibmcloud/networking-go-sdk/cachingapiv1"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"time"
)

var _ = Describe(`CachingApiV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		crn := "testString"
		zoneID := "testString"
		It(`Instantiate service client`, func() {
			testService, testServiceErr := cachingapiv1.NewCachingApiV1(&cachingapiv1.CachingApiV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
				Crn:           core.StringPtr(crn),
				ZoneID:        core.StringPtr(zoneID),
			})
			Expect(testService).ToNot(BeNil())
			Expect(testServiceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			testService, testServiceErr := cachingapiv1.NewCachingApiV1(&cachingapiv1.CachingApiV1Options{
				URL:    "{BAD_URL_STRING",
				Crn:    core.StringPtr(crn),
				ZoneID: core.StringPtr(zoneID),
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			testService, testServiceErr := cachingapiv1.NewCachingApiV1(&cachingapiv1.CachingApiV1Options{
				URL:    "https://cachingapiv1/api",
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
			testService, testServiceErr := cachingapiv1.NewCachingApiV1(&cachingapiv1.CachingApiV1Options{})
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
				"CACHING_API_URL":       "https://cachingapiv1/api",
				"CACHING_API_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := cachingapiv1.NewCachingApiV1UsingExternalConfig(&cachingapiv1.CachingApiV1Options{
					Crn:    core.StringPtr(crn),
					ZoneID: core.StringPtr(zoneID),
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := cachingapiv1.NewCachingApiV1UsingExternalConfig(&cachingapiv1.CachingApiV1Options{
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
				testService, testServiceErr := cachingapiv1.NewCachingApiV1UsingExternalConfig(&cachingapiv1.CachingApiV1Options{
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
				"CACHING_API_URL":       "https://cachingapiv1/api",
				"CACHING_API_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := cachingapiv1.NewCachingApiV1UsingExternalConfig(&cachingapiv1.CachingApiV1Options{
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
				"CACHING_API_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := cachingapiv1.NewCachingApiV1UsingExternalConfig(&cachingapiv1.CachingApiV1Options{
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
	Describe(`PurgeAll(purgeAllOptions *PurgeAllOptions) - Operation response error`, func() {
		crn := "testString"
		zoneID := "testString"
		purgeAllPath := "/v1/testString/zones/testString/purge_cache/purge_all"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(purgeAllPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke PurgeAll with error: Operation response processing error`, func() {
				testService, testServiceErr := cachingapiv1.NewCachingApiV1(&cachingapiv1.CachingApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the PurgeAllOptions model
				purgeAllOptionsModel := new(cachingapiv1.PurgeAllOptions)
				purgeAllOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.PurgeAll(purgeAllOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`PurgeAll(purgeAllOptions *PurgeAllOptions)`, func() {
		crn := "testString"
		zoneID := "testString"
		purgeAllPath := "/v1/testString/zones/testString/purge_cache/purge_all"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(purgeAllPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "62d26b178b67c0eda0613891f3f51b0a"}}`)
				}))
			})
			It(`Invoke PurgeAll successfully`, func() {
				testService, testServiceErr := cachingapiv1.NewCachingApiV1(&cachingapiv1.CachingApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.PurgeAll(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the PurgeAllOptions model
				purgeAllOptionsModel := new(cachingapiv1.PurgeAllOptions)
				purgeAllOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.PurgeAll(purgeAllOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke PurgeAll with error: Operation request error`, func() {
				testService, testServiceErr := cachingapiv1.NewCachingApiV1(&cachingapiv1.CachingApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the PurgeAllOptions model
				purgeAllOptionsModel := new(cachingapiv1.PurgeAllOptions)
				purgeAllOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.PurgeAll(purgeAllOptionsModel)
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
	Describe(`PurgeByUrls(purgeByUrlsOptions *PurgeByUrlsOptions) - Operation response error`, func() {
		crn := "testString"
		zoneID := "testString"
		purgeByUrlsPath := "/v1/testString/zones/testString/purge_cache/purge_by_urls"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(purgeByUrlsPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke PurgeByUrls with error: Operation response processing error`, func() {
				testService, testServiceErr := cachingapiv1.NewCachingApiV1(&cachingapiv1.CachingApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the PurgeByUrlsOptions model
				purgeByUrlsOptionsModel := new(cachingapiv1.PurgeByUrlsOptions)
				purgeByUrlsOptionsModel.Files = []string{"http://www.example.com/cat_picture.jpg"}
				purgeByUrlsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.PurgeByUrls(purgeByUrlsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`PurgeByUrls(purgeByUrlsOptions *PurgeByUrlsOptions)`, func() {
		crn := "testString"
		zoneID := "testString"
		purgeByUrlsPath := "/v1/testString/zones/testString/purge_cache/purge_by_urls"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(purgeByUrlsPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "62d26b178b67c0eda0613891f3f51b0a"}}`)
				}))
			})
			It(`Invoke PurgeByUrls successfully`, func() {
				testService, testServiceErr := cachingapiv1.NewCachingApiV1(&cachingapiv1.CachingApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.PurgeByUrls(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the PurgeByUrlsOptions model
				purgeByUrlsOptionsModel := new(cachingapiv1.PurgeByUrlsOptions)
				purgeByUrlsOptionsModel.Files = []string{"http://www.example.com/cat_picture.jpg"}
				purgeByUrlsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.PurgeByUrls(purgeByUrlsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke PurgeByUrls with error: Operation request error`, func() {
				testService, testServiceErr := cachingapiv1.NewCachingApiV1(&cachingapiv1.CachingApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the PurgeByUrlsOptions model
				purgeByUrlsOptionsModel := new(cachingapiv1.PurgeByUrlsOptions)
				purgeByUrlsOptionsModel.Files = []string{"http://www.example.com/cat_picture.jpg"}
				purgeByUrlsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.PurgeByUrls(purgeByUrlsOptionsModel)
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
	Describe(`PurgeByCacheTags(purgeByCacheTagsOptions *PurgeByCacheTagsOptions) - Operation response error`, func() {
		crn := "testString"
		zoneID := "testString"
		purgeByCacheTagsPath := "/v1/testString/zones/testString/purge_cache/purge_by_cache_tags"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(purgeByCacheTagsPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke PurgeByCacheTags with error: Operation response processing error`, func() {
				testService, testServiceErr := cachingapiv1.NewCachingApiV1(&cachingapiv1.CachingApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the PurgeByCacheTagsOptions model
				purgeByCacheTagsOptionsModel := new(cachingapiv1.PurgeByCacheTagsOptions)
				purgeByCacheTagsOptionsModel.Tags = []string{"some-tag"}
				purgeByCacheTagsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.PurgeByCacheTags(purgeByCacheTagsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`PurgeByCacheTags(purgeByCacheTagsOptions *PurgeByCacheTagsOptions)`, func() {
		crn := "testString"
		zoneID := "testString"
		purgeByCacheTagsPath := "/v1/testString/zones/testString/purge_cache/purge_by_cache_tags"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(purgeByCacheTagsPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "62d26b178b67c0eda0613891f3f51b0a"}}`)
				}))
			})
			It(`Invoke PurgeByCacheTags successfully`, func() {
				testService, testServiceErr := cachingapiv1.NewCachingApiV1(&cachingapiv1.CachingApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.PurgeByCacheTags(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the PurgeByCacheTagsOptions model
				purgeByCacheTagsOptionsModel := new(cachingapiv1.PurgeByCacheTagsOptions)
				purgeByCacheTagsOptionsModel.Tags = []string{"some-tag"}
				purgeByCacheTagsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.PurgeByCacheTags(purgeByCacheTagsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke PurgeByCacheTags with error: Operation request error`, func() {
				testService, testServiceErr := cachingapiv1.NewCachingApiV1(&cachingapiv1.CachingApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the PurgeByCacheTagsOptions model
				purgeByCacheTagsOptionsModel := new(cachingapiv1.PurgeByCacheTagsOptions)
				purgeByCacheTagsOptionsModel.Tags = []string{"some-tag"}
				purgeByCacheTagsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.PurgeByCacheTags(purgeByCacheTagsOptionsModel)
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
	Describe(`PurgeByHosts(purgeByHostsOptions *PurgeByHostsOptions) - Operation response error`, func() {
		crn := "testString"
		zoneID := "testString"
		purgeByHostsPath := "/v1/testString/zones/testString/purge_cache/purge_by_hosts"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(purgeByHostsPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke PurgeByHosts with error: Operation response processing error`, func() {
				testService, testServiceErr := cachingapiv1.NewCachingApiV1(&cachingapiv1.CachingApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the PurgeByHostsOptions model
				purgeByHostsOptionsModel := new(cachingapiv1.PurgeByHostsOptions)
				purgeByHostsOptionsModel.Hosts = []string{"www.example.com"}
				purgeByHostsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.PurgeByHosts(purgeByHostsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`PurgeByHosts(purgeByHostsOptions *PurgeByHostsOptions)`, func() {
		crn := "testString"
		zoneID := "testString"
		purgeByHostsPath := "/v1/testString/zones/testString/purge_cache/purge_by_hosts"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(purgeByHostsPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "62d26b178b67c0eda0613891f3f51b0a"}}`)
				}))
			})
			It(`Invoke PurgeByHosts successfully`, func() {
				testService, testServiceErr := cachingapiv1.NewCachingApiV1(&cachingapiv1.CachingApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.PurgeByHosts(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the PurgeByHostsOptions model
				purgeByHostsOptionsModel := new(cachingapiv1.PurgeByHostsOptions)
				purgeByHostsOptionsModel.Hosts = []string{"www.example.com"}
				purgeByHostsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.PurgeByHosts(purgeByHostsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke PurgeByHosts with error: Operation request error`, func() {
				testService, testServiceErr := cachingapiv1.NewCachingApiV1(&cachingapiv1.CachingApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the PurgeByHostsOptions model
				purgeByHostsOptionsModel := new(cachingapiv1.PurgeByHostsOptions)
				purgeByHostsOptionsModel.Hosts = []string{"www.example.com"}
				purgeByHostsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.PurgeByHosts(purgeByHostsOptionsModel)
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
	Describe(`GetBrowserCacheTTL(getBrowserCacheTtlOptions *GetBrowserCacheTtlOptions) - Operation response error`, func() {
		crn := "testString"
		zoneID := "testString"
		getBrowserCacheTTLPath := "/v1/testString/zones/testString/settings/browser_cache_ttl"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getBrowserCacheTTLPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetBrowserCacheTTL with error: Operation response processing error`, func() {
				testService, testServiceErr := cachingapiv1.NewCachingApiV1(&cachingapiv1.CachingApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetBrowserCacheTtlOptions model
				getBrowserCacheTtlOptionsModel := new(cachingapiv1.GetBrowserCacheTtlOptions)
				getBrowserCacheTtlOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.GetBrowserCacheTTL(getBrowserCacheTtlOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetBrowserCacheTTL(getBrowserCacheTtlOptions *GetBrowserCacheTtlOptions)`, func() {
		crn := "testString"
		zoneID := "testString"
		getBrowserCacheTTLPath := "/v1/testString/zones/testString/settings/browser_cache_ttl"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getBrowserCacheTTLPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "browser_cache_ttl", "value": 14400, "editable": true, "modified_on": "2014-01-01T05:20:00.12345Z"}}`)
				}))
			})
			It(`Invoke GetBrowserCacheTTL successfully`, func() {
				testService, testServiceErr := cachingapiv1.NewCachingApiV1(&cachingapiv1.CachingApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetBrowserCacheTTL(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetBrowserCacheTtlOptions model
				getBrowserCacheTtlOptionsModel := new(cachingapiv1.GetBrowserCacheTtlOptions)
				getBrowserCacheTtlOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetBrowserCacheTTL(getBrowserCacheTtlOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetBrowserCacheTTL with error: Operation request error`, func() {
				testService, testServiceErr := cachingapiv1.NewCachingApiV1(&cachingapiv1.CachingApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetBrowserCacheTtlOptions model
				getBrowserCacheTtlOptionsModel := new(cachingapiv1.GetBrowserCacheTtlOptions)
				getBrowserCacheTtlOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.GetBrowserCacheTTL(getBrowserCacheTtlOptionsModel)
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
	Describe(`UpdateBrowserCacheTTL(updateBrowserCacheTtlOptions *UpdateBrowserCacheTtlOptions) - Operation response error`, func() {
		crn := "testString"
		zoneID := "testString"
		updateBrowserCacheTTLPath := "/v1/testString/zones/testString/settings/browser_cache_ttl"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(updateBrowserCacheTTLPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateBrowserCacheTTL with error: Operation response processing error`, func() {
				testService, testServiceErr := cachingapiv1.NewCachingApiV1(&cachingapiv1.CachingApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the UpdateBrowserCacheTtlOptions model
				updateBrowserCacheTtlOptionsModel := new(cachingapiv1.UpdateBrowserCacheTtlOptions)
				updateBrowserCacheTtlOptionsModel.Value = core.Int64Ptr(int64(14400))
				updateBrowserCacheTtlOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.UpdateBrowserCacheTTL(updateBrowserCacheTtlOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateBrowserCacheTTL(updateBrowserCacheTtlOptions *UpdateBrowserCacheTtlOptions)`, func() {
		crn := "testString"
		zoneID := "testString"
		updateBrowserCacheTTLPath := "/v1/testString/zones/testString/settings/browser_cache_ttl"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(updateBrowserCacheTTLPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "browser_cache_ttl", "value": 14400, "editable": true, "modified_on": "2014-01-01T05:20:00.12345Z"}}`)
				}))
			})
			It(`Invoke UpdateBrowserCacheTTL successfully`, func() {
				testService, testServiceErr := cachingapiv1.NewCachingApiV1(&cachingapiv1.CachingApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.UpdateBrowserCacheTTL(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateBrowserCacheTtlOptions model
				updateBrowserCacheTtlOptionsModel := new(cachingapiv1.UpdateBrowserCacheTtlOptions)
				updateBrowserCacheTtlOptionsModel.Value = core.Int64Ptr(int64(14400))
				updateBrowserCacheTtlOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.UpdateBrowserCacheTTL(updateBrowserCacheTtlOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke UpdateBrowserCacheTTL with error: Operation request error`, func() {
				testService, testServiceErr := cachingapiv1.NewCachingApiV1(&cachingapiv1.CachingApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the UpdateBrowserCacheTtlOptions model
				updateBrowserCacheTtlOptionsModel := new(cachingapiv1.UpdateBrowserCacheTtlOptions)
				updateBrowserCacheTtlOptionsModel.Value = core.Int64Ptr(int64(14400))
				updateBrowserCacheTtlOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.UpdateBrowserCacheTTL(updateBrowserCacheTtlOptionsModel)
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
	Describe(`GetDevelopmentMode(getDevelopmentModeOptions *GetDevelopmentModeOptions) - Operation response error`, func() {
		crn := "testString"
		zoneID := "testString"
		getDevelopmentModePath := "/v1/testString/zones/testString/settings/development_mode"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getDevelopmentModePath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetDevelopmentMode with error: Operation response processing error`, func() {
				testService, testServiceErr := cachingapiv1.NewCachingApiV1(&cachingapiv1.CachingApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetDevelopmentModeOptions model
				getDevelopmentModeOptionsModel := new(cachingapiv1.GetDevelopmentModeOptions)
				getDevelopmentModeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.GetDevelopmentMode(getDevelopmentModeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetDevelopmentMode(getDevelopmentModeOptions *GetDevelopmentModeOptions)`, func() {
		crn := "testString"
		zoneID := "testString"
		getDevelopmentModePath := "/v1/testString/zones/testString/settings/development_mode"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getDevelopmentModePath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "development_mode", "value": "off", "editable": true, "modified_on": "2014-01-01T05:20:00.12345Z"}}`)
				}))
			})
			It(`Invoke GetDevelopmentMode successfully`, func() {
				testService, testServiceErr := cachingapiv1.NewCachingApiV1(&cachingapiv1.CachingApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetDevelopmentMode(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetDevelopmentModeOptions model
				getDevelopmentModeOptionsModel := new(cachingapiv1.GetDevelopmentModeOptions)
				getDevelopmentModeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetDevelopmentMode(getDevelopmentModeOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetDevelopmentMode with error: Operation request error`, func() {
				testService, testServiceErr := cachingapiv1.NewCachingApiV1(&cachingapiv1.CachingApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetDevelopmentModeOptions model
				getDevelopmentModeOptionsModel := new(cachingapiv1.GetDevelopmentModeOptions)
				getDevelopmentModeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.GetDevelopmentMode(getDevelopmentModeOptionsModel)
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
	Describe(`UpdateDevelopmentMode(updateDevelopmentModeOptions *UpdateDevelopmentModeOptions) - Operation response error`, func() {
		crn := "testString"
		zoneID := "testString"
		updateDevelopmentModePath := "/v1/testString/zones/testString/settings/development_mode"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(updateDevelopmentModePath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateDevelopmentMode with error: Operation response processing error`, func() {
				testService, testServiceErr := cachingapiv1.NewCachingApiV1(&cachingapiv1.CachingApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the UpdateDevelopmentModeOptions model
				updateDevelopmentModeOptionsModel := new(cachingapiv1.UpdateDevelopmentModeOptions)
				updateDevelopmentModeOptionsModel.Value = core.StringPtr("off")
				updateDevelopmentModeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.UpdateDevelopmentMode(updateDevelopmentModeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateDevelopmentMode(updateDevelopmentModeOptions *UpdateDevelopmentModeOptions)`, func() {
		crn := "testString"
		zoneID := "testString"
		updateDevelopmentModePath := "/v1/testString/zones/testString/settings/development_mode"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(updateDevelopmentModePath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "development_mode", "value": "off", "editable": true, "modified_on": "2014-01-01T05:20:00.12345Z"}}`)
				}))
			})
			It(`Invoke UpdateDevelopmentMode successfully`, func() {
				testService, testServiceErr := cachingapiv1.NewCachingApiV1(&cachingapiv1.CachingApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.UpdateDevelopmentMode(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateDevelopmentModeOptions model
				updateDevelopmentModeOptionsModel := new(cachingapiv1.UpdateDevelopmentModeOptions)
				updateDevelopmentModeOptionsModel.Value = core.StringPtr("off")
				updateDevelopmentModeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.UpdateDevelopmentMode(updateDevelopmentModeOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke UpdateDevelopmentMode with error: Operation request error`, func() {
				testService, testServiceErr := cachingapiv1.NewCachingApiV1(&cachingapiv1.CachingApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the UpdateDevelopmentModeOptions model
				updateDevelopmentModeOptionsModel := new(cachingapiv1.UpdateDevelopmentModeOptions)
				updateDevelopmentModeOptionsModel.Value = core.StringPtr("off")
				updateDevelopmentModeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.UpdateDevelopmentMode(updateDevelopmentModeOptionsModel)
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
	Describe(`GetQueryStringSort(getQueryStringSortOptions *GetQueryStringSortOptions) - Operation response error`, func() {
		crn := "testString"
		zoneID := "testString"
		getQueryStringSortPath := "/v1/testString/zones/testString/settings/sort_query_string_for_cache"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getQueryStringSortPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetQueryStringSort with error: Operation response processing error`, func() {
				testService, testServiceErr := cachingapiv1.NewCachingApiV1(&cachingapiv1.CachingApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetQueryStringSortOptions model
				getQueryStringSortOptionsModel := new(cachingapiv1.GetQueryStringSortOptions)
				getQueryStringSortOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.GetQueryStringSort(getQueryStringSortOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetQueryStringSort(getQueryStringSortOptions *GetQueryStringSortOptions)`, func() {
		crn := "testString"
		zoneID := "testString"
		getQueryStringSortPath := "/v1/testString/zones/testString/settings/sort_query_string_for_cache"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getQueryStringSortPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "sort_query_string_for_cache", "value": "off", "editable": true, "modified_on": "2014-01-01T05:20:00.12345Z"}}`)
				}))
			})
			It(`Invoke GetQueryStringSort successfully`, func() {
				testService, testServiceErr := cachingapiv1.NewCachingApiV1(&cachingapiv1.CachingApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetQueryStringSort(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetQueryStringSortOptions model
				getQueryStringSortOptionsModel := new(cachingapiv1.GetQueryStringSortOptions)
				getQueryStringSortOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetQueryStringSort(getQueryStringSortOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetQueryStringSort with error: Operation request error`, func() {
				testService, testServiceErr := cachingapiv1.NewCachingApiV1(&cachingapiv1.CachingApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetQueryStringSortOptions model
				getQueryStringSortOptionsModel := new(cachingapiv1.GetQueryStringSortOptions)
				getQueryStringSortOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.GetQueryStringSort(getQueryStringSortOptionsModel)
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
	Describe(`UpdateQueryStringSort(updateQueryStringSortOptions *UpdateQueryStringSortOptions) - Operation response error`, func() {
		crn := "testString"
		zoneID := "testString"
		updateQueryStringSortPath := "/v1/testString/zones/testString/settings/sort_query_string_for_cache"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(updateQueryStringSortPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateQueryStringSort with error: Operation response processing error`, func() {
				testService, testServiceErr := cachingapiv1.NewCachingApiV1(&cachingapiv1.CachingApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the UpdateQueryStringSortOptions model
				updateQueryStringSortOptionsModel := new(cachingapiv1.UpdateQueryStringSortOptions)
				updateQueryStringSortOptionsModel.Value = core.StringPtr("off")
				updateQueryStringSortOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.UpdateQueryStringSort(updateQueryStringSortOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateQueryStringSort(updateQueryStringSortOptions *UpdateQueryStringSortOptions)`, func() {
		crn := "testString"
		zoneID := "testString"
		updateQueryStringSortPath := "/v1/testString/zones/testString/settings/sort_query_string_for_cache"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(updateQueryStringSortPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "sort_query_string_for_cache", "value": "off", "editable": true, "modified_on": "2014-01-01T05:20:00.12345Z"}}`)
				}))
			})
			It(`Invoke UpdateQueryStringSort successfully`, func() {
				testService, testServiceErr := cachingapiv1.NewCachingApiV1(&cachingapiv1.CachingApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.UpdateQueryStringSort(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateQueryStringSortOptions model
				updateQueryStringSortOptionsModel := new(cachingapiv1.UpdateQueryStringSortOptions)
				updateQueryStringSortOptionsModel.Value = core.StringPtr("off")
				updateQueryStringSortOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.UpdateQueryStringSort(updateQueryStringSortOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke UpdateQueryStringSort with error: Operation request error`, func() {
				testService, testServiceErr := cachingapiv1.NewCachingApiV1(&cachingapiv1.CachingApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the UpdateQueryStringSortOptions model
				updateQueryStringSortOptionsModel := new(cachingapiv1.UpdateQueryStringSortOptions)
				updateQueryStringSortOptionsModel.Value = core.StringPtr("off")
				updateQueryStringSortOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.UpdateQueryStringSort(updateQueryStringSortOptionsModel)
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
	Describe(`Service constructor tests`, func() {
		crn := "testString"
		zoneID := "testString"
		It(`Instantiate service client`, func() {
			testService, testServiceErr := cachingapiv1.NewCachingApiV1(&cachingapiv1.CachingApiV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
				Crn:           core.StringPtr(crn),
				ZoneID:        core.StringPtr(zoneID),
			})
			Expect(testService).ToNot(BeNil())
			Expect(testServiceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			testService, testServiceErr := cachingapiv1.NewCachingApiV1(&cachingapiv1.CachingApiV1Options{
				URL:    "{BAD_URL_STRING",
				Crn:    core.StringPtr(crn),
				ZoneID: core.StringPtr(zoneID),
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			testService, testServiceErr := cachingapiv1.NewCachingApiV1(&cachingapiv1.CachingApiV1Options{
				URL:    "https://cachingapiv1/api",
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
			testService, testServiceErr := cachingapiv1.NewCachingApiV1(&cachingapiv1.CachingApiV1Options{})
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
				"CACHING_API_URL":       "https://cachingapiv1/api",
				"CACHING_API_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := cachingapiv1.NewCachingApiV1UsingExternalConfig(&cachingapiv1.CachingApiV1Options{
					Crn:    core.StringPtr(crn),
					ZoneID: core.StringPtr(zoneID),
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := cachingapiv1.NewCachingApiV1UsingExternalConfig(&cachingapiv1.CachingApiV1Options{
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
				testService, testServiceErr := cachingapiv1.NewCachingApiV1UsingExternalConfig(&cachingapiv1.CachingApiV1Options{
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
				"CACHING_API_URL":       "https://cachingapiv1/api",
				"CACHING_API_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := cachingapiv1.NewCachingApiV1UsingExternalConfig(&cachingapiv1.CachingApiV1Options{
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
				"CACHING_API_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := cachingapiv1.NewCachingApiV1UsingExternalConfig(&cachingapiv1.CachingApiV1Options{
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
	Describe(`GetCacheLevel(getCacheLevelOptions *GetCacheLevelOptions) - Operation response error`, func() {
		crn := "testString"
		zoneID := "testString"
		getCacheLevelPath := "/v1/testString/zones/testString/settings/cache_level"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getCacheLevelPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetCacheLevel with error: Operation response processing error`, func() {
				testService, testServiceErr := cachingapiv1.NewCachingApiV1(&cachingapiv1.CachingApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetCacheLevelOptions model
				getCacheLevelOptionsModel := new(cachingapiv1.GetCacheLevelOptions)
				getCacheLevelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.GetCacheLevel(getCacheLevelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetCacheLevel(getCacheLevelOptions *GetCacheLevelOptions)`, func() {
		crn := "testString"
		zoneID := "testString"
		getCacheLevelPath := "/v1/testString/zones/testString/settings/cache_level"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getCacheLevelPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "cache_level", "value": "aggressive", "editable": true, "modified_on": "2014-01-01T05:20:00.12345Z"}}`)
				}))
			})
			It(`Invoke GetCacheLevel successfully`, func() {
				testService, testServiceErr := cachingapiv1.NewCachingApiV1(&cachingapiv1.CachingApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetCacheLevel(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetCacheLevelOptions model
				getCacheLevelOptionsModel := new(cachingapiv1.GetCacheLevelOptions)
				getCacheLevelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetCacheLevel(getCacheLevelOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetCacheLevel with error: Operation request error`, func() {
				testService, testServiceErr := cachingapiv1.NewCachingApiV1(&cachingapiv1.CachingApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetCacheLevelOptions model
				getCacheLevelOptionsModel := new(cachingapiv1.GetCacheLevelOptions)
				getCacheLevelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.GetCacheLevel(getCacheLevelOptionsModel)
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
	Describe(`UpdateCacheLevel(updateCacheLevelOptions *UpdateCacheLevelOptions) - Operation response error`, func() {
		crn := "testString"
		zoneID := "testString"
		updateCacheLevelPath := "/v1/testString/zones/testString/settings/cache_level"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(updateCacheLevelPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateCacheLevel with error: Operation response processing error`, func() {
				testService, testServiceErr := cachingapiv1.NewCachingApiV1(&cachingapiv1.CachingApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the UpdateCacheLevelOptions model
				updateCacheLevelOptionsModel := new(cachingapiv1.UpdateCacheLevelOptions)
				updateCacheLevelOptionsModel.Value = core.StringPtr("aggressive")
				updateCacheLevelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.UpdateCacheLevel(updateCacheLevelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateCacheLevel(updateCacheLevelOptions *UpdateCacheLevelOptions)`, func() {
		crn := "testString"
		zoneID := "testString"
		updateCacheLevelPath := "/v1/testString/zones/testString/settings/cache_level"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(updateCacheLevelPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "cache_level", "value": "aggressive", "editable": true, "modified_on": "2014-01-01T05:20:00.12345Z"}}`)
				}))
			})
			It(`Invoke UpdateCacheLevel successfully`, func() {
				testService, testServiceErr := cachingapiv1.NewCachingApiV1(&cachingapiv1.CachingApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.UpdateCacheLevel(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateCacheLevelOptions model
				updateCacheLevelOptionsModel := new(cachingapiv1.UpdateCacheLevelOptions)
				updateCacheLevelOptionsModel.Value = core.StringPtr("aggressive")
				updateCacheLevelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.UpdateCacheLevel(updateCacheLevelOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke UpdateCacheLevel with error: Operation request error`, func() {
				testService, testServiceErr := cachingapiv1.NewCachingApiV1(&cachingapiv1.CachingApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn:           core.StringPtr(crn),
					ZoneID:        core.StringPtr(zoneID),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the UpdateCacheLevelOptions model
				updateCacheLevelOptionsModel := new(cachingapiv1.UpdateCacheLevelOptions)
				updateCacheLevelOptionsModel.Value = core.StringPtr("aggressive")
				updateCacheLevelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.UpdateCacheLevel(updateCacheLevelOptionsModel)
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
			testService, _ := cachingapiv1.NewCachingApiV1(&cachingapiv1.CachingApiV1Options{
				URL:           "http://cachingapiv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
				Crn:           core.StringPtr(crn),
				ZoneID:        core.StringPtr(zoneID),
			})
			It(`Invoke NewGetBrowserCacheTtlOptions successfully`, func() {
				// Construct an instance of the GetBrowserCacheTtlOptions model
				getBrowserCacheTtlOptionsModel := testService.NewGetBrowserCacheTtlOptions()
				getBrowserCacheTtlOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getBrowserCacheTtlOptionsModel).ToNot(BeNil())
				Expect(getBrowserCacheTtlOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetCacheLevelOptions successfully`, func() {
				// Construct an instance of the GetCacheLevelOptions model
				getCacheLevelOptionsModel := testService.NewGetCacheLevelOptions()
				getCacheLevelOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getCacheLevelOptionsModel).ToNot(BeNil())
				Expect(getCacheLevelOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetDevelopmentModeOptions successfully`, func() {
				// Construct an instance of the GetDevelopmentModeOptions model
				getDevelopmentModeOptionsModel := testService.NewGetDevelopmentModeOptions()
				getDevelopmentModeOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getDevelopmentModeOptionsModel).ToNot(BeNil())
				Expect(getDevelopmentModeOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetQueryStringSortOptions successfully`, func() {
				// Construct an instance of the GetQueryStringSortOptions model
				getQueryStringSortOptionsModel := testService.NewGetQueryStringSortOptions()
				getQueryStringSortOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getQueryStringSortOptionsModel).ToNot(BeNil())
				Expect(getQueryStringSortOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPurgeAllOptions successfully`, func() {
				// Construct an instance of the PurgeAllOptions model
				purgeAllOptionsModel := testService.NewPurgeAllOptions()
				purgeAllOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(purgeAllOptionsModel).ToNot(BeNil())
				Expect(purgeAllOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPurgeByCacheTagsOptions successfully`, func() {
				// Construct an instance of the PurgeByCacheTagsOptions model
				purgeByCacheTagsOptionsModel := testService.NewPurgeByCacheTagsOptions()
				purgeByCacheTagsOptionsModel.SetTags([]string{"some-tag"})
				purgeByCacheTagsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(purgeByCacheTagsOptionsModel).ToNot(BeNil())
				Expect(purgeByCacheTagsOptionsModel.Tags).To(Equal([]string{"some-tag"}))
				Expect(purgeByCacheTagsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPurgeByHostsOptions successfully`, func() {
				// Construct an instance of the PurgeByHostsOptions model
				purgeByHostsOptionsModel := testService.NewPurgeByHostsOptions()
				purgeByHostsOptionsModel.SetHosts([]string{"www.example.com"})
				purgeByHostsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(purgeByHostsOptionsModel).ToNot(BeNil())
				Expect(purgeByHostsOptionsModel.Hosts).To(Equal([]string{"www.example.com"}))
				Expect(purgeByHostsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPurgeByUrlsOptions successfully`, func() {
				// Construct an instance of the PurgeByUrlsOptions model
				purgeByUrlsOptionsModel := testService.NewPurgeByUrlsOptions()
				purgeByUrlsOptionsModel.SetFiles([]string{"http://www.example.com/cat_picture.jpg"})
				purgeByUrlsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(purgeByUrlsOptionsModel).ToNot(BeNil())
				Expect(purgeByUrlsOptionsModel.Files).To(Equal([]string{"http://www.example.com/cat_picture.jpg"}))
				Expect(purgeByUrlsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateBrowserCacheTtlOptions successfully`, func() {
				// Construct an instance of the UpdateBrowserCacheTtlOptions model
				updateBrowserCacheTtlOptionsModel := testService.NewUpdateBrowserCacheTtlOptions()
				updateBrowserCacheTtlOptionsModel.SetValue(int64(14400))
				updateBrowserCacheTtlOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateBrowserCacheTtlOptionsModel).ToNot(BeNil())
				Expect(updateBrowserCacheTtlOptionsModel.Value).To(Equal(core.Int64Ptr(int64(14400))))
				Expect(updateBrowserCacheTtlOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateCacheLevelOptions successfully`, func() {
				// Construct an instance of the UpdateCacheLevelOptions model
				updateCacheLevelOptionsModel := testService.NewUpdateCacheLevelOptions()
				updateCacheLevelOptionsModel.SetValue("aggressive")
				updateCacheLevelOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateCacheLevelOptionsModel).ToNot(BeNil())
				Expect(updateCacheLevelOptionsModel.Value).To(Equal(core.StringPtr("aggressive")))
				Expect(updateCacheLevelOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateDevelopmentModeOptions successfully`, func() {
				// Construct an instance of the UpdateDevelopmentModeOptions model
				updateDevelopmentModeOptionsModel := testService.NewUpdateDevelopmentModeOptions()
				updateDevelopmentModeOptionsModel.SetValue("off")
				updateDevelopmentModeOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateDevelopmentModeOptionsModel).ToNot(BeNil())
				Expect(updateDevelopmentModeOptionsModel.Value).To(Equal(core.StringPtr("off")))
				Expect(updateDevelopmentModeOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateQueryStringSortOptions successfully`, func() {
				// Construct an instance of the UpdateQueryStringSortOptions model
				updateQueryStringSortOptionsModel := testService.NewUpdateQueryStringSortOptions()
				updateQueryStringSortOptionsModel.SetValue("off")
				updateQueryStringSortOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateQueryStringSortOptionsModel).ToNot(BeNil())
				Expect(updateQueryStringSortOptionsModel.Value).To(Equal(core.StringPtr("off")))
				Expect(updateQueryStringSortOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
