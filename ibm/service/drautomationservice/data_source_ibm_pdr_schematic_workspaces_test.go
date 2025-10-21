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

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/drautomationservice"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/stretchr/testify/assert"
	"github.ibm.com/DRAutomation/dra-go-sdk/drautomationservicev1"
	. "github.com/IBM-Cloud/terraform-provider-ibm/ibm/unittest"
	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmPdrSchematicWorkspacesDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmPdrSchematicWorkspacesDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_pdr_schematic_workspaces.pdr_schematic_workspaces_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_pdr_schematic_workspaces.pdr_schematic_workspaces_instance", "instance_id"),
					resource.TestCheckResourceAttrSet("data.ibm_pdr_schematic_workspaces.pdr_schematic_workspaces_instance", "workspaces.#"),
				),
			},
		},
	})
}

func testAccCheckIbmPdrSchematicWorkspacesDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_pdr_schematic_workspaces" "pdr_schematic_workspaces_instance" {
			instance_id = "crn:v1:staging:public:power-dr-automation:global:a/a123456fb04ceebfb4a9fd38c22334455:123456d3-1122-3344-b67d-4389b44b7bf9::"
			If-None-Match = "If-None-Match"
		}
	`)
}


func TestDataSourceIbmPdrSchematicWorkspacesDrAutomationSchematicsWorkspaceToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		drAutomationCatalogRefModel := make(map[string]interface{})
		drAutomationCatalogRefModel["item_name"] = "Power Virtual Server with VPC landing zone"

		model := make(map[string]interface{})
		model["catalog_ref"] = []map[string]interface{}{drAutomationCatalogRefModel}
		model["created_at"] = "2025-02-24T11:18:49.819Z"
		model["created_by"] = "crn:v1:bluemix:public:project:eu-gb:a/094f4214c75941f94c2w1g1b001df1fe:fbe45gs9-d3c6-hny3-898d-ff2e6rfes257::"
		model["crn"] = "crn:v1:bluemix:public:schematics:eu-de:a/094f42141234567891da601b001df1fe:59389f45-a1d7-085g-8abe-7a28327e5574:workspace:eu-gb.workspace.projects-service.42a7ab33"
		model["description"] = "A configuration of the 54f31234-e74f-4567-a0a8-367b45658765 project"
		model["id"] = "A configuration of the fbe4a122-d3c6-4543-898d-ff2e6df74123 project"
		model["location"] = "us-south"
		model["name"] = "testWorkspace"
		model["status"] = "testString"

		assert.Equal(t, result, model)
	}

	drAutomationCatalogRefModel := new(drautomationservicev1.DrAutomationCatalogRef)
	drAutomationCatalogRefModel.ItemName = core.StringPtr("Power Virtual Server with VPC landing zone")

	model := new(drautomationservicev1.DrAutomationSchematicsWorkspace)
	model.CatalogRef = drAutomationCatalogRefModel
	model.CreatedAt = CreateMockDateTime("2025-02-24T11:18:49.819Z")
	model.CreatedBy = core.StringPtr("crn:v1:bluemix:public:project:eu-gb:a/094f4214c75941f94c2w1g1b001df1fe:fbe45gs9-d3c6-hny3-898d-ff2e6rfes257::")
	model.CRN = core.StringPtr("crn:v1:bluemix:public:schematics:eu-de:a/094f42141234567891da601b001df1fe:59389f45-a1d7-085g-8abe-7a28327e5574:workspace:eu-gb.workspace.projects-service.42a7ab33")
	model.Description = core.StringPtr("A configuration of the 54f31234-e74f-4567-a0a8-367b45658765 project")
	model.ID = core.StringPtr("A configuration of the fbe4a122-d3c6-4543-898d-ff2e6df74123 project")
	model.Location = core.StringPtr("us-south")
	model.Name = core.StringPtr("testWorkspace")
	model.Status = core.StringPtr("testString")

	result, err := drautomationservice.DataSourceIbmPdrSchematicWorkspacesDrAutomationSchematicsWorkspaceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmPdrSchematicWorkspacesDrAutomationCatalogRefToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["item_name"] = "Power Virtual Server with VPC landing zone"

		assert.Equal(t, result, model)
	}

	model := new(drautomationservicev1.DrAutomationCatalogRef)
	model.ItemName = core.StringPtr("Power Virtual Server with VPC landing zone")

	result, err := drautomationservice.DataSourceIbmPdrSchematicWorkspacesDrAutomationCatalogRefToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
