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

func TestAccIBMPrivateDNSCustomResolverDataSource_basic(t *testing.T) {
	node := "data.ibm_dns_custom_resolvers.test-cr"
	crname := fmt.Sprintf("tf-pdns-custom-resolver-%d", acctest.RandIntRange(100, 200))
	crdescription := fmt.Sprintf("tf-pdns-custom-resolver-tf-test%d", acctest.RandIntRange(100, 200))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPrivateDNSCustomResolverDataSourceConfig(crname, crdescription),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(node, "custom_resolvers.0.name"),
					resource.TestCheckResourceAttrSet(node, "custom_resolvers.0.description"),
					resource.TestCheckResourceAttrSet(node, "custom_resolvers.0.enabled"),
					resource.TestCheckResourceAttrSet(node, "custom_resolvers.0.locations.0.subnet_crn"),
					resource.TestCheckResourceAttrSet(node, "custom_resolvers.0.locations.0.enabled"),
				),
			},
		},
	})
}

func testAccCheckIBMPrivateDNSCustomResolverDataSourceConfig(crname, crdescription string) string {
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
		enabled 	= true
		locations {
			subnet_crn	= ibm_is_subnet.test-pdns-cr-subnet1.crn
			enabled		= true
		}
		locations {
			subnet_crn	= ibm_is_subnet.test-pdns-cr-subnet2.crn
			enabled     = true
		}
	}
	data "ibm_dns_custom_resolvers" "test-cr" {
		depends_on  = [ibm_dns_custom_resolver.test]
		instance_id	= ibm_dns_custom_resolver.test.instance_id
	}`, crname, crdescription)
}
