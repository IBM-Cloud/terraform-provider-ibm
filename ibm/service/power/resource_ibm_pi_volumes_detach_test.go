// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/IBM-Cloud/power-go-client/clients/instance"
	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMPIVolumesDetach(t *testing.T) {
	name := fmt.Sprintf("tf-pi-instance-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPIVolumesDetachDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIVolumesDetachAllVolumesConfig(name, "WARNING"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("ibm_pi_volumes_detach.pi_volumes_detach_instance", "id"),
				),
			},
		},
	})
}

func testAccCheckIBMPIVolumesDetachAllVolumesConfig(name, instanceHealthStatus string) string {
	return fmt.Sprintf(`
	resource "ibm_pi_key" "key" {
		pi_cloud_instance_id = "%[1]s"
		pi_key_name          = "%[2]s"
		pi_ssh_key           = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR"
	}
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
		pi_volume_name       = "%[2]s"
		pi_volume_pool       = data.ibm_pi_image.power_image.storage_pool
		pi_volume_shareable  = true
		pi_volume_size       = 20
		pi_volume_type       = "%[6]s"
	}
	resource "ibm_pi_volume" "power_volume_2" {
		pi_cloud_instance_id = "%[1]s"
		pi_volume_name       = "%[2]s-2"
		pi_volume_pool       = data.ibm_pi_image.power_image.storage_pool
		pi_volume_shareable  = true
		pi_volume_size       = 20
		pi_volume_type       = "%[6]s"
	}
	resource "ibm_pi_volume" "power_volume_3" {
		pi_cloud_instance_id = "%[1]s"
		pi_volume_name       = "%[2]s-3"
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
		pi_key_pair_name      = ibm_pi_key.key.name
		pi_memory             = "2"
		pi_proc_type          = "shared"
		pi_processors         = "0.25"
		pi_storage_pool       = data.ibm_pi_image.power_image.storage_pool
		pi_storage_type       = "%[6]s"
		pi_sys_type           = "s922"
		pi_volume_ids         = [ibm_pi_volume.power_volume.volume_id, ibm_pi_volume.power_volume_2.volume_id, ibm_pi_volume.power_volume_3.volume_id]
		pi_network {
			network_id = data.ibm_pi_network.power_networks.id
		}
	}
	data "ibm_pi_instance" "power_instance" {
		pi_cloud_instance_id = "%[1]s"
		pi_instance_name      = resource.ibm_pi_instance.power_instance.pi_instance_name
	}
	resource "ibm_pi_volumes_detach" "pi_volumes_detach_instance" {
		pi_cloud_instance_id = "%[1]s"
		pi_detach_all_volumes = true
		pi_instance_id       = data.ibm_pi_instance.power_instance.id
	}
	`, acc.Pi_cloud_instance_id, name, acc.Pi_image, acc.Pi_network_name, instanceHealthStatus, acc.PiStorageType)
}

func testAccCheckIBMPIVolumesDetachDestroy(s *terraform.State) error {
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
		cloudInstanceID, pvmInstanceID := ids[0], ids[1]
		client := instance.NewIBMPIInstanceClient(context.Background(), sess, cloudInstanceID)
		pvm, err := client.Get(pvmInstanceID)
		if err == nil {
			if len(pvm.VolumeIDs) > 1 {
				return fmt.Errorf("PI Volumes Attach still exists: %s", rs.Primary.ID)
			}

		}
	}

	return nil
}
