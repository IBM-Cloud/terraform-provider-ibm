// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMAtrackerEndpointsDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
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
	return fmt.Sprintf(`
		data "ibm_atracker_endpoints" "atracker_endpoints" {
		}
	`)
}
