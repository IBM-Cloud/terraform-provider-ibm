// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/IBM-Cloud/power-go-client/clients/instance"
	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccIBMIHostBasic(t *testing.T) {
	displayName := fmt.Sprintf("tf_display_name_%d", acctest.RandIntRange(10, 100))
	hostRes := "ibm_pi_host.host"
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acc.TestAccPreCheck(t)
		},
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPIHostDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIHostConfig(displayName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMPIHostExists(hostRes),
					resource.TestCheckResourceAttr(hostRes, "display_name", displayName),
				),
			},
		},
	})
}

func testAccCheckIBMPIHostConfig(name string) string {
	return fmt.Sprintf(`
	resource "ibm_pi_host" "host" {
		pi_cloud_instance_id = "%[1]s"
		pi_host            {
		  display_name = "%[2]s"
		  sys_type = "s922"
		}
		pi_host_group_id = "%[3]s"
	  }
	`, acc.Pi_cloud_instance_id, name, acc.Pi_host_group_id)
}

func TestAccIBMIHostUserTags(t *testing.T) {
	displayName := fmt.Sprintf("tf_display_name_%d", acctest.RandIntRange(10, 100))
	hostRes := "ibm_pi_host.host"
	userTagsString := `["env:dev", "test_tag1"]`
	userTagsStringUpdated := `["env:dev", "test_tag1", "ibm"]`
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acc.TestAccPreCheck(t)
		},
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPIHostDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIHostUserTagsConfig(displayName, userTagsString),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMPIHostExists(hostRes),
					resource.TestCheckResourceAttr(hostRes, "display_name", displayName),
					resource.TestCheckResourceAttr(hostRes, "user_tags.#", "2"),
					resource.TestCheckTypeSetElemAttr(hostRes, "user_tags.*", "env:dev"),
					resource.TestCheckTypeSetElemAttr(hostRes, "user_tags.*", "test_tag1"),
				),
			},
			{
				Config: testAccCheckIBMPIHostUserTagsConfig(displayName, userTagsStringUpdated),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMPIHostExists(hostRes),
					resource.TestCheckResourceAttr(hostRes, "display_name", displayName),
					resource.TestCheckResourceAttr(hostRes, "user_tags.#", "3"),
					resource.TestCheckTypeSetElemAttr(hostRes, "user_tags.*", "env:dev"),
					resource.TestCheckTypeSetElemAttr(hostRes, "user_tags.*", "test_tag1"),
					resource.TestCheckTypeSetElemAttr(hostRes, "user_tags.*", "ibm"),
				),
			},
		},
	})
}

func testAccCheckIBMPIHostUserTagsConfig(name string, userTagsString string) string {
	return fmt.Sprintf(`
	resource "ibm_pi_host" "host" {
		pi_cloud_instance_id = "%[1]s"
		pi_host            {
		  display_name = "%[2]s"
		  sys_type = "s922"
		  user_tags = %[4]s
		}
		pi_host_group_id = "%[3]s"
	  }
	`, acc.Pi_cloud_instance_id, name, acc.Pi_host_group_id, userTagsString)
}

func testAccCheckIBMPIHostDestroy(s *terraform.State) error {
	sess, err := acc.TestAccProvider.Meta().(conns.ClientSession).IBMPISession()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_pi_host" {
			continue
		}
		cloudInstanceID, hostID, err := splitID(rs.Primary.ID)
		if err != nil {
			return err
		}
		client := instance.NewIBMPIHostGroupsClient(context.Background(), sess, cloudInstanceID)
		_, err = client.GetHost(hostID)
		if err == nil {
			return fmt.Errorf("PI dedicated host still exists: %s", rs.Primary.ID)
		}
	}
	return nil
}

func testAccCheckIBMPIHostExists(n string) resource.TestCheckFunc {
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
		cloudInstanceID, hostID, err := splitID(rs.Primary.ID)
		if err != nil {
			return err
		}
		client := instance.NewIBMPIHostGroupsClient(context.Background(), sess, cloudInstanceID)
		_, err = client.GetHost(hostID)
		if err != nil {
			return err
		}

		return nil
	}
}
