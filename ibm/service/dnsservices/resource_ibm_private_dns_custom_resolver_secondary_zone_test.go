// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package dnsservices_test

import (
	"fmt"
	"strings"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMPrivateDNSSecondaryZone_basic(t *testing.T) {
	vpcname := fmt.Sprintf("seczone-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("seczone-subnet-name-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("seczone-cr-%s", acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum))
	description := "new test CR - TF"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() {},
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPrivateDNSSecondaryZoneDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPrivateDNSSecondaryZoneResource(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, name, description),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_dns_custom_resolver_secondary_zone.test", "zone", "seczone-terraform-plugin-test.com"),
				),
			},
		},
	})
}

func testAccCheckIBMPrivateDNSSecondaryZoneDestroy(s *terraform.State) error {
	// conn := testAccProvider.Meta().(*sdk.Client)
	pdnsClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).PrivateDNSClientSession()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_dns_secondary_zone" {
			continue
		}
		if rs.Primary.ID == "" {
			return flex.FmtErrorf("No resource primary ID is set")
		}

		partslist := strings.Split(rs.Primary.ID, "/")
		if len(partslist) < 3 {
			return flex.FmtErrorf("Invalid resource primary ID. Must contain 3 parts.")
		}
		instanceID := partslist[0]
		customResolverID := partslist[1]
		secondaryZoneID := partslist[2]
		getSecondaryZoneOptions := pdnsClient.NewGetSecondaryZoneOptions(
			instanceID,
			customResolverID,
			secondaryZoneID,
		)
		_, _, err := pdnsClient.GetSecondaryZone(getSecondaryZoneOptions)
		if err != nil {
			return flex.FmtErrorf("testAccCheckIBMPrivateDNSZoneDestroy: Error checking if instance (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}
	return nil
}

func testAccCheckIBMPrivateDNSSecondaryZoneResource(vpcname, subnetname, zone, cidr, name, description string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "rg" {
		is_default	= true
	}
	resource "ibm_is_vpc" "test-pdns-cr-vpc" {
		name			= "%s"
		resource_group	= data.ibm_resource_group.rg.id
	}
	resource "ibm_is_subnet" "test-pdns-cr-subnet1" {
		name			= "%s"
		vpc				= ibm_is_vpc.test-pdns-cr-vpc.id
		zone			= "%s"
		ipv4_cidr_block	= "%s"
		resource_group	= data.ibm_resource_group.rg.id
	}
	resource "ibm_resource_instance" "test-pdns-cr-instance" {
		name				= "test-pdns-cr-instance"
		resource_group_id	= data.ibm_resource_group.rg.id
		location			= "global"
		service				= "dns-svcs"
		plan				= "standard-dns"
	}
	resource "ibm_dns_custom_resolver" "test" {
		name		= "%s"
		instance_id = ibm_resource_instance.test-pdns-cr-instance.guid
		description = "%s"
		high_availability = false
		enabled 	= true
		locations {
			subnet_crn	= ibm_is_subnet.test-pdns-cr-subnet1.crn
			enabled		= true
		}
	}

	resource "ibm_dns_zone" "pdns-1-zone" {
		name        = "seczone-terraform-plugin-test.com"
		instance_id = ibm_resource_instance.test-pdns-cr-instance.guid
		description = "testdescription"
		label       = "testlabel"
	}
	
	resource "ibm_dns_custom_resolver_secondary_zone" "test" {
		instance_id   = ibm_resource_instance.test-pdns-cr-instance.guid
		resolver_id   = ibm_dns_custom_resolver.test.custom_resolver_id
		description   = "seczone terraform plugin test"
		zone          = "seczone-terraform-plugin-test.com"
		enabled       = false
		transfer_from = ["10.0.0.8"]
	}
	`, vpcname, subnetname, zone, cidr, name, description)
}
