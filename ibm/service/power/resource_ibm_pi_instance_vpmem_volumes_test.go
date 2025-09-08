// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/power"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/IBM-Cloud/power-go-client/clients/instance"
)

func TestAccIBMPIInstancevpmemVolumesBasic(t *testing.T) {
	name := fmt.Sprintf("tf-pvm-vpmem-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPIInstancevpmemVolumesDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIInstancevpmemVolumesConfigBasic(name, power.OK, power.Action_Stop),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMPIInstancevpmemVolumesExists("ibm_pi_instance_vpmem_volumes.vpmem_volumes_instance"),
					resource.TestCheckResourceAttr("ibm_pi_instance_vpmem_volumes.vpmem_volumes_instance", "volumes.#", "1"),
				),
			},
		},
	})
}

func testAccCheckIBMPIInstancevpmemVolumesConfigBasic(name, instanceHealthStatus, action string) string {
	return fmt.Sprintf(`
		
 	data "ibm_pi_image" "power_image" {
		pi_cloud_instance_id = "%[1]s"
		pi_image_name        = "%[3]s"
	}
	data "ibm_pi_network" "power_networks" {
		pi_cloud_instance_id = "%[1]s"
		pi_network_name      = "%[4]s"
	}
	resource "ibm_pi_volume" "power_volume" {
		pi_cloud_instance_id = "%[1]s"
		pi_volume_name       = "%[2]s-vol"
		pi_volume_pool       = data.ibm_pi_image.power_image.storage_pool
		pi_volume_shareable  = true
		pi_volume_size       = 20
		pi_volume_type       = "%[6]s"
	}
	resource "ibm_pi_instance" "power_instance" {
		pi_cloud_instance_id  = "%[1]s"
		pi_health_status      = "%[5]s"
		pi_image_id           = data.ibm_pi_image.power_image.id
		pi_instance_name      = "%[2]s"
		pi_memory             = "2"
		pi_proc_type          = "shared"
		pi_processors         = "0.25"
		pi_storage_pool       = data.ibm_pi_image.power_image.storage_pool
		pi_storage_type       = "%[6]s"
		pi_sys_type           = "s1022"
		pi_volume_ids         = [ibm_pi_volume.power_volume.volume_id]
		pi_network {
			network_id = data.ibm_pi_network.power_networks.id
		}
	}
	resource "ibm_pi_instance_action" "vm" {
		pi_action				= "%[7]s"
		pi_cloud_instance_id	= "%[1]s"
		pi_instance_id			= ibm_pi_instance.power_instance.instance_id 
	}	
	resource "ibm_pi_instance_vpmem_volumes" "vpmem_volumes" {
		pi_cloud_instance_id = "%[1]s"
		pi_pvm_instance_id = ibm_pi_instance_action.vm.pi_instance_id
		pi_vpmem_volumes {
			name = "%[2]s-1"
			size = 1
		}
	}
	`, acc.Pi_cloud_instance_id, name, acc.Pi_image, acc.Pi_network_name, instanceHealthStatus, acc.PiStorageType, action)
}

func testAccCheckIBMPIInstancevpmemVolumesExists(n string) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}
		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}
		sess, err := acc.TestAccProvider.Meta().(conns.ClientSession).IBMPISession()
		if err != nil {
			return err
		}

		idArr, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}
		cloudInstanceID := idArr[0]
		pvmInstanceID := idArr[1]
		for _, vpmemVolumeID := range idArr[2:] {
			client := instance.NewIBMPIVPMEMClient(context.Background(), sess, cloudInstanceID)
			_, err := client.GetPvmVpmemVolume(pvmInstanceID, vpmemVolumeID)
			if err == nil {
				return err
			}
		}

		return nil
	}
}

func testAccCheckIBMPIInstancevpmemVolumesDestroy(s *terraform.State) error {
	sess, err := acc.TestAccProvider.Meta().(conns.ClientSession).IBMPISession()
	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_pi_instance_vpmem_volumes" {
			continue
		}
		idArr, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}

		cloudInstanceID := idArr[0]
		pvmInstanceID := idArr[1]
		for _, vpmemVolumeID := range idArr[2:] {
			client := instance.NewIBMPIVPMEMClient(context.Background(), sess, cloudInstanceID)
			_, err := client.GetPvmVpmemVolume(pvmInstanceID, vpmemVolumeID)
			if err == nil {
				return fmt.Errorf("vpmemVolumeID still exists: %s", vpmemVolumeID)
			}
		}

	}

	return nil
}
