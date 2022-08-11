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

func TestAccIBMISSubnetReservedIPResource_basic(t *testing.T) {
	var reservedIPID string
	vpcName := fmt.Sprintf("tfresip-vpc-%d", acctest.RandIntRange(10, 100))
	subnetName := fmt.Sprintf("tfresip-subnet-%d", acctest.RandIntRange(10, 100))
	reservedIPName := fmt.Sprintf("tfresip-reservedip-%d", acctest.RandIntRange(10, 100))
	reservedIPName2 := fmt.Sprintf("tfresip-reservedip-%d", acctest.RandIntRange(10, 100))
	terraformTag := "ibm_is_subnet_reserved_ip.resIP1"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckisSubnetReservedIPDestroy,
		Steps: []resource.TestStep{
			{
				// Tests create
				Config: testAccCheckISSubnetReservedIPConfigBasic(vpcName, subnetName, reservedIPName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckISSubnetReservedIPExists(terraformTag, &reservedIPID),
					resource.TestCheckResourceAttr(terraformTag, "name", reservedIPName),
				),
			},
			{
				// Tests Update
				Config: testAccCheckISSubnetReservedIPConfigBasic(vpcName, subnetName, reservedIPName2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckISSubnetReservedIPExists(terraformTag, &reservedIPID),
					resource.TestCheckResourceAttr(terraformTag, "name", reservedIPName2),
				),
			},
		},
	})
}
func TestAccIBMISSubnetReservedIPResource_address(t *testing.T) {
	var reservedIPID string
	vpcName := fmt.Sprintf("tfresip-vpc-%d", acctest.RandIntRange(10, 100))
	subnetName := fmt.Sprintf("tfresip-subnet-%d", acctest.RandIntRange(10, 100))
	reservedIPName := fmt.Sprintf("tfresip-reservedip-%d", acctest.RandIntRange(10, 100))
	terraformTag := "ibm_is_subnet_reserved_ip.resIP1"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckisSubnetReservedIPDestroy,
		Steps: []resource.TestStep{
			{
				// Tests create
				Config: testAccCheckISSubnetReservedIPConfigAddress(vpcName, subnetName, reservedIPName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckISSubnetReservedIPExists(terraformTag, &reservedIPID),
					resource.TestCheckResourceAttr(terraformTag, "name", reservedIPName),
					resource.TestCheckResourceAttr(terraformTag, "address", "10.240.0.14"),
				),
			},
		},
	})
}

func testAccCheckisSubnetReservedIPDestroy(s *terraform.State) error {
	sess, err := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_subnet_reserved_ip" {
			continue
		}

		parts, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}

		opt := sess.NewGetSubnetReservedIPOptions(parts[0], parts[1])
		_, response, err := sess.GetSubnetReservedIP(opt)
		if err == nil {
			return fmt.Errorf("Reserved IP still exists: %v", response)
		}
	}
	return nil
}

func testAccCheckISSubnetReservedIPExists(resIPName string, reservedIPID *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resIPName]
		if !ok {
			return fmt.Errorf("Not Found (reserved IP): %s", resIPName)
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

func testAccCheckISSubnetReservedIPConfigBasic(vpcName, subnetName, resIPName string) string {
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
		name = "%s"
		target = ibm_is_virtual_endpoint_gateway.endpoint_gateway.id
	  }
	`, vpcName, subnetName, resIPName)
}
func testAccCheckISSubnetReservedIPConfigAddress(vpcName, subnetName, resIPName string) string {
	return fmt.Sprintf(`
	  resource "ibm_is_vpc" "vpc1" {
		name = "%s"
	  }

	  resource "ibm_is_subnet" "subnet1" {
		name                     = "%s"
		vpc                      = ibm_is_vpc.vpc1.id
		zone 					 = "%s"
		ipv4_cidr_block 		 = "%s"
	  }

	  resource "ibm_is_subnet_reserved_ip" "resIP1" {
		subnet 		= ibm_is_subnet.subnet1.id
		name 		= "%s"
		address		= "${replace(ibm_is_subnet.subnet1.ipv4_cidr_block, "0/24", "14")}"
	  }
	`, vpcName, subnetName, acc.ISZoneName, acc.ISCIDR, resIPName)
}
