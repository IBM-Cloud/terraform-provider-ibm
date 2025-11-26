// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power_test

import (
	"context"
	"errors"
	"fmt"
	"log"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccIBMPIVolumeAttachbasic(t *testing.T) {
	name := fmt.Sprintf("tf-pi-volume-attach-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPIVolumeAttachDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIVolumeAttachConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPIVolumeAttachExists("ibm_pi_volume_attach.power_attach_volume"),
					resource.TestCheckResourceAttrSet("ibm_pi_volume_attach.power_attach_volume", "id"),
					resource.TestCheckResourceAttrSet("ibm_pi_volume_attach.power_attach_volume", "status"),
				),
			},
		},
	})
}

func TestAccIBMPIShareableVolumeAttachbasic(t *testing.T) {
	name := fmt.Sprintf("tf-pi-shareable-volume-attach-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPIVolumeAttachDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIShareableVolumeAttachConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPIVolumeAttachExists("ibm_pi_volume_attach.power_attach_volume"),
					resource.TestCheckResourceAttrSet("ibm_pi_volume_attach.power_attach_volume", "id"),
					resource.TestCheckResourceAttrSet("ibm_pi_volume_attach.power_attach_volume", "status"),
				),
			},
		},
	})
}

func testAccCheckIBMPIVolumeAttachDestroy(s *terraform.State) error {
	sess, err := acc.TestAccProvider.Meta().(conns.ClientSession).IBMPISession()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_pi_volume_attach" {
			continue
		}

		ids, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}
		cloudInstanceID, pvmInstanceID, volumeID := ids[0], ids[1], ids[2]
		client := instance.NewIBMPIVolumeClient(context.Background(), sess, cloudInstanceID)
		volumeAttach, err := client.CheckVolumeAttach(pvmInstanceID, volumeID)
		if err == nil {
			log.Println("volume attach*****", volumeAttach.State)
			return fmt.Errorf("PI Volume Attach still exists: %s", rs.Primary.ID)
		}
	}

	return nil
}

func testAccCheckIBMPIVolumeAttachExists(n string) resource.TestCheckFunc {
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

		ids, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}
		cloudInstanceID, pvmInstanceID, volumeID := ids[0], ids[1], ids[2]
		client := instance.NewIBMPIVolumeClient(context.Background(), sess, cloudInstanceID)

		_, err = client.CheckVolumeAttach(pvmInstanceID, volumeID)
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccCheckIBMPIVolumeAttachConfig(name string) string {
	return fmt.Sprintf(`
		resource "ibm_pi_volume" "power_volume" {
			pi_volume_size       = 2
			pi_volume_name       = "%[2]s"
			pi_volume_shareable  = true
			pi_volume_pool       = "%[5]s"
			pi_cloud_instance_id = "%[1]s"
		}
		resource "ibm_pi_instance" "power_instance" {
			pi_memory             = "2"
			pi_processors         = "0.25"
			pi_instance_name      = "%[2]s"
			pi_proc_type          = "shared"
			pi_image_id           = "%[3]s"
			pi_sys_type           = "s922"
			pi_cloud_instance_id  = "%[1]s"
			pi_storage_pool       = "%[5]s"
			pi_network {
				network_id = "%[4]s"
			}
		}
		resource "ibm_pi_volume_attach" "power_attach_volume"{
			pi_cloud_instance_id 	= "%[1]s"
			pi_volume_id			= ibm_pi_volume.power_volume.volume_id
			pi_instance_id 			= ibm_pi_instance.power_instance.instance_id
		}
	`, acc.Pi_cloud_instance_id, name, acc.Pi_image, acc.Pi_network_name, acc.PiStoragePool)
}

func testAccCheckIBMPIShareableVolumeAttachConfig(name string) string {
	return fmt.Sprintf(`
		resource "ibm_pi_volume" "power_volume" {
			pi_volume_size       = 2
			pi_volume_name       = "%[2]s"
			pi_volume_shareable  = true
			pi_volume_pool       = "%[5]s"
			pi_cloud_instance_id = "%[1]s"
		}
		resource "ibm_pi_instance" "power_instance" {
			count                 = 2
			pi_memory             = "2"
			pi_processors         = "0.25"
			pi_instance_name      = "%[2]s-${count.index}"
			pi_proc_type          = "shared"
			pi_image_id           = "%[3]s"
			pi_sys_type           = "s922"
			pi_cloud_instance_id  = "%[1]s"
			pi_storage_pool       = "%[5]s"
			pi_volume_ids         =  count.index == 0 ? [ibm_pi_volume.power_volume.volume_id] : null
			pi_network {
				network_id = "%[4]s"
			}
		}
		resource "ibm_pi_volume_attach" "power_attach_volume"{
			pi_cloud_instance_id 	= "%[1]s"
			pi_volume_id 			= ibm_pi_volume.power_volume.volume_id
			pi_instance_id 			= ibm_pi_instance.power_instance[1].instance_id
		}
	`, acc.Pi_cloud_instance_id, name, acc.Pi_image, acc.Pi_network_name, acc.PiStoragePool)
}
