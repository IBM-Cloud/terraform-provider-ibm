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

func TestAccIbmBaasAgentUpgradeTaskBasic(t *testing.T) {
	var conf backuprecoveryv1.AgentUpgradeTaskStates
	xIbmTenantID := fmt.Sprintf("tf_x_ibm_tenant_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmBaasAgentUpgradeTaskDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBaasAgentUpgradeTaskConfigBasic(xIbmTenantID),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmBaasAgentUpgradeTaskExists("ibm_baas_agent_upgrade_task.baas_agent_upgrade_task_instance", conf),
					resource.TestCheckResourceAttr("ibm_baas_agent_upgrade_task.baas_agent_upgrade_task_instance", "x_ibm_tenant_id", xIbmTenantID),
				),
			},
		},
	})
}

func TestAccIbmBaasAgentUpgradeTaskAllArgs(t *testing.T) {
	var conf backuprecoveryv1.AgentUpgradeTaskStates
	xIbmTenantID := fmt.Sprintf("tf_x_ibm_tenant_id_%d", acctest.RandIntRange(10, 100))
	description := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	scheduleEndTimeUsecs := fmt.Sprintf("%d", acctest.RandIntRange(10, 100))
	scheduleTimeUsecs := fmt.Sprintf("%d", acctest.RandIntRange(10, 100))
	retryTaskId := fmt.Sprintf("%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmBaasAgentUpgradeTaskDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBaasAgentUpgradeTaskConfig(xIbmTenantID, description, name, scheduleEndTimeUsecs, scheduleTimeUsecs, retryTaskId),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmBaasAgentUpgradeTaskExists("ibm_baas_agent_upgrade_task.baas_agent_upgrade_task_instance", conf),
					resource.TestCheckResourceAttr("ibm_baas_agent_upgrade_task.baas_agent_upgrade_task_instance", "x_ibm_tenant_id", xIbmTenantID),
					resource.TestCheckResourceAttr("ibm_baas_agent_upgrade_task.baas_agent_upgrade_task_instance", "description", description),
					resource.TestCheckResourceAttr("ibm_baas_agent_upgrade_task.baas_agent_upgrade_task_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_baas_agent_upgrade_task.baas_agent_upgrade_task_instance", "schedule_end_time_usecs", scheduleEndTimeUsecs),
					resource.TestCheckResourceAttr("ibm_baas_agent_upgrade_task.baas_agent_upgrade_task_instance", "schedule_time_usecs", scheduleTimeUsecs),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_baas_agent_upgrade_task.baas_agent_upgrade_task",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmBaasAgentUpgradeTaskConfigBasic(xIbmTenantID string) string {
	return fmt.Sprintf(`
		resource "ibm_baas_agent_upgrade_task" "baas_agent_upgrade_task_instance" {
			x_ibm_tenant_id = "%s"
		}
	`, xIbmTenantID)
}

func testAccCheckIbmBaasAgentUpgradeTaskConfig(xIbmTenantID string, description string, name string, scheduleEndTimeUsecs string, scheduleTimeUsecs string, retryTaskId string) string {
	return fmt.Sprintf(`

		resource "ibm_baas_agent_upgrade_task" "baas_agent_upgrade_task_instance" {
			x_ibm_tenant_id = "%s"
			agent_ids = "FIXME"
			description = "%s"
			name = "%s"
			schedule_end_time_usecs = %s
			schedule_time_usecs = %s
			retry_task_id = "%s"
		}
	`, xIbmTenantID, description, name, scheduleEndTimeUsecs, scheduleTimeUsecs, retryTaskId)
}

func testAccCheckIbmBaasAgentUpgradeTaskExists(n string, obj backuprecoveryv1.AgentUpgradeTaskStates) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		_, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		backupRecoveryClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).BackupRecoveryV1()
		if err != nil {
			return err
		}

		getUpgradeTasksOptions := &backuprecoveryv1.GetUpgradeTasksOptions{}

		agentUpgradeTaskState, _, err := backupRecoveryClient.GetUpgradeTasks(getUpgradeTasksOptions)
		if err != nil {
			return err
		}

		obj = *agentUpgradeTaskState
		return nil
	}
}

func testAccCheckIbmBaasAgentUpgradeTaskDestroy(s *terraform.State) error {
	backupRecoveryClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).BackupRecoveryV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_baas_agent_upgrade_task" {
			continue
		}

		getUpgradeTasksOptions := &backuprecoveryv1.GetUpgradeTasksOptions{}

		// Try to find the key
		_, response, err := backupRecoveryClient.GetUpgradeTasks(getUpgradeTasksOptions)

		if err == nil {
			return fmt.Errorf("Agent upgrade task state still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for Agent upgrade task state (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
