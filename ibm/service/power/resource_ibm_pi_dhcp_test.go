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

func TestAccIBMPIDhcpbasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPIDhcpDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIDhcpConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPIDhcpExists("ibm_pi_dhcp.dhcp_service"),
					resource.TestCheckResourceAttrSet(
						"ibm_pi_dhcp.dhcp_service", "dhcp_id"),
				),
			},
		},
	})
}

func testAccCheckIBMPIDhcpDestroy(s *terraform.State) error {
	sess, err := acc.TestAccProvider.Meta().(conns.ClientSession).IBMPISession()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_pi_dhcp" {
			continue
		}

		cloudInstanceID, dhcpID, err := splitID(rs.Primary.ID)
		if err != nil {
			return err
		}

		client := instance.NewIBMPIDhcpClient(context.Background(), sess, cloudInstanceID)
		_, err = client.Get(dhcpID)
		if err == nil {
			return fmt.Errorf("PI DHCP still exists: %s", rs.Primary.ID)
		}
	}

	return nil
}

func testAccCheckIBMPIDhcpExists(n string) resource.TestCheckFunc {
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

		cloudInstanceID, dhcpID, err := splitID(rs.Primary.ID)
		if err != nil {
			return err
		}
		client := instance.NewIBMPIDhcpClient(context.Background(), sess, cloudInstanceID)

		_, err = client.Get(dhcpID)
		if err != nil {
			return err
		}

		return nil
	}
}

func testAccCheckIBMPIDhcpConfig() string {
	return fmt.Sprintf(`
		resource "ibm_pi_dhcp" "dhcp_service" {
			pi_cloud_instance_id = "%s"
		}`, acc.Pi_cloud_instance_id)
}

func TestAccIBMPIDhcpWithCidrName(t *testing.T) {
	name := fmt.Sprintf("tf-dhcp-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPIDhcpDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIDhcpWithCidrNameConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPIDhcpExists("ibm_pi_dhcp.dhcp_service"),
					resource.TestCheckResourceAttrSet(
						"ibm_pi_dhcp.dhcp_service", "dhcp_id"),
					resource.TestCheckResourceAttrSet(
						"ibm_pi_dhcp.dhcp_service", "status"),
					resource.TestCheckResourceAttrSet(
						"ibm_pi_dhcp.dhcp_service", "network_id"),
					resource.TestCheckResourceAttrSet(
						"ibm_pi_dhcp.dhcp_service", "network_name"),
				),
			},
		},
	})
}

func testAccCheckIBMPIDhcpWithCidrNameConfig(name string) string {
	return fmt.Sprintf(`
		resource "ibm_pi_dhcp" "dhcp_service" {
			pi_cloud_instance_id 	= "%[1]s"
			pi_dhcp_name = "%[2]s"
			pi_cidr = "192.168.103.0/24"
		}`, acc.Pi_cloud_instance_id, name)
}

func TestAccIBMPIDhcpSNAT(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPIDhcpDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIDhcpConfigWithSNATDisabled(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPIDhcpExists("ibm_pi_dhcp.dhcp_service"),
					resource.TestCheckResourceAttrSet(
						"ibm_pi_dhcp.dhcp_service", "dhcp_id"),
				),
			},
		},
	})
}

func testAccCheckIBMPIDhcpConfigWithSNATDisabled() string {
	return fmt.Sprintf(`
		resource "ibm_pi_dhcp" "dhcp_service" {
			pi_cloud_instance_id = "%s"
			pi_dhcp_snat_enabled = false
		}`, acc.Pi_cloud_instance_id)
}
