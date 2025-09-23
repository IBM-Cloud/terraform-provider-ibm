// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMPINetworkAddressGroupDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPINetworkAddressGroupDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_pi_network_address_group.network_address_group", "id"),
				),
			},
		},
	})
}

func testAccCheckIBMPINetworkAddressGroupDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_pi_network_address_group" "network_address_group" {
			pi_cloud_instance_id = "%s"
			pi_network_address_group_id = "%s"
		}`, acc.Pi_cloud_instance_id, acc.Pi_network_address_group_id)
}
