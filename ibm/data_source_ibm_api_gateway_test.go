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

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMAPIGatewayDataSource_basic(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMAPIGatewayDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(
						"data.ibm_api_gateway.apigateway", "service_instance_crn"),
				),
			},
		},
	})
}

func testAccCheckIBMAPIGatewayDataSourceConfig() string {
	return fmt.Sprintf(`
		resource "ibm_resource_instance" "apigateway"{
			name     = "testname"
			location = "global"
			service  = "api-gateway"
			plan     = "lite"
		}
		data "ibm_api_gateway" "apigateway" {
			service_instance_crn =ibm_resource_instance.apigateway.id
		}`)

}
