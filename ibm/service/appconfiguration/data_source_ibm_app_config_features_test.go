// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package appconfiguration_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIbmAppConfigFeaturesDataSourceBasic(t *testing.T) {
	environmentID := "dev"
	featureType := "BOOLEAN"
	tags := "development feature"
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	featureID := fmt.Sprintf("tf_feature_id_%d", acctest.RandIntRange(10, 100))
	description := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	instanceName := fmt.Sprintf("tf_app_config_test_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmAppConfigFeaturesDataSourceConfigBasic(instanceName, name, environmentID, featureID, featureType, description, tags),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_app_config_features.app_config_features_data2", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_features.app_config_features_data2", "first.#"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_features.app_config_features_data2", "last.#"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_features.app_config_features_data2", "total_count"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_features.app_config_features_data2", "features.#"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_features.app_config_features_data2", "features.0.segment_exists"),
					resource.TestCheckResourceAttr("data.ibm_app_config_features.app_config_features_data2", "features.0.name", name),
					resource.TestCheckResourceAttr("data.ibm_app_config_features.app_config_features_data2", "features.0.type", featureType),
					resource.TestCheckResourceAttr("data.ibm_app_config_features.app_config_features_data2", "features.0.feature_id", featureID),
				),
			},
		},
	})
}

func testAccCheckIbmAppConfigFeaturesDataSourceConfigBasic(instanceName, name, environmentID, featureID, featureType, description, tags string) string {
	return fmt.Sprintf(`
		resource "ibm_resource_instance" "app_config_terraform_test487" {
			name     = "%s"
			location = "us-south"
			service  = "apprapp"
			plan     = "lite"
		}
		
		resource "ibm_app_config_feature" "app_config_feature_resource2" {
			guid           	= ibm_resource_instance.app_config_terraform_test487.guid
			name           	= "%s"
			environment_id  = "%s"
			feature_id     	= "%s"
			type           	= "%s"
			enabled_value  	= true
			disabled_value 	= false
			description    	= "%s"
			tags    		= "%s"
			rollout_percentage  = "80"
		}
		
		data "ibm_app_config_features" "app_config_features_data2" {
			expand 				= true
			guid          = ibm_app_config_feature.app_config_feature_resource2.guid
			environment_id = ibm_app_config_feature.app_config_feature_resource2.environment_id
		}
		`, instanceName, name, environmentID, featureID, featureType, description, tags)
}
