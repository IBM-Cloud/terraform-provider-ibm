// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.105.0-3c13b041-20250605-193116
 */

package drautomationservice_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMPdrGetMachineTypesDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMPdrGetMachineTypesDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_pdr_get_machine_types.pdr_get_machine_types_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_pdr_get_machine_types.pdr_get_machine_types_instance", "instance_id"),
					resource.TestCheckResourceAttrSet("data.ibm_pdr_get_machine_types.pdr_get_machine_types_instance", "primary_workspace_name"),
				),
			},
		},
	})
}

func testAccCheckIBMPdrGetMachineTypesDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_pdr_get_machine_types" "pdr_get_machine_types_instance" {
			instance_id = "crn:v1:staging:public:power-dr-automation:global:a/b68c234e719144b18598ae4a7b80c44c:492fef47-3ebf-4090-b089-e9b4199878b6::"
			primary_workspace_name = "Test-workspace-wdc06"
			standby_workspace_name = "Test-workspace-wdc07"
		}
	`)
}
