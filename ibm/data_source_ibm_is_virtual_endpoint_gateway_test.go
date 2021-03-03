// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMISVirtualEndpointGatewayDataSource_basic(t *testing.T) {
	resName := "data.ibm_is_virtual_endpoint_gateway.data_test"
	vpcname1 := fmt.Sprintf("tfvpngw-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname1 := fmt.Sprintf("tfvpngw-subnet-%d", acctest.RandIntRange(10, 100))
	name1 := fmt.Sprintf("tfvpngw-createname-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckisVirtualEndpointGatewayDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVirtualEndpointGatewayDataSourceConfig(vpcname1, subnetname1, name1),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						resName, "name", name1),
				),
			},
		},
	})
}

func testAccCheckIBMISVirtualEndpointGatewayDataSourceConfig(vpcname1, subnetname1, name1 string) string {
	// status filter defaults to empty
	return testAccCheckisVirtualEndpointGatewayConfigBasic(vpcname1, subnetname1, name1) + fmt.Sprintf(`
	data "ibm_is_virtual_endpoint_gateway" "data_test" {
        name = ibm_is_virtual_endpoint_gateway.endpoint_gateway.name
	}`)
}
