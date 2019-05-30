package ibm

import (
	"errors"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.ibm.com/Bluemix/riaas-go-client/clients/vpn"
	"github.ibm.com/riaas/rias-api/riaas/models"
)

func TestAccIBMISVPNGateway_basic(t *testing.T) {
	var vpnGateway *models.VPNGateway
	vpcname := fmt.Sprintf("terraformvpnuat_vpc_%d", acctest.RandInt())
	subnetname := fmt.Sprintf("terraformvpnuat_subnet_%d", acctest.RandInt())
	name1 := fmt.Sprintf("terraformvpnGatewayuat_create_step_name_%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMISVPNGatewayDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVPNGatewayConfig(vpcname, subnetname, name1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPNGatewayExists("ibm_is_vpn_gateway.testacc_vpnGateway", &vpnGateway),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway.testacc_vpnGateway", "name", name1),
				),
			},
		},
	})
}

func testAccCheckIBMISVPNGatewayDestroy(s *terraform.State) error {
	sess, _ := testAccProvider.Meta().(ClientSession).ISSession()

	vpnGatewayC := vpn.NewVpnClient(sess)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_vpn_gateway" {
			continue
		}

		_, err := vpnGatewayC.Get(rs.Primary.ID)

		if err == nil {
			return fmt.Errorf("vpnGateway still exists: %s", rs.Primary.ID)
		}
	}

	return nil
}

func testAccCheckIBMISVPNGatewayExists(n string, vpnGateway **models.VPNGateway) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}

		sess, _ := testAccProvider.Meta().(ClientSession).ISSession()
		vpnGatewayC := vpn.NewVpnClient(sess)
		foundvpnGateway, err := vpnGatewayC.Get(rs.Primary.ID)

		if err != nil {
			return err
		}

		*vpnGateway = foundvpnGateway
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
