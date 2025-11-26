// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIBMPIVolumeDataSource_basic(t *testing.T) {
	volumeRes := "data.ibm_pi_volume.testacc_ds_volume"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIVolumeDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(volumeRes, "id"),
				),
			},
		},
	})
}

func testAccCheckIBMPIVolumeDataSourceConfig() string {
	return fmt.Sprintf(`
		data "ibm_pi_volume" "testacc_ds_volume" {
			pi_cloud_instance_id = "%[1]s"
			pi_volume_id         = "%[2]s"
		}`, acc.Pi_cloud_instance_id, acc.Pi_volume_id)
}

func TestAccIBMPIVolumeDataSource_replication(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIVolumeDataSourceReplicationConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_pi_volume.testacc_ds_volume", "id"),
					resource.TestCheckResourceAttr("data.ibm_pi_volume.testacc_ds_volume", "replication_enabled", "true"),
					resource.TestCheckResourceAttrSet("data.ibm_pi_volume.testacc_ds_volume", "replication_status"),
				),
			},
		},
	})
}

func testAccCheckIBMPIVolumeDataSourceReplicationConfig() string {
	return fmt.Sprintf(`
		data "ibm_pi_volume" "testacc_ds_volume" {
			pi_cloud_instance_id = "%[1]s"
			pi_volume_id         = "%[2]s"
		}`, acc.Pi_cloud_instance_id, acc.Pi_replication_volume_id)
}
