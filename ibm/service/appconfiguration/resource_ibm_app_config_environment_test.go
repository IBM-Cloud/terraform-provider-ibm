// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package appconfiguration_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/IBM/appconfiguration-go-admin-sdk/appconfigurationv1"
)

func TestAccIbmAppConfigEnvironmentBasic(t *testing.T) {
	var conf appconfigurationv1.Environment
	colorCode := "#e2a222"
	newColorCode := "#431133"
	name := fmt.Sprintf("name_%d", acctest.RandIntRange(10, 100))
	envName := fmt.Sprintf("env_%d", acctest.RandIntRange(10, 100))
	newEnvName := fmt.Sprintf("env_%d", acctest.RandIntRange(10, 100))
	description := fmt.Sprintf("description_%d", acctest.RandIntRange(10, 100))
	newDescription := fmt.Sprintf("description_%d", acctest.RandIntRange(10, 100))
	environmentID := fmt.Sprintf("environment_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmAppConfigEnvironmentDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmAppConfigEnvironmentConfigBasic(name, envName, environmentID, description, colorCode),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmAppConfigEnvironmentExists("ibm_app_config_environment.app_config_environment_resource1", conf),
					resource.TestCheckResourceAttrSet("ibm_app_config_environment.app_config_environment_resource1", "id"),
					resource.TestCheckResourceAttrSet("ibm_app_config_environment.app_config_environment_resource1", "name"),
					resource.TestCheckResourceAttrSet("ibm_app_config_environment.app_config_environment_resource1", "tags"),
					resource.TestCheckResourceAttrSet("ibm_app_config_environment.app_config_environment_resource1", "href"),
					resource.TestCheckResourceAttrSet("ibm_app_config_environment.app_config_environment_resource1", "color_code"),
					resource.TestCheckResourceAttrSet("ibm_app_config_environment.app_config_environment_resource1", "description"),
					resource.TestCheckResourceAttrSet("ibm_app_config_environment.app_config_environment_resource1", "created_time"),
					resource.TestCheckResourceAttrSet("ibm_app_config_environment.app_config_environment_resource1", "updated_time"),
					resource.TestCheckResourceAttrSet("ibm_app_config_environment.app_config_environment_resource1", "environment_id"),
				),
			},
			{
				Config: testAccCheckIbmAppConfigEnvironmentConfigBasic(name, newEnvName, environmentID, newDescription, newColorCode),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_app_config_environment.app_config_environment_resource1", "name", newEnvName),
					resource.TestCheckResourceAttr("ibm_app_config_environment.app_config_environment_resource1", "color_code", newColorCode),
					resource.TestCheckResourceAttr("ibm_app_config_environment.app_config_environment_resource1", "description", newDescription),
				),
			},
		},
	})
}

func testAccCheckIbmAppConfigEnvironmentConfigBasic(name, envName, environmentID, description, colorCode string) string {
	return fmt.Sprintf(`
		 resource "ibm_resource_instance" "app_config_terraform_test454"{
			 name     = "%s"
			 location = "us-south"
			 service  = "apprapp"
			 plan     = "lite"
		 }
		 resource "ibm_app_config_environment" "app_config_environment_resource1" {
			 name          		= "%s"
			 environment_id   = "%s"
			 description      = "%s"
			 color_code       = "%s"
			 tags							= "version v1"
			 guid 						= ibm_resource_instance.app_config_terraform_test454.guid
		 }`, name, envName, environmentID, description, colorCode)
}
func getAppConfigClient(meta interface{}, guid string) (*appconfigurationv1.AppConfigurationV1, error) {
	appconfigClient, err := meta.(conns.ClientSession).AppConfigurationV1()
	if err != nil {
		return nil, err
	}
	bluemixSession, err := meta.(conns.ClientSession).BluemixSession()
	if err != nil {
		return nil, err
	}
	appConfigURL := fmt.Sprintf("https://%s.apprapp.cloud.ibm.com/apprapp/feature/v1/instances/%s", bluemixSession.Config.Region, guid)
	url := conns.EnvFallBack([]string{"IBMCLOUD_APP_CONFIG_API_ENDPOINT"}, appConfigURL)
	appconfigClient.Service.Options.URL = url
	return appconfigClient, nil
}

func testAccCheckIbmAppConfigEnvironmentExists(n string, obj appconfigurationv1.Environment) resource.TestCheckFunc {

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

		options := &appconfigurationv1.GetEnvironmentOptions{}
		options.SetEnvironmentID(parts[1])

		result, _, err := appconfigClient.GetEnvironment(options)
		if err != nil {
			return flex.FmtErrorf(fmt.Sprintf("%s", err))
		}

		obj = *result
		return nil
	}
}

func testAccCheckIbmAppConfigEnvironmentDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "app_config_environment_resource1" {
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
		options := &appconfigurationv1.GetEnvironmentOptions{}
		options.SetEnvironmentID(parts[1])

		// Try to find the key
		_, response, err := appconfigClient.GetEnvironment(options)

		if err == nil {
			return flex.FmtErrorf("Environment still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return flex.FmtErrorf("[ERROR] Error checking for environment (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
