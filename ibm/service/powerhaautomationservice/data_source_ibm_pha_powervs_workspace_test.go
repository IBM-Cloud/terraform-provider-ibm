// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.116.0-df613dbc-20260803-154903
*/

package powerhaautomationservice_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/powerhaautomationservice"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/stretchr/testify/assert"
	"github.ibm.com/DRAutomation/dra-go-sdk/powerhaautomationservicev1"
	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMPhaPowervsWorkspaceDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMPhaPowervsWorkspaceDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_pha_powervs_workspace.pha_powervs_workspace_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_powervs_workspace.pha_powervs_workspace_instance", "pha_instance_id"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_powervs_workspace.pha_powervs_workspace_instance", "location_id"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_powervs_workspace.pha_powervs_workspace_instance", "workspaces.#"),
				),
			},
		},
	})
}

func testAccCheckIBMPhaPowervsWorkspaceDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_pha_powervs_workspace" "pha_powervs_workspace_instance" {
			pha_instance_id = "8eefautr-4c02-0009-0086-8bd4d8cf61b6"
			location_id = "us-south"
			Accept-Language = "en-US"
			If-None-Match = "abcdef"
		}
	`)
}


func TestDataSourceIBMPhaPowervsWorkspacePhaWorkspaceSummaryToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["id"] = "09845"
		model["name"] = "Primary_PowerHA_Workspace"

		assert.Equal(t, result, model)
	}

	model := new(powerhaautomationservicev1.PhaWorkspaceSummary)
	model.ID = core.StringPtr("09845")
	model.Name = core.StringPtr("Primary_PowerHA_Workspace")

	result, err := powerhaautomationservice.DataSourceIBMPhaPowervsWorkspacePhaWorkspaceSummaryToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
