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

func TestAccIbmBackupRecoveryConnectorsMetadataDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBackupRecoveryConnectorsMetadataDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_connectors_metadata.baas_connectors_metadata_instance", "id"),
					resource.TestCheckResourceAttr("data.ibm_backup_recovery_connectors_metadata.baas_connectors_metadata_instance", "x_ibm_tenant_id", tenantId),
					resource.TestCheckResourceAttr("data.ibm_backup_recovery_connectors_metadata.baas_connectors_metadata_instance", "connector_image_metadata.#", "1"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_connectors_metadata.baas_connectors_metadata_instance", "connector_image_metadata.0.connector_image_file_list.#"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_connectors_metadata.baas_connectors_metadata_instance", "connector_image_metadata.0.connector_image_file_list.0.image_type"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_connectors_metadata.baas_connectors_metadata_instance", "connector_image_metadata.0.connector_image_file_list.0.url"),
				),
			},
		},
	})
}

func testAccCheckIbmBackupRecoveryConnectorsMetadataDataSourceConfigBasic() string {
	return fmt.Sprintf(`
	data "ibm_backup_recovery_connectors_metadata" "baas_connectors_metadata_instance" {
		x_ibm_tenant_id = "%s"
		backup_recovery_endpoint = "https://protectiondomain0103.us-east.backup-recovery-tests.cloud.ibm.com/v2"
	  }
	`, tenantId)
}
