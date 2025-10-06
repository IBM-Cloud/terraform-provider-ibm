// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package iamidentity_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/IBM/platform-services-go-sdk/iamidentityv1"
)

func TestAccIBMIAMTrustedProfileClaimRuleBasic(t *testing.T) {
	var conf iamidentityv1.ProfileClaimRule
	profileName := fmt.Sprintf("tf_profile_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIamTrustedProfileClaimRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIamTrustedProfileClaimRuleConfigBasic(profileName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIamTrustedProfileClaimRuleExists("ibm_iam_trusted_profile_claim_rule.iam_trusted_profile_claim_rule", conf),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_claim_rule.iam_trusted_profile_claim_rule", "type", "Profile-SAML"),
				),
			},
		},
	})
}

func TestAccIBMIAMTrustedProfileClaimRuleAllArgs(t *testing.T) {
	var conf iamidentityv1.ProfileClaimRule
	profileName := fmt.Sprintf("tf_profile_%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIamTrustedProfileClaimRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIamTrustedProfileClaimRuleConfig(profileName, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIamTrustedProfileClaimRuleExists("ibm_iam_trusted_profile_claim_rule.iam_trusted_profile_claim_rule", conf),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_claim_rule.iam_trusted_profile_claim_rule", "type", "Profile-CR"),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_claim_rule.iam_trusted_profile_claim_rule", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_claim_rule.iam_trusted_profile_claim_rule", "cr_type", "IKS_SA"),
				),
			},
			{
				ResourceName:      "ibm_iam_trusted_profile_claim_rule.iam_trusted_profile_claim_rule",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMIamTrustedProfileClaimRuleConfigBasic(profileName string) string {
	return fmt.Sprintf(`
		resource "ibm_iam_trusted_profile" "iam_trusted_profile" {
			name = "%s"
			lifecycle {
              ignore_changes = [history]
            }
		}
		resource "ibm_iam_trusted_profile_claim_rule" "iam_trusted_profile_claim_rule" {
			profile_id = ibm_iam_trusted_profile.iam_trusted_profile.id 
			type = "Profile-SAML"
			name = "%s"
			realm_name = "testString"
			expiration = 43200
			conditions {
				claim = "blueGroups"
				operator = "NOT_EQUALS_IGNORE_CASE"
				value = "\"cloud-docs-dev\""
			}
		}
	`, profileName, profileName)
}

func testAccCheckIBMIamTrustedProfileClaimRuleConfig(profileName string, name string) string {
	return fmt.Sprintf(`
		resource "ibm_iam_trusted_profile" "iam_trusted_profile" {
			name = "%s"
			lifecycle {
              ignore_changes = [history]
            }
		}
		resource "ibm_iam_trusted_profile_claim_rule" "iam_trusted_profile_claim_rule" {
			profile_id = ibm_iam_trusted_profile.iam_trusted_profile.id
			type = "Profile-CR"
			conditions {
				claim = "blueGroups"
				operator = "CONTAINS"
				value = "\"cloud-docs-dev\""
			}
			name = "%s"
			cr_type = "IKS_SA"
		}
	`, profileName, name)
}

func testAccCheckIBMIamTrustedProfileClaimRuleExists(n string, obj iamidentityv1.ProfileClaimRule) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		iamIdentityClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).IAMIdentityV1API()
		if err != nil {
			return err
		}

		getClaimRuleOptions := &iamidentityv1.GetClaimRuleOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getClaimRuleOptions.SetProfileID(parts[0])
		getClaimRuleOptions.SetRuleID(parts[1])

		profileClaimRule, _, err := iamIdentityClient.GetClaimRule(getClaimRuleOptions)
		if err != nil {
			return err
		}

		obj = *profileClaimRule
		return nil
	}
}

func testAccCheckIBMIamTrustedProfileClaimRuleDestroy(s *terraform.State) error {
	iamIdentityClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).IAMIdentityV1API()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_iam_trusted_profile_claim_rule" {
			continue
		}

		getClaimRuleOptions := &iamidentityv1.GetClaimRuleOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getClaimRuleOptions.SetProfileID(parts[0])
		getClaimRuleOptions.SetRuleID(parts[1])

		// Try to find the key
		_, response, err := iamIdentityClient.GetClaimRule(getClaimRuleOptions)

		if err == nil {
			return fmt.Errorf("iam_trusted_profile_claim_rule still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("[ERROR] Error checking for iam_trusted_profile_claim_rule (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
