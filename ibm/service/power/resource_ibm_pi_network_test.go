// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
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
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccIBMPINetworkbasic(t *testing.T) {
	name := fmt.Sprintf("tf-pi-network-%d", acctest.RandIntRange(10, 100))
	networkRes := "ibm_pi_network.power_networks"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPINetworkDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPINetworkConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPINetworkExists(networkRes),
					resource.TestCheckResourceAttr(networkRes, "pi_network_name", name),
					resource.TestCheckResourceAttrSet(networkRes, "id"),
					resource.TestCheckResourceAttrSet(networkRes, "pi_gateway"),
					resource.TestCheckResourceAttrSet(networkRes, "pi_ipaddress_range.#"),
				),
			},
			{
				Config: testAccCheckIBMPINetworkConfigUpdateDNS(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPINetworkExists(networkRes),
					resource.TestCheckResourceAttr(networkRes, "pi_network_name", name),
					resource.TestCheckResourceAttr(networkRes, "pi_dns.#", "1"),
					resource.TestCheckResourceAttrSet(networkRes, "id"),
					resource.TestCheckResourceAttrSet(networkRes, "pi_gateway"),
					resource.TestCheckResourceAttrSet(networkRes, "pi_ipaddress_range.#"),
				),
			},
		},
	})
}

func TestAccIBMPINetworkGatewaybasic(t *testing.T) {
	name := fmt.Sprintf("tf-pi-network-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPINetworkDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPINetworkGatewayConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPINetworkExists("ibm_pi_network.power_networks"),
					resource.TestCheckResourceAttr(
						"ibm_pi_network.power_networks", "pi_network_name", name),
					resource.TestCheckResourceAttrSet("ibm_pi_network.power_networks", "pi_gateway"),
					resource.TestCheckResourceAttrSet("ibm_pi_network.power_networks", "id"),
					resource.TestCheckResourceAttrSet("ibm_pi_network.power_networks", "pi_ipaddress_range.#"),
				),
			},
			{
				Config: testAccCheckIBMPINetworkConfigGatewayUpdateDNS(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPINetworkExists("ibm_pi_network.power_networks"),
					resource.TestCheckResourceAttr(
						"ibm_pi_network.power_networks", "pi_network_name", name),
					resource.TestCheckResourceAttr(
						"ibm_pi_network.power_networks", "pi_gateway", "192.168.17.2"),
					resource.TestCheckResourceAttr(
						"ibm_pi_network.power_networks", "pi_ipaddress_range.0.pi_ending_ip_address", "192.168.17.254"),
					resource.TestCheckResourceAttr(
						"ibm_pi_network.power_networks", "pi_ipaddress_range.0.pi_starting_ip_address", "192.168.17.3"),
				),
			},
		},
	})
}

func TestAccIBMPINetworkGatewaybasicSatellite(t *testing.T) {
	name := fmt.Sprintf("tf-pi-network-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPINetworkDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPINetworkGatewayConfigSatellite(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPINetworkExists("ibm_pi_network.power_networks"),
					resource.TestCheckResourceAttr(
						"ibm_pi_network.power_networks", "pi_network_name", name),
					resource.TestCheckResourceAttrSet("ibm_pi_network.power_networks", "id"),
					resource.TestCheckResourceAttrSet("ibm_pi_network.power_networks", "pi_ipaddress_range.#"),
				),
			},
			{
				Config: testAccCheckIBMPINetworkConfigGatewayUpdateDNS(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPINetworkExists("ibm_pi_network.power_networks"),
					resource.TestCheckResourceAttr(
						"ibm_pi_network.power_networks", "pi_network_name", name),
					resource.TestCheckResourceAttr(
						"ibm_pi_network.power_networks", "pi_ipaddress_range.0.pi_ending_ip_address", "192.168.17.254"),
					resource.TestCheckResourceAttr(
						"ibm_pi_network.power_networks", "pi_ipaddress_range.0.pi_starting_ip_address", "192.168.17.3"),
				),
			},
		},
	})
}

func TestAccIBMPINetworkUserTags(t *testing.T) {
	name := fmt.Sprintf("tf-pi-network-%d", acctest.RandIntRange(10, 100))
	networkRes := "ibm_pi_network.power_networks"
	userTagsString := `["env:dev","test_tag"]`
	userTagsStringUpdated := `["env:dev","test_tag","test_tag2"]`
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPINetworkDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPINetworkUserTagsConfig(name, userTagsString),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPINetworkExists(networkRes),
					resource.TestCheckResourceAttr(networkRes, "pi_network_name", name),
					resource.TestCheckResourceAttrSet(networkRes, "id"),
					resource.TestCheckResourceAttr(networkRes, "pi_user_tags.#", "2"),
					resource.TestCheckTypeSetElemAttr(networkRes, "pi_user_tags.*", "env:dev"),
					resource.TestCheckTypeSetElemAttr(networkRes, "pi_user_tags.*", "test_tag"),
				),
			},
			{
				Config: testAccCheckIBMPINetworkUserTagsConfig(name, userTagsStringUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPINetworkExists(networkRes),
					resource.TestCheckResourceAttr(networkRes, "pi_network_name", name),
					resource.TestCheckResourceAttrSet(networkRes, "id"),
					resource.TestCheckResourceAttr(networkRes, "pi_user_tags.#", "3"),
					resource.TestCheckTypeSetElemAttr(networkRes, "pi_user_tags.*", "env:dev"),
					resource.TestCheckTypeSetElemAttr(networkRes, "pi_user_tags.*", "test_tag"),
					resource.TestCheckTypeSetElemAttr(networkRes, "pi_user_tags.*", "test_tag2"),
				),
			},
		},
	})
}

