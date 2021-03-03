// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"errors"
	"fmt"
	"log"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	st "github.com/IBM-Cloud/power-go-client/clients/instance"
)

func TestAccIBMPIVolumebasic(t *testing.T) {
	name := fmt.Sprintf("tf-pi-volume-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMPIVolumeDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIVolumeConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPIVolumeExists("ibm_pi_volume.power_volume"),
					resource.TestCheckResourceAttr(
						"ibm_pi_volume.power_volume", "pi_volume_name", name),
				),
			},
		},
	})
}
func testAccCheckIBMPIVolumeDestroy(s *terraform.State) error {

	sess, err := testAccProvider.Meta().(ClientSession).IBMPISession()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_pi_volume" {
			continue
		}
		parts, err := idParts(rs.Primary.ID)
		powerinstanceid := parts[0]
		volumeC := st.NewIBMPIVolumeClient(sess, powerinstanceid)
		volume, err := volumeC.Get(parts[1], powerinstanceid, volGetTimeOut)
		if err == nil {
			log.Println("volume*****", volume.State)
			return fmt.Errorf("PI Volume still exists: %s", rs.Primary.ID)
		}
	}

	return nil
}
func testAccCheckIBMPIVolumeExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		fmt.Println("siv ", s.RootModule().Resources)
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}

		sess, err := testAccProvider.Meta().(ClientSession).IBMPISession()
		if err != nil {
			return err
		}
		parts, err := idParts(rs.Primary.ID)
		if err != nil {
			return err
		}
		powerinstanceid := parts[0]
		client := st.NewIBMPIVolumeClient(sess, powerinstanceid)

		volume, err := client.Get(parts[1], powerinstanceid, volGetTimeOut)
		if err != nil {
			return err
		}
		parts[1] = *volume.VolumeID
		return nil

	}
}

func testAccCheckIBMPIVolumeConfig(name string) string {
	return fmt.Sprintf(`
	resource "ibm_pi_volume" "power_volume"{
		pi_volume_size       = 20
		pi_volume_name       = "%s"
		pi_volume_type       = "tier1"
		pi_volume_shareable  = true
		pi_cloud_instance_id = "%s"
	  }
	`, name, pi_cloud_instance_id)
}
