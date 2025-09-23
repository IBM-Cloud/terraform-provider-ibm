// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package iamidentity_test

import (
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMIAMTrustedProfileClaimRuleDataSourceBasic(t *testing.T) {
	profileClaimRuleProfileID := fmt.Sprintf("tf_profile_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckIAMTrustedProfile(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIamTrustedProfileClaimRuleDataSourceConfigBasic(profileClaimRuleProfileID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_claim_rule.iam_trusted_profile_claim_rule", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_claim_rule.iam_trusted_profile_claim_rule", "entity_tag"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_claim_rule.iam_trusted_profile_claim_rule", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_claim_rule.iam_trusted_profile_claim_rule", "type"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_claim_rule.iam_trusted_profile_claim_rule", "expiration"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_claim_rule.iam_trusted_profile_claim_rule", "conditions.#"),
				),
			},
		},
	})
}

func TestAccIBMIAMTrustedProfileClaimRuleDataSourceAllArgs(t *testing.T) {
	profileClaimRuleProfileID := fmt.Sprintf("tf_profile_id_%d", acctest.RandIntRange(10, 100))
	profileClaimRuleName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckIAMTrustedProfile(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIamTrustedProfileClaimRuleDataSourceConfig(profileClaimRuleProfileID, profileClaimRuleName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_claim_rule.iam_trusted_profile_claim_rule", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_claim_rule.iam_trusted_profile_claim_rule", "entity_tag"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_claim_rule.iam_trusted_profile_claim_rule", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_claim_rule.iam_trusted_profile_claim_rule", "modified_at"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_claim_rule.iam_trusted_profile_claim_rule", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_claim_rule.iam_trusted_profile_claim_rule", "type"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_claim_rule.iam_trusted_profile_claim_rule", "conditions.#"),
				),
			},
		},
	})
}

func testAccCheckIBMIamTrustedProfileClaimRuleDataSourceConfigBasic(profileClaimRuleProfileID string) string {
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
			name = "%[1]s"
			realm_name = "%s"
			expiration = 43200
			conditions {
				claim = "blueGroups"
				operator = "NOT_EQUALS_IGNORE_CASE"
				value = "\"cloud-docs-dev\""
			}
		}
		data "ibm_iam_trusted_profile_claim_rule" "iam_trusted_profile_claim_rule" {
			profile_id = ibm_iam_trusted_profile_claim_rule.iam_trusted_profile_claim_rule.profile_id
			rule_id = ibm_iam_trusted_profile_claim_rule.iam_trusted_profile_claim_rule.rule_id
		}
	`, profileClaimRuleProfileID, acc.RealmName)
}

func testAccCheckIBMIamTrustedProfileClaimRuleDataSourceConfig(profileClaimRuleProfileID string, profileClaimRuleName string) string {
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

		data "ibm_iam_trusted_profile_claim_rule" "iam_trusted_profile_claim_rule" {
			profile_id = ibm_iam_trusted_profile_claim_rule.iam_trusted_profile_claim_rule.profile_id
			rule_id = ibm_iam_trusted_profile_claim_rule.iam_trusted_profile_claim_rule.rule_id
		}
	`, profileClaimRuleProfileID, profileClaimRuleName)
}