func TestAccIBMPINetworkAdvertiseArpBroadcast(t *testing.T) {
	name := fmt.Sprintf("tf-pi-network-%d", acctest.RandIntRange(10, 100))
	networkRes := "ibm_pi_network.power_network_advertise_arpbroadcast"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPINetworkDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPINetworkAdvertiseArpBroadcast(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPINetworkExists(networkRes),
					resource.TestCheckResourceAttr(networkRes, "pi_network_name", name),
					resource.TestCheckResourceAttrSet(networkRes, "id"),
					resource.TestCheckResourceAttr(networkRes, "pi_advertise", "disable"),
					resource.TestCheckResourceAttr(networkRes, "pi_arp_broadcast", "enable"),
				),
			},
			{
				Config: testAccCheckIBMPINetworkAdvertiseArpBroadcastUpdate(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPINetworkExists(networkRes),
					resource.TestCheckResourceAttr(networkRes, "pi_network_name", name),
					resource.TestCheckResourceAttrSet(networkRes, "id"),
					resource.TestCheckResourceAttr(networkRes, "pi_advertise", "enable"),
					resource.TestCheckResourceAttr(networkRes, "pi_arp_broadcast", "disable"),
				),
			},
		},
	})
}

func testAccCheckIBMPINetworkDestroy(s *terraform.State) error {
	sess, err := acc.TestAccProvider.Meta().(conns.ClientSession).IBMPISession()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_pi_network" {
			continue
		}
		cloudInstanceID, networkID, err := splitID(rs.Primary.ID)
		if err != nil {
			return err
		}
		networkC := instance.NewIBMPINetworkClient(context.Background(), sess, cloudInstanceID)
		_, err = networkC.Get(networkID)
		if err == nil {
			return fmt.Errorf("PI Network still exists: %s", rs.Primary.ID)
		}
	}

	return nil
}

func testAccCheckIBMPINetworkExists(n string) resource.TestCheckFunc {
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
		cloudInstanceID, networkID, err := splitID(rs.Primary.ID)
		if err != nil {
			return err
		}
		client := instance.NewIBMPINetworkClient(context.Background(), sess, cloudInstanceID)

		_, err = client.Get(networkID)
		if err != nil {
			return err
		}
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
	`, acc.Pi_cloud_instance_id, name)
}

func testAccCheckIBMPINetworkConfigUpdateDNS(name string) string {
	return fmt.Sprintf(`
		resource "ibm_pi_network" "power_networks" {
			pi_cloud_instance_id = "%s"
			pi_network_name      = "%s"
			pi_network_type      = "pub-vlan"
			pi_dns               = ["127.0.0.1"]
		}
	`, acc.Pi_cloud_instance_id, name)
}

func testAccCheckIBMPINetworkGatewayConfig(name string) string {
	return fmt.Sprintf(`
		resource "ibm_pi_network" "power_networks" {
			pi_cloud_instance_id = "%s"		
			pi_cidr              = "192.168.17.0/24"
			pi_gateway           = "192.168.17.1"
			pi_network_name      = "%s"
			pi_network_type      = "vlan"
		}
	`, acc.Pi_cloud_instance_id, name)
}

func testAccCheckIBMPINetworkGatewayConfigSatellite(name string) string {
	return fmt.Sprintf(`
		resource "ibm_pi_network" "power_networks" {
			pi_cloud_instance_id 		= "%s"
			pi_network_name      		= "%s"
			pi_network_type      		= "vlan"
			pi_cidr              		= "192.168.17.0/24"
			pi_network_mtu		 		= 6500
		}
	`, acc.Pi_cloud_instance_id, name)
}

func testAccCheckIBMPINetworkConfigGatewayUpdateDNS(name string) string {
	return fmt.Sprintf(`
		resource "ibm_pi_network" "power_networks" {
			pi_cloud_instance_id = "%s"
			pi_cidr              = "192.168.17.0/24"
			pi_dns               = ["127.0.0.1"]
			pi_gateway           = "192.168.17.2"
			pi_network_name      = "%s"
			pi_network_type      = "vlan"
			pi_ipaddress_range {
				pi_ending_ip_address = "192.168.17.254"
				pi_starting_ip_address = "192.168.17.3"
			}
		}
	`, acc.Pi_cloud_instance_id, name)
}

func testAccCheckIBMPINetworkUserTagsConfig(name string, userTagsString string) string {
	return fmt.Sprintf(`
		resource "ibm_pi_network" "power_networks" {
			pi_cloud_instance_id = "%s"
			pi_network_name      = "%s"
			pi_network_type      = "pub-vlan"
			pi_user_tags         = %s
		}
	`, acc.Pi_cloud_instance_id, name, userTagsString)
}

func testAccCheckIBMPINetworkAdvertiseArpBroadcast(name string) string {
	return fmt.Sprintf(`
		resource "ibm_pi_network" "power_network_advertise_arpbroadcast" {
			pi_cloud_instance_id 		= "%s"
			pi_advertise 				= "disable"
			pi_arp_broadcast 			= "enable"
			pi_cidr                     = "192.168.17.0/24"
			pi_network_name      		= "%s"
			pi_network_type      		= "vlan"
		}
	`, acc.Pi_cloud_instance_id, name)
}

func testAccCheckIBMPINetworkAdvertiseArpBroadcastUpdate(name string) string {
	return fmt.Sprintf(`
		resource "ibm_pi_network" "power_network_advertise_arpbroadcast" {
			pi_cloud_instance_id 		= "%s"
			pi_advertise 				= "enable"
			pi_arp_broadcast 			= "disable"
			pi_cidr                     = "192.168.17.0/24"
			pi_network_name      		= "%s"
			pi_network_type      		= "vlan"
		}
	`, acc.Pi_cloud_instance_id, name)
}
