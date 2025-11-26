// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package codeengine_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/code-engine-go-sdk/codeenginev2"
)

func TestAccIbmCodeEngineBuildBasic(t *testing.T) {
	var conf codeenginev2.Build
	name := fmt.Sprintf("tf-build-basic-%d", acctest.RandIntRange(10, 1000))
	outputImage := fmt.Sprintf("private.us.icr.io/ce-terraform-test/%s", name)
	outputSecret := "ce-terraform-test"
	sourceURL := "https://github.com/IBM/CodeEngine"
	strategyType := "dockerfile"

	projectID := acc.CeProjectId

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmCodeEngineBuildDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmCodeEngineBuildConfigBasic(projectID, name, outputImage, outputSecret, sourceURL, strategyType),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmCodeEngineBuildExists("ibm_code_engine_build.code_engine_build_instance", conf),
					resource.TestCheckResourceAttrSet("ibm_code_engine_build.code_engine_build_instance", "build_id"),
					resource.TestCheckResourceAttr("ibm_code_engine_build.code_engine_build_instance", "project_id", projectID),
					resource.TestCheckResourceAttr("ibm_code_engine_build.code_engine_build_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_code_engine_build.code_engine_build_instance", "output_image", outputImage),
					resource.TestCheckResourceAttr("ibm_code_engine_build.code_engine_build_instance", "output_secret", outputSecret),
					resource.TestCheckResourceAttr("ibm_code_engine_build.code_engine_build_instance", "source_url", sourceURL),
					resource.TestCheckResourceAttr("ibm_code_engine_build.code_engine_build_instance", "strategy_type", strategyType),
				),
			},
		},
	})
}

func testAccCheckIbmCodeEngineBuildConfigBasic(projectID string, name string, outputImage string, outputSecret string, sourceURL string, strategyType string) string {
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
	`, projectID, name, outputImage, outputSecret, sourceURL, strategyType)
}

func testAccCheckIbmCodeEngineBuildExists(n string, obj codeenginev2.Build) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		codeEngineClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CodeEngineV2()
		if err != nil {
			return err
		}

		getBuildOptions := &codeenginev2.GetBuildOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getBuildOptions.SetProjectID(parts[0])
		getBuildOptions.SetName(parts[1])

		build, _, err := codeEngineClient.GetBuild(getBuildOptions)
		if err != nil {
			return err
		}

		obj = *build
		return nil
	}
}

func testAccCheckIbmCodeEngineBuildDestroy(s *terraform.State) error {
	codeEngineClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CodeEngineV2()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_code_engine_build" {
			continue
		}

		getBuildOptions := &codeenginev2.GetBuildOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getBuildOptions.SetProjectID(parts[0])
		getBuildOptions.SetName(parts[1])

		// Try to find the key
		_, response, err := codeEngineClient.GetBuild(getBuildOptions)

		if err == nil {
			return fmt.Errorf("code_engine_build still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for code_engine_build (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
