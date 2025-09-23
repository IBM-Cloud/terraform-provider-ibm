// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/IBM-Cloud/power-go-client/clients/instance"
	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/flex"
)

func TestAccIBMPIVolumeGroupActionbasic(t *testing.T) {
	name := fmt.Sprintf("tf-pi-volume-group-action-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIVolumeGroupStopActionConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPIVolumeGroupActionExists("ibm_pi_volume_group_action.power_volume_group_action"),
					resource.TestCheckResourceAttrSet("ibm_pi_volume_group_action.power_volume_group_action", "id"),
					resource.TestCheckResourceAttrSet("ibm_pi_volume_group_action.power_volume_group_action", "volume_group_status"),
				),
			},
			{
				Config: testAccCheckIBMPIVolumeGroupStartActionConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPIVolumeGroupActionExists("ibm_pi_volume_group_action.power_volume_group_action"),
					resource.TestCheckResourceAttrSet("ibm_pi_volume_group_action.power_volume_group_action", "id"),
					resource.TestCheckResourceAttrSet("ibm_pi_volume_group_action.power_volume_group_action", "volume_group_status"),
				),
			},
		},
	})
}

func testAccCheckIBMPIVolumeGroupActionExists(n string) resource.TestCheckFunc {
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
		cloudInstanceID, vgID := ids[0], ids[1]
		client := instance.NewIBMPIVolumeGroupClient(context.Background(), sess, cloudInstanceID)

		_, err = client.Get(vgID)
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccCheckIBMPIVolumeGroupStopActionConfig(name string) string {
	return testAccCheckIBMPIVolumeGroupConfig(name) + fmt.Sprintf(`
	  resource "ibm_pi_volume_group_action" "power_volume_group_action" {
		pi_cloud_instance_id   = "%[1]s"
		pi_volume_group_id     = ibm_pi_volume_group.power_volume_group.volume_group_id
		pi_volume_group_action {
			stop {
				access = true
			}
		}
	  }
	`, acc.Pi_cloud_instance_id)
}

func testAccCheckIBMPIVolumeGroupStartActionConfig(name string) string {
	return testAccCheckIBMPIVolumeGroupConfig(name) + fmt.Sprintf(`
	  resource "ibm_pi_volume_group_action" "power_volume_group_action" {
		pi_cloud_instance_id   = "%[1]s"
		pi_volume_group_id     = ibm_pi_volume_group.power_volume_group.volume_group_id
		pi_volume_group_action {
			start {
				source = "master"
			}
		}
	  }
	`, acc.Pi_cloud_instance_id)
}
