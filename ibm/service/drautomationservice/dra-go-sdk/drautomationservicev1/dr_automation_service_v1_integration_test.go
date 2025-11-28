//go:build integration

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

package drautomationservicev1_test

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.ibm.com/DRAutomation/dra-go-sdk/drautomationservicev1"
)

/**
 * This file contains an integration test for the drautomationservicev1 package.
 *
 * Notes:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 */

var _ = Describe(`DrAutomationServiceV1 Integration Tests`, func() {
	const externalConfigFile = "../dr_automation_service_v1.env"

	var (
		err                        error
		drAutomationServiceService *drautomationservicev1.DrAutomationServiceV1
		serviceURL                 string
		config                     map[string]string
	)

	var shouldSkipTest = func() {
		Skip("External configuration is not available, skipping tests...")
	}

	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping tests: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(drautomationservicev1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}
			serviceURL = config["URL"]
			if serviceURL == "" {
				Skip("Unable to load service URL configuration property, skipping tests")
			}

			fmt.Fprintf(GinkgoWriter, "Service URL: %v\n", serviceURL)
			shouldSkipTest = func() {}
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {
			drAutomationServiceServiceOptions := &drautomationservicev1.DrAutomationServiceV1Options{}

			drAutomationServiceService, err = drautomationservicev1.NewDrAutomationServiceV1UsingExternalConfig(drAutomationServiceServiceOptions)
			Expect(err).To(BeNil())
			Expect(drAutomationServiceService).ToNot(BeNil())
			Expect(drAutomationServiceService.Service.Options.URL).To(Equal(serviceURL))

			core.SetLogger(core.NewLogger(core.LevelDebug, log.New(GinkgoWriter, "", log.LstdFlags), log.New(GinkgoWriter, "", log.LstdFlags)))
			drAutomationServiceService.EnableRetries(4, 30*time.Second)
		})
	})

	Describe(`GetApikey - Validates whether current apikey is valid or not`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetApikey(getApikeyOptions *GetApikeyOptions)`, func() {
			getApikeyOptions := &drautomationservicev1.GetApikeyOptions{
				InstanceID:     core.StringPtr("crn:v1:staging:public:power-dr-automation:global:a/a123456fb04ceebfb4a9fd38c22334455:123456d3-1122-3344-b67d-4389b44b7bf9::"),
				AcceptLanguage: core.StringPtr("testString"),
				IfNoneMatch:    core.StringPtr("testString"),
			}

			validationKeyResponse, response, err := drAutomationServiceService.GetApikey(getApikeyOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(validationKeyResponse).ToNot(BeNil())
		})
	})

	Describe(`CreateApikey - validate key api`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateApikey(createApikeyOptions *CreateApikeyOptions)`, func() {
			createApikeyOptions := &drautomationservicev1.CreateApikeyOptions{
				InstanceID:     core.StringPtr("crn:v1:staging:public:power-dr-automation:global:a/a123456fb04ceebfb4a9fd38c22334455:123456d3-1122-3344-b67d-4389b44b7bf9::"),
				APIKey:         core.StringPtr("abcdefrg_izklmnop_fxbEED"),
				AcceptLanguage: core.StringPtr("testString"),
				IfNoneMatch:    core.StringPtr("testString"),
			}

			validationKeyResponse, response, err := drAutomationServiceService.CreateApikey(createApikeyOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(validationKeyResponse).ToNot(BeNil())
		})
	})

	Describe(`UpdateApikey - Updates the API key for the specified service instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateApikey(updateApikeyOptions *UpdateApikeyOptions)`, func() {
			updateApikeyOptions := &drautomationservicev1.UpdateApikeyOptions{
				InstanceID:     core.StringPtr("crn:v1:staging:public:power-dr-automation:global:a/a123456fb04ceebfb4a9fd38c22334455:123456d3-1122-3344-b67d-4389b44b7bf9::"),
				APIKey:         core.StringPtr("adfadfdsafsdfdsf"),
				AcceptLanguage: core.StringPtr("testString"),
				IfNoneMatch:    core.StringPtr("testString"),
			}

			validationKeyResponse, response, err := drAutomationServiceService.UpdateApikey(updateApikeyOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(validationKeyResponse).ToNot(BeNil())
		})
	})

	Describe(`GetDrGrsLocationPair - Get GRS location pairs based on managed vms`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetDrGrsLocationPair(getDrGrsLocationPairOptions *GetDrGrsLocationPairOptions)`, func() {
			getDrGrsLocationPairOptions := &drautomationservicev1.GetDrGrsLocationPairOptions{
				InstanceID:     core.StringPtr("crn:v1:staging:public:power-dr-automation:global:a/a123456fb04ceebfb4a9fd38c22334455:123456d3-1122-3344-b67d-4389b44b7bf9::"),
				AcceptLanguage: core.StringPtr("testString"),
				IfNoneMatch:    core.StringPtr("testString"),
			}

			getGrsLocationPairResponse, response, err := drAutomationServiceService.GetDrGrsLocationPair(getDrGrsLocationPairOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(getGrsLocationPairResponse).ToNot(BeNil())
		})
	})

	Describe(`GetDrLocations - Get Disaster recovery locations`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetDrLocations(getDrLocationsOptions *GetDrLocationsOptions)`, func() {
			getDrLocationsOptions := &drautomationservicev1.GetDrLocationsOptions{
				InstanceID:     core.StringPtr("crn:v1:staging:public:power-dr-automation:global:a/a123456fb04ceebfb4a9fd38c22334455:123456d3-1122-3344-b67d-4389b44b7bf9::"),
				AcceptLanguage: core.StringPtr("testString"),
				IfNoneMatch:    core.StringPtr("testString"),
			}

			getDrLocationsResponse, response, err := drAutomationServiceService.GetDrLocations(getDrLocationsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(getDrLocationsResponse).ToNot(BeNil())
		})
	})

	Describe(`GetDrManagedVM - Get managed vms for the instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetDrManagedVM(getDrManagedVMOptions *GetDrManagedVMOptions)`, func() {
			getDrManagedVMOptions := &drautomationservicev1.GetDrManagedVMOptions{
				InstanceID:     core.StringPtr("crn:v1:staging:public:power-dr-automation:global:a/a123456fb04ceebfb4a9fd38c22334455:123456d3-1122-3344-b67d-4389b44b7bf9::"),
				AcceptLanguage: core.StringPtr("testString"),
				IfNoneMatch:    core.StringPtr("testString"),
			}

			managedVMMapResponse, response, err := drAutomationServiceService.GetDrManagedVM(getDrManagedVMOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(managedVMMapResponse).ToNot(BeNil())
		})
	})

	Describe(`GetDrSummary - Disaster recovery deployment details`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetDrSummary(getDrSummaryOptions *GetDrSummaryOptions)`, func() {
			getDrSummaryOptions := &drautomationservicev1.GetDrSummaryOptions{
				InstanceID:     core.StringPtr("crn:v1:staging:public:power-dr-automation:global:a/a123456fb04ceebfb4a9fd38c22334455:123456d3-1122-3344-b67d-4389b44b7bf9::"),
				AcceptLanguage: core.StringPtr("testString"),
				IfNoneMatch:    core.StringPtr("testString"),
			}

			drAutomationGetSummaryResponse, response, err := drAutomationServiceService.GetDrSummary(getDrSummaryOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(drAutomationGetSummaryResponse).ToNot(BeNil())
		})
	})

	Describe(`GetMachineType - Get MachineTypes based on selected workspaces`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetMachineType(getMachineTypeOptions *GetMachineTypeOptions)`, func() {
			getMachineTypeOptions := &drautomationservicev1.GetMachineTypeOptions{
				InstanceID:           core.StringPtr("crn:v1:staging:public:power-dr-automation:global:a/a123456fb04ceebfb4a9fd38c22334455:123456d3-1122-3344-b67d-4389b44b7bf9::"),
				PrimaryWorkspaceName: core.StringPtr("Test-workspace-wdc06"),
				AcceptLanguage:       core.StringPtr("testString"),
				IfNoneMatch:          core.StringPtr("testString"),
				StandbyWorkspaceName: core.StringPtr("Test-workspace-wdc07"),
			}

			machineTypesByWorkspace, response, err := drAutomationServiceService.GetMachineType(getMachineTypeOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(machineTypesByWorkspace).ToNot(BeNil())
		})
	})

	Describe(`GetPowervsWorkspaces - List of primary and standby powervs workspaces`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetPowervsWorkspaces(getPowervsWorkspacesOptions *GetPowervsWorkspacesOptions)`, func() {
			getPowervsWorkspacesOptions := &drautomationservicev1.GetPowervsWorkspacesOptions{
				InstanceID:  core.StringPtr("crn:v1:staging:public:power-dr-automation:global:a/a123456fb04ceebfb4a9fd38c22334455:123456d3-1122-3344-b67d-4389b44b7bf9::"),
				LocationID:  core.StringPtr("testString"),
				IfNoneMatch: core.StringPtr("testString"),
			}

			drData, response, err := drAutomationServiceService.GetPowervsWorkspaces(getPowervsWorkspacesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(drData).ToNot(BeNil())
		})
	})

	Describe(`GetManageDr - View configured DR automation details`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetManageDr(getManageDrOptions *GetManageDrOptions)`, func() {
			getManageDrOptions := &drautomationservicev1.GetManageDrOptions{
				InstanceID:     core.StringPtr("crn:v1:staging:public:power-dr-automation:global:a/a123456fb04ceebfb4a9fd38c22334455:123456d3-1122-3344-b67d-4389b44b7bf9::"),
				AcceptLanguage: core.StringPtr("testString"),
				IfNoneMatch:    core.StringPtr("testString"),
			}

			serviceInstanceManageDr, response, err := drAutomationServiceService.GetManageDr(getManageDrOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(serviceInstanceManageDr).ToNot(BeNil())
		})
	})

	Describe(`CreateManageDr - Create DR Deployment`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateManageDr(createManageDrOptions *CreateManageDrOptions)`, func() {
			createManageDrOptions := &drautomationservicev1.CreateManageDrOptions{
				InstanceID:                           core.StringPtr("crn:v1:staging:public:power-dr-automation:global:a/a123456fb04ceebfb4a9fd38c22334455:123456d3-1122-3344-b67d-4389b44b7bf9::"),
				APIKey:                               core.StringPtr("testString"),
				ClientID:                             core.StringPtr("d9f2c83a-97d2-4b14-bf62-8eaecc67a122"),
				ClientSecret:                         core.StringPtr("N8lQ4tP2xM1yT5rS8wK6qR9dD7vF1hU4sA3bE2jG0pL9oX7yC"),
				GUID:                                 core.StringPtr("123e4567-e89b-12d3-a456-426614174000"),
				LocationID:                           core.StringPtr("dal10"),
				MachineType:                          core.StringPtr("bx2-4x16"),
				OrchestratorHa:                       core.BoolPtr(true),
				OrchestratorLocationType:             core.StringPtr("off-premises"),
				OrchestratorName:                     core.StringPtr("adminUser"),
				OrchestratorPassword:                 core.StringPtr("testString"),
				OrchestratorWorkspaceID:              core.StringPtr("orch-workspace-01"),
				OrchestratorWorkspaceLocation:        core.StringPtr("us-south"),
				ProxyIP:                              core.StringPtr("10.40.30.10:8888"),
				RegionID:                             core.StringPtr("us-south"),
				ResourceInstance:                     core.StringPtr("crn:v1:bluemix:public:resource-controller::res123"),
				SecondaryWorkspaceID:                 core.StringPtr("secondary-workspace789"),
				Secret:                               core.StringPtr("testString"),
				SecretGroup:                          core.StringPtr("default-secret-group"),
				SSHKeyName:                           core.StringPtr("my-ssh-key"),
				StandbyMachineType:                   core.StringPtr("bx2-8x32"),
				StandbyOrchestratorName:              core.StringPtr("standbyAdmin"),
				StandbyOrchestratorWorkspaceID:       core.StringPtr("orch-standby-02"),
				StandbyOrchestratorWorkspaceLocation: core.StringPtr("us-east"),
				StandbyTier:                          core.StringPtr("Premium"),
				TenantName:                           core.StringPtr("xxx.ibm.com"),
				Tier:                                 core.StringPtr("Standard"),
				StandByRedeploy:                      core.StringPtr("testString"),
				AcceptLanguage:                       core.StringPtr("testString"),
				IfNoneMatch:                          core.StringPtr("testString"),
				AcceptsIncomplete:                    core.BoolPtr(true),
			}

			serviceInstanceManageDr, response, err := drAutomationServiceService.CreateManageDr(createManageDrOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(serviceInstanceManageDr).ToNot(BeNil())
		})
	})

	Describe(`GetLastOperation - View details of Last operation performed on the instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetLastOperation(getLastOperationOptions *GetLastOperationOptions)`, func() {
			getLastOperationOptions := &drautomationservicev1.GetLastOperationOptions{
				InstanceID:     core.StringPtr("crn:v1:staging:public:power-dr-automation:global:a/a123456fb04ceebfb4a9fd38c22334455:123456d3-1122-3344-b67d-4389b44b7bf9::"),
				AcceptLanguage: core.StringPtr("testString"),
				IfNoneMatch:    core.StringPtr("testString"),
			}

			serviceInstanceStatus, response, err := drAutomationServiceService.GetLastOperation(getLastOperationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(serviceInstanceStatus).ToNot(BeNil())
		})
	})

	Describe(`ListEvents - Get events from the cloud instance since a specific timestamp`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListEvents(listEventsOptions *ListEventsOptions)`, func() {
			listEventsOptions := &drautomationservicev1.ListEventsOptions{
				InstanceID:     core.StringPtr("crn:v1:staging:public:power-dr-automation:global:a/a123456fb04ceebfb4a9fd38c22334455:123456d3-1122-3344-b67d-4389b44b7bf9::"),
				Time:           core.StringPtr("2025-06-19T23:59:59Z"),
				FromTime:       core.StringPtr("2025-06-19T00:00:00Z"),
				ToTime:         core.StringPtr("2025-06-19T23:59:59Z"),
				AcceptLanguage: core.StringPtr("testString"),
				IfNoneMatch:    core.StringPtr("testString"),
			}

			eventCollection, response, err := drAutomationServiceService.ListEvents(listEventsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(eventCollection).ToNot(BeNil())
		})
	})

	Describe(`GetEvent - Get a single event`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetEvent(getEventOptions *GetEventOptions)`, func() {
			getEventOptions := &drautomationservicev1.GetEventOptions{
				InstanceID:     core.StringPtr("crn:v1:staging:public:power-dr-automation:global:a/a123456fb04ceebfb4a9fd38c22334455:123456d3-1122-3344-b67d-4389b44b7bf9::"),
				EventID:        core.StringPtr("00116b2a-9326-4024-839e-fb5364b76898"),
				AcceptLanguage: core.StringPtr("testString"),
				IfNoneMatch:    core.StringPtr("testString"),
			}

			event, response, err := drAutomationServiceService.GetEvent(getEventOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(event).ToNot(BeNil())
		})
	})
})

//
// Utility functions are declared in the unit test file
//
