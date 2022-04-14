// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMPIStoragePoolCapacityDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIStoragePoolCapacityDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_pi_storage_pool_capacity.pool", "id"),
				),
			},
		},
	})
}

func testAccCheckIBMPIStoragePoolCapacityDataSourceConfig() string {
	return fmt.Sprintf(`
		data "ibm_pi_storage_pool_capacity" "pool" {
			pi_cloud_instance_id = "%s"
			pi_storage_pool = "%s"
		}
	`, acc.Pi_cloud_instance_id, acc.PiStoragePool)
}
