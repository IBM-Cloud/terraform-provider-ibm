// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMPIRoutesDataSource_basic(t *testing.T) {
	routesData := "data.ibm_pi_routes.routes_data"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIRoutesDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(routesData, "id"),
				),
			},
		},
	})
}

func testAccCheckIBMPIRoutesDataSourceConfig() string {
	return fmt.Sprintf(`
		data "ibm_pi_routes" "routes_data" {
			pi_cloud_instance_id = "%s"
		}`, acc.Pi_cloud_instance_id)
}
