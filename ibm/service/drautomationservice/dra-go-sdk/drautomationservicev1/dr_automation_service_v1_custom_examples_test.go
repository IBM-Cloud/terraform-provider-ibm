//go:build customexamples
// +build customexamples

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

		// -------- HA Cases ---------

		// 1. HA with sshkey
		It(`ServiceInstanceManageDr HA with sshkey`, func() {
			fmt.Println("\ncreate_manage_dr_ha_with_sshkey() result:")
			// begin-create_manage_dr_ha_with_sshkey
			createManageDrOptions := drAutomationServiceService.NewCreateManageDrOptions(
				"crn:v1:staging:public:power-dr-automation:global:a/a09202c1bfb04ceebfb4a9fd38c87721:050ebe3b-13f4-4db8-8ece-501a3c13be80mh1::",
			)
			createManageDrOptions.SetOrchestratorHa(true)
			createManageDrOptions.SetOrchestratorLocationType("off-premises")
			createManageDrOptions.SetLocationID("dal10")
			createManageDrOptions.SetOrchestratorWorkspaceID("75cbf05b-78f6-406e-afe7-a904f646d798")
			createManageDrOptions.SetOrchestratorName("drautomationprimarymh1")
			createManageDrOptions.SetOrchestratorPassword("EverytimeNewPassword@1")
			createManageDrOptions.SetMachineType("s922")
			createManageDrOptions.SetTier("tier1")
			createManageDrOptions.SetSSHKeyName("vijaykey")
			createManageDrOptions.SetAPIKey("apikey is required")
			// Standby fields (only for HA)
			createManageDrOptions.SetStandbyOrchestratorName("drautomationstandbymh1")
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
			// end-create_manage_dr_ha_with_sshkey

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(serviceInstanceManageDr).ToNot(BeNil())
		})

		// 2. HA with secrets
		It(`ServiceInstanceManageDr HA with secrets`, func() {
			fmt.Println("\ncreate_manage_dr_ha_with_secrets() result:")
			// begin-create_manage_dr_ha_with_secrets

			createManageDrOptions := drAutomationServiceService.NewCreateManageDrOptions(
				"crn:v1:staging:public:power-dr-automation:global:a/a09202c1bfb04ceebfb4a9fd38c87721:050ebe3b-13f4-4db8-8ece-501a3c13be80mh3::",
			)
			createManageDrOptions.SetOrchestratorHa(true)
			createManageDrOptions.SetOrchestratorLocationType("off-premises")
			createManageDrOptions.SetLocationID("dal10")
			createManageDrOptions.SetOrchestratorWorkspaceID("75cbf05b-78f6-406e-afe7-a904f646d798")
			createManageDrOptions.SetOrchestratorName("drautomationprimarymh3")
			createManageDrOptions.SetOrchestratorPassword("EverytimeNewPassword@1")
			createManageDrOptions.SetMachineType("s922")
			createManageDrOptions.SetTier("tier1")
			createManageDrOptions.SetGUID("397dc20d-9f66-46dc-a750-d15392872023")
			createManageDrOptions.SetSecretGroup("12345-714f-86a6-6a50-2f128a4e7ac2")
			createManageDrOptions.SetSecret("12345-997c-1d0d-5503-27ca856f2b5a")
			createManageDrOptions.SetRegionID("us-south")
			createManageDrOptions.SetAPIKey("apikey is required")
			// Standby fields (only for HA)
			createManageDrOptions.SetStandbyOrchestratorName("drautomationstandbymh3")
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
			// end-create_manage_dr_ha_with_secrets

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(serviceInstanceManageDr).ToNot(BeNil())
		})

		// -------- Non-HA Cases ---------

		// 3. Non-HA with sshkey
		It(`ServiceInstanceManageDr Non-HA with sshkey`, func() {
			fmt.Println("\ncreate_manage_dr_nonha_with_sshkey() result:")
			// begin-create_manage_dr_nonha_with_sshkey

			createManageDrOptions := drAutomationServiceService.NewCreateManageDrOptions(
				"crn:v1:staging:public:power-dr-automation:global:a/a09202c1bfb04ceebfb4a9fd38c87721:050ebe3b-13f4-4db8-8ece-501a3c13be80mnh5::",
			)
			createManageDrOptions.SetOrchestratorHa(false)
			createManageDrOptions.SetOrchestratorLocationType("off-premises")
			createManageDrOptions.SetLocationID("dal10")
			createManageDrOptions.SetOrchestratorWorkspaceID("75cbf05b-78f6-406e-afe7-a904f646d798")
			createManageDrOptions.SetOrchestratorName("drautomationprimarymnh5")
			createManageDrOptions.SetOrchestratorPassword("EverytimeNewPassword@1")
			createManageDrOptions.SetMachineType("s922")
			createManageDrOptions.SetTier("tier1")
			createManageDrOptions.SetSSHKeyName("vijaykey")
			createManageDrOptions.SetAPIKey("apikey is required")
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
			// end-create_manage_dr_nonha_with_sshkey

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(serviceInstanceManageDr).ToNot(BeNil())
		})

		// 4. Non-HA with secrets
		It(`ServiceInstanceManageDr Non-HA with secrets`, func() {
			fmt.Println("\ncreate_manage_dr_nonha_with_secrets() result:")
			// begin-create_manage_dr_nonha_with_secrets

			createManageDrOptions := drAutomationServiceService.NewCreateManageDrOptions(
				"crn:v1:staging:public:power-dr-automation:global:a/a09202c1bfb04ceebfb4a9fd38c87721:050ebe3b-13f4-4db8-8ece-501a3c13be80mnh7::",
			)
			createManageDrOptions.SetOrchestratorHa(false)
			createManageDrOptions.SetOrchestratorLocationType("off-premises")
			createManageDrOptions.SetLocationID("dal10")
			createManageDrOptions.SetOrchestratorWorkspaceID("75cbf05b-78f6-406e-afe7-a904f646d798")
			createManageDrOptions.SetOrchestratorName("drautomationprimarymnh7")
			createManageDrOptions.SetOrchestratorPassword("EverytimeNewPassword@1")
			createManageDrOptions.SetMachineType("s922")
			createManageDrOptions.SetTier("tier1")
			createManageDrOptions.SetGUID("397dc20d-9f66-46dc-a750-d15392872023")
			createManageDrOptions.SetSecretGroup("12345-714f-86a6-6a50-2f128a4e7ac2")
			createManageDrOptions.SetSecret("12345-997c-1d0d-5503-27ca856f2b5a")
			createManageDrOptions.SetRegionID("us-south")
			createManageDrOptions.SetAPIKey("apikey is required")
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
			// end-create_manage_dr_nonha_with_secrets

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(serviceInstanceManageDr).ToNot(BeNil())
		})
	})
})
