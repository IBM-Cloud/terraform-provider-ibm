// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"errors"
	"fmt"
	"testing"

	"github.com/IBM/vpc-go-sdk/vpcclassicv1"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccIBMISPublicGateway_basic(t *testing.T) {
	var publicgw string
	vpcname := fmt.Sprintf("tfpgw-vpc-%d", acctest.RandIntRange(10, 100))
	name1 := fmt.Sprintf("tf-create-name-%d", acctest.RandIntRange(10, 100))
	//name2 := fmt.Sprintf("tfpgw-update-name-%d", acctest.RandIntRange(10, 100))

	zone := "us-south-1"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
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

func testAccCheckIBMISPublicGatewayDestroy(s *terraform.State) error {
	userDetails, _ := testAccProvider.Meta().(ClientSession).BluemixUserDetails()

	if userDetails.generation == 1 {
		sess, _ := testAccProvider.Meta().(ClientSession).VpcClassicV1API()
		for _, rs := range s.RootModule().Resources {
			if rs.Type != "ibm_is_public_gateway" {
				continue
			}

			getpgwoptions := &vpcclassicv1.GetPublicGatewayOptions{
				ID: &rs.Primary.ID,
			}
			_, _, err := sess.GetPublicGateway(getpgwoptions)
			if err == nil {
				return fmt.Errorf("publicgw still exists: %s", rs.Primary.ID)
			}
		}
	} else {
		sess, _ := testAccProvider.Meta().(ClientSession).VpcV1API()
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
	}

	return nil
}

func testAccCheckIBMISPublicGatewayExists(n, publicgw string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		fmt.Println("siv ", s.RootModule().Resources)
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
			getpgwoptions := &vpcclassicv1.GetPublicGatewayOptions{
				ID: &rs.Primary.ID,
			}
			foundpublicgw, _, err := sess.GetPublicGateway(getpgwoptions)
			if err != nil {
				return err
			}
			publicgw = *foundpublicgw.ID
		} else {
			sess, _ := testAccProvider.Meta().(ClientSession).VpcV1API()
			getpgwoptions := &vpcv1.GetPublicGatewayOptions{
				ID: &rs.Primary.ID,
			}
			foundpublicgw, _, err := sess.GetPublicGateway(getpgwoptions)
			if err != nil {
				return err
			}
			publicgw = *foundpublicgw.ID
		}
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
