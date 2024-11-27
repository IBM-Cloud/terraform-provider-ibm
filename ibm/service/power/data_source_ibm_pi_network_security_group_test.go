// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMPINetworkSecurityGroupDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPINetworkSecurityGroupDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_pi_network_security_group.network_security_group", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_pi_network_security_group.network_security_group", "name"),
				),
			},
		},
	})
}

func testAccCheckIBMPINetworkSecurityGroupDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_pi_network_security_group" "network_security_group" {
			pi_cloud_instance_id = "%s"
			pi_network_security_group_id = "%s"
		}`, acc.Pi_cloud_instance_id, acc.Pi_network_security_group_id)
}
