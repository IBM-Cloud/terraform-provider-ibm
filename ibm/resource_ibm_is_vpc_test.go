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

func TestAccIBMISVPC_basic(t *testing.T) {
	var vpc *models.Vpc
	name1 := fmt.Sprintf("terraformvpcuat_create_step_name_%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMISVPCDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVPCConfig(name1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPCExists("ibm_is_vpc.testacc_vpc", &vpc),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc.testacc_vpc", "name", name1),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc.testacc_vpc", "tags.#", "2"),
				),
			},
			{
				Config: testAccCheckIBMISVPCConfigUpdate(name1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPCExists("ibm_is_vpc.testacc_vpc", &vpc),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc.testacc_vpc", "name", name1),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc.testacc_vpc", "tags.#", "1"),
				),
			},
		},
	})
}

func testAccCheckIBMISVPCDestroy(s *terraform.State) error {
	sess, _ := testAccProvider.Meta().(ClientSession).ISSession()

	vpcC := network.NewVPCClient(sess)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_vpc" {
			continue
		}

		_, err := vpcC.Get(rs.Primary.ID)

		if err == nil {
			return fmt.Errorf("vpc still exists: %s", rs.Primary.ID)
		}
	}

	return nil
}

func testAccCheckIBMISVPCExists(n string, vpc **models.Vpc) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}

		sess, _ := testAccProvider.Meta().(ClientSession).ISSession()
		vpcC := network.NewVPCClient(sess)
		foundvpc, err := vpcC.Get(rs.Primary.ID)

		if err != nil {
			return err
		}

		*vpc = foundvpc
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
