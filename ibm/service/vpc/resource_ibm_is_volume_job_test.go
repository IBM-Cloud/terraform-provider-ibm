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
	volumeName := fmt.Sprintf("tf-volume-%d", acctest.RandIntRange(10, 100))
	volumeJobName := fmt.Sprintf("tf-volume-job-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIsVolumeJobDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsVolumeJobConfigBasic(volumeName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_is_volume.gen_one_volume", "name", volumeName),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMIsVolumeJobUpdateConfigBasic(volumeName, volumeJobName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsVolumeJobExists("ibm_is_volume_job.gen_one_to_gen_two", conf),
					resource.TestCheckResourceAttr("ibm_is_volume.gen_one_volume", "name", volumeName),
					resource.TestCheckResourceAttr("ibm_is_volume_job.gen_one_to_gen_two", "name", volumeJobName),
					resource.TestCheckResourceAttr("ibm_is_volume_job.gen_one_to_gen_two", "job_type", "migrate"),
				),
			},
		},
	})
}

func TestAccIBMIsVolumeJobAllArgs(t *testing.T) {
	var conf vpcv1.VolumeJob
	volumeID := fmt.Sprintf("tf_volume_id_%d", acctest.RandIntRange(10, 100))
	jobType := "migrate"
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	jobTypeUpdate := "migrate"
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIsVolumeJobDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsVolumeJobConfig(volumeID, jobType, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsVolumeJobExists("ibm_is_volume_job.is_volume_job_instance", conf),
					resource.TestCheckResourceAttr("ibm_is_volume_job.is_volume_job_instance", "volume_id", volumeID),
					resource.TestCheckResourceAttr("ibm_is_volume_job.is_volume_job_instance", "job_type", jobType),
					resource.TestCheckResourceAttr("ibm_is_volume_job.is_volume_job_instance", "name", name),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMIsVolumeJobConfig(volumeID, jobTypeUpdate, nameUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_is_volume_job.is_volume_job_instance", "volume_id", volumeID),
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

func testAccCheckIBMIsVolumeJobConfigBasic(volumeName string) string {
	return fmt.Sprintf(`

		resource "ibm_is_volume" "gen_one_volume" {
			name     = "%s"
			capacity = 200
			profile  = "general-purpose"
			zone     = "%s"
		}
	`, volumeName, acc.ISZoneName)
}
func testAccCheckIBMIsVolumeJobUpdateConfigBasic(volumeName, volumeJobName string) string {
	return fmt.Sprintf(`

		resource "ibm_is_volume" "gen_one_volume" {
			name     = "%s"
			capacity = 200
			profile  = "general-purpose"
			zone     = "%s"
			lifecycle {
				ignore_changes = [ profile ]
			}
		}
		resource "ibm_is_volume_job" "gen_one_to_gen_two" {
			volume_id = ibm_is_volume.gen_one_volume.id
			job_type  = "migrate"
			name      = "%s"
			parameters {
				profile {
					name = "sdp"
				}
				bandwidth = 2000
				iops      = 5000
			}
		}
	`, volumeName, acc.ISZoneName, volumeJobName)
}

func testAccCheckIBMIsVolumeJobConfig(volumeID string, jobType string, name string) string {
	return fmt.Sprintf(`

		resource "ibm_is_volume_job" "is_volume_job_instance" {
			volume_id = "%s"
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
	`, volumeID, jobType, name)
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
