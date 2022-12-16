// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package metricsrouter_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMMetricsRouterTargetsDataSourceBasic(t *testing.T) {
	targetName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	targetDestinationCRN := "crn:v1:bluemix:public:sysdig-monitor:us-south:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMMetricsRouterTargetsDataSourceConfigBasic(targetName, targetDestinationCRN),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_metrics_router_targets.metrics_router_targets_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_metrics_router_targets.metrics_router_targets_instance", "targets.#"),
					resource.TestCheckResourceAttr("data.ibm_metrics_router_targets.metrics_router_targets_instance", "targets.0.name", targetName),
					resource.TestCheckResourceAttr("data.ibm_metrics_router_targets.metrics_router_targets_instance", "targets.0.destination_crn", targetDestinationCRN),
				),
			},
		},
	})
}

func TestAccIBMMetricsRouterTargetsDataSourceAllArgs(t *testing.T) {
	targetName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	targetDestinationCRN := "crn:v1:bluemix:public:sysdig-monitor:us-south:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::"
	targetRegion := "us-south"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMMetricsRouterTargetsDataSourceConfig(targetName, targetDestinationCRN, targetRegion),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_metrics_router_targets.metrics_router_targets_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_metrics_router_targets.metrics_router_targets_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_metrics_router_targets.metrics_router_targets_instance", "targets.#"),
					resource.TestCheckResourceAttrSet("data.ibm_metrics_router_targets.metrics_router_targets_instance", "targets.0.id"),
					resource.TestCheckResourceAttr("data.ibm_metrics_router_targets.metrics_router_targets_instance", "targets.0.name", targetName),
					resource.TestCheckResourceAttrSet("data.ibm_metrics_router_targets.metrics_router_targets_instance", "targets.0.crn"),
					resource.TestCheckResourceAttr("data.ibm_metrics_router_targets.metrics_router_targets_instance", "targets.0.destination_crn", targetDestinationCRN),
					resource.TestCheckResourceAttrSet("data.ibm_metrics_router_targets.metrics_router_targets_instance", "targets.0.target_type"),
					resource.TestCheckResourceAttr("data.ibm_metrics_router_targets.metrics_router_targets_instance", "targets.0.region", targetRegion),
					resource.TestCheckResourceAttrSet("data.ibm_metrics_router_targets.metrics_router_targets_instance", "targets.0.created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_metrics_router_targets.metrics_router_targets_instance", "targets.0.updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_metrics_router_targets.metrics_router_targets_instance", "targets.0.api_version"),
				),
			},
		},
	})
}

func testAccCheckIBMMetricsRouterTargetsDataSourceConfigBasic(targetName string, targetDestinationCRN string) string {
	return fmt.Sprintf(`
		resource "ibm_metrics_router_target" "metrics_router_target_instance" {
			name = "%s"
			destination_crn = "%s"
		}

		data "ibm_metrics_router_targets" "metrics_router_targets_instance" {
			name = ibm_metrics_router_target.metrics_router_target_instance.name
		}
	`, targetName, targetDestinationCRN)
}

func testAccCheckIBMMetricsRouterTargetsDataSourceConfig(targetName string, targetDestinationCRN string, targetRegion string) string {
	return fmt.Sprintf(`
		resource "ibm_metrics_router_target" "metrics_router_target_instance" {
			name = "%s"
			destination_crn = "%s"
			region = "%s"
		}

		data "ibm_metrics_router_targets" "metrics_router_targets_instance" {
			name = ibm_metrics_router_target.metrics_router_target_instance.name
		}
	`, targetName, targetDestinationCRN, targetRegion)
}
