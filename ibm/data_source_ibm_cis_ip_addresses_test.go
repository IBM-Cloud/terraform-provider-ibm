// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccIBMCisIPDataSource_Basic(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testAccCheckIBMCisIPDataSourceConfigBasic),
				Check: resource.ComposeTestCheckFunc(
					testAccIBMCisIPAddrs("data.ibm_cis_ip_addresses.test_acc"),
				),
			},
		},
	})
}

func testAccIBMCisIPAddrs(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		r := s.RootModule().Resources[n]
		a := r.Primary.Attributes

		cidrs, _ := strconv.Atoi(a["ipv4_cidrs.#"])
		if cidrs == 0 {
			return fmt.Errorf("No ipv4 cidrs returned")
		}
		cidrs, _ = strconv.Atoi(a["ipv6_cidrs.#"])
		if cidrs == 0 {
			return fmt.Errorf("No ipv6 cidrs returned")
		}
		return nil
	}
}

const testAccCheckIBMCisIPDataSourceConfigBasic = `
data "ibm_cis_ip_addresses" "test_acc" {}
`
