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

func TestAccIbmBaasUpdateProtectionGroupRunRequestBasic(t *testing.T) {
	var conf backuprecoveryv1.ProtectionGroupRunsResponse
	xIbmTenantID := fmt.Sprintf("tf_x_ibm_tenant_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmBaasUpdateProtectionGroupRunRequestDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBaasUpdateProtectionGroupRunRequestConfigBasic(xIbmTenantID),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmBaasUpdateProtectionGroupRunRequestExists("ibm_baas_update_protection_group_run_request.baas_update_protection_group_run_request_instance", conf),
					resource.TestCheckResourceAttr("ibm_baas_update_protection_group_run_request.baas_update_protection_group_run_request_instance", "x_ibm_tenant_id", xIbmTenantID),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_baas_update_protection_group_run_request.baas_update_protection_group_run_request",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmBaasUpdateProtectionGroupRunRequestConfigBasic(xIbmTenantID string) string {
	return fmt.Sprintf(`
		resource "ibm_baas_update_protection_group_run_request" "baas_update_protection_group_run_request_instance" {
			x_ibm_tenant_id = "%s"
			update_protection_group_run_params {
				run_id = "run_id"
				local_snapshot_config {
					enable_legal_hold = true
					delete_snapshot = true
					data_lock = "Compliance"
					days_to_keep = 1
				}
				replication_snapshot_config {
					new_snapshot_config {
						id = 1
						retention {
							unit = "Days"
							duration = 1
							data_lock_config {
								mode = "Compliance"
								unit = "Days"
								duration = 1
								enable_worm_on_external_target = true
							}
						}
					}
					update_existing_snapshot_config {
						id = 1
						name = "name"
						enable_legal_hold = true
						delete_snapshot = true
						resync = true
						data_lock = "Compliance"
						days_to_keep = 1
					}
				}
				archival_snapshot_config {
					new_snapshot_config {
						id = 1
						archival_target_type = "Tape"
						retention {
							unit = "Days"
							duration = 1
							data_lock_config {
								mode = "Compliance"
								unit = "Days"
								duration = 1
								enable_worm_on_external_target = true
							}
						}
						copy_only_fully_successful = true
					}
					update_existing_snapshot_config {
						id = 1
						name = "name"
						archival_target_type = "Tape"
						enable_legal_hold = true
						delete_snapshot = true
						resync = true
						data_lock = "Compliance"
						days_to_keep = 1
					}
				}
			}
		}
	`, xIbmTenantID)
}

func testAccCheckIbmBaasUpdateProtectionGroupRunRequestExists(n string, obj backuprecoveryv1.ProtectionGroupRunsResponse) resource.TestCheckFunc {

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

		updateProtectionGroupRunRequest, _, err := backupRecoveryClient.GetProtectionGroupRuns(getProtectionGroupRunsOptions)
		if err != nil {
			return err
		}

		obj = *updateProtectionGroupRunRequest
		return nil
	}
}

func testAccCheckIbmBaasUpdateProtectionGroupRunRequestDestroy(s *terraform.State) error {
	backupRecoveryClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).BackupRecoveryV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_baas_update_protection_group_run_request" {
			continue
		}

		getProtectionGroupRunsOptions := &backuprecoveryv1.GetProtectionGroupRunsOptions{}

		getProtectionGroupRunsOptions.SetID(rs.Primary.ID)

		// Try to find the key
		_, response, err := backupRecoveryClient.GetProtectionGroupRuns(getProtectionGroupRunsOptions)

		if err == nil {
			return fmt.Errorf("Update Protection Group Run Request Body. still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for Update Protection Group Run Request Body. (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
