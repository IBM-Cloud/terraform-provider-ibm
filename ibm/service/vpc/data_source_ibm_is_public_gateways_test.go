// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMISPublicGatewaysDatasource_basic(t *testing.T) {
	var publicgw string
	vpcname := fmt.Sprintf("tfpgw-vpc-%d", acctest.RandIntRange(10, 100))
	name1 := fmt.Sprintf("tfpgw-name-%d", acctest.RandIntRange(10, 100))
	zone := "us-south-1"

	resName := "data.ibm_is_public_gateways.test1"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISPublicGatewayDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISPublicGatewayDataSourceConfig(vpcname, name1, zone),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISPublicGatewayExists("ibm_is_public_gateway.testacc_public_gateway", publicgw),
					resource.TestCheckResourceAttr(
						"ibm_is_public_gateway.testacc_public_gateway", "name", name1),
					resource.TestCheckResourceAttr(
						"ibm_is_public_gateway.testacc_public_gateway", "zone", zone),
				),
			},
			{
				Config: testAccCheckIBMISPublicGatewaysDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resName, "public_gateways.0.id"),
					resource.TestCheckResourceAttrSet(resName, "public_gateways.0.name"),
					resource.TestCheckResourceAttrSet(resName, "public_gateways.0.zone"),
				),
			},
		},
	})
}

func testAccCheckIBMISPublicGatewaysDataSourceConfig() string {
	return fmt.Sprintf(`
	data "ibm_is_public_gateways" "test1"{
	}`)
}
