// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package logsrouter_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/logsrouter"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/logsrouterv3"
	"github.com/stretchr/testify/assert"
)

func TestAccIBMLogsRouterSettingsBasic(t *testing.T) {
	var conf logsrouterv3.Setting

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMLogsRouterSettingsDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMLogsRouterSettingsConfigBasic(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMLogsRouterSettingsExists("ibm_logs_router_settings.logs_router_settings_instance", conf),
				),
			},
		},
	})
}

func TestAccIBMLogsRouterSettingsAllArgs(t *testing.T) {
	var conf logsrouterv3.Setting
	primaryMetadataRegion := fmt.Sprintf("tf_primary_metadata_region_%d", acctest.RandIntRange(10, 100))
	backupMetadataRegion := fmt.Sprintf("tf_backup_metadata_region_%d", acctest.RandIntRange(10, 100))
	privateAPIEndpointOnly := "false"
	primaryMetadataRegionUpdate := fmt.Sprintf("tf_primary_metadata_region_%d", acctest.RandIntRange(10, 100))
	backupMetadataRegionUpdate := fmt.Sprintf("tf_backup_metadata_region_%d", acctest.RandIntRange(10, 100))
	privateAPIEndpointOnlyUpdate := "true"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMLogsRouterSettingsDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMLogsRouterSettingsConfig(primaryMetadataRegion, backupMetadataRegion, privateAPIEndpointOnly),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMLogsRouterSettingsExists("ibm_logs_router_settings.logs_router_settings_instance", conf),
					resource.TestCheckResourceAttr("ibm_logs_router_settings.logs_router_settings_instance", "primary_metadata_region", primaryMetadataRegion),
					resource.TestCheckResourceAttr("ibm_logs_router_settings.logs_router_settings_instance", "backup_metadata_region", backupMetadataRegion),
					resource.TestCheckResourceAttr("ibm_logs_router_settings.logs_router_settings_instance", "private_api_endpoint_only", privateAPIEndpointOnly),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMLogsRouterSettingsConfig(primaryMetadataRegionUpdate, backupMetadataRegionUpdate, privateAPIEndpointOnlyUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_logs_router_settings.logs_router_settings_instance", "primary_metadata_region", primaryMetadataRegionUpdate),
					resource.TestCheckResourceAttr("ibm_logs_router_settings.logs_router_settings_instance", "backup_metadata_region", backupMetadataRegionUpdate),
					resource.TestCheckResourceAttr("ibm_logs_router_settings.logs_router_settings_instance", "private_api_endpoint_only", privateAPIEndpointOnlyUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_logs_router_settings.logs_router_settings_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMLogsRouterSettingsConfigBasic() string {
	return fmt.Sprintf(`
		resource "ibm_logs_router_target" "logs_router_target_instance" {
			name = "my-lr-target"
			destination_crn = "crn:v1:bluemix:public:logs:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::"
		}
		resource "ibm_logs_router_settings" "logs_router_settings_instance" {
		}
	`)
}

func testAccCheckIBMLogsRouterSettingsConfig(primaryMetadataRegion string, backupMetadataRegion string, privateAPIEndpointOnly string) string {
	return fmt.Sprintf(`

		resource "ibm_logs_router_target" "logs_router_target_instance" {
			name = "my-lr-target"
			destination_crn = "crn:v1:bluemix:public:logs:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::"
		}

		resource "ibm_logs_router_settings" "logs_router_settings_instance" {
			default_targets {
				id = "c3af557f-fb0e-4476-85c3-0889e7fe7bc4"
				crn = "crn:v1:bluemix:public:logs:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::"
				name = "a-lr-target-us-south"
				target_type = "cloud_logs"
			}
			permitted_target_regions = "FIXME"
			primary_metadata_region = "%s"
			backup_metadata_region = "%s"
			private_api_endpoint_only = %s
		}
	`, primaryMetadataRegion, backupMetadataRegion, privateAPIEndpointOnly)
}

func testAccCheckIBMLogsRouterSettingsExists(n string, obj logsrouterv3.Setting) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		logsRouterClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).LogsRouterV3()
		if err != nil {
			return err
		}

		getSettingsOptions := &logsrouterv3.GetSettingsOptions{}

		getSettingsOptions.SetPrimaryMetadataRegion(rs.Primary.ID)

		setting, _, err := logsRouterClient.GetSettings(getSettingsOptions)
		if err != nil {
			return err
		}

		obj = *setting
		return nil
	}
}

func testAccCheckIBMLogsRouterSettingsDestroy(s *terraform.State) error {
	logsRouterClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).LogsRouterV3()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_logs_router_settings" {
			continue
		}

		getSettingsOptions := &logsrouterv3.GetSettingsOptions{}

		getSettingsOptions.SetPrimaryMetadataRegion(rs.Primary.ID)

		// Try to find the key
		_, response, err := logsRouterClient.GetSettings(getSettingsOptions)

		if err == nil {
			return fmt.Errorf("logs_router_settings still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for logs_router_settings (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func TestResourceIBMLogsRouterSettingsTargetReferenceToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["id"] = "c3af557f-fb0e-4476-85c3-0889e7fe7bc4"
		model["crn"] = "crn:v1:bluemix:public:logs:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::"
		model["name"] = "a-lr-target-us-south"
		model["target_type"] = "cloud_logs"

		assert.Equal(t, result, model)
	}

	model := new(logsrouterv3.TargetReference)
	model.ID = core.StringPtr("c3af557f-fb0e-4476-85c3-0889e7fe7bc4")
	model.CRN = core.StringPtr("crn:v1:bluemix:public:logs:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::")
	model.Name = core.StringPtr("a-lr-target-us-south")
	model.TargetType = core.StringPtr("cloud_logs")

	result, err := logsrouter.ResourceIBMLogsRouterSettingsTargetReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMLogsRouterSettingsMapToTargetIdentity(t *testing.T) {
	checkResult := func(result *logsrouterv3.TargetIdentity) {
		model := new(logsrouterv3.TargetIdentity)
		model.ID = core.StringPtr("c3af557f-fb0e-4476-85c3-0889e7fe7bc4")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["id"] = "c3af557f-fb0e-4476-85c3-0889e7fe7bc4"

	result, err := logsrouter.ResourceIBMLogsRouterSettingsMapToTargetIdentity(model)
	assert.Nil(t, err)
	checkResult(result)
}
