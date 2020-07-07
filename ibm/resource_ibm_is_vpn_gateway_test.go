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

func TestAccIBMISVPNGateway_basic(t *testing.T) {
	var vpnGateway string
	vpcname := fmt.Sprintf("tfvpnuat-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tfvpnuat-subnet-%d", acctest.RandIntRange(10, 100))
	name1 := fmt.Sprintf("tfvpnuat-createname-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMISVPNGatewayDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVPNGatewayConfig(vpcname, subnetname, name1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPNGatewayExists("ibm_is_vpn_gateway.testacc_vpnGateway", vpnGateway),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway.testacc_vpnGateway", "name", name1),
				),
			},
		},
	})
}

func testAccCheckIBMISVPNGatewayDestroy(s *terraform.State) error {
	userDetails, _ := testAccProvider.Meta().(ClientSession).BluemixUserDetails()

	if userDetails.generation == 1 {
		sess, _ := testAccProvider.Meta().(ClientSession).VpcClassicV1API()
		for _, rs := range s.RootModule().Resources {
			if rs.Type != "ibm_is_vpn_gateway" {
				continue
			}

			getvpngcptions := &vpcclassicv1.GetVpnGatewayConnectionOptions{
				ID: &rs.Primary.ID,
			}
			_, _, err := sess.GetVpnGatewayConnection(getvpngcptions)

			if err == nil {
				return fmt.Errorf("vpnGateway still exists: %s", rs.Primary.ID)
			}
		}
	} else {
		sess, _ := testAccProvider.Meta().(ClientSession).VpcV1API()
		for _, rs := range s.RootModule().Resources {
			if rs.Type != "ibm_is_vpn_gateway" {
				continue
			}

			getvpngcptions := &vpcv1.GetVpnGatewayConnectionOptions{
				ID: &rs.Primary.ID,
			}
			_, _, err := sess.GetVpnGatewayConnection(getvpngcptions)

			if err == nil {
				return fmt.Errorf("vpnGateway still exists: %s", rs.Primary.ID)
			}
		}
	}

	return nil
}

func testAccCheckIBMISVPNGatewayExists(n, vpnGatewayID string) resource.TestCheckFunc {
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
			getvpngcptions := &vpcclassicv1.GetVpnGatewayOptions{
				ID: &rs.Primary.ID,
			}
			foundvpnGateway, _, err := sess.GetVpnGateway(getvpngcptions)
			if err != nil {
				return err
			}
			vpnGatewayID = *foundvpnGateway.ID
		} else {
			sess, _ := testAccProvider.Meta().(ClientSession).VpcV1API()
			getvpngcptions := &vpcv1.GetVpnGatewayOptions{
				ID: &rs.Primary.ID,
			}
			foundvpnGateway, _, err := sess.GetVpnGateway(getvpngcptions)
			if err != nil {
				return err
			}
			vpnGatewayID = *foundvpnGateway.ID
		}
		return nil
	}
}

func testAccCheckIBMISVPNGatewayConfig(vpc, subnet, name string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	}
	
	resource "ibm_is_subnet" "testacc_subnet" {
		name = "%s"
		vpc = "${ibm_is_vpc.testacc_vpc.id}"
		zone = "%s"
		ipv4_cidr_block = "%s"
	}
	resource "ibm_is_vpn_gateway" "testacc_vpnGateway" {
	name = "%s"
	subnet = "${ibm_is_subnet.testacc_subnet.id}"
	}`, vpc, subnet, ISZoneName, ISCIDR, name)

}
