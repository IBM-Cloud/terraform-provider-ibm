// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"errors"
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"

	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMISPublicGateway_basic(t *testing.T) {
	var publicgw string
	vpcname := fmt.Sprintf("tfpgw-vpc-%d", acctest.RandIntRange(10, 100))
	name1 := fmt.Sprintf("tf-create-name-%d", acctest.RandIntRange(10, 100))
	//name2 := fmt.Sprintf("tfpgw-update-name-%d", acctest.RandIntRange(10, 100))

	zone := "us-south-1"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISPublicGatewayDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISPublicGatewayConfig(vpcname, name1, zone),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISPublicGatewayExists("ibm_is_public_gateway.testacc_public_gateway", publicgw),
					resource.TestCheckResourceAttr(
						"ibm_is_public_gateway.testacc_public_gateway", "name", name1),
					resource.TestCheckResourceAttr(
						"ibm_is_public_gateway.testacc_public_gateway", "zone", zone),
				),
			},

			/*			{
						Config: testAccCheckIBMISPublicGatewayConfig(vpcname, name2, zone, cidr),
						Check: resource.ComposeTestCheckFunc(
							testAccCheckIBMISPublicGatewayExists("ibm_is_publicgw.testacc_publicgw", publicgw),
							resource.TestCheckResourceAttr(
								"ibm_is_publicgw.testacc_publicgw", "name", name2),
							resource.TestCheckResourceAttr(
								"ibm_is_publicgw.testacc_publicgw", "zone", zone),
							resource.TestCheckResourceAttr(
								"ibm_is_publicgw.testacc_publicgw", "ipv4_cidr_block", cidr),
						),
					},*/
		},
	})
}
func TestAccIBMISPublicGateway_floatingip(t *testing.T) {
	var publicgw string
	vpcname := fmt.Sprintf("tfpgw-vpc-%d", acctest.RandIntRange(10, 100))
	name1 := fmt.Sprintf("tfpgw-pg-%d", acctest.RandIntRange(10, 100))
	//name2 := fmt.Sprintf("tfpgw-update-name-%d", acctest.RandIntRange(10, 100))

	fipname := fmt.Sprintf("tfpgw-fip-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISPublicGatewayDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISPublicGatewayFloatingIpConfig(vpcname, name1, fipname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISPublicGatewayExists("ibm_is_public_gateway.testacc_public_gateway", publicgw),
					resource.TestCheckResourceAttr(
						"ibm_is_public_gateway.testacc_public_gateway", "name", name1),
					resource.TestCheckResourceAttr(
						"ibm_is_public_gateway.testacc_public_gateway", "zone", acc.ISZoneName),
					resource.TestCheckResourceAttrSet(
						"ibm_is_public_gateway.testacc_public_gateway", "floating_ip.id"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_public_gateway.testacc_public_gateway", "floating_ip.address"),
					resource.TestCheckResourceAttr(
						"ibm_is_floating_ip.testacc_fip", "name", fipname),
					resource.TestCheckResourceAttrSet(
						"ibm_is_floating_ip.testacc_fip", "id"),
				),
			},
		},
	})
}
func TestAccIBMISPublicGateway_resource_group_change(t *testing.T) {
	var publicgw string
	vpcname := fmt.Sprintf("tfpgw-vpc-%d", acctest.RandIntRange(10, 100))
	name1 := fmt.Sprintf("tf-create-name-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tfpgw-subnet-%d", acctest.RandIntRange(10, 100))
	flag := true

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISPublicGatewayDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISPublicGatewayRgChangeConfig(vpcname, subnetname, name1, !flag),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISPublicGatewayExists("ibm_is_public_gateway.testacc_public_gateway", publicgw),
					resource.TestCheckResourceAttr(
						"ibm_is_public_gateway.testacc_public_gateway", "name", name1),
					resource.TestCheckResourceAttr(
						"ibm_is_public_gateway.testacc_public_gateway", "zone", acc.ISZoneName),
				),
			},
			{
				Config: testAccCheckIBMISPublicGatewayRgChangeConfig(vpcname, subnetname, name1, flag),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISPublicGatewayExists("ibm_is_public_gateway.testacc_public_gateway", publicgw),
					resource.TestCheckResourceAttr(
						"ibm_is_public_gateway.testacc_public_gateway", "name", name1),
					resource.TestCheckResourceAttr(
						"ibm_is_public_gateway.testacc_public_gateway", "zone", acc.ISZoneName),
				),
			},
		},
	})
}

func testAccCheckIBMISPublicGatewayDestroy(s *terraform.State) error {
	sess, _ := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_public_gateway" {
			continue
		}

		getpgwoptions := &vpcv1.GetPublicGatewayOptions{
			ID: &rs.Primary.ID,
		}
		_, _, err := sess.GetPublicGateway(getpgwoptions)
		if err == nil {
			return fmt.Errorf("publicgw still exists: %s", rs.Primary.ID)
		}
	}

	return nil
}

func testAccCheckIBMISPublicGatewayExists(n, publicgw string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}

		sess, _ := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
		getpgwoptions := &vpcv1.GetPublicGatewayOptions{
			ID: &rs.Primary.ID,
		}
		foundpublicgw, _, err := sess.GetPublicGateway(getpgwoptions)
		if err != nil {
			return err
		}
		publicgw = *foundpublicgw.ID
		return nil
	}
}

func testAccCheckIBMISPublicGatewayConfig(vpcname, name, zone string) string {
	return fmt.Sprintf(`
		resource "ibm_is_vpc" "testacc_vpc" {
			name = "%s"
		}

		resource "ibm_is_public_gateway" "testacc_public_gateway" {
			name = "%s"
			vpc = "${ibm_is_vpc.testacc_vpc.id}"
			zone = "%s"
		}`, vpcname, name, zone)

}
func testAccCheckIBMISPublicGatewayFloatingIpConfig(vpcname, name, fipname string) string {
	return fmt.Sprintf(`
		resource "ibm_is_vpc" "testacc_vpc" {
			name = "%s"
		}

		resource "ibm_is_floating_ip" "testacc_fip" {
			name = "%s"
			zone = "%s"
		}
	  

		resource "ibm_is_public_gateway" "testacc_public_gateway" {
			name 	= "%s"
			vpc 	= "${ibm_is_vpc.testacc_vpc.id}"
			zone 	= "%s"
		}
		
		`, vpcname, fipname, acc.ISZoneName, name, acc.ISZoneName)

}
func testAccCheckIBMISPublicGatewayRgChangeConfig(vpcname, subnetname, name string, flag bool) string {
	return fmt.Sprintf(`
		resource "ibm_is_vpc" "testacc_vpc" {
			name = "%s"
		}
		resource ibm_is_subnet subnet {
			name            = "%s"
			vpc             = ibm_is_vpc.testacc_vpc.id
			zone            = "%s"
			total_ipv4_address_count 	= 16
			public_gateway  = ibm_is_public_gateway.testacc_public_gateway.id
		}
		  
		resource "ibm_is_public_gateway" "testacc_public_gateway" {
			name 					= "%s"
			resource_group  		= %t ? "%s": null
			vpc 					= "${ibm_is_vpc.testacc_vpc.id}"
			zone 					= "%s"
		}
		
		`, vpcname, subnetname, acc.ISZoneName, name, flag, acc.IsResourceGroupID, acc.ISZoneName)

}
