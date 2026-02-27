// Copyright IBM Corp. 2024 All Rights Reserved.
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
	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
)

func TestAccIBMPIHostGroupBasic(t *testing.T) {
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	displayName := fmt.Sprintf("tf_display_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPIHostGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIHostGroupConfigBasic(name, displayName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMPIHostGroupExists("ibm_pi_host_group.hostGroup"),
					resource.TestCheckResourceAttr("ibm_pi_host_group.hostGroup", "pi_name", name),
				),
			},
		},
	})
}

func testAccCheckIBMPIHostGroupConfigBasic(name, displayName string) string {
	return fmt.Sprintf(`
		resource "ibm_pi_host_group" "hostGroup" {
			pi_cloud_instance_id = "%[1]s"
			pi_hosts {
				display_name = "%[2]s"
				sys_type = "s922"
			}
			pi_name = "%[3]s"
		}
	`, acc.Pi_cloud_instance_id, displayName, name)
}
func TestAccIBMPIHostGroupUpdateSecondaries(t *testing.T) {
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	displayName := fmt.Sprintf("tf_display_name_%d", acctest.RandIntRange(10, 100))
	hostGroupRes := "ibm_pi_host_group.hostGroup"

	resource.Test(t, resource.TestCase{
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPIHostGroupDestroy,
		Steps: []resource.TestStep{
			{
				// Create host group with one secondary.
				Config: testAccCheckIBMPIHostGroupConfigOneSecondary(name, displayName, acc.Pi_secondary_workspace_id_1),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMPIHostGroupExists(hostGroupRes),
					resource.TestCheckResourceAttr(hostGroupRes, "pi_name", name),
					resource.TestCheckResourceAttr(hostGroupRes, "pi_secondaries.#", "1"),
				),
			},
			{
				// Add a second secondary workspace.
				Config: testAccCheckIBMPIHostGroupConfigTwoSecondaries(name, displayName, acc.Pi_secondary_workspace_id_1, acc.Pi_secondary_workspace_id_2),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMPIHostGroupExists(hostGroupRes),
					resource.TestCheckResourceAttr(hostGroupRes, "secondaries.#", "2"),
				),
			},
			{
				// Remove the first secondary, keeping only the second.
				Config: testAccCheckIBMPIHostGroupConfigOneSecondary(name, displayName, acc.Pi_secondary_workspace_id_2),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMPIHostGroupExists(hostGroupRes),
					resource.TestCheckResourceAttr(hostGroupRes, "secondaries.#", "1"),
				),
			},
			{
				// Remove all secondaries.
				Config: testAccCheckIBMPIHostGroupConfigBasic(name, displayName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMPIHostGroupExists(hostGroupRes),
					resource.TestCheckResourceAttr(hostGroupRes, "secondaries.#", "0"),
				),
			},
		},
	})
}

func testAccCheckIBMPIHostGroupConfigOneSecondary(name, displayName, secondaryWorkspaceID string) string {
	return fmt.Sprintf(`
		resource "ibm_pi_host_group" "hostGroup" {
			pi_cloud_instance_id = "%[1]s"
			pi_hosts {
				display_name = "%[2]s"
				sys_type = "s922"
			}
			pi_name = "%[3]s"
			pi_secondaries {
				name      = "%[3]s-secondary"
				workspace = "%[4]s"
			}
		}
	`, acc.Pi_cloud_instance_id, displayName, name, secondaryWorkspaceID)
}

func testAccCheckIBMPIHostGroupConfigTwoSecondaries(name, displayName, secondaryWorkspaceID1, secondaryWorkspaceID2 string) string {
	return fmt.Sprintf(`
		resource "ibm_pi_host_group" "hostGroup" {
			pi_cloud_instance_id = "%[1]s"
			pi_hosts {
				display_name = "%[2]s"
				sys_type = "s922"
			}
			pi_name = "%[3]s"
			pi_secondaries {
				name      = "%[3]s-secondary-1"
				workspace = "%[4]s"
			}
			pi_secondaries {
				name      = "%[3]s-secondary-2"
				workspace = "%[5]s"
			}
		}
	`, acc.Pi_cloud_instance_id, displayName, name, secondaryWorkspaceID1, secondaryWorkspaceID2)
}
func testAccCheckIBMPIHostGroupExists(n string) resource.TestCheckFunc {

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
		cloudInstanceID, hostGroupID, err := splitID(rs.Primary.ID)
		if err != nil {
			return err
		}
		client := instance.NewIBMPIHostGroupsClient(context.Background(), sess, cloudInstanceID)
		_, err = client.GetHostGroup(hostGroupID)
		if err != nil {
			return err
		}

		return nil
	}
}

func testAccCheckIBMPIHostGroupDestroy(s *terraform.State) error {
	sess, err := acc.TestAccProvider.Meta().(conns.ClientSession).IBMPISession()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_pi_host_group" {
			continue
		}
		cloudInstanceID, hostGroupID, err := splitID(rs.Primary.ID)
		if err != nil {
			return err
		}
		client := instance.NewIBMPIHostGroupsClient(context.Background(), sess, cloudInstanceID)
		_, err = client.GetHostGroup(hostGroupID)
		if err == nil {
			return fmt.Errorf("PI host group still exists: %s", rs.Primary.ID)
		}
	}
	return nil
}
