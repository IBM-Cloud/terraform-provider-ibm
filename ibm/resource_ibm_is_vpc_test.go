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

func TestAccIBMISVPC_basic(t *testing.T) {
	var vpc string
	name1 := fmt.Sprintf("terraformvpcuat-%d", acctest.RandIntRange(10, 100))
	name2 := fmt.Sprintf("terraformvpcuat-%d", acctest.RandIntRange(10, 100))
	apm := "manual"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMISVPCDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVPCConfig(name1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPCExists("ibm_is_vpc.testacc_vpc", vpc),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc.testacc_vpc", "name", name1),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc.testacc_vpc", "tags.#", "2"),
				),
			},
			{
				Config: testAccCheckIBMISVPCConfigUpdate(name1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPCExists("ibm_is_vpc.testacc_vpc", vpc),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc.testacc_vpc", "name", name1),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc.testacc_vpc", "tags.#", "1"),
				),
			},
			{
				Config: testAccCheckIBMISVPCConfig1(name2, apm),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPCExists("ibm_is_vpc.testacc_vpc1", vpc),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc.testacc_vpc1", "name", name2),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc.testacc_vpc1", "tags.#", "2"),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc.testacc_vpc1", "cse_source_addresses.#", "3"),
				),
			},
		},
	})
}

func testAccCheckIBMISVPCDestroy(s *terraform.State) error {
	userDetails, _ := testAccProvider.Meta().(ClientSession).BluemixUserDetails()

	if userDetails.generation == 1 {
		sess, _ := testAccProvider.Meta().(ClientSession).VpcClassicV1API()
		for _, rs := range s.RootModule().Resources {
			if rs.Type != "ibm_is_vpc" {
				continue
			}

			getvpcoptions := &vpcclassicv1.GetVpcOptions{
				ID: &rs.Primary.ID,
			}
			_, _, err := sess.GetVpc(getvpcoptions)

			if err == nil {
				return fmt.Errorf("vpc still exists: %s", rs.Primary.ID)
			}
		}
	} else {
		sess, _ := testAccProvider.Meta().(ClientSession).VpcV1API()
		for _, rs := range s.RootModule().Resources {
			if rs.Type != "ibm_is_vpc" {
				continue
			}

			getvpcoptions := &vpcv1.GetVpcOptions{
				ID: &rs.Primary.ID,
			}
			_, _, err := sess.GetVpc(getvpcoptions)

			if err == nil {
				return fmt.Errorf("vpc still exists: %s", rs.Primary.ID)
			}
		}
	}
	return nil
}

func testAccCheckIBMISVPCExists(n, vpcID string) resource.TestCheckFunc {
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
			getvpcoptions := &vpcclassicv1.GetVpcOptions{
				ID: &rs.Primary.ID,
			}
			foundvpc, _, err := sess.GetVpc(getvpcoptions)
			if err != nil {
				return err
			}
			vpcID = *foundvpc.ID
		} else {
			sess, _ := testAccProvider.Meta().(ClientSession).VpcV1API()
			getvpcoptions := &vpcv1.GetVpcOptions{
				ID: &rs.Primary.ID,
			}
			foundvpc, _, err := sess.GetVpc(getvpcoptions)
			if err != nil {
				return err
			}
			vpcID = *foundvpc.ID
		}
		return nil
	}
}

func testAccCheckIBMISVPCConfig(name string) string {
	return fmt.Sprintf(`
resource "ibm_is_vpc" "testacc_vpc" {
	name = "%s"
	tags = ["Tag1", "tag2"]
}`, name)

}

func testAccCheckIBMISVPCConfigUpdate(name string) string {
	return fmt.Sprintf(`
resource "ibm_is_vpc" "testacc_vpc" {
	name = "%s"
	tags = ["tag1"]
}`, name)

}

func testAccCheckIBMISVPCConfig1(name string, apm string) string {
	return fmt.Sprintf(`
resource "ibm_is_vpc" "testacc_vpc1" {
	name = "%s"
	address_prefix_management = "%s"
	tags = ["Tag1", "tag2"]
}`, name, apm)

}
