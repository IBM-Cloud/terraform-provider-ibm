// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/power"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMPIInstanceDataSource_basic(t *testing.T) {
	instanceResData := "data.ibm_pi_instance.testacc_ds_instance"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIInstanceDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(instanceResData, "id"),
				),
			},
		},
	})
}

func testAccCheckIBMPIInstanceDataSourceConfig() string {
	return fmt.Sprintf(`
		data "ibm_pi_instance" "testacc_ds_instance" {
			pi_cloud_instance_id = "%[1]s"
			pi_instance_id       = "%[2]s"
		}`, acc.Pi_cloud_instance_id, acc.Pi_instance_id)
}

func TestAccIBMPIInstanceDataSource_IBMiPHAFSM(t *testing.T) {
	instanceRes := "ibm_pi_instance.power_instance"
	dsRes := "data.ibm_pi_instance.ds_instance_fsm"
	name := fmt.Sprintf("tf-pi-ds-ibmi-pha-fsm-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				// Create an IBM i instance with FSM enabled via count=3
				Config: testAccCheckIBMPIInstanceDataSourcePHAFSMConfig(name, power.OK, 3),
				Check: resource.ComposeTestCheckFunc(
					// Ensure the resource exists
					resource.TestCheckResourceAttr(instanceRes, "pi_instance_name", name),
					// Data source reflects computed boolean and computed count
					resource.TestCheckResourceAttrSet(dsRes, "id"),
					resource.TestCheckResourceAttr(dsRes, "ibmi_pha_fsm", "true"),
					resource.TestCheckResourceAttr(dsRes, "ibmi_pha_fsm_count", "3"),
				),
			},
		},
	})
}

func testAccCheckIBMPIInstanceDataSourcePHAFSMConfig(name, instanceHealthStatus string, fsmCount int) string {
	return fmt.Sprintf(`
      data "ibm_pi_image" "power_image" {
        pi_cloud_instance_id = "%[1]s"
        pi_image_name        = "%[3]s"
      }
      data "ibm_pi_network" "power_networks" {
        pi_cloud_instance_id = "%[1]s"
        pi_network_name      = "%[4]s"
      }
      resource "ibm_pi_instance" "power_instance" {
        pi_cloud_instance_id  = "%[1]s"
        pi_health_status      = "OK"
        pi_image_id           = data.ibm_pi_image.power_image.id
        pi_instance_name      = "%[2]s"
        pi_memory             = "2"
        pi_proc_type          = "shared"
        pi_processors         = "0.25"
        pi_storage_type       = "%[6]s"
        pi_sys_type           = "s922"

        # Set IBMi PHA FSM License
        pi_ibmi_pha_fsm_count = %[7]d

        pi_network {
            network_id = data.ibm_pi_network.power_networks.id
        }
      }

      data "ibm_pi_instance" "ds_instance_fsm" {
        pi_cloud_instance_id = "%[1]s"
        pi_instance_id       = ibm_pi_instance.power_instance.instance_id
      }
    `, acc.Pi_cloud_instance_id, name, acc.Pi_image, acc.Pi_network_name, instanceHealthStatus, acc.PiStorageType, fsmCount)
}
