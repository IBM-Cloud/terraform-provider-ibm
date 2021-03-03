// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"errors"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	st "github.com/IBM-Cloud/power-go-client/clients/instance"
)

func TestAccIBMPIInstanceSnapshotbasic(t *testing.T) {

	name := fmt.Sprintf("tf-pi-instance-snapshot-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMPIInstanceSnapshotDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIInstanceSnapshotConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPIInstanceSnapshotExists("ibm_pi_snapshot.power_snapshot"),
					resource.TestCheckResourceAttr(
						"ibm_pi_snapshot.power_snapshot", "pi_snap_shot_name", name),
				),
			},
		},
	})
}
func testAccCheckIBMPIInstanceSnapshotDestroy(s *terraform.State) error {

	sess, err := testAccProvider.Meta().(ClientSession).IBMPISession()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_pi_snapshot" {
			continue
		}
		parts, err := idParts(rs.Primary.ID)
		powerinstanceid := parts[0]
		networkC := st.NewIBMPISnapshotClient(sess, powerinstanceid)
		_, err = networkC.Get(parts[1], powerinstanceid, getTimeOut)
		if err == nil {
			return fmt.Errorf("PI Instance Snapshot still exists: %s", rs.Primary.ID)
		}
	}

	return nil
}
func testAccCheckIBMPIInstanceSnapshotExists(n string) resource.TestCheckFunc {
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
		client := st.NewIBMPISnapshotClient(sess, powerinstanceid)

		snapshot, err := client.Get(parts[1], powerinstanceid, getTimeOut)
		if err != nil {
			return err
		}
		parts[1] = *snapshot.SnapshotID
		return nil

	}
}

func testAccCheckIBMPIInstanceSnapshotConfig(name string) string {
	return testAccCheckIBMPIInstanceConfig(name) + fmt.Sprintf(`
	  resource "ibm_pi_snapshot" "power_snapshot"{
		depends_on=[ibm_pi_instance.power_instance]
		pi_instance_name       = ibm_pi_instance.power_instance.pi_instance_name
		pi_cloud_instance_id = "%s"
		pi_snap_shot_name       = "%s"
		pi_volume_ids         = [ibm_pi_volume.power_volume.volume_id]
	  }
	`, pi_cloud_instance_id, name)
}
