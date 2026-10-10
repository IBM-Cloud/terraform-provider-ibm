// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package database_test

import (
	"fmt"
	"os"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMDatabasePointInTimeRecoveryGen2DataSourceBasic(t *testing.T) {
	sourceID := os.Getenv("GEN2_POINT_IN_TIME_RECOVERY_SOURCE_ID")
	if sourceID == "" {
		t.Skip("Set GEN2_POINT_IN_TIME_RECOVERY_SOURCE_ID to the CRN of a Gen2 instance that can be restored to a point in time")
	}
	const dataSource = "data.ibm_database_point_in_time_recovery.source"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acc.TestAccPreCheck(t) },
		ProviderFactories: acc.TestAccProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMDatabasePointInTimeRecoveryGen2DataSourceConfig(sourceID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSource, "deployment_id", sourceID),
					resource.TestCheckResourceAttrSet(dataSource, "earliest_point_in_time_recovery_time"),
					resource.TestCheckResourceAttrSet(dataSource, "latest_point_in_time_recovery_time"),
					resource.TestCheckResourceAttrSet(dataSource, "retention_days"),
					resource.TestCheckResourceAttrSet(dataSource, "archiving_status"),
					resource.TestCheckResourceAttr(dataSource, "not_restorable_reason", ""),
				),
			},
		},
	})
}

func testAccCheckIBMDatabasePointInTimeRecoveryGen2DataSourceConfig(sourceID string) string {
	return fmt.Sprintf(`
		data "ibm_database_point_in_time_recovery" "source" {
			deployment_id = %q
		}
	`, sourceID)
}
