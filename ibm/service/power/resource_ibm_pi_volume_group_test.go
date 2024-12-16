// Copyright IBM Corp. 2022 All Rights Reserved.
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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMPIVolumeGroupUpdate(t *testing.T) {
	name := fmt.Sprintf("tf-pi-volume-group-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPIVolumeGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIVolumeGroupConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPIVolumeGroupExists("ibm_pi_volume_group.power_volume_group"),
					resource.TestCheckResourceAttr(
						"ibm_pi_volume_group.power_volume_group", "pi_volume_group_name", name),
				),
			},
			{
				Config: testAccCheckIBMPIVolumeGroupUpdateConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPIVolumeGroupExists("ibm_pi_volume_group.power_volume_group"),
					resource.TestCheckResourceAttr(
						"ibm_pi_volume_group.power_volume_group", "pi_volume_group_name", name),
				),
			},
			{
				Config: testAccCheckIBMPIVolumeGroupEmptyVolumeConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPIVolumeGroupExists("ibm_pi_volume_group.power_volume_group"),
					resource.TestCheckResourceAttr(
						"ibm_pi_volume_group.power_volume_group", "pi_volume_group_name", name),
				),
			},
		},
	})
}

func testAccCheckIBMPIVolumeGroupDestroy(s *terraform.State) error {
	sess, err := acc.TestAccProvider.Meta().(conns.ClientSession).IBMPISession()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_pi_volume_group" {
			continue
		}
		cloudInstanceID, vgID, err := splitID(rs.Primary.ID)
		if err != nil {
			return err
		}
		vgC := instance.NewIBMPIVolumeGroupClient(context.Background(), sess, cloudInstanceID)
		vg, err := vgC.Get(vgID)
		if err == nil {
			log.Println("volume-group*****", vg.Status)
			return fmt.Errorf("PI Volume Group still exists: %s", rs.Primary.ID)
		}
	}

	return nil
}

func testAccCheckIBMPIVolumeGroupExists(n string) resource.TestCheckFunc {
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
		cloudInstanceID, vgID, err := splitID(rs.Primary.ID)
		if err != nil {
			return err
		}
		client := instance.NewIBMPIVolumeGroupClient(context.Background(), sess, cloudInstanceID)

		_, err = client.Get(vgID)
		if err != nil {
			return err
		}
		return nil

	}
}

func testAccCheckIBMPIVolumeGroupConfig(name string) string {
	return volumeConfig(name, acc.Pi_cloud_instance_id) + fmt.Sprintf(`
		resource "ibm_pi_volume_group" "power_volume_group" {
			pi_volume_group_name       = "%[1]s"
			pi_cloud_instance_id 	   = "%[2]s"
			pi_volume_ids              = [ibm_pi_volume.power_volume[0].volume_id,ibm_pi_volume.power_volume[1].volume_id]
		}
	`, name, acc.Pi_cloud_instance_id)
}

func testAccCheckIBMPIVolumeGroupUpdateConfig(name string) string {
	return volumeConfig(name, acc.Pi_cloud_instance_id) + fmt.Sprintf(`
		resource "ibm_pi_volume_group" "power_volume_group" {
			pi_volume_group_name       = "%[1]s"
			pi_cloud_instance_id 	   = "%[2]s"
			pi_volume_ids              = [ibm_pi_volume.power_volume[2].volume_id]
		}
	`, name, acc.Pi_cloud_instance_id)
}

func testAccCheckIBMPIVolumeGroupEmptyVolumeConfig(name string) string {
	return volumeConfig(name, acc.Pi_cloud_instance_id) + fmt.Sprintf(`
		resource "ibm_pi_volume_group" "power_volume_group" {
			pi_volume_group_name       = "%[1]s"
			pi_cloud_instance_id 	   = "%[2]s"
			pi_volume_ids              = []
		}
	`, name, acc.Pi_cloud_instance_id)
}

func volumeConfig(name, cloud_instance_id string) string {
	return fmt.Sprintf(`
		resource "ibm_pi_volume" "power_volume" {
			count = 3
			pi_volume_size         = 2
			pi_volume_name         = "%[1]s-${count.index}"
			pi_volume_shareable    = true
			pi_volume_pool         = "%[3]s"
			pi_cloud_instance_id   = "%[2]s"
			pi_replication_enabled = true
		}
	`, name, cloud_instance_id, acc.PiStoragePool)
}
