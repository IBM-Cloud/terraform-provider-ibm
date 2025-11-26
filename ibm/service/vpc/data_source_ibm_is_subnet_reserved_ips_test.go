// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIBMISSubnetReservedIPs_basic(t *testing.T) {
	terraformTagData := "data.ibm_is_subnet_reserved_ips.data_resips"
	terraformTagRes1 := "ibm_is_subnet_reserved_ip.resIP1"
	terraformTagRes2 := "ibm_is_subnet_reserved_ip.resIP2"
	var resIP1 string
	var resIP2 string
	vpcName := fmt.Sprintf("tfresip-vpc-%d", acctest.RandIntRange(10, 100))
	subnetName := fmt.Sprintf("tfresip-subnet-%d", acctest.RandIntRange(10, 100))
	reservedIPName := fmt.Sprintf("tfresip-reservedip-%d", acctest.RandIntRange(10, 100))
	reservedIPName2 := fmt.Sprintf("tfresip-reservedip-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccIBMISReservedIPSResoruceConfig2(vpcName, subnetName,
					reservedIPName, reservedIPName2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckISSubnetReservedIPExists(terraformTagRes1, &resIP1),
					resource.TestCheckResourceAttr(
						terraformTagRes1, "name", reservedIPName),
					testAccCheckISSubnetReservedIPExists(terraformTagRes2, &resIP2),
					resource.TestCheckResourceAttr(
						terraformTagRes2, "name", reservedIPName2),
					resource.TestCheckResourceAttrSet(terraformTagData, "reserved_ips.0.address"),
					resource.TestCheckResourceAttrSet(terraformTagData, "reserved_ips.0.auto_delete"),
					resource.TestCheckResourceAttrSet(terraformTagData, "reserved_ips.0.created_at"),
					resource.TestCheckResourceAttrSet(terraformTagData, "reserved_ips.0.href"),
					resource.TestCheckResourceAttrSet(terraformTagData, "reserved_ips.0.reserved_ip"),
					resource.TestCheckResourceAttrSet(terraformTagData, "reserved_ips.0.name"),
					resource.TestCheckResourceAttrSet(terraformTagData, "reserved_ips.0.owner"),
					resource.TestCheckResourceAttrSet(terraformTagData, "reserved_ips.0.resource_type"),

					resource.TestCheckResourceAttrSet(terraformTagData, "reserved_ips.1.address"),
					resource.TestCheckResourceAttrSet(terraformTagData, "reserved_ips.1.auto_delete"),
					resource.TestCheckResourceAttrSet(terraformTagData, "reserved_ips.1.created_at"),
					resource.TestCheckResourceAttrSet(terraformTagData, "reserved_ips.1.href"),
					resource.TestCheckResourceAttrSet(terraformTagData, "reserved_ips.1.reserved_ip"),
					resource.TestCheckResourceAttrSet(terraformTagData, "reserved_ips.1.name"),
					resource.TestCheckResourceAttrSet(terraformTagData, "reserved_ips.1.owner"),
					resource.TestCheckResourceAttrSet(terraformTagData, "reserved_ips.1.resource_type"),

					resource.TestCheckResourceAttr(terraformTagData, "total_count", "2"),
				),
			},
		},
	})
}

func testAccIBMISReservedIPSResoruceConfig2(vpcName, subnetName, reservedIPName, reservedIPName2 string) string {
	return testAccCheckISSubnetReservedIPConfigBasic(vpcName, subnetName, reservedIPName) + fmt.Sprintf(`
	
		resource "ibm_is_subnet_reserved_ip" "resIP2" {
			subnet = ibm_is_subnet.subnet1.id
			name = "%s"
		}
		data "ibm_is_subnet_reserved_ips" "data_resips" {
			subnet = ibm_is_subnet_reserved_ip.resIP2.subnet
		}
      `, reservedIPName2)
}
