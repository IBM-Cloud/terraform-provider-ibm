// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package satellite_test

import (
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIbmSatelliteStorageAssignmentDataSourceBasic(t *testing.T) {
	uuid := fmt.Sprintf("tf-uuid-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmSatelliteStorageAssignmentDataSourceConfigBasic(uuid),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_satellite_storage_assignment.satellite_storage_assignment", "uuid"),
					resource.TestCheckResourceAttrSet("data.ibm_satellite_storage_assignment.satellite_storage_assignment", "assignment_name"),
					resource.TestCheckResourceAttrSet("data.ibm_satellite_storage_assignment.satellite_storage_assignment", "owner"),
					resource.TestCheckResourceAttrSet("data.ibm_satellite_storage_assignment.satellite_storage_assignment", "groups"),
					resource.TestCheckResourceAttrSet("data.ibm_satellite_storage_assignment.satellite_storage_assignment", "cluster"),
					resource.TestCheckResourceAttrSet("data.ibm_satellite_storage_assignment.satellite_storage_assignment", "svc_cluster"),
					resource.TestCheckResourceAttrSet("data.ibm_satellite_storage_assignment.satellite_storage_assignment", "sat_cluster"),
					resource.TestCheckResourceAttrSet("data.ibm_satellite_storage_assignment.satellite_storage_assignment", "config"),
					resource.TestCheckResourceAttrSet("data.ibm_satellite_storage_assignment.satellite_storage_assignment", "config_uuid"),
					resource.TestCheckResourceAttrSet("data.ibm_satellite_storage_assignment.satellite_storage_assignment", "config_version"),
					resource.TestCheckResourceAttrSet("data.ibm_satellite_storage_assignment.satellite_storage_assignment", "config_version_uuid"),
					resource.TestCheckResourceAttrSet("data.ibm_satellite_storage_assignment.satellite_storage_assignment", "assignment_type"),
					resource.TestCheckResourceAttrSet("data.ibm_satellite_storage_assignment.satellite_storage_assignment", "created"),
					resource.TestCheckResourceAttrSet("data.ibm_satellite_storage_assignment.satellite_storage_assignment", "rollout_success_count"),
					resource.TestCheckResourceAttrSet("data.ibm_satellite_storage_assignment.satellite_storage_assignment", "rollout_error_count"),
					resource.TestCheckResourceAttrSet("data.ibm_satellite_storage_assignment.satellite_storage_assignment", "is_assignment_upgrade_available"),
				),
			},
		},
	})
}

func testAccCheckIbmSatelliteStorageAssignmentDataSourceConfigBasic(uuid string) string {
	return fmt.Sprintf(`
		data "ibm_satellite_storage_assignment" "satellite_storage_assignment" {
			uuid = "%s"
		}
	`, uuid)
}
