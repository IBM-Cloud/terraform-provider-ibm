// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.114.3-943fbc81-20260603-173645
*/

package brsmigration_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/brsmigration"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/stretchr/testify/assert"
	"github.ibm.com/BackupAndRecovery/brs-migration-orchestrator/brsmigrationv2"
	. "github.com/IBM-Cloud/terraform-provider-ibm/ibm/unittest"
	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmBrsMigrationWorkloadsDataSourceBasic(t *testing.T) {
	workloadMigrationID := fmt.Sprintf("tf_migration_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBrsMigrationWorkloadsDataSourceConfigBasic(workloadMigrationID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_brs_migration_workloads.brs_migration_workloads_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_brs_migration_workloads.brs_migration_workloads_instance", "migration_id"),
					resource.TestCheckResourceAttrSet("data.ibm_brs_migration_workloads.brs_migration_workloads_instance", "workloads.#"),
				),
			},
		},
	})
}

func TestAccIbmBrsMigrationWorkloadsDataSourceAllArgs(t *testing.T) {
	workloadMigrationID := fmt.Sprintf("tf_migration_id_%d", acctest.RandIntRange(10, 100))
	workloadName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBrsMigrationWorkloadsDataSourceConfig(workloadMigrationID, workloadName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_brs_migration_workloads.brs_migration_workloads_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_brs_migration_workloads.brs_migration_workloads_instance", "migration_id"),
					resource.TestCheckResourceAttrSet("data.ibm_brs_migration_workloads.brs_migration_workloads_instance", "workloads.#"),
					resource.TestCheckResourceAttrSet("data.ibm_brs_migration_workloads.brs_migration_workloads_instance", "workloads.0.id"),
					resource.TestCheckResourceAttr("data.ibm_brs_migration_workloads.brs_migration_workloads_instance", "workloads.0.name", workloadName),
					resource.TestCheckResourceAttrSet("data.ibm_brs_migration_workloads.brs_migration_workloads_instance", "workloads.0.state"),
					resource.TestCheckResourceAttrSet("data.ibm_brs_migration_workloads.brs_migration_workloads_instance", "workloads.0.scheduling_error"),
					resource.TestCheckResourceAttrSet("data.ibm_brs_migration_workloads.brs_migration_workloads_instance", "workloads.0.created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_brs_migration_workloads.brs_migration_workloads_instance", "workloads.0.updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_brs_migration_workloads.brs_migration_workloads_instance", "workloads.0.completed_at"),
				),
			},
		},
	})
}

func testAccCheckIbmBrsMigrationWorkloadsDataSourceConfigBasic(workloadMigrationID string) string {
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

		data "ibm_brs_migration_workloads" "brs_migration_workloads_instance" {
			migration_id = ibm_brs_migration_workload.brs_migration_workload_instance.migration_id
		}
	`, workloadMigrationID)
}

func testAccCheckIbmBrsMigrationWorkloadsDataSourceConfig(workloadMigrationID string, workloadName string) string {
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

		data "ibm_brs_migration_workloads" "brs_migration_workloads_instance" {
			migration_id = ibm_brs_migration_workload.brs_migration_workload_instance.migration_id
		}
	`, workloadMigrationID, workloadName)
}

