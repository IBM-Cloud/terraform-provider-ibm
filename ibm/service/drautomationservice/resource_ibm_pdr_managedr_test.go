// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package drautomationservice_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/drautomationservice"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/stretchr/testify/assert"
	"github.ibm.com/DRAutomation/dra-go-sdk/drautomationservicev1"
)

func TestAccIBMPdrManagedrBasic(t *testing.T) {
	var conf drautomationservicev1.ServiceInstanceManageDr
	instanceID := "4a550c84-2bef-401f-bf6d-f6776eb3ec22"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPdrManagedrDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMPdrManagedrConfigBasic(instanceID),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMPdrManagedrExists("ibm_pdr_managedr.pdr_managedr_instance", conf),
					resource.TestCheckResourceAttr("ibm_pdr_managedr.pdr_managedr_instance", "instance_id", instanceID),
				),
			},
		},
	})
}

func TestAccIBMPdrManagedrAllArgs(t *testing.T) {
	var conf drautomationservicev1.ServiceInstanceManageDr
	instanceID := "4a550c84-2bef-401f-bf6d-f6776eb3ec22"
	standByRedeploy := "false"
	acceptLanguage := "en-us"
	acceptsIncomplete := "true"
	orchestratorLocationType := fmt.Sprintf("tf_orchestrator_location_type_%d", acctest.RandIntRange(10, 100))
	locationID := fmt.Sprintf("tf_location_id_%d", acctest.RandIntRange(10, 100))
	sshKeyName := fmt.Sprintf("tf_ssh_key_name_%d", acctest.RandIntRange(10, 100))
	standbySSHKeyName := fmt.Sprintf("tf_standby_ssh_key_name_%d", acctest.RandIntRange(10, 100))
	orchestratorName := fmt.Sprintf("tf_orchestrator_name_%d", acctest.RandIntRange(10, 100))
	orchestratorWorkspaceID := fmt.Sprintf("tf_orchestrator_workspace_id_%d", acctest.RandIntRange(10, 100))
	standbyOrchestratorName := fmt.Sprintf("tf_standby_orchestrator_name_%d", acctest.RandIntRange(10, 100))
	standbyOrchestratorWorkspaceID := fmt.Sprintf("tf_standby_orchestrator_workspace_id_%d", acctest.RandIntRange(10, 100))
	orchestratorHa := "false"
	resourceInstance := fmt.Sprintf("tf_resource_instance_%d", acctest.RandIntRange(10, 100))
	secretGroup := fmt.Sprintf("tf_secret_group_%d", acctest.RandIntRange(10, 100))
	secret := fmt.Sprintf("tf_secret_%d", acctest.RandIntRange(10, 100))
	regionID := fmt.Sprintf("tf_region_id_%d", acctest.RandIntRange(10, 100))
	guid := fmt.Sprintf("tf_guid_%d", acctest.RandIntRange(10, 100))
	machineType := fmt.Sprintf("tf_machine_type_%d", acctest.RandIntRange(10, 100))
	tier := fmt.Sprintf("tf_tier_%d", acctest.RandIntRange(10, 100))
	standbyTier := fmt.Sprintf("tf_standby_tier_%d", acctest.RandIntRange(10, 100))
	standbyMachineType := fmt.Sprintf("tf_standby_machine_type_%d", acctest.RandIntRange(10, 100))
	tenantName := fmt.Sprintf("tf_tenant_name_%d", acctest.RandIntRange(10, 100))
	proxyIP := fmt.Sprintf("tf_proxy_ip_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPdrManagedrDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMPdrManagedrConfig(instanceID, standByRedeploy, acceptLanguage, acceptsIncomplete, orchestratorLocationType, locationID, sshKeyName, standbySSHKeyName, orchestratorName, orchestratorWorkspaceID, standbyOrchestratorName, standbyOrchestratorWorkspaceID, orchestratorHa, resourceInstance, secretGroup, secret, regionID, guid, machineType, tier, standbyTier, standbyMachineType, tenantName, proxyIP),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMPdrManagedrExists("ibm_pdr_managedr.pdr_managedr_instance", conf),
					resource.TestCheckResourceAttr("ibm_pdr_managedr.pdr_managedr_instance", "instance_id", instanceID),
					resource.TestCheckResourceAttr("ibm_pdr_managedr.pdr_managedr_instance", "stand_by_redeploy", standByRedeploy),
					resource.TestCheckResourceAttr("ibm_pdr_managedr.pdr_managedr_instance", "accept_language", acceptLanguage),
					resource.TestCheckResourceAttr("ibm_pdr_managedr.pdr_managedr_instance", "accepts_incomplete", acceptsIncomplete),
					resource.TestCheckResourceAttr("ibm_pdr_managedr.pdr_managedr_instance", "orchestrator_location_type", orchestratorLocationType),
					resource.TestCheckResourceAttr("ibm_pdr_managedr.pdr_managedr_instance", "location_id", locationID),
					resource.TestCheckResourceAttr("ibm_pdr_managedr.pdr_managedr_instance", "ssh_key_name", sshKeyName),
					resource.TestCheckResourceAttr("ibm_pdr_managedr.pdr_managedr_instance", "standby_ssh_key_name", standbySSHKeyName),
					resource.TestCheckResourceAttr("ibm_pdr_managedr.pdr_managedr_instance", "orchestrator_name", orchestratorName),
					resource.TestCheckResourceAttr("ibm_pdr_managedr.pdr_managedr_instance", "orchestrator_workspace_id", orchestratorWorkspaceID),
					resource.TestCheckResourceAttr("ibm_pdr_managedr.pdr_managedr_instance", "standby_orchestrator_name", standbyOrchestratorName),
					resource.TestCheckResourceAttr("ibm_pdr_managedr.pdr_managedr_instance", "standby_orchestrator_workspace_id", standbyOrchestratorWorkspaceID),
					resource.TestCheckResourceAttr("ibm_pdr_managedr.pdr_managedr_instance", "orchestrator_ha", orchestratorHa),
					resource.TestCheckResourceAttr("ibm_pdr_managedr.pdr_managedr_instance", "resource_instance", resourceInstance),
					resource.TestCheckResourceAttr("ibm_pdr_managedr.pdr_managedr_instance", "secret_group", secretGroup),
					resource.TestCheckResourceAttr("ibm_pdr_managedr.pdr_managedr_instance", "secret", secret),
					resource.TestCheckResourceAttr("ibm_pdr_managedr.pdr_managedr_instance", "region_id", regionID),
					resource.TestCheckResourceAttr("ibm_pdr_managedr.pdr_managedr_instance", "guid", guid),
					resource.TestCheckResourceAttr("ibm_pdr_managedr.pdr_managedr_instance", "machine_type", machineType),
					resource.TestCheckResourceAttr("ibm_pdr_managedr.pdr_managedr_instance", "tier", tier),
					resource.TestCheckResourceAttr("ibm_pdr_managedr.pdr_managedr_instance", "standby_tier", standbyTier),
					resource.TestCheckResourceAttr("ibm_pdr_managedr.pdr_managedr_instance", "standby_machine_type", standbyMachineType),
					resource.TestCheckResourceAttr("ibm_pdr_managedr.pdr_managedr_instance", "tenant_name", tenantName),
					resource.TestCheckResourceAttr("ibm_pdr_managedr.pdr_managedr_instance", "proxy_ip", proxyIP),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_pdr_managedr.pdr_managedr_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMPdrManagedrConfigBasic(instanceID string) string {
	return fmt.Sprintf(`
		resource "ibm_pdr_managedr" "pdr_managedr_instance" {
			instance_id = "%s"
		}
	`, instanceID)
}

func testAccCheckIBMPdrManagedrConfig(instanceID string, standByRedeploy string, acceptLanguage string, acceptsIncomplete string, orchestratorLocationType string, locationID string, sshKeyName string, standbySSHKeyName string, orchestratorName string, orchestratorWorkspaceID string, standbyOrchestratorName string, standbyOrchestratorWorkspaceID string, orchestratorHa string, resourceInstance string, secretGroup string, secret string, regionID string, guid string, machineType string, tier string, standbyTier string, standbyMachineType string, tenantName string, proxyIP string) string {
	return fmt.Sprintf(`

		resource "ibm_pdr_managedr" "pdr_managedr_instance" {
			instance_id = "%s"
			stand_by_redeploy = "%s"
			accept_language = "%s"
			accepts_incomplete = %s
			orchestrator_location_type = "%s"
			location_id = "%s"
			ssh_key_name = "%s"
			standby_ssh_key_name = "%s"
			orchestrator_name = "%s"
			orchestrator_workspace_id = "%s"
			standby_orchestrator_name = "%s"
			standby_orchestrator_workspace_id = "%s"
			orchestrator_ha = %s
			resource_instance = "%s"
			secret_group = "%s"
			secret = "%s"
			region_id = "%s"
			guid = "%s"
			machine_type = "%s"
			tier = "%s"
			standby_tier = "%s"
			standby_machine_type = "%s"
			tenant_name = "%s"
			proxy_ip = "%s"
		}
	`, instanceID, standByRedeploy, acceptLanguage, acceptsIncomplete, orchestratorLocationType, locationID, sshKeyName, standbySSHKeyName, orchestratorName, orchestratorWorkspaceID, standbyOrchestratorName, standbyOrchestratorWorkspaceID, orchestratorHa, resourceInstance, secretGroup, secret, regionID, guid, machineType, tier, standbyTier, standbyMachineType, tenantName, proxyIP)
}

func testAccCheckIBMPdrManagedrExists(n string, obj drautomationservicev1.ServiceInstanceManageDr) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		drAutomationServiceClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).DrAutomationServiceV1()
		if err != nil {
			return err
		}

		getManageDrOptions := &drautomationservicev1.GetManageDrOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getManageDrOptions.SetInstanceID(parts[0])
		getManageDrOptions.SetInstanceID(parts[1])

		serviceInstanceManageDr, _, err := drAutomationServiceClient.GetManageDr(getManageDrOptions)
		if err != nil {
			return err
		}

		obj = *serviceInstanceManageDr
		return nil
	}
}

