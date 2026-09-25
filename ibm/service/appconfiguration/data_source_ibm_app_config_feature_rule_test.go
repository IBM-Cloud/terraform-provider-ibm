// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package appconfiguration_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIbmAppConfigFeatureRuleDataSource(t *testing.T) {
	instanceName := fmt.Sprintf("tf_app_config_test_%d", acctest.RandIntRange(10, 100))
	environmentID := "dev"
	featureID := fmt.Sprintf("tf_feature_%d", acctest.RandIntRange(10, 100))
	segmentID := fmt.Sprintf("tf_segment_%d", acctest.RandIntRange(10, 100))
	ruleID := fmt.Sprintf("tf-rule-%d", acctest.RandIntRange(10, 100))
	ruleName := fmt.Sprintf("tf_rule_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmAppConfigFeatureRuleDataSourceConfig(instanceName, environmentID, featureID, segmentID, ruleID, ruleName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_app_config_feature_rule.test_ds", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_feature_rule.test_ds", "rule_id"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_feature_rule.test_ds", "rule_name"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_feature_rule.test_ds", "value"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_feature_rule.test_ds", "order"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_feature_rule.test_ds", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_feature_rule.test_ds", "rules.#"),
				),
			},
		},
	})
}

func testAccCheckIbmAppConfigFeatureRuleDataSourceConfig(instanceName, environmentID, featureID, segmentID, ruleID, ruleName string) string {
	return fmt.Sprintf(`
		data "ibm_resource_group" "default" {
			is_default = true
		}

		resource "ibm_resource_instance" "app_config_fr_ds_test" {
			name              = "%s"
			location          = "us-south"
			service           = "apprapp"
			plan              = "lite"
			resource_group_id = data.ibm_resource_group.default.id
		}

		resource "ibm_app_config_segment" "fr_ds_segment" {
			guid       = ibm_resource_instance.app_config_fr_ds_test.guid
			segment_id = "%s"
			name       = "TF Segment for Rule DS"
			rules {
				attribute_name = "email"
				operator       = "endsWith"
				values         = ["@ibm.com"]
			}
		}

		resource "ibm_app_config_feature" "fr_ds_feature" {
			guid           = ibm_resource_instance.app_config_fr_ds_test.guid
			environment_id = "%s"
			feature_id     = "%s"
			name           = "TF Feature for Rule DS"
			type           = "BOOLEAN"
			enabled_value  = true
			disabled_value = false
			enabled        = false
			lifecycle {
				ignore_changes = [rollout_type, rollout_percentage, segment_rules]
			}
		}

		resource "ibm_app_config_feature_rule" "fr_ds_rule" {
			guid           = ibm_resource_instance.app_config_fr_ds_test.guid
			environment_id = "%s"
			feature_id     = ibm_app_config_feature.fr_ds_feature.feature_id
			rule_id        = "%s"
			rule_name      = "%s"
			value          = "true"
			rules {
				segments = [ibm_app_config_segment.fr_ds_segment.segment_id]
			}
			lifecycle {
				ignore_changes = [rollout_type, rollout_percentage]
			}
		}

		data "ibm_app_config_feature_rule" "test_ds" {
			guid           = ibm_resource_instance.app_config_fr_ds_test.guid
			environment_id = ibm_app_config_feature_rule.fr_ds_rule.environment_id
			feature_id     = ibm_app_config_feature_rule.fr_ds_rule.feature_id
			rule_id        = ibm_app_config_feature_rule.fr_ds_rule.rule_id
		}`, instanceName, segmentID, environmentID, featureID, environmentID, ruleID, ruleName)
}
