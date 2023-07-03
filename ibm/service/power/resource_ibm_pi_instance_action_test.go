// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMPIInstanceAction(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIInstanceActionConfig("stop"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_pi_instance_action.example", "status", "SHUTOFF"),
				),
			},
			{
				// Try to stop already stopped instance
				Config: testAccCheckIBMPIInstanceActionConfig("stop"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_pi_instance_action.example", "status", "SHUTOFF"),
				),
			},
			{
				Config: testAccCheckIBMPIInstanceActionConfig("start"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_pi_instance_action.example", "status", "ACTIVE"),
				),
			},
			{
				// Try to start already started instance
				Config: testAccCheckIBMPIInstanceActionConfig("start"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_pi_instance_action.example", "status", "ACTIVE"),
				),
			},
		},
	})
}

func TestAccIBMPIInstanceActionHardReboot(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIInstanceActionWithHealthStatusConfig("hard-reboot"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_pi_instance_action.example", "status", "ACTIVE"),
					resource.TestCheckResourceAttr(
						"ibm_pi_instance_action.example", "health_status", "WARNING"),
				),
			},
		},
	})
}

func TestAccIBMPIInstanceActionResetState(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIInstanceActionWithHealthStatusConfig("reset-state"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_pi_instance_action.example", "status", "ACTIVE"),
					resource.TestCheckResourceAttr(
						"ibm_pi_instance_action.example", "health_status", "CRITICAL"),
				),
			},
		},
	})
}

func testAccCheckIBMPIInstanceActionConfig(action string) string {
	return fmt.Sprintf(`
	resource "ibm_pi_instance_action" "example" {
		pi_cloud_instance_id	= "%s"
		pi_instance_id			= "%s"
		pi_action				= "%s"
	}
	`, acc.Pi_cloud_instance_id, acc.Pi_instance_name, action)
}

func testAccCheckIBMPIInstanceActionWithHealthStatusConfig(action string) string {
	return fmt.Sprintf(`
	resource "ibm_pi_instance_action" "example" {
		pi_cloud_instance_id	= "%s"
		pi_instance_id			= "%s"
		pi_action				= "%s"
		pi_health_status		= "WARNING"
	}
	`, acc.Pi_cloud_instance_id, acc.Pi_instance_name, action)
}
