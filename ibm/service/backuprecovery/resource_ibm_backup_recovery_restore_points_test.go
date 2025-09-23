// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package backuprecovery_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/IBM/ibm-backup-recovery-sdk-go/backuprecoveryv1"
	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"
)

func TestAccIbmBackupRecoveryRestorePointsBasic(t *testing.T) {
	xIbmTenantID := fmt.Sprintf("tf_x_ibm_tenant_id_%d", acctest.RandIntRange(10, 100))
	endTimeUsecs := fmt.Sprintf("%d", acctest.RandIntRange(10, 100))
	startTimeUsecs := fmt.Sprintf("%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBackupRecoveryRestorePointsConfigBasic(xIbmTenantID, endTimeUsecs, startTimeUsecs),
				Check:  resource.ComposeAggregateTestCheckFunc(),
			},
		},
	})
}

func testAccCheckIbmBackupRecoveryRestorePointsConfigBasic(xIbmTenantID string, endTimeUsecs string, startTimeUsecs string) string {
	return fmt.Sprintf(`
	`)
}

func testAccCheckIbmBackupRecoveryRestorePointsExists(n string, obj backuprecoveryv1.GetRestorePointsInTimeRangeResponse) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		_, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		backupRecoveryClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).BackupRecoveryV1()
		if err != nil {
			return err
		}

		getRestorePointsInTimeRangeOptions := &backuprecoveryv1.GetRestorePointsInTimeRangeOptions{}

		getRestorePointsInTimeRangeParams, _, err := backupRecoveryClient.GetRestorePointsInTimeRange(getRestorePointsInTimeRangeOptions)
		if err != nil {
			return err
		}

		obj = *getRestorePointsInTimeRangeParams
		return nil
	}
}
