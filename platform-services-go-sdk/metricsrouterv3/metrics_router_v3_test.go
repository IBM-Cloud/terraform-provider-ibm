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

package metricsrouterv3_test

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

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/metricsrouterv3"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`MetricsRouterV3`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			metricsRouterService, serviceErr := metricsrouterv3.NewMetricsRouterV3(&metricsrouterv3.MetricsRouterV3Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(metricsRouterService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			metricsRouterService, serviceErr := metricsrouterv3.NewMetricsRouterV3(&metricsrouterv3.MetricsRouterV3Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(metricsRouterService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			metricsRouterService, serviceErr := metricsrouterv3.NewMetricsRouterV3(&metricsrouterv3.MetricsRouterV3Options{
				URL: "https://metricsrouterv3/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(metricsRouterService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"METRICS_ROUTER_URL":       "https://metricsrouterv3/api",
				"METRICS_ROUTER_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				metricsRouterService, serviceErr := metricsrouterv3.NewMetricsRouterV3UsingExternalConfig(&metricsrouterv3.MetricsRouterV3Options{})
				Expect(metricsRouterService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := metricsRouterService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != metricsRouterService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(metricsRouterService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(metricsRouterService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				metricsRouterService, serviceErr := metricsrouterv3.NewMetricsRouterV3UsingExternalConfig(&metricsrouterv3.MetricsRouterV3Options{
					URL: "https://testService/api",
				})
				Expect(metricsRouterService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(metricsRouterService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := metricsRouterService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != metricsRouterService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(metricsRouterService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(metricsRouterService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				metricsRouterService, serviceErr := metricsrouterv3.NewMetricsRouterV3UsingExternalConfig(&metricsrouterv3.MetricsRouterV3Options{})
				err := metricsRouterService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(metricsRouterService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(metricsRouterService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := metricsRouterService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != metricsRouterService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(metricsRouterService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(metricsRouterService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"METRICS_ROUTER_URL":       "https://metricsrouterv3/api",
				"METRICS_ROUTER_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			metricsRouterService, serviceErr := metricsrouterv3.NewMetricsRouterV3UsingExternalConfig(&metricsrouterv3.MetricsRouterV3Options{})

			It(`Instantiate service client with error`, func() {
				Expect(metricsRouterService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"METRICS_ROUTER_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			metricsRouterService, serviceErr := metricsrouterv3.NewMetricsRouterV3UsingExternalConfig(&metricsrouterv3.MetricsRouterV3Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(metricsRouterService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = metricsrouterv3.GetServiceURLForRegion("au-syd")
			Expect(url).To(Equal("https://au-syd.metrics-router.cloud.ibm.com/api/v3"))
			Expect(err).To(BeNil())

			url, err = metricsrouterv3.GetServiceURLForRegion("private.au-syd")
			Expect(url).To(Equal("https://private.au-syd.metrics-router.cloud.ibm.com/api/v3"))
			Expect(err).To(BeNil())

			url, err = metricsrouterv3.GetServiceURLForRegion("br-sao")
			Expect(url).To(Equal("https://br-sao.metrics-router.cloud.ibm.com/api/v3"))
			Expect(err).To(BeNil())

			url, err = metricsrouterv3.GetServiceURLForRegion("private.br-sao")
			Expect(url).To(Equal("https://private.br-sao.metrics-router.cloud.ibm.com/api/v3"))
			Expect(err).To(BeNil())

			url, err = metricsrouterv3.GetServiceURLForRegion("ca-tor")
			Expect(url).To(Equal("https://ca-tor.metrics-router.cloud.ibm.com/api/v3"))
			Expect(err).To(BeNil())

			url, err = metricsrouterv3.GetServiceURLForRegion("private.ca-tor")
			Expect(url).To(Equal("https://private.ca-tor.metrics-router.cloud.ibm.com/api/v3"))
			Expect(err).To(BeNil())

			url, err = metricsrouterv3.GetServiceURLForRegion("eu-de")
			Expect(url).To(Equal("https://eu-de.metrics-router.cloud.ibm.com/api/v3"))
			Expect(err).To(BeNil())

			url, err = metricsrouterv3.GetServiceURLForRegion("private.eu-de")
			Expect(url).To(Equal("https://private.eu-de.metrics-router.cloud.ibm.com/api/v3"))
			Expect(err).To(BeNil())

			url, err = metricsrouterv3.GetServiceURLForRegion("eu-gb")
			Expect(url).To(Equal("https://eu-gb.metrics-router.cloud.ibm.com/api/v3"))
			Expect(err).To(BeNil())

			url, err = metricsrouterv3.GetServiceURLForRegion("private.eu-gb")
			Expect(url).To(Equal("https://private.eu-gb.metrics-router.cloud.ibm.com/api/v3"))
			Expect(err).To(BeNil())

			url, err = metricsrouterv3.GetServiceURLForRegion("eu-es")
			Expect(url).To(Equal("https://eu-es.metrics-router.cloud.ibm.com/api/v3"))
			Expect(err).To(BeNil())

			url, err = metricsrouterv3.GetServiceURLForRegion("private.eu-es")
			Expect(url).To(Equal("https://private.eu-es.metrics-router.cloud.ibm.com/api/v3"))
			Expect(err).To(BeNil())

			url, err = metricsrouterv3.GetServiceURLForRegion("eu-fr2")
			Expect(url).To(Equal("https://eu-fr2.metrics-router.cloud.ibm.com/api/v3"))
			Expect(err).To(BeNil())

			url, err = metricsrouterv3.GetServiceURLForRegion("private.eu-fr2")
			Expect(url).To(Equal("https://private.eu-fr2.metrics-router.cloud.ibm.com/api/v3"))
			Expect(err).To(BeNil())

			url, err = metricsrouterv3.GetServiceURLForRegion("jp-osa")
			Expect(url).To(Equal("https://jp-osa.metrics-router.cloud.ibm.com/api/v3"))
			Expect(err).To(BeNil())

			url, err = metricsrouterv3.GetServiceURLForRegion("private.jp-osa")
			Expect(url).To(Equal("https://private.jp-osa.metrics-router.cloud.ibm.com/api/v3"))
			Expect(err).To(BeNil())

			url, err = metricsrouterv3.GetServiceURLForRegion("jp-tok")
			Expect(url).To(Equal("https://jp-tok.metrics-router.cloud.ibm.com/api/v3"))
			Expect(err).To(BeNil())

			url, err = metricsrouterv3.GetServiceURLForRegion("private.jp-tok")
			Expect(url).To(Equal("https://private.jp-tok.metrics-router.cloud.ibm.com/api/v3"))
			Expect(err).To(BeNil())

			url, err = metricsrouterv3.GetServiceURLForRegion("us-east")
			Expect(url).To(Equal("https://us-east.metrics-router.cloud.ibm.com/api/v3"))
			Expect(err).To(BeNil())

			url, err = metricsrouterv3.GetServiceURLForRegion("private.us-east")
			Expect(url).To(Equal("https://private.us-east.metrics-router.cloud.ibm.com/api/v3"))
			Expect(err).To(BeNil())

			url, err = metricsrouterv3.GetServiceURLForRegion("us-south")
			Expect(url).To(Equal("https://us-south.metrics-router.cloud.ibm.com/api/v3"))
			Expect(err).To(BeNil())

			url, err = metricsrouterv3.GetServiceURLForRegion("private.us-south")
			Expect(url).To(Equal("https://private.us-south.metrics-router.cloud.ibm.com/api/v3"))
			Expect(err).To(BeNil())

			url, err = metricsrouterv3.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`CreateTarget(createTargetOptions *CreateTargetOptions) - Operation response error`, func() {
		createTargetPath := "/targets"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createTargetPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateTarget with error: Operation response processing error`, func() {
				metricsRouterService, serviceErr := metricsrouterv3.NewMetricsRouterV3(&metricsrouterv3.MetricsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(metricsRouterService).ToNot(BeNil())

				// Construct an instance of the CreateTargetOptions model
				createTargetOptionsModel := new(metricsrouterv3.CreateTargetOptions)
				createTargetOptionsModel.Name = core.StringPtr("my-mr-target")
				createTargetOptionsModel.DestinationCRN = core.StringPtr("crn:v1:bluemix:public:sysdig-monitor:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::")
				createTargetOptionsModel.Region = core.StringPtr("us-south")
				createTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := metricsRouterService.CreateTarget(createTargetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				metricsRouterService.EnableRetries(0, 0)
				result, response, operationErr = metricsRouterService.CreateTarget(createTargetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateTarget(createTargetOptions *CreateTargetOptions)`, func() {
		createTargetPath := "/targets"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createTargetPath))
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
					fmt.Fprintf(res, "%s", `{"id": "f7dcfae6-e7c5-08ca-451b-fdfa696c9bb6", "name": "a-mr-target-us-south", "crn": "crn:v1:bluemix:public:metrics-router:us-south:a/0be5ad401ae913d8ff665d92680664ed:b6eec08b-5201-08ca-451b-cd71523e3626:target:f7dcfae6-e7c5-08ca-451b-fdfa696c9bb6", "destination_crn": "crn:v1:bluemix:public:sysdig-monitor:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::", "target_type": "sysdig_monitor", "region": "us-south", "created_at": "2021-05-18T20:15:12.353Z", "updated_at": "2021-05-18T20:15:12.353Z"}`)
				}))
			})
			It(`Invoke CreateTarget successfully with retries`, func() {
				metricsRouterService, serviceErr := metricsrouterv3.NewMetricsRouterV3(&metricsrouterv3.MetricsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(metricsRouterService).ToNot(BeNil())
				metricsRouterService.EnableRetries(0, 0)

				// Construct an instance of the CreateTargetOptions model
				createTargetOptionsModel := new(metricsrouterv3.CreateTargetOptions)
				createTargetOptionsModel.Name = core.StringPtr("my-mr-target")
				createTargetOptionsModel.DestinationCRN = core.StringPtr("crn:v1:bluemix:public:sysdig-monitor:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::")
				createTargetOptionsModel.Region = core.StringPtr("us-south")
				createTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := metricsRouterService.CreateTargetWithContext(ctx, createTargetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				metricsRouterService.DisableRetries()
				result, response, operationErr := metricsRouterService.CreateTarget(createTargetOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = metricsRouterService.CreateTargetWithContext(ctx, createTargetOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createTargetPath))
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
					fmt.Fprintf(res, "%s", `{"id": "f7dcfae6-e7c5-08ca-451b-fdfa696c9bb6", "name": "a-mr-target-us-south", "crn": "crn:v1:bluemix:public:metrics-router:us-south:a/0be5ad401ae913d8ff665d92680664ed:b6eec08b-5201-08ca-451b-cd71523e3626:target:f7dcfae6-e7c5-08ca-451b-fdfa696c9bb6", "destination_crn": "crn:v1:bluemix:public:sysdig-monitor:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::", "target_type": "sysdig_monitor", "region": "us-south", "created_at": "2021-05-18T20:15:12.353Z", "updated_at": "2021-05-18T20:15:12.353Z"}`)
				}))
			})
			It(`Invoke CreateTarget successfully`, func() {
				metricsRouterService, serviceErr := metricsrouterv3.NewMetricsRouterV3(&metricsrouterv3.MetricsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(metricsRouterService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := metricsRouterService.CreateTarget(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateTargetOptions model
				createTargetOptionsModel := new(metricsrouterv3.CreateTargetOptions)
				createTargetOptionsModel.Name = core.StringPtr("my-mr-target")
				createTargetOptionsModel.DestinationCRN = core.StringPtr("crn:v1:bluemix:public:sysdig-monitor:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::")
				createTargetOptionsModel.Region = core.StringPtr("us-south")
				createTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = metricsRouterService.CreateTarget(createTargetOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateTarget with error: Operation validation and request error`, func() {
				metricsRouterService, serviceErr := metricsrouterv3.NewMetricsRouterV3(&metricsrouterv3.MetricsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(metricsRouterService).ToNot(BeNil())

				// Construct an instance of the CreateTargetOptions model
				createTargetOptionsModel := new(metricsrouterv3.CreateTargetOptions)
				createTargetOptionsModel.Name = core.StringPtr("my-mr-target")
				createTargetOptionsModel.DestinationCRN = core.StringPtr("crn:v1:bluemix:public:sysdig-monitor:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::")
				createTargetOptionsModel.Region = core.StringPtr("us-south")
				createTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := metricsRouterService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := metricsRouterService.CreateTarget(createTargetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateTargetOptions model with no property values
				createTargetOptionsModelNew := new(metricsrouterv3.CreateTargetOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = metricsRouterService.CreateTarget(createTargetOptionsModelNew)
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
			It(`Invoke CreateTarget successfully`, func() {
				metricsRouterService, serviceErr := metricsrouterv3.NewMetricsRouterV3(&metricsrouterv3.MetricsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(metricsRouterService).ToNot(BeNil())

				// Construct an instance of the CreateTargetOptions model
				createTargetOptionsModel := new(metricsrouterv3.CreateTargetOptions)
				createTargetOptionsModel.Name = core.StringPtr("my-mr-target")
				createTargetOptionsModel.DestinationCRN = core.StringPtr("crn:v1:bluemix:public:sysdig-monitor:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::")
				createTargetOptionsModel.Region = core.StringPtr("us-south")
				createTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := metricsRouterService.CreateTarget(createTargetOptionsModel)
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
	Describe(`ListTargets(listTargetsOptions *ListTargetsOptions) - Operation response error`, func() {
		listTargetsPath := "/targets"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listTargetsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListTargets with error: Operation response processing error`, func() {
				metricsRouterService, serviceErr := metricsrouterv3.NewMetricsRouterV3(&metricsrouterv3.MetricsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(metricsRouterService).ToNot(BeNil())

				// Construct an instance of the ListTargetsOptions model
				listTargetsOptionsModel := new(metricsrouterv3.ListTargetsOptions)
				listTargetsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := metricsRouterService.ListTargets(listTargetsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				metricsRouterService.EnableRetries(0, 0)
				result, response, operationErr = metricsRouterService.ListTargets(listTargetsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListTargets(listTargetsOptions *ListTargetsOptions)`, func() {
		listTargetsPath := "/targets"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listTargetsPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"targets": [{"id": "f7dcfae6-e7c5-08ca-451b-fdfa696c9bb6", "name": "a-mr-target-us-south", "crn": "crn:v1:bluemix:public:metrics-router:us-south:a/0be5ad401ae913d8ff665d92680664ed:b6eec08b-5201-08ca-451b-cd71523e3626:target:f7dcfae6-e7c5-08ca-451b-fdfa696c9bb6", "destination_crn": "crn:v1:bluemix:public:sysdig-monitor:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::", "target_type": "sysdig_monitor", "region": "us-south", "created_at": "2021-05-18T20:15:12.353Z", "updated_at": "2021-05-18T20:15:12.353Z"}]}`)
				}))
			})
			It(`Invoke ListTargets successfully with retries`, func() {
				metricsRouterService, serviceErr := metricsrouterv3.NewMetricsRouterV3(&metricsrouterv3.MetricsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(metricsRouterService).ToNot(BeNil())
				metricsRouterService.EnableRetries(0, 0)

				// Construct an instance of the ListTargetsOptions model
				listTargetsOptionsModel := new(metricsrouterv3.ListTargetsOptions)
				listTargetsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := metricsRouterService.ListTargetsWithContext(ctx, listTargetsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				metricsRouterService.DisableRetries()
				result, response, operationErr := metricsRouterService.ListTargets(listTargetsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = metricsRouterService.ListTargetsWithContext(ctx, listTargetsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listTargetsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"targets": [{"id": "f7dcfae6-e7c5-08ca-451b-fdfa696c9bb6", "name": "a-mr-target-us-south", "crn": "crn:v1:bluemix:public:metrics-router:us-south:a/0be5ad401ae913d8ff665d92680664ed:b6eec08b-5201-08ca-451b-cd71523e3626:target:f7dcfae6-e7c5-08ca-451b-fdfa696c9bb6", "destination_crn": "crn:v1:bluemix:public:sysdig-monitor:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::", "target_type": "sysdig_monitor", "region": "us-south", "created_at": "2021-05-18T20:15:12.353Z", "updated_at": "2021-05-18T20:15:12.353Z"}]}`)
				}))
			})
			It(`Invoke ListTargets successfully`, func() {
				metricsRouterService, serviceErr := metricsrouterv3.NewMetricsRouterV3(&metricsrouterv3.MetricsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(metricsRouterService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := metricsRouterService.ListTargets(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListTargetsOptions model
				listTargetsOptionsModel := new(metricsrouterv3.ListTargetsOptions)
				listTargetsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = metricsRouterService.ListTargets(listTargetsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListTargets with error: Operation request error`, func() {
				metricsRouterService, serviceErr := metricsrouterv3.NewMetricsRouterV3(&metricsrouterv3.MetricsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(metricsRouterService).ToNot(BeNil())

				// Construct an instance of the ListTargetsOptions model
				listTargetsOptionsModel := new(metricsrouterv3.ListTargetsOptions)
				listTargetsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := metricsRouterService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := metricsRouterService.ListTargets(listTargetsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
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
			It(`Invoke ListTargets successfully`, func() {
				metricsRouterService, serviceErr := metricsrouterv3.NewMetricsRouterV3(&metricsrouterv3.MetricsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(metricsRouterService).ToNot(BeNil())

				// Construct an instance of the ListTargetsOptions model
				listTargetsOptionsModel := new(metricsrouterv3.ListTargetsOptions)
				listTargetsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := metricsRouterService.ListTargets(listTargetsOptionsModel)
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
	Describe(`GetTarget(getTargetOptions *GetTargetOptions) - Operation response error`, func() {
		getTargetPath := "/targets/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTargetPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetTarget with error: Operation response processing error`, func() {
				metricsRouterService, serviceErr := metricsrouterv3.NewMetricsRouterV3(&metricsrouterv3.MetricsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(metricsRouterService).ToNot(BeNil())

				// Construct an instance of the GetTargetOptions model
				getTargetOptionsModel := new(metricsrouterv3.GetTargetOptions)
				getTargetOptionsModel.ID = core.StringPtr("testString")
				getTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := metricsRouterService.GetTarget(getTargetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				metricsRouterService.EnableRetries(0, 0)
				result, response, operationErr = metricsRouterService.GetTarget(getTargetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetTarget(getTargetOptions *GetTargetOptions)`, func() {
		getTargetPath := "/targets/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTargetPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "f7dcfae6-e7c5-08ca-451b-fdfa696c9bb6", "name": "a-mr-target-us-south", "crn": "crn:v1:bluemix:public:metrics-router:us-south:a/0be5ad401ae913d8ff665d92680664ed:b6eec08b-5201-08ca-451b-cd71523e3626:target:f7dcfae6-e7c5-08ca-451b-fdfa696c9bb6", "destination_crn": "crn:v1:bluemix:public:sysdig-monitor:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::", "target_type": "sysdig_monitor", "region": "us-south", "created_at": "2021-05-18T20:15:12.353Z", "updated_at": "2021-05-18T20:15:12.353Z"}`)
				}))
			})
			It(`Invoke GetTarget successfully with retries`, func() {
				metricsRouterService, serviceErr := metricsrouterv3.NewMetricsRouterV3(&metricsrouterv3.MetricsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(metricsRouterService).ToNot(BeNil())
				metricsRouterService.EnableRetries(0, 0)

				// Construct an instance of the GetTargetOptions model
				getTargetOptionsModel := new(metricsrouterv3.GetTargetOptions)
				getTargetOptionsModel.ID = core.StringPtr("testString")
				getTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := metricsRouterService.GetTargetWithContext(ctx, getTargetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				metricsRouterService.DisableRetries()
				result, response, operationErr := metricsRouterService.GetTarget(getTargetOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = metricsRouterService.GetTargetWithContext(ctx, getTargetOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getTargetPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "f7dcfae6-e7c5-08ca-451b-fdfa696c9bb6", "name": "a-mr-target-us-south", "crn": "crn:v1:bluemix:public:metrics-router:us-south:a/0be5ad401ae913d8ff665d92680664ed:b6eec08b-5201-08ca-451b-cd71523e3626:target:f7dcfae6-e7c5-08ca-451b-fdfa696c9bb6", "destination_crn": "crn:v1:bluemix:public:sysdig-monitor:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::", "target_type": "sysdig_monitor", "region": "us-south", "created_at": "2021-05-18T20:15:12.353Z", "updated_at": "2021-05-18T20:15:12.353Z"}`)
				}))
			})
			It(`Invoke GetTarget successfully`, func() {
				metricsRouterService, serviceErr := metricsrouterv3.NewMetricsRouterV3(&metricsrouterv3.MetricsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(metricsRouterService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := metricsRouterService.GetTarget(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetTargetOptions model
				getTargetOptionsModel := new(metricsrouterv3.GetTargetOptions)
				getTargetOptionsModel.ID = core.StringPtr("testString")
				getTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = metricsRouterService.GetTarget(getTargetOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetTarget with error: Operation validation and request error`, func() {
				metricsRouterService, serviceErr := metricsrouterv3.NewMetricsRouterV3(&metricsrouterv3.MetricsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(metricsRouterService).ToNot(BeNil())

				// Construct an instance of the GetTargetOptions model
				getTargetOptionsModel := new(metricsrouterv3.GetTargetOptions)
				getTargetOptionsModel.ID = core.StringPtr("testString")
				getTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := metricsRouterService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := metricsRouterService.GetTarget(getTargetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetTargetOptions model with no property values
				getTargetOptionsModelNew := new(metricsrouterv3.GetTargetOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = metricsRouterService.GetTarget(getTargetOptionsModelNew)
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
			It(`Invoke GetTarget successfully`, func() {
				metricsRouterService, serviceErr := metricsrouterv3.NewMetricsRouterV3(&metricsrouterv3.MetricsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(metricsRouterService).ToNot(BeNil())

				// Construct an instance of the GetTargetOptions model
				getTargetOptionsModel := new(metricsrouterv3.GetTargetOptions)
				getTargetOptionsModel.ID = core.StringPtr("testString")
				getTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := metricsRouterService.GetTarget(getTargetOptionsModel)
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
	Describe(`UpdateTarget(updateTargetOptions *UpdateTargetOptions) - Operation response error`, func() {
		updateTargetPath := "/targets/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateTargetPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateTarget with error: Operation response processing error`, func() {
				metricsRouterService, serviceErr := metricsrouterv3.NewMetricsRouterV3(&metricsrouterv3.MetricsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(metricsRouterService).ToNot(BeNil())

				// Construct an instance of the UpdateTargetOptions model
				updateTargetOptionsModel := new(metricsrouterv3.UpdateTargetOptions)
				updateTargetOptionsModel.ID = core.StringPtr("testString")
				updateTargetOptionsModel.Name = core.StringPtr("my-mr-target")
				updateTargetOptionsModel.DestinationCRN = core.StringPtr("crn:v1:bluemix:public:sysdig-monitor:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::")
				updateTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := metricsRouterService.UpdateTarget(updateTargetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				metricsRouterService.EnableRetries(0, 0)
				result, response, operationErr = metricsRouterService.UpdateTarget(updateTargetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateTarget(updateTargetOptions *UpdateTargetOptions)`, func() {
		updateTargetPath := "/targets/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateTargetPath))
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

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "f7dcfae6-e7c5-08ca-451b-fdfa696c9bb6", "name": "a-mr-target-us-south", "crn": "crn:v1:bluemix:public:metrics-router:us-south:a/0be5ad401ae913d8ff665d92680664ed:b6eec08b-5201-08ca-451b-cd71523e3626:target:f7dcfae6-e7c5-08ca-451b-fdfa696c9bb6", "destination_crn": "crn:v1:bluemix:public:sysdig-monitor:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::", "target_type": "sysdig_monitor", "region": "us-south", "created_at": "2021-05-18T20:15:12.353Z", "updated_at": "2021-05-18T20:15:12.353Z"}`)
				}))
			})
			It(`Invoke UpdateTarget successfully with retries`, func() {
				metricsRouterService, serviceErr := metricsrouterv3.NewMetricsRouterV3(&metricsrouterv3.MetricsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(metricsRouterService).ToNot(BeNil())
				metricsRouterService.EnableRetries(0, 0)

				// Construct an instance of the UpdateTargetOptions model
				updateTargetOptionsModel := new(metricsrouterv3.UpdateTargetOptions)
				updateTargetOptionsModel.ID = core.StringPtr("testString")
				updateTargetOptionsModel.Name = core.StringPtr("my-mr-target")
				updateTargetOptionsModel.DestinationCRN = core.StringPtr("crn:v1:bluemix:public:sysdig-monitor:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::")
				updateTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := metricsRouterService.UpdateTargetWithContext(ctx, updateTargetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				metricsRouterService.DisableRetries()
				result, response, operationErr := metricsRouterService.UpdateTarget(updateTargetOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = metricsRouterService.UpdateTargetWithContext(ctx, updateTargetOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateTargetPath))
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

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "f7dcfae6-e7c5-08ca-451b-fdfa696c9bb6", "name": "a-mr-target-us-south", "crn": "crn:v1:bluemix:public:metrics-router:us-south:a/0be5ad401ae913d8ff665d92680664ed:b6eec08b-5201-08ca-451b-cd71523e3626:target:f7dcfae6-e7c5-08ca-451b-fdfa696c9bb6", "destination_crn": "crn:v1:bluemix:public:sysdig-monitor:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::", "target_type": "sysdig_monitor", "region": "us-south", "created_at": "2021-05-18T20:15:12.353Z", "updated_at": "2021-05-18T20:15:12.353Z"}`)
				}))
			})
			It(`Invoke UpdateTarget successfully`, func() {
				metricsRouterService, serviceErr := metricsrouterv3.NewMetricsRouterV3(&metricsrouterv3.MetricsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(metricsRouterService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := metricsRouterService.UpdateTarget(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateTargetOptions model
				updateTargetOptionsModel := new(metricsrouterv3.UpdateTargetOptions)
				updateTargetOptionsModel.ID = core.StringPtr("testString")
				updateTargetOptionsModel.Name = core.StringPtr("my-mr-target")
				updateTargetOptionsModel.DestinationCRN = core.StringPtr("crn:v1:bluemix:public:sysdig-monitor:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::")
				updateTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = metricsRouterService.UpdateTarget(updateTargetOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateTarget with error: Operation validation and request error`, func() {
				metricsRouterService, serviceErr := metricsrouterv3.NewMetricsRouterV3(&metricsrouterv3.MetricsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(metricsRouterService).ToNot(BeNil())

				// Construct an instance of the UpdateTargetOptions model
				updateTargetOptionsModel := new(metricsrouterv3.UpdateTargetOptions)
				updateTargetOptionsModel.ID = core.StringPtr("testString")
				updateTargetOptionsModel.Name = core.StringPtr("my-mr-target")
				updateTargetOptionsModel.DestinationCRN = core.StringPtr("crn:v1:bluemix:public:sysdig-monitor:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::")
				updateTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := metricsRouterService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := metricsRouterService.UpdateTarget(updateTargetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateTargetOptions model with no property values
				updateTargetOptionsModelNew := new(metricsrouterv3.UpdateTargetOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = metricsRouterService.UpdateTarget(updateTargetOptionsModelNew)
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
			It(`Invoke UpdateTarget successfully`, func() {
				metricsRouterService, serviceErr := metricsrouterv3.NewMetricsRouterV3(&metricsrouterv3.MetricsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(metricsRouterService).ToNot(BeNil())

				// Construct an instance of the UpdateTargetOptions model
				updateTargetOptionsModel := new(metricsrouterv3.UpdateTargetOptions)
				updateTargetOptionsModel.ID = core.StringPtr("testString")
				updateTargetOptionsModel.Name = core.StringPtr("my-mr-target")
				updateTargetOptionsModel.DestinationCRN = core.StringPtr("crn:v1:bluemix:public:sysdig-monitor:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::")
				updateTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := metricsRouterService.UpdateTarget(updateTargetOptionsModel)
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
	Describe(`DeleteTarget(deleteTargetOptions *DeleteTargetOptions)`, func() {
		deleteTargetPath := "/targets/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteTargetPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteTarget successfully`, func() {
				metricsRouterService, serviceErr := metricsrouterv3.NewMetricsRouterV3(&metricsrouterv3.MetricsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(metricsRouterService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := metricsRouterService.DeleteTarget(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteTargetOptions model
				deleteTargetOptionsModel := new(metricsrouterv3.DeleteTargetOptions)
				deleteTargetOptionsModel.ID = core.StringPtr("testString")
				deleteTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = metricsRouterService.DeleteTarget(deleteTargetOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteTarget with error: Operation validation and request error`, func() {
				metricsRouterService, serviceErr := metricsrouterv3.NewMetricsRouterV3(&metricsrouterv3.MetricsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(metricsRouterService).ToNot(BeNil())

				// Construct an instance of the DeleteTargetOptions model
				deleteTargetOptionsModel := new(metricsrouterv3.DeleteTargetOptions)
				deleteTargetOptionsModel.ID = core.StringPtr("testString")
				deleteTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := metricsRouterService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := metricsRouterService.DeleteTarget(deleteTargetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteTargetOptions model with no property values
				deleteTargetOptionsModelNew := new(metricsrouterv3.DeleteTargetOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = metricsRouterService.DeleteTarget(deleteTargetOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateRoute(createRouteOptions *CreateRouteOptions) - Operation response error`, func() {
		createRoutePath := "/routes"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createRoutePath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateRoute with error: Operation response processing error`, func() {
				metricsRouterService, serviceErr := metricsrouterv3.NewMetricsRouterV3(&metricsrouterv3.MetricsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(metricsRouterService).ToNot(BeNil())

				// Construct an instance of the TargetIdentity model
				targetIdentityModel := new(metricsrouterv3.TargetIdentity)
				targetIdentityModel.ID = core.StringPtr("c3af557f-fb0e-4476-85c3-0889e7fe7bc4")

				// Construct an instance of the InclusionFilterPrototype model
				inclusionFilterPrototypeModel := new(metricsrouterv3.InclusionFilterPrototype)
				inclusionFilterPrototypeModel.Operand = core.StringPtr("location")
				inclusionFilterPrototypeModel.Operator = core.StringPtr("is")
				inclusionFilterPrototypeModel.Values = []string{"us-south"}

				// Construct an instance of the RulePrototype model
				rulePrototypeModel := new(metricsrouterv3.RulePrototype)
				rulePrototypeModel.Action = core.StringPtr("send")
				rulePrototypeModel.Targets = []metricsrouterv3.TargetIdentity{*targetIdentityModel}
				rulePrototypeModel.InclusionFilters = []metricsrouterv3.InclusionFilterPrototype{*inclusionFilterPrototypeModel}

				// Construct an instance of the CreateRouteOptions model
				createRouteOptionsModel := new(metricsrouterv3.CreateRouteOptions)
				createRouteOptionsModel.Name = core.StringPtr("my-route")
				createRouteOptionsModel.Rules = []metricsrouterv3.RulePrototype{*rulePrototypeModel}
				createRouteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := metricsRouterService.CreateRoute(createRouteOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				metricsRouterService.EnableRetries(0, 0)
				result, response, operationErr = metricsRouterService.CreateRoute(createRouteOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateRoute(createRouteOptions *CreateRouteOptions)`, func() {
		createRoutePath := "/routes"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createRoutePath))
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
					fmt.Fprintf(res, "%s", `{"id": "c3af557f-fb0e-4476-85c3-0889e7fe7bc4", "name": "my-route", "crn": "crn:v1:bluemix:public:metrics-router:global:a/0be5ad401ae913d8ff665d92680664ed:b6eec08b-5201-08ca-451b-cd71523e3626:route:c3af557f-fb0e-4476-85c3-0889e7fe7bc4", "rules": [{"action": "send", "targets": [{"id": "c3af557f-fb0e-4476-85c3-0889e7fe7bc4", "crn": "crn:v1:bluemix:public:sysdig-monitor:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::", "name": "a-mr-target-us-south", "target_type": "sysdig_monitor"}], "inclusion_filters": [{"operand": "location", "operator": "is", "values": ["us-south"]}]}], "created_at": "2021-05-18T20:15:12.353Z", "updated_at": "2021-05-18T20:15:12.353Z"}`)
				}))
			})
			It(`Invoke CreateRoute successfully with retries`, func() {
				metricsRouterService, serviceErr := metricsrouterv3.NewMetricsRouterV3(&metricsrouterv3.MetricsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(metricsRouterService).ToNot(BeNil())
				metricsRouterService.EnableRetries(0, 0)

				// Construct an instance of the TargetIdentity model
				targetIdentityModel := new(metricsrouterv3.TargetIdentity)
				targetIdentityModel.ID = core.StringPtr("c3af557f-fb0e-4476-85c3-0889e7fe7bc4")

				// Construct an instance of the InclusionFilterPrototype model
				inclusionFilterPrototypeModel := new(metricsrouterv3.InclusionFilterPrototype)
				inclusionFilterPrototypeModel.Operand = core.StringPtr("location")
				inclusionFilterPrototypeModel.Operator = core.StringPtr("is")
				inclusionFilterPrototypeModel.Values = []string{"us-south"}

				// Construct an instance of the RulePrototype model
				rulePrototypeModel := new(metricsrouterv3.RulePrototype)
				rulePrototypeModel.Action = core.StringPtr("send")
				rulePrototypeModel.Targets = []metricsrouterv3.TargetIdentity{*targetIdentityModel}
				rulePrototypeModel.InclusionFilters = []metricsrouterv3.InclusionFilterPrototype{*inclusionFilterPrototypeModel}

				// Construct an instance of the CreateRouteOptions model
				createRouteOptionsModel := new(metricsrouterv3.CreateRouteOptions)
				createRouteOptionsModel.Name = core.StringPtr("my-route")
				createRouteOptionsModel.Rules = []metricsrouterv3.RulePrototype{*rulePrototypeModel}
				createRouteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := metricsRouterService.CreateRouteWithContext(ctx, createRouteOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				metricsRouterService.DisableRetries()
				result, response, operationErr := metricsRouterService.CreateRoute(createRouteOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = metricsRouterService.CreateRouteWithContext(ctx, createRouteOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createRoutePath))
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
					fmt.Fprintf(res, "%s", `{"id": "c3af557f-fb0e-4476-85c3-0889e7fe7bc4", "name": "my-route", "crn": "crn:v1:bluemix:public:metrics-router:global:a/0be5ad401ae913d8ff665d92680664ed:b6eec08b-5201-08ca-451b-cd71523e3626:route:c3af557f-fb0e-4476-85c3-0889e7fe7bc4", "rules": [{"action": "send", "targets": [{"id": "c3af557f-fb0e-4476-85c3-0889e7fe7bc4", "crn": "crn:v1:bluemix:public:sysdig-monitor:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::", "name": "a-mr-target-us-south", "target_type": "sysdig_monitor"}], "inclusion_filters": [{"operand": "location", "operator": "is", "values": ["us-south"]}]}], "created_at": "2021-05-18T20:15:12.353Z", "updated_at": "2021-05-18T20:15:12.353Z"}`)
				}))
			})
			It(`Invoke CreateRoute successfully`, func() {
				metricsRouterService, serviceErr := metricsrouterv3.NewMetricsRouterV3(&metricsrouterv3.MetricsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(metricsRouterService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := metricsRouterService.CreateRoute(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the TargetIdentity model
				targetIdentityModel := new(metricsrouterv3.TargetIdentity)
				targetIdentityModel.ID = core.StringPtr("c3af557f-fb0e-4476-85c3-0889e7fe7bc4")

				// Construct an instance of the InclusionFilterPrototype model
				inclusionFilterPrototypeModel := new(metricsrouterv3.InclusionFilterPrototype)
				inclusionFilterPrototypeModel.Operand = core.StringPtr("location")
				inclusionFilterPrototypeModel.Operator = core.StringPtr("is")
				inclusionFilterPrototypeModel.Values = []string{"us-south"}

				// Construct an instance of the RulePrototype model
				rulePrototypeModel := new(metricsrouterv3.RulePrototype)
				rulePrototypeModel.Action = core.StringPtr("send")
				rulePrototypeModel.Targets = []metricsrouterv3.TargetIdentity{*targetIdentityModel}
				rulePrototypeModel.InclusionFilters = []metricsrouterv3.InclusionFilterPrototype{*inclusionFilterPrototypeModel}

				// Construct an instance of the CreateRouteOptions model
				createRouteOptionsModel := new(metricsrouterv3.CreateRouteOptions)
				createRouteOptionsModel.Name = core.StringPtr("my-route")
				createRouteOptionsModel.Rules = []metricsrouterv3.RulePrototype{*rulePrototypeModel}
				createRouteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = metricsRouterService.CreateRoute(createRouteOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateRoute with error: Operation validation and request error`, func() {
				metricsRouterService, serviceErr := metricsrouterv3.NewMetricsRouterV3(&metricsrouterv3.MetricsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(metricsRouterService).ToNot(BeNil())

				// Construct an instance of the TargetIdentity model
				targetIdentityModel := new(metricsrouterv3.TargetIdentity)
				targetIdentityModel.ID = core.StringPtr("c3af557f-fb0e-4476-85c3-0889e7fe7bc4")

				// Construct an instance of the InclusionFilterPrototype model
				inclusionFilterPrototypeModel := new(metricsrouterv3.InclusionFilterPrototype)
				inclusionFilterPrototypeModel.Operand = core.StringPtr("location")
				inclusionFilterPrototypeModel.Operator = core.StringPtr("is")
				inclusionFilterPrototypeModel.Values = []string{"us-south"}

				// Construct an instance of the RulePrototype model
				rulePrototypeModel := new(metricsrouterv3.RulePrototype)
				rulePrototypeModel.Action = core.StringPtr("send")
				rulePrototypeModel.Targets = []metricsrouterv3.TargetIdentity{*targetIdentityModel}
				rulePrototypeModel.InclusionFilters = []metricsrouterv3.InclusionFilterPrototype{*inclusionFilterPrototypeModel}

				// Construct an instance of the CreateRouteOptions model
				createRouteOptionsModel := new(metricsrouterv3.CreateRouteOptions)
				createRouteOptionsModel.Name = core.StringPtr("my-route")
				createRouteOptionsModel.Rules = []metricsrouterv3.RulePrototype{*rulePrototypeModel}
				createRouteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := metricsRouterService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := metricsRouterService.CreateRoute(createRouteOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateRouteOptions model with no property values
				createRouteOptionsModelNew := new(metricsrouterv3.CreateRouteOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = metricsRouterService.CreateRoute(createRouteOptionsModelNew)
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
			It(`Invoke CreateRoute successfully`, func() {
				metricsRouterService, serviceErr := metricsrouterv3.NewMetricsRouterV3(&metricsrouterv3.MetricsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(metricsRouterService).ToNot(BeNil())

				// Construct an instance of the TargetIdentity model
				targetIdentityModel := new(metricsrouterv3.TargetIdentity)
				targetIdentityModel.ID = core.StringPtr("c3af557f-fb0e-4476-85c3-0889e7fe7bc4")

				// Construct an instance of the InclusionFilterPrototype model
				inclusionFilterPrototypeModel := new(metricsrouterv3.InclusionFilterPrototype)
				inclusionFilterPrototypeModel.Operand = core.StringPtr("location")
				inclusionFilterPrototypeModel.Operator = core.StringPtr("is")
				inclusionFilterPrototypeModel.Values = []string{"us-south"}

				// Construct an instance of the RulePrototype model
				rulePrototypeModel := new(metricsrouterv3.RulePrototype)
				rulePrototypeModel.Action = core.StringPtr("send")
				rulePrototypeModel.Targets = []metricsrouterv3.TargetIdentity{*targetIdentityModel}
				rulePrototypeModel.InclusionFilters = []metricsrouterv3.InclusionFilterPrototype{*inclusionFilterPrototypeModel}

				// Construct an instance of the CreateRouteOptions model
				createRouteOptionsModel := new(metricsrouterv3.CreateRouteOptions)
				createRouteOptionsModel.Name = core.StringPtr("my-route")
				createRouteOptionsModel.Rules = []metricsrouterv3.RulePrototype{*rulePrototypeModel}
				createRouteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := metricsRouterService.CreateRoute(createRouteOptionsModel)
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
	Describe(`ListRoutes(listRoutesOptions *ListRoutesOptions) - Operation response error`, func() {
		listRoutesPath := "/routes"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listRoutesPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListRoutes with error: Operation response processing error`, func() {
				metricsRouterService, serviceErr := metricsrouterv3.NewMetricsRouterV3(&metricsrouterv3.MetricsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(metricsRouterService).ToNot(BeNil())

				// Construct an instance of the ListRoutesOptions model
				listRoutesOptionsModel := new(metricsrouterv3.ListRoutesOptions)
				listRoutesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := metricsRouterService.ListRoutes(listRoutesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				metricsRouterService.EnableRetries(0, 0)
				result, response, operationErr = metricsRouterService.ListRoutes(listRoutesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListRoutes(listRoutesOptions *ListRoutesOptions)`, func() {
		listRoutesPath := "/routes"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listRoutesPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"routes": [{"id": "c3af557f-fb0e-4476-85c3-0889e7fe7bc4", "name": "my-route", "crn": "crn:v1:bluemix:public:metrics-router:global:a/0be5ad401ae913d8ff665d92680664ed:b6eec08b-5201-08ca-451b-cd71523e3626:route:c3af557f-fb0e-4476-85c3-0889e7fe7bc4", "rules": [{"action": "send", "targets": [{"id": "c3af557f-fb0e-4476-85c3-0889e7fe7bc4", "crn": "crn:v1:bluemix:public:sysdig-monitor:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::", "name": "a-mr-target-us-south", "target_type": "sysdig_monitor"}], "inclusion_filters": [{"operand": "location", "operator": "is", "values": ["us-south"]}]}], "created_at": "2021-05-18T20:15:12.353Z", "updated_at": "2021-05-18T20:15:12.353Z"}]}`)
				}))
			})
			It(`Invoke ListRoutes successfully with retries`, func() {
				metricsRouterService, serviceErr := metricsrouterv3.NewMetricsRouterV3(&metricsrouterv3.MetricsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(metricsRouterService).ToNot(BeNil())
				metricsRouterService.EnableRetries(0, 0)

				// Construct an instance of the ListRoutesOptions model
				listRoutesOptionsModel := new(metricsrouterv3.ListRoutesOptions)
				listRoutesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := metricsRouterService.ListRoutesWithContext(ctx, listRoutesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				metricsRouterService.DisableRetries()
				result, response, operationErr := metricsRouterService.ListRoutes(listRoutesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = metricsRouterService.ListRoutesWithContext(ctx, listRoutesOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listRoutesPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"routes": [{"id": "c3af557f-fb0e-4476-85c3-0889e7fe7bc4", "name": "my-route", "crn": "crn:v1:bluemix:public:metrics-router:global:a/0be5ad401ae913d8ff665d92680664ed:b6eec08b-5201-08ca-451b-cd71523e3626:route:c3af557f-fb0e-4476-85c3-0889e7fe7bc4", "rules": [{"action": "send", "targets": [{"id": "c3af557f-fb0e-4476-85c3-0889e7fe7bc4", "crn": "crn:v1:bluemix:public:sysdig-monitor:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::", "name": "a-mr-target-us-south", "target_type": "sysdig_monitor"}], "inclusion_filters": [{"operand": "location", "operator": "is", "values": ["us-south"]}]}], "created_at": "2021-05-18T20:15:12.353Z", "updated_at": "2021-05-18T20:15:12.353Z"}]}`)
				}))
			})
			It(`Invoke ListRoutes successfully`, func() {
				metricsRouterService, serviceErr := metricsrouterv3.NewMetricsRouterV3(&metricsrouterv3.MetricsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(metricsRouterService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := metricsRouterService.ListRoutes(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListRoutesOptions model
				listRoutesOptionsModel := new(metricsrouterv3.ListRoutesOptions)
				listRoutesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = metricsRouterService.ListRoutes(listRoutesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListRoutes with error: Operation request error`, func() {
				metricsRouterService, serviceErr := metricsrouterv3.NewMetricsRouterV3(&metricsrouterv3.MetricsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(metricsRouterService).ToNot(BeNil())

				// Construct an instance of the ListRoutesOptions model
				listRoutesOptionsModel := new(metricsrouterv3.ListRoutesOptions)
				listRoutesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := metricsRouterService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := metricsRouterService.ListRoutes(listRoutesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
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
			It(`Invoke ListRoutes successfully`, func() {
				metricsRouterService, serviceErr := metricsrouterv3.NewMetricsRouterV3(&metricsrouterv3.MetricsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(metricsRouterService).ToNot(BeNil())

				// Construct an instance of the ListRoutesOptions model
				listRoutesOptionsModel := new(metricsrouterv3.ListRoutesOptions)
				listRoutesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := metricsRouterService.ListRoutes(listRoutesOptionsModel)
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
	Describe(`GetRoute(getRouteOptions *GetRouteOptions) - Operation response error`, func() {
		getRoutePath := "/routes/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getRoutePath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetRoute with error: Operation response processing error`, func() {
				metricsRouterService, serviceErr := metricsrouterv3.NewMetricsRouterV3(&metricsrouterv3.MetricsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(metricsRouterService).ToNot(BeNil())

				// Construct an instance of the GetRouteOptions model
				getRouteOptionsModel := new(metricsrouterv3.GetRouteOptions)
				getRouteOptionsModel.ID = core.StringPtr("testString")
				getRouteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := metricsRouterService.GetRoute(getRouteOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				metricsRouterService.EnableRetries(0, 0)
				result, response, operationErr = metricsRouterService.GetRoute(getRouteOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetRoute(getRouteOptions *GetRouteOptions)`, func() {
		getRoutePath := "/routes/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getRoutePath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "c3af557f-fb0e-4476-85c3-0889e7fe7bc4", "name": "my-route", "crn": "crn:v1:bluemix:public:metrics-router:global:a/0be5ad401ae913d8ff665d92680664ed:b6eec08b-5201-08ca-451b-cd71523e3626:route:c3af557f-fb0e-4476-85c3-0889e7fe7bc4", "rules": [{"action": "send", "targets": [{"id": "c3af557f-fb0e-4476-85c3-0889e7fe7bc4", "crn": "crn:v1:bluemix:public:sysdig-monitor:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::", "name": "a-mr-target-us-south", "target_type": "sysdig_monitor"}], "inclusion_filters": [{"operand": "location", "operator": "is", "values": ["us-south"]}]}], "created_at": "2021-05-18T20:15:12.353Z", "updated_at": "2021-05-18T20:15:12.353Z"}`)
				}))
			})
			It(`Invoke GetRoute successfully with retries`, func() {
				metricsRouterService, serviceErr := metricsrouterv3.NewMetricsRouterV3(&metricsrouterv3.MetricsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(metricsRouterService).ToNot(BeNil())
				metricsRouterService.EnableRetries(0, 0)

				// Construct an instance of the GetRouteOptions model
				getRouteOptionsModel := new(metricsrouterv3.GetRouteOptions)
				getRouteOptionsModel.ID = core.StringPtr("testString")
				getRouteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := metricsRouterService.GetRouteWithContext(ctx, getRouteOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				metricsRouterService.DisableRetries()
				result, response, operationErr := metricsRouterService.GetRoute(getRouteOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = metricsRouterService.GetRouteWithContext(ctx, getRouteOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getRoutePath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "c3af557f-fb0e-4476-85c3-0889e7fe7bc4", "name": "my-route", "crn": "crn:v1:bluemix:public:metrics-router:global:a/0be5ad401ae913d8ff665d92680664ed:b6eec08b-5201-08ca-451b-cd71523e3626:route:c3af557f-fb0e-4476-85c3-0889e7fe7bc4", "rules": [{"action": "send", "targets": [{"id": "c3af557f-fb0e-4476-85c3-0889e7fe7bc4", "crn": "crn:v1:bluemix:public:sysdig-monitor:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::", "name": "a-mr-target-us-south", "target_type": "sysdig_monitor"}], "inclusion_filters": [{"operand": "location", "operator": "is", "values": ["us-south"]}]}], "created_at": "2021-05-18T20:15:12.353Z", "updated_at": "2021-05-18T20:15:12.353Z"}`)
				}))
			})
			It(`Invoke GetRoute successfully`, func() {
				metricsRouterService, serviceErr := metricsrouterv3.NewMetricsRouterV3(&metricsrouterv3.MetricsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(metricsRouterService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := metricsRouterService.GetRoute(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetRouteOptions model
				getRouteOptionsModel := new(metricsrouterv3.GetRouteOptions)
				getRouteOptionsModel.ID = core.StringPtr("testString")
				getRouteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = metricsRouterService.GetRoute(getRouteOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetRoute with error: Operation validation and request error`, func() {
				metricsRouterService, serviceErr := metricsrouterv3.NewMetricsRouterV3(&metricsrouterv3.MetricsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(metricsRouterService).ToNot(BeNil())

				// Construct an instance of the GetRouteOptions model
				getRouteOptionsModel := new(metricsrouterv3.GetRouteOptions)
				getRouteOptionsModel.ID = core.StringPtr("testString")
				getRouteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := metricsRouterService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := metricsRouterService.GetRoute(getRouteOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetRouteOptions model with no property values
				getRouteOptionsModelNew := new(metricsrouterv3.GetRouteOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = metricsRouterService.GetRoute(getRouteOptionsModelNew)
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
			It(`Invoke GetRoute successfully`, func() {
				metricsRouterService, serviceErr := metricsrouterv3.NewMetricsRouterV3(&metricsrouterv3.MetricsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(metricsRouterService).ToNot(BeNil())

				// Construct an instance of the GetRouteOptions model
				getRouteOptionsModel := new(metricsrouterv3.GetRouteOptions)
				getRouteOptionsModel.ID = core.StringPtr("testString")
				getRouteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := metricsRouterService.GetRoute(getRouteOptionsModel)
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
	Describe(`UpdateRoute(updateRouteOptions *UpdateRouteOptions) - Operation response error`, func() {
		updateRoutePath := "/routes/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateRoutePath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateRoute with error: Operation response processing error`, func() {
				metricsRouterService, serviceErr := metricsrouterv3.NewMetricsRouterV3(&metricsrouterv3.MetricsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(metricsRouterService).ToNot(BeNil())

				// Construct an instance of the TargetIdentity model
				targetIdentityModel := new(metricsrouterv3.TargetIdentity)
				targetIdentityModel.ID = core.StringPtr("c3af557f-fb0e-4476-85c3-0889e7fe7bc4")

				// Construct an instance of the InclusionFilterPrototype model
				inclusionFilterPrototypeModel := new(metricsrouterv3.InclusionFilterPrototype)
				inclusionFilterPrototypeModel.Operand = core.StringPtr("location")
				inclusionFilterPrototypeModel.Operator = core.StringPtr("is")
				inclusionFilterPrototypeModel.Values = []string{"us-south"}

				// Construct an instance of the RulePrototype model
				rulePrototypeModel := new(metricsrouterv3.RulePrototype)
				rulePrototypeModel.Action = core.StringPtr("send")
				rulePrototypeModel.Targets = []metricsrouterv3.TargetIdentity{*targetIdentityModel}
				rulePrototypeModel.InclusionFilters = []metricsrouterv3.InclusionFilterPrototype{*inclusionFilterPrototypeModel}

				// Construct an instance of the UpdateRouteOptions model
				updateRouteOptionsModel := new(metricsrouterv3.UpdateRouteOptions)
				updateRouteOptionsModel.ID = core.StringPtr("testString")
				updateRouteOptionsModel.Name = core.StringPtr("my-route")
				updateRouteOptionsModel.Rules = []metricsrouterv3.RulePrototype{*rulePrototypeModel}
				updateRouteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := metricsRouterService.UpdateRoute(updateRouteOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				metricsRouterService.EnableRetries(0, 0)
				result, response, operationErr = metricsRouterService.UpdateRoute(updateRouteOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateRoute(updateRouteOptions *UpdateRouteOptions)`, func() {
		updateRoutePath := "/routes/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateRoutePath))
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

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "c3af557f-fb0e-4476-85c3-0889e7fe7bc4", "name": "my-route", "crn": "crn:v1:bluemix:public:metrics-router:global:a/0be5ad401ae913d8ff665d92680664ed:b6eec08b-5201-08ca-451b-cd71523e3626:route:c3af557f-fb0e-4476-85c3-0889e7fe7bc4", "rules": [{"action": "send", "targets": [{"id": "c3af557f-fb0e-4476-85c3-0889e7fe7bc4", "crn": "crn:v1:bluemix:public:sysdig-monitor:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::", "name": "a-mr-target-us-south", "target_type": "sysdig_monitor"}], "inclusion_filters": [{"operand": "location", "operator": "is", "values": ["us-south"]}]}], "created_at": "2021-05-18T20:15:12.353Z", "updated_at": "2021-05-18T20:15:12.353Z"}`)
				}))
			})
			It(`Invoke UpdateRoute successfully with retries`, func() {
				metricsRouterService, serviceErr := metricsrouterv3.NewMetricsRouterV3(&metricsrouterv3.MetricsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(metricsRouterService).ToNot(BeNil())
				metricsRouterService.EnableRetries(0, 0)

				// Construct an instance of the TargetIdentity model
				targetIdentityModel := new(metricsrouterv3.TargetIdentity)
				targetIdentityModel.ID = core.StringPtr("c3af557f-fb0e-4476-85c3-0889e7fe7bc4")

				// Construct an instance of the InclusionFilterPrototype model
				inclusionFilterPrototypeModel := new(metricsrouterv3.InclusionFilterPrototype)
				inclusionFilterPrototypeModel.Operand = core.StringPtr("location")
				inclusionFilterPrototypeModel.Operator = core.StringPtr("is")
				inclusionFilterPrototypeModel.Values = []string{"us-south"}

				// Construct an instance of the RulePrototype model
				rulePrototypeModel := new(metricsrouterv3.RulePrototype)
				rulePrototypeModel.Action = core.StringPtr("send")
				rulePrototypeModel.Targets = []metricsrouterv3.TargetIdentity{*targetIdentityModel}
				rulePrototypeModel.InclusionFilters = []metricsrouterv3.InclusionFilterPrototype{*inclusionFilterPrototypeModel}

				// Construct an instance of the UpdateRouteOptions model
				updateRouteOptionsModel := new(metricsrouterv3.UpdateRouteOptions)
				updateRouteOptionsModel.ID = core.StringPtr("testString")
				updateRouteOptionsModel.Name = core.StringPtr("my-route")
				updateRouteOptionsModel.Rules = []metricsrouterv3.RulePrototype{*rulePrototypeModel}
				updateRouteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := metricsRouterService.UpdateRouteWithContext(ctx, updateRouteOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				metricsRouterService.DisableRetries()
				result, response, operationErr := metricsRouterService.UpdateRoute(updateRouteOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = metricsRouterService.UpdateRouteWithContext(ctx, updateRouteOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateRoutePath))
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

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "c3af557f-fb0e-4476-85c3-0889e7fe7bc4", "name": "my-route", "crn": "crn:v1:bluemix:public:metrics-router:global:a/0be5ad401ae913d8ff665d92680664ed:b6eec08b-5201-08ca-451b-cd71523e3626:route:c3af557f-fb0e-4476-85c3-0889e7fe7bc4", "rules": [{"action": "send", "targets": [{"id": "c3af557f-fb0e-4476-85c3-0889e7fe7bc4", "crn": "crn:v1:bluemix:public:sysdig-monitor:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::", "name": "a-mr-target-us-south", "target_type": "sysdig_monitor"}], "inclusion_filters": [{"operand": "location", "operator": "is", "values": ["us-south"]}]}], "created_at": "2021-05-18T20:15:12.353Z", "updated_at": "2021-05-18T20:15:12.353Z"}`)
				}))
			})
			It(`Invoke UpdateRoute successfully`, func() {
				metricsRouterService, serviceErr := metricsrouterv3.NewMetricsRouterV3(&metricsrouterv3.MetricsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(metricsRouterService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := metricsRouterService.UpdateRoute(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the TargetIdentity model
				targetIdentityModel := new(metricsrouterv3.TargetIdentity)
				targetIdentityModel.ID = core.StringPtr("c3af557f-fb0e-4476-85c3-0889e7fe7bc4")

				// Construct an instance of the InclusionFilterPrototype model
				inclusionFilterPrototypeModel := new(metricsrouterv3.InclusionFilterPrototype)
				inclusionFilterPrototypeModel.Operand = core.StringPtr("location")
				inclusionFilterPrototypeModel.Operator = core.StringPtr("is")
				inclusionFilterPrototypeModel.Values = []string{"us-south"}

				// Construct an instance of the RulePrototype model
				rulePrototypeModel := new(metricsrouterv3.RulePrototype)
				rulePrototypeModel.Action = core.StringPtr("send")
				rulePrototypeModel.Targets = []metricsrouterv3.TargetIdentity{*targetIdentityModel}
				rulePrototypeModel.InclusionFilters = []metricsrouterv3.InclusionFilterPrototype{*inclusionFilterPrototypeModel}

				// Construct an instance of the UpdateRouteOptions model
				updateRouteOptionsModel := new(metricsrouterv3.UpdateRouteOptions)
				updateRouteOptionsModel.ID = core.StringPtr("testString")
				updateRouteOptionsModel.Name = core.StringPtr("my-route")
				updateRouteOptionsModel.Rules = []metricsrouterv3.RulePrototype{*rulePrototypeModel}
				updateRouteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = metricsRouterService.UpdateRoute(updateRouteOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateRoute with error: Operation validation and request error`, func() {
				metricsRouterService, serviceErr := metricsrouterv3.NewMetricsRouterV3(&metricsrouterv3.MetricsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(metricsRouterService).ToNot(BeNil())

				// Construct an instance of the TargetIdentity model
				targetIdentityModel := new(metricsrouterv3.TargetIdentity)
				targetIdentityModel.ID = core.StringPtr("c3af557f-fb0e-4476-85c3-0889e7fe7bc4")

				// Construct an instance of the InclusionFilterPrototype model
				inclusionFilterPrototypeModel := new(metricsrouterv3.InclusionFilterPrototype)
				inclusionFilterPrototypeModel.Operand = core.StringPtr("location")
				inclusionFilterPrototypeModel.Operator = core.StringPtr("is")
				inclusionFilterPrototypeModel.Values = []string{"us-south"}

				// Construct an instance of the RulePrototype model
				rulePrototypeModel := new(metricsrouterv3.RulePrototype)
				rulePrototypeModel.Action = core.StringPtr("send")
				rulePrototypeModel.Targets = []metricsrouterv3.TargetIdentity{*targetIdentityModel}
				rulePrototypeModel.InclusionFilters = []metricsrouterv3.InclusionFilterPrototype{*inclusionFilterPrototypeModel}

				// Construct an instance of the UpdateRouteOptions model
				updateRouteOptionsModel := new(metricsrouterv3.UpdateRouteOptions)
				updateRouteOptionsModel.ID = core.StringPtr("testString")
				updateRouteOptionsModel.Name = core.StringPtr("my-route")
				updateRouteOptionsModel.Rules = []metricsrouterv3.RulePrototype{*rulePrototypeModel}
				updateRouteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := metricsRouterService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := metricsRouterService.UpdateRoute(updateRouteOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateRouteOptions model with no property values
				updateRouteOptionsModelNew := new(metricsrouterv3.UpdateRouteOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = metricsRouterService.UpdateRoute(updateRouteOptionsModelNew)
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
			It(`Invoke UpdateRoute successfully`, func() {
				metricsRouterService, serviceErr := metricsrouterv3.NewMetricsRouterV3(&metricsrouterv3.MetricsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(metricsRouterService).ToNot(BeNil())

				// Construct an instance of the TargetIdentity model
				targetIdentityModel := new(metricsrouterv3.TargetIdentity)
				targetIdentityModel.ID = core.StringPtr("c3af557f-fb0e-4476-85c3-0889e7fe7bc4")

				// Construct an instance of the InclusionFilterPrototype model
				inclusionFilterPrototypeModel := new(metricsrouterv3.InclusionFilterPrototype)
				inclusionFilterPrototypeModel.Operand = core.StringPtr("location")
				inclusionFilterPrototypeModel.Operator = core.StringPtr("is")
				inclusionFilterPrototypeModel.Values = []string{"us-south"}

				// Construct an instance of the RulePrototype model
				rulePrototypeModel := new(metricsrouterv3.RulePrototype)
				rulePrototypeModel.Action = core.StringPtr("send")
				rulePrototypeModel.Targets = []metricsrouterv3.TargetIdentity{*targetIdentityModel}
				rulePrototypeModel.InclusionFilters = []metricsrouterv3.InclusionFilterPrototype{*inclusionFilterPrototypeModel}

				// Construct an instance of the UpdateRouteOptions model
				updateRouteOptionsModel := new(metricsrouterv3.UpdateRouteOptions)
				updateRouteOptionsModel.ID = core.StringPtr("testString")
				updateRouteOptionsModel.Name = core.StringPtr("my-route")
				updateRouteOptionsModel.Rules = []metricsrouterv3.RulePrototype{*rulePrototypeModel}
				updateRouteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := metricsRouterService.UpdateRoute(updateRouteOptionsModel)
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
	Describe(`DeleteRoute(deleteRouteOptions *DeleteRouteOptions)`, func() {
		deleteRoutePath := "/routes/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteRoutePath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteRoute successfully`, func() {
				metricsRouterService, serviceErr := metricsrouterv3.NewMetricsRouterV3(&metricsrouterv3.MetricsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(metricsRouterService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := metricsRouterService.DeleteRoute(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteRouteOptions model
				deleteRouteOptionsModel := new(metricsrouterv3.DeleteRouteOptions)
				deleteRouteOptionsModel.ID = core.StringPtr("testString")
				deleteRouteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = metricsRouterService.DeleteRoute(deleteRouteOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteRoute with error: Operation validation and request error`, func() {
				metricsRouterService, serviceErr := metricsrouterv3.NewMetricsRouterV3(&metricsrouterv3.MetricsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(metricsRouterService).ToNot(BeNil())

				// Construct an instance of the DeleteRouteOptions model
				deleteRouteOptionsModel := new(metricsrouterv3.DeleteRouteOptions)
				deleteRouteOptionsModel.ID = core.StringPtr("testString")
				deleteRouteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := metricsRouterService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := metricsRouterService.DeleteRoute(deleteRouteOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteRouteOptions model with no property values
				deleteRouteOptionsModelNew := new(metricsrouterv3.DeleteRouteOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = metricsRouterService.DeleteRoute(deleteRouteOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetSettings(getSettingsOptions *GetSettingsOptions) - Operation response error`, func() {
		getSettingsPath := "/settings"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSettingsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetSettings with error: Operation response processing error`, func() {
				metricsRouterService, serviceErr := metricsrouterv3.NewMetricsRouterV3(&metricsrouterv3.MetricsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(metricsRouterService).ToNot(BeNil())

				// Construct an instance of the GetSettingsOptions model
				getSettingsOptionsModel := new(metricsrouterv3.GetSettingsOptions)
				getSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := metricsRouterService.GetSettings(getSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				metricsRouterService.EnableRetries(0, 0)
				result, response, operationErr = metricsRouterService.GetSettings(getSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetSettings(getSettingsOptions *GetSettingsOptions)`, func() {
		getSettingsPath := "/settings"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSettingsPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"default_targets": [{"id": "c3af557f-fb0e-4476-85c3-0889e7fe7bc4", "crn": "crn:v1:bluemix:public:sysdig-monitor:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::", "name": "a-mr-target-us-south", "target_type": "sysdig_monitor"}], "permitted_target_regions": ["us-south"], "primary_metadata_region": "us-south", "backup_metadata_region": "us-east", "private_api_endpoint_only": false}`)
				}))
			})
			It(`Invoke GetSettings successfully with retries`, func() {
				metricsRouterService, serviceErr := metricsrouterv3.NewMetricsRouterV3(&metricsrouterv3.MetricsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(metricsRouterService).ToNot(BeNil())
				metricsRouterService.EnableRetries(0, 0)

				// Construct an instance of the GetSettingsOptions model
				getSettingsOptionsModel := new(metricsrouterv3.GetSettingsOptions)
				getSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := metricsRouterService.GetSettingsWithContext(ctx, getSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				metricsRouterService.DisableRetries()
				result, response, operationErr := metricsRouterService.GetSettings(getSettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = metricsRouterService.GetSettingsWithContext(ctx, getSettingsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getSettingsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"default_targets": [{"id": "c3af557f-fb0e-4476-85c3-0889e7fe7bc4", "crn": "crn:v1:bluemix:public:sysdig-monitor:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::", "name": "a-mr-target-us-south", "target_type": "sysdig_monitor"}], "permitted_target_regions": ["us-south"], "primary_metadata_region": "us-south", "backup_metadata_region": "us-east", "private_api_endpoint_only": false}`)
				}))
			})
			It(`Invoke GetSettings successfully`, func() {
				metricsRouterService, serviceErr := metricsrouterv3.NewMetricsRouterV3(&metricsrouterv3.MetricsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(metricsRouterService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := metricsRouterService.GetSettings(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetSettingsOptions model
				getSettingsOptionsModel := new(metricsrouterv3.GetSettingsOptions)
				getSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = metricsRouterService.GetSettings(getSettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetSettings with error: Operation request error`, func() {
				metricsRouterService, serviceErr := metricsrouterv3.NewMetricsRouterV3(&metricsrouterv3.MetricsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(metricsRouterService).ToNot(BeNil())

				// Construct an instance of the GetSettingsOptions model
				getSettingsOptionsModel := new(metricsrouterv3.GetSettingsOptions)
				getSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := metricsRouterService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := metricsRouterService.GetSettings(getSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
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
			It(`Invoke GetSettings successfully`, func() {
				metricsRouterService, serviceErr := metricsrouterv3.NewMetricsRouterV3(&metricsrouterv3.MetricsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(metricsRouterService).ToNot(BeNil())

				// Construct an instance of the GetSettingsOptions model
				getSettingsOptionsModel := new(metricsrouterv3.GetSettingsOptions)
				getSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := metricsRouterService.GetSettings(getSettingsOptionsModel)
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
	Describe(`UpdateSettings(updateSettingsOptions *UpdateSettingsOptions) - Operation response error`, func() {
		updateSettingsPath := "/settings"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateSettingsPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateSettings with error: Operation response processing error`, func() {
				metricsRouterService, serviceErr := metricsrouterv3.NewMetricsRouterV3(&metricsrouterv3.MetricsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(metricsRouterService).ToNot(BeNil())

				// Construct an instance of the TargetIdentity model
				targetIdentityModel := new(metricsrouterv3.TargetIdentity)
				targetIdentityModel.ID = core.StringPtr("c3af557f-fb0e-4476-85c3-0889e7fe7bc4")

				// Construct an instance of the UpdateSettingsOptions model
				updateSettingsOptionsModel := new(metricsrouterv3.UpdateSettingsOptions)
				updateSettingsOptionsModel.DefaultTargets = []metricsrouterv3.TargetIdentity{*targetIdentityModel}
				updateSettingsOptionsModel.PermittedTargetRegions = []string{"us-south"}
				updateSettingsOptionsModel.PrimaryMetadataRegion = core.StringPtr("us-south")
				updateSettingsOptionsModel.BackupMetadataRegion = core.StringPtr("us-east")
				updateSettingsOptionsModel.PrivateAPIEndpointOnly = core.BoolPtr(false)
				updateSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := metricsRouterService.UpdateSettings(updateSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				metricsRouterService.EnableRetries(0, 0)
				result, response, operationErr = metricsRouterService.UpdateSettings(updateSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateSettings(updateSettingsOptions *UpdateSettingsOptions)`, func() {
		updateSettingsPath := "/settings"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateSettingsPath))
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

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"default_targets": [{"id": "c3af557f-fb0e-4476-85c3-0889e7fe7bc4", "crn": "crn:v1:bluemix:public:sysdig-monitor:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::", "name": "a-mr-target-us-south", "target_type": "sysdig_monitor"}], "permitted_target_regions": ["us-south"], "primary_metadata_region": "us-south", "backup_metadata_region": "us-east", "private_api_endpoint_only": false}`)
				}))
			})
			It(`Invoke UpdateSettings successfully with retries`, func() {
				metricsRouterService, serviceErr := metricsrouterv3.NewMetricsRouterV3(&metricsrouterv3.MetricsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(metricsRouterService).ToNot(BeNil())
				metricsRouterService.EnableRetries(0, 0)

				// Construct an instance of the TargetIdentity model
				targetIdentityModel := new(metricsrouterv3.TargetIdentity)
				targetIdentityModel.ID = core.StringPtr("c3af557f-fb0e-4476-85c3-0889e7fe7bc4")

				// Construct an instance of the UpdateSettingsOptions model
				updateSettingsOptionsModel := new(metricsrouterv3.UpdateSettingsOptions)
				updateSettingsOptionsModel.DefaultTargets = []metricsrouterv3.TargetIdentity{*targetIdentityModel}
				updateSettingsOptionsModel.PermittedTargetRegions = []string{"us-south"}
				updateSettingsOptionsModel.PrimaryMetadataRegion = core.StringPtr("us-south")
				updateSettingsOptionsModel.BackupMetadataRegion = core.StringPtr("us-east")
				updateSettingsOptionsModel.PrivateAPIEndpointOnly = core.BoolPtr(false)
				updateSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := metricsRouterService.UpdateSettingsWithContext(ctx, updateSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				metricsRouterService.DisableRetries()
				result, response, operationErr := metricsRouterService.UpdateSettings(updateSettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = metricsRouterService.UpdateSettingsWithContext(ctx, updateSettingsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateSettingsPath))
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

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"default_targets": [{"id": "c3af557f-fb0e-4476-85c3-0889e7fe7bc4", "crn": "crn:v1:bluemix:public:sysdig-monitor:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::", "name": "a-mr-target-us-south", "target_type": "sysdig_monitor"}], "permitted_target_regions": ["us-south"], "primary_metadata_region": "us-south", "backup_metadata_region": "us-east", "private_api_endpoint_only": false}`)
				}))
			})
			It(`Invoke UpdateSettings successfully`, func() {
				metricsRouterService, serviceErr := metricsrouterv3.NewMetricsRouterV3(&metricsrouterv3.MetricsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(metricsRouterService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := metricsRouterService.UpdateSettings(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the TargetIdentity model
				targetIdentityModel := new(metricsrouterv3.TargetIdentity)
				targetIdentityModel.ID = core.StringPtr("c3af557f-fb0e-4476-85c3-0889e7fe7bc4")

				// Construct an instance of the UpdateSettingsOptions model
				updateSettingsOptionsModel := new(metricsrouterv3.UpdateSettingsOptions)
				updateSettingsOptionsModel.DefaultTargets = []metricsrouterv3.TargetIdentity{*targetIdentityModel}
				updateSettingsOptionsModel.PermittedTargetRegions = []string{"us-south"}
				updateSettingsOptionsModel.PrimaryMetadataRegion = core.StringPtr("us-south")
				updateSettingsOptionsModel.BackupMetadataRegion = core.StringPtr("us-east")
				updateSettingsOptionsModel.PrivateAPIEndpointOnly = core.BoolPtr(false)
				updateSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = metricsRouterService.UpdateSettings(updateSettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateSettings with error: Operation request error`, func() {
				metricsRouterService, serviceErr := metricsrouterv3.NewMetricsRouterV3(&metricsrouterv3.MetricsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(metricsRouterService).ToNot(BeNil())

				// Construct an instance of the TargetIdentity model
				targetIdentityModel := new(metricsrouterv3.TargetIdentity)
				targetIdentityModel.ID = core.StringPtr("c3af557f-fb0e-4476-85c3-0889e7fe7bc4")

				// Construct an instance of the UpdateSettingsOptions model
				updateSettingsOptionsModel := new(metricsrouterv3.UpdateSettingsOptions)
				updateSettingsOptionsModel.DefaultTargets = []metricsrouterv3.TargetIdentity{*targetIdentityModel}
				updateSettingsOptionsModel.PermittedTargetRegions = []string{"us-south"}
				updateSettingsOptionsModel.PrimaryMetadataRegion = core.StringPtr("us-south")
				updateSettingsOptionsModel.BackupMetadataRegion = core.StringPtr("us-east")
				updateSettingsOptionsModel.PrivateAPIEndpointOnly = core.BoolPtr(false)
				updateSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := metricsRouterService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := metricsRouterService.UpdateSettings(updateSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
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
			It(`Invoke UpdateSettings successfully`, func() {
				metricsRouterService, serviceErr := metricsrouterv3.NewMetricsRouterV3(&metricsrouterv3.MetricsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(metricsRouterService).ToNot(BeNil())

				// Construct an instance of the TargetIdentity model
				targetIdentityModel := new(metricsrouterv3.TargetIdentity)
				targetIdentityModel.ID = core.StringPtr("c3af557f-fb0e-4476-85c3-0889e7fe7bc4")

				// Construct an instance of the UpdateSettingsOptions model
				updateSettingsOptionsModel := new(metricsrouterv3.UpdateSettingsOptions)
				updateSettingsOptionsModel.DefaultTargets = []metricsrouterv3.TargetIdentity{*targetIdentityModel}
				updateSettingsOptionsModel.PermittedTargetRegions = []string{"us-south"}
				updateSettingsOptionsModel.PrimaryMetadataRegion = core.StringPtr("us-south")
				updateSettingsOptionsModel.BackupMetadataRegion = core.StringPtr("us-east")
				updateSettingsOptionsModel.PrivateAPIEndpointOnly = core.BoolPtr(false)
				updateSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := metricsRouterService.UpdateSettings(updateSettingsOptionsModel)
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
			metricsRouterService, _ := metricsrouterv3.NewMetricsRouterV3(&metricsrouterv3.MetricsRouterV3Options{
				URL:           "http://metricsrouterv3modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewCreateRouteOptions successfully`, func() {
				// Construct an instance of the TargetIdentity model
				targetIdentityModel := new(metricsrouterv3.TargetIdentity)
				Expect(targetIdentityModel).ToNot(BeNil())
				targetIdentityModel.ID = core.StringPtr("c3af557f-fb0e-4476-85c3-0889e7fe7bc4")
				Expect(targetIdentityModel.ID).To(Equal(core.StringPtr("c3af557f-fb0e-4476-85c3-0889e7fe7bc4")))

				// Construct an instance of the InclusionFilterPrototype model
				inclusionFilterPrototypeModel := new(metricsrouterv3.InclusionFilterPrototype)
				Expect(inclusionFilterPrototypeModel).ToNot(BeNil())
				inclusionFilterPrototypeModel.Operand = core.StringPtr("location")
				inclusionFilterPrototypeModel.Operator = core.StringPtr("is")
				inclusionFilterPrototypeModel.Values = []string{"us-south"}
				Expect(inclusionFilterPrototypeModel.Operand).To(Equal(core.StringPtr("location")))
				Expect(inclusionFilterPrototypeModel.Operator).To(Equal(core.StringPtr("is")))
				Expect(inclusionFilterPrototypeModel.Values).To(Equal([]string{"us-south"}))

				// Construct an instance of the RulePrototype model
				rulePrototypeModel := new(metricsrouterv3.RulePrototype)
				Expect(rulePrototypeModel).ToNot(BeNil())
				rulePrototypeModel.Action = core.StringPtr("send")
				rulePrototypeModel.Targets = []metricsrouterv3.TargetIdentity{*targetIdentityModel}
				rulePrototypeModel.InclusionFilters = []metricsrouterv3.InclusionFilterPrototype{*inclusionFilterPrototypeModel}
				Expect(rulePrototypeModel.Action).To(Equal(core.StringPtr("send")))
				Expect(rulePrototypeModel.Targets).To(Equal([]metricsrouterv3.TargetIdentity{*targetIdentityModel}))
				Expect(rulePrototypeModel.InclusionFilters).To(Equal([]metricsrouterv3.InclusionFilterPrototype{*inclusionFilterPrototypeModel}))

				// Construct an instance of the CreateRouteOptions model
				createRouteOptionsName := "my-route"
				createRouteOptionsRules := []metricsrouterv3.RulePrototype{}
				createRouteOptionsModel := metricsRouterService.NewCreateRouteOptions(createRouteOptionsName, createRouteOptionsRules)
				createRouteOptionsModel.SetName("my-route")
				createRouteOptionsModel.SetRules([]metricsrouterv3.RulePrototype{*rulePrototypeModel})
				createRouteOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createRouteOptionsModel).ToNot(BeNil())
				Expect(createRouteOptionsModel.Name).To(Equal(core.StringPtr("my-route")))
				Expect(createRouteOptionsModel.Rules).To(Equal([]metricsrouterv3.RulePrototype{*rulePrototypeModel}))
				Expect(createRouteOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateTargetOptions successfully`, func() {
				// Construct an instance of the CreateTargetOptions model
				createTargetOptionsName := "my-mr-target"
				createTargetOptionsDestinationCRN := "crn:v1:bluemix:public:sysdig-monitor:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::"
				createTargetOptionsModel := metricsRouterService.NewCreateTargetOptions(createTargetOptionsName, createTargetOptionsDestinationCRN)
				createTargetOptionsModel.SetName("my-mr-target")
				createTargetOptionsModel.SetDestinationCRN("crn:v1:bluemix:public:sysdig-monitor:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::")
				createTargetOptionsModel.SetRegion("us-south")
				createTargetOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createTargetOptionsModel).ToNot(BeNil())
				Expect(createTargetOptionsModel.Name).To(Equal(core.StringPtr("my-mr-target")))
				Expect(createTargetOptionsModel.DestinationCRN).To(Equal(core.StringPtr("crn:v1:bluemix:public:sysdig-monitor:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::")))
				Expect(createTargetOptionsModel.Region).To(Equal(core.StringPtr("us-south")))
				Expect(createTargetOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteRouteOptions successfully`, func() {
				// Construct an instance of the DeleteRouteOptions model
				id := "testString"
				deleteRouteOptionsModel := metricsRouterService.NewDeleteRouteOptions(id)
				deleteRouteOptionsModel.SetID("testString")
				deleteRouteOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteRouteOptionsModel).ToNot(BeNil())
				Expect(deleteRouteOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(deleteRouteOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteTargetOptions successfully`, func() {
				// Construct an instance of the DeleteTargetOptions model
				id := "testString"
				deleteTargetOptionsModel := metricsRouterService.NewDeleteTargetOptions(id)
				deleteTargetOptionsModel.SetID("testString")
				deleteTargetOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteTargetOptionsModel).ToNot(BeNil())
				Expect(deleteTargetOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(deleteTargetOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetRouteOptions successfully`, func() {
				// Construct an instance of the GetRouteOptions model
				id := "testString"
				getRouteOptionsModel := metricsRouterService.NewGetRouteOptions(id)
				getRouteOptionsModel.SetID("testString")
				getRouteOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getRouteOptionsModel).ToNot(BeNil())
				Expect(getRouteOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getRouteOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetSettingsOptions successfully`, func() {
				// Construct an instance of the GetSettingsOptions model
				getSettingsOptionsModel := metricsRouterService.NewGetSettingsOptions()
				getSettingsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getSettingsOptionsModel).ToNot(BeNil())
				Expect(getSettingsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetTargetOptions successfully`, func() {
				// Construct an instance of the GetTargetOptions model
				id := "testString"
				getTargetOptionsModel := metricsRouterService.NewGetTargetOptions(id)
				getTargetOptionsModel.SetID("testString")
				getTargetOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getTargetOptionsModel).ToNot(BeNil())
				Expect(getTargetOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getTargetOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewInclusionFilterPrototype successfully`, func() {
				operand := "location"
				operator := "is"
				values := []string{"us-south"}
				_model, err := metricsRouterService.NewInclusionFilterPrototype(operand, operator, values)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewListRoutesOptions successfully`, func() {
				// Construct an instance of the ListRoutesOptions model
				listRoutesOptionsModel := metricsRouterService.NewListRoutesOptions()
				listRoutesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listRoutesOptionsModel).ToNot(BeNil())
				Expect(listRoutesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListTargetsOptions successfully`, func() {
				// Construct an instance of the ListTargetsOptions model
				listTargetsOptionsModel := metricsRouterService.NewListTargetsOptions()
				listTargetsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listTargetsOptionsModel).ToNot(BeNil())
				Expect(listTargetsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewRulePrototype successfully`, func() {
				targets := []metricsrouterv3.TargetIdentity{}
				inclusionFilters := []metricsrouterv3.InclusionFilterPrototype{}
				_model, err := metricsRouterService.NewRulePrototype(targets, inclusionFilters)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewTargetIdentity successfully`, func() {
				id := "c3af557f-fb0e-4476-85c3-0889e7fe7bc4"
				_model, err := metricsRouterService.NewTargetIdentity(id)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewUpdateRouteOptions successfully`, func() {
				// Construct an instance of the TargetIdentity model
				targetIdentityModel := new(metricsrouterv3.TargetIdentity)
				Expect(targetIdentityModel).ToNot(BeNil())
				targetIdentityModel.ID = core.StringPtr("c3af557f-fb0e-4476-85c3-0889e7fe7bc4")
				Expect(targetIdentityModel.ID).To(Equal(core.StringPtr("c3af557f-fb0e-4476-85c3-0889e7fe7bc4")))

				// Construct an instance of the InclusionFilterPrototype model
				inclusionFilterPrototypeModel := new(metricsrouterv3.InclusionFilterPrototype)
				Expect(inclusionFilterPrototypeModel).ToNot(BeNil())
				inclusionFilterPrototypeModel.Operand = core.StringPtr("location")
				inclusionFilterPrototypeModel.Operator = core.StringPtr("is")
				inclusionFilterPrototypeModel.Values = []string{"us-south"}
				Expect(inclusionFilterPrototypeModel.Operand).To(Equal(core.StringPtr("location")))
				Expect(inclusionFilterPrototypeModel.Operator).To(Equal(core.StringPtr("is")))
				Expect(inclusionFilterPrototypeModel.Values).To(Equal([]string{"us-south"}))

				// Construct an instance of the RulePrototype model
				rulePrototypeModel := new(metricsrouterv3.RulePrototype)
				Expect(rulePrototypeModel).ToNot(BeNil())
				rulePrototypeModel.Action = core.StringPtr("send")
				rulePrototypeModel.Targets = []metricsrouterv3.TargetIdentity{*targetIdentityModel}
				rulePrototypeModel.InclusionFilters = []metricsrouterv3.InclusionFilterPrototype{*inclusionFilterPrototypeModel}
				Expect(rulePrototypeModel.Action).To(Equal(core.StringPtr("send")))
				Expect(rulePrototypeModel.Targets).To(Equal([]metricsrouterv3.TargetIdentity{*targetIdentityModel}))
				Expect(rulePrototypeModel.InclusionFilters).To(Equal([]metricsrouterv3.InclusionFilterPrototype{*inclusionFilterPrototypeModel}))

				// Construct an instance of the UpdateRouteOptions model
				id := "testString"
				updateRouteOptionsModel := metricsRouterService.NewUpdateRouteOptions(id)
				updateRouteOptionsModel.SetID("testString")
				updateRouteOptionsModel.SetName("my-route")
				updateRouteOptionsModel.SetRules([]metricsrouterv3.RulePrototype{*rulePrototypeModel})
				updateRouteOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateRouteOptionsModel).ToNot(BeNil())
				Expect(updateRouteOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(updateRouteOptionsModel.Name).To(Equal(core.StringPtr("my-route")))
				Expect(updateRouteOptionsModel.Rules).To(Equal([]metricsrouterv3.RulePrototype{*rulePrototypeModel}))
				Expect(updateRouteOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateSettingsOptions successfully`, func() {
				// Construct an instance of the TargetIdentity model
				targetIdentityModel := new(metricsrouterv3.TargetIdentity)
				Expect(targetIdentityModel).ToNot(BeNil())
				targetIdentityModel.ID = core.StringPtr("c3af557f-fb0e-4476-85c3-0889e7fe7bc4")
				Expect(targetIdentityModel.ID).To(Equal(core.StringPtr("c3af557f-fb0e-4476-85c3-0889e7fe7bc4")))

				// Construct an instance of the UpdateSettingsOptions model
				updateSettingsOptionsModel := metricsRouterService.NewUpdateSettingsOptions()
				updateSettingsOptionsModel.SetDefaultTargets([]metricsrouterv3.TargetIdentity{*targetIdentityModel})
				updateSettingsOptionsModel.SetPermittedTargetRegions([]string{"us-south"})
				updateSettingsOptionsModel.SetPrimaryMetadataRegion("us-south")
				updateSettingsOptionsModel.SetBackupMetadataRegion("us-east")
				updateSettingsOptionsModel.SetPrivateAPIEndpointOnly(false)
				updateSettingsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateSettingsOptionsModel).ToNot(BeNil())
				Expect(updateSettingsOptionsModel.DefaultTargets).To(Equal([]metricsrouterv3.TargetIdentity{*targetIdentityModel}))
				Expect(updateSettingsOptionsModel.PermittedTargetRegions).To(Equal([]string{"us-south"}))
				Expect(updateSettingsOptionsModel.PrimaryMetadataRegion).To(Equal(core.StringPtr("us-south")))
				Expect(updateSettingsOptionsModel.BackupMetadataRegion).To(Equal(core.StringPtr("us-east")))
				Expect(updateSettingsOptionsModel.PrivateAPIEndpointOnly).To(Equal(core.BoolPtr(false)))
				Expect(updateSettingsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateTargetOptions successfully`, func() {
				// Construct an instance of the UpdateTargetOptions model
				id := "testString"
				updateTargetOptionsModel := metricsRouterService.NewUpdateTargetOptions(id)
				updateTargetOptionsModel.SetID("testString")
				updateTargetOptionsModel.SetName("my-mr-target")
				updateTargetOptionsModel.SetDestinationCRN("crn:v1:bluemix:public:sysdig-monitor:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::")
				updateTargetOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateTargetOptionsModel).ToNot(BeNil())
				Expect(updateTargetOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(updateTargetOptionsModel.Name).To(Equal(core.StringPtr("my-mr-target")))
				Expect(updateTargetOptionsModel.DestinationCRN).To(Equal(core.StringPtr("crn:v1:bluemix:public:sysdig-monitor:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::")))
				Expect(updateTargetOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
		})
	})
	Describe(`Model unmarshaling tests`, func() {
		It(`Invoke UnmarshalInclusionFilterPrototype successfully`, func() {
			// Construct an instance of the model.
			model := new(metricsrouterv3.InclusionFilterPrototype)
			model.Operand = core.StringPtr("location")
			model.Operator = core.StringPtr("is")
			model.Values = []string{"us-south"}

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *metricsrouterv3.InclusionFilterPrototype
			err = metricsrouterv3.UnmarshalInclusionFilterPrototype(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalRulePrototype successfully`, func() {
			// Construct an instance of the model.
			model := new(metricsrouterv3.RulePrototype)
			model.Action = core.StringPtr("send")
			model.Targets = nil
			model.InclusionFilters = nil

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *metricsrouterv3.RulePrototype
			err = metricsrouterv3.UnmarshalRulePrototype(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalTargetIdentity successfully`, func() {
			// Construct an instance of the model.
			model := new(metricsrouterv3.TargetIdentity)
			model.ID = core.StringPtr("c3af557f-fb0e-4476-85c3-0889e7fe7bc4")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *metricsrouterv3.TargetIdentity
			err = metricsrouterv3.UnmarshalTargetIdentity(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
	})
	Describe(`Utility function tests`, func() {
		It(`Invoke CreateMockByteArray() successfully`, func() {
			mockByteArray := CreateMockByteArray("VGhpcyBpcyBhIHRlc3Qgb2YgdGhlIGVtZXJnZW5jeSBicm9hZGNhc3Qgc3lzdGVt")
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
