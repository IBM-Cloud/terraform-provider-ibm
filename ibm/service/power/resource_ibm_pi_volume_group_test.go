// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power_test

import (
	"context"
	"errors"
	"fmt"
	"log"
	"testing"

	st "github.com/IBM-Cloud/power-go-client/clients/instance"
	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
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
		vgC := st.NewIBMPIVolumeGroupClient(context.Background(), sess, cloudInstanceID)
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
		client := st.NewIBMPIVolumeGroupClient(context.Background(), sess, cloudInstanceID)

		_, err = client.Get(vgID)
		if err != nil {
			return err
		}
		return nil

	}
}

func testAccCheckIBMPIVolumeGroupConfig(name string) string {
	return fmt.Sprintf(`
	resource "ibm_pi_volume_group" "power_volume_group"{
		pi_volume_group_name       = "%[1]s"
		pi_cloud_instance_id 	   = "%[2]s"
		pi_volume_ids              = ["5f5c4c58-1657-433d-9556-85dc8fd97583","8tec4c58-1657-433d-9556-85dc8fd97583"]
	  }
	`, name, acc.Pi_cloud_instance_id)
}

func testAccCheckIBMPIVolumeGroupUpdateConfig(name string) string {
	return fmt.Sprintf(`
	resource "ibm_pi_volume_group" "power_volume_group"{
		pi_volume_group_name       = "%[1]s"
		pi_cloud_instance_id 	   = "%[2]s"
		pi_volume_ids              = ["q2mc4c58-1657-433d-9556-85dc8fd97583"]
	  }
	`, name, acc.Pi_cloud_instance_id)
}
