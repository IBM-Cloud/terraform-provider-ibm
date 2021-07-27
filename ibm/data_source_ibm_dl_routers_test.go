// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMDLRoutersDataSource_basic(t *testing.T) {
	node := "data.ibm_dl_routers.test1"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMDLRoutersDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(node, "cross_connect_routers.0.capabilities.#"),
					resource.TestCheckResourceAttrSet(node, "cross_connect_routers.0.router_name"),
					resource.TestCheckResourceAttrSet(node, "cross_connect_routers.0.total_connections"),
				),
			},
		},
	})
}

func testAccCheckIBMDLRoutersDataSourceConfig() string {
	return fmt.Sprintf(`
	data "ibm_dl_routers" "test1" {
		offering_type = "dedicated"
		location_name = "dal10"
	}
	`)
}
