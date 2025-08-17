// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMISVPCsDatasource_basic(t *testing.T) {
	node := "data.ibm_is_vpcs.test1"
	vpcname := fmt.Sprintf("tf-vpcname-%d", acctest.RandIntRange(100, 200))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testDSCheckIBMISVPCsConfig(vpcname),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(node, "vpcs.#"),
				),
			},
		},
	})
}
func TestAccIBMISVPCsDatasource_basicDefaultAddressPrefixes(t *testing.T) {
	node := "data.ibm_is_vpcs.test1"
	apm := "manual"
	vpcname := fmt.Sprintf("tf-vpcname-%d", acctest.RandIntRange(100, 200))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testDSCheckIBMISVPCsConfig(vpcname),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(node, "vpcs.#"),
					resource.TestCheckResourceAttrSet(
						node, "vpcs.0.default_address_prefixes.%"),
				),
			},
			{
				Config: testDSCheckIBMISVPCsConfig1(vpcname, apm),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(node, "vpcs.#"),
					resource.TestCheckResourceAttr(
						node, "vpcs.0.default_address_prefixes.#", "0"),
				),
			},
		},
	})
}

func testDSCheckIBMISVPCsConfig(vpcname string) string {
	return fmt.Sprintf(`

	resource "ibm_is_vpc" "test_vpc1" {
  		name = "%s"
	}

	data "ibm_is_vpcs" "test1" {
	}

	`, vpcname)
}
func testDSCheckIBMISVPCsConfig1(vpcname, apm string) string {
	return fmt.Sprintf(`

	resource "ibm_is_vpc" "test_vpc1" {
  		name 						= "%s"
		address_prefix_management 	= "%s"
	}

	data "ibm_is_vpcs" "test1" {
		depends_on = [ ibm_is_vpc.test_vpc1 ]
	}

	`, vpcname, apm)
}
