// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIBMPIPISharedProcessorPoolDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIPISharedProcessorPoolDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_pi_shared_processor_pool.test_pool", "id"),
				),
			},
		},
	})
}

func testAccCheckIBMPIPISharedProcessorPoolDataSourceConfig() string {
	return fmt.Sprintf(`
		data "ibm_pi_shared_processor_pool" "test_pool" {
			pi_shared_processor_pool_id = "%s"
			pi_cloud_instance_id = "%s"
		}`, acc.Pi_shared_processor_pool_id, acc.Pi_cloud_instance_id)
}
