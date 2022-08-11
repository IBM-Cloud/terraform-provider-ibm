// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package appconfiguration_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/IBM/appconfiguration-go-admin-sdk/appconfigurationv1"
)

func TestAccIbmIbmAppConfigFeatureBasic(t *testing.T) {
	var conf appconfigurationv1.Feature
	instanceName := fmt.Sprintf("tf_app_config_test_%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	featureID := fmt.Sprintf("tf_feature_id_%d", acctest.RandIntRange(10, 100))
	featureType := "BOOLEAN"
	description := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	descriptionUpdate := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	tags := fmt.Sprintf("tags_%d", acctest.RandIntRange(10, 100))
	tagsUpdated := fmt.Sprintf("tags_updated_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmAppConfigFeatureDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmAppConfigFeatureConfigBasic(instanceName, name, featureID, featureType, description, tags),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmAppConfigFeatureExists("ibm_app_config_feature.ibm_app_config_feature_resource1", conf),
					resource.TestCheckResourceAttrSet("ibm_app_config_feature.ibm_app_config_feature_resource1", "id"),
					resource.TestCheckResourceAttrSet("ibm_app_config_feature.ibm_app_config_feature_resource1", "type"),
					resource.TestCheckResourceAttrSet("ibm_app_config_feature.ibm_app_config_feature_resource1", "name"),
					resource.TestCheckResourceAttrSet("ibm_app_config_feature.ibm_app_config_feature_resource1", "tags"),
					resource.TestCheckResourceAttrSet("ibm_app_config_feature.ibm_app_config_feature_resource1", "feature_id"),
					resource.TestCheckResourceAttrSet("ibm_app_config_feature.ibm_app_config_feature_resource1", "description"),
					resource.TestCheckResourceAttrSet("ibm_app_config_feature.ibm_app_config_feature_resource1", "enabled_value"),
					resource.TestCheckResourceAttrSet("ibm_app_config_feature.ibm_app_config_feature_resource1", "disabled_value"),
					resource.TestCheckResourceAttrSet("ibm_app_config_feature.ibm_app_config_feature_resource1", "rollout_percentage"),
				),
			},
			{
				Config: testAccCheckIbmAppConfigFeatureConfigBasic(instanceName, nameUpdate, featureID, featureType, descriptionUpdate, tagsUpdated),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_app_config_feature.ibm_app_config_feature_resource1", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_app_config_feature.ibm_app_config_feature_resource1", "tags", tagsUpdated),
					resource.TestCheckResourceAttr("ibm_app_config_feature.ibm_app_config_feature_resource1", "description", descriptionUpdate),
				),
			},
		},
	})
}

func testAccCheckIbmAppConfigFeatureConfigBasic(name, envName, featureID, featureType, description, tags string) string {
	return fmt.Sprintf(`
		resource "ibm_resource_instance" "app_config_terraform_test456" {
			name     = "%s"
			location = "us-south"
			service  = "apprapp"
			plan     = "lite"
		}
		resource "ibm_app_config_feature" "ibm_app_config_feature_resource1" {
			guid           	    = ibm_resource_instance.app_config_terraform_test456.guid
			name            	= "%s"
			environment_id      = "dev"
			feature_id     	    = "%s"
			type            	= "%s"
			enabled_value  	    = true
			disabled_value 	    = false
			description    	    = "%s"
			tags    		    = "%s"
			rollout_percentage  = "80"
		}`, name, envName, featureID, featureType, description, tags)
}

func testAccCheckIbmAppConfigFeatureExists(n string, obj appconfigurationv1.Feature) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}
		parts, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}
		appconfigClient, err := getAppConfigClient(acc.TestAccProvider.Meta(), parts[0])
		if err != nil {
			return err
		}

		options := &appconfigurationv1.GetFeatureOptions{}

		options.SetEnvironmentID(parts[1])
		options.SetFeatureID(parts[2])

		result, _, err := appconfigClient.GetFeature(options)
		if err != nil {
			return err
		}

		obj = *result
		return nil
	}
}

func testAccCheckIbmAppConfigFeatureDestroy(s *terraform.State) error {

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_app_config_feature_resource1" {
			continue
		}
		parts, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}
		appconfigClient, err := getAppConfigClient(acc.TestAccProvider.Meta(), parts[0])
		if err != nil {
			return err
		}
		options := &appconfigurationv1.GetFeatureOptions{}

		options.SetEnvironmentID(parts[1])
		options.SetFeatureID(parts[2])

		// Try to find the key
		_, response, err := appconfigClient.GetFeature(options)

		if err == nil {
			return fmt.Errorf("Feature still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("[ERROR] Error checking for Feature (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
