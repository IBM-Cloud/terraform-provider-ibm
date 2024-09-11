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

func TestAccIbmBaasSearchProtectedObjectsDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBaasSearchProtectedObjectsDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_baas_search_protected_objects.baas_search_protected_objects_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_search_protected_objects.baas_search_protected_objects_instance", "x_ibm_tenant_id"),
				),
			},
		},
	})
}

func testAccCheckIbmBaasSearchProtectedObjectsDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_baas_search_protected_objects" "baas_search_protected_objects_instance" {
			X-IBM-Tenant-Id = "X-IBM-Tenant-Id"
			requestInitiatorType = "UIUser"
			searchString = "searchString"
			environments = [ "kPhysical" ]
			snapshotActions = [ "RecoverVMs" ]
			objectActionKey = "kPhysical"
			protectionGroupIds = [ "protectionGroupIds" ]
			objectIds = [ 1 ]
			subResultSize = 1
			filterSnapshotFromUsecs = 1
			filterSnapshotToUsecs = 1
			osTypes = [ "kLinux" ]
			sourceIds = [ 1 ]
			runInstanceIds = [ 1 ]
			cdpProtectedOnly = true
			useCachedData = true
		}
	`)
}
