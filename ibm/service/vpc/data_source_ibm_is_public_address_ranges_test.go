// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMIsPublicAddressRangesDataSourceBasic(t *testing.T) {
	ipv4AddressCount := "16"
	name := fmt.Sprintf("tf-name-par%d", acctest.RandIntRange(10, 100))
	vpcName := fmt.Sprintf("tf-name-vpc%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsPublicAddressRangesDataSourceConfigBasic(vpcName, name, ipv4AddressCount),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_public_address_ranges.is_public_address_ranges_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_public_address_ranges.is_public_address_ranges_instance", "public_address_ranges.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_public_address_ranges.is_public_address_ranges_instance", "public_address_ranges.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_public_address_ranges.is_public_address_ranges_instance", "public_address_ranges.0.cidr"),
					resource.TestCheckResourceAttrSet("data.ibm_is_public_address_ranges.is_public_address_ranges_instance", "public_address_ranges.0.created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_public_address_ranges.is_public_address_ranges_instance", "public_address_ranges.0.crn"),
					resource.TestCheckResourceAttrSet("data.ibm_is_public_address_ranges.is_public_address_ranges_instance", "public_address_ranges.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_public_address_ranges.is_public_address_ranges_instance", "public_address_ranges.0.ipv4_address_count"),
					resource.TestCheckResourceAttrSet("data.ibm_is_public_address_ranges.is_public_address_ranges_instance", "public_address_ranges.0.lifecycle_state"),
					resource.TestCheckResourceAttrSet("data.ibm_is_public_address_ranges.is_public_address_ranges_instance", "public_address_ranges.0.name"),
				),
			},
		},
	})
}

func testAccCheckIBMIsPublicAddressRangesDataSourceConfigBasic(vpcName, name, ipv4AddressCount string) string {
	return testAccCheckIBMPublicAddressRangeConfigBasic(vpcName, name, ipv4AddressCount) + fmt.Sprintf(`
	data "ibm_is_public_address_ranges" "is_public_address_ranges_instance" {
	  depends_on = [ibm_is_public_address_range.public_address_range_instance]
	}
`)
}
