// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"

	"github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/flex"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMPIVolumeClone(t *testing.T) {
	resVolumeClone := "ibm_pi_volume_clone.power_volume_clone"
	name := fmt.Sprintf("tf-pi-volume-clone-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIVolumeCloneConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPIVolumeCloneExists(resVolumeClone),
					resource.TestCheckResourceAttrSet(resVolumeClone, "id"),
					resource.TestCheckResourceAttrSet(resVolumeClone, "status"),
					resource.TestCheckResourceAttr(resVolumeClone, "status", "completed"),
					resource.TestCheckResourceAttrSet(resVolumeClone, "percent_complete"),
					resource.TestCheckResourceAttr(resVolumeClone, "percent_complete", "100"),
				),
			},
		},
	})
}

func testAccCheckIBMPIVolumeCloneExists(n string) resource.TestCheckFunc {
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
		cloudInstanceID, vcTaskID := ids[0], ids[1]
		client := instance.NewIBMPICloneVolumeClient(context.Background(), sess, cloudInstanceID)

		_, err = client.Get(vcTaskID)
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccCheckIBMPIVolumeCloneConfig(name string) string {
	return volumesCloneConfig(name, true) + fmt.Sprintf(`
		resource "ibm_pi_volume_clone" "power_volume_clone" {
			pi_cloud_instance_id   			= "%[1]s"
			pi_replication_enabled 			= %[4]v
			pi_target_storage_tier 			= "%[3]s"
			pi_volume_clone_name     		= "%[2]s"
			pi_volume_ids 					= ibm_pi_volume.power_volume.*.volume_id
		} `, acc.Pi_cloud_instance_id, name, acc.Pi_target_storage_tier, false)
}

func volumesCloneConfig(name string, volumeReplicationEnabled bool) string {
	return fmt.Sprintf(`
		resource "ibm_pi_volume" "power_volume" {
			count = 2
			pi_cloud_instance_id   = "%[2]s"
			pi_replication_enabled = %[4]v
			pi_volume_name         = "%[1]s-${count.index}"
			pi_volume_pool         = "%[3]s"
			pi_volume_size         = 2
		} `, name, acc.Pi_cloud_instance_id, acc.PiStoragePool, volumeReplicationEnabled)
}
