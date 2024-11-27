// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package codeengine_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmCodeEngineProjectDataSourceBasic(t *testing.T) {
	projectID := acc.CeProjectId

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmCodeEngineProjectDataSourceConfigBasic(projectID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_code_engine_project.code_engine_project_instance", "resource_group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_code_engine_project.code_engine_project_instance", "project_id"),
					resource.TestCheckResourceAttrSet("data.ibm_code_engine_project.code_engine_project_instance", "account_id"),
					resource.TestCheckResourceAttrSet("data.ibm_code_engine_project.code_engine_project_instance", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_code_engine_project.code_engine_project_instance", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_code_engine_project.code_engine_project_instance", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_code_engine_project.code_engine_project_instance", "region"),
					resource.TestCheckResourceAttrSet("data.ibm_code_engine_project.code_engine_project_instance", "status"),
					resource.TestCheckResourceAttr("data.ibm_code_engine_project.code_engine_project_instance", "id", projectID),
					resource.TestCheckResourceAttr("data.ibm_code_engine_project.code_engine_project_instance", "resource_type", "project_v2"),
				),
			},
		},
	})
}

func testAccCheckIbmCodeEngineProjectDataSourceConfigBasic(projectID string) string {
	return fmt.Sprintf(`
		data "ibm_code_engine_project" "code_engine_project_instance" {
			project_id = "%s"
		}
	`, projectID)
}
