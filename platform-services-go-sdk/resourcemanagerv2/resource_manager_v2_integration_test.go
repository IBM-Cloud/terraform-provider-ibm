//go:build integration

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
package resourcemanagerv2_test

import (
	"log"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/resourcemanagerv2"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Resource Manager - Integration Tests", func() {

	const externalConfigFile = "../resource_manager.env"

	var (
		service            *resourcemanagerv2.ResourceManagerV2
		altService         *resourcemanagerv2.ResourceManagerV2
		err                error
		config             map[string]string
		testQuotaID        string
		testUserAccountID  string
		newResourceGroupID string
		configLoaded       bool = false
	)

	var shouldSkipTest = func() {
		if !configLoaded {
			Skip("External configuration is not available, skipping...")
		}
	}

	It("Successfully load the configuration", func() {
		_, err = os.Stat(externalConfigFile)
		if err == nil {
			err = os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			if err == nil {
				configLoaded = true
			}
			config, err = core.GetServiceProperties(resourcemanagerv2.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}
			testQuotaID = config["QUOTA_ID"]
			if testQuotaID == "" {
				Skip("Unable to load test quota ID configuration property, skipping tests")
			}
			testUserAccountID = config["USER_ACCOUNT_ID"]
			if testUserAccountID == "" {
				Skip("Unable to test user account ID configuration property, skipping tests")
			}
		}
		if !configLoaded {
			Skip("External configuration could not be loaded, skipping...")
		}
	})
	It(`Successfully created ResourceManagerV2 service instances`, func() {
		shouldSkipTest()
		options := &resourcemanagerv2.ResourceManagerV2Options{
			ServiceName: resourcemanagerv2.DefaultServiceName,
		}
		service, err = resourcemanagerv2.NewResourceManagerV2UsingExternalConfig(options)
		Expect(err).To(BeNil())
		Expect(service).ToNot(BeNil())

		core.SetLogger(core.NewLogger(core.LevelDebug, log.New(GinkgoWriter, "", log.LstdFlags), log.New(GinkgoWriter, "", log.LstdFlags)))
		service.EnableRetries(4, 30*time.Second)

		optionsUser := &resourcemanagerv2.ResourceManagerV2Options{
			ServiceName: "ALT_RESOURCE_MANAGER",
		}
		altService, err = resourcemanagerv2.NewResourceManagerV2UsingExternalConfig(optionsUser)
		Expect(err).To(BeNil())
		Expect(altService).ToNot(BeNil())

		altService.EnableRetries(4, 30*time.Second)
	})

	It("Get list of all quota definition", func() {
		shouldSkipTest()
		listQuotaDefinitionOptionsModel := service.NewListQuotaDefinitionsOptions()
		result, detailedResponse, err := service.ListQuotaDefinitions(listQuotaDefinitionOptionsModel)
		Expect(err).To(BeNil())
		Expect(detailedResponse.StatusCode).To(Equal(200))
		Expect(result.Resources).NotTo(BeNil())
	})

	It("Get a quota definition by id", func() {
		shouldSkipTest()
		getQuotaDefinitionOptionsModel := service.NewGetQuotaDefinitionOptions(testQuotaID)
		result, detailedResponse, err := service.GetQuotaDefinition(getQuotaDefinitionOptionsModel)
		Expect(err).To(BeNil())
		Expect(detailedResponse.StatusCode).To(Equal(200))
		Expect(result).NotTo(BeNil())
	})

	Describe("Get a List of all resource groups in an account", func() {
		It("Successfully retrieved list of resource groups in an account", func() {
			shouldSkipTest()

			listResourceGroupsOptionsModel := service.NewListResourceGroupsOptions()
			listResourceGroupsOptionsModel.SetAccountID(testUserAccountID)
			result, detailedResponse, err := service.ListResourceGroups(listResourceGroupsOptionsModel)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(len(result.Resources)).To(BeNumerically(">=", 1))
			Expect(result.Resources[0]).NotTo(BeNil())
			Expect(result.Resources[0].ID).NotTo(BeNil())
			Expect(result.Resources[0].Name).NotTo(BeNil())
			Expect(result.Resources[0].CRN).NotTo(BeNil())
			Expect(result.Resources[0].AccountID).NotTo(BeNil())
			Expect(result.Resources[0].QuotaID).NotTo(BeNil())
			Expect(result.Resources[0].QuotaURL).NotTo(BeNil())
			Expect(result.Resources[0].CreatedAt).NotTo(BeNil())
			Expect(result.Resources[0].UpdatedAt).NotTo(BeNil())
		})
	})

	Describe("Create a new resource group in an account", func() {
		It("Successfully created new resource group in an account", func() {
			shouldSkipTest()

			createResourceGroupOptionsModel := service.NewCreateResourceGroupOptions()
			createResourceGroupOptionsModel.SetAccountID(testUserAccountID)
			createResourceGroupOptionsModel.SetName("TestGroup")
			result, detailedResponse, err := service.CreateResourceGroup(createResourceGroupOptionsModel)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(201))
			Expect(result).NotTo(BeNil())
			Expect(result.ID).NotTo(BeNil())
			newResourceGroupID = *result.ID
		})
	})

	Describe("Get a resource group by ID", func() {
		It("Successfully retrieved resource group by ID", func() {
			shouldSkipTest()

			getResourceGroupOptionsModel := service.NewGetResourceGroupOptions(newResourceGroupID)
			result, detailedResponse, err := service.GetResourceGroup(getResourceGroupOptionsModel)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(result).NotTo(BeNil())
		})
	})

	Describe("Update a resource group by ID", func() {
		It("Successfully updated resource group", func() {
			shouldSkipTest()

			updateResourceGroupOptionsModel := service.NewUpdateResourceGroupOptions(newResourceGroupID)
			result, detailedResponse, err := service.UpdateResourceGroup(updateResourceGroupOptionsModel)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(result).NotTo(BeNil())
		})
	})

	Describe("Delete a resource group by ID", func() {
		It("Successfully deleted resource group", func() {
			shouldSkipTest()

			deleteResourceGroupOptionsModel := altService.NewDeleteResourceGroupOptions(newResourceGroupID)
			detailedResponse, err := altService.DeleteResourceGroup(deleteResourceGroupOptionsModel)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(204))
		})
	})
})
