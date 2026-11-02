// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.114.3-943fbc81-20260603-173645
*/

package brsmigration_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/brsmigration"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/stretchr/testify/assert"
	"github.com/IBM/ibm-brs-migration-sdk-go/brsmigrationv1"
	. "github.com/IBM-Cloud/terraform-provider-ibm/ibm/unittest"
	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmBrsMigrationWorkloadHistoryDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBrsMigrationWorkloadHistoryDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_brs_migration_workload_history.brs_migration_workload_history_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_brs_migration_workload_history.brs_migration_workload_history_instance", "migration_id"),
					resource.TestCheckResourceAttrSet("data.ibm_brs_migration_workload_history.brs_migration_workload_history_instance", "workload_id"),
					resource.TestCheckResourceAttrSet("data.ibm_brs_migration_workload_history.brs_migration_workload_history_instance", "history.#"),
				),
			},
		},
	})
}

func testAccCheckIbmBrsMigrationWorkloadHistoryDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_brs_migration_workload_history" "brs_migration_workload_history_instance" {
			migration_id = "mgr-a1b2c3d4-e5f6-7890-abcd-ef12345678ab"
			workload_id = "wl-a1b2c3d4-e5f6-7890-abcd-ef1234567890"
		}
	`)
}


func TestDataSourceIbmBrsMigrationWorkloadHistoryWorkloadHistoryEntryToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["id"] = "hist-f6a7b8c9-d0e1-2345-f012-345678901234"
		model["state"] = "scheduled"
		model["started_at"] = "2019-01-01T12:00:00.000Z"
		model["completed_at"] = "2024-06-15T14:32:00.000Z"
		model["message"] = "Workload completed successfully."

		assert.Equal(t, result, model)
	}

	model := new(brsmigrationv1.WorkloadHistoryEntry)
	model.ID = core.StringPtr("hist-f6a7b8c9-d0e1-2345-f012-345678901234")
	model.State = core.StringPtr("scheduled")
	model.StartedAt = CreateMockDateTime("2019-01-01T12:00:00.000Z")
	model.CompletedAt = CreateMockDateTime("2024-06-15T14:32:00.000Z")
	model.Message = core.StringPtr("Workload completed successfully.")

	result, err := brsmigration.DataSourceIbmBrsMigrationWorkloadHistoryWorkloadHistoryEntryToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
