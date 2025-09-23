// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package directlink_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMDLGatewayMacsecCakDataSource_basic(t *testing.T) {
	node := "data.ibm_dl_gateway_macsec_cak.test_dl_gateway_macsec_cak"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMDLGatewayMacsecCakDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(node, "id"),
					resource.TestCheckResourceAttrSet(node, "cak_id"),
					resource.TestCheckResourceAttrSet(node, "status"),
					resource.TestCheckResourceAttrSet(node, "created_at"),
				),
			},
		},
	})
}

func testAccCheckIBMDLGatewayMacsecCakDataSourceConfig() string {
	return `
	data "ibm_dl_gateway_macsec_cak" "test_dl_gateway_macsec_cak" {
    	gateway = "9c95f464-1ba9-471e-85b4-d2bf188cb273"
        cak_id = "1ec87f08-95f6-4a85-9518-b2094c5a2520"
	}
	`
}
