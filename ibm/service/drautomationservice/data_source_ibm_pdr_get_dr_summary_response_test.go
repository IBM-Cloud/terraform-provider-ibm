// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.105.0-3c13b041-20250605-193116
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

func TestAccIBMPdrGetDrSummaryResponseDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMPdrGetDrSummaryResponseDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_pdr_get_dr_summary_response.pdr_get_dr_summary_response_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_pdr_get_dr_summary_response.pdr_get_dr_summary_response_instance", "instance_id"),
					resource.TestCheckResourceAttrSet("data.ibm_pdr_get_dr_summary_response.pdr_get_dr_summary_response_instance", "managed_vm_list.%"),
					resource.TestCheckResourceAttrSet("data.ibm_pdr_get_dr_summary_response.pdr_get_dr_summary_response_instance", "orchestrator_details.#"),
					resource.TestCheckResourceAttrSet("data.ibm_pdr_get_dr_summary_response.pdr_get_dr_summary_response_instance", "service_details.#"),
				),
			},
		},
	})
}

func testAccCheckIBMPdrGetDrSummaryResponseDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_pdr_get_dr_summary_response" "pdr_get_dr_summary_response_instance" {
			instance_id = "crn:v1:staging:public:power-dr-automation:global:a/a123456fb04ceebfb4a9fd38c22334455:123456d3-1122-3344-b67d-4389b44b7bf9::"
		}
	`)
}

func TestDataSourceIBMPdrGetDrSummaryResponseOrchestratorDetailsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["last_updated_orchestrator_deployment_time"] = "2025-10-16T09:28:13.696Z"
		model["last_updated_standby_orchestrator_deployment_time"] = "2025-10-16T09:28:13.696Z"
		model["latest_orchestrator_time"] = "2025-10-16T09:28:13.696Z"
		model["location_id"] = "testString"
		model["mfa_enabled"] = "testString"
		model["orch_ext_connectivity_status"] = "testString"
		model["orch_standby_node_addition_status"] = "testString"
		model["orchestrator_cluster_message"] = "testString"
		model["orchestrator_config_status"] = "testString"
		model["orchestrator_group_leader"] = "testString"
		model["orchestrator_location_type"] = "testString"
		model["orchestrator_name"] = "testString"
		model["orchestrator_status"] = "testString"
		model["orchestrator_workspace_name"] = "testString"
		model["proxy_ip"] = "testString"
		model["schematic_workspace_name"] = "testString"
		model["schematic_workspace_status"] = "testString"
		model["ssh_key_name"] = "testString"
		model["standby_orchestrator_name"] = "testString"
		model["standby_orchestrator_status"] = "testString"
		model["standby_orchestrator_workspace_name"] = "testString"
		model["transit_gateway_name"] = "testString"
		model["vpc_name"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(drautomationservicev1.OrchestratorDetails)
	model.LastUpdatedOrchestratorDeploymentTime = CreateMockDateTime("2025-10-16T09:28:13.696Z")
	model.LastUpdatedStandbyOrchestratorDeploymentTime = CreateMockDateTime("2025-10-16T09:28:13.696Z")
	model.LatestOrchestratorTime = CreateMockDateTime("2025-10-16T09:28:13.696Z")
	model.LocationID = core.StringPtr("testString")
	model.MfaEnabled = core.StringPtr("testString")
	model.OrchExtConnectivityStatus = core.StringPtr("testString")
	model.OrchStandbyNodeAdditionStatus = core.StringPtr("testString")
	model.OrchestratorClusterMessage = core.StringPtr("testString")
	model.OrchestratorConfigStatus = core.StringPtr("testString")
	model.OrchestratorGroupLeader = core.StringPtr("testString")
	model.OrchestratorLocationType = core.StringPtr("testString")
	model.OrchestratorName = core.StringPtr("testString")
	model.OrchestratorStatus = core.StringPtr("testString")
	model.OrchestratorWorkspaceName = core.StringPtr("testString")
	model.ProxyIP = core.StringPtr("testString")
	model.SchematicWorkspaceName = core.StringPtr("testString")
	model.SchematicWorkspaceStatus = core.StringPtr("testString")
	model.SSHKeyName = core.StringPtr("testString")
	model.StandbyOrchestratorName = core.StringPtr("testString")
	model.StandbyOrchestratorStatus = core.StringPtr("testString")
	model.StandbyOrchestratorWorkspaceName = core.StringPtr("testString")
	model.TransitGatewayName = core.StringPtr("testString")
	model.VPCName = core.StringPtr("testString")

	result, err := drautomationservice.DataSourceIBMPdrGetDrSummaryResponseOrchestratorDetailsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMPdrGetDrSummaryResponseServiceDetailsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["crn"] = "testString"
		model["deployment_name"] = "testString"
		model["description"] = "testString"
		model["orchestrator_ha"] = true
		model["plan_name"] = "testString"
		model["primary_ip_address"] = "testString"
		model["primary_orchestrator_dashboard_url"] = "testString"
		model["recovery_location"] = "testString"
		model["resource_group"] = "testString"
		model["standby_description"] = "testString"
		model["standby_ip_address"] = "testString"
		model["standby_orchestrator_dashboard_url"] = "testString"
		model["standby_status"] = "testString"
		model["status"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(drautomationservicev1.ServiceDetails)
	model.CRN = core.StringPtr("testString")
	model.DeploymentName = core.StringPtr("testString")
	model.Description = core.StringPtr("testString")
	model.OrchestratorHa = core.BoolPtr(true)
	model.PlanName = core.StringPtr("testString")
	model.PrimaryIPAddress = core.StringPtr("testString")
	model.PrimaryOrchestratorDashboardURL = core.StringPtr("testString")
	model.RecoveryLocation = core.StringPtr("testString")
	model.ResourceGroup = core.StringPtr("testString")
	model.StandbyDescription = core.StringPtr("testString")
	model.StandbyIPAddress = core.StringPtr("testString")
	model.StandbyOrchestratorDashboardURL = core.StringPtr("testString")
	model.StandbyStatus = core.StringPtr("testString")
	model.Status = core.StringPtr("testString")

	result, err := drautomationservice.DataSourceIBMPdrGetDrSummaryResponseServiceDetailsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
