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

func TestAccIBMISSubnetRoutingTableAttachment_basic(t *testing.T) {
	var subnetRT string
	rtname := fmt.Sprintf("tfrt-%d", acctest.RandIntRange(10, 100))
	vpcname := fmt.Sprintf("tfvpc-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tfsubnet-vpc-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: checkSubnetRoutingTableAttachmentDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMISSubnetRoutingTableAttachmentConfig(rtname, subnetname, vpcname, acc.ISZoneName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISSubnetRoutingTableAttachmentExists("ibm_is_subnet_routing_table_attachment.attach", subnetRT),
					resource.TestCheckResourceAttrSet(
						"ibm_is_subnet_routing_table_attachment.attach", "lifecycle_state"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_subnet_routing_table_attachment.attach", "resource_type"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_subnet_routing_table_attachment.attach", "id"),
				),
			},
		},
	})
}

func checkSubnetRoutingTableAttachmentDestroy(s *terraform.State) error {

	sess, _ := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_subnet_routing_table_attachment" {
			continue
		}

		getSubnetRoutingTableOptionsModel := &vpcv1.GetSubnetRoutingTableOptions{
			ID: &rs.Primary.ID,
		}
		_, _, err := sess.GetSubnetRoutingTable(getSubnetRoutingTableOptionsModel)

		if err == nil {
			return fmt.Errorf("subnet routing table attachment still exists: %s", rs.Primary.ID)
		}
	}
	return nil
}

func testAccCheckIBMISSubnetRoutingTableAttachmentExists(n, subnetRT string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}

		sess, _ := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
		getSubnetRoutingTableOptionsModel := &vpcv1.GetSubnetRoutingTableOptions{
			ID: &rs.Primary.ID,
		}
		foundsubnetRT, _, err := sess.GetSubnetRoutingTable(getSubnetRoutingTableOptionsModel)
		if err != nil {
			return err
		}
		subnetRT = *foundsubnetRT.ID
		return nil
	}
}

func testAccCheckIBMISSubnetRoutingTableAttachmentConfig(rtname, subnetname, vpcname, zone string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	}
	resource "ibm_is_vpc_routing_table" "testacc_vpc_routing_table" {
		vpc = ibm_is_vpc.testacc_vpc.id
		name = "%s"
	}

	resource "ibm_is_subnet" "testacc_subnet" {
		name = "%s"
		vpc = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		total_ipv4_address_count  = 16
	}

	resource "ibm_is_subnet_routing_table_attachment" "attach" {
		depends_on = [ibm_is_vpc_routing_table.testacc_vpc_routing_table, ibm_is_subnet.testacc_subnet]
		subnet      = ibm_is_subnet.testacc_subnet.id
		routing_table = ibm_is_vpc_routing_table.testacc_vpc_routing_table.routing_table
	  }

	`, vpcname, rtname, subnetname, zone)
}
