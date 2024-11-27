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
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccIBMISSubnetNetworkACLAttachment_basic(t *testing.T) {
	var subnetNwACL string
	nwaclname := fmt.Sprintf("tfnw-acl-%d", acctest.RandIntRange(10, 100))
	vpcname := fmt.Sprintf("tfsubnet-vpc-%d", acctest.RandIntRange(10, 100))
	name1 := fmt.Sprintf("tfsubnet-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: checkSubnetNetworkACLAttachmentDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISSubnetNetworkACLAttachmentConfig(nwaclname, vpcname, name1, acc.ISZoneName, acc.ISCIDR),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISSubnetNetworkACLAttachmentExists("ibm_is_subnet_network_acl_attachment.attach", subnetNwACL),
				),
			},
		},
	})
}

func checkSubnetNetworkACLAttachmentDestroy(s *terraform.State) error {

	sess, _ := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_subnet_network_acl_attachment" {
			continue
		}

		getSubnetNetworkACLOptionsModel := &vpcv1.GetSubnetNetworkACLOptions{
			ID: &rs.Primary.ID,
		}
		_, _, err := sess.GetSubnetNetworkACL(getSubnetNetworkACLOptionsModel)

		if err == nil {
			return fmt.Errorf("subnet network acl attachment still exists: %s", rs.Primary.ID)
		}
	}
	return nil
}

func testAccCheckIBMISSubnetNetworkACLAttachmentExists(n, subnetNwACL string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}

		sess, _ := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
		getSubnetNetworkACLOptionsModel := &vpcv1.GetSubnetNetworkACLOptions{
			ID: &rs.Primary.ID,
		}
		foundSubnetNwACL, _, err := sess.GetSubnetNetworkACL(getSubnetNetworkACLOptionsModel)
		if err != nil {
			return err
		}
		subnetNwACL = *foundSubnetNwACL.ID
		return nil
	}
}

func testAccCheckIBMISSubnetNetworkACLAttachmentConfig(nwaclname, vpcname, name, zone, cidr string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	}

	resource "ibm_is_network_acl" "isExampleACL" {
		name = "%s"
		vpc  = ibm_is_vpc.testacc_vpc.id

		rules {
		  name        = "outbound"
		  action      = "allow"
		  source      = "0.0.0.0/0"
		  destination = "0.0.0.0/0"
		  direction   = "outbound"
		  icmp {
			code = 1
			type = 1
		  }
		  # Optionals :
		  # port_max =
		  # port_min =
		}
		rules {
		  name        = "inbound"
		  action      = "allow"
		  source      = "0.0.0.0/0"
		  destination = "0.0.0.0/0"
		  direction   = "inbound"
		  icmp {
			code = 1
			type = 1
		  }
		  # Optionals :
		  # port_max =
		  # port_min =
		}
	  }

	resource "ibm_is_subnet" "testacc_subnet" {
		name = "%s"
		vpc = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		ipv4_cidr_block = "%s"
	}

	resource "ibm_is_subnet_network_acl_attachment" attach {
		depends_on = [ibm_is_network_acl.isExampleACL, ibm_is_subnet.testacc_subnet]
		subnet      = ibm_is_subnet.testacc_subnet.id
		network_acl = ibm_is_network_acl.isExampleACL.id
	  }

	`, vpcname, nwaclname, name, zone, cidr)
}
