// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package codeengine_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM/code-engine-go-sdk/codeenginev2"
)

func TestAccIbmCodeEngineProjectBasic(t *testing.T) {
	var conf codeenginev2.Project
	projectName := fmt.Sprintf("tf-project-basic-%d", acctest.RandIntRange(10, 100))
	resourceGroupID := acc.CeResourceGroupID

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmCodeEngineProjectDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmCodeEngineProjectConfig(projectName, resourceGroupID),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmCodeEngineProjectExists("ibm_code_engine_project.code_engine_project_instance", conf),
					resource.TestCheckResourceAttrSet("ibm_code_engine_project.code_engine_project_instance", "project_id"),
					resource.TestCheckResourceAttrSet("ibm_code_engine_project.code_engine_project_instance", "account_id"),
					resource.TestCheckResourceAttrSet("ibm_code_engine_project.code_engine_project_instance", "created_at"),
					resource.TestCheckResourceAttrSet("ibm_code_engine_project.code_engine_project_instance", "crn"),
					resource.TestCheckResourceAttrSet("ibm_code_engine_project.code_engine_project_instance", "href"),
					resource.TestCheckResourceAttrSet("ibm_code_engine_project.code_engine_project_instance", "region"),
					resource.TestCheckResourceAttrSet("ibm_code_engine_project.code_engine_project_instance", "status"),
					resource.TestCheckResourceAttr("ibm_code_engine_project.code_engine_project_instance", "name", projectName),
					resource.TestCheckResourceAttr("ibm_code_engine_project.code_engine_project_instance", "resource_group_id", resourceGroupID),
					resource.TestCheckResourceAttr("ibm_code_engine_project.code_engine_project_instance", "resource_type", "project_v2"),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_code_engine_project.code_engine_project_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmCodeEngineProjectConfig(projectName string, resourceGroupID string) string {
	return fmt.Sprintf(`
		resource "ibm_code_engine_project" "code_engine_project_instance" {
			name = "%s"
			resource_group_id = "%s"
		}
	`, projectName, resourceGroupID)
}

func testAccCheckIbmCodeEngineProjectExists(n string, obj codeenginev2.Project) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		codeEngineClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CodeEngineV2()
		if err != nil {
			return err
		}

		getProjectOptions := &codeenginev2.GetProjectOptions{}

		getProjectOptions.SetID(rs.Primary.ID)

		project, _, err := codeEngineClient.GetProject(getProjectOptions)
		if err != nil {
			return err
		}

		obj = *project
		return nil
	}
}

func testAccCheckIbmCodeEngineProjectDestroy(s *terraform.State) error {
	codeEngineClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CodeEngineV2()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_code_engine_project" {
			continue
		}

		getProjectOptions := &codeenginev2.GetProjectOptions{}

		getProjectOptions.SetID(rs.Primary.ID)

		// Try to find the key
		res, response, err := codeEngineClient.GetProject(getProjectOptions)

		if *res.Status != "soft_deleted" {
			return fmt.Errorf("code_engine_project `%s` hasn't changed to correct status: '%s'", rs.Primary.ID, *res.Status)
		} else if err != nil {
			return fmt.Errorf("An error occured during clean up: '%s'", err)
		} else if response.StatusCode != 200 {
			return fmt.Errorf("Error checking for code_engine_project ('%s') has been destroyed", rs.Primary.ID)
		}
	}

	return nil
}
