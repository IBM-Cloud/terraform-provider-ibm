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

func TestAccIbmProjectInstanceBasic(t *testing.T) {
	var conf projectv1.Project
	resourceGroup := fmt.Sprintf("tf_resource_group_%d", acctest.RandIntRange(10, 100))
	location := fmt.Sprintf("tf_location_%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	resourceGroupUpdate := fmt.Sprintf("tf_resource_group_%d", acctest.RandIntRange(10, 100))
	locationUpdate := fmt.Sprintf("tf_location_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmProjectInstanceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmProjectInstanceConfigBasic(resourceGroup, location, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmProjectInstanceExists("ibm_project_instance.project_instance", conf),
					resource.TestCheckResourceAttr("ibm_project_instance.project_instance", "resource_group", resourceGroup),
					resource.TestCheckResourceAttr("ibm_project_instance.project_instance", "location", location),
					resource.TestCheckResourceAttr("ibm_project_instance.project_instance", "name", name),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmProjectInstanceConfigBasic(resourceGroupUpdate, locationUpdate, nameUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_project_instance.project_instance", "resource_group", resourceGroupUpdate),
					resource.TestCheckResourceAttr("ibm_project_instance.project_instance", "location", locationUpdate),
					resource.TestCheckResourceAttr("ibm_project_instance.project_instance", "name", nameUpdate),
				),
			},
		},
	})
}

func TestAccIbmProjectInstanceAllArgs(t *testing.T) {
	var conf projectv1.Project
	resourceGroup := fmt.Sprintf("tf_resource_group_%d", acctest.RandIntRange(10, 100))
	location := fmt.Sprintf("tf_location_%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	description := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	resourceGroupUpdate := fmt.Sprintf("tf_resource_group_%d", acctest.RandIntRange(10, 100))
	locationUpdate := fmt.Sprintf("tf_location_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	descriptionUpdate := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmProjectInstanceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmProjectInstanceConfig(resourceGroup, location, name, description),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmProjectInstanceExists("ibm_project_instance.project_instance", conf),
					resource.TestCheckResourceAttr("ibm_project_instance.project_instance", "resource_group", resourceGroup),
					resource.TestCheckResourceAttr("ibm_project_instance.project_instance", "location", location),
					resource.TestCheckResourceAttr("ibm_project_instance.project_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_project_instance.project_instance", "description", description),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmProjectInstanceConfig(resourceGroupUpdate, locationUpdate, nameUpdate, descriptionUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_project_instance.project_instance", "resource_group", resourceGroupUpdate),
					resource.TestCheckResourceAttr("ibm_project_instance.project_instance", "location", locationUpdate),
					resource.TestCheckResourceAttr("ibm_project_instance.project_instance", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_project_instance.project_instance", "description", descriptionUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_project_instance.project_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmProjectInstanceConfigBasic(resourceGroup string, location string, name string) string {
	return fmt.Sprintf(`

		resource "ibm_project_instance" "project_instance_instance" {
			resource_group = "%s"
			location = "%s"
			name = "%s"
		}
	`, resourceGroup, location, name)
}

func testAccCheckIbmProjectInstanceConfig(resourceGroup string, location string, name string, description string) string {
	return fmt.Sprintf(`

		resource "ibm_project_instance" "project_instance_instance" {
			resource_group = "%s"
			location = "%s"
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
					value = "anything as a string"
				}
				setting {
					name = "name"
					value = "value"
				}
			}
		}
	`, resourceGroup, location, name, description)
}

func testAccCheckIbmProjectInstanceExists(n string, obj projectv1.Project) resource.TestCheckFunc {

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

func testAccCheckIbmProjectInstanceDestroy(s *terraform.State) error {
	projectClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).ProjectV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_project_instance" {
			continue
		}

		getProjectOptions := &projectv1.GetProjectOptions{}

		getProjectOptions.SetID(rs.Primary.ID)

		// Try to find the key
		_, response, err := projectClient.GetProject(getProjectOptions)

		if err == nil {
			return fmt.Errorf("Project definition still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for Project definition (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
