// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.94.0-fa797aec-20240814-142622
 */

package backuprecovery_test

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmBackupRecoveryDownloadAgentDataSourceBasic(t *testing.T) {
	filePath := "./temp/Cohesity_Agent_ibm_rm_20240824_Win_x64_Installer_test_datasource.exe"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBackupRecoveryDownloadAgentDataSourceConfigBasic(filePath),
				Check: resource.ComposeTestCheckFunc(
					testCheckFileNameExists(filePath),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_download_agent.baas_download_agent_instance", "id"),
					resource.TestCheckResourceAttr("data.ibm_backup_recovery_download_agent.baas_download_agent_instance", "x_ibm_tenant_id", tenantId),
					resource.TestCheckResourceAttr("data.ibm_backup_recovery_download_agent.baas_download_agent_instance", "platform", "kWindows"),
					resource.TestCheckResourceAttr("data.ibm_backup_recovery_download_agent.baas_download_agent_instance", "file_path", filePath),
				),
			},
		},
	})
}

func testCheckFileNameExists(path string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		timeout := time.Now().Add(2 * time.Minute)
		for time.Now().Before(timeout) {

			_, err := os.Stat(path)
			if err != nil {
				if os.IsNotExist(err) {
					time.Sleep(5 * time.Second)
					continue
				} else {
					return err
				}
			} else {
				return nil
			}
		}
		return nil
	}
}

func testAccCheckIbmBackupRecoveryDownloadAgentDataSourceConfigBasic(filePath string) string {
	return fmt.Sprintf(`
		data "ibm_backup_recovery_download_agent" "baas_download_agent_instance" {
			x_ibm_tenant_id = "%s"
			
			platform = "kWindows"
			file_path = "%s"
		}
	`, tenantId, filePath)
}
