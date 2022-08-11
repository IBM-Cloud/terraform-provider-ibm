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

func TestAccIBMPICloudConnectionbasic(t *testing.T) {
	name := fmt.Sprintf("tf-cloudconnection-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPICloudConnectionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPICloudConnectionConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPICloudConnectionExists("ibm_pi_cloud_connection.cloud_connection"),
					resource.TestCheckResourceAttr("ibm_pi_cloud_connection.cloud_connection",
						"pi_cloud_connection_name", name),
					resource.TestCheckResourceAttr("ibm_pi_cloud_connection.cloud_connection",
						"pi_cloud_connection_speed", "100"),
					resource.TestCheckResourceAttr("ibm_pi_cloud_connection.cloud_connection",
						"pi_cloud_connection_networks.#", "1"),
				),
			},
		},
	})
}
func testAccCheckIBMPICloudConnectionDestroy(s *terraform.State) error {
	sess, err := acc.TestAccProvider.Meta().(conns.ClientSession).IBMPISession()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_pi_cloud_connection" {
			continue
		}
		cloudInstanceID, cloudConnectionID, err := splitID(rs.Primary.ID)
		if err != nil {
			return err
		}
		client := st.NewIBMPICloudConnectionClient(context.Background(), sess, cloudInstanceID)
		_, err = client.Get(cloudConnectionID)
		if err == nil {
			return fmt.Errorf("Cloud Connection still exists: %s", rs.Primary.ID)
		}
	}
	return nil
}
func splitID(id string) (id1, id2 string, err error) {
	parts, err := flex.IdParts(id)
	if err != nil {
		return
	}
	id1 = parts[0]
	id2 = parts[1]
	return
}
func testAccCheckIBMPICloudConnectionExists(n string) resource.TestCheckFunc {
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
		cloudInstanceID, cloudConnectionID, err := splitID(rs.Primary.ID)
		if err != nil {
			return err
		}
		client := st.NewIBMPICloudConnectionClient(context.Background(), sess, cloudInstanceID)

		_, err = client.Get(cloudConnectionID)
		if err != nil {
			return err
		}

		return nil
	}
}
func testAccCheckIBMPICloudConnectionConfig(name string) string {
	return fmt.Sprintf(`
	resource "ibm_pi_cloud_connection" "cloud_connection" {
		pi_cloud_instance_id         = "%[1]s"
		pi_cloud_connection_name     = "%[2]s"
		pi_cloud_connection_speed    = 100
		pi_cloud_connection_networks = [ibm_pi_network.network1.network_id]
	}
	resource "ibm_pi_network" "network1" {
		pi_cloud_instance_id = "%[1]s"
		pi_network_name      = "%[2]s"
		pi_network_type      = "vlan"
		pi_cidr              = "192.112.111.0/24"
	}
	`, acc.Pi_cloud_instance_id, name)
}

func TestAccIBMPICloudConnectionNetworks(t *testing.T) {
	name := fmt.Sprintf("tf-cloudconnection-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPICloudConnectionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPICloudConnectionNetworkConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPICloudConnectionExists("ibm_pi_cloud_connection.cc_network"),
					resource.TestCheckResourceAttr("ibm_pi_cloud_connection.cc_network",
						"pi_cloud_connection_name", name),
					resource.TestCheckResourceAttr("ibm_pi_cloud_connection.cc_network",
						"pi_cloud_connection_networks.#", "1"),
				),
			},
			{
				Config: testAccCheckIBMPICloudConnectionNetworkUpdateConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPICloudConnectionExists("ibm_pi_cloud_connection.cc_network"),
					resource.TestCheckResourceAttr("ibm_pi_cloud_connection.cc_network",
						"pi_cloud_connection_name", name),
					resource.TestCheckResourceAttr("ibm_pi_cloud_connection.cc_network",
						"pi_cloud_connection_networks.#", "1"),
				),
			},
		},
	})
}

func testAccCheckIBMPICloudConnectionNetworkConfig(name string) string {
	return fmt.Sprintf(`
	resource "ibm_pi_cloud_connection" "cc_network" {
		pi_cloud_instance_id         = "%[1]s"
		pi_cloud_connection_name     = "%[2]s"
		pi_cloud_connection_speed    = 1000
		pi_cloud_connection_networks = [ibm_pi_network.network1.network_id]
	}
	resource "ibm_pi_network" "network1" {
		pi_cloud_instance_id = "%[1]s"
		pi_network_name      = "%[2]s_net1"
		pi_network_type      = "vlan"
		pi_cidr              = "192.112.112.0/24"
	}
	resource "ibm_pi_network" "network2" {
		pi_cloud_instance_id = "%[1]s"
		pi_network_name      = "%[2]s_net2"
		pi_network_type      = "vlan"
		pi_cidr              = "192.112.113.0/24"
	}
	`, acc.Pi_cloud_instance_id, name)
}

