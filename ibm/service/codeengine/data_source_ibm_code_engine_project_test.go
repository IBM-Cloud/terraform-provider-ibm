// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package codeengine_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmCodeEngineProjectDataSourceBasic(t *testing.T) {
	projectName := fmt.Sprintf("tf-project-data-basic-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmCodeEngineProjectDataSourceConfigBasic(projectName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_code_engine_project.code_engine_project", "resource_group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_code_engine_project.code_engine_project", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_code_engine_project.code_engine_project", "account_id"),
					resource.TestCheckResourceAttrSet("data.ibm_code_engine_project.code_engine_project", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_code_engine_project.code_engine_project", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_code_engine_project.code_engine_project", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_code_engine_project.code_engine_project", "region"),
					resource.TestCheckResourceAttrSet("data.ibm_code_engine_project.code_engine_project", "status"),
					resource.TestCheckResourceAttr("data.ibm_code_engine_project.code_engine_project", "name", projectName),
					resource.TestCheckResourceAttr("data.ibm_code_engine_project.code_engine_project", "resource_type", "project_v2"),
				),
			},
		},
	})
}

func testAccCheckIbmCodeEngineProjectDataSourceConfigBasic(projectName string) string {
	return fmt.Sprintf(`
		resource "ibm_code_engine_project" "code_engine_project_instance" {
			name = "%s"
		}

		data "ibm_code_engine_project" "code_engine_project_instance" {
			id = ibm_code_engine_project.code_engine_project_instance.id
		}
	`, projectName)
}
