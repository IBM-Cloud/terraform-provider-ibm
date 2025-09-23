// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power_test

import (
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMPIVolumeClone_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIVolumeCloneBasicConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_pi_volume_clone.testacc_ds_volume_clone", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_pi_volume_clone.testacc_ds_volume_clone", "status"),
				),
			},
		},
	})
}

func testAccCheckIBMPIVolumeCloneBasicConfig() string {
	return fmt.Sprintf(`
		data "ibm_pi_volume_clone" "testacc_ds_volume_clone" {
			pi_cloud_instance_id		= "%s"
			pi_volume_clone_task_id		= "%s"
		}`, acc.Pi_cloud_instance_id, acc.Pi_volume_clone_task_id)
}
