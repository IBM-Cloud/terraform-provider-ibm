// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMPIPublicNetworkDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIPublicNetworkDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_pi_public_network.testacc_ds_public_network", "id"),
				),
			},
		},
	})
}

func testAccCheckIBMPIPublicNetworkDataSourceConfig() string {
	return fmt.Sprintf(`
data "ibm_pi_public_network" "testacc_ds_public_network" {
    pi_cloud_instance_id = "%s"
}`, acc.Pi_cloud_instance_id)

}
