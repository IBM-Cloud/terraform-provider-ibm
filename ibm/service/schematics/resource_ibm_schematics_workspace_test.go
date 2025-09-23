// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package schematics_test

import (
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/IBM/schematics-go-sdk/schematicsv1"
)

func TestAccIBMSchematicsWorkspaceBasic(t *testing.T) {
	var conf schematicsv1.WorkspaceResponse

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMSchematicsWorkspaceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMSchematicsWorkspaceConfigBasic(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMSchematicsWorkspaceExists("ibm_schematics_workspace.schematics_workspace", conf),
				),
			},
		},
	})
}

func TestAccIBMSchematicsWorkspaceAllArgs(t *testing.T) {
	var conf schematicsv1.WorkspaceResponse
	description := fmt.Sprintf("tf-acc-test-schematics-all-args_%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-acc-test-schematics_%d", acctest.RandIntRange(10, 100))
	templateType := "terraform_v1.6"

	descriptionUpdate := fmt.Sprintf("description_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf-acc-test-schematics_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMSchematicsWorkspaceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMSchematicsWorkspaceConfig(description, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMSchematicsWorkspaceExists("ibm_schematics_workspace.schematics_workspace", conf),
					resource.TestCheckResourceAttr("ibm_schematics_workspace.schematics_workspace", "description", description),
					resource.TestCheckResourceAttr("ibm_schematics_workspace.schematics_workspace", "name", name),
					resource.TestCheckResourceAttr("ibm_schematics_workspace.schematics_workspace", "template_type", templateType),
					resource.TestCheckResourceAttrSet("ibm_schematics_workspace.schematics_workspace", "created_at"),
					resource.TestCheckResourceAttrSet("ibm_schematics_workspace.schematics_workspace", "created_by"),
					resource.TestCheckResourceAttrSet("ibm_schematics_workspace.schematics_workspace", "crn"),
					resource.TestCheckResourceAttrSet("ibm_schematics_workspace.schematics_workspace", "runtime_data.#"),
				),
			},
			{
				Config: testAccCheckIBMSchematicsWorkspaceConfigUpdate(descriptionUpdate, nameUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_schematics_workspace.schematics_workspace", "description", descriptionUpdate),
					resource.TestCheckResourceAttr("ibm_schematics_workspace.schematics_workspace", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_schematics_workspace.schematics_workspace", "template_type", templateType),
					resource.TestCheckResourceAttrSet("ibm_schematics_workspace.schematics_workspace", "created_at"),
					resource.TestCheckResourceAttrSet("ibm_schematics_workspace.schematics_workspace", "created_by"),
					resource.TestCheckResourceAttrSet("ibm_schematics_workspace.schematics_workspace", "crn"),
					resource.TestCheckResourceAttrSet("ibm_schematics_workspace.schematics_workspace", "runtime_data.#"),
				),
			},
			{
				Config: testAccCheckIBMSchematicsWorkspaceConfigUpdateRepoURL(descriptionUpdate, nameUpdate, acc.RepoURL),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMSchematicsWorkspaceExists("ibm_schematics_workspace.schematics_workspace", conf),
					resource.TestCheckResourceAttr("ibm_schematics_workspace.schematics_workspace", "description", descriptionUpdate),
					resource.TestCheckResourceAttr("ibm_schematics_workspace.schematics_workspace", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_schematics_workspace.schematics_workspace", "template_type", templateType),
					resource.TestCheckResourceAttrSet("ibm_schematics_workspace.schematics_workspace", "created_at"),
					resource.TestCheckResourceAttrSet("ibm_schematics_workspace.schematics_workspace", "created_by"),
					resource.TestCheckResourceAttrSet("ibm_schematics_workspace.schematics_workspace", "crn"),
					resource.TestCheckResourceAttrSet("ibm_schematics_workspace.schematics_workspace", "runtime_data.#"),
				),
			},
			{
				Config: testAccCheckIBMSchematicsWorkspaceConfigUpdateRepoBranch(descriptionUpdate, nameUpdate, acc.RepoURL, acc.RepoBranch),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMSchematicsWorkspaceExists("ibm_schematics_workspace.schematics_workspace", conf),
					resource.TestCheckResourceAttr("ibm_schematics_workspace.schematics_workspace", "description", descriptionUpdate),
					resource.TestCheckResourceAttr("ibm_schematics_workspace.schematics_workspace", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_schematics_workspace.schematics_workspace", "template_type", templateType),
					resource.TestCheckResourceAttrSet("ibm_schematics_workspace.schematics_workspace", "created_at"),
					resource.TestCheckResourceAttrSet("ibm_schematics_workspace.schematics_workspace", "created_by"),
					resource.TestCheckResourceAttrSet("ibm_schematics_workspace.schematics_workspace", "crn"),
					resource.TestCheckResourceAttrSet("ibm_schematics_workspace.schematics_workspace", "runtime_data.#"),
				),
			},
			{
				ResourceName:      "ibm_schematics_workspace.schematics_workspace",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMSchematicsWorkspaceConfigBasic() string {
	return `

		resource "ibm_schematics_workspace" "schematics_workspace" {
			description = "tf-acc-test-schematics"
			name = "tf-acc-test-schematics"
			location = "us-east"
			resource_group = "Default"
			template_type = "terraform_v1.6"
		}
	`
}

func testAccCheckIBMSchematicsWorkspaceConfig(description string, name string) string {
	return fmt.Sprintf(`

		resource "ibm_schematics_workspace" "schematics_workspace" {
			description = "%s"
			location = "us-east"
			name = "%s"
			resource_group = "Default"
			template_type = "terraform_v1.6"
		}
	`, description, name)
}

func testAccCheckIBMSchematicsWorkspaceConfigUpdate(description string, name string) string {
	return fmt.Sprintf(`

		resource "ibm_schematics_workspace" "schematics_workspace" {
			description = "%s"
			location = "us-east"
			name = "%s"
			resource_group = "Default"
			template_type = "terraform_v1.6"
			template_inputs {
				name = "testinput"
				value = "test"
				type = "string"
			  }
		}
	`, description, name)
}

func testAccCheckIBMSchematicsWorkspaceConfigUpdateRepoURL(description string, name string, repoURL string) string {
	return fmt.Sprintf(`

		resource "ibm_schematics_workspace" "schematics_workspace" {
			description = "%s"
			location = "us-east"
			name = "%s"
			resource_group = "Default"
			template_type = "terraform_v1.6"
			template_git_url = "%s"
			template_inputs {
				name = "testinput"
				value = "test"
				type = "string"
			  }
		}
	`, description, name, repoURL)
}

func testAccCheckIBMSchematicsWorkspaceConfigUpdateRepoBranch(description string, name string, repoURL string, repoBranch string) string {
	return fmt.Sprintf(`

		resource "ibm_schematics_workspace" "schematics_workspace" {
			description = "%s"
			location = "us-east"
			name = "%s"
			resource_group = "Default"
			template_type = "terraform_v1.6"
			template_git_url = "%s"
			template_inputs {
				name = "testinput"
				value = "test"
				type = "string"
			  }
			template_git_branch = "%s"
		}
	`, description, name, repoURL, repoBranch)
}

func testAccCheckIBMSchematicsWorkspaceExists(n string, obj schematicsv1.WorkspaceResponse) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		schematicsClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SchematicsV1()
		if err != nil {
			return err
		}

		getWorkspaceOptions := &schematicsv1.GetWorkspaceOptions{}

		getWorkspaceOptions.SetWID(rs.Primary.ID)

		workspaceResponse, _, err := schematicsClient.GetWorkspace(getWorkspaceOptions)
		if err != nil {
			return err
		}

		obj = *workspaceResponse
		return nil
	}
}

func testAccCheckIBMSchematicsWorkspaceDestroy(s *terraform.State) error {
	schematicsClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SchematicsV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_schematics_workspace" {
			continue
		}

		getWorkspaceOptions := &schematicsv1.GetWorkspaceOptions{}

		getWorkspaceOptions.SetWID(rs.Primary.ID)

		// Try to find the key
		_, response, err := schematicsClient.GetWorkspace(getWorkspaceOptions)

		if err == nil {
			return fmt.Errorf("schematics_workspace still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("[ERROR] Error checking for schematics_workspace (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
