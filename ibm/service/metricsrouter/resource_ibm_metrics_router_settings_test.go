// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package metricsrouter_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM/platform-services-go-sdk/metricsrouterv3"
)

func TestAccIBMMetricsRouterSettingsBasic(t *testing.T) {
	var conf metricsrouterv3.Settings
	metadataRegionPrimary := "us-east"
	privateAPIEndpointOnly := "false"
	metadataRegionPrimaryUpdate := "us-south"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMMetricsRouterSettingsDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMMetricsRouterSettingsConfigBasic(metadataRegionPrimary, privateAPIEndpointOnly),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMMetricsRouterSettingsExists("ibm_metrics_router_settings.metrics_router_settings_instance", conf),
					resource.TestCheckResourceAttr("ibm_metrics_router_settings.metrics_router_settings_instance", "metadata_region_primary", metadataRegionPrimary),
					resource.TestCheckResourceAttr("ibm_metrics_router_settings.metrics_router_settings_instance", "private_api_endpoint_only", privateAPIEndpointOnly),
				),
			},
			{
				Config: testAccCheckIBMMetricsRouterSettingsConfigBasic(metadataRegionPrimaryUpdate, privateAPIEndpointOnly),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMMetricsRouterSettingsExists("ibm_metrics_router_settings.metrics_router_settings_instance", conf),
					resource.TestCheckResourceAttr("ibm_metrics_router_settings.metrics_router_settings_instance", "metadata_region_primary", metadataRegionPrimaryUpdate),
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
	var conf metricsrouterv3.Settings
	metadataRegionPrimary := "us-south"
	privateAPIEndpointOnly := "false"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMMetricsRouterSettingsDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMMetricsRouterSettingsAllArgs(metadataRegionPrimary, privateAPIEndpointOnly),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMMetricsRouterSettingsExists("ibm_metrics_router_settings.metrics_router_settings", conf),
					resource.TestCheckResourceAttr("ibm_metrics_router_settings.metrics_router_settings", "metadata_region_primary", metadataRegionPrimary),
					resource.TestCheckResourceAttr("ibm_metrics_router_settings.metrics_router_settings", "private_api_endpoint_only", privateAPIEndpointOnly),
				),
			},
			{
				ResourceName:      "ibm_metrics_router_settings.metrics_router_settings",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMMetricsRouterSettingsConfigBasic(metadataRegionPrimary string, privateAPIEndpointOnly string) string {
	return fmt.Sprintf(`
		resource "ibm_metrics_router_settings" "metrics_router_settings_instance" {
			metadata_region_primary = "%s"
			private_api_endpoint_only = %s
		}
	`, metadataRegionPrimary, privateAPIEndpointOnly)
}

func testAccCheckIBMMetricsRouterSettingsAllArgs(metadataRegionPrimary string, privateAPIEndpointOnly string) string {
	return fmt.Sprintf(`

		resource "ibm_metrics_router_target" "metrics_router_target" {
		    name = "my_target_updated2"
		    destination_crn = "crn:v1:bluemix:public:sysdig-monitor:us-south:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::"
		}
		resource "ibm_metrics_router_settings" "metrics_router_settings" {
			metadata_region_primary = "%s"
			private_api_endpoint_only = %s
			default_targets = [ibm_metrics_router_target.metrics_router_target.id]
		}
	`, metadataRegionPrimary, privateAPIEndpointOnly)
}

func testAccCheckIBMMetricsRouterSettingsExists(n string, obj metricsrouterv3.Settings) resource.TestCheckFunc {

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

		settings, _, err := metricsRouterClient.GetSettings(getSettingsOptions)
		if err != nil {
			return err
		}

		obj = *settings
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
			// Settings can never really truely be deleted (at least for MetaRegionPrimary) but the other fields will be cleared
			if *settings.MetadataRegionPrimary == rs.Primary.ID && len(*&settings.DefaultTargets) == 0 && len(*&settings.PermittedTargetRegions) == 0 {
				return nil
			}
			return fmt.Errorf("[ERROR] Metrics Router Settings still exists and other fields not deleted: %s %s, Targets: %v, PermittedRegions: %v", *settings.MetadataRegionPrimary, rs.Primary.ID, *&settings.DefaultTargets, *&settings.PermittedTargetRegions)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for Metrics Router Settings (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
