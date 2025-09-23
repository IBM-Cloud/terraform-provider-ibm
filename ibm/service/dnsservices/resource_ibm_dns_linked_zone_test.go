// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package dnsservices_test

import (
	"fmt"
	"strings"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMDNSLinkedZone_basic(t *testing.T) {
	name := fmt.Sprintf("seczone-cr-%s", acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum))
	ownerInstanceId := "OWNER Instance ID"
	ownerZoneId := "OWNER ZONE ID"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() {},
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMDNSLinkedZoneDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMDNSLinkedZoneBasic(name, ownerInstanceId, ownerZoneId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_dns_linked_zone.test", "name", name),
					resource.TestCheckResourceAttr("ibm_dns_linked_zone.test", "owner_instance_id", ownerInstanceId),
					resource.TestCheckResourceAttr("ibm_dns_linked_zone.test", "owner_zone_id", ownerZoneId),
				),
			},
		},
	})
}

func TestAccIBMDNSLinkedZone_update(t *testing.T) {
	name := fmt.Sprintf("seczone-cr-%s", acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum))
	ownerInstanceId := "OWNER Instance ID"
	ownerZoneId := "OWNER ZONE ID"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() {},
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMDNSLinkedZoneDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMDNSLinkedZoneBasic(name, ownerInstanceId, ownerZoneId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_dns_linked_zone.test", "name", name),
				),
			},
			{
				Config: testAccCheckIBMDNSLinkedZoneUpdateConfig(name, "test description", "label"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_dns_linked_zone.test", "description", "test description"),
					resource.TestCheckResourceAttr("ibm_dns_linked_zone.test", "label", "label"),
				),
			},
		},
	})
}

func testAccCheckIBMDNSLinkedZoneUpdateConfig(name string, description string, label string) string {
	return fmt.Sprintf(`
	resource "ibm_resource_group" "rg" {
		name		= "DNS_linked_zone_test"
	}
	resource "ibm_resource_instance" "test-pdns-cr-instance" {
		name				= "test-pdns-cr-instance"
		resource_group_id	= ibm_resource_group.rg.id
		service				= "dns-svcs"
		plan				= "standard"
	}
	resource "ibm_dns_linked_zone" "test" {
		name = "%s"
		instance_id	= ibm_resource_instance.test-pdns-cr-instance.id
		owner_instance_id = "OWNER Instance ID"
                owner_zone_id = "OWNER ZONE ID"
		description	= "%s"
		label		= "%s"
	}
	`, name, description, label)
}

func TestAccIBMDNSLinkedZoneImport(t *testing.T) {
	var linkedZoneID string
	name := fmt.Sprintf("seczone-vpc-%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	ownerInstanceId := "OWNER Instance ID"
	ownerZoneId := "OWNER ZONE ID"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMDNSLinkedZoneDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMDNSLinkedZoneBasic(name, ownerInstanceId, ownerZoneId),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMDNSLinkedZoneExists("ibm_dns_linked_zone.test", &linkedZoneID),
					resource.TestCheckResourceAttr("ibm_dns_linked_zone.test", "name", name),
				),
			},
			{
				ResourceName:      "ibm_dns_linked_zone.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMDNSLinkedZoneDestroy(s *terraform.State) error {
	// conn := testAccProvider.Meta().(*sdk.Client)
	pdnsClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).PrivateDNSClientSession()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_dns_linked_zone" {
			continue
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No resource primary ID is set")
		}

		partslist := strings.Split(rs.Primary.ID, "/")
		if len(partslist) < 2 {
			return fmt.Errorf("Invalid resource primary ID. Must contain 2 parts.")
		}
		instanceID := partslist[0]
		linkedDnsZoneID := partslist[1]
		getLinkedZoneOptions := pdnsClient.NewGetLinkedZoneOptions(
			instanceID,
			linkedDnsZoneID,
		)
		_, _, err := pdnsClient.GetLinkedZone(getLinkedZoneOptions)
		if err != nil {
			return fmt.Errorf("testAccCheckIBMDNSLinkedZoneDestroy: Error checking if instance (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}
	return nil
}

func testAccCheckIBMDNSLinkedZoneBasic(name string, ownerInstanceId string, ownerZoneId string) string {
	return fmt.Sprintf(`
	resource "ibm_resource_group" "rg" {
		name		= "DNS_linked_zone_test"
	}
	resource "ibm_resource_instance" "test-pdns-cr-instance" {
		name				= "test-pdns-cr-instance"
		resource_group_id	= ibm_resource_group.rg.id
		location			= "global"
		service				= "dns-svcs"
		plan				= "standard-dns"
	}
	resource "ibm_dns_linked_zone" "test" {
		name		= "%s"
		instance_id = ibm_resource_instance.test-pdns-cr-instance.id
                owner_instance_id = "%s"
                owner_zone_id = "%s"
	}
	`, name, ownerInstanceId, ownerZoneId)
}

//func testAccCheckIBMDNSLinkedZoneResource(vpcname, subnetname, zone, cidr, name, description string) string {
//	return fmt.Sprintf(`
//	data "ibm_resource_group" "rg" {
//		is_default	= true
//	}
//	resource "ibm_resource_instance" "test-pdns-cr-instance" {
//		name				= "test-pdns-cr-instance"
//		resource_group_id	= data.ibm_resource_group.rg.id
//		location			= "global"
//		service				= "dns-svcs"
//		plan				= "standard-dns"
//	}
//	resource "ibm_dns_zone" "pdns-1-zone" {
//		name        = "seczone-terraform-plugin-test.com"
//		instance_id = ibm_resource_instance.test-pdns-cr-instance.guid
//		description = "testdescription"
//		label       = "testlabel"
//	}
//
//	resource "ibm_dns_linked_zone" "test" {
//		instance_id   = ibm_resource_instance.test-pdns-cr-instance.guid
//		description   = "seczone terraform plugin test"
//	}
//	`, name)
//}

func testAccCheckIBMDNSLinkedZoneExists(n string, linkedZoneID *string) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Linked Zone not found: %s", n)
		}
		pdnsClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).PrivateDNSClientSession()
		if err != nil {
			return err
		}
		parts := rs.Primary.ID
		partslist := strings.Split(parts, "/")

		getLinkedZoneOptions := pdnsClient.NewGetLinkedZoneOptions(
			partslist[0],
			partslist[1],
		)
		linkedZone, _, err := pdnsClient.GetLinkedZone(getLinkedZoneOptions)
		if err != nil {
			return fmt.Errorf("Error Fetching Linked Zone: %s", err)
		}
		*linkedZoneID = *linkedZone.ID
		return nil
	}
}
