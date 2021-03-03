// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMPIInstanceDataSource_basic(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIInstanceDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_pi_instance.testacc_ds_instance", "id"),
				),
			},
		},
	})
}

func testAccCheckIBMPIInstanceDataSourceConfig() string {
	return fmt.Sprintf(`
	
data "ibm_pi_instance" "testacc_ds_instance" {
	pi_instance_name="%s"
    pi_cloud_instance_id = "%s"
}`, pi_instance_name, pi_cloud_instance_id)

}
