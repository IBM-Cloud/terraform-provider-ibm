// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/stretchr/testify/assert"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	vpc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/vpc"
	. "github.com/IBM-Cloud/terraform-provider-ibm/ibm/unittest"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
)

func TestAccIBMIsVolumeJobsDataSourceBasic(t *testing.T) {
	volumeJobVolumeID := fmt.Sprintf("tf_volume_id_%d", acctest.RandIntRange(10, 100))
	volumeJobJobType := "migrate"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsVolumeJobsDataSourceConfigBasic(volumeJobVolumeID, volumeJobJobType),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_jobs.is_volume_jobs_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_jobs.is_volume_jobs_instance", "volume_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_jobs.is_volume_jobs_instance", "jobs.#"),
					resource.TestCheckResourceAttr("data.ibm_is_volume_jobs.is_volume_jobs_instance", "jobs.0.job_type", volumeJobJobType),
				),
			},
		},
	})
}

func TestAccIBMIsVolumeJobsDataSourceAllArgs(t *testing.T) {
	volumeJobVolumeID := fmt.Sprintf("tf_volume_id_%d", acctest.RandIntRange(10, 100))
	volumeJobStart := fmt.Sprintf("tf_start_%d", acctest.RandIntRange(10, 100))
	volumeJobLimit := fmt.Sprintf("%d", acctest.RandIntRange(1, 100))
	volumeJobJobType := "migrate"
	volumeJobName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsVolumeJobsDataSourceConfig(volumeJobVolumeID, volumeJobStart, volumeJobLimit, volumeJobJobType, volumeJobName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_jobs.is_volume_jobs_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_jobs.is_volume_jobs_instance", "volume_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_jobs.is_volume_jobs_instance", "jobs.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_jobs.is_volume_jobs_instance", "jobs.0.auto_delete"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_jobs.is_volume_jobs_instance", "jobs.0.completed_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_jobs.is_volume_jobs_instance", "jobs.0.created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_jobs.is_volume_jobs_instance", "jobs.0.estimated_completion_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_jobs.is_volume_jobs_instance", "jobs.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_jobs.is_volume_jobs_instance", "jobs.0.id"),
					resource.TestCheckResourceAttr("data.ibm_is_volume_jobs.is_volume_jobs_instance", "jobs.0.job_type", volumeJobJobType),
					resource.TestCheckResourceAttr("data.ibm_is_volume_jobs.is_volume_jobs_instance", "jobs.0.name", volumeJobName),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_jobs.is_volume_jobs_instance", "jobs.0.resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_jobs.is_volume_jobs_instance", "jobs.0.started_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_jobs.is_volume_jobs_instance", "jobs.0.status"),
				),
			},
		},
	})
}

func testAccCheckIBMIsVolumeJobsDataSourceConfigBasic(volumeJobVolumeID string, volumeJobJobType string) string {
	return fmt.Sprintf(`
		resource "ibm_is_volume_job" "is_volume_job_instance" {
			volume_id = "%s"
			job_type = "%s"
		}

		data "ibm_is_volume_jobs" "is_volume_jobs_instance" {
			volume_id = ibm_is_volume_job.is_volume_job_instance.volume_id
		}
	`, volumeJobVolumeID, volumeJobJobType)
}

func testAccCheckIBMIsVolumeJobsDataSourceConfig(volumeJobVolumeID string, volumeJobStart string, volumeJobLimit string, volumeJobJobType string, volumeJobName string) string {
	return fmt.Sprintf(`
		resource "ibm_is_volume_job" "is_volume_job_instance" {
			volume_id = "%s"
			start = "%s"
			limit = %s
			job_type = "%s"
			name = "%s"
			parameters {
				bandwidth = 1000
				iops = 10000
				profile {
					name = "general-purpose"
				}
			}
		}

		data "ibm_is_volume_jobs" "is_volume_jobs_instance" {
			volume_id = ibm_is_volume_job.is_volume_job_instance.volume_id
		}
	`, volumeJobVolumeID, volumeJobStart, volumeJobLimit, volumeJobJobType, volumeJobName)
}

