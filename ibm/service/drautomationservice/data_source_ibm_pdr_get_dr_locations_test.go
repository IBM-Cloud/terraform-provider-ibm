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

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/drautomationservice"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/stretchr/testify/assert"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM/dra-go-sdk/drautomationservicev1"
)

func TestAccIBMPdrGetDrLocationsDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMPdrGetDrLocationsDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_pdr_get_dr_locations.pdr_get_dr_locations_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_pdr_get_dr_locations.pdr_get_dr_locations_instance", "instance_id"),
					resource.TestCheckResourceAttrSet("data.ibm_pdr_get_dr_locations.pdr_get_dr_locations_instance", "dr_locations.#"),
				),
			},
		},
	})
}

func testAccCheckIBMPdrGetDrLocationsDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_pdr_get_dr_locations" "pdr_get_dr_locations_instance" {
			instance_id = "ac645fe5-fba1-4cb3-952e-e1b09fa0df26"
		}
	`)
}

func TestDataSourceIBMPdrGetDrLocationsDrLocationToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["id"] = "loc123"
		model["name"] = "US-East-1"

		assert.Equal(t, result, model)
	}

	model := new(drautomationservicev1.DrLocation)
	model.ID = core.StringPtr("loc123")
	model.Name = core.StringPtr("US-East-1")

	result, err := drautomationservice.DataSourceIBMPdrGetDrLocationsDrLocationToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
