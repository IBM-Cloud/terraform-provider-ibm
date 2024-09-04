// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMISSubnetReservedIPPatchResource_basic(t *testing.T) {
	var reservedIPID string
	vpcName := fmt.Sprintf("tfresip-vpc-%d", acctest.RandIntRange(10, 100))
	subnetName := fmt.Sprintf("tfresip-subnet-%d", acctest.RandIntRange(10, 100))
	reservedIPName := fmt.Sprintf("tfresip-reservedip-%d", acctest.RandIntRange(10, 100))
	reservedIPName2 := fmt.Sprintf("tfresip-reservedip-%d", acctest.RandIntRange(10, 100))
	terraformTag := "ibm_is_subnet_reserved_ip_patch.resIP1"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				// Tests create
				Config: testAccCheckISSubnetReservedIPPatchConfigBasic(vpcName, subnetName, reservedIPName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckISSubnetReservedIPPatchExists(terraformTag, &reservedIPID),
					resource.TestCheckResourceAttrSet(terraformTag, "name"),
					resource.TestCheckResourceAttr(terraformTag, "name", reservedIPName),
				),
			},
			{
				// Tests Update
				Config: testAccCheckISSubnetReservedIPPatchConfigBasic(vpcName, subnetName, reservedIPName2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckISSubnetReservedIPExists(terraformTag, &reservedIPID),
					resource.TestCheckResourceAttr(terraformTag, "name", reservedIPName2),
				),
			},
		},
	})
}

func testAccCheckISSubnetReservedIPPatchExists(resIPName string, reservedIPID *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resIPName]
		if !ok {
			return fmt.Errorf("Not Found (reserved IP patch): %s", resIPName)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("[ERROR] No reserved IP ID is set")
		}

		sess, err := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
		if err != nil {
			return err
		}
		parts, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}
		opt := sess.NewGetSubnetReservedIPOptions(parts[0], parts[1])
		result, response, err := sess.GetSubnetReservedIP(opt)
		if err != nil {
			return fmt.Errorf("Reserved IP does not exist: %s", response)
		}
		*reservedIPID = *result.ID
		return nil
	}
}

func testAccCheckISSubnetReservedIPPatchConfigBasic(vpcName, subnetName, resIPName string) string {
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

	  resource "ibm_is_virtual_endpoint_gateway" "endpoint_gateway" {
		name = "my-endpoint-gateway-1"
		target {
		  name          = "ibm-ntp-server"
		  resource_type = "provider_infrastructure_service"
		}
		vpc = ibm_is_vpc.vpc1.id
	  }

	  resource "ibm_is_subnet_reserved_ip" "resIP1" {
		subnet = ibm_is_subnet.subnet1.id
		target = ibm_is_virtual_endpoint_gateway.endpoint_gateway.id
	  }

	resource "ibm_is_subnet_reserved_ip_patch" "resIP1" {
		subnet = ibm_is_subnet.subnet1.id
		name = "%s"
		reserved_ip = ibm_is_subnet_reserved_ip.resIP1.reserved_ip
	  }
	`, vpcName, subnetName, resIPName)
}
