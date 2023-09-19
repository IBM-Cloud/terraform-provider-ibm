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
	projectResourceGroup := fmt.Sprintf("tf_resource_group_%d", acctest.RandIntRange(10, 100))
	projectLocation := fmt.Sprintf("tf_location_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmProjectDataSourceConfigBasic(projectResourceGroup, projectLocation),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_project.project", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_project.project", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_project.project", "definition.#"),
				),
			},
		},
	})
}

func testAccCheckIbmProjectDataSourceConfigBasic(projectResourceGroup string, projectLocation string) string {
	return fmt.Sprintf(`
		resource "ibm_project" "project_instance" {
			resource_group = "%s"
			location = "%s"
		}

		data "ibm_project" "project_instance" {
			id = ibm_project.project_instance.project_id
		}
	`, projectResourceGroup, projectLocation)
}
