// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/IBM/platform-services-go-sdk/iamidentityv1"
)

func TestAccIBMIamTrustedProfileClaimRuleBasic(t *testing.T) {
	var conf iamidentityv1.ProfileClaimRule
	profileID := fmt.Sprintf("tf_profile_id_%d", acctest.RandIntRange(10, 100))
	typeVar := fmt.Sprintf("tf_type_%d", acctest.RandIntRange(10, 100))
	typeVarUpdate := fmt.Sprintf("tf_type_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMIamTrustedProfileClaimRuleDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIamTrustedProfileClaimRuleConfigBasic(profileID, typeVar),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIamTrustedProfileClaimRuleExists("ibm_iam_trusted_profile_claim_rule.iam_trusted_profile_claim_rule", conf),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_claim_rule.iam_trusted_profile_claim_rule", "profile_id", profileID),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_claim_rule.iam_trusted_profile_claim_rule", "type", typeVar),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMIamTrustedProfileClaimRuleConfigBasic(profileID, typeVarUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_claim_rule.iam_trusted_profile_claim_rule", "profile_id", profileID),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_claim_rule.iam_trusted_profile_claim_rule", "type", typeVarUpdate),
				),
			},
		},
	})
}

func TestAccIBMIamTrustedProfileClaimRuleAllArgs(t *testing.T) {
	var conf iamidentityv1.ProfileClaimRule
	profileID := fmt.Sprintf("tf_profile_id_%d", acctest.RandIntRange(10, 100))
	typeVar := fmt.Sprintf("tf_type_%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	realmName := fmt.Sprintf("tf_realm_name_%d", acctest.RandIntRange(10, 100))
	crType := fmt.Sprintf("tf_cr_type_%d", acctest.RandIntRange(10, 100))
	expiration := fmt.Sprintf("%d", acctest.RandIntRange(10, 100))
	typeVarUpdate := fmt.Sprintf("tf_type_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	realmNameUpdate := fmt.Sprintf("tf_realm_name_%d", acctest.RandIntRange(10, 100))
	crTypeUpdate := fmt.Sprintf("tf_cr_type_%d", acctest.RandIntRange(10, 100))
	expirationUpdate := fmt.Sprintf("%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMIamTrustedProfileClaimRuleDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIamTrustedProfileClaimRuleConfig(profileID, typeVar, name, realmName, crType, expiration),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIamTrustedProfileClaimRuleExists("ibm_iam_trusted_profile_claim_rule.iam_trusted_profile_claim_rule", conf),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_claim_rule.iam_trusted_profile_claim_rule", "profile_id", profileID),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_claim_rule.iam_trusted_profile_claim_rule", "type", typeVar),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_claim_rule.iam_trusted_profile_claim_rule", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_claim_rule.iam_trusted_profile_claim_rule", "realm_name", realmName),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_claim_rule.iam_trusted_profile_claim_rule", "cr_type", crType),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_claim_rule.iam_trusted_profile_claim_rule", "expiration", expiration),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMIamTrustedProfileClaimRuleConfig(profileID, typeVarUpdate, nameUpdate, realmNameUpdate, crTypeUpdate, expirationUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_claim_rule.iam_trusted_profile_claim_rule", "profile_id", profileID),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_claim_rule.iam_trusted_profile_claim_rule", "type", typeVarUpdate),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_claim_rule.iam_trusted_profile_claim_rule", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_claim_rule.iam_trusted_profile_claim_rule", "realm_name", realmNameUpdate),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_claim_rule.iam_trusted_profile_claim_rule", "cr_type", crTypeUpdate),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_claim_rule.iam_trusted_profile_claim_rule", "expiration", expirationUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_iam_trusted_profile_claim_rule.iam_trusted_profile_claim_rule",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMIamTrustedProfileClaimRuleConfigBasic(profileID string, typeVar string) string {
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
	`, profileID, typeVar)
}

func testAccCheckIBMIamTrustedProfileClaimRuleConfig(profileID string, typeVar string, name string, realmName string, crType string, expiration string) string {
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
	`, profileID, typeVar, name, realmName, crType, expiration)
}

func testAccCheckIBMIamTrustedProfileClaimRuleExists(n string, obj iamidentityv1.ProfileClaimRule) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		iamIdentityClient, err := testAccProvider.Meta().(ClientSession).IAMIdentityV1API()
		if err != nil {
			return err
		}

		getClaimRuleOptions := &iamidentityv1.GetClaimRuleOptions{}

		parts, err := sepIdParts(rs.Primary.ID, "/")
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
	iamIdentityClient, err := testAccProvider.Meta().(ClientSession).IAMIdentityV1API()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_iam_trusted_profile_claim_rule" {
			continue
		}

		getClaimRuleOptions := &iamidentityv1.GetClaimRuleOptions{}

		parts, err := sepIdParts(rs.Primary.ID, "/")
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
			return fmt.Errorf("Error checking for iam_trusted_profile_claim_rule (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
