// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package appconfiguration_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/appconfiguration-go-admin-sdk/appconfigurationv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIbmAppConfigFeatureRuleBasic(t *testing.T) {
	var conf appconfigurationv1.FeatureSegmentRuleWithRuleID

	instanceName := fmt.Sprintf("tf_app_config_test_%d", acctest.RandIntRange(10, 100))
	environmentID := "dev"
	featureID := fmt.Sprintf("tf_feature_%d", acctest.RandIntRange(10, 100))
	segmentID := fmt.Sprintf("tf_segment_%d", acctest.RandIntRange(10, 100))
	ruleID := fmt.Sprintf("tf-rule-%d", acctest.RandIntRange(10, 100))
	ruleName := fmt.Sprintf("tf_rule_name_%d", acctest.RandIntRange(10, 100))
	ruleNameUpdate := fmt.Sprintf("tf_rule_name_updated_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmAppConfigFeatureRuleDestroy,
		Steps: []resource.TestStep{
			{
				// Step 1: CREATE
				Config: testAccCheckIbmAppConfigFeatureRuleConfig(instanceName, environmentID, featureID, segmentID, ruleID, ruleName, "true"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmAppConfigFeatureRuleExists("ibm_app_config_feature_rule.test", conf),
					resource.TestCheckResourceAttrSet("ibm_app_config_feature_rule.test", "id"),
					resource.TestCheckResourceAttrSet("ibm_app_config_feature_rule.test", "rule_id"),
					resource.TestCheckResourceAttrSet("ibm_app_config_feature_rule.test", "href"),
					resource.TestCheckResourceAttrSet("ibm_app_config_feature_rule.test", "order"),
					resource.TestCheckResourceAttr("ibm_app_config_feature_rule.test", "rule_name", ruleName),
					resource.TestCheckResourceAttr("ibm_app_config_feature_rule.test", "value", "true"),
				),
			},
			{
				// Step 2: UPDATE — change rule_name
				Config: testAccCheckIbmAppConfigFeatureRuleConfig(instanceName, environmentID, featureID, segmentID, ruleID, ruleNameUpdate, "true"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_app_config_feature_rule.test", "rule_name", ruleNameUpdate),
				),
			},
		},
	})
}

func testAccCheckIbmAppConfigFeatureRuleConfig(instanceName, environmentID, featureID, segmentID, ruleID, ruleName, value string) string {
	return fmt.Sprintf(`
		data "ibm_resource_group" "default" {
			is_default = true
		}

		resource "ibm_resource_instance" "app_config_feature_rule_test" {
			name              = "%s"
			location          = "us-south"
			service           = "apprapp"
			plan              = "lite"
			resource_group_id = data.ibm_resource_group.default.id
		}

		resource "ibm_app_config_segment" "feature_rule_segment" {
			guid       = ibm_resource_instance.app_config_feature_rule_test.guid
			segment_id = "%s"
			name       = "TF Test Segment for Rule"
			rules {
				attribute_name = "email"
				operator       = "endsWith"
				values         = ["@ibm.com"]
			}
		}

		resource "ibm_app_config_feature" "feature_rule_feature" {
			guid           = ibm_resource_instance.app_config_feature_rule_test.guid
			environment_id = "%s"
			feature_id     = "%s"
			name           = "TF Test Feature for Rule"
			type           = "BOOLEAN"
			enabled_value  = true
			disabled_value = false
			enabled        = false
			lifecycle {
				ignore_changes = [rollout_type, rollout_percentage, segment_rules]
			}
		}

		resource "ibm_app_config_feature_rule" "test" {
			guid           = ibm_resource_instance.app_config_feature_rule_test.guid
			environment_id = "%s"
			feature_id     = ibm_app_config_feature.feature_rule_feature.feature_id
			rule_id        = "%s"
			rule_name      = "%s"
			value          = "%s"
			rules {
				segments = [ibm_app_config_segment.feature_rule_segment.segment_id]
			}
			lifecycle {
				ignore_changes = [rollout_type, rollout_percentage]
			}
		}`, instanceName, segmentID, environmentID, featureID, environmentID, ruleID, ruleName, value)
}

func testAccCheckIbmAppConfigFeatureRuleExists(n string, obj appconfigurationv1.FeatureSegmentRuleWithRuleID) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return flex.FmtErrorf("Not found: %s", n)
		}
		// ID format: guid/environment_id/feature_id/rule_id
		parts, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return flex.FmtErrorf("%s", err)
		}
		if len(parts) < 4 {
			return flex.FmtErrorf("invalid ID format, expected guid/environment_id/feature_id/rule_id")
		}

		appconfigClient, err := getAppConfigClient(acc.TestAccProvider.Meta(), parts[0])
		if err != nil {
			return flex.FmtErrorf("%s", err)
		}

		options := &appconfigurationv1.GetFeatureRuleOptions{}
		options.SetEnvironmentID(parts[1])
		options.SetFeatureID(parts[2])
		options.SetRuleID(parts[3])

		result, _, err := appconfigClient.GetFeatureRule(options)
		if err != nil {
			return flex.FmtErrorf("%s", err)
		}

		obj = *result
		return nil
	}
}

func testAccCheckIbmAppConfigFeatureRuleDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_app_config_feature_rule" {
			continue
		}
		parts, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return flex.FmtErrorf("%s", err)
		}
		if len(parts) < 4 {
			continue
		}

		appconfigClient, err := getAppConfigClient(acc.TestAccProvider.Meta(), parts[0])
		if err != nil {
			return flex.FmtErrorf("%s", err)
		}

		options := &appconfigurationv1.GetFeatureRuleOptions{}
		options.SetEnvironmentID(parts[1])
		options.SetFeatureID(parts[2])
		options.SetRuleID(parts[3])

		_, response, err := appconfigClient.GetFeatureRule(options)
		if err == nil {
			return flex.FmtErrorf("FeatureRule still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return flex.FmtErrorf("[ERROR] Error checking for FeatureRule (%s) destroyed: %s", rs.Primary.ID, err)
		}
	}
	return nil
}
