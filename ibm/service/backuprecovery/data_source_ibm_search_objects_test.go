// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package backuprecovery_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmSearchObjectsDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSearchObjectsDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_search_objects.search_objects_instance", "id"),
				),
			},
		},
	})
}

func testAccCheckIbmSearchObjectsDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_search_objects" "search_objects_instance" {
			requestInitiatorType = "UIUser"
			searchString = "searchString"
			environments = [ "kPhysical" ]
			protectionTypes = [ "kAgent" ]
			tenantIds = [ "tenantIds" ]
			includeTenants = true
			protectionGroupIds = [ "protectionGroupIds" ]
			objectIds = [ 1 ]
			osTypes = [ "kLinux" ]
			sourceIds = [ 1 ]
			sourceUuids = [ "sourceUuids" ]
			isProtected = true
			isDeleted = true
			lastRunStatusList = [ "Accepted" ]
			regionIds = [ "regionIds" ]
			clusterIdentifiers = [ "clusterIdentifiers" ]
			storageDomainIds = [ "storageDomainIds" ]
			includeDeletedObjects = true
			paginationCookie = "paginationCookie"
			count = 1
			mustHaveTagIds = [ "mustHaveTagIds" ]
			mightHaveTagIds = [ "mightHaveTagIds" ]
			mustHaveSnapshotTagIds = [ "mustHaveSnapshotTagIds" ]
			mightHaveSnapshotTagIds = [ "mightHaveSnapshotTagIds" ]
		}
	`)
}
