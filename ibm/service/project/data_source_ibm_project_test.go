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
	projectCanonicalResourceGroup := fmt.Sprintf("Default")
	projectCanonicalLocation := fmt.Sprintf("us-south")
	projectCanonicalName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmProjectDataSourceConfigBasic(projectCanonicalResourceGroup, projectCanonicalLocation, projectCanonicalName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_project.project_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_project.project_instance", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_project.project_instance", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_project.project_instance", "location"),
					resource.TestCheckResourceAttrSet("data.ibm_project.project_instance", "resource_group"),
					resource.TestCheckResourceAttrSet("data.ibm_project.project_instance", "state"),
					resource.TestCheckResourceAttrSet("data.ibm_project.project_instance", "definition.#"),
					resource.TestCheckResourceAttrSet("data.ibm_project.project_instance", "definition.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_project.project_instance", "destroy_on_delete"),
				),
			},
		},
	})
}

func TestAccIbmProjectDataSourceAllArgs(t *testing.T) {
	projectCanonicalResourceGroup := fmt.Sprintf("Default")
	projectCanonicalLocation := fmt.Sprintf("us-south")
	projectCanonicalName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	projectCanonicalDescription := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	projectCanonicalDestroyOnDelete := "true"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmProjectDataSourceConfig(projectCanonicalResourceGroup, projectCanonicalLocation, projectCanonicalName, projectCanonicalDescription, projectCanonicalDestroyOnDelete),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_project.project_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_project.project_instance", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_project.project_instance", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_project.project_instance", "location"),
					resource.TestCheckResourceAttrSet("data.ibm_project.project_instance", "resource_group"),
					resource.TestCheckResourceAttrSet("data.ibm_project.project_instance", "state"),
					resource.TestCheckResourceAttrSet("data.ibm_project.project_instance", "definition.#"),
					resource.TestCheckResourceAttrSet("data.ibm_project.project_instance", "definition.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_project.project_instance", "definition.0.description"),
					resource.TestCheckResourceAttrSet("data.ibm_project.project_instance", "destroy_on_delete"),
				),
			},
		},
	})
}

func testAccCheckIbmProjectDataSourceConfigBasic(projectCanonicalResourceGroup string, projectCanonicalLocation string, projectCanonicalName string) string {
	return fmt.Sprintf(`
		resource "ibm_project" "project_instance" {
			resource_group = "%s"
			location = "%s"
			name = "%s"
		}

		data "ibm_project" "project_instance" {
			id = ibm_project.project_instance.id
		}
	`, projectCanonicalResourceGroup, projectCanonicalLocation, projectCanonicalName)
}

func testAccCheckIbmProjectDataSourceConfig(projectCanonicalResourceGroup string, projectCanonicalLocation string, projectCanonicalName string, projectCanonicalDescription string, projectCanonicalDestroyOnDelete string) string {
	return fmt.Sprintf(`
		resource "ibm_project" "project_instance" {
			resource_group = "%s"
			location = "%s"
			name = "%s"
			description = "%s"
			destroy_on_delete = %s
		}

		data "ibm_project" "project_instance" {
			id = ibm_project.project_instance.id
		}
	`, projectCanonicalResourceGroup, projectCanonicalLocation, projectCanonicalName, projectCanonicalDescription, projectCanonicalDestroyOnDelete)
}
