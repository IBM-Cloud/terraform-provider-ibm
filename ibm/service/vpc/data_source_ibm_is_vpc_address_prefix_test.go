// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMIsVPCAddressPrefixDataSourceBasic(t *testing.T) {
	name := fmt.Sprintf("tfvpcuat-%d", acctest.RandIntRange(10, 100))
	prefixName := fmt.Sprintf("tfaddprename-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIsVPCAddressPrefixDataSourceConfigBasic(name, prefixName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc_address_prefix.is_vpc_address_prefix", "cidr"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc_address_prefix.is_vpc_address_prefix", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc_address_prefix.is_vpc_address_prefix", "has_subnets"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc_address_prefix.is_vpc_address_prefix", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc_address_prefix.is_vpc_address_prefix", "is_default"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc_address_prefix.is_vpc_address_prefix", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc_address_prefix.is_vpc_address_prefix", "zone.#"),
				),
			},
			{
				Config: testAccCheckIBMIsVPCAddressPrefixDataSourceConfigBasic(name, prefixName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc_address_prefix.is_vpc_address_prefix1", "cidr"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc_address_prefix.is_vpc_address_prefix1", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc_address_prefix.is_vpc_address_prefix1", "has_subnets"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc_address_prefix.is_vpc_address_prefix1", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc_address_prefix.is_vpc_address_prefix1", "is_default"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc_address_prefix.is_vpc_address_prefix1", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc_address_prefix.is_vpc_address_prefix1", "zone.#"),
				),
			},
			{
				Config: testAccCheckIBMIsVPCAddressPrefixDataSourceConfigBasic(name, prefixName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc_address_prefix.is_vpc_address_prefix2", "cidr"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc_address_prefix.is_vpc_address_prefix2", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc_address_prefix.is_vpc_address_prefix2", "has_subnets"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc_address_prefix.is_vpc_address_prefix2", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc_address_prefix.is_vpc_address_prefix2", "is_default"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc_address_prefix.is_vpc_address_prefix2", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc_address_prefix.is_vpc_address_prefix2", "zone.#"),
				),
			},
			{
				Config: testAccCheckIBMIsVPCAddressPrefixDataSourceConfigBasic(name, prefixName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc_address_prefix.is_vpc_address_prefix3", "cidr"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc_address_prefix.is_vpc_address_prefix3", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc_address_prefix.is_vpc_address_prefix3", "has_subnets"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc_address_prefix.is_vpc_address_prefix3", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc_address_prefix.is_vpc_address_prefix3", "is_default"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc_address_prefix.is_vpc_address_prefix3", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc_address_prefix.is_vpc_address_prefix3", "zone.#"),
				),
			},
		},
	})
}

func testAccCheckIBMIsVPCAddressPrefixDataSourceConfigBasic(name, prefixName string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
		address_prefix_management = "manual"
	}
	resource "ibm_is_vpc_address_prefix" "testacc_vpc_address_prefix" {
		name = "%s"
		zone = "%s"
		vpc = "${ibm_is_vpc.testacc_vpc.id}"
		cidr = "%s"
		is_default = true
	}
	data "ibm_is_vpc_address_prefix" "is_vpc_address_prefix" {
		vpc = ibm_is_vpc.testacc_vpc.id
		address_prefix = ibm_is_vpc_address_prefix.testacc_vpc_address_prefix.address_prefix
	}
	data "ibm_is_vpc_address_prefix" "is_vpc_address_prefix1" {
		vpc_name = ibm_is_vpc.testacc_vpc.name
		address_prefix = ibm_is_vpc_address_prefix.testacc_vpc_address_prefix.address_prefix
	}
	data "ibm_is_vpc_address_prefix" "is_vpc_address_prefix2" {
		vpc = ibm_is_vpc.testacc_vpc.id
		address_prefix_name = ibm_is_vpc_address_prefix.testacc_vpc_address_prefix.name
	}
	data "ibm_is_vpc_address_prefix" "is_vpc_address_prefix3" {
		vpc_name = ibm_is_vpc.testacc_vpc.name
		address_prefix_name = ibm_is_vpc_address_prefix.testacc_vpc_address_prefix.name
	}
	`, name, prefixName, acc.ISZoneName, acc.ISAddressPrefixCIDR)
}
