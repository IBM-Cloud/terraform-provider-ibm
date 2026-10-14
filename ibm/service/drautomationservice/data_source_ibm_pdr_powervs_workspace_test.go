// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.116.0-df613dbc-20260803-154903
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

func TestAccIBMPdrPowervsWorkspaceDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMPdrPowervsWorkspaceDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_pdr_powervs_workspace.pdr_powervs_workspace_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_pdr_powervs_workspace.pdr_powervs_workspace_instance", "instance_id"),
					resource.TestCheckResourceAttrSet("data.ibm_pdr_powervs_workspace.pdr_powervs_workspace_instance", "location_id"),
					resource.TestCheckResourceAttrSet("data.ibm_pdr_powervs_workspace.pdr_powervs_workspace_instance", "dr_workspaces.#"),
					resource.TestCheckResourceAttrSet("data.ibm_pdr_powervs_workspace.pdr_powervs_workspace_instance", "dr_standby_workspaces.#"),
				),
			},
		},
	})
}

func testAccCheckIBMPdrPowervsWorkspaceDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_pdr_powervs_workspace" "pdr_powervs_workspace_instance" {
			instance_id = "123456d3-1122-3344-b67d-4389b44b7bf9"
			location_id = "us-south"
		}
	`)
}

func TestDataSourceIBMPdrPowervsWorkspaceDrWorkspaceToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		detailsDrModel := make(map[string]interface{})
		detailsDrModel["crn"] = "crn:v1:abc123:public:power-iaas:us-south:resource/path-123:instance-456::"

		locationDrModel := make(map[string]interface{})
		locationDrModel["region"] = "lon06"
		locationDrModel["type"] = "data-center"
		locationDrModel["url"] = "https://lon.power-iaas.cloud.ibm.com"

		model := make(map[string]interface{})
		model["details"] = []map[string]interface{}{detailsDrModel}
		model["id"] = "testString"
		model["location"] = []map[string]interface{}{locationDrModel}
		model["name"] = "testString"
		model["status"] = "active"
		model["default"] = true

		assert.Equal(t, result, model)
	}

	detailsDrModel := new(drautomationservicev1.DetailsDr)
	detailsDrModel.CRN = core.StringPtr("crn:v1:abc123:public:power-iaas:us-south:resource/path-123:instance-456::")

	locationDrModel := new(drautomationservicev1.LocationDr)
	locationDrModel.Region = core.StringPtr("lon06")
	locationDrModel.Type = core.StringPtr("data-center")
	locationDrModel.URL = core.StringPtr("https://lon.power-iaas.cloud.ibm.com")

	model := new(drautomationservicev1.DrWorkspace)
	model.Details = detailsDrModel
	model.ID = core.StringPtr("testString")
	model.Location = locationDrModel
	model.Name = core.StringPtr("testString")
	model.Status = core.StringPtr("active")
	model.Default = core.BoolPtr(true)

	result, err := drautomationservice.DataSourceIBMPdrPowervsWorkspaceDrWorkspaceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMPdrPowervsWorkspaceDetailsDrToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["crn"] = "crn:v1:abc123:public:power-iaas:us-south:resource/path-123:instance-456::"

		assert.Equal(t, result, model)
	}

	model := new(drautomationservicev1.DetailsDr)
	model.CRN = core.StringPtr("crn:v1:abc123:public:power-iaas:us-south:resource/path-123:instance-456::")

	result, err := drautomationservice.DataSourceIBMPdrPowervsWorkspaceDetailsDrToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMPdrPowervsWorkspaceLocationDrToMap(t *testing.T) {
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

	result, err := drautomationservice.DataSourceIBMPdrPowervsWorkspaceLocationDrToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMPdrPowervsWorkspaceDrStandbyWorkspaceToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		detailsDrModel := make(map[string]interface{})
		detailsDrModel["crn"] = "crn:v1:abc123:public:power-iaas:us-south:resource/path-123:instance-456::"

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
	detailsDrModel.CRN = core.StringPtr("crn:v1:abc123:public:power-iaas:us-south:resource/path-123:instance-456::")

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

	result, err := drautomationservice.DataSourceIBMPdrPowervsWorkspaceDrStandbyWorkspaceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
