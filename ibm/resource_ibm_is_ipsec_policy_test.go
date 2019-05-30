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

func TestAccIBMISIPSecPolicy_basic(t *testing.T) {
	name := fmt.Sprintf("terraformipsecuat_create_step_name_%d", acctest.RandInt())
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: checkPolicyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMISIPSecPolicyConfig(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_is_ipsec_policy.example", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_ipsec_policy.example", "authentication_algorithm", "md5"),
					resource.TestCheckResourceAttr(
						"ibm_is_ipsec_policy.example", "encryption_algorithm", "3des"),
					resource.TestCheckResourceAttr(
						"ibm_is_ipsec_policy.example", "pfs", "disabled"),
				),
			},
			{
				Config: testAccCheckIBMISIPSecPolicyConfigUpdate(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_is_ipsec_policy.example", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_ipsec_policy.example", "authentication_algorithm", "sha1"),
					resource.TestCheckResourceAttr(
						"ibm_is_ipsec_policy.example", "encryption_algorithm", "aes128"),
					resource.TestCheckResourceAttr(
						"ibm_is_ipsec_policy.example", "pfs", "group_2"),
				),
			},
		},
	})
}

func checkPolicyDestroy(s *terraform.State) error {
	sess, _ := testAccProvider.Meta().(ClientSession).ISSession()

	vpnC := vpn.NewVpnClient(sess)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_ipsec_policy" {
			continue
		}

		_, err := vpnC.GetIpsecPolicy(rs.Primary.ID)

		if err == nil {
			return fmt.Errorf("policy still exists: %s", rs.Primary.ID)
		}
	}

	return nil
}

func testAccCheckIBMISIpSecPolicyExists(n string, policy **models.IpsecPolicy) resource.TestCheckFunc {
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
		ipSecPolicy, err := vpnC.GetIpsecPolicy(rs.Primary.ID)

		if err != nil {
			return err
		}

		*policy = ipSecPolicy
		return nil
	}
}

func testAccCheckIBMISIPSecPolicyConfig(name string) string {
	return fmt.Sprintf(`
		resource "ibm_is_ipsec_policy" "example" {
			name = "%s"
			authentication_algorithm = "md5"
			encryption_algorithm = "3des"
			pfs = "disabled"
		}
	`, name)
}

func testAccCheckIBMISIPSecPolicyConfigUpdate(name string) string {
	return fmt.Sprintf(`
		resource "ibm_is_ipsec_policy" "example" {
			name = "%s"
			authentication_algorithm = "sha1"
			encryption_algorithm = "aes128"
			pfs = "group_2"
		}
	`, name)
}
