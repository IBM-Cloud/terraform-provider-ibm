// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package dnsservices_test

import (
	"fmt"
	"strings"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMPrivateDNSPermittedNetwork_Basic(t *testing.T) {
	var resultprivatedns string
	name := fmt.Sprintf("testpdnspn%s.com", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPrivateDNSPermittedNetworkDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPrivateDNSPermittedNetworkBasic(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPrivateDNSPermittedNetworkExists("ibm_dns_permitted_network.test-pdns-permitted-network-nw", resultprivatedns),
					resource.TestCheckResourceAttr("ibm_dns_permitted_network.test-pdns-permitted-network-nw", "type", "vpc"),
				),
			},
		},
	})
}

func TestAccIBMPrivateDNSPermittedNetworkImport(t *testing.T) {
	var resultprivatedns string
	name := fmt.Sprintf("testpdnszone%s.com", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPrivateDNSPermittedNetworkDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPrivateDNSPermittedNetworkBasic(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPrivateDNSPermittedNetworkExists("ibm_dns_permitted_network.test-pdns-permitted-network-nw", resultprivatedns),
					resource.TestCheckResourceAttr("ibm_dns_permitted_network.test-pdns-permitted-network-nw", "type", "vpc"),
				),
			},
			{
				ResourceName:      "ibm_dns_permitted_network.test-pdns-permitted-network-nw",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"vpc_crn"},
			},
		},
	})
}

func testAccCheckIBMPrivateDNSPermittedNetworkBasic(name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "rg" {
		is_default=true
	}
	resource "ibm_is_vpc" "test-pdns-permitted-network" {
		name = "test-pdns-permitted-network"
		resource_group = data.ibm_resource_group.rg.id
	}
	resource "ibm_resource_instance" "test-pdns-permitted-network-instance" {
		name = "test-pdns-permitted-network-instance"
		resource_group_id = data.ibm_resource_group.rg.id
		location = "global"
		service = "dns-svcs"
		plan = "standard-dns"
	}
	resource "ibm_dns_zone" "test-pdns-permitted-network-zone" {
		name = "%s"
		instance_id = ibm_resource_instance.test-pdns-permitted-network-instance.guid
		description = "testdescription"
		label = "testlabel"
	}
	resource "ibm_dns_permitted_network" "test-pdns-permitted-network-nw" {
		instance_id = ibm_resource_instance.test-pdns-permitted-network-instance.guid
		zone_id = ibm_dns_zone.test-pdns-permitted-network-zone.zone_id
		vpc_crn = ibm_is_vpc.test-pdns-permitted-network.resource_crn
		type = "vpc"
	}
	  `, name)
}

func testAccCheckIBMPrivateDNSPermittedNetworkDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_permitted_network" {
			continue
		}

		pdnsClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).PrivateDNSClientSession()
		if err != nil {
			return err
		}

		parts := rs.Primary.ID
		partslist := strings.Split(parts, "/")

		getPermittedNetworkOptions := pdnsClient.NewGetPermittedNetworkOptions(partslist[0], partslist[1], partslist[2])
		_, res, err := pdnsClient.GetPermittedNetwork(getPermittedNetworkOptions)

		if err != nil &&
			res.StatusCode != 403 &&
			!strings.Contains(err.Error(), "The service instance was disabled, any access is not allowed.") {

			return fmt.Errorf("testAccCheckIBMPrivateDNSZoneDestroy: Error checking if instance (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}
	return nil
}

func testAccCheckIBMPrivateDNSPermittedNetworkExists(n string, result string) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		pdnsClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).PrivateDNSClientSession()
		if err != nil {
			return err
		}

		parts := rs.Primary.ID
		partslist := strings.Split(parts, "/")

		getPermittedNetworkOptions := pdnsClient.NewGetPermittedNetworkOptions(partslist[0], partslist[1], partslist[2])
		r, res, err := pdnsClient.GetPermittedNetwork(getPermittedNetworkOptions)

		if err != nil &&
			res.StatusCode != 403 &&
			!strings.Contains(err.Error(), "The service instance was disabled, any access is not allowed.") {
			return fmt.Errorf("testAccCheckIBMPrivateDNSZoneExists: Error checking if instance (%s) has been destroyed: %s", rs.Primary.ID, err)
		}

		result = *r.ID
		return nil
	}
}
