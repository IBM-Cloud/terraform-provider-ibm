// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package project_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM/project-go-sdk/projectv1"
)

func TestAccIbmProjectBasic(t *testing.T) {
	var conf projectv1.ProjectSummary
	resourceGroup := fmt.Sprintf("Default")
	location := fmt.Sprintf("us-south")
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmProjectDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmProjectConfigBasic(resourceGroup, location, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmProjectExists("ibm_project.project_instance", conf),
					resource.TestCheckResourceAttr("ibm_project.project_instance", "resource_group", resourceGroup),
					resource.TestCheckResourceAttr("ibm_project.project_instance", "location", location),
					resource.TestCheckResourceAttr("ibm_project.project_instance", "name", name),
				),
			},
		},
	})
}

func TestAccIbmProjectAllArgs(t *testing.T) {
	var conf projectv1.ProjectSummary
	resourceGroup := fmt.Sprintf("Default")
	location := fmt.Sprintf("us-south")
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	description := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	destroyOnDelete := "true"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmProjectDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmProjectConfig(resourceGroup, location, name, description, destroyOnDelete),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmProjectExists("ibm_project.project_instance", conf),
					resource.TestCheckResourceAttr("ibm_project.project_instance", "resource_group", resourceGroup),
					resource.TestCheckResourceAttr("ibm_project.project_instance", "location", location),
					resource.TestCheckResourceAttr("ibm_project.project_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_project.project_instance", "description", description),
					resource.TestCheckResourceAttr("ibm_project.project_instance", "destroy_on_delete", destroyOnDelete),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_project.project_instance",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"location",
					"resource_group",
					"configs",
				},
			},
		},
	})
}

func testAccCheckIbmProjectConfigBasic(resourceGroup string, location string, name string) string {
	return fmt.Sprintf(`

		resource "ibm_project" "project_instance" {
			resource_group = "%s"
			location = "%s"
			name = "%s"
		}
	`, resourceGroup, location, name)
}

func testAccCheckIbmProjectConfig(resourceGroup string, location string, name string, description string, destroyOnDelete string) string {
	return fmt.Sprintf(`

		resource "ibm_project" "project_instance" {
			resource_group = "%s"
			location = "%s"
			name = "%s"
			description = "%s"
			destroy_on_delete = %s
		}
	`, resourceGroup, location, name, description, destroyOnDelete)
}

func testAccCheckIbmProjectExists(n string, obj projectv1.ProjectSummary) resource.TestCheckFunc {

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
