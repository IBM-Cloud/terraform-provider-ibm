/**
 * (C) Copyright IBM Corp. 2025.
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

package distributionlistapiv1_test

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	distributionlistapiv1 "github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/distlistapi"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`DistributionListApiV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			distributionListApiService, serviceErr := distributionlistapiv1.NewDistributionListApiV1(&distributionlistapiv1.DistributionListApiV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(distributionListApiService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			distributionListApiService, serviceErr := distributionlistapiv1.NewDistributionListApiV1(&distributionlistapiv1.DistributionListApiV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(distributionListApiService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			distributionListApiService, serviceErr := distributionlistapiv1.NewDistributionListApiV1(&distributionlistapiv1.DistributionListApiV1Options{
				URL: "https://distributionlistapiv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "", // pragma: allowlist secret
					Password: "", // pragma: allowlist secret
				},
			})
			Expect(distributionListApiService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"DISTRIBUTION_LIST_API_URL":       "https://distributionlistapiv1/api",
				"DISTRIBUTION_LIST_API_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				distributionListApiService, serviceErr := distributionlistapiv1.NewDistributionListApiV1UsingExternalConfig(&distributionlistapiv1.DistributionListApiV1Options{})
				Expect(distributionListApiService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := distributionListApiService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != distributionListApiService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(distributionListApiService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(distributionListApiService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				distributionListApiService, serviceErr := distributionlistapiv1.NewDistributionListApiV1UsingExternalConfig(&distributionlistapiv1.DistributionListApiV1Options{
					URL: "https://testService/api",
				})
				Expect(distributionListApiService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(distributionListApiService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := distributionListApiService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != distributionListApiService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(distributionListApiService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(distributionListApiService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				distributionListApiService, serviceErr := distributionlistapiv1.NewDistributionListApiV1UsingExternalConfig(&distributionlistapiv1.DistributionListApiV1Options{})
				err := distributionListApiService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(distributionListApiService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(distributionListApiService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := distributionListApiService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != distributionListApiService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(distributionListApiService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(distributionListApiService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"DISTRIBUTION_LIST_API_URL":       "https://distributionlistapiv1/api",
				"DISTRIBUTION_LIST_API_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			distributionListApiService, serviceErr := distributionlistapiv1.NewDistributionListApiV1UsingExternalConfig(&distributionlistapiv1.DistributionListApiV1Options{})

			It(`Instantiate service client with error`, func() {
				Expect(distributionListApiService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"DISTRIBUTION_LIST_API_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			distributionListApiService, serviceErr := distributionlistapiv1.NewDistributionListApiV1UsingExternalConfig(&distributionlistapiv1.DistributionListApiV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(distributionListApiService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = distributionlistapiv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`GetAllDestinationEntries(getAllDestinationEntriesOptions *GetAllDestinationEntriesOptions) - Operation response error`, func() {
		getAllDestinationEntriesPath := "/notification-api/v1/distribution_lists/a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6/destinations" // pragma: allowlist secret
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAllDestinationEntriesPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetAllDestinationEntries with error: Operation response processing error`, func() {
				distributionListApiService, serviceErr := distributionlistapiv1.NewDistributionListApiV1(&distributionlistapiv1.DistributionListApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(distributionListApiService).ToNot(BeNil())

				// Construct an instance of the GetAllDestinationEntriesOptions model
				getAllDestinationEntriesOptionsModel := new(distributionlistapiv1.GetAllDestinationEntriesOptions)
				getAllDestinationEntriesOptionsModel.AccountID = core.StringPtr("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6") // pragma: allowlist secret
				getAllDestinationEntriesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := distributionListApiService.GetAllDestinationEntries(getAllDestinationEntriesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				distributionListApiService.EnableRetries(0, 0)
				result, response, operationErr = distributionListApiService.GetAllDestinationEntries(getAllDestinationEntriesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetAllDestinationEntries(getAllDestinationEntriesOptions *GetAllDestinationEntriesOptions)`, func() {
		getAllDestinationEntriesPath := "/notification-api/v1/distribution_lists/a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6/destinations" // pragma: allowlist secret
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAllDestinationEntriesPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `[{"id": "12345678-1234-1234-1234-123456789012", "destination_type": "event_notifications"}]`)
				}))
			})
			It(`Invoke GetAllDestinationEntries successfully with retries`, func() {
				distributionListApiService, serviceErr := distributionlistapiv1.NewDistributionListApiV1(&distributionlistapiv1.DistributionListApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(distributionListApiService).ToNot(BeNil())
				distributionListApiService.EnableRetries(0, 0)

				// Construct an instance of the GetAllDestinationEntriesOptions model
				getAllDestinationEntriesOptionsModel := new(distributionlistapiv1.GetAllDestinationEntriesOptions)
				getAllDestinationEntriesOptionsModel.AccountID = core.StringPtr("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6") // pragma: allowlist secret
				getAllDestinationEntriesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := distributionListApiService.GetAllDestinationEntriesWithContext(ctx, getAllDestinationEntriesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				distributionListApiService.DisableRetries()
				result, response, operationErr := distributionListApiService.GetAllDestinationEntries(getAllDestinationEntriesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = distributionListApiService.GetAllDestinationEntriesWithContext(ctx, getAllDestinationEntriesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAllDestinationEntriesPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `[{"id": "12345678-1234-1234-1234-123456789012", "destination_type": "event_notifications"}]`)
				}))
			})
			It(`Invoke GetAllDestinationEntries successfully`, func() {
				distributionListApiService, serviceErr := distributionlistapiv1.NewDistributionListApiV1(&distributionlistapiv1.DistributionListApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(distributionListApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := distributionListApiService.GetAllDestinationEntries(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetAllDestinationEntriesOptions model
				getAllDestinationEntriesOptionsModel := new(distributionlistapiv1.GetAllDestinationEntriesOptions)
				getAllDestinationEntriesOptionsModel.AccountID = core.StringPtr("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6") // pragma: allowlist secret
				getAllDestinationEntriesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = distributionListApiService.GetAllDestinationEntries(getAllDestinationEntriesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetAllDestinationEntries with error: Operation validation and request error`, func() {
				distributionListApiService, serviceErr := distributionlistapiv1.NewDistributionListApiV1(&distributionlistapiv1.DistributionListApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(distributionListApiService).ToNot(BeNil())

				// Construct an instance of the GetAllDestinationEntriesOptions model
				getAllDestinationEntriesOptionsModel := new(distributionlistapiv1.GetAllDestinationEntriesOptions)
				getAllDestinationEntriesOptionsModel.AccountID = core.StringPtr("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6") // pragma: allowlist secret
				getAllDestinationEntriesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := distributionListApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := distributionListApiService.GetAllDestinationEntries(getAllDestinationEntriesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetAllDestinationEntriesOptions model with no property values
				getAllDestinationEntriesOptionsModelNew := new(distributionlistapiv1.GetAllDestinationEntriesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = distributionListApiService.GetAllDestinationEntries(getAllDestinationEntriesOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetAllDestinationEntries successfully`, func() {
				distributionListApiService, serviceErr := distributionlistapiv1.NewDistributionListApiV1(&distributionlistapiv1.DistributionListApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(distributionListApiService).ToNot(BeNil())

				// Construct an instance of the GetAllDestinationEntriesOptions model
				getAllDestinationEntriesOptionsModel := new(distributionlistapiv1.GetAllDestinationEntriesOptions)
				getAllDestinationEntriesOptionsModel.AccountID = core.StringPtr("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6") // pragma: allowlist secret
				getAllDestinationEntriesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := distributionListApiService.GetAllDestinationEntries(getAllDestinationEntriesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`AddDestinationEntry(addDestinationEntryOptions *AddDestinationEntryOptions) - Operation response error`, func() {
		addDestinationEntryPath := "/notification-api/v1/distribution_lists/a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6/destinations" // pragma: allowlist secret
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(addDestinationEntryPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke AddDestinationEntry with error: Operation response processing error`, func() {
				distributionListApiService, serviceErr := distributionlistapiv1.NewDistributionListApiV1(&distributionlistapiv1.DistributionListApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(distributionListApiService).ToNot(BeNil())

				// Construct an instance of the AddDestinationEntryRequestEventNotificationDestination model
				addDestinationEntryRequestModel := new(distributionlistapiv1.AddDestinationEntryRequestEventNotificationDestination)
				addDestinationEntryRequestModel.ID = CreateMockUUID("12345678-1234-1234-1234-123456789012")
				addDestinationEntryRequestModel.DestinationType = core.StringPtr("event_notifications")

				// Construct an instance of the AddDestinationEntryOptions model
				addDestinationEntryOptionsModel := new(distributionlistapiv1.AddDestinationEntryOptions)
				addDestinationEntryOptionsModel.AccountID = core.StringPtr("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6") // pragma: allowlist secret
				addDestinationEntryOptionsModel.AddDestinationEntryRequest = addDestinationEntryRequestModel
				addDestinationEntryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := distributionListApiService.AddDestinationEntry(addDestinationEntryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				distributionListApiService.EnableRetries(0, 0)
				result, response, operationErr = distributionListApiService.AddDestinationEntry(addDestinationEntryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`AddDestinationEntry(addDestinationEntryOptions *AddDestinationEntryOptions)`, func() {
		addDestinationEntryPath := "/notification-api/v1/distribution_lists/a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6/destinations" // pragma: allowlist secret
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(addDestinationEntryPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "12345678-1234-1234-1234-123456789012", "destination_type": "event_notifications"}`)
				}))
			})
			It(`Invoke AddDestinationEntry successfully with retries`, func() {
				distributionListApiService, serviceErr := distributionlistapiv1.NewDistributionListApiV1(&distributionlistapiv1.DistributionListApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(distributionListApiService).ToNot(BeNil())
				distributionListApiService.EnableRetries(0, 0)

				// Construct an instance of the AddDestinationEntryRequestEventNotificationDestination model
				addDestinationEntryRequestModel := new(distributionlistapiv1.AddDestinationEntryRequestEventNotificationDestination)
				addDestinationEntryRequestModel.ID = CreateMockUUID("12345678-1234-1234-1234-123456789012")
				addDestinationEntryRequestModel.DestinationType = core.StringPtr("event_notifications")

				// Construct an instance of the AddDestinationEntryOptions model
				addDestinationEntryOptionsModel := new(distributionlistapiv1.AddDestinationEntryOptions)
				addDestinationEntryOptionsModel.AccountID = core.StringPtr("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6") // pragma: allowlist secret
				addDestinationEntryOptionsModel.AddDestinationEntryRequest = addDestinationEntryRequestModel
				addDestinationEntryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := distributionListApiService.AddDestinationEntryWithContext(ctx, addDestinationEntryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				distributionListApiService.DisableRetries()
				result, response, operationErr := distributionListApiService.AddDestinationEntry(addDestinationEntryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = distributionListApiService.AddDestinationEntryWithContext(ctx, addDestinationEntryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(addDestinationEntryPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "12345678-1234-1234-1234-123456789012", "destination_type": "event_notifications"}`)
				}))
			})
			It(`Invoke AddDestinationEntry successfully`, func() {
				distributionListApiService, serviceErr := distributionlistapiv1.NewDistributionListApiV1(&distributionlistapiv1.DistributionListApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(distributionListApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := distributionListApiService.AddDestinationEntry(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the AddDestinationEntryRequestEventNotificationDestination model
				addDestinationEntryRequestModel := new(distributionlistapiv1.AddDestinationEntryRequestEventNotificationDestination)
				addDestinationEntryRequestModel.ID = CreateMockUUID("12345678-1234-1234-1234-123456789012")
				addDestinationEntryRequestModel.DestinationType = core.StringPtr("event_notifications")

				// Construct an instance of the AddDestinationEntryOptions model
				addDestinationEntryOptionsModel := new(distributionlistapiv1.AddDestinationEntryOptions)
				addDestinationEntryOptionsModel.AccountID = core.StringPtr("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6") // pragma: allowlist secret
				addDestinationEntryOptionsModel.AddDestinationEntryRequest = addDestinationEntryRequestModel
				addDestinationEntryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = distributionListApiService.AddDestinationEntry(addDestinationEntryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke AddDestinationEntry with error: Operation validation and request error`, func() {
				distributionListApiService, serviceErr := distributionlistapiv1.NewDistributionListApiV1(&distributionlistapiv1.DistributionListApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(distributionListApiService).ToNot(BeNil())

				// Construct an instance of the AddDestinationEntryRequestEventNotificationDestination model
				addDestinationEntryRequestModel := new(distributionlistapiv1.AddDestinationEntryRequestEventNotificationDestination)
				addDestinationEntryRequestModel.ID = CreateMockUUID("12345678-1234-1234-1234-123456789012")
				addDestinationEntryRequestModel.DestinationType = core.StringPtr("event_notifications")

				// Construct an instance of the AddDestinationEntryOptions model
				addDestinationEntryOptionsModel := new(distributionlistapiv1.AddDestinationEntryOptions)
				addDestinationEntryOptionsModel.AccountID = core.StringPtr("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6") // pragma: allowlist secret
				addDestinationEntryOptionsModel.AddDestinationEntryRequest = addDestinationEntryRequestModel
				addDestinationEntryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := distributionListApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := distributionListApiService.AddDestinationEntry(addDestinationEntryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the AddDestinationEntryOptions model with no property values
				addDestinationEntryOptionsModelNew := new(distributionlistapiv1.AddDestinationEntryOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = distributionListApiService.AddDestinationEntry(addDestinationEntryOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(201)
				}))
			})
			It(`Invoke AddDestinationEntry successfully`, func() {
				distributionListApiService, serviceErr := distributionlistapiv1.NewDistributionListApiV1(&distributionlistapiv1.DistributionListApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(distributionListApiService).ToNot(BeNil())

				// Construct an instance of the AddDestinationEntryRequestEventNotificationDestination model
				addDestinationEntryRequestModel := new(distributionlistapiv1.AddDestinationEntryRequestEventNotificationDestination)
				addDestinationEntryRequestModel.ID = CreateMockUUID("12345678-1234-1234-1234-123456789012")
				addDestinationEntryRequestModel.DestinationType = core.StringPtr("event_notifications")

				// Construct an instance of the AddDestinationEntryOptions model
				addDestinationEntryOptionsModel := new(distributionlistapiv1.AddDestinationEntryOptions)
				addDestinationEntryOptionsModel.AccountID = core.StringPtr("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6") // pragma: allowlist secret
				addDestinationEntryOptionsModel.AddDestinationEntryRequest = addDestinationEntryRequestModel
				addDestinationEntryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := distributionListApiService.AddDestinationEntry(addDestinationEntryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetDestinationEntry(getDestinationEntryOptions *GetDestinationEntryOptions) - Operation response error`, func() {
		getDestinationEntryPath := "/notification-api/v1/distribution_lists/a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6/destinations/12345678-1234-1234-1234-123456789012" // pragma: allowlist secret
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDestinationEntryPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetDestinationEntry with error: Operation response processing error`, func() {
				distributionListApiService, serviceErr := distributionlistapiv1.NewDistributionListApiV1(&distributionlistapiv1.DistributionListApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(distributionListApiService).ToNot(BeNil())

				// Construct an instance of the GetDestinationEntryOptions model
				getDestinationEntryOptionsModel := new(distributionlistapiv1.GetDestinationEntryOptions)
				getDestinationEntryOptionsModel.AccountID = core.StringPtr("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6") // pragma: allowlist secret
				getDestinationEntryOptionsModel.DestinationID = CreateMockUUID("12345678-1234-1234-1234-123456789012")
				getDestinationEntryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := distributionListApiService.GetDestinationEntry(getDestinationEntryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				distributionListApiService.EnableRetries(0, 0)
				result, response, operationErr = distributionListApiService.GetDestinationEntry(getDestinationEntryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetDestinationEntry(getDestinationEntryOptions *GetDestinationEntryOptions)`, func() {
		getDestinationEntryPath := "/notification-api/v1/distribution_lists/a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6/destinations/12345678-1234-1234-1234-123456789012" // pragma: allowlist secret
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDestinationEntryPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "12345678-1234-1234-1234-123456789012", "destination_type": "event_notifications"}`)
				}))
			})
			It(`Invoke GetDestinationEntry successfully with retries`, func() {
				distributionListApiService, serviceErr := distributionlistapiv1.NewDistributionListApiV1(&distributionlistapiv1.DistributionListApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(distributionListApiService).ToNot(BeNil())
				distributionListApiService.EnableRetries(0, 0)

				// Construct an instance of the GetDestinationEntryOptions model
				getDestinationEntryOptionsModel := new(distributionlistapiv1.GetDestinationEntryOptions)
				getDestinationEntryOptionsModel.AccountID = core.StringPtr("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6") // pragma: allowlist secret
				getDestinationEntryOptionsModel.DestinationID = CreateMockUUID("12345678-1234-1234-1234-123456789012")
				getDestinationEntryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := distributionListApiService.GetDestinationEntryWithContext(ctx, getDestinationEntryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				distributionListApiService.DisableRetries()
				result, response, operationErr := distributionListApiService.GetDestinationEntry(getDestinationEntryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = distributionListApiService.GetDestinationEntryWithContext(ctx, getDestinationEntryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDestinationEntryPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "12345678-1234-1234-1234-123456789012", "destination_type": "event_notifications"}`)
				}))
			})
			It(`Invoke GetDestinationEntry successfully`, func() {
				distributionListApiService, serviceErr := distributionlistapiv1.NewDistributionListApiV1(&distributionlistapiv1.DistributionListApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(distributionListApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := distributionListApiService.GetDestinationEntry(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetDestinationEntryOptions model
				getDestinationEntryOptionsModel := new(distributionlistapiv1.GetDestinationEntryOptions)
				getDestinationEntryOptionsModel.AccountID = core.StringPtr("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6") // pragma: allowlist secret
				getDestinationEntryOptionsModel.DestinationID = CreateMockUUID("12345678-1234-1234-1234-123456789012")
				getDestinationEntryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = distributionListApiService.GetDestinationEntry(getDestinationEntryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetDestinationEntry with error: Operation validation and request error`, func() {
				distributionListApiService, serviceErr := distributionlistapiv1.NewDistributionListApiV1(&distributionlistapiv1.DistributionListApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(distributionListApiService).ToNot(BeNil())

				// Construct an instance of the GetDestinationEntryOptions model
				getDestinationEntryOptionsModel := new(distributionlistapiv1.GetDestinationEntryOptions)
				getDestinationEntryOptionsModel.AccountID = core.StringPtr("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6") // pragma: allowlist secret
				getDestinationEntryOptionsModel.DestinationID = CreateMockUUID("12345678-1234-1234-1234-123456789012")
				getDestinationEntryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := distributionListApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := distributionListApiService.GetDestinationEntry(getDestinationEntryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetDestinationEntryOptions model with no property values
				getDestinationEntryOptionsModelNew := new(distributionlistapiv1.GetDestinationEntryOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = distributionListApiService.GetDestinationEntry(getDestinationEntryOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetDestinationEntry successfully`, func() {
				distributionListApiService, serviceErr := distributionlistapiv1.NewDistributionListApiV1(&distributionlistapiv1.DistributionListApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(distributionListApiService).ToNot(BeNil())

				// Construct an instance of the GetDestinationEntryOptions model
				getDestinationEntryOptionsModel := new(distributionlistapiv1.GetDestinationEntryOptions)
				getDestinationEntryOptionsModel.AccountID = core.StringPtr("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6") // pragma: allowlist secret
				getDestinationEntryOptionsModel.DestinationID = CreateMockUUID("12345678-1234-1234-1234-123456789012")
				getDestinationEntryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := distributionListApiService.GetDestinationEntry(getDestinationEntryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteDestinationEntry(deleteDestinationEntryOptions *DeleteDestinationEntryOptions)`, func() {
		deleteDestinationEntryPath := "/notification-api/v1/distribution_lists/a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6/destinations/12345678-1234-1234-1234-123456789012" // pragma: allowlist secret
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteDestinationEntryPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteDestinationEntry successfully`, func() {
				distributionListApiService, serviceErr := distributionlistapiv1.NewDistributionListApiV1(&distributionlistapiv1.DistributionListApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(distributionListApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := distributionListApiService.DeleteDestinationEntry(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteDestinationEntryOptions model
				deleteDestinationEntryOptionsModel := new(distributionlistapiv1.DeleteDestinationEntryOptions)
				deleteDestinationEntryOptionsModel.AccountID = core.StringPtr("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6") // pragma: allowlist secret
				deleteDestinationEntryOptionsModel.DestinationID = CreateMockUUID("12345678-1234-1234-1234-123456789012")
				deleteDestinationEntryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = distributionListApiService.DeleteDestinationEntry(deleteDestinationEntryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteDestinationEntry with error: Operation validation and request error`, func() {
				distributionListApiService, serviceErr := distributionlistapiv1.NewDistributionListApiV1(&distributionlistapiv1.DistributionListApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(distributionListApiService).ToNot(BeNil())

				// Construct an instance of the DeleteDestinationEntryOptions model
				deleteDestinationEntryOptionsModel := new(distributionlistapiv1.DeleteDestinationEntryOptions)
				deleteDestinationEntryOptionsModel.AccountID = core.StringPtr("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6") // pragma: allowlist secret
				deleteDestinationEntryOptionsModel.DestinationID = CreateMockUUID("12345678-1234-1234-1234-123456789012")
				deleteDestinationEntryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := distributionListApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := distributionListApiService.DeleteDestinationEntry(deleteDestinationEntryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteDestinationEntryOptions model with no property values
				deleteDestinationEntryOptionsModelNew := new(distributionlistapiv1.DeleteDestinationEntryOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = distributionListApiService.DeleteDestinationEntry(deleteDestinationEntryOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`TestDestinationEntry(testDestinationEntryOptions *TestDestinationEntryOptions) - Operation response error`, func() {
		testDestinationEntryPath := "/notification-api/v1/distribution_lists/a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6/destinations/12345678-1234-1234-1234-123456789012/test" // pragma: allowlist secret
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(testDestinationEntryPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke TestDestinationEntry with error: Operation response processing error`, func() {
				distributionListApiService, serviceErr := distributionlistapiv1.NewDistributionListApiV1(&distributionlistapiv1.DistributionListApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(distributionListApiService).ToNot(BeNil())

				// Construct an instance of the TestDestinationEntryRequestTestEventNotificationDestination model
				testDestinationEntryRequestModel := new(distributionlistapiv1.TestDestinationEntryRequestTestEventNotificationDestination)
				testDestinationEntryRequestModel.DestinationType = core.StringPtr("event_notifications")
				testDestinationEntryRequestModel.NotificationType = core.StringPtr("incident")

				// Construct an instance of the TestDestinationEntryOptions model
				testDestinationEntryOptionsModel := new(distributionlistapiv1.TestDestinationEntryOptions)
				testDestinationEntryOptionsModel.AccountID = core.StringPtr("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6") // pragma: allowlist secret
				testDestinationEntryOptionsModel.DestinationID = CreateMockUUID("12345678-1234-1234-1234-123456789012")
				testDestinationEntryOptionsModel.TestDestinationEntryRequest = testDestinationEntryRequestModel
				testDestinationEntryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := distributionListApiService.TestDestinationEntry(testDestinationEntryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				distributionListApiService.EnableRetries(0, 0)
				result, response, operationErr = distributionListApiService.TestDestinationEntry(testDestinationEntryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`TestDestinationEntry(testDestinationEntryOptions *TestDestinationEntryOptions)`, func() {
		testDestinationEntryPath := "/notification-api/v1/distribution_lists/a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6/destinations/12345678-1234-1234-1234-123456789012/test" // pragma: allowlist secret
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(testDestinationEntryPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"message": "success"}`)
				}))
			})
			It(`Invoke TestDestinationEntry successfully with retries`, func() {
				distributionListApiService, serviceErr := distributionlistapiv1.NewDistributionListApiV1(&distributionlistapiv1.DistributionListApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(distributionListApiService).ToNot(BeNil())
				distributionListApiService.EnableRetries(0, 0)

				// Construct an instance of the TestDestinationEntryRequestTestEventNotificationDestination model
				testDestinationEntryRequestModel := new(distributionlistapiv1.TestDestinationEntryRequestTestEventNotificationDestination)
				testDestinationEntryRequestModel.DestinationType = core.StringPtr("event_notifications")
				testDestinationEntryRequestModel.NotificationType = core.StringPtr("incident")

				// Construct an instance of the TestDestinationEntryOptions model
				testDestinationEntryOptionsModel := new(distributionlistapiv1.TestDestinationEntryOptions)
				testDestinationEntryOptionsModel.AccountID = core.StringPtr("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6") // pragma: allowlist secret
				testDestinationEntryOptionsModel.DestinationID = CreateMockUUID("12345678-1234-1234-1234-123456789012")
				testDestinationEntryOptionsModel.TestDestinationEntryRequest = testDestinationEntryRequestModel
				testDestinationEntryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := distributionListApiService.TestDestinationEntryWithContext(ctx, testDestinationEntryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				distributionListApiService.DisableRetries()
				result, response, operationErr := distributionListApiService.TestDestinationEntry(testDestinationEntryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = distributionListApiService.TestDestinationEntryWithContext(ctx, testDestinationEntryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(testDestinationEntryPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"message": "success"}`)
				}))
			})
			It(`Invoke TestDestinationEntry successfully`, func() {
				distributionListApiService, serviceErr := distributionlistapiv1.NewDistributionListApiV1(&distributionlistapiv1.DistributionListApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(distributionListApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := distributionListApiService.TestDestinationEntry(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the TestDestinationEntryRequestTestEventNotificationDestination model
				testDestinationEntryRequestModel := new(distributionlistapiv1.TestDestinationEntryRequestTestEventNotificationDestination)
				testDestinationEntryRequestModel.DestinationType = core.StringPtr("event_notifications")
				testDestinationEntryRequestModel.NotificationType = core.StringPtr("incident")

				// Construct an instance of the TestDestinationEntryOptions model
				testDestinationEntryOptionsModel := new(distributionlistapiv1.TestDestinationEntryOptions)
				testDestinationEntryOptionsModel.AccountID = core.StringPtr("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6") // pragma: allowlist secret
				testDestinationEntryOptionsModel.DestinationID = CreateMockUUID("12345678-1234-1234-1234-123456789012")
				testDestinationEntryOptionsModel.TestDestinationEntryRequest = testDestinationEntryRequestModel
				testDestinationEntryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = distributionListApiService.TestDestinationEntry(testDestinationEntryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke TestDestinationEntry with error: Operation validation and request error`, func() {
				distributionListApiService, serviceErr := distributionlistapiv1.NewDistributionListApiV1(&distributionlistapiv1.DistributionListApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(distributionListApiService).ToNot(BeNil())

				// Construct an instance of the TestDestinationEntryRequestTestEventNotificationDestination model
				testDestinationEntryRequestModel := new(distributionlistapiv1.TestDestinationEntryRequestTestEventNotificationDestination)
				testDestinationEntryRequestModel.DestinationType = core.StringPtr("event_notifications")
				testDestinationEntryRequestModel.NotificationType = core.StringPtr("incident")

				// Construct an instance of the TestDestinationEntryOptions model
				testDestinationEntryOptionsModel := new(distributionlistapiv1.TestDestinationEntryOptions)
				testDestinationEntryOptionsModel.AccountID = core.StringPtr("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6") // pragma: allowlist secret
				testDestinationEntryOptionsModel.DestinationID = CreateMockUUID("12345678-1234-1234-1234-123456789012")
				testDestinationEntryOptionsModel.TestDestinationEntryRequest = testDestinationEntryRequestModel
				testDestinationEntryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := distributionListApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := distributionListApiService.TestDestinationEntry(testDestinationEntryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the TestDestinationEntryOptions model with no property values
				testDestinationEntryOptionsModelNew := new(distributionlistapiv1.TestDestinationEntryOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = distributionListApiService.TestDestinationEntry(testDestinationEntryOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke TestDestinationEntry successfully`, func() {
				distributionListApiService, serviceErr := distributionlistapiv1.NewDistributionListApiV1(&distributionlistapiv1.DistributionListApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(distributionListApiService).ToNot(BeNil())

				// Construct an instance of the TestDestinationEntryRequestTestEventNotificationDestination model
				testDestinationEntryRequestModel := new(distributionlistapiv1.TestDestinationEntryRequestTestEventNotificationDestination)
				testDestinationEntryRequestModel.DestinationType = core.StringPtr("event_notifications")
				testDestinationEntryRequestModel.NotificationType = core.StringPtr("incident")

				// Construct an instance of the TestDestinationEntryOptions model
				testDestinationEntryOptionsModel := new(distributionlistapiv1.TestDestinationEntryOptions)
				testDestinationEntryOptionsModel.AccountID = core.StringPtr("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6") // pragma: allowlist secret
				testDestinationEntryOptionsModel.DestinationID = CreateMockUUID("12345678-1234-1234-1234-123456789012")
				testDestinationEntryOptionsModel.TestDestinationEntryRequest = testDestinationEntryRequestModel
				testDestinationEntryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := distributionListApiService.TestDestinationEntry(testDestinationEntryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Model constructor tests`, func() {
		Context(`Using a service client instance`, func() {
			distributionListApiService, _ := distributionlistapiv1.NewDistributionListApiV1(&distributionlistapiv1.DistributionListApiV1Options{
				URL:           "http://distributionlistapiv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewAddDestinationEntryOptions successfully`, func() {
				// Construct an instance of the AddDestinationEntryRequestEventNotificationDestination model
				addDestinationEntryRequestModel := new(distributionlistapiv1.AddDestinationEntryRequestEventNotificationDestination)
				Expect(addDestinationEntryRequestModel).ToNot(BeNil())
				addDestinationEntryRequestModel.ID = CreateMockUUID("12345678-1234-1234-1234-123456789012")
				addDestinationEntryRequestModel.DestinationType = core.StringPtr("event_notifications")
				Expect(addDestinationEntryRequestModel.ID).To(Equal(CreateMockUUID("12345678-1234-1234-1234-123456789012")))
				Expect(addDestinationEntryRequestModel.DestinationType).To(Equal(core.StringPtr("event_notifications")))

				// Construct an instance of the AddDestinationEntryOptions model
				accountID := "a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6" // pragma: allowlist secret
				var addDestinationEntryRequest distributionlistapiv1.AddDestinationEntryRequestIntf = nil
				addDestinationEntryOptionsModel := distributionListApiService.NewAddDestinationEntryOptions(accountID, addDestinationEntryRequest)
				addDestinationEntryOptionsModel.SetAccountID("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6") // pragma: allowlist secret
				addDestinationEntryOptionsModel.SetAddDestinationEntryRequest(addDestinationEntryRequestModel)
				addDestinationEntryOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(addDestinationEntryOptionsModel).ToNot(BeNil())
				Expect(addDestinationEntryOptionsModel.AccountID).To(Equal(core.StringPtr("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6"))) // pragma: allowlist secret
				Expect(addDestinationEntryOptionsModel.AddDestinationEntryRequest).To(Equal(addDestinationEntryRequestModel))
				Expect(addDestinationEntryOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteDestinationEntryOptions successfully`, func() {
				// Construct an instance of the DeleteDestinationEntryOptions model
				accountID := "a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6" // pragma: allowlist secret
				destinationID := CreateMockUUID("12345678-1234-1234-1234-123456789012")
				deleteDestinationEntryOptionsModel := distributionListApiService.NewDeleteDestinationEntryOptions(accountID, destinationID)
				deleteDestinationEntryOptionsModel.SetAccountID("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6") // pragma: allowlist secret
				deleteDestinationEntryOptionsModel.SetDestinationID(CreateMockUUID("12345678-1234-1234-1234-123456789012"))
				deleteDestinationEntryOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteDestinationEntryOptionsModel).ToNot(BeNil())
				Expect(deleteDestinationEntryOptionsModel.AccountID).To(Equal(core.StringPtr("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6"))) // pragma: allowlist secret
				Expect(deleteDestinationEntryOptionsModel.DestinationID).To(Equal(CreateMockUUID("12345678-1234-1234-1234-123456789012")))
				Expect(deleteDestinationEntryOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetAllDestinationEntriesOptions successfully`, func() {
				// Construct an instance of the GetAllDestinationEntriesOptions model
				accountID := "a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6" // pragma: allowlist secret
				getAllDestinationEntriesOptionsModel := distributionListApiService.NewGetAllDestinationEntriesOptions(accountID)
				getAllDestinationEntriesOptionsModel.SetAccountID("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6") // pragma: allowlist secret
				getAllDestinationEntriesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getAllDestinationEntriesOptionsModel).ToNot(BeNil())
				Expect(getAllDestinationEntriesOptionsModel.AccountID).To(Equal(core.StringPtr("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6"))) // pragma: allowlist secret
				Expect(getAllDestinationEntriesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetDestinationEntryOptions successfully`, func() {
				// Construct an instance of the GetDestinationEntryOptions model
				accountID := "a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6" // pragma: allowlist secret
				destinationID := CreateMockUUID("12345678-1234-1234-1234-123456789012")
				getDestinationEntryOptionsModel := distributionListApiService.NewGetDestinationEntryOptions(accountID, destinationID)
				getDestinationEntryOptionsModel.SetAccountID("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6") // pragma: allowlist secret
				getDestinationEntryOptionsModel.SetDestinationID(CreateMockUUID("12345678-1234-1234-1234-123456789012"))
				getDestinationEntryOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getDestinationEntryOptionsModel).ToNot(BeNil())
				Expect(getDestinationEntryOptionsModel.AccountID).To(Equal(core.StringPtr("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6"))) // pragma: allowlist secret
				Expect(getDestinationEntryOptionsModel.DestinationID).To(Equal(CreateMockUUID("12345678-1234-1234-1234-123456789012")))
				Expect(getDestinationEntryOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewTestDestinationEntryOptions successfully`, func() {
				// Construct an instance of the TestDestinationEntryRequestTestEventNotificationDestination model
				testDestinationEntryRequestModel := new(distributionlistapiv1.TestDestinationEntryRequestTestEventNotificationDestination)
				Expect(testDestinationEntryRequestModel).ToNot(BeNil())
				testDestinationEntryRequestModel.DestinationType = core.StringPtr("event_notifications")
				testDestinationEntryRequestModel.NotificationType = core.StringPtr("incident")
				Expect(testDestinationEntryRequestModel.DestinationType).To(Equal(core.StringPtr("event_notifications")))
				Expect(testDestinationEntryRequestModel.NotificationType).To(Equal(core.StringPtr("incident")))

				// Construct an instance of the TestDestinationEntryOptions model
				accountID := "a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6" // pragma: allowlist secret
				destinationID := CreateMockUUID("12345678-1234-1234-1234-123456789012")
				var testDestinationEntryRequest distributionlistapiv1.TestDestinationEntryRequestIntf = nil
				testDestinationEntryOptionsModel := distributionListApiService.NewTestDestinationEntryOptions(accountID, destinationID, testDestinationEntryRequest)
				testDestinationEntryOptionsModel.SetAccountID("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6") // pragma: allowlist secret
				testDestinationEntryOptionsModel.SetDestinationID(CreateMockUUID("12345678-1234-1234-1234-123456789012"))
				testDestinationEntryOptionsModel.SetTestDestinationEntryRequest(testDestinationEntryRequestModel)
				testDestinationEntryOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(testDestinationEntryOptionsModel).ToNot(BeNil())
				Expect(testDestinationEntryOptionsModel.AccountID).To(Equal(core.StringPtr("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6"))) // pragma: allowlist secret
				Expect(testDestinationEntryOptionsModel.DestinationID).To(Equal(CreateMockUUID("12345678-1234-1234-1234-123456789012")))
				Expect(testDestinationEntryOptionsModel.TestDestinationEntryRequest).To(Equal(testDestinationEntryRequestModel))
				Expect(testDestinationEntryOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewAddDestinationEntryRequestEventNotificationDestination successfully`, func() {
				id := CreateMockUUID("12345678-1234-1234-1234-123456789012")
				destinationType := "event_notifications"
				_model, err := distributionListApiService.NewAddDestinationEntryRequestEventNotificationDestination(id, destinationType)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewTestDestinationEntryRequestTestEventNotificationDestination successfully`, func() {
				destinationType := "event_notifications"
				notificationType := "incident"
				_model, err := distributionListApiService.NewTestDestinationEntryRequestTestEventNotificationDestination(destinationType, notificationType)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
		})
	})
	Describe(`Model unmarshaling tests`, func() {
		It(`Invoke UnmarshalAddDestinationEntryRequest successfully`, func() {
			// Construct an instance of the model.
			model := new(distributionlistapiv1.AddDestinationEntryRequest)
			model.ID = CreateMockUUID("12345678-1234-1234-1234-123456789012")
			model.DestinationType = core.StringPtr("event_notifications")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result interface{}
			err = distributionlistapiv1.UnmarshalAddDestinationEntryRequest(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
		})
		It(`Invoke UnmarshalTestDestinationEntryRequest successfully`, func() {
			// Construct an instance of the model.
			model := new(distributionlistapiv1.TestDestinationEntryRequest)
			model.DestinationType = core.StringPtr("event_notifications")
			model.NotificationType = core.StringPtr("incident")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result interface{}
			err = distributionlistapiv1.UnmarshalTestDestinationEntryRequest(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
		})
		It(`Invoke UnmarshalAddDestinationEntryRequestEventNotificationDestination successfully`, func() {
			// Construct an instance of the model.
			model := new(distributionlistapiv1.AddDestinationEntryRequestEventNotificationDestination)
			model.ID = CreateMockUUID("12345678-1234-1234-1234-123456789012")
			model.DestinationType = core.StringPtr("event_notifications")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *distributionlistapiv1.AddDestinationEntryRequestEventNotificationDestination
			err = distributionlistapiv1.UnmarshalAddDestinationEntryRequestEventNotificationDestination(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalTestDestinationEntryRequestTestEventNotificationDestination successfully`, func() {
			// Construct an instance of the model.
			model := new(distributionlistapiv1.TestDestinationEntryRequestTestEventNotificationDestination)
			model.DestinationType = core.StringPtr("event_notifications")
			model.NotificationType = core.StringPtr("incident")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *distributionlistapiv1.TestDestinationEntryRequestTestEventNotificationDestination
			err = distributionlistapiv1.UnmarshalTestDestinationEntryRequestTestEventNotificationDestination(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
	})
	Describe(`Utility function tests`, func() {
		It(`Invoke CreateMockByteArray() successfully`, func() {
			mockByteArray := CreateMockByteArray("VGhpcyBpcyBhIHRlc3Qgb2YgdGhlIGVtZXJnZW5jeSBicm9hZGNhc3Qgc3lzdGVt") // pragma: allowlist secret
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
			mockDate := CreateMockDate("2019-01-01")
			Expect(mockDate).ToNot(BeNil())
		})
		It(`Invoke CreateMockDateTime() successfully`, func() {
			mockDateTime := CreateMockDateTime("2019-01-01T12:00:00.000Z")
			Expect(mockDateTime).ToNot(BeNil())
		})
	})
})

//
// Utility functions used by the generated test code
//

func CreateMockByteArray(encodedString string) *[]byte {
	ba, err := base64.StdEncoding.DecodeString(encodedString)
	if err != nil {
		panic(err)
	}
	return &ba
}

func CreateMockUUID(mockData string) *strfmt.UUID {
	uuid := strfmt.UUID(mockData)
	return &uuid
}

func CreateMockReader(mockData string) io.ReadCloser {
	return io.NopCloser(bytes.NewReader([]byte(mockData)))
}

func CreateMockDate(mockData string) *strfmt.Date {
	d, err := core.ParseDate(mockData)
	if err != nil {
		return nil
	}
	return &d
}

func CreateMockDateTime(mockData string) *strfmt.DateTime {
	d, err := core.ParseDateTime(mockData)
	if err != nil {
		return nil
	}
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
