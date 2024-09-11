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

func TestAccIbmBaasAgentUpgradeTasksDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBaasAgentUpgradeTasksDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_baas_agent_upgrade_tasks.baas_agent_upgrade_tasks_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_agent_upgrade_tasks.baas_agent_upgrade_tasks_instance", "x_ibm_tenant_id"),
				),
			},
		},
	})
}

func testAccCheckIbmBaasAgentUpgradeTasksDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_baas_agent_upgrade_tasks" "baas_agent_upgrade_tasks_instance" {
			X-IBM-Tenant-Id = "X-IBM-Tenant-Id"
			ids = [ 1 ]
		}
	`)
}

// func TestDataSourceIbmBaasAgentUpgradeTasksAgentUpgradeTaskStateToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		errorModel := make(map[string]interface{})
// 		errorModel["error_code"] = "testString"
// 		errorModel["message"] = "testString"
// 		errorModel["task_log_id"] = "testString"

// 		agentInfoObjectModel := make(map[string]interface{})
// 		agentInfoObjectModel["end_time_usecs"] = int(26)
// 		agentInfoObjectModel["error"] = []map[string]interface{}{errorModel}
// 		agentInfoObjectModel["name"] = "testString"
// 		agentInfoObjectModel["previous_software_version"] = "testString"
// 		agentInfoObjectModel["start_time_usecs"] = int(26)
// 		agentInfoObjectModel["status"] = "Scheduled"

// 		agentUpgradeInfoObjectModel := make(map[string]interface{})
// 		agentUpgradeInfoObjectModel["id"] = int(26)
// 		agentUpgradeInfoObjectModel["info"] = []map[string]interface{}{agentInfoObjectModel}

// 		model := make(map[string]interface{})
// 		model["agent_i_ds"] = []int64{int64(26)}
// 		model["agents"] = []map[string]interface{}{agentUpgradeInfoObjectModel}
// 		model["cluster_version"] = "testString"
// 		model["description"] = "testString"
// 		model["end_time_usecs"] = int(26)
// 		model["error"] = []map[string]interface{}{errorModel}
// 		model["id"] = int(26)
// 		model["is_retryable"] = true
// 		model["name"] = "testString"
// 		model["retried_task_id"] = int(26)
// 		model["schedule_end_time_usecs"] = int(26)
// 		model["schedule_time_usecs"] = int(26)
// 		model["start_time_usecs"] = int(26)
// 		model["status"] = "Scheduled"
// 		model["type"] = "Auto"

// 		assert.Equal(t, result, model)
// 	}

// 	errorModel := new(backuprecoveryv1.Error)
// 	errorModel.ErrorCode = core.StringPtr("testString")
// 	errorModel.Message = core.StringPtr("testString")
// 	errorModel.TaskLogID = core.StringPtr("testString")

// 	agentInfoObjectModel := new(backuprecoveryv1.AgentInfoObject)
// 	agentInfoObjectModel.EndTimeUsecs = core.Int64Ptr(int64(26))
// 	agentInfoObjectModel.Error = errorModel
// 	agentInfoObjectModel.Name = core.StringPtr("testString")
// 	agentInfoObjectModel.PreviousSoftwareVersion = core.StringPtr("testString")
// 	agentInfoObjectModel.StartTimeUsecs = core.Int64Ptr(int64(26))
// 	agentInfoObjectModel.Status = core.StringPtr("Scheduled")

// 	agentUpgradeInfoObjectModel := new(backuprecoveryv1.AgentUpgradeInfoObject)
// 	agentUpgradeInfoObjectModel.ID = core.Int64Ptr(int64(26))
// 	agentUpgradeInfoObjectModel.Info = agentInfoObjectModel

// 	model := new(backuprecoveryv1.AgentUpgradeTaskState)
// 	model.AgentIDs = []int64{int64(26)}
// 	model.Agents = []backuprecoveryv1.AgentUpgradeInfoObject{*agentUpgradeInfoObjectModel}
// 	model.ClusterVersion = core.StringPtr("testString")
// 	model.Description = core.StringPtr("testString")
// 	model.EndTimeUsecs = core.Int64Ptr(int64(26))
// 	model.Error = errorModel
// 	model.ID = core.Int64Ptr(int64(26))
// 	model.IsRetryable = core.BoolPtr(true)
// 	model.Name = core.StringPtr("testString")
// 	model.RetriedTaskID = core.Int64Ptr(int64(26))
// 	model.ScheduleEndTimeUsecs = core.Int64Ptr(int64(26))
// 	model.ScheduleTimeUsecs = core.Int64Ptr(int64(26))
// 	model.StartTimeUsecs = core.Int64Ptr(int64(26))
// 	model.Status = core.StringPtr("Scheduled")
// 	model.Type = core.StringPtr("Auto")

// 	result, err := backuprecovery.DataSourceIbmBaasAgentUpgradeTasksAgentUpgradeTaskStateToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmBaasAgentUpgradeTasksAgentUpgradeInfoObjectToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		errorModel := make(map[string]interface{})
// 		errorModel["error_code"] = "testString"
// 		errorModel["message"] = "testString"
// 		errorModel["task_log_id"] = "testString"

// 		agentInfoObjectModel := make(map[string]interface{})
// 		agentInfoObjectModel["end_time_usecs"] = int(26)
// 		agentInfoObjectModel["error"] = []map[string]interface{}{errorModel}
// 		agentInfoObjectModel["name"] = "testString"
// 		agentInfoObjectModel["previous_software_version"] = "testString"
// 		agentInfoObjectModel["start_time_usecs"] = int(26)
// 		agentInfoObjectModel["status"] = "Scheduled"

// 		model := make(map[string]interface{})
// 		model["id"] = int(26)
// 		model["info"] = []map[string]interface{}{agentInfoObjectModel}

// 		assert.Equal(t, result, model)
// 	}

// 	errorModel := new(backuprecoveryv1.Error)
// 	errorModel.ErrorCode = core.StringPtr("testString")
// 	errorModel.Message = core.StringPtr("testString")
// 	errorModel.TaskLogID = core.StringPtr("testString")

// 	agentInfoObjectModel := new(backuprecoveryv1.AgentInfoObject)
// 	agentInfoObjectModel.EndTimeUsecs = core.Int64Ptr(int64(26))
// 	agentInfoObjectModel.Error = errorModel
// 	agentInfoObjectModel.Name = core.StringPtr("testString")
// 	agentInfoObjectModel.PreviousSoftwareVersion = core.StringPtr("testString")
// 	agentInfoObjectModel.StartTimeUsecs = core.Int64Ptr(int64(26))
// 	agentInfoObjectModel.Status = core.StringPtr("Scheduled")

// 	model := new(backuprecoveryv1.AgentUpgradeInfoObject)
// 	model.ID = core.Int64Ptr(int64(26))
// 	model.Info = agentInfoObjectModel

// 	result, err := backuprecovery.DataSourceIbmBaasAgentUpgradeTasksAgentUpgradeInfoObjectToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmBaasAgentUpgradeTasksAgentInfoObjectToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		errorModel := make(map[string]interface{})
// 		errorModel["error_code"] = "testString"
// 		errorModel["message"] = "testString"
// 		errorModel["task_log_id"] = "testString"

// 		model := make(map[string]interface{})
// 		model["end_time_usecs"] = int(26)
// 		model["error"] = []map[string]interface{}{errorModel}
// 		model["name"] = "testString"
// 		model["previous_software_version"] = "testString"
// 		model["start_time_usecs"] = int(26)
// 		model["status"] = "Scheduled"

// 		assert.Equal(t, result, model)
// 	}

// 	errorModel := new(backuprecoveryv1.Error)
// 	errorModel.ErrorCode = core.StringPtr("testString")
// 	errorModel.Message = core.StringPtr("testString")
// 	errorModel.TaskLogID = core.StringPtr("testString")

// 	model := new(backuprecoveryv1.AgentInfoObject)
// 	model.EndTimeUsecs = core.Int64Ptr(int64(26))
// 	model.Error = errorModel
// 	model.Name = core.StringPtr("testString")
// 	model.PreviousSoftwareVersion = core.StringPtr("testString")
// 	model.StartTimeUsecs = core.Int64Ptr(int64(26))
// 	model.Status = core.StringPtr("Scheduled")

// 	result, err := backuprecovery.DataSourceIbmBaasAgentUpgradeTasksAgentInfoObjectToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmBaasAgentUpgradeTasksErrorToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		model := make(map[string]interface{})
// 		model["error_code"] = "testString"
// 		model["message"] = "testString"
// 		model["task_log_id"] = "testString"

// 		assert.Equal(t, result, model)
// 	}

// 	model := new(backuprecoveryv1.Error)
// 	model.ErrorCode = core.StringPtr("testString")
// 	model.Message = core.StringPtr("testString")
// 	model.TaskLogID = core.StringPtr("testString")

// 	result, err := backuprecovery.DataSourceIbmBaasAgentUpgradeTasksErrorToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }
