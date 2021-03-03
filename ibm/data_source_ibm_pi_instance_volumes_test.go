// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMPIVolumesDataSource_basic(t *testing.T) {
	name := fmt.Sprintf("tf-pi-volume-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIVolumesDataSourceConfig(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_pi_instance_volumes.testacc_ds_volumes", "id"),
				),
			},
		},
	})
}

func testAccCheckIBMPIVolumesDataSourceConfig(name string) string {
	return fmt.Sprintf(`
data "ibm_pi_instance_volumes" "testacc_ds_volumes" {
    pi_instance_name = "%s"
    pi_cloud_instance_id = "%s"
}`, pi_instance_name, pi_cloud_instance_id)

}
