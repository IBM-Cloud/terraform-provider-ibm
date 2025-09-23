// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"errors"
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"

	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMISSubnet_basic(t *testing.T) {
	var subnet string
	vpcname := fmt.Sprintf("tfsubnet-vpc-%d", acctest.RandIntRange(10, 100))
	gwname := fmt.Sprintf("tfsubnet-gw-%d", acctest.RandIntRange(10, 100))
	name1 := fmt.Sprintf("tfsubnet-%d", acctest.RandIntRange(10, 100))
	name2 := fmt.Sprintf("tfsubnet-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISSubnetDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISSubnetConfig(vpcname, name1, acc.ISZoneName, acc.ISCIDR),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISSubnetExists("ibm_is_subnet.testacc_subnet", subnet),
					resource.TestCheckResourceAttr(
						"ibm_is_subnet.testacc_subnet", "name", name1),
					resource.TestCheckResourceAttr(
						"ibm_is_subnet.testacc_subnet", "zone", acc.ISZoneName),
					resource.TestCheckResourceAttr(
						"ibm_is_subnet.testacc_subnet", "ipv4_cidr_block", acc.ISCIDR),
					resource.TestCheckResourceAttr(
						"ibm_is_subnet.testacc_subnet", "tags.#", "2"),
				),
			},
			{
				Config: testAccCheckIBMISSubnetConfigUpdate(vpcname, name2, acc.ISZoneName, acc.ISCIDR, gwname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISSubnetExists("ibm_is_subnet.testacc_subnet", subnet),
					resource.TestCheckResourceAttr(
						"ibm_is_subnet.testacc_subnet", "name", name2),
					resource.TestCheckResourceAttr(
						"ibm_is_subnet.testacc_subnet", "zone", acc.ISZoneName),
					resource.TestCheckResourceAttr(
						"ibm_is_subnet.testacc_subnet", "ipv4_cidr_block", acc.ISCIDR),
					resource.TestCheckResourceAttrSet(
						"ibm_is_subnet.testacc_subnet", "public_gateway"),
					resource.TestCheckResourceAttr(
						"ibm_is_subnet.testacc_subnet", "tags.#", "1"),
				),
			},
		},
	})
}

func testAccCheckIBMISSubnetDestroy(s *terraform.State) error {

	sess, _ := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_subnet" {
			continue
		}

		getsubnetoptions := &vpcv1.GetSubnetOptions{
			ID: &rs.Primary.ID,
		}
		_, _, err := sess.GetSubnet(getsubnetoptions)

		if err == nil {
			return fmt.Errorf("subnet still exists: %s", rs.Primary.ID)
		}
	}

	return nil
}

func testAccCheckIBMISSubnetExists(n, subnetID string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}

		sess, _ := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
		getsubnetoptions := &vpcv1.GetSubnetOptions{
			ID: &rs.Primary.ID,
		}
		foundsubnet, _, err := sess.GetSubnet(getsubnetoptions)
		if err != nil {
			return err
		}
		subnetID = *foundsubnet.ID
		return nil
	}
}

func testAccCheckIBMISSubnetConfig(vpcname, name, zone, cidr string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	}

	resource "ibm_is_subnet" "testacc_subnet" {
		name = "%s"
		vpc = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		ipv4_cidr_block = "%s"
		tags = ["Tag1", "tag2"]
	}`, vpcname, name, zone, cidr)
}

func testAccCheckIBMISSubnetConfigUpdate(vpcname, name, zone, cidr, gwname string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	}

	resource "ibm_is_public_gateway" "testacc_gw" {
		name = "%s"
		vpc  = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
	}

	resource "ibm_is_subnet" "testacc_subnet" {
		name = "%s"
		vpc = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		ipv4_cidr_block = "%s"
		public_gateway = ibm_is_public_gateway.testacc_gw.id
		tags = ["tag1"]
	}`, vpcname, gwname, zone, name, zone, cidr)
}
