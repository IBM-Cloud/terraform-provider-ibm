// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package codeengine_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmCodeEngineJobDataSourceBasic(t *testing.T) {
	jobName := fmt.Sprintf("tf-data-job-basic-%d", acctest.RandIntRange(10, 1000))
	jobImageReference := "icr.io/codeengine/helloworld"

	projectID := acc.CeProjectId

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmCodeEngineJobDataSourceConfigBasic(projectID, jobImageReference, jobName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_code_engine_job.code_engine_job_instance", "job_id"),
					resource.TestCheckResourceAttr("data.ibm_code_engine_job.code_engine_job_instance", "project_id", projectID),
					resource.TestCheckResourceAttr("data.ibm_code_engine_job.code_engine_job_instance", "name", jobName),
					resource.TestCheckResourceAttr("data.ibm_code_engine_job.code_engine_job_instance", "resource_type", "job_v2"),
					resource.TestCheckResourceAttr("data.ibm_code_engine_job.code_engine_job_instance", "image_reference", jobImageReference),
					resource.TestCheckResourceAttr("data.ibm_code_engine_job.code_engine_job_instance", "run_service_account", "default"),
					resource.TestCheckResourceAttr("data.ibm_code_engine_job.code_engine_job_instance", "scale_cpu_limit", "1"),
					resource.TestCheckResourceAttr("data.ibm_code_engine_job.code_engine_job_instance", "scale_ephemeral_storage_limit", "400M"),
					resource.TestCheckResourceAttr("data.ibm_code_engine_job.code_engine_job_instance", "scale_max_execution_time", "7200"),
					resource.TestCheckResourceAttr("data.ibm_code_engine_job.code_engine_job_instance", "scale_memory_limit", "4G"),
					resource.TestCheckResourceAttr("data.ibm_code_engine_job.code_engine_job_instance", "scale_retry_limit", "3"),
					resource.TestCheckResourceAttr("data.ibm_code_engine_job.code_engine_job_instance", "computed_env_variables.#", "3"),
				),
			},
		},
	})
}

func TestAccIbmCodeEngineJobDataSourceExtended(t *testing.T) {
	jobName := fmt.Sprintf("tf-data-job-extended-%d", acctest.RandIntRange(10, 1000))
	jobImageReference := "icr.io/codeengine/helloworld"
	jobRunMode := "task"
	jobRunServiceAccount := "default"
	jobScaleCpuLimit := "0.5"
	jobScaleEphemeralStorageLimit := "500M"
	jobScaleMaxExecutionTime := "3600"
	jobScaleMemoryLimit := "1G"
	jobScaleRetryLimit := "2"

	projectID := acc.CeProjectId

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmCodeEngineJobDataSourceConfig(projectID, jobImageReference, jobName, jobRunMode, jobRunServiceAccount, jobScaleCpuLimit, jobScaleEphemeralStorageLimit, jobScaleMaxExecutionTime, jobScaleMemoryLimit, jobScaleRetryLimit),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_code_engine_job.code_engine_job_instance", "job_id"),
					resource.TestCheckResourceAttr("data.ibm_code_engine_job.code_engine_job_instance", "project_id", projectID),
					resource.TestCheckResourceAttr("data.ibm_code_engine_job.code_engine_job_instance", "name", jobName),
					resource.TestCheckResourceAttr("data.ibm_code_engine_job.code_engine_job_instance", "image_reference", jobImageReference),
					resource.TestCheckResourceAttr("data.ibm_code_engine_job.code_engine_job_instance", "run_mode", jobRunMode),
					resource.TestCheckResourceAttr("data.ibm_code_engine_job.code_engine_job_instance", "run_service_account", jobRunServiceAccount),
					resource.TestCheckResourceAttr("data.ibm_code_engine_job.code_engine_job_instance", "scale_cpu_limit", jobScaleCpuLimit),
					resource.TestCheckResourceAttr("data.ibm_code_engine_job.code_engine_job_instance", "scale_ephemeral_storage_limit", jobScaleEphemeralStorageLimit),
					resource.TestCheckResourceAttr("data.ibm_code_engine_job.code_engine_job_instance", "scale_max_execution_time", jobScaleMaxExecutionTime),
					resource.TestCheckResourceAttr("data.ibm_code_engine_job.code_engine_job_instance", "scale_memory_limit", jobScaleMemoryLimit),
					resource.TestCheckResourceAttr("data.ibm_code_engine_job.code_engine_job_instance", "scale_retry_limit", jobScaleRetryLimit),
					resource.TestCheckResourceAttr("data.ibm_code_engine_job.code_engine_job_instance", "computed_env_variables.#", "3"),
				),
			},
		},
	})
}

func testAccCheckIbmCodeEngineJobDataSourceConfigBasic(projectID string, jobImageReference string, jobName string) string {
	return fmt.Sprintf(`
		data "ibm_code_engine_project" "code_engine_project_instance" {
			project_id = "%s"
		}

		resource "ibm_code_engine_job" "code_engine_job_instance" {
			project_id = data.ibm_code_engine_project.code_engine_project_instance.project_id
			image_reference = "%s"
			name = "%s"
		}

		data "ibm_code_engine_job" "code_engine_job_instance" {
			project_id = ibm_code_engine_job.code_engine_job_instance.project_id
			name = ibm_code_engine_job.code_engine_job_instance.name
		}
	`, projectID, jobImageReference, jobName)
}

func testAccCheckIbmCodeEngineJobDataSourceConfig(projectID string, jobImageReference string, jobName string, jobRunMode string, jobRunServiceAccount string, jobScaleCpuLimit string, jobScaleEphemeralStorageLimit string, jobScaleMaxExecutionTime string, jobScaleMemoryLimit string, jobScaleRetryLimit string) string {
	return fmt.Sprintf(`
		data "ibm_code_engine_project" "code_engine_project_instance" {
			project_id = "%s"
		}

		resource "ibm_code_engine_job" "code_engine_job_instance" {
			project_id =  data.ibm_code_engine_project.code_engine_project_instance.project_id
			image_reference = "%s"
			name = "%s"
			run_mode = "%s"
			run_service_account = "%s"
			scale_cpu_limit = "%s"
			scale_ephemeral_storage_limit = "%s"
			scale_max_execution_time = %s
			scale_memory_limit = "%s"
			scale_retry_limit = %s
		}

		data "ibm_code_engine_job" "code_engine_job_instance" {
			project_id = ibm_code_engine_job.code_engine_job_instance.project_id
			name = ibm_code_engine_job.code_engine_job_instance.name
		}
	`, projectID, jobImageReference, jobName, jobRunMode, jobRunServiceAccount, jobScaleCpuLimit, jobScaleEphemeralStorageLimit, jobScaleMaxExecutionTime, jobScaleMemoryLimit, jobScaleRetryLimit)
}
