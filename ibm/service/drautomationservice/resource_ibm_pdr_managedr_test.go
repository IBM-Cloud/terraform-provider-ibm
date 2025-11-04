// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package drautomationservice_test

import (
	"fmt"
	"testing"


	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"

	"github.com/IBM/dra-go-sdk/drautomationservicev1"
	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMPdrManagedrBasic(t *testing.T) {
	var conf drautomationservicev1.ServiceInstanceManageDr
	instanceID := "xxxx2ec4-xxxx-4f84-xxxx-c2aa834dd4ed"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
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
	instanceID := "xxxxxfe5-fba1-4cb3-xxxx-e1b09fa0df26"
	acceptLanguage := "it"
	standByRedeploy := "false"
	acceptsIncomplete := "true"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMPdrManagedrConfig(instanceID, standByRedeploy, acceptLanguage, acceptsIncomplete),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMPdrManagedrExists("ibm_pdr_managedr.pdr_managedr_instance", conf),
					resource.TestCheckResourceAttr("ibm_pdr_managedr.pdr_managedr_instance", "instance_id", instanceID),
					resource.TestCheckResourceAttr("ibm_pdr_managedr.pdr_managedr_instance", "stand_by_redeploy", standByRedeploy),
					resource.TestCheckResourceAttr("ibm_pdr_managedr.pdr_managedr_instance", "accept_language", acceptLanguage),
					resource.TestCheckResourceAttr("ibm_pdr_managedr.pdr_managedr_instance", "accepts_incomplete", acceptsIncomplete),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_pdr_managedr.pdr_managedr",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMPdrManagedrConfigBasic(instanceID string) string {
	apiKey := acc.DRApiKey
	return fmt.Sprintf(`
		resource "ibm_pdr_managedr" "pdr_managedr_instance" {
			instance_id = "%s"
			orchestrator_ha             = "false"
			orchestrator_location_type  = "off-premises"
			location_id                 = "mad04"
			orchestrator_workspace_id   = "dbc77621-e7cb-4d70-b662-776ff2c3317f"
			orchestrator_name           = "terraform_orch_vm_04"
			orchestrator_password       = "Password1234567"
			machine_type                = "e1080"
			tier                        = "tier1"
			ssh_key_name                = "vijaykey"
			action                      = "done"
			api_key                     = "%s"
		}
	`, instanceID, apiKey)
}

func testAccCheckIBMPdrManagedrConfig(instanceID string, standByRedeploy string, acceptLanguage string, acceptsIncomplete string) string {
	apiKey := acc.DRApiKey
	return fmt.Sprintf(`

		resource "ibm_pdr_managedr" "pdr_managedr_instance" {
			instance_id = "%s"
			accept_language = "%s"
			accepts_incomplete = %s
			orchestrator_ha             = "false"
			orchestrator_location_type  = "off-premises"
			location_id                 = "mad04"
			orchestrator_workspace_id   = "dbc77621-e7cb-4d70-b662-776ff2c3317f"
			orchestrator_name           = "terraform_orch_vm_04"
			orchestrator_password       = "Password1234567"
			machine_type                = "e1080"
			tier                        = "tier1"
			ssh_key_name                = "vijaykey"
			action                      = "done"
			api_key                     = "%s"
		}
	`, instanceID, acceptLanguage, acceptsIncomplete, apiKey)
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
