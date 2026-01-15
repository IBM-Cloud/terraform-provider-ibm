// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package logsrouter_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/logsrouter"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/logsrouterv3"
	"github.com/stretchr/testify/assert"
)

func TestAccIBMLogsRouterSettingsBasic(t *testing.T) {
	var conf logsrouterv3.Setting
	primaryMetadataRegion := "us-south"
	backupMetadataRegion := "us-east"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMLogsRouterSettingsDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMLogsRouterSettingsConfigBasic(primaryMetadataRegion, backupMetadataRegion),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMLogsRouterSettingsExists("ibm_logs_router_v3_settings.logs_router_settings_instance", conf),
				),
			},
		},
	})
}

func TestAccIBMLogsRouterSettingsAllArgs(t *testing.T) {
	var conf logsrouterv3.Setting
	primaryMetadataRegion := "us-south"
	backupMetadataRegion := "us-east"
	privateAPIEndpointOnly := "false"

	primaryMetadataRegionUpdate := "us-south"
	backupMetadataRegionUpdate := "eu-de"
	privateAPIEndpointOnlyUpdate := "false"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMLogsRouterSettingsDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMLogsRouterSettingsConfig(primaryMetadataRegion, backupMetadataRegion, privateAPIEndpointOnly),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMLogsRouterSettingsExists("ibm_logs_router_v3_settings.logs_router_settings_instance", conf),
					resource.TestCheckResourceAttr("ibm_logs_router_v3_settings.logs_router_settings_instance", "primary_metadata_region", primaryMetadataRegion),
					resource.TestCheckResourceAttr("ibm_logs_router_v3_settings.logs_router_settings_instance", "backup_metadata_region", backupMetadataRegion),
					resource.TestCheckResourceAttr("ibm_logs_router_v3_settings.logs_router_settings_instance", "private_api_endpoint_only", privateAPIEndpointOnly),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMLogsRouterSettingsConfig(primaryMetadataRegionUpdate, backupMetadataRegionUpdate, privateAPIEndpointOnlyUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_logs_router_v3_settings.logs_router_settings_instance", "primary_metadata_region", primaryMetadataRegionUpdate),
					resource.TestCheckResourceAttr("ibm_logs_router_v3_settings.logs_router_settings_instance", "backup_metadata_region", backupMetadataRegionUpdate),
					resource.TestCheckResourceAttr("ibm_logs_router_v3_settings.logs_router_settings_instance", "private_api_endpoint_only", privateAPIEndpointOnlyUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_logs_router_v3_settings.logs_router_settings_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

// Add this function to configure primary_metadata_region before creating targets
func testSettingsPrimaryMetadataRegion(primaryMetadataRegion, backupMetadataRegion, privateAPIEndpointOnly string) string {
	return fmt.Sprintf(`
		resource "ibm_logs_router_v3_settings" "logs_router_settings_instance" {
			primary_metadata_region = "%s"
			backup_metadata_region = "%s"
			private_api_endpoint_only = %s
		}
	`, primaryMetadataRegion, backupMetadataRegion, privateAPIEndpointOnly)
}

func testAccCheckIBMLogsRouterSettingsConfigBasic(primaryMetadataRegion, backupMetadataRegion string) string {
	return fmt.Sprintf(`
		resource "ibm_logs_router_v3_settings" "logs_router_settings_instance" {
			primary_metadata_region = "%s"
			backup_metadata_region = "%s"
		}
	`, primaryMetadataRegion, backupMetadataRegion)
}

func testAccCheckIBMLogsRouterSettingsConfig(primaryMetadataRegion string, backupMetadataRegion string, privateAPIEndpointOnly string) string {
	return fmt.Sprintf(`

		resource "ibm_logs_router_v3_target" "logs_router_target_instance" {
			name = "my-lr-target"
			destination_crn = "%s"
			region = "%s"
			managed_by = "account"
		}

		resource "ibm_logs_router_v3_settings" "logs_router_settings_instance" {
			default_targets {
				id = ibm_logs_router_v3_target.logs_router_target_instance.id
			}
			permitted_target_regions = ["%s"]
			primary_metadata_region = "%s"
			backup_metadata_region = "%s"
			private_api_endpoint_only = %s
		}
		`, iclDestinationCRN, primaryMetadataRegion, primaryMetadataRegion, primaryMetadataRegion, backupMetadataRegion, privateAPIEndpointOnly)
}

func testAccCheckIBMLogsRouterSettingsExists(n string, obj logsrouterv3.Setting) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		_, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		logsRouterClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).LogsRouterV3()
		if err != nil {
			return err
		}

		getSettingsOptions := &logsrouterv3.GetSettingsOptions{}

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
		if rs.Type != "ibm_logs_router_v3_settings" {
			continue
		}

		getSettingsOptions := &logsrouterv3.GetSettingsOptions{}

		// Try to find the key
		settings, response, err := logsRouterClient.GetSettings(getSettingsOptions)

		if err == nil {
			// Settings can never really truely be deleted (at least for PrimaryMetadataRegion and BackupMetadataRegion) but the other fields will be cleared
			if *settings.PrimaryMetadataRegion == rs.Primary.ID && len(settings.DefaultTargets) == 0 && len(settings.PermittedTargetRegions) == 0 {
				return nil
			}
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
