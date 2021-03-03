// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMISVirtualEndpointGatewaysDataSource_basic(t *testing.T) {
	t.Skip()
	resName := "data.ibm_is_virtual_endpoint_gateways.test1"
	vpcname1 := fmt.Sprintf("tfvpngw-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname1 := fmt.Sprintf("tfvpngw-subnet-%d", acctest.RandIntRange(10, 100))
	name1 := fmt.Sprintf("tfvpngw-createname-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVirtualEndpointGatewaysDataSourceConfig(vpcname1, subnetname1, name1),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(
						resName, "virtual_endpoint_gateways.0.name"),
				),
			},
		},
	})
}

func testAccCheckIBMISVirtualEndpointGatewaysDataSourceConfig(vpcname1, subnetname1, name1 string) string {
	// status filter defaults to empty
	return testAccCheckisVirtualEndpointGatewayConfigBasic(vpcname1, subnetname1, name1) + fmt.Sprintf(`
	data "ibm_is_virtual_endpoint_gateways" "test1" {
		
	}`)
}
