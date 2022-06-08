// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package dnsservices_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMPrivateDNSCustomResolverLocations_basic(t *testing.T) {

	vpcname := fmt.Sprintf("cr-loc-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("cr-loc-subnet-name-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("testpdnscustomresolver%s", acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum))
	description := "new test CR Locations - TF"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPrivateDNSCRLocationsBasic(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, name, description),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_dns_custom_resolver.test", "enabled", "false"),
				),
			},
		},
	})
}

func TestAccIBMPrivateDNSCustomResolverLocations_Import(t *testing.T) {

	name := fmt.Sprintf("testpdnscustomresolver%s", acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum))
	description := "new test CR Locations - TF"
	vpcname := fmt.Sprintf("cr-loc-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("cr-loc-subnet-name-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPrivateDNSCRLocationsBasic(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, name, description),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_dns_custom_resolver.test", "enabled", "false"),
				),
			},
			{
				ResourceName:      "ibm_dns_custom_resolver.test",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"enabled",
					"instance_id",
				},
			},
		},
	})
}

func testAccCheckIBMPrivateDNSCRLocationsBasic(vpcname, subnetname, zone, cidr, name, description string) string {
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
		enabled 	= false
		locations {
			subnet_crn = ibm_is_subnet.test-pdns-cr-subnet1.crn
			enabled    = false
		}
		locations {
			subnet_crn = ibm_is_subnet.test-pdns-cr-subnet1.crn
			enabled    = false
		}
	}
	`, vpcname, subnetname, zone, cidr, name, description)

}
