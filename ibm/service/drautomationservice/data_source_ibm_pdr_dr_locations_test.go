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

func TestAccIBMPdrDrLocationsDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMPdrDrLocationsDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_pdr_dr_locations.pdr_dr_locations_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_pdr_dr_locations.pdr_dr_locations_instance", "instance_id"),
					resource.TestCheckResourceAttrSet("data.ibm_pdr_dr_locations.pdr_dr_locations_instance", "dr_locations.#"),
				),
			},
		},
	})
}

func testAccCheckIBMPdrDrLocationsDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_pdr_dr_locations" "pdr_dr_locations_instance" {
			instance_id = "123456d3-1122-3344-b67d-4389b44b7bf9"
			Accept-Language = "en-US"
		}
	`)
}

func TestDataSourceIBMPdrDrLocationsDrLocationToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["id"] = "loc123"
		model["name"] = "US-East-1"
		model["power_edge_router"] = false

		assert.Equal(t, result, model)
	}

	model := new(drautomationservicev1.DrLocation)
	model.ID = core.StringPtr("loc123")
	model.Name = core.StringPtr("US-East-1")
	model.PowerEdgeRouter = core.BoolPtr(false)

	result, err := drautomationservice.DataSourceIBMPdrDrLocationsDrLocationToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
