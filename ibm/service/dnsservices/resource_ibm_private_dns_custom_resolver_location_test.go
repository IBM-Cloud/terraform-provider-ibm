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
	name := fmt.Sprintf("testpdnscustomresolver%s", acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum))
	description := "new test CR Locations - TF"
	subnet_crn := "crn:v1:bluemix:public:is:us-south-3:a/bcf1865e99742d38d2d5fc3fb80a5496::subnet:0737-0d198509-3221-4162-b2d8-4a9326d3d7ad"
	subnet_crn_new := "crn:v1:bluemix:public:is:us-south-2:a/bcf1865e99742d38d2d5fc3fb80a5496::subnet:0727-f17967f2-2bbe-427c-bcf6-22f8c2395285"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPrivateDNSCRLocationsBasic(name, description, subnet_crn, subnet_crn_new),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_dns_custom_resolver_location.test1", "subnet_crn", subnet_crn),
					resource.TestCheckResourceAttr("ibm_dns_custom_resolver_location.test1", "enabled", "true"),
					resource.TestCheckResourceAttr("ibm_dns_custom_resolver_location.test2", "enabled", "false"),
					resource.TestCheckResourceAttr("ibm_dns_custom_resolver_location.test1", "cr_enabled", "false"),
					resource.TestCheckResourceAttr("ibm_dns_custom_resolver_location.test2", "cr_enabled", "false"),
					resource.TestCheckResourceAttr("ibm_dns_custom_resolver_location.test2", "subnet_crn", subnet_crn_new),
				),
			},
		},
	})
}

func TestAccIBMPrivateDNSCustomResolverLocations_Import(t *testing.T) {

	name := fmt.Sprintf("testpdnscustomresolver%s", acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum))
	description := "new test CR Locations - TF"
	subnet_crn := "crn:v1:bluemix:public:is:us-south-3:a/bcf1865e99742d38d2d5fc3fb80a5496::subnet:0737-0d198509-3221-4162-b2d8-4a9326d3d7ad"
	subnet_crn_new := "crn:v1:bluemix:public:is:us-south-2:a/bcf1865e99742d38d2d5fc3fb80a5496::subnet:0727-f17967f2-2bbe-427c-bcf6-22f8c2395285"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPrivateDNSCRLocationsBasic(name, description, subnet_crn, subnet_crn_new),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_dns_custom_resolver_location.test1", "subnet_crn", subnet_crn),
					resource.TestCheckResourceAttr("ibm_dns_custom_resolver_location.test1", "enabled", "true"),
					resource.TestCheckResourceAttr("ibm_dns_custom_resolver_location.test2", "enabled", "false"),
					resource.TestCheckResourceAttr("ibm_dns_custom_resolver_location.test1", "cr_enabled", "false"),
					resource.TestCheckResourceAttr("ibm_dns_custom_resolver_location.test2", "cr_enabled", "false"),
					resource.TestCheckResourceAttr("ibm_dns_custom_resolver_location.test2", "subnet_crn", subnet_crn_new),
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

func testAccCheckIBMPrivateDNSCRLocationsBasic(name, description, subnet_crn, subnet_crn_new string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "rg" {
		is_default	= true
	}
	resource "ibm_is_vpc" "test-pdns-cr-vpc" {
		name			= "test-pdns-custom-resolver-vpc"
		resource_group	= data.ibm_resource_group.rg.id
	}
	resource "ibm_is_subnet" "test-pdns-cr-subnet1" {
		name			= "test-pdns-cr-subnet1"
		vpc				= ibm_is_vpc.test-pdns-cr-vpc.id
		zone			= "us-south-1"
		ipv4_cidr_block	= "10.240.0.0/24"
		resource_group	= data.ibm_resource_group.rg.id
	}
	resource "ibm_is_subnet" "test-pdns-cr-subnet2" {
		name			= "test-pdns-cr-subnet2"
		vpc				= ibm_is_vpc.test-pdns-cr-vpc.id
		zone			= "us-south-1"
		ipv4_cidr_block	= "10.240.64.0/24"
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
		enabled = false
	}
	resource "ibm_dns_custom_resolver_location" "test1" {
		instance_id = ibm_resource_instance.test-pdns-cr-instance.guid
		resolver_id = ibm_dns_custom_resolver.test.custom_resolver_id
		subnet_crn = ibm_is_subnet.test-pdns-cr-subnet1.crn
		enabled    = true
		cr_enabled = false
	}
	resource "ibm_dns_custom_resolver_location" "test2" {
		instance_id   = ibm_resource_instance.test-pdns-cr-instance.guid
		resolver_id   = ibm_dns_custom_resolver.test.custom_resolver_id
		subnet_crn    = ibm_is_subnet.test-pdns-cr-subnet2.crn
		enabled       = false
		cr_enabled    = false 
	}
	`, name, description)
}
