// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMPIInstanceIPDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIInstanceIPDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_pi_instance_ip.testacc_ds_instance_ip", "id"),
				),
			},
		},
	})
}

func testAccCheckIBMPIInstanceIPDataSourceConfig() string {
	return fmt.Sprintf(`
		data "ibm_pi_instance_ip" "testacc_ds_instance_ip" {
			pi_cloud_instance_id = "%[1]s"
			pi_instance_id       = "%[2]s"
			pi_network_name      = "%[3]s"
		}`, acc.Pi_cloud_instance_id, acc.Pi_instance_id, acc.Pi_network_name)
}
