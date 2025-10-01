// Copyright IBM Corp. 2025 All Rights Reserved.
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
	"github.ibm.com/DRAutomation/dra-go-sdk/drautomationservicev1"
)

func TestAccIbmPdrManagedrBasic(t *testing.T) {
	var conf drautomationservicev1.ServiceInstanceManageDR
	instanceID := "crn:v1:staging:public:power-dr-automation:global:a/a09202c1bfb04ceebfb4a9fd38c87721:050ebe3b-13f4-4db8-8ece-501a3c13be0965822::"
	standByRedeploy := "false"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmPdrManagedrDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmPdrManagedrConfigBasic(instanceID, standByRedeploy),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmPdrManagedrExists("ibm_pdr_managedr.pdr_managedr_instance", conf),
					resource.TestCheckResourceAttr("ibm_pdr_managedr.pdr_managedr_instance", "instance_id", instanceID),
					resource.TestCheckResourceAttr("ibm_pdr_managedr.pdr_managedr_instance", "stand_by_redeploy", standByRedeploy),
				),
			},
		},
	})
}

func TestAccIbmPdrManagedrAllArgs(t *testing.T) {
	var conf drautomationservicev1.ServiceInstanceManageDR
	instanceID := "crn:v1:staging:public:power-dr-automation:global:a/a09202c1bfb04ceebfb4a9fd38c87721:050ebe3b-13f4-4db8-8ece-501a3c13be0965822::"
	standByRedeploy := "false"
	acceptLanguage := fmt.Sprintf("tf_accept_language_%d", acctest.RandIntRange(10, 100))
	ifNoneMatch := fmt.Sprintf("tf_if_none_match_%d", acctest.RandIntRange(10, 100))
	acceptsIncomplete := "true"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmPdrManagedrDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmPdrManagedrConfig(instanceID, standByRedeploy, acceptLanguage, ifNoneMatch, acceptsIncomplete),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmPdrManagedrExists("ibm_pdr_managedr.pdr_managedr_instance", conf),
					resource.TestCheckResourceAttr("ibm_pdr_managedr.pdr_managedr_instance", "instance_id", instanceID),
					resource.TestCheckResourceAttr("ibm_pdr_managedr.pdr_managedr_instance", "stand_by_redeploy", standByRedeploy),
					resource.TestCheckResourceAttr("ibm_pdr_managedr.pdr_managedr_instance", "accept_language", acceptLanguage),
					resource.TestCheckResourceAttr("ibm_pdr_managedr.pdr_managedr_instance", "if_none_match", ifNoneMatch),
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

func testAccCheckIbmPdrManagedrConfigBasic(instanceID string, standByRedeploy string) string {
	return fmt.Sprintf(`
		resource "ibm_pdr_managedr" "pdr_managedr_instance" {
			# Path / query parameters
			instance_id        = "%s"
			stand_by_redeploy  = "%s"

			# Request body parameters (from ServiceInstanceManageDRRequest)
			action                          = ""
			api_key                         = ""
			enable_ha                       = false
			guid                            = ""
			location_id                     = "dal10"
			machine_type                    = "s922"
			orchestrator_cluster_type       = "off-premises"
			orchestrator_name               = "vindhya_NONHAtest1232"
			orchestrator_password           = "vindhyaSri@123hyuuu"
			orchestrator_workspace_id       = "75cbf05b-78f6-406e-afe7-a904f646d798"
			ssh_key_name					= "vijaykey"
			orchestrator_workspace_location = ""
			proxy_ip                        = "10.30.40.4:3128"
			region_id                       = ""
			resource_instance               = "crn:v1:bluemix:public:resource-controller::res123"
			schematic_workspace_id          = ""
			secondary_workspace_id          = ""
			secret                          = ""
			secret_group                    = ""
			standby_machine_type            = "s922"
			standby_orchestrator_name       = "drautomationstandby"
			standby_orchestrator_workspace_id = "71027b79-0e31-44f6-a499-63eca1a66feb"
			standby_orchestrator_workspace_location = ""
			standby_schematic_workspace_id  = ""
			standby_tier                    = "tier1"
			tier                            = "tier1"
			transit_gateway_id              = "024fcff9-c676-46e4-ad42-3b2d349c9f8f"
			vpc_id                          = "r006-2f3b3ab9-2149-49cc-83a1-30a5d93d59b2"
		}
	`, instanceID, standByRedeploy)
}

func testAccCheckIbmPdrManagedrConfig(instanceID string, standByRedeploy string, acceptLanguage string, ifNoneMatch string, acceptsIncomplete string) string {
	return fmt.Sprintf(`

		resource "ibm_pdr_managedr" "pdr_managedr_instance" {
			instance_id = "%s"
			stand_by_redeploy = "%s"
			accept_language = "%s"
			if_none_match = "%s"
			accepts_incomplete = %s
		}
	`, instanceID, standByRedeploy, acceptLanguage, ifNoneMatch, acceptsIncomplete)
}

func testAccCheckIbmPdrManagedrExists(n string, obj drautomationservicev1.ServiceInstanceManageDR) resource.TestCheckFunc {

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

		serviceInstanceManageDR, _, err := drAutomationServiceClient.GetManageDr(getManageDrOptions)
		if err != nil {
			return err
		}

		obj = *serviceInstanceManageDR
		return nil
	}
}

func testAccCheckIbmPdrManagedrDestroy(s *terraform.State) error {
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

		getManageDrOptions.SetInstanceID(fmt.Sprintf("%s/%s", parts[0], parts[1]))

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
