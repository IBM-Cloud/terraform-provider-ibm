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

func TestAccIBMPhaSupportedLocationDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMPhaSupportedLocationDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_pha_supported_locations.pha_supported_location_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_supported_locations.pha_supported_location_instance", "instance_id"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_supported_locations.pha_supported_location_instance", "locations.#"),
				),
			},
		},
	})
}

func testAccCheckIBMPhaSupportedLocationDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_pha_supported_locations" "pha_supported_location_instance" {
			instance_id = "8eefautr-4c02-0009-0086-8bd4d8cf61b6"
			If-None-Match = "abcdef"
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
