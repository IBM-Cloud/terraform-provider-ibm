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

func TestAccIBMPdrLastOperationDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMPdrLastOperationDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_pdr_last_operation.pdr_last_operation_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_pdr_last_operation.pdr_last_operation_instance", "instance_id"),
					resource.TestCheckResourceAttrSet("data.ibm_pdr_last_operation.pdr_last_operation_instance", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_pdr_last_operation.pdr_last_operation_instance", "deployment_name"),
					resource.TestCheckResourceAttrSet("data.ibm_pdr_last_operation.pdr_last_operation_instance", "is_ksys_ha"),
					resource.TestCheckResourceAttrSet("data.ibm_pdr_last_operation.pdr_last_operation_instance", "last_updated_orchestrator_deployment_time"),
					resource.TestCheckResourceAttrSet("data.ibm_pdr_last_operation.pdr_last_operation_instance", "last_updated_standby_orchestrator_deployment_time"),
					resource.TestCheckResourceAttrSet("data.ibm_pdr_last_operation.pdr_last_operation_instance", "mfa_enabled"),
					resource.TestCheckResourceAttrSet("data.ibm_pdr_last_operation.pdr_last_operation_instance", "orch_standby_node_addtion_status"),
					resource.TestCheckResourceAttrSet("data.ibm_pdr_last_operation.pdr_last_operation_instance", "orchestrator_cluster_message"),
					resource.TestCheckResourceAttrSet("data.ibm_pdr_last_operation.pdr_last_operation_instance", "orchestrator_config_status"),
					resource.TestCheckResourceAttrSet("data.ibm_pdr_last_operation.pdr_last_operation_instance", "plan_name"),
					resource.TestCheckResourceAttrSet("data.ibm_pdr_last_operation.pdr_last_operation_instance", "primary_description"),
					resource.TestCheckResourceAttrSet("data.ibm_pdr_last_operation.pdr_last_operation_instance", "primary_ip_address"),
					resource.TestCheckResourceAttrSet("data.ibm_pdr_last_operation.pdr_last_operation_instance", "primary_orchestrator_status"),
					resource.TestCheckResourceAttrSet("data.ibm_pdr_last_operation.pdr_last_operation_instance", "recovery_location"),
					resource.TestCheckResourceAttrSet("data.ibm_pdr_last_operation.pdr_last_operation_instance", "resource_group"),
					resource.TestCheckResourceAttrSet("data.ibm_pdr_last_operation.pdr_last_operation_instance", "standby_description"),
					resource.TestCheckResourceAttrSet("data.ibm_pdr_last_operation.pdr_last_operation_instance", "standby_ip_address"),
					resource.TestCheckResourceAttrSet("data.ibm_pdr_last_operation.pdr_last_operation_instance", "standby_status"),
					resource.TestCheckResourceAttrSet("data.ibm_pdr_last_operation.pdr_last_operation_instance", "status"),
				),
			},
		},
	})
}

func testAccCheckIBMPdrLastOperationDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_pdr_last_operation" "pdr_last_operation_instance" {
			instance_id = "crn:v1:staging:public:power-dr-automation:global:a/a123456fb04ceebfb4a9fd38c22334455:123456d3-1122-3344-b67d-4389b44b7bf9::"
		}
	`)
}

