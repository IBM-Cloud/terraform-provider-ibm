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

func TestAccIBMISSubnet_basic(t *testing.T) {
	var subnet *models.Subnet
	vpcname := fmt.Sprintf("terraformsubnetuat_vpc_%d", acctest.RandInt())
	name1 := fmt.Sprintf("terraformsubnetuat_create_step_name_%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMISSubnetDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISSubnetConfig(vpcname, name1, ISZoneName, ISCIDR),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISSubnetExists("ibm_is_subnet.testacc_subnet", &subnet),
					resource.TestCheckResourceAttr(
						"ibm_is_subnet.testacc_subnet", "name", name1),
					resource.TestCheckResourceAttr(
						"ibm_is_subnet.testacc_subnet", "zone", ISZoneName),
					resource.TestCheckResourceAttr(
						"ibm_is_subnet.testacc_subnet", "ipv4_cidr_block", ISCIDR),
				),
			},
		},
	})
}

func testAccCheckIBMISSubnetDestroy(s *terraform.State) error {
	sess, _ := testAccProvider.Meta().(ClientSession).ISSession()

	subnetC := network.NewSubnetClient(sess)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_subnet" {
			continue
		}

		_, err := subnetC.Get(rs.Primary.ID)

		if err == nil {
			return fmt.Errorf("subnet still exists: %s", rs.Primary.ID)
		}
	}

	return nil
}

func testAccCheckIBMISSubnetExists(n string, subnet **models.Subnet) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}

		sess, _ := testAccProvider.Meta().(ClientSession).ISSession()
		subnetC := network.NewSubnetClient(sess)
		foundsubnet, err := subnetC.Get(rs.Primary.ID)

		if err != nil {
			return err
		}

		*subnet = foundsubnet
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
	vpc = "${ibm_is_vpc.testacc_vpc.id}"
	zone = "%s"
	ipv4_cidr_block = "%s"
}`, vpcname, name, zone, cidr)

}
