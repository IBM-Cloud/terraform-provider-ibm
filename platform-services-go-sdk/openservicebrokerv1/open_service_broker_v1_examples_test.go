//go:build examples

/**
 * (C) Copyright IBM Corp. 2020, 2021.
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
	"encoding/json"
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/openservicebrokerv1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//
// This file provides an example of how to use the Open Service Broker service.
//
// The following configuration properties are assumed to be defined:
//
// OPEN_SERVICE_BROKER_URL=<Service broker's URL>
// OPEN_SERVICE_BROKER_AUTH_TYPE=basic
// OPEN_SERVICE_BROKER_USERNAME=<username>
// OPEN_SERVICE_BROKER_PASSWORD=<password>
// OPEN_SERVICE_BROKER_PLAN_ID=<The ID of the plan associated with the service offering>
// OPEN_SERVICE_BROKER_RESOURCE_INSTANCE_ID=<The ID of a previously provisioned service instance>
// OPEN_SERVICE_BROKER_SERVICE_ID=<The ID of the service being offered>
// OPEN_SERVICE_BROKER_ACCOUNT_ID=<User's account ID>
// OPEN_SERVICE_BROKER_BINDING_ID=<The ID of a previously provisioned binding for that service instance>
// OPEN_SERVICE_BROKER_SPACE_GUID=<The identifier for the project space within the IBM Cloud platform organization>
// OPEN_SERVICE_BROKER_APPLICATION_GUID=<GUID of an application associated with the binding>
// OPEN_SERVICE_BROKER_ORGANIZATION_GUID=<The IBM Cloud platform GUID for the organization under which the service instance is to be provisioned>
//
// These configuration properties can be exported as environment variables, or stored
// in a configuration file and then:
// export IBM_CREDENTIALS_FILE=<name of configuration file>
//

var _ = Describe(`OpenServiceBrokerV1 Examples Tests`, func() {
	const externalConfigFile = "../open_service_broker.env"

	var (
		openServiceBrokerService *openservicebrokerv1.OpenServiceBrokerV1
		config                   map[string]string
		configLoaded             bool = false

		instanceId  string
		orgGUID     string
		planId      string
		serviceId   string
		spaceGUID   string
		accountId   string
		bindingId   string
		appGUID     string
		initiatorId string = "null"
		operation   string = "Privision_45"
	)

	var shouldSkipTest = func() {
		if !configLoaded {
			Skip("External configuration is not available, skipping tests...")
		}
	}

	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			var err error
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping tests: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(openservicebrokerv1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}

			configLoaded = len(config) > 0

			instanceId = config["RESOURCE_INSTANCE_ID"]
			Expect(instanceId).ToNot(BeEmpty())

			orgGUID = config["ORGANIZATION_GUID"]
			Expect(orgGUID).ToNot(BeEmpty())

			planId = config["PLAN_ID"]
			Expect(planId).ToNot(BeEmpty())

			serviceId = config["SERVICE_ID"]
			Expect(serviceId).ToNot(BeEmpty())

			spaceGUID = config["SPACE_GUID"]
			Expect(spaceGUID).ToNot(BeEmpty())

			accountId = config["ACCOUNT_ID"]
			Expect(accountId).ToNot(BeEmpty())

			bindingId = config["BINDING_ID"]
			Expect(bindingId).ToNot(BeEmpty())

			appGUID = config["APPLICATION_GUID"]
			Expect(appGUID).ToNot(BeEmpty())
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {
			var err error

			// begin-common

			options := &openservicebrokerv1.OpenServiceBrokerV1Options{}

			openServiceBrokerService, err = openservicebrokerv1.NewOpenServiceBrokerV1UsingExternalConfig(options)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(openServiceBrokerService).ToNot(BeNil())
		})
	})

	Describe(`OpenServiceBrokerV1 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetServiceInstanceState request example`, func() {
			fmt.Println("\nGetServiceInstanceState() result:")
			// begin-get_service_instance_state

			options := openServiceBrokerService.NewGetServiceInstanceStateOptions(
				instanceId,
			)

			result, response, err := openServiceBrokerService.GetServiceInstanceState(options)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(result, "", "  ")
			fmt.Println(string(b))

			// end-get_service_instance_state

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())

		})
		It(`ReplaceServiceInstanceState request example`, func() {
			fmt.Println("\nReplaceServiceInstanceState() result:")
			// begin-replace_service_instance_state

			options := openServiceBrokerService.NewReplaceServiceInstanceStateOptions(
				instanceId,
			)
			options = options.SetEnabled(false)
			options = options.SetInitiatorID(initiatorId)

			result, response, err := openServiceBrokerService.ReplaceServiceInstanceState(options)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(result, "", "  ")
			fmt.Println(string(b))

			// end-replace_service_instance_state

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())

		})
		It(`ReplaceServiceInstance request example`, func() {
			fmt.Println("\nReplaceServiceInstance() result:")
			// begin-replace_service_instance

			contextOpt := &openservicebrokerv1.Context{
				AccountID: &accountId,
				CRN:       &instanceId,
				Platform:  core.StringPtr("ibmcloud"),
			}
			paramsOpt := make(map[string]string, 0)
			paramsOpt["example"] = "property"

			options := openServiceBrokerService.NewReplaceServiceInstanceOptions(
				instanceId,
			)
			options = options.SetPlanID(planId)
			options = options.SetServiceID(serviceId)
			options = options.SetOrganizationGUID(orgGUID)
			options = options.SetSpaceGUID(spaceGUID)
			options = options.SetContext(contextOpt)
			options = options.SetParameters(paramsOpt)
			options = options.SetAcceptsIncomplete(true)

			result, response, err := openServiceBrokerService.ReplaceServiceInstance(options)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(result, "", "  ")
			fmt.Println(string(b))

			// end-replace_service_instance

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(result).ToNot(BeNil())

		})
		It(`UpdateServiceInstance request example`, func() {
			fmt.Println("\nUpdateServiceInstance() result:")
			// begin-update_service_instance

			contextOpt := &openservicebrokerv1.Context{
				AccountID: &accountId,
				CRN:       &instanceId,
				Platform:  core.StringPtr("ibmcloud"),
			}

			paramsOpt := make(map[string]string, 0)
			paramsOpt["example"] = "property"

			previousValues := make(map[string]string, 0)
			previousValues["plan_id"] = planId

			options := openServiceBrokerService.NewUpdateServiceInstanceOptions(
				instanceId,
			)
			options = options.SetPlanID(planId)
			options = options.SetServiceID(serviceId)
			options = options.SetContext(contextOpt)
			options = options.SetParameters(paramsOpt)
			options = options.SetAcceptsIncomplete(true)

			result, response, err := openServiceBrokerService.UpdateServiceInstance(options)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(result, "", "  ")
			fmt.Println(string(b))

			// end-update_service_instance

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())

		})
		It(`ListCatalog request example`, func() {
			fmt.Println("\nListCatalog() result:")
			// begin-list_catalog

			options := openServiceBrokerService.NewListCatalogOptions()

			result, response, err := openServiceBrokerService.ListCatalog(options)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(result, "", "  ")
			fmt.Println(string(b))

			// end-list_catalog

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())

		})
		It(`GetLastOperation request example`, func() {
			fmt.Println("\nGetLastOperation() result:")
			// begin-get_last_operation

			options := openServiceBrokerService.NewGetLastOperationOptions(
				instanceId,
			)
			options = options.SetOperation(operation)
			options = options.SetPlanID(planId)
			options = options.SetServiceID(serviceId)

			result, response, err := openServiceBrokerService.GetLastOperation(options)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(result, "", "  ")
			fmt.Println(string(b))

			// end-get_last_operation

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())

		})
		It(`ReplaceServiceBinding request example`, func() {
			fmt.Println("\nReplaceServiceBinding() result:")
			// begin-replace_service_binding

			paramsOpt := make(map[string]string, 0)
			paramsOpt["example"] = "property"

			bindResource := &openservicebrokerv1.BindResource{
				AccountID:    &accountId,
				ServiceidCRN: &appGUID,
			}

			options := openServiceBrokerService.NewReplaceServiceBindingOptions(
				bindingId,
				instanceId,
			)
			options = options.SetPlanID(planId)
			options = options.SetServiceID(serviceId)
			options = options.SetParameters(paramsOpt)
			options = options.SetBindResource(bindResource)

			result, response, err := openServiceBrokerService.ReplaceServiceBinding(options)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(result, "", "  ")
			fmt.Println(string(b))

			// end-replace_service_binding

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(result).ToNot(BeNil())

		})
		It(`DeleteServiceInstance request example`, func() {
			fmt.Println("\nDeleteServiceInstance() result:")
			// begin-delete_service_instance

			options := openServiceBrokerService.NewDeleteServiceInstanceOptions(
				serviceId,
				planId,
				instanceId,
			)

			result, response, err := openServiceBrokerService.DeleteServiceInstance(options)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(result, "", "  ")
			fmt.Println(string(b))

			// end-delete_service_instance

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())

		})
		It(`DeleteServiceBinding request example`, func() {
			// begin-delete_service_binding

			deleteServiceBindingOptions := openServiceBrokerService.NewDeleteServiceBindingOptions(
				bindingId,
				instanceId,
				planId,
				serviceId,
			)

			response, err := openServiceBrokerService.DeleteServiceBinding(deleteServiceBindingOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_service_binding
			fmt.Printf("\nDeleteServiceBinding() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

		})
	})
})
