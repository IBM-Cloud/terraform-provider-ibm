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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMISIKEPolicy_basic(t *testing.T) {
	name := fmt.Sprintf("tfike-name-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: checkIKEPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISIKEPolicyConfig(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_is_ike_policy.example", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_ike_policy.example", "authentication_algorithm", "sha256"),
					resource.TestCheckResourceAttr(
						"ibm_is_ike_policy.example", "encryption_algorithm", "aes128"),
					resource.TestCheckResourceAttr(
						"ibm_is_ike_policy.example", "dh_group", "14"),
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
						"ibm_is_ike_policy.example", "authentication_algorithm", "sha384"),
					resource.TestCheckResourceAttr(
						"ibm_is_ike_policy.example", "encryption_algorithm", "aes256"),
					resource.TestCheckResourceAttr(
						"ibm_is_ike_policy.example", "dh_group", "15"),
					resource.TestCheckResourceAttr(
						"ibm_is_ike_policy.example", "ike_version", "2"),
				),
			},
		},
	})
}

func checkIKEPolicyDestroy(s *terraform.State) error {

	sess, _ := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
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
	return nil
}

func testAccCheckIBMISIKEPolicyExists(n, policy string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}

		sess, _ := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
		getikepoptions := &vpcv1.GetIkePolicyOptions{
			ID: &rs.Primary.ID,
		}
		ikePolicy, _, err := sess.GetIkePolicy(getikepoptions)
		if err != nil {
			return err
		}
		policy = *ikePolicy.ID
		return nil
	}
}

func testAccCheckIBMISIKEPolicyConfig(name string) string {
	return fmt.Sprintf(`
		resource "ibm_is_ike_policy" "example" {
			name = "%s"
			authentication_algorithm = "sha256"
			encryption_algorithm = "aes128"
			dh_group = 14
			ike_version = 1
		}
	`, name)
}

func testAccCheckIBMISIKEPolicyConfigUpdate(name string) string {
	return fmt.Sprintf(`
		resource "ibm_is_ike_policy" "example" {
			name = "%s"
			authentication_algorithm = "sha384"
			encryption_algorithm = "aes256"
			dh_group = 15
			ike_version = 2
			key_lifetime = 1800
		}
	`, name)
}
