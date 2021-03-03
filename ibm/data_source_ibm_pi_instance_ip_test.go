// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMPIInstanceIPDataSource_basic(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
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
    pi_network_name = "%[1]s"
    pi_instance_name = "%[2]s"
    pi_cloud_instance_id = "%[3]s"
}`, pi_network_name, pi_instance_name, pi_cloud_instance_id)

}
