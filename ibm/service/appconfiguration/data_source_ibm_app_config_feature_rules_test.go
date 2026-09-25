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

func TestAccIbmAppConfigFeatureRulesDataSource(t *testing.T) {
	instanceName := fmt.Sprintf("tf_app_config_test_%d", acctest.RandIntRange(10, 100))
	environmentID := "dev"
	featureID := fmt.Sprintf("tf_feature_%d", acctest.RandIntRange(10, 100))
	segmentID := fmt.Sprintf("tf_segment_%d", acctest.RandIntRange(10, 100))
	ruleID := fmt.Sprintf("tf-rule-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmAppConfigFeatureRulesDataSourceConfig(instanceName, environmentID, featureID, segmentID, ruleID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_app_config_feature_rules.test_ds", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_feature_rules.test_ds", "total_count"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_feature_rules.test_ds", "rules.#"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_feature_rules.test_ds", "rules.0.rule_id"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_feature_rules.test_ds", "rules.0.value"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_feature_rules.test_ds", "rules.0.order"),
				),
			},
		},
	})
}

func testAccCheckIbmAppConfigFeatureRulesDataSourceConfig(instanceName, environmentID, featureID, segmentID, ruleID string) string {
	return fmt.Sprintf(`
		data "ibm_resource_group" "default" {
			is_default = true
		}

		resource "ibm_resource_instance" "app_config_frs_ds_test" {
			name              = "%s"
			location          = "us-south"
			service           = "apprapp"
			plan              = "lite"
			resource_group_id = data.ibm_resource_group.default.id
		}

		resource "ibm_app_config_segment" "frs_ds_segment" {
			guid       = ibm_resource_instance.app_config_frs_ds_test.guid
			segment_id = "%s"
			name       = "TF Segment for Rules DS"
			rules {
				attribute_name = "email"
				operator       = "endsWith"
				values         = ["@ibm.com"]
			}
		}

		resource "ibm_app_config_feature" "frs_ds_feature" {
			guid           = ibm_resource_instance.app_config_frs_ds_test.guid
			environment_id = "%s"
			feature_id     = "%s"
			name           = "TF Feature for Rules DS"
			type           = "BOOLEAN"
			enabled_value  = true
			disabled_value = false
			enabled        = false
			lifecycle {
				ignore_changes = [rollout_type, rollout_percentage, segment_rules]
			}
		}

		resource "ibm_app_config_feature_rule" "frs_ds_rule" {
			guid           = ibm_resource_instance.app_config_frs_ds_test.guid
			environment_id = "%s"
			feature_id     = ibm_app_config_feature.frs_ds_feature.feature_id
			rule_id        = "%s"
			rule_name      = "TF Rule for Rules DS"
			value          = "true"
			rules {
				segments = [ibm_app_config_segment.frs_ds_segment.segment_id]
			}
			lifecycle {
				ignore_changes = [rollout_type, rollout_percentage]
			}
		}

		data "ibm_app_config_feature_rules" "test_ds" {
			guid           = ibm_resource_instance.app_config_frs_ds_test.guid
			environment_id = ibm_app_config_feature_rule.frs_ds_rule.environment_id
			feature_id     = ibm_app_config_feature_rule.frs_ds_rule.feature_id
		}`, instanceName, segmentID, environmentID, featureID, environmentID, ruleID)
}
