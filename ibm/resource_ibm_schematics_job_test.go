// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/IBM/schematics-go-sdk/schematicsv1"
)

func TestAccIBMSchematicsJobBasic(t *testing.T) {
	var conf schematicsv1.Job
	commandObject := "action"
	commandObjectID := actionID
	commandName := "ansible_playbook_run"
	commandParameter := "ssh_user.yml"
	commandObjectUpdate := "action"
	commandObjectIDUpdate := actionID
	commandNameUpdate := "ansible_playbook_run"
	commandParameterUpdate := "ssh_user_updated.yml"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMSchematicsJobDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMSchematicsJobConfig(commandObject, commandObjectID, commandName, commandParameter),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMSchematicsJobExists("ibm_schematics_job.schematics_job", conf),
					resource.TestCheckResourceAttr("ibm_schematics_job.schematics_job", "command_object", commandObject),
					resource.TestCheckResourceAttr("ibm_schematics_job.schematics_job", "command_object_id", commandObjectID),
					resource.TestCheckResourceAttr("ibm_schematics_job.schematics_job", "command_name", commandName),
					resource.TestCheckResourceAttr("ibm_schematics_job.schematics_job", "command_parameter", commandParameter),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMSchematicsJobConfig(commandObjectUpdate, commandObjectIDUpdate, commandNameUpdate, commandParameterUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_schematics_job.schematics_job", "command_object", commandObjectUpdate),
					resource.TestCheckResourceAttr("ibm_schematics_job.schematics_job", "command_object_id", commandObjectIDUpdate),
					resource.TestCheckResourceAttr("ibm_schematics_job.schematics_job", "command_name", commandNameUpdate),
					resource.TestCheckResourceAttr("ibm_schematics_job.schematics_job", "command_parameter", commandParameterUpdate),
				),
			},
		},
	})
}

func testAccCheckIBMSchematicsJobConfig(commandObject string, commandObjectID string, commandName string, commandParameter string) string {
	return fmt.Sprintf(`

		resource "ibm_schematics_job" "schematics_job" {
			command_object = "%s"
			command_object_id = "%s"
			command_name = "%s"
			command_parameter = "%s"
			location = "us"
		}
	`, commandObject, commandObjectID, commandName, commandParameter)
}

func testAccCheckIBMSchematicsJobExists(n string, obj schematicsv1.Job) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		schematicsClient, err := testAccProvider.Meta().(ClientSession).SchematicsV1()
		if err != nil {
			return err
		}

		getJobOptions := &schematicsv1.GetJobOptions{}

		getJobOptions.SetJobID(rs.Primary.ID)

		job, _, err := schematicsClient.GetJob(getJobOptions)
		if err != nil {
			return err
		}

		obj = *job
		return nil
	}
}

func testAccCheckIBMSchematicsJobDestroy(s *terraform.State) error {
	schematicsClient, err := testAccProvider.Meta().(ClientSession).SchematicsV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_schematics_job" {
			continue
		}

		getJobOptions := &schematicsv1.GetJobOptions{}

		getJobOptions.SetJobID(rs.Primary.ID)

		// Try to find the key
		_, response, err := schematicsClient.GetJob(getJobOptions)

		if err == nil {
			return fmt.Errorf("schematics_job still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for schematics_job (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
