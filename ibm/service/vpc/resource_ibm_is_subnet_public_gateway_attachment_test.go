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

func TestAccIBMISSubnetPublicGatewayAttachment_basic(t *testing.T) {
	var subnetPublicGateway string
	pgname := fmt.Sprintf("tfnw-pg-%d", acctest.RandIntRange(10, 100))
	vpcname := fmt.Sprintf("tfsubnet-vpc-%d", acctest.RandIntRange(10, 100))
	name1 := fmt.Sprintf("tfsubnet-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: checkSubnetPublicGatewayAttachmentDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMISSubnetPublicGatewayAttachmentConfig(vpcname, name1, acc.ISZoneName, pgname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISSubnetPublicGatewayAttachmentExists("ibm_is_subnet_public_gateway_attachment.attach", subnetPublicGateway),
					resource.TestCheckResourceAttrSet(
						"ibm_is_subnet_public_gateway_attachment.attach", "crn"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_subnet_public_gateway_attachment.attach", "id"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_subnet_public_gateway_attachment.attach", "resource_type"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_subnet_public_gateway_attachment.attach", "status"),
				),
			},
		},
	})
}

func checkSubnetPublicGatewayAttachmentDestroy(s *terraform.State) error {

	sess, _ := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_subnet_public_gateway_attachment" {
			continue
		}
		getSubnetPublicGatewayOptionsModel := &vpcv1.GetSubnetPublicGatewayOptions{
			ID: &rs.Primary.ID,
		}
		_, _, err := sess.GetSubnetPublicGateway(getSubnetPublicGatewayOptionsModel)

		if err == nil {
			return fmt.Errorf("subnet public gateway attachment still exists: %s", rs.Primary.ID)
		}
	}
	return nil
}

func testAccCheckIBMISSubnetPublicGatewayAttachmentExists(n, subnetPublicGateway string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}

		sess, _ := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
		getSubnetPublicGatewayOptionsModel := &vpcv1.GetSubnetPublicGatewayOptions{
			ID: &rs.Primary.ID,
		}
		foundSubnetPG, _, err := sess.GetSubnetPublicGateway(getSubnetPublicGatewayOptionsModel)
		if err != nil {
			return err
		}
		subnetPublicGateway = *foundSubnetPG.ID
		return nil
	}
}

func testAccCheckIBMISSubnetPublicGatewayAttachmentConfig(vpcname, name, zone, pgname string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	}

	resource "ibm_is_subnet" "testacc_subnet" {
		name 						= "%s"
		vpc 						= ibm_is_vpc.testacc_vpc.id
		zone 						= "%s"
		total_ipv4_address_count 	= 16
	}

	resource "ibm_is_public_gateway" "testacc_pg" {
		name = "%s"
		vpc  = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
	}

	resource "ibm_is_subnet_public_gateway_attachment" "attach" {
		depends_on 		= [ibm_is_public_gateway.testacc_pg, ibm_is_subnet.testacc_subnet]
		subnet      	= ibm_is_subnet.testacc_subnet.id
		public_gateway 	= ibm_is_public_gateway.testacc_pg.id
	}

	`, vpcname, name, zone, pgname, zone)
}
