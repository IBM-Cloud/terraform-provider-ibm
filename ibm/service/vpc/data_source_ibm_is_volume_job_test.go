// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.107.1-41b0fbd0-20250825-080732
 */

package vpc_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/stretchr/testify/assert"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	vpc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/vpc"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
)

func TestAccIBMIsVolumeJobDataSourceBasic(t *testing.T) {
	volumeJobVolumeID := fmt.Sprintf("tf_volume_id_%d", acctest.RandIntRange(10, 100))
	volumeJobJobType := "migrate"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsVolumeJobDataSourceConfigBasic(volumeJobVolumeID, volumeJobJobType),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_job.is_volume_job_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_job.is_volume_job_instance", "volume_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_job.is_volume_job_instance", "volume_job_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_job.is_volume_job_instance", "auto_delete"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_job.is_volume_job_instance", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_job.is_volume_job_instance", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_job.is_volume_job_instance", "job_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_job.is_volume_job_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_job.is_volume_job_instance", "resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_job.is_volume_job_instance", "status"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_job.is_volume_job_instance", "status_reasons.#"),
				),
			},
		},
	})
}

func TestAccIBMIsVolumeJobDataSourceAllArgs(t *testing.T) {
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
				Config: testAccCheckIBMIsVolumeJobDataSourceConfig(volumeJobVolumeID, volumeJobStart, volumeJobLimit, volumeJobJobType, volumeJobName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_job.is_volume_job_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_job.is_volume_job_instance", "volume_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_job.is_volume_job_instance", "volume_job_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_job.is_volume_job_instance", "auto_delete"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_job.is_volume_job_instance", "completed_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_job.is_volume_job_instance", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_job.is_volume_job_instance", "estimated_completion_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_job.is_volume_job_instance", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_job.is_volume_job_instance", "job_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_job.is_volume_job_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_job.is_volume_job_instance", "resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_job.is_volume_job_instance", "started_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_job.is_volume_job_instance", "status"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_job.is_volume_job_instance", "status_reasons.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_job.is_volume_job_instance", "status_reasons.0.code"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_job.is_volume_job_instance", "status_reasons.0.message"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_job.is_volume_job_instance", "status_reasons.0.more_info"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_job.is_volume_job_instance", "parameters.#"),
				),
			},
		},
	})
}

func testAccCheckIBMIsVolumeJobDataSourceConfigBasic(volumeJobVolumeID string, volumeJobJobType string) string {
	return fmt.Sprintf(`
		resource "ibm_is_volume_job" "is_volume_job_instance" {
			volume_id = "%s"
			job_type = "%s"
		}

		data "ibm_is_volume_job" "is_volume_job_instance" {
			volume_id = ibm_is_volume_job.is_volume_job_instance.volume_id
			volume_job_id = ibm_is_volume_job.is_volume_job_instance.volume_job_id
		}
	`, volumeJobVolumeID, volumeJobJobType)
}

func testAccCheckIBMIsVolumeJobDataSourceConfig(volumeJobVolumeID string, volumeJobStart string, volumeJobLimit string, volumeJobJobType string, volumeJobName string) string {
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

		data "ibm_is_volume_job" "is_volume_job_instance" {
			volume_id = ibm_is_volume_job.is_volume_job_instance.volume_id
			volume_job_id = ibm_is_volume_job.is_volume_job_instance.volume_job_id
		}
	`, volumeJobVolumeID, volumeJobStart, volumeJobLimit, volumeJobJobType, volumeJobName)
}

func TestDataSourceIBMIsVolumeJobVolumeJobStatusReasonToMap(t *testing.T) {
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

	result, err := vpc.DataSourceIBMIsVolumeJobVolumeJobStatusReasonToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsVolumeJobVolumeJobTypeMigrateParametersToMap(t *testing.T) {
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

	result, err := vpc.DataSourceIBMIsVolumeJobVolumeJobTypeMigrateParametersToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsVolumeJobVolumeProfileIdentityToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["name"] = "general-purpose"
		model["href"] = "https://us-south.iaas.cloud.ibm.com/v1/volume/profiles/general-purpose"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.VolumeProfileIdentity)
	model.Name = core.StringPtr("general-purpose")
	model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/volume/profiles/general-purpose")

	result, err := vpc.DataSourceIBMIsVolumeJobVolumeProfileIdentityToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsVolumeJobVolumeProfileIdentityByNameToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["name"] = "general-purpose"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.VolumeProfileIdentityByName)
	model.Name = core.StringPtr("general-purpose")

	result, err := vpc.DataSourceIBMIsVolumeJobVolumeProfileIdentityByNameToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsVolumeJobVolumeProfileIdentityByHrefToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["href"] = "https://us-south.iaas.cloud.ibm.com/v1/volume/profiles/general-purpose"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.VolumeProfileIdentityByHref)
	model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/volume/profiles/general-purpose")

	result, err := vpc.DataSourceIBMIsVolumeJobVolumeProfileIdentityByHrefToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
