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

package ibmcloudfunctionsnamespaceapiv1_test

import (
	"fmt"
	"github.com/IBM/go-sdk-core/v3/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
	"strings"
)

var _ = Describe(`IbmCloudFunctionsNamespaceApiV1`, func() {
	Describe(`GetNamespaces(getNamespacesOptions *GetNamespacesOptions)`, func() {
		getNamespacesPath := "/namespaces"
		Context(`Successfully - Retrieve all IBM Cloud Functions namespaces (classic and IAM)`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getNamespacesPath))
				Expect(req.Method).To(Equal("GET"))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"limit": 5, "namespaces": [], "offset": 6, "total_count": 10}`)
			}))
			It(`Succeed to call GetNamespaces`, func() {
				defer testServer.Close()

				testService, testServiceErr := ibmcloudfunctionsnamespaceapiv1.NewIbmCloudFunctionsNamespaceApiV1(&ibmcloudfunctionsnamespaceapiv1.IbmCloudFunctionsNamespaceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.GetNamespaces(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				getNamespacesOptions := testService.NewGetNamespacesOptions()
				result, response, operationErr = testService.GetNamespaces(getNamespacesOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`CreateNamespace(createNamespaceOptions *CreateNamespaceOptions)`, func() {
		createNamespacePath := "/namespaces"
		name := "exampleString"
		resourceGroupID := "exampleString"
		resourcePlanID := "exampleString"
		Context(`Successfully - Create an IBM Cloud Functions namespace`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createNamespacePath))
				Expect(req.Method).To(Equal("POST"))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"api_key_created": "2017-05-16T13:56:54.957Z", "api_key_id": "fake_ApiKeyID", "crn": "fake_Crn", "description": "fake_Description", "id": "fake_ID", "location": "fake_Location", "name": "fake_Name", "resource_group_id": "fake_ResourceGroupID", "resource_plan_id": "fake_ResourcePlanID", "service_id": "fake_ServiceID"}`)
			}))
			It(`Succeed to call CreateNamespace`, func() {
				defer testServer.Close()

				testService, testServiceErr := ibmcloudfunctionsnamespaceapiv1.NewIbmCloudFunctionsNamespaceApiV1(&ibmcloudfunctionsnamespaceapiv1.IbmCloudFunctionsNamespaceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.CreateNamespace(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				createNamespaceOptions := testService.NewCreateNamespaceOptions(name, resourceGroupID, resourcePlanID)
				result, response, operationErr = testService.CreateNamespace(createNamespaceOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`GetNamespace(getNamespaceOptions *GetNamespaceOptions)`, func() {
		getNamespacePath := "/namespaces/{id}"
		id := "exampleString"
		getNamespacePath = strings.Replace(getNamespacePath, "{id}", id, 1)
		Context(`Successfully - Retrieve an IBM Cloud Functions namespace`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getNamespacePath))
				Expect(req.Method).To(Equal("GET"))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"id": "fake_ID", "location": "fake_Location"}`)
			}))
			It(`Succeed to call GetNamespace`, func() {
				defer testServer.Close()

				testService, testServiceErr := ibmcloudfunctionsnamespaceapiv1.NewIbmCloudFunctionsNamespaceApiV1(&ibmcloudfunctionsnamespaceapiv1.IbmCloudFunctionsNamespaceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.GetNamespace(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				getNamespaceOptions := testService.NewGetNamespaceOptions(id)
				result, response, operationErr = testService.GetNamespace(getNamespaceOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`DeleteNamespace(deleteNamespaceOptions *DeleteNamespaceOptions)`, func() {
		deleteNamespacePath := "/namespaces/{id}"
		id := "exampleString"
		deleteNamespacePath = strings.Replace(deleteNamespacePath, "{id}", id, 1)
		Context(`Successfully - Delete an IBM Cloud Functions namespace`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteNamespacePath))
				Expect(req.Method).To(Equal("DELETE"))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"id": "fake_ID"}`)
			}))
			It(`Succeed to call DeleteNamespace`, func() {
				defer testServer.Close()

				testService, testServiceErr := ibmcloudfunctionsnamespaceapiv1.NewIbmCloudFunctionsNamespaceApiV1(&ibmcloudfunctionsnamespaceapiv1.IbmCloudFunctionsNamespaceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.DeleteNamespace(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				deleteNamespaceOptions := testService.NewDeleteNamespaceOptions(id)
				result, response, operationErr = testService.DeleteNamespace(deleteNamespaceOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`UpdateNamespace(updateNamespaceOptions *UpdateNamespaceOptions)`, func() {
		updateNamespacePath := "/namespaces/{id}"
		id := "exampleString"
		updateNamespacePath = strings.Replace(updateNamespacePath, "{id}", id, 1)
		Context(`Successfully - Update an IBM Cloud Functions namespace`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(updateNamespacePath))
				Expect(req.Method).To(Equal("PATCH"))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"id": "fake_ID", "location": "fake_Location"}`)
			}))
			It(`Succeed to call UpdateNamespace`, func() {
				defer testServer.Close()

				testService, testServiceErr := ibmcloudfunctionsnamespaceapiv1.NewIbmCloudFunctionsNamespaceApiV1(&ibmcloudfunctionsnamespaceapiv1.IbmCloudFunctionsNamespaceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.UpdateNamespace(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				updateNamespaceOptions := testService.NewUpdateNamespaceOptions(id)
				result, response, operationErr = testService.UpdateNamespace(updateNamespaceOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`UpdateNamespaceAPIKey(updateNamespaceAPIKeyOptions *UpdateNamespaceAPIKeyOptions)`, func() {
		updateNamespaceApiKeyPath := "/namespaces/{id}/apikey"
		id := "exampleString"
		updateNamespaceApiKeyPath = strings.Replace(updateNamespaceApiKeyPath, "{id}", id, 1)
		Context(`Successfully - Update an IBM Cloud Functions namespace API key`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(updateNamespaceApiKeyPath))
				Expect(req.Method).To(Equal("PATCH"))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"id": "fake_ID", "location": "fake_Location"}`)
			}))
			It(`Succeed to call UpdateNamespaceAPIKey`, func() {
				defer testServer.Close()

				testService, testServiceErr := ibmcloudfunctionsnamespaceapiv1.NewIbmCloudFunctionsNamespaceApiV1(&ibmcloudfunctionsnamespaceapiv1.IbmCloudFunctionsNamespaceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.UpdateNamespaceAPIKey(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				updateNamespaceApiKeyOptions := testService.NewUpdateNamespaceAPIKeyOptions(id)
				result, response, operationErr = testService.UpdateNamespaceAPIKey(updateNamespaceApiKeyOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`MigrateNamespace(migrateNamespaceOptions *MigrateNamespaceOptions)`, func() {
		migrateNamespacePath := "/namespaces/{id}/migrate"
		id := "exampleString"
		name := "exampleString"
		resourceGroupID := "exampleString"
		resourcePlanID := "exampleString"
		migrateNamespacePath = strings.Replace(migrateNamespacePath, "{id}", id, 1)
		Context(`Successfully - Migrate a classic namespace and create an IAM enabled IBM Cloud Functions namespace`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(migrateNamespacePath))
				Expect(req.Method).To(Equal("POST"))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"api_key_created": "2017-05-16T13:56:54.957Z", "api_key_id": "fake_ApiKeyID", "crn": "fake_Crn", "description": "fake_Description", "id": "fake_ID", "location": "fake_Location", "name": "fake_Name", "resource_group_id": "fake_ResourceGroupID", "resource_plan_id": "fake_ResourcePlanID", "service_id": "fake_ServiceID"}`)
			}))
			It(`Succeed to call MigrateNamespace`, func() {
				defer testServer.Close()

				testService, testServiceErr := ibmcloudfunctionsnamespaceapiv1.NewIbmCloudFunctionsNamespaceApiV1(&ibmcloudfunctionsnamespaceapiv1.IbmCloudFunctionsNamespaceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.MigrateNamespace(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				migrateNamespaceOptions := testService.NewMigrateNamespaceOptions(id, name, resourceGroupID, resourcePlanID)
				result, response, operationErr = testService.MigrateNamespace(migrateNamespaceOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
})
