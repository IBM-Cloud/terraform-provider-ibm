// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package brsmigration_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/brsmigration"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/stretchr/testify/assert"
	"github.com/IBM/ibm-brs-migration-sdk-go/brsmigrationv1"
	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmBrsMigrationWorkloadBasic(t *testing.T) {
	var conf brsmigrationv1.Workload
	migrationID := fmt.Sprintf("tf_migration_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmBrsMigrationWorkloadDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBrsMigrationWorkloadConfigBasic(migrationID),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmBrsMigrationWorkloadExists("ibm_brs_migration_workload.brs_migration_workload_instance", conf),
					resource.TestCheckResourceAttr("ibm_brs_migration_workload.brs_migration_workload_instance", "migration_id", migrationID),
				),
			},
		},
	})
}

func TestAccIbmBrsMigrationWorkloadAllArgs(t *testing.T) {
	var conf brsmigrationv1.Workload
	migrationID := fmt.Sprintf("tf_migration_id_%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmBrsMigrationWorkloadDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBrsMigrationWorkloadConfig(migrationID, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmBrsMigrationWorkloadExists("ibm_brs_migration_workload.brs_migration_workload_instance", conf),
					resource.TestCheckResourceAttr("ibm_brs_migration_workload.brs_migration_workload_instance", "migration_id", migrationID),
					resource.TestCheckResourceAttr("ibm_brs_migration_workload.brs_migration_workload_instance", "name", name),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_brs_migration_workload.brs_migration_workload_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmBrsMigrationWorkloadConfigBasic(migrationID string) string {
	return fmt.Sprintf(`
		resource "ibm_brs_migration_workload" "brs_migration_workload_instance" {
			migration_id = "%s"
			payloads {
				id = "pl-c3d4e5f6-a7b8-9012-cdef-012345678901"
				source_host_id = "host-a1b2c3d4-e5f6-7890-abcd-ef1234567890"
				destination_host_id = "host-b2c3d4e5-f6a7-8901-bcde-f01234567890"
				data_specs {
					data_format = "raw"
					source {
						volume_id = "vol-b1c2d3e4-f5a6-7890-bcde-f01234567890"
						type = "ext4"
						path = "/mnt/data"
					}
					destination {
						volume_id = "vol-b1c2d3e4-f5a6-7890-bcde-f01234567890"
						type = "ext4"
						path = "/mnt/data"
					}
				}
			}
		}
	`, migrationID)
}

func testAccCheckIbmBrsMigrationWorkloadConfig(migrationID string, name string) string {
	return fmt.Sprintf(`

		resource "ibm_brs_migration_workload" "brs_migration_workload_instance" {
			migration_id = "%s"
			name = "%s"
			payloads {
				id = "pl-c3d4e5f6-a7b8-9012-cdef-012345678901"
				source_host_id = "host-a1b2c3d4-e5f6-7890-abcd-ef1234567890"
				destination_host_id = "host-b2c3d4e5-f6a7-8901-bcde-f01234567890"
				data_specs {
					data_format = "raw"
					source {
						volume_id = "vol-b1c2d3e4-f5a6-7890-bcde-f01234567890"
						type = "ext4"
						path = "/mnt/data"
					}
					destination {
						volume_id = "vol-b1c2d3e4-f5a6-7890-bcde-f01234567890"
						type = "ext4"
						path = "/mnt/data"
					}
				}
			}
		}
	`, migrationID, name)
}

func testAccCheckIbmBrsMigrationWorkloadExists(n string, obj brsmigrationv1.Workload) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		brsMigrationClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).BrsMigrationV1()
		if err != nil {
			return err
		}

		getWorkloadOptions := &brsmigrationv1.GetWorkloadOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getWorkloadOptions.SetMigrationID(parts[0])
		getWorkloadOptions.SetWorkloadID(parts[1])

		workload, _, err := brsMigrationClient.GetWorkload(getWorkloadOptions)
		if err != nil {
			return err
		}

		obj = *workload
		return nil
	}
}

