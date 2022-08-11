// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package atracker_test

import (
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMAtrackerEndpointsDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMAtrackerEndpointsDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_atracker_endpoints.atracker_endpoints", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_atracker_endpoints.atracker_endpoints", "api_endpoint.#"),
				),
			},
		},
	})
}

func testAccCheckIBMAtrackerEndpointsDataSourceConfigBasic() string {
	return `
		data "ibm_atracker_endpoints" "atracker_endpoints" {
		}
	`
}
