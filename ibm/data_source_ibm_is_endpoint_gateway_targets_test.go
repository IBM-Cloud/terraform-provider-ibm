// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMISEndpointGatewayTargetsDataSource_basics(t *testing.T) {
	resName := "data.ibm_is_endpoint_gateway_targets.test"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISEgtsDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resName, "resources.0.name"),
					resource.TestCheckResourceAttrSet(resName, "resources.0.crn"),
				),
			},
		},
	})
}

func testAccCheckIBMISEgtsDataSourceConfig() string {

	return fmt.Sprintf(`

	data "ibm_is_endpoint_gateway_targets" "test" { 
	  }
	  `)
}
