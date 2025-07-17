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

package atrackerv2_test

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
	"github.com/IBM/platform-services-go-sdk/atrackerv2"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`AtrackerV2`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			atrackerService, serviceErr := atrackerv2.NewAtrackerV2(&atrackerv2.AtrackerV2Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(atrackerService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			atrackerService, serviceErr := atrackerv2.NewAtrackerV2(&atrackerv2.AtrackerV2Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(atrackerService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			atrackerService, serviceErr := atrackerv2.NewAtrackerV2(&atrackerv2.AtrackerV2Options{
				URL: "https://atrackerv2/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(atrackerService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"ATRACKER_URL":       "https://atrackerv2/api",
				"ATRACKER_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				atrackerService, serviceErr := atrackerv2.NewAtrackerV2UsingExternalConfig(&atrackerv2.AtrackerV2Options{})
				Expect(atrackerService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := atrackerService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != atrackerService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(atrackerService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(atrackerService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				atrackerService, serviceErr := atrackerv2.NewAtrackerV2UsingExternalConfig(&atrackerv2.AtrackerV2Options{
					URL: "https://testService/api",
				})
				Expect(atrackerService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := atrackerService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != atrackerService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(atrackerService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(atrackerService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				atrackerService, serviceErr := atrackerv2.NewAtrackerV2UsingExternalConfig(&atrackerv2.AtrackerV2Options{})
				err := atrackerService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := atrackerService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != atrackerService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(atrackerService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(atrackerService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"ATRACKER_URL":       "https://atrackerv2/api",
				"ATRACKER_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			atrackerService, serviceErr := atrackerv2.NewAtrackerV2UsingExternalConfig(&atrackerv2.AtrackerV2Options{})

			It(`Instantiate service client with error`, func() {
				Expect(atrackerService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"ATRACKER_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			atrackerService, serviceErr := atrackerv2.NewAtrackerV2UsingExternalConfig(&atrackerv2.AtrackerV2Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(atrackerService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = atrackerv2.GetServiceURLForRegion("us-south")
			Expect(url).To(Equal("https://us-south.atracker.cloud.ibm.com"))
			Expect(err).To(BeNil())

			url, err = atrackerv2.GetServiceURLForRegion("private.us-south")
			Expect(url).To(Equal("https://private.us-south.atracker.cloud.ibm.com"))
			Expect(err).To(BeNil())

			url, err = atrackerv2.GetServiceURLForRegion("us-east")
			Expect(url).To(Equal("https://us-east.atracker.cloud.ibm.com"))
			Expect(err).To(BeNil())

			url, err = atrackerv2.GetServiceURLForRegion("private.us-east")
			Expect(url).To(Equal("https://private.us-east.atracker.cloud.ibm.com"))
			Expect(err).To(BeNil())

			url, err = atrackerv2.GetServiceURLForRegion("eu-de")
			Expect(url).To(Equal("https://eu-de.atracker.cloud.ibm.com"))
			Expect(err).To(BeNil())

			url, err = atrackerv2.GetServiceURLForRegion("private.eu-de")
			Expect(url).To(Equal("https://private.eu-de.atracker.cloud.ibm.com"))
			Expect(err).To(BeNil())

			url, err = atrackerv2.GetServiceURLForRegion("eu-gb")
			Expect(url).To(Equal("https://eu-gb.atracker.cloud.ibm.com"))
			Expect(err).To(BeNil())

			url, err = atrackerv2.GetServiceURLForRegion("private.eu-gb")
			Expect(url).To(Equal("https://private.eu-gb.atracker.cloud.ibm.com"))
			Expect(err).To(BeNil())

			url, err = atrackerv2.GetServiceURLForRegion("eu-es")
			Expect(url).To(Equal("https://eu-es.atracker.cloud.ibm.com"))
			Expect(err).To(BeNil())

			url, err = atrackerv2.GetServiceURLForRegion("private.eu-es")
			Expect(url).To(Equal("https://private.eu-es.atracker.cloud.ibm.com"))
			Expect(err).To(BeNil())

			url, err = atrackerv2.GetServiceURLForRegion("au-syd")
			Expect(url).To(Equal("https://au-syd.atracker.cloud.ibm.com"))
			Expect(err).To(BeNil())

			url, err = atrackerv2.GetServiceURLForRegion("private.au-syd")
			Expect(url).To(Equal("https://private.au-syd.atracker.cloud.ibm.com"))
			Expect(err).To(BeNil())

			url, err = atrackerv2.GetServiceURLForRegion("ca-tor")
			Expect(url).To(Equal("https://us-east.atracker.cloud.ibm.com"))
			Expect(err).To(BeNil())

			url, err = atrackerv2.GetServiceURLForRegion("private.ca-tor")
			Expect(url).To(Equal("https://private.us-east.atracker.cloud.ibm.com"))
			Expect(err).To(BeNil())

			url, err = atrackerv2.GetServiceURLForRegion("br-sao")
			Expect(url).To(Equal("https://us-south.atracker.cloud.ibm.com"))
			Expect(err).To(BeNil())

			url, err = atrackerv2.GetServiceURLForRegion("private.br-sao")
			Expect(url).To(Equal("https://private.us-south.atracker.cloud.ibm.com"))
			Expect(err).To(BeNil())

			url, err = atrackerv2.GetServiceURLForRegion("eu-fr2")
			Expect(url).To(Equal("https://eu-de.atracker.cloud.ibm.com"))
			Expect(err).To(BeNil())

			url, err = atrackerv2.GetServiceURLForRegion("private.eu-fr2")
			Expect(url).To(Equal("https://private.eu-de.atracker.cloud.ibm.com"))
			Expect(err).To(BeNil())

			url, err = atrackerv2.GetServiceURLForRegion("jp-tok")
			Expect(url).To(Equal("https://eu-de.atracker.cloud.ibm.com"))
			Expect(err).To(BeNil())

			url, err = atrackerv2.GetServiceURLForRegion("private.jp-tok")
			Expect(url).To(Equal("https://private.eu-de.atracker.cloud.ibm.com"))
			Expect(err).To(BeNil())

			url, err = atrackerv2.GetServiceURLForRegion("jp-osa")
			Expect(url).To(Equal("https://eu-de.atracker.cloud.ibm.com"))
			Expect(err).To(BeNil())

			url, err = atrackerv2.GetServiceURLForRegion("private.jp-osa")
			Expect(url).To(Equal("https://private.eu-de.atracker.cloud.ibm.com"))
			Expect(err).To(BeNil())

			url, err = atrackerv2.GetServiceURLForRegion("in-che")
			Expect(url).To(Equal("https://eu-de.atracker.cloud.ibm.com"))
			Expect(err).To(BeNil())

			url, err = atrackerv2.GetServiceURLForRegion("private.in-che")
			Expect(url).To(Equal("https://private.eu-de.atracker.cloud.ibm.com"))
			Expect(err).To(BeNil())

			url, err = atrackerv2.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`CreateTarget(createTargetOptions *CreateTargetOptions) - Operation response error`, func() {
		createTargetPath := "/api/v2/targets"
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
				atrackerService, serviceErr := atrackerv2.NewAtrackerV2(&atrackerv2.AtrackerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())

				// Construct an instance of the CosEndpointPrototype model
				cosEndpointPrototypeModel := new(atrackerv2.CosEndpointPrototype)
				cosEndpointPrototypeModel.Endpoint = core.StringPtr("s3.private.us-east.cloud-object-storage.appdomain.cloud")
				cosEndpointPrototypeModel.TargetCRN = core.StringPtr("crn:v1:bluemix:public:cloud-object-storage:global:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::")
				cosEndpointPrototypeModel.Bucket = core.StringPtr("my-atracker-bucket")
				cosEndpointPrototypeModel.APIKey = core.StringPtr("xxxxxxxxxxxxxx")
				cosEndpointPrototypeModel.ServiceToServiceEnabled = core.BoolPtr(true)

				// Construct an instance of the EventstreamsEndpointPrototype model
				eventstreamsEndpointPrototypeModel := new(atrackerv2.EventstreamsEndpointPrototype)
				eventstreamsEndpointPrototypeModel.TargetCRN = core.StringPtr("crn:v1:bluemix:public:messagehub:us-south:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::")
				eventstreamsEndpointPrototypeModel.Brokers = []string{"kafka-x:9094"}
				eventstreamsEndpointPrototypeModel.Topic = core.StringPtr("my-topic")
				eventstreamsEndpointPrototypeModel.APIKey = core.StringPtr("xxxxxxxxxxxxxx")
				eventstreamsEndpointPrototypeModel.ServiceToServiceEnabled = core.BoolPtr(false)

				// Construct an instance of the CloudLogsEndpointPrototype model
				cloudLogsEndpointPrototypeModel := new(atrackerv2.CloudLogsEndpointPrototype)
				cloudLogsEndpointPrototypeModel.TargetCRN = core.StringPtr("crn:v1:bluemix:public:logs:eu-es:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::")

				// Construct an instance of the CreateTargetOptions model
				createTargetOptionsModel := new(atrackerv2.CreateTargetOptions)
				createTargetOptionsModel.Name = core.StringPtr("my-cos-target")
				createTargetOptionsModel.TargetType = core.StringPtr("cloud_object_storage")
				createTargetOptionsModel.CosEndpoint = cosEndpointPrototypeModel
				createTargetOptionsModel.EventstreamsEndpoint = eventstreamsEndpointPrototypeModel
				createTargetOptionsModel.CloudlogsEndpoint = cloudLogsEndpointPrototypeModel
				createTargetOptionsModel.Region = core.StringPtr("us-south")
				createTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := atrackerService.CreateTarget(createTargetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				atrackerService.EnableRetries(0, 0)
				result, response, operationErr = atrackerService.CreateTarget(createTargetOptionsModel)
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
		createTargetPath := "/api/v2/targets"
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
					fmt.Fprintf(res, "%s", `{"id": "f7dcfae6-e7c5-08ca-451b-fdfa696c9bb6", "name": "a-cos-target-us-south", "crn": "crn:v1:bluemix:public:atracker:us-south:a/11111111111111111111111111111111:b6eec08b-5201-08ca-451b-cd71523e3626:target:f7dcfae6-e7c5-08ca-451b-fdfa696c9bb6", "target_type": "cloud_object_storage", "region": "us-south", "cos_endpoint": {"endpoint": "s3.private.us-east.cloud-object-storage.appdomain.cloud", "target_crn": "crn:v1:bluemix:public:cloud-object-storage:global:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::", "bucket": "my-atracker-bucket", "service_to_service_enabled": true}, "eventstreams_endpoint": {"target_crn": "crn:v1:bluemix:public:messagehub:us-south:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::", "brokers": ["kafka-x:9094"], "topic": "my-topic", "api_key": "xxxxxxxxxxxxxx", "service_to_service_enabled": false}, "cloudlogs_endpoint": {"target_crn": "crn:v1:bluemix:public:eu-es:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::"}, "write_status": {"status": "success", "last_failure": "2021-05-18T20:15:12.353Z", "reason_for_last_failure": "Provided API key could not be found"}, "created_at": "2021-05-18T20:15:12.353Z", "updated_at": "2021-05-18T20:15:12.353Z", "message": "This is a valid target. However, there is another target already defined with the same target endpoint.", "api_version": 2}`)
				}))
			})
			It(`Invoke CreateTarget successfully with retries`, func() {
				atrackerService, serviceErr := atrackerv2.NewAtrackerV2(&atrackerv2.AtrackerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())
				atrackerService.EnableRetries(0, 0)

				// Construct an instance of the CosEndpointPrototype model
				cosEndpointPrototypeModel := new(atrackerv2.CosEndpointPrototype)
				cosEndpointPrototypeModel.Endpoint = core.StringPtr("s3.private.us-east.cloud-object-storage.appdomain.cloud")
				cosEndpointPrototypeModel.TargetCRN = core.StringPtr("crn:v1:bluemix:public:cloud-object-storage:global:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::")
				cosEndpointPrototypeModel.Bucket = core.StringPtr("my-atracker-bucket")
				cosEndpointPrototypeModel.APIKey = core.StringPtr("xxxxxxxxxxxxxx")
				cosEndpointPrototypeModel.ServiceToServiceEnabled = core.BoolPtr(true)

				// Construct an instance of the EventstreamsEndpointPrototype model
				eventstreamsEndpointPrototypeModel := new(atrackerv2.EventstreamsEndpointPrototype)
				eventstreamsEndpointPrototypeModel.TargetCRN = core.StringPtr("crn:v1:bluemix:public:messagehub:us-south:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::")
				eventstreamsEndpointPrototypeModel.Brokers = []string{"kafka-x:9094"}
				eventstreamsEndpointPrototypeModel.Topic = core.StringPtr("my-topic")
				eventstreamsEndpointPrototypeModel.APIKey = core.StringPtr("xxxxxxxxxxxxxx")
				eventstreamsEndpointPrototypeModel.ServiceToServiceEnabled = core.BoolPtr(false)

				// Construct an instance of the CloudLogsEndpointPrototype model
				cloudLogsEndpointPrototypeModel := new(atrackerv2.CloudLogsEndpointPrototype)
				cloudLogsEndpointPrototypeModel.TargetCRN = core.StringPtr("crn:v1:bluemix:public:logs:eu-es:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::")

				// Construct an instance of the CreateTargetOptions model
				createTargetOptionsModel := new(atrackerv2.CreateTargetOptions)
				createTargetOptionsModel.Name = core.StringPtr("my-cos-target")
				createTargetOptionsModel.TargetType = core.StringPtr("cloud_object_storage")
				createTargetOptionsModel.CosEndpoint = cosEndpointPrototypeModel
				createTargetOptionsModel.EventstreamsEndpoint = eventstreamsEndpointPrototypeModel
				createTargetOptionsModel.CloudlogsEndpoint = cloudLogsEndpointPrototypeModel
				createTargetOptionsModel.Region = core.StringPtr("us-south")
				createTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := atrackerService.CreateTargetWithContext(ctx, createTargetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				atrackerService.DisableRetries()
				result, response, operationErr := atrackerService.CreateTarget(createTargetOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = atrackerService.CreateTargetWithContext(ctx, createTargetOptionsModel)
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
					fmt.Fprintf(res, "%s", `{"id": "f7dcfae6-e7c5-08ca-451b-fdfa696c9bb6", "name": "a-cos-target-us-south", "crn": "crn:v1:bluemix:public:atracker:us-south:a/11111111111111111111111111111111:b6eec08b-5201-08ca-451b-cd71523e3626:target:f7dcfae6-e7c5-08ca-451b-fdfa696c9bb6", "target_type": "cloud_object_storage", "region": "us-south", "cos_endpoint": {"endpoint": "s3.private.us-east.cloud-object-storage.appdomain.cloud", "target_crn": "crn:v1:bluemix:public:cloud-object-storage:global:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::", "bucket": "my-atracker-bucket", "service_to_service_enabled": true}, "eventstreams_endpoint": {"target_crn": "crn:v1:bluemix:public:messagehub:us-south:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::", "brokers": ["kafka-x:9094"], "topic": "my-topic", "api_key": "xxxxxxxxxxxxxx", "service_to_service_enabled": false}, "cloudlogs_endpoint": {"target_crn": "crn:v1:bluemix:public:eu-es:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::"}, "write_status": {"status": "success", "last_failure": "2021-05-18T20:15:12.353Z", "reason_for_last_failure": "Provided API key could not be found"}, "created_at": "2021-05-18T20:15:12.353Z", "updated_at": "2021-05-18T20:15:12.353Z", "message": "This is a valid target. However, there is another target already defined with the same target endpoint.", "api_version": 2}`)
				}))
			})
			It(`Invoke CreateTarget successfully`, func() {
				atrackerService, serviceErr := atrackerv2.NewAtrackerV2(&atrackerv2.AtrackerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := atrackerService.CreateTarget(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CosEndpointPrototype model
				cosEndpointPrototypeModel := new(atrackerv2.CosEndpointPrototype)
				cosEndpointPrototypeModel.Endpoint = core.StringPtr("s3.private.us-east.cloud-object-storage.appdomain.cloud")
				cosEndpointPrototypeModel.TargetCRN = core.StringPtr("crn:v1:bluemix:public:cloud-object-storage:global:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::")
				cosEndpointPrototypeModel.Bucket = core.StringPtr("my-atracker-bucket")
				cosEndpointPrototypeModel.APIKey = core.StringPtr("xxxxxxxxxxxxxx")
				cosEndpointPrototypeModel.ServiceToServiceEnabled = core.BoolPtr(true)

				// Construct an instance of the EventstreamsEndpointPrototype model
				eventstreamsEndpointPrototypeModel := new(atrackerv2.EventstreamsEndpointPrototype)
				eventstreamsEndpointPrototypeModel.TargetCRN = core.StringPtr("crn:v1:bluemix:public:messagehub:us-south:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::")
				eventstreamsEndpointPrototypeModel.Brokers = []string{"kafka-x:9094"}
				eventstreamsEndpointPrototypeModel.Topic = core.StringPtr("my-topic")
				eventstreamsEndpointPrototypeModel.APIKey = core.StringPtr("xxxxxxxxxxxxxx")
				eventstreamsEndpointPrototypeModel.ServiceToServiceEnabled = core.BoolPtr(false)

				// Construct an instance of the CloudLogsEndpointPrototype model
				cloudLogsEndpointPrototypeModel := new(atrackerv2.CloudLogsEndpointPrototype)
				cloudLogsEndpointPrototypeModel.TargetCRN = core.StringPtr("crn:v1:bluemix:public:logs:eu-es:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::")

				// Construct an instance of the CreateTargetOptions model
				createTargetOptionsModel := new(atrackerv2.CreateTargetOptions)
				createTargetOptionsModel.Name = core.StringPtr("my-cos-target")
				createTargetOptionsModel.TargetType = core.StringPtr("cloud_object_storage")
				createTargetOptionsModel.CosEndpoint = cosEndpointPrototypeModel
				createTargetOptionsModel.EventstreamsEndpoint = eventstreamsEndpointPrototypeModel
				createTargetOptionsModel.CloudlogsEndpoint = cloudLogsEndpointPrototypeModel
				createTargetOptionsModel.Region = core.StringPtr("us-south")
				createTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = atrackerService.CreateTarget(createTargetOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateTarget with error: Operation validation and request error`, func() {
				atrackerService, serviceErr := atrackerv2.NewAtrackerV2(&atrackerv2.AtrackerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())

				// Construct an instance of the CosEndpointPrototype model
				cosEndpointPrototypeModel := new(atrackerv2.CosEndpointPrototype)
				cosEndpointPrototypeModel.Endpoint = core.StringPtr("s3.private.us-east.cloud-object-storage.appdomain.cloud")
				cosEndpointPrototypeModel.TargetCRN = core.StringPtr("crn:v1:bluemix:public:cloud-object-storage:global:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::")
				cosEndpointPrototypeModel.Bucket = core.StringPtr("my-atracker-bucket")
				cosEndpointPrototypeModel.APIKey = core.StringPtr("xxxxxxxxxxxxxx")
				cosEndpointPrototypeModel.ServiceToServiceEnabled = core.BoolPtr(true)

				// Construct an instance of the EventstreamsEndpointPrototype model
				eventstreamsEndpointPrototypeModel := new(atrackerv2.EventstreamsEndpointPrototype)
				eventstreamsEndpointPrototypeModel.TargetCRN = core.StringPtr("crn:v1:bluemix:public:messagehub:us-south:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::")
				eventstreamsEndpointPrototypeModel.Brokers = []string{"kafka-x:9094"}
				eventstreamsEndpointPrototypeModel.Topic = core.StringPtr("my-topic")
				eventstreamsEndpointPrototypeModel.APIKey = core.StringPtr("xxxxxxxxxxxxxx")
				eventstreamsEndpointPrototypeModel.ServiceToServiceEnabled = core.BoolPtr(false)

				// Construct an instance of the CloudLogsEndpointPrototype model
				cloudLogsEndpointPrototypeModel := new(atrackerv2.CloudLogsEndpointPrototype)
				cloudLogsEndpointPrototypeModel.TargetCRN = core.StringPtr("crn:v1:bluemix:public:logs:eu-es:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::")

				// Construct an instance of the CreateTargetOptions model
				createTargetOptionsModel := new(atrackerv2.CreateTargetOptions)
				createTargetOptionsModel.Name = core.StringPtr("my-cos-target")
				createTargetOptionsModel.TargetType = core.StringPtr("cloud_object_storage")
				createTargetOptionsModel.CosEndpoint = cosEndpointPrototypeModel
				createTargetOptionsModel.EventstreamsEndpoint = eventstreamsEndpointPrototypeModel
				createTargetOptionsModel.CloudlogsEndpoint = cloudLogsEndpointPrototypeModel
				createTargetOptionsModel.Region = core.StringPtr("us-south")
				createTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := atrackerService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := atrackerService.CreateTarget(createTargetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateTargetOptions model with no property values
				createTargetOptionsModelNew := new(atrackerv2.CreateTargetOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = atrackerService.CreateTarget(createTargetOptionsModelNew)
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
				atrackerService, serviceErr := atrackerv2.NewAtrackerV2(&atrackerv2.AtrackerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())

				// Construct an instance of the CosEndpointPrototype model
				cosEndpointPrototypeModel := new(atrackerv2.CosEndpointPrototype)
				cosEndpointPrototypeModel.Endpoint = core.StringPtr("s3.private.us-east.cloud-object-storage.appdomain.cloud")
				cosEndpointPrototypeModel.TargetCRN = core.StringPtr("crn:v1:bluemix:public:cloud-object-storage:global:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::")
				cosEndpointPrototypeModel.Bucket = core.StringPtr("my-atracker-bucket")
				cosEndpointPrototypeModel.APIKey = core.StringPtr("xxxxxxxxxxxxxx")
				cosEndpointPrototypeModel.ServiceToServiceEnabled = core.BoolPtr(true)

				// Construct an instance of the EventstreamsEndpointPrototype model
				eventstreamsEndpointPrototypeModel := new(atrackerv2.EventstreamsEndpointPrototype)
				eventstreamsEndpointPrototypeModel.TargetCRN = core.StringPtr("crn:v1:bluemix:public:messagehub:us-south:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::")
				eventstreamsEndpointPrototypeModel.Brokers = []string{"kafka-x:9094"}
				eventstreamsEndpointPrototypeModel.Topic = core.StringPtr("my-topic")
				eventstreamsEndpointPrototypeModel.APIKey = core.StringPtr("xxxxxxxxxxxxxx")
				eventstreamsEndpointPrototypeModel.ServiceToServiceEnabled = core.BoolPtr(false)

				// Construct an instance of the CloudLogsEndpointPrototype model
				cloudLogsEndpointPrototypeModel := new(atrackerv2.CloudLogsEndpointPrototype)
				cloudLogsEndpointPrototypeModel.TargetCRN = core.StringPtr("crn:v1:bluemix:public:logs:eu-es:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::")

				// Construct an instance of the CreateTargetOptions model
				createTargetOptionsModel := new(atrackerv2.CreateTargetOptions)
				createTargetOptionsModel.Name = core.StringPtr("my-cos-target")
				createTargetOptionsModel.TargetType = core.StringPtr("cloud_object_storage")
				createTargetOptionsModel.CosEndpoint = cosEndpointPrototypeModel
				createTargetOptionsModel.EventstreamsEndpoint = eventstreamsEndpointPrototypeModel
				createTargetOptionsModel.CloudlogsEndpoint = cloudLogsEndpointPrototypeModel
				createTargetOptionsModel.Region = core.StringPtr("us-south")
				createTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := atrackerService.CreateTarget(createTargetOptionsModel)
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
		listTargetsPath := "/api/v2/targets"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listTargetsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["region"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListTargets with error: Operation response processing error`, func() {
				atrackerService, serviceErr := atrackerv2.NewAtrackerV2(&atrackerv2.AtrackerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())

				// Construct an instance of the ListTargetsOptions model
				listTargetsOptionsModel := new(atrackerv2.ListTargetsOptions)
				listTargetsOptionsModel.Region = core.StringPtr("testString")
				listTargetsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := atrackerService.ListTargets(listTargetsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				atrackerService.EnableRetries(0, 0)
				result, response, operationErr = atrackerService.ListTargets(listTargetsOptionsModel)
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
		listTargetsPath := "/api/v2/targets"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listTargetsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["region"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"targets": [{"id": "f7dcfae6-e7c5-08ca-451b-fdfa696c9bb6", "name": "a-cos-target-us-south", "crn": "crn:v1:bluemix:public:atracker:us-south:a/11111111111111111111111111111111:b6eec08b-5201-08ca-451b-cd71523e3626:target:f7dcfae6-e7c5-08ca-451b-fdfa696c9bb6", "target_type": "cloud_object_storage", "region": "us-south", "cos_endpoint": {"endpoint": "s3.private.us-east.cloud-object-storage.appdomain.cloud", "target_crn": "crn:v1:bluemix:public:cloud-object-storage:global:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::", "bucket": "my-atracker-bucket", "service_to_service_enabled": true}, "eventstreams_endpoint": {"target_crn": "crn:v1:bluemix:public:messagehub:us-south:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::", "brokers": ["kafka-x:9094"], "topic": "my-topic", "api_key": "xxxxxxxxxxxxxx", "service_to_service_enabled": false}, "cloudlogs_endpoint": {"target_crn": "crn:v1:bluemix:public:eu-es:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::"}, "write_status": {"status": "success", "last_failure": "2021-05-18T20:15:12.353Z", "reason_for_last_failure": "Provided API key could not be found"}, "created_at": "2021-05-18T20:15:12.353Z", "updated_at": "2021-05-18T20:15:12.353Z", "message": "This is a valid target. However, there is another target already defined with the same target endpoint.", "api_version": 2}]}`)
				}))
			})
			It(`Invoke ListTargets successfully with retries`, func() {
				atrackerService, serviceErr := atrackerv2.NewAtrackerV2(&atrackerv2.AtrackerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())
				atrackerService.EnableRetries(0, 0)

				// Construct an instance of the ListTargetsOptions model
				listTargetsOptionsModel := new(atrackerv2.ListTargetsOptions)
				listTargetsOptionsModel.Region = core.StringPtr("testString")
				listTargetsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := atrackerService.ListTargetsWithContext(ctx, listTargetsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				atrackerService.DisableRetries()
				result, response, operationErr := atrackerService.ListTargets(listTargetsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = atrackerService.ListTargetsWithContext(ctx, listTargetsOptionsModel)
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

					Expect(req.URL.Query()["region"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"targets": [{"id": "f7dcfae6-e7c5-08ca-451b-fdfa696c9bb6", "name": "a-cos-target-us-south", "crn": "crn:v1:bluemix:public:atracker:us-south:a/11111111111111111111111111111111:b6eec08b-5201-08ca-451b-cd71523e3626:target:f7dcfae6-e7c5-08ca-451b-fdfa696c9bb6", "target_type": "cloud_object_storage", "region": "us-south", "cos_endpoint": {"endpoint": "s3.private.us-east.cloud-object-storage.appdomain.cloud", "target_crn": "crn:v1:bluemix:public:cloud-object-storage:global:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::", "bucket": "my-atracker-bucket", "service_to_service_enabled": true}, "eventstreams_endpoint": {"target_crn": "crn:v1:bluemix:public:messagehub:us-south:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::", "brokers": ["kafka-x:9094"], "topic": "my-topic", "api_key": "xxxxxxxxxxxxxx", "service_to_service_enabled": false}, "cloudlogs_endpoint": {"target_crn": "crn:v1:bluemix:public:eu-es:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::"}, "write_status": {"status": "success", "last_failure": "2021-05-18T20:15:12.353Z", "reason_for_last_failure": "Provided API key could not be found"}, "created_at": "2021-05-18T20:15:12.353Z", "updated_at": "2021-05-18T20:15:12.353Z", "message": "This is a valid target. However, there is another target already defined with the same target endpoint.", "api_version": 2}]}`)
				}))
			})
			It(`Invoke ListTargets successfully`, func() {
				atrackerService, serviceErr := atrackerv2.NewAtrackerV2(&atrackerv2.AtrackerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := atrackerService.ListTargets(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListTargetsOptions model
				listTargetsOptionsModel := new(atrackerv2.ListTargetsOptions)
				listTargetsOptionsModel.Region = core.StringPtr("testString")
				listTargetsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = atrackerService.ListTargets(listTargetsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListTargets with error: Operation request error`, func() {
				atrackerService, serviceErr := atrackerv2.NewAtrackerV2(&atrackerv2.AtrackerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())

				// Construct an instance of the ListTargetsOptions model
				listTargetsOptionsModel := new(atrackerv2.ListTargetsOptions)
				listTargetsOptionsModel.Region = core.StringPtr("testString")
				listTargetsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := atrackerService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := atrackerService.ListTargets(listTargetsOptionsModel)
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
				atrackerService, serviceErr := atrackerv2.NewAtrackerV2(&atrackerv2.AtrackerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())

				// Construct an instance of the ListTargetsOptions model
				listTargetsOptionsModel := new(atrackerv2.ListTargetsOptions)
				listTargetsOptionsModel.Region = core.StringPtr("testString")
				listTargetsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := atrackerService.ListTargets(listTargetsOptionsModel)
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
		getTargetPath := "/api/v2/targets/testString"
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
				atrackerService, serviceErr := atrackerv2.NewAtrackerV2(&atrackerv2.AtrackerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())

				// Construct an instance of the GetTargetOptions model
				getTargetOptionsModel := new(atrackerv2.GetTargetOptions)
				getTargetOptionsModel.ID = core.StringPtr("testString")
				getTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := atrackerService.GetTarget(getTargetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				atrackerService.EnableRetries(0, 0)
				result, response, operationErr = atrackerService.GetTarget(getTargetOptionsModel)
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
		getTargetPath := "/api/v2/targets/testString"
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
					fmt.Fprintf(res, "%s", `{"id": "f7dcfae6-e7c5-08ca-451b-fdfa696c9bb6", "name": "a-cos-target-us-south", "crn": "crn:v1:bluemix:public:atracker:us-south:a/11111111111111111111111111111111:b6eec08b-5201-08ca-451b-cd71523e3626:target:f7dcfae6-e7c5-08ca-451b-fdfa696c9bb6", "target_type": "cloud_object_storage", "region": "us-south", "cos_endpoint": {"endpoint": "s3.private.us-east.cloud-object-storage.appdomain.cloud", "target_crn": "crn:v1:bluemix:public:cloud-object-storage:global:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::", "bucket": "my-atracker-bucket", "service_to_service_enabled": true}, "eventstreams_endpoint": {"target_crn": "crn:v1:bluemix:public:messagehub:us-south:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::", "brokers": ["kafka-x:9094"], "topic": "my-topic", "api_key": "xxxxxxxxxxxxxx", "service_to_service_enabled": false}, "cloudlogs_endpoint": {"target_crn": "crn:v1:bluemix:public:eu-es:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::"}, "write_status": {"status": "success", "last_failure": "2021-05-18T20:15:12.353Z", "reason_for_last_failure": "Provided API key could not be found"}, "created_at": "2021-05-18T20:15:12.353Z", "updated_at": "2021-05-18T20:15:12.353Z", "message": "This is a valid target. However, there is another target already defined with the same target endpoint.", "api_version": 2}`)
				}))
			})
			It(`Invoke GetTarget successfully with retries`, func() {
				atrackerService, serviceErr := atrackerv2.NewAtrackerV2(&atrackerv2.AtrackerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())
				atrackerService.EnableRetries(0, 0)

				// Construct an instance of the GetTargetOptions model
				getTargetOptionsModel := new(atrackerv2.GetTargetOptions)
				getTargetOptionsModel.ID = core.StringPtr("testString")
				getTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := atrackerService.GetTargetWithContext(ctx, getTargetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				atrackerService.DisableRetries()
				result, response, operationErr := atrackerService.GetTarget(getTargetOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = atrackerService.GetTargetWithContext(ctx, getTargetOptionsModel)
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
					fmt.Fprintf(res, "%s", `{"id": "f7dcfae6-e7c5-08ca-451b-fdfa696c9bb6", "name": "a-cos-target-us-south", "crn": "crn:v1:bluemix:public:atracker:us-south:a/11111111111111111111111111111111:b6eec08b-5201-08ca-451b-cd71523e3626:target:f7dcfae6-e7c5-08ca-451b-fdfa696c9bb6", "target_type": "cloud_object_storage", "region": "us-south", "cos_endpoint": {"endpoint": "s3.private.us-east.cloud-object-storage.appdomain.cloud", "target_crn": "crn:v1:bluemix:public:cloud-object-storage:global:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::", "bucket": "my-atracker-bucket", "service_to_service_enabled": true}, "eventstreams_endpoint": {"target_crn": "crn:v1:bluemix:public:messagehub:us-south:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::", "brokers": ["kafka-x:9094"], "topic": "my-topic", "api_key": "xxxxxxxxxxxxxx", "service_to_service_enabled": false}, "cloudlogs_endpoint": {"target_crn": "crn:v1:bluemix:public:eu-es:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::"}, "write_status": {"status": "success", "last_failure": "2021-05-18T20:15:12.353Z", "reason_for_last_failure": "Provided API key could not be found"}, "created_at": "2021-05-18T20:15:12.353Z", "updated_at": "2021-05-18T20:15:12.353Z", "message": "This is a valid target. However, there is another target already defined with the same target endpoint.", "api_version": 2}`)
				}))
			})
			It(`Invoke GetTarget successfully`, func() {
				atrackerService, serviceErr := atrackerv2.NewAtrackerV2(&atrackerv2.AtrackerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := atrackerService.GetTarget(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetTargetOptions model
				getTargetOptionsModel := new(atrackerv2.GetTargetOptions)
				getTargetOptionsModel.ID = core.StringPtr("testString")
				getTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = atrackerService.GetTarget(getTargetOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetTarget with error: Operation validation and request error`, func() {
				atrackerService, serviceErr := atrackerv2.NewAtrackerV2(&atrackerv2.AtrackerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())

				// Construct an instance of the GetTargetOptions model
				getTargetOptionsModel := new(atrackerv2.GetTargetOptions)
				getTargetOptionsModel.ID = core.StringPtr("testString")
				getTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := atrackerService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := atrackerService.GetTarget(getTargetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetTargetOptions model with no property values
				getTargetOptionsModelNew := new(atrackerv2.GetTargetOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = atrackerService.GetTarget(getTargetOptionsModelNew)
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
				atrackerService, serviceErr := atrackerv2.NewAtrackerV2(&atrackerv2.AtrackerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())

				// Construct an instance of the GetTargetOptions model
				getTargetOptionsModel := new(atrackerv2.GetTargetOptions)
				getTargetOptionsModel.ID = core.StringPtr("testString")
				getTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := atrackerService.GetTarget(getTargetOptionsModel)
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
	Describe(`ReplaceTarget(replaceTargetOptions *ReplaceTargetOptions) - Operation response error`, func() {
		replaceTargetPath := "/api/v2/targets/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceTargetPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ReplaceTarget with error: Operation response processing error`, func() {
				atrackerService, serviceErr := atrackerv2.NewAtrackerV2(&atrackerv2.AtrackerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())

				// Construct an instance of the CosEndpointPrototype model
				cosEndpointPrototypeModel := new(atrackerv2.CosEndpointPrototype)
				cosEndpointPrototypeModel.Endpoint = core.StringPtr("s3.private.us-east.cloud-object-storage.appdomain.cloud")
				cosEndpointPrototypeModel.TargetCRN = core.StringPtr("crn:v1:bluemix:public:cloud-object-storage:global:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::")
				cosEndpointPrototypeModel.Bucket = core.StringPtr("my-atracker-bucket")
				cosEndpointPrototypeModel.APIKey = core.StringPtr("xxxxxxxxxxxxxx")
				cosEndpointPrototypeModel.ServiceToServiceEnabled = core.BoolPtr(true)

				// Construct an instance of the EventstreamsEndpointPrototype model
				eventstreamsEndpointPrototypeModel := new(atrackerv2.EventstreamsEndpointPrototype)
				eventstreamsEndpointPrototypeModel.TargetCRN = core.StringPtr("crn:v1:bluemix:public:messagehub:us-south:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::")
				eventstreamsEndpointPrototypeModel.Brokers = []string{"kafka-x:9094"}
				eventstreamsEndpointPrototypeModel.Topic = core.StringPtr("my-topic")
				eventstreamsEndpointPrototypeModel.APIKey = core.StringPtr("xxxxxxxxxxxxxx")
				eventstreamsEndpointPrototypeModel.ServiceToServiceEnabled = core.BoolPtr(false)

				// Construct an instance of the CloudLogsEndpointPrototype model
				cloudLogsEndpointPrototypeModel := new(atrackerv2.CloudLogsEndpointPrototype)
				cloudLogsEndpointPrototypeModel.TargetCRN = core.StringPtr("crn:v1:bluemix:public:logs:eu-es:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::")

				// Construct an instance of the ReplaceTargetOptions model
				replaceTargetOptionsModel := new(atrackerv2.ReplaceTargetOptions)
				replaceTargetOptionsModel.ID = core.StringPtr("testString")
				replaceTargetOptionsModel.Name = core.StringPtr("my-cos-target")
				replaceTargetOptionsModel.CosEndpoint = cosEndpointPrototypeModel
				replaceTargetOptionsModel.EventstreamsEndpoint = eventstreamsEndpointPrototypeModel
				replaceTargetOptionsModel.CloudlogsEndpoint = cloudLogsEndpointPrototypeModel
				replaceTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := atrackerService.ReplaceTarget(replaceTargetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				atrackerService.EnableRetries(0, 0)
				result, response, operationErr = atrackerService.ReplaceTarget(replaceTargetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ReplaceTarget(replaceTargetOptions *ReplaceTargetOptions)`, func() {
		replaceTargetPath := "/api/v2/targets/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceTargetPath))
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
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "f7dcfae6-e7c5-08ca-451b-fdfa696c9bb6", "name": "a-cos-target-us-south", "crn": "crn:v1:bluemix:public:atracker:us-south:a/11111111111111111111111111111111:b6eec08b-5201-08ca-451b-cd71523e3626:target:f7dcfae6-e7c5-08ca-451b-fdfa696c9bb6", "target_type": "cloud_object_storage", "region": "us-south", "cos_endpoint": {"endpoint": "s3.private.us-east.cloud-object-storage.appdomain.cloud", "target_crn": "crn:v1:bluemix:public:cloud-object-storage:global:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::", "bucket": "my-atracker-bucket", "service_to_service_enabled": true}, "eventstreams_endpoint": {"target_crn": "crn:v1:bluemix:public:messagehub:us-south:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::", "brokers": ["kafka-x:9094"], "topic": "my-topic", "api_key": "xxxxxxxxxxxxxx", "service_to_service_enabled": false}, "cloudlogs_endpoint": {"target_crn": "crn:v1:bluemix:public:eu-es:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::"}, "write_status": {"status": "success", "last_failure": "2021-05-18T20:15:12.353Z", "reason_for_last_failure": "Provided API key could not be found"}, "created_at": "2021-05-18T20:15:12.353Z", "updated_at": "2021-05-18T20:15:12.353Z", "message": "This is a valid target. However, there is another target already defined with the same target endpoint.", "api_version": 2}`)
				}))
			})
			It(`Invoke ReplaceTarget successfully with retries`, func() {
				atrackerService, serviceErr := atrackerv2.NewAtrackerV2(&atrackerv2.AtrackerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())
				atrackerService.EnableRetries(0, 0)

				// Construct an instance of the CosEndpointPrototype model
				cosEndpointPrototypeModel := new(atrackerv2.CosEndpointPrototype)
				cosEndpointPrototypeModel.Endpoint = core.StringPtr("s3.private.us-east.cloud-object-storage.appdomain.cloud")
				cosEndpointPrototypeModel.TargetCRN = core.StringPtr("crn:v1:bluemix:public:cloud-object-storage:global:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::")
				cosEndpointPrototypeModel.Bucket = core.StringPtr("my-atracker-bucket")
				cosEndpointPrototypeModel.APIKey = core.StringPtr("xxxxxxxxxxxxxx")
				cosEndpointPrototypeModel.ServiceToServiceEnabled = core.BoolPtr(true)

				// Construct an instance of the EventstreamsEndpointPrototype model
				eventstreamsEndpointPrototypeModel := new(atrackerv2.EventstreamsEndpointPrototype)
				eventstreamsEndpointPrototypeModel.TargetCRN = core.StringPtr("crn:v1:bluemix:public:messagehub:us-south:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::")
				eventstreamsEndpointPrototypeModel.Brokers = []string{"kafka-x:9094"}
				eventstreamsEndpointPrototypeModel.Topic = core.StringPtr("my-topic")
				eventstreamsEndpointPrototypeModel.APIKey = core.StringPtr("xxxxxxxxxxxxxx")
				eventstreamsEndpointPrototypeModel.ServiceToServiceEnabled = core.BoolPtr(false)

				// Construct an instance of the CloudLogsEndpointPrototype model
				cloudLogsEndpointPrototypeModel := new(atrackerv2.CloudLogsEndpointPrototype)
				cloudLogsEndpointPrototypeModel.TargetCRN = core.StringPtr("crn:v1:bluemix:public:logs:eu-es:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::")

				// Construct an instance of the ReplaceTargetOptions model
				replaceTargetOptionsModel := new(atrackerv2.ReplaceTargetOptions)
				replaceTargetOptionsModel.ID = core.StringPtr("testString")
				replaceTargetOptionsModel.Name = core.StringPtr("my-cos-target")
				replaceTargetOptionsModel.CosEndpoint = cosEndpointPrototypeModel
				replaceTargetOptionsModel.EventstreamsEndpoint = eventstreamsEndpointPrototypeModel
				replaceTargetOptionsModel.CloudlogsEndpoint = cloudLogsEndpointPrototypeModel
				replaceTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := atrackerService.ReplaceTargetWithContext(ctx, replaceTargetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				atrackerService.DisableRetries()
				result, response, operationErr := atrackerService.ReplaceTarget(replaceTargetOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = atrackerService.ReplaceTargetWithContext(ctx, replaceTargetOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(replaceTargetPath))
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

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "f7dcfae6-e7c5-08ca-451b-fdfa696c9bb6", "name": "a-cos-target-us-south", "crn": "crn:v1:bluemix:public:atracker:us-south:a/11111111111111111111111111111111:b6eec08b-5201-08ca-451b-cd71523e3626:target:f7dcfae6-e7c5-08ca-451b-fdfa696c9bb6", "target_type": "cloud_object_storage", "region": "us-south", "cos_endpoint": {"endpoint": "s3.private.us-east.cloud-object-storage.appdomain.cloud", "target_crn": "crn:v1:bluemix:public:cloud-object-storage:global:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::", "bucket": "my-atracker-bucket", "service_to_service_enabled": true}, "eventstreams_endpoint": {"target_crn": "crn:v1:bluemix:public:messagehub:us-south:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::", "brokers": ["kafka-x:9094"], "topic": "my-topic", "api_key": "xxxxxxxxxxxxxx", "service_to_service_enabled": false}, "cloudlogs_endpoint": {"target_crn": "crn:v1:bluemix:public:eu-es:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::"}, "write_status": {"status": "success", "last_failure": "2021-05-18T20:15:12.353Z", "reason_for_last_failure": "Provided API key could not be found"}, "created_at": "2021-05-18T20:15:12.353Z", "updated_at": "2021-05-18T20:15:12.353Z", "message": "This is a valid target. However, there is another target already defined with the same target endpoint.", "api_version": 2}`)
				}))
			})
			It(`Invoke ReplaceTarget successfully`, func() {
				atrackerService, serviceErr := atrackerv2.NewAtrackerV2(&atrackerv2.AtrackerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := atrackerService.ReplaceTarget(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CosEndpointPrototype model
				cosEndpointPrototypeModel := new(atrackerv2.CosEndpointPrototype)
				cosEndpointPrototypeModel.Endpoint = core.StringPtr("s3.private.us-east.cloud-object-storage.appdomain.cloud")
				cosEndpointPrototypeModel.TargetCRN = core.StringPtr("crn:v1:bluemix:public:cloud-object-storage:global:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::")
				cosEndpointPrototypeModel.Bucket = core.StringPtr("my-atracker-bucket")
				cosEndpointPrototypeModel.APIKey = core.StringPtr("xxxxxxxxxxxxxx")
				cosEndpointPrototypeModel.ServiceToServiceEnabled = core.BoolPtr(true)

				// Construct an instance of the EventstreamsEndpointPrototype model
				eventstreamsEndpointPrototypeModel := new(atrackerv2.EventstreamsEndpointPrototype)
				eventstreamsEndpointPrototypeModel.TargetCRN = core.StringPtr("crn:v1:bluemix:public:messagehub:us-south:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::")
				eventstreamsEndpointPrototypeModel.Brokers = []string{"kafka-x:9094"}
				eventstreamsEndpointPrototypeModel.Topic = core.StringPtr("my-topic")
				eventstreamsEndpointPrototypeModel.APIKey = core.StringPtr("xxxxxxxxxxxxxx")
				eventstreamsEndpointPrototypeModel.ServiceToServiceEnabled = core.BoolPtr(false)

				// Construct an instance of the CloudLogsEndpointPrototype model
				cloudLogsEndpointPrototypeModel := new(atrackerv2.CloudLogsEndpointPrototype)
				cloudLogsEndpointPrototypeModel.TargetCRN = core.StringPtr("crn:v1:bluemix:public:logs:eu-es:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::")

				// Construct an instance of the ReplaceTargetOptions model
				replaceTargetOptionsModel := new(atrackerv2.ReplaceTargetOptions)
				replaceTargetOptionsModel.ID = core.StringPtr("testString")
				replaceTargetOptionsModel.Name = core.StringPtr("my-cos-target")
				replaceTargetOptionsModel.CosEndpoint = cosEndpointPrototypeModel
				replaceTargetOptionsModel.EventstreamsEndpoint = eventstreamsEndpointPrototypeModel
				replaceTargetOptionsModel.CloudlogsEndpoint = cloudLogsEndpointPrototypeModel
				replaceTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = atrackerService.ReplaceTarget(replaceTargetOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ReplaceTarget with error: Operation validation and request error`, func() {
				atrackerService, serviceErr := atrackerv2.NewAtrackerV2(&atrackerv2.AtrackerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())

				// Construct an instance of the CosEndpointPrototype model
				cosEndpointPrototypeModel := new(atrackerv2.CosEndpointPrototype)
				cosEndpointPrototypeModel.Endpoint = core.StringPtr("s3.private.us-east.cloud-object-storage.appdomain.cloud")
				cosEndpointPrototypeModel.TargetCRN = core.StringPtr("crn:v1:bluemix:public:cloud-object-storage:global:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::")
				cosEndpointPrototypeModel.Bucket = core.StringPtr("my-atracker-bucket")
				cosEndpointPrototypeModel.APIKey = core.StringPtr("xxxxxxxxxxxxxx")
				cosEndpointPrototypeModel.ServiceToServiceEnabled = core.BoolPtr(true)

				// Construct an instance of the EventstreamsEndpointPrototype model
				eventstreamsEndpointPrototypeModel := new(atrackerv2.EventstreamsEndpointPrototype)
				eventstreamsEndpointPrototypeModel.TargetCRN = core.StringPtr("crn:v1:bluemix:public:messagehub:us-south:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::")
				eventstreamsEndpointPrototypeModel.Brokers = []string{"kafka-x:9094"}
				eventstreamsEndpointPrototypeModel.Topic = core.StringPtr("my-topic")
				eventstreamsEndpointPrototypeModel.APIKey = core.StringPtr("xxxxxxxxxxxxxx")
				eventstreamsEndpointPrototypeModel.ServiceToServiceEnabled = core.BoolPtr(false)

				// Construct an instance of the CloudLogsEndpointPrototype model
				cloudLogsEndpointPrototypeModel := new(atrackerv2.CloudLogsEndpointPrototype)
				cloudLogsEndpointPrototypeModel.TargetCRN = core.StringPtr("crn:v1:bluemix:public:logs:eu-es:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::")

				// Construct an instance of the ReplaceTargetOptions model
				replaceTargetOptionsModel := new(atrackerv2.ReplaceTargetOptions)
				replaceTargetOptionsModel.ID = core.StringPtr("testString")
				replaceTargetOptionsModel.Name = core.StringPtr("my-cos-target")
				replaceTargetOptionsModel.CosEndpoint = cosEndpointPrototypeModel
				replaceTargetOptionsModel.EventstreamsEndpoint = eventstreamsEndpointPrototypeModel
				replaceTargetOptionsModel.CloudlogsEndpoint = cloudLogsEndpointPrototypeModel
				replaceTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := atrackerService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := atrackerService.ReplaceTarget(replaceTargetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ReplaceTargetOptions model with no property values
				replaceTargetOptionsModelNew := new(atrackerv2.ReplaceTargetOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = atrackerService.ReplaceTarget(replaceTargetOptionsModelNew)
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
			It(`Invoke ReplaceTarget successfully`, func() {
				atrackerService, serviceErr := atrackerv2.NewAtrackerV2(&atrackerv2.AtrackerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())

				// Construct an instance of the CosEndpointPrototype model
				cosEndpointPrototypeModel := new(atrackerv2.CosEndpointPrototype)
				cosEndpointPrototypeModel.Endpoint = core.StringPtr("s3.private.us-east.cloud-object-storage.appdomain.cloud")
				cosEndpointPrototypeModel.TargetCRN = core.StringPtr("crn:v1:bluemix:public:cloud-object-storage:global:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::")
				cosEndpointPrototypeModel.Bucket = core.StringPtr("my-atracker-bucket")
				cosEndpointPrototypeModel.APIKey = core.StringPtr("xxxxxxxxxxxxxx")
				cosEndpointPrototypeModel.ServiceToServiceEnabled = core.BoolPtr(true)

				// Construct an instance of the EventstreamsEndpointPrototype model
				eventstreamsEndpointPrototypeModel := new(atrackerv2.EventstreamsEndpointPrototype)
				eventstreamsEndpointPrototypeModel.TargetCRN = core.StringPtr("crn:v1:bluemix:public:messagehub:us-south:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::")
				eventstreamsEndpointPrototypeModel.Brokers = []string{"kafka-x:9094"}
				eventstreamsEndpointPrototypeModel.Topic = core.StringPtr("my-topic")
				eventstreamsEndpointPrototypeModel.APIKey = core.StringPtr("xxxxxxxxxxxxxx")
				eventstreamsEndpointPrototypeModel.ServiceToServiceEnabled = core.BoolPtr(false)

				// Construct an instance of the CloudLogsEndpointPrototype model
				cloudLogsEndpointPrototypeModel := new(atrackerv2.CloudLogsEndpointPrototype)
				cloudLogsEndpointPrototypeModel.TargetCRN = core.StringPtr("crn:v1:bluemix:public:logs:eu-es:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::")

				// Construct an instance of the ReplaceTargetOptions model
				replaceTargetOptionsModel := new(atrackerv2.ReplaceTargetOptions)
				replaceTargetOptionsModel.ID = core.StringPtr("testString")
				replaceTargetOptionsModel.Name = core.StringPtr("my-cos-target")
				replaceTargetOptionsModel.CosEndpoint = cosEndpointPrototypeModel
				replaceTargetOptionsModel.EventstreamsEndpoint = eventstreamsEndpointPrototypeModel
				replaceTargetOptionsModel.CloudlogsEndpoint = cloudLogsEndpointPrototypeModel
				replaceTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := atrackerService.ReplaceTarget(replaceTargetOptionsModel)
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
	Describe(`DeleteTarget(deleteTargetOptions *DeleteTargetOptions) - Operation response error`, func() {
		deleteTargetPath := "/api/v2/targets/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteTargetPath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteTarget with error: Operation response processing error`, func() {
				atrackerService, serviceErr := atrackerv2.NewAtrackerV2(&atrackerv2.AtrackerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())

				// Construct an instance of the DeleteTargetOptions model
				deleteTargetOptionsModel := new(atrackerv2.DeleteTargetOptions)
				deleteTargetOptionsModel.ID = core.StringPtr("testString")
				deleteTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := atrackerService.DeleteTarget(deleteTargetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				atrackerService.EnableRetries(0, 0)
				result, response, operationErr = atrackerService.DeleteTarget(deleteTargetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteTarget(deleteTargetOptions *DeleteTargetOptions)`, func() {
		deleteTargetPath := "/api/v2/targets/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteTargetPath))
					Expect(req.Method).To(Equal("DELETE"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"status_code": 10, "trace": "Trace", "warnings": [{"code": "Code", "message": "Message"}]}`)
				}))
			})
			It(`Invoke DeleteTarget successfully with retries`, func() {
				atrackerService, serviceErr := atrackerv2.NewAtrackerV2(&atrackerv2.AtrackerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())
				atrackerService.EnableRetries(0, 0)

				// Construct an instance of the DeleteTargetOptions model
				deleteTargetOptionsModel := new(atrackerv2.DeleteTargetOptions)
				deleteTargetOptionsModel.ID = core.StringPtr("testString")
				deleteTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := atrackerService.DeleteTargetWithContext(ctx, deleteTargetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				atrackerService.DisableRetries()
				result, response, operationErr := atrackerService.DeleteTarget(deleteTargetOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = atrackerService.DeleteTargetWithContext(ctx, deleteTargetOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(deleteTargetPath))
					Expect(req.Method).To(Equal("DELETE"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"status_code": 10, "trace": "Trace", "warnings": [{"code": "Code", "message": "Message"}]}`)
				}))
			})
			It(`Invoke DeleteTarget successfully`, func() {
				atrackerService, serviceErr := atrackerv2.NewAtrackerV2(&atrackerv2.AtrackerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := atrackerService.DeleteTarget(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteTargetOptions model
				deleteTargetOptionsModel := new(atrackerv2.DeleteTargetOptions)
				deleteTargetOptionsModel.ID = core.StringPtr("testString")
				deleteTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = atrackerService.DeleteTarget(deleteTargetOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke DeleteTarget with error: Operation validation and request error`, func() {
				atrackerService, serviceErr := atrackerv2.NewAtrackerV2(&atrackerv2.AtrackerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())

				// Construct an instance of the DeleteTargetOptions model
				deleteTargetOptionsModel := new(atrackerv2.DeleteTargetOptions)
				deleteTargetOptionsModel.ID = core.StringPtr("testString")
				deleteTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := atrackerService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := atrackerService.DeleteTarget(deleteTargetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteTargetOptions model with no property values
				deleteTargetOptionsModelNew := new(atrackerv2.DeleteTargetOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = atrackerService.DeleteTarget(deleteTargetOptionsModelNew)
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
			It(`Invoke DeleteTarget successfully`, func() {
				atrackerService, serviceErr := atrackerv2.NewAtrackerV2(&atrackerv2.AtrackerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())

				// Construct an instance of the DeleteTargetOptions model
				deleteTargetOptionsModel := new(atrackerv2.DeleteTargetOptions)
				deleteTargetOptionsModel.ID = core.StringPtr("testString")
				deleteTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := atrackerService.DeleteTarget(deleteTargetOptionsModel)
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
	Describe(`ValidateTarget(validateTargetOptions *ValidateTargetOptions) - Operation response error`, func() {
		validateTargetPath := "/api/v2/targets/testString/validate"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(validateTargetPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ValidateTarget with error: Operation response processing error`, func() {
				atrackerService, serviceErr := atrackerv2.NewAtrackerV2(&atrackerv2.AtrackerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())

				// Construct an instance of the ValidateTargetOptions model
				validateTargetOptionsModel := new(atrackerv2.ValidateTargetOptions)
				validateTargetOptionsModel.ID = core.StringPtr("testString")
				validateTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := atrackerService.ValidateTarget(validateTargetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				atrackerService.EnableRetries(0, 0)
				result, response, operationErr = atrackerService.ValidateTarget(validateTargetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ValidateTarget(validateTargetOptions *ValidateTargetOptions)`, func() {
		validateTargetPath := "/api/v2/targets/testString/validate"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(validateTargetPath))
					Expect(req.Method).To(Equal("POST"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "f7dcfae6-e7c5-08ca-451b-fdfa696c9bb6", "name": "a-cos-target-us-south", "crn": "crn:v1:bluemix:public:atracker:us-south:a/11111111111111111111111111111111:b6eec08b-5201-08ca-451b-cd71523e3626:target:f7dcfae6-e7c5-08ca-451b-fdfa696c9bb6", "target_type": "cloud_object_storage", "region": "us-south", "cos_endpoint": {"endpoint": "s3.private.us-east.cloud-object-storage.appdomain.cloud", "target_crn": "crn:v1:bluemix:public:cloud-object-storage:global:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::", "bucket": "my-atracker-bucket", "service_to_service_enabled": true}, "eventstreams_endpoint": {"target_crn": "crn:v1:bluemix:public:messagehub:us-south:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::", "brokers": ["kafka-x:9094"], "topic": "my-topic", "api_key": "xxxxxxxxxxxxxx", "service_to_service_enabled": false}, "cloudlogs_endpoint": {"target_crn": "crn:v1:bluemix:public:eu-es:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::"}, "write_status": {"status": "success", "last_failure": "2021-05-18T20:15:12.353Z", "reason_for_last_failure": "Provided API key could not be found"}, "created_at": "2021-05-18T20:15:12.353Z", "updated_at": "2021-05-18T20:15:12.353Z", "message": "This is a valid target. However, there is another target already defined with the same target endpoint.", "api_version": 2}`)
				}))
			})
			It(`Invoke ValidateTarget successfully with retries`, func() {
				atrackerService, serviceErr := atrackerv2.NewAtrackerV2(&atrackerv2.AtrackerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())
				atrackerService.EnableRetries(0, 0)

				// Construct an instance of the ValidateTargetOptions model
				validateTargetOptionsModel := new(atrackerv2.ValidateTargetOptions)
				validateTargetOptionsModel.ID = core.StringPtr("testString")
				validateTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := atrackerService.ValidateTargetWithContext(ctx, validateTargetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				atrackerService.DisableRetries()
				result, response, operationErr := atrackerService.ValidateTarget(validateTargetOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = atrackerService.ValidateTargetWithContext(ctx, validateTargetOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(validateTargetPath))
					Expect(req.Method).To(Equal("POST"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "f7dcfae6-e7c5-08ca-451b-fdfa696c9bb6", "name": "a-cos-target-us-south", "crn": "crn:v1:bluemix:public:atracker:us-south:a/11111111111111111111111111111111:b6eec08b-5201-08ca-451b-cd71523e3626:target:f7dcfae6-e7c5-08ca-451b-fdfa696c9bb6", "target_type": "cloud_object_storage", "region": "us-south", "cos_endpoint": {"endpoint": "s3.private.us-east.cloud-object-storage.appdomain.cloud", "target_crn": "crn:v1:bluemix:public:cloud-object-storage:global:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::", "bucket": "my-atracker-bucket", "service_to_service_enabled": true}, "eventstreams_endpoint": {"target_crn": "crn:v1:bluemix:public:messagehub:us-south:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::", "brokers": ["kafka-x:9094"], "topic": "my-topic", "api_key": "xxxxxxxxxxxxxx", "service_to_service_enabled": false}, "cloudlogs_endpoint": {"target_crn": "crn:v1:bluemix:public:eu-es:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::"}, "write_status": {"status": "success", "last_failure": "2021-05-18T20:15:12.353Z", "reason_for_last_failure": "Provided API key could not be found"}, "created_at": "2021-05-18T20:15:12.353Z", "updated_at": "2021-05-18T20:15:12.353Z", "message": "This is a valid target. However, there is another target already defined with the same target endpoint.", "api_version": 2}`)
				}))
			})
			It(`Invoke ValidateTarget successfully`, func() {
				atrackerService, serviceErr := atrackerv2.NewAtrackerV2(&atrackerv2.AtrackerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := atrackerService.ValidateTarget(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ValidateTargetOptions model
				validateTargetOptionsModel := new(atrackerv2.ValidateTargetOptions)
				validateTargetOptionsModel.ID = core.StringPtr("testString")
				validateTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = atrackerService.ValidateTarget(validateTargetOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ValidateTarget with error: Operation validation and request error`, func() {
				atrackerService, serviceErr := atrackerv2.NewAtrackerV2(&atrackerv2.AtrackerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())

				// Construct an instance of the ValidateTargetOptions model
				validateTargetOptionsModel := new(atrackerv2.ValidateTargetOptions)
				validateTargetOptionsModel.ID = core.StringPtr("testString")
				validateTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := atrackerService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := atrackerService.ValidateTarget(validateTargetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ValidateTargetOptions model with no property values
				validateTargetOptionsModelNew := new(atrackerv2.ValidateTargetOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = atrackerService.ValidateTarget(validateTargetOptionsModelNew)
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
			It(`Invoke ValidateTarget successfully`, func() {
				atrackerService, serviceErr := atrackerv2.NewAtrackerV2(&atrackerv2.AtrackerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())

				// Construct an instance of the ValidateTargetOptions model
				validateTargetOptionsModel := new(atrackerv2.ValidateTargetOptions)
				validateTargetOptionsModel.ID = core.StringPtr("testString")
				validateTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := atrackerService.ValidateTarget(validateTargetOptionsModel)
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
	Describe(`CreateRoute(createRouteOptions *CreateRouteOptions) - Operation response error`, func() {
		createRoutePath := "/api/v2/routes"
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
				atrackerService, serviceErr := atrackerv2.NewAtrackerV2(&atrackerv2.AtrackerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())

				// Construct an instance of the RulePrototype model
				rulePrototypeModel := new(atrackerv2.RulePrototype)
				rulePrototypeModel.TargetIds = []string{"c3af557f-fb0e-4476-85c3-0889e7fe7bc4"}
				rulePrototypeModel.Locations = []string{"us-south"}

				// Construct an instance of the CreateRouteOptions model
				createRouteOptionsModel := new(atrackerv2.CreateRouteOptions)
				createRouteOptionsModel.Name = core.StringPtr("my-route")
				createRouteOptionsModel.Rules = []atrackerv2.RulePrototype{*rulePrototypeModel}
				createRouteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := atrackerService.CreateRoute(createRouteOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				atrackerService.EnableRetries(0, 0)
				result, response, operationErr = atrackerService.CreateRoute(createRouteOptionsModel)
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
		createRoutePath := "/api/v2/routes"
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
					fmt.Fprintf(res, "%s", `{"id": "c3af557f-fb0e-4476-85c3-0889e7fe7bc4", "name": "my-route", "crn": "crn:v1:bluemix:public:atracker:global:a/11111111111111111111111111111111:b6eec08b-5201-08ca-451b-cd71523e3626:route:c3af557f-fb0e-4476-85c3-0889e7fe7bc4", "version": 0, "rules": [{"target_ids": ["c3af557f-fb0e-4476-85c3-0889e7fe7bc4"], "locations": ["us-south"]}], "created_at": "2021-05-18T20:15:12.353Z", "updated_at": "2021-05-18T20:15:12.353Z", "api_version": 2, "message": "Route was created successfully."}`)
				}))
			})
			It(`Invoke CreateRoute successfully with retries`, func() {
				atrackerService, serviceErr := atrackerv2.NewAtrackerV2(&atrackerv2.AtrackerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())
				atrackerService.EnableRetries(0, 0)

				// Construct an instance of the RulePrototype model
				rulePrototypeModel := new(atrackerv2.RulePrototype)
				rulePrototypeModel.TargetIds = []string{"c3af557f-fb0e-4476-85c3-0889e7fe7bc4"}
				rulePrototypeModel.Locations = []string{"us-south"}

				// Construct an instance of the CreateRouteOptions model
				createRouteOptionsModel := new(atrackerv2.CreateRouteOptions)
				createRouteOptionsModel.Name = core.StringPtr("my-route")
				createRouteOptionsModel.Rules = []atrackerv2.RulePrototype{*rulePrototypeModel}
				createRouteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := atrackerService.CreateRouteWithContext(ctx, createRouteOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				atrackerService.DisableRetries()
				result, response, operationErr := atrackerService.CreateRoute(createRouteOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = atrackerService.CreateRouteWithContext(ctx, createRouteOptionsModel)
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
					fmt.Fprintf(res, "%s", `{"id": "c3af557f-fb0e-4476-85c3-0889e7fe7bc4", "name": "my-route", "crn": "crn:v1:bluemix:public:atracker:global:a/11111111111111111111111111111111:b6eec08b-5201-08ca-451b-cd71523e3626:route:c3af557f-fb0e-4476-85c3-0889e7fe7bc4", "version": 0, "rules": [{"target_ids": ["c3af557f-fb0e-4476-85c3-0889e7fe7bc4"], "locations": ["us-south"]}], "created_at": "2021-05-18T20:15:12.353Z", "updated_at": "2021-05-18T20:15:12.353Z", "api_version": 2, "message": "Route was created successfully."}`)
				}))
			})
			It(`Invoke CreateRoute successfully`, func() {
				atrackerService, serviceErr := atrackerv2.NewAtrackerV2(&atrackerv2.AtrackerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := atrackerService.CreateRoute(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the RulePrototype model
				rulePrototypeModel := new(atrackerv2.RulePrototype)
				rulePrototypeModel.TargetIds = []string{"c3af557f-fb0e-4476-85c3-0889e7fe7bc4"}
				rulePrototypeModel.Locations = []string{"us-south"}

				// Construct an instance of the CreateRouteOptions model
				createRouteOptionsModel := new(atrackerv2.CreateRouteOptions)
				createRouteOptionsModel.Name = core.StringPtr("my-route")
				createRouteOptionsModel.Rules = []atrackerv2.RulePrototype{*rulePrototypeModel}
				createRouteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = atrackerService.CreateRoute(createRouteOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateRoute with error: Operation validation and request error`, func() {
				atrackerService, serviceErr := atrackerv2.NewAtrackerV2(&atrackerv2.AtrackerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())

				// Construct an instance of the RulePrototype model
				rulePrototypeModel := new(atrackerv2.RulePrototype)
				rulePrototypeModel.TargetIds = []string{"c3af557f-fb0e-4476-85c3-0889e7fe7bc4"}
				rulePrototypeModel.Locations = []string{"us-south"}

				// Construct an instance of the CreateRouteOptions model
				createRouteOptionsModel := new(atrackerv2.CreateRouteOptions)
				createRouteOptionsModel.Name = core.StringPtr("my-route")
				createRouteOptionsModel.Rules = []atrackerv2.RulePrototype{*rulePrototypeModel}
				createRouteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := atrackerService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := atrackerService.CreateRoute(createRouteOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateRouteOptions model with no property values
				createRouteOptionsModelNew := new(atrackerv2.CreateRouteOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = atrackerService.CreateRoute(createRouteOptionsModelNew)
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
				atrackerService, serviceErr := atrackerv2.NewAtrackerV2(&atrackerv2.AtrackerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())

				// Construct an instance of the RulePrototype model
				rulePrototypeModel := new(atrackerv2.RulePrototype)
				rulePrototypeModel.TargetIds = []string{"c3af557f-fb0e-4476-85c3-0889e7fe7bc4"}
				rulePrototypeModel.Locations = []string{"us-south"}

				// Construct an instance of the CreateRouteOptions model
				createRouteOptionsModel := new(atrackerv2.CreateRouteOptions)
				createRouteOptionsModel.Name = core.StringPtr("my-route")
				createRouteOptionsModel.Rules = []atrackerv2.RulePrototype{*rulePrototypeModel}
				createRouteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := atrackerService.CreateRoute(createRouteOptionsModel)
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
		listRoutesPath := "/api/v2/routes"
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
				atrackerService, serviceErr := atrackerv2.NewAtrackerV2(&atrackerv2.AtrackerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())

				// Construct an instance of the ListRoutesOptions model
				listRoutesOptionsModel := new(atrackerv2.ListRoutesOptions)
				listRoutesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := atrackerService.ListRoutes(listRoutesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				atrackerService.EnableRetries(0, 0)
				result, response, operationErr = atrackerService.ListRoutes(listRoutesOptionsModel)
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
		listRoutesPath := "/api/v2/routes"
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
					fmt.Fprintf(res, "%s", `{"routes": [{"id": "c3af557f-fb0e-4476-85c3-0889e7fe7bc4", "name": "my-route", "crn": "crn:v1:bluemix:public:atracker:global:a/11111111111111111111111111111111:b6eec08b-5201-08ca-451b-cd71523e3626:route:c3af557f-fb0e-4476-85c3-0889e7fe7bc4", "version": 0, "rules": [{"target_ids": ["c3af557f-fb0e-4476-85c3-0889e7fe7bc4"], "locations": ["us-south"]}], "created_at": "2021-05-18T20:15:12.353Z", "updated_at": "2021-05-18T20:15:12.353Z", "api_version": 2, "message": "Route was created successfully."}]}`)
				}))
			})
			It(`Invoke ListRoutes successfully with retries`, func() {
				atrackerService, serviceErr := atrackerv2.NewAtrackerV2(&atrackerv2.AtrackerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())
				atrackerService.EnableRetries(0, 0)

				// Construct an instance of the ListRoutesOptions model
				listRoutesOptionsModel := new(atrackerv2.ListRoutesOptions)
				listRoutesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := atrackerService.ListRoutesWithContext(ctx, listRoutesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				atrackerService.DisableRetries()
				result, response, operationErr := atrackerService.ListRoutes(listRoutesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = atrackerService.ListRoutesWithContext(ctx, listRoutesOptionsModel)
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
					fmt.Fprintf(res, "%s", `{"routes": [{"id": "c3af557f-fb0e-4476-85c3-0889e7fe7bc4", "name": "my-route", "crn": "crn:v1:bluemix:public:atracker:global:a/11111111111111111111111111111111:b6eec08b-5201-08ca-451b-cd71523e3626:route:c3af557f-fb0e-4476-85c3-0889e7fe7bc4", "version": 0, "rules": [{"target_ids": ["c3af557f-fb0e-4476-85c3-0889e7fe7bc4"], "locations": ["us-south"]}], "created_at": "2021-05-18T20:15:12.353Z", "updated_at": "2021-05-18T20:15:12.353Z", "api_version": 2, "message": "Route was created successfully."}]}`)
				}))
			})
			It(`Invoke ListRoutes successfully`, func() {
				atrackerService, serviceErr := atrackerv2.NewAtrackerV2(&atrackerv2.AtrackerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := atrackerService.ListRoutes(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListRoutesOptions model
				listRoutesOptionsModel := new(atrackerv2.ListRoutesOptions)
				listRoutesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = atrackerService.ListRoutes(listRoutesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListRoutes with error: Operation request error`, func() {
				atrackerService, serviceErr := atrackerv2.NewAtrackerV2(&atrackerv2.AtrackerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())

				// Construct an instance of the ListRoutesOptions model
				listRoutesOptionsModel := new(atrackerv2.ListRoutesOptions)
				listRoutesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := atrackerService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := atrackerService.ListRoutes(listRoutesOptionsModel)
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
				atrackerService, serviceErr := atrackerv2.NewAtrackerV2(&atrackerv2.AtrackerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())

				// Construct an instance of the ListRoutesOptions model
				listRoutesOptionsModel := new(atrackerv2.ListRoutesOptions)
				listRoutesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := atrackerService.ListRoutes(listRoutesOptionsModel)
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
		getRoutePath := "/api/v2/routes/testString"
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
				atrackerService, serviceErr := atrackerv2.NewAtrackerV2(&atrackerv2.AtrackerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())

				// Construct an instance of the GetRouteOptions model
				getRouteOptionsModel := new(atrackerv2.GetRouteOptions)
				getRouteOptionsModel.ID = core.StringPtr("testString")
				getRouteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := atrackerService.GetRoute(getRouteOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				atrackerService.EnableRetries(0, 0)
				result, response, operationErr = atrackerService.GetRoute(getRouteOptionsModel)
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
		getRoutePath := "/api/v2/routes/testString"
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
					fmt.Fprintf(res, "%s", `{"id": "c3af557f-fb0e-4476-85c3-0889e7fe7bc4", "name": "my-route", "crn": "crn:v1:bluemix:public:atracker:global:a/11111111111111111111111111111111:b6eec08b-5201-08ca-451b-cd71523e3626:route:c3af557f-fb0e-4476-85c3-0889e7fe7bc4", "version": 0, "rules": [{"target_ids": ["c3af557f-fb0e-4476-85c3-0889e7fe7bc4"], "locations": ["us-south"]}], "created_at": "2021-05-18T20:15:12.353Z", "updated_at": "2021-05-18T20:15:12.353Z", "api_version": 2, "message": "Route was created successfully."}`)
				}))
			})
			It(`Invoke GetRoute successfully with retries`, func() {
				atrackerService, serviceErr := atrackerv2.NewAtrackerV2(&atrackerv2.AtrackerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())
				atrackerService.EnableRetries(0, 0)

				// Construct an instance of the GetRouteOptions model
				getRouteOptionsModel := new(atrackerv2.GetRouteOptions)
				getRouteOptionsModel.ID = core.StringPtr("testString")
				getRouteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := atrackerService.GetRouteWithContext(ctx, getRouteOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				atrackerService.DisableRetries()
				result, response, operationErr := atrackerService.GetRoute(getRouteOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = atrackerService.GetRouteWithContext(ctx, getRouteOptionsModel)
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
					fmt.Fprintf(res, "%s", `{"id": "c3af557f-fb0e-4476-85c3-0889e7fe7bc4", "name": "my-route", "crn": "crn:v1:bluemix:public:atracker:global:a/11111111111111111111111111111111:b6eec08b-5201-08ca-451b-cd71523e3626:route:c3af557f-fb0e-4476-85c3-0889e7fe7bc4", "version": 0, "rules": [{"target_ids": ["c3af557f-fb0e-4476-85c3-0889e7fe7bc4"], "locations": ["us-south"]}], "created_at": "2021-05-18T20:15:12.353Z", "updated_at": "2021-05-18T20:15:12.353Z", "api_version": 2, "message": "Route was created successfully."}`)
				}))
			})
			It(`Invoke GetRoute successfully`, func() {
				atrackerService, serviceErr := atrackerv2.NewAtrackerV2(&atrackerv2.AtrackerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := atrackerService.GetRoute(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetRouteOptions model
				getRouteOptionsModel := new(atrackerv2.GetRouteOptions)
				getRouteOptionsModel.ID = core.StringPtr("testString")
				getRouteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = atrackerService.GetRoute(getRouteOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetRoute with error: Operation validation and request error`, func() {
				atrackerService, serviceErr := atrackerv2.NewAtrackerV2(&atrackerv2.AtrackerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())

				// Construct an instance of the GetRouteOptions model
				getRouteOptionsModel := new(atrackerv2.GetRouteOptions)
				getRouteOptionsModel.ID = core.StringPtr("testString")
				getRouteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := atrackerService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := atrackerService.GetRoute(getRouteOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetRouteOptions model with no property values
				getRouteOptionsModelNew := new(atrackerv2.GetRouteOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = atrackerService.GetRoute(getRouteOptionsModelNew)
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
				atrackerService, serviceErr := atrackerv2.NewAtrackerV2(&atrackerv2.AtrackerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())

				// Construct an instance of the GetRouteOptions model
				getRouteOptionsModel := new(atrackerv2.GetRouteOptions)
				getRouteOptionsModel.ID = core.StringPtr("testString")
				getRouteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := atrackerService.GetRoute(getRouteOptionsModel)
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
	Describe(`ReplaceRoute(replaceRouteOptions *ReplaceRouteOptions) - Operation response error`, func() {
		replaceRoutePath := "/api/v2/routes/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceRoutePath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ReplaceRoute with error: Operation response processing error`, func() {
				atrackerService, serviceErr := atrackerv2.NewAtrackerV2(&atrackerv2.AtrackerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())

				// Construct an instance of the RulePrototype model
				rulePrototypeModel := new(atrackerv2.RulePrototype)
				rulePrototypeModel.TargetIds = []string{"c3af557f-fb0e-4476-85c3-0889e7fe7bc4"}
				rulePrototypeModel.Locations = []string{"us-south"}

				// Construct an instance of the ReplaceRouteOptions model
				replaceRouteOptionsModel := new(atrackerv2.ReplaceRouteOptions)
				replaceRouteOptionsModel.ID = core.StringPtr("testString")
				replaceRouteOptionsModel.Name = core.StringPtr("my-route")
				replaceRouteOptionsModel.Rules = []atrackerv2.RulePrototype{*rulePrototypeModel}
				replaceRouteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := atrackerService.ReplaceRoute(replaceRouteOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				atrackerService.EnableRetries(0, 0)
				result, response, operationErr = atrackerService.ReplaceRoute(replaceRouteOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ReplaceRoute(replaceRouteOptions *ReplaceRouteOptions)`, func() {
		replaceRoutePath := "/api/v2/routes/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceRoutePath))
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
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "c3af557f-fb0e-4476-85c3-0889e7fe7bc4", "name": "my-route", "crn": "crn:v1:bluemix:public:atracker:global:a/11111111111111111111111111111111:b6eec08b-5201-08ca-451b-cd71523e3626:route:c3af557f-fb0e-4476-85c3-0889e7fe7bc4", "version": 0, "rules": [{"target_ids": ["c3af557f-fb0e-4476-85c3-0889e7fe7bc4"], "locations": ["us-south"]}], "created_at": "2021-05-18T20:15:12.353Z", "updated_at": "2021-05-18T20:15:12.353Z", "api_version": 2, "message": "Route was created successfully."}`)
				}))
			})
			It(`Invoke ReplaceRoute successfully with retries`, func() {
				atrackerService, serviceErr := atrackerv2.NewAtrackerV2(&atrackerv2.AtrackerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())
				atrackerService.EnableRetries(0, 0)

				// Construct an instance of the RulePrototype model
				rulePrototypeModel := new(atrackerv2.RulePrototype)
				rulePrototypeModel.TargetIds = []string{"c3af557f-fb0e-4476-85c3-0889e7fe7bc4"}
				rulePrototypeModel.Locations = []string{"us-south"}

				// Construct an instance of the ReplaceRouteOptions model
				replaceRouteOptionsModel := new(atrackerv2.ReplaceRouteOptions)
				replaceRouteOptionsModel.ID = core.StringPtr("testString")
				replaceRouteOptionsModel.Name = core.StringPtr("my-route")
				replaceRouteOptionsModel.Rules = []atrackerv2.RulePrototype{*rulePrototypeModel}
				replaceRouteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := atrackerService.ReplaceRouteWithContext(ctx, replaceRouteOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				atrackerService.DisableRetries()
				result, response, operationErr := atrackerService.ReplaceRoute(replaceRouteOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = atrackerService.ReplaceRouteWithContext(ctx, replaceRouteOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(replaceRoutePath))
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

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "c3af557f-fb0e-4476-85c3-0889e7fe7bc4", "name": "my-route", "crn": "crn:v1:bluemix:public:atracker:global:a/11111111111111111111111111111111:b6eec08b-5201-08ca-451b-cd71523e3626:route:c3af557f-fb0e-4476-85c3-0889e7fe7bc4", "version": 0, "rules": [{"target_ids": ["c3af557f-fb0e-4476-85c3-0889e7fe7bc4"], "locations": ["us-south"]}], "created_at": "2021-05-18T20:15:12.353Z", "updated_at": "2021-05-18T20:15:12.353Z", "api_version": 2, "message": "Route was created successfully."}`)
				}))
			})
			It(`Invoke ReplaceRoute successfully`, func() {
				atrackerService, serviceErr := atrackerv2.NewAtrackerV2(&atrackerv2.AtrackerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := atrackerService.ReplaceRoute(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the RulePrototype model
				rulePrototypeModel := new(atrackerv2.RulePrototype)
				rulePrototypeModel.TargetIds = []string{"c3af557f-fb0e-4476-85c3-0889e7fe7bc4"}
				rulePrototypeModel.Locations = []string{"us-south"}

				// Construct an instance of the ReplaceRouteOptions model
				replaceRouteOptionsModel := new(atrackerv2.ReplaceRouteOptions)
				replaceRouteOptionsModel.ID = core.StringPtr("testString")
				replaceRouteOptionsModel.Name = core.StringPtr("my-route")
				replaceRouteOptionsModel.Rules = []atrackerv2.RulePrototype{*rulePrototypeModel}
				replaceRouteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = atrackerService.ReplaceRoute(replaceRouteOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ReplaceRoute with error: Operation validation and request error`, func() {
				atrackerService, serviceErr := atrackerv2.NewAtrackerV2(&atrackerv2.AtrackerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())

				// Construct an instance of the RulePrototype model
				rulePrototypeModel := new(atrackerv2.RulePrototype)
				rulePrototypeModel.TargetIds = []string{"c3af557f-fb0e-4476-85c3-0889e7fe7bc4"}
				rulePrototypeModel.Locations = []string{"us-south"}

				// Construct an instance of the ReplaceRouteOptions model
				replaceRouteOptionsModel := new(atrackerv2.ReplaceRouteOptions)
				replaceRouteOptionsModel.ID = core.StringPtr("testString")
				replaceRouteOptionsModel.Name = core.StringPtr("my-route")
				replaceRouteOptionsModel.Rules = []atrackerv2.RulePrototype{*rulePrototypeModel}
				replaceRouteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := atrackerService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := atrackerService.ReplaceRoute(replaceRouteOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ReplaceRouteOptions model with no property values
				replaceRouteOptionsModelNew := new(atrackerv2.ReplaceRouteOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = atrackerService.ReplaceRoute(replaceRouteOptionsModelNew)
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
			It(`Invoke ReplaceRoute successfully`, func() {
				atrackerService, serviceErr := atrackerv2.NewAtrackerV2(&atrackerv2.AtrackerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())

				// Construct an instance of the RulePrototype model
				rulePrototypeModel := new(atrackerv2.RulePrototype)
				rulePrototypeModel.TargetIds = []string{"c3af557f-fb0e-4476-85c3-0889e7fe7bc4"}
				rulePrototypeModel.Locations = []string{"us-south"}

				// Construct an instance of the ReplaceRouteOptions model
				replaceRouteOptionsModel := new(atrackerv2.ReplaceRouteOptions)
				replaceRouteOptionsModel.ID = core.StringPtr("testString")
				replaceRouteOptionsModel.Name = core.StringPtr("my-route")
				replaceRouteOptionsModel.Rules = []atrackerv2.RulePrototype{*rulePrototypeModel}
				replaceRouteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := atrackerService.ReplaceRoute(replaceRouteOptionsModel)
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
		deleteRoutePath := "/api/v2/routes/testString"
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
				atrackerService, serviceErr := atrackerv2.NewAtrackerV2(&atrackerv2.AtrackerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := atrackerService.DeleteRoute(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteRouteOptions model
				deleteRouteOptionsModel := new(atrackerv2.DeleteRouteOptions)
				deleteRouteOptionsModel.ID = core.StringPtr("testString")
				deleteRouteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = atrackerService.DeleteRoute(deleteRouteOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteRoute with error: Operation validation and request error`, func() {
				atrackerService, serviceErr := atrackerv2.NewAtrackerV2(&atrackerv2.AtrackerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())

				// Construct an instance of the DeleteRouteOptions model
				deleteRouteOptionsModel := new(atrackerv2.DeleteRouteOptions)
				deleteRouteOptionsModel.ID = core.StringPtr("testString")
				deleteRouteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := atrackerService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := atrackerService.DeleteRoute(deleteRouteOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteRouteOptions model with no property values
				deleteRouteOptionsModelNew := new(atrackerv2.DeleteRouteOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = atrackerService.DeleteRoute(deleteRouteOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetSettings(getSettingsOptions *GetSettingsOptions) - Operation response error`, func() {
		getSettingsPath := "/api/v2/settings"
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
				atrackerService, serviceErr := atrackerv2.NewAtrackerV2(&atrackerv2.AtrackerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())

				// Construct an instance of the GetSettingsOptions model
				getSettingsOptionsModel := new(atrackerv2.GetSettingsOptions)
				getSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := atrackerService.GetSettings(getSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				atrackerService.EnableRetries(0, 0)
				result, response, operationErr = atrackerService.GetSettings(getSettingsOptionsModel)
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
		getSettingsPath := "/api/v2/settings"
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
					fmt.Fprintf(res, "%s", `{"default_targets": ["c3af557f-fb0e-4476-85c3-0889e7fe7bc4"], "permitted_target_regions": ["us-south"], "metadata_region_primary": "us-south", "metadata_region_backup": "eu-de", "private_api_endpoint_only": false, "api_version": 2, "message": "The route and target audit logs can be found in the metadata primary region and everything else can be found in the region it is being called from."}`)
				}))
			})
			It(`Invoke GetSettings successfully with retries`, func() {
				atrackerService, serviceErr := atrackerv2.NewAtrackerV2(&atrackerv2.AtrackerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())
				atrackerService.EnableRetries(0, 0)

				// Construct an instance of the GetSettingsOptions model
				getSettingsOptionsModel := new(atrackerv2.GetSettingsOptions)
				getSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := atrackerService.GetSettingsWithContext(ctx, getSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				atrackerService.DisableRetries()
				result, response, operationErr := atrackerService.GetSettings(getSettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = atrackerService.GetSettingsWithContext(ctx, getSettingsOptionsModel)
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
					fmt.Fprintf(res, "%s", `{"default_targets": ["c3af557f-fb0e-4476-85c3-0889e7fe7bc4"], "permitted_target_regions": ["us-south"], "metadata_region_primary": "us-south", "metadata_region_backup": "eu-de", "private_api_endpoint_only": false, "api_version": 2, "message": "The route and target audit logs can be found in the metadata primary region and everything else can be found in the region it is being called from."}`)
				}))
			})
			It(`Invoke GetSettings successfully`, func() {
				atrackerService, serviceErr := atrackerv2.NewAtrackerV2(&atrackerv2.AtrackerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := atrackerService.GetSettings(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetSettingsOptions model
				getSettingsOptionsModel := new(atrackerv2.GetSettingsOptions)
				getSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = atrackerService.GetSettings(getSettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetSettings with error: Operation request error`, func() {
				atrackerService, serviceErr := atrackerv2.NewAtrackerV2(&atrackerv2.AtrackerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())

				// Construct an instance of the GetSettingsOptions model
				getSettingsOptionsModel := new(atrackerv2.GetSettingsOptions)
				getSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := atrackerService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := atrackerService.GetSettings(getSettingsOptionsModel)
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
				atrackerService, serviceErr := atrackerv2.NewAtrackerV2(&atrackerv2.AtrackerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())

				// Construct an instance of the GetSettingsOptions model
				getSettingsOptionsModel := new(atrackerv2.GetSettingsOptions)
				getSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := atrackerService.GetSettings(getSettingsOptionsModel)
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
	Describe(`PutSettings(putSettingsOptions *PutSettingsOptions) - Operation response error`, func() {
		putSettingsPath := "/api/v2/settings"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(putSettingsPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke PutSettings with error: Operation response processing error`, func() {
				atrackerService, serviceErr := atrackerv2.NewAtrackerV2(&atrackerv2.AtrackerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())

				// Construct an instance of the PutSettingsOptions model
				putSettingsOptionsModel := new(atrackerv2.PutSettingsOptions)
				putSettingsOptionsModel.MetadataRegionPrimary = core.StringPtr("us-south")
				putSettingsOptionsModel.PrivateAPIEndpointOnly = core.BoolPtr(false)
				putSettingsOptionsModel.DefaultTargets = []string{"c3af557f-fb0e-4476-85c3-0889e7fe7bc4"}
				putSettingsOptionsModel.PermittedTargetRegions = []string{"us-south"}
				putSettingsOptionsModel.MetadataRegionBackup = core.StringPtr("eu-de")
				putSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := atrackerService.PutSettings(putSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				atrackerService.EnableRetries(0, 0)
				result, response, operationErr = atrackerService.PutSettings(putSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`PutSettings(putSettingsOptions *PutSettingsOptions)`, func() {
		putSettingsPath := "/api/v2/settings"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(putSettingsPath))
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
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"default_targets": ["c3af557f-fb0e-4476-85c3-0889e7fe7bc4"], "permitted_target_regions": ["us-south"], "metadata_region_primary": "us-south", "metadata_region_backup": "eu-de", "private_api_endpoint_only": false, "api_version": 2, "message": "The route and target audit logs can be found in the metadata primary region and everything else can be found in the region it is being called from."}`)
				}))
			})
			It(`Invoke PutSettings successfully with retries`, func() {
				atrackerService, serviceErr := atrackerv2.NewAtrackerV2(&atrackerv2.AtrackerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())
				atrackerService.EnableRetries(0, 0)

				// Construct an instance of the PutSettingsOptions model
				putSettingsOptionsModel := new(atrackerv2.PutSettingsOptions)
				putSettingsOptionsModel.MetadataRegionPrimary = core.StringPtr("us-south")
				putSettingsOptionsModel.PrivateAPIEndpointOnly = core.BoolPtr(false)
				putSettingsOptionsModel.DefaultTargets = []string{"c3af557f-fb0e-4476-85c3-0889e7fe7bc4"}
				putSettingsOptionsModel.PermittedTargetRegions = []string{"us-south"}
				putSettingsOptionsModel.MetadataRegionBackup = core.StringPtr("eu-de")
				putSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := atrackerService.PutSettingsWithContext(ctx, putSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				atrackerService.DisableRetries()
				result, response, operationErr := atrackerService.PutSettings(putSettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = atrackerService.PutSettingsWithContext(ctx, putSettingsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(putSettingsPath))
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

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"default_targets": ["c3af557f-fb0e-4476-85c3-0889e7fe7bc4"], "permitted_target_regions": ["us-south"], "metadata_region_primary": "us-south", "metadata_region_backup": "eu-de", "private_api_endpoint_only": false, "api_version": 2, "message": "The route and target audit logs can be found in the metadata primary region and everything else can be found in the region it is being called from."}`)
				}))
			})
			It(`Invoke PutSettings successfully`, func() {
				atrackerService, serviceErr := atrackerv2.NewAtrackerV2(&atrackerv2.AtrackerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := atrackerService.PutSettings(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the PutSettingsOptions model
				putSettingsOptionsModel := new(atrackerv2.PutSettingsOptions)
				putSettingsOptionsModel.MetadataRegionPrimary = core.StringPtr("us-south")
				putSettingsOptionsModel.PrivateAPIEndpointOnly = core.BoolPtr(false)
				putSettingsOptionsModel.DefaultTargets = []string{"c3af557f-fb0e-4476-85c3-0889e7fe7bc4"}
				putSettingsOptionsModel.PermittedTargetRegions = []string{"us-south"}
				putSettingsOptionsModel.MetadataRegionBackup = core.StringPtr("eu-de")
				putSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = atrackerService.PutSettings(putSettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke PutSettings with error: Operation validation and request error`, func() {
				atrackerService, serviceErr := atrackerv2.NewAtrackerV2(&atrackerv2.AtrackerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())

				// Construct an instance of the PutSettingsOptions model
				putSettingsOptionsModel := new(atrackerv2.PutSettingsOptions)
				putSettingsOptionsModel.MetadataRegionPrimary = core.StringPtr("us-south")
				putSettingsOptionsModel.PrivateAPIEndpointOnly = core.BoolPtr(false)
				putSettingsOptionsModel.DefaultTargets = []string{"c3af557f-fb0e-4476-85c3-0889e7fe7bc4"}
				putSettingsOptionsModel.PermittedTargetRegions = []string{"us-south"}
				putSettingsOptionsModel.MetadataRegionBackup = core.StringPtr("eu-de")
				putSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := atrackerService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := atrackerService.PutSettings(putSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the PutSettingsOptions model with no property values
				putSettingsOptionsModelNew := new(atrackerv2.PutSettingsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = atrackerService.PutSettings(putSettingsOptionsModelNew)
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
			It(`Invoke PutSettings successfully`, func() {
				atrackerService, serviceErr := atrackerv2.NewAtrackerV2(&atrackerv2.AtrackerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())

				// Construct an instance of the PutSettingsOptions model
				putSettingsOptionsModel := new(atrackerv2.PutSettingsOptions)
				putSettingsOptionsModel.MetadataRegionPrimary = core.StringPtr("us-south")
				putSettingsOptionsModel.PrivateAPIEndpointOnly = core.BoolPtr(false)
				putSettingsOptionsModel.DefaultTargets = []string{"c3af557f-fb0e-4476-85c3-0889e7fe7bc4"}
				putSettingsOptionsModel.PermittedTargetRegions = []string{"us-south"}
				putSettingsOptionsModel.MetadataRegionBackup = core.StringPtr("eu-de")
				putSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := atrackerService.PutSettings(putSettingsOptionsModel)
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
			atrackerService, _ := atrackerv2.NewAtrackerV2(&atrackerv2.AtrackerV2Options{
				URL:           "http://atrackerv2modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewCloudLogsEndpointPrototype successfully`, func() {
				targetCRN := "crn:v1:bluemix:public:logs:eu-es:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::"
				_model, err := atrackerService.NewCloudLogsEndpointPrototype(targetCRN)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewCosEndpointPrototype successfully`, func() {
				endpoint := "s3.private.us-east.cloud-object-storage.appdomain.cloud"
				targetCRN := "crn:v1:bluemix:public:cloud-object-storage:global:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::"
				bucket := "my-atracker-bucket"
				_model, err := atrackerService.NewCosEndpointPrototype(endpoint, targetCRN, bucket)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewCreateRouteOptions successfully`, func() {
				// Construct an instance of the RulePrototype model
				rulePrototypeModel := new(atrackerv2.RulePrototype)
				Expect(rulePrototypeModel).ToNot(BeNil())
				rulePrototypeModel.TargetIds = []string{"c3af557f-fb0e-4476-85c3-0889e7fe7bc4"}
				rulePrototypeModel.Locations = []string{"us-south"}
				Expect(rulePrototypeModel.TargetIds).To(Equal([]string{"c3af557f-fb0e-4476-85c3-0889e7fe7bc4"}))
				Expect(rulePrototypeModel.Locations).To(Equal([]string{"us-south"}))

				// Construct an instance of the CreateRouteOptions model
				createRouteOptionsName := "my-route"
				createRouteOptionsRules := []atrackerv2.RulePrototype{}
				createRouteOptionsModel := atrackerService.NewCreateRouteOptions(createRouteOptionsName, createRouteOptionsRules)
				createRouteOptionsModel.SetName("my-route")
				createRouteOptionsModel.SetRules([]atrackerv2.RulePrototype{*rulePrototypeModel})
				createRouteOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createRouteOptionsModel).ToNot(BeNil())
				Expect(createRouteOptionsModel.Name).To(Equal(core.StringPtr("my-route")))
				Expect(createRouteOptionsModel.Rules).To(Equal([]atrackerv2.RulePrototype{*rulePrototypeModel}))
				Expect(createRouteOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateTargetOptions successfully`, func() {
				// Construct an instance of the CosEndpointPrototype model
				cosEndpointPrototypeModel := new(atrackerv2.CosEndpointPrototype)
				Expect(cosEndpointPrototypeModel).ToNot(BeNil())
				cosEndpointPrototypeModel.Endpoint = core.StringPtr("s3.private.us-east.cloud-object-storage.appdomain.cloud")
				cosEndpointPrototypeModel.TargetCRN = core.StringPtr("crn:v1:bluemix:public:cloud-object-storage:global:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::")
				cosEndpointPrototypeModel.Bucket = core.StringPtr("my-atracker-bucket")
				cosEndpointPrototypeModel.APIKey = core.StringPtr("xxxxxxxxxxxxxx")
				cosEndpointPrototypeModel.ServiceToServiceEnabled = core.BoolPtr(true)
				Expect(cosEndpointPrototypeModel.Endpoint).To(Equal(core.StringPtr("s3.private.us-east.cloud-object-storage.appdomain.cloud")))
				Expect(cosEndpointPrototypeModel.TargetCRN).To(Equal(core.StringPtr("crn:v1:bluemix:public:cloud-object-storage:global:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::")))
				Expect(cosEndpointPrototypeModel.Bucket).To(Equal(core.StringPtr("my-atracker-bucket")))
				Expect(cosEndpointPrototypeModel.APIKey).To(Equal(core.StringPtr("xxxxxxxxxxxxxx")))
				Expect(cosEndpointPrototypeModel.ServiceToServiceEnabled).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the EventstreamsEndpointPrototype model
				eventstreamsEndpointPrototypeModel := new(atrackerv2.EventstreamsEndpointPrototype)
				Expect(eventstreamsEndpointPrototypeModel).ToNot(BeNil())
				eventstreamsEndpointPrototypeModel.TargetCRN = core.StringPtr("crn:v1:bluemix:public:messagehub:us-south:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::")
				eventstreamsEndpointPrototypeModel.Brokers = []string{"kafka-x:9094"}
				eventstreamsEndpointPrototypeModel.Topic = core.StringPtr("my-topic")
				eventstreamsEndpointPrototypeModel.APIKey = core.StringPtr("xxxxxxxxxxxxxx")
				eventstreamsEndpointPrototypeModel.ServiceToServiceEnabled = core.BoolPtr(false)
				Expect(eventstreamsEndpointPrototypeModel.TargetCRN).To(Equal(core.StringPtr("crn:v1:bluemix:public:messagehub:us-south:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::")))
				Expect(eventstreamsEndpointPrototypeModel.Brokers).To(Equal([]string{"kafka-x:9094"}))
				Expect(eventstreamsEndpointPrototypeModel.Topic).To(Equal(core.StringPtr("my-topic")))
				Expect(eventstreamsEndpointPrototypeModel.APIKey).To(Equal(core.StringPtr("xxxxxxxxxxxxxx")))
				Expect(eventstreamsEndpointPrototypeModel.ServiceToServiceEnabled).To(Equal(core.BoolPtr(false)))

				// Construct an instance of the CloudLogsEndpointPrototype model
				cloudLogsEndpointPrototypeModel := new(atrackerv2.CloudLogsEndpointPrototype)
				Expect(cloudLogsEndpointPrototypeModel).ToNot(BeNil())
				cloudLogsEndpointPrototypeModel.TargetCRN = core.StringPtr("crn:v1:bluemix:public:logs:eu-es:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::")
				Expect(cloudLogsEndpointPrototypeModel.TargetCRN).To(Equal(core.StringPtr("crn:v1:bluemix:public:logs:eu-es:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::")))

				// Construct an instance of the CreateTargetOptions model
				createTargetOptionsName := "my-cos-target"
				createTargetOptionsTargetType := "cloud_object_storage"
				createTargetOptionsModel := atrackerService.NewCreateTargetOptions(createTargetOptionsName, createTargetOptionsTargetType)
				createTargetOptionsModel.SetName("my-cos-target")
				createTargetOptionsModel.SetTargetType("cloud_object_storage")
				createTargetOptionsModel.SetCosEndpoint(cosEndpointPrototypeModel)
				createTargetOptionsModel.SetEventstreamsEndpoint(eventstreamsEndpointPrototypeModel)
				createTargetOptionsModel.SetCloudlogsEndpoint(cloudLogsEndpointPrototypeModel)
				createTargetOptionsModel.SetRegion("us-south")
				createTargetOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createTargetOptionsModel).ToNot(BeNil())
				Expect(createTargetOptionsModel.Name).To(Equal(core.StringPtr("my-cos-target")))
				Expect(createTargetOptionsModel.TargetType).To(Equal(core.StringPtr("cloud_object_storage")))
				Expect(createTargetOptionsModel.CosEndpoint).To(Equal(cosEndpointPrototypeModel))
				Expect(createTargetOptionsModel.EventstreamsEndpoint).To(Equal(eventstreamsEndpointPrototypeModel))
				Expect(createTargetOptionsModel.CloudlogsEndpoint).To(Equal(cloudLogsEndpointPrototypeModel))
				Expect(createTargetOptionsModel.Region).To(Equal(core.StringPtr("us-south")))
				Expect(createTargetOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteRouteOptions successfully`, func() {
				// Construct an instance of the DeleteRouteOptions model
				id := "testString"
				deleteRouteOptionsModel := atrackerService.NewDeleteRouteOptions(id)
				deleteRouteOptionsModel.SetID("testString")
				deleteRouteOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteRouteOptionsModel).ToNot(BeNil())
				Expect(deleteRouteOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(deleteRouteOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteTargetOptions successfully`, func() {
				// Construct an instance of the DeleteTargetOptions model
				id := "testString"
				deleteTargetOptionsModel := atrackerService.NewDeleteTargetOptions(id)
				deleteTargetOptionsModel.SetID("testString")
				deleteTargetOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteTargetOptionsModel).ToNot(BeNil())
				Expect(deleteTargetOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(deleteTargetOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewEventstreamsEndpointPrototype successfully`, func() {
				targetCRN := "crn:v1:bluemix:public:messagehub:us-south:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::"
				brokers := []string{"kafka-x:9094"}
				topic := "my-topic"
				_model, err := atrackerService.NewEventstreamsEndpointPrototype(targetCRN, brokers, topic)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewGetRouteOptions successfully`, func() {
				// Construct an instance of the GetRouteOptions model
				id := "testString"
				getRouteOptionsModel := atrackerService.NewGetRouteOptions(id)
				getRouteOptionsModel.SetID("testString")
				getRouteOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getRouteOptionsModel).ToNot(BeNil())
				Expect(getRouteOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getRouteOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetSettingsOptions successfully`, func() {
				// Construct an instance of the GetSettingsOptions model
				getSettingsOptionsModel := atrackerService.NewGetSettingsOptions()
				getSettingsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getSettingsOptionsModel).ToNot(BeNil())
				Expect(getSettingsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetTargetOptions successfully`, func() {
				// Construct an instance of the GetTargetOptions model
				id := "testString"
				getTargetOptionsModel := atrackerService.NewGetTargetOptions(id)
				getTargetOptionsModel.SetID("testString")
				getTargetOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getTargetOptionsModel).ToNot(BeNil())
				Expect(getTargetOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getTargetOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListRoutesOptions successfully`, func() {
				// Construct an instance of the ListRoutesOptions model
				listRoutesOptionsModel := atrackerService.NewListRoutesOptions()
				listRoutesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listRoutesOptionsModel).ToNot(BeNil())
				Expect(listRoutesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListTargetsOptions successfully`, func() {
				// Construct an instance of the ListTargetsOptions model
				listTargetsOptionsModel := atrackerService.NewListTargetsOptions()
				listTargetsOptionsModel.SetRegion("testString")
				listTargetsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listTargetsOptionsModel).ToNot(BeNil())
				Expect(listTargetsOptionsModel.Region).To(Equal(core.StringPtr("testString")))
				Expect(listTargetsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPutSettingsOptions successfully`, func() {
				// Construct an instance of the PutSettingsOptions model
				putSettingsOptionsMetadataRegionPrimary := "us-south"
				putSettingsOptionsPrivateAPIEndpointOnly := false
				putSettingsOptionsModel := atrackerService.NewPutSettingsOptions(putSettingsOptionsMetadataRegionPrimary, putSettingsOptionsPrivateAPIEndpointOnly)
				putSettingsOptionsModel.SetMetadataRegionPrimary("us-south")
				putSettingsOptionsModel.SetPrivateAPIEndpointOnly(false)
				putSettingsOptionsModel.SetDefaultTargets([]string{"c3af557f-fb0e-4476-85c3-0889e7fe7bc4"})
				putSettingsOptionsModel.SetPermittedTargetRegions([]string{"us-south"})
				putSettingsOptionsModel.SetMetadataRegionBackup("eu-de")
				putSettingsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(putSettingsOptionsModel).ToNot(BeNil())
				Expect(putSettingsOptionsModel.MetadataRegionPrimary).To(Equal(core.StringPtr("us-south")))
				Expect(putSettingsOptionsModel.PrivateAPIEndpointOnly).To(Equal(core.BoolPtr(false)))
				Expect(putSettingsOptionsModel.DefaultTargets).To(Equal([]string{"c3af557f-fb0e-4476-85c3-0889e7fe7bc4"}))
				Expect(putSettingsOptionsModel.PermittedTargetRegions).To(Equal([]string{"us-south"}))
				Expect(putSettingsOptionsModel.MetadataRegionBackup).To(Equal(core.StringPtr("eu-de")))
				Expect(putSettingsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewReplaceRouteOptions successfully`, func() {
				// Construct an instance of the RulePrototype model
				rulePrototypeModel := new(atrackerv2.RulePrototype)
				Expect(rulePrototypeModel).ToNot(BeNil())
				rulePrototypeModel.TargetIds = []string{"c3af557f-fb0e-4476-85c3-0889e7fe7bc4"}
				rulePrototypeModel.Locations = []string{"us-south"}
				Expect(rulePrototypeModel.TargetIds).To(Equal([]string{"c3af557f-fb0e-4476-85c3-0889e7fe7bc4"}))
				Expect(rulePrototypeModel.Locations).To(Equal([]string{"us-south"}))

				// Construct an instance of the ReplaceRouteOptions model
				id := "testString"
				replaceRouteOptionsName := "my-route"
				replaceRouteOptionsRules := []atrackerv2.RulePrototype{}
				replaceRouteOptionsModel := atrackerService.NewReplaceRouteOptions(id, replaceRouteOptionsName, replaceRouteOptionsRules)
				replaceRouteOptionsModel.SetID("testString")
				replaceRouteOptionsModel.SetName("my-route")
				replaceRouteOptionsModel.SetRules([]atrackerv2.RulePrototype{*rulePrototypeModel})
				replaceRouteOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(replaceRouteOptionsModel).ToNot(BeNil())
				Expect(replaceRouteOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(replaceRouteOptionsModel.Name).To(Equal(core.StringPtr("my-route")))
				Expect(replaceRouteOptionsModel.Rules).To(Equal([]atrackerv2.RulePrototype{*rulePrototypeModel}))
				Expect(replaceRouteOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewReplaceTargetOptions successfully`, func() {
				// Construct an instance of the CosEndpointPrototype model
				cosEndpointPrototypeModel := new(atrackerv2.CosEndpointPrototype)
				Expect(cosEndpointPrototypeModel).ToNot(BeNil())
				cosEndpointPrototypeModel.Endpoint = core.StringPtr("s3.private.us-east.cloud-object-storage.appdomain.cloud")
				cosEndpointPrototypeModel.TargetCRN = core.StringPtr("crn:v1:bluemix:public:cloud-object-storage:global:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::")
				cosEndpointPrototypeModel.Bucket = core.StringPtr("my-atracker-bucket")
				cosEndpointPrototypeModel.APIKey = core.StringPtr("xxxxxxxxxxxxxx")
				cosEndpointPrototypeModel.ServiceToServiceEnabled = core.BoolPtr(true)
				Expect(cosEndpointPrototypeModel.Endpoint).To(Equal(core.StringPtr("s3.private.us-east.cloud-object-storage.appdomain.cloud")))
				Expect(cosEndpointPrototypeModel.TargetCRN).To(Equal(core.StringPtr("crn:v1:bluemix:public:cloud-object-storage:global:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::")))
				Expect(cosEndpointPrototypeModel.Bucket).To(Equal(core.StringPtr("my-atracker-bucket")))
				Expect(cosEndpointPrototypeModel.APIKey).To(Equal(core.StringPtr("xxxxxxxxxxxxxx")))
				Expect(cosEndpointPrototypeModel.ServiceToServiceEnabled).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the EventstreamsEndpointPrototype model
				eventstreamsEndpointPrototypeModel := new(atrackerv2.EventstreamsEndpointPrototype)
				Expect(eventstreamsEndpointPrototypeModel).ToNot(BeNil())
				eventstreamsEndpointPrototypeModel.TargetCRN = core.StringPtr("crn:v1:bluemix:public:messagehub:us-south:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::")
				eventstreamsEndpointPrototypeModel.Brokers = []string{"kafka-x:9094"}
				eventstreamsEndpointPrototypeModel.Topic = core.StringPtr("my-topic")
				eventstreamsEndpointPrototypeModel.APIKey = core.StringPtr("xxxxxxxxxxxxxx")
				eventstreamsEndpointPrototypeModel.ServiceToServiceEnabled = core.BoolPtr(false)
				Expect(eventstreamsEndpointPrototypeModel.TargetCRN).To(Equal(core.StringPtr("crn:v1:bluemix:public:messagehub:us-south:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::")))
				Expect(eventstreamsEndpointPrototypeModel.Brokers).To(Equal([]string{"kafka-x:9094"}))
				Expect(eventstreamsEndpointPrototypeModel.Topic).To(Equal(core.StringPtr("my-topic")))
				Expect(eventstreamsEndpointPrototypeModel.APIKey).To(Equal(core.StringPtr("xxxxxxxxxxxxxx")))
				Expect(eventstreamsEndpointPrototypeModel.ServiceToServiceEnabled).To(Equal(core.BoolPtr(false)))

				// Construct an instance of the CloudLogsEndpointPrototype model
				cloudLogsEndpointPrototypeModel := new(atrackerv2.CloudLogsEndpointPrototype)
				Expect(cloudLogsEndpointPrototypeModel).ToNot(BeNil())
				cloudLogsEndpointPrototypeModel.TargetCRN = core.StringPtr("crn:v1:bluemix:public:logs:eu-es:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::")
				Expect(cloudLogsEndpointPrototypeModel.TargetCRN).To(Equal(core.StringPtr("crn:v1:bluemix:public:logs:eu-es:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::")))

				// Construct an instance of the ReplaceTargetOptions model
				id := "testString"
				replaceTargetOptionsModel := atrackerService.NewReplaceTargetOptions(id)
				replaceTargetOptionsModel.SetID("testString")
				replaceTargetOptionsModel.SetName("my-cos-target")
				replaceTargetOptionsModel.SetCosEndpoint(cosEndpointPrototypeModel)
				replaceTargetOptionsModel.SetEventstreamsEndpoint(eventstreamsEndpointPrototypeModel)
				replaceTargetOptionsModel.SetCloudlogsEndpoint(cloudLogsEndpointPrototypeModel)
				replaceTargetOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(replaceTargetOptionsModel).ToNot(BeNil())
				Expect(replaceTargetOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(replaceTargetOptionsModel.Name).To(Equal(core.StringPtr("my-cos-target")))
				Expect(replaceTargetOptionsModel.CosEndpoint).To(Equal(cosEndpointPrototypeModel))
				Expect(replaceTargetOptionsModel.EventstreamsEndpoint).To(Equal(eventstreamsEndpointPrototypeModel))
				Expect(replaceTargetOptionsModel.CloudlogsEndpoint).To(Equal(cloudLogsEndpointPrototypeModel))
				Expect(replaceTargetOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewRulePrototype successfully`, func() {
				targetIds := []string{"c3af557f-fb0e-4476-85c3-0889e7fe7bc4"}
				_model, err := atrackerService.NewRulePrototype(targetIds)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewValidateTargetOptions successfully`, func() {
				// Construct an instance of the ValidateTargetOptions model
				id := "testString"
				validateTargetOptionsModel := atrackerService.NewValidateTargetOptions(id)
				validateTargetOptionsModel.SetID("testString")
				validateTargetOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(validateTargetOptionsModel).ToNot(BeNil())
				Expect(validateTargetOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(validateTargetOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
		})
	})
	Describe(`Model unmarshaling tests`, func() {
		It(`Invoke UnmarshalCloudLogsEndpointPrototype successfully`, func() {
			// Construct an instance of the model.
			model := new(atrackerv2.CloudLogsEndpointPrototype)
			model.TargetCRN = core.StringPtr("crn:v1:bluemix:public:logs:eu-es:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *atrackerv2.CloudLogsEndpointPrototype
			err = atrackerv2.UnmarshalCloudLogsEndpointPrototype(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalCosEndpointPrototype successfully`, func() {
			// Construct an instance of the model.
			model := new(atrackerv2.CosEndpointPrototype)
			model.Endpoint = core.StringPtr("s3.private.us-east.cloud-object-storage.appdomain.cloud")
			model.TargetCRN = core.StringPtr("crn:v1:bluemix:public:cloud-object-storage:global:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::")
			model.Bucket = core.StringPtr("my-atracker-bucket")
			model.APIKey = core.StringPtr("xxxxxxxxxxxxxx")
			model.ServiceToServiceEnabled = core.BoolPtr(true)

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *atrackerv2.CosEndpointPrototype
			err = atrackerv2.UnmarshalCosEndpointPrototype(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalEventstreamsEndpointPrototype successfully`, func() {
			// Construct an instance of the model.
			model := new(atrackerv2.EventstreamsEndpointPrototype)
			model.TargetCRN = core.StringPtr("crn:v1:bluemix:public:messagehub:us-south:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::")
			model.Brokers = []string{"kafka-x:9094"}
			model.Topic = core.StringPtr("my-topic")
			model.APIKey = core.StringPtr("xxxxxxxxxxxxxx")
			model.ServiceToServiceEnabled = core.BoolPtr(false)

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *atrackerv2.EventstreamsEndpointPrototype
			err = atrackerv2.UnmarshalEventstreamsEndpointPrototype(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalRulePrototype successfully`, func() {
			// Construct an instance of the model.
			model := new(atrackerv2.RulePrototype)
			model.TargetIds = []string{"c3af557f-fb0e-4476-85c3-0889e7fe7bc4"}
			model.Locations = []string{"us-south"}

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *atrackerv2.RulePrototype
			err = atrackerv2.UnmarshalRulePrototype(raw, &result)
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
