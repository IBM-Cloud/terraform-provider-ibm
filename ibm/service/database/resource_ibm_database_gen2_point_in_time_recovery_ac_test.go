// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package database_test

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

const testAccGen2PointInTimeSampleSource = "crn:v1:bluemix:public:databases-for-postgresql:ca-mon:a/23b09aee04da4545b6e32805fa93249d:00000000-0000-0000-0000-000000000000::"

func TestAccIBMDatabaseInstancePostgresGen2PointInTimeRecovery(t *testing.T) {
	if acc.Gen2PointInTimeSourceId == "" || acc.Gen2PointInTime == "" {
		t.Skip("GEN2_POINT_IN_TIME_RECOVERY_SOURCE_ID and GEN2_POINT_IN_TIME_RECOVERY_TIME are not set")
	}
	t.Parallel()

	name := fmt.Sprintf("tf-gen2-pitr-%s", acctest.RandString(10))
	resourceName := "ibm_database." + name
	location := strings.Split(acc.Gen2PointInTimeSourceId, ":")[5]

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMDatabaseInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccIBMDatabaseGen2PointInTimeRecoveryConfig(name, location, acc.Gen2PointInTimeSourceId, acc.Gen2PointInTime, 14),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(resourceName, new(string)),
					resource.TestCheckResourceAttr(resourceName, "plan", "standard-gen2"),
					resource.TestCheckResourceAttr(resourceName, "point_in_time_recovery_deployment_id", acc.Gen2PointInTimeSourceId),
					resource.TestCheckResourceAttr(resourceName, "backups.0.point_in_time_recovery.0.retention_days", "14"),
				),
			},
			{
				Config: testAccIBMDatabaseGen2PointInTimeRecoveryConfig(name, location, acc.Gen2PointInTimeSourceId, acc.Gen2PointInTime, 10),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "backups.0.point_in_time_recovery.0.retention_days", "10"),
				),
			},
		},
	})
}

func TestAccIBMDatabaseInstancePostgresGen2PointInTimeRecoveryInvalid(t *testing.T) {
	t.Parallel()

	name := fmt.Sprintf("tf-gen2-pitr-%s", acctest.RandString(10))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      testAccIBMDatabaseGen2PointInTimeRecoveryInvalidConfig(name, "standard-gen2", "ca-mon", `point_in_time_recovery_time = ""`),
				PlanOnly:    true,
				ExpectError: regexp.MustCompile("no restore to the latest time"),
			},
			{
				Config:      testAccIBMDatabaseGen2PointInTimeRecoveryInvalidConfig(name, "standard-gen2", "us-east", `point_in_time_recovery_time = "2026-09-27T09:30:00Z"`),
				PlanOnly:    true,
				ExpectError: regexp.MustCompile("restores into the source's region only"),
			},
			{
				Config: testAccIBMDatabaseGen2PointInTimeRecoveryInvalidConfig(name, "standard-gen2", "ca-mon", `point_in_time_recovery_time = "2026-09-27T09:30:00Z"
  backup_id                   = "crn:v1:bluemix:public:databases-independent-backups:ca-mon:a/23b09aee04da4545b6e32805fa93249d:00000000-0000-0000-0000-000000000001::"`),
				PlanOnly:    true,
				ExpectError: regexp.MustCompile("backup_id cannot be combined"),
			},
			{
				Config: testAccIBMDatabaseGen2PointInTimeRecoveryInvalidConfig(name, "standard", "ca-mon", `service_endpoints = "private"
  backups {
    point_in_time_recovery {
      retention_days = 14
    }
  }`),
				PlanOnly:    true,
				ExpectError: regexp.MustCompile("not supported for Classic databases"),
			},
		},
	})
}

func testAccIBMDatabaseGen2PointInTimeRecoveryConfig(name, location, sourceCRN, pointInTime string, retentionDays int) string {
	return fmt.Sprintf(`
resource "ibm_database" %[1]q {
  name              = %[1]q
  service           = "databases-for-postgresql"
  plan              = "standard-gen2"
  location          = %[2]q
  service_endpoints = "private"

  point_in_time_recovery_deployment_id = %[3]q
  point_in_time_recovery_time          = %[4]q

  backups {
    point_in_time_recovery {
      retention_days = %[5]d
    }
  }

  timeouts {
    create = "180m"
  }
}
`, name, location, sourceCRN, pointInTime, retentionDays)
}

func testAccIBMDatabaseGen2PointInTimeRecoveryInvalidConfig(name, plan, location, extra string) string {
	return fmt.Sprintf(`
resource "ibm_database" %[1]q {
  name     = %[1]q
  service  = "databases-for-postgresql"
  plan     = %[2]q
  location = %[3]q

  point_in_time_recovery_deployment_id = %[4]q
  %[5]s
}
`, name, plan, location, testAccGen2PointInTimeSampleSource, extra)
}
