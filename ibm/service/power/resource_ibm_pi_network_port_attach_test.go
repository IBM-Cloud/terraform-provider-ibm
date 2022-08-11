// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	st "github.com/IBM-Cloud/power-go-client/clients/instance"
)

func TestAccIBMPINetworkPortAttachbasic(t *testing.T) {
	name := fmt.Sprintf("tf-pi-network-port-attach-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPINetworkPortAttachDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPINetworkPortAttachConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPINetworkPortAttachExists("ibm_pi_network_port_attach.power_network_port_attach"),
					resource.TestCheckResourceAttr(
						"ibm_pi_network_port_attach.power_network_port_attach", "pi_network_name", name),
					resource.TestCheckResourceAttrSet("ibm_pi_network_port_attach.power_network_port_attach", "id"),
					resource.TestCheckResourceAttrSet("ibm_pi_network_port_attach.power_network_port_attach", "network_port_id"),
					resource.TestCheckResourceAttrSet("ibm_pi_network_port_attach.power_network_port_attach", "public_ip"),
				),
			},
		},
	})
}

func TestAccIBMPINetworkPortAttachVlanbasic(t *testing.T) {
	name := fmt.Sprintf("tf-pi-network-port-attach-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPINetworkPortAttachDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPINetworkPortAttachVlanConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPINetworkPortAttachExists("ibm_pi_network_port_attach.power_network_port_attach"),
					resource.TestCheckResourceAttr(
						"ibm_pi_network_port_attach.power_network_port_attach", "pi_network_name", name),
					resource.TestCheckResourceAttrSet("ibm_pi_network_port_attach.power_network_port_attach", "id"),
					resource.TestCheckResourceAttrSet("ibm_pi_network_port_attach.power_network_port_attach", "network_port_id"),
				),
			},
		},
	})
}
func testAccCheckIBMPINetworkPortAttachDestroy(s *terraform.State) error {
	sess, err := acc.TestAccProvider.Meta().(conns.ClientSession).IBMPISession()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_pi_network_port_attach" {
			continue
		}
		parts, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}
		cloudInstanceID := parts[0]
		networkname := parts[1]
		portID := parts[2]
		networkC := st.NewIBMPINetworkClient(context.Background(), sess, cloudInstanceID)
		_, err = networkC.GetPort(networkname, portID)
		if err == nil {
			return fmt.Errorf("PI Network Port still exists: %s", rs.Primary.ID)
		}
	}

	return nil
}
func testAccCheckIBMPINetworkPortAttachExists(n string) resource.TestCheckFunc {
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
		parts, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}
		cloudInstanceID := parts[0]
		networkname := parts[1]
		portID := parts[2]
		client := st.NewIBMPINetworkClient(context.Background(), sess, cloudInstanceID)

		_, err = client.GetPort(networkname, portID)
		if err != nil {
			return err
		}
		return nil

	}
}

func testAccCheckIBMPINetworkPortAttachConfig(name string) string {
	return testAccCheckIBMPINetworkConfig(name) + fmt.Sprintf(`
	resource "ibm_pi_network_port_attach" "power_network_port_attach" {
		pi_cloud_instance_id  = "%s"
		pi_network_name       = ibm_pi_network.power_networks.pi_network_name
		pi_network_port_description = "IP Reserved for Test UAT"
		pi_instance_id = "%s"
	}
	`, acc.Pi_cloud_instance_id, acc.Pi_instance_name)
}

func testAccCheckIBMPINetworkPortAttachVlanConfig(name string) string {
	return testAccCheckIBMPINetworkGatewayConfig(name) + fmt.Sprintf(`
	resource "ibm_pi_network_port_attach" "power_network_port_attach" {
		pi_cloud_instance_id  = "%s"
		pi_network_name       = ibm_pi_network.power_networks.pi_network_name
		pi_network_port_description = "IP Reserved for Test UAT"
		pi_instance_id = "%s"
	}
	`, acc.Pi_cloud_instance_id, acc.Pi_instance_name)
}
