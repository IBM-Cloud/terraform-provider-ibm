// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	vpc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/vpc"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/assert"
)

func TestAccIBMIsVolumeJobBasic(t *testing.T) {
	var conf vpcv1.VolumeJob
	volumeID := fmt.Sprintf("tf_volume_id_%d", acctest.RandIntRange(10, 100))
	jobType := "migrate"
	jobTypeUpdate := "migrate"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIsVolumeJobDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsVolumeJobConfigBasic(volumeID, jobType),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsVolumeJobExists("ibm_is_volume_job.is_volume_job_instance", conf),
					resource.TestCheckResourceAttr("ibm_is_volume_job.is_volume_job_instance", "volume_id", volumeID),
					resource.TestCheckResourceAttr("ibm_is_volume_job.is_volume_job_instance", "job_type", jobType),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMIsVolumeJobConfigBasic(volumeID, jobTypeUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_is_volume_job.is_volume_job_instance", "volume_id", volumeID),
					resource.TestCheckResourceAttr("ibm_is_volume_job.is_volume_job_instance", "job_type", jobTypeUpdate),
				),
			},
		},
	})
}

func TestAccIBMIsVolumeJobAllArgs(t *testing.T) {
	var conf vpcv1.VolumeJob
	volumeID := fmt.Sprintf("tf_volume_id_%d", acctest.RandIntRange(10, 100))
	start := fmt.Sprintf("tf_start_%d", acctest.RandIntRange(10, 100))
	limit := fmt.Sprintf("%d", acctest.RandIntRange(1, 100))
	jobType := "migrate"
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	startUpdate := fmt.Sprintf("tf_start_%d", acctest.RandIntRange(10, 100))
	limitUpdate := fmt.Sprintf("%d", acctest.RandIntRange(1, 100))
	jobTypeUpdate := "migrate"
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIsVolumeJobDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsVolumeJobConfig(volumeID, start, limit, jobType, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsVolumeJobExists("ibm_is_volume_job.is_volume_job_instance", conf),
					resource.TestCheckResourceAttr("ibm_is_volume_job.is_volume_job_instance", "volume_id", volumeID),
					resource.TestCheckResourceAttr("ibm_is_volume_job.is_volume_job_instance", "start", start),
					resource.TestCheckResourceAttr("ibm_is_volume_job.is_volume_job_instance", "limit", limit),
					resource.TestCheckResourceAttr("ibm_is_volume_job.is_volume_job_instance", "job_type", jobType),
					resource.TestCheckResourceAttr("ibm_is_volume_job.is_volume_job_instance", "name", name),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMIsVolumeJobConfig(volumeID, startUpdate, limitUpdate, jobTypeUpdate, nameUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_is_volume_job.is_volume_job_instance", "volume_id", volumeID),
					resource.TestCheckResourceAttr("ibm_is_volume_job.is_volume_job_instance", "start", startUpdate),
					resource.TestCheckResourceAttr("ibm_is_volume_job.is_volume_job_instance", "limit", limitUpdate),
					resource.TestCheckResourceAttr("ibm_is_volume_job.is_volume_job_instance", "job_type", jobTypeUpdate),
					resource.TestCheckResourceAttr("ibm_is_volume_job.is_volume_job_instance", "name", nameUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_is_volume_job.is_volume_job_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMIsVolumeJobConfigBasic(volumeID string, jobType string) string {
	return fmt.Sprintf(`
		resource "ibm_is_volume_job" "is_volume_job_instance" {
			volume_id = "%s"
			job_type = "%s"
		}
	`, volumeID, jobType)
}

func testAccCheckIBMIsVolumeJobConfig(volumeID string, start string, limit string, jobType string, name string) string {
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
	`, volumeID, start, limit, jobType, name)
}

func testAccCheckIBMIsVolumeJobExists(n string, obj vpcv1.VolumeJob) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		vpcClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
		if err != nil {
			return err
		}

		getVolumeJobOptions := &vpcv1.GetVolumeJobOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getVolumeJobOptions.SetVolumeID(parts[0])
		getVolumeJobOptions.SetID(parts[1])

		volumeJobIntf, _, err := vpcClient.GetVolumeJob(getVolumeJobOptions)
		if err != nil {
			return err
		}

		volumeJob := volumeJobIntf.(*vpcv1.VolumeJob)
		obj = *volumeJob
		return nil
	}
}

func testAccCheckIBMIsVolumeJobDestroy(s *terraform.State) error {
	vpcClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_volume_job" {
			continue
		}

		getVolumeJobOptions := &vpcv1.GetVolumeJobOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getVolumeJobOptions.SetVolumeID(parts[0])
		getVolumeJobOptions.SetID(parts[1])

		// Try to find the key
		_, response, err := vpcClient.GetVolumeJob(getVolumeJobOptions)

		if err == nil {
			return fmt.Errorf("is_volume_job still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for is_volume_job (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func TestResourceIBMIsVolumeJobVolumeJobTypeMigrateParametersToMap(t *testing.T) {
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

	result, err := vpc.ResourceIBMIsVolumeJobVolumeJobTypeMigrateParametersToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsVolumeJobVolumeProfileIdentityToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["name"] = "general-purpose"
		model["href"] = "https://us-south.iaas.cloud.ibm.com/v1/volume/profiles/general-purpose"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.VolumeProfileIdentity)
	model.Name = core.StringPtr("general-purpose")
	model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/volume/profiles/general-purpose")

	result, err := vpc.ResourceIBMIsVolumeJobVolumeProfileIdentityToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsVolumeJobVolumeProfileIdentityByNameToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["name"] = "general-purpose"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.VolumeProfileIdentityByName)
	model.Name = core.StringPtr("general-purpose")

	result, err := vpc.ResourceIBMIsVolumeJobVolumeProfileIdentityByNameToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsVolumeJobVolumeProfileIdentityByHrefToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["href"] = "https://us-south.iaas.cloud.ibm.com/v1/volume/profiles/general-purpose"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.VolumeProfileIdentityByHref)
	model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/volume/profiles/general-purpose")

	result, err := vpc.ResourceIBMIsVolumeJobVolumeProfileIdentityByHrefToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsVolumeJobVolumeJobStatusReasonToMap(t *testing.T) {
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

	result, err := vpc.ResourceIBMIsVolumeJobVolumeJobStatusReasonToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsVolumeJobMapToVolumeJobTypeMigrateParameters(t *testing.T) {
	checkResult := func(result *vpcv1.VolumeJobTypeMigrateParameters) {
		volumeProfileIdentityModel := new(vpcv1.VolumeProfileIdentityByName)
		volumeProfileIdentityModel.Name = core.StringPtr("general-purpose")

		model := new(vpcv1.VolumeJobTypeMigrateParameters)
		model.Bandwidth = core.Int64Ptr(int64(1000))
		model.Iops = core.Int64Ptr(int64(10000))
		model.Profile = volumeProfileIdentityModel

		assert.Equal(t, result, model)
	}

	volumeProfileIdentityModel := make(map[string]interface{})
	volumeProfileIdentityModel["name"] = "general-purpose"

	model := make(map[string]interface{})
	model["bandwidth"] = int(1000)
	model["iops"] = int(10000)
	model["profile"] = []interface{}{volumeProfileIdentityModel}

	result, err := vpc.ResourceIBMIsVolumeJobMapToVolumeJobTypeMigrateParameters(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsVolumeJobMapToVolumeProfileIdentity(t *testing.T) {
	checkResult := func(result vpcv1.VolumeProfileIdentityIntf) {
		model := new(vpcv1.VolumeProfileIdentity)
		model.Name = core.StringPtr("general-purpose")
		model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/volume/profiles/general-purpose")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["name"] = "general-purpose"
	model["href"] = "https://us-south.iaas.cloud.ibm.com/v1/volume/profiles/general-purpose"

	result, err := vpc.ResourceIBMIsVolumeJobMapToVolumeProfileIdentity(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsVolumeJobMapToVolumeProfileIdentityByName(t *testing.T) {
	checkResult := func(result *vpcv1.VolumeProfileIdentityByName) {
		model := new(vpcv1.VolumeProfileIdentityByName)
		model.Name = core.StringPtr("general-purpose")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["name"] = "general-purpose"

	result, err := vpc.ResourceIBMIsVolumeJobMapToVolumeProfileIdentityByName(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsVolumeJobMapToVolumeProfileIdentityByHref(t *testing.T) {
	checkResult := func(result *vpcv1.VolumeProfileIdentityByHref) {
		model := new(vpcv1.VolumeProfileIdentityByHref)
		model.Href = core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/volume/profiles/general-purpose")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["href"] = "https://us-south.iaas.cloud.ibm.com/v1/volume/profiles/general-purpose"

	result, err := vpc.ResourceIBMIsVolumeJobMapToVolumeProfileIdentityByHref(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsVolumeJobMapToVolumeJobPrototype(t *testing.T) {
	checkResult := func(result vpcv1.VolumeJobPrototypeIntf) {
		volumeProfileIdentityModel := new(vpcv1.VolumeProfileIdentityByName)
		volumeProfileIdentityModel.Name = core.StringPtr("general-purpose")

		volumeJobTypeMigrateParametersModel := new(vpcv1.VolumeJobTypeMigrateParameters)
		volumeJobTypeMigrateParametersModel.Bandwidth = core.Int64Ptr(int64(1000))
		volumeJobTypeMigrateParametersModel.Iops = core.Int64Ptr(int64(10000))
		volumeJobTypeMigrateParametersModel.Profile = volumeProfileIdentityModel

		model := new(vpcv1.VolumeJobPrototype)
		model.JobType = core.StringPtr("migrate")
		model.Name = core.StringPtr("my-volume-job")
		model.Parameters = volumeJobTypeMigrateParametersModel

		assert.Equal(t, result, model)
	}

	volumeProfileIdentityModel := make(map[string]interface{})
	volumeProfileIdentityModel["name"] = "general-purpose"

	volumeJobTypeMigrateParametersModel := make(map[string]interface{})
	volumeJobTypeMigrateParametersModel["bandwidth"] = int(1000)
	volumeJobTypeMigrateParametersModel["iops"] = int(10000)
	volumeJobTypeMigrateParametersModel["profile"] = []interface{}{volumeProfileIdentityModel}

	model := make(map[string]interface{})
	model["job_type"] = "migrate"
	model["name"] = "my-volume-job"
	model["parameters"] = []interface{}{volumeJobTypeMigrateParametersModel}

	result, err := vpc.ResourceIBMIsVolumeJobMapToVolumeJobPrototype(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMIsVolumeJobMapToVolumeJobPrototypeVolumeJobTypeMigratePrototype(t *testing.T) {
	checkResult := func(result *vpcv1.VolumeJobPrototypeVolumeJobTypeMigratePrototype) {
		volumeProfileIdentityModel := new(vpcv1.VolumeProfileIdentityByName)
		volumeProfileIdentityModel.Name = core.StringPtr("general-purpose")

		volumeJobTypeMigrateParametersModel := new(vpcv1.VolumeJobTypeMigrateParameters)
		volumeJobTypeMigrateParametersModel.Bandwidth = core.Int64Ptr(int64(1000))
		volumeJobTypeMigrateParametersModel.Iops = core.Int64Ptr(int64(10000))
		volumeJobTypeMigrateParametersModel.Profile = volumeProfileIdentityModel

		model := new(vpcv1.VolumeJobPrototypeVolumeJobTypeMigratePrototype)
		model.Name = core.StringPtr("my-volume-job")
		model.JobType = core.StringPtr("migrate")
		model.Parameters = volumeJobTypeMigrateParametersModel

		assert.Equal(t, result, model)
	}

	volumeProfileIdentityModel := make(map[string]interface{})
	volumeProfileIdentityModel["name"] = "general-purpose"

	volumeJobTypeMigrateParametersModel := make(map[string]interface{})
	volumeJobTypeMigrateParametersModel["bandwidth"] = int(1000)
	volumeJobTypeMigrateParametersModel["iops"] = int(10000)
	volumeJobTypeMigrateParametersModel["profile"] = []interface{}{volumeProfileIdentityModel}

	model := make(map[string]interface{})
	model["name"] = "my-volume-job"
	model["job_type"] = "migrate"
	model["parameters"] = []interface{}{volumeJobTypeMigrateParametersModel}

	result, err := vpc.ResourceIBMIsVolumeJobMapToVolumeJobPrototypeVolumeJobTypeMigratePrototype(model)
	assert.Nil(t, err)
	checkResult(result)
}
