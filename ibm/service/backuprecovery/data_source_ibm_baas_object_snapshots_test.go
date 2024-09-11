// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.94.0-fa797aec-20240814-142622
 */

package backuprecovery_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmBaasObjectSnapshotsDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBaasObjectSnapshotsDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_baas_object_snapshots.baas_object_snapshots_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_object_snapshots.baas_object_snapshots_instance", "baas_object_snapshots_id"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_object_snapshots.baas_object_snapshots_instance", "x_ibm_tenant_id"),
				),
			},
		},
	})
}

func testAccCheckIbmBaasObjectSnapshotsDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_baas_object_snapshots" "baas_object_snapshots_instance" {
			id = 1
			X-IBM-Tenant-Id = "X-IBM-Tenant-Id"
			fromTimeUsecs = 1
			toTimeUsecs = 1
			runStartFromTimeUsecs = 1
			runStartToTimeUsecs = 1
			snapshotActions = [ "RecoverVMs" ]
			runTypes = [ "kRegular" ]
			protectionGroupIds = [ "protectionGroupIds" ]
			runInstanceIds = [ 1 ]
			regionIds = [ "regionIds" ]
			objectActionKeys = [ "kVMware" ]
		}
	`)
}
