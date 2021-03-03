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

func TestAccIBMISIKEPolicy_basic(t *testing.T) {
	name := fmt.Sprintf("tfike-name-%d", acctest.RandIntRange(10, 100))
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
						"ibm_is_ike_policy.example", "encryption_algorithm", "triple_des"),
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
	userDetails, _ := testAccProvider.Meta().(ClientSession).BluemixUserDetails()

	if userDetails.generation == 1 {
		sess, _ := testAccProvider.Meta().(ClientSession).VpcClassicV1API()
		for _, rs := range s.RootModule().Resources {
			if rs.Type != "ibm_is_ike_policy" {
				continue
			}

			getikepoptions := &vpcclassicv1.GetIkePolicyOptions{
				ID: &rs.Primary.ID,
			}
			_, _, err := sess.GetIkePolicy(getikepoptions)
			if err == nil {
				return fmt.Errorf("policy still exists: %s", rs.Primary.ID)
			}
		}
	} else {
		sess, _ := testAccProvider.Meta().(ClientSession).VpcV1API()
		for _, rs := range s.RootModule().Resources {
			if rs.Type != "ibm_is_ike_policy" {
				continue
			}

			getikepoptions := &vpcv1.GetIkePolicyOptions{
				ID: &rs.Primary.ID,
			}
			_, _, err := sess.GetIkePolicy(getikepoptions)
			if err == nil {
				return fmt.Errorf("policy still exists: %s", rs.Primary.ID)
			}
		}
	}

	return nil
}

func testAccCheckIBMISIKEPolicyExists(n, policy string) resource.TestCheckFunc {
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
			getikepoptions := &vpcclassicv1.GetIkePolicyOptions{
				ID: &rs.Primary.ID,
			}
			ikePolicy, _, err := sess.GetIkePolicy(getikepoptions)
			if err != nil {
				return err
			}
			policy = *ikePolicy.ID
		} else {
			sess, _ := testAccProvider.Meta().(ClientSession).VpcV1API()
			getikepoptions := &vpcv1.GetIkePolicyOptions{
				ID: &rs.Primary.ID,
			}
			ikePolicy, _, err := sess.GetIkePolicy(getikepoptions)
			if err != nil {
				return err
			}
			policy = *ikePolicy.ID
		}
		return nil
	}
}

func testAccCheckIBMISIKEPolicyConfig(name string) string {
	return fmt.Sprintf(`
		resource "ibm_is_ike_policy" "example" {
			name = "%s"
			authentication_algorithm = "md5"
			encryption_algorithm = "triple_des"
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
			key_lifetime = 1800
		}
	`, name)
}
