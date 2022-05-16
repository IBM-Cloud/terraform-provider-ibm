// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package atracker_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM/platform-services-go-sdk/atrackerv2"
)

func TestAccIBMAtrackerSettingsBasic(t *testing.T) {
	var conf atrackerv2.Settings
	metadataRegionPrimary := "us-east"
	privateAPIEndpointOnly := "false"
	metadataRegionPrimaryUpdate := "us-south"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMAtrackerSettingsDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMAtrackerSettingsConfigBasic(metadataRegionPrimary, privateAPIEndpointOnly),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMAtrackerSettingsExists("ibm_atracker_settings.atracker_settings", conf),
					resource.TestCheckResourceAttr("ibm_atracker_settings.atracker_settings", "metadata_region_primary", metadataRegionPrimary),
					resource.TestCheckResourceAttr("ibm_atracker_settings.atracker_settings", "private_api_endpoint_only", privateAPIEndpointOnly),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMAtrackerSettingsConfigBasic(metadataRegionPrimaryUpdate, privateAPIEndpointOnly),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_atracker_settings.atracker_settings", "metadata_region_primary", metadataRegionPrimaryUpdate),
					resource.TestCheckResourceAttr("ibm_atracker_settings.atracker_settings", "private_api_endpoint_only", privateAPIEndpointOnly),
				),
			},
		},
	})
}

func testAccCheckIBMAtrackerSettingsConfigBasic(metadataRegionPrimary string, privateAPIEndpointOnly string) string {
	return fmt.Sprintf(`
		resource "ibm_atracker_target" "atracker_target" {
			name = "my-cos-target"
			target_type = "cloud_object_storage"
			cos_endpoint {
				endpoint = "s3.private.us-east.cloud-object-storage.appdomain.cloud"
				target_crn = "crn:v1:bluemix:public:cloud-object-storage:global:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::"
				bucket = "my-atracker-bucket"
				api_key = "xxxxxxxxxxxxxx" # pragma: whitelist secret
				service_to_service_enabled = false
			}
		}

		resource "ibm_atracker_settings" "atracker_settings" {
			metadata_region_primary = "%s"
			private_api_endpoint_only = %s
		}
	`, metadataRegionPrimary, privateAPIEndpointOnly)
}

func testAccCheckIBMAtrackerSettingsConfig(metadataRegionPrimary string, privateAPIEndpointOnly string) string {
	return fmt.Sprintf(`

		resource "ibm_atracker_target" "atracker_target" {
			name = "my-cos-target"
			target_type = "cloud_object_storage"
			cos_endpoint {
				endpoint = "s3.private.us-east.cloud-object-storage.appdomain.cloud"
				target_crn = "crn:v1:bluemix:public:cloud-object-storage:global:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::"
				bucket = "my-atracker-bucket"
				api_key = "xxxxxxxxxxxxxx" # pragma: whitelist secret
				service_to_service_enabled = false
			}
		}

		resource "ibm_atracker_settings" "atracker_settings" {
			metadata_region_primary = "%s"
			private_api_endpoint_only = %s
			default_targets = [ ibm_atracker_target.atracker_target.id ]
			permitted_target_regions = ["us-south", "us-east"]
		}
	`, metadataRegionPrimary, privateAPIEndpointOnly)
}

func testAccCheckIBMAtrackerSettingsExists(n string, obj atrackerv2.Settings) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		atrackerClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).AtrackerV2()
		if err != nil {
			return err
		}

		getSettingsOptions := &atrackerv2.GetSettingsOptions{}

		settings, _, err := atrackerClient.GetSettings(getSettingsOptions)
		if err != nil {
			return err
		}

		obj = *settings
		return nil
	}
}

func testAccCheckIBMAtrackerSettingsDestroy(s *terraform.State) error {
	atrackerClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).AtrackerV2()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_atracker_settings" {
			continue
		}

		getSettingsOptions := &atrackerv2.GetSettingsOptions{}

		// Try to find the key
		settings, response, err := atrackerClient.GetSettings(getSettingsOptions)

		if err == nil {
			// Settings can never really truely be deleted (at least for MetaRegionPrimary) but the other fields will be cleared
			if *settings.MetadataRegionPrimary == rs.Primary.ID && len(*&settings.DefaultTargets) == 0 && len(*&settings.DefaultTargets) == 0 {
				return nil
			}
			return fmt.Errorf("[ERROR] Activity Tracker Settings still exists but other fields not deleted: %s, Targets: %v, PermittedRegions: %v", rs.Primary.ID, *&settings.DefaultTargets, *&settings.PermittedTargetRegions)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("[ERROR] Error checking for Activity Tracker Settings (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
