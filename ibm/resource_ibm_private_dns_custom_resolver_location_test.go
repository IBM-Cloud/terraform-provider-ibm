// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMPrivateDNSCustomResolverLocations_Basic(t *testing.T) {
	name := fmt.Sprintf("testpdnscustomresolver%s", acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum))
	description := "new test CR Locations - TF"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPrivateDNSCRLocationsBasic(name, description),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_dns_custom_resolver_location.test", "enabled", "true"),
				),
			},
		},
	})
}

func TestAccIBMPrivateDNSCustomResolverLocations_Import(t *testing.T) {
	name := fmt.Sprintf("testpdnscustomresolver%s", acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum))
	description := "new test CR Locations - TF"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPrivateDNSCRLocationsBasic(name, description),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_dns_custom_resolver_location.test", "enabled", "true"),
				),
			},
			{
				ResourceName:      "ibm_dns_custom_resolver_location.test",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"enabled",
					"instance_id",
					"resolver_id",
					"subnet_crn",
				},
			},
		},
	})
}

func testAccCheckIBMPrivateDNSCRLocationsBasic(name, description string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "rg" {
		is_default	= true
	}
	resource "ibm_is_vpc" "test-pdns-cr-vpc" {
		name			= "test-pdns-custom-resolver-locations-vpc"
		resource_group	= data.ibm_resource_group.rg.id
	}
	resource "ibm_is_subnet" "test-pdns-cr-subnet1" {
		name                    = "test-pdns-cr-subnet1"
		vpc                     = ibm_is_vpc.test-pdns-cr-vpc.id
		zone            		= "us-south-1"
		ipv4_cidr_block 		= "10.240.0.0/24"
		resource_group 			= data.ibm_resource_group.rg.id
	}
	resource "ibm_is_subnet" "test-pdns-cr-subnet2" {
		name                    = "test-pdns-cr-subnet2"
		vpc                     = ibm_is_vpc.test-pdns-cr-vpc.id
		zone            		= "us-south-1"
		ipv4_cidr_block 		= "10.240.64.0/24"
		resource_group 			= data.ibm_resource_group.rg.id
	}
	resource "ibm_resource_instance" "test-pdns-cr-instance" {
		name				= "test-pdns-cr-location-instance"
		resource_group_id	= data.ibm_resource_group.rg.id
		location			= "global"
		service				= "dns-svcs"
		plan				= "standard-dns"
	}
	resource "ibm_dns_custom_resolver" "test" {
		name			= "%s"
		instance_id		= ibm_resource_instance.test-pdns-cr-instance.guid
		description 	= "%s"
		high_availability = false
		locations	{
			subnet_crn = ibm_is_subnet.test-pdns-cr-subnet1.crn
			enabled     = true
		}
	}
	resource "ibm_dns_custom_resolver_location" "test" {
		instance_id = ibm_resource_instance.test-pdns-cr-instance.guid
		resolver_id = ibm_dns_custom_resolver.test.custom_resolver_id
		subnet_crn  = ibm_is_subnet.test-pdns-cr-subnet2.crn
		enabled     = true
	}
	  	`, name, description)
}
