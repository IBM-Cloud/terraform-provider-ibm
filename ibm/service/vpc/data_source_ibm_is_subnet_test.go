// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMISSubnetDatasource_basic(t *testing.T) {
	vpcname := fmt.Sprintf("tfsubnet-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tfsubnet-name-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISVPCDestroy,
		Steps: []resource.TestStep{
			{
				Config: testDSCheckIBMISSubnetConfig(vpcname, name, acc.ISZoneName, acc.ISCIDR),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.ibm_is_subnet.ds_subnet", "name", name),
					resource.TestCheckResourceAttr(
						"data.ibm_is_subnet.ds_subnet", "vpc_name", vpcname),
				),
			},
		},
	})
}

func testDSCheckIBMISSubnetConfig(vpcname, name, zone, cidr string) string {
	return fmt.Sprintf(`
resource "ibm_is_vpc" "testacc_vpc" {
	name = "%s"
}

resource "ibm_is_subnet" "testacc_subnet" {
	name = "%s"
	vpc = "${ibm_is_vpc.testacc_vpc.id}"
	zone = "%s"
	ipv4_cidr_block = "%s"
}
data "ibm_is_subnet" "ds_subnet" {
	identifier = "${ibm_is_subnet.testacc_subnet.id}"
}`, vpcname, name, zone, cidr)
}
