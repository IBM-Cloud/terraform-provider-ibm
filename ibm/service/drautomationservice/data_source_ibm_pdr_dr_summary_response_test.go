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

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/drautomationservice"
	. "github.com/IBM-Cloud/terraform-provider-ibm/ibm/unittest"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/stretchr/testify/assert"
	"github.ibm.com/DRAutomation/dra-go-sdk/drautomationservicev1"
)

func TestAccIBMPdrDrSummaryResponseDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMPdrDrSummaryResponseDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_pdr_dr_summary_response.pdr_dr_summary_response_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_pdr_dr_summary_response.pdr_dr_summary_response_instance", "instance_id"),
					resource.TestCheckResourceAttrSet("data.ibm_pdr_dr_summary_response.pdr_dr_summary_response_instance", "managed_vm_list.%"),
					resource.TestCheckResourceAttrSet("data.ibm_pdr_dr_summary_response.pdr_dr_summary_response_instance", "orchestrator_details.#"),
					resource.TestCheckResourceAttrSet("data.ibm_pdr_dr_summary_response.pdr_dr_summary_response_instance", "service_details.#"),
				),
			},
		},
	})
}

func testAccCheckIBMPdrDrSummaryResponseDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_pdr_dr_summary_response" "pdr_dr_summary_response_instance" {
			instance_id = "123456d3-1122-3344-b67d-4389b44b7bf9"
			Accept-Language = "en-US"
		}
	`)
}

func TestDataSourceIBMPdrDrSummaryResponseServiceDetailsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["status"] = "testString"
		model["description"] = "testString"
		model["orchestrator_ha"] = true
		model["deployment_name"] = "testString"
		model["recovery_location"] = "testString"
		model["resource_group"] = "testString"
		model["crn"] = "testString"
		model["primary_ip_address"] = "testString"
		model["standby_ip_address"] = "testString"
		model["standby_description"] = "testString"
		model["standby_status"] = "testString"
		model["primary_orchestrator_dashboard_url"] = "testString"
		model["standby_orchestrator_dashboard_url"] = "testString"
		model["plan_name"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(drautomationservicev1.ServiceDetails)
	model.Status = core.StringPtr("testString")
	model.Description = core.StringPtr("testString")
	model.OrchestratorHa = core.BoolPtr(true)
	model.DeploymentName = core.StringPtr("testString")
	model.RecoveryLocation = core.StringPtr("testString")
	model.ResourceGroup = core.StringPtr("testString")
	model.CRN = core.StringPtr("testString")
	model.PrimaryIPAddress = core.StringPtr("testString")
	model.StandbyIPAddress = core.StringPtr("testString")
	model.StandbyDescription = core.StringPtr("testString")
	model.StandbyStatus = core.StringPtr("testString")
	model.PrimaryOrchestratorDashboardURL = core.StringPtr("testString")
	model.StandbyOrchestratorDashboardURL = core.StringPtr("testString")
	model.PlanName = core.StringPtr("testString")

	result, err := drautomationservice.DataSourceIBMPdrDrSummaryResponseServiceDetailsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMPdrDrSummaryResponseOrchestratorDetailsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["orchestrator_name"] = "testString"
		model["orchestrator_status"] = "testString"
		model["ssh_key_name"] = "testString"
		model["standby_ssh_key_name"] = "testString"
		model["schematic_workspace_status"] = "testString"
		model["schematic_workspace_name"] = "testString"
		model["standby_orchestrator_name"] = "testString"
		model["standby_orchestrator_status"] = "testString"
		model["orchestrator_config_status"] = "testString"
		model["orchestrator_group_leader"] = "testString"
		model["orchestrator_cluster_message"] = "testString"
		model["orch_standby_node_addition_status"] = "testString"
		model["orchestrator_location_type"] = "testString"
		model["location_id"] = "testString"
		model["vpc_name"] = "testString"
		model["transit_gateway_name"] = "testString"
		model["proxy_ip"] = "testString"
		model["orchestrator_workspace_name"] = "testString"
		model["standby_orchestrator_workspace_name"] = "testString"
		model["orch_ext_connectivity_status"] = "testString"
		model["last_updated_orchestrator_deployment_time"] = "2025-10-16T09:28:13.696Z"
		model["last_updated_standby_orchestrator_deployment_time"] = "2025-10-16T09:28:13.696Z"
		model["latest_orchestrator_time"] = "2025-10-16T09:28:13.696Z"
		model["mfa_enabled"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(drautomationservicev1.OrchestratorDetails)
	model.OrchestratorName = core.StringPtr("testString")
	model.OrchestratorStatus = core.StringPtr("testString")
	model.SSHKeyName = core.StringPtr("testString")
	model.StandbySSHKeyName = core.StringPtr("testString")
	model.SchematicWorkspaceStatus = core.StringPtr("testString")
	model.SchematicWorkspaceName = core.StringPtr("testString")
	model.StandbyOrchestratorName = core.StringPtr("testString")
	model.StandbyOrchestratorStatus = core.StringPtr("testString")
	model.OrchestratorConfigStatus = core.StringPtr("testString")
	model.OrchestratorGroupLeader = core.StringPtr("testString")
	model.OrchestratorClusterMessage = core.StringPtr("testString")
	model.OrchStandbyNodeAdditionStatus = core.StringPtr("testString")
	model.OrchestratorLocationType = core.StringPtr("testString")
	model.LocationID = core.StringPtr("testString")
	model.VPCName = core.StringPtr("testString")
	model.TransitGatewayName = core.StringPtr("testString")
	model.ProxyIP = core.StringPtr("testString")
	model.OrchestratorWorkspaceName = core.StringPtr("testString")
	model.StandbyOrchestratorWorkspaceName = core.StringPtr("testString")
	model.OrchExtConnectivityStatus = core.StringPtr("testString")
	model.LastUpdatedOrchestratorDeploymentTime = CreateMockDateTime("2025-10-16T09:28:13.696Z")
	model.LastUpdatedStandbyOrchestratorDeploymentTime = CreateMockDateTime("2025-10-16T09:28:13.696Z")
	model.LatestOrchestratorTime = CreateMockDateTime("2025-10-16T09:28:13.696Z")
	model.MfaEnabled = core.StringPtr("testString")

	result, err := drautomationservice.DataSourceIBMPdrDrSummaryResponseOrchestratorDetailsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
