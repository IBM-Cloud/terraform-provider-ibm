// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.116.0-df613dbc-20260803-154903
*/

package drautomationservice_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/drautomationservice"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/stretchr/testify/assert"
	"github.ibm.com/DRAutomation/dra-go-sdk/drautomationservicev1"
	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMPdrIBMMaintainedOrchDetailsDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMPdrIBMMaintainedOrchDetailsDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_pdr_ibm_maintained_orch_details.pdr_ibm_maintained_orch_details_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_pdr_ibm_maintained_orch_details.pdr_ibm_maintained_orch_details_instance", "instance_id"),
				),
			},
		},
	})
}

func testAccCheckIBMPdrIBMMaintainedOrchDetailsDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_pdr_ibm_maintained_orch_details" "pdr_ibm_maintained_orch_details_instance" {
			instance_id = "123456d3-1122-3344-b67d-4389b44b7bf9"
			Accept-Language = "en-US"
		}
	`)
}


func TestDataSourceIBMPdrIBMMaintainedOrchDetailsIBMMaintainedOrchServiceDetailsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["status"] = "testString"
		model["deployment_name"] = "testString"
		model["resource_group"] = "testString"
		model["crn"] = "testString"
		model["plan_name"] = "testString"
		model["region"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(drautomationservicev1.IBMMaintainedOrchServiceDetails)
	model.Status = core.StringPtr("testString")
	model.DeploymentName = core.StringPtr("testString")
	model.ResourceGroup = core.StringPtr("testString")
	model.CRN = core.StringPtr("testString")
	model.PlanName = core.StringPtr("testString")
	model.Region = core.StringPtr("testString")

	result, err := drautomationservice.DataSourceIBMPdrIBMMaintainedOrchDetailsIBMMaintainedOrchServiceDetailsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMPdrIBMMaintainedOrchDetailsIBMMaintainedOrchOrchestratorDetailsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["orchestrator_gui_status"] = "testString"
		model["orchestrator_name"] = "testString"
		model["orchestrator_status"] = "testString"
		model["orchestrator_ip"] = "testString"
		model["orchestrator_id"] = "testString"
		model["orchestrator_gui_url"] = "testString"
		model["orchestrator_cluster_status"] = "testString"
		model["orchestrator_cluster_config_message"] = "testString"
		model["orchestrator_location_type"] = "testString"
		model["orchestrator_workspace_name"] = "testString"
		model["orchestrator_description"] = "2/5: Creating orchestrator VM."
		model["orchestrator_username"] = "testString"
		model["standby_orchestrator_gui_status"] = "testString"
		model["standby_orchestrator_name"] = "testString"
		model["standby_orchestrator_id"] = "testString"
		model["standby_orchestrator_ip"] = "testString"
		model["standby_orchestrator_status"] = "testString"
		model["standby_orchestrator_description"] = "testString"
		model["standby_orchestrator_gui_url"] = "testString"
		model["standby_orchestrator_username"] = "testString"
		model["standby_orchestrator_node_addition_status"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(drautomationservicev1.IBMMaintainedOrchOrchestratorDetails)
	model.OrchestratorGuiStatus = core.StringPtr("testString")
	model.OrchestratorName = core.StringPtr("testString")
	model.OrchestratorStatus = core.StringPtr("testString")
	model.OrchestratorIP = core.StringPtr("testString")
	model.OrchestratorID = core.StringPtr("testString")
	model.OrchestratorGuiURL = core.StringPtr("testString")
	model.OrchestratorClusterStatus = core.StringPtr("testString")
	model.OrchestratorClusterConfigMessage = core.StringPtr("testString")
	model.OrchestratorLocationType = core.StringPtr("testString")
	model.OrchestratorWorkspaceName = core.StringPtr("testString")
	model.OrchestratorDescription = core.StringPtr("2/5: Creating orchestrator VM.")
	model.OrchestratorUsername = core.StringPtr("testString")
	model.StandbyOrchestratorGuiStatus = core.StringPtr("testString")
	model.StandbyOrchestratorName = core.StringPtr("testString")
	model.StandbyOrchestratorID = core.StringPtr("testString")
	model.StandbyOrchestratorIP = core.StringPtr("testString")
	model.StandbyOrchestratorStatus = core.StringPtr("testString")
	model.StandbyOrchestratorDescription = core.StringPtr("testString")
	model.StandbyOrchestratorGuiURL = core.StringPtr("testString")
	model.StandbyOrchestratorUsername = core.StringPtr("testString")
	model.StandbyOrchestratorNodeAdditionStatus = core.StringPtr("testString")

	result, err := drautomationservice.DataSourceIBMPdrIBMMaintainedOrchDetailsIBMMaintainedOrchOrchestratorDetailsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
