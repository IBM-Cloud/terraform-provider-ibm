// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package codeengine_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/IBM/code-engine-go-sdk/codeenginev2"
	"github.com/IBM/go-sdk-core/v5/core"
	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"
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

func RetrieveProjectIdByName(projectName string, apiKey string, apiEndpoint string) (projectId *string, err error) {
	iamEndpoint := determineEnvironment(apiEndpoint)

	codeEngineService, err := codeenginev2.NewCodeEngineV2(&codeenginev2.CodeEngineV2Options{
		Authenticator: &core.IamAuthenticator{
			ApiKey:       apiKey,
			ClientId:     "bx",
			ClientSecret: "bx",
			URL:          iamEndpoint,
		},
		URL: apiEndpoint,
	})

	if err != nil {
		return
	}

	limit := int64(100)
	listProjectsOptions := &codeenginev2.ListProjectsOptions{
		Limit: &limit,
	}
	pager, err := codeEngineService.NewProjectsPager(listProjectsOptions)
	if err != nil {
		panic(err)
	}

	var allResults []codeenginev2.Project
	for pager.HasNext() {
		nextPage, err := pager.GetNext()
		if err != nil {
			panic(err)
		}
		allResults = append(allResults, nextPage...)
	}

	for _, project := range allResults {
		if project.Name == &projectName {
			if *project.Status == "soft_delete" {
				err = fmt.Errorf("Error project '%s' is in 'soft_delete' status, please clean it up first", projectName)
				break
			}

			projectId = project.ID
			break
		}
	}

	// if projectId == nil {
	// 	createProjectOptions := codeenginev2.CreateProjectOptions{
	// 		Name: &projectName,
	// 	}

	// 	createdProject, res, err := codeEngineService.CreateProject(&createProjectOptions)

	// 	if err != nil {
	// 		err = fmt.Errorf("Error created project '%s': '%s'", projectName, err)
	// 	} else if res.StatusCode != 202 {
	// 		err = fmt.Errorf("Error created project '%s'", projectName)
	// 	} else if createdProject == nil {
	// 		err = fmt.Errorf("Error created project '%s'", projectName)
	// 	}

	// 	createdProjectId := string(*createdProject.ID)

	// 	println("!!!createdProject.ID: ", createdProjectId)
	// 	attempts := 25
	// 	getProjectOptions := codeenginev2.GetProjectOptions{
	// 		ID: createdProject.ID,
	// 	}
	// 	for i := 0; i < attempts; i++ {
	// 		retrievedProject, res, err := codeEngineService.GetProject(&getProjectOptions)
	// 		if err != nil || res.StatusCode != 200 || *retrievedProject.Status != "active" {
	// 			log.Println("Project not ready yet, waiting....", err)
	// 			time.Sleep(5 * time.Second)
	// 		} else {
	// 			projectId = retrievedProject.ID
	// 			break
	// 		}
	// 	}
	// }

	return
}

func determineEnvironment(apiEndpoint string) string {
	if strings.Contains(apiEndpoint, "codeengine.test.cloud") || strings.Contains(apiEndpoint, "codeengine.dev.cloud") {
		return "https://iam.test.cloud.ibm.com"
	} else {
		return "https://iam.cloud.ibm.com"
	}
}
