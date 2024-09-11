// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package backuprecovery_test

import (
	"fmt"
	"io"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.ibm.com/BackupAndRecovery/ibm-backup-recovery-sdk-go/backuprecoveryv1"
)

func TestAccIbmBaasDownloadAgentBasic(t *testing.T) {
	var conf io.ReadCloser
	xIbmTenantID := fmt.Sprintf("tf_x_ibm_tenant_id_%d", acctest.RandIntRange(10, 100))
	platform := "kWindows"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmBaasDownloadAgentDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBaasDownloadAgentConfigBasic(xIbmTenantID, platform),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmBaasDownloadAgentExists("ibm_baas_download_agent.baas_download_agent_instance", conf),
					resource.TestCheckResourceAttr("ibm_baas_download_agent.baas_download_agent_instance", "x_ibm_tenant_id", xIbmTenantID),
					resource.TestCheckResourceAttr("ibm_baas_download_agent.baas_download_agent_instance", "platform", platform),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_baas_download_agent.baas_download_agent",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmBaasDownloadAgentConfigBasic(xIbmTenantID string, platform string) string {
	return fmt.Sprintf(`
		resource "ibm_baas_download_agent" "baas_download_agent_instance" {
			x_ibm_tenant_id = "%s"
			platform = "%s"
		}
	`, xIbmTenantID, platform)
}

func testAccCheckIbmBaasDownloadAgentExists(n string, obj io.ReadCloser) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		_, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		backupRecoveryClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).BackupRecoveryV1()
		if err != nil {
			return err
		}

		downloadAgentOptions := &backuprecoveryv1.DownloadAgentOptions{}

		downloadAgentRequestParams, _, err := backupRecoveryClient.DownloadAgent(downloadAgentOptions)
		if err != nil {
			return err
		}

		obj = downloadAgentRequestParams
		return nil
	}
}

func testAccCheckIbmBaasDownloadAgentDestroy(s *terraform.State) error {
	backupRecoveryClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).BackupRecoveryV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_baas_download_agent" {
			continue
		}

		downloadAgentOptions := &backuprecoveryv1.DownloadAgentOptions{}

		// Try to find the key
		_, response, err := backupRecoveryClient.DownloadAgent(downloadAgentOptions)

		if err == nil {
			return fmt.Errorf("baas_download_agent still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for baas_download_agent (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
