/* IBM Confidential
*  Object Code Only Source Materials
*  5747-SM3
*  (c) Copyright IBM Corp. 2017,2021
*
*  The source code for this program is not published or otherwise divested
*  of its trade secrets, irrespective of what has been deposited with the
*  U.S. Copyright Office. */

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

func TestAccIBMISSecurityGroup_basic(t *testing.T) {
	var securityGroup string

	vpcname := fmt.Sprintf("tfsg-vpc-%d", acctest.RandIntRange(10, 100))
	name1 := fmt.Sprintf("tfsg-createname-%d", acctest.RandIntRange(10, 100))
	//name2 := fmt.Sprintf("tfsg-updatename-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMISSecurityGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISsecurityGroupConfig(vpcname, name1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISSecurityGroupExists("ibm_is_security_group.testacc_security_group", securityGroup),
					resource.TestCheckResourceAttr(
						"ibm_is_security_group.testacc_security_group", "name", name1),
				),
			},
		},
	})
}

func testAccCheckIBMISSecurityGroupDestroy(s *terraform.State) error {
	userDetails, _ := testAccProvider.Meta().(ClientSession).BluemixUserDetails()

	if userDetails.generation == 1 {
		sess, _ := testAccProvider.Meta().(ClientSession).VpcClassicV1API()
		for _, rs := range s.RootModule().Resources {
			if rs.Type != "ibm_is_security_group" {
				continue
			}

			getsgoptions := &vpcclassicv1.GetSecurityGroupOptions{
				ID: &rs.Primary.ID,
			}
			_, _, err := sess.GetSecurityGroup(getsgoptions)
			if err == nil {
				return fmt.Errorf("securitygroup still exists: %s", rs.Primary.ID)
			}
		}
	} else {
		sess, _ := testAccProvider.Meta().(ClientSession).VpcV1API()
		for _, rs := range s.RootModule().Resources {
			if rs.Type != "ibm_is_security_group" {
				continue
			}

			getsgoptions := &vpcv1.GetSecurityGroupOptions{
				ID: &rs.Primary.ID,
			}
			_, _, err := sess.GetSecurityGroup(getsgoptions)

			if err == nil {
				return fmt.Errorf("securitygroup still exists: %s", rs.Primary.ID)
			}
		}
	}

	return nil
}

func testAccCheckIBMISSecurityGroupExists(n, securityGroupID string) resource.TestCheckFunc {
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
			getsgoptions := &vpcclassicv1.GetSecurityGroupOptions{
				ID: &rs.Primary.ID,
			}
			foundsecurityGroup, _, err := sess.GetSecurityGroup(getsgoptions)
			if err != nil {
				return err
			}
			securityGroupID = *foundsecurityGroup.ID
		} else {
			sess, _ := testAccProvider.Meta().(ClientSession).VpcV1API()
			getsgoptions := &vpcv1.GetSecurityGroupOptions{
				ID: &rs.Primary.ID,
			}
			foundsecurityGroup, _, err := sess.GetSecurityGroup(getsgoptions)
			if err != nil {
				return err
			}
			securityGroupID = *foundsecurityGroup.ID
		}
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
