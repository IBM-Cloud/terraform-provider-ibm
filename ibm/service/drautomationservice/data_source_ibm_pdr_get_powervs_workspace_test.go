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
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/stretchr/testify/assert"
	"github.ibm.com/DRAutomation/dra-go-sdk/drautomationservicev1"
)

func TestAccIBMPdrGetPowervsWorkspaceDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMPdrGetPowervsWorkspaceDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_pdr_get_powervs_workspace.pdr_get_powervs_workspace_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_pdr_get_powervs_workspace.pdr_get_powervs_workspace_instance", "instance_id"),
					resource.TestCheckResourceAttrSet("data.ibm_pdr_get_powervs_workspace.pdr_get_powervs_workspace_instance", "location_id"),
					resource.TestCheckResourceAttrSet("data.ibm_pdr_get_powervs_workspace.pdr_get_powervs_workspace_instance", "dr_standby_workspaces.#"),
					resource.TestCheckResourceAttrSet("data.ibm_pdr_get_powervs_workspace.pdr_get_powervs_workspace_instance", "dr_workspaces.#"),
				),
			},
		},
	})
}

func testAccCheckIBMPdrGetPowervsWorkspaceDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_pdr_get_powervs_workspace" "pdr_get_powervs_workspace_instance" {
			instance_id = "crn:v1:staging:public:power-dr-automation:global:a/a123456fb04ceebfb4a9fd38c22334455:123456d3-1122-3344-b67d-4389b44b7bf9::"
			location_id = "location_id"
			If-None-Match = "If-None-Match"
		}
	`)
}

func TestDataSourceIBMPdrGetPowervsWorkspaceDrStandbyWorkspaceToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		detailsDrModel := make(map[string]interface{})
		detailsDrModel["crn"] = "crn:v1:bluemix:public:power-iaas:lon06:a/094f4214c75941f991da601b001df1fe:b6297e60-d0fe-4e24-8b15-276cf0645737::"

		locationDrModel := make(map[string]interface{})
		locationDrModel["region"] = "lon06"
		locationDrModel["type"] = "data-center"
		locationDrModel["url"] = "https://lon.power-iaas.cloud.ibm.com"

		model := make(map[string]interface{})
		model["details"] = []map[string]interface{}{detailsDrModel}
		model["id"] = "testString"
		model["location"] = []map[string]interface{}{locationDrModel}
		model["name"] = "testString"
		model["status"] = "testString"

		assert.Equal(t, result, model)
	}

	detailsDrModel := new(drautomationservicev1.DetailsDr)
	detailsDrModel.CRN = core.StringPtr("crn:v1:bluemix:public:power-iaas:lon06:a/094f4214c75941f991da601b001df1fe:b6297e60-d0fe-4e24-8b15-276cf0645737::")

	locationDrModel := new(drautomationservicev1.LocationDr)
	locationDrModel.Region = core.StringPtr("lon06")
	locationDrModel.Type = core.StringPtr("data-center")
	locationDrModel.URL = core.StringPtr("https://lon.power-iaas.cloud.ibm.com")

	model := new(drautomationservicev1.DrStandbyWorkspace)
	model.Details = detailsDrModel
	model.ID = core.StringPtr("testString")
	model.Location = locationDrModel
	model.Name = core.StringPtr("testString")
	model.Status = core.StringPtr("testString")

	result, err := drautomationservice.DataSourceIBMPdrGetPowervsWorkspaceDrStandbyWorkspaceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMPdrGetPowervsWorkspaceDetailsDrToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["crn"] = "crn:v1:bluemix:public:power-iaas:lon06:a/094f4214c75941f991da601b001df1fe:b6297e60-d0fe-4e24-8b15-276cf0645737::"

		assert.Equal(t, result, model)
	}

	model := new(drautomationservicev1.DetailsDr)
	model.CRN = core.StringPtr("crn:v1:bluemix:public:power-iaas:lon06:a/094f4214c75941f991da601b001df1fe:b6297e60-d0fe-4e24-8b15-276cf0645737::")

	result, err := drautomationservice.DataSourceIBMPdrGetPowervsWorkspaceDetailsDrToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMPdrGetPowervsWorkspaceLocationDrToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["region"] = "lon06"
		model["type"] = "data-center"
		model["url"] = "https://lon.power-iaas.cloud.ibm.com"

		assert.Equal(t, result, model)
	}

	model := new(drautomationservicev1.LocationDr)
	model.Region = core.StringPtr("lon06")
	model.Type = core.StringPtr("data-center")
	model.URL = core.StringPtr("https://lon.power-iaas.cloud.ibm.com")

	result, err := drautomationservice.DataSourceIBMPdrGetPowervsWorkspaceLocationDrToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMPdrGetPowervsWorkspaceDrWorkspaceToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		detailsDrModel := make(map[string]interface{})
		detailsDrModel["crn"] = "crn:v1:bluemix:public:power-iaas:lon06:a/094f4214c75941f991da601b001df1fe:b6297e60-d0fe-4e24-8b15-276cf0645737::"

		locationDrModel := make(map[string]interface{})
		locationDrModel["region"] = "lon06"
		locationDrModel["type"] = "data-center"
		locationDrModel["url"] = "https://lon.power-iaas.cloud.ibm.com"

		model := make(map[string]interface{})
		model["default"] = true
		model["details"] = []map[string]interface{}{detailsDrModel}
		model["id"] = "testString"
		model["location"] = []map[string]interface{}{locationDrModel}
		model["name"] = "testString"
		model["status"] = "active"

		assert.Equal(t, result, model)
	}

	detailsDrModel := new(drautomationservicev1.DetailsDr)
	detailsDrModel.CRN = core.StringPtr("crn:v1:bluemix:public:power-iaas:lon06:a/094f4214c75941f991da601b001df1fe:b6297e60-d0fe-4e24-8b15-276cf0645737::")

	locationDrModel := new(drautomationservicev1.LocationDr)
	locationDrModel.Region = core.StringPtr("lon06")
	locationDrModel.Type = core.StringPtr("data-center")
	locationDrModel.URL = core.StringPtr("https://lon.power-iaas.cloud.ibm.com")

	model := new(drautomationservicev1.DrWorkspace)
	model.Default = core.BoolPtr(true)
	model.Details = detailsDrModel
	model.ID = core.StringPtr("testString")
	model.Location = locationDrModel
	model.Name = core.StringPtr("testString")
	model.Status = core.StringPtr("active")

	result, err := drautomationservice.DataSourceIBMPdrGetPowervsWorkspaceDrWorkspaceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
