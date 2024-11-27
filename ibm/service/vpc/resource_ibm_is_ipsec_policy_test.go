// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"errors"
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"

	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccIBMISIPSecPolicy_basic(t *testing.T) {
	name := fmt.Sprintf("tfipsecc-name-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: checkPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISIPSecPolicyConfig(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_is_ipsec_policy.example", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_ipsec_policy.example", "authentication_algorithm", "sha256"),
					resource.TestCheckResourceAttr(
						"ibm_is_ipsec_policy.example", "encryption_algorithm", "aes128"),
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
						"ibm_is_ipsec_policy.example", "authentication_algorithm", "sha512"),
					resource.TestCheckResourceAttr(
						"ibm_is_ipsec_policy.example", "encryption_algorithm", "aes256"),
					resource.TestCheckResourceAttr(
						"ibm_is_ipsec_policy.example", "pfs", "group_14"),
				),
			},
		},
	})
}

func checkPolicyDestroy(s *terraform.State) error {

	sess, _ := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_ipsec_policy" {
			continue
		}

		getipsecpoptions := &vpcv1.GetIpsecPolicyOptions{
			ID: &rs.Primary.ID,
		}
		_, _, err := sess.GetIpsecPolicy(getipsecpoptions)
		if err == nil {
			return fmt.Errorf("policy still exists: %s", rs.Primary.ID)
		}
	}

	return nil
}

func testAccCheckIBMISIpSecPolicyExists(n, policy string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}

		sess, _ := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
		getipsecpoptions := &vpcv1.GetIpsecPolicyOptions{
			ID: &rs.Primary.ID,
		}
		ipSecPolicy, _, err := sess.GetIpsecPolicy(getipsecpoptions)
		if err != nil {
			return err
		}
		policy = *ipSecPolicy.ID

		return nil
	}
}

func testAccCheckIBMISIPSecPolicyConfig(name string) string {
	return fmt.Sprintf(`
		resource "ibm_is_ipsec_policy" "example" {
			name = "%s"
			authentication_algorithm = "sha256"
			encryption_algorithm = "aes128"
			pfs = "disabled"
		}
	`, name)
}

func testAccCheckIBMISIPSecPolicyConfigUpdate(name string) string {
	return fmt.Sprintf(`
		resource "ibm_is_ipsec_policy" "example" {
			name = "%s"
			authentication_algorithm = "sha512"
			encryption_algorithm = "aes256"
			pfs = "group_2"
		}
	`, name)
}
