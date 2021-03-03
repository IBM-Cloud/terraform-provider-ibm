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

func TestAccIBMPINetworkbasic(t *testing.T) {
	name := fmt.Sprintf("tf-pi-network-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMPINetworkDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPINetworkConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPINetworkExists("ibm_pi_network.power_networks"),
					resource.TestCheckResourceAttr(
						"ibm_pi_network.power_networks", "pi_network_name", name),
				),
			},
		},
	})
}
func testAccCheckIBMPINetworkDestroy(s *terraform.State) error {

	sess, err := testAccProvider.Meta().(ClientSession).IBMPISession()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_pi_network" {
			continue
		}
		parts, err := idParts(rs.Primary.ID)
		powerinstanceid := parts[0]
		networkC := st.NewIBMPINetworkClient(sess, powerinstanceid)
		_, err = networkC.Get(parts[1], powerinstanceid, getTimeOut)
		if err == nil {
			return fmt.Errorf("PI Network still exists: %s", rs.Primary.ID)
		}
	}

	return nil
}
func testAccCheckIBMPINetworkExists(n string) resource.TestCheckFunc {
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
		client := st.NewIBMPINetworkClient(sess, powerinstanceid)

		network, err := client.Get(parts[1], powerinstanceid, getTimeOut)
		if err != nil {
			return err
		}
		parts[1] = *network.NetworkID
		return nil

	}
}

func testAccCheckIBMPINetworkConfig(name string) string {
	return fmt.Sprintf(`
		resource "ibm_pi_network" "power_networks" {
			pi_cloud_instance_id = "%s"
			pi_network_name      = "%s"
			pi_network_type      = "pub-vlan"
		}
	`, pi_cloud_instance_id, name)
}
