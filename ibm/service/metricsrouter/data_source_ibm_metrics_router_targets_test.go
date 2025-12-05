// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.108.0-56772134-20251111-102802
 */

package metricsrouter_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/metricsrouter"
	. "github.com/IBM-Cloud/terraform-provider-ibm/ibm/unittest"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/metricsrouterv3"
	"github.com/stretchr/testify/assert"
)

func TestAccIBMMetricsRouterTargetsDataSourceBasic(t *testing.T) {
	targetName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMMetricsRouterTargetsDataSourceConfigBasic(targetName, destinationCRN),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_metrics_router_targets.metrics_router_targets_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_metrics_router_targets.metrics_router_targets_instance", "targets.#"),
					resource.TestCheckResourceAttr("data.ibm_metrics_router_targets.metrics_router_targets_instance", "targets.0.name", targetName),
					resource.TestCheckResourceAttr("data.ibm_metrics_router_targets.metrics_router_targets_instance", "targets.0.destination_crn", destinationCRN),
				),
			},
		},
	})
}

func TestAccIBMMetricsRouterTargetsDataSourceAllArgs(t *testing.T) {
	targetName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	targetRegion := "us-south"
	targetManagedBy := "enterprise"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMMetricsRouterTargetsDataSourceConfig(targetName, destinationCRN, targetRegion, targetManagedBy),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_metrics_router_targets.metrics_router_targets_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_metrics_router_targets.metrics_router_targets_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_metrics_router_targets.metrics_router_targets_instance", "targets.#"),
					resource.TestCheckResourceAttrSet("data.ibm_metrics_router_targets.metrics_router_targets_instance", "targets.0.id"),
					resource.TestCheckResourceAttr("data.ibm_metrics_router_targets.metrics_router_targets_instance", "targets.0.name", targetName),
					resource.TestCheckResourceAttrSet("data.ibm_metrics_router_targets.metrics_router_targets_instance", "targets.0.crn"),
					resource.TestCheckResourceAttr("data.ibm_metrics_router_targets.metrics_router_targets_instance", "targets.0.destination_crn", destinationCRN),
					resource.TestCheckResourceAttrSet("data.ibm_metrics_router_targets.metrics_router_targets_instance", "targets.0.target_type"),
					resource.TestCheckResourceAttr("data.ibm_metrics_router_targets.metrics_router_targets_instance", "targets.0.region", targetRegion),
					resource.TestCheckResourceAttrSet("data.ibm_metrics_router_targets.metrics_router_targets_instance", "targets.0.created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_metrics_router_targets.metrics_router_targets_instance", "targets.0.updated_at"),
					resource.TestCheckResourceAttr("data.ibm_metrics_router_targets.metrics_router_targets_instance", "targets.0.managed_by", targetManagedBy),
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

func testAccCheckIBMMetricsRouterTargetsDataSourceConfig(targetName string, targetDestinationCRN string, targetRegion string, targetManagedBy string) string {
	return fmt.Sprintf(`
		resource "ibm_metrics_router_target" "metrics_router_target_instance" {
			name = "%s"
			destination_crn = "%s"
			region = "%s"
			managed_by = "%s"
		}

		data "ibm_metrics_router_targets" "metrics_router_targets_instance" {
			name = ibm_metrics_router_target.metrics_router_target_instance.name
		}
	`, targetName, targetDestinationCRN, targetRegion, targetManagedBy)
}

func TestDataSourceIBMMetricsRouterTargetsTargetToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["id"] = "f7dcfae6-e7c5-08ca-451b-fdfa696c9bb6"
		model["name"] = "a-mr-target-us-south"
		model["crn"] = "crn:v1:bluemix:public:metrics-router:us-south:a/0be5ad401ae913d8ff665d92680664ed:b6eec08b-5201-08ca-451b-cd71523e3626:target:f7dcfae6-e7c5-08ca-451b-fdfa696c9bb6"
		model["destination_crn"] = "crn:v1:bluemix:public:sysdig-monitor:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::"
		model["target_type"] = "sysdig_monitor"
		model["region"] = "us-south"
		model["created_at"] = "2021-05-18T20:15:12.353Z"
		model["updated_at"] = "2021-05-18T20:15:12.353Z"
		model["managed_by"] = "enterprise"

		assert.Equal(t, result, model)
	}

	model := new(metricsrouterv3.Target)
	model.ID = core.StringPtr("f7dcfae6-e7c5-08ca-451b-fdfa696c9bb6")
	model.Name = core.StringPtr("a-mr-target-us-south")
	model.CRN = core.StringPtr("crn:v1:bluemix:public:metrics-router:us-south:a/0be5ad401ae913d8ff665d92680664ed:b6eec08b-5201-08ca-451b-cd71523e3626:target:f7dcfae6-e7c5-08ca-451b-fdfa696c9bb6")
	model.DestinationCRN = core.StringPtr("crn:v1:bluemix:public:sysdig-monitor:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::")
	model.TargetType = core.StringPtr("sysdig_monitor")
	model.Region = core.StringPtr("us-south")
	model.CreatedAt = CreateMockDateTime("2021-05-18T20:15:12.353Z")
	model.UpdatedAt = CreateMockDateTime("2021-05-18T20:15:12.353Z")
	model.ManagedBy = core.StringPtr("enterprise")

	result, err := metricsrouter.DataSourceIBMMetricsRouterTargetsTargetToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
