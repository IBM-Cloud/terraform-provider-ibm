package ibm

import (
	"errors"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.ibm.com/ibmcloud/vpc-go-sdk/vpcclassicv1"
	"github.ibm.com/ibmcloud/vpc-go-sdk/vpcv1"
)

func TestAccIBMISSubnet_basic(t *testing.T) {
	var subnet string
	vpcname := fmt.Sprintf("tfsubnet-vpc-%d", acctest.RandIntRange(10, 100))
	gwname := fmt.Sprintf("tfsubnet-gw-%d", acctest.RandIntRange(10, 100))
	name1 := fmt.Sprintf("tfsubnet-%d", acctest.RandIntRange(10, 100))
	name2 := fmt.Sprintf("tfsubnet-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMISSubnetDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISSubnetConfig(vpcname, name1, ISZoneName, ISCIDR),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISSubnetExists("ibm_is_subnet.testacc_subnet", subnet),
					resource.TestCheckResourceAttr(
						"ibm_is_subnet.testacc_subnet", "name", name1),
					resource.TestCheckResourceAttr(
						"ibm_is_subnet.testacc_subnet", "zone", ISZoneName),
					resource.TestCheckResourceAttr(
						"ibm_is_subnet.testacc_subnet", "ipv4_cidr_block", ISCIDR),
				),
			},
			{
				Config: testAccCheckIBMISSubnetConfigUpdate(vpcname, name2, ISZoneName, ISCIDR, gwname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISSubnetExists("ibm_is_subnet.testacc_subnet", subnet),
					resource.TestCheckResourceAttr(
						"ibm_is_subnet.testacc_subnet", "name", name2),
					resource.TestCheckResourceAttr(
						"ibm_is_subnet.testacc_subnet", "zone", ISZoneName),
					resource.TestCheckResourceAttr(
						"ibm_is_subnet.testacc_subnet", "ipv4_cidr_block", ISCIDR),
					resource.TestCheckResourceAttrSet(
						"ibm_is_subnet.testacc_subnet", "public_gateway"),
				),
			},
		},
	})
}

func testAccCheckIBMISSubnetDestroy(s *terraform.State) error {
	userDetails, _ := testAccProvider.Meta().(ClientSession).BluemixUserDetails()

	if userDetails.generation == 1 {
		sess, _ := testAccProvider.Meta().(ClientSession).VpcClassicV1API()
		for _, rs := range s.RootModule().Resources {
			if rs.Type != "ibm_is_subnet" {
				continue
			}

			getsubnetoptions := &vpcclassicv1.GetSubnetOptions{
				ID: &rs.Primary.ID,
			}
			_, _, err := sess.GetSubnet(getsubnetoptions)
			if err == nil {
				return fmt.Errorf("subnet still exists: %s", rs.Primary.ID)
			}
		}
	} else {
		sess, _ := testAccProvider.Meta().(ClientSession).VpcV1API()
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

		userDetails, _ := testAccProvider.Meta().(ClientSession).BluemixUserDetails()

		if userDetails.generation == 1 {
			sess, _ := testAccProvider.Meta().(ClientSession).VpcClassicV1API()
			getsubnetoptions := &vpcclassicv1.GetSubnetOptions{
				ID: &rs.Primary.ID,
			}
			foundsubnet, _, err := sess.GetSubnet(getsubnetoptions)
			if err != nil {
				return err
			}
			subnetID = *foundsubnet.ID
		} else {
			sess, _ := testAccProvider.Meta().(ClientSession).VpcV1API()
			getsubnetoptions := &vpcv1.GetSubnetOptions{
				ID: &rs.Primary.ID,
			}
			foundsubnet, _, err := sess.GetSubnet(getsubnetoptions)
			if err != nil {
				return err
			}
			subnetID = *foundsubnet.ID
		}
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
	}`, vpcname, gwname, zone, name, zone, cidr)
}