func TestDataSourceIBMIsVolumeJobsVolumeJobToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		volumeJobStatusReasonModel := make(map[string]interface{})
		volumeJobStatusReasonModel["code"] = "volume_detached_from_virtual_instance"
		volumeJobStatusReasonModel["message"] = "testString"
		volumeJobStatusReasonModel["more_info"] = "https://cloud.ibm.com/docs/vpc?topic=vpc-block-storage"

		volumeProfileIdentityModel := make(map[string]interface{})
		volumeProfileIdentityModel["name"] = "sdp"

		volumeJobTypeMigrateParametersModel := make(map[string]interface{})
		volumeJobTypeMigrateParametersModel["bandwidth"] = int(1000)
		volumeJobTypeMigrateParametersModel["iops"] = int(1000)
		volumeJobTypeMigrateParametersModel["profile"] = []map[string]interface{}{volumeProfileIdentityModel}

		model := make(map[string]interface{})
		model["auto_delete"] = false
		model["completed_at"] = "2024-12-09T02:09:02.000Z"
		model["created_at"] = "2024-12-09T01:09:02.000Z"
		model["estimated_completion_at"] = "2019-01-01T12:00:00.000Z"
		model["href"] = "https://us-south.iaas.cloud.ibm.com/v1/volumes/r006-72b27b5c-f4b0-48bb-b954-5becc7c1dcb8/jobs/r006-095e9baf-01d4-4e29-986e-20d26606b82a"
		model["id"] = "r006-095e9baf-01d4-4e29-986e-20d26606b82a"
		model["job_type"] = "migrate"
		model["name"] = "my-volume-job"
		model["resource_type"] = "volume_job"
		model["started_at"] = "2024-12-09T01:10:02.000Z"
		model["status"] = "queued"
		model["status_reasons"] = []map[string]interface{}{volumeJobStatusReasonModel}
		model["parameters"] = []map[string]interface{}{volumeJobTypeMigrateParametersModel}

		assert.Equal(t, result, model)
	}

	volumeJobStatusReasonModel := new(vpcv1.VolumeJobStatusReason)
	volumeJobStatusReasonModel.Code = core.StringPtr("volume_detached_from_virtual_instance")
	volumeJobStatusReasonModel.Message = core.StringPtr("testString")
	volumeJobStatusReasonModel.MoreInfo = core.StringPtr("https://cloud.ibm.com/docs/vpc?topic=vpc-block-storage")

	volumeProfileIdentityModel := new(vpcv1.VolumeProfileIdentityByName)
	volumeProfileIdentityModel.Name = core.StringPtr("sdp")

	volumeJobTypeMigrateParametersModel := new(vpcv1.VolumeJobTypeMigrateParameters)
	volumeJobTypeMigrateParametersModel.Bandwidth = core.Int64Ptr(int64(1000))
	volumeJobTypeMigrateParametersModel.Iops = core.Int64Ptr(int64(1000))
	volumeJobTypeMigrateParametersModel.Profile = volumeProfileIdentityModel

	model := new(vpcv1.VolumeJob)
	model.AutoDelete = core.BoolPtr(false)
	model.CompletedAt = CreateMockDateTime("2024-12-09T02:09:02.000Z")
	model.CreatedAt = CreateMockDateTime("2024-12-09T01:09:02.000Z")
	model.EstimatedCompletionAt = CreateMockDateTime("2019-01-01T12:00:00.000Z")
	model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/volumes/r006-72b27b5c-f4b0-48bb-b954-5becc7c1dcb8/jobs/r006-095e9baf-01d4-4e29-986e-20d26606b82a")
	model.ID = core.StringPtr("r006-095e9baf-01d4-4e29-986e-20d26606b82a")
	model.JobType = core.StringPtr("migrate")
	model.Name = core.StringPtr("my-volume-job")
	model.ResourceType = core.StringPtr("volume_job")
	model.StartedAt = CreateMockDateTime("2024-12-09T01:10:02.000Z")
	model.Status = core.StringPtr("queued")
	model.StatusReasons = []vpcv1.VolumeJobStatusReason{*volumeJobStatusReasonModel}
	model.Parameters = volumeJobTypeMigrateParametersModel

	result, err := vpc.DataSourceIBMIsVolumeJobsVolumeJobToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsVolumeJobsVolumeJobStatusReasonToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["code"] = "volume_detached_from_virtual_instance"
		model["message"] = "testString"
		model["more_info"] = "https://cloud.ibm.com/docs/vpc?topic=vpc-block-storage"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.VolumeJobStatusReason)
	model.Code = core.StringPtr("volume_detached_from_virtual_instance")
	model.Message = core.StringPtr("testString")
	model.MoreInfo = core.StringPtr("https://cloud.ibm.com/docs/vpc?topic=vpc-block-storage")

	result, err := vpc.DataSourceIBMIsVolumeJobsVolumeJobStatusReasonToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsVolumeJobsVolumeJobTypeMigrateParametersToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		volumeProfileIdentityModel := make(map[string]interface{})
		volumeProfileIdentityModel["name"] = "general-purpose"

		model := make(map[string]interface{})
		model["bandwidth"] = int(1000)
		model["iops"] = int(10000)
		model["profile"] = []map[string]interface{}{volumeProfileIdentityModel}

		assert.Equal(t, result, model)
	}

	volumeProfileIdentityModel := new(vpcv1.VolumeProfileIdentityByName)
	volumeProfileIdentityModel.Name = core.StringPtr("general-purpose")

	model := new(vpcv1.VolumeJobTypeMigrateParameters)
	model.Bandwidth = core.Int64Ptr(int64(1000))
	model.Iops = core.Int64Ptr(int64(10000))
	model.Profile = volumeProfileIdentityModel

	result, err := vpc.DataSourceIBMIsVolumeJobsVolumeJobTypeMigrateParametersToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsVolumeJobsVolumeProfileIdentityToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["name"] = "general-purpose"
		model["href"] = "https://us-south.iaas.cloud.ibm.com/v1/volume/profiles/general-purpose"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.VolumeProfileIdentity)
	model.Name = core.StringPtr("general-purpose")
	model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/volume/profiles/general-purpose")

	result, err := vpc.DataSourceIBMIsVolumeJobsVolumeProfileIdentityToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsVolumeJobsVolumeProfileIdentityByNameToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["name"] = "general-purpose"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.VolumeProfileIdentityByName)
	model.Name = core.StringPtr("general-purpose")

	result, err := vpc.DataSourceIBMIsVolumeJobsVolumeProfileIdentityByNameToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsVolumeJobsVolumeProfileIdentityByHrefToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["href"] = "https://us-south.iaas.cloud.ibm.com/v1/volume/profiles/general-purpose"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.VolumeProfileIdentityByHref)
	model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/volume/profiles/general-purpose")

	result, err := vpc.DataSourceIBMIsVolumeJobsVolumeProfileIdentityByHrefToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsVolumeJobsVolumeJobTypeMigrateToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		volumeJobStatusReasonModel := make(map[string]interface{})
		volumeJobStatusReasonModel["code"] = "volume_detached_from_virtual_instance"
		volumeJobStatusReasonModel["message"] = "testString"
		volumeJobStatusReasonModel["more_info"] = "https://cloud.ibm.com/docs/vpc?topic=vpc-block-storage"

		volumeProfileIdentityModel := make(map[string]interface{})
		volumeProfileIdentityModel["name"] = "general-purpose"

		volumeJobTypeMigrateParametersModel := make(map[string]interface{})
		volumeJobTypeMigrateParametersModel["bandwidth"] = int(1000)
		volumeJobTypeMigrateParametersModel["iops"] = int(10000)
		volumeJobTypeMigrateParametersModel["profile"] = []map[string]interface{}{volumeProfileIdentityModel}

		model := make(map[string]interface{})
		model["auto_delete"] = false
		model["completed_at"] = "2019-01-01T12:00:00.000Z"
		model["created_at"] = "2019-01-01T12:00:00.000Z"
		model["estimated_completion_at"] = "2019-01-01T12:00:00.000Z"
		model["href"] = "https://us-south.iaas.cloud.ibm.com/v1/volumes/r006-72b27b5c-f4b0-48bb-b954-5becc7c1dcb8/jobs/r006-095e9baf-01d4-4e29-986e-20d26606b82a"
		model["id"] = "r006-095e9baf-01d4-4e29-986e-20d26606b82a"
		model["job_type"] = "migrate"
		model["name"] = "my-volume-job"
		model["resource_type"] = "volume_job"
		model["started_at"] = "2019-01-01T12:00:00.000Z"
		model["status"] = "canceled"
		model["status_reasons"] = []map[string]interface{}{volumeJobStatusReasonModel}
		model["parameters"] = []map[string]interface{}{volumeJobTypeMigrateParametersModel}

		assert.Equal(t, result, model)
	}

	volumeJobStatusReasonModel := new(vpcv1.VolumeJobStatusReason)
	volumeJobStatusReasonModel.Code = core.StringPtr("volume_detached_from_virtual_instance")
	volumeJobStatusReasonModel.Message = core.StringPtr("testString")
	volumeJobStatusReasonModel.MoreInfo = core.StringPtr("https://cloud.ibm.com/docs/vpc?topic=vpc-block-storage")

	volumeProfileIdentityModel := new(vpcv1.VolumeProfileIdentityByName)
	volumeProfileIdentityModel.Name = core.StringPtr("general-purpose")

	volumeJobTypeMigrateParametersModel := new(vpcv1.VolumeJobTypeMigrateParameters)
	volumeJobTypeMigrateParametersModel.Bandwidth = core.Int64Ptr(int64(1000))
	volumeJobTypeMigrateParametersModel.Iops = core.Int64Ptr(int64(10000))
	volumeJobTypeMigrateParametersModel.Profile = volumeProfileIdentityModel

	model := new(vpcv1.VolumeJobTypeMigrate)
	model.AutoDelete = core.BoolPtr(false)
	model.CompletedAt = CreateMockDateTime("2019-01-01T12:00:00.000Z")
	model.CreatedAt = CreateMockDateTime("2019-01-01T12:00:00.000Z")
	model.EstimatedCompletionAt = CreateMockDateTime("2019-01-01T12:00:00.000Z")
	model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/volumes/r006-72b27b5c-f4b0-48bb-b954-5becc7c1dcb8/jobs/r006-095e9baf-01d4-4e29-986e-20d26606b82a")
	model.ID = core.StringPtr("r006-095e9baf-01d4-4e29-986e-20d26606b82a")
	model.JobType = core.StringPtr("migrate")
	model.Name = core.StringPtr("my-volume-job")
	model.ResourceType = core.StringPtr("volume_job")
	model.StartedAt = CreateMockDateTime("2019-01-01T12:00:00.000Z")
	model.Status = core.StringPtr("canceled")
	model.StatusReasons = []vpcv1.VolumeJobStatusReason{*volumeJobStatusReasonModel}
	model.Parameters = volumeJobTypeMigrateParametersModel

	result, err := vpc.DataSourceIBMIsVolumeJobsVolumeJobTypeMigrateToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
