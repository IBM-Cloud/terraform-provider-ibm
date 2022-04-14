// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package atracker_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM/platform-services-go-sdk/atrackerv2"
)

func TestAccIBMAtrackerSettingsBasic(t *testing.T) {
	var conf atrackerv2.Settings
	metadataRegionPrimary := fmt.Sprintf("tf_metadata_region_primary_%d", acctest.RandIntRange(10, 100))
	privateAPIEndpointOnly := "false"
	metadataRegionPrimaryUpdate := fmt.Sprintf("tf_metadata_region_primary_%d", acctest.RandIntRange(10, 100))
	privateAPIEndpointOnlyUpdate := "true"

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
				Config: testAccCheckIBMAtrackerSettingsConfigBasic(metadataRegionPrimaryUpdate, privateAPIEndpointOnlyUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_atracker_settings.atracker_settings", "metadata_region_primary", metadataRegionPrimaryUpdate),
					resource.TestCheckResourceAttr("ibm_atracker_settings.atracker_settings", "private_api_endpoint_only", privateAPIEndpointOnlyUpdate),
				),
			},
		},
	})
}

func TestAccIBMAtrackerSettingsAllArgs(t *testing.T) {
	var conf atrackerv2.Settings
	metadataRegionPrimary := "us-south"
	privateAPIEndpointOnly := "false"
	metadataRegionPrimaryUpdate := "us-east"
	privateAPIEndpointOnlyUpdate := "true"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMAtrackerSettingsDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMAtrackerSettingsConfig(metadataRegionPrimary, privateAPIEndpointOnly),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMAtrackerSettingsExists("ibm_atracker_settings.atracker_settings", conf),
					resource.TestCheckResourceAttr("ibm_atracker_settings.atracker_settings", "metadata_region_primary", metadataRegionPrimary),
					resource.TestCheckResourceAttr("ibm_atracker_settings.atracker_settings", "private_api_endpoint_only", privateAPIEndpointOnly),
					// resource.TestCheckResourceAttr("ibm_atracker_settings.atracker_settings", "metadata_region_backup", metadataRegionBackup),
				),
			},
			{
				Config: testAccCheckIBMAtrackerSettingsConfig(metadataRegionPrimaryUpdate, privateAPIEndpointOnlyUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_atracker_settings.atracker_settings", "metadata_region_primary", metadataRegionPrimaryUpdate),
					resource.TestCheckResourceAttr("ibm_atracker_settings.atracker_settings", "private_api_endpoint_only", privateAPIEndpointOnlyUpdate),
					// resource.TestCheckResourceAttr("ibm_atracker_settings.atracker_settings", "metadata_region_backup", metadataRegionBackupUpdate),
				),
			},
			{
				ResourceName:      "ibm_atracker_settings.atracker_settings",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMAtrackerSettingsConfigBasic(metadataRegionPrimary string, privateAPIEndpointOnly string) string {
	return fmt.Sprintf(`
		resource "ibm_atracker_target" "atracker_target" {
			name = "my-cos-target"
			target_type = "cloud_object_storage"
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
		_, response, err := atrackerClient.GetSettings(getSettingsOptions)

		if err == nil {
			return fmt.Errorf("[ERROR] Activity Tracker Settings still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("[ERROR] Error checking for Activity Tracker Settings (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
