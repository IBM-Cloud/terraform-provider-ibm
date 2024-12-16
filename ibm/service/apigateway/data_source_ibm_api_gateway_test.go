// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package apigateway_test

import (
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIBMAPIGatewayDataSource_basic(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
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
	return `
		resource "ibm_resource_instance" "apigateway"{
			name     = "testname"
			location = "global"
			service  = "api-gateway"
			plan     = "lite"
		}
		data "ibm_api_gateway" "apigateway" {
			service_instance_crn =ibm_resource_instance.apigateway.id
		}`

}
