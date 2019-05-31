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

func TestAccIBMISIKEPolicy_basic(t *testing.T) {
	name := fmt.Sprintf("terraformIkeuat_create_step_name_%d", acctest.RandInt())
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: checkIKEPolicyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMISIKEPolicyConfig(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_is_ike_policy.example", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_ike_policy.example", "authentication_algorithm", "md5"),
					resource.TestCheckResourceAttr(
						"ibm_is_ike_policy.example", "encryption_algorithm", "3des"),
					resource.TestCheckResourceAttr(
						"ibm_is_ike_policy.example", "dh_group", "2"),
					resource.TestCheckResourceAttr(
						"ibm_is_ike_policy.example", "ike_version", "1"),
				),
			},
			{
				Config: testAccCheckIBMISIKEPolicyConfigUpdate(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_is_ike_policy.example", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_ike_policy.example", "authentication_algorithm", "sha1"),
					resource.TestCheckResourceAttr(
						"ibm_is_ike_policy.example", "encryption_algorithm", "aes128"),
					resource.TestCheckResourceAttr(
						"ibm_is_ike_policy.example", "dh_group", "5"),
					resource.TestCheckResourceAttr(
						"ibm_is_ike_policy.example", "ike_version", "2"),
				),
			},
		},
	})
}

func checkIKEPolicyDestroy(s *terraform.State) error {
	sess, _ := testAccProvider.Meta().(ClientSession).ISSession()

	vpnC := vpn.NewVpnClient(sess)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_ike_policy" {
			continue
		}

		_, err := vpnC.GetIkePolicy(rs.Primary.ID)

		if err == nil {
			return fmt.Errorf("policy still exists: %s", rs.Primary.ID)
		}
	}

	return nil
}

func testAccCheckIBMISIKEPolicyExists(n string, policy **models.IKEPolicy) resource.TestCheckFunc {
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
		vpnC := vpn.NewVpnClient(sess)
		ikePolicy, err := vpnC.GetIkePolicy(rs.Primary.ID)

		if err != nil {
			return err
		}

		*policy = ikePolicy
		return nil
	}
}

func testAccCheckIBMISIKEPolicyConfig(name string) string {
	return fmt.Sprintf(`
		resource "ibm_is_ike_policy" "example" {
			name = "%s"
			authentication_algorithm = "md5"
			encryption_algorithm = "3des"
			dh_group = 2
			ike_version = 1
		}
	`, name)
}

func testAccCheckIBMISIKEPolicyConfigUpdate(name string) string {
	return fmt.Sprintf(`
		resource "ibm_is_ike_policy" "example" {
			name = "%s"
			authentication_algorithm = "sha1"
			encryption_algorithm = "aes128"
			dh_group = 5
			ike_version = 2
			key_lifetime = 600
		}
	`, name)
}
