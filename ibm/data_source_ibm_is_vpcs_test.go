// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMISVPCsDatasource_basic(t *testing.T) {
	node := "data.ibm_is_vpcs.test1"
	vpcname := fmt.Sprintf("tf-vpcname-%d", acctest.RandIntRange(100, 200))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
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

func testDSCheckIBMISVPCsConfig(vpcname string) string {
	return fmt.Sprintf(`

	resource "ibm_is_vpc" "test_vpc1" {
  		name = "%s"
	}

	data "ibm_is_vpcs" "test1" {
	}

	`, vpcname)
}
