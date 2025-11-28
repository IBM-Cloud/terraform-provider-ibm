// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMPIInstanceNetworkDataSource_basic(t *testing.T) {
	ds := "data.ibm_pi_instance_network.testacc_ds_instance_network"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIInstanceNetworkDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(ds, "id"),
				),
			},
		},
	})
}

func testAccCheckIBMPIInstanceNetworkDataSourceConfig() string {
	return fmt.Sprintf(`
		data "ibm_pi_instance_network" "testacc_ds_instance_network" {
  			pi_cloud_instance_id = "%[1]s"
  			pi_instance_id       = "%[2]s"
  			pi_network_id        = "%[3]s"
		}`, acc.Pi_cloud_instance_id, acc.Pi_instance_id, acc.Pi_network_id)
}
