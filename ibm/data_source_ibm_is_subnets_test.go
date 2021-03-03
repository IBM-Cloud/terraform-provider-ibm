// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMISSubnetsDataSource_basic(t *testing.T) {
	var subnet string
	resName := "data.ibm_is_subnets.test1"
	vpcname := fmt.Sprintf("tfsubnet-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tfsubnet-name-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISSubnetConfig(vpcname, name, ISZoneName, ISCIDR),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISSubnetExists("ibm_is_subnet.testacc_subnet", subnet),
					resource.TestCheckResourceAttr(
						"ibm_is_subnet.testacc_subnet", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_subnet.testacc_subnet", "zone", ISZoneName),
					resource.TestCheckResourceAttr(
						"ibm_is_subnet.testacc_subnet", "ipv4_cidr_block", ISCIDR),
				),
			},
			{
				Config: testAccCheckIBMISSubnetsDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resName, "subnets.0.name"),
					resource.TestCheckResourceAttrSet(resName, "subnets.0.status"),
					resource.TestCheckResourceAttrSet(resName, "subnets.0.status"),
					resource.TestCheckResourceAttrSet(resName, "subnets.0.zone"),
					resource.TestCheckResourceAttrSet(resName, "subnets.0.crn"),
					resource.TestCheckResourceAttrSet(resName, "subnets.0.ipv4_cidr_block"),
					resource.TestCheckResourceAttrSet(resName, "subnets.0.available_ipv4_address_count"),
					resource.TestCheckResourceAttrSet(resName, "subnets.0.network_acl"),
					resource.TestCheckResourceAttrSet(resName, "subnets.0.total_ipv4_address_count"),
					resource.TestCheckResourceAttrSet(resName, "subnets.0.vpc"),
				),
			},
		},
	})
}

func testAccCheckIBMISSubnetsDataSourceConfig() string {
	// status filter defaults to empty
	return fmt.Sprintf(`
	data "ibm_is_subnets" "test1" {
	}`)
}
