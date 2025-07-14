/**
 * (C) Copyright IBM Corp. 2024.
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

package enterprisemanagementv1_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/enterprisemanagementv1"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`EnterpriseManagementV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			enterpriseManagementService, serviceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(enterpriseManagementService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			enterpriseManagementService, serviceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(enterpriseManagementService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			enterpriseManagementService, serviceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
				URL: "https://enterprisemanagementv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(enterpriseManagementService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"ENTERPRISE_MANAGEMENT_URL":       "https://enterprisemanagementv1/api",
				"ENTERPRISE_MANAGEMENT_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				enterpriseManagementService, serviceErr := enterprisemanagementv1.NewEnterpriseManagementV1UsingExternalConfig(&enterprisemanagementv1.EnterpriseManagementV1Options{})
				Expect(enterpriseManagementService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := enterpriseManagementService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != enterpriseManagementService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(enterpriseManagementService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(enterpriseManagementService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				enterpriseManagementService, serviceErr := enterprisemanagementv1.NewEnterpriseManagementV1UsingExternalConfig(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL: "https://testService/api",
				})
				Expect(enterpriseManagementService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(enterpriseManagementService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := enterpriseManagementService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != enterpriseManagementService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(enterpriseManagementService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(enterpriseManagementService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				enterpriseManagementService, serviceErr := enterprisemanagementv1.NewEnterpriseManagementV1UsingExternalConfig(&enterprisemanagementv1.EnterpriseManagementV1Options{})
				err := enterpriseManagementService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(enterpriseManagementService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(enterpriseManagementService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := enterpriseManagementService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != enterpriseManagementService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(enterpriseManagementService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(enterpriseManagementService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"ENTERPRISE_MANAGEMENT_URL":       "https://enterprisemanagementv1/api",
				"ENTERPRISE_MANAGEMENT_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			enterpriseManagementService, serviceErr := enterprisemanagementv1.NewEnterpriseManagementV1UsingExternalConfig(&enterprisemanagementv1.EnterpriseManagementV1Options{})

			It(`Instantiate service client with error`, func() {
				Expect(enterpriseManagementService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"ENTERPRISE_MANAGEMENT_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			enterpriseManagementService, serviceErr := enterprisemanagementv1.NewEnterpriseManagementV1UsingExternalConfig(&enterprisemanagementv1.EnterpriseManagementV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(enterpriseManagementService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = enterprisemanagementv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`CreateEnterprise(createEnterpriseOptions *CreateEnterpriseOptions) - Operation response error`, func() {
		createEnterprisePath := "/enterprises"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createEnterprisePath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateEnterprise with error: Operation response processing error`, func() {
				enterpriseManagementService, serviceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(enterpriseManagementService).ToNot(BeNil())

				// Construct an instance of the CreateEnterpriseOptions model
				createEnterpriseOptionsModel := new(enterprisemanagementv1.CreateEnterpriseOptions)
				createEnterpriseOptionsModel.SourceAccountID = core.StringPtr("testString")
				createEnterpriseOptionsModel.Name = core.StringPtr("testString")
				createEnterpriseOptionsModel.PrimaryContactIamID = core.StringPtr("testString")
				createEnterpriseOptionsModel.Domain = core.StringPtr("testString")
				createEnterpriseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := enterpriseManagementService.CreateEnterprise(createEnterpriseOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				enterpriseManagementService.EnableRetries(0, 0)
				result, response, operationErr = enterpriseManagementService.CreateEnterprise(createEnterpriseOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateEnterprise(createEnterpriseOptions *CreateEnterpriseOptions)`, func() {
		createEnterprisePath := "/enterprises"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createEnterprisePath))
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
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"enterprise_id": "EnterpriseID", "enterprise_account_id": "EnterpriseAccountID"}`)
				}))
			})
			It(`Invoke CreateEnterprise successfully with retries`, func() {
				enterpriseManagementService, serviceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(enterpriseManagementService).ToNot(BeNil())
				enterpriseManagementService.EnableRetries(0, 0)

				// Construct an instance of the CreateEnterpriseOptions model
				createEnterpriseOptionsModel := new(enterprisemanagementv1.CreateEnterpriseOptions)
				createEnterpriseOptionsModel.SourceAccountID = core.StringPtr("testString")
				createEnterpriseOptionsModel.Name = core.StringPtr("testString")
				createEnterpriseOptionsModel.PrimaryContactIamID = core.StringPtr("testString")
				createEnterpriseOptionsModel.Domain = core.StringPtr("testString")
				createEnterpriseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := enterpriseManagementService.CreateEnterpriseWithContext(ctx, createEnterpriseOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				enterpriseManagementService.DisableRetries()
				result, response, operationErr := enterpriseManagementService.CreateEnterprise(createEnterpriseOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = enterpriseManagementService.CreateEnterpriseWithContext(ctx, createEnterpriseOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createEnterprisePath))
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
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"enterprise_id": "EnterpriseID", "enterprise_account_id": "EnterpriseAccountID"}`)
				}))
			})
			It(`Invoke CreateEnterprise successfully`, func() {
				enterpriseManagementService, serviceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(enterpriseManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := enterpriseManagementService.CreateEnterprise(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateEnterpriseOptions model
				createEnterpriseOptionsModel := new(enterprisemanagementv1.CreateEnterpriseOptions)
				createEnterpriseOptionsModel.SourceAccountID = core.StringPtr("testString")
				createEnterpriseOptionsModel.Name = core.StringPtr("testString")
				createEnterpriseOptionsModel.PrimaryContactIamID = core.StringPtr("testString")
				createEnterpriseOptionsModel.Domain = core.StringPtr("testString")
				createEnterpriseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = enterpriseManagementService.CreateEnterprise(createEnterpriseOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateEnterprise with error: Operation validation and request error`, func() {
				enterpriseManagementService, serviceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(enterpriseManagementService).ToNot(BeNil())

				// Construct an instance of the CreateEnterpriseOptions model
				createEnterpriseOptionsModel := new(enterprisemanagementv1.CreateEnterpriseOptions)
				createEnterpriseOptionsModel.SourceAccountID = core.StringPtr("testString")
				createEnterpriseOptionsModel.Name = core.StringPtr("testString")
				createEnterpriseOptionsModel.PrimaryContactIamID = core.StringPtr("testString")
				createEnterpriseOptionsModel.Domain = core.StringPtr("testString")
				createEnterpriseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := enterpriseManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := enterpriseManagementService.CreateEnterprise(createEnterpriseOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateEnterpriseOptions model with no property values
				createEnterpriseOptionsModelNew := new(enterprisemanagementv1.CreateEnterpriseOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = enterpriseManagementService.CreateEnterprise(createEnterpriseOptionsModelNew)
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
					res.WriteHeader(202)
				}))
			})
			It(`Invoke CreateEnterprise successfully`, func() {
				enterpriseManagementService, serviceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(enterpriseManagementService).ToNot(BeNil())

				// Construct an instance of the CreateEnterpriseOptions model
				createEnterpriseOptionsModel := new(enterprisemanagementv1.CreateEnterpriseOptions)
				createEnterpriseOptionsModel.SourceAccountID = core.StringPtr("testString")
				createEnterpriseOptionsModel.Name = core.StringPtr("testString")
				createEnterpriseOptionsModel.PrimaryContactIamID = core.StringPtr("testString")
				createEnterpriseOptionsModel.Domain = core.StringPtr("testString")
				createEnterpriseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := enterpriseManagementService.CreateEnterprise(createEnterpriseOptionsModel)
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
	Describe(`ListEnterprises(listEnterprisesOptions *ListEnterprisesOptions) - Operation response error`, func() {
		listEnterprisesPath := "/enterprises"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listEnterprisesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["enterprise_account_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["account_group_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["next_docid"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListEnterprises with error: Operation response processing error`, func() {
				enterpriseManagementService, serviceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(enterpriseManagementService).ToNot(BeNil())

				// Construct an instance of the ListEnterprisesOptions model
				listEnterprisesOptionsModel := new(enterprisemanagementv1.ListEnterprisesOptions)
				listEnterprisesOptionsModel.EnterpriseAccountID = core.StringPtr("testString")
				listEnterprisesOptionsModel.AccountGroupID = core.StringPtr("testString")
				listEnterprisesOptionsModel.AccountID = core.StringPtr("testString")
				listEnterprisesOptionsModel.NextDocid = core.StringPtr("testString")
				listEnterprisesOptionsModel.Limit = core.Int64Ptr(int64(10))
				listEnterprisesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := enterpriseManagementService.ListEnterprises(listEnterprisesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				enterpriseManagementService.EnableRetries(0, 0)
				result, response, operationErr = enterpriseManagementService.ListEnterprises(listEnterprisesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListEnterprises(listEnterprisesOptions *ListEnterprisesOptions)`, func() {
		listEnterprisesPath := "/enterprises"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listEnterprisesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["enterprise_account_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["account_group_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["next_docid"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"rows_count": 9, "next_url": "NextURL", "resources": [{"url": "URL", "id": "ID", "enterprise_account_id": "EnterpriseAccountID", "crn": "CRN", "name": "Name", "domain": "Domain", "state": "State", "primary_contact_iam_id": "PrimaryContactIamID", "primary_contact_email": "PrimaryContactEmail", "source_account_id": "SourceAccountID", "created_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_at": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy"}]}`)
				}))
			})
			It(`Invoke ListEnterprises successfully with retries`, func() {
				enterpriseManagementService, serviceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(enterpriseManagementService).ToNot(BeNil())
				enterpriseManagementService.EnableRetries(0, 0)

				// Construct an instance of the ListEnterprisesOptions model
				listEnterprisesOptionsModel := new(enterprisemanagementv1.ListEnterprisesOptions)
				listEnterprisesOptionsModel.EnterpriseAccountID = core.StringPtr("testString")
				listEnterprisesOptionsModel.AccountGroupID = core.StringPtr("testString")
				listEnterprisesOptionsModel.AccountID = core.StringPtr("testString")
				listEnterprisesOptionsModel.NextDocid = core.StringPtr("testString")
				listEnterprisesOptionsModel.Limit = core.Int64Ptr(int64(10))
				listEnterprisesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := enterpriseManagementService.ListEnterprisesWithContext(ctx, listEnterprisesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				enterpriseManagementService.DisableRetries()
				result, response, operationErr := enterpriseManagementService.ListEnterprises(listEnterprisesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = enterpriseManagementService.ListEnterprisesWithContext(ctx, listEnterprisesOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listEnterprisesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["enterprise_account_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["account_group_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["next_docid"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"rows_count": 9, "next_url": "NextURL", "resources": [{"url": "URL", "id": "ID", "enterprise_account_id": "EnterpriseAccountID", "crn": "CRN", "name": "Name", "domain": "Domain", "state": "State", "primary_contact_iam_id": "PrimaryContactIamID", "primary_contact_email": "PrimaryContactEmail", "source_account_id": "SourceAccountID", "created_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_at": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy"}]}`)
				}))
			})
			It(`Invoke ListEnterprises successfully`, func() {
				enterpriseManagementService, serviceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(enterpriseManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := enterpriseManagementService.ListEnterprises(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListEnterprisesOptions model
				listEnterprisesOptionsModel := new(enterprisemanagementv1.ListEnterprisesOptions)
				listEnterprisesOptionsModel.EnterpriseAccountID = core.StringPtr("testString")
				listEnterprisesOptionsModel.AccountGroupID = core.StringPtr("testString")
				listEnterprisesOptionsModel.AccountID = core.StringPtr("testString")
				listEnterprisesOptionsModel.NextDocid = core.StringPtr("testString")
				listEnterprisesOptionsModel.Limit = core.Int64Ptr(int64(10))
				listEnterprisesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = enterpriseManagementService.ListEnterprises(listEnterprisesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListEnterprises with error: Operation request error`, func() {
				enterpriseManagementService, serviceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(enterpriseManagementService).ToNot(BeNil())

				// Construct an instance of the ListEnterprisesOptions model
				listEnterprisesOptionsModel := new(enterprisemanagementv1.ListEnterprisesOptions)
				listEnterprisesOptionsModel.EnterpriseAccountID = core.StringPtr("testString")
				listEnterprisesOptionsModel.AccountGroupID = core.StringPtr("testString")
				listEnterprisesOptionsModel.AccountID = core.StringPtr("testString")
				listEnterprisesOptionsModel.NextDocid = core.StringPtr("testString")
				listEnterprisesOptionsModel.Limit = core.Int64Ptr(int64(10))
				listEnterprisesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := enterpriseManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := enterpriseManagementService.ListEnterprises(listEnterprisesOptionsModel)
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
			It(`Invoke ListEnterprises successfully`, func() {
				enterpriseManagementService, serviceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(enterpriseManagementService).ToNot(BeNil())

				// Construct an instance of the ListEnterprisesOptions model
				listEnterprisesOptionsModel := new(enterprisemanagementv1.ListEnterprisesOptions)
				listEnterprisesOptionsModel.EnterpriseAccountID = core.StringPtr("testString")
				listEnterprisesOptionsModel.AccountGroupID = core.StringPtr("testString")
				listEnterprisesOptionsModel.AccountID = core.StringPtr("testString")
				listEnterprisesOptionsModel.NextDocid = core.StringPtr("testString")
				listEnterprisesOptionsModel.Limit = core.Int64Ptr(int64(10))
				listEnterprisesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := enterpriseManagementService.ListEnterprises(listEnterprisesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Test pagination helper method on response`, func() {
			It(`Invoke GetNextNextDocid successfully`, func() {
				responseObject := new(enterprisemanagementv1.ListEnterprisesResponse)
				responseObject.NextURL = core.StringPtr("ibm.com?next_docid=abc-123")

				value, err := responseObject.GetNextNextDocid()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.StringPtr("abc-123")))
			})
			It(`Invoke GetNextNextDocid without a "NextURL" property in the response`, func() {
				responseObject := new(enterprisemanagementv1.ListEnterprisesResponse)

				value, err := responseObject.GetNextNextDocid()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
			It(`Invoke GetNextNextDocid without any query params in the "NextURL" URL`, func() {
				responseObject := new(enterprisemanagementv1.ListEnterprisesResponse)
				responseObject.NextURL = core.StringPtr("ibm.com")

				value, err := responseObject.GetNextNextDocid()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
		})
		Context(`Using mock server endpoint - paginated response`, func() {
			BeforeEach(func() {
				var requestNumber int = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listEnterprisesPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					requestNumber++
					if requestNumber == 1 {
						fmt.Fprintf(res, "%s", `{"total_count":2,"limit":1,"next_url":"https://myhost.com/somePath?next_docid=1","resources":[{"url":"URL","id":"ID","enterprise_account_id":"EnterpriseAccountID","crn":"CRN","name":"Name","domain":"Domain","state":"State","primary_contact_iam_id":"PrimaryContactIamID","primary_contact_email":"PrimaryContactEmail","source_account_id":"SourceAccountID","created_at":"2019-01-01T12:00:00.000Z","created_by":"CreatedBy","updated_at":"2019-01-01T12:00:00.000Z","updated_by":"UpdatedBy"}]}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"total_count":2,"limit":1,"resources":[{"url":"URL","id":"ID","enterprise_account_id":"EnterpriseAccountID","crn":"CRN","name":"Name","domain":"Domain","state":"State","primary_contact_iam_id":"PrimaryContactIamID","primary_contact_email":"PrimaryContactEmail","source_account_id":"SourceAccountID","created_at":"2019-01-01T12:00:00.000Z","created_by":"CreatedBy","updated_at":"2019-01-01T12:00:00.000Z","updated_by":"UpdatedBy"}]}`)
					} else {
						res.WriteHeader(400)
					}
				}))
			})
			It(`Use EnterprisesPager.GetNext successfully`, func() {
				enterpriseManagementService, serviceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(enterpriseManagementService).ToNot(BeNil())

				listEnterprisesOptionsModel := &enterprisemanagementv1.ListEnterprisesOptions{
					EnterpriseAccountID: core.StringPtr("testString"),
					AccountGroupID:      core.StringPtr("testString"),
					AccountID:           core.StringPtr("testString"),
					Limit:               core.Int64Ptr(int64(10)),
				}

				pager, err := enterpriseManagementService.NewEnterprisesPager(listEnterprisesOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				var allResults []enterprisemanagementv1.Enterprise
				for pager.HasNext() {
					nextPage, err := pager.GetNext()
					Expect(err).To(BeNil())
					Expect(nextPage).ToNot(BeNil())
					allResults = append(allResults, nextPage...)
				}
				Expect(len(allResults)).To(Equal(2))
			})
			It(`Use EnterprisesPager.GetAll successfully`, func() {
				enterpriseManagementService, serviceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(enterpriseManagementService).ToNot(BeNil())

				listEnterprisesOptionsModel := &enterprisemanagementv1.ListEnterprisesOptions{
					EnterpriseAccountID: core.StringPtr("testString"),
					AccountGroupID:      core.StringPtr("testString"),
					AccountID:           core.StringPtr("testString"),
					Limit:               core.Int64Ptr(int64(10)),
				}

				pager, err := enterpriseManagementService.NewEnterprisesPager(listEnterprisesOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				allResults, err := pager.GetAll()
				Expect(err).To(BeNil())
				Expect(allResults).ToNot(BeNil())
				Expect(len(allResults)).To(Equal(2))
			})
		})
	})
	Describe(`GetEnterprise(getEnterpriseOptions *GetEnterpriseOptions) - Operation response error`, func() {
		getEnterprisePath := "/enterprises/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getEnterprisePath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetEnterprise with error: Operation response processing error`, func() {
				enterpriseManagementService, serviceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(enterpriseManagementService).ToNot(BeNil())

				// Construct an instance of the GetEnterpriseOptions model
				getEnterpriseOptionsModel := new(enterprisemanagementv1.GetEnterpriseOptions)
				getEnterpriseOptionsModel.EnterpriseID = core.StringPtr("testString")
				getEnterpriseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := enterpriseManagementService.GetEnterprise(getEnterpriseOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				enterpriseManagementService.EnableRetries(0, 0)
				result, response, operationErr = enterpriseManagementService.GetEnterprise(getEnterpriseOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetEnterprise(getEnterpriseOptions *GetEnterpriseOptions)`, func() {
		getEnterprisePath := "/enterprises/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getEnterprisePath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"url": "URL", "id": "ID", "enterprise_account_id": "EnterpriseAccountID", "crn": "CRN", "name": "Name", "domain": "Domain", "state": "State", "primary_contact_iam_id": "PrimaryContactIamID", "primary_contact_email": "PrimaryContactEmail", "source_account_id": "SourceAccountID", "created_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_at": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy"}`)
				}))
			})
			It(`Invoke GetEnterprise successfully with retries`, func() {
				enterpriseManagementService, serviceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(enterpriseManagementService).ToNot(BeNil())
				enterpriseManagementService.EnableRetries(0, 0)

				// Construct an instance of the GetEnterpriseOptions model
				getEnterpriseOptionsModel := new(enterprisemanagementv1.GetEnterpriseOptions)
				getEnterpriseOptionsModel.EnterpriseID = core.StringPtr("testString")
				getEnterpriseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := enterpriseManagementService.GetEnterpriseWithContext(ctx, getEnterpriseOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				enterpriseManagementService.DisableRetries()
				result, response, operationErr := enterpriseManagementService.GetEnterprise(getEnterpriseOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = enterpriseManagementService.GetEnterpriseWithContext(ctx, getEnterpriseOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getEnterprisePath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"url": "URL", "id": "ID", "enterprise_account_id": "EnterpriseAccountID", "crn": "CRN", "name": "Name", "domain": "Domain", "state": "State", "primary_contact_iam_id": "PrimaryContactIamID", "primary_contact_email": "PrimaryContactEmail", "source_account_id": "SourceAccountID", "created_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_at": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy"}`)
				}))
			})
			It(`Invoke GetEnterprise successfully`, func() {
				enterpriseManagementService, serviceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(enterpriseManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := enterpriseManagementService.GetEnterprise(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetEnterpriseOptions model
				getEnterpriseOptionsModel := new(enterprisemanagementv1.GetEnterpriseOptions)
				getEnterpriseOptionsModel.EnterpriseID = core.StringPtr("testString")
				getEnterpriseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = enterpriseManagementService.GetEnterprise(getEnterpriseOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetEnterprise with error: Operation validation and request error`, func() {
				enterpriseManagementService, serviceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(enterpriseManagementService).ToNot(BeNil())

				// Construct an instance of the GetEnterpriseOptions model
				getEnterpriseOptionsModel := new(enterprisemanagementv1.GetEnterpriseOptions)
				getEnterpriseOptionsModel.EnterpriseID = core.StringPtr("testString")
				getEnterpriseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := enterpriseManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := enterpriseManagementService.GetEnterprise(getEnterpriseOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetEnterpriseOptions model with no property values
				getEnterpriseOptionsModelNew := new(enterprisemanagementv1.GetEnterpriseOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = enterpriseManagementService.GetEnterprise(getEnterpriseOptionsModelNew)
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
			It(`Invoke GetEnterprise successfully`, func() {
				enterpriseManagementService, serviceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(enterpriseManagementService).ToNot(BeNil())

				// Construct an instance of the GetEnterpriseOptions model
				getEnterpriseOptionsModel := new(enterprisemanagementv1.GetEnterpriseOptions)
				getEnterpriseOptionsModel.EnterpriseID = core.StringPtr("testString")
				getEnterpriseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := enterpriseManagementService.GetEnterprise(getEnterpriseOptionsModel)
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
	Describe(`UpdateEnterprise(updateEnterpriseOptions *UpdateEnterpriseOptions)`, func() {
		updateEnterprisePath := "/enterprises/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateEnterprisePath))
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

					res.WriteHeader(204)
				}))
			})
			It(`Invoke UpdateEnterprise successfully`, func() {
				enterpriseManagementService, serviceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(enterpriseManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := enterpriseManagementService.UpdateEnterprise(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the UpdateEnterpriseOptions model
				updateEnterpriseOptionsModel := new(enterprisemanagementv1.UpdateEnterpriseOptions)
				updateEnterpriseOptionsModel.EnterpriseID = core.StringPtr("testString")
				updateEnterpriseOptionsModel.Name = core.StringPtr("testString")
				updateEnterpriseOptionsModel.Domain = core.StringPtr("testString")
				updateEnterpriseOptionsModel.PrimaryContactIamID = core.StringPtr("testString")
				updateEnterpriseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = enterpriseManagementService.UpdateEnterprise(updateEnterpriseOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke UpdateEnterprise with error: Operation validation and request error`, func() {
				enterpriseManagementService, serviceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(enterpriseManagementService).ToNot(BeNil())

				// Construct an instance of the UpdateEnterpriseOptions model
				updateEnterpriseOptionsModel := new(enterprisemanagementv1.UpdateEnterpriseOptions)
				updateEnterpriseOptionsModel.EnterpriseID = core.StringPtr("testString")
				updateEnterpriseOptionsModel.Name = core.StringPtr("testString")
				updateEnterpriseOptionsModel.Domain = core.StringPtr("testString")
				updateEnterpriseOptionsModel.PrimaryContactIamID = core.StringPtr("testString")
				updateEnterpriseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := enterpriseManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := enterpriseManagementService.UpdateEnterprise(updateEnterpriseOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the UpdateEnterpriseOptions model with no property values
				updateEnterpriseOptionsModelNew := new(enterprisemanagementv1.UpdateEnterpriseOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = enterpriseManagementService.UpdateEnterprise(updateEnterpriseOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ImportAccountToEnterprise(importAccountToEnterpriseOptions *ImportAccountToEnterpriseOptions)`, func() {
		importAccountToEnterprisePath := "/enterprises/testString/import/accounts/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(importAccountToEnterprisePath))
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

					res.WriteHeader(202)
				}))
			})
			It(`Invoke ImportAccountToEnterprise successfully`, func() {
				enterpriseManagementService, serviceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(enterpriseManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := enterpriseManagementService.ImportAccountToEnterprise(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the ImportAccountToEnterpriseOptions model
				importAccountToEnterpriseOptionsModel := new(enterprisemanagementv1.ImportAccountToEnterpriseOptions)
				importAccountToEnterpriseOptionsModel.EnterpriseID = core.StringPtr("testString")
				importAccountToEnterpriseOptionsModel.AccountID = core.StringPtr("testString")
				importAccountToEnterpriseOptionsModel.Parent = core.StringPtr("testString")
				importAccountToEnterpriseOptionsModel.BillingUnitID = core.StringPtr("testString")
				importAccountToEnterpriseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = enterpriseManagementService.ImportAccountToEnterprise(importAccountToEnterpriseOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke ImportAccountToEnterprise with error: Operation validation and request error`, func() {
				enterpriseManagementService, serviceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(enterpriseManagementService).ToNot(BeNil())

				// Construct an instance of the ImportAccountToEnterpriseOptions model
				importAccountToEnterpriseOptionsModel := new(enterprisemanagementv1.ImportAccountToEnterpriseOptions)
				importAccountToEnterpriseOptionsModel.EnterpriseID = core.StringPtr("testString")
				importAccountToEnterpriseOptionsModel.AccountID = core.StringPtr("testString")
				importAccountToEnterpriseOptionsModel.Parent = core.StringPtr("testString")
				importAccountToEnterpriseOptionsModel.BillingUnitID = core.StringPtr("testString")
				importAccountToEnterpriseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := enterpriseManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := enterpriseManagementService.ImportAccountToEnterprise(importAccountToEnterpriseOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the ImportAccountToEnterpriseOptions model with no property values
				importAccountToEnterpriseOptionsModelNew := new(enterprisemanagementv1.ImportAccountToEnterpriseOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = enterpriseManagementService.ImportAccountToEnterprise(importAccountToEnterpriseOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateAccount(createAccountOptions *CreateAccountOptions) - Operation response error`, func() {
		createAccountPath := "/accounts"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createAccountPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateAccount with error: Operation response processing error`, func() {
				enterpriseManagementService, serviceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(enterpriseManagementService).ToNot(BeNil())

				// Construct an instance of the CreateAccountRequestTraits model
				createAccountRequestTraitsModel := new(enterprisemanagementv1.CreateAccountRequestTraits)
				createAccountRequestTraitsModel.Mfa = core.StringPtr("testString")
				createAccountRequestTraitsModel.EnterpriseIamManaged = core.BoolPtr(true)

				// Construct an instance of the CreateAccountRequestOptions model
				createAccountRequestOptionsModel := new(enterprisemanagementv1.CreateAccountRequestOptions)
				createAccountRequestOptionsModel.CreateIamServiceIDWithApikeyAndOwnerPolicies = core.BoolPtr(true)

				// Construct an instance of the CreateAccountOptions model
				createAccountOptionsModel := new(enterprisemanagementv1.CreateAccountOptions)
				createAccountOptionsModel.Parent = core.StringPtr("testString")
				createAccountOptionsModel.Name = core.StringPtr("testString")
				createAccountOptionsModel.OwnerIamID = core.StringPtr("testString")
				createAccountOptionsModel.Traits = createAccountRequestTraitsModel
				createAccountOptionsModel.Options = createAccountRequestOptionsModel
				createAccountOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := enterpriseManagementService.CreateAccount(createAccountOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				enterpriseManagementService.EnableRetries(0, 0)
				result, response, operationErr = enterpriseManagementService.CreateAccount(createAccountOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateAccount(createAccountOptions *CreateAccountOptions)`, func() {
		createAccountPath := "/accounts"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createAccountPath))
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
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"account_id": "AccountID", "iam_service_id": "IamServiceID", "iam_apikey_id": "IamApikeyID", "iam_apikey": "IamApikey"}`)
				}))
			})
			It(`Invoke CreateAccount successfully with retries`, func() {
				enterpriseManagementService, serviceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(enterpriseManagementService).ToNot(BeNil())
				enterpriseManagementService.EnableRetries(0, 0)

				// Construct an instance of the CreateAccountRequestTraits model
				createAccountRequestTraitsModel := new(enterprisemanagementv1.CreateAccountRequestTraits)
				createAccountRequestTraitsModel.Mfa = core.StringPtr("testString")
				createAccountRequestTraitsModel.EnterpriseIamManaged = core.BoolPtr(true)

				// Construct an instance of the CreateAccountRequestOptions model
				createAccountRequestOptionsModel := new(enterprisemanagementv1.CreateAccountRequestOptions)
				createAccountRequestOptionsModel.CreateIamServiceIDWithApikeyAndOwnerPolicies = core.BoolPtr(true)

				// Construct an instance of the CreateAccountOptions model
				createAccountOptionsModel := new(enterprisemanagementv1.CreateAccountOptions)
				createAccountOptionsModel.Parent = core.StringPtr("testString")
				createAccountOptionsModel.Name = core.StringPtr("testString")
				createAccountOptionsModel.OwnerIamID = core.StringPtr("testString")
				createAccountOptionsModel.Traits = createAccountRequestTraitsModel
				createAccountOptionsModel.Options = createAccountRequestOptionsModel
				createAccountOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := enterpriseManagementService.CreateAccountWithContext(ctx, createAccountOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				enterpriseManagementService.DisableRetries()
				result, response, operationErr := enterpriseManagementService.CreateAccount(createAccountOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = enterpriseManagementService.CreateAccountWithContext(ctx, createAccountOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createAccountPath))
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
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"account_id": "AccountID", "iam_service_id": "IamServiceID", "iam_apikey_id": "IamApikeyID", "iam_apikey": "IamApikey"}`)
				}))
			})
			It(`Invoke CreateAccount successfully`, func() {
				enterpriseManagementService, serviceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(enterpriseManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := enterpriseManagementService.CreateAccount(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateAccountRequestTraits model
				createAccountRequestTraitsModel := new(enterprisemanagementv1.CreateAccountRequestTraits)
				createAccountRequestTraitsModel.Mfa = core.StringPtr("testString")
				createAccountRequestTraitsModel.EnterpriseIamManaged = core.BoolPtr(true)

				// Construct an instance of the CreateAccountRequestOptions model
				createAccountRequestOptionsModel := new(enterprisemanagementv1.CreateAccountRequestOptions)
				createAccountRequestOptionsModel.CreateIamServiceIDWithApikeyAndOwnerPolicies = core.BoolPtr(true)

				// Construct an instance of the CreateAccountOptions model
				createAccountOptionsModel := new(enterprisemanagementv1.CreateAccountOptions)
				createAccountOptionsModel.Parent = core.StringPtr("testString")
				createAccountOptionsModel.Name = core.StringPtr("testString")
				createAccountOptionsModel.OwnerIamID = core.StringPtr("testString")
				createAccountOptionsModel.Traits = createAccountRequestTraitsModel
				createAccountOptionsModel.Options = createAccountRequestOptionsModel
				createAccountOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = enterpriseManagementService.CreateAccount(createAccountOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateAccount with error: Operation validation and request error`, func() {
				enterpriseManagementService, serviceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(enterpriseManagementService).ToNot(BeNil())

				// Construct an instance of the CreateAccountRequestTraits model
				createAccountRequestTraitsModel := new(enterprisemanagementv1.CreateAccountRequestTraits)
				createAccountRequestTraitsModel.Mfa = core.StringPtr("testString")
				createAccountRequestTraitsModel.EnterpriseIamManaged = core.BoolPtr(true)

				// Construct an instance of the CreateAccountRequestOptions model
				createAccountRequestOptionsModel := new(enterprisemanagementv1.CreateAccountRequestOptions)
				createAccountRequestOptionsModel.CreateIamServiceIDWithApikeyAndOwnerPolicies = core.BoolPtr(true)

				// Construct an instance of the CreateAccountOptions model
				createAccountOptionsModel := new(enterprisemanagementv1.CreateAccountOptions)
				createAccountOptionsModel.Parent = core.StringPtr("testString")
				createAccountOptionsModel.Name = core.StringPtr("testString")
				createAccountOptionsModel.OwnerIamID = core.StringPtr("testString")
				createAccountOptionsModel.Traits = createAccountRequestTraitsModel
				createAccountOptionsModel.Options = createAccountRequestOptionsModel
				createAccountOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := enterpriseManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := enterpriseManagementService.CreateAccount(createAccountOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateAccountOptions model with no property values
				createAccountOptionsModelNew := new(enterprisemanagementv1.CreateAccountOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = enterpriseManagementService.CreateAccount(createAccountOptionsModelNew)
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
					res.WriteHeader(202)
				}))
			})
			It(`Invoke CreateAccount successfully`, func() {
				enterpriseManagementService, serviceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(enterpriseManagementService).ToNot(BeNil())

				// Construct an instance of the CreateAccountRequestTraits model
				createAccountRequestTraitsModel := new(enterprisemanagementv1.CreateAccountRequestTraits)
				createAccountRequestTraitsModel.Mfa = core.StringPtr("testString")
				createAccountRequestTraitsModel.EnterpriseIamManaged = core.BoolPtr(true)

				// Construct an instance of the CreateAccountRequestOptions model
				createAccountRequestOptionsModel := new(enterprisemanagementv1.CreateAccountRequestOptions)
				createAccountRequestOptionsModel.CreateIamServiceIDWithApikeyAndOwnerPolicies = core.BoolPtr(true)

				// Construct an instance of the CreateAccountOptions model
				createAccountOptionsModel := new(enterprisemanagementv1.CreateAccountOptions)
				createAccountOptionsModel.Parent = core.StringPtr("testString")
				createAccountOptionsModel.Name = core.StringPtr("testString")
				createAccountOptionsModel.OwnerIamID = core.StringPtr("testString")
				createAccountOptionsModel.Traits = createAccountRequestTraitsModel
				createAccountOptionsModel.Options = createAccountRequestOptionsModel
				createAccountOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := enterpriseManagementService.CreateAccount(createAccountOptionsModel)
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
	Describe(`ListAccounts(listAccountsOptions *ListAccountsOptions) - Operation response error`, func() {
		listAccountsPath := "/accounts"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listAccountsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["enterprise_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["account_group_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["next_docid"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["parent"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					// TODO: Add check for include_deleted query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListAccounts with error: Operation response processing error`, func() {
				enterpriseManagementService, serviceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(enterpriseManagementService).ToNot(BeNil())

				// Construct an instance of the ListAccountsOptions model
				listAccountsOptionsModel := new(enterprisemanagementv1.ListAccountsOptions)
				listAccountsOptionsModel.EnterpriseID = core.StringPtr("testString")
				listAccountsOptionsModel.AccountGroupID = core.StringPtr("testString")
				listAccountsOptionsModel.NextDocid = core.StringPtr("testString")
				listAccountsOptionsModel.Parent = core.StringPtr("testString")
				listAccountsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listAccountsOptionsModel.IncludeDeleted = core.BoolPtr(true)
				listAccountsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := enterpriseManagementService.ListAccounts(listAccountsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				enterpriseManagementService.EnableRetries(0, 0)
				result, response, operationErr = enterpriseManagementService.ListAccounts(listAccountsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListAccounts(listAccountsOptions *ListAccountsOptions)`, func() {
		listAccountsPath := "/accounts"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listAccountsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["enterprise_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["account_group_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["next_docid"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["parent"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					// TODO: Add check for include_deleted query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"rows_count": 9, "next_url": "NextURL", "resources": [{"url": "URL", "id": "ID", "crn": "CRN", "parent": "Parent", "enterprise_account_id": "EnterpriseAccountID", "enterprise_id": "EnterpriseID", "enterprise_path": "EnterprisePath", "name": "Name", "state": "State", "owner_iam_id": "OwnerIamID", "paid": true, "owner_email": "OwnerEmail", "is_enterprise_account": false, "created_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_at": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy"}]}`)
				}))
			})
			It(`Invoke ListAccounts successfully with retries`, func() {
				enterpriseManagementService, serviceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(enterpriseManagementService).ToNot(BeNil())
				enterpriseManagementService.EnableRetries(0, 0)

				// Construct an instance of the ListAccountsOptions model
				listAccountsOptionsModel := new(enterprisemanagementv1.ListAccountsOptions)
				listAccountsOptionsModel.EnterpriseID = core.StringPtr("testString")
				listAccountsOptionsModel.AccountGroupID = core.StringPtr("testString")
				listAccountsOptionsModel.NextDocid = core.StringPtr("testString")
				listAccountsOptionsModel.Parent = core.StringPtr("testString")
				listAccountsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listAccountsOptionsModel.IncludeDeleted = core.BoolPtr(true)
				listAccountsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := enterpriseManagementService.ListAccountsWithContext(ctx, listAccountsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				enterpriseManagementService.DisableRetries()
				result, response, operationErr := enterpriseManagementService.ListAccounts(listAccountsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = enterpriseManagementService.ListAccountsWithContext(ctx, listAccountsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listAccountsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["enterprise_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["account_group_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["next_docid"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["parent"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					// TODO: Add check for include_deleted query parameter
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"rows_count": 9, "next_url": "NextURL", "resources": [{"url": "URL", "id": "ID", "crn": "CRN", "parent": "Parent", "enterprise_account_id": "EnterpriseAccountID", "enterprise_id": "EnterpriseID", "enterprise_path": "EnterprisePath", "name": "Name", "state": "State", "owner_iam_id": "OwnerIamID", "paid": true, "owner_email": "OwnerEmail", "is_enterprise_account": false, "created_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_at": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy"}]}`)
				}))
			})
			It(`Invoke ListAccounts successfully`, func() {
				enterpriseManagementService, serviceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(enterpriseManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := enterpriseManagementService.ListAccounts(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListAccountsOptions model
				listAccountsOptionsModel := new(enterprisemanagementv1.ListAccountsOptions)
				listAccountsOptionsModel.EnterpriseID = core.StringPtr("testString")
				listAccountsOptionsModel.AccountGroupID = core.StringPtr("testString")
				listAccountsOptionsModel.NextDocid = core.StringPtr("testString")
				listAccountsOptionsModel.Parent = core.StringPtr("testString")
				listAccountsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listAccountsOptionsModel.IncludeDeleted = core.BoolPtr(true)
				listAccountsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = enterpriseManagementService.ListAccounts(listAccountsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListAccounts with error: Operation request error`, func() {
				enterpriseManagementService, serviceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(enterpriseManagementService).ToNot(BeNil())

				// Construct an instance of the ListAccountsOptions model
				listAccountsOptionsModel := new(enterprisemanagementv1.ListAccountsOptions)
				listAccountsOptionsModel.EnterpriseID = core.StringPtr("testString")
				listAccountsOptionsModel.AccountGroupID = core.StringPtr("testString")
				listAccountsOptionsModel.NextDocid = core.StringPtr("testString")
				listAccountsOptionsModel.Parent = core.StringPtr("testString")
				listAccountsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listAccountsOptionsModel.IncludeDeleted = core.BoolPtr(true)
				listAccountsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := enterpriseManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := enterpriseManagementService.ListAccounts(listAccountsOptionsModel)
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
			It(`Invoke ListAccounts successfully`, func() {
				enterpriseManagementService, serviceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(enterpriseManagementService).ToNot(BeNil())

				// Construct an instance of the ListAccountsOptions model
				listAccountsOptionsModel := new(enterprisemanagementv1.ListAccountsOptions)
				listAccountsOptionsModel.EnterpriseID = core.StringPtr("testString")
				listAccountsOptionsModel.AccountGroupID = core.StringPtr("testString")
				listAccountsOptionsModel.NextDocid = core.StringPtr("testString")
				listAccountsOptionsModel.Parent = core.StringPtr("testString")
				listAccountsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listAccountsOptionsModel.IncludeDeleted = core.BoolPtr(true)
				listAccountsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := enterpriseManagementService.ListAccounts(listAccountsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Test pagination helper method on response`, func() {
			It(`Invoke GetNextNextDocid successfully`, func() {
				responseObject := new(enterprisemanagementv1.ListAccountsResponse)
				responseObject.NextURL = core.StringPtr("ibm.com?next_docid=abc-123")

				value, err := responseObject.GetNextNextDocid()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.StringPtr("abc-123")))
			})
			It(`Invoke GetNextNextDocid without a "NextURL" property in the response`, func() {
				responseObject := new(enterprisemanagementv1.ListAccountsResponse)

				value, err := responseObject.GetNextNextDocid()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
			It(`Invoke GetNextNextDocid without any query params in the "NextURL" URL`, func() {
				responseObject := new(enterprisemanagementv1.ListAccountsResponse)
				responseObject.NextURL = core.StringPtr("ibm.com")

				value, err := responseObject.GetNextNextDocid()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
		})
		Context(`Using mock server endpoint - paginated response`, func() {
			BeforeEach(func() {
				var requestNumber int = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listAccountsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					requestNumber++
					if requestNumber == 1 {
						fmt.Fprintf(res, "%s", `{"total_count":2,"limit":1,"next_url":"https://myhost.com/somePath?next_docid=1","resources":[{"url":"URL","id":"ID","crn":"CRN","parent":"Parent","enterprise_account_id":"EnterpriseAccountID","enterprise_id":"EnterpriseID","enterprise_path":"EnterprisePath","name":"Name","state":"State","owner_iam_id":"OwnerIamID","paid":true,"owner_email":"OwnerEmail","is_enterprise_account":false,"created_at":"2019-01-01T12:00:00.000Z","created_by":"CreatedBy","updated_at":"2019-01-01T12:00:00.000Z","updated_by":"UpdatedBy"}]}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"total_count":2,"limit":1,"resources":[{"url":"URL","id":"ID","crn":"CRN","parent":"Parent","enterprise_account_id":"EnterpriseAccountID","enterprise_id":"EnterpriseID","enterprise_path":"EnterprisePath","name":"Name","state":"State","owner_iam_id":"OwnerIamID","paid":true,"owner_email":"OwnerEmail","is_enterprise_account":false,"created_at":"2019-01-01T12:00:00.000Z","created_by":"CreatedBy","updated_at":"2019-01-01T12:00:00.000Z","updated_by":"UpdatedBy"}]}`)
					} else {
						res.WriteHeader(400)
					}
				}))
			})
			It(`Use AccountsPager.GetNext successfully`, func() {
				enterpriseManagementService, serviceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(enterpriseManagementService).ToNot(BeNil())

				listAccountsOptionsModel := &enterprisemanagementv1.ListAccountsOptions{
					EnterpriseID:   core.StringPtr("testString"),
					AccountGroupID: core.StringPtr("testString"),
					Parent:         core.StringPtr("testString"),
					Limit:          core.Int64Ptr(int64(10)),
					IncludeDeleted: core.BoolPtr(true),
				}

				pager, err := enterpriseManagementService.NewAccountsPager(listAccountsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				var allResults []enterprisemanagementv1.Account
				for pager.HasNext() {
					nextPage, err := pager.GetNext()
					Expect(err).To(BeNil())
					Expect(nextPage).ToNot(BeNil())
					allResults = append(allResults, nextPage...)
				}
				Expect(len(allResults)).To(Equal(2))
			})
			It(`Use AccountsPager.GetAll successfully`, func() {
				enterpriseManagementService, serviceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(enterpriseManagementService).ToNot(BeNil())

				listAccountsOptionsModel := &enterprisemanagementv1.ListAccountsOptions{
					EnterpriseID:   core.StringPtr("testString"),
					AccountGroupID: core.StringPtr("testString"),
					Parent:         core.StringPtr("testString"),
					Limit:          core.Int64Ptr(int64(10)),
					IncludeDeleted: core.BoolPtr(true),
				}

				pager, err := enterpriseManagementService.NewAccountsPager(listAccountsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				allResults, err := pager.GetAll()
				Expect(err).To(BeNil())
				Expect(allResults).ToNot(BeNil())
				Expect(len(allResults)).To(Equal(2))
			})
		})
	})
	Describe(`GetAccount(getAccountOptions *GetAccountOptions) - Operation response error`, func() {
		getAccountPath := "/accounts/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAccountPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetAccount with error: Operation response processing error`, func() {
				enterpriseManagementService, serviceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(enterpriseManagementService).ToNot(BeNil())

				// Construct an instance of the GetAccountOptions model
				getAccountOptionsModel := new(enterprisemanagementv1.GetAccountOptions)
				getAccountOptionsModel.AccountID = core.StringPtr("testString")
				getAccountOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := enterpriseManagementService.GetAccount(getAccountOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				enterpriseManagementService.EnableRetries(0, 0)
				result, response, operationErr = enterpriseManagementService.GetAccount(getAccountOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetAccount(getAccountOptions *GetAccountOptions)`, func() {
		getAccountPath := "/accounts/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAccountPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"url": "URL", "id": "ID", "crn": "CRN", "parent": "Parent", "enterprise_account_id": "EnterpriseAccountID", "enterprise_id": "EnterpriseID", "enterprise_path": "EnterprisePath", "name": "Name", "state": "State", "owner_iam_id": "OwnerIamID", "paid": true, "owner_email": "OwnerEmail", "is_enterprise_account": false, "created_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_at": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy"}`)
				}))
			})
			It(`Invoke GetAccount successfully with retries`, func() {
				enterpriseManagementService, serviceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(enterpriseManagementService).ToNot(BeNil())
				enterpriseManagementService.EnableRetries(0, 0)

				// Construct an instance of the GetAccountOptions model
				getAccountOptionsModel := new(enterprisemanagementv1.GetAccountOptions)
				getAccountOptionsModel.AccountID = core.StringPtr("testString")
				getAccountOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := enterpriseManagementService.GetAccountWithContext(ctx, getAccountOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				enterpriseManagementService.DisableRetries()
				result, response, operationErr := enterpriseManagementService.GetAccount(getAccountOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = enterpriseManagementService.GetAccountWithContext(ctx, getAccountOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getAccountPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"url": "URL", "id": "ID", "crn": "CRN", "parent": "Parent", "enterprise_account_id": "EnterpriseAccountID", "enterprise_id": "EnterpriseID", "enterprise_path": "EnterprisePath", "name": "Name", "state": "State", "owner_iam_id": "OwnerIamID", "paid": true, "owner_email": "OwnerEmail", "is_enterprise_account": false, "created_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_at": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy"}`)
				}))
			})
			It(`Invoke GetAccount successfully`, func() {
				enterpriseManagementService, serviceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(enterpriseManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := enterpriseManagementService.GetAccount(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetAccountOptions model
				getAccountOptionsModel := new(enterprisemanagementv1.GetAccountOptions)
				getAccountOptionsModel.AccountID = core.StringPtr("testString")
				getAccountOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = enterpriseManagementService.GetAccount(getAccountOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetAccount with error: Operation validation and request error`, func() {
				enterpriseManagementService, serviceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(enterpriseManagementService).ToNot(BeNil())

				// Construct an instance of the GetAccountOptions model
				getAccountOptionsModel := new(enterprisemanagementv1.GetAccountOptions)
				getAccountOptionsModel.AccountID = core.StringPtr("testString")
				getAccountOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := enterpriseManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := enterpriseManagementService.GetAccount(getAccountOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetAccountOptions model with no property values
				getAccountOptionsModelNew := new(enterprisemanagementv1.GetAccountOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = enterpriseManagementService.GetAccount(getAccountOptionsModelNew)
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
			It(`Invoke GetAccount successfully`, func() {
				enterpriseManagementService, serviceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(enterpriseManagementService).ToNot(BeNil())

				// Construct an instance of the GetAccountOptions model
				getAccountOptionsModel := new(enterprisemanagementv1.GetAccountOptions)
				getAccountOptionsModel.AccountID = core.StringPtr("testString")
				getAccountOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := enterpriseManagementService.GetAccount(getAccountOptionsModel)
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
	Describe(`UpdateAccount(updateAccountOptions *UpdateAccountOptions)`, func() {
		updateAccountPath := "/accounts/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateAccountPath))
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

					res.WriteHeader(202)
				}))
			})
			It(`Invoke UpdateAccount successfully`, func() {
				enterpriseManagementService, serviceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(enterpriseManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := enterpriseManagementService.UpdateAccount(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the UpdateAccountOptions model
				updateAccountOptionsModel := new(enterprisemanagementv1.UpdateAccountOptions)
				updateAccountOptionsModel.AccountID = core.StringPtr("testString")
				updateAccountOptionsModel.Parent = core.StringPtr("testString")
				updateAccountOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = enterpriseManagementService.UpdateAccount(updateAccountOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke UpdateAccount with error: Operation validation and request error`, func() {
				enterpriseManagementService, serviceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(enterpriseManagementService).ToNot(BeNil())

				// Construct an instance of the UpdateAccountOptions model
				updateAccountOptionsModel := new(enterprisemanagementv1.UpdateAccountOptions)
				updateAccountOptionsModel.AccountID = core.StringPtr("testString")
				updateAccountOptionsModel.Parent = core.StringPtr("testString")
				updateAccountOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := enterpriseManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := enterpriseManagementService.UpdateAccount(updateAccountOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the UpdateAccountOptions model with no property values
				updateAccountOptionsModelNew := new(enterprisemanagementv1.UpdateAccountOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = enterpriseManagementService.UpdateAccount(updateAccountOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteAccount(deleteAccountOptions *DeleteAccountOptions)`, func() {
		deleteAccountPath := "/accounts/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteAccountPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteAccount successfully`, func() {
				enterpriseManagementService, serviceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(enterpriseManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := enterpriseManagementService.DeleteAccount(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteAccountOptions model
				deleteAccountOptionsModel := new(enterprisemanagementv1.DeleteAccountOptions)
				deleteAccountOptionsModel.AccountID = core.StringPtr("testString")
				deleteAccountOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = enterpriseManagementService.DeleteAccount(deleteAccountOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteAccount with error: Operation validation and request error`, func() {
				enterpriseManagementService, serviceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(enterpriseManagementService).ToNot(BeNil())

				// Construct an instance of the DeleteAccountOptions model
				deleteAccountOptionsModel := new(enterprisemanagementv1.DeleteAccountOptions)
				deleteAccountOptionsModel.AccountID = core.StringPtr("testString")
				deleteAccountOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := enterpriseManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := enterpriseManagementService.DeleteAccount(deleteAccountOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteAccountOptions model with no property values
				deleteAccountOptionsModelNew := new(enterprisemanagementv1.DeleteAccountOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = enterpriseManagementService.DeleteAccount(deleteAccountOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateAccountGroup(createAccountGroupOptions *CreateAccountGroupOptions) - Operation response error`, func() {
		createAccountGroupPath := "/account-groups"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createAccountGroupPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateAccountGroup with error: Operation response processing error`, func() {
				enterpriseManagementService, serviceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(enterpriseManagementService).ToNot(BeNil())

				// Construct an instance of the CreateAccountGroupOptions model
				createAccountGroupOptionsModel := new(enterprisemanagementv1.CreateAccountGroupOptions)
				createAccountGroupOptionsModel.Parent = core.StringPtr("testString")
				createAccountGroupOptionsModel.Name = core.StringPtr("testString")
				createAccountGroupOptionsModel.PrimaryContactIamID = core.StringPtr("testString")
				createAccountGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := enterpriseManagementService.CreateAccountGroup(createAccountGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				enterpriseManagementService.EnableRetries(0, 0)
				result, response, operationErr = enterpriseManagementService.CreateAccountGroup(createAccountGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateAccountGroup(createAccountGroupOptions *CreateAccountGroupOptions)`, func() {
		createAccountGroupPath := "/account-groups"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createAccountGroupPath))
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
					fmt.Fprintf(res, "%s", `{"account_group_id": "AccountGroupID"}`)
				}))
			})
			It(`Invoke CreateAccountGroup successfully with retries`, func() {
				enterpriseManagementService, serviceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(enterpriseManagementService).ToNot(BeNil())
				enterpriseManagementService.EnableRetries(0, 0)

				// Construct an instance of the CreateAccountGroupOptions model
				createAccountGroupOptionsModel := new(enterprisemanagementv1.CreateAccountGroupOptions)
				createAccountGroupOptionsModel.Parent = core.StringPtr("testString")
				createAccountGroupOptionsModel.Name = core.StringPtr("testString")
				createAccountGroupOptionsModel.PrimaryContactIamID = core.StringPtr("testString")
				createAccountGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := enterpriseManagementService.CreateAccountGroupWithContext(ctx, createAccountGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				enterpriseManagementService.DisableRetries()
				result, response, operationErr := enterpriseManagementService.CreateAccountGroup(createAccountGroupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = enterpriseManagementService.CreateAccountGroupWithContext(ctx, createAccountGroupOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createAccountGroupPath))
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
					fmt.Fprintf(res, "%s", `{"account_group_id": "AccountGroupID"}`)
				}))
			})
			It(`Invoke CreateAccountGroup successfully`, func() {
				enterpriseManagementService, serviceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(enterpriseManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := enterpriseManagementService.CreateAccountGroup(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateAccountGroupOptions model
				createAccountGroupOptionsModel := new(enterprisemanagementv1.CreateAccountGroupOptions)
				createAccountGroupOptionsModel.Parent = core.StringPtr("testString")
				createAccountGroupOptionsModel.Name = core.StringPtr("testString")
				createAccountGroupOptionsModel.PrimaryContactIamID = core.StringPtr("testString")
				createAccountGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = enterpriseManagementService.CreateAccountGroup(createAccountGroupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateAccountGroup with error: Operation validation and request error`, func() {
				enterpriseManagementService, serviceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(enterpriseManagementService).ToNot(BeNil())

				// Construct an instance of the CreateAccountGroupOptions model
				createAccountGroupOptionsModel := new(enterprisemanagementv1.CreateAccountGroupOptions)
				createAccountGroupOptionsModel.Parent = core.StringPtr("testString")
				createAccountGroupOptionsModel.Name = core.StringPtr("testString")
				createAccountGroupOptionsModel.PrimaryContactIamID = core.StringPtr("testString")
				createAccountGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := enterpriseManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := enterpriseManagementService.CreateAccountGroup(createAccountGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateAccountGroupOptions model with no property values
				createAccountGroupOptionsModelNew := new(enterprisemanagementv1.CreateAccountGroupOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = enterpriseManagementService.CreateAccountGroup(createAccountGroupOptionsModelNew)
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
			It(`Invoke CreateAccountGroup successfully`, func() {
				enterpriseManagementService, serviceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(enterpriseManagementService).ToNot(BeNil())

				// Construct an instance of the CreateAccountGroupOptions model
				createAccountGroupOptionsModel := new(enterprisemanagementv1.CreateAccountGroupOptions)
				createAccountGroupOptionsModel.Parent = core.StringPtr("testString")
				createAccountGroupOptionsModel.Name = core.StringPtr("testString")
				createAccountGroupOptionsModel.PrimaryContactIamID = core.StringPtr("testString")
				createAccountGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := enterpriseManagementService.CreateAccountGroup(createAccountGroupOptionsModel)
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
	Describe(`ListAccountGroups(listAccountGroupsOptions *ListAccountGroupsOptions) - Operation response error`, func() {
		listAccountGroupsPath := "/account-groups"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listAccountGroupsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["enterprise_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["parent_account_group_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["next_docid"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["parent"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					// TODO: Add check for include_deleted query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListAccountGroups with error: Operation response processing error`, func() {
				enterpriseManagementService, serviceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(enterpriseManagementService).ToNot(BeNil())

				// Construct an instance of the ListAccountGroupsOptions model
				listAccountGroupsOptionsModel := new(enterprisemanagementv1.ListAccountGroupsOptions)
				listAccountGroupsOptionsModel.EnterpriseID = core.StringPtr("testString")
				listAccountGroupsOptionsModel.ParentAccountGroupID = core.StringPtr("testString")
				listAccountGroupsOptionsModel.NextDocid = core.StringPtr("testString")
				listAccountGroupsOptionsModel.Parent = core.StringPtr("testString")
				listAccountGroupsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listAccountGroupsOptionsModel.IncludeDeleted = core.BoolPtr(true)
				listAccountGroupsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := enterpriseManagementService.ListAccountGroups(listAccountGroupsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				enterpriseManagementService.EnableRetries(0, 0)
				result, response, operationErr = enterpriseManagementService.ListAccountGroups(listAccountGroupsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListAccountGroups(listAccountGroupsOptions *ListAccountGroupsOptions)`, func() {
		listAccountGroupsPath := "/account-groups"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listAccountGroupsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["enterprise_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["parent_account_group_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["next_docid"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["parent"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					// TODO: Add check for include_deleted query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"rows_count": 9, "next_url": "NextURL", "resources": [{"url": "URL", "id": "ID", "crn": "CRN", "parent": "Parent", "enterprise_account_id": "EnterpriseAccountID", "enterprise_id": "EnterpriseID", "enterprise_path": "EnterprisePath", "name": "Name", "state": "State", "primary_contact_iam_id": "PrimaryContactIamID", "primary_contact_email": "PrimaryContactEmail", "created_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_at": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy"}]}`)
				}))
			})
			It(`Invoke ListAccountGroups successfully with retries`, func() {
				enterpriseManagementService, serviceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(enterpriseManagementService).ToNot(BeNil())
				enterpriseManagementService.EnableRetries(0, 0)

				// Construct an instance of the ListAccountGroupsOptions model
				listAccountGroupsOptionsModel := new(enterprisemanagementv1.ListAccountGroupsOptions)
				listAccountGroupsOptionsModel.EnterpriseID = core.StringPtr("testString")
				listAccountGroupsOptionsModel.ParentAccountGroupID = core.StringPtr("testString")
				listAccountGroupsOptionsModel.NextDocid = core.StringPtr("testString")
				listAccountGroupsOptionsModel.Parent = core.StringPtr("testString")
				listAccountGroupsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listAccountGroupsOptionsModel.IncludeDeleted = core.BoolPtr(true)
				listAccountGroupsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := enterpriseManagementService.ListAccountGroupsWithContext(ctx, listAccountGroupsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				enterpriseManagementService.DisableRetries()
				result, response, operationErr := enterpriseManagementService.ListAccountGroups(listAccountGroupsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = enterpriseManagementService.ListAccountGroupsWithContext(ctx, listAccountGroupsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listAccountGroupsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["enterprise_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["parent_account_group_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["next_docid"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["parent"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					// TODO: Add check for include_deleted query parameter
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"rows_count": 9, "next_url": "NextURL", "resources": [{"url": "URL", "id": "ID", "crn": "CRN", "parent": "Parent", "enterprise_account_id": "EnterpriseAccountID", "enterprise_id": "EnterpriseID", "enterprise_path": "EnterprisePath", "name": "Name", "state": "State", "primary_contact_iam_id": "PrimaryContactIamID", "primary_contact_email": "PrimaryContactEmail", "created_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_at": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy"}]}`)
				}))
			})
			It(`Invoke ListAccountGroups successfully`, func() {
				enterpriseManagementService, serviceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(enterpriseManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := enterpriseManagementService.ListAccountGroups(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListAccountGroupsOptions model
				listAccountGroupsOptionsModel := new(enterprisemanagementv1.ListAccountGroupsOptions)
				listAccountGroupsOptionsModel.EnterpriseID = core.StringPtr("testString")
				listAccountGroupsOptionsModel.ParentAccountGroupID = core.StringPtr("testString")
				listAccountGroupsOptionsModel.NextDocid = core.StringPtr("testString")
				listAccountGroupsOptionsModel.Parent = core.StringPtr("testString")
				listAccountGroupsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listAccountGroupsOptionsModel.IncludeDeleted = core.BoolPtr(true)
				listAccountGroupsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = enterpriseManagementService.ListAccountGroups(listAccountGroupsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListAccountGroups with error: Operation request error`, func() {
				enterpriseManagementService, serviceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(enterpriseManagementService).ToNot(BeNil())

				// Construct an instance of the ListAccountGroupsOptions model
				listAccountGroupsOptionsModel := new(enterprisemanagementv1.ListAccountGroupsOptions)
				listAccountGroupsOptionsModel.EnterpriseID = core.StringPtr("testString")
				listAccountGroupsOptionsModel.ParentAccountGroupID = core.StringPtr("testString")
				listAccountGroupsOptionsModel.NextDocid = core.StringPtr("testString")
				listAccountGroupsOptionsModel.Parent = core.StringPtr("testString")
				listAccountGroupsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listAccountGroupsOptionsModel.IncludeDeleted = core.BoolPtr(true)
				listAccountGroupsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := enterpriseManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := enterpriseManagementService.ListAccountGroups(listAccountGroupsOptionsModel)
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
			It(`Invoke ListAccountGroups successfully`, func() {
				enterpriseManagementService, serviceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(enterpriseManagementService).ToNot(BeNil())

				// Construct an instance of the ListAccountGroupsOptions model
				listAccountGroupsOptionsModel := new(enterprisemanagementv1.ListAccountGroupsOptions)
				listAccountGroupsOptionsModel.EnterpriseID = core.StringPtr("testString")
				listAccountGroupsOptionsModel.ParentAccountGroupID = core.StringPtr("testString")
				listAccountGroupsOptionsModel.NextDocid = core.StringPtr("testString")
				listAccountGroupsOptionsModel.Parent = core.StringPtr("testString")
				listAccountGroupsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listAccountGroupsOptionsModel.IncludeDeleted = core.BoolPtr(true)
				listAccountGroupsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := enterpriseManagementService.ListAccountGroups(listAccountGroupsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Test pagination helper method on response`, func() {
			It(`Invoke GetNextNextDocid successfully`, func() {
				responseObject := new(enterprisemanagementv1.ListAccountGroupsResponse)
				responseObject.NextURL = core.StringPtr("ibm.com?next_docid=abc-123")

				value, err := responseObject.GetNextNextDocid()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.StringPtr("abc-123")))
			})
			It(`Invoke GetNextNextDocid without a "NextURL" property in the response`, func() {
				responseObject := new(enterprisemanagementv1.ListAccountGroupsResponse)

				value, err := responseObject.GetNextNextDocid()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
			It(`Invoke GetNextNextDocid without any query params in the "NextURL" URL`, func() {
				responseObject := new(enterprisemanagementv1.ListAccountGroupsResponse)
				responseObject.NextURL = core.StringPtr("ibm.com")

				value, err := responseObject.GetNextNextDocid()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
		})
		Context(`Using mock server endpoint - paginated response`, func() {
			BeforeEach(func() {
				var requestNumber int = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listAccountGroupsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					requestNumber++
					if requestNumber == 1 {
						fmt.Fprintf(res, "%s", `{"total_count":2,"limit":1,"next_url":"https://myhost.com/somePath?next_docid=1","resources":[{"url":"URL","id":"ID","crn":"CRN","parent":"Parent","enterprise_account_id":"EnterpriseAccountID","enterprise_id":"EnterpriseID","enterprise_path":"EnterprisePath","name":"Name","state":"State","primary_contact_iam_id":"PrimaryContactIamID","primary_contact_email":"PrimaryContactEmail","created_at":"2019-01-01T12:00:00.000Z","created_by":"CreatedBy","updated_at":"2019-01-01T12:00:00.000Z","updated_by":"UpdatedBy"}]}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"total_count":2,"limit":1,"resources":[{"url":"URL","id":"ID","crn":"CRN","parent":"Parent","enterprise_account_id":"EnterpriseAccountID","enterprise_id":"EnterpriseID","enterprise_path":"EnterprisePath","name":"Name","state":"State","primary_contact_iam_id":"PrimaryContactIamID","primary_contact_email":"PrimaryContactEmail","created_at":"2019-01-01T12:00:00.000Z","created_by":"CreatedBy","updated_at":"2019-01-01T12:00:00.000Z","updated_by":"UpdatedBy"}]}`)
					} else {
						res.WriteHeader(400)
					}
				}))
			})
			It(`Use AccountGroupsPager.GetNext successfully`, func() {
				enterpriseManagementService, serviceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(enterpriseManagementService).ToNot(BeNil())

				listAccountGroupsOptionsModel := &enterprisemanagementv1.ListAccountGroupsOptions{
					EnterpriseID:         core.StringPtr("testString"),
					ParentAccountGroupID: core.StringPtr("testString"),
					Parent:               core.StringPtr("testString"),
					Limit:                core.Int64Ptr(int64(10)),
					IncludeDeleted:       core.BoolPtr(true),
				}

				pager, err := enterpriseManagementService.NewAccountGroupsPager(listAccountGroupsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				var allResults []enterprisemanagementv1.AccountGroup
				for pager.HasNext() {
					nextPage, err := pager.GetNext()
					Expect(err).To(BeNil())
					Expect(nextPage).ToNot(BeNil())
					allResults = append(allResults, nextPage...)
				}
				Expect(len(allResults)).To(Equal(2))
			})
			It(`Use AccountGroupsPager.GetAll successfully`, func() {
				enterpriseManagementService, serviceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(enterpriseManagementService).ToNot(BeNil())

				listAccountGroupsOptionsModel := &enterprisemanagementv1.ListAccountGroupsOptions{
					EnterpriseID:         core.StringPtr("testString"),
					ParentAccountGroupID: core.StringPtr("testString"),
					Parent:               core.StringPtr("testString"),
					Limit:                core.Int64Ptr(int64(10)),
					IncludeDeleted:       core.BoolPtr(true),
				}

				pager, err := enterpriseManagementService.NewAccountGroupsPager(listAccountGroupsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				allResults, err := pager.GetAll()
				Expect(err).To(BeNil())
				Expect(allResults).ToNot(BeNil())
				Expect(len(allResults)).To(Equal(2))
			})
		})
	})
	Describe(`GetAccountGroup(getAccountGroupOptions *GetAccountGroupOptions) - Operation response error`, func() {
		getAccountGroupPath := "/account-groups/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAccountGroupPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetAccountGroup with error: Operation response processing error`, func() {
				enterpriseManagementService, serviceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(enterpriseManagementService).ToNot(BeNil())

				// Construct an instance of the GetAccountGroupOptions model
				getAccountGroupOptionsModel := new(enterprisemanagementv1.GetAccountGroupOptions)
				getAccountGroupOptionsModel.AccountGroupID = core.StringPtr("testString")
				getAccountGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := enterpriseManagementService.GetAccountGroup(getAccountGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				enterpriseManagementService.EnableRetries(0, 0)
				result, response, operationErr = enterpriseManagementService.GetAccountGroup(getAccountGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetAccountGroup(getAccountGroupOptions *GetAccountGroupOptions)`, func() {
		getAccountGroupPath := "/account-groups/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAccountGroupPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"url": "URL", "id": "ID", "crn": "CRN", "parent": "Parent", "enterprise_account_id": "EnterpriseAccountID", "enterprise_id": "EnterpriseID", "enterprise_path": "EnterprisePath", "name": "Name", "state": "State", "primary_contact_iam_id": "PrimaryContactIamID", "primary_contact_email": "PrimaryContactEmail", "created_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_at": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy"}`)
				}))
			})
			It(`Invoke GetAccountGroup successfully with retries`, func() {
				enterpriseManagementService, serviceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(enterpriseManagementService).ToNot(BeNil())
				enterpriseManagementService.EnableRetries(0, 0)

				// Construct an instance of the GetAccountGroupOptions model
				getAccountGroupOptionsModel := new(enterprisemanagementv1.GetAccountGroupOptions)
				getAccountGroupOptionsModel.AccountGroupID = core.StringPtr("testString")
				getAccountGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := enterpriseManagementService.GetAccountGroupWithContext(ctx, getAccountGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				enterpriseManagementService.DisableRetries()
				result, response, operationErr := enterpriseManagementService.GetAccountGroup(getAccountGroupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = enterpriseManagementService.GetAccountGroupWithContext(ctx, getAccountGroupOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getAccountGroupPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"url": "URL", "id": "ID", "crn": "CRN", "parent": "Parent", "enterprise_account_id": "EnterpriseAccountID", "enterprise_id": "EnterpriseID", "enterprise_path": "EnterprisePath", "name": "Name", "state": "State", "primary_contact_iam_id": "PrimaryContactIamID", "primary_contact_email": "PrimaryContactEmail", "created_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "updated_at": "2019-01-01T12:00:00.000Z", "updated_by": "UpdatedBy"}`)
				}))
			})
			It(`Invoke GetAccountGroup successfully`, func() {
				enterpriseManagementService, serviceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(enterpriseManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := enterpriseManagementService.GetAccountGroup(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetAccountGroupOptions model
				getAccountGroupOptionsModel := new(enterprisemanagementv1.GetAccountGroupOptions)
				getAccountGroupOptionsModel.AccountGroupID = core.StringPtr("testString")
				getAccountGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = enterpriseManagementService.GetAccountGroup(getAccountGroupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetAccountGroup with error: Operation validation and request error`, func() {
				enterpriseManagementService, serviceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(enterpriseManagementService).ToNot(BeNil())

				// Construct an instance of the GetAccountGroupOptions model
				getAccountGroupOptionsModel := new(enterprisemanagementv1.GetAccountGroupOptions)
				getAccountGroupOptionsModel.AccountGroupID = core.StringPtr("testString")
				getAccountGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := enterpriseManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := enterpriseManagementService.GetAccountGroup(getAccountGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetAccountGroupOptions model with no property values
				getAccountGroupOptionsModelNew := new(enterprisemanagementv1.GetAccountGroupOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = enterpriseManagementService.GetAccountGroup(getAccountGroupOptionsModelNew)
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
			It(`Invoke GetAccountGroup successfully`, func() {
				enterpriseManagementService, serviceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(enterpriseManagementService).ToNot(BeNil())

				// Construct an instance of the GetAccountGroupOptions model
				getAccountGroupOptionsModel := new(enterprisemanagementv1.GetAccountGroupOptions)
				getAccountGroupOptionsModel.AccountGroupID = core.StringPtr("testString")
				getAccountGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := enterpriseManagementService.GetAccountGroup(getAccountGroupOptionsModel)
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
	Describe(`UpdateAccountGroup(updateAccountGroupOptions *UpdateAccountGroupOptions)`, func() {
		updateAccountGroupPath := "/account-groups/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateAccountGroupPath))
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

					res.WriteHeader(204)
				}))
			})
			It(`Invoke UpdateAccountGroup successfully`, func() {
				enterpriseManagementService, serviceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(enterpriseManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := enterpriseManagementService.UpdateAccountGroup(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the UpdateAccountGroupOptions model
				updateAccountGroupOptionsModel := new(enterprisemanagementv1.UpdateAccountGroupOptions)
				updateAccountGroupOptionsModel.AccountGroupID = core.StringPtr("testString")
				updateAccountGroupOptionsModel.Name = core.StringPtr("testString")
				updateAccountGroupOptionsModel.PrimaryContactIamID = core.StringPtr("testString")
				updateAccountGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = enterpriseManagementService.UpdateAccountGroup(updateAccountGroupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke UpdateAccountGroup with error: Operation validation and request error`, func() {
				enterpriseManagementService, serviceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(enterpriseManagementService).ToNot(BeNil())

				// Construct an instance of the UpdateAccountGroupOptions model
				updateAccountGroupOptionsModel := new(enterprisemanagementv1.UpdateAccountGroupOptions)
				updateAccountGroupOptionsModel.AccountGroupID = core.StringPtr("testString")
				updateAccountGroupOptionsModel.Name = core.StringPtr("testString")
				updateAccountGroupOptionsModel.PrimaryContactIamID = core.StringPtr("testString")
				updateAccountGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := enterpriseManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := enterpriseManagementService.UpdateAccountGroup(updateAccountGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the UpdateAccountGroupOptions model with no property values
				updateAccountGroupOptionsModelNew := new(enterprisemanagementv1.UpdateAccountGroupOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = enterpriseManagementService.UpdateAccountGroup(updateAccountGroupOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteAccountGroup(deleteAccountGroupOptions *DeleteAccountGroupOptions)`, func() {
		deleteAccountGroupPath := "/account-groups/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteAccountGroupPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteAccountGroup successfully`, func() {
				enterpriseManagementService, serviceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(enterpriseManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := enterpriseManagementService.DeleteAccountGroup(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteAccountGroupOptions model
				deleteAccountGroupOptionsModel := new(enterprisemanagementv1.DeleteAccountGroupOptions)
				deleteAccountGroupOptionsModel.AccountGroupID = core.StringPtr("testString")
				deleteAccountGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = enterpriseManagementService.DeleteAccountGroup(deleteAccountGroupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteAccountGroup with error: Operation validation and request error`, func() {
				enterpriseManagementService, serviceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(enterpriseManagementService).ToNot(BeNil())

				// Construct an instance of the DeleteAccountGroupOptions model
				deleteAccountGroupOptionsModel := new(enterprisemanagementv1.DeleteAccountGroupOptions)
				deleteAccountGroupOptionsModel.AccountGroupID = core.StringPtr("testString")
				deleteAccountGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := enterpriseManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := enterpriseManagementService.DeleteAccountGroup(deleteAccountGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteAccountGroupOptions model with no property values
				deleteAccountGroupOptionsModelNew := new(enterprisemanagementv1.DeleteAccountGroupOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = enterpriseManagementService.DeleteAccountGroup(deleteAccountGroupOptionsModelNew)
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
			enterpriseManagementService, _ := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
				URL:           "http://enterprisemanagementv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewCreateAccountGroupOptions successfully`, func() {
				// Construct an instance of the CreateAccountGroupOptions model
				createAccountGroupOptionsParent := "testString"
				createAccountGroupOptionsName := "testString"
				createAccountGroupOptionsPrimaryContactIamID := "testString"
				createAccountGroupOptionsModel := enterpriseManagementService.NewCreateAccountGroupOptions(createAccountGroupOptionsParent, createAccountGroupOptionsName, createAccountGroupOptionsPrimaryContactIamID)
				createAccountGroupOptionsModel.SetParent("testString")
				createAccountGroupOptionsModel.SetName("testString")
				createAccountGroupOptionsModel.SetPrimaryContactIamID("testString")
				createAccountGroupOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createAccountGroupOptionsModel).ToNot(BeNil())
				Expect(createAccountGroupOptionsModel.Parent).To(Equal(core.StringPtr("testString")))
				Expect(createAccountGroupOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(createAccountGroupOptionsModel.PrimaryContactIamID).To(Equal(core.StringPtr("testString")))
				Expect(createAccountGroupOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateAccountOptions successfully`, func() {
				// Construct an instance of the CreateAccountRequestTraits model
				createAccountRequestTraitsModel := new(enterprisemanagementv1.CreateAccountRequestTraits)
				Expect(createAccountRequestTraitsModel).ToNot(BeNil())
				createAccountRequestTraitsModel.Mfa = core.StringPtr("testString")
				createAccountRequestTraitsModel.EnterpriseIamManaged = core.BoolPtr(true)
				Expect(createAccountRequestTraitsModel.Mfa).To(Equal(core.StringPtr("testString")))
				Expect(createAccountRequestTraitsModel.EnterpriseIamManaged).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the CreateAccountRequestOptions model
				createAccountRequestOptionsModel := new(enterprisemanagementv1.CreateAccountRequestOptions)
				Expect(createAccountRequestOptionsModel).ToNot(BeNil())
				createAccountRequestOptionsModel.CreateIamServiceIDWithApikeyAndOwnerPolicies = core.BoolPtr(true)
				Expect(createAccountRequestOptionsModel.CreateIamServiceIDWithApikeyAndOwnerPolicies).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the CreateAccountOptions model
				createAccountOptionsParent := "testString"
				createAccountOptionsName := "testString"
				createAccountOptionsOwnerIamID := "testString"
				createAccountOptionsModel := enterpriseManagementService.NewCreateAccountOptions(createAccountOptionsParent, createAccountOptionsName, createAccountOptionsOwnerIamID)
				createAccountOptionsModel.SetParent("testString")
				createAccountOptionsModel.SetName("testString")
				createAccountOptionsModel.SetOwnerIamID("testString")
				createAccountOptionsModel.SetTraits(createAccountRequestTraitsModel)
				createAccountOptionsModel.SetOptions(createAccountRequestOptionsModel)
				createAccountOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createAccountOptionsModel).ToNot(BeNil())
				Expect(createAccountOptionsModel.Parent).To(Equal(core.StringPtr("testString")))
				Expect(createAccountOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(createAccountOptionsModel.OwnerIamID).To(Equal(core.StringPtr("testString")))
				Expect(createAccountOptionsModel.Traits).To(Equal(createAccountRequestTraitsModel))
				Expect(createAccountOptionsModel.Options).To(Equal(createAccountRequestOptionsModel))
				Expect(createAccountOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateEnterpriseOptions successfully`, func() {
				// Construct an instance of the CreateEnterpriseOptions model
				createEnterpriseOptionsSourceAccountID := "testString"
				createEnterpriseOptionsName := "testString"
				createEnterpriseOptionsPrimaryContactIamID := "testString"
				createEnterpriseOptionsModel := enterpriseManagementService.NewCreateEnterpriseOptions(createEnterpriseOptionsSourceAccountID, createEnterpriseOptionsName, createEnterpriseOptionsPrimaryContactIamID)
				createEnterpriseOptionsModel.SetSourceAccountID("testString")
				createEnterpriseOptionsModel.SetName("testString")
				createEnterpriseOptionsModel.SetPrimaryContactIamID("testString")
				createEnterpriseOptionsModel.SetDomain("testString")
				createEnterpriseOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createEnterpriseOptionsModel).ToNot(BeNil())
				Expect(createEnterpriseOptionsModel.SourceAccountID).To(Equal(core.StringPtr("testString")))
				Expect(createEnterpriseOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(createEnterpriseOptionsModel.PrimaryContactIamID).To(Equal(core.StringPtr("testString")))
				Expect(createEnterpriseOptionsModel.Domain).To(Equal(core.StringPtr("testString")))
				Expect(createEnterpriseOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteAccountGroupOptions successfully`, func() {
				// Construct an instance of the DeleteAccountGroupOptions model
				accountGroupID := "testString"
				deleteAccountGroupOptionsModel := enterpriseManagementService.NewDeleteAccountGroupOptions(accountGroupID)
				deleteAccountGroupOptionsModel.SetAccountGroupID("testString")
				deleteAccountGroupOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteAccountGroupOptionsModel).ToNot(BeNil())
				Expect(deleteAccountGroupOptionsModel.AccountGroupID).To(Equal(core.StringPtr("testString")))
				Expect(deleteAccountGroupOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteAccountOptions successfully`, func() {
				// Construct an instance of the DeleteAccountOptions model
				accountID := "testString"
				deleteAccountOptionsModel := enterpriseManagementService.NewDeleteAccountOptions(accountID)
				deleteAccountOptionsModel.SetAccountID("testString")
				deleteAccountOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteAccountOptionsModel).ToNot(BeNil())
				Expect(deleteAccountOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(deleteAccountOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetAccountGroupOptions successfully`, func() {
				// Construct an instance of the GetAccountGroupOptions model
				accountGroupID := "testString"
				getAccountGroupOptionsModel := enterpriseManagementService.NewGetAccountGroupOptions(accountGroupID)
				getAccountGroupOptionsModel.SetAccountGroupID("testString")
				getAccountGroupOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getAccountGroupOptionsModel).ToNot(BeNil())
				Expect(getAccountGroupOptionsModel.AccountGroupID).To(Equal(core.StringPtr("testString")))
				Expect(getAccountGroupOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetAccountOptions successfully`, func() {
				// Construct an instance of the GetAccountOptions model
				accountID := "testString"
				getAccountOptionsModel := enterpriseManagementService.NewGetAccountOptions(accountID)
				getAccountOptionsModel.SetAccountID("testString")
				getAccountOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getAccountOptionsModel).ToNot(BeNil())
				Expect(getAccountOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(getAccountOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetEnterpriseOptions successfully`, func() {
				// Construct an instance of the GetEnterpriseOptions model
				enterpriseID := "testString"
				getEnterpriseOptionsModel := enterpriseManagementService.NewGetEnterpriseOptions(enterpriseID)
				getEnterpriseOptionsModel.SetEnterpriseID("testString")
				getEnterpriseOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getEnterpriseOptionsModel).ToNot(BeNil())
				Expect(getEnterpriseOptionsModel.EnterpriseID).To(Equal(core.StringPtr("testString")))
				Expect(getEnterpriseOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewImportAccountToEnterpriseOptions successfully`, func() {
				// Construct an instance of the ImportAccountToEnterpriseOptions model
				enterpriseID := "testString"
				accountID := "testString"
				importAccountToEnterpriseOptionsModel := enterpriseManagementService.NewImportAccountToEnterpriseOptions(enterpriseID, accountID)
				importAccountToEnterpriseOptionsModel.SetEnterpriseID("testString")
				importAccountToEnterpriseOptionsModel.SetAccountID("testString")
				importAccountToEnterpriseOptionsModel.SetParent("testString")
				importAccountToEnterpriseOptionsModel.SetBillingUnitID("testString")
				importAccountToEnterpriseOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(importAccountToEnterpriseOptionsModel).ToNot(BeNil())
				Expect(importAccountToEnterpriseOptionsModel.EnterpriseID).To(Equal(core.StringPtr("testString")))
				Expect(importAccountToEnterpriseOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(importAccountToEnterpriseOptionsModel.Parent).To(Equal(core.StringPtr("testString")))
				Expect(importAccountToEnterpriseOptionsModel.BillingUnitID).To(Equal(core.StringPtr("testString")))
				Expect(importAccountToEnterpriseOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListAccountGroupsOptions successfully`, func() {
				// Construct an instance of the ListAccountGroupsOptions model
				listAccountGroupsOptionsModel := enterpriseManagementService.NewListAccountGroupsOptions()
				listAccountGroupsOptionsModel.SetEnterpriseID("testString")
				listAccountGroupsOptionsModel.SetParentAccountGroupID("testString")
				listAccountGroupsOptionsModel.SetNextDocid("testString")
				listAccountGroupsOptionsModel.SetParent("testString")
				listAccountGroupsOptionsModel.SetLimit(int64(10))
				listAccountGroupsOptionsModel.SetIncludeDeleted(true)
				listAccountGroupsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listAccountGroupsOptionsModel).ToNot(BeNil())
				Expect(listAccountGroupsOptionsModel.EnterpriseID).To(Equal(core.StringPtr("testString")))
				Expect(listAccountGroupsOptionsModel.ParentAccountGroupID).To(Equal(core.StringPtr("testString")))
				Expect(listAccountGroupsOptionsModel.NextDocid).To(Equal(core.StringPtr("testString")))
				Expect(listAccountGroupsOptionsModel.Parent).To(Equal(core.StringPtr("testString")))
				Expect(listAccountGroupsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(10))))
				Expect(listAccountGroupsOptionsModel.IncludeDeleted).To(Equal(core.BoolPtr(true)))
				Expect(listAccountGroupsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListAccountsOptions successfully`, func() {
				// Construct an instance of the ListAccountsOptions model
				listAccountsOptionsModel := enterpriseManagementService.NewListAccountsOptions()
				listAccountsOptionsModel.SetEnterpriseID("testString")
				listAccountsOptionsModel.SetAccountGroupID("testString")
				listAccountsOptionsModel.SetNextDocid("testString")
				listAccountsOptionsModel.SetParent("testString")
				listAccountsOptionsModel.SetLimit(int64(10))
				listAccountsOptionsModel.SetIncludeDeleted(true)
				listAccountsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listAccountsOptionsModel).ToNot(BeNil())
				Expect(listAccountsOptionsModel.EnterpriseID).To(Equal(core.StringPtr("testString")))
				Expect(listAccountsOptionsModel.AccountGroupID).To(Equal(core.StringPtr("testString")))
				Expect(listAccountsOptionsModel.NextDocid).To(Equal(core.StringPtr("testString")))
				Expect(listAccountsOptionsModel.Parent).To(Equal(core.StringPtr("testString")))
				Expect(listAccountsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(10))))
				Expect(listAccountsOptionsModel.IncludeDeleted).To(Equal(core.BoolPtr(true)))
				Expect(listAccountsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListEnterprisesOptions successfully`, func() {
				// Construct an instance of the ListEnterprisesOptions model
				listEnterprisesOptionsModel := enterpriseManagementService.NewListEnterprisesOptions()
				listEnterprisesOptionsModel.SetEnterpriseAccountID("testString")
				listEnterprisesOptionsModel.SetAccountGroupID("testString")
				listEnterprisesOptionsModel.SetAccountID("testString")
				listEnterprisesOptionsModel.SetNextDocid("testString")
				listEnterprisesOptionsModel.SetLimit(int64(10))
				listEnterprisesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listEnterprisesOptionsModel).ToNot(BeNil())
				Expect(listEnterprisesOptionsModel.EnterpriseAccountID).To(Equal(core.StringPtr("testString")))
				Expect(listEnterprisesOptionsModel.AccountGroupID).To(Equal(core.StringPtr("testString")))
				Expect(listEnterprisesOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(listEnterprisesOptionsModel.NextDocid).To(Equal(core.StringPtr("testString")))
				Expect(listEnterprisesOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(10))))
				Expect(listEnterprisesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateAccountGroupOptions successfully`, func() {
				// Construct an instance of the UpdateAccountGroupOptions model
				accountGroupID := "testString"
				updateAccountGroupOptionsModel := enterpriseManagementService.NewUpdateAccountGroupOptions(accountGroupID)
				updateAccountGroupOptionsModel.SetAccountGroupID("testString")
				updateAccountGroupOptionsModel.SetName("testString")
				updateAccountGroupOptionsModel.SetPrimaryContactIamID("testString")
				updateAccountGroupOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateAccountGroupOptionsModel).ToNot(BeNil())
				Expect(updateAccountGroupOptionsModel.AccountGroupID).To(Equal(core.StringPtr("testString")))
				Expect(updateAccountGroupOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(updateAccountGroupOptionsModel.PrimaryContactIamID).To(Equal(core.StringPtr("testString")))
				Expect(updateAccountGroupOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateAccountOptions successfully`, func() {
				// Construct an instance of the UpdateAccountOptions model
				accountID := "testString"
				updateAccountOptionsParent := "testString"
				updateAccountOptionsModel := enterpriseManagementService.NewUpdateAccountOptions(accountID, updateAccountOptionsParent)
				updateAccountOptionsModel.SetAccountID("testString")
				updateAccountOptionsModel.SetParent("testString")
				updateAccountOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateAccountOptionsModel).ToNot(BeNil())
				Expect(updateAccountOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(updateAccountOptionsModel.Parent).To(Equal(core.StringPtr("testString")))
				Expect(updateAccountOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateEnterpriseOptions successfully`, func() {
				// Construct an instance of the UpdateEnterpriseOptions model
				enterpriseID := "testString"
				updateEnterpriseOptionsModel := enterpriseManagementService.NewUpdateEnterpriseOptions(enterpriseID)
				updateEnterpriseOptionsModel.SetEnterpriseID("testString")
				updateEnterpriseOptionsModel.SetName("testString")
				updateEnterpriseOptionsModel.SetDomain("testString")
				updateEnterpriseOptionsModel.SetPrimaryContactIamID("testString")
				updateEnterpriseOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateEnterpriseOptionsModel).ToNot(BeNil())
				Expect(updateEnterpriseOptionsModel.EnterpriseID).To(Equal(core.StringPtr("testString")))
				Expect(updateEnterpriseOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(updateEnterpriseOptionsModel.Domain).To(Equal(core.StringPtr("testString")))
				Expect(updateEnterpriseOptionsModel.PrimaryContactIamID).To(Equal(core.StringPtr("testString")))
				Expect(updateEnterpriseOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
		})
	})
	Describe(`Model unmarshaling tests`, func() {
		It(`Invoke UnmarshalCreateAccountRequestOptions successfully`, func() {
			// Construct an instance of the model.
			model := new(enterprisemanagementv1.CreateAccountRequestOptions)
			model.CreateIamServiceIDWithApikeyAndOwnerPolicies = core.BoolPtr(true)

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *enterprisemanagementv1.CreateAccountRequestOptions
			err = enterprisemanagementv1.UnmarshalCreateAccountRequestOptions(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalCreateAccountRequestTraits successfully`, func() {
			// Construct an instance of the model.
			model := new(enterprisemanagementv1.CreateAccountRequestTraits)
			model.Mfa = core.StringPtr("testString")
			model.EnterpriseIamManaged = core.BoolPtr(true)

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *enterprisemanagementv1.CreateAccountRequestTraits
			err = enterprisemanagementv1.UnmarshalCreateAccountRequestTraits(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
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

func CreateMockByteArray(mockData string) *[]byte {
	ba := []byte(mockData)
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
