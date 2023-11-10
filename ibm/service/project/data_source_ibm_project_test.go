// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package project_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmProjectDataSourceBasic(t *testing.T) {
	projectLocation := fmt.Sprintf("us-south")
	projectResourceGroup := fmt.Sprintf("Default")

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmProjectDataSourceConfigBasic(projectLocation, projectResourceGroup),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_project.project_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_project.project_instance", "project_id"),
					resource.TestCheckResourceAttrSet("data.ibm_project.project_instance", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_project.project_instance", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_project.project_instance", "location"),
					resource.TestCheckResourceAttrSet("data.ibm_project.project_instance", "resource_group"),
					resource.TestCheckResourceAttrSet("data.ibm_project.project_instance", "state"),
					resource.TestCheckResourceAttrSet("data.ibm_project.project_instance", "definition.#"),
				),
			},
		},
	})
}

func testAccCheckIbmProjectDataSourceConfigBasic(projectLocation string, projectResourceGroup string) string {
	return fmt.Sprintf(`
		resource "ibm_project" "project_instance" {
			location = "%s"
			resource_group = "%s"
			definition {
                name = "acme-microservice"
                description = "acme-microservice description"
                destroy_on_delete = true
            }
		}

		data "ibm_project" "project_instance" {
			project_id = ibm_project.project_instance.id
		}
	`, projectLocation, projectResourceGroup)
}
