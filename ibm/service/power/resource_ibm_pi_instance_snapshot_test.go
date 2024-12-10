// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/power"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMPIInstanceSnapshotbasic(t *testing.T) {
	name := fmt.Sprintf("tf-pi-instance-snapshot-%d", acctest.RandIntRange(10, 100))
	snapshotRes := "ibm_pi_instance_snapshot.power_snapshot"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPIInstanceSnapshotDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIInstanceSnapshotConfig(name, power.OK),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPIInstanceSnapshotExists(snapshotRes),
					resource.TestCheckResourceAttr(snapshotRes, "pi_snapshot_name", name),
					resource.TestCheckResourceAttr(snapshotRes, "status", power.State_Available),
					resource.TestCheckResourceAttrSet(snapshotRes, "id"),
				),
			},
		},
	})
}

func TestAccIBMPIInstanceSnapshotUserTags(t *testing.T) {
	name := fmt.Sprintf("tf-pi-instance-snapshot-%d", acctest.RandIntRange(10, 100))
	snapshotRes := "ibm_pi_instance_snapshot.power_snapshot"
	userTagsString := `["env:dev","test_tag"]`
	userTagsStringUpdated := `["env:dev","test_tag","test_tag2"]`
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPIInstanceSnapshotDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIInstanceSnapshotUserTagsConfig(name, power.OK, userTagsString),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPIInstanceSnapshotExists(snapshotRes),
					resource.TestCheckResourceAttr(snapshotRes, "pi_snapshot_name", name),
					resource.TestCheckResourceAttr(snapshotRes, "status", power.State_Available),
					resource.TestCheckResourceAttrSet(snapshotRes, "id"),
					resource.TestCheckResourceAttr(snapshotRes, "pi_user_tags.#", "2"),
					resource.TestCheckTypeSetElemAttr(snapshotRes, "pi_user_tags.*", "env:dev"),
					resource.TestCheckTypeSetElemAttr(snapshotRes, "pi_user_tags.*", "test_tag"),
				),
			},
			{
				Config: testAccCheckIBMPIInstanceSnapshotUserTagsConfig(name, power.OK, userTagsStringUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPIInstanceSnapshotExists(snapshotRes),
					resource.TestCheckResourceAttr(snapshotRes, "pi_snapshot_name", name),
					resource.TestCheckResourceAttr(snapshotRes, "status", power.State_Available),
					resource.TestCheckResourceAttrSet(snapshotRes, "id"),
					resource.TestCheckResourceAttr(snapshotRes, "pi_user_tags.#", "3"),
					resource.TestCheckTypeSetElemAttr(snapshotRes, "pi_user_tags.*", "env:dev"),
					resource.TestCheckTypeSetElemAttr(snapshotRes, "pi_user_tags.*", "test_tag"),
					resource.TestCheckTypeSetElemAttr(snapshotRes, "pi_user_tags.*", "test_tag2"),
				),
			},
		},
	})
}

func testAccCheckIBMPIInstanceSnapshotDestroy(s *terraform.State) error {
	sess, err := acc.TestAccProvider.Meta().(conns.ClientSession).IBMPISession()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_pi_instance_snapshot" {
			continue
		}
		cloudInstanceID, snapshotID, err := splitID(rs.Primary.ID)
		if err != nil {
			return err
		}
		snapshotC := instance.NewIBMPISnapshotClient(context.Background(), sess, cloudInstanceID)
		_, err = snapshotC.Get(snapshotID)
		if err == nil {
			return fmt.Errorf("PI Instance Snapshot still exists: %s", rs.Primary.ID)
		}
	}
	return nil
}

func testAccCheckIBMPIInstanceSnapshotExists(n string) resource.TestCheckFunc {
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

		cloudInstanceID, snapshotID, err := splitID(rs.Primary.ID)
		if err != nil {
			return err
		}

		client := instance.NewIBMPISnapshotClient(context.Background(), sess, cloudInstanceID)

		_, err = client.Get(snapshotID)
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccCheckIBMPIInstanceSnapshotConfig(name, healthStatus string) string {
	return testAccCheckIBMPIInstanceConfig(name, healthStatus) + fmt.Sprintf(`
		resource "ibm_pi_instance_snapshot" "power_snapshot"{
			depends_on=[ibm_pi_instance.power_instance]
			pi_instance_name       = ibm_pi_instance.power_instance.pi_instance_name
			pi_cloud_instance_id   = "%s"
			pi_snapshot_name       = "%s"
			pi_volume_ids          = [ibm_pi_volume.power_volume.volume_id]
		}`, acc.Pi_cloud_instance_id, name)
}

func testAccCheckIBMPIInstanceSnapshotUserTagsConfig(name, healthStatus string, userTagsString string) string {
	return testAccCheckIBMPIInstanceConfig(name, healthStatus) + fmt.Sprintf(`
		resource "ibm_pi_instance_snapshot" "power_snapshot"{
			depends_on=[ibm_pi_instance.power_instance]
			pi_instance_name       = ibm_pi_instance.power_instance.pi_instance_name
			pi_cloud_instance_id   = "%s"
			pi_snapshot_name       = "%s"
			pi_user_tags           =  %s
			pi_volume_ids          = [ibm_pi_volume.power_volume.volume_id]
		}`, acc.Pi_cloud_instance_id, name, userTagsString)
}
