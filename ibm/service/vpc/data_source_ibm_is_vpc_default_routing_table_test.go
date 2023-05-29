// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMISVPCDefaultRoutingTableDataSource_basic(t *testing.T) {
	node := "data.ibm_is_vpc_default_routing_table.def_route_table"
	vpcname := fmt.Sprintf("tf-vpcname-%d", acctest.RandIntRange(100, 200))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVPCDefaultRoutingTableDataSourceConfig(vpcname),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(node, "id"),
					resource.TestCheckResourceAttrSet(node, "name"),
					resource.TestCheckResourceAttrSet(node, "lifecycle_state"),
					resource.TestCheckResourceAttrSet(node, "route_internet_ingress"),
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
