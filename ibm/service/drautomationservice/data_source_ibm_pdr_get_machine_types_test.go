// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.108.0-56772134-20251111-102802
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
			instance_id = "xxxx2ec4-xxxx-4f84-xxxx-c2aa834dd4ed"
			primary_workspace_name = "Test-workspace-wdc06"
			standby_workspace_name = "Test-workspace-wdc07"
		}
	`)
}
