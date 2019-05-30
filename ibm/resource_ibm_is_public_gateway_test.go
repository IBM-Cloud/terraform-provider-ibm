package ibm

import (
	"errors"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.ibm.com/Bluemix/riaas-go-client/clients/network"
	"github.ibm.com/riaas/rias-api/riaas/models"
)

func TestAccIBMISPublicGateway_basic(t *testing.T) {
	var publicgw *models.PublicGateway
	vpcname := fmt.Sprintf("terraformpublicgwuat_vpc_%d", acctest.RandInt())
	name1 := fmt.Sprintf("terraformpublicgwuat_create_step_name_%d", acctest.RandInt())
	//name2 := fmt.Sprintf("terraformpublicgwuat_update_step_name_%d", acctest.RandInt())

	zone := "us-south-1"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMISPublicGatewayDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISPublicGatewayConfig(vpcname, name1, zone),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISPublicGatewayExists("ibm_is_public_gateway.testacc_public_gateway", &publicgw),
					resource.TestCheckResourceAttr(
						"ibm_is_public_gateway.testacc_public_gateway", "name", name1),
					resource.TestCheckResourceAttr(
						"ibm_is_public_gateway.testacc_public_gateway", "zone", zone),
				),
			},

			/*			{
						Config: testAccCheckIBMISPublicGatewayConfig(vpcname, name2, zone, cidr),
						Check: resource.ComposeTestCheckFunc(
							testAccCheckIBMISPublicGatewayExists("ibm_is_publicgw.testacc_publicgw", &publicgw),
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

func testAccCheckIBMISPublicGatewayDestroy(s *terraform.State) error {
	sess, _ := testAccProvider.Meta().(ClientSession).ISSession()

	publicgwC := network.NewPublicGatewayClient(sess)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_public_gateway" {
			continue
		}

		_, err := publicgwC.Get(rs.Primary.ID)

		if err == nil {
			return fmt.Errorf("publicgw still exists: %s", rs.Primary.ID)
		}
	}

	return nil
}

func testAccCheckIBMISPublicGatewayExists(n string, publicgw **models.PublicGateway) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		fmt.Println("siv ", s.RootModule().Resources)
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}

		sess, _ := testAccProvider.Meta().(ClientSession).ISSession()
		publicgwC := network.NewPublicGatewayClient(sess)
		foundpublicgw, err := publicgwC.Get(rs.Primary.ID)

		if err != nil {
			return err
		}

		*publicgw = foundpublicgw
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