func TestDataSourceIbmBrsMigrationWorkloadsWorkloadToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		dataPayloadModel := make(map[string]interface{})
		dataPayloadModel["volume_id"] = "vol-b1c2d3e4-f5a6-7890-bcde-f01234567890"
		dataPayloadModel["type"] = "ext4"
		dataPayloadModel["path"] = "/mnt/data"

		dataSpecModel := make(map[string]interface{})
		dataSpecModel["data_format"] = "raw"
		dataSpecModel["source"] = []map[string]interface{}{dataPayloadModel}
		dataSpecModel["destination"] = []map[string]interface{}{dataPayloadModel}

		workloadPayloadMappingModel := make(map[string]interface{})
		workloadPayloadMappingModel["id"] = "pl-c3d4e5f6-a7b8-9012-cdef-012345678901"
		workloadPayloadMappingModel["source_host_id"] = "host-a1b2c3d4-e5f6-7890-abcd-ef1234567890"
		workloadPayloadMappingModel["destination_host_id"] = "host-b2c3d4e5-f6a7-8901-bcde-f01234567890"
		workloadPayloadMappingModel["data_specs"] = []map[string]interface{}{dataSpecModel}

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

		workloadScheduleModel := make(map[string]interface{})
		workloadScheduleModel["name"] = "daily-incremental-weekly-full"
		workloadScheduleModel["description"] = "Daily incremental with weekly full backup and 30-day retention"
		workloadScheduleModel["backup_policy"] = []map[string]interface{}{backupPolicyModel}
		workloadScheduleModel["blackout_window"] = []map[string]interface{}{blackoutWindowModel}
		workloadScheduleModel["extended_retention"] = []map[string]interface{}{extendedRetentionPolicyModel}
		workloadScheduleModel["retry_options"] = []map[string]interface{}{workloadScheduleRetryOptionsModel}

		model := make(map[string]interface{})
		model["id"] = "wl-a1b2c3d4-e5f6-7890-abcd-ef1234567890"
		model["name"] = "prod-to-vpc-migration"
		model["volume_ownership_map"] = map[string]interface{}{"key1": "testString"}
		model["state"] = "created"
		model["payloads"] = []map[string]interface{}{workloadPayloadMappingModel}
		model["schedule"] = []map[string]interface{}{workloadScheduleModel}
		model["scheduling_error"] = "BRS protection group creation failed: policy not found"
		model["created_at"] = "2019-01-01T12:00:00.000Z"
		model["updated_at"] = "2019-01-01T12:00:00.000Z"
		model["completed_at"] = "2019-01-01T12:00:00.000Z"

		assert.Equal(t, result, model)
	}

	dataPayloadModel := new(brsmigrationv2.DataPayload)
	dataPayloadModel.VolumeID = core.StringPtr("vol-b1c2d3e4-f5a6-7890-bcde-f01234567890")
	dataPayloadModel.Type = core.StringPtr("ext4")
	dataPayloadModel.Path = core.StringPtr("/mnt/data")

	dataSpecModel := new(brsmigrationv2.DataSpec)
	dataSpecModel.DataFormat = core.StringPtr("raw")
	dataSpecModel.Source = dataPayloadModel
	dataSpecModel.Destination = dataPayloadModel

	workloadPayloadMappingModel := new(brsmigrationv2.WorkloadPayloadMapping)
	workloadPayloadMappingModel.ID = core.StringPtr("pl-c3d4e5f6-a7b8-9012-cdef-012345678901")
	workloadPayloadMappingModel.SourceHostID = core.StringPtr("host-a1b2c3d4-e5f6-7890-abcd-ef1234567890")
	workloadPayloadMappingModel.DestinationHostID = core.StringPtr("host-b2c3d4e5-f6a7-8901-bcde-f01234567890")
	workloadPayloadMappingModel.DataSpecs = []brsmigrationv2.DataSpec{*dataSpecModel}

	incrementalScheduleMinuteScheduleModel := new(brsmigrationv2.IncrementalScheduleMinuteSchedule)
	incrementalScheduleMinuteScheduleModel.Frequency = core.Int64Ptr(int64(1))

	incrementalScheduleHourScheduleModel := new(brsmigrationv2.IncrementalScheduleHourSchedule)
	incrementalScheduleHourScheduleModel.Frequency = core.Int64Ptr(int64(1))

	incrementalScheduleDayScheduleModel := new(brsmigrationv2.IncrementalScheduleDaySchedule)
	incrementalScheduleDayScheduleModel.Frequency = core.Int64Ptr(int64(1))

	incrementalScheduleWeekScheduleModel := new(brsmigrationv2.IncrementalScheduleWeekSchedule)
	incrementalScheduleWeekScheduleModel.DayOfWeek = []string{"sunday"}

	incrementalScheduleMonthScheduleModel := new(brsmigrationv2.IncrementalScheduleMonthSchedule)
	incrementalScheduleMonthScheduleModel.DayOfWeek = []string{"sunday"}
	incrementalScheduleMonthScheduleModel.WeekOfMonth = core.StringPtr("first")
	incrementalScheduleMonthScheduleModel.DayOfMonth = core.Int64Ptr(int64(1))

	incrementalScheduleYearScheduleModel := new(brsmigrationv2.IncrementalScheduleYearSchedule)
	incrementalScheduleYearScheduleModel.DayOfYear = core.StringPtr("first")

	incrementalScheduleModel := new(brsmigrationv2.IncrementalSchedule)
	incrementalScheduleModel.Unit = core.StringPtr("minutes")
	incrementalScheduleModel.MinuteSchedule = incrementalScheduleMinuteScheduleModel
	incrementalScheduleModel.HourSchedule = incrementalScheduleHourScheduleModel
	incrementalScheduleModel.DaySchedule = incrementalScheduleDayScheduleModel
	incrementalScheduleModel.WeekSchedule = incrementalScheduleWeekScheduleModel
	incrementalScheduleModel.MonthSchedule = incrementalScheduleMonthScheduleModel
	incrementalScheduleModel.YearSchedule = incrementalScheduleYearScheduleModel

	regularBackupPolicyIncrementalModel := new(brsmigrationv2.RegularBackupPolicyIncremental)
	regularBackupPolicyIncrementalModel.Schedule = incrementalScheduleModel

	fullScheduleDayScheduleModel := new(brsmigrationv2.FullScheduleDaySchedule)
	fullScheduleDayScheduleModel.Frequency = core.Int64Ptr(int64(1))

	fullScheduleWeekScheduleModel := new(brsmigrationv2.FullScheduleWeekSchedule)
	fullScheduleWeekScheduleModel.DayOfWeek = []string{"sunday"}

	fullScheduleMonthScheduleModel := new(brsmigrationv2.FullScheduleMonthSchedule)
	fullScheduleMonthScheduleModel.DayOfWeek = []string{"sunday"}
	fullScheduleMonthScheduleModel.WeekOfMonth = core.StringPtr("first")
	fullScheduleMonthScheduleModel.DayOfMonth = core.Int64Ptr(int64(1))

	fullScheduleYearScheduleModel := new(brsmigrationv2.FullScheduleYearSchedule)
	fullScheduleYearScheduleModel.DayOfYear = core.StringPtr("first")

	fullBackupPolicyScheduleModel := new(brsmigrationv2.FullBackupPolicySchedule)
	fullBackupPolicyScheduleModel.Unit = core.StringPtr("days")
	fullBackupPolicyScheduleModel.DaySchedule = fullScheduleDayScheduleModel
	fullBackupPolicyScheduleModel.WeekSchedule = fullScheduleWeekScheduleModel
	fullBackupPolicyScheduleModel.MonthSchedule = fullScheduleMonthScheduleModel
	fullBackupPolicyScheduleModel.YearSchedule = fullScheduleYearScheduleModel

	regularBackupPolicyFullModel := new(brsmigrationv2.RegularBackupPolicyFull)
	regularBackupPolicyFullModel.Schedule = fullBackupPolicyScheduleModel

	regularBackupPolicyRetentionModel := new(brsmigrationv2.RegularBackupPolicyRetention)
	regularBackupPolicyRetentionModel.Unit = core.StringPtr("days")
	regularBackupPolicyRetentionModel.Duration = core.Int64Ptr(int64(1))

	regularBackupPolicyModel := new(brsmigrationv2.RegularBackupPolicy)
	regularBackupPolicyModel.Incremental = regularBackupPolicyIncrementalModel
	regularBackupPolicyModel.Full = regularBackupPolicyFullModel
	regularBackupPolicyModel.Retention = regularBackupPolicyRetentionModel

	logScheduleMinuteScheduleModel := new(brsmigrationv2.LogScheduleMinuteSchedule)
	logScheduleMinuteScheduleModel.Frequency = core.Int64Ptr(int64(1))

	logScheduleHourScheduleModel := new(brsmigrationv2.LogScheduleHourSchedule)
	logScheduleHourScheduleModel.Frequency = core.Int64Ptr(int64(1))

	logBackupPolicyScheduleModel := new(brsmigrationv2.LogBackupPolicySchedule)
	logBackupPolicyScheduleModel.Unit = core.StringPtr("minutes")
	logBackupPolicyScheduleModel.MinuteSchedule = logScheduleMinuteScheduleModel
	logBackupPolicyScheduleModel.HourSchedule = logScheduleHourScheduleModel

	backupPolicyLogModel := new(brsmigrationv2.BackupPolicyLog)
	backupPolicyLogModel.Schedule = logBackupPolicyScheduleModel

	backupPolicyModel := new(brsmigrationv2.BackupPolicy)
	backupPolicyModel.Regular = regularBackupPolicyModel
	backupPolicyModel.Log = backupPolicyLogModel

	timeOfDayModel := new(brsmigrationv2.TimeOfDay)
	timeOfDayModel.Hour = core.Int64Ptr(int64(0))
	timeOfDayModel.Minute = core.Int64Ptr(int64(0))
	timeOfDayModel.Timezone = core.StringPtr("America/New_York")

	blackoutWindowModel := new(brsmigrationv2.BlackoutWindow)
	blackoutWindowModel.Day = core.StringPtr("sunday")
	blackoutWindowModel.StartTime = timeOfDayModel
	blackoutWindowModel.EndTime = timeOfDayModel

	extendedRetentionScheduleModel := new(brsmigrationv2.ExtendedRetentionSchedule)
	extendedRetentionScheduleModel.Unit = core.StringPtr("runs")
	extendedRetentionScheduleModel.Frequency = core.Int64Ptr(int64(1))

	retentionModel := new(brsmigrationv2.Retention)
	retentionModel.Unit = core.StringPtr("days")
	retentionModel.Duration = core.Int64Ptr(int64(1))

	extendedRetentionPolicyModel := new(brsmigrationv2.ExtendedRetentionPolicy)
	extendedRetentionPolicyModel.Schedule = extendedRetentionScheduleModel
	extendedRetentionPolicyModel.Retention = retentionModel
	extendedRetentionPolicyModel.RunType = core.StringPtr("regular")

	workloadScheduleRetryOptionsModel := new(brsmigrationv2.WorkloadScheduleRetryOptions)
	workloadScheduleRetryOptionsModel.Retries = core.Int64Ptr(int64(0))
	workloadScheduleRetryOptionsModel.RetryIntervalMins = core.Int64Ptr(int64(1))

	workloadScheduleModel := new(brsmigrationv2.WorkloadSchedule)
	workloadScheduleModel.Name = core.StringPtr("daily-incremental-weekly-full")
	workloadScheduleModel.Description = core.StringPtr("Daily incremental with weekly full backup and 30-day retention")
	workloadScheduleModel.BackupPolicy = backupPolicyModel
	workloadScheduleModel.BlackoutWindow = []brsmigrationv2.BlackoutWindow{*blackoutWindowModel}
	workloadScheduleModel.ExtendedRetention = []brsmigrationv2.ExtendedRetentionPolicy{*extendedRetentionPolicyModel}
	workloadScheduleModel.RetryOptions = workloadScheduleRetryOptionsModel

	model := new(brsmigrationv2.Workload)
	model.ID = core.StringPtr("wl-a1b2c3d4-e5f6-7890-abcd-ef1234567890")
	model.Name = core.StringPtr("prod-to-vpc-migration")
	model.VolumeOwnershipMap = map[string]string{"key1": "testString"}
	model.State = core.StringPtr("created")
	model.Payloads = []brsmigrationv2.WorkloadPayloadMapping{*workloadPayloadMappingModel}
	model.Schedule = workloadScheduleModel
	model.SchedulingError = core.StringPtr("BRS protection group creation failed: policy not found")
	model.CreatedAt = CreateMockDateTime("2019-01-01T12:00:00.000Z")
	model.UpdatedAt = CreateMockDateTime("2019-01-01T12:00:00.000Z")
	model.CompletedAt = CreateMockDateTime("2019-01-01T12:00:00.000Z")

	result, err := brsmigration.DataSourceIbmBrsMigrationWorkloadsWorkloadToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBrsMigrationWorkloadsWorkloadPayloadMappingToMap(t *testing.T) {
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

	dataPayloadModel := new(brsmigrationv2.DataPayload)
	dataPayloadModel.VolumeID = core.StringPtr("vol-b1c2d3e4-f5a6-7890-bcde-f01234567890")
	dataPayloadModel.Type = core.StringPtr("ext4")
	dataPayloadModel.Path = core.StringPtr("/mnt/data")

	dataSpecModel := new(brsmigrationv2.DataSpec)
	dataSpecModel.DataFormat = core.StringPtr("raw")
	dataSpecModel.Source = dataPayloadModel
	dataSpecModel.Destination = dataPayloadModel

	model := new(brsmigrationv2.WorkloadPayloadMapping)
	model.ID = core.StringPtr("pl-c3d4e5f6-a7b8-9012-cdef-012345678901")
	model.SourceHostID = core.StringPtr("host-a1b2c3d4-e5f6-7890-abcd-ef1234567890")
	model.DestinationHostID = core.StringPtr("host-b2c3d4e5-f6a7-8901-bcde-f01234567890")
	model.DataSpecs = []brsmigrationv2.DataSpec{*dataSpecModel}

	result, err := brsmigration.DataSourceIbmBrsMigrationWorkloadsWorkloadPayloadMappingToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBrsMigrationWorkloadsDataSpecToMap(t *testing.T) {
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

	dataPayloadModel := new(brsmigrationv2.DataPayload)
	dataPayloadModel.VolumeID = core.StringPtr("vol-b1c2d3e4-f5a6-7890-bcde-f01234567890")
	dataPayloadModel.Type = core.StringPtr("ext4")
	dataPayloadModel.Path = core.StringPtr("/mnt/data")

	model := new(brsmigrationv2.DataSpec)
	model.DataFormat = core.StringPtr("raw")
	model.Source = dataPayloadModel
	model.Destination = dataPayloadModel

	result, err := brsmigration.DataSourceIbmBrsMigrationWorkloadsDataSpecToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBrsMigrationWorkloadsDataPayloadToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["volume_id"] = "vol-b1c2d3e4-f5a6-7890-bcde-f01234567890"
		model["type"] = "ext4"
		model["path"] = "/mnt/data"

		assert.Equal(t, result, model)
	}

	model := new(brsmigrationv2.DataPayload)
	model.VolumeID = core.StringPtr("vol-b1c2d3e4-f5a6-7890-bcde-f01234567890")
	model.Type = core.StringPtr("ext4")
	model.Path = core.StringPtr("/mnt/data")

	result, err := brsmigration.DataSourceIbmBrsMigrationWorkloadsDataPayloadToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBrsMigrationWorkloadsWorkloadScheduleToMap(t *testing.T) {
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

	incrementalScheduleMinuteScheduleModel := new(brsmigrationv2.IncrementalScheduleMinuteSchedule)
	incrementalScheduleMinuteScheduleModel.Frequency = core.Int64Ptr(int64(1))

	incrementalScheduleHourScheduleModel := new(brsmigrationv2.IncrementalScheduleHourSchedule)
	incrementalScheduleHourScheduleModel.Frequency = core.Int64Ptr(int64(1))

	incrementalScheduleDayScheduleModel := new(brsmigrationv2.IncrementalScheduleDaySchedule)
	incrementalScheduleDayScheduleModel.Frequency = core.Int64Ptr(int64(1))

	incrementalScheduleWeekScheduleModel := new(brsmigrationv2.IncrementalScheduleWeekSchedule)
	incrementalScheduleWeekScheduleModel.DayOfWeek = []string{"sunday"}

	incrementalScheduleMonthScheduleModel := new(brsmigrationv2.IncrementalScheduleMonthSchedule)
	incrementalScheduleMonthScheduleModel.DayOfWeek = []string{"sunday"}
	incrementalScheduleMonthScheduleModel.WeekOfMonth = core.StringPtr("first")
	incrementalScheduleMonthScheduleModel.DayOfMonth = core.Int64Ptr(int64(1))

	incrementalScheduleYearScheduleModel := new(brsmigrationv2.IncrementalScheduleYearSchedule)
	incrementalScheduleYearScheduleModel.DayOfYear = core.StringPtr("first")

	incrementalScheduleModel := new(brsmigrationv2.IncrementalSchedule)
	incrementalScheduleModel.Unit = core.StringPtr("minutes")
	incrementalScheduleModel.MinuteSchedule = incrementalScheduleMinuteScheduleModel
	incrementalScheduleModel.HourSchedule = incrementalScheduleHourScheduleModel
	incrementalScheduleModel.DaySchedule = incrementalScheduleDayScheduleModel
	incrementalScheduleModel.WeekSchedule = incrementalScheduleWeekScheduleModel
	incrementalScheduleModel.MonthSchedule = incrementalScheduleMonthScheduleModel
	incrementalScheduleModel.YearSchedule = incrementalScheduleYearScheduleModel

	regularBackupPolicyIncrementalModel := new(brsmigrationv2.RegularBackupPolicyIncremental)
	regularBackupPolicyIncrementalModel.Schedule = incrementalScheduleModel

	fullScheduleDayScheduleModel := new(brsmigrationv2.FullScheduleDaySchedule)
	fullScheduleDayScheduleModel.Frequency = core.Int64Ptr(int64(1))

	fullScheduleWeekScheduleModel := new(brsmigrationv2.FullScheduleWeekSchedule)
	fullScheduleWeekScheduleModel.DayOfWeek = []string{"sunday"}

	fullScheduleMonthScheduleModel := new(brsmigrationv2.FullScheduleMonthSchedule)
	fullScheduleMonthScheduleModel.DayOfWeek = []string{"sunday"}
	fullScheduleMonthScheduleModel.WeekOfMonth = core.StringPtr("first")
	fullScheduleMonthScheduleModel.DayOfMonth = core.Int64Ptr(int64(1))

	fullScheduleYearScheduleModel := new(brsmigrationv2.FullScheduleYearSchedule)
	fullScheduleYearScheduleModel.DayOfYear = core.StringPtr("first")

	fullBackupPolicyScheduleModel := new(brsmigrationv2.FullBackupPolicySchedule)
	fullBackupPolicyScheduleModel.Unit = core.StringPtr("days")
	fullBackupPolicyScheduleModel.DaySchedule = fullScheduleDayScheduleModel
	fullBackupPolicyScheduleModel.WeekSchedule = fullScheduleWeekScheduleModel
	fullBackupPolicyScheduleModel.MonthSchedule = fullScheduleMonthScheduleModel
	fullBackupPolicyScheduleModel.YearSchedule = fullScheduleYearScheduleModel

	regularBackupPolicyFullModel := new(brsmigrationv2.RegularBackupPolicyFull)
	regularBackupPolicyFullModel.Schedule = fullBackupPolicyScheduleModel

	regularBackupPolicyRetentionModel := new(brsmigrationv2.RegularBackupPolicyRetention)
	regularBackupPolicyRetentionModel.Unit = core.StringPtr("days")
	regularBackupPolicyRetentionModel.Duration = core.Int64Ptr(int64(1))

	regularBackupPolicyModel := new(brsmigrationv2.RegularBackupPolicy)
	regularBackupPolicyModel.Incremental = regularBackupPolicyIncrementalModel
	regularBackupPolicyModel.Full = regularBackupPolicyFullModel
	regularBackupPolicyModel.Retention = regularBackupPolicyRetentionModel

	logScheduleMinuteScheduleModel := new(brsmigrationv2.LogScheduleMinuteSchedule)
	logScheduleMinuteScheduleModel.Frequency = core.Int64Ptr(int64(1))

	logScheduleHourScheduleModel := new(brsmigrationv2.LogScheduleHourSchedule)
	logScheduleHourScheduleModel.Frequency = core.Int64Ptr(int64(1))

	logBackupPolicyScheduleModel := new(brsmigrationv2.LogBackupPolicySchedule)
	logBackupPolicyScheduleModel.Unit = core.StringPtr("minutes")
	logBackupPolicyScheduleModel.MinuteSchedule = logScheduleMinuteScheduleModel
	logBackupPolicyScheduleModel.HourSchedule = logScheduleHourScheduleModel

	backupPolicyLogModel := new(brsmigrationv2.BackupPolicyLog)
	backupPolicyLogModel.Schedule = logBackupPolicyScheduleModel

	backupPolicyModel := new(brsmigrationv2.BackupPolicy)
	backupPolicyModel.Regular = regularBackupPolicyModel
	backupPolicyModel.Log = backupPolicyLogModel

	timeOfDayModel := new(brsmigrationv2.TimeOfDay)
	timeOfDayModel.Hour = core.Int64Ptr(int64(0))
	timeOfDayModel.Minute = core.Int64Ptr(int64(0))
	timeOfDayModel.Timezone = core.StringPtr("America/New_York")

	blackoutWindowModel := new(brsmigrationv2.BlackoutWindow)
	blackoutWindowModel.Day = core.StringPtr("sunday")
	blackoutWindowModel.StartTime = timeOfDayModel
	blackoutWindowModel.EndTime = timeOfDayModel

	extendedRetentionScheduleModel := new(brsmigrationv2.ExtendedRetentionSchedule)
	extendedRetentionScheduleModel.Unit = core.StringPtr("runs")
	extendedRetentionScheduleModel.Frequency = core.Int64Ptr(int64(1))

	retentionModel := new(brsmigrationv2.Retention)
	retentionModel.Unit = core.StringPtr("days")
	retentionModel.Duration = core.Int64Ptr(int64(1))

	extendedRetentionPolicyModel := new(brsmigrationv2.ExtendedRetentionPolicy)
	extendedRetentionPolicyModel.Schedule = extendedRetentionScheduleModel
	extendedRetentionPolicyModel.Retention = retentionModel
	extendedRetentionPolicyModel.RunType = core.StringPtr("regular")

	workloadScheduleRetryOptionsModel := new(brsmigrationv2.WorkloadScheduleRetryOptions)
	workloadScheduleRetryOptionsModel.Retries = core.Int64Ptr(int64(0))
	workloadScheduleRetryOptionsModel.RetryIntervalMins = core.Int64Ptr(int64(1))

	model := new(brsmigrationv2.WorkloadSchedule)
	model.Name = core.StringPtr("daily-incremental-weekly-full")
	model.Description = core.StringPtr("Daily incremental with weekly full backup and 30-day retention")
	model.BackupPolicy = backupPolicyModel
	model.BlackoutWindow = []brsmigrationv2.BlackoutWindow{*blackoutWindowModel}
	model.ExtendedRetention = []brsmigrationv2.ExtendedRetentionPolicy{*extendedRetentionPolicyModel}
	model.RetryOptions = workloadScheduleRetryOptionsModel

	result, err := brsmigration.DataSourceIbmBrsMigrationWorkloadsWorkloadScheduleToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBrsMigrationWorkloadsBackupPolicyToMap(t *testing.T) {
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

	incrementalScheduleMinuteScheduleModel := new(brsmigrationv2.IncrementalScheduleMinuteSchedule)
	incrementalScheduleMinuteScheduleModel.Frequency = core.Int64Ptr(int64(1))

	incrementalScheduleHourScheduleModel := new(brsmigrationv2.IncrementalScheduleHourSchedule)
	incrementalScheduleHourScheduleModel.Frequency = core.Int64Ptr(int64(1))

	incrementalScheduleDayScheduleModel := new(brsmigrationv2.IncrementalScheduleDaySchedule)
	incrementalScheduleDayScheduleModel.Frequency = core.Int64Ptr(int64(1))

	incrementalScheduleWeekScheduleModel := new(brsmigrationv2.IncrementalScheduleWeekSchedule)
	incrementalScheduleWeekScheduleModel.DayOfWeek = []string{"sunday"}

	incrementalScheduleMonthScheduleModel := new(brsmigrationv2.IncrementalScheduleMonthSchedule)
	incrementalScheduleMonthScheduleModel.DayOfWeek = []string{"sunday"}
	incrementalScheduleMonthScheduleModel.WeekOfMonth = core.StringPtr("first")
	incrementalScheduleMonthScheduleModel.DayOfMonth = core.Int64Ptr(int64(1))

	incrementalScheduleYearScheduleModel := new(brsmigrationv2.IncrementalScheduleYearSchedule)
	incrementalScheduleYearScheduleModel.DayOfYear = core.StringPtr("first")

	incrementalScheduleModel := new(brsmigrationv2.IncrementalSchedule)
	incrementalScheduleModel.Unit = core.StringPtr("minutes")
	incrementalScheduleModel.MinuteSchedule = incrementalScheduleMinuteScheduleModel
	incrementalScheduleModel.HourSchedule = incrementalScheduleHourScheduleModel
	incrementalScheduleModel.DaySchedule = incrementalScheduleDayScheduleModel
	incrementalScheduleModel.WeekSchedule = incrementalScheduleWeekScheduleModel
	incrementalScheduleModel.MonthSchedule = incrementalScheduleMonthScheduleModel
	incrementalScheduleModel.YearSchedule = incrementalScheduleYearScheduleModel

	regularBackupPolicyIncrementalModel := new(brsmigrationv2.RegularBackupPolicyIncremental)
	regularBackupPolicyIncrementalModel.Schedule = incrementalScheduleModel

	fullScheduleDayScheduleModel := new(brsmigrationv2.FullScheduleDaySchedule)
	fullScheduleDayScheduleModel.Frequency = core.Int64Ptr(int64(1))

	fullScheduleWeekScheduleModel := new(brsmigrationv2.FullScheduleWeekSchedule)
	fullScheduleWeekScheduleModel.DayOfWeek = []string{"sunday"}

	fullScheduleMonthScheduleModel := new(brsmigrationv2.FullScheduleMonthSchedule)
	fullScheduleMonthScheduleModel.DayOfWeek = []string{"sunday"}
	fullScheduleMonthScheduleModel.WeekOfMonth = core.StringPtr("first")
	fullScheduleMonthScheduleModel.DayOfMonth = core.Int64Ptr(int64(1))

	fullScheduleYearScheduleModel := new(brsmigrationv2.FullScheduleYearSchedule)
	fullScheduleYearScheduleModel.DayOfYear = core.StringPtr("first")

	fullBackupPolicyScheduleModel := new(brsmigrationv2.FullBackupPolicySchedule)
	fullBackupPolicyScheduleModel.Unit = core.StringPtr("days")
	fullBackupPolicyScheduleModel.DaySchedule = fullScheduleDayScheduleModel
	fullBackupPolicyScheduleModel.WeekSchedule = fullScheduleWeekScheduleModel
	fullBackupPolicyScheduleModel.MonthSchedule = fullScheduleMonthScheduleModel
	fullBackupPolicyScheduleModel.YearSchedule = fullScheduleYearScheduleModel

	regularBackupPolicyFullModel := new(brsmigrationv2.RegularBackupPolicyFull)
	regularBackupPolicyFullModel.Schedule = fullBackupPolicyScheduleModel

	regularBackupPolicyRetentionModel := new(brsmigrationv2.RegularBackupPolicyRetention)
	regularBackupPolicyRetentionModel.Unit = core.StringPtr("days")
	regularBackupPolicyRetentionModel.Duration = core.Int64Ptr(int64(1))

	regularBackupPolicyModel := new(brsmigrationv2.RegularBackupPolicy)
	regularBackupPolicyModel.Incremental = regularBackupPolicyIncrementalModel
	regularBackupPolicyModel.Full = regularBackupPolicyFullModel
	regularBackupPolicyModel.Retention = regularBackupPolicyRetentionModel

	logScheduleMinuteScheduleModel := new(brsmigrationv2.LogScheduleMinuteSchedule)
	logScheduleMinuteScheduleModel.Frequency = core.Int64Ptr(int64(1))

	logScheduleHourScheduleModel := new(brsmigrationv2.LogScheduleHourSchedule)
	logScheduleHourScheduleModel.Frequency = core.Int64Ptr(int64(1))

	logBackupPolicyScheduleModel := new(brsmigrationv2.LogBackupPolicySchedule)
	logBackupPolicyScheduleModel.Unit = core.StringPtr("minutes")
	logBackupPolicyScheduleModel.MinuteSchedule = logScheduleMinuteScheduleModel
	logBackupPolicyScheduleModel.HourSchedule = logScheduleHourScheduleModel

	backupPolicyLogModel := new(brsmigrationv2.BackupPolicyLog)
	backupPolicyLogModel.Schedule = logBackupPolicyScheduleModel

	model := new(brsmigrationv2.BackupPolicy)
	model.Regular = regularBackupPolicyModel
	model.Log = backupPolicyLogModel

	result, err := brsmigration.DataSourceIbmBrsMigrationWorkloadsBackupPolicyToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBrsMigrationWorkloadsRegularBackupPolicyToMap(t *testing.T) {
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

	incrementalScheduleMinuteScheduleModel := new(brsmigrationv2.IncrementalScheduleMinuteSchedule)
	incrementalScheduleMinuteScheduleModel.Frequency = core.Int64Ptr(int64(1))

	incrementalScheduleHourScheduleModel := new(brsmigrationv2.IncrementalScheduleHourSchedule)
	incrementalScheduleHourScheduleModel.Frequency = core.Int64Ptr(int64(1))

	incrementalScheduleDayScheduleModel := new(brsmigrationv2.IncrementalScheduleDaySchedule)
	incrementalScheduleDayScheduleModel.Frequency = core.Int64Ptr(int64(1))

	incrementalScheduleWeekScheduleModel := new(brsmigrationv2.IncrementalScheduleWeekSchedule)
	incrementalScheduleWeekScheduleModel.DayOfWeek = []string{"sunday"}

	incrementalScheduleMonthScheduleModel := new(brsmigrationv2.IncrementalScheduleMonthSchedule)
	incrementalScheduleMonthScheduleModel.DayOfWeek = []string{"sunday"}
	incrementalScheduleMonthScheduleModel.WeekOfMonth = core.StringPtr("first")
	incrementalScheduleMonthScheduleModel.DayOfMonth = core.Int64Ptr(int64(1))

	incrementalScheduleYearScheduleModel := new(brsmigrationv2.IncrementalScheduleYearSchedule)
	incrementalScheduleYearScheduleModel.DayOfYear = core.StringPtr("first")

	incrementalScheduleModel := new(brsmigrationv2.IncrementalSchedule)
	incrementalScheduleModel.Unit = core.StringPtr("minutes")
	incrementalScheduleModel.MinuteSchedule = incrementalScheduleMinuteScheduleModel
	incrementalScheduleModel.HourSchedule = incrementalScheduleHourScheduleModel
	incrementalScheduleModel.DaySchedule = incrementalScheduleDayScheduleModel
	incrementalScheduleModel.WeekSchedule = incrementalScheduleWeekScheduleModel
	incrementalScheduleModel.MonthSchedule = incrementalScheduleMonthScheduleModel
	incrementalScheduleModel.YearSchedule = incrementalScheduleYearScheduleModel

	regularBackupPolicyIncrementalModel := new(brsmigrationv2.RegularBackupPolicyIncremental)
	regularBackupPolicyIncrementalModel.Schedule = incrementalScheduleModel

	fullScheduleDayScheduleModel := new(brsmigrationv2.FullScheduleDaySchedule)
	fullScheduleDayScheduleModel.Frequency = core.Int64Ptr(int64(1))

	fullScheduleWeekScheduleModel := new(brsmigrationv2.FullScheduleWeekSchedule)
	fullScheduleWeekScheduleModel.DayOfWeek = []string{"sunday"}

	fullScheduleMonthScheduleModel := new(brsmigrationv2.FullScheduleMonthSchedule)
	fullScheduleMonthScheduleModel.DayOfWeek = []string{"sunday"}
	fullScheduleMonthScheduleModel.WeekOfMonth = core.StringPtr("first")
	fullScheduleMonthScheduleModel.DayOfMonth = core.Int64Ptr(int64(1))

	fullScheduleYearScheduleModel := new(brsmigrationv2.FullScheduleYearSchedule)
	fullScheduleYearScheduleModel.DayOfYear = core.StringPtr("first")

	fullBackupPolicyScheduleModel := new(brsmigrationv2.FullBackupPolicySchedule)
	fullBackupPolicyScheduleModel.Unit = core.StringPtr("days")
	fullBackupPolicyScheduleModel.DaySchedule = fullScheduleDayScheduleModel
	fullBackupPolicyScheduleModel.WeekSchedule = fullScheduleWeekScheduleModel
	fullBackupPolicyScheduleModel.MonthSchedule = fullScheduleMonthScheduleModel
	fullBackupPolicyScheduleModel.YearSchedule = fullScheduleYearScheduleModel

	regularBackupPolicyFullModel := new(brsmigrationv2.RegularBackupPolicyFull)
	regularBackupPolicyFullModel.Schedule = fullBackupPolicyScheduleModel

	regularBackupPolicyRetentionModel := new(brsmigrationv2.RegularBackupPolicyRetention)
	regularBackupPolicyRetentionModel.Unit = core.StringPtr("days")
	regularBackupPolicyRetentionModel.Duration = core.Int64Ptr(int64(1))

	model := new(brsmigrationv2.RegularBackupPolicy)
	model.Incremental = regularBackupPolicyIncrementalModel
	model.Full = regularBackupPolicyFullModel
	model.Retention = regularBackupPolicyRetentionModel

	result, err := brsmigration.DataSourceIbmBrsMigrationWorkloadsRegularBackupPolicyToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBrsMigrationWorkloadsRegularBackupPolicyIncrementalToMap(t *testing.T) {
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

	incrementalScheduleMinuteScheduleModel := new(brsmigrationv2.IncrementalScheduleMinuteSchedule)
	incrementalScheduleMinuteScheduleModel.Frequency = core.Int64Ptr(int64(1))

	incrementalScheduleHourScheduleModel := new(brsmigrationv2.IncrementalScheduleHourSchedule)
	incrementalScheduleHourScheduleModel.Frequency = core.Int64Ptr(int64(1))

	incrementalScheduleDayScheduleModel := new(brsmigrationv2.IncrementalScheduleDaySchedule)
	incrementalScheduleDayScheduleModel.Frequency = core.Int64Ptr(int64(1))

	incrementalScheduleWeekScheduleModel := new(brsmigrationv2.IncrementalScheduleWeekSchedule)
	incrementalScheduleWeekScheduleModel.DayOfWeek = []string{"sunday"}

	incrementalScheduleMonthScheduleModel := new(brsmigrationv2.IncrementalScheduleMonthSchedule)
	incrementalScheduleMonthScheduleModel.DayOfWeek = []string{"sunday"}
	incrementalScheduleMonthScheduleModel.WeekOfMonth = core.StringPtr("first")
	incrementalScheduleMonthScheduleModel.DayOfMonth = core.Int64Ptr(int64(1))

	incrementalScheduleYearScheduleModel := new(brsmigrationv2.IncrementalScheduleYearSchedule)
	incrementalScheduleYearScheduleModel.DayOfYear = core.StringPtr("first")

	incrementalScheduleModel := new(brsmigrationv2.IncrementalSchedule)
	incrementalScheduleModel.Unit = core.StringPtr("minutes")
	incrementalScheduleModel.MinuteSchedule = incrementalScheduleMinuteScheduleModel
	incrementalScheduleModel.HourSchedule = incrementalScheduleHourScheduleModel
	incrementalScheduleModel.DaySchedule = incrementalScheduleDayScheduleModel
	incrementalScheduleModel.WeekSchedule = incrementalScheduleWeekScheduleModel
	incrementalScheduleModel.MonthSchedule = incrementalScheduleMonthScheduleModel
	incrementalScheduleModel.YearSchedule = incrementalScheduleYearScheduleModel

	model := new(brsmigrationv2.RegularBackupPolicyIncremental)
	model.Schedule = incrementalScheduleModel

	result, err := brsmigration.DataSourceIbmBrsMigrationWorkloadsRegularBackupPolicyIncrementalToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBrsMigrationWorkloadsIncrementalScheduleToMap(t *testing.T) {
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

	incrementalScheduleMinuteScheduleModel := new(brsmigrationv2.IncrementalScheduleMinuteSchedule)
	incrementalScheduleMinuteScheduleModel.Frequency = core.Int64Ptr(int64(1))

	incrementalScheduleHourScheduleModel := new(brsmigrationv2.IncrementalScheduleHourSchedule)
	incrementalScheduleHourScheduleModel.Frequency = core.Int64Ptr(int64(1))

	incrementalScheduleDayScheduleModel := new(brsmigrationv2.IncrementalScheduleDaySchedule)
	incrementalScheduleDayScheduleModel.Frequency = core.Int64Ptr(int64(1))

	incrementalScheduleWeekScheduleModel := new(brsmigrationv2.IncrementalScheduleWeekSchedule)
	incrementalScheduleWeekScheduleModel.DayOfWeek = []string{"sunday"}

	incrementalScheduleMonthScheduleModel := new(brsmigrationv2.IncrementalScheduleMonthSchedule)
	incrementalScheduleMonthScheduleModel.DayOfWeek = []string{"sunday"}
	incrementalScheduleMonthScheduleModel.WeekOfMonth = core.StringPtr("first")
	incrementalScheduleMonthScheduleModel.DayOfMonth = core.Int64Ptr(int64(1))

	incrementalScheduleYearScheduleModel := new(brsmigrationv2.IncrementalScheduleYearSchedule)
	incrementalScheduleYearScheduleModel.DayOfYear = core.StringPtr("first")

	model := new(brsmigrationv2.IncrementalSchedule)
	model.Unit = core.StringPtr("minutes")
	model.MinuteSchedule = incrementalScheduleMinuteScheduleModel
	model.HourSchedule = incrementalScheduleHourScheduleModel
	model.DaySchedule = incrementalScheduleDayScheduleModel
	model.WeekSchedule = incrementalScheduleWeekScheduleModel
	model.MonthSchedule = incrementalScheduleMonthScheduleModel
	model.YearSchedule = incrementalScheduleYearScheduleModel

	result, err := brsmigration.DataSourceIbmBrsMigrationWorkloadsIncrementalScheduleToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBrsMigrationWorkloadsIncrementalScheduleMinuteScheduleToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["frequency"] = int(1)

		assert.Equal(t, result, model)
	}

	model := new(brsmigrationv2.IncrementalScheduleMinuteSchedule)
	model.Frequency = core.Int64Ptr(int64(1))

	result, err := brsmigration.DataSourceIbmBrsMigrationWorkloadsIncrementalScheduleMinuteScheduleToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBrsMigrationWorkloadsIncrementalScheduleHourScheduleToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["frequency"] = int(1)

		assert.Equal(t, result, model)
	}

	model := new(brsmigrationv2.IncrementalScheduleHourSchedule)
	model.Frequency = core.Int64Ptr(int64(1))

	result, err := brsmigration.DataSourceIbmBrsMigrationWorkloadsIncrementalScheduleHourScheduleToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBrsMigrationWorkloadsIncrementalScheduleDayScheduleToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["frequency"] = int(1)

		assert.Equal(t, result, model)
	}

	model := new(brsmigrationv2.IncrementalScheduleDaySchedule)
	model.Frequency = core.Int64Ptr(int64(1))

	result, err := brsmigration.DataSourceIbmBrsMigrationWorkloadsIncrementalScheduleDayScheduleToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBrsMigrationWorkloadsIncrementalScheduleWeekScheduleToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["day_of_week"] = []string{"sunday"}

		assert.Equal(t, result, model)
	}

	model := new(brsmigrationv2.IncrementalScheduleWeekSchedule)
	model.DayOfWeek = []string{"sunday"}

	result, err := brsmigration.DataSourceIbmBrsMigrationWorkloadsIncrementalScheduleWeekScheduleToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBrsMigrationWorkloadsIncrementalScheduleMonthScheduleToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["day_of_week"] = []string{"sunday"}
		model["week_of_month"] = "first"
		model["day_of_month"] = int(1)

		assert.Equal(t, result, model)
	}

	model := new(brsmigrationv2.IncrementalScheduleMonthSchedule)
	model.DayOfWeek = []string{"sunday"}
	model.WeekOfMonth = core.StringPtr("first")
	model.DayOfMonth = core.Int64Ptr(int64(1))

	result, err := brsmigration.DataSourceIbmBrsMigrationWorkloadsIncrementalScheduleMonthScheduleToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBrsMigrationWorkloadsIncrementalScheduleYearScheduleToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["day_of_year"] = "first"

		assert.Equal(t, result, model)
	}

	model := new(brsmigrationv2.IncrementalScheduleYearSchedule)
	model.DayOfYear = core.StringPtr("first")

	result, err := brsmigration.DataSourceIbmBrsMigrationWorkloadsIncrementalScheduleYearScheduleToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBrsMigrationWorkloadsRegularBackupPolicyFullToMap(t *testing.T) {
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

	fullScheduleDayScheduleModel := new(brsmigrationv2.FullScheduleDaySchedule)
	fullScheduleDayScheduleModel.Frequency = core.Int64Ptr(int64(1))

	fullScheduleWeekScheduleModel := new(brsmigrationv2.FullScheduleWeekSchedule)
	fullScheduleWeekScheduleModel.DayOfWeek = []string{"sunday"}

	fullScheduleMonthScheduleModel := new(brsmigrationv2.FullScheduleMonthSchedule)
	fullScheduleMonthScheduleModel.DayOfWeek = []string{"sunday"}
	fullScheduleMonthScheduleModel.WeekOfMonth = core.StringPtr("first")
	fullScheduleMonthScheduleModel.DayOfMonth = core.Int64Ptr(int64(1))

	fullScheduleYearScheduleModel := new(brsmigrationv2.FullScheduleYearSchedule)
	fullScheduleYearScheduleModel.DayOfYear = core.StringPtr("first")

	fullBackupPolicyScheduleModel := new(brsmigrationv2.FullBackupPolicySchedule)
	fullBackupPolicyScheduleModel.Unit = core.StringPtr("days")
	fullBackupPolicyScheduleModel.DaySchedule = fullScheduleDayScheduleModel
	fullBackupPolicyScheduleModel.WeekSchedule = fullScheduleWeekScheduleModel
	fullBackupPolicyScheduleModel.MonthSchedule = fullScheduleMonthScheduleModel
	fullBackupPolicyScheduleModel.YearSchedule = fullScheduleYearScheduleModel

	model := new(brsmigrationv2.RegularBackupPolicyFull)
	model.Schedule = fullBackupPolicyScheduleModel

	result, err := brsmigration.DataSourceIbmBrsMigrationWorkloadsRegularBackupPolicyFullToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBrsMigrationWorkloadsFullBackupPolicyScheduleToMap(t *testing.T) {
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

	fullScheduleDayScheduleModel := new(brsmigrationv2.FullScheduleDaySchedule)
	fullScheduleDayScheduleModel.Frequency = core.Int64Ptr(int64(1))

	fullScheduleWeekScheduleModel := new(brsmigrationv2.FullScheduleWeekSchedule)
	fullScheduleWeekScheduleModel.DayOfWeek = []string{"sunday"}

	fullScheduleMonthScheduleModel := new(brsmigrationv2.FullScheduleMonthSchedule)
	fullScheduleMonthScheduleModel.DayOfWeek = []string{"sunday"}
	fullScheduleMonthScheduleModel.WeekOfMonth = core.StringPtr("first")
	fullScheduleMonthScheduleModel.DayOfMonth = core.Int64Ptr(int64(1))

	fullScheduleYearScheduleModel := new(brsmigrationv2.FullScheduleYearSchedule)
	fullScheduleYearScheduleModel.DayOfYear = core.StringPtr("first")

	model := new(brsmigrationv2.FullBackupPolicySchedule)
	model.Unit = core.StringPtr("days")
	model.DaySchedule = fullScheduleDayScheduleModel
	model.WeekSchedule = fullScheduleWeekScheduleModel
	model.MonthSchedule = fullScheduleMonthScheduleModel
	model.YearSchedule = fullScheduleYearScheduleModel

	result, err := brsmigration.DataSourceIbmBrsMigrationWorkloadsFullBackupPolicyScheduleToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBrsMigrationWorkloadsFullScheduleDayScheduleToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["frequency"] = int(1)

		assert.Equal(t, result, model)
	}

	model := new(brsmigrationv2.FullScheduleDaySchedule)
	model.Frequency = core.Int64Ptr(int64(1))

	result, err := brsmigration.DataSourceIbmBrsMigrationWorkloadsFullScheduleDayScheduleToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBrsMigrationWorkloadsFullScheduleWeekScheduleToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["day_of_week"] = []string{"sunday"}

		assert.Equal(t, result, model)
	}

	model := new(brsmigrationv2.FullScheduleWeekSchedule)
	model.DayOfWeek = []string{"sunday"}

	result, err := brsmigration.DataSourceIbmBrsMigrationWorkloadsFullScheduleWeekScheduleToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBrsMigrationWorkloadsFullScheduleMonthScheduleToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["day_of_week"] = []string{"sunday"}
		model["week_of_month"] = "first"
		model["day_of_month"] = int(1)

		assert.Equal(t, result, model)
	}

	model := new(brsmigrationv2.FullScheduleMonthSchedule)
	model.DayOfWeek = []string{"sunday"}
	model.WeekOfMonth = core.StringPtr("first")
	model.DayOfMonth = core.Int64Ptr(int64(1))

	result, err := brsmigration.DataSourceIbmBrsMigrationWorkloadsFullScheduleMonthScheduleToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBrsMigrationWorkloadsFullScheduleYearScheduleToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["day_of_year"] = "first"

		assert.Equal(t, result, model)
	}

	model := new(brsmigrationv2.FullScheduleYearSchedule)
	model.DayOfYear = core.StringPtr("first")

	result, err := brsmigration.DataSourceIbmBrsMigrationWorkloadsFullScheduleYearScheduleToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBrsMigrationWorkloadsRegularBackupPolicyRetentionToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["unit"] = "days"
		model["duration"] = int(1)

		assert.Equal(t, result, model)
	}

	model := new(brsmigrationv2.RegularBackupPolicyRetention)
	model.Unit = core.StringPtr("days")
	model.Duration = core.Int64Ptr(int64(1))

	result, err := brsmigration.DataSourceIbmBrsMigrationWorkloadsRegularBackupPolicyRetentionToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBrsMigrationWorkloadsBackupPolicyLogToMap(t *testing.T) {
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

	logScheduleMinuteScheduleModel := new(brsmigrationv2.LogScheduleMinuteSchedule)
	logScheduleMinuteScheduleModel.Frequency = core.Int64Ptr(int64(1))

	logScheduleHourScheduleModel := new(brsmigrationv2.LogScheduleHourSchedule)
	logScheduleHourScheduleModel.Frequency = core.Int64Ptr(int64(1))

	logBackupPolicyScheduleModel := new(brsmigrationv2.LogBackupPolicySchedule)
	logBackupPolicyScheduleModel.Unit = core.StringPtr("minutes")
	logBackupPolicyScheduleModel.MinuteSchedule = logScheduleMinuteScheduleModel
	logBackupPolicyScheduleModel.HourSchedule = logScheduleHourScheduleModel

	model := new(brsmigrationv2.BackupPolicyLog)
	model.Schedule = logBackupPolicyScheduleModel

	result, err := brsmigration.DataSourceIbmBrsMigrationWorkloadsBackupPolicyLogToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBrsMigrationWorkloadsLogBackupPolicyScheduleToMap(t *testing.T) {
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

	logScheduleMinuteScheduleModel := new(brsmigrationv2.LogScheduleMinuteSchedule)
	logScheduleMinuteScheduleModel.Frequency = core.Int64Ptr(int64(1))

	logScheduleHourScheduleModel := new(brsmigrationv2.LogScheduleHourSchedule)
	logScheduleHourScheduleModel.Frequency = core.Int64Ptr(int64(1))

	model := new(brsmigrationv2.LogBackupPolicySchedule)
	model.Unit = core.StringPtr("minutes")
	model.MinuteSchedule = logScheduleMinuteScheduleModel
	model.HourSchedule = logScheduleHourScheduleModel

	result, err := brsmigration.DataSourceIbmBrsMigrationWorkloadsLogBackupPolicyScheduleToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBrsMigrationWorkloadsLogScheduleMinuteScheduleToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["frequency"] = int(1)

		assert.Equal(t, result, model)
	}

	model := new(brsmigrationv2.LogScheduleMinuteSchedule)
	model.Frequency = core.Int64Ptr(int64(1))

	result, err := brsmigration.DataSourceIbmBrsMigrationWorkloadsLogScheduleMinuteScheduleToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBrsMigrationWorkloadsLogScheduleHourScheduleToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["frequency"] = int(1)

		assert.Equal(t, result, model)
	}

	model := new(brsmigrationv2.LogScheduleHourSchedule)
	model.Frequency = core.Int64Ptr(int64(1))

	result, err := brsmigration.DataSourceIbmBrsMigrationWorkloadsLogScheduleHourScheduleToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBrsMigrationWorkloadsBlackoutWindowToMap(t *testing.T) {
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

	timeOfDayModel := new(brsmigrationv2.TimeOfDay)
	timeOfDayModel.Hour = core.Int64Ptr(int64(0))
	timeOfDayModel.Minute = core.Int64Ptr(int64(0))
	timeOfDayModel.Timezone = core.StringPtr("America/New_York")

	model := new(brsmigrationv2.BlackoutWindow)
	model.Day = core.StringPtr("sunday")
	model.StartTime = timeOfDayModel
	model.EndTime = timeOfDayModel

	result, err := brsmigration.DataSourceIbmBrsMigrationWorkloadsBlackoutWindowToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBrsMigrationWorkloadsTimeOfDayToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["hour"] = int(0)
		model["minute"] = int(0)
		model["timezone"] = "America/New_York"

		assert.Equal(t, result, model)
	}

	model := new(brsmigrationv2.TimeOfDay)
	model.Hour = core.Int64Ptr(int64(0))
	model.Minute = core.Int64Ptr(int64(0))
	model.Timezone = core.StringPtr("America/New_York")

	result, err := brsmigration.DataSourceIbmBrsMigrationWorkloadsTimeOfDayToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBrsMigrationWorkloadsExtendedRetentionPolicyToMap(t *testing.T) {
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

	extendedRetentionScheduleModel := new(brsmigrationv2.ExtendedRetentionSchedule)
	extendedRetentionScheduleModel.Unit = core.StringPtr("runs")
	extendedRetentionScheduleModel.Frequency = core.Int64Ptr(int64(1))

	retentionModel := new(brsmigrationv2.Retention)
	retentionModel.Unit = core.StringPtr("days")
	retentionModel.Duration = core.Int64Ptr(int64(1))

	model := new(brsmigrationv2.ExtendedRetentionPolicy)
	model.Schedule = extendedRetentionScheduleModel
	model.Retention = retentionModel
	model.RunType = core.StringPtr("regular")

	result, err := brsmigration.DataSourceIbmBrsMigrationWorkloadsExtendedRetentionPolicyToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBrsMigrationWorkloadsExtendedRetentionScheduleToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["unit"] = "runs"
		model["frequency"] = int(1)

		assert.Equal(t, result, model)
	}

	model := new(brsmigrationv2.ExtendedRetentionSchedule)
	model.Unit = core.StringPtr("runs")
	model.Frequency = core.Int64Ptr(int64(1))

	result, err := brsmigration.DataSourceIbmBrsMigrationWorkloadsExtendedRetentionScheduleToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBrsMigrationWorkloadsRetentionToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["unit"] = "days"
		model["duration"] = int(1)

		assert.Equal(t, result, model)
	}

	model := new(brsmigrationv2.Retention)
	model.Unit = core.StringPtr("days")
	model.Duration = core.Int64Ptr(int64(1))

	result, err := brsmigration.DataSourceIbmBrsMigrationWorkloadsRetentionToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBrsMigrationWorkloadsWorkloadScheduleRetryOptionsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["retries"] = int(0)
		model["retry_interval_mins"] = int(1)

		assert.Equal(t, result, model)
	}

	model := new(brsmigrationv2.WorkloadScheduleRetryOptions)
	model.Retries = core.Int64Ptr(int64(0))
	model.RetryIntervalMins = core.Int64Ptr(int64(1))

	result, err := brsmigration.DataSourceIbmBrsMigrationWorkloadsWorkloadScheduleRetryOptionsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
