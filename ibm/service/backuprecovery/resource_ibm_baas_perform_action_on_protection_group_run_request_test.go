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
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/backuprecovery"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/stretchr/testify/assert"
	"github.ibm.com/BackupAndRecovery/ibm-backup-recovery-sdk-go/backuprecoveryv1"
)

func TestAccIbmBaasPerformActionOnProtectionGroupRunRequestBasic(t *testing.T) {
	var conf backuprecoveryv1.ProtectionGroupRunsResponse
	xIbmTenantID := fmt.Sprintf("tf_x_ibm_tenant_id_%d", acctest.RandIntRange(10, 100))
	action := "Pause"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmBaasPerformActionOnProtectionGroupRunRequestDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBaasPerformActionOnProtectionGroupRunRequestConfigBasic(xIbmTenantID, action),
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

func testAccCheckIbmBaasPerformActionOnProtectionGroupRunRequestConfigBasic(xIbmTenantID string, action string) string {
	return fmt.Sprintf(`
		resource "ibm_baas_perform_action_on_protection_group_run_request" "baas_perform_action_on_protection_group_run_request_instance" {
			x_ibm_tenant_id = "%s"
			action = "%s"
		}
	`, xIbmTenantID, action)
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

func TestResourceIbmBaasPerformActionOnProtectionGroupRunRequestPauseProtectionRunActionParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["run_id"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.PauseProtectionRunActionResponseParams)
	model.RunID = core.StringPtr("testString")

	result, err := backuprecovery.ResourceIbmBaasPerformActionOnProtectionGroupRunRequestPauseProtectionRunActionParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasPerformActionOnProtectionGroupRunRequestResumeProtectionRunActionParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["run_id"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.ResumeProtectionRunActionResponseParams)
	model.RunID = core.StringPtr("testString")

	result, err := backuprecovery.ResourceIbmBaasPerformActionOnProtectionGroupRunRequestResumeProtectionRunActionParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasPerformActionOnProtectionGroupRunRequestCancelProtectionGroupRunRequestToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["run_id"] = "testString"
		model["local_task_id"] = "testString"
		model["object_ids"] = []int64{int64(26)}
		model["replication_task_id"] = []string{"testString"}
		model["archival_task_id"] = []string{"testString"}
		model["cloud_spin_task_id"] = []string{"testString"}

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.CancelProtectionGroupRunResponseParams)
	model.RunID = core.StringPtr("testString")

	result, err := backuprecovery.ResourceIbmBaasPerformActionOnProtectionGroupRunRequestCancelProtectionGroupRunRequestToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasPerformActionOnProtectionGroupRunRequestMapToPauseProtectionRunActionParams(t *testing.T) {
	checkResult := func(result *backuprecoveryv1.PauseProtectionRunActionParams) {
		model := new(backuprecoveryv1.PauseProtectionRunActionParams)
		model.RunID = core.StringPtr("testString")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["run_id"] = "testString"

	result, err := backuprecovery.ResourceIbmBaasPerformActionOnProtectionGroupRunRequestMapToPauseProtectionRunActionParams(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasPerformActionOnProtectionGroupRunRequestMapToResumeProtectionRunActionParams(t *testing.T) {
	checkResult := func(result *backuprecoveryv1.ResumeProtectionRunActionParams) {
		model := new(backuprecoveryv1.ResumeProtectionRunActionParams)
		model.RunID = core.StringPtr("testString")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["run_id"] = "testString"

	result, err := backuprecovery.ResourceIbmBaasPerformActionOnProtectionGroupRunRequestMapToResumeProtectionRunActionParams(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasPerformActionOnProtectionGroupRunRequestMapToCancelProtectionGroupRunRequest(t *testing.T) {
	checkResult := func(result *backuprecoveryv1.CancelProtectionGroupRunRequest) {
		model := new(backuprecoveryv1.CancelProtectionGroupRunRequest)
		model.RunID = core.StringPtr("testString")
		model.LocalTaskID = core.StringPtr("testString")
		model.ObjectIds = []int64{int64(26)}
		model.ReplicationTaskID = []string{"testString"}
		model.ArchivalTaskID = []string{"testString"}
		model.CloudSpinTaskID = []string{"testString"}

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["run_id"] = "testString"
	model["local_task_id"] = "testString"
	model["object_ids"] = []interface{}{int(26)}
	model["replication_task_id"] = []interface{}{"testString"}
	model["archival_task_id"] = []interface{}{"testString"}
	model["cloud_spin_task_id"] = []interface{}{"testString"}

	result, err := backuprecovery.ResourceIbmBaasPerformActionOnProtectionGroupRunRequestMapToCancelProtectionGroupRunRequest(model)
	assert.Nil(t, err)
	checkResult(result)
}
