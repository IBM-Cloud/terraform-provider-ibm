// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package codeengine_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM/code-engine-go-sdk/codeenginev2"
)

// TestAccIbmCodeEngineBuildRunActionBasic tests successful build run invocation
// This test verifies that:
// - Action can be invoked with required parameters via lifecycle trigger
// - Build run is created successfully
// - No errors are returned
func TestAccIbmCodeEngineBuildRunActionBasic(t *testing.T) {
	projectID := acc.CeProjectId
	buildName := fmt.Sprintf("tf-build-test-%d", acctest.RandIntRange(10, 1000))

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acc.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acc.TestAccProtoV6ProviderFactories(),
		Steps: []resource.TestStep{
			{
				// Create build and trigger action via lifecycle
				Config: testAccCheckIbmCodeEngineBuildRunActionBasic(projectID, buildName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("ibm_code_engine_build.test_build", "build_id"),
					testAccCheckIbmCodeEngineBuildRunActionInvoked(projectID, buildName),
				),
			},
		},
	})
}

// TestAccIbmCodeEngineBuildRunActionTimeout tests timeout handling with subtests
// This test verifies that:
// - Action accepts custom timeout parameter
// - Action respects timeout for successful builds
// - Appropriate error is returned when build run exceeds timeout
func TestAccIbmCodeEngineBuildRunActionTimeout(t *testing.T) {
	projectID := acc.CeProjectId

	t.Run("Success", func(t *testing.T) {
		buildName := fmt.Sprintf("tf-build-timeout-success-%d", acctest.RandIntRange(10, 1000))
		timeout := int64(600) // 10 minutes

		resource.Test(t, resource.TestCase{
			PreCheck:                 func() { acc.TestAccPreCheckCodeEngine(t) },
			ProtoV6ProviderFactories: acc.TestAccProtoV6ProviderFactories(),
			Steps: []resource.TestStep{
				{
					Config: testAccCheckIbmCodeEngineBuildRunActionPrerequisites(projectID, buildName),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttrSet("ibm_code_engine_build.test_build", "build_id"),
					),
				},
				{
					Config: testAccCheckIbmCodeEngineBuildRunActionWithTimeout(projectID, buildName, timeout),
					Check: resource.ComposeTestCheckFunc(
						testAccCheckIbmCodeEngineBuildRunActionInvoked(projectID, buildName),
					),
				},
			},
		})
	})

	t.Run("Failure", func(t *testing.T) {
		buildName := fmt.Sprintf("tf-build-timeout-fail-%d", acctest.RandIntRange(10, 1000))
		shortTimeout := int64(30) // 30 seconds - likely to timeout for real builds

		resource.Test(t, resource.TestCase{
			PreCheck:                 func() { acc.TestAccPreCheckCodeEngine(t) },
			ProtoV6ProviderFactories: acc.TestAccProtoV6ProviderFactories(),
			Steps: []resource.TestStep{
				{
					Config: testAccCheckIbmCodeEngineBuildRunActionPrerequisites(projectID, buildName),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttrSet("ibm_code_engine_build.test_build", "build_id"),
					),
				},
				{
					Config:      testAccCheckIbmCodeEngineBuildRunActionWithTimeout(projectID, buildName, shortTimeout),
					ExpectError: regexp.MustCompile("timeout|timed out"),
				},
			},
		})
	})
}

// TestAccIbmCodeEngineBuildRunActionBuildNotFound tests error handling for non-existent build
// This test verifies that:
// - Action returns appropriate error when build doesn't exist
// - Error message is clear and actionable
func TestAccIbmCodeEngineBuildRunActionBuildNotFound(t *testing.T) {
	projectID := acc.CeProjectId
	nonExistentBuild := "non-existent-build-12345"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acc.TestAccPreCheckCodeEngine(t) },
		ProtoV6ProviderFactories: acc.TestAccProtoV6ProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckIbmCodeEngineBuildRunActionBasic(projectID, nonExistentBuild),
				ExpectError: regexp.MustCompile("Build .* not found"),
			},
		},
	})
}

// Helper function to verify action was invoked by checking for recent build runs
func testAccCheckIbmCodeEngineBuildRunActionInvoked(projectID, buildName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		codeEngineClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CodeEngineV2()
		if err != nil {
			return fmt.Errorf("Error getting Code Engine client: %s", err)
		}

		// List build runs for this build to verify one was created
		listBuildRunsOptions := &codeenginev2.ListBuildRunsOptions{}
		listBuildRunsOptions.SetProjectID(projectID)
		listBuildRunsOptions.SetBuildName(buildName)
		listBuildRunsOptions.SetLimit(10)

		buildRunsList, _, err := codeEngineClient.ListBuildRuns(listBuildRunsOptions)
		if err != nil {
			return fmt.Errorf("Error listing build runs: %s", err)
		}

		if buildRunsList == nil || len(buildRunsList.BuildRuns) == 0 {
			return fmt.Errorf("No build runs found for build %s - action may not have been invoked", buildName)
		}

		// Verify at least one build run exists (action was invoked)
		return nil
	}
}

// Configuration helpers

func testAccCheckIbmCodeEngineBuildRunActionPrerequisites(projectID, buildName string) string {
	return fmt.Sprintf(`
		data "ibm_code_engine_project" "test_project" {
			project_id = "%s"
		}

		resource "ibm_code_engine_build" "test_build" {
			project_id    = data.ibm_code_engine_project.test_project.project_id
			name          = "%s"
			output_image  = "private.us.icr.io/ce-terraform-test/%s"
			output_secret = "ce-terraform-test"
			source_url    = "https://github.com/IBM/CodeEngine"
			strategy_type = "dockerfile"
		}
	`, projectID, buildName, buildName)
}

func testAccCheckIbmCodeEngineBuildRunActionBasic(projectID, buildName string) string {
	return fmt.Sprintf(`
		data "ibm_code_engine_project" "test_project" {
			project_id = "%s"
		}

		action "ibm_code_engine_build_run" "test_action" {
			config {
				project_id = "%s"
				build_name = "%s"
			}
		}

		resource "ibm_code_engine_build" "test_build" {
			project_id    = data.ibm_code_engine_project.test_project.project_id
			name          = "%s"
			output_image  = "private.us.icr.io/ce-terraform-test/%s"
			output_secret = "ce-terraform-test"
			source_url    = "https://github.com/IBM/CodeEngine"
			strategy_type = "dockerfile"

			lifecycle {
				action_trigger {
					events  = [after_create]
					actions = [action.ibm_code_engine_build_run.test_action]
				}
			}
		}
	`, projectID, projectID, buildName, buildName, buildName)
}

func testAccCheckIbmCodeEngineBuildRunActionWithTimeout(projectID, buildName string, timeout int64) string {
	return fmt.Sprintf(`
		data "ibm_code_engine_project" "test_project" {
			project_id = "%s"
		}

		action "ibm_code_engine_build_run" "test_action" {
			config {
				project_id = "%s"
				build_name = "%s"
				timeout    = %d
			}
		}

		resource "ibm_code_engine_build" "test_build" {
			project_id    = data.ibm_code_engine_project.test_project.project_id
			name          = "%s"
			output_image  = "private.us.icr.io/ce-terraform-test/%s"
			output_secret = "ce-terraform-test"
			source_url    = "https://github.com/IBM/CodeEngine"
			strategy_type = "dockerfile"

			lifecycle {
				action_trigger {
					events  = [after_create]
					actions = [action.ibm_code_engine_build_run.test_action]
				}
			}
		}
	`, projectID, projectID, buildName, timeout, buildName, buildName)
}
