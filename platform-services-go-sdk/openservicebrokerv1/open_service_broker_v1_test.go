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

package openservicebrokerv1_test

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/openservicebrokerv1"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`OpenServiceBrokerV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			openServiceBrokerService, serviceErr := openservicebrokerv1.NewOpenServiceBrokerV1(&openservicebrokerv1.OpenServiceBrokerV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(openServiceBrokerService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			openServiceBrokerService, serviceErr := openservicebrokerv1.NewOpenServiceBrokerV1(&openservicebrokerv1.OpenServiceBrokerV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(openServiceBrokerService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			openServiceBrokerService, serviceErr := openservicebrokerv1.NewOpenServiceBrokerV1(&openservicebrokerv1.OpenServiceBrokerV1Options{
				URL: "https://openservicebrokerv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(openServiceBrokerService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"OPEN_SERVICE_BROKER_URL":       "https://openservicebrokerv1/api",
				"OPEN_SERVICE_BROKER_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				openServiceBrokerService, serviceErr := openservicebrokerv1.NewOpenServiceBrokerV1UsingExternalConfig(&openservicebrokerv1.OpenServiceBrokerV1Options{})
				Expect(openServiceBrokerService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := openServiceBrokerService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != openServiceBrokerService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(openServiceBrokerService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(openServiceBrokerService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				openServiceBrokerService, serviceErr := openservicebrokerv1.NewOpenServiceBrokerV1UsingExternalConfig(&openservicebrokerv1.OpenServiceBrokerV1Options{
					URL: "https://testService/api",
				})
				Expect(openServiceBrokerService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(openServiceBrokerService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := openServiceBrokerService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != openServiceBrokerService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(openServiceBrokerService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(openServiceBrokerService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				openServiceBrokerService, serviceErr := openservicebrokerv1.NewOpenServiceBrokerV1UsingExternalConfig(&openservicebrokerv1.OpenServiceBrokerV1Options{})
				err := openServiceBrokerService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(openServiceBrokerService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(openServiceBrokerService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := openServiceBrokerService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != openServiceBrokerService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(openServiceBrokerService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(openServiceBrokerService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"OPEN_SERVICE_BROKER_URL":       "https://openservicebrokerv1/api",
				"OPEN_SERVICE_BROKER_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			openServiceBrokerService, serviceErr := openservicebrokerv1.NewOpenServiceBrokerV1UsingExternalConfig(&openservicebrokerv1.OpenServiceBrokerV1Options{})

			It(`Instantiate service client with error`, func() {
				Expect(openServiceBrokerService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"OPEN_SERVICE_BROKER_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			openServiceBrokerService, serviceErr := openservicebrokerv1.NewOpenServiceBrokerV1UsingExternalConfig(&openservicebrokerv1.OpenServiceBrokerV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(openServiceBrokerService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = openservicebrokerv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`GetServiceInstanceState(getServiceInstanceStateOptions *GetServiceInstanceStateOptions) - Operation response error`, func() {
		getServiceInstanceStatePath := "/bluemix_v1/service_instances/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getServiceInstanceStatePath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetServiceInstanceState with error: Operation response processing error`, func() {
				openServiceBrokerService, serviceErr := openservicebrokerv1.NewOpenServiceBrokerV1(&openservicebrokerv1.OpenServiceBrokerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(openServiceBrokerService).ToNot(BeNil())

				// Construct an instance of the GetServiceInstanceStateOptions model
				getServiceInstanceStateOptionsModel := new(openservicebrokerv1.GetServiceInstanceStateOptions)
				getServiceInstanceStateOptionsModel.InstanceID = core.StringPtr("testString")
				getServiceInstanceStateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := openServiceBrokerService.GetServiceInstanceState(getServiceInstanceStateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				openServiceBrokerService.EnableRetries(0, 0)
				result, response, operationErr = openServiceBrokerService.GetServiceInstanceState(getServiceInstanceStateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetServiceInstanceState(getServiceInstanceStateOptions *GetServiceInstanceStateOptions)`, func() {
		getServiceInstanceStatePath := "/bluemix_v1/service_instances/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getServiceInstanceStatePath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"active": true, "enabled": false, "last_active": 10}`)
				}))
			})
			It(`Invoke GetServiceInstanceState successfully`, func() {
				openServiceBrokerService, serviceErr := openservicebrokerv1.NewOpenServiceBrokerV1(&openservicebrokerv1.OpenServiceBrokerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(openServiceBrokerService).ToNot(BeNil())
				openServiceBrokerService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := openServiceBrokerService.GetServiceInstanceState(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetServiceInstanceStateOptions model
				getServiceInstanceStateOptionsModel := new(openservicebrokerv1.GetServiceInstanceStateOptions)
				getServiceInstanceStateOptionsModel.InstanceID = core.StringPtr("testString")
				getServiceInstanceStateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = openServiceBrokerService.GetServiceInstanceState(getServiceInstanceStateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = openServiceBrokerService.GetServiceInstanceStateWithContext(ctx, getServiceInstanceStateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				openServiceBrokerService.DisableRetries()
				result, response, operationErr = openServiceBrokerService.GetServiceInstanceState(getServiceInstanceStateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = openServiceBrokerService.GetServiceInstanceStateWithContext(ctx, getServiceInstanceStateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetServiceInstanceState with error: Operation validation and request error`, func() {
				openServiceBrokerService, serviceErr := openservicebrokerv1.NewOpenServiceBrokerV1(&openservicebrokerv1.OpenServiceBrokerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(openServiceBrokerService).ToNot(BeNil())

				// Construct an instance of the GetServiceInstanceStateOptions model
				getServiceInstanceStateOptionsModel := new(openservicebrokerv1.GetServiceInstanceStateOptions)
				getServiceInstanceStateOptionsModel.InstanceID = core.StringPtr("testString")
				getServiceInstanceStateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := openServiceBrokerService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := openServiceBrokerService.GetServiceInstanceState(getServiceInstanceStateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetServiceInstanceStateOptions model with no property values
				getServiceInstanceStateOptionsModelNew := new(openservicebrokerv1.GetServiceInstanceStateOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = openServiceBrokerService.GetServiceInstanceState(getServiceInstanceStateOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ReplaceServiceInstanceState(replaceServiceInstanceStateOptions *ReplaceServiceInstanceStateOptions) - Operation response error`, func() {
		replaceServiceInstanceStatePath := "/bluemix_v1/service_instances/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceServiceInstanceStatePath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ReplaceServiceInstanceState with error: Operation response processing error`, func() {
				openServiceBrokerService, serviceErr := openservicebrokerv1.NewOpenServiceBrokerV1(&openservicebrokerv1.OpenServiceBrokerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(openServiceBrokerService).ToNot(BeNil())

				// Construct an instance of the ReplaceServiceInstanceStateOptions model
				replaceServiceInstanceStateOptionsModel := new(openservicebrokerv1.ReplaceServiceInstanceStateOptions)
				replaceServiceInstanceStateOptionsModel.InstanceID = core.StringPtr("testString")
				replaceServiceInstanceStateOptionsModel.Enabled = core.BoolPtr(false)
				replaceServiceInstanceStateOptionsModel.InitiatorID = core.StringPtr("null")
				replaceServiceInstanceStateOptionsModel.ReasonCode = core.StringPtr("null")
				replaceServiceInstanceStateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := openServiceBrokerService.ReplaceServiceInstanceState(replaceServiceInstanceStateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				openServiceBrokerService.EnableRetries(0, 0)
				result, response, operationErr = openServiceBrokerService.ReplaceServiceInstanceState(replaceServiceInstanceStateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ReplaceServiceInstanceState(replaceServiceInstanceStateOptions *ReplaceServiceInstanceStateOptions)`, func() {
		replaceServiceInstanceStatePath := "/bluemix_v1/service_instances/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceServiceInstanceStatePath))
					Expect(req.Method).To(Equal("PUT"))

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
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"active": true, "enabled": false, "last_active": 10}`)
				}))
			})
			It(`Invoke ReplaceServiceInstanceState successfully`, func() {
				openServiceBrokerService, serviceErr := openservicebrokerv1.NewOpenServiceBrokerV1(&openservicebrokerv1.OpenServiceBrokerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(openServiceBrokerService).ToNot(BeNil())
				openServiceBrokerService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := openServiceBrokerService.ReplaceServiceInstanceState(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ReplaceServiceInstanceStateOptions model
				replaceServiceInstanceStateOptionsModel := new(openservicebrokerv1.ReplaceServiceInstanceStateOptions)
				replaceServiceInstanceStateOptionsModel.InstanceID = core.StringPtr("testString")
				replaceServiceInstanceStateOptionsModel.Enabled = core.BoolPtr(false)
				replaceServiceInstanceStateOptionsModel.InitiatorID = core.StringPtr("null")
				replaceServiceInstanceStateOptionsModel.ReasonCode = core.StringPtr("null")
				replaceServiceInstanceStateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = openServiceBrokerService.ReplaceServiceInstanceState(replaceServiceInstanceStateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = openServiceBrokerService.ReplaceServiceInstanceStateWithContext(ctx, replaceServiceInstanceStateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				openServiceBrokerService.DisableRetries()
				result, response, operationErr = openServiceBrokerService.ReplaceServiceInstanceState(replaceServiceInstanceStateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = openServiceBrokerService.ReplaceServiceInstanceStateWithContext(ctx, replaceServiceInstanceStateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke ReplaceServiceInstanceState with error: Operation validation and request error`, func() {
				openServiceBrokerService, serviceErr := openservicebrokerv1.NewOpenServiceBrokerV1(&openservicebrokerv1.OpenServiceBrokerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(openServiceBrokerService).ToNot(BeNil())

				// Construct an instance of the ReplaceServiceInstanceStateOptions model
				replaceServiceInstanceStateOptionsModel := new(openservicebrokerv1.ReplaceServiceInstanceStateOptions)
				replaceServiceInstanceStateOptionsModel.InstanceID = core.StringPtr("testString")
				replaceServiceInstanceStateOptionsModel.Enabled = core.BoolPtr(false)
				replaceServiceInstanceStateOptionsModel.InitiatorID = core.StringPtr("null")
				replaceServiceInstanceStateOptionsModel.ReasonCode = core.StringPtr("null")
				replaceServiceInstanceStateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := openServiceBrokerService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := openServiceBrokerService.ReplaceServiceInstanceState(replaceServiceInstanceStateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ReplaceServiceInstanceStateOptions model with no property values
				replaceServiceInstanceStateOptionsModelNew := new(openservicebrokerv1.ReplaceServiceInstanceStateOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = openServiceBrokerService.ReplaceServiceInstanceState(replaceServiceInstanceStateOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			openServiceBrokerService, serviceErr := openservicebrokerv1.NewOpenServiceBrokerV1(&openservicebrokerv1.OpenServiceBrokerV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(openServiceBrokerService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			openServiceBrokerService, serviceErr := openservicebrokerv1.NewOpenServiceBrokerV1(&openservicebrokerv1.OpenServiceBrokerV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(openServiceBrokerService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			openServiceBrokerService, serviceErr := openservicebrokerv1.NewOpenServiceBrokerV1(&openservicebrokerv1.OpenServiceBrokerV1Options{
				URL: "https://openservicebrokerv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(openServiceBrokerService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"OPEN_SERVICE_BROKER_URL":       "https://openservicebrokerv1/api",
				"OPEN_SERVICE_BROKER_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				openServiceBrokerService, serviceErr := openservicebrokerv1.NewOpenServiceBrokerV1UsingExternalConfig(&openservicebrokerv1.OpenServiceBrokerV1Options{})
				Expect(openServiceBrokerService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := openServiceBrokerService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != openServiceBrokerService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(openServiceBrokerService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(openServiceBrokerService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				openServiceBrokerService, serviceErr := openservicebrokerv1.NewOpenServiceBrokerV1UsingExternalConfig(&openservicebrokerv1.OpenServiceBrokerV1Options{
					URL: "https://testService/api",
				})
				Expect(openServiceBrokerService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(openServiceBrokerService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := openServiceBrokerService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != openServiceBrokerService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(openServiceBrokerService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(openServiceBrokerService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				openServiceBrokerService, serviceErr := openservicebrokerv1.NewOpenServiceBrokerV1UsingExternalConfig(&openservicebrokerv1.OpenServiceBrokerV1Options{})
				err := openServiceBrokerService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(openServiceBrokerService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(openServiceBrokerService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := openServiceBrokerService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != openServiceBrokerService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(openServiceBrokerService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(openServiceBrokerService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"OPEN_SERVICE_BROKER_URL":       "https://openservicebrokerv1/api",
				"OPEN_SERVICE_BROKER_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			openServiceBrokerService, serviceErr := openservicebrokerv1.NewOpenServiceBrokerV1UsingExternalConfig(&openservicebrokerv1.OpenServiceBrokerV1Options{})

			It(`Instantiate service client with error`, func() {
				Expect(openServiceBrokerService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"OPEN_SERVICE_BROKER_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			openServiceBrokerService, serviceErr := openservicebrokerv1.NewOpenServiceBrokerV1UsingExternalConfig(&openservicebrokerv1.OpenServiceBrokerV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(openServiceBrokerService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = openservicebrokerv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`ReplaceServiceInstance(replaceServiceInstanceOptions *ReplaceServiceInstanceOptions) - Operation response error`, func() {
		replaceServiceInstancePath := "/v2/service_instances/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceServiceInstancePath))
					Expect(req.Method).To(Equal("PUT"))

					// TODO: Add check for accepts_incomplete query parameter

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ReplaceServiceInstance with error: Operation response processing error`, func() {
				openServiceBrokerService, serviceErr := openservicebrokerv1.NewOpenServiceBrokerV1(&openservicebrokerv1.OpenServiceBrokerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(openServiceBrokerService).ToNot(BeNil())

				// Construct an instance of the Context model
				contextModel := new(openservicebrokerv1.Context)
				contextModel.AccountID = core.StringPtr("null")
				contextModel.CRN = core.StringPtr("null")
				contextModel.Platform = core.StringPtr("null")

				// Construct an instance of the ReplaceServiceInstanceOptions model
				replaceServiceInstanceOptionsModel := new(openservicebrokerv1.ReplaceServiceInstanceOptions)
				replaceServiceInstanceOptionsModel.InstanceID = core.StringPtr("testString")
				replaceServiceInstanceOptionsModel.Context = contextModel
				replaceServiceInstanceOptionsModel.OrganizationGUID = core.StringPtr("null")
				replaceServiceInstanceOptionsModel.Parameters = make(map[string]string)
				replaceServiceInstanceOptionsModel.PlanID = core.StringPtr("null")
				replaceServiceInstanceOptionsModel.ServiceID = core.StringPtr("null")
				replaceServiceInstanceOptionsModel.SpaceGUID = core.StringPtr("null")
				replaceServiceInstanceOptionsModel.AcceptsIncomplete = core.BoolPtr(true)
				replaceServiceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := openServiceBrokerService.ReplaceServiceInstance(replaceServiceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				openServiceBrokerService.EnableRetries(0, 0)
				result, response, operationErr = openServiceBrokerService.ReplaceServiceInstance(replaceServiceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ReplaceServiceInstance(replaceServiceInstanceOptions *ReplaceServiceInstanceOptions)`, func() {
		replaceServiceInstancePath := "/v2/service_instances/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceServiceInstancePath))
					Expect(req.Method).To(Equal("PUT"))

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

					// TODO: Add check for accepts_incomplete query parameter

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"dashboard_url": "DashboardURL", "operation": "Operation"}`)
				}))
			})
			It(`Invoke ReplaceServiceInstance successfully`, func() {
				openServiceBrokerService, serviceErr := openservicebrokerv1.NewOpenServiceBrokerV1(&openservicebrokerv1.OpenServiceBrokerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(openServiceBrokerService).ToNot(BeNil())
				openServiceBrokerService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := openServiceBrokerService.ReplaceServiceInstance(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the Context model
				contextModel := new(openservicebrokerv1.Context)
				contextModel.AccountID = core.StringPtr("null")
				contextModel.CRN = core.StringPtr("null")
				contextModel.Platform = core.StringPtr("null")

				// Construct an instance of the ReplaceServiceInstanceOptions model
				replaceServiceInstanceOptionsModel := new(openservicebrokerv1.ReplaceServiceInstanceOptions)
				replaceServiceInstanceOptionsModel.InstanceID = core.StringPtr("testString")
				replaceServiceInstanceOptionsModel.Context = contextModel
				replaceServiceInstanceOptionsModel.OrganizationGUID = core.StringPtr("null")
				replaceServiceInstanceOptionsModel.Parameters = make(map[string]string)
				replaceServiceInstanceOptionsModel.PlanID = core.StringPtr("null")
				replaceServiceInstanceOptionsModel.ServiceID = core.StringPtr("null")
				replaceServiceInstanceOptionsModel.SpaceGUID = core.StringPtr("null")
				replaceServiceInstanceOptionsModel.AcceptsIncomplete = core.BoolPtr(true)
				replaceServiceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = openServiceBrokerService.ReplaceServiceInstance(replaceServiceInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = openServiceBrokerService.ReplaceServiceInstanceWithContext(ctx, replaceServiceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				openServiceBrokerService.DisableRetries()
				result, response, operationErr = openServiceBrokerService.ReplaceServiceInstance(replaceServiceInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = openServiceBrokerService.ReplaceServiceInstanceWithContext(ctx, replaceServiceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke ReplaceServiceInstance with error: Operation validation and request error`, func() {
				openServiceBrokerService, serviceErr := openservicebrokerv1.NewOpenServiceBrokerV1(&openservicebrokerv1.OpenServiceBrokerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(openServiceBrokerService).ToNot(BeNil())

				// Construct an instance of the Context model
				contextModel := new(openservicebrokerv1.Context)
				contextModel.AccountID = core.StringPtr("null")
				contextModel.CRN = core.StringPtr("null")
				contextModel.Platform = core.StringPtr("null")

				// Construct an instance of the ReplaceServiceInstanceOptions model
				replaceServiceInstanceOptionsModel := new(openservicebrokerv1.ReplaceServiceInstanceOptions)
				replaceServiceInstanceOptionsModel.InstanceID = core.StringPtr("testString")
				replaceServiceInstanceOptionsModel.Context = contextModel
				replaceServiceInstanceOptionsModel.OrganizationGUID = core.StringPtr("null")
				replaceServiceInstanceOptionsModel.Parameters = make(map[string]string)
				replaceServiceInstanceOptionsModel.PlanID = core.StringPtr("null")
				replaceServiceInstanceOptionsModel.ServiceID = core.StringPtr("null")
				replaceServiceInstanceOptionsModel.SpaceGUID = core.StringPtr("null")
				replaceServiceInstanceOptionsModel.AcceptsIncomplete = core.BoolPtr(true)
				replaceServiceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := openServiceBrokerService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := openServiceBrokerService.ReplaceServiceInstance(replaceServiceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ReplaceServiceInstanceOptions model with no property values
				replaceServiceInstanceOptionsModelNew := new(openservicebrokerv1.ReplaceServiceInstanceOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = openServiceBrokerService.ReplaceServiceInstance(replaceServiceInstanceOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateServiceInstance(updateServiceInstanceOptions *UpdateServiceInstanceOptions) - Operation response error`, func() {
		updateServiceInstancePath := "/v2/service_instances/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateServiceInstancePath))
					Expect(req.Method).To(Equal("PATCH"))

					// TODO: Add check for accepts_incomplete query parameter

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateServiceInstance with error: Operation response processing error`, func() {
				openServiceBrokerService, serviceErr := openservicebrokerv1.NewOpenServiceBrokerV1(&openservicebrokerv1.OpenServiceBrokerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(openServiceBrokerService).ToNot(BeNil())

				// Construct an instance of the Context model
				contextModel := new(openservicebrokerv1.Context)
				contextModel.AccountID = core.StringPtr("null")
				contextModel.CRN = core.StringPtr("null")
				contextModel.Platform = core.StringPtr("null")

				// Construct an instance of the UpdateServiceInstanceOptions model
				updateServiceInstanceOptionsModel := new(openservicebrokerv1.UpdateServiceInstanceOptions)
				updateServiceInstanceOptionsModel.InstanceID = core.StringPtr("testString")
				updateServiceInstanceOptionsModel.Context = contextModel
				updateServiceInstanceOptionsModel.Parameters = make(map[string]string)
				updateServiceInstanceOptionsModel.PlanID = core.StringPtr("null")
				updateServiceInstanceOptionsModel.PreviousValues = make(map[string]string)
				updateServiceInstanceOptionsModel.ServiceID = core.StringPtr("null")
				updateServiceInstanceOptionsModel.AcceptsIncomplete = core.BoolPtr(true)
				updateServiceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := openServiceBrokerService.UpdateServiceInstance(updateServiceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				openServiceBrokerService.EnableRetries(0, 0)
				result, response, operationErr = openServiceBrokerService.UpdateServiceInstance(updateServiceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateServiceInstance(updateServiceInstanceOptions *UpdateServiceInstanceOptions)`, func() {
		updateServiceInstancePath := "/v2/service_instances/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateServiceInstancePath))
					Expect(req.Method).To(Equal("PATCH"))

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

					// TODO: Add check for accepts_incomplete query parameter

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"operation": "Operation"}`)
				}))
			})
			It(`Invoke UpdateServiceInstance successfully`, func() {
				openServiceBrokerService, serviceErr := openservicebrokerv1.NewOpenServiceBrokerV1(&openservicebrokerv1.OpenServiceBrokerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(openServiceBrokerService).ToNot(BeNil())
				openServiceBrokerService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := openServiceBrokerService.UpdateServiceInstance(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the Context model
				contextModel := new(openservicebrokerv1.Context)
				contextModel.AccountID = core.StringPtr("null")
				contextModel.CRN = core.StringPtr("null")
				contextModel.Platform = core.StringPtr("null")

				// Construct an instance of the UpdateServiceInstanceOptions model
				updateServiceInstanceOptionsModel := new(openservicebrokerv1.UpdateServiceInstanceOptions)
				updateServiceInstanceOptionsModel.InstanceID = core.StringPtr("testString")
				updateServiceInstanceOptionsModel.Context = contextModel
				updateServiceInstanceOptionsModel.Parameters = make(map[string]string)
				updateServiceInstanceOptionsModel.PlanID = core.StringPtr("null")
				updateServiceInstanceOptionsModel.PreviousValues = make(map[string]string)
				updateServiceInstanceOptionsModel.ServiceID = core.StringPtr("null")
				updateServiceInstanceOptionsModel.AcceptsIncomplete = core.BoolPtr(true)
				updateServiceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = openServiceBrokerService.UpdateServiceInstance(updateServiceInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = openServiceBrokerService.UpdateServiceInstanceWithContext(ctx, updateServiceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				openServiceBrokerService.DisableRetries()
				result, response, operationErr = openServiceBrokerService.UpdateServiceInstance(updateServiceInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = openServiceBrokerService.UpdateServiceInstanceWithContext(ctx, updateServiceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke UpdateServiceInstance with error: Operation validation and request error`, func() {
				openServiceBrokerService, serviceErr := openservicebrokerv1.NewOpenServiceBrokerV1(&openservicebrokerv1.OpenServiceBrokerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(openServiceBrokerService).ToNot(BeNil())

				// Construct an instance of the Context model
				contextModel := new(openservicebrokerv1.Context)
				contextModel.AccountID = core.StringPtr("null")
				contextModel.CRN = core.StringPtr("null")
				contextModel.Platform = core.StringPtr("null")

				// Construct an instance of the UpdateServiceInstanceOptions model
				updateServiceInstanceOptionsModel := new(openservicebrokerv1.UpdateServiceInstanceOptions)
				updateServiceInstanceOptionsModel.InstanceID = core.StringPtr("testString")
				updateServiceInstanceOptionsModel.Context = contextModel
				updateServiceInstanceOptionsModel.Parameters = make(map[string]string)
				updateServiceInstanceOptionsModel.PlanID = core.StringPtr("null")
				updateServiceInstanceOptionsModel.PreviousValues = make(map[string]string)
				updateServiceInstanceOptionsModel.ServiceID = core.StringPtr("null")
				updateServiceInstanceOptionsModel.AcceptsIncomplete = core.BoolPtr(true)
				updateServiceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := openServiceBrokerService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := openServiceBrokerService.UpdateServiceInstance(updateServiceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateServiceInstanceOptions model with no property values
				updateServiceInstanceOptionsModelNew := new(openservicebrokerv1.UpdateServiceInstanceOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = openServiceBrokerService.UpdateServiceInstance(updateServiceInstanceOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteServiceInstance(deleteServiceInstanceOptions *DeleteServiceInstanceOptions) - Operation response error`, func() {
		deleteServiceInstancePath := "/v2/service_instances/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteServiceInstancePath))
					Expect(req.Method).To(Equal("DELETE"))
					Expect(req.URL.Query()["service_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["plan_id"]).To(Equal([]string{"testString"}))

					// TODO: Add check for accepts_incomplete query parameter

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteServiceInstance with error: Operation response processing error`, func() {
				openServiceBrokerService, serviceErr := openservicebrokerv1.NewOpenServiceBrokerV1(&openservicebrokerv1.OpenServiceBrokerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(openServiceBrokerService).ToNot(BeNil())

				// Construct an instance of the DeleteServiceInstanceOptions model
				deleteServiceInstanceOptionsModel := new(openservicebrokerv1.DeleteServiceInstanceOptions)
				deleteServiceInstanceOptionsModel.ServiceID = core.StringPtr("testString")
				deleteServiceInstanceOptionsModel.PlanID = core.StringPtr("testString")
				deleteServiceInstanceOptionsModel.InstanceID = core.StringPtr("testString")
				deleteServiceInstanceOptionsModel.AcceptsIncomplete = core.BoolPtr(true)
				deleteServiceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := openServiceBrokerService.DeleteServiceInstance(deleteServiceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				openServiceBrokerService.EnableRetries(0, 0)
				result, response, operationErr = openServiceBrokerService.DeleteServiceInstance(deleteServiceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteServiceInstance(deleteServiceInstanceOptions *DeleteServiceInstanceOptions)`, func() {
		deleteServiceInstancePath := "/v2/service_instances/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteServiceInstancePath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.URL.Query()["service_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["plan_id"]).To(Equal([]string{"testString"}))

					// TODO: Add check for accepts_incomplete query parameter

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"operation": "Operation"}`)
				}))
			})
			It(`Invoke DeleteServiceInstance successfully`, func() {
				openServiceBrokerService, serviceErr := openservicebrokerv1.NewOpenServiceBrokerV1(&openservicebrokerv1.OpenServiceBrokerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(openServiceBrokerService).ToNot(BeNil())
				openServiceBrokerService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := openServiceBrokerService.DeleteServiceInstance(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteServiceInstanceOptions model
				deleteServiceInstanceOptionsModel := new(openservicebrokerv1.DeleteServiceInstanceOptions)
				deleteServiceInstanceOptionsModel.ServiceID = core.StringPtr("testString")
				deleteServiceInstanceOptionsModel.PlanID = core.StringPtr("testString")
				deleteServiceInstanceOptionsModel.InstanceID = core.StringPtr("testString")
				deleteServiceInstanceOptionsModel.AcceptsIncomplete = core.BoolPtr(true)
				deleteServiceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = openServiceBrokerService.DeleteServiceInstance(deleteServiceInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = openServiceBrokerService.DeleteServiceInstanceWithContext(ctx, deleteServiceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				openServiceBrokerService.DisableRetries()
				result, response, operationErr = openServiceBrokerService.DeleteServiceInstance(deleteServiceInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = openServiceBrokerService.DeleteServiceInstanceWithContext(ctx, deleteServiceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke DeleteServiceInstance with error: Operation validation and request error`, func() {
				openServiceBrokerService, serviceErr := openservicebrokerv1.NewOpenServiceBrokerV1(&openservicebrokerv1.OpenServiceBrokerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(openServiceBrokerService).ToNot(BeNil())

				// Construct an instance of the DeleteServiceInstanceOptions model
				deleteServiceInstanceOptionsModel := new(openservicebrokerv1.DeleteServiceInstanceOptions)
				deleteServiceInstanceOptionsModel.ServiceID = core.StringPtr("testString")
				deleteServiceInstanceOptionsModel.PlanID = core.StringPtr("testString")
				deleteServiceInstanceOptionsModel.InstanceID = core.StringPtr("testString")
				deleteServiceInstanceOptionsModel.AcceptsIncomplete = core.BoolPtr(true)
				deleteServiceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := openServiceBrokerService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := openServiceBrokerService.DeleteServiceInstance(deleteServiceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteServiceInstanceOptions model with no property values
				deleteServiceInstanceOptionsModelNew := new(openservicebrokerv1.DeleteServiceInstanceOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = openServiceBrokerService.DeleteServiceInstance(deleteServiceInstanceOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			openServiceBrokerService, serviceErr := openservicebrokerv1.NewOpenServiceBrokerV1(&openservicebrokerv1.OpenServiceBrokerV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(openServiceBrokerService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			openServiceBrokerService, serviceErr := openservicebrokerv1.NewOpenServiceBrokerV1(&openservicebrokerv1.OpenServiceBrokerV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(openServiceBrokerService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			openServiceBrokerService, serviceErr := openservicebrokerv1.NewOpenServiceBrokerV1(&openservicebrokerv1.OpenServiceBrokerV1Options{
				URL: "https://openservicebrokerv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(openServiceBrokerService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"OPEN_SERVICE_BROKER_URL":       "https://openservicebrokerv1/api",
				"OPEN_SERVICE_BROKER_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				openServiceBrokerService, serviceErr := openservicebrokerv1.NewOpenServiceBrokerV1UsingExternalConfig(&openservicebrokerv1.OpenServiceBrokerV1Options{})
				Expect(openServiceBrokerService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := openServiceBrokerService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != openServiceBrokerService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(openServiceBrokerService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(openServiceBrokerService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				openServiceBrokerService, serviceErr := openservicebrokerv1.NewOpenServiceBrokerV1UsingExternalConfig(&openservicebrokerv1.OpenServiceBrokerV1Options{
					URL: "https://testService/api",
				})
				Expect(openServiceBrokerService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(openServiceBrokerService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := openServiceBrokerService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != openServiceBrokerService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(openServiceBrokerService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(openServiceBrokerService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				openServiceBrokerService, serviceErr := openservicebrokerv1.NewOpenServiceBrokerV1UsingExternalConfig(&openservicebrokerv1.OpenServiceBrokerV1Options{})
				err := openServiceBrokerService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(openServiceBrokerService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(openServiceBrokerService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := openServiceBrokerService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != openServiceBrokerService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(openServiceBrokerService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(openServiceBrokerService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"OPEN_SERVICE_BROKER_URL":       "https://openservicebrokerv1/api",
				"OPEN_SERVICE_BROKER_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			openServiceBrokerService, serviceErr := openservicebrokerv1.NewOpenServiceBrokerV1UsingExternalConfig(&openservicebrokerv1.OpenServiceBrokerV1Options{})

			It(`Instantiate service client with error`, func() {
				Expect(openServiceBrokerService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"OPEN_SERVICE_BROKER_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			openServiceBrokerService, serviceErr := openservicebrokerv1.NewOpenServiceBrokerV1UsingExternalConfig(&openservicebrokerv1.OpenServiceBrokerV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(openServiceBrokerService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = openservicebrokerv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`ListCatalog(listCatalogOptions *ListCatalogOptions) - Operation response error`, func() {
		listCatalogPath := "/v2/catalog"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listCatalogPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListCatalog with error: Operation response processing error`, func() {
				openServiceBrokerService, serviceErr := openservicebrokerv1.NewOpenServiceBrokerV1(&openservicebrokerv1.OpenServiceBrokerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(openServiceBrokerService).ToNot(BeNil())

				// Construct an instance of the ListCatalogOptions model
				listCatalogOptionsModel := new(openservicebrokerv1.ListCatalogOptions)
				listCatalogOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := openServiceBrokerService.ListCatalog(listCatalogOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				openServiceBrokerService.EnableRetries(0, 0)
				result, response, operationErr = openServiceBrokerService.ListCatalog(listCatalogOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListCatalog(listCatalogOptions *ListCatalogOptions)`, func() {
		listCatalogPath := "/v2/catalog"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listCatalogPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"services": [{"bindable": true, "description": "Description", "id": "ID", "name": "Name", "plan_updateable": true, "plans": [{"description": "Description", "free": true, "id": "ID", "name": "Name"}]}]}`)
				}))
			})
			It(`Invoke ListCatalog successfully`, func() {
				openServiceBrokerService, serviceErr := openservicebrokerv1.NewOpenServiceBrokerV1(&openservicebrokerv1.OpenServiceBrokerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(openServiceBrokerService).ToNot(BeNil())
				openServiceBrokerService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := openServiceBrokerService.ListCatalog(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListCatalogOptions model
				listCatalogOptionsModel := new(openservicebrokerv1.ListCatalogOptions)
				listCatalogOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = openServiceBrokerService.ListCatalog(listCatalogOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = openServiceBrokerService.ListCatalogWithContext(ctx, listCatalogOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				openServiceBrokerService.DisableRetries()
				result, response, operationErr = openServiceBrokerService.ListCatalog(listCatalogOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = openServiceBrokerService.ListCatalogWithContext(ctx, listCatalogOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke ListCatalog with error: Operation request error`, func() {
				openServiceBrokerService, serviceErr := openservicebrokerv1.NewOpenServiceBrokerV1(&openservicebrokerv1.OpenServiceBrokerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(openServiceBrokerService).ToNot(BeNil())

				// Construct an instance of the ListCatalogOptions model
				listCatalogOptionsModel := new(openservicebrokerv1.ListCatalogOptions)
				listCatalogOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := openServiceBrokerService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := openServiceBrokerService.ListCatalog(listCatalogOptionsModel)
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
		It(`Instantiate service client`, func() {
			openServiceBrokerService, serviceErr := openservicebrokerv1.NewOpenServiceBrokerV1(&openservicebrokerv1.OpenServiceBrokerV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(openServiceBrokerService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			openServiceBrokerService, serviceErr := openservicebrokerv1.NewOpenServiceBrokerV1(&openservicebrokerv1.OpenServiceBrokerV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(openServiceBrokerService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			openServiceBrokerService, serviceErr := openservicebrokerv1.NewOpenServiceBrokerV1(&openservicebrokerv1.OpenServiceBrokerV1Options{
				URL: "https://openservicebrokerv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(openServiceBrokerService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"OPEN_SERVICE_BROKER_URL":       "https://openservicebrokerv1/api",
				"OPEN_SERVICE_BROKER_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				openServiceBrokerService, serviceErr := openservicebrokerv1.NewOpenServiceBrokerV1UsingExternalConfig(&openservicebrokerv1.OpenServiceBrokerV1Options{})
				Expect(openServiceBrokerService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := openServiceBrokerService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != openServiceBrokerService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(openServiceBrokerService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(openServiceBrokerService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				openServiceBrokerService, serviceErr := openservicebrokerv1.NewOpenServiceBrokerV1UsingExternalConfig(&openservicebrokerv1.OpenServiceBrokerV1Options{
					URL: "https://testService/api",
				})
				Expect(openServiceBrokerService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(openServiceBrokerService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := openServiceBrokerService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != openServiceBrokerService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(openServiceBrokerService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(openServiceBrokerService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				openServiceBrokerService, serviceErr := openservicebrokerv1.NewOpenServiceBrokerV1UsingExternalConfig(&openservicebrokerv1.OpenServiceBrokerV1Options{})
				err := openServiceBrokerService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(openServiceBrokerService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(openServiceBrokerService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := openServiceBrokerService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != openServiceBrokerService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(openServiceBrokerService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(openServiceBrokerService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"OPEN_SERVICE_BROKER_URL":       "https://openservicebrokerv1/api",
				"OPEN_SERVICE_BROKER_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			openServiceBrokerService, serviceErr := openservicebrokerv1.NewOpenServiceBrokerV1UsingExternalConfig(&openservicebrokerv1.OpenServiceBrokerV1Options{})

			It(`Instantiate service client with error`, func() {
				Expect(openServiceBrokerService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"OPEN_SERVICE_BROKER_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			openServiceBrokerService, serviceErr := openservicebrokerv1.NewOpenServiceBrokerV1UsingExternalConfig(&openservicebrokerv1.OpenServiceBrokerV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(openServiceBrokerService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = openservicebrokerv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`GetLastOperation(getLastOperationOptions *GetLastOperationOptions) - Operation response error`, func() {
		getLastOperationPath := "/v2/service_instances/testString/last_operation"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getLastOperationPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["operation"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["plan_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["service_id"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetLastOperation with error: Operation response processing error`, func() {
				openServiceBrokerService, serviceErr := openservicebrokerv1.NewOpenServiceBrokerV1(&openservicebrokerv1.OpenServiceBrokerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(openServiceBrokerService).ToNot(BeNil())

				// Construct an instance of the GetLastOperationOptions model
				getLastOperationOptionsModel := new(openservicebrokerv1.GetLastOperationOptions)
				getLastOperationOptionsModel.InstanceID = core.StringPtr("testString")
				getLastOperationOptionsModel.Operation = core.StringPtr("testString")
				getLastOperationOptionsModel.PlanID = core.StringPtr("testString")
				getLastOperationOptionsModel.ServiceID = core.StringPtr("testString")
				getLastOperationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := openServiceBrokerService.GetLastOperation(getLastOperationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				openServiceBrokerService.EnableRetries(0, 0)
				result, response, operationErr = openServiceBrokerService.GetLastOperation(getLastOperationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetLastOperation(getLastOperationOptions *GetLastOperationOptions)`, func() {
		getLastOperationPath := "/v2/service_instances/testString/last_operation"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getLastOperationPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["operation"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["plan_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["service_id"]).To(Equal([]string{"testString"}))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"description": "Description", "state": "State"}`)
				}))
			})
			It(`Invoke GetLastOperation successfully`, func() {
				openServiceBrokerService, serviceErr := openservicebrokerv1.NewOpenServiceBrokerV1(&openservicebrokerv1.OpenServiceBrokerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(openServiceBrokerService).ToNot(BeNil())
				openServiceBrokerService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := openServiceBrokerService.GetLastOperation(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetLastOperationOptions model
				getLastOperationOptionsModel := new(openservicebrokerv1.GetLastOperationOptions)
				getLastOperationOptionsModel.InstanceID = core.StringPtr("testString")
				getLastOperationOptionsModel.Operation = core.StringPtr("testString")
				getLastOperationOptionsModel.PlanID = core.StringPtr("testString")
				getLastOperationOptionsModel.ServiceID = core.StringPtr("testString")
				getLastOperationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = openServiceBrokerService.GetLastOperation(getLastOperationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = openServiceBrokerService.GetLastOperationWithContext(ctx, getLastOperationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				openServiceBrokerService.DisableRetries()
				result, response, operationErr = openServiceBrokerService.GetLastOperation(getLastOperationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = openServiceBrokerService.GetLastOperationWithContext(ctx, getLastOperationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetLastOperation with error: Operation validation and request error`, func() {
				openServiceBrokerService, serviceErr := openservicebrokerv1.NewOpenServiceBrokerV1(&openservicebrokerv1.OpenServiceBrokerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(openServiceBrokerService).ToNot(BeNil())

				// Construct an instance of the GetLastOperationOptions model
				getLastOperationOptionsModel := new(openservicebrokerv1.GetLastOperationOptions)
				getLastOperationOptionsModel.InstanceID = core.StringPtr("testString")
				getLastOperationOptionsModel.Operation = core.StringPtr("testString")
				getLastOperationOptionsModel.PlanID = core.StringPtr("testString")
				getLastOperationOptionsModel.ServiceID = core.StringPtr("testString")
				getLastOperationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := openServiceBrokerService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := openServiceBrokerService.GetLastOperation(getLastOperationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetLastOperationOptions model with no property values
				getLastOperationOptionsModelNew := new(openservicebrokerv1.GetLastOperationOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = openServiceBrokerService.GetLastOperation(getLastOperationOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			openServiceBrokerService, serviceErr := openservicebrokerv1.NewOpenServiceBrokerV1(&openservicebrokerv1.OpenServiceBrokerV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(openServiceBrokerService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			openServiceBrokerService, serviceErr := openservicebrokerv1.NewOpenServiceBrokerV1(&openservicebrokerv1.OpenServiceBrokerV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(openServiceBrokerService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			openServiceBrokerService, serviceErr := openservicebrokerv1.NewOpenServiceBrokerV1(&openservicebrokerv1.OpenServiceBrokerV1Options{
				URL: "https://openservicebrokerv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(openServiceBrokerService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"OPEN_SERVICE_BROKER_URL":       "https://openservicebrokerv1/api",
				"OPEN_SERVICE_BROKER_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				openServiceBrokerService, serviceErr := openservicebrokerv1.NewOpenServiceBrokerV1UsingExternalConfig(&openservicebrokerv1.OpenServiceBrokerV1Options{})
				Expect(openServiceBrokerService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := openServiceBrokerService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != openServiceBrokerService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(openServiceBrokerService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(openServiceBrokerService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				openServiceBrokerService, serviceErr := openservicebrokerv1.NewOpenServiceBrokerV1UsingExternalConfig(&openservicebrokerv1.OpenServiceBrokerV1Options{
					URL: "https://testService/api",
				})
				Expect(openServiceBrokerService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(openServiceBrokerService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := openServiceBrokerService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != openServiceBrokerService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(openServiceBrokerService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(openServiceBrokerService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				openServiceBrokerService, serviceErr := openservicebrokerv1.NewOpenServiceBrokerV1UsingExternalConfig(&openservicebrokerv1.OpenServiceBrokerV1Options{})
				err := openServiceBrokerService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(openServiceBrokerService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(openServiceBrokerService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := openServiceBrokerService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != openServiceBrokerService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(openServiceBrokerService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(openServiceBrokerService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"OPEN_SERVICE_BROKER_URL":       "https://openservicebrokerv1/api",
				"OPEN_SERVICE_BROKER_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			openServiceBrokerService, serviceErr := openservicebrokerv1.NewOpenServiceBrokerV1UsingExternalConfig(&openservicebrokerv1.OpenServiceBrokerV1Options{})

			It(`Instantiate service client with error`, func() {
				Expect(openServiceBrokerService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"OPEN_SERVICE_BROKER_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			openServiceBrokerService, serviceErr := openservicebrokerv1.NewOpenServiceBrokerV1UsingExternalConfig(&openservicebrokerv1.OpenServiceBrokerV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(openServiceBrokerService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = openservicebrokerv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`ReplaceServiceBinding(replaceServiceBindingOptions *ReplaceServiceBindingOptions) - Operation response error`, func() {
		replaceServiceBindingPath := "/v2/service_instances/testString/service_bindings/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceServiceBindingPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ReplaceServiceBinding with error: Operation response processing error`, func() {
				openServiceBrokerService, serviceErr := openservicebrokerv1.NewOpenServiceBrokerV1(&openservicebrokerv1.OpenServiceBrokerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(openServiceBrokerService).ToNot(BeNil())

				// Construct an instance of the BindResource model
				bindResourceModel := new(openservicebrokerv1.BindResource)
				bindResourceModel.AccountID = core.StringPtr("null")
				bindResourceModel.ServiceidCRN = core.StringPtr("null")
				bindResourceModel.TargetCRN = core.StringPtr("null")
				bindResourceModel.AppGUID = core.StringPtr("null")
				bindResourceModel.Route = core.StringPtr("null")

				// Construct an instance of the ReplaceServiceBindingOptions model
				replaceServiceBindingOptionsModel := new(openservicebrokerv1.ReplaceServiceBindingOptions)
				replaceServiceBindingOptionsModel.BindingID = core.StringPtr("testString")
				replaceServiceBindingOptionsModel.InstanceID = core.StringPtr("testString")
				replaceServiceBindingOptionsModel.BindResource = bindResourceModel
				replaceServiceBindingOptionsModel.Parameters = make(map[string]string)
				replaceServiceBindingOptionsModel.PlanID = core.StringPtr("null")
				replaceServiceBindingOptionsModel.ServiceID = core.StringPtr("null")
				replaceServiceBindingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := openServiceBrokerService.ReplaceServiceBinding(replaceServiceBindingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				openServiceBrokerService.EnableRetries(0, 0)
				result, response, operationErr = openServiceBrokerService.ReplaceServiceBinding(replaceServiceBindingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ReplaceServiceBinding(replaceServiceBindingOptions *ReplaceServiceBindingOptions)`, func() {
		replaceServiceBindingPath := "/v2/service_instances/testString/service_bindings/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceServiceBindingPath))
					Expect(req.Method).To(Equal("PUT"))

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
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"credentials": {"anyKey": "anyValue"}, "syslog_drain_url": "SyslogDrainURL", "route_service_url": "RouteServiceURL", "volume_mounts": [{"driver": "Driver", "container_dir": "ContainerDir", "mode": "Mode", "device_type": "DeviceType", "device": "Device"}]}`)
				}))
			})
			It(`Invoke ReplaceServiceBinding successfully`, func() {
				openServiceBrokerService, serviceErr := openservicebrokerv1.NewOpenServiceBrokerV1(&openservicebrokerv1.OpenServiceBrokerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(openServiceBrokerService).ToNot(BeNil())
				openServiceBrokerService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := openServiceBrokerService.ReplaceServiceBinding(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the BindResource model
				bindResourceModel := new(openservicebrokerv1.BindResource)
				bindResourceModel.AccountID = core.StringPtr("null")
				bindResourceModel.ServiceidCRN = core.StringPtr("null")
				bindResourceModel.TargetCRN = core.StringPtr("null")
				bindResourceModel.AppGUID = core.StringPtr("null")
				bindResourceModel.Route = core.StringPtr("null")

				// Construct an instance of the ReplaceServiceBindingOptions model
				replaceServiceBindingOptionsModel := new(openservicebrokerv1.ReplaceServiceBindingOptions)
				replaceServiceBindingOptionsModel.BindingID = core.StringPtr("testString")
				replaceServiceBindingOptionsModel.InstanceID = core.StringPtr("testString")
				replaceServiceBindingOptionsModel.BindResource = bindResourceModel
				replaceServiceBindingOptionsModel.Parameters = make(map[string]string)
				replaceServiceBindingOptionsModel.PlanID = core.StringPtr("null")
				replaceServiceBindingOptionsModel.ServiceID = core.StringPtr("null")
				replaceServiceBindingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = openServiceBrokerService.ReplaceServiceBinding(replaceServiceBindingOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = openServiceBrokerService.ReplaceServiceBindingWithContext(ctx, replaceServiceBindingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				openServiceBrokerService.DisableRetries()
				result, response, operationErr = openServiceBrokerService.ReplaceServiceBinding(replaceServiceBindingOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = openServiceBrokerService.ReplaceServiceBindingWithContext(ctx, replaceServiceBindingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke ReplaceServiceBinding with error: Operation validation and request error`, func() {
				openServiceBrokerService, serviceErr := openservicebrokerv1.NewOpenServiceBrokerV1(&openservicebrokerv1.OpenServiceBrokerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(openServiceBrokerService).ToNot(BeNil())

				// Construct an instance of the BindResource model
				bindResourceModel := new(openservicebrokerv1.BindResource)
				bindResourceModel.AccountID = core.StringPtr("null")
				bindResourceModel.ServiceidCRN = core.StringPtr("null")
				bindResourceModel.TargetCRN = core.StringPtr("null")
				bindResourceModel.AppGUID = core.StringPtr("null")
				bindResourceModel.Route = core.StringPtr("null")

				// Construct an instance of the ReplaceServiceBindingOptions model
				replaceServiceBindingOptionsModel := new(openservicebrokerv1.ReplaceServiceBindingOptions)
				replaceServiceBindingOptionsModel.BindingID = core.StringPtr("testString")
				replaceServiceBindingOptionsModel.InstanceID = core.StringPtr("testString")
				replaceServiceBindingOptionsModel.BindResource = bindResourceModel
				replaceServiceBindingOptionsModel.Parameters = make(map[string]string)
				replaceServiceBindingOptionsModel.PlanID = core.StringPtr("null")
				replaceServiceBindingOptionsModel.ServiceID = core.StringPtr("null")
				replaceServiceBindingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := openServiceBrokerService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := openServiceBrokerService.ReplaceServiceBinding(replaceServiceBindingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ReplaceServiceBindingOptions model with no property values
				replaceServiceBindingOptionsModelNew := new(openservicebrokerv1.ReplaceServiceBindingOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = openServiceBrokerService.ReplaceServiceBinding(replaceServiceBindingOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteServiceBinding(deleteServiceBindingOptions *DeleteServiceBindingOptions)`, func() {
		deleteServiceBindingPath := "/v2/service_instances/testString/service_bindings/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteServiceBindingPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.URL.Query()["plan_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["service_id"]).To(Equal([]string{"testString"}))

					res.WriteHeader(200)
				}))
			})
			It(`Invoke DeleteServiceBinding successfully`, func() {
				openServiceBrokerService, serviceErr := openservicebrokerv1.NewOpenServiceBrokerV1(&openservicebrokerv1.OpenServiceBrokerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(openServiceBrokerService).ToNot(BeNil())
				openServiceBrokerService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				response, operationErr := openServiceBrokerService.DeleteServiceBinding(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteServiceBindingOptions model
				deleteServiceBindingOptionsModel := new(openservicebrokerv1.DeleteServiceBindingOptions)
				deleteServiceBindingOptionsModel.BindingID = core.StringPtr("testString")
				deleteServiceBindingOptionsModel.InstanceID = core.StringPtr("testString")
				deleteServiceBindingOptionsModel.PlanID = core.StringPtr("testString")
				deleteServiceBindingOptionsModel.ServiceID = core.StringPtr("testString")
				deleteServiceBindingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = openServiceBrokerService.DeleteServiceBinding(deleteServiceBindingOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Disable retries and test again
				openServiceBrokerService.DisableRetries()
				response, operationErr = openServiceBrokerService.DeleteServiceBinding(deleteServiceBindingOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteServiceBinding with error: Operation validation and request error`, func() {
				openServiceBrokerService, serviceErr := openservicebrokerv1.NewOpenServiceBrokerV1(&openservicebrokerv1.OpenServiceBrokerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(openServiceBrokerService).ToNot(BeNil())

				// Construct an instance of the DeleteServiceBindingOptions model
				deleteServiceBindingOptionsModel := new(openservicebrokerv1.DeleteServiceBindingOptions)
				deleteServiceBindingOptionsModel.BindingID = core.StringPtr("testString")
				deleteServiceBindingOptionsModel.InstanceID = core.StringPtr("testString")
				deleteServiceBindingOptionsModel.PlanID = core.StringPtr("testString")
				deleteServiceBindingOptionsModel.ServiceID = core.StringPtr("testString")
				deleteServiceBindingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := openServiceBrokerService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := openServiceBrokerService.DeleteServiceBinding(deleteServiceBindingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteServiceBindingOptions model with no property values
				deleteServiceBindingOptionsModelNew := new(openservicebrokerv1.DeleteServiceBindingOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = openServiceBrokerService.DeleteServiceBinding(deleteServiceBindingOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Model constructor tests`, func() {
		Context(`Using a service client instance`, func() {
			openServiceBrokerService, _ := openservicebrokerv1.NewOpenServiceBrokerV1(&openservicebrokerv1.OpenServiceBrokerV1Options{
				URL:           "http://openservicebrokerv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewDeleteServiceBindingOptions successfully`, func() {
				// Construct an instance of the DeleteServiceBindingOptions model
				bindingID := "testString"
				instanceID := "testString"
				planID := "testString"
				serviceID := "testString"
				deleteServiceBindingOptionsModel := openServiceBrokerService.NewDeleteServiceBindingOptions(bindingID, instanceID, planID, serviceID)
				deleteServiceBindingOptionsModel.SetBindingID("testString")
				deleteServiceBindingOptionsModel.SetInstanceID("testString")
				deleteServiceBindingOptionsModel.SetPlanID("testString")
				deleteServiceBindingOptionsModel.SetServiceID("testString")
				deleteServiceBindingOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteServiceBindingOptionsModel).ToNot(BeNil())
				Expect(deleteServiceBindingOptionsModel.BindingID).To(Equal(core.StringPtr("testString")))
				Expect(deleteServiceBindingOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(deleteServiceBindingOptionsModel.PlanID).To(Equal(core.StringPtr("testString")))
				Expect(deleteServiceBindingOptionsModel.ServiceID).To(Equal(core.StringPtr("testString")))
				Expect(deleteServiceBindingOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteServiceInstanceOptions successfully`, func() {
				// Construct an instance of the DeleteServiceInstanceOptions model
				serviceID := "testString"
				planID := "testString"
				instanceID := "testString"
				deleteServiceInstanceOptionsModel := openServiceBrokerService.NewDeleteServiceInstanceOptions(serviceID, planID, instanceID)
				deleteServiceInstanceOptionsModel.SetServiceID("testString")
				deleteServiceInstanceOptionsModel.SetPlanID("testString")
				deleteServiceInstanceOptionsModel.SetInstanceID("testString")
				deleteServiceInstanceOptionsModel.SetAcceptsIncomplete(true)
				deleteServiceInstanceOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteServiceInstanceOptionsModel).ToNot(BeNil())
				Expect(deleteServiceInstanceOptionsModel.ServiceID).To(Equal(core.StringPtr("testString")))
				Expect(deleteServiceInstanceOptionsModel.PlanID).To(Equal(core.StringPtr("testString")))
				Expect(deleteServiceInstanceOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(deleteServiceInstanceOptionsModel.AcceptsIncomplete).To(Equal(core.BoolPtr(true)))
				Expect(deleteServiceInstanceOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetLastOperationOptions successfully`, func() {
				// Construct an instance of the GetLastOperationOptions model
				instanceID := "testString"
				getLastOperationOptionsModel := openServiceBrokerService.NewGetLastOperationOptions(instanceID)
				getLastOperationOptionsModel.SetInstanceID("testString")
				getLastOperationOptionsModel.SetOperation("testString")
				getLastOperationOptionsModel.SetPlanID("testString")
				getLastOperationOptionsModel.SetServiceID("testString")
				getLastOperationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getLastOperationOptionsModel).ToNot(BeNil())
				Expect(getLastOperationOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(getLastOperationOptionsModel.Operation).To(Equal(core.StringPtr("testString")))
				Expect(getLastOperationOptionsModel.PlanID).To(Equal(core.StringPtr("testString")))
				Expect(getLastOperationOptionsModel.ServiceID).To(Equal(core.StringPtr("testString")))
				Expect(getLastOperationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetServiceInstanceStateOptions successfully`, func() {
				// Construct an instance of the GetServiceInstanceStateOptions model
				instanceID := "testString"
				getServiceInstanceStateOptionsModel := openServiceBrokerService.NewGetServiceInstanceStateOptions(instanceID)
				getServiceInstanceStateOptionsModel.SetInstanceID("testString")
				getServiceInstanceStateOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getServiceInstanceStateOptionsModel).ToNot(BeNil())
				Expect(getServiceInstanceStateOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(getServiceInstanceStateOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListCatalogOptions successfully`, func() {
				// Construct an instance of the ListCatalogOptions model
				listCatalogOptionsModel := openServiceBrokerService.NewListCatalogOptions()
				listCatalogOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listCatalogOptionsModel).ToNot(BeNil())
				Expect(listCatalogOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewReplaceServiceBindingOptions successfully`, func() {
				// Construct an instance of the BindResource model
				bindResourceModel := new(openservicebrokerv1.BindResource)
				Expect(bindResourceModel).ToNot(BeNil())
				bindResourceModel.AccountID = core.StringPtr("null")
				bindResourceModel.ServiceidCRN = core.StringPtr("null")
				bindResourceModel.TargetCRN = core.StringPtr("null")
				bindResourceModel.AppGUID = core.StringPtr("null")
				bindResourceModel.Route = core.StringPtr("null")
				Expect(bindResourceModel.AccountID).To(Equal(core.StringPtr("null")))
				Expect(bindResourceModel.ServiceidCRN).To(Equal(core.StringPtr("null")))
				Expect(bindResourceModel.TargetCRN).To(Equal(core.StringPtr("null")))
				Expect(bindResourceModel.AppGUID).To(Equal(core.StringPtr("null")))
				Expect(bindResourceModel.Route).To(Equal(core.StringPtr("null")))

				// Construct an instance of the ReplaceServiceBindingOptions model
				bindingID := "testString"
				instanceID := "testString"
				replaceServiceBindingOptionsModel := openServiceBrokerService.NewReplaceServiceBindingOptions(bindingID, instanceID)
				replaceServiceBindingOptionsModel.SetBindingID("testString")
				replaceServiceBindingOptionsModel.SetInstanceID("testString")
				replaceServiceBindingOptionsModel.SetBindResource(bindResourceModel)
				replaceServiceBindingOptionsModel.SetParameters(make(map[string]string))
				replaceServiceBindingOptionsModel.SetPlanID("null")
				replaceServiceBindingOptionsModel.SetServiceID("null")
				replaceServiceBindingOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(replaceServiceBindingOptionsModel).ToNot(BeNil())
				Expect(replaceServiceBindingOptionsModel.BindingID).To(Equal(core.StringPtr("testString")))
				Expect(replaceServiceBindingOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(replaceServiceBindingOptionsModel.BindResource).To(Equal(bindResourceModel))
				Expect(replaceServiceBindingOptionsModel.Parameters).To(Equal(make(map[string]string)))
				Expect(replaceServiceBindingOptionsModel.PlanID).To(Equal(core.StringPtr("null")))
				Expect(replaceServiceBindingOptionsModel.ServiceID).To(Equal(core.StringPtr("null")))
				Expect(replaceServiceBindingOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewReplaceServiceInstanceOptions successfully`, func() {
				// Construct an instance of the Context model
				contextModel := new(openservicebrokerv1.Context)
				Expect(contextModel).ToNot(BeNil())
				contextModel.AccountID = core.StringPtr("null")
				contextModel.CRN = core.StringPtr("null")
				contextModel.Platform = core.StringPtr("null")
				Expect(contextModel.AccountID).To(Equal(core.StringPtr("null")))
				Expect(contextModel.CRN).To(Equal(core.StringPtr("null")))
				Expect(contextModel.Platform).To(Equal(core.StringPtr("null")))

				// Construct an instance of the ReplaceServiceInstanceOptions model
				instanceID := "testString"
				replaceServiceInstanceOptionsModel := openServiceBrokerService.NewReplaceServiceInstanceOptions(instanceID)
				replaceServiceInstanceOptionsModel.SetInstanceID("testString")
				replaceServiceInstanceOptionsModel.SetContext(contextModel)
				replaceServiceInstanceOptionsModel.SetOrganizationGUID("null")
				replaceServiceInstanceOptionsModel.SetParameters(make(map[string]string))
				replaceServiceInstanceOptionsModel.SetPlanID("null")
				replaceServiceInstanceOptionsModel.SetServiceID("null")
				replaceServiceInstanceOptionsModel.SetSpaceGUID("null")
				replaceServiceInstanceOptionsModel.SetAcceptsIncomplete(true)
				replaceServiceInstanceOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(replaceServiceInstanceOptionsModel).ToNot(BeNil())
				Expect(replaceServiceInstanceOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(replaceServiceInstanceOptionsModel.Context).To(Equal(contextModel))
				Expect(replaceServiceInstanceOptionsModel.OrganizationGUID).To(Equal(core.StringPtr("null")))
				Expect(replaceServiceInstanceOptionsModel.Parameters).To(Equal(make(map[string]string)))
				Expect(replaceServiceInstanceOptionsModel.PlanID).To(Equal(core.StringPtr("null")))
				Expect(replaceServiceInstanceOptionsModel.ServiceID).To(Equal(core.StringPtr("null")))
				Expect(replaceServiceInstanceOptionsModel.SpaceGUID).To(Equal(core.StringPtr("null")))
				Expect(replaceServiceInstanceOptionsModel.AcceptsIncomplete).To(Equal(core.BoolPtr(true)))
				Expect(replaceServiceInstanceOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewReplaceServiceInstanceStateOptions successfully`, func() {
				// Construct an instance of the ReplaceServiceInstanceStateOptions model
				instanceID := "testString"
				replaceServiceInstanceStateOptionsModel := openServiceBrokerService.NewReplaceServiceInstanceStateOptions(instanceID)
				replaceServiceInstanceStateOptionsModel.SetInstanceID("testString")
				replaceServiceInstanceStateOptionsModel.SetEnabled(false)
				replaceServiceInstanceStateOptionsModel.SetInitiatorID("null")
				replaceServiceInstanceStateOptionsModel.SetReasonCode("null")
				replaceServiceInstanceStateOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(replaceServiceInstanceStateOptionsModel).ToNot(BeNil())
				Expect(replaceServiceInstanceStateOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(replaceServiceInstanceStateOptionsModel.Enabled).To(Equal(core.BoolPtr(false)))
				Expect(replaceServiceInstanceStateOptionsModel.InitiatorID).To(Equal(core.StringPtr("null")))
				Expect(replaceServiceInstanceStateOptionsModel.ReasonCode).To(Equal(core.StringPtr("null")))
				Expect(replaceServiceInstanceStateOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateServiceInstanceOptions successfully`, func() {
				// Construct an instance of the Context model
				contextModel := new(openservicebrokerv1.Context)
				Expect(contextModel).ToNot(BeNil())
				contextModel.AccountID = core.StringPtr("null")
				contextModel.CRN = core.StringPtr("null")
				contextModel.Platform = core.StringPtr("null")
				Expect(contextModel.AccountID).To(Equal(core.StringPtr("null")))
				Expect(contextModel.CRN).To(Equal(core.StringPtr("null")))
				Expect(contextModel.Platform).To(Equal(core.StringPtr("null")))

				// Construct an instance of the UpdateServiceInstanceOptions model
				instanceID := "testString"
				updateServiceInstanceOptionsModel := openServiceBrokerService.NewUpdateServiceInstanceOptions(instanceID)
				updateServiceInstanceOptionsModel.SetInstanceID("testString")
				updateServiceInstanceOptionsModel.SetContext(contextModel)
				updateServiceInstanceOptionsModel.SetParameters(make(map[string]string))
				updateServiceInstanceOptionsModel.SetPlanID("null")
				updateServiceInstanceOptionsModel.SetPreviousValues(make(map[string]string))
				updateServiceInstanceOptionsModel.SetServiceID("null")
				updateServiceInstanceOptionsModel.SetAcceptsIncomplete(true)
				updateServiceInstanceOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateServiceInstanceOptionsModel).ToNot(BeNil())
				Expect(updateServiceInstanceOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(updateServiceInstanceOptionsModel.Context).To(Equal(contextModel))
				Expect(updateServiceInstanceOptionsModel.Parameters).To(Equal(make(map[string]string)))
				Expect(updateServiceInstanceOptionsModel.PlanID).To(Equal(core.StringPtr("null")))
				Expect(updateServiceInstanceOptionsModel.PreviousValues).To(Equal(make(map[string]string)))
				Expect(updateServiceInstanceOptionsModel.ServiceID).To(Equal(core.StringPtr("null")))
				Expect(updateServiceInstanceOptionsModel.AcceptsIncomplete).To(Equal(core.BoolPtr(true)))
				Expect(updateServiceInstanceOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
	return io.NopCloser(bytes.NewReader([]byte(mockData)))
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
