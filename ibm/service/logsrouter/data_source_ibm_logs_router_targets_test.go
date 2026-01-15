// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.108.0-56772134-20251111-102802
 */

package logsrouter_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/logsrouter"
	. "github.com/IBM-Cloud/terraform-provider-ibm/ibm/unittest"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/logsrouterv3"
	"github.com/stretchr/testify/assert"
)

func TestAccIBMLogsRouterTargetsDataSourceBasic(t *testing.T) {
	targetName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	targetDestinationCRN := iclDestinationCRN

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMLogsRouterTargetsDataSourceConfigBasic(targetName, targetDestinationCRN),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_logs_router_v3_targets.logs_router_targets_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_router_v3_targets.logs_router_targets_instance", "targets.#"),
					resource.TestCheckResourceAttr("data.ibm_logs_router_v3_targets.logs_router_targets_instance", "targets.0.name", targetName),
					resource.TestCheckResourceAttr("data.ibm_logs_router_v3_targets.logs_router_targets_instance", "targets.0.destination_crn", targetDestinationCRN),
				),
			},
		},
	})
}

func TestAccIBMLogsRouterTargetsDataSourceAllArgs(t *testing.T) {
	targetName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	targetDestinationCRN := iclDestinationCRN
	targetRegion := "us-south"
	targetManagedBy := "enterprise"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMLogsRouterTargetsDataSourceConfig(targetName, targetDestinationCRN, targetRegion, targetManagedBy),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_logs_router_v3_targets.logs_router_targets_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_router_v3_targets.logs_router_targets_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_router_v3_targets.logs_router_targets_instance", "targets.#"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_router_v3_targets.logs_router_targets_instance", "targets.0.id"),
					resource.TestCheckResourceAttr("data.ibm_logs_router_v3_targets.logs_router_targets_instance", "targets.0.name", targetName),
					resource.TestCheckResourceAttrSet("data.ibm_logs_router_v3_targets.logs_router_targets_instance", "targets.0.crn"),
					resource.TestCheckResourceAttr("data.ibm_logs_router_v3_targets.logs_router_targets_instance", "targets.0.destination_crn", targetDestinationCRN),
					resource.TestCheckResourceAttrSet("data.ibm_logs_router_v3_targets.logs_router_targets_instance", "targets.0.target_type"),
					resource.TestCheckResourceAttr("data.ibm_logs_router_v3_targets.logs_router_targets_instance", "targets.0.region", targetRegion),
					resource.TestCheckResourceAttrSet("data.ibm_logs_router_v3_targets.logs_router_targets_instance", "targets.0.created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_router_v3_targets.logs_router_targets_instance", "targets.0.updated_at"),
					resource.TestCheckResourceAttr("data.ibm_logs_router_v3_targets.logs_router_targets_instance", "targets.0.managed_by", targetManagedBy),
				),
			},
		},
	})
}

func testAccCheckIBMLogsRouterTargetsDataSourceConfigBasic(targetName string, targetDestinationCRN string) string {
	return fmt.Sprintf(`
		resource "ibm_logs_router_v3_target" "logs_router_target_instance" {
			name = "%s"
			destination_crn = "%s"
			region = "us-south"
			managed_by = "account"
		}

		data "ibm_logs_router_v3_targets" "logs_router_targets_instance" {
			name = ibm_logs_router_v3_target.logs_router_target_instance.name
		}
	`, targetName, targetDestinationCRN)
}

func testAccCheckIBMLogsRouterTargetsDataSourceConfig(targetName string, targetDestinationCRN string, targetRegion string, targetManagedBy string) string {
	return fmt.Sprintf(`
		resource "ibm_logs_router_v3_target" "logs_router_target_instance" {
			name = "%s"
			destination_crn = "%s"
			region = "%s"
			managed_by = "%s"
		}

		data "ibm_logs_router_v3_targets" "logs_router_targets_instance" {
			name = ibm_logs_router_v3_target.logs_router_target_instance.name
		}
	`, targetName, targetDestinationCRN, targetRegion, targetManagedBy)
}

func TestDataSourceIBMLogsRouterTargetsTargetToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		writeStatusModel := make(map[string]interface{})
		writeStatusModel["status"] = "success"
		writeStatusModel["last_failure"] = "2025-05-18T20:15:12.353Z"
		writeStatusModel["reason_for_last_failure"] = "Provided API key could not be found"

		model := make(map[string]interface{})
		model["id"] = "f7dcfae6-e7c5-08ca-451b-fdfa696c9bb6"
		model["name"] = "a-lr-target-us-south"
		model["crn"] = "crn:v1:bluemix:public:logs-router:us-south:a/0be5ad401ae913d8ff665d92680664ed:b6eec08b-5201-08ca-451b-cd71523e3626:target:f7dcfae6-e7c5-08ca-451b-fdfa696c9bb6"
		model["destination_crn"] = "crn:v1:bluemix:public:logs:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::"
		model["target_type"] = "cloud_logs"
		model["region"] = "us-south"
		model["write_status"] = []map[string]interface{}{writeStatusModel}
		model["created_at"] = "2021-05-18T20:15:12.353Z"
		model["updated_at"] = "2021-05-18T20:15:12.353Z"
		model["managed_by"] = "enterprise"

		assert.Equal(t, result, model)
	}

	writeStatusModel := new(logsrouterv3.WriteStatus)
	writeStatusModel.Status = core.StringPtr("success")
	writeStatusModel.LastFailure = CreateMockDateTime("2025-05-18T20:15:12.353Z")
	writeStatusModel.ReasonForLastFailure = core.StringPtr("Provided API key could not be found")

	model := new(logsrouterv3.Target)
	model.ID = core.StringPtr("f7dcfae6-e7c5-08ca-451b-fdfa696c9bb6")
	model.Name = core.StringPtr("a-lr-target-us-south")
	model.CRN = core.StringPtr("crn:v1:bluemix:public:logs-router:us-south:a/0be5ad401ae913d8ff665d92680664ed:b6eec08b-5201-08ca-451b-cd71523e3626:target:f7dcfae6-e7c5-08ca-451b-fdfa696c9bb6")
	model.DestinationCRN = core.StringPtr("crn:v1:bluemix:public:logs:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::")
	model.TargetType = core.StringPtr("cloud_logs")
	model.Region = core.StringPtr("us-south")
	model.WriteStatus = writeStatusModel
	model.CreatedAt = CreateMockDateTime("2021-05-18T20:15:12.353Z")
	model.UpdatedAt = CreateMockDateTime("2021-05-18T20:15:12.353Z")
	model.ManagedBy = core.StringPtr("enterprise")

	result, err := logsrouter.DataSourceIBMLogsRouterTargetsTargetToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMLogsRouterTargetsWriteStatusToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["status"] = "success"
		model["last_failure"] = "2025-05-18T20:15:12.353Z"
		model["reason_for_last_failure"] = "Provided API key could not be found"

		assert.Equal(t, result, model)
	}

	model := new(logsrouterv3.WriteStatus)
	model.Status = core.StringPtr("success")
	model.LastFailure = CreateMockDateTime("2025-05-18T20:15:12.353Z")
	model.ReasonForLastFailure = core.StringPtr("Provided API key could not be found")

	result, err := logsrouter.DataSourceIBMLogsRouterTargetsWriteStatusToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
