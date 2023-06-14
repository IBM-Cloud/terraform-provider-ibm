// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package metricsrouter_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM/platform-services-go-sdk/metricsrouterv3"
)

func TestAccIBMMetricsRouterSettingsBasic(t *testing.T) {
	var conf metricsrouterv3.Setting
	primaryMetadataRegion := "us-south"
	backupMetadataRegion := "us-east"
	permittedTargetRegions := "us-south"
	privateAPIEndpointOnly := "false"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMMetricsRouterSettingsDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMMetricsRouterSettingsConfigBasic(permittedTargetRegions, primaryMetadataRegion, backupMetadataRegion, privateAPIEndpointOnly),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMMetricsRouterSettingsExists("ibm_metrics_router_settings.metrics_router_settings_instance", conf),
					resource.TestCheckResourceAttr("ibm_metrics_router_settings.metrics_router_settings_instance", "permitted_target_regions.0", permittedTargetRegions),
					resource.TestCheckResourceAttr("ibm_metrics_router_settings.metrics_router_settings_instance", "primary_metadata_region", primaryMetadataRegion),
					resource.TestCheckResourceAttr("ibm_metrics_router_settings.metrics_router_settings_instance", "backup_metadata_region", backupMetadataRegion),
					resource.TestCheckResourceAttr("ibm_metrics_router_settings.metrics_router_settings_instance", "private_api_endpoint_only", privateAPIEndpointOnly),
				),
			},
			{
				ResourceName:      "ibm_metrics_router_settings.metrics_router_settings_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIBMMetricsRouterSettingsAllArgs(t *testing.T) {
	var conf metricsrouterv3.Setting
	primaryMetadataRegion := "us-south"
	backupMetadataRegion := "us-east"
	permittedTargetRegions := "us-south"
	privateAPIEndpointOnly := "false"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMMetricsRouterSettingsDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMMetricsRouterSettingsConfig(permittedTargetRegions, primaryMetadataRegion, backupMetadataRegion, privateAPIEndpointOnly),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMMetricsRouterSettingsExists("ibm_metrics_router_settings.metrics_router_settings_instance", conf),
					resource.TestCheckResourceAttr("ibm_metrics_router_settings.metrics_router_settings_instance", "permitted_target_regions.0", permittedTargetRegions),
					resource.TestCheckResourceAttr("ibm_metrics_router_settings.metrics_router_settings_instance", "primary_metadata_region", primaryMetadataRegion),
					resource.TestCheckResourceAttr("ibm_metrics_router_settings.metrics_router_settings_instance", "backup_metadata_region", backupMetadataRegion),
					resource.TestCheckResourceAttr("ibm_metrics_router_settings.metrics_router_settings_instance", "private_api_endpoint_only", privateAPIEndpointOnly),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_metrics_router_settings.metrics_router_settings_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIBMMetricsRouterSettingsEmptyTarget(t *testing.T) {
	primaryMetadataRegion := "us-south"
	backupMetadataRegion := "us-east"
	permittedTargetRegions := "us-south"
	privateAPIEndpointOnly := "false"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMMetricsRouterSettingsDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckIBMMetricsRouterSettingsEmptyTarget(permittedTargetRegions, primaryMetadataRegion, backupMetadataRegion, privateAPIEndpointOnly),
				ExpectError: regexp.MustCompile("should match regexp"),
			},
		},
	})
}

func testAccCheckIBMMetricsRouterSettingsConfigBasic(permittedTargetRegions, primaryMetadataRegion, backupMetadataRegion, privateAPIEndpointOnly string) string {
	return fmt.Sprintf(`
		resource "ibm_metrics_router_target" "metrics_router_target_instance" {
			name = "my-mr-target"
			destination_crn = "%s"
		}

		resource "ibm_metrics_router_settings" "metrics_router_settings_instance" {
			default_targets {
				id = ibm_metrics_router_target.metrics_router_target_instance.id
			}
			permitted_target_regions = ["%s"]
			primary_metadata_region = "%s"
			backup_metadata_region = "%s"
			private_api_endpoint_only = %s
		}
	`, destinationCRN, permittedTargetRegions, primaryMetadataRegion, backupMetadataRegion, privateAPIEndpointOnly)
}

func testAccCheckIBMMetricsRouterSettingsConfig(permittedTargetRegions, primaryMetadataRegion, backupMetadataRegion, privateAPIEndpointOnly string) string {
	return fmt.Sprintf(`

		resource "ibm_metrics_router_target" "metrics_router_target_instance" {
			name = "my-mr-target"
			destination_crn = "%s"
		}

		resource "ibm_metrics_router_settings" "metrics_router_settings_instance" {
			default_targets {
				id = ibm_metrics_router_target.metrics_router_target_instance.id
			}
			permitted_target_regions = ["%s"]
			primary_metadata_region = "%s"
			backup_metadata_region = "%s"
			private_api_endpoint_only = %s
		}
	`, destinationCRN, permittedTargetRegions, primaryMetadataRegion, backupMetadataRegion, privateAPIEndpointOnly)
}

func testAccCheckIBMMetricsRouterSettingsEmptyTarget(permittedTargetRegions, primaryMetadataRegion, backupMetadataRegion, privateAPIEndpointOnly string) string {
	return fmt.Sprintf(`
        resource "ibm_metrics_router_settings" "metrics_router_settings_instance" {
            default_targets {
                id = ""
            }
            permitted_target_regions = ["%s"]
            primary_metadata_region = "%s"
            backup_metadata_region = "%s"
            private_api_endpoint_only = %s
        }
    `, permittedTargetRegions, primaryMetadataRegion, backupMetadataRegion, privateAPIEndpointOnly)
}

func testAccCheckIBMMetricsRouterSettingsExists(n string, obj metricsrouterv3.Setting) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		_, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		metricsRouterClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).MetricsRouterV3()
		if err != nil {
			return err
		}

		getSettingsOptions := &metricsrouterv3.GetSettingsOptions{}

		setting, _, err := metricsRouterClient.GetSettings(getSettingsOptions)
		if err != nil {
			return err
		}

		obj = *setting
		return nil
	}
}

func testAccCheckIBMMetricsRouterSettingsDestroy(s *terraform.State) error {
	metricsRouterClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).MetricsRouterV3()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_metrics_router_settings" {
			continue
		}

		getSettingsOptions := &metricsrouterv3.GetSettingsOptions{}

		// Try to find the key
		settings, response, err := metricsRouterClient.GetSettings(getSettingsOptions)

		if err == nil {
			// Settings can never really truely be deleted (at least for PrimaryMetadataRegion and BackupMetadataRegion) but the other fields will be cleared
			if *settings.PrimaryMetadataRegion == rs.Primary.ID && len(*&settings.DefaultTargets) == 0 && len(*&settings.PermittedTargetRegions) == 0 {
				return nil
			}
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for metrics_router_settings (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
