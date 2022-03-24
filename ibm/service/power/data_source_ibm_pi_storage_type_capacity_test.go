// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMPIStorageTypeCapacityDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIStorageTypeCapacityDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_pi_storage_type_capacity.type", "id"),
				),
			},
		},
	})
}

func testAccCheckIBMPIStorageTypeCapacityDataSourceConfig() string {
	return fmt.Sprintf(`
		data "ibm_pi_storage_type_capacity" "type" {
			pi_cloud_instance_id = "%s"
			pi_storage_type = "%s"
		}
	`, acc.Pi_cloud_instance_id, acc.PiStorageType)
}
