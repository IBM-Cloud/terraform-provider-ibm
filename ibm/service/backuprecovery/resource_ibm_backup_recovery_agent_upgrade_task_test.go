// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package backuprecovery_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM/ibm-backup-recovery-sdk-go/backuprecoveryv1"
)

func TestAccIbmBackupRecoveryAgentUpgradeTaskBasic(t *testing.T) {
	var conf backuprecoveryv1.AgentUpgradeTaskStates
	name := fmt.Sprintf("tf_name_upgarde_task_%d", acctest.RandIntRange(10, 100))
	agentId := 346

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmBackupRecoveryAgentUpgradeTaskDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBackupRecoveryAgentUpgradeTaskConfigBasic(name, agentId),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmBackupRecoveryAgentUpgradeTaskExists("ibm_backup_recovery_agent_upgrade_task.baas_agent_upgrade_task_instance", conf),
					resource.TestCheckResourceAttr("ibm_backup_recovery_agent_upgrade_task.baas_agent_upgrade_task_instance", "x_ibm_tenant_id", tenantId),
				),
				Destroy: false,
			},
		},
	})
}

func testAccCheckIbmBackupRecoveryAgentUpgradeTaskConfigBasic(name string, agentId int) string {
	return fmt.Sprintf(`
		resource "ibm_backup_recovery_agent_upgrade_task" "baas_agent_upgrade_task_instance" {
			x_ibm_tenant_id = "%s"
			
			agent_ids = [%d]
			name = "%s"
		}
	`, tenantId, agentId, name)
}

func testAccCheckIbmBackupRecoveryAgentUpgradeTaskExists(n string, obj backuprecoveryv1.AgentUpgradeTaskStates) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		backupRecoveryClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).BackupRecoveryV1()
		if err != nil {
			return err
		}

		getUpgradeTasksOptions := &backuprecoveryv1.GetUpgradeTasksOptions{}
		taskId, _ := strconv.Atoi(rs.Primary.ID)
		getUpgradeTasksOptions.SetXIBMTenantID(tenantId)
		getUpgradeTasksOptions.SetIds([]int64{int64(taskId)})

		agentUpgradeTaskState, _, err := backupRecoveryClient.GetUpgradeTasks(getUpgradeTasksOptions)
		if err != nil {
			return err
		}

		if len(agentUpgradeTaskState.Tasks) > 0 {
			return nil
		} else {
			return fmt.Errorf("Not found: %s", n)
		}
	}
}

func testAccCheckIbmBackupRecoveryAgentUpgradeTaskDestroy(s *terraform.State) error {
	return nil
}
