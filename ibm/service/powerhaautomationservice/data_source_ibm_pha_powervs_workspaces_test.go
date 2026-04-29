// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.113.1-d76630af-20260320-135953
 */

package powerhaautomationservice_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/powerhaautomationservice"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/stretchr/testify/assert"
	"github.com/IBM/dra-go-sdk/powerhaautomationservicev1"
)

func TestAccIBMPhaGetPowervsWorkspaceDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMPhaPowervsWorkspaceDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_pha_powervs_workspaces.pha_powervs_workspace_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_powervs_workspaces.pha_powervs_workspace_instance", "pha_instance_id"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_powervs_workspaces.pha_powervs_workspace_instance", "location_id"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_powervs_workspaces.pha_powervs_workspace_instance", "workspaces.#"),
				),
			},
		},
	})
}

func testAccCheckIBMPhaPowervsWorkspaceDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_pha_powervs_workspaces" "pha_powervs_workspace_instance" {
			instance_id = "8eefautr-4c02-0009-0086-8bd4d8cf61b6"
			location_id = "us-south"
			Accept-Language = "en-US"
			If-None-Match = "abcdef"
		}
	`)
}

func TestDataSourceIBMPhaGetPowervsWorkspacePhaWorkspaceSummaryToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["id"] = "ws-001"
		model["name"] = "primary-workspace"

		assert.Equal(t, result, model)
	}

	model := new(powerhaautomationservicev1.PhaWorkspaceSummary)
	model.ID = core.StringPtr("ws-001")
	model.Name = core.StringPtr("primary-workspace")

	result, err := powerhaautomationservice.DataSourceIBMPhaGetPowervsWorkspacePhaWorkspaceSummaryToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
