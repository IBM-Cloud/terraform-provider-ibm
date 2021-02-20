/* IBM Confidential
*  Object Code Only Source Materials
*  5747-SM3
*  (c) Copyright IBM Corp. 2017,2021
*
*  The source code for this program is not published or otherwise divested
*  of its trade secrets, irrespective of what has been deposited with the
*  U.S. Copyright Office. */

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMISVPCRoutingTablesDataSource_basic(t *testing.T) {
	node := "data.ibm_is_vpc_routing_tables.test1"
	vpcname := fmt.Sprintf("tf-vpcname-%d", acctest.RandIntRange(100, 200))
	routetablename := fmt.Sprintf("tf-routetable-%d", acctest.RandIntRange(100, 200))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVPCRoutingTablesDataSourceConfig(vpcname, routetablename),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(node, "routing_tables.#"),
				),
			},
		},
	})
}

func testAccCheckIBMISVPCRoutingTablesDataSourceConfig(vpcname, routetablename string) string {
	return fmt.Sprintf(`

	resource "ibm_is_vpc" "test_custom_route_vpc" {
  		name = "%s"
	}
    
	resource "ibm_is_vpc_routing_table" "test_route_table" {
  		name = "%s"
  		vpc =  ibm_is_vpc.test_custom_route_vpc.id
	}

	data "ibm_is_vpc_routing_tables" "test1" {
		vpc =  ibm_is_vpc.test_custom_route_vpc.id
	}

	`, vpcname, routetablename)
}
