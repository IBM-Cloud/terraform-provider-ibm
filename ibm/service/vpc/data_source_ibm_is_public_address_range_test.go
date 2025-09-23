// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMIsPublicAddressRangeDataSourceBasic(t *testing.T) {
	ipv4AddressCount := "16"
	// ipv4AddressCountUpdate := "8"
	name := fmt.Sprintf("tf-name-par%d", acctest.RandIntRange(10, 100))
	vpcName := fmt.Sprintf("tf-name-vpc%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsPublicAddressRangeDataSourceConfigBasic(vpcName, name, ipv4AddressCount),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_public_address_range.is_public_address_range_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_public_address_range.is_public_address_range_instance", "cidr"),
					resource.TestCheckResourceAttrSet("data.ibm_is_public_address_range.is_public_address_range_instance", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_public_address_range.is_public_address_range_instance", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_is_public_address_range.is_public_address_range_instance", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_public_address_range.is_public_address_range_instance", "ipv4_address_count"),
					resource.TestCheckResourceAttrSet("data.ibm_is_public_address_range.is_public_address_range_instance", "lifecycle_state"),
					resource.TestCheckResourceAttrSet("data.ibm_is_public_address_range.is_public_address_range_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_public_address_range.is_public_address_range_instance", "resource_group.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_public_address_range.is_public_address_range_instance", "resource_type"),
				),
			},
		},
	})
}

func testAccCheckIBMIsPublicAddressRangeDataSourceConfigBasic(vpcName, name, ipv4AddressCount string) string {
	return testAccCheckIBMPublicAddressRangeConfigBasic(vpcName, name, ipv4AddressCount) + fmt.Sprintf(`
	data "ibm_is_public_address_range" "is_public_address_range_instance" {
		identifier = ibm_is_public_address_range.public_address_range_instance.id
	}
`)
}
