// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package appconfiguration_test

import (
	"fmt"
	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"testing"

	"github.com/IBM/appconfiguration-go-admin-sdk/appconfigurationv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
)

func TestAccIbmIbmAppConfigPropertyBasic(t *testing.T) {
	var conf appconfigurationv1.Property
	instanceName := fmt.Sprintf("tf_app_config_test_%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	propertyID := fmt.Sprintf("tf_property_id_%d", acctest.RandIntRange(10, 100))
	typeVar := "BOOLEAN"
	description := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	tags := fmt.Sprintf("tf_tags_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	descriptionUpdate := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	tagsUpdate := fmt.Sprintf("tf_tags_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmAppConfigPropertyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmAppConfigPropertyConfigBasic(instanceName, name, propertyID, typeVar, description, tags),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmAppConfigPropertyExists("ibm_app_config_property.ibm_app_config_property_resource1", conf),
					resource.TestCheckResourceAttrSet("ibm_app_config_property.ibm_app_config_property_resource1", "environment_id"),
					resource.TestCheckResourceAttrSet("ibm_app_config_property.ibm_app_config_property_resource1", "name"),
					resource.TestCheckResourceAttrSet("ibm_app_config_property.ibm_app_config_property_resource1", "property_id"),
					resource.TestCheckResourceAttrSet("ibm_app_config_property.ibm_app_config_property_resource1", "type"),
					resource.TestCheckResourceAttrSet("ibm_app_config_property.ibm_app_config_property_resource1", "description"),
					resource.TestCheckResourceAttrSet("ibm_app_config_property.ibm_app_config_property_resource1", "tags"),
					resource.TestCheckResourceAttrSet("ibm_app_config_property.ibm_app_config_property_resource1", "created_time"),
					resource.TestCheckResourceAttrSet("ibm_app_config_property.ibm_app_config_property_resource1", "updated_time"),
					resource.TestCheckResourceAttrSet("ibm_app_config_property.ibm_app_config_property_resource1", "href"),
					resource.TestCheckResourceAttrSet("ibm_app_config_property.ibm_app_config_property_resource1", "segment_exists"),
				),
			},
			{
				Config: testAccCheckIbmAppConfigPropertyConfigBasic(instanceName, nameUpdate, propertyID, typeVar, descriptionUpdate, tagsUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_app_config_property.ibm_app_config_property_resource1", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_app_config_property.ibm_app_config_property_resource1", "description", descriptionUpdate),
					resource.TestCheckResourceAttr("ibm_app_config_property.ibm_app_config_property_resource1", "tags", tagsUpdate),
				),
			},
		},
	})
}

func testAccCheckIbmAppConfigPropertyConfigBasic(instanceName, name, propertyID, typeVar, description, tags string) string {
	return fmt.Sprintf(`
    	resource "ibm_resource_instance" "app_config_terraform_test476" {
    		name     = "%s"
    		location = "us-south"
    		service  = "apprapp"
    		plan     = "lite"
    	}
    		resource "ibm_app_config_property" "ibm_app_config_property_resource1" {
    			guid 					= ibm_resource_instance.app_config_terraform_test476.guid
    			environment_id = "dev"
    			name = "%s"
    			property_id = "%s"
    			type = "%s"
    			value = "false"
    			description = "%s"
    			tags = "%s"
    		}`, instanceName, name, propertyID, typeVar, description, tags)
}

func testAccCheckIbmAppConfigPropertyExists(n string, obj appconfigurationv1.Property) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return flex.FmtErrorf("Not found: %s", n)
		}

		parts, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return flex.FmtErrorf(fmt.Sprintf("%s", err))
		}

		appconfigClient, err := getAppConfigClient(acc.TestAccProvider.Meta(), parts[0])
		if err != nil {
			return flex.FmtErrorf(fmt.Sprintf("%s", err))
		}
		options := &appconfigurationv1.GetPropertyOptions{}

		options.SetEnvironmentID(parts[1])
		options.SetPropertyID(parts[2])

		property, _, err := appconfigClient.GetProperty(options)
		if err != nil {
			return flex.FmtErrorf(fmt.Sprintf("%s", err))
		}

		obj = *property
		return nil
	}
}

func testAccCheckIbmAppConfigPropertyDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_app-config-property" {
			continue
		}
		parts, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return flex.FmtErrorf(fmt.Sprintf("%s", err))
		}

		appconfigClient, err := getAppConfigClient(acc.TestAccProvider.Meta(), parts[0])
		if err != nil {
			return flex.FmtErrorf(fmt.Sprintf("%s", err))
		}
		options := &appconfigurationv1.GetPropertyOptions{}

		options.SetEnvironmentID(parts[1])
		options.SetPropertyID(parts[2])

		// Try to find the key
		_, response, err := appconfigClient.GetProperty(options)

		if err == nil {
			return flex.FmtErrorf("app_config_property still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return flex.FmtErrorf("Error checking for app_config_property (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