func testAccCheckIbmBrsMigrationWorkloadDestroy(s *terraform.State) error {
	brsMigrationClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).BrsMigrationV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_brs_migration_workload" {
			continue
		}

		getWorkloadOptions := &brsmigrationv1.GetWorkloadOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getWorkloadOptions.SetMigrationID(parts[0])
		getWorkloadOptions.SetWorkloadID(parts[1])

		// Try to find the key
		_, response, err := brsMigrationClient.GetWorkload(getWorkloadOptions)

		if err == nil {
			return fmt.Errorf("brs_migration_workload still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for brs_migration_workload (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func TestResourceIbmBrsMigrationWorkloadWorkloadPayloadMappingToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		dataPayloadModel := make(map[string]interface{})
		dataPayloadModel["volume_id"] = "vol-b1c2d3e4-f5a6-7890-bcde-f01234567890"
		dataPayloadModel["type"] = "ext4"
		dataPayloadModel["path"] = "/mnt/data"

		dataSpecModel := make(map[string]interface{})
		dataSpecModel["data_format"] = "raw"
		dataSpecModel["source"] = []map[string]interface{}{dataPayloadModel}
		dataSpecModel["destination"] = []map[string]interface{}{dataPayloadModel}

		model := make(map[string]interface{})
		model["id"] = "pl-c3d4e5f6-a7b8-9012-cdef-012345678901"
		model["source_host_id"] = "host-a1b2c3d4-e5f6-7890-abcd-ef1234567890"
		model["destination_host_id"] = "host-b2c3d4e5-f6a7-8901-bcde-f01234567890"
		model["data_specs"] = []map[string]interface{}{dataSpecModel}

		assert.Equal(t, result, model)
	}

	dataPayloadModel := new(brsmigrationv1.DataPayload)
	dataPayloadModel.VolumeID = core.StringPtr("vol-b1c2d3e4-f5a6-7890-bcde-f01234567890")
	dataPayloadModel.Type = core.StringPtr("ext4")
	dataPayloadModel.Path = core.StringPtr("/mnt/data")

	dataSpecModel := new(brsmigrationv1.DataSpec)
	dataSpecModel.DataFormat = core.StringPtr("raw")
	dataSpecModel.Source = dataPayloadModel
	dataSpecModel.Destination = dataPayloadModel

	model := new(brsmigrationv1.WorkloadPayloadMapping)
	model.ID = core.StringPtr("pl-c3d4e5f6-a7b8-9012-cdef-012345678901")
	model.SourceHostID = core.StringPtr("host-a1b2c3d4-e5f6-7890-abcd-ef1234567890")
	model.DestinationHostID = core.StringPtr("host-b2c3d4e5-f6a7-8901-bcde-f01234567890")
	model.DataSpecs = []brsmigrationv1.DataSpec{*dataSpecModel}

	result, err := brsmigration.ResourceIbmBrsMigrationWorkloadWorkloadPayloadMappingToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBrsMigrationWorkloadDataSpecToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		dataPayloadModel := make(map[string]interface{})
		dataPayloadModel["volume_id"] = "vol-b1c2d3e4-f5a6-7890-bcde-f01234567890"
		dataPayloadModel["type"] = "ext4"
		dataPayloadModel["path"] = "/mnt/data"

		model := make(map[string]interface{})
		model["data_format"] = "raw"
		model["source"] = []map[string]interface{}{dataPayloadModel}
		model["destination"] = []map[string]interface{}{dataPayloadModel}

		assert.Equal(t, result, model)
	}

	dataPayloadModel := new(brsmigrationv1.DataPayload)
	dataPayloadModel.VolumeID = core.StringPtr("vol-b1c2d3e4-f5a6-7890-bcde-f01234567890")
	dataPayloadModel.Type = core.StringPtr("ext4")
	dataPayloadModel.Path = core.StringPtr("/mnt/data")

	model := new(brsmigrationv1.DataSpec)
	model.DataFormat = core.StringPtr("raw")
	model.Source = dataPayloadModel
	model.Destination = dataPayloadModel

	result, err := brsmigration.ResourceIbmBrsMigrationWorkloadDataSpecToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBrsMigrationWorkloadDataPayloadToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["volume_id"] = "vol-b1c2d3e4-f5a6-7890-bcde-f01234567890"
		model["type"] = "ext4"
		model["path"] = "/mnt/data"

		assert.Equal(t, result, model)
	}

	model := new(brsmigrationv1.DataPayload)
	model.VolumeID = core.StringPtr("vol-b1c2d3e4-f5a6-7890-bcde-f01234567890")
	model.Type = core.StringPtr("ext4")
	model.Path = core.StringPtr("/mnt/data")

	result, err := brsmigration.ResourceIbmBrsMigrationWorkloadDataPayloadToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBrsMigrationWorkloadWorkloadScheduleToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		incrementalScheduleMinuteScheduleModel := make(map[string]interface{})
		incrementalScheduleMinuteScheduleModel["frequency"] = int(1)

		incrementalScheduleHourScheduleModel := make(map[string]interface{})
		incrementalScheduleHourScheduleModel["frequency"] = int(1)

		incrementalScheduleDayScheduleModel := make(map[string]interface{})
		incrementalScheduleDayScheduleModel["frequency"] = int(1)

		incrementalScheduleWeekScheduleModel := make(map[string]interface{})
		incrementalScheduleWeekScheduleModel["day_of_week"] = []string{"sunday"}

		incrementalScheduleMonthScheduleModel := make(map[string]interface{})
		incrementalScheduleMonthScheduleModel["day_of_week"] = []string{"sunday"}
		incrementalScheduleMonthScheduleModel["week_of_month"] = "first"
		incrementalScheduleMonthScheduleModel["day_of_month"] = int(1)

		incrementalScheduleYearScheduleModel := make(map[string]interface{})
		incrementalScheduleYearScheduleModel["day_of_year"] = "first"

		incrementalScheduleModel := make(map[string]interface{})
		incrementalScheduleModel["unit"] = "minutes"
		incrementalScheduleModel["minute_schedule"] = []map[string]interface{}{incrementalScheduleMinuteScheduleModel}
		incrementalScheduleModel["hour_schedule"] = []map[string]interface{}{incrementalScheduleHourScheduleModel}
		incrementalScheduleModel["day_schedule"] = []map[string]interface{}{incrementalScheduleDayScheduleModel}
		incrementalScheduleModel["week_schedule"] = []map[string]interface{}{incrementalScheduleWeekScheduleModel}
		incrementalScheduleModel["month_schedule"] = []map[string]interface{}{incrementalScheduleMonthScheduleModel}
		incrementalScheduleModel["year_schedule"] = []map[string]interface{}{incrementalScheduleYearScheduleModel}

		regularBackupPolicyIncrementalModel := make(map[string]interface{})
		regularBackupPolicyIncrementalModel["schedule"] = []map[string]interface{}{incrementalScheduleModel}

		fullScheduleDayScheduleModel := make(map[string]interface{})
		fullScheduleDayScheduleModel["frequency"] = int(1)

		fullScheduleWeekScheduleModel := make(map[string]interface{})
		fullScheduleWeekScheduleModel["day_of_week"] = []string{"sunday"}

		fullScheduleMonthScheduleModel := make(map[string]interface{})
		fullScheduleMonthScheduleModel["day_of_week"] = []string{"sunday"}
		fullScheduleMonthScheduleModel["week_of_month"] = "first"
		fullScheduleMonthScheduleModel["day_of_month"] = int(1)

		fullScheduleYearScheduleModel := make(map[string]interface{})
		fullScheduleYearScheduleModel["day_of_year"] = "first"

		fullBackupPolicyScheduleModel := make(map[string]interface{})
		fullBackupPolicyScheduleModel["unit"] = "days"
		fullBackupPolicyScheduleModel["day_schedule"] = []map[string]interface{}{fullScheduleDayScheduleModel}
		fullBackupPolicyScheduleModel["week_schedule"] = []map[string]interface{}{fullScheduleWeekScheduleModel}
		fullBackupPolicyScheduleModel["month_schedule"] = []map[string]interface{}{fullScheduleMonthScheduleModel}
		fullBackupPolicyScheduleModel["year_schedule"] = []map[string]interface{}{fullScheduleYearScheduleModel}

		regularBackupPolicyFullModel := make(map[string]interface{})
		regularBackupPolicyFullModel["schedule"] = []map[string]interface{}{fullBackupPolicyScheduleModel}

		regularBackupPolicyRetentionModel := make(map[string]interface{})
		regularBackupPolicyRetentionModel["unit"] = "days"
		regularBackupPolicyRetentionModel["duration"] = int(1)

		regularBackupPolicyModel := make(map[string]interface{})
		regularBackupPolicyModel["incremental"] = []map[string]interface{}{regularBackupPolicyIncrementalModel}
		regularBackupPolicyModel["full"] = []map[string]interface{}{regularBackupPolicyFullModel}
		regularBackupPolicyModel["retention"] = []map[string]interface{}{regularBackupPolicyRetentionModel}

		logScheduleMinuteScheduleModel := make(map[string]interface{})
		logScheduleMinuteScheduleModel["frequency"] = int(1)

		logScheduleHourScheduleModel := make(map[string]interface{})
		logScheduleHourScheduleModel["frequency"] = int(1)

		logBackupPolicyScheduleModel := make(map[string]interface{})
		logBackupPolicyScheduleModel["unit"] = "minutes"
		logBackupPolicyScheduleModel["minute_schedule"] = []map[string]interface{}{logScheduleMinuteScheduleModel}
		logBackupPolicyScheduleModel["hour_schedule"] = []map[string]interface{}{logScheduleHourScheduleModel}

		backupPolicyLogModel := make(map[string]interface{})
		backupPolicyLogModel["schedule"] = []map[string]interface{}{logBackupPolicyScheduleModel}

		backupPolicyModel := make(map[string]interface{})
		backupPolicyModel["regular"] = []map[string]interface{}{regularBackupPolicyModel}
		backupPolicyModel["log"] = []map[string]interface{}{backupPolicyLogModel}

		timeOfDayModel := make(map[string]interface{})
		timeOfDayModel["hour"] = int(0)
		timeOfDayModel["minute"] = int(0)
		timeOfDayModel["timezone"] = "America/New_York"

		blackoutWindowModel := make(map[string]interface{})
		blackoutWindowModel["day"] = "sunday"
		blackoutWindowModel["start_time"] = []map[string]interface{}{timeOfDayModel}
		blackoutWindowModel["end_time"] = []map[string]interface{}{timeOfDayModel}

		extendedRetentionScheduleModel := make(map[string]interface{})
		extendedRetentionScheduleModel["unit"] = "runs"
		extendedRetentionScheduleModel["frequency"] = int(1)

		retentionModel := make(map[string]interface{})
		retentionModel["unit"] = "days"
		retentionModel["duration"] = int(1)

		extendedRetentionPolicyModel := make(map[string]interface{})
		extendedRetentionPolicyModel["schedule"] = []map[string]interface{}{extendedRetentionScheduleModel}
		extendedRetentionPolicyModel["retention"] = []map[string]interface{}{retentionModel}
		extendedRetentionPolicyModel["run_type"] = "regular"

		workloadScheduleRetryOptionsModel := make(map[string]interface{})
		workloadScheduleRetryOptionsModel["retries"] = int(0)
		workloadScheduleRetryOptionsModel["retry_interval_mins"] = int(1)

		model := make(map[string]interface{})
		model["name"] = "daily-incremental-weekly-full"
		model["description"] = "Daily incremental with weekly full backup and 30-day retention"
		model["backup_policy"] = []map[string]interface{}{backupPolicyModel}
		model["blackout_window"] = []map[string]interface{}{blackoutWindowModel}
		model["extended_retention"] = []map[string]interface{}{extendedRetentionPolicyModel}
		model["retry_options"] = []map[string]interface{}{workloadScheduleRetryOptionsModel}

		assert.Equal(t, result, model)
	}

	incrementalScheduleMinuteScheduleModel := new(brsmigrationv1.IncrementalScheduleMinuteSchedule)
	incrementalScheduleMinuteScheduleModel.Frequency = core.Int64Ptr(int64(1))

	incrementalScheduleHourScheduleModel := new(brsmigrationv1.IncrementalScheduleHourSchedule)
	incrementalScheduleHourScheduleModel.Frequency = core.Int64Ptr(int64(1))

	incrementalScheduleDayScheduleModel := new(brsmigrationv1.IncrementalScheduleDaySchedule)
	incrementalScheduleDayScheduleModel.Frequency = core.Int64Ptr(int64(1))

	incrementalScheduleWeekScheduleModel := new(brsmigrationv1.IncrementalScheduleWeekSchedule)
	incrementalScheduleWeekScheduleModel.DayOfWeek = []string{"sunday"}

	incrementalScheduleMonthScheduleModel := new(brsmigrationv1.IncrementalScheduleMonthSchedule)
	incrementalScheduleMonthScheduleModel.DayOfWeek = []string{"sunday"}
	incrementalScheduleMonthScheduleModel.WeekOfMonth = core.StringPtr("first")
	incrementalScheduleMonthScheduleModel.DayOfMonth = core.Int64Ptr(int64(1))

	incrementalScheduleYearScheduleModel := new(brsmigrationv1.IncrementalScheduleYearSchedule)
	incrementalScheduleYearScheduleModel.DayOfYear = core.StringPtr("first")

	incrementalScheduleModel := new(brsmigrationv1.IncrementalSchedule)
	incrementalScheduleModel.Unit = core.StringPtr("minutes")
	incrementalScheduleModel.MinuteSchedule = incrementalScheduleMinuteScheduleModel
	incrementalScheduleModel.HourSchedule = incrementalScheduleHourScheduleModel
	incrementalScheduleModel.DaySchedule = incrementalScheduleDayScheduleModel
	incrementalScheduleModel.WeekSchedule = incrementalScheduleWeekScheduleModel
	incrementalScheduleModel.MonthSchedule = incrementalScheduleMonthScheduleModel
	incrementalScheduleModel.YearSchedule = incrementalScheduleYearScheduleModel

	regularBackupPolicyIncrementalModel := new(brsmigrationv1.RegularBackupPolicyIncremental)
	regularBackupPolicyIncrementalModel.Schedule = incrementalScheduleModel

	fullScheduleDayScheduleModel := new(brsmigrationv1.FullScheduleDaySchedule)
	fullScheduleDayScheduleModel.Frequency = core.Int64Ptr(int64(1))

	fullScheduleWeekScheduleModel := new(brsmigrationv1.FullScheduleWeekSchedule)
	fullScheduleWeekScheduleModel.DayOfWeek = []string{"sunday"}

	fullScheduleMonthScheduleModel := new(brsmigrationv1.FullScheduleMonthSchedule)
	fullScheduleMonthScheduleModel.DayOfWeek = []string{"sunday"}
	fullScheduleMonthScheduleModel.WeekOfMonth = core.StringPtr("first")
	fullScheduleMonthScheduleModel.DayOfMonth = core.Int64Ptr(int64(1))

	fullScheduleYearScheduleModel := new(brsmigrationv1.FullScheduleYearSchedule)
	fullScheduleYearScheduleModel.DayOfYear = core.StringPtr("first")

	fullBackupPolicyScheduleModel := new(brsmigrationv1.FullBackupPolicySchedule)
	fullBackupPolicyScheduleModel.Unit = core.StringPtr("days")
	fullBackupPolicyScheduleModel.DaySchedule = fullScheduleDayScheduleModel
	fullBackupPolicyScheduleModel.WeekSchedule = fullScheduleWeekScheduleModel
	fullBackupPolicyScheduleModel.MonthSchedule = fullScheduleMonthScheduleModel
	fullBackupPolicyScheduleModel.YearSchedule = fullScheduleYearScheduleModel

	regularBackupPolicyFullModel := new(brsmigrationv1.RegularBackupPolicyFull)
	regularBackupPolicyFullModel.Schedule = fullBackupPolicyScheduleModel

	regularBackupPolicyRetentionModel := new(brsmigrationv1.RegularBackupPolicyRetention)
	regularBackupPolicyRetentionModel.Unit = core.StringPtr("days")
	regularBackupPolicyRetentionModel.Duration = core.Int64Ptr(int64(1))

	regularBackupPolicyModel := new(brsmigrationv1.RegularBackupPolicy)
	regularBackupPolicyModel.Incremental = regularBackupPolicyIncrementalModel
	regularBackupPolicyModel.Full = regularBackupPolicyFullModel
	regularBackupPolicyModel.Retention = regularBackupPolicyRetentionModel

	logScheduleMinuteScheduleModel := new(brsmigrationv1.LogScheduleMinuteSchedule)
	logScheduleMinuteScheduleModel.Frequency = core.Int64Ptr(int64(1))

	logScheduleHourScheduleModel := new(brsmigrationv1.LogScheduleHourSchedule)
	logScheduleHourScheduleModel.Frequency = core.Int64Ptr(int64(1))

	logBackupPolicyScheduleModel := new(brsmigrationv1.LogBackupPolicySchedule)
	logBackupPolicyScheduleModel.Unit = core.StringPtr("minutes")
	logBackupPolicyScheduleModel.MinuteSchedule = logScheduleMinuteScheduleModel
	logBackupPolicyScheduleModel.HourSchedule = logScheduleHourScheduleModel

	backupPolicyLogModel := new(brsmigrationv1.BackupPolicyLog)
	backupPolicyLogModel.Schedule = logBackupPolicyScheduleModel

	backupPolicyModel := new(brsmigrationv1.BackupPolicy)
	backupPolicyModel.Regular = regularBackupPolicyModel
	backupPolicyModel.Log = backupPolicyLogModel

	timeOfDayModel := new(brsmigrationv1.TimeOfDay)
	timeOfDayModel.Hour = core.Int64Ptr(int64(0))
	timeOfDayModel.Minute = core.Int64Ptr(int64(0))
	timeOfDayModel.Timezone = core.StringPtr("America/New_York")

	blackoutWindowModel := new(brsmigrationv1.BlackoutWindow)
	blackoutWindowModel.Day = core.StringPtr("sunday")
	blackoutWindowModel.StartTime = timeOfDayModel
	blackoutWindowModel.EndTime = timeOfDayModel

	extendedRetentionScheduleModel := new(brsmigrationv1.ExtendedRetentionSchedule)
	extendedRetentionScheduleModel.Unit = core.StringPtr("runs")
	extendedRetentionScheduleModel.Frequency = core.Int64Ptr(int64(1))

	retentionModel := new(brsmigrationv1.Retention)
	retentionModel.Unit = core.StringPtr("days")
	retentionModel.Duration = core.Int64Ptr(int64(1))

	extendedRetentionPolicyModel := new(brsmigrationv1.ExtendedRetentionPolicy)
	extendedRetentionPolicyModel.Schedule = extendedRetentionScheduleModel
	extendedRetentionPolicyModel.Retention = retentionModel
	extendedRetentionPolicyModel.RunType = core.StringPtr("regular")

	workloadScheduleRetryOptionsModel := new(brsmigrationv1.WorkloadScheduleRetryOptions)
	workloadScheduleRetryOptionsModel.Retries = core.Int64Ptr(int64(0))
	workloadScheduleRetryOptionsModel.RetryIntervalMins = core.Int64Ptr(int64(1))

	model := new(brsmigrationv1.WorkloadSchedule)
	model.Name = core.StringPtr("daily-incremental-weekly-full")
	model.Description = core.StringPtr("Daily incremental with weekly full backup and 30-day retention")
	model.BackupPolicy = backupPolicyModel
	model.BlackoutWindow = []brsmigrationv1.BlackoutWindow{*blackoutWindowModel}
	model.ExtendedRetention = []brsmigrationv1.ExtendedRetentionPolicy{*extendedRetentionPolicyModel}
	model.RetryOptions = workloadScheduleRetryOptionsModel

	result, err := brsmigration.ResourceIbmBrsMigrationWorkloadWorkloadScheduleToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBrsMigrationWorkloadBackupPolicyToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		incrementalScheduleMinuteScheduleModel := make(map[string]interface{})
		incrementalScheduleMinuteScheduleModel["frequency"] = int(1)

		incrementalScheduleHourScheduleModel := make(map[string]interface{})
		incrementalScheduleHourScheduleModel["frequency"] = int(1)

		incrementalScheduleDayScheduleModel := make(map[string]interface{})
		incrementalScheduleDayScheduleModel["frequency"] = int(1)

		incrementalScheduleWeekScheduleModel := make(map[string]interface{})
		incrementalScheduleWeekScheduleModel["day_of_week"] = []string{"sunday"}

		incrementalScheduleMonthScheduleModel := make(map[string]interface{})
		incrementalScheduleMonthScheduleModel["day_of_week"] = []string{"sunday"}
		incrementalScheduleMonthScheduleModel["week_of_month"] = "first"
		incrementalScheduleMonthScheduleModel["day_of_month"] = int(1)

		incrementalScheduleYearScheduleModel := make(map[string]interface{})
		incrementalScheduleYearScheduleModel["day_of_year"] = "first"

		incrementalScheduleModel := make(map[string]interface{})
		incrementalScheduleModel["unit"] = "minutes"
		incrementalScheduleModel["minute_schedule"] = []map[string]interface{}{incrementalScheduleMinuteScheduleModel}
		incrementalScheduleModel["hour_schedule"] = []map[string]interface{}{incrementalScheduleHourScheduleModel}
		incrementalScheduleModel["day_schedule"] = []map[string]interface{}{incrementalScheduleDayScheduleModel}
		incrementalScheduleModel["week_schedule"] = []map[string]interface{}{incrementalScheduleWeekScheduleModel}
		incrementalScheduleModel["month_schedule"] = []map[string]interface{}{incrementalScheduleMonthScheduleModel}
		incrementalScheduleModel["year_schedule"] = []map[string]interface{}{incrementalScheduleYearScheduleModel}

		regularBackupPolicyIncrementalModel := make(map[string]interface{})
		regularBackupPolicyIncrementalModel["schedule"] = []map[string]interface{}{incrementalScheduleModel}

		fullScheduleDayScheduleModel := make(map[string]interface{})
		fullScheduleDayScheduleModel["frequency"] = int(1)

		fullScheduleWeekScheduleModel := make(map[string]interface{})
		fullScheduleWeekScheduleModel["day_of_week"] = []string{"sunday"}

		fullScheduleMonthScheduleModel := make(map[string]interface{})
		fullScheduleMonthScheduleModel["day_of_week"] = []string{"sunday"}
		fullScheduleMonthScheduleModel["week_of_month"] = "first"
		fullScheduleMonthScheduleModel["day_of_month"] = int(1)

		fullScheduleYearScheduleModel := make(map[string]interface{})
		fullScheduleYearScheduleModel["day_of_year"] = "first"

		fullBackupPolicyScheduleModel := make(map[string]interface{})
		fullBackupPolicyScheduleModel["unit"] = "days"
		fullBackupPolicyScheduleModel["day_schedule"] = []map[string]interface{}{fullScheduleDayScheduleModel}
		fullBackupPolicyScheduleModel["week_schedule"] = []map[string]interface{}{fullScheduleWeekScheduleModel}
		fullBackupPolicyScheduleModel["month_schedule"] = []map[string]interface{}{fullScheduleMonthScheduleModel}
		fullBackupPolicyScheduleModel["year_schedule"] = []map[string]interface{}{fullScheduleYearScheduleModel}

		regularBackupPolicyFullModel := make(map[string]interface{})
		regularBackupPolicyFullModel["schedule"] = []map[string]interface{}{fullBackupPolicyScheduleModel}

		regularBackupPolicyRetentionModel := make(map[string]interface{})
		regularBackupPolicyRetentionModel["unit"] = "days"
		regularBackupPolicyRetentionModel["duration"] = int(1)

		regularBackupPolicyModel := make(map[string]interface{})
		regularBackupPolicyModel["incremental"] = []map[string]interface{}{regularBackupPolicyIncrementalModel}
		regularBackupPolicyModel["full"] = []map[string]interface{}{regularBackupPolicyFullModel}
		regularBackupPolicyModel["retention"] = []map[string]interface{}{regularBackupPolicyRetentionModel}

		logScheduleMinuteScheduleModel := make(map[string]interface{})
		logScheduleMinuteScheduleModel["frequency"] = int(1)

		logScheduleHourScheduleModel := make(map[string]interface{})
		logScheduleHourScheduleModel["frequency"] = int(1)

		logBackupPolicyScheduleModel := make(map[string]interface{})
		logBackupPolicyScheduleModel["unit"] = "minutes"
		logBackupPolicyScheduleModel["minute_schedule"] = []map[string]interface{}{logScheduleMinuteScheduleModel}
		logBackupPolicyScheduleModel["hour_schedule"] = []map[string]interface{}{logScheduleHourScheduleModel}

		backupPolicyLogModel := make(map[string]interface{})
		backupPolicyLogModel["schedule"] = []map[string]interface{}{logBackupPolicyScheduleModel}

		model := make(map[string]interface{})
		model["regular"] = []map[string]interface{}{regularBackupPolicyModel}
		model["log"] = []map[string]interface{}{backupPolicyLogModel}

		assert.Equal(t, result, model)
	}

	incrementalScheduleMinuteScheduleModel := new(brsmigrationv1.IncrementalScheduleMinuteSchedule)
	incrementalScheduleMinuteScheduleModel.Frequency = core.Int64Ptr(int64(1))

	incrementalScheduleHourScheduleModel := new(brsmigrationv1.IncrementalScheduleHourSchedule)
	incrementalScheduleHourScheduleModel.Frequency = core.Int64Ptr(int64(1))

	incrementalScheduleDayScheduleModel := new(brsmigrationv1.IncrementalScheduleDaySchedule)
	incrementalScheduleDayScheduleModel.Frequency = core.Int64Ptr(int64(1))

	incrementalScheduleWeekScheduleModel := new(brsmigrationv1.IncrementalScheduleWeekSchedule)
	incrementalScheduleWeekScheduleModel.DayOfWeek = []string{"sunday"}

	incrementalScheduleMonthScheduleModel := new(brsmigrationv1.IncrementalScheduleMonthSchedule)
	incrementalScheduleMonthScheduleModel.DayOfWeek = []string{"sunday"}
	incrementalScheduleMonthScheduleModel.WeekOfMonth = core.StringPtr("first")
	incrementalScheduleMonthScheduleModel.DayOfMonth = core.Int64Ptr(int64(1))

	incrementalScheduleYearScheduleModel := new(brsmigrationv1.IncrementalScheduleYearSchedule)
	incrementalScheduleYearScheduleModel.DayOfYear = core.StringPtr("first")

	incrementalScheduleModel := new(brsmigrationv1.IncrementalSchedule)
	incrementalScheduleModel.Unit = core.StringPtr("minutes")
	incrementalScheduleModel.MinuteSchedule = incrementalScheduleMinuteScheduleModel
	incrementalScheduleModel.HourSchedule = incrementalScheduleHourScheduleModel
	incrementalScheduleModel.DaySchedule = incrementalScheduleDayScheduleModel
	incrementalScheduleModel.WeekSchedule = incrementalScheduleWeekScheduleModel
	incrementalScheduleModel.MonthSchedule = incrementalScheduleMonthScheduleModel
	incrementalScheduleModel.YearSchedule = incrementalScheduleYearScheduleModel

	regularBackupPolicyIncrementalModel := new(brsmigrationv1.RegularBackupPolicyIncremental)
	regularBackupPolicyIncrementalModel.Schedule = incrementalScheduleModel

	fullScheduleDayScheduleModel := new(brsmigrationv1.FullScheduleDaySchedule)
	fullScheduleDayScheduleModel.Frequency = core.Int64Ptr(int64(1))

	fullScheduleWeekScheduleModel := new(brsmigrationv1.FullScheduleWeekSchedule)
	fullScheduleWeekScheduleModel.DayOfWeek = []string{"sunday"}

	fullScheduleMonthScheduleModel := new(brsmigrationv1.FullScheduleMonthSchedule)
	fullScheduleMonthScheduleModel.DayOfWeek = []string{"sunday"}
	fullScheduleMonthScheduleModel.WeekOfMonth = core.StringPtr("first")
	fullScheduleMonthScheduleModel.DayOfMonth = core.Int64Ptr(int64(1))

	fullScheduleYearScheduleModel := new(brsmigrationv1.FullScheduleYearSchedule)
	fullScheduleYearScheduleModel.DayOfYear = core.StringPtr("first")

	fullBackupPolicyScheduleModel := new(brsmigrationv1.FullBackupPolicySchedule)
	fullBackupPolicyScheduleModel.Unit = core.StringPtr("days")
	fullBackupPolicyScheduleModel.DaySchedule = fullScheduleDayScheduleModel
	fullBackupPolicyScheduleModel.WeekSchedule = fullScheduleWeekScheduleModel
	fullBackupPolicyScheduleModel.MonthSchedule = fullScheduleMonthScheduleModel
	fullBackupPolicyScheduleModel.YearSchedule = fullScheduleYearScheduleModel

	regularBackupPolicyFullModel := new(brsmigrationv1.RegularBackupPolicyFull)
	regularBackupPolicyFullModel.Schedule = fullBackupPolicyScheduleModel

	regularBackupPolicyRetentionModel := new(brsmigrationv1.RegularBackupPolicyRetention)
	regularBackupPolicyRetentionModel.Unit = core.StringPtr("days")
	regularBackupPolicyRetentionModel.Duration = core.Int64Ptr(int64(1))

	regularBackupPolicyModel := new(brsmigrationv1.RegularBackupPolicy)
	regularBackupPolicyModel.Incremental = regularBackupPolicyIncrementalModel
	regularBackupPolicyModel.Full = regularBackupPolicyFullModel
	regularBackupPolicyModel.Retention = regularBackupPolicyRetentionModel

	logScheduleMinuteScheduleModel := new(brsmigrationv1.LogScheduleMinuteSchedule)
	logScheduleMinuteScheduleModel.Frequency = core.Int64Ptr(int64(1))

	logScheduleHourScheduleModel := new(brsmigrationv1.LogScheduleHourSchedule)
	logScheduleHourScheduleModel.Frequency = core.Int64Ptr(int64(1))

	logBackupPolicyScheduleModel := new(brsmigrationv1.LogBackupPolicySchedule)
	logBackupPolicyScheduleModel.Unit = core.StringPtr("minutes")
	logBackupPolicyScheduleModel.MinuteSchedule = logScheduleMinuteScheduleModel
	logBackupPolicyScheduleModel.HourSchedule = logScheduleHourScheduleModel

	backupPolicyLogModel := new(brsmigrationv1.BackupPolicyLog)
	backupPolicyLogModel.Schedule = logBackupPolicyScheduleModel

	model := new(brsmigrationv1.BackupPolicy)
	model.Regular = regularBackupPolicyModel
	model.Log = backupPolicyLogModel

	result, err := brsmigration.ResourceIbmBrsMigrationWorkloadBackupPolicyToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBrsMigrationWorkloadRegularBackupPolicyToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		incrementalScheduleMinuteScheduleModel := make(map[string]interface{})
		incrementalScheduleMinuteScheduleModel["frequency"] = int(1)

		incrementalScheduleHourScheduleModel := make(map[string]interface{})
		incrementalScheduleHourScheduleModel["frequency"] = int(1)

		incrementalScheduleDayScheduleModel := make(map[string]interface{})
		incrementalScheduleDayScheduleModel["frequency"] = int(1)

		incrementalScheduleWeekScheduleModel := make(map[string]interface{})
		incrementalScheduleWeekScheduleModel["day_of_week"] = []string{"sunday"}

		incrementalScheduleMonthScheduleModel := make(map[string]interface{})
		incrementalScheduleMonthScheduleModel["day_of_week"] = []string{"sunday"}
		incrementalScheduleMonthScheduleModel["week_of_month"] = "first"
		incrementalScheduleMonthScheduleModel["day_of_month"] = int(1)

		incrementalScheduleYearScheduleModel := make(map[string]interface{})
		incrementalScheduleYearScheduleModel["day_of_year"] = "first"

		incrementalScheduleModel := make(map[string]interface{})
		incrementalScheduleModel["unit"] = "minutes"
		incrementalScheduleModel["minute_schedule"] = []map[string]interface{}{incrementalScheduleMinuteScheduleModel}
		incrementalScheduleModel["hour_schedule"] = []map[string]interface{}{incrementalScheduleHourScheduleModel}
		incrementalScheduleModel["day_schedule"] = []map[string]interface{}{incrementalScheduleDayScheduleModel}
		incrementalScheduleModel["week_schedule"] = []map[string]interface{}{incrementalScheduleWeekScheduleModel}
		incrementalScheduleModel["month_schedule"] = []map[string]interface{}{incrementalScheduleMonthScheduleModel}
		incrementalScheduleModel["year_schedule"] = []map[string]interface{}{incrementalScheduleYearScheduleModel}

		regularBackupPolicyIncrementalModel := make(map[string]interface{})
		regularBackupPolicyIncrementalModel["schedule"] = []map[string]interface{}{incrementalScheduleModel}

		fullScheduleDayScheduleModel := make(map[string]interface{})
		fullScheduleDayScheduleModel["frequency"] = int(1)

		fullScheduleWeekScheduleModel := make(map[string]interface{})
		fullScheduleWeekScheduleModel["day_of_week"] = []string{"sunday"}

		fullScheduleMonthScheduleModel := make(map[string]interface{})
		fullScheduleMonthScheduleModel["day_of_week"] = []string{"sunday"}
		fullScheduleMonthScheduleModel["week_of_month"] = "first"
		fullScheduleMonthScheduleModel["day_of_month"] = int(1)

		fullScheduleYearScheduleModel := make(map[string]interface{})
		fullScheduleYearScheduleModel["day_of_year"] = "first"

		fullBackupPolicyScheduleModel := make(map[string]interface{})
		fullBackupPolicyScheduleModel["unit"] = "days"
		fullBackupPolicyScheduleModel["day_schedule"] = []map[string]interface{}{fullScheduleDayScheduleModel}
		fullBackupPolicyScheduleModel["week_schedule"] = []map[string]interface{}{fullScheduleWeekScheduleModel}
		fullBackupPolicyScheduleModel["month_schedule"] = []map[string]interface{}{fullScheduleMonthScheduleModel}
		fullBackupPolicyScheduleModel["year_schedule"] = []map[string]interface{}{fullScheduleYearScheduleModel}

		regularBackupPolicyFullModel := make(map[string]interface{})
		regularBackupPolicyFullModel["schedule"] = []map[string]interface{}{fullBackupPolicyScheduleModel}

		regularBackupPolicyRetentionModel := make(map[string]interface{})
		regularBackupPolicyRetentionModel["unit"] = "days"
		regularBackupPolicyRetentionModel["duration"] = int(1)

		model := make(map[string]interface{})
		model["incremental"] = []map[string]interface{}{regularBackupPolicyIncrementalModel}
		model["full"] = []map[string]interface{}{regularBackupPolicyFullModel}
		model["retention"] = []map[string]interface{}{regularBackupPolicyRetentionModel}

		assert.Equal(t, result, model)
	}

	incrementalScheduleMinuteScheduleModel := new(brsmigrationv1.IncrementalScheduleMinuteSchedule)
	incrementalScheduleMinuteScheduleModel.Frequency = core.Int64Ptr(int64(1))

	incrementalScheduleHourScheduleModel := new(brsmigrationv1.IncrementalScheduleHourSchedule)
	incrementalScheduleHourScheduleModel.Frequency = core.Int64Ptr(int64(1))

	incrementalScheduleDayScheduleModel := new(brsmigrationv1.IncrementalScheduleDaySchedule)
	incrementalScheduleDayScheduleModel.Frequency = core.Int64Ptr(int64(1))

	incrementalScheduleWeekScheduleModel := new(brsmigrationv1.IncrementalScheduleWeekSchedule)
	incrementalScheduleWeekScheduleModel.DayOfWeek = []string{"sunday"}

	incrementalScheduleMonthScheduleModel := new(brsmigrationv1.IncrementalScheduleMonthSchedule)
	incrementalScheduleMonthScheduleModel.DayOfWeek = []string{"sunday"}
	incrementalScheduleMonthScheduleModel.WeekOfMonth = core.StringPtr("first")
	incrementalScheduleMonthScheduleModel.DayOfMonth = core.Int64Ptr(int64(1))

	incrementalScheduleYearScheduleModel := new(brsmigrationv1.IncrementalScheduleYearSchedule)
	incrementalScheduleYearScheduleModel.DayOfYear = core.StringPtr("first")

	incrementalScheduleModel := new(brsmigrationv1.IncrementalSchedule)
	incrementalScheduleModel.Unit = core.StringPtr("minutes")
	incrementalScheduleModel.MinuteSchedule = incrementalScheduleMinuteScheduleModel
	incrementalScheduleModel.HourSchedule = incrementalScheduleHourScheduleModel
	incrementalScheduleModel.DaySchedule = incrementalScheduleDayScheduleModel
	incrementalScheduleModel.WeekSchedule = incrementalScheduleWeekScheduleModel
	incrementalScheduleModel.MonthSchedule = incrementalScheduleMonthScheduleModel
	incrementalScheduleModel.YearSchedule = incrementalScheduleYearScheduleModel

	regularBackupPolicyIncrementalModel := new(brsmigrationv1.RegularBackupPolicyIncremental)
	regularBackupPolicyIncrementalModel.Schedule = incrementalScheduleModel

	fullScheduleDayScheduleModel := new(brsmigrationv1.FullScheduleDaySchedule)
	fullScheduleDayScheduleModel.Frequency = core.Int64Ptr(int64(1))

	fullScheduleWeekScheduleModel := new(brsmigrationv1.FullScheduleWeekSchedule)
	fullScheduleWeekScheduleModel.DayOfWeek = []string{"sunday"}

	fullScheduleMonthScheduleModel := new(brsmigrationv1.FullScheduleMonthSchedule)
	fullScheduleMonthScheduleModel.DayOfWeek = []string{"sunday"}
	fullScheduleMonthScheduleModel.WeekOfMonth = core.StringPtr("first")
	fullScheduleMonthScheduleModel.DayOfMonth = core.Int64Ptr(int64(1))

	fullScheduleYearScheduleModel := new(brsmigrationv1.FullScheduleYearSchedule)
	fullScheduleYearScheduleModel.DayOfYear = core.StringPtr("first")

	fullBackupPolicyScheduleModel := new(brsmigrationv1.FullBackupPolicySchedule)
	fullBackupPolicyScheduleModel.Unit = core.StringPtr("days")
	fullBackupPolicyScheduleModel.DaySchedule = fullScheduleDayScheduleModel
	fullBackupPolicyScheduleModel.WeekSchedule = fullScheduleWeekScheduleModel
	fullBackupPolicyScheduleModel.MonthSchedule = fullScheduleMonthScheduleModel
	fullBackupPolicyScheduleModel.YearSchedule = fullScheduleYearScheduleModel

	regularBackupPolicyFullModel := new(brsmigrationv1.RegularBackupPolicyFull)
	regularBackupPolicyFullModel.Schedule = fullBackupPolicyScheduleModel

	regularBackupPolicyRetentionModel := new(brsmigrationv1.RegularBackupPolicyRetention)
	regularBackupPolicyRetentionModel.Unit = core.StringPtr("days")
	regularBackupPolicyRetentionModel.Duration = core.Int64Ptr(int64(1))

	model := new(brsmigrationv1.RegularBackupPolicy)
	model.Incremental = regularBackupPolicyIncrementalModel
	model.Full = regularBackupPolicyFullModel
	model.Retention = regularBackupPolicyRetentionModel

	result, err := brsmigration.ResourceIbmBrsMigrationWorkloadRegularBackupPolicyToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBrsMigrationWorkloadRegularBackupPolicyIncrementalToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		incrementalScheduleMinuteScheduleModel := make(map[string]interface{})
		incrementalScheduleMinuteScheduleModel["frequency"] = int(1)

		incrementalScheduleHourScheduleModel := make(map[string]interface{})
		incrementalScheduleHourScheduleModel["frequency"] = int(1)

		incrementalScheduleDayScheduleModel := make(map[string]interface{})
		incrementalScheduleDayScheduleModel["frequency"] = int(1)

		incrementalScheduleWeekScheduleModel := make(map[string]interface{})
		incrementalScheduleWeekScheduleModel["day_of_week"] = []string{"sunday"}

		incrementalScheduleMonthScheduleModel := make(map[string]interface{})
		incrementalScheduleMonthScheduleModel["day_of_week"] = []string{"sunday"}
		incrementalScheduleMonthScheduleModel["week_of_month"] = "first"
		incrementalScheduleMonthScheduleModel["day_of_month"] = int(1)

		incrementalScheduleYearScheduleModel := make(map[string]interface{})
		incrementalScheduleYearScheduleModel["day_of_year"] = "first"

		incrementalScheduleModel := make(map[string]interface{})
		incrementalScheduleModel["unit"] = "minutes"
		incrementalScheduleModel["minute_schedule"] = []map[string]interface{}{incrementalScheduleMinuteScheduleModel}
		incrementalScheduleModel["hour_schedule"] = []map[string]interface{}{incrementalScheduleHourScheduleModel}
		incrementalScheduleModel["day_schedule"] = []map[string]interface{}{incrementalScheduleDayScheduleModel}
		incrementalScheduleModel["week_schedule"] = []map[string]interface{}{incrementalScheduleWeekScheduleModel}
		incrementalScheduleModel["month_schedule"] = []map[string]interface{}{incrementalScheduleMonthScheduleModel}
		incrementalScheduleModel["year_schedule"] = []map[string]interface{}{incrementalScheduleYearScheduleModel}

		model := make(map[string]interface{})
		model["schedule"] = []map[string]interface{}{incrementalScheduleModel}

		assert.Equal(t, result, model)
	}

	incrementalScheduleMinuteScheduleModel := new(brsmigrationv1.IncrementalScheduleMinuteSchedule)
	incrementalScheduleMinuteScheduleModel.Frequency = core.Int64Ptr(int64(1))

	incrementalScheduleHourScheduleModel := new(brsmigrationv1.IncrementalScheduleHourSchedule)
	incrementalScheduleHourScheduleModel.Frequency = core.Int64Ptr(int64(1))

	incrementalScheduleDayScheduleModel := new(brsmigrationv1.IncrementalScheduleDaySchedule)
	incrementalScheduleDayScheduleModel.Frequency = core.Int64Ptr(int64(1))

	incrementalScheduleWeekScheduleModel := new(brsmigrationv1.IncrementalScheduleWeekSchedule)
	incrementalScheduleWeekScheduleModel.DayOfWeek = []string{"sunday"}

	incrementalScheduleMonthScheduleModel := new(brsmigrationv1.IncrementalScheduleMonthSchedule)
	incrementalScheduleMonthScheduleModel.DayOfWeek = []string{"sunday"}
	incrementalScheduleMonthScheduleModel.WeekOfMonth = core.StringPtr("first")
	incrementalScheduleMonthScheduleModel.DayOfMonth = core.Int64Ptr(int64(1))

	incrementalScheduleYearScheduleModel := new(brsmigrationv1.IncrementalScheduleYearSchedule)
	incrementalScheduleYearScheduleModel.DayOfYear = core.StringPtr("first")

	incrementalScheduleModel := new(brsmigrationv1.IncrementalSchedule)
	incrementalScheduleModel.Unit = core.StringPtr("minutes")
	incrementalScheduleModel.MinuteSchedule = incrementalScheduleMinuteScheduleModel
	incrementalScheduleModel.HourSchedule = incrementalScheduleHourScheduleModel
	incrementalScheduleModel.DaySchedule = incrementalScheduleDayScheduleModel
	incrementalScheduleModel.WeekSchedule = incrementalScheduleWeekScheduleModel
	incrementalScheduleModel.MonthSchedule = incrementalScheduleMonthScheduleModel
	incrementalScheduleModel.YearSchedule = incrementalScheduleYearScheduleModel

	model := new(brsmigrationv1.RegularBackupPolicyIncremental)
	model.Schedule = incrementalScheduleModel

	result, err := brsmigration.ResourceIbmBrsMigrationWorkloadRegularBackupPolicyIncrementalToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBrsMigrationWorkloadIncrementalScheduleToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		incrementalScheduleMinuteScheduleModel := make(map[string]interface{})
		incrementalScheduleMinuteScheduleModel["frequency"] = int(1)

		incrementalScheduleHourScheduleModel := make(map[string]interface{})
		incrementalScheduleHourScheduleModel["frequency"] = int(1)

		incrementalScheduleDayScheduleModel := make(map[string]interface{})
		incrementalScheduleDayScheduleModel["frequency"] = int(1)

		incrementalScheduleWeekScheduleModel := make(map[string]interface{})
		incrementalScheduleWeekScheduleModel["day_of_week"] = []string{"sunday"}

		incrementalScheduleMonthScheduleModel := make(map[string]interface{})
		incrementalScheduleMonthScheduleModel["day_of_week"] = []string{"sunday"}
		incrementalScheduleMonthScheduleModel["week_of_month"] = "first"
		incrementalScheduleMonthScheduleModel["day_of_month"] = int(1)

		incrementalScheduleYearScheduleModel := make(map[string]interface{})
		incrementalScheduleYearScheduleModel["day_of_year"] = "first"

		model := make(map[string]interface{})
		model["unit"] = "minutes"
		model["minute_schedule"] = []map[string]interface{}{incrementalScheduleMinuteScheduleModel}
		model["hour_schedule"] = []map[string]interface{}{incrementalScheduleHourScheduleModel}
		model["day_schedule"] = []map[string]interface{}{incrementalScheduleDayScheduleModel}
		model["week_schedule"] = []map[string]interface{}{incrementalScheduleWeekScheduleModel}
		model["month_schedule"] = []map[string]interface{}{incrementalScheduleMonthScheduleModel}
		model["year_schedule"] = []map[string]interface{}{incrementalScheduleYearScheduleModel}

		assert.Equal(t, result, model)
	}

	incrementalScheduleMinuteScheduleModel := new(brsmigrationv1.IncrementalScheduleMinuteSchedule)
	incrementalScheduleMinuteScheduleModel.Frequency = core.Int64Ptr(int64(1))

	incrementalScheduleHourScheduleModel := new(brsmigrationv1.IncrementalScheduleHourSchedule)
	incrementalScheduleHourScheduleModel.Frequency = core.Int64Ptr(int64(1))

	incrementalScheduleDayScheduleModel := new(brsmigrationv1.IncrementalScheduleDaySchedule)
	incrementalScheduleDayScheduleModel.Frequency = core.Int64Ptr(int64(1))

	incrementalScheduleWeekScheduleModel := new(brsmigrationv1.IncrementalScheduleWeekSchedule)
	incrementalScheduleWeekScheduleModel.DayOfWeek = []string{"sunday"}

	incrementalScheduleMonthScheduleModel := new(brsmigrationv1.IncrementalScheduleMonthSchedule)
	incrementalScheduleMonthScheduleModel.DayOfWeek = []string{"sunday"}
	incrementalScheduleMonthScheduleModel.WeekOfMonth = core.StringPtr("first")
	incrementalScheduleMonthScheduleModel.DayOfMonth = core.Int64Ptr(int64(1))

	incrementalScheduleYearScheduleModel := new(brsmigrationv1.IncrementalScheduleYearSchedule)
	incrementalScheduleYearScheduleModel.DayOfYear = core.StringPtr("first")

	model := new(brsmigrationv1.IncrementalSchedule)
	model.Unit = core.StringPtr("minutes")
	model.MinuteSchedule = incrementalScheduleMinuteScheduleModel
	model.HourSchedule = incrementalScheduleHourScheduleModel
	model.DaySchedule = incrementalScheduleDayScheduleModel
	model.WeekSchedule = incrementalScheduleWeekScheduleModel
	model.MonthSchedule = incrementalScheduleMonthScheduleModel
	model.YearSchedule = incrementalScheduleYearScheduleModel

	result, err := brsmigration.ResourceIbmBrsMigrationWorkloadIncrementalScheduleToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBrsMigrationWorkloadIncrementalScheduleMinuteScheduleToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["frequency"] = int(1)

		assert.Equal(t, result, model)
	}

	model := new(brsmigrationv1.IncrementalScheduleMinuteSchedule)
	model.Frequency = core.Int64Ptr(int64(1))

	result, err := brsmigration.ResourceIbmBrsMigrationWorkloadIncrementalScheduleMinuteScheduleToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBrsMigrationWorkloadIncrementalScheduleHourScheduleToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["frequency"] = int(1)

		assert.Equal(t, result, model)
	}

	model := new(brsmigrationv1.IncrementalScheduleHourSchedule)
	model.Frequency = core.Int64Ptr(int64(1))

	result, err := brsmigration.ResourceIbmBrsMigrationWorkloadIncrementalScheduleHourScheduleToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBrsMigrationWorkloadIncrementalScheduleDayScheduleToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["frequency"] = int(1)

		assert.Equal(t, result, model)
	}

	model := new(brsmigrationv1.IncrementalScheduleDaySchedule)
	model.Frequency = core.Int64Ptr(int64(1))

	result, err := brsmigration.ResourceIbmBrsMigrationWorkloadIncrementalScheduleDayScheduleToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBrsMigrationWorkloadIncrementalScheduleWeekScheduleToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["day_of_week"] = []string{"sunday"}

		assert.Equal(t, result, model)
	}

	model := new(brsmigrationv1.IncrementalScheduleWeekSchedule)
	model.DayOfWeek = []string{"sunday"}

	result, err := brsmigration.ResourceIbmBrsMigrationWorkloadIncrementalScheduleWeekScheduleToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBrsMigrationWorkloadIncrementalScheduleMonthScheduleToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["day_of_week"] = []string{"sunday"}
		model["week_of_month"] = "first"
		model["day_of_month"] = int(1)

		assert.Equal(t, result, model)
	}

	model := new(brsmigrationv1.IncrementalScheduleMonthSchedule)
	model.DayOfWeek = []string{"sunday"}
	model.WeekOfMonth = core.StringPtr("first")
	model.DayOfMonth = core.Int64Ptr(int64(1))

	result, err := brsmigration.ResourceIbmBrsMigrationWorkloadIncrementalScheduleMonthScheduleToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBrsMigrationWorkloadIncrementalScheduleYearScheduleToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["day_of_year"] = "first"

		assert.Equal(t, result, model)
	}

	model := new(brsmigrationv1.IncrementalScheduleYearSchedule)
	model.DayOfYear = core.StringPtr("first")

	result, err := brsmigration.ResourceIbmBrsMigrationWorkloadIncrementalScheduleYearScheduleToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBrsMigrationWorkloadRegularBackupPolicyFullToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		fullScheduleDayScheduleModel := make(map[string]interface{})
		fullScheduleDayScheduleModel["frequency"] = int(1)

		fullScheduleWeekScheduleModel := make(map[string]interface{})
		fullScheduleWeekScheduleModel["day_of_week"] = []string{"sunday"}

		fullScheduleMonthScheduleModel := make(map[string]interface{})
		fullScheduleMonthScheduleModel["day_of_week"] = []string{"sunday"}
		fullScheduleMonthScheduleModel["week_of_month"] = "first"
		fullScheduleMonthScheduleModel["day_of_month"] = int(1)

		fullScheduleYearScheduleModel := make(map[string]interface{})
		fullScheduleYearScheduleModel["day_of_year"] = "first"

		fullBackupPolicyScheduleModel := make(map[string]interface{})
		fullBackupPolicyScheduleModel["unit"] = "days"
		fullBackupPolicyScheduleModel["day_schedule"] = []map[string]interface{}{fullScheduleDayScheduleModel}
		fullBackupPolicyScheduleModel["week_schedule"] = []map[string]interface{}{fullScheduleWeekScheduleModel}
		fullBackupPolicyScheduleModel["month_schedule"] = []map[string]interface{}{fullScheduleMonthScheduleModel}
		fullBackupPolicyScheduleModel["year_schedule"] = []map[string]interface{}{fullScheduleYearScheduleModel}

		model := make(map[string]interface{})
		model["schedule"] = []map[string]interface{}{fullBackupPolicyScheduleModel}

		assert.Equal(t, result, model)
	}

	fullScheduleDayScheduleModel := new(brsmigrationv1.FullScheduleDaySchedule)
	fullScheduleDayScheduleModel.Frequency = core.Int64Ptr(int64(1))

	fullScheduleWeekScheduleModel := new(brsmigrationv1.FullScheduleWeekSchedule)
	fullScheduleWeekScheduleModel.DayOfWeek = []string{"sunday"}

	fullScheduleMonthScheduleModel := new(brsmigrationv1.FullScheduleMonthSchedule)
	fullScheduleMonthScheduleModel.DayOfWeek = []string{"sunday"}
	fullScheduleMonthScheduleModel.WeekOfMonth = core.StringPtr("first")
	fullScheduleMonthScheduleModel.DayOfMonth = core.Int64Ptr(int64(1))

	fullScheduleYearScheduleModel := new(brsmigrationv1.FullScheduleYearSchedule)
	fullScheduleYearScheduleModel.DayOfYear = core.StringPtr("first")

	fullBackupPolicyScheduleModel := new(brsmigrationv1.FullBackupPolicySchedule)
	fullBackupPolicyScheduleModel.Unit = core.StringPtr("days")
	fullBackupPolicyScheduleModel.DaySchedule = fullScheduleDayScheduleModel
	fullBackupPolicyScheduleModel.WeekSchedule = fullScheduleWeekScheduleModel
	fullBackupPolicyScheduleModel.MonthSchedule = fullScheduleMonthScheduleModel
	fullBackupPolicyScheduleModel.YearSchedule = fullScheduleYearScheduleModel

	model := new(brsmigrationv1.RegularBackupPolicyFull)
	model.Schedule = fullBackupPolicyScheduleModel

	result, err := brsmigration.ResourceIbmBrsMigrationWorkloadRegularBackupPolicyFullToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBrsMigrationWorkloadFullBackupPolicyScheduleToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		fullScheduleDayScheduleModel := make(map[string]interface{})
		fullScheduleDayScheduleModel["frequency"] = int(1)

		fullScheduleWeekScheduleModel := make(map[string]interface{})
		fullScheduleWeekScheduleModel["day_of_week"] = []string{"sunday"}

		fullScheduleMonthScheduleModel := make(map[string]interface{})
		fullScheduleMonthScheduleModel["day_of_week"] = []string{"sunday"}
		fullScheduleMonthScheduleModel["week_of_month"] = "first"
		fullScheduleMonthScheduleModel["day_of_month"] = int(1)

		fullScheduleYearScheduleModel := make(map[string]interface{})
		fullScheduleYearScheduleModel["day_of_year"] = "first"

		model := make(map[string]interface{})
		model["unit"] = "days"
		model["day_schedule"] = []map[string]interface{}{fullScheduleDayScheduleModel}
		model["week_schedule"] = []map[string]interface{}{fullScheduleWeekScheduleModel}
		model["month_schedule"] = []map[string]interface{}{fullScheduleMonthScheduleModel}
		model["year_schedule"] = []map[string]interface{}{fullScheduleYearScheduleModel}

		assert.Equal(t, result, model)
	}

	fullScheduleDayScheduleModel := new(brsmigrationv1.FullScheduleDaySchedule)
	fullScheduleDayScheduleModel.Frequency = core.Int64Ptr(int64(1))

	fullScheduleWeekScheduleModel := new(brsmigrationv1.FullScheduleWeekSchedule)
	fullScheduleWeekScheduleModel.DayOfWeek = []string{"sunday"}

	fullScheduleMonthScheduleModel := new(brsmigrationv1.FullScheduleMonthSchedule)
	fullScheduleMonthScheduleModel.DayOfWeek = []string{"sunday"}
	fullScheduleMonthScheduleModel.WeekOfMonth = core.StringPtr("first")
	fullScheduleMonthScheduleModel.DayOfMonth = core.Int64Ptr(int64(1))

	fullScheduleYearScheduleModel := new(brsmigrationv1.FullScheduleYearSchedule)
	fullScheduleYearScheduleModel.DayOfYear = core.StringPtr("first")

	model := new(brsmigrationv1.FullBackupPolicySchedule)
	model.Unit = core.StringPtr("days")
	model.DaySchedule = fullScheduleDayScheduleModel
	model.WeekSchedule = fullScheduleWeekScheduleModel
	model.MonthSchedule = fullScheduleMonthScheduleModel
	model.YearSchedule = fullScheduleYearScheduleModel

	result, err := brsmigration.ResourceIbmBrsMigrationWorkloadFullBackupPolicyScheduleToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBrsMigrationWorkloadFullScheduleDayScheduleToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["frequency"] = int(1)

		assert.Equal(t, result, model)
	}

	model := new(brsmigrationv1.FullScheduleDaySchedule)
	model.Frequency = core.Int64Ptr(int64(1))

	result, err := brsmigration.ResourceIbmBrsMigrationWorkloadFullScheduleDayScheduleToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBrsMigrationWorkloadFullScheduleWeekScheduleToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["day_of_week"] = []string{"sunday"}

		assert.Equal(t, result, model)
	}

	model := new(brsmigrationv1.FullScheduleWeekSchedule)
	model.DayOfWeek = []string{"sunday"}

	result, err := brsmigration.ResourceIbmBrsMigrationWorkloadFullScheduleWeekScheduleToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBrsMigrationWorkloadFullScheduleMonthScheduleToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["day_of_week"] = []string{"sunday"}
		model["week_of_month"] = "first"
		model["day_of_month"] = int(1)

		assert.Equal(t, result, model)
	}

	model := new(brsmigrationv1.FullScheduleMonthSchedule)
	model.DayOfWeek = []string{"sunday"}
	model.WeekOfMonth = core.StringPtr("first")
	model.DayOfMonth = core.Int64Ptr(int64(1))

	result, err := brsmigration.ResourceIbmBrsMigrationWorkloadFullScheduleMonthScheduleToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBrsMigrationWorkloadFullScheduleYearScheduleToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["day_of_year"] = "first"

		assert.Equal(t, result, model)
	}

	model := new(brsmigrationv1.FullScheduleYearSchedule)
	model.DayOfYear = core.StringPtr("first")

	result, err := brsmigration.ResourceIbmBrsMigrationWorkloadFullScheduleYearScheduleToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBrsMigrationWorkloadRegularBackupPolicyRetentionToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["unit"] = "days"
		model["duration"] = int(1)

		assert.Equal(t, result, model)
	}

	model := new(brsmigrationv1.RegularBackupPolicyRetention)
	model.Unit = core.StringPtr("days")
	model.Duration = core.Int64Ptr(int64(1))

	result, err := brsmigration.ResourceIbmBrsMigrationWorkloadRegularBackupPolicyRetentionToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBrsMigrationWorkloadBackupPolicyLogToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		logScheduleMinuteScheduleModel := make(map[string]interface{})
		logScheduleMinuteScheduleModel["frequency"] = int(1)

		logScheduleHourScheduleModel := make(map[string]interface{})
		logScheduleHourScheduleModel["frequency"] = int(1)

		logBackupPolicyScheduleModel := make(map[string]interface{})
		logBackupPolicyScheduleModel["unit"] = "minutes"
		logBackupPolicyScheduleModel["minute_schedule"] = []map[string]interface{}{logScheduleMinuteScheduleModel}
		logBackupPolicyScheduleModel["hour_schedule"] = []map[string]interface{}{logScheduleHourScheduleModel}

		model := make(map[string]interface{})
		model["schedule"] = []map[string]interface{}{logBackupPolicyScheduleModel}

		assert.Equal(t, result, model)
	}

	logScheduleMinuteScheduleModel := new(brsmigrationv1.LogScheduleMinuteSchedule)
	logScheduleMinuteScheduleModel.Frequency = core.Int64Ptr(int64(1))

	logScheduleHourScheduleModel := new(brsmigrationv1.LogScheduleHourSchedule)
	logScheduleHourScheduleModel.Frequency = core.Int64Ptr(int64(1))

	logBackupPolicyScheduleModel := new(brsmigrationv1.LogBackupPolicySchedule)
	logBackupPolicyScheduleModel.Unit = core.StringPtr("minutes")
	logBackupPolicyScheduleModel.MinuteSchedule = logScheduleMinuteScheduleModel
	logBackupPolicyScheduleModel.HourSchedule = logScheduleHourScheduleModel

	model := new(brsmigrationv1.BackupPolicyLog)
	model.Schedule = logBackupPolicyScheduleModel

	result, err := brsmigration.ResourceIbmBrsMigrationWorkloadBackupPolicyLogToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBrsMigrationWorkloadLogBackupPolicyScheduleToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		logScheduleMinuteScheduleModel := make(map[string]interface{})
		logScheduleMinuteScheduleModel["frequency"] = int(1)

		logScheduleHourScheduleModel := make(map[string]interface{})
		logScheduleHourScheduleModel["frequency"] = int(1)

		model := make(map[string]interface{})
		model["unit"] = "minutes"
		model["minute_schedule"] = []map[string]interface{}{logScheduleMinuteScheduleModel}
		model["hour_schedule"] = []map[string]interface{}{logScheduleHourScheduleModel}

		assert.Equal(t, result, model)
	}

	logScheduleMinuteScheduleModel := new(brsmigrationv1.LogScheduleMinuteSchedule)
	logScheduleMinuteScheduleModel.Frequency = core.Int64Ptr(int64(1))

	logScheduleHourScheduleModel := new(brsmigrationv1.LogScheduleHourSchedule)
	logScheduleHourScheduleModel.Frequency = core.Int64Ptr(int64(1))

	model := new(brsmigrationv1.LogBackupPolicySchedule)
	model.Unit = core.StringPtr("minutes")
	model.MinuteSchedule = logScheduleMinuteScheduleModel
	model.HourSchedule = logScheduleHourScheduleModel

	result, err := brsmigration.ResourceIbmBrsMigrationWorkloadLogBackupPolicyScheduleToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBrsMigrationWorkloadLogScheduleMinuteScheduleToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["frequency"] = int(1)

		assert.Equal(t, result, model)
	}

	model := new(brsmigrationv1.LogScheduleMinuteSchedule)
	model.Frequency = core.Int64Ptr(int64(1))

	result, err := brsmigration.ResourceIbmBrsMigrationWorkloadLogScheduleMinuteScheduleToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBrsMigrationWorkloadLogScheduleHourScheduleToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["frequency"] = int(1)

		assert.Equal(t, result, model)
	}

	model := new(brsmigrationv1.LogScheduleHourSchedule)
	model.Frequency = core.Int64Ptr(int64(1))

	result, err := brsmigration.ResourceIbmBrsMigrationWorkloadLogScheduleHourScheduleToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBrsMigrationWorkloadBlackoutWindowToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		timeOfDayModel := make(map[string]interface{})
		timeOfDayModel["hour"] = int(0)
		timeOfDayModel["minute"] = int(0)
		timeOfDayModel["timezone"] = "America/New_York"

		model := make(map[string]interface{})
		model["day"] = "sunday"
		model["start_time"] = []map[string]interface{}{timeOfDayModel}
		model["end_time"] = []map[string]interface{}{timeOfDayModel}

		assert.Equal(t, result, model)
	}

	timeOfDayModel := new(brsmigrationv1.TimeOfDay)
	timeOfDayModel.Hour = core.Int64Ptr(int64(0))
	timeOfDayModel.Minute = core.Int64Ptr(int64(0))
	timeOfDayModel.Timezone = core.StringPtr("America/New_York")

	model := new(brsmigrationv1.BlackoutWindow)
	model.Day = core.StringPtr("sunday")
	model.StartTime = timeOfDayModel
	model.EndTime = timeOfDayModel

	result, err := brsmigration.ResourceIbmBrsMigrationWorkloadBlackoutWindowToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBrsMigrationWorkloadTimeOfDayToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["hour"] = int(0)
		model["minute"] = int(0)
		model["timezone"] = "America/New_York"

		assert.Equal(t, result, model)
	}

	model := new(brsmigrationv1.TimeOfDay)
	model.Hour = core.Int64Ptr(int64(0))
	model.Minute = core.Int64Ptr(int64(0))
	model.Timezone = core.StringPtr("America/New_York")

	result, err := brsmigration.ResourceIbmBrsMigrationWorkloadTimeOfDayToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBrsMigrationWorkloadExtendedRetentionPolicyToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		extendedRetentionScheduleModel := make(map[string]interface{})
		extendedRetentionScheduleModel["unit"] = "runs"
		extendedRetentionScheduleModel["frequency"] = int(1)

		retentionModel := make(map[string]interface{})
		retentionModel["unit"] = "days"
		retentionModel["duration"] = int(1)

		model := make(map[string]interface{})
		model["schedule"] = []map[string]interface{}{extendedRetentionScheduleModel}
		model["retention"] = []map[string]interface{}{retentionModel}
		model["run_type"] = "regular"

		assert.Equal(t, result, model)
	}

	extendedRetentionScheduleModel := new(brsmigrationv1.ExtendedRetentionSchedule)
	extendedRetentionScheduleModel.Unit = core.StringPtr("runs")
	extendedRetentionScheduleModel.Frequency = core.Int64Ptr(int64(1))

	retentionModel := new(brsmigrationv1.Retention)
	retentionModel.Unit = core.StringPtr("days")
	retentionModel.Duration = core.Int64Ptr(int64(1))

	model := new(brsmigrationv1.ExtendedRetentionPolicy)
	model.Schedule = extendedRetentionScheduleModel
	model.Retention = retentionModel
	model.RunType = core.StringPtr("regular")

	result, err := brsmigration.ResourceIbmBrsMigrationWorkloadExtendedRetentionPolicyToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBrsMigrationWorkloadExtendedRetentionScheduleToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["unit"] = "runs"
		model["frequency"] = int(1)

		assert.Equal(t, result, model)
	}

	model := new(brsmigrationv1.ExtendedRetentionSchedule)
	model.Unit = core.StringPtr("runs")
	model.Frequency = core.Int64Ptr(int64(1))

	result, err := brsmigration.ResourceIbmBrsMigrationWorkloadExtendedRetentionScheduleToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBrsMigrationWorkloadRetentionToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["unit"] = "days"
		model["duration"] = int(1)

		assert.Equal(t, result, model)
	}

	model := new(brsmigrationv1.Retention)
	model.Unit = core.StringPtr("days")
	model.Duration = core.Int64Ptr(int64(1))

	result, err := brsmigration.ResourceIbmBrsMigrationWorkloadRetentionToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBrsMigrationWorkloadWorkloadScheduleRetryOptionsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["retries"] = int(0)
		model["retry_interval_mins"] = int(1)

		assert.Equal(t, result, model)
	}

	model := new(brsmigrationv1.WorkloadScheduleRetryOptions)
	model.Retries = core.Int64Ptr(int64(0))
	model.RetryIntervalMins = core.Int64Ptr(int64(1))

	result, err := brsmigration.ResourceIbmBrsMigrationWorkloadWorkloadScheduleRetryOptionsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBrsMigrationWorkloadMapToWorkloadPayloadMappingInput(t *testing.T) {
	checkResult := func(result *brsmigrationv1.WorkloadPayloadMappingInput) {
		dataPayloadModel := new(brsmigrationv1.DataPayload)
		dataPayloadModel.VolumeID = core.StringPtr("vol-b1c2d3e4-f5a6-7890-bcde-f01234567890")
		dataPayloadModel.Type = core.StringPtr("ext4")
		dataPayloadModel.Path = core.StringPtr("/mnt/data")

		dataSpecModel := new(brsmigrationv1.DataSpec)
		dataSpecModel.DataFormat = core.StringPtr("raw")
		dataSpecModel.Source = dataPayloadModel
		dataSpecModel.Destination = dataPayloadModel

		model := new(brsmigrationv1.WorkloadPayloadMappingInput)
		model.SourceHostID = core.StringPtr("host-a1b2c3d4-e5f6-7890-abcd-ef1234567890")
		model.DestinationHostID = core.StringPtr("host-b2c3d4e5-f6a7-8901-bcde-f01234567890")
		model.DataSpecs = []brsmigrationv1.DataSpec{*dataSpecModel}

		assert.Equal(t, result, model)
	}

	dataPayloadModel := make(map[string]interface{})
	dataPayloadModel["volume_id"] = "vol-b1c2d3e4-f5a6-7890-bcde-f01234567890"
	dataPayloadModel["type"] = "ext4"
	dataPayloadModel["path"] = "/mnt/data"

	dataSpecModel := make(map[string]interface{})
	dataSpecModel["data_format"] = "raw"
	dataSpecModel["source"] = []interface{}{dataPayloadModel}
	dataSpecModel["destination"] = []interface{}{dataPayloadModel}

	model := make(map[string]interface{})
	model["source_host_id"] = "host-a1b2c3d4-e5f6-7890-abcd-ef1234567890"
	model["destination_host_id"] = "host-b2c3d4e5-f6a7-8901-bcde-f01234567890"
	model["data_specs"] = []interface{}{dataSpecModel}

	result, err := brsmigration.ResourceIbmBrsMigrationWorkloadMapToWorkloadPayloadMappingInput(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBrsMigrationWorkloadMapToDataSpec(t *testing.T) {
	checkResult := func(result *brsmigrationv1.DataSpec) {
		dataPayloadModel := new(brsmigrationv1.DataPayload)
		dataPayloadModel.VolumeID = core.StringPtr("vol-b1c2d3e4-f5a6-7890-bcde-f01234567890")
		dataPayloadModel.Type = core.StringPtr("ext4")
		dataPayloadModel.Path = core.StringPtr("/mnt/data")

		model := new(brsmigrationv1.DataSpec)
		model.DataFormat = core.StringPtr("raw")
		model.Source = dataPayloadModel
		model.Destination = dataPayloadModel

		assert.Equal(t, result, model)
	}

	dataPayloadModel := make(map[string]interface{})
	dataPayloadModel["volume_id"] = "vol-b1c2d3e4-f5a6-7890-bcde-f01234567890"
	dataPayloadModel["type"] = "ext4"
	dataPayloadModel["path"] = "/mnt/data"

	model := make(map[string]interface{})
	model["data_format"] = "raw"
	model["source"] = []interface{}{dataPayloadModel}
	model["destination"] = []interface{}{dataPayloadModel}

	result, err := brsmigration.ResourceIbmBrsMigrationWorkloadMapToDataSpec(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBrsMigrationWorkloadMapToDataPayload(t *testing.T) {
	checkResult := func(result *brsmigrationv1.DataPayload) {
		model := new(brsmigrationv1.DataPayload)
		model.VolumeID = core.StringPtr("vol-b1c2d3e4-f5a6-7890-bcde-f01234567890")
		model.Type = core.StringPtr("ext4")
		model.Path = core.StringPtr("/mnt/data")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["volume_id"] = "vol-b1c2d3e4-f5a6-7890-bcde-f01234567890"
	model["type"] = "ext4"
	model["path"] = "/mnt/data"

	result, err := brsmigration.ResourceIbmBrsMigrationWorkloadMapToDataPayload(model)
	assert.Nil(t, err)
	checkResult(result)
}
