// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIBMPINetworksDataSource_basic(t *testing.T) {
	networksResData := "data.ibm_pi_networks.testacc_ds_networks"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPINetworksDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(networksResData, "id"),
				),
			},
		},
	})
}

func testAccCheckIBMPINetworksDataSourceConfig() string {
	return fmt.Sprintf(`
		data "ibm_pi_networks" "testacc_ds_networks" {
			pi_cloud_instance_id = "%s"
		}`, acc.Pi_cloud_instance_id)
}
