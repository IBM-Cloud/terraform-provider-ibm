// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package codeengine_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/code-engine-go-sdk/codeenginev2"
)

func TestAccIbmCodeEngineJobRunBasic(t *testing.T) {
	var conf codeenginev2.JobRun
	jobName := fmt.Sprintf("tf-data-job-basic-%d", acctest.RandIntRange(10, 1000))
	jobRunName := fmt.Sprintf("tf-data-job-run-basic-%d", acctest.RandIntRange(10, 1000))
	jobImageReference := "icr.io/codeengine/helloworld"

	projectID := acc.CeProjectId

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmCodeEngineJobRunDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmCodeEngineJobRunConfigBasic(projectID, jobImageReference, jobName, jobRunName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmCodeEngineJobRunExists("ibm_code_engine_job_run.code_engine_job_run_instance", conf),
					resource.TestCheckResourceAttrSet("ibm_code_engine_job_run.code_engine_job_run_instance", "job_run_id"),
					resource.TestCheckResourceAttr("ibm_code_engine_job_run.code_engine_job_run_instance", "project_id", projectID),
					resource.TestCheckResourceAttr("ibm_code_engine_job_run.code_engine_job_run_instance", "name", jobRunName),
					resource.TestCheckResourceAttr("ibm_code_engine_job_run.code_engine_job_run_instance", "job_name", jobName),
					resource.TestCheckResourceAttr("ibm_code_engine_job_run.code_engine_job_run_instance", "resource_type", "job_run_v2"),
				),
			},
		},
	})
}

func testAccCheckIbmCodeEngineJobRunConfigBasic(projectID string, imageReference string, jobName string, jobRunName string) string {
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
			project_id = data.ibm_code_engine_project.code_engine_project_instance.project_id
			name       = "%s"
			job_name   = ibm_code_engine_job.code_engine_job_instance.name
		}
	`, projectID, imageReference, jobName, jobRunName)
}

func testAccCheckIbmCodeEngineJobRunExists(n string, obj codeenginev2.JobRun) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		codeEngineClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CodeEngineV2()
		if err != nil {
			return err
		}

		getJobRunOptions := &codeenginev2.GetJobRunOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getJobRunOptions.SetProjectID(parts[0])
		getJobRunOptions.SetName(parts[1])

		job, _, err := codeEngineClient.GetJobRun(getJobRunOptions)
		if err != nil {
			return err
		}

		obj = *job
		return nil
	}
}

func testAccCheckIbmCodeEngineJobRunDestroy(s *terraform.State) error {
	codeEngineClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CodeEngineV2()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_code_engine_job_run" {
			continue
		}

		getJobRunOptions := &codeenginev2.GetJobRunOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getJobRunOptions.SetProjectID(parts[0])
		getJobRunOptions.SetName(parts[1])

		// Try to find the key
		_, response, err := codeEngineClient.GetJobRun(getJobRunOptions)

		if err == nil {
			return fmt.Errorf("code_engine_job_run still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for code_engine_job_run (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
