// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package backuprecovery_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmSearchProtectedObjectsDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSearchProtectedObjectsDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_search_protected_objects.search_protected_objects_instance", "id"),
				),
			},
		},
	})
}

func testAccCheckIbmSearchProtectedObjectsDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_search_protected_objects" "search_protected_objects_instance" {
			requestInitiatorType = "UIUser"
			searchString = "searchString"
			environments = [ "kPhysical" ]
			snapshotActions = [ "RecoverVMs" ]
			objectActionKey = "kPhysical"
			tenantIds = [ "tenantIds" ]
			includeTenants = true
			protectionGroupIds = [ "protectionGroupIds" ]
			objectIds = [ 1 ]
			storageDomainIds = [ 1 ]
			subResultSize = 1
			filterSnapshotFromUsecs = 1
			filterSnapshotToUsecs = 1
			osTypes = [ "kLinux" ]
			sourceIds = [ 1 ]
			runInstanceIds = [ 1 ]
			cdpProtectedOnly = true
			regionIds = [ "regionIds" ]
			useCachedData = true
		}
	`)
}
