// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0
package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMISSubnetReservedIP_basic(t *testing.T) {
	vpcName := fmt.Sprintf("tfresip-vpc-%d", acctest.RandIntRange(10, 100))
	subnetName := fmt.Sprintf("tfresip-subnet-%d", acctest.RandIntRange(10, 100))
	resIPName := fmt.Sprintf("tfresip-reservedip-%d", acctest.RandIntRange(10, 100))
	terraformTag := "data.ibm_is_subnet_reserved_ip.data_resip1"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccIBMISReservedIPdataSoruceConfig(vpcName, subnetName, resIPName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(terraformTag, isReservedIPName, resIPName),
				),
			},
		},
	})
}

func testAccIBMISReservedIPdataSoruceConfig(vpcName, subnetName, reservedIPName string) string {
	return fmt.Sprintf(`
		resource "ibm_is_vpc" "vpc1" {
			name = "%s"
		}

		resource "ibm_is_subnet" "subnet1" {
			name                     = "%s"
			vpc                      = ibm_is_vpc.vpc1.id
			zone                     = "us-south-1"
			total_ipv4_address_count = 256
		}

		resource "ibm_is_subnet_reserved_ip" "resip1" {
			subnet = ibm_is_subnet.subnet1.id
			name = "%s"
		}

		data "ibm_is_subnet_reserved_ip" "data_resip1" {
			subnet = ibm_is_subnet.subnet1.id
			reserved_ip = ibm_is_subnet_reserved_ip.resip1.reserved_ip
		}
      `, vpcName, subnetName, reservedIPName)
}
