// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIBMPIDhcpDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIDhcpDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_pi_dhcp.dhcp", "id"),
				),
			},
		},
	})
}

func testAccCheckIBMPIDhcpDataSourceConfig() string {
	return fmt.Sprintf(`
		data "ibm_pi_dhcp" "dhcp" {
			pi_cloud_instance_id = "%s"
			pi_dhcp_id = "%s"
		}`, acc.Pi_cloud_instance_id, acc.Pi_dhcp_id)
}