func testAccCheckIBMPdrManagedrDestroy(s *terraform.State) error {
	drAutomationServiceClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).DrAutomationServiceV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_pdr_managedr" {
			continue
		}

		getManageDrOptions := &drautomationservicev1.GetManageDrOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getManageDrOptions.SetInstanceID(parts[0])
		getManageDrOptions.SetInstanceID(parts[1])

		// Try to find the key
		_, response, err := drAutomationServiceClient.GetManageDr(getManageDrOptions)

		if err == nil {
			return fmt.Errorf("pdr_managedr still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for pdr_managedr (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func TestResourceIBMPdrManagedrMapToPrimaryOrchCaSecrets(t *testing.T) {
	checkResult := func(result *drautomationservicev1.PrimaryOrchCaSecrets) {
		model := new(drautomationservicev1.PrimaryOrchCaSecrets)
		model.CaCertificateSecretID = core.StringPtr("12345678-1234-1234-1234-123456789012")
		model.CaSecretManagerGUID = core.StringPtr("12345678-1234-1234-1234-123456789012")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["ca_certificate_secret_id"] = "12345678-1234-1234-1234-123456789012"
	model["ca_secret_manager_guid"] = "12345678-1234-1234-1234-123456789012"

	result, err := drautomationservice.ResourceIBMPdrManagedrMapToPrimaryOrchCaSecrets(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMPdrManagedrMapToStandbyOrchCaSecrets(t *testing.T) {
	checkResult := func(result *drautomationservicev1.StandbyOrchCaSecrets) {
		model := new(drautomationservicev1.StandbyOrchCaSecrets)
		model.CaCertificateSecretID = core.StringPtr("12345678-1234-1234-1234-123456789012")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["ca_certificate_secret_id"] = "12345678-1234-1234-1234-123456789012"

	result, err := drautomationservice.ResourceIBMPdrManagedrMapToStandbyOrchCaSecrets(model)
	assert.Nil(t, err)
	checkResult(result)
}
