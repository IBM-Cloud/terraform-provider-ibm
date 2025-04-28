// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package schematics_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMSchematicsWorkspaceDataSourceBasic(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMSchematicsWorkspaceDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "workspace_id"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "created_by"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "description"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "last_health_check_at"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "location"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "resource_group"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "runtime_data.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "shared_data.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "status"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "template_type"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "template_env_settings.#"),
				),
			},
		},
	})
}

func TestAccIBMSchematicsWorkspaceDataSourceAllArgs(t *testing.T) {
	workspaceResponseDescription := fmt.Sprintf("description_%d", acctest.RandIntRange(10, 100))
	workspaceResponseLocation := "us-east"
	workspaceTemplateType := "terraform_v1.6"
	workspaceResponseName := fmt.Sprintf("name_%d", acctest.RandIntRange(10, 100))
	workspaceResponseResourceGroup := "Default"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMSchematicsWorkspaceDataSourceConfig(workspaceResponseDescription, workspaceResponseLocation, workspaceResponseName, workspaceResponseResourceGroup),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "workspace_id"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "created_by"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "crn"),
					resource.TestCheckResourceAttr("data.ibm_schematics_workspace.schematics_workspace", "description", workspaceResponseDescription),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "last_health_check_at"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "location"),
					resource.TestCheckResourceAttr("data.ibm_schematics_workspace.schematics_workspace", "name", workspaceResponseName),
					resource.TestCheckResourceAttr("data.ibm_schematics_workspace.schematics_workspace", "template_type", workspaceTemplateType),
					resource.TestCheckResourceAttr("data.ibm_schematics_workspace.schematics_workspace", "resource_group", workspaceResponseResourceGroup),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "runtime_data.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "shared_data.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "status"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "template_type"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "template_env_settings.#"),
				),
			},
		},
	})
}

func testAccCheckIBMSchematicsWorkspaceDataSourceConfigBasic() string {
	return `
		resource "ibm_schematics_workspace" "schematics_workspace" {
			description = "tf-acc-test-schematics"
			name = "tf-acc-test-schematics"
			location = "us-east"
			resource_group = "Default"
			template_type = "terraform_v1.6"
			template_env_settings = [
				{
					IBMCLOUD_ENV_VAR = "ENV_VALUE",
				}
			]
		}

		data "ibm_schematics_workspace" "schematics_workspace" {
			workspace_id = ibm_schematics_workspace.schematics_workspace.id
		}
	`
}

func testAccCheckIBMSchematicsWorkspaceDataSourceConfig(workspaceResponseDescription string, workspaceResponseLocation string, workspaceResponseName string, workspaceResponseResourceGroup string) string {
	return fmt.Sprintf(`
		 resource "ibm_schematics_workspace" "schematics_workspace" {
			 description = "%s"
			 location = "%s"
			 name = "%s"
			 resource_group = "%s"
			 template_type = "terraform_v1.6"
		 }
 
		 data "ibm_schematics_workspace" "schematics_workspace" {
			 workspace_id = ibm_schematics_workspace.schematics_workspace.id
		 }
	 `, workspaceResponseDescription, workspaceResponseLocation, workspaceResponseName, workspaceResponseResourceGroup)
}
