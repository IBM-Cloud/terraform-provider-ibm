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

func TestAccIBMISIPSecPolicy_basic(t *testing.T) {
	name := fmt.Sprintf("tfipsecc-name-%d", acctest.RandIntRange(10, 100))
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
						"ibm_is_ipsec_policy.example", "encryption_algorithm", "triple_des"),
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
	userDetails, _ := testAccProvider.Meta().(ClientSession).BluemixUserDetails()

	if userDetails.generation == 1 {
		sess, _ := testAccProvider.Meta().(ClientSession).VpcClassicV1API()
		for _, rs := range s.RootModule().Resources {
			if rs.Type != "ibm_is_ipsec_policy" {
				continue
			}

			getipsecpoptions := &vpcclassicv1.GetIpsecPolicyOptions{
				ID: &rs.Primary.ID,
			}
			_, _, err := sess.GetIpsecPolicy(getipsecpoptions)
			if err == nil {
				return fmt.Errorf("policy still exists: %s", rs.Primary.ID)
			}
		}
	} else {
		sess, _ := testAccProvider.Meta().(ClientSession).VpcV1API()
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
	}
	return nil
}

func testAccCheckIBMISIpSecPolicyExists(n, policy string) resource.TestCheckFunc {
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
			getipsecpoptions := &vpcclassicv1.GetIpsecPolicyOptions{
				ID: &rs.Primary.ID,
			}
			ipSecPolicy, _, err := sess.GetIpsecPolicy(getipsecpoptions)
			if err != nil {
				return err
			}
			policy = *ipSecPolicy.ID
		} else {
			sess, _ := testAccProvider.Meta().(ClientSession).VpcV1API()
			getipsecpoptions := &vpcv1.GetIpsecPolicyOptions{
				ID: &rs.Primary.ID,
			}
			ipSecPolicy, _, err := sess.GetIpsecPolicy(getipsecpoptions)
			if err != nil {
				return err
			}
			policy = *ipSecPolicy.ID
		}
		return nil
	}
}

func testAccCheckIBMISIPSecPolicyConfig(name string) string {
	return fmt.Sprintf(`
		resource "ibm_is_ipsec_policy" "example" {
			name = "%s"
			authentication_algorithm = "md5"
			encryption_algorithm = "triple_des"
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
