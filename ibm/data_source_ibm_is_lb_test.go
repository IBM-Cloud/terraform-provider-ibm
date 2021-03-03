// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMISLBDatasource_basic(t *testing.T) {
	name := fmt.Sprintf("tflb-name-%d", acctest.RandIntRange(10, 100))
	vpcname := fmt.Sprintf("tflb-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tflb-subnet-name-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testDSCheckIBMISLBConfig(vpcname, subnetname, ISZoneName, ISCIDR, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.ibm_is_lb.ds_lb", "name", name),
				),
			},
		},
	})
}

func testDSCheckIBMISLBConfig(vpcname, subnetname, zone, cidr, name string) string {
	return fmt.Sprintf(`
resource "ibm_is_vpc" "testacc_vpc" {
  name = "%s"
}

resource "ibm_is_subnet" "testacc_subnet" {
  name            = "%s"
  vpc             = ibm_is_vpc.testacc_vpc.id
  zone            = "%s"
  ipv4_cidr_block = "%s"
}
resource "ibm_is_lb" "testacc_lb" {
  name    = "%s"
  subnets = [ibm_is_subnet.testacc_subnet.id]
}
data "ibm_is_lb" "ds_lb" {
  name = ibm_is_lb.testacc_lb.name
}`, vpcname, subnetname, zone, cidr, name)
}
