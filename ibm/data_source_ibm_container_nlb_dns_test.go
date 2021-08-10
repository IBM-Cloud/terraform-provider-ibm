// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMContainerNLBDNSDatasourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerNLBDNSDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_container_nlb_dns.dns", "id"),
				),
			},
		},
	})
}

func testAccCheckIBMContainerNLBDNSDataSourceConfig() string {
	return `
	data "ibm_container_nlb_dns" "dns" {
	    name = "c48mibfw00d5bj2919pg"
	}
`
}
