// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.108.0-56772134-20251111-102802
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
	"github.ibm.com/DRAutomation/dra-go-sdk/powerhaautomationservicev1"
)

func TestAccIBMPhaGetSupportedLocationDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMPhaGetSupportedLocationDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_pha_get_supported_location.pha_get_supported_location_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_get_supported_location.pha_get_supported_location_instance", "pha_instance_id"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_get_supported_location.pha_get_supported_location_instance", "dr_locations.#"),
				),
			},
		},
	})
}

func testAccCheckIBMPhaGetSupportedLocationDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_pha_get_supported_location" "pha_get_supported_location_instance" {
			pha_instance_id = "8ce2a099-a463-479a-9a1d-eedc19287a62"
		}
	`)
}

func TestDataSourceIBMPhaGetSupportedLocationPhaLocationToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["id"] = "loc-us-south-01"
		model["name"] = "Dallas (us-south)"

		assert.Equal(t, result, model)
	}

	model := new(powerhaautomationservicev1.PhaLocation)
	model.ID = core.StringPtr("loc-us-south-01")
	model.Name = core.StringPtr("Dallas (us-south)")

	result, err := powerhaautomationservice.DataSourceIBMPhaGetSupportedLocationPhaLocationToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
