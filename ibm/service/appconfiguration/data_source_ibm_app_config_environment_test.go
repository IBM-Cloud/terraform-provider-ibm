// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package appconfiguration_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIbmAppConfigEnvironmentDataSourceBasic(t *testing.T) {
	name := fmt.Sprintf("name_%d", acctest.RandIntRange(10, 100))
	colorCode := "#e23433"
	tags := fmt.Sprintf("tags_%d", acctest.RandIntRange(10, 100))
	description := fmt.Sprintf("description_%d", acctest.RandIntRange(10, 100))
	envName := fmt.Sprintf("env_%d", acctest.RandIntRange(10, 100))
	environmentID := fmt.Sprintf("environment_id_%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmAppConfigEnvironmentDataSourceConfigBasic(name, envName, environmentID, description, colorCode, tags),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_app_config_environment.app_config_environment_data1", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_environment.app_config_environment_data1", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_environment.app_config_environment_data1", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_environment.app_config_environment_data1", "tags"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_environment.app_config_environment_data1", "color_code"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_environment.app_config_environment_data1", "description"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_environment.app_config_environment_data1", "updated_time"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_environment.app_config_environment_data1", "created_time"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_environment.app_config_environment_data1", "environment_id"),
				),
			},
		},
	})
}

func testAccCheckIbmAppConfigEnvironmentDataSourceConfigBasic(name, envName, environmentID, description, colorCode, tags string) string {
	return fmt.Sprintf(`
		 resource "ibm_resource_instance" "app_config_terraform_test45"{
			 name     = "%s"
			 location = "us-south"
			 service  = "apprapp"
			 plan     = "lite"
		 }
		 resource "ibm_app_config_environment" "app_config_environment_resource1" {
			 name          		= "%s"
			 environment_id    = "%s"
			 description       = "%s"
			 color_code        = "%s"
			 tags        			= "%s"
			 guid = ibm_resource_instance.app_config_terraform_test45.guid
		 }
		 data "ibm_app_config_environment" "app_config_environment_data1" {
			 expand						= true
			 guid 							= ibm_app_config_environment.app_config_environment_resource1.guid
			 environment_id    = ibm_app_config_environment.app_config_environment_resource1.environment_id
		 }`, name, envName, environmentID, description, colorCode, tags)
}
