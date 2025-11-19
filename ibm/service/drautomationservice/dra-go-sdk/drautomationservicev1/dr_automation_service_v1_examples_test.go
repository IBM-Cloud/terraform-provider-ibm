//go:build examples

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
	"encoding/json"
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.ibm.com/DRAutomation/dra-go-sdk/drautomationservicev1"
)

// This file provides an example of how to use the DrAutomation Service service.
//
// The following configuration properties are assumed to be defined:
// DR_AUTOMATION_SERVICE_URL=<service base url>
// DR_AUTOMATION_SERVICE_AUTH_TYPE=iam
// DR_AUTOMATION_SERVICE_APIKEY=<IAM apikey>
// DR_AUTOMATION_SERVICE_AUTH_URL=<IAM token service base URL - omit this if using the production environment>
//
// These configuration properties can be exported as environment variables, or stored
// in a configuration file and then:
// export IBM_CREDENTIALS_FILE=<name of configuration file>
var _ = Describe(`DrAutomationServiceV1 Examples Tests`, func() {

	const externalConfigFile = "../dr_automation_service_v1.env"

	var (
		drAutomationServiceService *drautomationservicev1.DrAutomationServiceV1
		config                     map[string]string
	)

	var shouldSkipTest = func() {
		Skip("External configuration is not available, skipping examples...")
	}

	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			var err error
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping examples: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(drautomationservicev1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping examples: " + err.Error())
			} else if len(config) == 0 {
				Skip("Unable to load service properties, skipping examples")
			}

			shouldSkipTest = func() {}
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {
			var err error

			// begin-common

			drAutomationServiceServiceOptions := &drautomationservicev1.DrAutomationServiceV1Options{}

			drAutomationServiceService, err = drautomationservicev1.NewDrAutomationServiceV1UsingExternalConfig(drAutomationServiceServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(drAutomationServiceService).ToNot(BeNil())
		})
	})

	Describe(`DrAutomationServiceV1 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetApikey request example`, func() {
			fmt.Println("\nGetApikey() result:")
			// begin-get_apikey

			getApikeyOptions := drAutomationServiceService.NewGetApikeyOptions(
				"crn:v1:staging:public:power-dr-automation:global:a/a123456fb04ceebfb4a9fd38c22334455:123456d3-1122-3344-b67d-4389b44b7bf9::",
			)

			validationKeyResponse, response, err := drAutomationServiceService.GetApikey(getApikeyOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(validationKeyResponse, "", "  ")
			fmt.Println(string(b))

			// end-get_apikey

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(validationKeyResponse).ToNot(BeNil())
		})
		It(`CreateApikey request example`, func() {
			fmt.Println("\nCreateApikey() result:")
			// begin-create_apikey

			createApikeyOptions := drAutomationServiceService.NewCreateApikeyOptions(
				"crn:v1:staging:public:power-dr-automation:global:a/a123456fb04ceebfb4a9fd38c22334455:123456d3-1122-3344-b67d-4389b44b7bf9::",
				"abcdefrg_izklmnop_fxbEED",
			)

			validationKeyResponse, response, err := drAutomationServiceService.CreateApikey(createApikeyOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(validationKeyResponse, "", "  ")
			fmt.Println(string(b))

			// end-create_apikey

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(validationKeyResponse).ToNot(BeNil())
		})
		It(`UpdateApikey request example`, func() {
			fmt.Println("\nUpdateApikey() result:")
			// begin-update_apikey

			updateApikeyOptions := drAutomationServiceService.NewUpdateApikeyOptions(
				"crn:v1:staging:public:power-dr-automation:global:a/a123456fb04ceebfb4a9fd38c22334455:123456d3-1122-3344-b67d-4389b44b7bf9::",
				"adfadfdsafsdfdsf",
			)

			validationKeyResponse, response, err := drAutomationServiceService.UpdateApikey(updateApikeyOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(validationKeyResponse, "", "  ")
			fmt.Println(string(b))

			// end-update_apikey

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(validationKeyResponse).ToNot(BeNil())
		})
		It(`GetDrGrsLocationPair request example`, func() {
			fmt.Println("\nGetDrGrsLocationPair() result:")
			// begin-get_dr_grs_location_pair

			getDrGrsLocationPairOptions := drAutomationServiceService.NewGetDrGrsLocationPairOptions(
				"crn:v1:staging:public:power-dr-automation:global:a/a123456fb04ceebfb4a9fd38c22334455:123456d3-1122-3344-b67d-4389b44b7bf9::",
			)

			getGrsLocationPairResponse, response, err := drAutomationServiceService.GetDrGrsLocationPair(getDrGrsLocationPairOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(getGrsLocationPairResponse, "", "  ")
			fmt.Println(string(b))

			// end-get_dr_grs_location_pair

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(getGrsLocationPairResponse).ToNot(BeNil())
		})
		It(`GetDrLocations request example`, func() {
			fmt.Println("\nGetDrLocations() result:")
			// begin-get_dr_locations

			getDrLocationsOptions := drAutomationServiceService.NewGetDrLocationsOptions(
				"crn:v1:staging:public:power-dr-automation:global:a/a123456fb04ceebfb4a9fd38c22334455:123456d3-1122-3344-b67d-4389b44b7bf9::",
			)

			getDrLocationsResponse, response, err := drAutomationServiceService.GetDrLocations(getDrLocationsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(getDrLocationsResponse, "", "  ")
			fmt.Println(string(b))

			// end-get_dr_locations

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(getDrLocationsResponse).ToNot(BeNil())
		})
		It(`GetDrManagedVM request example`, func() {
			fmt.Println("\nGetDrManagedVM() result:")
			// begin-get_dr_managed_vm

			getDrManagedVMOptions := drAutomationServiceService.NewGetDrManagedVMOptions(
				"crn:v1:staging:public:power-dr-automation:global:a/a123456fb04ceebfb4a9fd38c22334455:123456d3-1122-3344-b67d-4389b44b7bf9::",
			)

			managedVMMapResponse, response, err := drAutomationServiceService.GetDrManagedVM(getDrManagedVMOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(managedVMMapResponse, "", "  ")
			fmt.Println(string(b))

			// end-get_dr_managed_vm

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(managedVMMapResponse).ToNot(BeNil())
		})
		It(`GetDrSummary request example`, func() {
			fmt.Println("\nGetDrSummary() result:")
			// begin-get_dr_summary

			getDrSummaryOptions := drAutomationServiceService.NewGetDrSummaryOptions(
				"crn:v1:staging:public:power-dr-automation:global:a/a123456fb04ceebfb4a9fd38c22334455:123456d3-1122-3344-b67d-4389b44b7bf9::",
			)

			drAutomationGetSummaryResponse, response, err := drAutomationServiceService.GetDrSummary(getDrSummaryOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(drAutomationGetSummaryResponse, "", "  ")
			fmt.Println(string(b))

			// end-get_dr_summary

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(drAutomationGetSummaryResponse).ToNot(BeNil())
		})
		It(`GetMachineType request example`, func() {
			fmt.Println("\nGetMachineType() result:")
			// begin-get_machine_type

			getMachineTypeOptions := drAutomationServiceService.NewGetMachineTypeOptions(
				"crn:v1:staging:public:power-dr-automation:global:a/a123456fb04ceebfb4a9fd38c22334455:123456d3-1122-3344-b67d-4389b44b7bf9::",
				"Test-workspace-wdc06",
			)
			getMachineTypeOptions.SetStandbyWorkspaceName("Test-workspace-wdc07")

			machineTypesByWorkspace, response, err := drAutomationServiceService.GetMachineType(getMachineTypeOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(machineTypesByWorkspace, "", "  ")
			fmt.Println(string(b))

			// end-get_machine_type

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(machineTypesByWorkspace).ToNot(BeNil())
		})
		It(`GetPowervsWorkspaces request example`, func() {
			fmt.Println("\nGetPowervsWorkspaces() result:")
			// begin-get_powervs_workspaces

			getPowervsWorkspacesOptions := drAutomationServiceService.NewGetPowervsWorkspacesOptions(
				"crn:v1:staging:public:power-dr-automation:global:a/a123456fb04ceebfb4a9fd38c22334455:123456d3-1122-3344-b67d-4389b44b7bf9::",
				"testString",
			)

			drData, response, err := drAutomationServiceService.GetPowervsWorkspaces(getPowervsWorkspacesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(drData, "", "  ")
			fmt.Println(string(b))

			// end-get_powervs_workspaces

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(drData).ToNot(BeNil())
		})
		It(`GetManageDr request example`, func() {
			fmt.Println("\nGetManageDr() result:")
			// begin-get_manage_dr

			getManageDrOptions := drAutomationServiceService.NewGetManageDrOptions(
				"crn:v1:staging:public:power-dr-automation:global:a/a123456fb04ceebfb4a9fd38c22334455:123456d3-1122-3344-b67d-4389b44b7bf9::",
			)

			serviceInstanceManageDr, response, err := drAutomationServiceService.GetManageDr(getManageDrOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(serviceInstanceManageDr, "", "  ")
			fmt.Println(string(b))

			// end-get_manage_dr

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(serviceInstanceManageDr).ToNot(BeNil())
		})
		It(`CreateManageDr request example`, func() {
			fmt.Println("\nCreateManageDr() result:")
			// begin-create_manage_dr

			createManageDrOptions := drAutomationServiceService.NewCreateManageDrOptions(
				"crn:v1:staging:public:power-dr-automation:global:a/a123456fb04ceebfb4a9fd38c22334455:123456d3-1122-3344-b67d-4389b44b7bf9::",
			)
			createManageDrOptions.SetOrchestratorHa(true)
			createManageDrOptions.SetOrchestratorLocationType("off-premises")
			createManageDrOptions.SetLocationID("dal10")
			createManageDrOptions.SetOrchestratorWorkspaceID("75cbf05b-78f6-406e-afe7-a904f646d798")
			createManageDrOptions.SetOrchestratorName("drautomationprimarybyh1105")
			createManageDrOptions.SetOrchestratorPassword("EverytimeNewPassword@1")
			createManageDrOptions.SetMachineType("s922")
			createManageDrOptions.SetTier("tier1")
			createManageDrOptions.SetSSHKeyName("vijaykey")
			createManageDrOptions.SetAPIKey("key should pass")
			// Standby fields (only for HA)
			createManageDrOptions.SetStandbyOrchestratorName("drautomationstandbyh1105")
			createManageDrOptions.SetStandbyOrchestratorWorkspaceID("71027b79-0e31-44f6-a499-63eca1a66feb")
			createManageDrOptions.SetStandbyMachineType("s922")
			createManageDrOptions.SetStandbyTier("tier1")
			createManageDrOptions.SetStandByRedeploy("false")
			// mfa
			createManageDrOptions.SetClientID("123abcd-97d2-4b14-bf62-8eaecc67a122")
			createManageDrOptions.SetClientSecret("abcdefgT5rS8wK6qR9dD7vF1hU4sA3bE2jG0pL9oX7yC")
			createManageDrOptions.SetTenantName("xxx.ibm.com")

			serviceInstanceManageDr, response, err := drAutomationServiceService.CreateManageDr(createManageDrOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(serviceInstanceManageDr, "", "  ")
			fmt.Println(string(b))

			// end-create_manage_dr

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(serviceInstanceManageDr).ToNot(BeNil())
		})
		It(`GetLastOperation request example`, func() {
			fmt.Println("\nGetLastOperation() result:")
			// begin-get_last_operation

			getLastOperationOptions := drAutomationServiceService.NewGetLastOperationOptions(
				"crn:v1:staging:public:power-dr-automation:global:a/a123456fb04ceebfb4a9fd38c22334455:123456d3-1122-3344-b67d-4389b44b7bf9::",
			)

			serviceInstanceStatus, response, err := drAutomationServiceService.GetLastOperation(getLastOperationOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(serviceInstanceStatus, "", "  ")
			fmt.Println(string(b))

			// end-get_last_operation

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(serviceInstanceStatus).ToNot(BeNil())
		})
		It(`ListEvents request example`, func() {
			fmt.Println("\nListEvents() result:")
			// begin-list_events

			listEventsOptions := drAutomationServiceService.NewListEventsOptions(
				"crn:v1:staging:public:power-dr-automation:global:a/a123456fb04ceebfb4a9fd38c22334455:123456d3-1122-3344-b67d-4389b44b7bf9::",
			)
			listEventsOptions.SetTime("2025-06-19T23:59:59Z")
			listEventsOptions.SetFromTime("2025-06-19T00:00:00Z")
			listEventsOptions.SetToTime("2025-06-19T23:59:59Z")

			eventCollection, response, err := drAutomationServiceService.ListEvents(listEventsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(eventCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_events

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(eventCollection).ToNot(BeNil())
		})
		It(`GetEvent request example`, func() {
			fmt.Println("\nGetEvent() result:")
			// begin-get_event

			getEventOptions := drAutomationServiceService.NewGetEventOptions(
				"crn:v1:staging:public:power-dr-automation:global:a/a123456fb04ceebfb4a9fd38c22334455:123456d3-1122-3344-b67d-4389b44b7bf9::",
				"00116b2a-9326-4024-839e-fb5364b76898",
			)

			event, response, err := drAutomationServiceService.GetEvent(getEventOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(event, "", "  ")
			fmt.Println(string(b))

			// end-get_event

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(event).ToNot(BeNil())
		})
	})
})
