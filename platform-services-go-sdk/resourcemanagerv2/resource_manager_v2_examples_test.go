//go:build examples

/**
 * (C) Copyright IBM Corp. 2021.
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

package resourcemanagerv2_test

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/resourcemanagerv2"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//
// This file provides an example of how to use the Resource Manager service.
//
// The following configuration properties are assumed to be defined:
// RESOURCE_MANAGER_URL=<service base url>
// RESOURCE_MANAGER_AUTH_TYPE=iam
// RESOURCE_MANAGER_APIKEY=<IAM apikey of the service>
// RESOURCE_MANAGER_AUTH_URL=<IAM token service base URL - omit this if using the production environment>
// RESOURCE_MANAGER_QUOTA_ID=<quota ID>
// RESOURCE_MANAGER_USER_ACCOUNT_ID=<account ID of the user with delete permission>
//
// ALT_RESOURCE_MANAGER_URL=<service base url>
// ALT_RESOURCE_MANAGER_AUTH_TYPE=iam
// ALT_RESOURCE_MANAGER_AUTH_URL=<IAM token service base URL - omit this if using the production environment>
// ALT_RESOURCE_MANAGER_APIKEY=<IAM apikey of the user with delete permission>
//
// These configuration properties can be exported as environment variables, or stored
// in a configuration file and then:
// export IBM_CREDENTIALS_FILE=<name of configuration file>
//

var _ = Describe(`ResourceManagerV2 Examples Tests`, func() {
	const externalConfigFile = "../resource_manager.env"

	var (
		resourceManagerService       *resourcemanagerv2.ResourceManagerV2
		deleteResourceManagerService *resourcemanagerv2.ResourceManagerV2
		config                       map[string]string
		configLoaded                 bool = false

		exampleQuotaID       string
		exampleUserAccountID string

		resourceGroupID string
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
			config, err = core.GetServiceProperties(resourcemanagerv2.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}

			exampleQuotaID = config["QUOTA_ID"]
			Expect(exampleQuotaID).ToNot(BeEmpty())

			exampleUserAccountID = config["USER_ACCOUNT_ID"]
			Expect(exampleUserAccountID).ToNot(BeEmpty())

			configLoaded = len(config) > 0
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {
			var err error

			// begin-common

			resourceManagerServiceOptions := &resourcemanagerv2.ResourceManagerV2Options{
				ServiceName: resourcemanagerv2.DefaultServiceName,
			}

			resourceManagerService, err = resourcemanagerv2.NewResourceManagerV2UsingExternalConfig(resourceManagerServiceOptions)

			if err != nil {
				panic(err)
			}

			deleteResourceManagerServiceOptions := &resourcemanagerv2.ResourceManagerV2Options{
				ServiceName: "ALT_RESOURCE_MANAGER",
			}

			deleteResourceManagerService, err = resourcemanagerv2.NewResourceManagerV2UsingExternalConfig(deleteResourceManagerServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(resourceManagerService).ToNot(BeNil())
		})
	})

	Describe(`ResourceManagerV2 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateResourceGroup request example`, func() {
			Expect(exampleUserAccountID).NotTo(BeNil())

			fmt.Println("\nCreateResourceGroup() result:")
			// begin-create_resource_group

			createResourceGroupOptions := resourceManagerService.NewCreateResourceGroupOptions()
			createResourceGroupOptions.SetAccountID(exampleUserAccountID)
			createResourceGroupOptions.SetName("ExampleGroup")

			resCreateResourceGroup, response, err := resourceManagerService.CreateResourceGroup(createResourceGroupOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(resCreateResourceGroup, "", "  ")
			fmt.Println(string(b))

			// end-create_resource_group

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(resCreateResourceGroup).ToNot(BeNil())

			resourceGroupID = *resCreateResourceGroup.ID
		})
		It(`GetResourceGroup request example`, func() {
			Expect(resourceGroupID).NotTo(BeNil())

			fmt.Println("\nGetResourceGroup() result:")
			// begin-get_resource_group

			getResourceGroupOptions := resourceManagerService.NewGetResourceGroupOptions(
				resourceGroupID,
			)

			resourceGroup, response, err := resourceManagerService.GetResourceGroup(getResourceGroupOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(resourceGroup, "", "  ")
			fmt.Println(string(b))

			// end-get_resource_group

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(resourceGroup).ToNot(BeNil())
		})
		It(`UpdateResourceGroup request example`, func() {
			Expect(resourceGroupID).NotTo(BeNil())

			fmt.Println("\nUpdateResourceGroup() result:")
			// begin-update_resource_group

			updateResourceGroupOptions := resourceManagerService.NewUpdateResourceGroupOptions(
				resourceGroupID,
			)
			updateResourceGroupOptions.SetName("RenamedExampleGroup")
			updateResourceGroupOptions.SetState("ACTIVE")

			resourceGroup, response, err := resourceManagerService.UpdateResourceGroup(updateResourceGroupOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(resourceGroup, "", "  ")
			fmt.Println(string(b))

			// end-update_resource_group

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(resourceGroup).ToNot(BeNil())
		})
		It(`ListResourceGroups request example`, func() {
			Expect(exampleUserAccountID).NotTo(BeNil())

			fmt.Println("\nListResourceGroups() result:")
			// begin-list_resource_groups

			listResourceGroupsOptions := resourceManagerService.NewListResourceGroupsOptions()
			listResourceGroupsOptions.SetAccountID(exampleUserAccountID)
			listResourceGroupsOptions.SetIncludeDeleted(true)

			resourceGroupList, response, err := resourceManagerService.ListResourceGroups(listResourceGroupsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(resourceGroupList, "", "  ")
			fmt.Println(string(b))

			// end-list_resource_groups

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(resourceGroupList).ToNot(BeNil())
		})
		It(`DeleteResourceGroup request example`, func() {
			Expect(resourceGroupID).NotTo(BeNil())

			// begin-delete_resource_group

			deleteResourceGroupOptions := resourceManagerService.NewDeleteResourceGroupOptions(
				resourceGroupID,
			)

			response, err := deleteResourceManagerService.DeleteResourceGroup(deleteResourceGroupOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_resource_group

			fmt.Printf("\nDeleteResourceGroup() response status code: %d\n", response.StatusCode)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`GetQuotaDefinition request example`, func() {
			Expect(exampleQuotaID).NotTo(BeNil())

			fmt.Println("\nGetQuotaDefinition() result:")
			// begin-get_quota_definition

			getQuotaDefinitionOptions := resourceManagerService.NewGetQuotaDefinitionOptions(
				exampleQuotaID,
			)

			quotaDefinition, response, err := resourceManagerService.GetQuotaDefinition(getQuotaDefinitionOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(quotaDefinition, "", "  ")
			fmt.Println(string(b))

			// end-get_quota_definition

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(quotaDefinition).ToNot(BeNil())
		})
		It(`ListQuotaDefinitions request example`, func() {
			fmt.Println("\nListQuotaDefinitions() result:")
			// begin-list_quota_definitions

			listQuotaDefinitionsOptions := resourceManagerService.NewListQuotaDefinitionsOptions()

			quotaDefinitionList, response, err := resourceManagerService.ListQuotaDefinitions(listQuotaDefinitionsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(quotaDefinitionList, "", "  ")
			fmt.Println(string(b))

			// end-list_quota_definitions

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(quotaDefinitionList).ToNot(BeNil())
		})
	})
})
