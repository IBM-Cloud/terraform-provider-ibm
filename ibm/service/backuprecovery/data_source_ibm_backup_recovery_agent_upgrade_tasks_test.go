// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.94.0-fa797aec-20240814-142622
 */

package backuprecovery_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmBackupRecoveryAgentUpgradeTasksDataSourceBasic(t *testing.T) {
	name := fmt.Sprintf("tf_name_upgarde_task_%d", acctest.RandIntRange(10, 100))
	agentId := 346
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config:  testAccCheckIbmBackupRecoveryAgentUpgradeTasksDataSourceConfigBasic(name, agentId),
				Destroy: false,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_agent_upgrade_tasks.baas_agent_upgrade_tasks_instance", "id"),
					resource.TestCheckResourceAttr("data.ibm_backup_recovery_agent_upgrade_tasks.baas_agent_upgrade_tasks_instance", "tasks.#", "1"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_agent_upgrade_tasks.baas_agent_upgrade_tasks_instance", "tasks.0.description"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_agent_upgrade_tasks.baas_agent_upgrade_tasks_instance", "tasks.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_agent_upgrade_tasks.baas_agent_upgrade_tasks_instance", "x_ibm_tenant_id"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_agent_upgrade_tasks.baas_agent_upgrade_tasks_instance", "tasks.0.status"),
					resource.TestCheckResourceAttr("data.ibm_backup_recovery_agent_upgrade_tasks.baas_agent_upgrade_tasks_instance", "tasks.0.agent_i_ds.#", "1"),
					resource.TestCheckResourceAttr("data.ibm_backup_recovery_agent_upgrade_tasks.baas_agent_upgrade_tasks_instance", "tasks.0.agent_i_ds.0", strconv.Itoa(agentId)),
				),
			},
		},
	})
}

func testAccCheckIbmBackupRecoveryAgentUpgradeTasksDataSourceConfigBasic(name string, agentId int) string {
	return fmt.Sprintf(`
		resource "ibm_backup_recovery_agent_upgrade_task" "baas_agent_upgrade_task_instance" {
			x_ibm_tenant_id = "%s"
			agent_ids = [%d]
			backup_recovery_endpoint = "https://protectiondomain0103.us-east.backup-recovery-tests.cloud.ibm.com/v2"
			name = "%s"
			description = "Includes Agents for Sources RHEL, Win Server and MS SQL"
		}
		data "ibm_backup_recovery_agent_upgrade_tasks" "baas_agent_upgrade_tasks_instance" {
			x_ibm_tenant_id = "%[1]s"
			backup_recovery_endpoint = "https://protectiondomain0103.us-east.backup-recovery-tests.cloud.ibm.com/v2"
			ids = [ibm_backup_recovery_agent_upgrade_task.baas_agent_upgrade_task_instance.id]
		}
	`, tenantId, agentId, name)
}
