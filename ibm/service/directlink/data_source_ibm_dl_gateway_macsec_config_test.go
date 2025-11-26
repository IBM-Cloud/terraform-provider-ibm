// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package directlink_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMDLGatewayMacsecConfigDataSource_basic(t *testing.T) {
	node := "data.ibm_dl_gateway_macsec_config.test_dl_gateway_macsec_config"
	gatewayID := "9c95f464-1ba9-471e-85b4-d2bf188cb273"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMDLGatewayMacsecConfigDataSourceConfig(gatewayID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(node, "id"),
					resource.TestCheckResourceAttrSet(node, "cipher_suite"),
					resource.TestCheckResourceAttrSet(node, "created_at"),
					resource.TestCheckResourceAttrSet(node, "status"),
				),
			},
		},
	})
}

func testAccCheckIBMDLGatewayMacsecConfigDataSourceConfig(gatewayName string) string {
	return fmt.Sprintf(`
	data "ibm_dl_gateway_macsec_config" "test_dl_gateway_macsec_config" {
		gateway = %q
	}
	`, gatewayName)
}
