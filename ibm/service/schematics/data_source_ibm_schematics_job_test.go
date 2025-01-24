// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package schematics_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMSchematicsJobDataSourceBasic(t *testing.T) {
	jobCommandObject := "action"
	//jobCommandObjectID := fmt.Sprintf("command_object_id_%d", acctest.RandIntRange(10, 100))
	jobCommandName := "ansible_playbook_run"
	jobCommandParameter := fmt.Sprintf("command_parameter_%d", acctest.RandIntRange(10, 100))
	jobLocation := "us-south"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMSchematicsJobDataSourceConfig(jobCommandObject, acc.ActionID, jobCommandName, jobCommandParameter, jobLocation),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "job_id"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "command_object"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "command_object_id"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "command_name"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "location"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "resource_group"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "submitted_at"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "submitted_by"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "start_at"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "end_at"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "status.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "log_summary.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "updated_at"),
				),
			},
		},
	})
}

func testAccCheckIBMSchematicsJobDataSourceConfig(jobCommandObject string, jobCommandObjectID, jobCommandName string, jobCommandParameter string, jobLocation string) string {
	return fmt.Sprintf(`
		resource "ibm_schematics_job" "schematics_job" {
			command_object = "%s"
			command_object_id = "%s"
			command_name = "%s"
			command_parameter = "%s"
			location = "%s"
		}

		data "ibm_schematics_job" "schematics_job" {
			job_id = ibm_schematics_job.schematics_job.id
		}
	`, jobCommandObject, jobCommandObjectID, jobCommandName, jobCommandParameter, jobLocation)
}
