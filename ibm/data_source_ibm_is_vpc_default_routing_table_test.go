/**
 * © Copyright IBM Corporation 2020. All Rights Reserved.
 *
 * Licensed under the Mozilla Public License, version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at https://mozilla.org/MPL/2.0/
 */

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMISVPCDefaultRoutingTableDataSource_basic(t *testing.T) {
	node := "data.ibm_is_vpc_default_routing_table.def_route_table"
	vpcname := fmt.Sprintf("tf-vpcname-%d", acctest.RandIntRange(100, 200))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVPCDefaultRoutingTableDataSourceConfig(vpcname),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(node, "id"),
					resource.TestCheckResourceAttrSet(node, "name"),
					resource.TestCheckResourceAttrSet(node, "lifecycle_state"),
				),
			},
		},
	})
}

func testAccCheckIBMISVPCDefaultRoutingTableDataSourceConfig(vpcname string) string {
	return fmt.Sprintf(`

	resource "ibm_is_vpc" "test_vpc" {
  		name = "%s"
	}
	
	data "ibm_is_vpc_default_routing_table" "def_route_table" {
		vpc = ibm_is_vpc.test_vpc.id
	}
	`, vpcname)
}
