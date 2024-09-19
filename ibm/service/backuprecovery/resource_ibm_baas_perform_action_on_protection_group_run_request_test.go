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

func TestAccIbmBaasPerformActionOnProtectionGroupRunRequestBasic(t *testing.T) {
	var conf backuprecoveryv1.ProtectionGroupRunsResponse
	xIbmTenantID := fmt.Sprintf("tf_x_ibm_tenant_id_%d", acctest.RandIntRange(10, 100))
	action := "Pause"
	objectId := 72

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmBaasPerformActionOnProtectionGroupRunRequestDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBaasPerformActionOnProtectionGroupRunRequestConfigBasic(action, objectId),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmBaasPerformActionOnProtectionGroupRunRequestExists("ibm_baas_perform_action_on_protection_group_run_request.baas_perform_action_on_protection_group_run_request_instance", conf),
					resource.TestCheckResourceAttr("ibm_baas_perform_action_on_protection_group_run_request.baas_perform_action_on_protection_group_run_request_instance", "x_ibm_tenant_id", xIbmTenantID),
					resource.TestCheckResourceAttr("ibm_baas_perform_action_on_protection_group_run_request.baas_perform_action_on_protection_group_run_request_instance", "action", action),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_baas_perform_action_on_protection_group_run_request.baas_perform_action_on_protection_group_run_request",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmBaasPerformActionOnProtectionGroupRunRequestConfigBasic(action string, objectId int) string {
	return fmt.Sprintf(`

			resource "ibm_baas_protection_policy" "baas_protection_policy_instance" {
				x_ibm_tenant_id = "%s"
				name = "tf-name-policy-test-1"
				backup_policy {
						regular {
							incremental{
								schedule{
										day_schedule {
											frequency = 1
										}
										unit = "Days"
									}
							}
							retention {
								duration = 1
								unit = "Weeks"
							}
							primary_backup_target {
								use_default_backup_target = true
							}
						}
				}
				retry_options {
				retries = 3
				retry_interval_mins = 5
				}
			}

		resource "ibm_baas_protection_group" "baas_protection_group_instance" {
			x_ibm_tenant_id = "%s"
			policy_id = ibm_baas_protection_policy.baas_protection_policy_instance.id
			name = "tf-name-group-test-1"
			environment = "kPhysical"
			physical_params {
				protection_type = "kFile"
				file_protection_type_params {
				objects {
					id = %d
					file_paths{
						included_path = "/"
					}
				}
				}
			}
		}

		resource "ibm_baas_perform_action_on_protection_group_run_request" "baas_perform_action_on_protection_group_run_request_instance" {
			x_ibm_tenant_id = "%s"
			action = "%s"
		}
	`, tenantId, tenantId, objectId, tenantId, action)
}

func testAccCheckIbmBaasPerformActionOnProtectionGroupRunRequestExists(n string, obj backuprecoveryv1.ProtectionGroupRunsResponse) resource.TestCheckFunc {

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

		performActionOnProtectionGroupRunRequest, _, err := backupRecoveryClient.GetProtectionGroupRuns(getProtectionGroupRunsOptions)
		if err != nil {
			return err
		}

		obj = *performActionOnProtectionGroupRunRequest
		return nil
	}
}

func testAccCheckIbmBaasPerformActionOnProtectionGroupRunRequestDestroy(s *terraform.State) error {
	backupRecoveryClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).BackupRecoveryV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_baas_perform_action_on_protection_group_run_request" {
			continue
		}

		getProtectionGroupRunsOptions := &backuprecoveryv1.GetProtectionGroupRunsOptions{}

		getProtectionGroupRunsOptions.SetID(rs.Primary.ID)

		// Try to find the key
		_, response, err := backupRecoveryClient.GetProtectionGroupRuns(getProtectionGroupRunsOptions)

		if err == nil {
			return fmt.Errorf("baas_perform_action_on_protection_group_run_request still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for baas_perform_action_on_protection_group_run_request (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
