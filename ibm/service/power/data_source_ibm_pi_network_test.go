// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIBMPINetworkDataSource_basic(t *testing.T) {
	networkRes := "data.ibm_pi_network.testacc_ds_network"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPINetworkDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(networkRes, "id"),
				),
			},
		},
	})
}

func testAccCheckIBMPINetworkDataSourceConfig() string {
	return fmt.Sprintf(`
		data "ibm_pi_network" "testacc_ds_network" {
			pi_cloud_instance_id = "%[2]s"
			pi_network_id        = "%[1]s"
		}`, acc.Pi_network_id, acc.Pi_cloud_instance_id)
}
