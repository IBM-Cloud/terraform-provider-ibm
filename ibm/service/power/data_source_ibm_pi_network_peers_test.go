// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMPINetworkPeersDataSourceBasic(t *testing.T) {
	networksResData := "data.ibm_pi_network_peers.network_peers"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPINetworkPeersDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(networksResData, "id"),
					resource.TestCheckResourceAttrSet(networksResData, "network_peers.#"),
				),
			},
		},
	})
}

func testAccCheckIBMPINetworkPeersDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_pi_network_peers" "network_peers" {
			pi_cloud_instance_id = "%s"
		}`, acc.Pi_cloud_instance_id)
}
