// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power_test

import (
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMPIVolumeRemoteCopyRelationshipDataSource_basics(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIVolumeRemoteCopyRelationshipsDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_pi_volume_remote_copy_relationship.testacc_volume_remote_copy_relationship", "id"),
				),
			},
		},
	})
}

func testAccCheckIBMPIVolumeRemoteCopyRelationshipsDataSourceConfig() string {
	return fmt.Sprintf(`
		data "ibm_pi_volume_remote_copy_relationship" "testacc_volume_remote_copy_relationship" {
			pi_cloud_instance_id = "%s"
			pi_volume_id         = "%s"
		}`, acc.Pi_cloud_instance_id, acc.Pi_volume_id)
}
