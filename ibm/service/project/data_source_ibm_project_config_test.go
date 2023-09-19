// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package project_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmProjectConfigDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmProjectConfigDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config", "project_id"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config", "project_config_id"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config", "version"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config", "is_draft"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config", "state"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config", "update_available"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config", "definition.#"),
				),
			},
		},
	})
}

func testAccCheckIbmProjectConfigDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		resource "ibm_project" "project_instance" {
			resource_group = "Default"
			location = "us-south"
		}

		resource "ibm_project_config" "project_config_instance" {
			project_id = ibm_project.project_instance.id
		}

		data "ibm_project_config" "project_config_instance" {
			project_id = ibm_project_config.project_config.project_id
			id = ibm_project_config.project_config_instance.projectConfig_id
		}
	`)
}
