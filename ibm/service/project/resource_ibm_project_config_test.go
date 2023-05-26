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
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/project-go-sdk/projectv1"
)

func TestAccIbmProjectConfigBasic(t *testing.T) {
	var conf projectv1.ProjectConfigDraft
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	locatorID := fmt.Sprintf("1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc.cd596f95-95a2-4f21-9b84-477f21fd1e95-global")
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmProjectConfigDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmProjectConfigConfigBasic(name, locatorID),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmProjectConfigExists("ibm_project_config.project_config_instance", conf),
					resource.TestCheckResourceAttr("ibm_project_config.project_config_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_project_config.project_config_instance", "locator_id", locatorID),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmProjectConfigConfigBasic(nameUpdate, locatorID),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_project_config.project_config_instance", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_project_config.project_config_instance", "locator_id", locatorID),
				),
			},
		},
	})
}

func TestAccIbmProjectConfigAllArgs(t *testing.T) {
	var conf projectv1.ProjectConfigDraft
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	locatorID := fmt.Sprintf("1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc.cd596f95-95a2-4f21-9b84-477f21fd1e95-global")
	description := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmProjectConfigDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmProjectConfigConfig(name, locatorID, description),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmProjectConfigExists("ibm_project_config.project_config_instance", conf),
					resource.TestCheckResourceAttr("ibm_project_config.project_config_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_project_config.project_config_instance", "locator_id", locatorID),
					resource.TestCheckResourceAttr("ibm_project_config.project_config_instance", "description", description),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmProjectConfigConfig(nameUpdate, locatorID, description),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_project_config.project_config_instance", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_project_config.project_config_instance", "locator_id", locatorID),
					resource.TestCheckResourceAttr("ibm_project_config.project_config_instance", "description", description),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_project_config.project_config_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmProjectConfigConfigBasic(name string, locatorID string) string {
	return fmt.Sprintf(`
		resource "ibm_project" "project_instance" {
			resource_group = "Default"
			location = "us-south"
			name = "acme-microservice-3"
		}

		resource "ibm_project_config" "project_config_instance" {
			project_id = ibm_project.project_instance.id
			name = "%s"
			locator_id = "%s"
		}
	`, name, locatorID)
}

func testAccCheckIbmProjectConfigConfig(name string, locatorID string, description string) string {
	return fmt.Sprintf(`

		resource "ibm_project" "project_instance" {
			resource_group = "Default"
			location = "us-south"
			name = "acme-microservice-4"
		}

		resource "ibm_project_config" "project_config_instance" {
			project_id = ibm_project.project_instance.id
			name = "%s"
			locator_id = "%s"
			labels = [ "labels" ]
            description = "%s"
            authorizations {
                method = "API_KEY"
                api_key = "xxx"
            }
			input {
				name = "name"
				value = "anything as a string"
			}
			setting {
				name = "name"
				value = "value"
			}
		}
	`, name, locatorID, description)
}

func testAccCheckIbmProjectConfigExists(n string, obj projectv1.ProjectConfigDraft) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		projectClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).ProjectV1()
		if err != nil {
			return err
		}

		getConfigOptions := &projectv1.GetConfigOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getConfigOptions.SetProjectID(parts[0])
		getConfigOptions.SetID(parts[1])
		getConfigOptions.SetVersion("draft")

		_, response, err := projectClient.GetConfig(getConfigOptions)
		if err != nil {
			return err
		} else if response.StatusCode != 200 {
			return fmt.Errorf("Error checking for project_config (%s) exists: %s", rs.Primary.ID, err)
		}

		return nil
	}
}

func testAccCheckIbmProjectConfigDestroy(s *terraform.State) error {
	projectClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).ProjectV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_project_config" {
			continue
		}

		getConfigOptions := &projectv1.GetConfigOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getConfigOptions.SetProjectID(parts[0])
		getConfigOptions.SetID(parts[1])

		// Try to find the key
		_, response, err := projectClient.GetConfig(getConfigOptions)

		if err == nil {
			return fmt.Errorf("project_config still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for project_config (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
