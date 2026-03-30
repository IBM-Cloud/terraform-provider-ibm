// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMIsVolumeJobDataSourceBasic(t *testing.T) {
	volumeName := fmt.Sprintf("tf-volume-%d", acctest.RandIntRange(10, 100))
	volumeJobName := fmt.Sprintf("tf-volume-job-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsVolumeJobDataSourceConfigBasic(volumeName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("ibm_is_volume.gen_one_volume", "id"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMIsVolumeJobDataSourceConfig(volumeName, volumeJobName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_job.is_volume_job_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_job.is_volume_job_instance", "volume_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_job.is_volume_job_instance", "volume_job_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_job.is_volume_job_instance", "auto_delete"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_job.is_volume_job_instance", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_job.is_volume_job_instance", "completed_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_job.is_volume_job_instance", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_job.is_volume_job_instance", "job_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_job.is_volume_job_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_job.is_volume_job_instance", "parameters.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_job.is_volume_job_instance", "resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_job.is_volume_job_instance", "status"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_job.is_volume_job_instance", "started_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_job.is_volume_job_instance", "status_reasons.#"),
				),
			},
		},
	})
}

func testAccCheckIBMIsVolumeJobDataSourceConfigBasic(volumeName string) string {
	return fmt.Sprintf(`
		resource "ibm_is_volume" "gen_one_volume" {
			name     = "%s"
			capacity = 10
			profile  = "general-purpose"
			zone     = "%s"
		}
	`, volumeName, acc.ISZoneName)
}

func testAccCheckIBMIsVolumeJobDataSourceConfig(volumeName string, volumeJobName string) string {
	return fmt.Sprintf(`
		resource "ibm_is_volume" "gen_one_volume" {
			name     = "%s"
			capacity = 10
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
				bandwidth = 1000
    			iops      = 3000
			}
		}

		data "ibm_is_volume_job" "is_volume_job_instance" {
			volume_id 		= ibm_is_volume_job.gen_one_to_gen_two.volume_id
			volume_job_id 	= ibm_is_volume_job.gen_one_to_gen_two.volume_job_id
		}
	`, volumeName, acc.ISZoneName, volumeJobName)
}
