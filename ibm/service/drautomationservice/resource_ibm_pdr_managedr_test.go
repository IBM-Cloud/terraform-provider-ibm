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

func TestAccIBMPdrManagedrBasic(t *testing.T) {
	var conf drautomationservicev1.ServiceInstanceManageDr
	instanceID := fmt.Sprintf("tf_instance_id_%d", acctest.RandIntRange(10, 100))

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
	instanceID := fmt.Sprintf("tf_instance_id_%d", acctest.RandIntRange(10, 100))
	standByRedeploy := fmt.Sprintf("tf_stand_by_redeploy_%d", acctest.RandIntRange(10, 100))
	acceptLanguage := fmt.Sprintf("tf_accept_language_%d", acctest.RandIntRange(10, 100))
	ifNoneMatch := fmt.Sprintf("tf_if_none_match_%d", acctest.RandIntRange(10, 100))
	acceptsIncomplete := "true"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMPdrManagedrConfig(instanceID, standByRedeploy, acceptLanguage, ifNoneMatch, acceptsIncomplete),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMPdrManagedrExists("ibm_pdr_managedr.pdr_managedr_instance", conf),
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

func testAccCheckIBMPdrManagedrConfigBasic(instanceID string) string {
	return fmt.Sprintf(`
		resource "ibm_pdr_managedr" "pdr_managedr_instance" {
			instance_id = "%s"
		}
	`, instanceID)
}

func testAccCheckIBMPdrManagedrConfig(instanceID string, standByRedeploy string, acceptLanguage string, ifNoneMatch string, acceptsIncomplete string) string {
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
