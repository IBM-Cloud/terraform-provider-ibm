// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package project_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmProjectDataSourceBasic(t *testing.T) {
	getProjectResponseName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmProjectDataSourceConfigBasic(getProjectResponseName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_project.project", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_project.project", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_project.project", "name"),
				),
			},
		},
	})
}

func TestAccIbmProjectDataSourceAllArgs(t *testing.T) {
	getProjectResponseName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	getProjectResponseDescription := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	getProjectResponseResourceGroup := fmt.Sprintf("tf_resource_group_%d", acctest.RandIntRange(10, 100))
	getProjectResponseLocation := fmt.Sprintf("tf_location_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmProjectDataSourceConfig(getProjectResponseName, getProjectResponseDescription, getProjectResponseResourceGroup, getProjectResponseLocation),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_project.project", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_project.project", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_project.project", "exclude_configs"),
					resource.TestCheckResourceAttrSet("data.ibm_project.project", "complete"),
					resource.TestCheckResourceAttrSet("data.ibm_project.project", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_project.project", "description"),
					resource.TestCheckResourceAttrSet("data.ibm_project.project", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_project.project", "configs.#"),
					resource.TestCheckResourceAttrSet("data.ibm_project.project", "configs.0.id"),
					resource.TestCheckResourceAttr("data.ibm_project.project", "configs.0.name", getProjectResponseName),
					resource.TestCheckResourceAttr("data.ibm_project.project", "configs.0.description", getProjectResponseDescription),
					resource.TestCheckResourceAttrSet("data.ibm_project.project", "configs.0.locator_id"),
					resource.TestCheckResourceAttrSet("data.ibm_project.project", "configs.0.type"),
					resource.TestCheckResourceAttrSet("data.ibm_project.project", "metadata.#"),
				),
			},
		},
	})
}

func testAccCheckIbmProjectDataSourceConfigBasic(getProjectResponseName string) string {
	return fmt.Sprintf(`
		resource "ibm_project" "project_instance" {
			name = "%s"
		}

		data "ibm_project" "project_instance" {
			id = projectIdLink
			exclude_configs = true
			complete = true
		}
	`, getProjectResponseName)
}

func testAccCheckIbmProjectDataSourceConfig(getProjectResponseName string, getProjectResponseDescription string, getProjectResponseResourceGroup string, getProjectResponseLocation string) string {
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

		data "ibm_project" "project_instance" {
			id = projectIdLink
			exclude_configs = true
			complete = true
		}
	`, getProjectResponseName, getProjectResponseDescription, getProjectResponseResourceGroup, getProjectResponseLocation)
}
