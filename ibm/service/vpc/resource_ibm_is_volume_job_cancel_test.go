// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"

	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMIsVolumeJobCancelBasic(t *testing.T) {
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
				Config: testAccCheckIBMIsVolumeJobCancelConfigBasic(volumeID, jobType),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsVolumeJobExists("ibm_is_volume_job.is_volume_job_instance", conf),
					resource.TestCheckResourceAttr("ibm_is_volume_job.is_volume_job_instance", "volume_id", volumeID),
					resource.TestCheckResourceAttr("ibm_is_volume_job.is_volume_job_instance", "job_type", jobType),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMIsVolumeJobCancelConfigBasic(volumeID, jobTypeUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_is_volume_job.is_volume_job_instance", "volume_id", volumeID),
					resource.TestCheckResourceAttr("ibm_is_volume_job.is_volume_job_instance", "job_type", jobTypeUpdate),
				),
			},
		},
	})
}

func testAccCheckIBMIsVolumeJobCancelConfigBasic(volumeID string, jobType string) string {
	return fmt.Sprintf(`
		resource "ibm_is_volume_job" "is_volume_job_instance" {
			volume_id = "%s"
			job_type = "%s"
		}
		
		resource "ibm_is_volume_job_cancel" "cancel_migration" {
			volume_id        = ibm_is_volume_job.is_volume_job_instance.volume_id
			volume_job_id    = ibm_is_volume_job.is_volume_job_instance.volume_job_id
		}

	`, volumeID, jobType)
}

func testAccCheckIBMIsVolumeJobCancelExists(n string, obj vpcv1.VolumeJob) resource.TestCheckFunc {

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

func testAccCheckIBMIsVolumeJobCancelDestroy(s *terraform.State) error {
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
