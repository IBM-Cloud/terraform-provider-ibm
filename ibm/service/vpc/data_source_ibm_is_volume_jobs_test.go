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

func TestAccIBMIsVolumeJobsDataSourceBasic(t *testing.T) {
	volumeName := fmt.Sprintf("tf-volume-%d", acctest.RandIntRange(10, 100))
	volumeJobName := fmt.Sprintf("tf-volume-job-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsVolumeJobsDataSourceConfigBasic(volumeName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("ibm_is_volume.gen_one_volume", "id"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMIsVolumeJobsDataSourceConfig(volumeName, volumeJobName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_jobs.is_volume_jobs_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_jobs.is_volume_jobs_instance", "volume_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_jobs.is_volume_jobs_instance", "jobs.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_jobs.is_volume_jobs_instance", "jobs.0.auto_delete"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_jobs.is_volume_jobs_instance", "jobs.0.completed_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_jobs.is_volume_jobs_instance", "jobs.0.created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_jobs.is_volume_jobs_instance", "jobs.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_jobs.is_volume_jobs_instance", "jobs.0.id"),
					resource.TestCheckResourceAttr("data.ibm_is_volume_jobs.is_volume_jobs_instance", "jobs.0.job_type", "migrate"),
					resource.TestCheckResourceAttr("data.ibm_is_volume_jobs.is_volume_jobs_instance", "jobs.0.name", volumeJobName),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_jobs.is_volume_jobs_instance", "jobs.0.parameters.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_jobs.is_volume_jobs_instance", "jobs.0.resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_jobs.is_volume_jobs_instance", "jobs.0.started_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_jobs.is_volume_jobs_instance", "jobs.0.status"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volume_jobs.is_volume_jobs_instance", "jobs.0.status_reasons.#"),
				),
			},
		},
	})
}

func testAccCheckIBMIsVolumeJobsDataSourceConfigBasic(volumeName string) string {
	return fmt.Sprintf(`
		resource "ibm_is_volume" "gen_one_volume" {
			name     = "%s"
			capacity = 10
			profile  = "general-purpose"
			zone     = "%s"
		}
	`, volumeName, acc.ISZoneName)
}

func testAccCheckIBMIsVolumeJobsDataSourceConfig(volumeName string, volumeJobName string) string {
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

		data "ibm_is_volume_jobs" "is_volume_jobs_instance" {
			volume_id = ibm_is_volume_job.gen_one_to_gen_two.volume_id
		}
	`, volumeName, acc.ISZoneName, volumeJobName)
}
