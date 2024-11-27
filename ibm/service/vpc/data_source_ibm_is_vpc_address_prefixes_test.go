// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIbmIsVpcAddressPrefixesDataSourceBasic(t *testing.T) {
	name := fmt.Sprintf("tfvpcuat-%d", acctest.RandIntRange(10, 100))
	prefixName := fmt.Sprintf("tfaddprename-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmIsVpcAddressPrefixDataSourceConfigBasic(name, prefixName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc_address_prefixes.is_vpc_address_prefix", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc_address_prefixes.is_vpc_address_prefix", "vpc"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpc_address_prefixes.is_vpc_address_prefix", "address_prefixes.#"),
				),
			},
		},
	})
}

func testAccCheckIbmIsVpcAddressPrefixDataSourceConfigBasic(name, prefixName string) string {
	return testAccCheckIBMISVPCAddressPrefixConfig(name, prefixName) + fmt.Sprintf(`
		data "ibm_is_vpc_address_prefixes" "is_vpc_address_prefix" {
			vpc = "${ibm_is_vpc.testacc_vpc.id}"
		}
	`)
}
