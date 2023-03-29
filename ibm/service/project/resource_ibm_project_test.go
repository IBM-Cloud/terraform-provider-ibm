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
	"github.com/damianovesperini/platform-services-go-sdk/projectv1"
)

func TestAccIbmProjectBasic(t *testing.T) {
	var conf projectv1.GetProjectResponse
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmProjectDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmProjectConfigBasic(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmProjectExists("ibm_project.project", conf),
					resource.TestCheckResourceAttr("ibm_project.project", "name", name),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmProjectConfigBasic(nameUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_project.project", "name", nameUpdate),
				),
			},
		},
	})
}

func TestAccIbmProjectAllArgs(t *testing.T) {
	var conf projectv1.GetProjectResponse
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	description := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	resourceGroup := fmt.Sprintf("tf_resource_group_%d", acctest.RandIntRange(10, 100))
	location := fmt.Sprintf("tf_location_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	descriptionUpdate := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	resourceGroupUpdate := fmt.Sprintf("tf_resource_group_%d", acctest.RandIntRange(10, 100))
	locationUpdate := fmt.Sprintf("tf_location_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmProjectDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmProjectConfig(name, description, resourceGroup, location),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmProjectExists("ibm_project.project", conf),
					resource.TestCheckResourceAttr("ibm_project.project", "name", name),
					resource.TestCheckResourceAttr("ibm_project.project", "description", description),
					resource.TestCheckResourceAttr("ibm_project.project", "resource_group", resourceGroup),
					resource.TestCheckResourceAttr("ibm_project.project", "location", location),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmProjectConfig(nameUpdate, descriptionUpdate, resourceGroupUpdate, locationUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_project.project", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_project.project", "description", descriptionUpdate),
					resource.TestCheckResourceAttr("ibm_project.project", "resource_group", resourceGroupUpdate),
					resource.TestCheckResourceAttr("ibm_project.project", "location", locationUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_project.project",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmProjectConfigBasic(name string) string {
	return fmt.Sprintf(`

		resource "ibm_project" "project_instance" {
			name = "%s"
		}
	`, name)
}

func testAccCheckIbmProjectConfig(name string, description string, resourceGroup string, location string) string {
	return fmt.Sprintf(`

		resource "ibm_project" "project_instance" {
			name = "%s"
			description = "%s"
			configs {
				id = "id"
				name = "name"
				labels = [ "labels" ]
				description = "description"
				locator_id = "locator_id"
				input {
					name = "name"
				}
				setting {
					name = "name"
					value = "value"
				}
			}
			resource_group = "%s"
			location = "%s"
		}
	`, name, description, resourceGroup, location)
}

func testAccCheckIbmProjectExists(n string, obj projectv1.GetProjectResponse) resource.TestCheckFunc {

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

		getProjectResponse, _, err := projectClient.GetProject(getProjectOptions)
		if err != nil {
			return err
		}

		obj = *getProjectResponse
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
