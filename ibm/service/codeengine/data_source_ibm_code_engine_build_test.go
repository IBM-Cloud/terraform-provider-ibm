// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package codeengine_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmCodeEngineBuildDataSourceBasic(t *testing.T) {
	buildName := fmt.Sprintf("tf-data-build-basic-%d", acctest.RandIntRange(10, 1000))
	buildOutputImage := fmt.Sprintf("private.us.icr.io/ce-terraform-test/%s", buildName)
	buildOutputSecret := "ce-terraform-test"
	buildSourceURL := "https://github.com/IBM/CodeEngine"
	buildStrategyType := "dockerfile"

	projectID := acc.CeProjectId

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmCodeEngineBuildDataSourceConfigBasic(projectID, buildName, buildOutputImage, buildOutputSecret, buildSourceURL, buildStrategyType),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_code_engine_build.code_engine_build_instance", "build_id"),
					resource.TestCheckResourceAttr("data.ibm_code_engine_build.code_engine_build_instance", "project_id", projectID),
					resource.TestCheckResourceAttr("data.ibm_code_engine_build.code_engine_build_instance", "name", buildName),
					resource.TestCheckResourceAttr("data.ibm_code_engine_build.code_engine_build_instance", "output_image", buildOutputImage),
					resource.TestCheckResourceAttr("data.ibm_code_engine_build.code_engine_build_instance", "output_secret", buildOutputSecret),
					resource.TestCheckResourceAttr("data.ibm_code_engine_build.code_engine_build_instance", "source_url", buildSourceURL),
					resource.TestCheckResourceAttr("data.ibm_code_engine_build.code_engine_build_instance", "strategy_type", buildStrategyType),
				),
			},
		},
	})
}

func testAccCheckIbmCodeEngineBuildDataSourceConfigBasic(projectID string, buildName string, buildOutputImage string, buildOutputSecret string, buildSourceURL string, buildStrategyType string) string {
	return fmt.Sprintf(`
		data "ibm_code_engine_project" "code_engine_project_instance" {
			project_id = "%s"
		}

		resource "ibm_code_engine_build" "code_engine_build_instance" {
			project_id = data.ibm_code_engine_project.code_engine_project_instance.project_id
			name = "%s"
			output_image = "%s"
			output_secret = "%s"
			source_url = "%s"
			strategy_type = "%s"
		}

		data "ibm_code_engine_build" "code_engine_build_instance" {
			project_id = ibm_code_engine_build.code_engine_build_instance.project_id
			name = ibm_code_engine_build.code_engine_build_instance.name
		}
	`, projectID, buildName, buildOutputImage, buildOutputSecret, buildSourceURL, buildStrategyType)
}
