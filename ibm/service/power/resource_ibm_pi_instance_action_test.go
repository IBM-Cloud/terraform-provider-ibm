// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power_test

import (
	"fmt"
	"strings"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/service/power"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMPIInstanceAction(t *testing.T) {
	name := fmt.Sprintf("tf-pi-instance-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIInstanceActionConfig(name, power.Action_Stop),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_pi_instance_action.example", "status", strings.ToUpper(power.State_Shutoff)),
				),
			},
			{
				// Try to stop already stopped instance
				Config: testAccCheckIBMPIInstanceActionConfig(name, power.Action_Stop),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_pi_instance_action.example", "status", strings.ToUpper(power.State_Shutoff)),
				),
			},
			{
				Config: testAccCheckIBMPIInstanceActionConfig(name, power.Action_Start),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_pi_instance_action.example", "status", strings.ToUpper(power.State_Active)),
				),
			},
			{
				// Try to start already started instance
				Config: testAccCheckIBMPIInstanceActionConfig(name, power.Action_Start),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_pi_instance_action.example", "status", strings.ToUpper(power.State_Active)),
				),
			},
		},
	})
}

func TestAccIBMPIInstanceActionHardReboot(t *testing.T) {
	name := fmt.Sprintf("tf-pi-instance-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIInstanceActionWithHealthStatusConfig(name, power.Action_HardReboot, power.Warning),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_pi_instance_action.example", "status", strings.ToUpper(power.State_Active)),
					resource.TestCheckResourceAttr(
						"ibm_pi_instance_action.example", "health_status", power.Warning),
				),
			},
		},
	})
}

func TestAccIBMPIInstanceActionResetState(t *testing.T) {
	name := fmt.Sprintf("tf-pi-instance-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIInstanceActionWithHealthStatusConfig(name, power.Action_ResetState, power.Warning),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_pi_instance_action.example", "status", strings.ToUpper(power.State_Active)),
					resource.TestCheckResourceAttr(
						"ibm_pi_instance_action.example", "health_status", power.Critical),
				),
			},
		},
	})
}

func testAccCheckIBMPIInstanceActionConfig(name, action string) string {
	return fmt.Sprintf(`
	data "ibm_pi_image" "power_image" {
		pi_cloud_instance_id = "%[1]s"
		pi_image_name        = "%[4]s"
	  }
	  data "ibm_pi_network" "power_networks" {
		pi_cloud_instance_id = "%[1]s"
		pi_network_name      = "%[5]s"
	  }
	  resource "ibm_pi_instance" "power_instance" {
		pi_cloud_instance_id  = "%[1]s"
		pi_image_id           = data.ibm_pi_image.power_image.id
		pi_instance_name      = "%[2]s"
		pi_memory             = "2"
		pi_proc_type          = "shared"
		pi_processors         = "0.25"
		pi_storage_pool       = data.ibm_pi_image.power_image.storage_pool
		pi_storage_type       = "%[6]s"
		pi_sys_type           = "s922"
		pi_network {
			network_id = data.ibm_pi_network.power_networks.id
		}
	  }

	resource "ibm_pi_instance_action" "example" {
		pi_action				= "%[3]s"
		pi_cloud_instance_id	= "%[1]s"
		pi_instance_id			= resource.ibm_pi_instance.power_instance.pi_instance_name 
	}
	`, acc.Pi_cloud_instance_id, name, action, acc.Pi_image, acc.Pi_network_name, acc.PiStorageType)
}

func testAccCheckIBMPIInstanceActionWithHealthStatusConfig(name, action, instanceHealthStatus string) string {
	return fmt.Sprintf(`
	data "ibm_pi_image" "power_image" {
		pi_cloud_instance_id = "%[1]s"
		pi_image_name        = "%[4]s"
	  }
	  data "ibm_pi_network" "power_networks" {
		pi_cloud_instance_id = "%[1]s"
		pi_network_name      = "%[5]s"
	  }
	  resource "ibm_pi_instance" "power_instance" {
		pi_cloud_instance_id  = "%[1]s"
		pi_health_status      = "%[7]s"
		pi_image_id           = data.ibm_pi_image.power_image.id
		pi_instance_name      = "%[2]s"
		pi_memory             = "2"
		pi_proc_type          = "shared"
		pi_processors         = "0.25"
		pi_storage_pool       = data.ibm_pi_image.power_image.storage_pool
		pi_storage_type       = "%[6]s"
		pi_sys_type           = "s922"
		pi_network {
			network_id = data.ibm_pi_network.power_networks.id
		}
	  }
	resource "ibm_pi_instance_action" "example" {
		pi_action				= "%[3]s"
		pi_cloud_instance_id	= "%[1]s"
		pi_health_status		= "%[7]s"
		pi_instance_id			= resource.ibm_pi_instance.power_instance.pi_instance_name 
	}
	`, acc.Pi_cloud_instance_id, name, action, acc.Pi_image, acc.Pi_network_name, acc.PiStorageType, instanceHealthStatus)
}
