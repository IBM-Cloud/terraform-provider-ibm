// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMPIVolumeGroupRemoteCopyRelationshipsDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIVolumeGroupRemoteCopyRelationshipsDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_pi_volume_group_remote_copy_relationships.testacc_volume_group_remote_copy_relationships", "id"),
				),
			},
		},
	})
}

func testAccCheckIBMPIVolumeGroupRemoteCopyRelationshipsDataSourceConfig() string {
	return fmt.Sprintf(`
		data "ibm_pi_volume_group_remote_copy_relationships" "testacc_volume_group_remote_copy_relationships" {
			pi_cloud_instance_id = "%s"
			pi_volume_group_id   = "%s"
		}`, acc.Pi_cloud_instance_id, acc.Pi_volume_group_id)
}
