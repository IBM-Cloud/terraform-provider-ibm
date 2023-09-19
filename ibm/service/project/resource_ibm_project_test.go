// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package project_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM/project-go-sdk/projectv1"
)

func TestAccIbmProjectBasic(t *testing.T) {
	var conf projectv1.Project
	resourceGroup := fmt.Sprintf("Default")
	location := fmt.Sprintf("us-south")
	resourceGroupUpdate := fmt.Sprintf("Default")
	locationUpdate := fmt.Sprintf("us-south")

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmProjectDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmProjectConfigBasic(resourceGroup, location),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmProjectExists("ibm_project.project_instance", conf),
					resource.TestCheckResourceAttr("ibm_project.project_instance", "resource_group", resourceGroup),
					resource.TestCheckResourceAttr("ibm_project.project_instance", "location", location),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmProjectConfigBasic(resourceGroupUpdate, locationUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_project.project_instance", "resource_group", resourceGroupUpdate),
					resource.TestCheckResourceAttr("ibm_project.project_instance", "location", locationUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_project.project_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmProjectConfigBasic(resourceGroup string, location string) string {
	return fmt.Sprintf(`
		resource "ibm_project" "project_instance" {
			resource_group = "%s"
			location = "%s"
			definition {
                name = "acme-microservice"
                description = "acme-microservice description"
                destroy_on_delete = true
            }
		}
	`, resourceGroup, location)
}

func testAccCheckIbmProjectExists(n string, obj projectv1.Project) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		projectClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).ProjectV1()
		if err != nil {
			return err
		}

		getProjectOptions := &projectv1.GetProjectOptions{}

		getProjectOptions.SetID(rs.Primary.ID)

		project, _, err := projectClient.GetProject(getProjectOptions)
		if err != nil {
			return err
		}

		obj = *project
		return nil
	}
}

func testAccCheckIbmProjectDestroy(s *terraform.State) error {
	projectClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).ProjectV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_project" {
			continue
		}

		getProjectOptions := &projectv1.GetProjectOptions{}

		getProjectOptions.SetID(rs.Primary.ID)

		// Try to find the key
		_, response, err := projectClient.GetProject(getProjectOptions)

		if err == nil {
			return fmt.Errorf("project still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for project (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
