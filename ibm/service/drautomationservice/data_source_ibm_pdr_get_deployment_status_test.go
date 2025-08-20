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
)

func TestAccIbmPdrGetDeploymentStatusDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmPdrGetDeploymentStatusDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_pdr_get_deployment_status.pdr_get_deployment_status_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_pdr_get_deployment_status.pdr_get_deployment_status_instance", "instance_id"),
					resource.TestCheckResourceAttrSet("data.ibm_pdr_get_deployment_status.pdr_get_deployment_status_instance", "orch_standby_node_addition_status"),
					resource.TestCheckResourceAttrSet("data.ibm_pdr_get_deployment_status.pdr_get_deployment_status_instance", "orchestrator_cluster_message"),
					resource.TestCheckResourceAttrSet("data.ibm_pdr_get_deployment_status.pdr_get_deployment_status_instance", "orchestrator_cluster_type"),
					resource.TestCheckResourceAttrSet("data.ibm_pdr_get_deployment_status.pdr_get_deployment_status_instance", "orchestrator_config_status"),
					resource.TestCheckResourceAttrSet("data.ibm_pdr_get_deployment_status.pdr_get_deployment_status_instance", "orchestrator_group_leader"),
					resource.TestCheckResourceAttrSet("data.ibm_pdr_get_deployment_status.pdr_get_deployment_status_instance", "orchestrator_name"),
					resource.TestCheckResourceAttrSet("data.ibm_pdr_get_deployment_status.pdr_get_deployment_status_instance", "orchestrator_status"),
					resource.TestCheckResourceAttrSet("data.ibm_pdr_get_deployment_status.pdr_get_deployment_status_instance", "schematic_workspace_name"),
					resource.TestCheckResourceAttrSet("data.ibm_pdr_get_deployment_status.pdr_get_deployment_status_instance", "schematic_workspace_status"),
					resource.TestCheckResourceAttrSet("data.ibm_pdr_get_deployment_status.pdr_get_deployment_status_instance", "ssh_key_name"),
					resource.TestCheckResourceAttrSet("data.ibm_pdr_get_deployment_status.pdr_get_deployment_status_instance", "standby_orchestrator_name"),
					resource.TestCheckResourceAttrSet("data.ibm_pdr_get_deployment_status.pdr_get_deployment_status_instance", "standby_orchestrator_status"),
				),
			},
		},
	})
}

func testAccCheckIbmPdrGetDeploymentStatusDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_pdr_get_deployment_status" "pdr_get_deployment_status_instance" {
			instance_id = "crn:v1:staging:public:power-dr-automation:global:a/a123456fb04ceebfb4a9fd38c22334455:123456d3-1122-3344-b67d-4389b44b7bf9::"
			If-None-Match = "If-None-Match"
		}
	`)
}
