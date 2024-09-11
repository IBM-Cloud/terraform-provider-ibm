// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package backuprecovery_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.ibm.com/BackupAndRecovery/ibm-backup-recovery-sdk-go/backuprecoveryv1"
)

func TestAccIbmBaasProtectionGroupRunRequestBasic(t *testing.T) {
	var conf backuprecoveryv1.ProtectionGroupRunsResponse
	xIbmTenantID := fmt.Sprintf("tf_x_ibm_tenant_id_%d", acctest.RandIntRange(10, 100))
	groupID := fmt.Sprintf("group_id_%d", acctest.RandIntRange(10, 100))
	runType := "kRegular"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmBaasProtectionGroupRunRequestDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBaasProtectionGroupRunRequestConfigBasic(xIbmTenantID, runType, groupID),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmBaasProtectionGroupRunRequestExists("ibm_baas_protection_group_run_request.baas_protection_group_run_request_instance", conf),
					resource.TestCheckResourceAttr("ibm_baas_protection_group_run_request.baas_protection_group_run_request_instance", "x_ibm_tenant_id", xIbmTenantID),
					resource.TestCheckResourceAttr("ibm_baas_protection_group_run_request.baas_protection_group_run_request_instance", "x_ibm_tenant_id", groupID),
					resource.TestCheckResourceAttr("ibm_baas_protection_group_run_request.baas_protection_group_run_request_instance", "group_id", runType),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_baas_protection_group_run_request.baas_protection_group_run_request",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmBaasProtectionGroupRunRequestConfigBasic(xIbmTenantID string, runType, groupID string) string {
	return fmt.Sprintf(`
		resource "ibm_baas_protection_group_run_request" "baas_protection_group_run_request_instance" {
			x_ibm_tenant_id = "%s"
			run_type = "%s"
			group_id = "%s"
		}
	`, xIbmTenantID, runType, groupID)
}

func testAccCheckIbmBaasProtectionGroupRunRequestExists(n string, obj backuprecoveryv1.ProtectionGroupRunsResponse) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		backupRecoveryClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).BackupRecoveryV1()
		if err != nil {
			return err
		}

		getProtectionGroupRunsOptions := &backuprecoveryv1.GetProtectionGroupRunsOptions{}

		getProtectionGroupRunsOptions.SetID(rs.Primary.ID)

		protectionGroupRunRequest, _, err := backupRecoveryClient.GetProtectionGroupRuns(getProtectionGroupRunsOptions)
		if err != nil {
			return err
		}

		obj = *protectionGroupRunRequest
		return nil
	}
}

func testAccCheckIbmBaasProtectionGroupRunRequestDestroy(s *terraform.State) error {
	backupRecoveryClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).BackupRecoveryV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_baas_protection_group_run_request" {
			continue
		}

		getProtectionGroupRunsOptions := &backuprecoveryv1.GetProtectionGroupRunsOptions{}

		getProtectionGroupRunsOptions.SetID(rs.Primary.ID)

		// Try to find the key
		_, response, err := backupRecoveryClient.GetProtectionGroupRuns(getProtectionGroupRunsOptions)

		if err == nil {
			return fmt.Errorf("baas_protection_group_run_request still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for baas_protection_group_run_request (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