func testAccCheckIBMPICloudConnectionNetworkUpdateConfig(name string) string {
	return fmt.Sprintf(`
	resource "ibm_pi_cloud_connection" "cc_network" {
		pi_cloud_instance_id         = "%[1]s"
		pi_cloud_connection_name     = "%[2]s"
		pi_cloud_connection_speed    = 1000
		pi_cloud_connection_networks = [ibm_pi_network.network2.network_id]
	}
	resource "ibm_pi_network" "network1" {
		pi_cloud_instance_id = "%[1]s"
		pi_network_name      = "%[2]s_net1"
		pi_network_type      = "vlan"
		pi_cidr              = "192.112.112.0/24"
	}
	resource "ibm_pi_network" "network2" {
		pi_cloud_instance_id = "%[1]s"
		pi_network_name      = "%[2]s_net2"
		pi_network_type      = "vlan"
		pi_cidr              = "192.112.113.0/24"
	}
	`, acc.Pi_cloud_instance_id, name)
}

func TestAccIBMPICloudConnectionClassic(t *testing.T) {
	name := fmt.Sprintf("tf-cloudconnection-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPICloudConnectionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPICloudConnectionClassicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPICloudConnectionExists("ibm_pi_cloud_connection.classic"),
					resource.TestCheckResourceAttr("ibm_pi_cloud_connection.classic",
						"pi_cloud_connection_name", name),
					resource.TestCheckResourceAttr("ibm_pi_cloud_connection.classic",
						"pi_cloud_connection_networks.#", "0"),
					resource.TestCheckResourceAttr("ibm_pi_cloud_connection.classic",
						"pi_cloud_connection_classic_enabled", "true"),
					resource.TestCheckResourceAttr("ibm_pi_cloud_connection.classic",
						"pi_cloud_connection_vpc_enabled", "false"),
				),
			},
		},
	})
}
func testAccCheckIBMPICloudConnectionClassicConfig(name string) string {
	return fmt.Sprintf(`
	resource "ibm_pi_cloud_connection" "classic" {
		pi_cloud_instance_id                = "%[1]s"
		pi_cloud_connection_name            = "%[2]s"
		pi_cloud_connection_speed           = 50
		pi_cloud_connection_classic_enabled = true
	}
	`, acc.Pi_cloud_instance_id, name)
}

func TestAccIBMPICloudConnectionVPC(t *testing.T) {
	name := fmt.Sprintf("tf-cloudconnection-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPICloudConnectionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPICloudConnectionVPCConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPICloudConnectionExists("ibm_pi_cloud_connection.vpc"),
					resource.TestCheckResourceAttr("ibm_pi_cloud_connection.vpc",
						"pi_cloud_connection_name", name),
					resource.TestCheckResourceAttr("ibm_pi_cloud_connection.vpc",
						"pi_cloud_connection_networks.#", "0"),
					resource.TestCheckResourceAttr("ibm_pi_cloud_connection.vpc",
						"pi_cloud_connection_classic_enabled", "false"),
					resource.TestCheckResourceAttr("ibm_pi_cloud_connection.vpc",
						"pi_cloud_connection_vpc_enabled", "true"),
					resource.TestCheckResourceAttr("ibm_pi_cloud_connection.vpc",
						"pi_cloud_connection_vpc_crns.#", "1"),
				),
			},
		},
	})
}

func testAccCheckIBMPICloudConnectionVPCConfig(name string) string {
	return fmt.Sprintf(`
	resource "ibm_pi_cloud_connection" "vpc" {
		pi_cloud_instance_id            = "%[1]s"
		pi_cloud_connection_name        = "%[2]s"
		pi_cloud_connection_speed       = 50
		pi_cloud_connection_vpc_enabled = true
		pi_cloud_connection_vpc_crns    = ["crn:v1:bluemix:public:is:us-south:a/d9cec80d0adc400ead8e2076afe26698::vpc:r006-6486cf73-451d-4d44-b90d-83dff504cbed"]
	}
	`, acc.Pi_cloud_instance_id, name)
}

func TestAccIBMPICloudConnectionTransitGateway(t *testing.T) {
	name := fmt.Sprintf("tf-cloudconnection-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPICloudConnectionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPICloudConnectionConfigTransitGateway(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPICloudConnectionExists("ibm_pi_cloud_connection.cloud_connection_transit"),
					resource.TestCheckResourceAttr("ibm_pi_cloud_connection.cloud_connection_transit",
						"pi_cloud_connection_name", name),
					resource.TestCheckResourceAttr("ibm_pi_cloud_connection.cloud_connection_transit",
						"pi_cloud_connection_speed", "100"),
					resource.TestCheckResourceAttr("ibm_pi_cloud_connection.cloud_connection_transit",
						"pi_cloud_connection_networks.#", "1"),
					resource.TestCheckResourceAttr("ibm_pi_cloud_connection.cloud_connection_transit",
						"pi_cloud_connection_transit_enabled", "true"),
				),
			},
		},
	})
}

func testAccCheckIBMPICloudConnectionConfigTransitGateway(name string) string {
	return fmt.Sprintf(`
	resource "ibm_pi_cloud_connection" "cloud_connection_transit" {
		pi_cloud_instance_id                = "%[1]s"
		pi_cloud_connection_name            = "%[2]s"
		pi_cloud_connection_speed           = 100
		pi_cloud_connection_networks        = [ibm_pi_network.network1.network_id]
		pi_cloud_connection_transit_enabled = true
	}
	resource "ibm_pi_network" "network1" {
		pi_cloud_instance_id = "%[1]s"
		pi_network_name      = "%[2]s"
		pi_network_type      = "vlan"
		pi_cidr              = "192.112.111.0/24"
	}
	`, acc.Pi_cloud_instance_id, name)
}
