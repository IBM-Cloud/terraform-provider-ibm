package ibm

import (
	"errors"
	"fmt"
	"testing"

	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccIBMISSubnetNetworkACLAttachment_basic(t *testing.T) {
	var subnetNwACL string
	nwaclname := fmt.Sprintf("tfnw-acl-%d", acctest.RandIntRange(10, 100))
	vpcname := fmt.Sprintf("tfsubnet-vpc-%d", acctest.RandIntRange(10, 100))
	name1 := fmt.Sprintf("tfsubnet-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: checkSubnetNetworkACLAttachmentDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMISSubnetNetworkACLAttachmentConfig(nwaclname, vpcname, name1, ISZoneName, ISCIDR),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISSubnetNetworkACLAttachmentExists("ibm_is_subnet_network_acl_attachment.attach", subnetNwACL),
				),
			},
		},
	})
}

func checkSubnetNetworkACLAttachmentDestroy(s *terraform.State) error {

	sess, _ := testAccProvider.Meta().(ClientSession).VpcV1API()
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
		fmt.Println("siv ", s.RootModule().Resources)
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}

		sess, _ := testAccProvider.Meta().(ClientSession).VpcV1API()
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
