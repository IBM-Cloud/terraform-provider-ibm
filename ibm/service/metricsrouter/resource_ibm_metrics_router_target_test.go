// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package metricsrouter_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM/platform-services-go-sdk/metricsrouterv3"
)

func TestAccIBMMetricsRouterTargetBasic(t *testing.T) {
	var conf metricsrouterv3.Target
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	destinationCRN := "crn:v1:bluemix:public:sysdig-monitor:us-south:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::"
	nameUpdate := fmt.Sprintf("updated_tf_name_%d", acctest.RandIntRange(10, 100))
	destinationCRNUpdate := "crn:v1:bluemix:public:sysdig-monitor:us-south:a/11111111111111111111111111111111:33333333-3333-3333-3333-333333333333::"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMMetricsRouterTargetDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMMetricsRouterTargetConfigBasic(name, destinationCRN),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMMetricsRouterTargetExists("ibm_metrics_router_target.metrics_router_target", conf),
					resource.TestCheckResourceAttr("ibm_metrics_router_target.metrics_router_target", "name", name),
					resource.TestCheckResourceAttr("ibm_metrics_router_target.metrics_router_target", "destination_crn", destinationCRN),
				),
			},
			{
				Config: testAccCheckIBMMetricsRouterTargetConfigBasic(nameUpdate, destinationCRNUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_metrics_router_target.metrics_router_target", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_metrics_router_target.metrics_router_target", "destination_crn", destinationCRNUpdate),
				),
			},
			{
				ResourceName:      "ibm_metrics_router_target.metrics_router_target",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIBMMetricsRouterTargetAllArgs(t *testing.T) {
	var conf metricsrouterv3.Target
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	destinationCRN := "crn:v1:bluemix:public:sysdig-monitor:us-south:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::"
	region := "us-south"
	nameUpdate := fmt.Sprintf("updated_tf_name_%d", acctest.RandIntRange(10, 100))
	destinationCRNUpdate := "crn:v1:bluemix:public:sysdig-monitor:us-south:a/11111111111111111111111111111111:33333333-3333-3333-3333-333333333333::"
	regionUpdate := "us-east"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMMetricsRouterTargetDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMMetricsRouterTargetConfig(name, destinationCRN, region),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMMetricsRouterTargetExists("ibm_metrics_router_target.metrics_router_target_instance", conf),
					resource.TestCheckResourceAttr("ibm_metrics_router_target.metrics_router_target_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_metrics_router_target.metrics_router_target_instance", "destination_crn", destinationCRN),
					resource.TestCheckResourceAttr("ibm_metrics_router_target.metrics_router_target_instance", "region", region),
				),
			},
			{
				Config: testAccCheckIBMMetricsRouterTargetConfig(nameUpdate, destinationCRNUpdate, regionUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_metrics_router_target.metrics_router_target_instance", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_metrics_router_target.metrics_router_target_instance", "destination_crn", destinationCRNUpdate),
					resource.TestCheckResourceAttr("ibm_metrics_router_target.metrics_router_target_instance", "region", regionUpdate),
				),
			},
		},
	})
}

func testAccCheckIBMMetricsRouterTargetConfigBasic(name string, destinationCRN string) string {
	return fmt.Sprintf(`

		resource "ibm_metrics_router_target" "metrics_router_target" {
			name = "%s"
			destination_crn = "%s"
		}
	`, name, destinationCRN)
}

func testAccCheckIBMMetricsRouterTargetConfig(name string, destinationCRN string, region string) string {
	return fmt.Sprintf(`

		resource "ibm_metrics_router_target" "metrics_router_target_instance" {
			name = "%s"
			destination_crn = "%s"
			region = "%s"
		}
	`, name, destinationCRN, region)
}

func testAccCheckIBMMetricsRouterTargetExists(n string, obj metricsrouterv3.Target) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		metricsRouterClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).MetricsRouterV3()
		if err != nil {
			return err
		}

		getTargetOptions := &metricsrouterv3.GetTargetOptions{}

		getTargetOptions.SetID(rs.Primary.ID)

		target, _, err := metricsRouterClient.GetTarget(getTargetOptions)
		if err != nil {
			return err
		}

		obj = *target
		return nil
	}
}

func testAccCheckIBMMetricsRouterTargetDestroy(s *terraform.State) error {
	metricsRouterClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).MetricsRouterV3()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_metrics_router_target" {
			continue
		}

		getTargetOptions := &metricsrouterv3.GetTargetOptions{}

		getTargetOptions.SetID(rs.Primary.ID)

		// Try to find the key
		_, response, err := metricsRouterClient.GetTarget(getTargetOptions)

		if err == nil {
			return fmt.Errorf("metrics_router_target still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for metrics_router_target (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
