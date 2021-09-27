// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMIamTrustedProfileClaimRuleDataSourceBasic(t *testing.T) {
	profileClaimRuleProfileID := fmt.Sprintf("tf_profile_id_%d", acctest.RandIntRange(10, 100))
	profileClaimRuleType := fmt.Sprintf("tf_type_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIamTrustedProfileClaimRuleDataSourceConfigBasic(profileClaimRuleProfileID, profileClaimRuleType),
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

func TestAccIBMIamTrustedProfileClaimRuleDataSourceAllArgs(t *testing.T) {
	profileClaimRuleProfileID := fmt.Sprintf("tf_profile_id_%d", acctest.RandIntRange(10, 100))
	profileClaimRuleType := fmt.Sprintf("tf_type_%d", acctest.RandIntRange(10, 100))
	profileClaimRuleName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	profileClaimRuleRealmName := fmt.Sprintf("tf_realm_name_%d", acctest.RandIntRange(10, 100))
	profileClaimRuleCrType := fmt.Sprintf("tf_cr_type_%d", acctest.RandIntRange(10, 100))
	profileClaimRuleExpiration := fmt.Sprintf("%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIamTrustedProfileClaimRuleDataSourceConfig(profileClaimRuleProfileID, profileClaimRuleType, profileClaimRuleName, profileClaimRuleRealmName, profileClaimRuleCrType, profileClaimRuleExpiration),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_claim_rule.iam_trusted_profile_claim_rule", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_claim_rule.iam_trusted_profile_claim_rule", "entity_tag"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_claim_rule.iam_trusted_profile_claim_rule", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_claim_rule.iam_trusted_profile_claim_rule", "modified_at"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_claim_rule.iam_trusted_profile_claim_rule", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_claim_rule.iam_trusted_profile_claim_rule", "type"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_claim_rule.iam_trusted_profile_claim_rule", "realm_name"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_claim_rule.iam_trusted_profile_claim_rule", "expiration"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_claim_rule.iam_trusted_profile_claim_rule", "cr_type"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_claim_rule.iam_trusted_profile_claim_rule", "conditions.#"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_claim_rule.iam_trusted_profile_claim_rule", "conditions.0.claim"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_claim_rule.iam_trusted_profile_claim_rule", "conditions.0.operator"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile_claim_rule.iam_trusted_profile_claim_rule", "conditions.0.value"),
				),
			},
		},
	})
}

func testAccCheckIBMIamTrustedProfileClaimRuleDataSourceConfigBasic(profileClaimRuleProfileID string, profileClaimRuleType string) string {
	return fmt.Sprintf(`
		resource "ibm_iam_trusted_profile_claim_rule" "iam_trusted_profile_claim_rule" {
			profile_id = "%s"
			type = "%s"
			conditions {
				claim = "claim"
				operator = "operator"
				value = "value"
			}
		}

		data "ibm_iam_trusted_profile_claim_rule" "iam_trusted_profile_claim_rule" {
			profile-id = ibm_iam_trusted_profile_claim_rule.iam_trusted_profile_claim_rule.profile_id
			rule-id = "rule-id"
		}
	`, profileClaimRuleProfileID, profileClaimRuleType)
}

func testAccCheckIBMIamTrustedProfileClaimRuleDataSourceConfig(profileClaimRuleProfileID string, profileClaimRuleType string, profileClaimRuleName string, profileClaimRuleRealmName string, profileClaimRuleCrType string, profileClaimRuleExpiration string) string {
	return fmt.Sprintf(`
		resource "ibm_iam_trusted_profile_claim_rule" "iam_trusted_profile_claim_rule" {
			profile_id = "%s"
			type = "%s"
			conditions {
				claim = "claim"
				operator = "operator"
				value = "value"
			}
			context {
				transaction_id = "transaction_id"
				operation = "operation"
				user_agent = "user_agent"
				url = "url"
				instance_id = "instance_id"
				thread_id = "thread_id"
				host = "host"
				start_time = "start_time"
				end_time = "end_time"
				elapsed_time = "elapsed_time"
				cluster_name = "cluster_name"
			}
			name = "%s"
			realm_name = "%s"
			cr_type = "%s"
			expiration = %s
		}

		data "ibm_iam_trusted_profile_claim_rule" "iam_trusted_profile_claim_rule" {
			profile-id = ibm_iam_trusted_profile_claim_rule.iam_trusted_profile_claim_rule.profile_id
			rule-id = "rule-id"
		}
	`, profileClaimRuleProfileID, profileClaimRuleType, profileClaimRuleName, profileClaimRuleRealmName, profileClaimRuleCrType, profileClaimRuleExpiration)
}
