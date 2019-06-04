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

func TestAccIBMISSecurityGroup_basic(t *testing.T) {
	var securityGroup *models.SecurityGroup

	vpcname := fmt.Sprintf("terraformsecurityuat_vpc_%d", acctest.RandInt())
	name1 := fmt.Sprintf("terraformsecurityuat_create_step_name_%d", acctest.RandInt())
	//name2 := fmt.Sprintf("terraformsecurityuat_update_step_name_%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMISSecurityGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISsecurityGroupConfig(vpcname, name1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISSecurityGroupExists("ibm_is_security_group.testacc_security_group", &securityGroup),
					resource.TestCheckResourceAttr(
						"ibm_is_security_group.testacc_security_group", "name", name1),
				),
			},
		},
	})
}

func testAccCheckIBMISSecurityGroupDestroy(s *terraform.State) error {
	sess, _ := testAccProvider.Meta().(ClientSession).ISSession()

	securityGroupC := network.NewSecurityGroupClient(sess)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_security_group" {
			continue
		}

		_, err := securityGroupC.Get(rs.Primary.ID)

		if err == nil {
			return fmt.Errorf("securitygroup still exists: %s", rs.Primary.ID)
		}
	}

	return nil
}

func testAccCheckIBMISSecurityGroupExists(n string, securityGroup **models.SecurityGroup) resource.TestCheckFunc {
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
		securityGroupC := network.NewSecurityGroupClient(sess)
		foundsecurityGroup, err := securityGroupC.Get(rs.Primary.ID)

		if err != nil {
			return err
		}

		*securityGroup = foundsecurityGroup
		return nil
	}
}

func testAccCheckIBMISsecurityGroupConfig(vpcname, name string) string {
	return fmt.Sprintf(`
resource "ibm_is_vpc" "testacc_vpc" {
	name = "%s"
}

resource "ibm_is_security_group" "testacc_security_group" {
	name = "%s"
	vpc = "${ibm_is_vpc.testacc_vpc.id}"
}`, vpcname, name)

}
