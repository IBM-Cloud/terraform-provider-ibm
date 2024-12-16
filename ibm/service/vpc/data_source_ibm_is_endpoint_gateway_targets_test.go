// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIBMISEndpointGatewayTargetsDataSource_basics(t *testing.T) {
	resName := "data.ibm_is_endpoint_gateway_targets.test"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
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
