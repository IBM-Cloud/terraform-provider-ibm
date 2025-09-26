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

func TestAccIbmPdrValidateWorkspaceDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmPdrValidateWorkspaceDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_pdr_validate_workspace.pdr_validate_workspace_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_pdr_validate_workspace.pdr_validate_workspace_instance", "instance_id"),
					resource.TestCheckResourceAttrSet("data.ibm_pdr_validate_workspace.pdr_validate_workspace_instance", "workspace_id"),
					resource.TestCheckResourceAttrSet("data.ibm_pdr_validate_workspace.pdr_validate_workspace_instance", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_pdr_validate_workspace.pdr_validate_workspace_instance", "location_url"),
				),
			},
		},
	})
}

func testAccCheckIbmPdrValidateWorkspaceDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_pdr_validate_workspace" "pdr_validate_workspace_instance" {
			instance_id = "crn:v1:staging:public:power-dr-automation:global:a/a123456fb04ceebfb4a9fd38c22334455:123456d3-1122-3344-b67d-4389b44b7bf9::"
			workspace_id = "75cbf05b-78f6-406e-afe7-a904f646d798"
			crn = "crn:v1:bluemix:public:power-iaas:dal10:a/094f4214c75941f991da601b001df1fe:75cbf05b-78f6-406e-afe7-a904f646d798::"
			location_url = "https://us-south.power-iaas.cloud.ibm.com"
		}
	`)
}
