// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package backuprecovery_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmBackupRecoverySearchIndexedObjectDataSourceBasic(t *testing.T) {
	objectType := "Files"
	objectId := 344
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBackupRecoverySearchIndexedObjectConfigBasic(objectType, objectId),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_backup_recovery_search_indexed_object.baas_search_indexed_object_instance", "x_ibm_tenant_id", tenantId),
					resource.TestCheckResourceAttr("data.ibm_backup_recovery_search_indexed_object.baas_search_indexed_object_instance", "object_type", objectType),
					resource.TestCheckResourceAttr("data.ibm_backup_recovery_search_indexed_object.baas_search_indexed_object_instance", "file_params.0.source_ids.0", strconv.Itoa(objectId)),
				),
			},
		},
	})
}

func testAccCheckIbmBackupRecoverySearchIndexedObjectConfigBasic(objectType string, objectId int) string {
	return fmt.Sprintf(`

		data "ibm_backup_recovery_search_indexed_object" "baas_search_indexed_object_instance" {
			x_ibm_tenant_id = "%s"
			
			object_type = "%s"
			file_params {
			  source_ids = [%d]
			}
		}
	`, tenantId, objectType, objectId)
}
