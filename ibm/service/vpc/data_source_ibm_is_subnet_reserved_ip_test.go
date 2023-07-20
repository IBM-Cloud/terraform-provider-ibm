// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0
package vpc_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMISSubnetReservedIP_basic(t *testing.T) {
	vpcName := fmt.Sprintf("tfresip-vpc-%d", acctest.RandIntRange(10, 100))
	subnetName := fmt.Sprintf("tfresip-subnet-%d", acctest.RandIntRange(10, 100))
	resIPName := fmt.Sprintf("tfresip-reservedip-%d", acctest.RandIntRange(10, 100))
	terraformTag := "data.ibm_is_subnet_reserved_ip.data_resip1"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccIBMISReservedIPdataSoruceConfig(vpcName, subnetName, resIPName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(terraformTag, "name", resIPName),
				),
			},
		},
	})
}
func TestAccIBMISSubnetReservedIP_targetCrn(t *testing.T) {
	vpcName := fmt.Sprintf("tfresip-vpc-%d", acctest.RandIntRange(10, 100))
	subnetName := fmt.Sprintf("tfresip-subnet-%d", acctest.RandIntRange(10, 100))
	resIPName := fmt.Sprintf("tfresip-reservedip-%d", acctest.RandIntRange(10, 100))
	gatewayName := fmt.Sprintf("tfresip-egateway-%d", acctest.RandIntRange(10, 100))
	terraformTag := "data.ibm_is_subnet_reserved_ip.data_resip1"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccIBMISReservedIPdataSoruceTargetCrnConfig(vpcName, subnetName, resIPName, gatewayName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(terraformTag, "name", resIPName),
					resource.TestCheckResourceAttrSet(
						terraformTag, "target_crn"),
					resource.TestCheckResourceAttrSet(
						terraformTag, "target"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_virtual_endpoint_gateway_ip.example", "target.#"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_virtual_endpoint_gateway.example", "target.#"),
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
func testAccIBMISReservedIPdataSoruceTargetCrnConfig(vpcName, subnetName, reservedIPName, gatewayName string) string {
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

		resource "ibm_is_virtual_endpoint_gateway" "example" {

			name = "%s"
			target {
			name          = "ibm-ntp-server"
			resource_type = "provider_infrastructure_service"
			}
			vpc            = ibm_is_vpc.vpc1.id
		}

		resource "ibm_is_virtual_endpoint_gateway_ip" "example" {
			gateway     = ibm_is_virtual_endpoint_gateway.example.id
			reserved_ip = ibm_is_subnet_reserved_ip.resip1.reserved_ip
		}
  

		data "ibm_is_subnet_reserved_ip" "data_resip1" {
			subnet = ibm_is_subnet.subnet1.id
			reserved_ip = ibm_is_subnet_reserved_ip.resip1.reserved_ip
			depends_on 	= 	[
	  				ibm_is_virtual_endpoint_gateway_ip.example
							]
		}

      `, vpcName, subnetName, reservedIPName, gatewayName)
}
