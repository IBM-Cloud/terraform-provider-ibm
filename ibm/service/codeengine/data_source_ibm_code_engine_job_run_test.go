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

func TestAccIbmCodeEngineJobRunDataSourceBasic(t *testing.T) {
	jobName := fmt.Sprintf("tf-data-job-basic-%d", acctest.RandIntRange(10, 1000))
	jobRunName := fmt.Sprintf("tf-data-job-run-basic-%d", acctest.RandIntRange(10, 1000))
	jobImageReference := "icr.io/codeengine/helloworld"

	projectID := acc.CeProjectId

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmCodeEngineJobRunDataSourceConfigBasic(projectID, jobImageReference, jobName, jobRunName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_code_engine_job_run.code_engine_job_run_instance", "job_run_id"),
					resource.TestCheckResourceAttr("data.ibm_code_engine_job_run.code_engine_job_run_instance", "project_id", projectID),
					resource.TestCheckResourceAttr("data.ibm_code_engine_job_run.code_engine_job_run_instance", "name", jobRunName),
					resource.TestCheckResourceAttr("data.ibm_code_engine_job_run.code_engine_job_run_instance", "jobName", jobName),
					resource.TestCheckResourceAttr("data.ibm_code_engine_job_run.code_engine_job_run_instance", "resource_type", "job_run_v2"),
				),
			},
		},
	})
}

func testAccCheckIbmCodeEngineJobRunDataSourceConfigBasic(projectID string, jobImageReference string, jobName string, jobRunName string) string {
	return fmt.Sprintf(`
		data "ibm_code_engine_project" "code_engine_project_instance" {
			project_id = "%s"
		}

		resource "ibm_code_engine_job" "code_engine_job_instance" {
			project_id = data.ibm_code_engine_project.code_engine_project_instance.project_id
			image_reference = "%s"
			name = "%s"
		}

		resource "ibm_code_engine_job_run" "code_engine_job_run_instance" {
			project_id = ibm_code_engine_project.ce_project.project_id
			name       = "%s"
			job_name   = ibm_code_engine_job.code_engine_job_instance.name
		}

		data "ibm_code_engine_job_run" "code_engine_job_run_instance" {
			project_id = ibm_code_engine_project.ce_project.project_id
			name       = data.ibm_code_engine_job_run.code_engine_job_run_instance.name
		}


	`, projectID, jobImageReference, jobName, jobRunName)
}
