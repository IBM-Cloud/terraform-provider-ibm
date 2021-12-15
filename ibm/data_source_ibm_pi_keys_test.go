// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMPIKeysDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIKeysDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_pi_keys.test", "id"),
				),
			},
		},
	})
}

func testAccCheckIBMPIKeysDataSourceConfig() string {
	return fmt.Sprintf(`
		data "ibm_pi_keys" "test" {
			pi_cloud_instance_id = "%s"
		}
	`, pi_cloud_instance_id)
}
